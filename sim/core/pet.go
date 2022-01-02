package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Extension of Agent interface, for Pets.
type PetAgent interface {
	Agent

	// The Pet controlled by this PetAgent.
	GetPet() *Pet
}

// Pet is an extension of Character, for any entity created by a player that can
// take actions on its own.
type Pet struct {
	Character

	Owner *Character

	baseStats stats.Stats

	// Coefficients for each stat that correspond to the proportion
	// of that stat inherited from the owner.
	statInheritanceCoeffs stats.Stats

	initialEnabled bool

	// Whether this pet is currently active. Pets which are active throughout a whole
	// encounter, like Hunter pets, are always enabled. Pets which are instead summoned,
	// such as Mage Water Elemental, begin as disabled and are enabled when summoned.
	enabled bool

	// The PendingAction corresponding to this Pet, or nil if this Pet is currently
	// disabled.
	pendingAction *PendingAction

	// Some pets expire after a certain duration. This is the pending action that disables
	// the pet on expiration.
	timeoutAction *PendingAction
}

func NewPet(name string, owner *Character, baseStats stats.Stats, statInheritanceCoeffs stats.Stats, enabledOnStart bool) Pet {
	pet := Pet{
		Character: Character{
			Name: name,
			PseudoStats: stats.PseudoStats{
				AttackSpeedMultiplier: 1,
				CastSpeedMultiplier:   1,
				SpiritRegenMultiplier: 1,
			},
			Party:       owner.Party,
			PartyIndex:  owner.PartyIndex,
			RaidIndex:   owner.RaidIndex,
			auraTracker: newAuraTracker(false),
			Metrics:     NewCharacterMetrics(),
		},
		Owner:                 owner,
		baseStats:             baseStats,
		statInheritanceCoeffs: statInheritanceCoeffs,
		initialEnabled:        enabledOnStart,
	}

	pet.AddStats(baseStats)

	return pet
}

// Updates the stats for this pet in response to a stat change on the owner.
// addedStats is the amount of stats added to the owner (will be negative if the
// owner lost stats).
func (pet *Pet) addOwnerStats(addedStats stats.Stats) {
	inheritedChange := addedStats.DotProduct(pet.statInheritanceCoeffs)
	pet.AddStats(inheritedChange)
}
func (pet *Pet) addOwnerStat(stat stats.Stat, addedAmount float64) {
	inheritedChange := addedAmount * pet.statInheritanceCoeffs[stat]
	pet.AddStat(stat, inheritedChange)
}

// This needs to be called after owner stats are finalized so we can inherit the
// final values.
func (pet *Pet) Finalize() {
	inheritedStats := pet.Owner.GetStats().DotProduct(pet.statInheritanceCoeffs)
	pet.AddStats(inheritedStats)
	pet.Character.Finalize()
}

func (pet *Pet) reset(sim *Simulation) {
	pet.Character.reset(sim)
	pet.pendingAction = nil
	pet.enabled = false
}
func (pet *Pet) advance(sim *Simulation, elapsedTime time.Duration) {
	pet.Character.advance(sim, elapsedTime)
}
func (pet *Pet) doneIteration(simDuration time.Duration) {
	pet.Character.doneIteration(simDuration)
}

func (pet *Pet) IsEnabled() bool {
	return pet.enabled
}

// petAgent should be the PetAgent which embeds this Pet.
func (pet *Pet) Enable(sim *Simulation, petAgent PetAgent) {
	if pet.enabled {
		panic("Pet is already enabled!")
	}

	if sim.Log != nil {
		pet.Log(sim, "Pet summoned")
	}

	pet.pendingAction = sim.newDefaultAgentAction(petAgent)
	sim.AddPendingAction(pet.pendingAction)

	pet.enabled = true
}
func (pet *Pet) Disable(sim *Simulation) {
	if !pet.enabled {
		panic("Pet is already disabled!")
	}

	pet.pendingAction.Cancel(sim)
	pet.pendingAction = nil
	pet.enabled = false

	if pet.timeoutAction != nil {
		pet.timeoutAction.Cancel(sim)
		pet.timeoutAction = nil
	}

	if sim.Log != nil {
		pet.Log(sim, "Pet dismissed")
	}
}

// Helper for enabling a pet that will expire after a certain duration.
func (pet *Pet) EnableWithTimeout(sim *Simulation, petAgent PetAgent, petDuration time.Duration) {
	pet.Enable(sim, petAgent)

	pet.timeoutAction = &PendingAction{
		Name:         "Pet Timeout",
		NextActionAt: sim.CurrentTime + petDuration,
		OnAction: func(sim *Simulation) {
			pet.Disable(sim)
		},
	}
	sim.AddPendingAction(pet.timeoutAction)
}

// Default implementations for some Agent functions which most Pets don't need.
func (pet *Pet) GetCharacter() *Character {
	return &pet.Character
}
func (pet *Pet) AddRaidBuffs(raidBuffs *proto.RaidBuffs)    {}
func (pet *Pet) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {}
