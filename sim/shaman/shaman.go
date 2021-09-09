package shaman

import (
	"github.com/wowsims/tbc/sim/core"
)

func NewShaman(player *core.Player, agentID int, options map[string]string) *Shaman {
	agent := NewAdaptiveAgent(nil)

	return &Shaman{
		agent:  agent,
		Player: player,
	}
}

// Shaman represents a shaman player.
type Shaman struct {
	agent        shamanAgent // Controller logic
	Talents      Talents     // Shaman Talents
	*core.Player             // State of player
}

// BuffUp lets you buff up all players in sim (and yourself)
func (s *Shaman) BuffUp(sim *core.Simulation, party *core.Party) {
	// if sim.Options.Talents.Concussion > 0 {
	// 	bonusdmg := (0.01 * sim.Options.Talents.Concussion)
	// }
	if s.Talents.LightningOverload > 0 {
		s.AddAura(sim, AuraLightningOverload(s.Talents.LightningOverload))
	}
}
func (s *Shaman) ChooseAction(sim *core.Simulation) core.AgentAction {
	return s.agent.ChooseAction(s, sim)
}
func (s *Shaman) OnActionAccepted(sim *core.Simulation, action core.AgentAction) {

}
func (s *Shaman) Reset(newsim *core.Simulation) {

}
func (s *Shaman) Cast() {

}

// Agent is shaman specific agent for behavior.
type shamanAgent interface {
	// Returns the action this Agent would like to take next.
	ChooseAction(*Shaman, *core.Simulation) core.AgentAction

	// This will be invoked if the chosen action is actually executed, so the Agent can update its state.
	OnActionAccepted(*core.Simulation, core.AgentAction)

	// Returns this Agent to its initial state.
	Reset(*core.Simulation)
}

type Totems struct {
	TotemOfWrath int
	WrathOfAir   bool
	ManaStream   bool
	Cyclone2PC   bool // Cyclone set 2pc bonus
}

// func (tt Totems) AddStats(s Stats) Stats {
// 	s[StatSpellCrit] += 66.24 * float64(tt.TotemOfWrath)
// 	s[StatSpellHit] += 37.8 * float64(tt.TotemOfWrath)
// 	if tt.WrathOfAir {
// 		s[StatSpellDmg] += 101
// 		if tt.Cyclone2PC {
// 			s[StatSpellDmg] += 20
// 		}
// 	}
// 	if tt.ManaStream {
// 		s[StatMP5] += 50
// 	}
// 	return s
// }

type Talents struct {
	LightningOverload  int
	ElementalPrecision int
	NaturesGuidance    int
	TidalMastery       int
	ElementalMastery   bool
	UnrelentingStorm   int
	CallOfThunder      int
	Convection         int

	Concussion float64 // temp hack to speed up not converting this to a int on every spell cast
}

// func (t Talents) AddStats(s Stats) Stats {
// 	s[StatSpellHit] += 25.2 * float64(t.ElementalPrecision)
// 	s[StatSpellHit] += 12.6 * float64(t.NaturesGuidance)
// 	s[StatSpellCrit] += 22.08 * float64(t.TidalMastery)
// 	s[StatSpellCrit] += 22.08 * float64(t.CallOfThunder)

// 	return s
// }

func ActivateChainLightningBounce() core.Aura {
	return core.Aura{
		ID:      core.MagicIDClBounce,
		Expires: core.NeverExpires,
		OnCastComplete: func(sim *core.Simulation, p *core.Player, c *core.Cast) {
			if c.Spell.ID != core.MagicIDCL6 || c.IsClBounce {
				return
			}

			dmgCoeff := 1.0
			if c.IsLO {
				dmgCoeff = 0.5
			}
			for i := 1; i < sim.Options.Encounter.NumTargets; i++ {
				if p.HasAura(core.MagicIDTidefury) {
					dmgCoeff *= 0.83
				} else {
					dmgCoeff *= 0.7
				}
				clone := &core.Cast{
					IsLO:       c.IsLO,
					IsClBounce: true,
					Spell:      c.Spell,
					Crit:       c.Crit,
					CritBonus:  c.CritBonus,
					Effect:     func(sim *core.Simulation, p *core.Player, c *core.Cast) { c.DidDmg *= dmgCoeff },
				}
				sim.Cast(p, clone)
			}
		},
	}
}
