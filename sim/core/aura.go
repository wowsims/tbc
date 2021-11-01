package core

import (
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

const NeverExpires = time.Duration(math.MaxInt64)

type AuraID int32

var numAuraIDs = 0
func NewAuraID() AuraID {
	newAuraID := AuraID(numAuraIDs)
	numAuraIDs++
	return newAuraID
}

// Offsensive trinkets put each other on CD, so they can all share 1 aura ID
var OffensiveTrinketActiveAuraID = NewAuraID()
// Defensive trinkets put each other on CD, so they can all share 1 aura ID
var DefensiveTrinketActiveAuraID = NewAuraID()

type CooldownID int32

var numCooldownIDs = 0
func NewCooldownID() CooldownID {
	newCooldownID := CooldownID(numCooldownIDs)
	numCooldownIDs++
	return newCooldownID
}

var GCDCooldownID = NewCooldownID()
var MainHandSwingCooldownID = NewCooldownID()
var OffHandSwingCooldownID = NewCooldownID()
var RangedSwingCooldownID = NewCooldownID()
var OffensiveTrinketSharedCooldownID = NewCooldownID()
var DefensiveTrinketSharedCooldownID = NewCooldownID()

type Aura struct {
	ID          AuraID
	Name        string        // Label used for logging.
	Expires     time.Duration // time at which aura will be removed
	activeIndex int32         // Position of this aura's index in the sim.activeAuraIDs array

	// The number of stacks, or charges, of this aura. If this aura doesn't care
	// about charges, is just 0.
	Stacks int32

	// Invoked at creation time for a spell cast.
	OnCast         func(sim *Simulation, cast DirectCastAction, castInput *DirectCastInput)

	// Invoked when a spell cast completes casting, before results are calculated.
	OnCastComplete func(sim *Simulation, cast DirectCastAction)

	// Invoked before a spell lands, but after the target is selected.
	OnBeforeSpellHit func(sim *Simulation, hitInput *DirectCastDamageInput)

	// Invoked when a spell is fully resisted.
	OnSpellMiss    func(sim *Simulation, cast DirectCastAction)

	// Invoked when a spell hits, after results are calculated. Results can be modified by changing
	// properties of result.
	OnSpellHit     func(sim *Simulation, cast DirectCastAction, result *DirectCastDamageResult)

	// Invoked when this Aura expires.
	OnExpire       func(sim *Simulation)
}

// This needs to be a function that returns an Aura rather than an Aura, so captured
// local variables can be reset on Sim reset.
type PermanentAura func(*Simulation) Aura

// auraTracker is a centralized implementation of CD and Aura tracking.
//  This is currently used by Player and Raid (for global debuffs)
type auraTracker struct {
	// Auras that never expire and should always be active.
	// These are automatically applied on each Sim reset.
	permanentAuras []PermanentAura

	// Used for logging.
	playerID int

	// Whether finalize() has been called for this object.
	finalized bool

  // Maps MagicIDs to sim duration at which CD is done. Using array for perf.
	cooldowns []time.Duration

	// Maps MagicIDs to aura for that ID. Using array for perf.
	auras []Aura

	// IDs of Auras that are active, in no particular order
	activeAuraIDs []AuraID
}

func newAuraTracker() auraTracker {
	return auraTracker{
		permanentAuras: []PermanentAura{},
		activeAuraIDs: make([]AuraID, 0, 5),
	}
}

// Registers a permanent aura to this Character which will be re-applied on
// every Sim reset.
func (at *auraTracker) AddPermanentAura(permAura PermanentAura) {
	if at.finalized {
		panic("Permanent auras may not be added once finalized!")
	}

	at.permanentAuras = append(at.permanentAuras, permAura)
}

func (at *auraTracker) finalize() {
	if at.finalized {
		return
	}
	at.finalized = true
}

func (at *auraTracker) reset(sim *Simulation) {
	at.auras = make([]Aura, numAuraIDs)
	at.cooldowns = make([]time.Duration, numCooldownIDs)
	at.activeAuraIDs = at.activeAuraIDs[:0]

	for _, permAura := range at.permanentAuras {
		aura := permAura(sim)
		aura.Expires = NeverExpires
		at.AddAura(sim, aura)
	}
}

func (at *auraTracker) advance(sim *Simulation, newTime time.Duration) {
	// Go in reverse order so we can safely delete while looping
	for i := len(at.activeAuraIDs) - 1; i >= 0; i-- {
		id := at.activeAuraIDs[i]
		if at.auras[id].Expires != 0 && at.auras[id].Expires <= newTime {
			at.RemoveAura(sim, id)
		}
	}
}

// addAura will add a new aura to the simulation. If there is a matching aura ID
// it will be replaced with the newer aura.
// Auras with duration of 0 will be logged as activating but never added to simulation Auras.
func (at *auraTracker) AddAura(sim *Simulation, newAura Aura) {
	if newAura.Expires <= sim.CurrentTime {
		return // no need to waste time adding aura that doesn't last.
	}

	if at.HasAura(newAura.ID) {
		at.RemoveAura(sim, newAura.ID)
	}

	at.auras[newAura.ID] = newAura
	at.auras[newAura.ID].activeIndex = int32(len(at.activeAuraIDs))
	at.activeAuraIDs = append(at.activeAuraIDs, newAura.ID)

	if sim.Log != nil {
		sim.Log("(%d) +%s\n", at.playerID, newAura.Name)
	}
}

// Remove an aura by its ID
func (at *auraTracker) RemoveAura(sim *Simulation, id AuraID) {
	if at.auras[id].OnExpire != nil {
		at.auras[id].OnExpire(sim)
	}

	if sim.Log != nil {
		sim.Log("(%d) -%s\n", at.playerID, at.auras[id].Name)
	}

	removeActiveIndex := at.auras[id].activeIndex
	at.auras[id] = Aura{}

	// Overwrite the element being removed with the last element
	otherAuraID := at.activeAuraIDs[len(at.activeAuraIDs)-1]
	if id != otherAuraID {
		at.activeAuraIDs[removeActiveIndex] = otherAuraID
		at.auras[otherAuraID].activeIndex = removeActiveIndex
	}

	// Now we can remove the last element, in constant time
	at.activeAuraIDs = at.activeAuraIDs[:len(at.activeAuraIDs)-1]
}

// Returns whether an aura with the given ID is currently active.
func (at *auraTracker) HasAura(id AuraID) bool {
	return at.auras[id].ID != 0
}

func (at *auraTracker) IsOnCD(id CooldownID, currentTime time.Duration) bool {
	return at.cooldowns[id] > currentTime
}

func (at *auraTracker) GetRemainingCD(id CooldownID, currentTime time.Duration) time.Duration {
	remainingCD := at.cooldowns[id] - currentTime
	if remainingCD > 0 {
		return remainingCD
	} else {
		return 0
	}
}

func (at *auraTracker) SetCD(id CooldownID, newCD time.Duration) {
	at.cooldowns[id] = newCD
}

// Invokes the OnCast event for all tracked Auras.
func (at *auraTracker) OnCast(sim *Simulation, cast DirectCastAction, castInput *DirectCastInput) {
	for _, id := range at.activeAuraIDs {
		if at.auras[id].OnCast != nil {
			at.auras[id].OnCast(sim, cast, castInput)
		}
	}
}

// Invokes the OnCastComplete event for all tracked Auras.
func (at *auraTracker) OnCastComplete(sim *Simulation, cast DirectCastAction) {
	for _, id := range at.activeAuraIDs {
		if at.auras[id].OnCastComplete != nil {
			at.auras[id].OnCastComplete(sim, cast)
		}
	}
}

// Invokes the OnBeforeSpellHit event for all tracked Auras.
func (at *auraTracker) OnBeforeSpellHit(sim *Simulation, hitInput *DirectCastDamageInput) {
	for _, id := range at.activeAuraIDs {
		if at.auras[id].OnBeforeSpellHit != nil {
			at.auras[id].OnBeforeSpellHit(sim, hitInput)
		}
	}
}

// Invokes the OnSpellMiss event for all tracked Auras.
func (at *auraTracker) OnSpellMiss(sim *Simulation, cast DirectCastAction) {
	for _, id := range at.activeAuraIDs {
		if at.auras[id].OnSpellMiss != nil {
			at.auras[id].OnSpellMiss(sim, cast)
		}
	}
}

// Invokes the OnSpellHit event for all tracked Auras.
func (at *auraTracker) OnSpellHit(sim *Simulation, cast DirectCastAction, result *DirectCastDamageResult) {
	for _, id := range at.activeAuraIDs {
		if at.auras[id].OnSpellHit != nil {
			at.auras[id].OnSpellHit(sim, cast, result)
		}
	}
}

// Stored value is the time at which the ICD will be off CD
type InternalCD time.Duration

func (icd InternalCD) IsOnCD(sim *Simulation) bool {
	return time.Duration(icd) > sim.CurrentTime
}

func NewICD() InternalCD {
	return InternalCD(0)
}

// Helper for the common case of adding an Aura that gives a temporary stat boost.
func (character *Character) AddAuraWithTemporaryStats(sim *Simulation, auraID AuraID, auraName string, stat stats.Stat, amount float64, duration time.Duration) {
	if sim.Log != nil {
		sim.Log(" +%0.0f %s from %s\n", amount, stat.StatName(), auraName)
	}
	character.AddStat(stat, amount)

	character.AddAura(sim, Aura{
		ID:      auraID,
		Name:    auraName,
		Expires: sim.CurrentTime + duration,
		OnExpire: func(sim *Simulation) {
			if sim.Log != nil {
				sim.Log(" -%0.0f %s from %s\n", amount, stat.StatName(), auraName)
			}
			character.AddStat(stat, -amount)
		},
	})
}
