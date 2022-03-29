package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

const SpellIDIgnite int32 = 12848

var IgniteDebuffID = core.NewDebuffID()

func (mage *Mage) newIgniteTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID: SpellIDIgnite,
				},
				Character:           &mage.Character,
				CritRollCategory:    core.CritRollCategoryNone,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolFire,
				SpellExtras:         SpellFlagMage | core.SpellExtrasBinary | core.SpellExtrasAlwaysHits,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier: 1,
				ThreatMultiplier: 1 - 0.05*float64(mage.Talents.BurningSoul),
			},
			DotInput: core.DotDamageInput{
				NumberOfTicks:         2,
				TickLength:            time.Second * 2,
				IgnoreDamageModifiers: true,
				DebuffID:              IgniteDebuffID,
			},
		},
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (mage *Mage) procIgnite(sim *core.Simulation, target *core.Target, damageFromProccingSpell float64) {
	newIgniteDamage := damageFromProccingSpell * float64(mage.Talents.Ignite) * 0.08
	ignite := &mage.igniteSpells[target.Index]

	if ignite.Effect.DotInput.IsTicking(sim) {
		newIgniteDamage += ignite.Effect.DotInput.RemainingDamage()
	}

	// Cancel the current ignite dot.
	ignite.Cancel(sim)

	mage.igniteCastTemplate.Apply(ignite)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ignite.Effect.Target = target
	ignite.Effect.DotInput.TickBaseDamage = core.DotSnapshotFuncMagic(newIgniteDamage/2, 0)
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
				if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}
				if spellCast.SpellSchool == core.SpellSchoolFire && spellEffect.Outcome.Matches(core.OutcomeCrit) {
					mage.procIgnite(sim, spellEffect.Target, spellEffect.Damage)
				}
			},
		}
	})
}
