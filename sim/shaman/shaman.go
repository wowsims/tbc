package shaman

import (
	"time"

	"github.com/wowsims/tbc/items"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

type AgentType int

const (
	AgentTypeUnknown       = 0
	AgentTypeFixedLBCL     = 1
	AgentTypeCLOnClearcast = 2
	AgentTypeAdaptive      = 3
	AgentTypeCLOnCD        = 4
)

func NewShaman(player *core.Player, party *core.Party, talents Talents, totems Totems, agentID AgentType, agentOptions map[string]int) *Shaman {
	var agent shamanAgent

	switch agentID {
	case AgentTypeAdaptive:
		agent = NewAdaptiveAgent()
	case AgentTypeCLOnClearcast:
		agent = NewCLOnClearcastAgent()
	case AgentTypeFixedLBCL:
		numLB := agentOptions["numLBtoCL"]
		if numLB == -1 {
			agent = NewLBOnlyAgent()
		} else {
			agent = NewFixedRotationAgent(numLB)
		}
	case AgentTypeCLOnCD:
		agent = NewCLOnCDAgent()
	}

	// if WaterShield {
	player.InitialStats[stats.MP5] += 50
	// }

	party.AddInitialStats(totems.Stats())

	return &Shaman{
		agent:   agent,
		Player:  player,
		Talents: talents,

		convectionBonus: 0.02 * float64(talents.Convection),
		concussionBonus: 1 + 0.01*float64(talents.Concussion),
	}
}

// Shaman represents a shaman player.
type Shaman struct {
	agent        shamanAgent // Controller logic
	Talents      Talents     // Shaman Talents
	*core.Player             // State of player

	// HACK HACK HACK
	// TODO: do we actually need a 'on start' method for agents?
	//   This particular use case could also be solved by the 'OnStatAdd' event...
	//    but are there other things we want to do once all buffs are applied right before starting?
	//   Unrelenting storm could also be calculated on the fly if we can allow agents to override the 'Advance' function.
	started bool

	// cache
	convectionBonus float64
	concussionBonus float64
}

// BuffUp lets you buff up all players in sim (and yourself)
func (s *Shaman) BuffUp(sim *core.Simulation, party *core.Party) {
	if s.Talents.LightningOverload > 0 {
		s.AddAura(sim, AuraLightningOverload(s.Talents.LightningOverload))
	}
}
func (s *Shaman) ChooseAction(sim *core.Simulation, party *core.Party) core.AgentAction {
	if !s.started {
		s.started = true
		// we need to apply regen once all buffs are applied.
		s.Stats[stats.MP5] += s.Stats[stats.Intellect] * (0.02 * float64(s.Talents.UnrelentingStorm))
	}
	// Before casting, activate shaman powers!
	TryActivateBloodlust(sim, party, s.Player)
	if s.Talents.ElementalMastery {
		TryActivateEleMastery(sim, s.Player)
	}

	return s.agent.ChooseAction(s, party, sim)
}
func (s *Shaman) OnActionAccepted(sim *core.Simulation, action core.AgentAction) {
	s.agent.OnActionAccepted(s, sim, action)
}
func (s *Shaman) Reset(newsim *core.Simulation) {
	s.started = false
	s.agent.Reset(newsim)
}
func (s *Shaman) OnSpellHit(sim *core.Simulation, _ core.PlayerAgent, cast *core.Cast) {
	if cast.Spell.ID == core.MagicIDTLCLB { // TLC does not benefit from shaman talents
		return
	}
	cast.DidDmg *= s.concussionBonus // add concussion

	if cast.DidCrit && s.Talents.ElementalFocus {
		a := core.Aura{
			ID:             core.MagicIDEleFocus,
			Expires:        sim.CurrentTime + time.Second*15,
			Stacks:         2,
			OnCast:         elementalFocusOnCast,
			OnCastComplete: elementalFocusOnCastComplete,
		}
		s.AddAura(sim, a)
	}
}

func elementalFocusOnCast(sim *core.Simulation, player core.PlayerAgent, c *core.Cast) {
	c.ManaCost *= .6 // reduced by 40%
}

func elementalFocusOnCastComplete(sim *core.Simulation, player core.PlayerAgent, c *core.Cast) {
	if c.ManaCost <= 0 {
		return // Don't consume charges from free spells.
	}

	player.Auras[core.MagicIDEleFocus].Stacks--
	if player.Auras[core.MagicIDEleFocus].Stacks == 0 {
		player.RemoveAura(sim, player, core.MagicIDEleFocus)
	}
}

// Agent is shaman specific agent for behavior.
type shamanAgent interface {
	// Returns the action this Agent would like to take next.
	ChooseAction(*Shaman, *core.Party, *core.Simulation) core.AgentAction

	// This will be invoked if the chosen action is actually executed, so the Agent can update its state.
	OnActionAccepted(*Shaman, *core.Simulation, core.AgentAction)

	// Returns this Agent to its initial state.
	Reset(*core.Simulation)
}

type Totems struct {
	TotemOfWrath int
	WrathOfAir   bool
	ManaStream   bool
}

func (tt Totems) Stats() stats.Stats {
	s := stats.Stats{
		stats.SpellCrit:  66.24 * float64(tt.TotemOfWrath),
		stats.SpellHit:   37.8 * float64(tt.TotemOfWrath),
		stats.SpellPower: 0,
		stats.MP5:        0,
	}
	if tt.WrathOfAir {
		s[stats.SpellPower] += 101
	}
	if tt.ManaStream {
		s[stats.MP5] += 50
	}
	return s
}

type Talents struct {
	ElementalFocus     bool
	LightningMastery   int
	LightningOverload  int
	ElementalPrecision int
	NaturesGuidance    int
	TidalMastery       int
	ElementalMastery   bool
	UnrelentingStorm   int
	CallOfThunder      int
	Convection         int
	Concussion         int
}

func TryActivateBloodlust(sim *core.Simulation, party *core.Party, player *core.Player) {
	if player.IsOnCD(core.MagicIDBloodlust, sim.CurrentTime) {
		return
	}

	dur := time.Second * 40 // assumes that multiple BLs are different shaman.
	player.SetCD(core.MagicIDBloodlust, time.Minute*10+sim.CurrentTime)

	party.AddAura(sim, core.Aura{
		ID:      core.MagicIDBloodlust,
		Expires: sim.CurrentTime + dur,
		OnCast: func(sim *core.Simulation, p core.PlayerAgent, c *core.Cast) {
			c.CastTime = (c.CastTime * 10) / 13 // 30% faster
		},
	})
}

// FUTURE: We can cache like 75% of the calculation for a spell cast ahead of time.
//   First time we cast we should create and cache this cast object for better performance.
//   This would get rid of the individual cached floats on Shaman.

// func createBaseCast(player *Shaman, sim *core.Simulation, sp *core.Spell) *core.Cast {
// 	cast := core.NewCast(sim, sp)

// 	if player.Talents.ElementalPrecision > 0 {
// 		// FUTURE: This only impacts "frost fire and nature" spells.
// 		//  We know it doesnt impact TLC.
// 		//  Are there any other spells that a shaman can cast?
// 		cast.BonusHit += float64(player.Talents.ElementalPrecision) * 0.02
// 	}
// 	if player.Talents.NaturesGuidance > 0 {
// 		cast.BonusHit += float64(player.Talents.NaturesGuidance) * 0.01
// 	}
// 	if player.Talents.TidalMastery > 0 {
// 		cast.BonusCrit += float64(player.Talents.TidalMastery) * 0.01
// 	}

// 	// TODO: Should we change these to be full auras?
// 	//   Doesnt seem needed since they can only be used by shaman right here.
// 	if player.Equip[items.ItemSlotRanged].ID == 28248 {
// 		cast.BonusSpellPower += 55
// 	} else if player.Equip[items.ItemSlotRanged].ID == 23199 {
// 		cast.BonusSpellPower += 33
// 	} else if player.Equip[items.ItemSlotRanged].ID == 32330 {
// 		cast.BonusSpellPower += 85
// 	}
// 	if player.Talents.CallOfThunder > 0 { // only applies to CL and LB
// 		cast.BonusCrit += float64(player.Talents.CallOfThunder) * 0.01
// 	}
// 	if sim.Options.Encounter.NumTargets > 1 {
// 		cast.DoItNow = ChainCast
// 	}
// 	cast.ManaCost *= player.convectionBonus

// 	if player.Talents.LightningMastery > 0 {
// 		cast.CastTime -= time.Millisecond * 100 * time.Duration(player.Talents.LightningMastery)
// 	}

// 	return cast
// }

// Totem Item IDs
const (
	TotemOfTheVoid           = 28248
	TotemOfStorms            = 23199
	TotemOfAncestralGuidance = 32330
)

// NewCastAction is how a shaman creates a new spell
//  TODO: Decide if we need separate functions for elemental and enhancement?
func NewCastAction(sim *core.Simulation, player *Shaman, sp *core.Spell) core.AgentAction {
	cast := core.NewCast(sim, sp)

	itsElectric := sp.ID == core.MagicIDCL6 || sp.ID == core.MagicIDLB12

	if player.Talents.ElementalPrecision > 0 {
		// FUTURE: This only impacts "frost fire and nature" spells.
		//  We know it doesnt impact TLC.
		//  Are there any other spells that a shaman can cast?
		cast.BonusHit += float64(player.Talents.ElementalPrecision) * 0.02
	}
	if player.Talents.NaturesGuidance > 0 {
		cast.BonusHit += float64(player.Talents.NaturesGuidance) * 0.01
	}
	if player.Talents.TidalMastery > 0 {
		cast.BonusCrit += float64(player.Talents.TidalMastery) * 0.01
	}

	if itsElectric {
		// TODO: Should we change these to be full auras?
		//   Doesnt seem needed since they can only be used by shaman right here.
		if player.Equip[items.ItemSlotRanged].ID == TotemOfTheVoid {
			cast.BonusSpellPower += 55
		} else if player.Equip[items.ItemSlotRanged].ID == TotemOfStorms {
			cast.BonusSpellPower += 33
		} else if player.Equip[items.ItemSlotRanged].ID == TotemOfAncestralGuidance {
			cast.BonusSpellPower += 85
		}
		if player.Talents.CallOfThunder > 0 { // only applies to CL and LB
			cast.BonusCrit += float64(player.Talents.CallOfThunder) * 0.01
		}
		if sp.ID == core.MagicIDCL6 && sim.Options.Encounter.NumTargets > 1 {
			cast.DoItNow = ChainCast
		}
		if player.Talents.LightningMastery > 0 {
			cast.CastTime -= time.Millisecond * 100 * time.Duration(player.Talents.LightningMastery)
		}
	}
	cast.CastTime = time.Duration(float64(cast.CastTime) / player.HasteBonus())

	// Apply any on cast effects.
	for _, id := range player.ActiveAuraIDs {
		if player.Auras[id].OnCast != nil {
			player.Auras[id].OnCast(sim, core.PlayerAgent{Agent: player, Player: player.Player}, cast)
		}
	}
	if itsElectric { // TODO: Add ElementalFury talent
		// This is written this way to deal with making CSD dmg increase correct.
		// The 'OnCast' auras include CSD
		cast.CritDamageMultipier *= 2 // This handles the 'Elemental Fury' talent which increases the crit bonus.
		cast.CritDamageMultipier -= 1 // reduce to multiplier instead of percent.

		// Convection applies against the base cost of the spell.
		cast.ManaCost -= sp.Mana * player.convectionBonus
	}

	return core.AgentAction{
		Wait: 0,
		Cast: cast,
	}
}
