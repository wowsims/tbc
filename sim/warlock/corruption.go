package warlock

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDCorruption8 int32 = 27216

var Corruption8ActionID = core.ActionID{SpellID: SpellIDCorruption8}

func (warlock *Warlock) registerCorruptionSpell(sim *core.Simulation) {
	baseCost := 370.0
	warlock.Corruption = warlock.RegisterSpell(core.SpellConfig{
		ActionID:     Corruption8ActionID,
		SpellSchool:  core.SpellSchoolShadow,
		ResourceType: stats.Mana,
		BaseCost:     baseCost,
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost:     baseCost,
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond*2000 - (time.Millisecond * 400 * time.Duration(warlock.Talents.ImprovedCorruption)),
			},
		},
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			OutcomeApplier: core.OutcomeFuncMagicHit(),
			OnSpellHit:     applyDotOnLanded(&warlock.CorruptionDot),
		}),
	})
	target := sim.GetPrimaryTarget()
	spellCoefficient := 0.156 + (0.12 * float64(warlock.Talents.EmpoweredCorruption))

	warlock.CorruptionDot = core.NewDot(core.Dot{
		Spell: warlock.Corruption,
		Aura: target.RegisterAura(core.Aura{
			Label:    "Corruption-" + strconv.Itoa(int(warlock.Index)),
			ActionID: Corruption8ActionID,
		}),
		NumberOfTicks: 6,
		TickLength:    time.Second * 3,
		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1 * (1 + 0.02*float64(warlock.Talents.ShadowMastery)) * (1 + 0.01*float64(warlock.Talents.Contagion)),
			ThreatMultiplier: 1 - 0.05*float64(warlock.Talents.ImprovedDrainSoul),
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(900/6, spellCoefficient),
			OutcomeApplier:   core.OutcomeFuncTick(),
			IsPeriodic:       true,
		}),
	})
}

func applyDotOnLanded(dot **core.Dot) func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
	return func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			(*dot).Apply(sim)
		}
	}
}
