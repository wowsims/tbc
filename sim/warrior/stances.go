package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Stance uint8

const (
	BattleStance Stance = 1 << iota
	DefensiveStance
	BerserkerStance
)

func (warrior *Warrior) StanceMatches(other Stance) bool {
	return (warrior.Stance & other) != 0
}

var StanceCooldownID = core.NewCooldownID()
var StanceCooldown = time.Second * 1

var BattleStanceAuraID = core.NewAuraID()
var DefensiveStanceAuraID = core.NewAuraID()
var BerserkerStanceAuraID = core.NewAuraID()

func (warrior *Warrior) makeCastStance(_ *core.Simulation, stance Stance, aura core.Aura) func(sim *core.Simulation) {
	maxRetainedRage := 10.0 + 5*float64(warrior.Talents.TacticalMastery)

	return func(sim *core.Simulation) {
		if warrior.Stance == stance {
			panic("Already in stance " + string(stance))
		}
		if warrior.IsOnCD(StanceCooldownID, sim.CurrentTime) {
			panic("Stance on CD")
		}

		warrior.SetCD(StanceCooldownID, sim.CurrentTime+StanceCooldown)
		warrior.Metrics.AddInstantCast(aura.ActionID)
		if warrior.CurrentRage() > maxRetainedRage {
			warrior.SpendRage(sim, warrior.CurrentRage()-maxRetainedRage, aura.ActionID)
		}

		// Remove old stance aura.
		if warrior.Stance == BattleStance {
			warrior.RemoveAura(sim, BattleStanceAuraID)
		} else if warrior.Stance == DefensiveStance {
			warrior.RemoveAura(sim, DefensiveStanceAuraID)
		} else {
			warrior.RemoveAura(sim, BerserkerStanceAuraID)
		}

		// Add new stance aura.
		warrior.AddAura(sim, aura)
		warrior.Stance = stance
	}
}

func (warrior *Warrior) BattleStanceAura() core.Aura {
	actionID := core.ActionID{SpellID: 2457}
	threatMult := 0.8

	return core.Aura{
		ID:       BattleStanceAuraID,
		ActionID: actionID,
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			warrior.PseudoStats.ThreatMultiplier *= threatMult
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			warrior.PseudoStats.ThreatMultiplier /= threatMult
		},
	}
}

func (warrior *Warrior) DefensiveStanceAura() core.Aura {
	actionID := core.ActionID{SpellID: 71}
	threatMult := 1.3 * (1 + 0.05*float64(warrior.Talents.Defiance))

	return core.Aura{
		ID:       DefensiveStanceAuraID,
		ActionID: actionID,
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			warrior.PseudoStats.ThreatMultiplier *= threatMult
			warrior.PseudoStats.DamageDealtMultiplier *= 0.9
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			warrior.PseudoStats.ThreatMultiplier /= threatMult
			warrior.PseudoStats.DamageDealtMultiplier /= 0.9
		},
	}
}

func (warrior *Warrior) BerserkerStanceAura() core.Aura {
	threatMult := 0.8 - 0.02*float64(warrior.Talents.ImprovedBerserkerStance)
	critBonus := core.MeleeCritRatingPerCritChance * 3

	return core.Aura{
		ID:       BerserkerStanceAuraID,
		ActionID: core.ActionID{SpellID: 2458},
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			warrior.PseudoStats.ThreatMultiplier *= threatMult
			warrior.AddStat(stats.MeleeCrit, critBonus)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			warrior.PseudoStats.ThreatMultiplier /= threatMult
			warrior.AddStat(stats.MeleeCrit, -critBonus)
		},
	}
}
