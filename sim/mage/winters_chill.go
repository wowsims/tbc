package mage

import (
	"github.com/wowsims/tbc/sim/core"
)

const SpellIDWintersChill int32 = 28595

// Winters Chill has a separate hit check from frostbolt, so it needs its own spell.
func (mage *Mage) registerWintersChillSpell(sim *core.Simulation) {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: SpellIDWintersChill},
				Character:   &mage.Character,
				SpellSchool: core.SpellSchoolFrost,
				SpellExtras: SpellFlagMage,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		},
	}

	spell.Effect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance

	spell.Effect.OnSpellHit = func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
		if !spellEffect.Landed() {
			return
		}

		// Don't overwrite the permanent version.
		if spellEffect.Target.RemainingAuraDuration(sim, core.WintersChillDebuffID) == core.NeverExpires {
			return
		}

		newNumStacks := core.MinInt32(5, spellEffect.Target.NumStacks(core.WintersChillDebuffID)+1)
		spellEffect.Target.AddAura(sim, core.WintersChillAura(spellEffect.Target, newNumStacks))
	}

	mage.WintersChill = mage.RegisterSpell(core.SpellConfig{
		Template:   spell,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

var WintersChillAuraID = core.NewAuraID()

func (mage *Mage) applyWintersChill() {
	if mage.Talents.WintersChill == 0 {
		return
	}

	procChance := float64(mage.Talents.WintersChill) / 5.0

	mage.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: WintersChillAuraID,
			OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}

				if spell.SpellSchool == core.SpellSchoolFrost && spell.ActionID.SpellID != SpellIDWintersChill {
					if procChance != 1.0 && sim.RandomFloat("Winters Chill") > procChance {
						return
					}

					mage.WintersChill.Cast(sim, spellEffect.Target)
				}
			},
		}
	})
}
