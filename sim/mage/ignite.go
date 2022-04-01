package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var IgniteActionID = core.ActionID{SpellID: 12848}
var IgniteDebuffID = core.NewDebuffID()

func (mage *Mage) newIgniteSpell(sim *core.Simulation) *core.Spell {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    IgniteActionID,
				Character:   &mage.Character,
				SpellSchool: core.SpellSchoolFire,
				SpellExtras: SpellFlagMage | core.SpellExtrasBinary | core.SpellExtrasAlwaysHits,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryNone,
			DamageMultiplier:    1,
			ThreatMultiplier:    1 - 0.05*float64(mage.Talents.BurningSoul),
			DotInput: core.DotDamageInput{
				NumberOfTicks:         2,
				TickLength:            time.Second * 2,
				IgnoreDamageModifiers: true,
				DebuffID:              IgniteDebuffID,
			},
		},
	}

	return mage.RegisterSpell(core.SpellConfig{
		Template:   spell,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

func (mage *Mage) procIgnite(sim *core.Simulation, target *core.Target, damageFromProccingSpell float64) {
	newIgniteDamage := damageFromProccingSpell * float64(mage.Talents.Ignite) * 0.08
	ignite := mage.Ignites[target.Index]

	if ignite.Instance.Effect.DotInput.IsTicking(sim) {
		newIgniteDamage += ignite.Instance.Effect.DotInput.RemainingDamage()
	}

	// Cancel the current ignite dot.
	ignite.Instance.Cancel(sim)
	ignite.Template.Effect.DotInput.TickBaseDamage = core.DotSnapshotFuncMagic(newIgniteDamage/2, 0)
	ignite.Cast(sim, target)
}

var IgniteAuraID = core.NewAuraID()

func (mage *Mage) applyIgnite() {
	if mage.Talents.Ignite == 0 {
		return
	}

	mage.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: IgniteAuraID,
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}
				if spell.SpellSchool == core.SpellSchoolFire && spellEffect.Outcome.Matches(core.OutcomeCrit) {
					mage.procIgnite(sim, spellEffect.Target, spellEffect.Damage)
				}
			},
		}
	})
}
