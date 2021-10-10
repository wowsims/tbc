package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func newShaman(character core.Character, talents Talents, selfBuffs SelfBuffs, agent shamanAgent) *Shaman {
	if selfBuffs.WaterShield {
		character.InitialStats[stats.MP5] += 50
	}

	return &Shaman{
		Character: character,
		agent:     agent,
		Talents:   talents,
		SelfBuffs: selfBuffs,

		convectionBonus: 0.02 * float64(talents.Convection),
		concussionBonus: 1 + 0.01*float64(talents.Concussion),
	}
}

// Which buffs this shaman is using.
type SelfBuffs struct {
	Bloodlust    bool
	WaterShield  bool
	TotemOfWrath bool
	WrathOfAir   bool
	ManaSpring   bool
}

// Agent is shaman specific agent for behavior.
type shamanAgent interface {
	// Returns the action this Agent would like to take next.
	ChooseAction(*Shaman, *core.Simulation) core.AgentAction

	// This will be invoked if the chosen action is actually executed, so the Agent can update its state.
	OnActionAccepted(*Shaman, *core.Simulation, core.AgentAction)

	// Returns this Agent to its initial state.
	Reset(*Shaman, *core.Simulation)
}

// Shaman represents a shaman character.
type Shaman struct {
	core.Character

	agent shamanAgent

	Talents   Talents
	SelfBuffs SelfBuffs

	// HACK HACK HACK
	// TODO: do we actually need a 'on start' method for agents?
	//   This particular use case could also be solved by the 'OnStatAdd' event...
	//    but are there other things we want to do once all buffs are applied right before starting?
	//   Unrelenting storm could also be calculated on the fly if we can allow agents to override the 'Advance' function.
	started bool

	// cache
	convectionBonus float64
	concussionBonus float64

	elementalFocusStacks byte
}

func (shaman *Shaman) GetCharacter() *core.Character {
	return &shaman.Character
}

func (shaman *Shaman) AddRaidBuffs(buffs *core.Buffs) {
}
func (shaman *Shaman) AddPartyBuffs(buffs *core.Buffs) {
	if shaman.SelfBuffs.Bloodlust {
		buffs.Bloodlust += 1
	}

	if shaman.SelfBuffs.TotemOfWrath {
		buffs.TotemOfWrath += 1
	}

	if shaman.SelfBuffs.ManaSpring {
		buffs.ManaSpringTotem = proto.TristateEffect_TristateEffectRegular
	}

	if shaman.SelfBuffs.WrathOfAir {
		// TODO: Check for t4 set bonus
		buffs.WrathOfAirTotem = proto.TristateEffect_TristateEffectRegular
	}
}

// BuffUp lets you buff up all characters in sim (and yourself)
func (shaman *Shaman) BuffUp(sim *core.Simulation) {
}
func (shaman *Shaman) ChooseAction(sim *core.Simulation) core.AgentAction {
	// TODO: Move this to BuffUp?
	if !shaman.started {
		shaman.started = true
		// we need to apply regen once all buffs are applied.
		shaman.Stats[stats.MP5] += shaman.Stats[stats.Intellect] * (0.02 * float64(shaman.Talents.UnrelentingStorm))
	}

	// Before casting, activate shaman powers!
	TryActivateBloodlust(sim, shaman)
	if shaman.Talents.ElementalMastery {
		TryActivateEleMastery(sim, shaman)
	}

	return shaman.agent.ChooseAction(shaman, sim)
}
func (shaman *Shaman) OnActionAccepted(sim *core.Simulation, action core.AgentAction) {
	shaman.agent.OnActionAccepted(shaman, sim, action)
}
func (shaman *Shaman) Reset(newsim *core.Simulation) {
	shaman.started = false
	shaman.agent.Reset(shaman, newsim)
}

type Talents struct {
	ElementalFocus     bool
	LightningMastery   int
	LightningOverload  int
	ElementalPrecision int
	NaturesGuidance    int
	TidalMastery       int
	ElementalMastery   bool
	ElementalFury      bool
	UnrelentingStorm   int
	CallOfThunder      int
	Convection         int
	Concussion         int
}

func convertShamTalents(t *proto.ShamanTalents) Talents {
	return Talents{
		LightningOverload:  int(t.LightningOverload),
		ElementalPrecision: int(t.ElementalPrecision),
		NaturesGuidance:    int(t.NaturesGuidance),
		TidalMastery:       int(t.TidalMastery),
		ElementalMastery:   t.ElementalMastery,
		ElementalFury:      t.ElementalFury,
		UnrelentingStorm:   int(t.UnrelentingStorm),
		CallOfThunder:      int(t.CallOfThunder),
		Convection:         int(t.Convection),
		Concussion:         int(t.Concussion),
		LightningMastery:   int(t.LightningMastery),
		ElementalFocus:     t.ElementalFocus,
	}
}

func TryActivateBloodlust(sim *core.Simulation, shaman *Shaman) {
	if shaman.IsOnCD(core.MagicIDBloodlust, sim.CurrentTime) {
		return
	}

	dur := time.Second * 40 // assumes that multiple BLs are different shaman.
	shaman.SetCD(core.MagicIDBloodlust, time.Minute*10+sim.CurrentTime)

	shaman.Party.AddAura(sim, core.Aura{
		ID:      core.MagicIDBloodlust,
		Expires: sim.CurrentTime + dur,
		OnCast: func(sim *core.Simulation, cast core.DirectCastAction, input *core.DirectCastInput) {
			input.CastTime = (input.CastTime * 10) / 13 // 30% faster
		},
	})
}

func TryActivateEleMastery(sim *core.Simulation, shaman *Shaman) {
	if shaman.IsOnCD(core.MagicIDEleMastery, sim.CurrentTime) {
		return
	}

	shaman.AddAura(sim, core.Aura{
		ID:      core.MagicIDEleMastery,
		Expires: core.NeverExpires,
		OnCast: func(sim *core.Simulation, cast core.DirectCastAction, input *core.DirectCastInput) {
			input.ManaCost = 0
			input.GuaranteedCrit = true
		},
		OnCastComplete: func(sim *core.Simulation, cast core.DirectCastAction) {
			// Remove the buff and put skill on CD
			shaman.SetCD(core.MagicIDEleMastery, time.Second*180+sim.CurrentTime)
			shaman.RemoveAura(sim, core.MagicIDEleMastery)
		},
	})
}
