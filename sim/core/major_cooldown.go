package core

import (
	"sort"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const (
	CooldownPriorityLow       = -1.0
	CooldownPriorityDefault   = 0.0
	CooldownPriorityDrums     = 2.0
	CooldownPriorityBloodlust = 1.0
)

const (
	CooldownTypeUnknown = 0
	CooldownTypeMana    = 1
	CooldownTypeDPS     = 2
)

// Condition for whether a cooldown can/should be activated.
// Returning false prevents the cooldown from being activated.
type CooldownActivationCondition func(*Simulation, *Character) bool

// Function for activating a cooldown.
// Returns whether the activation was successful.
type CooldownActivation func(*Simulation, *Character)

// Function for making a CooldownActivation.
//
// We need a function that returns a CooldownActivation rather than a
// CooldownActivation, so captured local variables can be reset on Sim reset.
type CooldownActivationFactory func(*Simulation) CooldownActivation

type MajorCooldown struct {
	// Unique ID for this cooldown, used to look it up.
	ActionID

	// Primary cooldown ID for checking whether this cooldown is ready.
	CooldownID CooldownID

	// Amount of time before activation can be used again, after a successful
	// activation.
	Cooldown time.Duration

	// Secondary cooldown ID, used for shared cooldowns.
	SharedCooldownID CooldownID

	// Duration of secondary cooldown.
	SharedCooldown time.Duration

	// How long before this cooldown takes effect after activation.
	// Not used yet, but will eventually be important for planning cooldown
	// schedules.
	CastTime time.Duration

	// Cooldowns with higher priority get used first. This is important when some
	// cooldowns have a non-zero cast time. For example, Drums should be used
	// before Bloodlust.
	Priority float64

	// Internal category, used for filtering. For example, mages want to disable
	// all DPS cooldowns during their regen rotation.
	Type int32

	// Whether the cooldown meets all hard requirements for activation (e.g. resource cost).
	// Note chat whether the cooldown is off CD is automatically checked, so it does not
	// need to be checked again by this function.
	CanActivate CooldownActivationCondition

	// Whether the cooldown meets all optional conditions for activation. These
	// conditions will be ignored when the user specifies their own activation time.
	// This is for things like mana thresholds, which are optimizations for better
	// automatic timing.
	ShouldActivate CooldownActivationCondition

	// Factory for creating the activate function on every Sim reset.
	ActivationFactory CooldownActivationFactory

	// Fixed timings at which to use this cooldown. If these are specified, they
	// are used instead of ShouldActivate.
	timings []time.Duration

	// Number of times this MCD was used so far in the current iteration.
	numUsages int

	// Internal lambda function to use the cooldown.
	activate CooldownActivation

	// Whether this MCD is currently disabled.
	disabled bool
}

func (mcd *MajorCooldown) IsOnCD(sim *Simulation, character *Character) bool {
	// Even if SharedCooldownID == 0 this will work since we never call SetCD(0, currentTime)
	return character.IsOnCD(mcd.CooldownID, sim.CurrentTime) || character.IsOnCD(mcd.SharedCooldownID, sim.CurrentTime)
}

func (mcd *MajorCooldown) GetRemainingCD(currentTime time.Duration, character *Character) time.Duration {
	return MaxDuration(
		character.GetRemainingCD(mcd.CooldownID, currentTime),
		character.GetRemainingCD(mcd.SharedCooldownID, currentTime))
}

func (mcd *MajorCooldown) IsEnabled() bool {
	return !mcd.disabled
}

// Activates this MCD, if all the conditions pass.
// Returns whether the MCD was activated.
func (mcd *MajorCooldown) tryActivate(sim *Simulation, character *Character) bool {
	if mcd.disabled {
		return false
	}

	if !mcd.CanActivate(sim, character) {
		return false
	}

	var shouldActivate bool
	if mcd.numUsages < len(mcd.timings) {
		shouldActivate = sim.CurrentTime >= mcd.timings[mcd.numUsages]
	} else {
		shouldActivate = mcd.ShouldActivate(sim, character)
	}

	if shouldActivate {
		mcd.activate(sim, character)
		mcd.numUsages++
		if sim.Log != nil {
			character.Log(sim, "Major cooldown used: %s", mcd.ActionID)
		}
	}

	return shouldActivate
}

type majorCooldownManager struct {
	// The Character whose cooldowns are being managed.
	character *Character

	// User-specified cooldown configs.
	cooldownConfigs proto.Cooldowns

	// Cached list of major cooldowns sorted by priority, for resetting quickly.
	initialMajorCooldowns []MajorCooldown

	// Whether finalize() has been called on this object.
	finalized bool

	// Major cooldowns, ordered by next available. This should always contain
	// the same cooldows as initialMajorCooldowns, but the order will change over
	// the course of the sim.
	majorCooldowns []*MajorCooldown
}

func newMajorCooldownManager(cooldowns *proto.Cooldowns) majorCooldownManager {
	cds := proto.Cooldowns{}
	if cooldowns != nil {
		cds = *cooldowns
	}

	return majorCooldownManager{
		cooldownConfigs: cds,
	}
}

func (mcdm *majorCooldownManager) finalize(character *Character) {
	if mcdm.finalized {
		return
	}
	mcdm.finalized = true

	mcdm.character = character

	if mcdm.initialMajorCooldowns == nil {
		mcdm.initialMajorCooldowns = []MajorCooldown{}
	}

	// Sort major cooldowns by descending priority so they get used in the correct order.
	sort.SliceStable(mcdm.initialMajorCooldowns, func(i, j int) bool {
		return mcdm.initialMajorCooldowns[i].Priority > mcdm.initialMajorCooldowns[j].Priority
	})

	// Match user-specified cooldown configs to existing cooldowns.
	for i, _ := range mcdm.initialMajorCooldowns {
		mcd := &mcdm.initialMajorCooldowns[i]
		mcd.timings = []time.Duration{}

		if mcdm.cooldownConfigs.Cooldowns != nil {
			for _, cooldownConfig := range mcdm.cooldownConfigs.Cooldowns {
				configID := ProtoToActionID(*cooldownConfig.Id)
				if configID.SameAction(mcd.ActionID) {
					mcd.timings = make([]time.Duration, len(cooldownConfig.Timings))
					for t, timing := range cooldownConfig.Timings {
						mcd.timings[t] = DurationFromSeconds(timing)
					}
					break
				}
			}
		}
	}

	mcdm.majorCooldowns = make([]*MajorCooldown, len(mcdm.initialMajorCooldowns))
}

func (mcdm *majorCooldownManager) reset(sim *Simulation) {
	// Need to create all cooldowns before calling ActivationFactory on any of them,
	// so that any cooldown can do lookups on other cooldowns.
	for i, _ := range mcdm.majorCooldowns {
		newMCD := &MajorCooldown{}
		*newMCD = mcdm.initialMajorCooldowns[i]
		mcdm.majorCooldowns[i] = newMCD
	}

	for i, _ := range mcdm.majorCooldowns {
		mcdm.majorCooldowns[i].activate = mcdm.majorCooldowns[i].ActivationFactory(sim)
		if mcdm.majorCooldowns[i].activate == nil {
			panic("Nil cooldown activation returned!")
		}
	}
}

// Registers a major cooldown to the Character, which will be automatically
// used when available.
func (mcdm *majorCooldownManager) AddMajorCooldown(mcd MajorCooldown) {
	if mcdm.finalized {
		panic("Major cooldowns may not be added once finalized!")
	}

	if mcd.IsEmptyAction() {
		panic("Major cooldown must have an ActionID!")
	}

	if mcd.CanActivate == nil || mcd.ShouldActivate == nil {
		panic("Major cooldown must provide CanActivate and ShouldActivate callbacks!")
	}

	mcdm.initialMajorCooldowns = append(mcdm.initialMajorCooldowns, mcd)
}

func (mcdm *majorCooldownManager) GetMajorCooldown(actionID ActionID) *MajorCooldown {
	for _, mcd := range mcdm.majorCooldowns {
		if mcd.SameAction(actionID) {
			return mcd
		}
	}

	return nil
}

// Returns all MCDs.
func (mcdm *majorCooldownManager) GetMajorCooldowns() []*MajorCooldown {
	return mcdm.majorCooldowns
}

func (mcdm *majorCooldownManager) GetMajorCooldownIDs() []*proto.ActionID {
	ids := make([]*proto.ActionID, len(mcdm.initialMajorCooldowns))
	for i, mcd := range mcdm.initialMajorCooldowns {
		ids[i] = mcd.ActionID.ToProto()
	}
	return ids
}

func (mcdm *majorCooldownManager) HasMajorCooldown(actionID ActionID) bool {
	return mcdm.GetMajorCooldown(actionID) != nil
}

func (mcdm *majorCooldownManager) DisableMajorCooldown(actionID ActionID) {
	mcd := mcdm.GetMajorCooldown(actionID)
	if mcd != nil {
		mcd.disabled = true
	}
}

func (mcdm *majorCooldownManager) EnableMajorCooldown(actionID ActionID) {
	mcd := mcdm.GetMajorCooldown(actionID)
	if mcd != nil {
		mcd.disabled = false
	}
}

func (mcdm *majorCooldownManager) TryUseCooldowns(sim *Simulation) {
	anyCooldownsUsed := false
	for curIdx := 0; curIdx < len(mcdm.majorCooldowns) && !mcdm.majorCooldowns[curIdx].IsOnCD(sim, mcdm.character); curIdx++ {
		mcd := mcdm.majorCooldowns[curIdx]
		if mcd.tryActivate(sim, mcdm.character) {
			anyCooldownsUsed = true
		}
	}

	if anyCooldownsUsed {
		// Re-sort by availability.
		// TODO: Probably a much faster way to do this, especially since we know which cooldowns need to be re-ordered.
		// Maybe not because MCDs with shared cooldowns put each other on CD.
		mcdm.UpdateMajorCooldowns()
	}
}

// This function should be called if the CD for a major cooldown changes outside
// of the TryActivate() call.
func (mcdm *majorCooldownManager) UpdateMajorCooldowns() {
	sort.Slice(mcdm.majorCooldowns, func(i, j int) bool {
		// Since we're just comparing and don't actually care about the remaining CD, ok to use 0 instead of sim.CurrentTime.
		return mcdm.majorCooldowns[i].GetRemainingCD(0, mcdm.character) < mcdm.majorCooldowns[j].GetRemainingCD(0, mcdm.character)
	})
}

// Add a major cooldown to the given agent, which provides a temporary boost to a single stat.
// This is use for effects like Icon of the Silver Crescent and Bloodlust Brooch.
func RegisterTemporaryStatsOnUseCD(agent Agent, auraID AuraID, stat stats.Stat, amount float64, duration time.Duration, mcd MajorCooldown) {
	// If shared cooldown ID is set but shared cooldown isn't, default to duration.
	// Most items on a shared cooldown put each other on that cooldown for the
	// duration of their active effect.
	if mcd.SharedCooldownID != 0 && mcd.SharedCooldown == 0 {
		mcd.SharedCooldown = duration
	}

	mcd.CanActivate = func(sim *Simulation, character *Character) bool { return true }
	mcd.ShouldActivate = func(sim *Simulation, character *Character) bool { return true }
	mcd.Type = CooldownTypeDPS

	mcd.ActivationFactory = func(sim *Simulation) CooldownActivation {
		return func(sim *Simulation, character *Character) {
			character.AddAuraWithTemporaryStats(sim, auraID, mcd.ActionID, stat, amount, duration)
			character.SetCD(mcd.CooldownID, sim.CurrentTime+mcd.Cooldown)
			if mcd.SharedCooldownID != 0 {
				character.SetCD(mcd.SharedCooldownID, sim.CurrentTime+mcd.SharedCooldown)
			}
		}
	}

	agent.GetCharacter().AddMajorCooldown(mcd)
}

// Helper function to make an ApplyEffect for a temporary stats on-use cooldown.
func MakeTemporaryStatsOnUseCDRegistration(auraID AuraID, stat stats.Stat, amount float64, duration time.Duration, mcd MajorCooldown) ApplyEffect {
	return func(agent Agent) {
		RegisterTemporaryStatsOnUseCD(agent, auraID, stat, amount, duration, mcd)
	}
}
