package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var MultiShotCooldownID = core.NewCooldownID()
var MultiShotActionID = core.ActionID{SpellID: 27021, CooldownID: MultiShotCooldownID}

func (hunter *Hunter) registerMultiShotSpell(sim *core.Simulation) {
	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:  MultiShotActionID,
				Character: hunter.GetCharacter(),
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 275,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 275,
				},
				// Cast time is affected by ranged attack speed so set it later.
				//CastTime:     time.Millisecond * 500,
				GCD:         core.GCDDefault + hunter.latency,
				Cooldown:    time.Second * 10,
				IgnoreHaste: true, // Hunter GCD is locked at 1.5s
				SpellSchool: core.SpellSchoolPhysical,
				SpellExtras: core.SpellExtrasMeleeMetrics,
			},
		},
	}
	ama.Cost.Value *= 1 - 0.02*float64(hunter.Talents.Efficiency)
	if ItemSetDemonStalker.CharacterHasSetBonus(&hunter.Character, 4) {
		ama.Cost.Value -= 275.0 * 0.1
	}

	baseEffect := core.SpellEffect{
		ProcMask: core.ProcMaskRangedSpecial,

		BonusCritRating:  float64(hunter.Talents.ImprovedBarrage) * 4 * core.MeleeCritRatingPerCritChance,
		DamageMultiplier: 1 + 0.04*float64(hunter.Talents.Barrage),
		ThreatMultiplier: 1,

		BaseDamage: hunter.talonOfAlarDamageMod(core.BaseDamageConfig{
			Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
				return (hitEffect.RangedAttackPower(spell.Character)+hitEffect.RangedAttackPowerOnTarget())*0.2 +
					hunter.AutoAttacks.Ranged.BaseDamage(sim) +
					hunter.AmmoDamageBonus +
					hitEffect.BonusWeaponDamage(spell.Character) +
					205
			},
			TargetSpellCoefficient: 1,
		}),
		OutcomeApplier: core.OutcomeFuncRangedHitAndCrit(hunter.critMultiplier(true, sim.GetPrimaryTarget())),

		OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			hunter.rotation(sim, false)
		},
	}

	numHits := core.MinInt32(3, sim.GetNumTargets())
	effects := make([]core.SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}

	hunter.MultiShot = hunter.RegisterSpell(core.SpellConfig{
		Template: ama,
		ModifyCast: func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
			instance.CastTime = hunter.MultiShotCastTime()
		},
		ApplyEffects: core.ApplyEffectFuncDamageMultiple(effects),
	})
}

func (hunter *Hunter) MultiShotCastTime() time.Duration {
	return time.Duration(float64(time.Millisecond*500)/hunter.RangedSwingSpeed()) + hunter.latency
}
