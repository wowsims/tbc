package priest

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
)

const SpellIDStarshards int32 = 25446

var SSCooldownID = core.NewCooldownID()
var StarshardsActionID = core.ActionID{SpellID: SpellIDStarshards, CooldownID: SSCooldownID}

func (priest *Priest) registerStarshardsSpell(sim *core.Simulation) {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    StarshardsActionID,
				Character:   &priest.Character,
				SpellSchool: core.SpellSchoolArcane,
				CastTime:    0,
				GCD:         core.GCDDefault,
				Cooldown:    time.Second * 30,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			DamageMultiplier:    1,
			ThreatMultiplier:    1,
			DotInput: core.DotDamageInput{
				NumberOfTicks:  5,
				TickLength:     time.Second * 3,
				TickBaseDamage: core.DotSnapshotFuncMagic(785/5, 0.167),
			},
		},
	}
	template.Cost.Value -= template.BaseCost.Value * float64(priest.Talents.MentalAgility) * 0.02

	priest.Starshards = priest.RegisterSpell(core.SpellConfig{
		Template: template,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ThreatMultiplier: 1,
			OutcomeApplier:   core.OutcomeFuncMagicHit(),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					priest.StarshardsDot.Apply(sim)
				}
			},
		}),
	})

	target := sim.GetPrimaryTarget()
	priest.StarshardsDot = core.NewDot(core.Dot{
		Spell: priest.Starshards,
		Aura: target.RegisterAura(&core.Aura{
			Label:    "Starshards-" + strconv.Itoa(int(priest.Index)),
			ActionID: StarshardsActionID,
		}),

		NumberOfTicks: 5,
		TickLength:    time.Second * 3,

		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			IsPeriodic:       true,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(785/5, 0.167),
			OutcomeApplier:   core.OutcomeFuncTick(),
		}),
	})
}
