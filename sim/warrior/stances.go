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

func (warrior *Warrior) makeCastStance(sim *core.Simulation, stance Stance, aura *core.Aura) func(sim *core.Simulation) {
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

		// Add new stance aura.
		aura.Activate(sim)
		warrior.Stance = stance
	}
}

func (warrior *Warrior) registerBattleStanceAura() {
	actionID := core.ActionID{SpellID: 2457}
	threatMult := 0.8

	warrior.BattleStanceAura = warrior.GetOrRegisterAura(core.Aura{
		Label:    "Battle Stance",
		Tag:      "Stance",
		Priority: 1,
		ActionID: actionID,
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.ThreatMultiplier *= threatMult
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.ThreatMultiplier /= threatMult
		},
	})
}

func (warrior *Warrior) registerDefensiveStanceAura() {
	actionID := core.ActionID{SpellID: 71}
	threatMult := 1.3 * (1 + 0.05*float64(warrior.Talents.Defiance))

	warrior.DefensiveStanceAura = warrior.GetOrRegisterAura(core.Aura{
		Label:    "Defensive Stance",
		Tag:      "Stance",
		Priority: 1,
		ActionID: actionID,
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.ThreatMultiplier *= threatMult
			aura.Unit.PseudoStats.DamageDealtMultiplier *= 0.9
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.ThreatMultiplier /= threatMult
			aura.Unit.PseudoStats.DamageDealtMultiplier /= 0.9
		},
	})
}

func (warrior *Warrior) registerBerserkerStanceAura() {
	threatMult := 0.8 - 0.02*float64(warrior.Talents.ImprovedBerserkerStance)
	critBonus := core.MeleeCritRatingPerCritChance * 3

	warrior.BerserkerStanceAura = warrior.GetOrRegisterAura(core.Aura{
		Label:    "Berserker Stance",
		Tag:      "Stance",
		Priority: 1,
		ActionID: core.ActionID{SpellID: 2458},
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.ThreatMultiplier *= threatMult
			aura.Unit.AddStat(stats.MeleeCrit, critBonus)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.ThreatMultiplier /= threatMult
			aura.Unit.AddStat(stats.MeleeCrit, -critBonus)
		},
	})
}
