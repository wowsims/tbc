package elemental

import (
	"time"

	"github.com/wowsims/tbc/sim/common/rotations"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
	. "github.com/wowsims/tbc/sim/shaman"
	googleProto "google.golang.org/protobuf/proto"
)

func RegisterElementalShaman() {
	core.RegisterAgentFactory(proto.PlayerOptions_ElementalShaman{}, func(character core.Character, options proto.PlayerOptions, isr proto.IndividualSimRequest) core.Agent {
		return NewElementalShaman(character, options, isr)
	})
}

func NewElementalShaman(character core.Character, options proto.PlayerOptions, isr proto.IndividualSimRequest) *Shaman {
	eleShamOptions := options.GetElementalShaman()

	selfBuffs := SelfBuffs{
		Bloodlust:    eleShamOptions.Options.Bloodlust,
		ManaSpring:   eleShamOptions.Options.ManaSpringTotem,
		TotemOfWrath: eleShamOptions.Options.TotemOfWrath,
		WrathOfAir:   eleShamOptions.Options.WrathOfAirTotem,
		WaterShield:  eleShamOptions.Options.WaterShield,
	}

	shaman := NewShaman(character, *eleShamOptions.Talents, selfBuffs)

	var rotation rotations.SimpleCasterRotationImpl

	switch eleShamOptions.Rotation.Type {
	case proto.ElementalShaman_Rotation_Adaptive:
		rotation = NewAdaptiveRotation(shaman, isr)
	case proto.ElementalShaman_Rotation_CLOnClearcast:
		rotation = NewCLOnClearcastRotation(shaman)
	case proto.ElementalShaman_Rotation_CLOnCD:
		rotation = NewCLOnCDRotation(shaman)
	case proto.ElementalShaman_Rotation_FixedLBCL:
		rotation = NewFixedRotation(shaman, eleShamOptions.Rotation.LbsPerCl)
	case proto.ElementalShaman_Rotation_LBOnly:
		rotation = NewLBOnlyRotation(shaman)
	}

	shaman.Rotation = rotations.NewSimpleCasterRotation(rotation)

	return shaman
}

// ################################################################
//                              LB ONLY
// ################################################################
type LBOnlyRotation struct {
	shamanAgent *Shaman
}

func (rotation *LBOnlyRotation) GetAgent() core.Agent {
	return rotation.shamanAgent
}

func (rotation *LBOnlyRotation) ChooseAction(sim *core.Simulation) core.AgentAction {
	return NewLightningBolt(sim, rotation.shamanAgent, false)
}

func (rotation *LBOnlyRotation) ChooseOOMAction(sim *core.Simulation, oomDuration time.Duration) core.AgentAction {
	return core.NewWaitAction(sim, rotation.shamanAgent.GetCharacter(), oomDuration)
}

func (rotation *LBOnlyRotation) OnActionAccepted(sim *core.Simulation, action core.AgentAction) {
}
func (rotation *LBOnlyRotation) Reset(sim *core.Simulation) {}

func NewLBOnlyRotation(shamanAgent *Shaman) *LBOnlyRotation {
	return &LBOnlyRotation{
		shamanAgent: shamanAgent,
	}
}

// ################################################################
//                             CL ON CD
// ################################################################
type CLOnCDRotation struct {
	shamanAgent *Shaman
}

func (rotation *CLOnCDRotation) GetAgent() core.Agent {
	return rotation.shamanAgent
}

func (rotation *CLOnCDRotation) ChooseAction(sim *core.Simulation) core.AgentAction {
	if rotation.shamanAgent.IsOnCD(ChainLightningCooldownID, sim.CurrentTime) {
		return NewLightningBolt(sim, rotation.shamanAgent, false)
	} else {
		return NewChainLightning(sim, rotation.shamanAgent, false)
	}
}

func (rotation *CLOnCDRotation) ChooseOOMAction(sim *core.Simulation, oomDuration time.Duration) core.AgentAction {
	return core.NewWaitAction(sim, rotation.shamanAgent.GetCharacter(), oomDuration)
}

func (rotation *CLOnCDRotation) OnActionAccepted(sim *core.Simulation, action core.AgentAction) {
}
func (rotation *CLOnCDRotation) Reset(sim *core.Simulation) {}

func NewCLOnCDRotation(shamanAgent *Shaman) *CLOnCDRotation {
	return &CLOnCDRotation{
		shamanAgent: shamanAgent,
	}
}

// ################################################################
//                          FIXED ROTATION
// ################################################################
type FixedRotation struct {
	shamanAgent *Shaman

	numLBsPerCL       int32
	numLBsSinceLastCL int32
}

func (rotation *FixedRotation) GetAgent() core.Agent {
	return rotation.shamanAgent
}

// Returns if any temporary haste buff is currently active.
// TODO: Figure out a way to make this automatic
func (rotation *FixedRotation) temporaryHasteActive(shamanAgent *Shaman) bool {
	return shamanAgent.HasAura(core.BloodlustAuraID) ||
		shamanAgent.HasAura(core.TrollBerserkingAuraID) ||
		shamanAgent.HasTemporaryBonusForStat(stats.SpellHaste)
}

func (rotation *FixedRotation) ChooseAction(sim *core.Simulation) core.AgentAction {
	if rotation.numLBsSinceLastCL < rotation.numLBsPerCL {
		return NewLightningBolt(sim, rotation.shamanAgent, false)
	}

	if !rotation.shamanAgent.IsOnCD(ChainLightningCooldownID, sim.CurrentTime) {
		return NewChainLightning(sim, rotation.shamanAgent, false)
	}

	// If we have a temporary haste effect (like bloodlust or quags eye) then
	// we should add LB casts instead of waiting
	if rotation.temporaryHasteActive(rotation.shamanAgent) {
		return NewLightningBolt(sim, rotation.shamanAgent, false)
	}

	return core.NewWaitAction(sim, rotation.shamanAgent.GetCharacter(), rotation.shamanAgent.GetRemainingCD(ChainLightningCooldownID, sim.CurrentTime))
}

func (rotation *FixedRotation) ChooseOOMAction(sim *core.Simulation, oomDuration time.Duration) core.AgentAction {
	return core.NewWaitAction(sim, rotation.shamanAgent.GetCharacter(), oomDuration)
}

func (rotation *FixedRotation) OnActionAccepted(sim *core.Simulation, action core.AgentAction) {
	cast, isCastAction := action.(*core.DirectCastAction)
	if !isCastAction {
		return
	}

	if cast.GetActionID().SpellID == SpellIDLB12 {
		rotation.numLBsSinceLastCL++
	} else if cast.GetActionID().SpellID == SpellIDCL6 {
		rotation.numLBsSinceLastCL = 0
	}
}

func (rotation *FixedRotation) Reset(sim *core.Simulation) {
	rotation.numLBsSinceLastCL = rotation.numLBsPerCL // This lets us cast CL first
}

func NewFixedRotation(shamanAgent *Shaman, numLBsPerCL int32) *FixedRotation {
	return &FixedRotation{
		shamanAgent: shamanAgent,
		numLBsPerCL: numLBsPerCL,
	}
}

// ################################################################
//                          CL ON CLEARCAST
// ################################################################
type CLOnClearcastRotation struct {
	shamanAgent *Shaman

	// Whether the second-to-last spell procced clearcasting
	prevPrevCastProccedCC bool
}

func (rotation *CLOnClearcastRotation) GetAgent() core.Agent {
	return rotation.shamanAgent
}

func (rotation *CLOnClearcastRotation) ChooseAction(sim *core.Simulation) core.AgentAction {
	if rotation.shamanAgent.IsOnCD(ChainLightningCooldownID, sim.CurrentTime) || !rotation.prevPrevCastProccedCC {
		return NewLightningBolt(sim, rotation.shamanAgent, false)
	}

	return NewChainLightning(sim, rotation.shamanAgent, false)
}

func (rotation *CLOnClearcastRotation) ChooseOOMAction(sim *core.Simulation, oomDuration time.Duration) core.AgentAction {
	return core.NewWaitAction(sim, rotation.shamanAgent.GetCharacter(), oomDuration)
}

func (rotation *CLOnClearcastRotation) OnActionAccepted(sim *core.Simulation, action core.AgentAction) {
	rotation.prevPrevCastProccedCC = rotation.shamanAgent.ElementalFocusStacks == 2
}

func (rotation *CLOnClearcastRotation) Reset(sim *core.Simulation) {
	rotation.prevPrevCastProccedCC = true // Lets us cast CL first
}

func NewCLOnClearcastRotation(shamanAgent *Shaman) *CLOnClearcastRotation {
	return &CLOnClearcastRotation{
		shamanAgent: shamanAgent,
	}
}

// ################################################################
//                             ADAPTIVE
// ################################################################
type AdaptiveRotation struct {
	shamanAgent *Shaman

	// Circular array buffer for recent mana snapshots, within a time window
	manaSnapshots      [manaSnapshotsBufferSize]ManaSnapshot
	numSnapshots       int32
	firstSnapshotIndex int32

	baseRotation    rotations.SimpleCasterRotationImpl // The rotation used most of the time
	surplusRotation rotations.SimpleCasterRotationImpl // The rotation used when we have extra mana
}

const manaSpendingWindowNumSeconds = 60
const manaSpendingWindow = time.Second * manaSpendingWindowNumSeconds

// 2 * (# of seconds) should be plenty of slots
const manaSnapshotsBufferSize = manaSpendingWindowNumSeconds * 2

type ManaSnapshot struct {
	time      time.Duration // time this snapshot was taken
	manaSpent float64       // total amount of mana spent up to this time
}

func (rotation *AdaptiveRotation) GetAgent() core.Agent {
	return rotation.shamanAgent
}

func (rotation *AdaptiveRotation) getOldestSnapshot() ManaSnapshot {
	return rotation.manaSnapshots[rotation.firstSnapshotIndex]
}

func (rotation *AdaptiveRotation) purgeExpiredSnapshots(sim *core.Simulation) {
	expirationCutoff := sim.CurrentTime - manaSpendingWindow

	curIndex := rotation.firstSnapshotIndex
	for rotation.numSnapshots > 0 && rotation.manaSnapshots[curIndex].time < expirationCutoff {
		curIndex = (curIndex + 1) % manaSnapshotsBufferSize
		rotation.numSnapshots--
	}
	rotation.firstSnapshotIndex = curIndex
}

func (rotation *AdaptiveRotation) takeSnapshot(sim *core.Simulation, shamanAgent *Shaman) {
	if rotation.numSnapshots >= manaSnapshotsBufferSize {
		panic("Rotation snapshot buffer full")
	}

	snapshot := ManaSnapshot{
		time:      sim.CurrentTime,
		manaSpent: sim.GetIndividualMetrics(shamanAgent.ID).ManaSpent,
	}

	nextIndex := (rotation.firstSnapshotIndex + rotation.numSnapshots) % manaSnapshotsBufferSize
	rotation.manaSnapshots[nextIndex] = snapshot
	rotation.numSnapshots++
}

func (rotation *AdaptiveRotation) ChooseAction(sim *core.Simulation) core.AgentAction {
	rotation.purgeExpiredSnapshots(sim)
	oldestSnapshot := rotation.getOldestSnapshot()

	manaSpent := sim.GetIndividualMetrics(rotation.shamanAgent.ID).ManaSpent - oldestSnapshot.manaSpent
	timeDelta := sim.CurrentTime - oldestSnapshot.time
	if timeDelta == 0 {
		timeDelta = 1
	}

	timeRemaining := sim.Duration - sim.CurrentTime
	projectedManaCost := manaSpent * (timeRemaining.Seconds() / timeDelta.Seconds())

	if sim.Log != nil {
		manaSpendingRate := manaSpent / timeDelta.Seconds()
		sim.Log("[AI] CL Ready: Mana/s: %0.1f, Est Mana Cost: %0.1f, CurrentMana: %0.1f\n", manaSpendingRate, projectedManaCost, rotation.shamanAgent.CurrentMana())
	}

	// If we have enough mana to burn, use the surplus agent.
	if projectedManaCost < rotation.shamanAgent.CurrentMana() {
		return rotation.surplusRotation.ChooseAction(sim)
	} else {
		return rotation.baseRotation.ChooseAction(sim)
	}
}

func (rotation *AdaptiveRotation) ChooseOOMAction(sim *core.Simulation, oomDuration time.Duration) core.AgentAction {
	return core.NewWaitAction(sim, rotation.shamanAgent.GetCharacter(), oomDuration)
}

func (rotation *AdaptiveRotation) OnActionAccepted(sim *core.Simulation, action core.AgentAction) {
	rotation.takeSnapshot(sim, rotation.shamanAgent)
	rotation.baseRotation.OnActionAccepted(sim, action)
	rotation.surplusRotation.OnActionAccepted(sim, action)
}

func (rotation *AdaptiveRotation) Reset(sim *core.Simulation) {
	rotation.manaSnapshots = [manaSnapshotsBufferSize]ManaSnapshot{}
	rotation.firstSnapshotIndex = 0
	rotation.numSnapshots = 0
	rotation.baseRotation.Reset(sim)
	rotation.surplusRotation.Reset(sim)
}

func NewAdaptiveRotation(shamanAgent *Shaman, isr proto.IndividualSimRequest) *AdaptiveRotation {
	rotation := &AdaptiveRotation{
		shamanAgent: shamanAgent,
	}

	// If no encounter is set, it means we aren't going to run a sim at all.
	// So just return something valid.
	// TODO: Probably need some organized way of doing presims so we dont have
	// to check these types of things.
	if isr.Encounter == nil || len(isr.Encounter.Targets) == 0 {
		rotation.baseRotation = NewLBOnlyRotation(shamanAgent)
		rotation.surplusRotation = NewCLOnClearcastRotation(shamanAgent)
		return rotation
	}

	clearcastRequest := googleProto.Clone(&isr).(*proto.IndividualSimRequest)
	clearcastRequest.SimOptions.Debug = false
	clearcastRequest.SimOptions.Iterations = 100
	clearcastRequest.Player.Options.Spec.(*proto.PlayerOptions_ElementalShaman).ElementalShaman.Rotation.Type = proto.ElementalShaman_Rotation_CLOnClearcast

	clearcastSim := core.NewIndividualSim(*clearcastRequest)
	clearcastResult := clearcastSim.Run()

	if clearcastResult.Agents[0].NumOom >= 5 {
		rotation.baseRotation = NewLBOnlyRotation(shamanAgent)
		rotation.surplusRotation = NewCLOnClearcastRotation(shamanAgent)
	} else {
		rotation.baseRotation = NewCLOnClearcastRotation(shamanAgent)
		rotation.surplusRotation = NewCLOnCDRotation(shamanAgent)
	}

	return rotation
}
