package elemental

import (
	"time"

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

func NewElementalShaman(character core.Character, options proto.PlayerOptions, isr proto.IndividualSimRequest) *ElementalShaman {
	eleShamOptions := options.GetElementalShaman()

	selfBuffs := SelfBuffs{
		Bloodlust:    eleShamOptions.Options.Bloodlust,
		ManaSpring:   eleShamOptions.Options.ManaSpringTotem,
		TotemOfWrath: eleShamOptions.Options.TotemOfWrath,
		WrathOfAir:   eleShamOptions.Options.WrathOfAirTotem,
		WaterShield:  eleShamOptions.Options.WaterShield,
	}

	var rotation Rotation

	switch eleShamOptions.Rotation.Type {
	case proto.ElementalShaman_Rotation_Adaptive:
		rotation = NewAdaptiveRotation(isr)
	case proto.ElementalShaman_Rotation_CLOnClearcast:
		rotation = NewCLOnClearcastRotation()
	case proto.ElementalShaman_Rotation_CLOnCD:
		rotation = NewCLOnCDRotation()
	case proto.ElementalShaman_Rotation_FixedLBCL:
		rotation = NewFixedRotation(eleShamOptions.Rotation.LbsPerCl)
	case proto.ElementalShaman_Rotation_LBOnly:
		rotation = NewLBOnlyRotation()
	}

	return &ElementalShaman{
		Shaman: NewShaman(character, *eleShamOptions.Talents, selfBuffs),
		rotation: rotation,
	}
}

type ElementalShaman struct {
	Shaman

	rotation Rotation
}

func (eleShaman *ElementalShaman) GetShaman() *Shaman {
	return &eleShaman.Shaman
}

func (eleShaman *ElementalShaman) Reset(sim *core.Simulation) {
	eleShaman.Shaman.Reset(sim)
	eleShaman.rotation.Reset(eleShaman, sim)
}

func (eleShaman *ElementalShaman) Act(sim *core.Simulation) time.Duration {
	newAction := eleShaman.rotation.ChooseAction(eleShaman, sim)

	actionSuccessful := newAction.Act(sim)
	if actionSuccessful {
		eleShaman.rotation.OnActionAccepted(eleShaman, sim, newAction)
		return sim.CurrentTime + core.MaxDuration(
				eleShaman.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime),
				newAction.GetDuration())
	} else {
		// Only way for a shaman spell to fail is due to mana cost.
		// Wait until we have enough mana to cast.
		regenTime := eleShaman.TimeUntilManaRegen(newAction.GetManaCost())
		newAction = core.NewWaitAction(sim, eleShaman.GetCharacter(), regenTime)
		eleShaman.rotation.OnActionAccepted(eleShaman, sim, newAction)
		return sim.CurrentTime + regenTime
	}
}

// Picks which attacks / abilities the Shaman does.
type Rotation interface {
	// Returns the action this rotation would like to take next.
	ChooseAction(*ElementalShaman, *core.Simulation) core.AgentAction

	// This will be invoked right before the chosen action is actually executed, so the rotation can update its state.
	// Note that the action may be different from the action chosen by this rotation.
	OnActionAccepted(*ElementalShaman, *core.Simulation, core.AgentAction)

	// Returns this rotation to its initial state. Called before each Sim iteration.
	Reset(*ElementalShaman, *core.Simulation)
}

// ################################################################
//                              LB ONLY
// ################################################################
type LBOnlyRotation struct {
}

func (rotation *LBOnlyRotation) ChooseAction(eleShaman *ElementalShaman, sim *core.Simulation) core.AgentAction {
	return eleShaman.NewLightningBolt(sim, sim.GetPrimaryTarget(), false)
}

func (rotation *LBOnlyRotation) OnActionAccepted(eleShaman *ElementalShaman, sim *core.Simulation, action core.AgentAction) {
}
func (rotation *LBOnlyRotation) Reset(eleShaman *ElementalShaman, sim *core.Simulation) {}

func NewLBOnlyRotation() *LBOnlyRotation {
	return &LBOnlyRotation{}
}

// ################################################################
//                             CL ON CD
// ################################################################
type CLOnCDRotation struct {
}

func (rotation *CLOnCDRotation) ChooseAction(eleShaman *ElementalShaman, sim *core.Simulation) core.AgentAction {
	if eleShaman.IsOnCD(ChainLightningCooldownID, sim.CurrentTime) {
		return eleShaman.NewLightningBolt(sim, sim.GetPrimaryTarget(), false)
	} else {
		return eleShaman.NewChainLightning(sim, sim.GetPrimaryTarget(), false)
	}
}

func (rotation *CLOnCDRotation) OnActionAccepted(eleShaman *ElementalShaman, sim *core.Simulation, action core.AgentAction) {
}
func (rotation *CLOnCDRotation) Reset(eleShaman *ElementalShaman, sim *core.Simulation) {}

func NewCLOnCDRotation() *CLOnCDRotation {
	return &CLOnCDRotation{}
}

// ################################################################
//                          FIXED ROTATION
// ################################################################
type FixedRotation struct {
	numLBsPerCL       int32
	numLBsSinceLastCL int32
}

// Returns if any temporary haste buff is currently active.
// TODO: Figure out a way to make this automatic
func (rotation *FixedRotation) temporaryHasteActive(eleShaman *ElementalShaman) bool {
	return eleShaman.HasAura(core.BloodlustAuraID) ||
		eleShaman.HasAura(core.TrollBerserkingAuraID) ||
		eleShaman.HasTemporaryBonusForStat(stats.SpellHaste)
}

func (rotation *FixedRotation) ChooseAction(eleShaman *ElementalShaman, sim *core.Simulation) core.AgentAction {
	if rotation.numLBsSinceLastCL < rotation.numLBsPerCL {
		return eleShaman.NewLightningBolt(sim, sim.GetPrimaryTarget(), false)
	}

	if !eleShaman.IsOnCD(ChainLightningCooldownID, sim.CurrentTime) {
		return eleShaman.NewChainLightning(sim, sim.GetPrimaryTarget(), false)
	}

	// If we have a temporary haste effect (like bloodlust or quags eye) then
	// we should add LB casts instead of waiting
	if rotation.temporaryHasteActive(eleShaman) {
		return eleShaman.NewLightningBolt(sim, sim.GetPrimaryTarget(), false)
	}

	return core.NewWaitAction(sim, eleShaman.GetCharacter(), eleShaman.GetRemainingCD(ChainLightningCooldownID, sim.CurrentTime))
}

func (rotation *FixedRotation) OnActionAccepted(eleShaman *ElementalShaman, sim *core.Simulation, action core.AgentAction) {
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

func (rotation *FixedRotation) Reset(eleShaman *ElementalShaman, sim *core.Simulation) {
	rotation.numLBsSinceLastCL = rotation.numLBsPerCL // This lets us cast CL first
}

func NewFixedRotation(numLBsPerCL int32) *FixedRotation {
	return &FixedRotation{
		numLBsPerCL: numLBsPerCL,
	}
}

// ################################################################
//                          CL ON CLEARCAST
// ################################################################
type CLOnClearcastRotation struct {
	// Whether the second-to-last spell procced clearcasting
	prevPrevCastProccedCC bool
}

func (rotation *CLOnClearcastRotation) ChooseAction(eleShaman *ElementalShaman, sim *core.Simulation) core.AgentAction {
	if eleShaman.IsOnCD(ChainLightningCooldownID, sim.CurrentTime) || !rotation.prevPrevCastProccedCC {
		return eleShaman.NewLightningBolt(sim, sim.GetPrimaryTarget(), false)
	}

	return eleShaman.NewChainLightning(sim, sim.GetPrimaryTarget(), false)
}

func (rotation *CLOnClearcastRotation) OnActionAccepted(eleShaman *ElementalShaman, sim *core.Simulation, action core.AgentAction) {
	rotation.prevPrevCastProccedCC = eleShaman.ElementalFocusStacks == 2
}

func (rotation *CLOnClearcastRotation) Reset(eleShaman *ElementalShaman, sim *core.Simulation) {
	rotation.prevPrevCastProccedCC = true // Lets us cast CL first
}

func NewCLOnClearcastRotation() *CLOnClearcastRotation {
	return &CLOnClearcastRotation{}
}

// ################################################################
//                             ADAPTIVE
// ################################################################
type AdaptiveRotation struct {
	// Circular array buffer for recent mana snapshots, within a time window
	manaSnapshots      [manaSnapshotsBufferSize]ManaSnapshot
	numSnapshots       int32
	firstSnapshotIndex int32

	baseRotation    Rotation // The rotation used most of the time
	surplusRotation Rotation // The rotation used when we have extra mana
}

const manaSpendingWindowNumSeconds = 60
const manaSpendingWindow = time.Second * manaSpendingWindowNumSeconds

// 2 * (# of seconds) should be plenty of slots
const manaSnapshotsBufferSize = manaSpendingWindowNumSeconds * 2

type ManaSnapshot struct {
	time      time.Duration // time this snapshot was taken
	manaSpent float64       // total amount of mana spent up to this time
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

func (rotation *AdaptiveRotation) takeSnapshot(sim *core.Simulation, eleShaman *ElementalShaman) {
	if rotation.numSnapshots >= manaSnapshotsBufferSize {
		panic("Rotation snapshot buffer full")
	}

	snapshot := ManaSnapshot{
		time:      sim.CurrentTime,
		manaSpent: sim.GetIndividualMetrics(eleShaman.ID).ManaSpent,
	}

	nextIndex := (rotation.firstSnapshotIndex + rotation.numSnapshots) % manaSnapshotsBufferSize
	rotation.manaSnapshots[nextIndex] = snapshot
	rotation.numSnapshots++
}

func (rotation *AdaptiveRotation) ChooseAction(eleShaman *ElementalShaman, sim *core.Simulation) core.AgentAction {
	rotation.purgeExpiredSnapshots(sim)
	oldestSnapshot := rotation.getOldestSnapshot()

	manaSpent := sim.GetIndividualMetrics(eleShaman.ID).ManaSpent - oldestSnapshot.manaSpent
	timeDelta := sim.CurrentTime - oldestSnapshot.time
	if timeDelta == 0 {
		timeDelta = 1
	}

	timeRemaining := sim.Duration - sim.CurrentTime
	projectedManaCost := manaSpent * (timeRemaining.Seconds() / timeDelta.Seconds())

	if sim.Log != nil {
		manaSpendingRate := manaSpent / timeDelta.Seconds()
		sim.Log("[AI] CL Ready: Mana/s: %0.1f, Est Mana Cost: %0.1f, CurrentMana: %0.1f\n", manaSpendingRate, projectedManaCost, eleShaman.CurrentMana())
	}

	// If we have enough mana to burn, use the surplus rotation.
	if projectedManaCost < eleShaman.CurrentMana() {
		return rotation.surplusRotation.ChooseAction(eleShaman, sim)
	} else {
		return rotation.baseRotation.ChooseAction(eleShaman, sim)
	}
}
func (rotation *AdaptiveRotation) OnActionAccepted(eleShaman *ElementalShaman, sim *core.Simulation, action core.AgentAction) {
	rotation.takeSnapshot(sim, eleShaman)
	rotation.baseRotation.OnActionAccepted(eleShaman, sim, action)
	rotation.surplusRotation.OnActionAccepted(eleShaman, sim, action)
}

func (rotation *AdaptiveRotation) Reset(eleShaman *ElementalShaman, sim *core.Simulation) {
	rotation.manaSnapshots = [manaSnapshotsBufferSize]ManaSnapshot{}
	rotation.firstSnapshotIndex = 0
	rotation.numSnapshots = 0
	rotation.baseRotation.Reset(eleShaman, sim)
	rotation.surplusRotation.Reset(eleShaman, sim)
}

func NewAdaptiveRotation(isr proto.IndividualSimRequest) *AdaptiveRotation {
	rotation := &AdaptiveRotation{}

	// If no encounter is set, it means we aren't going to run a sim at all.
	// So just return something valid.
	// TODO: Probably need some organized way of doing presims so we dont have
	// to check these types of things.
	if isr.Encounter == nil || len(isr.Encounter.Targets) == 0 {
		rotation.baseRotation = NewLBOnlyRotation()
		rotation.surplusRotation = NewCLOnClearcastRotation()
		return rotation
	}

	clearcastRequest := googleProto.Clone(&isr).(*proto.IndividualSimRequest)
	clearcastRequest.SimOptions.Debug = false
	clearcastRequest.SimOptions.Iterations = 100
	clearcastRequest.Player.Options.Spec.(*proto.PlayerOptions_ElementalShaman).ElementalShaman.Rotation.Type = proto.ElementalShaman_Rotation_CLOnClearcast

	clearcastSim := core.NewIndividualSim(*clearcastRequest)
	clearcastResult := clearcastSim.Run()

	if clearcastResult.Agents[0].NumOom >= 5 {
		rotation.baseRotation = NewLBOnlyRotation()
		rotation.surplusRotation = NewCLOnClearcastRotation()
	} else {
		rotation.baseRotation = NewCLOnClearcastRotation()
		rotation.surplusRotation = NewCLOnCDRotation()
	}

	return rotation
}
