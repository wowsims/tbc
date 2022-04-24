package mage

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var IgniteActionID = core.ActionID{SpellID: 12848}

func (mage *Mage) registerIgniteSpell(sim *core.Simulation) {
	mage.Ignite = mage.RegisterSpell(core.SpellConfig{
		ActionID:    IgniteActionID,
		SpellSchool: core.SpellSchoolFire,
		SpellExtras: SpellFlagMage | core.SpellExtrasIgnoreModifiers,
	})
}

func (mage *Mage) newIgniteDot(sim *core.Simulation, target *core.Target) *core.Dot {
	return core.NewDot(core.Dot{
		Spell: mage.Ignite,
		Aura: target.RegisterAura(core.Aura{
			Label:    "Ignite-" + strconv.Itoa(int(mage.Index)),
			ActionID: IgniteActionID,
		}),
		NumberOfTicks: 2,
		TickLength:    time.Second * 2,
	})
}

func (mage *Mage) procIgnite(sim *core.Simulation, target *core.Target, damageFromProccingSpell float64) {
	igniteDot := mage.IgniteDots[target.Index]

	newIgniteDamage := damageFromProccingSpell * float64(mage.Talents.Ignite) * 0.08
	if igniteDot.IsActive() {
		newIgniteDamage += mage.IgniteTickDamage[target.Index] * float64(2-igniteDot.TickCount)
	}

	newTickDamage := newIgniteDamage / 2
	mage.IgniteTickDamage[target.Index] = newTickDamage

	if sim.Log != nil {
		mage.Log(sim, "Casting %s (Cost = %0.03f, Cast Time = %s)", IgniteActionID, 0.0, time.Duration(0))
		mage.Log(sim, "Completed cast %s", IgniteActionID)
	}
	mage.Ignite.SpellMetrics[target.Index].Casts++
	mage.Ignite.SpellMetrics[target.Index].Hits++

	// Reassign the effect to apply the new damage value.
	igniteDot.TickEffects = core.TickFuncSnapshot(target, core.SpellEffect{
		DamageMultiplier: 1,
		ThreatMultiplier: 1 - 0.05*float64(mage.Talents.BurningSoul),
		IsPeriodic:       true,
		BaseDamage:       core.BaseDamageConfigFlat(newTickDamage),
		OutcomeApplier:   core.OutcomeFuncTick(),
	})
	igniteDot.Apply(sim)
}

func (mage *Mage) applyIgnite() {
	if mage.Talents.Ignite == 0 {
		return
	}

	mage.RegisterAura(core.Aura{
		Label:    "Ignite Talent",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
				return
			}
			if spell.SpellSchool == core.SpellSchoolFire && spellEffect.Outcome.Matches(core.OutcomeCrit) {
				mage.procIgnite(sim, spellEffect.Target, spellEffect.Damage)
			}
		},
	})
}
