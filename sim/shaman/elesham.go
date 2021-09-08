package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func NewElemental(agent int, options map[string]string) *Elemental {
	return &Elemental{}
}

func NewEnhancement(agent int, options map[string]string) *Enhancement {
	return &Enhancement{}
}

type Enhancement struct {
	Agent
}

// BuffUp lets you buff up all players in sim.
func (e *Enhancement) BuffUp(sim *core.Simulation) {

}

type Elemental struct {
	CurrentMana float64
	Talents     Talents

	Agent
}

// BuffUp lets you buff up all players in sim.
func (e *Elemental) BuffUp(sim *core.Simulation) {
	// if sim.Options.Talents.Concussion > 0 {
	// 	bonusdmg := (0.01 * sim.Options.Talents.Concussion)
	// }
}

// Agent is shaman specific agent for behavior.
type Agent interface {
	// Returns the action this Agent would like to take next.
	ChooseAction(*core.Simulation) core.AgentAction

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

// ################################################################
//                              LB ONLY
// ################################################################
// type LBOnlyAgent struct {
// 	lb *core.Spell
// }

// func (agent *LBOnlyAgent) ChooseAction(sim *core.Simulation) AgentAction {
// 	return core.NewCastAction(sim, agent.lb)
// }

// func (agent *LBOnlyAgent) OnActionAccepted(sim *core.Simulation, action AgentAction) {}
// func (agent *LBOnlyAgent) Reset(sim *core.Simulation)                                {}

// func NewLBOnlyAgent(sim *core.Simulation) *LBOnlyAgent {
// 	return &LBOnlyAgent{
// 		lb: core.Spells[MagicIDLB12],
// 	}
// }

// ################################################################
//                             CL ON CD
// ################################################################
// type CLOnCDAgent struct {
// 	lb *core.Spell
// 	cl *core.Spell
// }

// func (agent *CLOnCDAgent) ChooseAction(sim *core.Simulation) AgentAction {
// 	if sim.isOnCD(MagicIDCL6) {
// 		return NewCastAction(sim, agent.lb)
// 	} else {
// 		return NewCastAction(sim, agent.cl)
// 	}
// }

// func (agent *CLOnCDAgent) OnActionAccepted(sim *core.Simulation, action AgentAction) {}
// func (agent *CLOnCDAgent) Reset(sim *core.Simulation)                                {}

// func NewCLOnCDAgent(sim *core.Simulation) *CLOnCDAgent {
// 	return &CLOnCDAgent{
// 		lb: spellmap[MagicIDLB12],
// 		cl: spellmap[MagicIDCL6],
// 	}
// }

// ################################################################
//                          FIXED ROTATION
// ################################################################
// type FixedRotationAgent struct {
// 	numLBsPerCL       int
// 	numLBsSinceLastCL int
// 	lb                *core.Spell
// 	cl                *core.Spell
// }

// // Returns if any temporary haste buff is currently active.
// // TODO: Figure out a way to make this automatic
// func (agent *FixedRotationAgent) temporaryHasteActive(sim *core.Simulation) bool {
// 	return sim.hasAura(MagicIDBloodlust) ||
// 		sim.hasAura(MagicIDDrums) ||
// 		sim.hasAura(MagicIDTrollBerserking) ||
// 		sim.hasAura(MagicIDSkullGuldan) ||
// 		sim.hasAura(MagicIDFungalFrenzy)
// }

// func (agent *FixedRotationAgent) ChooseAction(sim *core.Simulation) AgentAction {
// 	if agent.numLBsSinceLastCL < agent.numLBsPerCL {
// 		return NewCastAction(sim, agent.lb)
// 	}

// 	if !sim.isOnCD(MagicIDCL6) {
// 		return NewCastAction(sim, agent.cl)
// 	}

// 	// If we have a temporary haste effect (like bloodlust or quags eye) then
// 	// we should add LB casts instead of waiting
// 	if agent.temporaryHasteActive(sim) {
// 		return NewCastAction(sim, agent.lb)
// 	}

// 	return NewWaitAction(sim.getRemainingCD(MagicIDCL6))
// }

// func (agent *FixedRotationAgent) OnActionAccepted(sim *core.Simulation, action AgentAction) {
// 	if action.Cast == nil {
// 		return
// 	}

// 	if action.Cast.Spell.ID == MagicIDLB12 {
// 		agent.numLBsSinceLastCL++
// 	} else if action.Cast.Spell.ID == MagicIDCL6 {
// 		agent.numLBsSinceLastCL = 0
// 	}
// }

// func (agent *FixedRotationAgent) Reset(sim *core.Simulation) {
// 	agent.numLBsSinceLastCL = agent.numLBsPerCL
// }

// func NewFixedRotationAgent(sim *core.Simulation, numLBsPerCL int) *FixedRotationAgent {
// 	return &FixedRotationAgent{
// 		numLBsPerCL:       numLBsPerCL,
// 		numLBsSinceLastCL: numLBsPerCL, // This lets us cast CL first
// 		lb:                spellmap[MagicIDLB12],
// 		cl:                spellmap[MagicIDCL6],
// 	}
// }

// ################################################################
//                          CL ON CLEARCAST
// ################################################################
type CLOnClearcastAgent struct {
	// Whether the second-to-last spell procced clearcasting
	prevPrevCastProccedCC bool

	lb *core.Spell
	cl *core.Spell
}

func (agent *CLOnClearcastAgent) ChooseAction(sim *core.Simulation) AgentAction {
	if sim.isOnCD(MagicIDCL6) || !agent.prevPrevCastProccedCC {
		return NewCastAction(sim, agent.lb)
	}

	return NewCastAction(sim, agent.cl)
}

func (agent *CLOnClearcastAgent) OnActionAccepted(sim *core.Simulation, action AgentAction) {
	agent.prevPrevCastProccedCC = sim.auras[MagicIDEleFocus].stacks == 2
}

func (agent *CLOnClearcastAgent) Reset(sim *core.Simulation) {
	agent.prevPrevCastProccedCC = true // Lets us cast CL first
}

func NewCLOnClearcastAgent(sim *core.Simulation) *CLOnClearcastAgent {
	return &CLOnClearcastAgent{
		lb: spellmap[MagicIDLB12],
		cl: spellmap[MagicIDCL6],
	}
}

// ################################################################
//                             ADAPTIVE
// ################################################################
type AdaptiveAgent struct {
	// Circular array buffer for recent mana snapshots, within a time window
	manaSnapshots      [manaSnapshotsBufferSize]ManaSnapshot
	numSnapshots       int32
	firstSnapshotIndex int32

	p *core.Player // two way connection seems bad.

	baseAgent    Agent // The agent used most of the time
	surplusAgent Agent // The agent used when we have extra mana
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

func (agent *AdaptiveAgent) takeSnapshot(sim *core.Simulation) {
	if agent.numSnapshots >= manaSnapshotsBufferSize {
		panic("Agent snapshot buffer full")
	}

	snapshot := ManaSnapshot{
		time:      sim.CurrentTime,
		manaSpent: sim.Metrics.ManaSpent,
	}

	nextIndex := (agent.firstSnapshotIndex + agent.numSnapshots) % manaSnapshotsBufferSize
	agent.manaSnapshots[nextIndex] = snapshot
	agent.numSnapshots++
}

func (agent *AdaptiveAgent) ChooseAction(sim *core.Simulation) core.AgentAction {
	agent.purgeExpiredSnapshots(sim)
	oldestSnapshot := agent.getOldestSnapshot()

	manaSpent := sim.Metrics.ManaSpent - oldestSnapshot.manaSpent
	timeDelta := sim.CurrentTime - oldestSnapshot.time
	if timeDelta == 0 {
		timeDelta = 1
	}

	timeRemaining := sim.Duration - sim.CurrentTime
	projectedManaCost := manaSpent * (timeRemaining.Seconds() / timeDelta.Seconds())

	if sim.Debug != nil {
		manaSpendingRate := manaSpent / timeDelta.Seconds()
		sim.Debug("[AI] CL Ready: Mana/s: %0.1f, Est Mana Cost: %0.1f, CurrentMana: %0.1f\n", manaSpendingRate, projectedManaCost, sim.CurrentMana)
	}

	// If we have enough mana to burn, use the surplus agent.
	if projectedManaCost < agent.p.CurrentMana {
		return agent.surplusAgent.ChooseAction(sim)
	} else {
		return agent.baseAgent.ChooseAction(sim)
	}
}
func (agent *AdaptiveAgent) OnActionAccepted(sim *core.Simulation, action core.AgentAction) {
	agent.takeSnapshot(sim)
	agent.baseAgent.OnActionAccepted(sim, action)
	agent.surplusAgent.OnActionAccepted(sim, action)
}

func (agent *AdaptiveAgent) Reset(sim *core.Simulation) {
	agent.manaSnapshots = [manaSnapshotsBufferSize]ManaSnapshot{}
	agent.firstSnapshotIndex = 0
	agent.numSnapshots = 0
	agent.baseAgent.Reset(sim)
	agent.surplusAgent.Reset(sim)
}

func NewAdaptiveAgent(sim *core.Simulation) *AdaptiveAgent {
	agent := &AdaptiveAgent{}

	clearcastSimRequest := core.SimRequest{
		Options:    sim.Options,
		Gear:       sim.EquipSpec,
		Iterations: 100,
	}
	clearcastSimRequest.Options.AgentType = AGENT_TYPE_CL_ON_CLEARCAST
	clearcastResult := core.RunSimulation(clearcastSimRequest)

	if clearcastResult.NumOom >= 5 {
		agent.baseAgent = NewAgent(sim, AGENT_TYPE_FIXED_LB_ONLY)
		agent.surplusAgent = NewAgent(sim, AGENT_TYPE_CL_ON_CLEARCAST)
	} else {
		agent.baseAgent = NewAgent(sim, AGENT_TYPE_CL_ON_CLEARCAST)
		agent.surplusAgent = NewAgent(sim, AGENT_TYPE_FIXED_CL_ON_CD)
	}

	return agent
}

// func NewAgent(sim *core.Simulation, agentType AgentType) Agent {
// 	switch agentType {
// 	case AGENT_TYPE_FIXED_3LB_1CL:
// 		return NewFixedRotationAgent(sim, 3)
// 	case AGENT_TYPE_FIXED_4LB_1CL:
// 		return NewFixedRotationAgent(sim, 4)
// 	case AGENT_TYPE_FIXED_5LB_1CL:
// 		return NewFixedRotationAgent(sim, 5)
// 	case AGENT_TYPE_FIXED_6LB_1CL:
// 		return NewFixedRotationAgent(sim, 6)
// 	case AGENT_TYPE_FIXED_7LB_1CL:
// 		return NewFixedRotationAgent(sim, 7)
// 	case AGENT_TYPE_FIXED_8LB_1CL:
// 		return NewFixedRotationAgent(sim, 8)
// 	case AGENT_TYPE_FIXED_9LB_1CL:
// 		return NewFixedRotationAgent(sim, 9)
// 	case AGENT_TYPE_FIXED_10LB_1CL:
// 		return NewFixedRotationAgent(sim, 10)
// 	case AGENT_TYPE_FIXED_LB_ONLY:
// 		return NewLBOnlyAgent(sim)
// 	case AGENT_TYPE_FIXED_CL_ON_CD:
// 		return NewCLOnCDAgent(sim)
// 	case AGENT_TYPE_ADAPTIVE:
// 		return NewAdaptiveAgent(sim)
// 	case AGENT_TYPE_CL_ON_CLEARCAST:
// 		return NewCLOnClearcastAgent(sim)
// 	default:
// 		fmt.Printf("[ERROR] No rotation given to sim.\n")
// 		return nil
// 	}
// }
