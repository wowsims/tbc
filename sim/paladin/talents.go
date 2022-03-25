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

var CrusadeAuraID = core.NewAuraID()

// Maybe don't make this an aura but we'll do it for now
func (paladin *Paladin) applyCrusade() {
	if paladin.Talents.Crusade == 0 {
		return
	}

	paladin.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: CrusadeAuraID,
			OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
				target := spellEffect.Target

				if target.MobType == proto.MobType_MobTypeDemon || target.MobType == proto.MobType_MobTypeHumanoid ||
					target.MobType == proto.MobType_MobTypeUndead || target.MobType == proto.MobType_MobTypeElemental {
					spellEffect.DamageMultiplier *= 1 + (0.01 * float64(paladin.Talents.Crusade)) // assume multiplicative scaling
				}
			},
			OnBeforePeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage *float64) {
				target := spellEffect.Target

				if target.MobType == proto.MobType_MobTypeDemon || target.MobType == proto.MobType_MobTypeHumanoid ||
					target.MobType == proto.MobType_MobTypeUndead || target.MobType == proto.MobType_MobTypeElemental {
					*tickDamage *= 1 + (0.01 * float64(paladin.Talents.Crusade))
				}
			},
		}
	})
}

var TwoHandedWeaponSpecializationAuraID = core.NewAuraID()

// Affects all physical damage or spells that can be rolled as physical
// It affects white, Windfury, Crusader Strike, Seals, and Judgement of Command / Blood
func (paladin *Paladin) applyTwoHandedWeaponSpecialization() {
	if paladin.Talents.TwoHandedWeaponSpecialization == 0 {
		return
	}

	if paladin.GetMHWeapon().HandType == proto.HandType_HandTypeTwoHand {
		paladin.AddPermanentAura(func(sim *core.Simulation) core.Aura {
			return core.Aura{
				ID: TwoHandedWeaponSpecializationAuraID,
				OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
					if spellCast.OutcomeRollCategory.Matches(core.OutcomeRollCategoryPhysical) {
						spellEffect.StaticDamageMultiplier *= 1 + (0.02 * float64(paladin.Talents.TwoHandedWeaponSpecialization)) // assume multiplicative scaling
					}
				},
			}
		})
	}
}

// Apply as permanent aura only to self for now
// Maybe should put this in the partybuff section instead at some point
func (paladin *Paladin) applySanctityAura() {
	if !paladin.Talents.SanctityAura {
		return
	}

	paladin.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.ImprovedSanctityAura(sim, float64(paladin.Talents.ImprovedSanctityAura))
	})

}

var VengeanceAuraID = core.NewAuraID()
var VengeanceActionID = core.ActionID{SpellID: 20059}

const VengeanceDuration = time.Second * 30

var VengeancePermAuraID = core.NewAuraID()

// I don't know if the new stack of vengeance applies to the crit that triggered it or not
// Need to check this
func (paladin *Paladin) applyVengeance() {
	if paladin.Talents.Vengeance == 0 {
		return
	}

	vng := core.Aura{
		ID:       VengeanceAuraID,
		ActionID: VengeanceActionID,
		Duration: VengeanceDuration,
		Stacks:   0,
	}

	// Maybe a better way to do this than a perm aura that applies the buff and then increases damage based on it?
	paladin.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: VengeancePermAuraID,
			OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
				if spellCast.SpellSchool.Matches(core.SpellSchoolHoly | core.SpellSchoolPhysical) {
					spellEffect.DamageMultiplier *= 1 + (0.01*float64(paladin.Talents.Vengeance))*float64(paladin.NumStacks(VengeanceAuraID))
				}
			},
			OnBeforePeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage *float64) {
				if spellCast.SpellSchool.Matches(core.SpellSchoolHoly | core.SpellSchoolPhysical) {
					*tickDamage *= 1 + (0.01*float64(paladin.Talents.Vengeance))*float64(paladin.NumStacks(VengeanceAuraID))
				}
			},
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellEffect.Outcome.Matches(core.OutcomeCrit) {
					vng.Stacks = core.MinInt32(3, paladin.NumStacks(VengeanceAuraID)+1)
					paladin.ReplaceAura(sim, vng)
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
