package warlock

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var SeedActionID = core.ActionID{SpellID: 27243}

func (warlock *Warlock) registerSeedSpell(sim *core.Simulation) {
	seedExplosion := warlock.RegisterSpell(core.SpellConfig{
		ActionID:    SeedActionID,
		SpellSchool: core.SpellSchoolShadow,
		Cast:        core.CastConfig{},
		ApplyEffects: core.ApplyEffectFuncAOEDamageCapped(sim, 13580, core.SpellEffect{
			ProcMask:         core.ProcMaskSpellDamage,
			DamageMultiplier: 1 * (1 + 0.02*float64(warlock.Talents.ShadowMastery)) * (1 + 0.01*float64(warlock.Talents.Contagion)),
			ThreatMultiplier: 1 - 0.05*float64(warlock.Talents.ImprovedDrainSoul),
			BaseDamage:       core.BaseDamageConfigMagic(1110, 1290, 0.143),
			OutcomeApplier:   warlock.OutcomeFuncMagicHitAndCrit(warlock.SpellCritMultiplier(1, core.TernaryFloat64(warlock.Talents.Ruin, 1, 0))),
		}),
	})
	numTargets := int(sim.GetNumTargets())
	warlock.Seeds = make([]*core.Spell, numTargets)
	warlock.SeedDots = make([]*core.Dot, numTargets)

	for i := 0; i < numTargets; i++ {
		warlock.makeSeed(sim, i, seedExplosion)
	}

}

func (warlock *Warlock) makeSeed(sim *core.Simulation, targetIdx int, seedExplosion *core.Spell) {
	baseCost := 882.0

	warlock.Seeds[targetIdx] = warlock.RegisterSpell(core.SpellConfig{
		ActionID:     SeedActionID,
		SpellSchool:  core.SpellSchoolShadow,
		ResourceType: stats.Mana,
		BaseCost:     baseCost,
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost:     baseCost,
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond * 2000,
			},
		},
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:       core.ProcMaskEmpty,
			OutcomeApplier: warlock.OutcomeFuncMagicHit(),
			OnSpellHit:     applyDotOnLanded(&warlock.SeedDots[targetIdx]),
		}),
	})
	target := sim.GetTarget(int32(targetIdx))

	seedDmgTracker := 0.0
	warlock.SeedDots[targetIdx] = core.NewDot(core.Dot{
		Spell: warlock.Seeds[targetIdx],
		Aura: target.RegisterAura(core.Aura{
			Label:    "Seed-" + strconv.Itoa(int(warlock.Index)),
			ActionID: SeedActionID,
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}
				seedDmgTracker += spellEffect.Damage
				if seedDmgTracker > 1044 {
					warlock.SeedDots[targetIdx].Deactivate(sim)
					seedExplosion.Cast(sim, target)
					seedDmgTracker = 0
				}
			},
			OnPeriodicDamage: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				seedDmgTracker += spellEffect.Damage
				if seedDmgTracker > 1044 {
					warlock.SeedDots[targetIdx].Deactivate(sim)
					seedExplosion.Cast(sim, target)
					seedDmgTracker = 0
				}
			},
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				seedDmgTracker = 0
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				seedDmgTracker = 0
			},
		}),

		NumberOfTicks: 6,
		TickLength:    time.Second * 3,
		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1 * (1 + 0.02*float64(warlock.Talents.ShadowMastery)) * (1 + 0.01*float64(warlock.Talents.Contagion)),
			ThreatMultiplier: 1 - 0.05*float64(warlock.Talents.ImprovedDrainSoul),
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(174, 0.25),
			OutcomeApplier:   warlock.OutcomeFuncTick(),
			IsPeriodic:       true,
			ProcMask:         core.ProcMaskPeriodicDamage,
		}),
	})
}
