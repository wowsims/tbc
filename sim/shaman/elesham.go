package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func RegisterElementalShaman() {
	core.RegisterAgentFactory(proto.PlayerOptions_ElementalShaman{}, func(sim *core.Simulation, character core.Character, options *proto.PlayerOptions) core.Agent {
		return NewElementalShaman(sim, character, options)
	})
}

func NewElementalShaman(sim *core.Simulation, character core.Character, options *proto.PlayerOptions) *Shaman {
	eleShamOptions := options.GetElementalShaman()
	talents := convertShamTalents(eleShamOptions.Talents)

	selfBuffs := SelfBuffs{
		Bloodlust:    eleShamOptions.Options.Bloodlust,
		ManaSpring:   eleShamOptions.Options.ManaSpringTotem,
		TotemOfWrath: eleShamOptions.Options.TotemOfWrath,
		WrathOfAir:   eleShamOptions.Options.WrathOfAirTotem,
		WaterShield:  eleShamOptions.Options.WaterShield,
	}

	var agent shamanAgent

	switch eleShamOptions.Agent.Type {
	case proto.ElementalShaman_Agent_Adaptive:
		agent = NewAdaptiveAgent(sim)
	case proto.ElementalShaman_Agent_CLOnClearcast:
		agent = NewCLOnClearcastAgent(sim)
	case proto.ElementalShaman_Agent_FixedLBCL:
		agent = NewLBOnlyAgent(sim)
		// TODO: Add option for this
		//numLB := agentOptions["numLBtoCL"]
		//if numLB == -1 {
		//	agent = NewLBOnlyAgent()
		//} else {
		//	agent = NewFixedRotationAgent(numLB)
		//}
	case proto.ElementalShaman_Agent_CLOnCD:
		agent = NewCLOnCDAgent(sim)
	}

	return newShaman(character, talents, selfBuffs, agent)
}

// ################################################################
//                              LB ONLY
// ################################################################
type LBOnlyAgent struct {
}

func (agent *LBOnlyAgent) ChooseAction(shaman *Shaman, sim *core.Simulation) core.AgentAction {
	return NewLightningBolt(sim, shaman, false)
}

func (agent *LBOnlyAgent) OnActionAccepted(shaman *Shaman, sim *core.Simulation, action core.AgentAction) {
}
func (agent *LBOnlyAgent) Reset(shaman *Shaman, sim *core.Simulation) {}

func NewLBOnlyAgent(sim *core.Simulation) *LBOnlyAgent {
	return &LBOnlyAgent{}
}

// ################################################################
//                             CL ON CD
// ################################################################
type CLOnCDAgent struct {
}

func (agent *CLOnCDAgent) ChooseAction(shaman *Shaman, sim *core.Simulation) core.AgentAction {
	if shaman.IsOnCD(core.MagicIDChainLightning6, sim.CurrentTime) {
		return NewLightningBolt(sim, shaman, false)
	} else {
		return NewChainLightning(sim, shaman, false)
	}
}

func (agent *CLOnCDAgent) OnActionAccepted(shaman *Shaman, sim *core.Simulation, action core.AgentAction) {
}
func (agent *CLOnCDAgent) Reset(shaman *Shaman, sim *core.Simulation) {}

func NewCLOnCDAgent(sim *core.Simulation) *CLOnCDAgent {
	return &CLOnCDAgent{}
}

// ################################################################
//                          FIXED ROTATION
// ################################################################
type FixedRotationAgent struct {
	numLBsPerCL       int
	numLBsSinceLastCL int
}

// Returns if any temporary haste buff is currently active.
// TODO: Figure out a way to make this automatic
func (agent *FixedRotationAgent) temporaryHasteActive(shaman *Shaman) bool {
	return shaman.HasAura(core.MagicIDBloodlust) ||
		shaman.HasAura(core.MagicIDDrums) ||
		shaman.HasAura(core.MagicIDTrollBerserking) ||
		shaman.HasAura(core.MagicIDSkullGuldan) ||
		shaman.HasAura(core.MagicIDFungalFrenzy)
}

func (agent *FixedRotationAgent) ChooseAction(shaman *Shaman, sim *core.Simulation) core.AgentAction {
	if agent.numLBsSinceLastCL < agent.numLBsPerCL {
		return NewLightningBolt(sim, shaman, false)
	}

	if !shaman.IsOnCD(core.MagicIDChainLightning6, sim.CurrentTime) {
		return NewChainLightning(sim, shaman, false)
	}

	// If we have a temporary haste effect (like bloodlust or quags eye) then
	// we should add LB casts instead of waiting
	if agent.temporaryHasteActive(shaman) {
		return NewLightningBolt(sim, shaman, false)
	}

	return core.NewWaitAction(sim, shaman, shaman.GetRemainingCD(core.MagicIDChainLightning6, sim.CurrentTime))
}

func (agent *FixedRotationAgent) OnActionAccepted(shaman *Shaman, sim *core.Simulation, action core.AgentAction) {
	cast, isCastAction := action.(*core.DirectCastAction)
	if !isCastAction {
		return
	}

	if cast.GetActionID().SpellID == SpellIDLB12 {
		agent.numLBsSinceLastCL++
	} else if cast.GetActionID().SpellID == SpellIDCL6 {
		agent.numLBsSinceLastCL = 0
	}
}

func (agent *FixedRotationAgent) Reset(shaman *Shaman, sim *core.Simulation) {
	agent.numLBsSinceLastCL = agent.numLBsPerCL // This lets us cast CL first
}

func NewFixedRotationAgent(sim *core.Simulation, numLBsPerCL int) *FixedRotationAgent {
	return &FixedRotationAgent{
		numLBsPerCL:       numLBsPerCL,
	}
}

// ################################################################
//                          CL ON CLEARCAST
// ################################################################
type CLOnClearcastAgent struct {
	// Whether the second-to-last spell procced clearcasting
	prevPrevCastProccedCC bool
}

func (agent *CLOnClearcastAgent) ChooseAction(shaman *Shaman, sim *core.Simulation) core.AgentAction {
	if shaman.IsOnCD(core.MagicIDChainLightning6, sim.CurrentTime) || !agent.prevPrevCastProccedCC {
		return NewLightningBolt(sim, shaman, false)
	}

	return NewChainLightning(sim, shaman, false)
}

func (agent *CLOnClearcastAgent) OnActionAccepted(shaman *Shaman, sim *core.Simulation, action core.AgentAction) {
	agent.prevPrevCastProccedCC = shaman.elementalFocusStacks == 2
}

func (agent *CLOnClearcastAgent) Reset(shaman *Shaman, sim *core.Simulation) {
	agent.prevPrevCastProccedCC = true // Lets us cast CL first
}

func NewCLOnClearcastAgent(sim *core.Simulation) *CLOnClearcastAgent {
	return &CLOnClearcastAgent{}
}

// ################################################################
//                             ADAPTIVE
// ################################################################
type AdaptiveAgent struct {
	// Circular array buffer for recent mana snapshots, within a time window
	manaSnapshots      [manaSnapshotsBufferSize]ManaSnapshot
	numSnapshots       int32
	firstSnapshotIndex int32

	baseAgent    shamanAgent // The agent used most of the time
	surplusAgent shamanAgent // The agent used when we have extra mana
}

const manaSpendingWindowNumSeconds = 60
const manaSpendingWindow = time.Second * manaSpendingWindowNumSeconds

// 2 * (# of seconds) should be plenty of slots
const manaSnapshotsBufferSize = manaSpendingWindowNumSeconds * 2

type ManaSnapshot struct {
	time      time.Duration // time this snapshot was taken
	manaSpent float64       // total amount of mana spent up to this time
}

func (agent *AdaptiveAgent) getOldestSnapshot() ManaSnapshot {
	return agent.manaSnapshots[agent.firstSnapshotIndex]
}

func (agent *AdaptiveAgent) purgeExpiredSnapshots(sim *core.Simulation) {
	expirationCutoff := sim.CurrentTime - manaSpendingWindow

	curIndex := agent.firstSnapshotIndex
	for agent.numSnapshots > 0 && agent.manaSnapshots[curIndex].time < expirationCutoff {
		curIndex = (curIndex + 1) % manaSnapshotsBufferSize
		agent.numSnapshots--
	}
	agent.firstSnapshotIndex = curIndex
}

func (agent *AdaptiveAgent) takeSnapshot(sim *core.Simulation, shaman *Shaman) {
	if agent.numSnapshots >= manaSnapshotsBufferSize {
		panic("Agent snapshot buffer full")
	}

	snapshot := ManaSnapshot{
		time:      sim.CurrentTime,
		manaSpent: sim.GetIndividualMetrics(shaman.ID).ManaSpent,
	}

	nextIndex := (agent.firstSnapshotIndex + agent.numSnapshots) % manaSnapshotsBufferSize
	agent.manaSnapshots[nextIndex] = snapshot
	agent.numSnapshots++
}

func (agent *AdaptiveAgent) ChooseAction(shaman *Shaman, sim *core.Simulation) core.AgentAction {
	agent.purgeExpiredSnapshots(sim)
	oldestSnapshot := agent.getOldestSnapshot()

	manaSpent := sim.GetIndividualMetrics(shaman.ID).ManaSpent - oldestSnapshot.manaSpent
	timeDelta := sim.CurrentTime - oldestSnapshot.time
	if timeDelta == 0 {
		timeDelta = 1
	}

	timeRemaining := sim.Duration - sim.CurrentTime
	projectedManaCost := manaSpent * (timeRemaining.Seconds() / timeDelta.Seconds())

	if sim.Log != nil {
		manaSpendingRate := manaSpent / timeDelta.Seconds()
		sim.Log("[AI] CL Ready: Mana/s: %0.1f, Est Mana Cost: %0.1f, CurrentMana: %0.1f\n", manaSpendingRate, projectedManaCost, shaman.Stats[stats.Mana])
	}

	// If we have enough mana to burn, use the surplus agent.
	if projectedManaCost < shaman.Stats[stats.Mana] {
		return agent.surplusAgent.ChooseAction(shaman, sim)
	} else {
		return agent.baseAgent.ChooseAction(shaman, sim)
	}
}
func (agent *AdaptiveAgent) OnActionAccepted(shaman *Shaman, sim *core.Simulation, action core.AgentAction) {
	agent.takeSnapshot(sim, shaman)
	agent.baseAgent.OnActionAccepted(shaman, sim, action)
	agent.surplusAgent.OnActionAccepted(shaman, sim, action)
}

func (agent *AdaptiveAgent) Reset(shaman *Shaman, sim *core.Simulation) {
	agent.manaSnapshots = [manaSnapshotsBufferSize]ManaSnapshot{}
	agent.firstSnapshotIndex = 0
	agent.numSnapshots = 0
	agent.baseAgent.Reset(shaman, sim)
	agent.surplusAgent.Reset(shaman, sim)
}

func NewAdaptiveAgent(sim *core.Simulation) *AdaptiveAgent {
	agent := &AdaptiveAgent{}

	clearcastParams := sim.IndividualParams
	clearcastParams.Options.Debug = false
	clearcastParams.Options.Iterations = 100

	// eleShamParams := *clearcastParams.PlayerOptions.GetElementalShaman()
	// eleShamParams.Agent.Type = proto.ElementalShaman_Agent_CLOnClearcast
	params := *clearcastParams.PlayerOptions.GetElementalShaman()

	eleShamParams := params                                                                             // clone
	eleShamParams.Agent = &proto.ElementalShaman_Agent{Type: proto.ElementalShaman_Agent_CLOnClearcast} // create new agent.

	// Assign new eleShamParams
	clearcastParams.PlayerOptions = &proto.PlayerOptions{
		Race: sim.IndividualParams.PlayerOptions.Race, //primitive, no pointer
		Spec: &proto.PlayerOptions_ElementalShaman{
			ElementalShaman: &eleShamParams,
		},
		// reuse pointer since this isn't mutated
		Consumes: sim.IndividualParams.PlayerOptions.Consumes,
	}

	clearcastSim := core.NewIndividualSim(clearcastParams)
	clearcastResult := clearcastSim.Run()

	if clearcastResult.Agents[0].NumOom >= 5 {
		agent.baseAgent = NewLBOnlyAgent(sim)
		agent.surplusAgent = NewCLOnClearcastAgent(sim)
	} else {
		agent.baseAgent = NewCLOnClearcastAgent(sim)
		agent.surplusAgent = NewCLOnCDAgent(sim)
	}

	return agent
}
