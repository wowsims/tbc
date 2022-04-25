package warlock

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var CurseOfElementsActionID = core.ActionID{SpellID: 27228}
var CurseOfRecklessnessActionID = core.ActionID{SpellID: 27226}
var CurseOfTonguesActionID = core.ActionID{SpellID: 11719}
var CurseOfAgonyActionID = core.ActionID{SpellID: 27218}
var CurseOfDoomActionID = core.ActionID{SpellID: 30910}

func (warlock *Warlock) registerCurseOfElementsSpell(sim *core.Simulation) {
	baseCost := 145.0
	auras := sim.GetPrimaryTarget().GetAurasWithTag("Curse of Elements")
	for _, aura := range auras {
		if int32(aura.Priority) >= warlock.Talents.Malediction {
			// Someone else with at least as good of curse is already doing it... lets not.
			warlock.Rotation.Curse = proto.Warlock_Rotation_NoCurse // TODO: swap to agony for dps?
			return
		}
	}
	warlock.CurseOfElementsAura = core.CurseOfElementsAura(sim.GetPrimaryTarget(), warlock.Talents.Malediction)
	warlock.CurseOfElementsAura.Duration = time.Minute * 5

	warlock.CurseOfElements = warlock.RegisterSpell(core.SpellConfig{
		ActionID:    CurseOfElementsActionID,
		SpellSchool: core.SpellSchoolShadow,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  core.GCDDefault,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ThreatMultiplier: 1,
			FlatThreatBonus:  0, // TODO
			OutcomeApplier:   core.OutcomeFuncMagicHit(),
			OnSpellHit:       applyAuraOnLanded(warlock.CurseOfElementsAura),
		}),
	})
}

func (warlock *Warlock) ShouldCastCurseOfElements(sim *core.Simulation, target *core.Target, curse proto.Warlock_Rotation_Curse) bool {
	return curse == proto.Warlock_Rotation_Elements && !warlock.CurseOfElementsAura.IsActive()
}

func (warlock *Warlock) registerCurseOfRecklessnessSpell(sim *core.Simulation) {
	baseCost := 160.0
	warlock.CurseOfRecklessnessAura = core.CurseOfRecklessnessAura(sim.GetPrimaryTarget())
	warlock.CurseOfRecklessnessAura.Duration = time.Minute * 2

	warlock.CurseOfRecklessness = warlock.RegisterSpell(core.SpellConfig{
		ActionID:     CurseOfElementsActionID,
		SpellSchool:  core.SpellSchoolShadow,
		ResourceType: stats.Mana,
		BaseCost:     baseCost,
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  core.GCDDefault,
			},
		},
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ThreatMultiplier: 1,
			FlatThreatBonus:  0, // TODO
			OutcomeApplier:   core.OutcomeFuncMagicHit(),
			OnSpellHit:       applyAuraOnLanded(warlock.CurseOfRecklessnessAura),
		}),
	})
}

// https://tbc.wowhead.com/spell=11719/curse-of-tongues
func (warlock *Warlock) registerCurseOfTonguesSpell(sim *core.Simulation) {
	baseCost := 110.0
	// Empty aura so we can simulate cost/time to keep tongues up
	warlock.CurseOfTonguesAura = sim.GetPrimaryTarget().GetOrRegisterAura(core.Aura{
		Label:    "Curse of Tongues",
		ActionID: CurseOfTonguesActionID,
		Duration: time.Second * 30,
	})

	warlock.CurseOfTongues = warlock.RegisterSpell(core.SpellConfig{
		ActionID:     CurseOfTonguesActionID,
		SpellSchool:  core.SpellSchoolShadow,
		ResourceType: stats.Mana,
		BaseCost:     baseCost,
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  core.GCDDefault,
			},
		},
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ThreatMultiplier: 1,
			FlatThreatBonus:  0, // TODO
			OutcomeApplier:   core.OutcomeFuncMagicHit(),
			OnSpellHit:       applyAuraOnLanded(warlock.CurseOfTonguesAura),
		}),
	})
}

// https://tbc.wowhead.com/spell=27218/curse-of-agony
func (warlock *Warlock) registerCurseOfAgonySpell(sim *core.Simulation) {
	baseCost := 265.0
	target := sim.GetPrimaryTarget()
	effect := core.SpellEffect{
		DamageMultiplier: 1 *
			(1 + 0.02*float64(warlock.Talents.ShadowMastery)) *
			(1 + 0.01*float64(warlock.Talents.Contagion)) *
			(1 + 0.02*float64(warlock.Talents.ImprovedCurseOfAgony)),
		ThreatMultiplier: 1 - 0.05*float64(warlock.Talents.ImprovedDrainSoul),
		BaseDamage:       core.BaseDamageConfigMagicNoRoll(1356/12, 0.1),
		OutcomeApplier:   core.OutcomeFuncTick(),
		IsPeriodic:       true,
	}
	// Amplify Curse talent
	if warlock.Talents.AmplifyCurse {
		effect.BaseDamage = core.WrapBaseDamageConfig(effect.BaseDamage, func(oldCalculator core.BaseDamageCalculator) core.BaseDamageCalculator {
			return func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
				if warlock.AmplifyCurseAura.IsActive() {
					panic("not implemented")
					return oldCalculator(sim, hitEffect, spell) * 1.5
				} else {
					return oldCalculator(sim, hitEffect, spell)
				}
			}
		})
		// Make sure a hit of this spell deactivates any active amp curse
		effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			warlock.AmplifyCurseAura.Deactivate(sim)
		}
	}
	warlock.CurseOfAgony = warlock.RegisterSpell(core.SpellConfig{
		ActionID:     CurseOfAgonyActionID,
		SpellSchool:  core.SpellSchoolShadow,
		ResourceType: stats.Mana,
		BaseCost:     baseCost,
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  core.GCDDefault,
			},
		},
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ThreatMultiplier: 1,
			FlatThreatBonus:  0, // TODO
			OutcomeApplier:   core.OutcomeFuncMagicHit(),
			OnSpellHit:       applyDotOnLanded(&warlock.CurseOfAgonyDot),
		}),
	})
	warlock.CurseOfAgonyDot = core.NewDot(core.Dot{
		Spell: warlock.CurseOfAgony,
		Aura: target.RegisterAura(core.Aura{
			Label:    "CurseofAgony-" + strconv.Itoa(int(warlock.Index)),
			ActionID: CurseOfAgonyActionID,
		}),
		NumberOfTicks: 12,
		TickLength:    time.Second * 2,
		TickEffects:   core.TickFuncSnapshot(target, effect),
	})
}

func (warlock *Warlock) registerCurseOfDoomSpell(sim *core.Simulation) {
	baseCost := 380.0
	target := sim.GetPrimaryTarget()
	effect := core.SpellEffect{
		DamageMultiplier: 1 *
			(1 + 0.02*float64(warlock.Talents.ShadowMastery)) *
			(1 + 0.01*float64(warlock.Talents.Contagion)),
		ThreatMultiplier: 1 - 0.05*float64(warlock.Talents.ImprovedDrainSoul),
		BaseDamage:       core.BaseDamageConfigMagicNoRoll(4200, 2),
		OutcomeApplier:   core.OutcomeFuncTick(),
		IsPeriodic:       true,
	}
	// Amplify Curse talent
	if warlock.Talents.AmplifyCurse {
		effect.BaseDamage = core.WrapBaseDamageConfig(effect.BaseDamage, func(oldCalculator core.BaseDamageCalculator) core.BaseDamageCalculator {
			return func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
				if warlock.AmplifyCurseAura.IsActive() {
					panic("not implemented")
					return oldCalculator(sim, hitEffect, spell) * 1.5
				} else {
					return oldCalculator(sim, hitEffect, spell)
				}
			}
		})
		// Make sure a hit of this spell deactivates any active amp curse
		effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			warlock.AmplifyCurseAura.Deactivate(sim)
		}
	}

	warlock.CurseOfDoom = warlock.RegisterSpell(core.SpellConfig{
		ActionID:     CurseOfDoomActionID,
		SpellSchool:  core.SpellSchoolShadow,
		ResourceType: stats.Mana,
		BaseCost:     baseCost,
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    warlock.NewTimer(),
				Duration: time.Minute,
			},
		},
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ThreatMultiplier: 1,
			FlatThreatBonus:  0, // TODO
			OutcomeApplier:   core.OutcomeFuncMagicHit(),
			OnSpellHit:       applyDotOnLanded(&warlock.CurseOfDoomDot),
		}),
	})

	warlock.CurseOfDoomDot = core.NewDot(core.Dot{
		Spell: warlock.CurseOfDoom,
		Aura: target.RegisterAura(core.Aura{
			Label:    "CurseofDoom-" + strconv.Itoa(int(warlock.Index)),
			ActionID: CurseOfDoomActionID,
		}),
		NumberOfTicks: 1,
		TickLength:    time.Minute,
		TickEffects:   core.TickFuncSnapshot(target, effect),
	})
}

func applyAuraOnLanded(aura *core.Aura) func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
	return func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			aura.Activate(sim)
		}
	}
}
