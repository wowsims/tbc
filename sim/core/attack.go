package core

import (
	"fmt"
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// OnBeforeSwingHit is called right before an auto attack lands.
//  if false is returned the weapon swing dmg is not calculated.
//  This allows for abilities that convert a white attack into yellow attack.
type OnBeforeSwingHit func(sim *Simulation, isOH bool) bool

// OnMeleeAttack is invoked on auto attacks and abilities.
//  Ability can be nil if this was activated by an ability.
type OnMeleeAttack func(sim *Simulation, target *Target, result MeleeHitType, ability *ActiveMeleeAbility, isOH bool)

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
func PerformAutoAttack(sim *Simulation, c *Character, target *Target, weapon *items.Item, dmgMult float64, isOH bool) {
	hit := PerformAttack(sim, c, target)

	hitStr := ""
	if hit == MeleeHitTypeGlance {
		dmgMult *= 0.75
		hitStr = "glances"
	} else if hit == MeleeHitTypeCrit {
		dmgMult *= 2.0
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

	// Main use of OnBeforeSwingHit is if the swing needs to turn into a yellow hit (skipping the white hit damage below)
	doSwing := c.OnBeforeSwingHit(sim, isOH)
	if !doSwing {
		return // skip the attack
	}

	dmg := meleeDamage(sim, weapon.WeaponDamageMin, weapon.WeaponDamageMax, 0, weapon.SwingSpeed, isOH, dmgMult, c.stats[stats.AttackPower], target.ArmorDamageReduction())
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
func PerformAttack(sim *Simulation, c *Character, target *Target) MeleeHitType {
	// 1. Single roll -> Miss				Dodge	Parry	Glance	Block	Crit / Hit
	// 3 				8.0%(9.0% hit cap)	6.5%	14.0%	24% 	5%		-4.8%

	roll := sim.RandomFloat("auto attack")

	level := float64(target.Level)
	skill := 350.0 // assume max skill level for now.

	levelMinusSkill := (level * 5) - skill
	// First check miss
	missChance := 0.05 + (levelMinusSkill)*0.002
	hitSuppression := (levelMinusSkill - 10) * 0.002

	if c.Equip[proto.ItemSlot_ItemSlotOffHand].WeaponType != 0 {
		missChance += 0.19
	}

	hitBonus := (c.stats[stats.MeleeHit] / (MeleeHitRatingPerHitChance * 100)) - hitSuppression
	if hitBonus > 0 {
		missChance -= hitBonus
	}

	chance := missChance
	if roll < chance {
		// log.Printf("miss")
		return MeleeHitTypeMiss
	}

	expertise := c.stats[stats.Expertise] / (ExpertisePerPercentReduction * 100)
	// Next Dodge
	chance += 0.05 + levelMinusSkill*0.001 - expertise

	if roll < chance {
		// log.Printf("dodge")
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
	critReduction := (level - 70*0.01) + 0.018
	chance += c.stats[stats.MeleeCrit]/(MeleeCritRatingPerCritChance*100) + (skill - (level*5)*0.002) - critReduction

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

// Attack will perform the attack
//  Returns false if unable to attack (due to resource lacking)
// TODO: add AbilityResult data to action metrics.
func (ama *ActiveMeleeAbility) Attack(sim *Simulation) bool {
	result := ama.performAttack(sim)
	if result {
		ama.Character.Metrics.AddMeleeAbility(ama)
	}
	return result
}

func (ama *ActiveMeleeAbility) performAttack(sim *Simulation) bool {
	if ama.MeleeAbility.Cost.Type != 0 {
		if ama.Character.stats[ama.MeleeAbility.Cost.Type] < ama.MeleeAbility.Cost.Value {
			return false
		}
		if ama.MeleeAbility.Cost.Type == stats.Mana {
			ama.Character.SpendMana(sim, ama.MeleeAbility.Cost.Value, ama.MeleeAbility.Name)
		} else {
			ama.Character.AddStat(ama.MeleeAbility.Cost.Type, -ama.MeleeAbility.Cost.Value)
		}
	}

	// Goes on CD on use
	if ama.ActionID.CooldownID != 0 {
		ama.Character.SetCD(ama.ActionID.CooldownID, sim.CurrentTime+ama.Cooldown)
	}

	// 1. Attack Roll
	hit := PerformAttack(sim, ama.Character, ama.Target)
	ama.Result = hit
	if hit != MeleeHitTypeCrit && hit != MeleeHitTypeGlance && hit != MeleeHitTypeHit {
		// TODO: add metrics
		if sim.Log != nil {
			sim.Log("%s did not hit.", ama.Name)
		}
		if ama.WeaponDamageInput.MainHand > 0 {
			ama.Misses++
		}
		if ama.WeaponDamageInput.Offhand > 0 {
			ama.Misses++
		}
		if ama.DirectDamageInput.FlatDamageBonus > 0 {
			ama.Misses++
		}
		// Not sure MH/OH Matters for an attack
		ama.Character.OnMeleeAttack(sim, ama.Target, hit, ama, false)
	}

	c := ama.Character
	skill := 350.0
	level := float64(ama.Target.Level)
	critReduction := (level - 70*0.01) + 0.018
	critChance := c.stats[stats.MeleeCrit]/(MeleeCritRatingPerCritChance*100) + (skill - (level*5)*0.002) - critReduction

	if ama.DirectDamageInput.FlatDamageBonus > 0 || ama.DirectDamageInput.MinBaseDamage > 0 {
		ama.applyFlatDamage(sim, critChance)
	}

	if ama.WeaponDamageInput.MainHand > 0 {
		ama.applySwingDamage(sim, proto.ItemSlot_ItemSlotMainHand, ama.WeaponDamageInput.MainHand, critChance)
	}

	if weapon := c.Equip[proto.ItemSlot_ItemSlotOffHand]; ama.WeaponDamageInput.Offhand > 0 && weapon.ID > 0 { // only attack if we have it
		ama.applySwingDamage(sim, proto.ItemSlot_ItemSlotOffHand, ama.WeaponDamageInput.Offhand, critChance)
	}

	return true
}

func (ama *ActiveMeleeAbility) applySwingDamage(sim *Simulation, slot proto.ItemSlot, dmgMult, critChance float64) {
	c := ama.Character
	roll := sim.RandomFloat("weapon swing")
	hit := MeleeHitTypeHit
	if roll < critChance {
		hit = MeleeHitTypeCrit
		dmgMult *= 2
		ama.Crits++
	}

	weapon := c.Equip[slot]
	speed := ama.NormalizeWeaponSpeed
	if speed == 0 {
		speed = weapon.SwingSpeed
	}
	dmg := meleeDamage(sim, weapon.WeaponDamageMin, weapon.WeaponDamageMax, ama.WeaponDamageInput.MainHandFlat, speed, false, ama.WeaponDamageInput.MainHand, c.stats[stats.AttackPower], ama.Target.ArmorDamageReduction())
	if sim.Log != nil {
		sim.Log("%s mainhand for %0.1f", ama.Name, dmg)
	}
	if ama.OnMeleeAttack != nil {
		ama.OnMeleeAttack(sim, ama.Target, hit, ama, false)
	}
	c.OnMeleeAttack(sim, ama.Target, ama.Result, ama, false)
	ama.Hits++
	ama.TotalDamage += dmg
}

func (ama *ActiveMeleeAbility) applyFlatDamage(sim *Simulation, critChance float64) {
	roll := sim.RandomFloat("weapon swing")
	hit := MeleeHitTypeHit
	dmgMult := 1.0
	if roll < critChance {
		hit = MeleeHitTypeCrit
		dmgMult = 2.0
		ama.Crits++
	}

	// Do a 'direct damage' if ability has it
	dmg := ama.DirectDamageInput.MinBaseDamage + (ama.DirectDamageInput.MaxBaseDamage-ama.DirectDamageInput.MinBaseDamage)*sim.RandomFloat("melee direct damage") + ama.DirectDamageInput.FlatDamageBonus
	ama.TotalDamage += dmg * dmgMult
	ama.Hits++
	if hit == MeleeHitTypeCrit {
		ama.Crits++
	}
	if sim.Log != nil {
		sim.Log("%s for %0.1f", ama.Name, dmg)
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

	// Bonus stats to be added to the spell.
	BonusSpellHitRating  float64
	BonusSpellPower      float64
	BonusSpellCritRating float64

	// Additional multiplier that is always applied.
	DamageMultiplier float64

	// applies fixed % increases to damage at cast time.
	//  Only use multipliers that don't change for the lifetime of the sim.
	//  This should probably only be mutated in a template and not changed in auras.
	StaticDamageMultiplier float64
}

func NewAutoAttacks(c *Character) AutoAttacks {
	st := AutoAttacks{
		c:                c,
		DamageMultiplier: 1.0,
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

	MainhandSwingAt  time.Duration
	OffhandSwingAt   time.Duration
	DamageMultiplier float64 // auto attack damage multiplier
}

func (aa *AutoAttacks) MainhandSwingSpeed() time.Duration {
	return time.Duration(float64(aa.mhbase) / aa.c.SwingSpeed())
}

func (aa *AutoAttacks) OffhandSwingSpeed() time.Duration {
	return time.Duration(float64(aa.ohbase) / aa.c.SwingSpeed())
}

// Swing will check any swing timers if they are up, and if so, swing!
func (aa *AutoAttacks) Swing(sim *Simulation, target *Target) {
	if aa.MainhandSwingAt <= sim.CurrentTime {
		// Make a MH swing!
		PerformAutoAttack(sim, aa.c, target, aa.mh, aa.DamageMultiplier, false)
		aa.MainhandSwingAt = sim.CurrentTime + aa.MainhandSwingSpeed()
	}
	if aa.OffhandSwingAt <= sim.CurrentTime {
		// Make a OH swing!
		PerformAutoAttack(sim, aa.c, target, aa.oh, aa.DamageMultiplier, true)
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
		aa.OffhandSwingAt = sim.CurrentTime + time.Duration(float64(ohSwingTime)/amount)
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
