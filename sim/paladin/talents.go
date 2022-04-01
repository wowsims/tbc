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

	paladin.RegisterResetEffect(func(sim *core.Simulation) {
		switch sim.GetPrimaryTarget().MobType {
		case proto.MobType_MobTypeHumanoid, proto.MobType_MobTypeDemon, proto.MobType_MobTypeUndead, proto.MobType_MobTypeElemental:
			paladin.PseudoStats.DamageDealtMultiplier *= damageMultiplier
		}
	})
}

// Affects all physical damage or spells that can be rolled as physical
// It affects white, Windfury, Crusader Strike, Seals, and Judgement of Command / Blood
func (paladin *Paladin) applyTwoHandedWeaponSpecialization() {
	if paladin.Talents.TwoHandedWeaponSpecialization == 0 {
		return
	}

	if paladin.GetMHWeapon().HandType == proto.HandType_HandTypeTwoHand {
		paladin.PseudoStats.PhysicalDamageDealtMultiplier *= 1 + (0.02 * float64(paladin.Talents.TwoHandedWeaponSpecialization)) // assume multiplicative scaling
		// TODO: Might need to additionally apply this to non-physical spells directly.
	}
}

// Apply as permanent aura only to self for now
// Maybe should put this in the partybuff section instead at some point
func (paladin *Paladin) applySanctityAura() {
	if !paladin.Talents.SanctityAura {
		return
	}

	paladin.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.ImprovedSanctityAura(&paladin.Character, float64(paladin.Talents.ImprovedSanctityAura))
	})

}

var VengeanceAuraID = core.NewAuraID()
var VengeanceProcAuraID = core.NewAuraID()
var VengeanceActionID = core.ActionID{SpellID: 20059}

const VengeanceDuration = time.Second * 30

// I don't know if the new stack of vengeance applies to the crit that triggered it or not
// Need to check this
func (paladin *Paladin) applyVengeance() {
	if paladin.Talents.Vengeance == 0 {
		return
	}

	multiplierPerStack := 1 + (0.01 * float64(paladin.Talents.Vengeance))

	makeProcAura := func(numStacks int32) core.Aura {
		multiplier := multiplierPerStack * float64(numStacks)
		return core.Aura{
			ID:       VengeanceProcAuraID,
			ActionID: VengeanceActionID,
			Duration: VengeanceDuration,
			Stacks:   numStacks,
			OnGain: func(sim *core.Simulation) {
				paladin.PseudoStats.DamageDealtMultiplier *= multiplier
			},
			OnExpire: func(sim *core.Simulation) {
				paladin.PseudoStats.DamageDealtMultiplier /= multiplier
			},
		}
	}

	paladin.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: VengeanceAuraID,
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Outcome.Matches(core.OutcomeCrit) {
					newStacks := core.MinInt32(3, paladin.NumStacks(VengeanceAuraID)+1)
					paladin.AddAura(sim, makeProcAura(newStacks))
				}
			},
		}
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
