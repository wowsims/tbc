package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func NewShaman(player *core.Player, agentID int, options map[string]string) *Shaman {
	agent := NewAdaptiveAgent(nil)

	// if WaterShield {
	// 	s[StatMP5] += 50
	// }

	// TODO: could we include Party as a constructor argument to add spec specific
	// buffs during construction instead of on every reset.

	// TODO: totem buffs

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

	// TODO: Should this be a special aura just for shaman? or should we add all the general aura
	//   OnXXX functions to agents (to allow them to apply non-aura effects for the class)
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
	s.agent.OnActionAccepted(s, sim, action)
}
func (s *Shaman) Reset(newsim *core.Simulation) {
	// Do we need to reset anything special?
}

// Agent is shaman specific agent for behavior.
type shamanAgent interface {
	// Returns the action this Agent would like to take next.
	ChooseAction(*Shaman, *core.Simulation) core.AgentAction

	// This will be invoked if the chosen action is actually executed, so the Agent can update its state.
	OnActionAccepted(*Shaman, *core.Simulation, core.AgentAction)

	// Returns this Agent to its initial state.
	Reset(*core.Simulation)
}

type Totems struct {
	TotemOfWrath int
	WrathOfAir   bool
	ManaStream   bool
	Cyclone2PC   bool // Cyclone set 2pc bonus
}

func (tt Totems) AddStats(s core.Stats) core.Stats {
	s[core.StatSpellCrit] += 66.24 * float64(tt.TotemOfWrath)
	s[core.StatSpellHit] += 37.8 * float64(tt.TotemOfWrath)
	if tt.WrathOfAir {
		s[core.StatSpellPower] += 101
		if tt.Cyclone2PC {
			s[core.StatSpellPower] += 20
		}
	}
	if tt.ManaStream {
		s[core.StatMP5] += 50
	}
	return s
}

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

// Mods for TLC spells that don't change within the sim.
// var hitMod = (-0.02 * float64(sim.Options.Talents.ElementalPrecision)) + (-0.01 * float64(sim.Options.Talents.NaturesGuidance))
// var critMod = (-0.01 * float64(sim.Options.Talents.TidalMastery)) + (-0.01 * float64(sim.Options.Talents.CallOfThunder))

// func (t Talents) AddStats(s Stats) Stats {
// TODO: make these an 'oncast' effect that specificaly checks the cast type.
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

func TryActivateBloodlust(sim *core.Simulation, party *core.Party, player *core.Player) {
	if player.IsOnCD(core.MagicIDBloodlust, sim.CurrentTime) {
		return
	}

	dur := time.Second * 40 // assumes that multiple BLs are different shaman.
	player.SetCD(core.MagicIDBloodlust, time.Minute*10)

	for _, p := range party.Players {
		p.AddAura(sim, core.Aura{
			ID:      core.MagicIDBloodlust,
			Expires: sim.CurrentTime + dur,
			OnCast: func(sim *core.Simulation, p *core.Player, c *core.Cast) {
				c.CastTime = (c.CastTime * 10) / 13 // 30% faster
			},
		})
	}
}
