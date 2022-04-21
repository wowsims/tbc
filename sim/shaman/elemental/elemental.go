package elemental

import (
	"time"

	"github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
	"github.com/wowsims/tbc/sim/shaman"
)

func RegisterElementalShaman() {
	core.RegisterAgentFactory(
		proto.Player_ElementalShaman{},
		proto.Spec_SpecElementalShaman,
		func(character core.Character, options proto.Player) core.Agent {
			return NewElementalShaman(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_ElementalShaman)
			if !ok {
				panic("Invalid spec value for Elemental Shaman!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewElementalShaman(character core.Character, options proto.Player) *ElementalShaman {
	eleShamOptions := options.GetElementalShaman()

	selfBuffs := shaman.SelfBuffs{
		Bloodlust:        eleShamOptions.Options.Bloodlust,
		WaterShield:      eleShamOptions.Options.WaterShield,
		SnapshotWOAT42Pc: eleShamOptions.Options.SnapshotT4_2Pc,
	}

	totems := proto.ShamanTotems{}
	if eleShamOptions.Rotation.Totems != nil {
		totems = *eleShamOptions.Rotation.Totems
	}

	var rotation Rotation

	switch eleShamOptions.Rotation.Type {
	case proto.ElementalShaman_Rotation_Adaptive:
		rotation = NewAdaptiveRotation()
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
		Shaman:   shaman.NewShaman(character, *eleShamOptions.Talents, totems, selfBuffs),
		rotation: rotation,
		has4pT6:  shaman.ItemSetSkyshatterRegalia.CharacterHasSetBonus(&character, 4),
	}
}

type ElementalShaman struct {
	*shaman.Shaman

	rotation Rotation

	has4pT6 bool
}

func (eleShaman *ElementalShaman) GetShaman() *shaman.Shaman {
	return eleShaman.Shaman
}

func (eleShaman *ElementalShaman) GetPresimOptions() *core.PresimOptions {
	return eleShaman.rotation.GetPresimOptions()
}

func (eleShaman *ElementalShaman) Reset(sim *core.Simulation) {
	eleShaman.Shaman.Reset(sim)
	eleShaman.rotation.Reset(eleShaman, sim)
}

func (eleShaman *ElementalShaman) OnGCDReady(sim *core.Simulation) {
	eleShaman.tryUseGCD(sim)
}

func (eleShaman *ElementalShaman) OnManaTick(sim *core.Simulation) {
	if eleShaman.FinishedWaitingForManaAndGCDReady(sim) {
		eleShaman.tryUseGCD(sim)
	}
}

func (eleShaman *ElementalShaman) tryUseGCD(sim *core.Simulation) {
	if eleShaman.TryDropTotems(sim) {
		return
	}

	eleShaman.rotation.DoAction(eleShaman, sim)
	//actionSuccessful := newAction.Cast(sim)
	//if actionSuccessful {
	//	eleShaman.rotation.OnActionAccepted(eleShaman, sim, newAction)
	//} else {
	//	// Only way for a shaman spell to fail is due to mana cost.
	//	// Wait until we have enough mana to cast.
	//	eleShaman.WaitForMana(sim, newAction.GetManaCost())
	//}
}

// Picks which attacks / abilities the Shaman does.
type Rotation interface {
	GetPresimOptions() *core.PresimOptions

	// Returns the action this rotation would like to take next.
	DoAction(*ElementalShaman, *core.Simulation)

	// Returns this rotation to its initial state. Called before each Sim iteration.
	Reset(*ElementalShaman, *core.Simulation)
}

// ################################################################
//                              LB ONLY
// ################################################################
type LBOnlyRotation struct {
}

func (rotation *LBOnlyRotation) DoAction(eleShaman *ElementalShaman, sim *core.Simulation) {
	if !eleShaman.LightningBolt.Cast(sim, sim.GetPrimaryTarget()) {
		eleShaman.WaitForMana(sim, eleShaman.LightningBolt.CurCast.Cost)
	}
}

func (rotation *LBOnlyRotation) Reset(eleShaman *ElementalShaman, sim *core.Simulation) {}
func (rotation *LBOnlyRotation) GetPresimOptions() *core.PresimOptions                  { return nil }

func NewLBOnlyRotation() *LBOnlyRotation {
	return &LBOnlyRotation{}
}

// ################################################################
//                             CL ON CD
// ################################################################
type CLOnCDRotation struct {
}

func (rotation *CLOnCDRotation) DoAction(eleShaman *ElementalShaman, sim *core.Simulation) {
	var spell *core.Spell
	if eleShaman.ChainLightning.IsReady(sim) {
		spell = eleShaman.ChainLightning
	} else {
		spell = eleShaman.LightningBolt
	}

	if !spell.Cast(sim, sim.GetPrimaryTarget()) {
		eleShaman.WaitForMana(sim, spell.CurCast.Cost)
	}
}

func (rotation *CLOnCDRotation) Reset(eleShaman *ElementalShaman, sim *core.Simulation) {}
func (rotation *CLOnCDRotation) GetPresimOptions() *core.PresimOptions                  { return nil }

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

func (rotation *FixedRotation) DoAction(eleShaman *ElementalShaman, sim *core.Simulation) {
	var spell *core.Spell
	if rotation.numLBsSinceLastCL < rotation.numLBsPerCL {
		spell = eleShaman.LightningBolt
		rotation.numLBsSinceLastCL++
	} else if eleShaman.ChainLightning.IsReady(sim) {
		spell = eleShaman.ChainLightning
		rotation.numLBsSinceLastCL = 0
	} else if eleShaman.HasTemporarySpellCastSpeedIncrease() {
		// If we have a temporary haste effect (like bloodlust or quags eye) then
		// we should add LB casts instead of waiting
		spell = eleShaman.LightningBolt
		rotation.numLBsSinceLastCL++
	}

	if spell == nil {
		common.NewWaitAction(sim, eleShaman.GetCharacter(), eleShaman.ChainLightning.TimeToReady(sim), common.WaitReasonRotation).Cast(sim)
	} else {
		if !spell.Cast(sim, sim.GetPrimaryTarget()) {
			eleShaman.WaitForMana(sim, spell.CurCast.Cost)
		}
	}
}

func (rotation *FixedRotation) Reset(eleShaman *ElementalShaman, sim *core.Simulation) {
	rotation.numLBsSinceLastCL = rotation.numLBsPerCL // This lets us cast CL first
}

func (rotation *FixedRotation) GetPresimOptions() *core.PresimOptions { return nil }

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

func (rotation *CLOnClearcastRotation) DoAction(eleShaman *ElementalShaman, sim *core.Simulation) {
	var spell *core.Spell
	if !eleShaman.ChainLightning.IsReady(sim) || !rotation.prevPrevCastProccedCC {
		spell = eleShaman.LightningBolt
	} else {
		spell = eleShaman.ChainLightning
	}

	if !spell.Cast(sim, sim.GetPrimaryTarget()) {
		eleShaman.WaitForMana(sim, spell.CurCast.Cost)
	} else {
		rotation.prevPrevCastProccedCC = eleShaman.ElementalFocusStacks == 2
	}
}

func (rotation *CLOnClearcastRotation) Reset(eleShaman *ElementalShaman, sim *core.Simulation) {
	rotation.prevPrevCastProccedCC = true // Lets us cast CL first
}

func (rotation *CLOnClearcastRotation) GetPresimOptions() *core.PresimOptions { return nil }

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

func (rotation *AdaptiveRotation) DoAction(eleShaman *ElementalShaman, sim *core.Simulation) {
	didLB := false
	if sim.GetNumTargets() == 1 {
		sp := eleShaman.GetStat(stats.NatureSpellPower) + eleShaman.GetStat(stats.SpellPower)
		castSpeed := eleShaman.CastSpeed()
		lb := ((612 + (sp * 0.794)) * 1.2) / (2 / castSpeed)
		cl := ((786 + (sp * 0.651)) * 1.0666) / core.MaxFloat((1.5/castSpeed), 1)
		if eleShaman.has4pT6 {
			lb *= 1.05
		}
		if lb+10 >= cl {
			eleShaman.LightningBolt.Cast(sim, sim.GetPrimaryTarget())
			didLB = true
		}
	}

	if !didLB {
		// If we have enough mana to burn, use the surplus rotation.
		if rotation.manaTracker.ProjectedManaSurplus(sim, eleShaman.GetCharacter()) {
			rotation.surplusRotation.DoAction(eleShaman, sim)
		} else {
			rotation.baseRotation.DoAction(eleShaman, sim)
		}
	}

	rotation.manaTracker.Update(sim, eleShaman.GetCharacter())
}

func (rotation *AdaptiveRotation) Reset(eleShaman *ElementalShaman, sim *core.Simulation) {
	rotation.manaTracker.Reset()
	rotation.baseRotation.Reset(eleShaman, sim)
	rotation.surplusRotation.Reset(eleShaman, sim)
}

func (rotation *AdaptiveRotation) GetPresimOptions() *core.PresimOptions {
	return &core.PresimOptions{
		SetPresimPlayerOptions: func(player *proto.Player) {
			player.Spec.(*proto.Player_ElementalShaman).ElementalShaman.Rotation.Type = proto.ElementalShaman_Rotation_CLOnClearcast
		},

		OnPresimResult: func(presimResult proto.PlayerMetrics, iterations int32, duration time.Duration) bool {
			if float64(presimResult.SecondsOomAvg) >= 0.03*duration.Seconds() {
				rotation.baseRotation = NewLBOnlyRotation()
				rotation.surplusRotation = NewCLOnClearcastRotation()
			} else {
				rotation.baseRotation = NewCLOnClearcastRotation()
				rotation.surplusRotation = NewCLOnCDRotation()
			}
			return true
		},
	}
}

func NewAdaptiveRotation() *AdaptiveRotation {
	return &AdaptiveRotation{
		manaTracker: common.NewManaSpendingRateTracker(),
	}
}

// A single action that an Agent can take.
type AgentAction interface {
	GetActionID() core.ActionID

	// TODO: Maybe change this to 'ResourceCost'
	// Amount of mana required to perform the action.
	GetManaCost() float64

	// Do the action. Returns whether the action was successful. An unsuccessful
	// action indicates that the prerequisites, like resource cost, were not met.
	Cast(sim *core.Simulation) bool
}
