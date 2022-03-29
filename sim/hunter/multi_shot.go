package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var MultiShotCooldownID = core.NewCooldownID()
var MultiShotActionID = core.ActionID{SpellID: 27021, CooldownID: MultiShotCooldownID}

func (hunter *Hunter) newMultiShotTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
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
				GCD:                 core.GCDDefault + hunter.latency,
				Cooldown:            time.Second * 10,
				IgnoreHaste:         true, // Hunter GCD is locked at 1.5s
				OutcomeRollCategory: core.OutcomeRollCategoryRanged,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				// TODO: If we ever allow multiple targets to have their own type, need to
				// update this.
				CritMultiplier: hunter.critMultiplier(true, sim.GetPrimaryTarget()),
			},
		},
	}

	ama.Cost.Value *= 1 - 0.02*float64(hunter.Talents.Efficiency)
	if ItemSetDemonStalker.CharacterHasSetBonus(&hunter.Character, 4) {
		ama.Cost.Value -= 275.0 * 0.1
	}

	baseEffect := core.SpellEffect{
		ProcMask:         core.ProcMaskRangedSpecial,
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BaseDamage: hunter.talonOfAlarDamageMod(core.BaseDamageConfig{
			Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spellCast *core.SpellCast) float64 {
				return (hitEffect.RangedAttackPower(spellCast)+hitEffect.RangedAttackPowerOnTarget())*0.2 +
					hunter.AutoAttacks.Ranged.BaseDamage(sim) +
					hunter.AmmoDamageBonus +
					hitEffect.BonusWeaponDamage(spellCast) +
					205
			},
			TargetSpellCoefficient: 1,
		}),
		OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			hunter.rotation(sim, false)
		},
	}

	baseEffect.DamageMultiplier *= 1 + 0.04*float64(hunter.Talents.Barrage)
	baseEffect.BonusCritRating += float64(hunter.Talents.ImprovedBarrage) * 4 * core.MeleeCritRatingPerCritChance

	numHits := core.MinInt32(3, sim.GetNumTargets())
	effects := make([]core.SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	ama.Effects = effects

	return core.NewSimpleSpellTemplate(ama)
}

func (hunter *Hunter) NewMultiShot(sim *core.Simulation) *core.SimpleSpell {
	ms := &hunter.multiShot
	hunter.multiShotTemplate.Apply(ms)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ms.CastTime = hunter.MultiShotCastTime()

	ms.Init(sim)
	return ms
}

func (hunter *Hunter) MultiShotCastTime() time.Duration {
	return time.Duration(float64(time.Millisecond*500)/hunter.RangedSwingSpeed()) + hunter.latency
}
