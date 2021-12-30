package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDFrostbolt int32 = 27072

func (mage *Mage) newFrostboltTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				Name:           "Frostbolt",
				CritMultiplier: 1.5 + 0.125*float64(mage.Talents.SpellPower) + 0.1*float64(mage.Talents.IceShards),
				SpellSchool:    stats.FrostSpellPower,
				Character:      &mage.Character,
				BaseManaCost:   330,
				ManaCost:       330,
				CastTime:       time.Millisecond * 3000,
				ActionID: core.ActionID{
					SpellID: SpellIDFrostbolt,
				},
			},
		},
		SpellHitEffect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: mage.spellDamageMultiplier,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage:    600,
				MaxBaseDamage:    647,
				SpellCoefficient: (3.0 / 3.5) * 0.95,
				// TODO: binary: true,
			},
		},
	}

	spell.CastTime -= time.Millisecond * 100 * time.Duration(mage.Talents.ImprovedFrostbolt)
	spell.ManaCost -= spell.BaseManaCost * float64(mage.Talents.FrostChanneling) * 0.05
	spell.SpellHitEffect.SpellEffect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance
	spell.SpellHitEffect.SpellEffect.BonusSpellCritRating += float64(mage.Talents.EmpoweredFrostbolt) * 1 * core.SpellCritRatingPerCritChance
	spell.SpellHitEffect.SpellEffect.StaticDamageMultiplier *= 1 + 0.02*float64(mage.Talents.PiercingIce)
	spell.SpellHitEffect.SpellEffect.StaticDamageMultiplier *= 1 + 0.01*float64(mage.Talents.ArcticWinds)
	spell.SpellHitEffect.DirectInput.SpellCoefficient += 0.02 * float64(mage.Talents.EmpoweredFrostbolt)

	if ItemSetTempestRegalia.CharacterHasSetBonus(&mage.Character, 4) {
		spell.SpellHitEffect.SpellEffect.StaticDamageMultiplier *= 1.05
	}

	// TODO: Winters chill applies to other frost spells also, but we don't implement them yet.
	if mage.Talents.WintersChill > 0 {
		procChance := float64(mage.Talents.WintersChill) / 5.0
		spell.SpellHitEffect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			// Don't overwrite the permanent version.
			if spellEffect.Target.RemainingAuraDuration(sim, core.WintersChillDebuffID) == core.NeverExpires {
				return
			}

			if procChance != 1.0 || sim.RandomFloat("Winters Chill") > procChance {
				return
			}

			newNumStacks := core.MinInt32(5, spellEffect.Target.NumStacks(core.WintersChillDebuffID)+1)
			spellEffect.Target.ReplaceAura(sim, core.WintersChillAura(sim, newNumStacks))
		}
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (mage *Mage) NewFrostbolt(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	frostbolt := &mage.frostboltSpell
	mage.frostboltCastTemplate.Apply(frostbolt)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	frostbolt.Target = target
	frostbolt.Init(sim)

	return frostbolt
}
