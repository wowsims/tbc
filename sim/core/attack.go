package core

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// OnBeforeSwingHit is called right before an auto attack fires
//  if false is returned the weapon swing is cancelled.
//  This allows for abilities that convert a white attack into yellow attack.
type OnBeforeSwing func(sim *Simulation, isOH bool) bool

// OnMeleeAttack is invoked on auto attacks and abilities.
//  Ability can be nil if this was activated by an ability.
type OnMeleeAttack func(sim *Simulation, target *Target, result MeleeHitType, ability *ActiveMeleeAbility, isOH bool)

// OnBeforeMelee is invoked before the hit/dmg rolls are made.
//  This is invoked on both auto attacks and melee abilities.
//
type OnBeforeMelee func(sim *Simulation, ability *ActiveMeleeAbility, isOH bool)

type ResourceCost struct {
	Type  stats.Stat // stats.Mana, stats.Energy, stats.Rage
	Value float64
}

type MeleeAbility struct {
	// ID for the action.
	ActionID ActionID

	// The name of the cast action, e.g. 'Shadowbolt'.
	Name string

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

	// If true will reset swing timers.
	ResetSwingTimer bool

	// NormalizeWeaponSpeed will override the weapon speed for damage normalization
	NormalizeWeaponSpeed float64

	// Internal field only, used to prevent pool objects from being used by
	// multiple attacks simultaneously.
	objectInUse bool
}

// PerformAutoAttack performs a basic weapon swing of the given type.
func PerformAutoAttack(sim *Simulation, c *Character, weapon *items.Item, effect *ActiveMeleeAbility, isOH bool) {
	target := effect.Target
	target.OnBeforeMelee(sim, effect, isOH)
	c.OnBeforeMelee(sim, effect, isOH)

	// Main use of OnBeforeSwing is if the swing needs to turn into a yellow hit (skipping the white hit damage below)
	if c.AutoAttacks.OnBeforeSwing != nil {
		if doSwing := c.AutoAttacks.OnBeforeSwing(sim, isOH); !doSwing {
			return // skip the attack, metrics should be recorded in the replaced attack.
		}
	}

	hit := PerformAttack(sim, c, target, effect.AbilityEffect)

	hitStr := ""
	dmgMult := effect.DamageMultiplier * effect.StaticDamageMultiplier
	if hit == MeleeHitTypeBlock {
		// TODO: How does block reduce damage.
		hitStr = "blocked"
	} else if hit == MeleeHitTypeGlance {
		dmgMult *= 0.75
		hitStr = "glances"
	} else if hit == MeleeHitTypeCrit {
		dmgMult *= effect.CritMultiplier
		hitStr = "crits"
	} else if hit == MeleeHitTypeHit {
		// no change to multiplier
		hitStr = "hits"
	} else {
		if sim.Log != nil {
			// TODO: log actual type of not-hit
			sim.Log("Melee auto attack did not hit.")
		}
		c.OnMeleeAttack(sim, target, hit, nil, isOH)
		c.Metrics.AddAutoAttack(weapon.ID, hit, 0, isOH)
		return // no damage from a block/miss
	}

	dmg := meleeDamage(sim, weapon.WeaponDamageMin, weapon.WeaponDamageMax, 0, weapon.SwingSpeed, isOH, dmgMult, c.stats[stats.AttackPower]+effect.BonusAttackPower, target.ArmorDamageReduction())
	if sim.Log != nil {
		sim.Log("Melee auto attack %s for %0.1f", hitStr, dmg)
	}
	c.OnMeleeAttack(sim, target, hit, nil, isOH)
	c.Metrics.AddAutoAttack(weapon.ID, hit, dmg, isOH)
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

// PerformAttack performs a basic weapon swing of the given type.
func PerformAttack(sim *Simulation, c *Character, target *Target, effect AbilityEffect) MeleeHitType {
	// 1. Single roll -> Miss				Dodge	Parry	Glance	Block	Crit / Hit
	// 3 				8.0%(9.0% hit cap)	6.5%	14.0%	24% 	5%		-4.8%

	// TODO: many calculations in here can be cached. For now its just written out fully.
	//  Once everything is working we can start caching values.

	roll := sim.RandomFloat("auto attack")
	level := float64(target.Level)
	skill := 350.0 // assume max skill level for now.

	levelMinusSkill := (level * 5) - skill
	// First check miss
	missChance := 0.05 + (levelMinusSkill)*0.002
	hitSuppression := (levelMinusSkill - 10) * 0.002

	if !effect.IgnoreDualWieldPenalty && c.Equip[proto.ItemSlot_ItemSlotOffHand].WeaponType != 0 {
		missChance += 0.19
	}

	hitBonus := ((c.stats[stats.MeleeHit] + effect.BonusHitRating) / (MeleeHitRatingPerHitChance * 100)) - hitSuppression
	if hitBonus > 0 {
		missChance -= math.Min(missChance, hitBonus)
	}

	chance := missChance
	if roll < chance {
		return MeleeHitTypeMiss
	}

	dodge := 0.05 + levelMinusSkill*0.001
	expertisePercentage := math.Min(math.Floor(c.stats[stats.Expertise]/(ExpertisePerQuarterPercentReduction))/400, dodge)
	// Next Dodge
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
	chance += math.Max(0.06+levelMinusSkill*0.012, 0)
	if roll < chance {
		return MeleeHitTypeGlance
	}
	// Crit Check
	chance += ((c.stats[stats.MeleeCrit] + effect.BonusCritRating) / (MeleeCritRatingPerCritChance * 100)) + ((skill - (level * 5)) * 0.002) - 0.018

	if roll < chance {
		return MeleeHitTypeCrit
	}

	return MeleeHitTypeHit
}

type ActiveMeleeAbility struct {
	MeleeAbility

	OnMeleeAttack OnMeleeAttack

	Result      MeleeHitType // Hit roll result
	Hits        int32
	Misses      int32
	Crits       int32
	TotalDamage float64 // Damage done by this cast.

	DirectDamageInput
	WeaponDamageInput
	AbilityEffect
}

// If MainHand or Offhand is non-zero the associated ability will create a weapon swing.
type WeaponDamageInput struct {
	MainHand     float64 // dmg multiplier on MH weapon damage.
	MainHandFlat float64 // flat bonus added to MH swing
	Offhand      float64 // dmg multiplier on OH weapon damage.
	OffhandFlat  float64 // Flat bonus added to OH swing
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
// TODO: add AbilityResult data to action metrics.
func (ability *ActiveMeleeAbility) Attack(sim *Simulation) bool {
	if !ability.IgnoreCooldowns && ability.Character.GetRemainingCD(GCDCooldownID, sim.CurrentTime) > 0 {
		log.Fatalf("Ability used while on GCD\n-------\nAbility %s: %#v\n-------\nCharacter: %#v", ability.Name, ability, ability.Character)
	}
	if ability.MeleeAbility.Cost.Type != 0 {
		if ability.Character.stats[ability.MeleeAbility.Cost.Type] < ability.MeleeAbility.Cost.Value {
			return false
		}
		if ability.MeleeAbility.Cost.Type == stats.Mana {
			ability.Character.SpendMana(sim, ability.MeleeAbility.Cost.Value, ability.MeleeAbility.Name)
		} else {
			ability.Character.AddStat(ability.MeleeAbility.Cost.Type, -ability.MeleeAbility.Cost.Value)
		}
	}

	ability.performAttack(sim)
	ability.Character.Metrics.AddMeleeAbility(ability)
	if !ability.IgnoreCooldowns {
		gcdCD := MaxDuration(ability.CalculatedGCD(ability.Character), ability.CastTime)
		ability.Character.SetCD(GCDCooldownID, sim.CurrentTime+gcdCD)
	}
	return true
}

func (ability *ActiveMeleeAbility) performAttack(sim *Simulation) {
	ability.Target.OnBeforeMelee(sim, ability, false)
	ability.Character.OnBeforeMelee(sim, ability, false)

	// Goes on CD on use
	if ability.ActionID.CooldownID != 0 {
		ability.Character.SetCD(ability.ActionID.CooldownID, sim.CurrentTime+ability.Cooldown)
	}

	// 1. Attack Roll
	hit := PerformAttack(sim, ability.Character, ability.Target, ability.AbilityEffect)
	ability.Result = hit
	if hit == MeleeHitTypeMiss || hit == MeleeHitTypeDodge || hit == MeleeHitTypeParry {
		if sim.Log != nil {
			sim.Log("%s did not hit.", ability.Name)
		}
		if ability.WeaponDamageInput.MainHand > 0 {
			ability.Misses++
		}
		if ability.WeaponDamageInput.Offhand > 0 {
			ability.Misses++
		}
		if ability.DirectDamageInput.FlatDamageBonus > 0 || ability.DirectDamageInput.MinBaseDamage > 0 {
			ability.Misses++
		}
		// Not sure MH/OH Matters for an attack
		ability.Character.OnMeleeAttack(sim, ability.Target, hit, ability, false)
		return // we know we missed.
	}

	c := ability.Character
	skill := 350.0
	level := float64(ability.Target.Level)
	critChance := ((c.stats[stats.MeleeCrit] + ability.BonusCritRating) / (MeleeCritRatingPerCritChance * 100)) + ((skill - (level * 5)) * 0.002) - 0.018

	if ability.DirectDamageInput.FlatDamageBonus > 0 || ability.DirectDamageInput.MinBaseDamage > 0 {
		ability.applyFlatDamage(sim, critChance)
	}

	if ability.WeaponDamageInput.MainHand > 0 {
		ability.applySwingDamage(sim, proto.ItemSlot_ItemSlotMainHand, ability.WeaponDamageInput.MainHand, critChance)
	}

	if weapon := c.Equip[proto.ItemSlot_ItemSlotOffHand]; ability.WeaponDamageInput.Offhand > 0 && weapon.ID > 0 { // only attack if we have it
		ability.applySwingDamage(sim, proto.ItemSlot_ItemSlotOffHand, ability.WeaponDamageInput.Offhand, critChance)
	}

	return
}

func (ability *ActiveMeleeAbility) applySwingDamage(sim *Simulation, slot proto.ItemSlot, dmgMult, critChance float64) {
	roll := sim.RandomFloat("weapon swing")
	hit := MeleeHitTypeHit
	if roll < critChance {
		hit = MeleeHitTypeCrit
		dmgMult *= ability.CritMultiplier
		ability.Crits++
	}
	ability.Hits++
	char := ability.Character // just to shorten usage.

	weapon := ability.Character.Equip[slot]
	speed := ability.NormalizeWeaponSpeed
	if speed == 0 {
		speed = weapon.SwingSpeed
	}

	isOH := slot != proto.ItemSlot_ItemSlotMainHand

	multiplier := ability.DamageMultiplier * ability.StaticDamageMultiplier
	flat := 0.0
	if !isOH {
		multiplier *= ability.WeaponDamageInput.MainHand
		flat += ability.MainHandFlat
	} else {
		multiplier *= ability.WeaponDamageInput.Offhand
		flat += ability.OffhandFlat
	}

	dmg := meleeDamage(sim, weapon.WeaponDamageMin, weapon.WeaponDamageMax, flat, speed, isOH, multiplier, char.stats[stats.AttackPower]+ability.BonusAttackPower, ability.Target.ArmorDamageReduction())
	if sim.Log != nil {
		sim.Log("%s mainhand for %0.1f", ability.Name, dmg)
	}
	if ability.OnMeleeAttack != nil {
		ability.OnMeleeAttack(sim, ability.Target, hit, ability, false)
	}
	char.OnMeleeAttack(sim, ability.Target, ability.Result, ability, false)
	ability.TotalDamage += dmg
}

func (ability *ActiveMeleeAbility) applyFlatDamage(sim *Simulation, critChance float64) {
	roll := sim.RandomFloat("weapon swing")
	dmgMult := 1.0
	if roll < critChance {
		dmgMult = ability.CritMultiplier
		ability.Crits++
	}
	ability.Hits++

	// Do a 'direct damage' if ability has it
	dmg := ability.DirectDamageInput.MinBaseDamage + (ability.DirectDamageInput.MaxBaseDamage-ability.DirectDamageInput.MinBaseDamage)*sim.RandomFloat("melee direct damage") + ability.DirectDamageInput.FlatDamageBonus
	ability.TotalDamage += dmg * dmgMult
	if sim.Log != nil {
		sim.Log("%s for %0.1f", ability.Name, dmg)
	}
}

func meleeDamage(sim *Simulation, weaponMin, weaponMax, flatBonus, speed float64, offhand bool, multiplier float64, attackPower float64, damageReduction float64) float64 {
	if offhand {
		multiplier *= 0.5
	}
	dmg := weaponMin + (weaponMax-weaponMin)*sim.RandomFloat("melee")
	dmg += (speed * attackPower) / MeleeAttackRatingPerDamage
	dmg *= multiplier
	dmg += flatBonus
	dmg *= 1 - damageReduction
	return dmg
}

type AbilityEffect struct {
	// Target of the spell.
	Target *Target

	// Bonus stats to be added to the attack.
	BonusHitRating   float64
	BonusAttackPower float64
	BonusCritRating  float64

	IgnoreDualWieldPenalty bool

	// Additional multiplier that is always applied.
	DamageMultiplier float64

	// applies fixed % increases to damage at cast time.
	//  Only use multipliers that don't change for the lifetime of the sim.
	//  This should probably only be mutated in a template and not changed in auras.
	StaticDamageMultiplier float64
}

func NewAutoAttacks(c *Character) AutoAttacks {
	st := AutoAttacks{
		c: c,
		AbilityEffect: AbilityEffect{
			DamageMultiplier:       1.0,
			StaticDamageMultiplier: 1.0,
		},
		MeleeAbility: MeleeAbility{
			Name:           "Auto Attacks",
			CritMultiplier: 2.0,
		},
	}

	// Setup initial swing timers
	if weapon := c.Equip[proto.ItemSlot_ItemSlotMainHand]; weapon.ID != 0 {
		st.mhbase = time.Duration(weapon.SwingSpeed * float64(time.Second))
		st.MainhandSwingAt = time.Duration(float64(st.mhbase) / c.SwingSpeed())
		st.mh = &weapon
	}
	if weapon := c.Equip[proto.ItemSlot_ItemSlotOffHand]; weapon.ID != 0 {
		st.ohbase = time.Duration(weapon.SwingSpeed * float64(time.Second))
		st.OffhandSwingAt = time.Duration(float64(st.ohbase) / c.SwingSpeed())
		st.oh = &weapon
	}

	return st
}

type AutoAttacks struct {
	// initialized
	c      *Character
	mh     *items.Item
	oh     *items.Item
	mhbase time.Duration
	ohbase time.Duration

	MainhandSwingAt time.Duration
	OffhandSwingAt  time.Duration
	AbilityEffect   // bonuses to auto attacks
	MeleeAbility
	active ActiveMeleeAbility // Mostly just for passing AbilityEffect to OnBeforeMelee to allow modification to auto attacks.

	OnBeforeSwing OnBeforeSwing
}

func (aa *AutoAttacks) MainhandSwingSpeed() time.Duration {
	return time.Duration(float64(aa.mhbase) / aa.c.SwingSpeed())
}

func (aa *AutoAttacks) OffhandSwingSpeed() time.Duration {
	return time.Duration(float64(aa.ohbase) / aa.c.SwingSpeed())
}

// Swing will check any swing timers if they are up, and if so, swing!
func (aa *AutoAttacks) Swing(sim *Simulation, target *Target) {
	aa.AbilityEffect.Target = target
	if aa.MainhandSwingAt <= sim.CurrentTime {
		// Make a MH swing!
		aa.active.AbilityEffect = aa.AbilityEffect
		aa.active.MeleeAbility = aa.MeleeAbility
		PerformAutoAttack(sim, aa.c, aa.mh, &aa.active, false)
		aa.MainhandSwingAt = sim.CurrentTime + aa.MainhandSwingSpeed()
	}
	if aa.OffhandSwingAt <= sim.CurrentTime {
		// Make a OH swing!
		aa.active.AbilityEffect = aa.AbilityEffect
		aa.active.MeleeAbility = aa.MeleeAbility
		PerformAutoAttack(sim, aa.c, aa.oh, &aa.active, true)
		aa.OffhandSwingAt = sim.CurrentTime + aa.OffhandSwingSpeed()
	}
}

func (aa *AutoAttacks) ModifySwingTime(sim *Simulation, amount float64) {
	if aa.mh == nil {
		return
	}
	mhSwingTime := aa.MainhandSwingAt - sim.CurrentTime
	if mhSwingTime > 0 {
		aa.MainhandSwingAt = sim.CurrentTime + time.Duration(float64(mhSwingTime)/amount)
	}

	if aa.oh == nil {
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

// TimeUntil compares swing timers to the next cast or attack and returns the time the next event occurs at.
//   This could probably be broken into TimeUntil(cast), TimeUntil(attack), TimeUntil(event)
func (aa *AutoAttacks) TimeUntil(sim *Simulation, cast *SimpleSpell, atk *ActiveMeleeAbility, event time.Duration) time.Duration {
	var nextEventTime time.Duration
	if event > 0 {
		nextEventTime = event
	}
	if cast != nil {
		if cast.CastTime > 0 {
			// Resume swings after cast is completed
			aa.MainhandSwingAt = sim.CurrentTime + cast.CastTime + aa.MainhandSwingSpeed()
			aa.OffhandSwingAt = sim.CurrentTime + cast.CastTime + aa.OffhandSwingSpeed()
		}
		nextEventTime = MaxDuration(cast.CastTime, cast.Character.GetRemainingCD(GCDCooldownID, sim.CurrentTime))
	}
	if atk != nil {
		if atk.ResetSwingTimer {
			aa.MainhandSwingAt = sim.CurrentTime + aa.MainhandSwingSpeed()
			aa.OffhandSwingAt = sim.CurrentTime + aa.OffhandSwingSpeed()
		}
		nextEventTime = MaxDuration(atk.CastTime, atk.Character.GetRemainingCD(GCDCooldownID, sim.CurrentTime))
	}
	mhswing := aa.MainhandSwingAt - sim.CurrentTime
	if mhswing < nextEventTime || nextEventTime == 0 {
		nextEventTime = mhswing
	}
	if aa.ohbase > 0 {
		ohswing := aa.OffhandSwingAt - sim.CurrentTime
		if ohswing < nextEventTime {
			nextEventTime = ohswing
		}
	}

	return nextEventTime
}

type MeleeAbilittyTemplate struct {
	template ActiveMeleeAbility
}

func (template *MeleeAbilittyTemplate) Apply(newAction *ActiveMeleeAbility) {
	if newAction.objectInUse {
		panic(fmt.Sprintf("Damage over time spell (%s) already in use", newAction.Name))
	}
	*newAction = template.template
}

// Takes in a cast template and returns a template, so you don't need to keep track of which things to allocate yourself.
func NewMeleeAbilittyTemplate(spellTemplate ActiveMeleeAbility) MeleeAbilittyTemplate {
	return MeleeAbilittyTemplate{
		template: spellTemplate,
	}
}

// PPMToChance converts a character proc-per-minute into mh/oh proc chances
func PPMToChance(char *Character, ppm float64) (float64, float64) {
	procChance := (char.Equip[proto.ItemSlot_ItemSlotMainHand].SwingSpeed * ppm) / 60.0
	ohProcChance := (char.Equip[proto.ItemSlot_ItemSlotOffHand].SwingSpeed * ppm) / 60.0
	return procChance, ohProcChance
}
