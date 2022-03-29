package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var SteadyShotActionID = core.ActionID{SpellID: 34120}

func (hunter *Hunter) newSteadyShotTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:  SteadyShotActionID,
				Character: hunter.GetCharacter(),
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 110,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 110,
				},
				// Cast time is affected by ranged attack speed so set it later.
				//CastTime:     time.Millisecond * 1500,
				GCD:                 core.GCDDefault + hunter.latency,
				IgnoreHaste:         true, // Hunter GCD is locked at 1.5s
				OutcomeRollCategory: core.OutcomeRollCategoryRanged,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				CritMultiplier:      hunter.critMultiplier(true, sim.GetPrimaryTarget()),
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:         core.ProcMaskRangedSpecial,
				DamageMultiplier: 1,
				ThreatMultiplier: 1,
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					hunter.killCommandBlocked = false
					hunter.TryKillCommand(sim, spellEffect.Target)
					hunter.rotation(sim, false)
				},
			},
			BaseDamage: hunter.talonOfAlarDamageMod(core.BaseDamageConfig{
				Calculator: func(sim *core.Simulation, hitEffect *core.SpellHitEffect, spellCast *core.SpellCast) float64 {
					return (hitEffect.RangedAttackPower(spellCast)+hitEffect.RangedAttackPowerOnTarget())*0.2 +
						hunter.AutoAttacks.Ranged.BaseDamage(sim)*2.8/hunter.AutoAttacks.Ranged.SwingSpeed +
						150
				},
				TargetSpellCoefficient: 1,
			}),
		},
	}

	ama.Cost.Value *= 1 - 0.02*float64(hunter.Talents.Efficiency)

	if ItemSetRiftStalker.CharacterHasSetBonus(&hunter.Character, 4) {
		ama.Effect.BonusCritRating += 5 * core.MeleeCritRatingPerCritChance
	}
	if ItemSetGronnstalker.CharacterHasSetBonus(&hunter.Character, 4) {
		ama.Effect.DamageMultiplier *= 1.1
	}

	return core.NewSimpleSpellTemplate(ama)
}

func (hunter *Hunter) NewSteadyShot(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	ss := &hunter.steadyShot
	hunter.steadyShotTemplate.Apply(ss)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ss.CastTime = hunter.SteadyShotCastTime()
	ss.Effect.Target = target

	ss.Init(sim)
	return ss
}

func (hunter *Hunter) SteadyShotCastTime() time.Duration {
	return time.Duration(float64(time.Millisecond*1500)/hunter.RangedSwingSpeed()) + hunter.latency
}
