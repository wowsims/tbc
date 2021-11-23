package core

import (
	"sort"
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

const (
	CooldownPriorityDefault   = 0.0
	CooldownPriorityDrums     = 2.0
	CooldownPriorityBloodlust = 1.0
)

// Function for activating a cooldown.
// Returns whether the activation was successful.
type CooldownActivation func(*Simulation, *Character) bool

// Function for making a CooldownActivation.
//
// We need a function that returns a CooldownActivation rather than a
// CooldownActivation, so captured local variables can be reset on Sim reset.
type CooldownActivationFactory func(*Simulation) CooldownActivation

type MajorCooldown struct {
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

	// Factory for creating the activate function on every Sim reset.
	ActivationFactory CooldownActivationFactory

	// Internal lambda function to use the cooldown.
	activate CooldownActivation
}

func (mcd MajorCooldown) IsOnCD(sim *Simulation, character *Character) bool {
	// Even if SharedCooldownID == 0 this will work since we never call SetCD(0, currentTime)
	return character.IsOnCD(mcd.CooldownID, sim.CurrentTime) || character.IsOnCD(mcd.SharedCooldownID, sim.CurrentTime)
}

func (mcd MajorCooldown) GetRemainingCD(sim *Simulation, character *Character) time.Duration {
	return MaxDuration(
		character.GetRemainingCD(mcd.CooldownID, sim.CurrentTime),
		character.GetRemainingCD(mcd.SharedCooldownID, sim.CurrentTime))
}

type majorCooldownManager struct {
	// The Character whose cooldowns are being managed.
	character *Character

	// Cached list of major cooldowns sorted by priority, for resetting quickly.
	initialMajorCooldowns []MajorCooldown

	// Whether finalize() has been called on this object.
	finalized bool

	// Major cooldowns, ordered by next available. This should always contain
	// the same cooldows as initialMajorCooldowns, but the order will change over
	// the course of the sim.
	majorCooldowns []MajorCooldown
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
}

func (mcdm *majorCooldownManager) reset(sim *Simulation) {
	mcdm.majorCooldowns = make([]MajorCooldown, len(mcdm.initialMajorCooldowns))
	copy(mcdm.majorCooldowns, mcdm.initialMajorCooldowns)

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

	if mcdm.initialMajorCooldowns == nil {
		mcdm.initialMajorCooldowns = []MajorCooldown{}
	}

	mcdm.initialMajorCooldowns = append(mcdm.initialMajorCooldowns, mcd)
}

func (mcdm *majorCooldownManager) TryUseCooldowns(sim *Simulation) {
	anyCooldownsUsed := false
	for curIdx := 0; curIdx < len(mcdm.majorCooldowns) && !mcdm.majorCooldowns[curIdx].IsOnCD(sim, mcdm.character); curIdx++ {
		success := mcdm.majorCooldowns[curIdx].activate(sim, mcdm.character)
		anyCooldownsUsed = anyCooldownsUsed || success
	}

	if anyCooldownsUsed {
		// Re-sort by availability.
		// TODO: Probably a much faster way to do this, especially since we know which cooldowns need to be re-ordered.
		sort.Slice(mcdm.majorCooldowns, func(i, j int) bool {
			return mcdm.majorCooldowns[i].GetRemainingCD(sim, mcdm.character) < mcdm.majorCooldowns[j].GetRemainingCD(sim, mcdm.character)
		})
	}
}

// This function should be called if the CD for a major cooldown changes outside
// of the TryActivate() call.
func (mcdm *majorCooldownManager) UpdateMajorCooldowns(sim *Simulation) {
	sort.Slice(mcdm.majorCooldowns, func(i, j int) bool {
		return mcdm.majorCooldowns[i].GetRemainingCD(sim, mcdm.character) < mcdm.majorCooldowns[j].GetRemainingCD(sim, mcdm.character)
	})
}

// Add a major cooldown to the given agent, which provides a temporary boost to a single stat.
// This is use for effects like Icon of the Silver Crescent and Bloodlust Brooch.
func RegisterTemporaryStatsOnUseCD(agent Agent, auraID AuraID, spellID int32, auraName string, stat stats.Stat, amount float64, duration time.Duration, mcd MajorCooldown) {
	// If shared cooldown ID is set but shared cooldown isn't, default to duration.
	// Most items on a shared cooldown put each other on that cooldown for the
	// duration of their active effect.
	if mcd.SharedCooldownID != 0 && mcd.SharedCooldown == 0 {
		mcd.SharedCooldown = duration
	}

	mcd.ActivationFactory = func(sim *Simulation) CooldownActivation {
		return func(sim *Simulation, character *Character) bool {
			character.AddAuraWithTemporaryStats(sim, auraID, spellID, auraName, stat, amount, duration)
			character.SetCD(mcd.CooldownID, sim.CurrentTime+mcd.Cooldown)
			if mcd.SharedCooldownID != 0 {
				character.SetCD(mcd.SharedCooldownID, sim.CurrentTime+mcd.SharedCooldown)
			}
			return true
		}
	}

	agent.GetCharacter().AddMajorCooldown(mcd)
}

// Helper function to make an ApplyEffect for a temporary stats on-use cooldown.
func MakeTemporaryStatsOnUseCDRegistration(auraID AuraID, spellID int32, auraName string, stat stats.Stat, amount float64, duration time.Duration, mcd MajorCooldown) ApplyEffect {
	return func(agent Agent) {
		RegisterTemporaryStatsOnUseCD(agent, auraID, spellID, auraName, stat, amount, duration, mcd)
	}
}
