package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (paladin *Paladin) registerConsecrationSpell(sim *core.Simulation) {
	cdTimer := paladin.NewTimer()
	paladin.registerConsecrationSpellRank6(sim, cdTimer)
	paladin.registerConsecrationSpellRank4(sim, cdTimer)
	paladin.registerConsecrationSpellRank1(sim, cdTimer)

}

const ConsecrationCDTime = time.Second * 8
const SpellIDConsecrationRank6 int32 = 27173

var ConsecrationRank6ActionID = core.ActionID{SpellID: SpellIDConsecrationRank6}

func (paladin *Paladin) registerConsecrationSpellRank6(sim *core.Simulation, cdTimer *core.Timer) {
	baseCost := 660.0

	consecrationDot := core.NewDot(core.Dot{
		Aura: paladin.RegisterAura(core.Aura{
			Label:    "ConsecrationRank6",
			ActionID: ConsecrationRank6ActionID,
		}),
		NumberOfTicks: 8,
		TickLength:    time.Second * 1,
		TickEffects: core.TickFuncAOESnapshot(sim, core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(64, 0.119),
			OutcomeApplier:   core.OutcomeFuncTick(),
			IsPeriodic:       true,
		}),
	})

	paladin.ConsecrationRank6 = paladin.RegisterSpell(core.SpellConfig{
		ActionID:    ConsecrationRank6ActionID,
		SpellSchool: core.SpellSchoolHoly,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    cdTimer,
				Duration: ConsecrationCDTime,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDot(consecrationDot),
	})
	consecrationDot.Spell = paladin.ConsecrationRank6
}

const SpellIDConsecrationRank4 int32 = 20923

var ConsecrationRank4ActionID = core.ActionID{SpellID: SpellIDConsecrationRank4}

func (paladin *Paladin) registerConsecrationSpellRank4(sim *core.Simulation, cdTimer *core.Timer) {
	baseCost := 390.0

	consecrationDot := core.NewDot(core.Dot{
		Aura: paladin.RegisterAura(core.Aura{
			Label:    "ConsecrationRank4",
			ActionID: ConsecrationRank4ActionID,
		}),
		NumberOfTicks: 8,
		TickLength:    time.Second * 1,
		TickEffects: core.TickFuncAOESnapshot(sim, core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(35, 0.119),
			OutcomeApplier:   core.OutcomeFuncTick(),
			IsPeriodic:       true,
		}),
	})

	paladin.ConsecrationRank4 = paladin.RegisterSpell(core.SpellConfig{
		ActionID:    ConsecrationRank4ActionID,
		SpellSchool: core.SpellSchoolHoly,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    cdTimer,
				Duration: ConsecrationCDTime,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDot(consecrationDot),
	})
	consecrationDot.Spell = paladin.ConsecrationRank4
}

const SpellIDConsecrationRank1 int32 = 26573

var ConsecrationRank1ActionID = core.ActionID{SpellID: SpellIDConsecrationRank1}

func (paladin *Paladin) registerConsecrationSpellRank1(sim *core.Simulation, cdTimer *core.Timer) {
	baseCost := 120.0

	consecrationDot := core.NewDot(core.Dot{
		Aura: paladin.RegisterAura(core.Aura{
			Label:    "ConsecrationRank1",
			ActionID: ConsecrationRank1ActionID,
		}),
		NumberOfTicks: 8,
		TickLength:    time.Second * 1,
		TickEffects: core.TickFuncAOESnapshot(sim, core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(8, 0.119),
			OutcomeApplier:   core.OutcomeFuncTick(),
			IsPeriodic:       true,
		}),
	})

	paladin.ConsecrationRank1 = paladin.RegisterSpell(core.SpellConfig{
		ActionID:    ConsecrationRank1ActionID,
		SpellSchool: core.SpellSchoolHoly,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    cdTimer,
				Duration: ConsecrationCDTime,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDot(consecrationDot),
	})
	consecrationDot.Spell = paladin.ConsecrationRank1
}

func (paladin *Paladin) CastConsecrationRank(sim *core.Simulation, target *core.Target, rank proto.RetributionPaladin_Rotation_ConsecrationRank) {
	switch rank {
	case proto.RetributionPaladin_Rotation_Rank1:
		paladin.ConsecrationRank1.Cast(sim, target)
	case proto.RetributionPaladin_Rotation_Rank4:
		paladin.ConsecrationRank4.Cast(sim, target)
	case proto.RetributionPaladin_Rotation_Rank6:
		paladin.ConsecrationRank6.Cast(sim, target)
	}
}
