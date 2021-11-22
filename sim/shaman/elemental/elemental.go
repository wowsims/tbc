package elemental

import (
	"time"

	"github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
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
		Shaman:   NewShaman(character, *eleShamOptions.Talents, selfBuffs),
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
	dropTime := eleShaman.TryDropTotems(sim)
	if dropTime > 0 {
		return dropTime
	}
	newAction := eleShaman.rotation.ChooseAction(eleShaman, sim)

	actionSuccessful := newAction.Cast(sim)
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
	ChooseAction(*ElementalShaman, *core.Simulation) AgentAction

	// This will be invoked right before the chosen action is actually executed, so the rotation can update its state.
	// Note that the action may be different from the action chosen by this rotation.
	OnActionAccepted(*ElementalShaman, *core.Simulation, AgentAction)

	// Returns this rotation to its initial state. Called before each Sim iteration.
	Reset(*ElementalShaman, *core.Simulation)
}

// ################################################################
//                              LB ONLY
// ################################################################
type LBOnlyRotation struct {
}

func (rotation *LBOnlyRotation) ChooseAction(eleShaman *ElementalShaman, sim *core.Simulation) AgentAction {
	return eleShaman.NewLightningBolt(sim, sim.GetPrimaryTarget(), false)
}

func (rotation *LBOnlyRotation) OnActionAccepted(eleShaman *ElementalShaman, sim *core.Simulation, action AgentAction) {
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

func (rotation *CLOnCDRotation) ChooseAction(eleShaman *ElementalShaman, sim *core.Simulation) AgentAction {
	if eleShaman.IsOnCD(ChainLightningCooldownID, sim.CurrentTime) {
		return eleShaman.NewLightningBolt(sim, sim.GetPrimaryTarget(), false)
	} else {
		return eleShaman.NewChainLightning(sim, sim.GetPrimaryTarget(), false)
	}
}

func (rotation *CLOnCDRotation) OnActionAccepted(eleShaman *ElementalShaman, sim *core.Simulation, action AgentAction) {
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

func (rotation *FixedRotation) ChooseAction(eleShaman *ElementalShaman, sim *core.Simulation) AgentAction {
	if rotation.numLBsSinceLastCL < rotation.numLBsPerCL {
		return eleShaman.NewLightningBolt(sim, sim.GetPrimaryTarget(), false)
	}

	if !eleShaman.IsOnCD(ChainLightningCooldownID, sim.CurrentTime) {
		return eleShaman.NewChainLightning(sim, sim.GetPrimaryTarget(), false)
	}

	// If we have a temporary haste effect (like bloodlust or quags eye) then
	// we should add LB casts instead of waiting
	if eleShaman.HasTemporarySpellCastSpeedIncrease() {
		return eleShaman.NewLightningBolt(sim, sim.GetPrimaryTarget(), false)
	}

	return core.NewWaitAction(sim, eleShaman.GetCharacter(), eleShaman.GetRemainingCD(ChainLightningCooldownID, sim.CurrentTime))
}

func (rotation *FixedRotation) OnActionAccepted(eleShaman *ElementalShaman, sim *core.Simulation, action AgentAction) {
	if action.GetActionID().SpellID == SpellIDLB12 {
		rotation.numLBsSinceLastCL++
	} else if action.GetActionID().SpellID == SpellIDCL6 {
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

func (rotation *CLOnClearcastRotation) ChooseAction(eleShaman *ElementalShaman, sim *core.Simulation) AgentAction {
	if eleShaman.IsOnCD(ChainLightningCooldownID, sim.CurrentTime) || !rotation.prevPrevCastProccedCC {
		return eleShaman.NewLightningBolt(sim, sim.GetPrimaryTarget(), false)
	}

	return eleShaman.NewChainLightning(sim, sim.GetPrimaryTarget(), false)
}

func (rotation *CLOnClearcastRotation) OnActionAccepted(eleShaman *ElementalShaman, sim *core.Simulation, action AgentAction) {
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
	manaTracker common.ManaSpendingRateTracker

	baseRotation    Rotation // The rotation used most of the time
	surplusRotation Rotation // The rotation used when we have extra mana
}

func (rotation *AdaptiveRotation) ChooseAction(eleShaman *ElementalShaman, sim *core.Simulation) AgentAction {
	projectedManaCost := rotation.manaTracker.ProjectedManaCost(sim, eleShaman.GetCharacter())

	// If we have enough mana to burn, use the surplus rotation.
	if projectedManaCost < eleShaman.CurrentMana() {
		return rotation.surplusRotation.ChooseAction(eleShaman, sim)
	} else {
		return rotation.baseRotation.ChooseAction(eleShaman, sim)
	}
}
func (rotation *AdaptiveRotation) OnActionAccepted(eleShaman *ElementalShaman, sim *core.Simulation, action AgentAction) {
	rotation.manaTracker.Update(sim, eleShaman.GetCharacter())
	rotation.baseRotation.OnActionAccepted(eleShaman, sim, action)
	rotation.surplusRotation.OnActionAccepted(eleShaman, sim, action)
}

func (rotation *AdaptiveRotation) Reset(eleShaman *ElementalShaman, sim *core.Simulation) {
	rotation.manaTracker.Reset()
	rotation.baseRotation.Reset(eleShaman, sim)
	rotation.surplusRotation.Reset(eleShaman, sim)
}

func NewAdaptiveRotation(isr proto.IndividualSimRequest) *AdaptiveRotation {
	rotation := &AdaptiveRotation{
		manaTracker: common.NewManaSpendingRateTracker(),
	}

	// If no encounter is set, it means we aren't going to run a sim at all.
	// So just return something valid.
	// TODO: Probably need some organized way of doing presims so we dont have
	// to check these types of things.
	if isr.Encounter == nil || len(isr.Encounter.Targets) == 0 {
		rotation.baseRotation = NewLBOnlyRotation()
		rotation.surplusRotation = NewCLOnClearcastRotation()
		return rotation
	}

	presimRequest := googleProto.Clone(&isr).(*proto.IndividualSimRequest)
	presimRequest.SimOptions.Debug = false
	presimRequest.SimOptions.Iterations = 100
	presimRequest.Player.Options.Spec.(*proto.PlayerOptions_ElementalShaman).ElementalShaman.Rotation.Type = proto.ElementalShaman_Rotation_CLOnClearcast

	presimResult := core.RunIndividualSim(presimRequest)

	if presimResult.PlayerMetrics.NumOom >= 5 {
		rotation.baseRotation = NewLBOnlyRotation()
		rotation.surplusRotation = NewCLOnClearcastRotation()
	} else {
		rotation.baseRotation = NewCLOnClearcastRotation()
		rotation.surplusRotation = NewCLOnCDRotation()
	}

	return rotation
}

// A single action that an Agent can take.
type AgentAction interface {
	GetActionID() core.ActionID

	// For logging / debugging.
	GetName() string

	// The Character performing this action.
	GetCharacter() *core.Character

	// How long this action takes to cast/channel/etc.
	// In other words, how long until another action should be chosen.
	GetDuration() time.Duration

	// TODO: Maybe change this to 'ResourceCost'
	// Amount of mana required to perform the action.
	GetManaCost() float64

	// Do the action. Returns whether the action was successful. An unsuccessful
	// action indicates that the prerequisites, like resource cost, were not met.
	Cast(sim *core.Simulation) bool
}
