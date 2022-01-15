package core

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// OnBeforeMHSwing is called right before an auto attack fires
//  if false is returned the weapon swing is cancelled.
//  This allows for abilities that convert a white attack into yellow attack.
type OnBeforeMHSwing func(sim *Simulation) bool

// OnBeforeMelee is invoked once for each ability, even if there are multiple hits.
//  This should be used for any effects that adjust the stats / multipliers of the attack.
type OnBeforeMelee func(sim *Simulation, ability *ActiveMeleeAbility)

// OnBeforeMelee is invoked before the hit/dmg rolls are made.
//  This is invoked on both auto attacks and melee abilities.
//  This should be used for any effects that adjust the stats / multipliers of the attack.
type OnBeforeMeleeHit func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect)

// OnMeleeAttack is invoked on auto attacks and abilities.
//  This should be used for any on-hit procs.
type OnMeleeAttack func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect)

type ResourceCost struct {
	Type  stats.Stat // stats.Mana, stats.Energy, stats.Rage
	Value float64
}

type MeleeAbility struct {
	// ID for the action.
	ActionID

	// The character performing this action.
	Character *Character

	// If set, this action will start a cooldown using its cooldown ID.
	// Note that the GCD CD will be activated even if this is not set.
	Cooldown time.Duration

	// If set, this will be used as the GCD instead of the default value (1.5s).
	GCDCooldown time.Duration

	// If set, CD for this action and GCD CD will be ignored, and this action
	// will not set new values for those CDs either.
	IgnoreCooldowns bool

	// If set, this spell will have its resource cost ignored.
	IgnoreCost bool

	Cost ResourceCost

	CastTime time.Duration // most melee skills are instant... are there any with a cast time?

	// E.g. for nature spells, set to stats.NatureSpellPower.
	SpellSchool stats.Stat

	// How much to multiply damage by, if this cast crits.
	CritMultiplier float64

	// If true, will force the cast to crit (if it doesnt miss).
	GuaranteedCrit bool

	// Internal field only, used to prevent pool objects from being used by
	// multiple attacks simultaneously.
	objectInUse bool
}

type MeleeHitType byte

const (
	MeleeHitTypeMiss MeleeHitType = iota
	MeleeHitTypeDodge
	MeleeHitTypeParry
	MeleeHitTypeGlance
	MeleeHitTypeBlock
	MeleeHitTypeCrit
	MeleeHitTypeHit
)

type AbilityEffect struct {
	// Target of the spell.
	Target *Target

	// Bonus stats to be added to the attack.
	BonusWeaponDamage    float64
	BonusHitRating       float64
	BonusAttackPower     float64
	BonusCritRating      float64
	BonusExpertiseRating float64

	IsWhiteHit bool

	// Causes the first roll for this hit to be copied from ActiveMeleeAbility.MainHit.HitType.
	// This is only used by Shaman Stormstrike.
	ReuseMainHitRoll bool

	// Additional multiplier that is always applied.
	DamageMultiplier float64

	// applies fixed % increases to damage at cast time.
	//  Only use multipliers that don't change for the lifetime of the sim.
	//  This should probably only be mutated in a template and not changed in auras.
	StaticDamageMultiplier float64

	// The type of hit this was, i.e. miss/dodge/block/crit/hit.
	HitType MeleeHitType

	// The damage done by this effect.
	Damage float64
}

// Represents a generic weapon. Pets / unarmed / various other cases dont use
// actual weapon items so this is an abstraction of a Weapon.
type Weapon struct {
	BaseDamageMin float64
	BaseDamageMax float64
	SwingSpeed    float64
	SwingDuration time.Duration // Duration between 2 swings.
}

func (weapon Weapon) calculateSwingDamage(sim *Simulation, bonusWeaponDamage float64, attackPower float64) float64 {
	dmg := weapon.BaseDamageMin + bonusWeaponDamage + (weapon.BaseDamageMax-weapon.BaseDamageMin)*sim.RandomFloat("melee")
	dmg += (weapon.SwingSpeed * attackPower) / MeleeAttackRatingPerDamage
	return dmg
}

// If MainHand or Offhand is non-zero the associated ability will create a weapon swing.
type WeaponDamageInput struct {
	// Whether this input corresponds to the OH weapon.
	// It's important that this be 'IsOH' instead of 'IsMH' so that MH is the default.
	IsOH bool

	DamageMultiplier float64 // Damage multiplier on weapon damage.
	FlatDamageBonus  float64 // Flat bonus added to swing.
}

type AbilityHitEffect struct {
	AbilityEffect
	DirectInput DirectDamageInput
	WeaponInput WeaponDamageInput
}

type ActiveMeleeAbility struct {
	MeleeAbility

	OnMeleeAttack OnMeleeAttack

	HitType     MeleeHitType // Hit roll result
	Hits        int32
	Misses      int32
	Crits       int32
	Dodges      int32
	Glances     int32
	Parries     int32
	Blocks      int32
	TotalDamage float64 // Damage done by this cast.

	// All abilities have at least 1 hit, so this should always be filled.
	MainHit AbilityHitEffect

	// For abilities that have more than 1 hit.
	AdditionalHits []AbilityHitEffect
}

func (effect *AbilityEffect) Landed() bool {
	return effect.HitType != MeleeHitTypeMiss && effect.HitType != MeleeHitTypeDodge && effect.HitType != MeleeHitTypeParry
}

// Computes an attack result using the white-hit table formula (single roll).
func (effect *AbilityEffect) WhiteHitTableResult(sim *Simulation, ability *ActiveMeleeAbility) MeleeHitType {
	// 1. Single roll -> Miss				Dodge	Parry	Glance	Block	Crit / Hit
	// 3 				8.0%(9.0% hit cap)	6.5%	14.0%	24% 	5%		-4.8%

	// TODO: many calculations in here can be cached. For now its just written out fully.
	//  Once everything is working we can start caching values.
	character := ability.Character

	roll := sim.RandomFloat("auto attack")
	level := float64(effect.Target.Level)
	skill := 350.0 // assume max skill level for now.

	// Difference between attacker's waepon skill and target's defense skill.
	skillDifference := (level * 5) - skill

	// First check miss
	missChance := 0.05 + skillDifference*0.002
	if effect.IsWhiteHit && character.AutoAttacks.IsDualWielding {
		missChance += 0.19
	}
	hitSuppression := (skillDifference - 10) * 0.002
	hitBonus := ((character.stats[stats.MeleeHit] + effect.BonusHitRating) / (MeleeHitRatingPerHitChance * 100)) - hitSuppression
	if hitBonus > 0 {
		missChance = math.Max(0, missChance-hitBonus)
	}

	chance := missChance
	if roll < chance {
		return MeleeHitTypeMiss
	}

	// Next Dodge
	dodge := 0.05 + skillDifference*0.001
	expertisePercentage := math.Min(math.Floor((character.stats[stats.Expertise]+effect.BonusExpertiseRating)/(ExpertisePerQuarterPercentReduction))/400, dodge)
	chance += dodge - expertisePercentage
	if roll < chance {
		return MeleeHitTypeDodge
	}

	// Parry (if in front)
	// If the target is a mob and defense minus weapon skill is 11 or more:
	// ParryChance = 5% + (TargetLevel*5 - AttackerSkill) * 0.6%

	// If the target is a mob and defense minus weapon skill is 10 or less:
	// ParryChance = 5% + (TargetLevel*5 - AttackerSkill) * 0.1%

	// Block (if in front)
	// If the target is a mob:
	// BlockChance = MIN(5%, 5% + (TargetLevel*5 - AttackerSkill) * 0.1%)

	// Glancing Check
	chance += math.Max(0.06+skillDifference*0.012, 0)
	if roll < chance {
		return MeleeHitTypeGlance
	}
	// Crit Check
	chance += ((character.stats[stats.MeleeCrit] + effect.BonusCritRating) / (MeleeCritRatingPerCritChance * 100)) - skillDifference*0.002 - 0.018

	if roll < chance {
		return MeleeHitTypeCrit
	}

	return MeleeHitTypeHit
}

func (effect *AbilityEffect) String() string {
	if effect.HitType == MeleeHitTypeMiss {
		return "Miss"
	} else if effect.HitType == MeleeHitTypeDodge {
		return "Dodge"
	} else if effect.HitType == MeleeHitTypeParry {
		return "Parry"
	}

	var sb strings.Builder

	if effect.HitType == MeleeHitTypeHit {
		sb.WriteString("Hit")
	} else if effect.HitType == MeleeHitTypeCrit {
		sb.WriteString("Crit")
	} else if effect.HitType == MeleeHitTypeGlance {
		sb.WriteString("Glance")
	} else { // Block
		sb.WriteString("Block")
	}

	fmt.Fprintf(&sb, " for %0.3f damage", effect.Damage)

	return sb.String()
}

func (ahe *AbilityHitEffect) calculateDamage(sim *Simulation, ability *ActiveMeleeAbility) {
	character := ability.Character

	attackPower := character.stats[stats.AttackPower] + ahe.BonusAttackPower
	bonusWeaponDamage := ahe.BonusWeaponDamage

	if ahe.AbilityEffect.ReuseMainHitRoll {
		ahe.HitType = ability.MainHit.HitType
	} else {
		ahe.HitType = ahe.AbilityEffect.WhiteHitTableResult(sim, ability)
	}

	if !ahe.Landed() {
		ahe.Damage = 0
		return
	}

	dmg := 0.0
	if ahe.WeaponInput.DamageMultiplier != 0 {
		if !ahe.WeaponInput.IsOH {
			dmg += character.AutoAttacks.mh.calculateSwingDamage(sim, bonusWeaponDamage, attackPower)
		} else {
			dmg += character.AutoAttacks.oh.calculateSwingDamage(sim, bonusWeaponDamage, attackPower) * 0.5
		}
		dmg *= ahe.WeaponInput.DamageMultiplier
		dmg += ahe.WeaponInput.FlatDamageBonus
	}

	// TODO: Add damage from DirectDamageInput

	// If this is a yellow attack, need a 2nd roll to decide crit. Otherwise just use existing hit result.
	if !ahe.AbilityEffect.IsWhiteHit {
		skill := 350.0
		level := float64(ahe.Target.Level)
		critChance := ((character.stats[stats.MeleeCrit] + ahe.BonusCritRating) / (MeleeCritRatingPerCritChance * 100)) + ((skill - (level * 5)) * 0.002) - 0.018

		roll := sim.RandomFloat("weapon swing")
		if roll < critChance {
			ahe.HitType = MeleeHitTypeCrit
		} else {
			ahe.HitType = MeleeHitTypeHit
		}
	}

	if ahe.HitType == MeleeHitTypeCrit {
		dmg *= ability.CritMultiplier
	} else if ahe.HitType == MeleeHitTypeGlance {
		dmg *= 0.75
	}

	// Apply armor reduction.
	dmg *= 1 - ahe.Target.ArmorDamageReduction(character.stats[stats.ArmorPenetration])
	if sim.Log != nil {
		character.Log(sim, "Target armor: %0.2f\n", ahe.Target.currentArmor)
	}

	// Apply all other effect multipliers.
	dmg *= ahe.DamageMultiplier * ahe.StaticDamageMultiplier

	ahe.Damage = dmg
}

// Returns whether this hit effect is associated with one of the character's
// weapons. This check is necessary to decide whether certains effects are eligible.
func (ahe *AbilityHitEffect) IsWeaponHit() bool {
	return ahe.WeaponInput.DamageMultiplier > 0
}

// Returns whether this hit effect is associated with the main-hand weapon.
func (ahe *AbilityHitEffect) IsMH() bool {
	return !ahe.WeaponInput.IsOH
}

// Returns whether this hit effect is associated with the off-hand weapon.
func (ahe *AbilityHitEffect) IsOH() bool {
	return ahe.WeaponInput.IsOH
}

func (ability *ActiveMeleeAbility) CalculatedGCD(char *Character) time.Duration {
	baseGCD := GCDDefault
	if ability.GCDCooldown != 0 {
		baseGCD = ability.GCDCooldown
	}
	return MaxDuration(GCDMin, time.Duration(float64(baseGCD)/char.SwingSpeed()))
}

// Attack will perform the attack
//  Returns false if unable to attack (due to resource lacking)
func (ability *ActiveMeleeAbility) Attack(sim *Simulation) bool {
	if !ability.IgnoreCooldowns && ability.Character.GetRemainingCD(GCDCooldownID, sim.CurrentTime) > 0 {
		log.Fatalf("Ability used while on GCD\n-------\nAbility %s: %#v\n", ability.ActionID, ability)
	}
	if ability.MeleeAbility.Cost.Type != 0 {
		if ability.Character.stats[ability.MeleeAbility.Cost.Type] < ability.MeleeAbility.Cost.Value {
			return false
		}
		if ability.MeleeAbility.Cost.Type == stats.Mana {
			ability.Character.SpendMana(sim, ability.MeleeAbility.Cost.Value, ability.MeleeAbility.ActionID)
		} else {
			ability.Character.AddStat(ability.MeleeAbility.Cost.Type, -ability.MeleeAbility.Cost.Value)
		}
	}

	ability.Character.OnBeforeMelee(sim, ability)

	ability.MainHit.performAttack(sim, ability)

	if len(ability.AdditionalHits) > 0 {
		for i, _ := range ability.AdditionalHits {
			ahe := &ability.AdditionalHits[i]
			ahe.performAttack(sim, ability)
		}
	}

	if !ability.IgnoreCooldowns {
		gcdCD := MaxDuration(ability.CalculatedGCD(ability.Character), ability.CastTime)
		ability.Character.SetCD(GCDCooldownID, sim.CurrentTime+gcdCD)

		if ability.ActionID.CooldownID != 0 {
			ability.Character.SetCD(ability.ActionID.CooldownID, sim.CurrentTime+ability.Cooldown)
		}
	}
	ability.Character.Metrics.AddMeleeAbility(ability)
	return true
}

func (ahe *AbilityHitEffect) performAttack(sim *Simulation, ability *ActiveMeleeAbility) {
	ability.Character.OnBeforeMeleeHit(sim, ability, ahe)
	ahe.Target.OnBeforeMeleeHit(sim, ability, ahe)

	ahe.calculateDamage(sim, ability)

	if ahe.HitType == MeleeHitTypeMiss {
		ability.Misses++
	} else if ahe.HitType == MeleeHitTypeDodge {
		ability.Dodges++
	} else if ahe.HitType == MeleeHitTypeGlance {
		ability.Glances++
	} else if ahe.HitType == MeleeHitTypeCrit {
		ability.Crits++
	} else if ahe.HitType == MeleeHitTypeHit {
		ability.Hits++
	} else if ahe.HitType == MeleeHitTypeParry {
		ability.Parries++
	} else if ahe.HitType == MeleeHitTypeBlock {
		ability.Blocks++
	}
	ability.TotalDamage += ahe.Damage

	if sim.Log != nil {
		ability.Character.Log(sim, "%s %s", ability.ActionID, ahe)
	}

	ability.Character.OnMeleeAttack(sim, ability, ahe)
	ahe.Target.OnMeleeAttack(sim, ability, ahe)
	if ability.OnMeleeAttack != nil {
		ability.OnMeleeAttack(sim, ability, ahe)
	}
}

type AutoAttacks struct {
	// initialized
	character *Character
	mh        Weapon
	oh        Weapon

	IsDualWielding bool

	// Set this to true to use the OH delay macro, mostly used by enhance shamans.
	// This will intentionally delay OH swings to that they always fall within the
	// 0.5s window following a MH swing.
	DelayOHSwings bool

	MainhandSwingAt time.Duration
	OffhandSwingAt  time.Duration

	ActiveMeleeAbility // Parameters for auto attacks.

	OnBeforeMHSwing OnBeforeMHSwing

	// The time at which the last MH swing occurred.
	previousMHSwingAt time.Duration
}

func NewAutoAttacks(c *Character, delayOHSwings bool) AutoAttacks {
	aa := AutoAttacks{
		character:     c,
		DelayOHSwings: delayOHSwings,
		ActiveMeleeAbility: ActiveMeleeAbility{
			MeleeAbility: MeleeAbility{
				ActionID:        ActionID{OtherID: proto.OtherAction_OtherActionAttack},
				CritMultiplier:  2.0,
				Character:       c,
				IgnoreCooldowns: true,
				IgnoreCost:      true,
			},
			MainHit: AbilityHitEffect{
				AbilityEffect: AbilityEffect{
					IsWhiteHit:             true,
					DamageMultiplier:       1.0,
					StaticDamageMultiplier: 1.0,
				},
				WeaponInput: WeaponDamageInput{
					DamageMultiplier: 1,
				},
			},
		},
	}

	if weapon := c.Equip[proto.ItemSlot_ItemSlotMainHand]; weapon.ID != 0 {
		aa.mh = Weapon{
			BaseDamageMin: weapon.WeaponDamageMin,
			BaseDamageMax: weapon.WeaponDamageMax,
			SwingSpeed:    weapon.SwingSpeed,
			SwingDuration: time.Duration(weapon.SwingSpeed * float64(time.Second)),
		}
	}
	if weapon := c.Equip[proto.ItemSlot_ItemSlotOffHand]; weapon.ID != 0 {
		aa.oh = Weapon{
			BaseDamageMin: weapon.WeaponDamageMin,
			BaseDamageMax: weapon.WeaponDamageMax,
			SwingSpeed:    weapon.SwingSpeed,
			SwingDuration: time.Duration(weapon.SwingSpeed * float64(time.Second)),
		}
	}
	aa.IsDualWielding = aa.mh.SwingSpeed != 0 && aa.oh.SwingSpeed != 0

	return aa
}

func (aa *AutoAttacks) IsEnabled() bool {
	return aa.mh.SwingSpeed != 0
}

func (aa *AutoAttacks) reset(sim *Simulation) {
	if !aa.IsEnabled() {
		return
	}

	aa.MainhandSwingAt = 0
	aa.OffhandSwingAt = 0

	// Set a fake value for previousMHSwing so that offhand swing delay works
	// properly at the start.
	aa.previousMHSwingAt = time.Second * -1

	// Apply random delay of 0 - 0.5s, to one of the weapons.
	delay := time.Duration(sim.RandomFloat("SwingResetDelay") * float64(time.Millisecond*500))
	isMHDelay := sim.RandomFloat("SwingResetWeapon") < 0.5

	if isMHDelay {
		aa.MainhandSwingAt = delay
	} else {
		aa.OffhandSwingAt = delay
	}
}

func (aa *AutoAttacks) MainhandSwingSpeed() time.Duration {
	return time.Duration(float64(aa.mh.SwingDuration) / aa.character.SwingSpeed())
}

func (aa *AutoAttacks) OffhandSwingSpeed() time.Duration {
	return time.Duration(float64(aa.oh.SwingDuration) / aa.character.SwingSpeed())
}

// Swing will check any swing timers if they are up, and if so, swing!
func (aa *AutoAttacks) Swing(sim *Simulation, target *Target) {
	aa.MainHit.Target = target
	if aa.MainhandSwingAt <= sim.CurrentTime {
		doSwing := true
		if aa.OnBeforeMHSwing != nil {
			doSwing = aa.OnBeforeMHSwing(sim)
		}

		if doSwing {
			// Make a MH swing!
			ama := aa.ActiveMeleeAbility
			ama.ActionID.Tag = 1
			ama.MainHit.WeaponInput.IsOH = false
			ama.Attack(sim)
			aa.MainhandSwingAt = sim.CurrentTime + aa.MainhandSwingSpeed()
			aa.previousMHSwingAt = sim.CurrentTime
		}
	}
	if aa.OffhandSwingAt <= sim.CurrentTime {
		if aa.DelayOHSwings && (sim.CurrentTime-aa.previousMHSwingAt) > time.Millisecond*500 {
			// Delay the OH swing for later, so it follows the MH swing.
			aa.OffhandSwingAt = aa.MainhandSwingAt + time.Millisecond*250
		} else {
			// Make a OH swing!
			ama := aa.ActiveMeleeAbility
			ama.ActionID.Tag = 2
			ama.MainHit.WeaponInput.IsOH = true
			ama.Attack(sim)
			aa.OffhandSwingAt = sim.CurrentTime + aa.OffhandSwingSpeed()
		}
	}
}

func (aa *AutoAttacks) ModifySwingTime(sim *Simulation, amount float64) {
	if aa.mh.SwingSpeed == 0 {
		return
	}
	mhSwingTime := aa.MainhandSwingAt - sim.CurrentTime
	if mhSwingTime > 0 {
		aa.MainhandSwingAt = sim.CurrentTime + time.Duration(float64(mhSwingTime)/amount)
	}

	if aa.oh.SwingSpeed == 0 {
		return
	}
	ohSwingTime := aa.OffhandSwingAt - sim.CurrentTime
	if ohSwingTime > 0 {
		newTime := time.Duration(float64(ohSwingTime) / amount)
		if newTime > 0 {
			aa.OffhandSwingAt = sim.CurrentTime + newTime
		}
	}
}

// Returns the time at which the next attack will occur.
func (aa *AutoAttacks) NextAttackAt() time.Duration {
	nextAttack := aa.MainhandSwingAt
	if aa.oh.SwingSpeed != 0 {
		nextAttack = MinDuration(nextAttack, aa.OffhandSwingAt)
	}
	return nextAttack
}

// Returns the time at which the next event will occur, considering both autos and the gcd.
func (aa *AutoAttacks) NextEventAt(sim *Simulation) time.Duration {
	return MinDuration(
		sim.CurrentTime+aa.Character.GetRemainingCD(GCDCooldownID, sim.CurrentTime),
		aa.NextAttackAt())
}

type PPMManager struct {
	mhProcChance float64
	ohProcChance float64
}

// For manually overriding proc chance.
func (ppmm *PPMManager) SetProcChance(isMH bool, newChance float64) {
	if isMH {
		ppmm.mhProcChance = newChance
	} else {
		ppmm.ohProcChance = newChance
	}
}

// Returns whether the effect procced.
func (ppmm *PPMManager) Proc(sim *Simulation, isMH bool, label string) bool {
	if isMH {
		return ppmm.ProcMH(sim, label)
	} else {
		return ppmm.ProcOH(sim, label)
	}
}

// Returns whether the effect procced, assuming MH.
func (ppmm *PPMManager) ProcMH(sim *Simulation, label string) bool {
	return ppmm.mhProcChance > 0 && sim.RandomFloat(label) < ppmm.mhProcChance
}

// Returns whether the effect procced, assuming OH.
func (ppmm *PPMManager) ProcOH(sim *Simulation, label string) bool {
	return ppmm.ohProcChance > 0 && sim.RandomFloat(label) < ppmm.ohProcChance
}

// PPMToChance converts a character proc-per-minute into mh/oh proc chances
func (aa *AutoAttacks) NewPPMManager(ppm float64) PPMManager {
	if aa.mh.SwingSpeed == 0 {
		// Means this character didn't enable autoattacks.
		return PPMManager{
			mhProcChance: 0,
			ohProcChance: 0,
		}
	}

	return PPMManager{
		mhProcChance: (aa.mh.SwingSpeed * ppm) / 60.0,
		ohProcChance: (aa.oh.SwingSpeed * ppm) / 60.0,
	}
}

type MeleeAbilityTemplate struct {
	template       ActiveMeleeAbility
	additionalHits []AbilityHitEffect
}

func (template *MeleeAbilityTemplate) Apply(newAction *ActiveMeleeAbility) {
	if newAction.objectInUse {
		panic(fmt.Sprintf("Melee ability (%s) already in use", newAction.ActionID))
	}
	*newAction = template.template
	newAction.AdditionalHits = template.additionalHits
	copy(newAction.AdditionalHits, template.template.AdditionalHits)
}

// Takes in a cast template and returns a template, so you don't need to keep track of which things to allocate yourself.
func NewMeleeAbilityTemplate(abilityTemplate ActiveMeleeAbility) MeleeAbilityTemplate {
	return MeleeAbilityTemplate{
		template:       abilityTemplate,
		additionalHits: make([]AbilityHitEffect, len(abilityTemplate.AdditionalHits)),
	}
}
