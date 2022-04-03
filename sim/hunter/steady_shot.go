package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var SteadyShotActionID = core.ActionID{SpellID: 34120}

func (hunter *Hunter) registerSteadyShotSpell(sim *core.Simulation) {
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
				GCD:         core.GCDDefault + hunter.latency,
				IgnoreHaste: true, // Hunter GCD is locked at 1.5s
				SpellSchool: core.SpellSchoolPhysical,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryRanged,
			CritRollCategory:    core.CritRollCategoryPhysical,
			CritMultiplier:      hunter.critMultiplier(true, sim.GetPrimaryTarget()),
			ProcMask:            core.ProcMaskRangedSpecial,
			DamageMultiplier:    1,
			ThreatMultiplier:    1,
			BaseDamage: hunter.talonOfAlarDamageMod(core.BaseDamageConfig{
				Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
					return (hitEffect.RangedAttackPower(spell.Character)+hitEffect.RangedAttackPowerOnTarget())*0.2 +
						hunter.AutoAttacks.Ranged.BaseDamage(sim)*2.8/hunter.AutoAttacks.Ranged.SwingSpeed +
						150
				},
				TargetSpellCoefficient: 1,
			}),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				hunter.killCommandBlocked = false
				hunter.TryKillCommand(sim, spellEffect.Target)
				hunter.rotation(sim, false)
			},
		},
	}

	ama.Cost.Value *= 1 - 0.02*float64(hunter.Talents.Efficiency)

	if ItemSetRiftStalker.CharacterHasSetBonus(&hunter.Character, 4) {
		ama.Effect.BonusCritRating += 5 * core.MeleeCritRatingPerCritChance
	}
	if ItemSetGronnstalker.CharacterHasSetBonus(&hunter.Character, 4) {
		ama.Effect.DamageMultiplier *= 1.1
	}

	hunter.SteadyShot = hunter.RegisterSpell(core.SpellConfig{
		Template: ama,
		ModifyCast: func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
			instance.Effect.Target = target
			instance.CastTime = hunter.SteadyShotCastTime()
		},
	})
}

func (hunter *Hunter) SteadyShotCastTime() time.Duration {
	return time.Duration(float64(time.Millisecond*1500)/hunter.RangedSwingSpeed()) + hunter.latency
}
