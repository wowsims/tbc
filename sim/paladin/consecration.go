package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const ConsecrationCDTime = time.Second * 8

const SpellIDConsecrationRank6 = 27173
const SpellIDConsecrationRank4 = 20923
const SpellIDConsecrationRank1 = 26573

var ConsecrationRank6ActionID = core.ActionID{SpellID: SpellIDConsecrationRank6}
var ConsecrationRank4ActionID = core.ActionID{SpellID: SpellIDConsecrationRank4}
var ConsecrationRank1ActionID = core.ActionID{SpellID: SpellIDConsecrationRank1}

// Maybe could switch "rank" parameter type to some proto thing. Would require updates to proto files.
// Prot guys do whatever you want here I guess
func (paladin *Paladin) RegisterConsecrationSpell(sim *core.Simulation, rank int32) {
	var manaCost float64
	var actionID core.ActionID
	var baseDamage float64

	switch rank {
	case 6:
		manaCost = 660
		actionID = ConsecrationRank6ActionID
		baseDamage = 64
	case 4:
		manaCost = 390
		actionID = ConsecrationRank4ActionID
		baseDamage = 35
	case 1:
		manaCost = 120
		actionID = ConsecrationRank1ActionID
		baseDamage = 8
	default:
		manaCost = 0.0
	}

	// Check for bad input
	if manaCost == 0.0 {
		panic("Undefined Consecration rank specified.")
	}

	consecrationDot := core.NewDot(core.Dot{
		Aura: paladin.RegisterAura(core.Aura{
			Label:    "Consecration",
			ActionID: actionID,
		}),
		NumberOfTicks: 8,
		TickLength:    time.Second * 1,
		TickEffects: core.TickFuncAOESnapshot(sim, core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(baseDamage, 0.119),
			OutcomeApplier:   core.OutcomeFuncTick(),
			IsPeriodic:       true,
		}),
	})

	paladin.Consecration = paladin.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolHoly,

		ResourceType: stats.Mana,
		BaseCost:     manaCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: manaCost,
				GCD:  core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    paladin.NewTimer(),
				Duration: ConsecrationCDTime,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDot(consecrationDot),
	})

	consecrationDot.Spell = paladin.Consecration
}
