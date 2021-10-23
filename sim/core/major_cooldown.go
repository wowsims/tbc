package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

const (
	CooldownPriorityDefault = 0.0
	CooldownPriorityDrums = 2.0
	CooldownPriorityBloodlust = 1.0
)

// Function for activating a cooldown.
// Returns whether the activation was successful.
type CooldownActivation func(*Simulation, *Character) bool

type MajorCooldown struct {
	// Primary cooldown ID for checking whether this cooldown is ready.
	CooldownID int32

	// Amount of time before activation can be used again, after a successful
	// activation.
	Cooldown time.Duration

	// Secondary cooldown ID, used for shared cooldowns.
	SharedCooldownID int32

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

	// Lambda function to use the cooldown.
	TryActivate CooldownActivation
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

// Add a major cooldown to the given agent, which provides a temporary boost to a single stat.
// This is use for effects like Icon of the Silver Crescent and Bloodlust Brooch.
func RegisterTemporaryStatsOnUseCD(agent Agent, id int32, stat stats.Stat, amount float64, duration time.Duration, mcd MajorCooldown) {
	// If shared cooldown ID is set but shared cooldown isn't, default to duration.
	// Most items on a shared cooldown put each other on that cooldown for the
	// duration of their active effect.
	if mcd.SharedCooldownID != 0 && mcd.SharedCooldown == 0 {
		mcd.SharedCooldown = duration
	}

	mcd.TryActivate = func(sim *Simulation, character *Character) bool {
		AddAuraWithTemporaryStats(sim, character, id, stat, amount, duration)
		return true
	}

	agent.GetCharacter().AddMajorCooldown(mcd)
}

// Helper function to make an ApplyEffect for a temporary stats on-use cooldown.
func MakeTemporaryStatsOnUseCDRegistration(id int32, stat stats.Stat, amount float64, duration time.Duration, mcd MajorCooldown) ApplyEffect {
	return func(agent Agent) {
		RegisterTemporaryStatsOnUseCD(agent, id, stat, amount, duration, mcd)
	}
}
