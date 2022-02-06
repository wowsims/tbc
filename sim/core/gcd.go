package core

import (
	"time"
)

func (character *Character) newGCDAction(sim *Simulation, agent Agent) *PendingAction {
	return &PendingAction{
		Priority: ActionPriorityGCD,
		OnAction: func(sim *Simulation) {
			character := agent.GetCharacter()
			character.TryUseCooldowns(sim)
			if !character.IsOnCD(GCDCooldownID, sim.CurrentTime) {
				agent.OnGCDReady(sim)
			}
		},
	}
}

// Note that this is only used when the hardcast and GCD actions
func (character *Character) newHardcastAction(sim *Simulation) *PendingAction {
	return &PendingAction{
		OnAction: func(sim *Simulation) {
			// Don't need to do anything, the Advance() call will take care of the hardcast.
		},
	}
}

func (character *Character) NextGCDAt() time.Duration {
	return character.gcdAction.NextActionAt
}

func (character *Character) SetGCDTimer(sim *Simulation, gcdReadyAt time.Duration) {
	character.SetCD(GCDCooldownID, gcdReadyAt)

	character.gcdAction.Cancel(sim)
	character.gcdAction = &PendingAction{
		Priority:     ActionPriorityGCD,
		OnAction:     character.gcdAction.OnAction,
		NextActionAt: gcdReadyAt,
	}
	sim.AddPendingAction(character.gcdAction)
}

// Call this to stop the GCD loop for a character.
// This is mostly used for pets that get summoned / expire.
func (character *Character) CancelGCDTimer(sim *Simulation) {
	character.gcdAction.Cancel(sim)
}

func (character *Character) IsWaiting() bool {
	return character.waitStartTime != 0
}
func (character *Character) IsWaitingForMana() bool {
	return character.waitingForMana != 0
}

// Assumes that IsWaitingForMana() == true
func (character *Character) DoneWaitingForMana(sim *Simulation) bool {
	if character.CurrentMana() >= character.waitingForMana {
		character.Metrics.MarkOOM(character, sim.CurrentTime-character.waitStartTime)
		character.waitStartTime = 0
		character.waitingForMana = 0
		return true
	}
	return false
}

// Returns true if the character was waiting for mana but is now finished AND
// the GCD is also ready.
func (character *Character) FinishedWaitingForManaAndGCDReady(sim *Simulation) bool {
	if !character.IsWaitingForMana() || !character.DoneWaitingForMana(sim) {
		return false
	}

	return !character.IsOnCD(GCDCooldownID, sim.CurrentTime)
}

func (character *Character) WaitUntil(sim *Simulation, readyTime time.Duration) {
	character.waitStartTime = sim.CurrentTime
	character.SetGCDTimer(sim, readyTime)
	if sim.Log != nil {
		character.Log(sim, "Pausing GCD for %s due to rotation / CDs.", readyTime-sim.CurrentTime)
	}
}

func (character *Character) WaitForMana(sim *Simulation, desiredMana float64) {
	character.waitStartTime = sim.CurrentTime
	character.waitingForMana = desiredMana
	if sim.Log != nil {
		character.Log(sim, "Not enough mana to cast, pausing GCD until mana >= %0.01f.", desiredMana)
	}
}

func (character *Character) doneIterationGCD(simDuration time.Duration) {
	if character.IsWaitingForMana() {
		character.Metrics.MarkOOM(character, simDuration-character.waitStartTime)
		character.waitStartTime = 0
		character.waitingForMana = 0
	} else if character.IsWaiting() {
		character.waitStartTime = 0
	}
}
