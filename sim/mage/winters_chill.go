package mage

import (
	"github.com/wowsims/tbc/sim/core"
)

const SpellIDWintersChill int32 = 28595

var WintersChillActionID = core.ActionID{SpellID: SpellIDWintersChill}

// Winters Chill has a separate hit check from frostbolt, so it needs its own spell.
func (mage *Mage) registerWintersChillSpell(sim *core.Simulation) {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    WintersChillActionID,
				Character:   &mage.Character,
				SpellSchool: core.SpellSchoolFrost,
				SpellExtras: SpellFlagMage,
			},
		},
	}

	effect := core.SpellEffect{
		BonusSpellHitRating: float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance,
		ThreatMultiplier:    1,
		OutcomeApplier:      core.OutcomeFuncMagicHit(),
	}

	if mage.Talents.WintersChill > 0 {
		wcAura := sim.GetPrimaryTarget().GetAura(core.WintersChillAuraLabel)
		if wcAura == nil {
			wcAura = core.WintersChillAura(sim.GetPrimaryTarget(), 0)
		}

		effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Landed() {
				wcAura.Activate(sim)
				wcAura.AddStack(sim)
			}
		}
	}

	mage.WintersChill = mage.RegisterSpell(core.SpellConfig{
		Template:     spell,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(effect),
	})
}

func (mage *Mage) applyWintersChill() {
	if mage.Talents.WintersChill == 0 {
		return
	}

	procChance := float64(mage.Talents.WintersChill) / 5.0

	mage.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return mage.GetOrRegisterAura(&core.Aura{
			Label: "Winters Chill Talent",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}

				if spell.SpellSchool == core.SpellSchoolFrost && !spell.SameAction(WintersChillActionID) {
					if procChance != 1.0 && sim.RandomFloat("Winters Chill") > procChance {
						return
					}

					mage.WintersChill.Cast(sim, spellEffect.Target)
				}
			},
		})
	})
}
