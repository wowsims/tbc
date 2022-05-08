package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (paladin *Paladin) ApplyTalents() {
	paladin.applyConviction()
	paladin.applyCrusade()
	paladin.applyTwoHandedWeaponSpecialization()
	paladin.applySanctityAura()
	paladin.applyVengeance()
	paladin.applySanctifiedSeals()
	paladin.applyPrecision()
	paladin.applyDivineStrength()
	paladin.applyDivineIntellect()
}

func (paladin *Paladin) applyConviction() {
	paladin.AddStat(stats.MeleeCrit, core.MeleeCritRatingPerCritChance*float64(paladin.Talents.Conviction))
}

func (paladin *Paladin) applyCrusade() {
	if paladin.Talents.Crusade == 0 {
		return
	}

	damageMultiplier := 1 + (0.01 * float64(paladin.Talents.Crusade)) // assume multiplicative scaling

	// TO-DO: This doesn't account for multiple targets
	switch paladin.Env.GetPrimaryTarget().MobType {
	case proto.MobType_MobTypeHumanoid, proto.MobType_MobTypeDemon, proto.MobType_MobTypeUndead, proto.MobType_MobTypeElemental:
		paladin.PseudoStats.DamageDealtMultiplier *= damageMultiplier
	}
}

// Affects all physical damage or spells that can be rolled as physical
// It affects white, Windfury, Crusader Strike, Seals, and Judgement of Command / Blood
func (paladin *Paladin) applyTwoHandedWeaponSpecialization() {
	if paladin.Talents.TwoHandedWeaponSpecialization == 0 {
		return
	}

	// This impacts Crusader Strike, Melee Attacks, WF attacks
	// Seals + Judgements need to be implemented separately
	if paladin.GetMHWeapon().HandType == proto.HandType_HandTypeTwoHand {
		paladin.PseudoStats.PhysicalDamageDealtMultiplier *= 1 + (0.02 * float64(paladin.Talents.TwoHandedWeaponSpecialization)) // assume multiplicative scaling
	}
}

func (paladin *Paladin) applyTwoHandedWeaponSpecializationToSpell(spellEffect *core.SpellEffect) {
	if paladin.GetMHWeapon().HandType == proto.HandType_HandTypeTwoHand {
		spellEffect.DamageMultiplier *= 1 + (0.02 * float64(paladin.Talents.TwoHandedWeaponSpecialization))
	}
}

// Apply as permanent aura only to self for now
// TO-DO: Maybe should put this in the partybuff section instead at some point
func (paladin *Paladin) applySanctityAura() {
	if paladin.Talents.SanctityAura {
		core.SanctityAura(&paladin.Character, float64(paladin.Talents.ImprovedSanctityAura))
	}
}

// I don't know if the new stack of vengeance applies to the crit that triggered it or not
// Need to check this
func (paladin *Paladin) applyVengeance() {
	if paladin.Talents.Vengeance == 0 {
		return
	}

	bonusPerStack := 0.01 * float64(paladin.Talents.Vengeance)
	procAura := paladin.RegisterAura(core.Aura{
		Label:     "Vengeance Proc",
		ActionID:  core.ActionID{SpellID: 20059},
		Duration:  time.Second * 30,
		MaxStacks: 3,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			aura.Unit.PseudoStats.DamageDealtMultiplier /= 1 + (bonusPerStack * float64(oldStacks))
			aura.Unit.PseudoStats.DamageDealtMultiplier *= 1 + (bonusPerStack * float64(newStacks))
		},
	})

	paladin.RegisterAura(core.Aura{
		Label:    "Vengeance",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Outcome.Matches(core.OutcomeCrit) {
				procAura.Activate(sim)
				procAura.AddStack(sim)
			}
		},
	})
}

func (paladin *Paladin) applySanctifiedSeals() {
	paladin.AddStat(stats.MeleeCrit, core.MeleeCritRatingPerCritChance*float64(paladin.Talents.SanctifiedSeals))
	paladin.AddStat(stats.SpellCrit, core.SpellCritRatingPerCritChance*float64(paladin.Talents.SanctifiedSeals))
}

func (paladin *Paladin) applyPrecision() {
	paladin.AddStat(stats.MeleeHit, core.MeleeHitRatingPerHitChance*float64(paladin.Talents.Precision))
	paladin.AddStat(stats.SpellHit, core.SpellHitRatingPerHitChance*float64(paladin.Talents.Precision))
}

func (paladin *Paladin) applyDivineStrength() {
	bonusStr := paladin.GetStat(stats.Strength) * 0.02 * float64(paladin.Talents.DivineStrength)
	paladin.AddStat(stats.Strength, bonusStr)
}

func (paladin *Paladin) applyDivineIntellect() {
	bonusInt := paladin.GetStat(stats.Intellect) * 0.02 * float64(paladin.Talents.DivineIntellect)
	paladin.AddStat(stats.Intellect, bonusInt)
}
