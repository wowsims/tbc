package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ThunderClapCooldownID = core.NewCooldownID()
var ThunderClapActionID = core.ActionID{SpellID: 25264, CooldownID: ThunderClapCooldownID}

func (warrior *Warrior) registerThunderClapSpell(sim *core.Simulation) {
	cost := 20.0 - float64(warrior.Talents.FocusedRage)
	impTCDamageMult := 1.0
	if warrior.Talents.ImprovedThunderClap == 1 {
		cost -= 1
		impTCDamageMult = 1.4
	} else if warrior.Talents.ImprovedThunderClap == 2 {
		cost -= 2
		impTCDamageMult = 1.7
	} else if warrior.Talents.ImprovedThunderClap == 3 {
		cost -= 4
		impTCDamageMult = 2
	}

	baseEffect := core.SpellEffect{
		DamageMultiplier: impTCDamageMult,
		ThreatMultiplier: 1.75,
		BaseDamage:       core.BaseDamageConfigFlat(123),
		OutcomeApplier:   core.OutcomeFuncMagicHitAndCrit(warrior.spellCritMultiplier(true)),
	}

	numHits := core.MinInt32(4, sim.GetNumTargets())
	effects := make([]core.SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)

		tcAura := core.ThunderClapAura(effects[i].Target, warrior.Talents.ImprovedThunderClap)
		if i == 0 {
			warrior.ThunderClapAura = tcAura
		}

		effects[i].OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Landed() {
				tcAura.Activate(sim)
			}
		}
	}

	warrior.ThunderClap = warrior.RegisterSpell(core.SpellConfig{
		ActionID:    ThunderClapActionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasBinary,

		ResourceType: stats.Rage,
		BaseCost:     cost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: cost,
				GCD:  core.GCDDefault,
			},
			IgnoreHaste: true,
			Cooldown:    time.Second * 4,
		},

		ApplyEffects: core.ApplyEffectFuncDamageMultiple(effects),
	})
}

func (warrior *Warrior) CanThunderClap(sim *core.Simulation) bool {
	return warrior.StanceMatches(BattleStance|DefensiveStance) && warrior.CurrentRage() >= warrior.ThunderClap.DefaultCast.Cost && !warrior.IsOnCD(ThunderClapCooldownID, sim.CurrentTime)
}
