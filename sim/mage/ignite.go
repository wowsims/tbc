package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDIgnite int32 = 12848

var IgniteDebuffID = core.NewDebuffID()

func (mage *Mage) newIgniteTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				SpellSchool: stats.FireSpellPower,
				Character:   &mage.Character,
				ActionID: core.ActionID{
					SpellID: SpellIDIgnite,
				},
				IgnoreCooldowns: true,
				IgnoreManaCost:  true,
				Binary:          true,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				IgnoreHitCheck:         true,
			},
			DotInput: core.DotDamageInput{
				NumberOfTicks:         2,
				TickLength:            time.Second * 2,
				TickBaseDamage:        0, // This is set dynamically
				TickSpellCoefficient:  0,
				IgnoreDamageModifiers: true,
				DebuffID:              IgniteDebuffID,
			},
		},
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (mage *Mage) procIgnite(sim *core.Simulation, target *core.Target, damageFromProccingSpell float64) {
	newIgniteDamage := damageFromProccingSpell * float64(mage.Talents.Ignite) * 0.08
	if mage.igniteSpell.Effect.DotInput.IsTicking(sim) {
		newIgniteDamage += mage.igniteSpell.Effect.DotInput.RemainingDamage()
	}

	// Cancel the current ignite dot.
	mage.igniteSpell.Cancel(sim)

	ignite := &mage.igniteSpell
	mage.igniteCastTemplate.Apply(ignite)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ignite.Effect.Target = target
	ignite.Effect.DotInput.TickBaseDamage = newIgniteDamage / 2
	ignite.Init(sim)
	ignite.Cast(sim)
}

var IgniteAuraID = core.NewAuraID()

func (mage *Mage) applyIgnite() {
	if mage.Talents.Ignite == 0 {
		return
	}

	mage.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: IgniteAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellCast.SpellSchool == stats.FireSpellPower && spellEffect.Crit {
					mage.procIgnite(sim, spellEffect.Target, spellEffect.Damage)
				}
			},
		}
	})
}
