package core

import (
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

const NeverExpires = time.Duration(math.MaxInt64)

type AuraID int32

// Reserve the default value so no aura uses it.
const UnknownAuraID = AuraID(0)

var numAuraIDs = 1

func NewAuraID() AuraID {
	newAuraID := AuraID(numAuraIDs)
	numAuraIDs++
	return newAuraID
}

// Offsensive trinkets put each other on CD, so they can all share 1 aura ID
var OffensiveTrinketActiveAuraID = NewAuraID()

// Defensive trinkets put each other on CD, so they can all share 1 aura ID
var DefensiveTrinketActiveAuraID = NewAuraID()

// Reserve the default value so no aura uses it.
const UnknownDebuffID = AuraID(0)

var numDebuffIDs = 1

func NewDebuffID() AuraID {
	newDebuffID := AuraID(numDebuffIDs)
	numDebuffIDs++
	return newDebuffID
}

type CooldownID int32

// Reserve the default value so no cooldown uses it.
const UnknownCooldownID = AuraID(0)

var numCooldownIDs = 1

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

type OnExpire func(sim *Simulation)

type Aura struct {
	ID          AuraID
	Name        string        // Label used for logging.
	Expires     time.Duration // time at which aura will be removed
	activeIndex int32         // Position of this aura's index in the sim.activeAuraIDs array

	onCastIndex           int32 // Position of this aura's index in the sim.onCastIDs array
	onCastCompleteIndex   int32 // Position of this aura's index in the sim.onCastCompleteIDs array
	onBeforeSpellHitIndex int32 // Position of this aura's index in the sim.onBeforeSpellHitIDs array
	onSpellHitIndex       int32 // Position of this aura's index in the sim.onSpellHitIDs array
	onSpellMissIndex      int32 // Position of this aura's index in the sim.onSpellMissIDs array

	// The number of stacks, or charges, of this aura. If this aura doesn't care
	// about charges, is just 0.
	Stacks int32

	// Invoked at creation time for a spell cast.
	OnCast OnCast

	// Invoked when a spell cast completes casting, before results are calculated.
	OnCastComplete OnCastComplete

	// Invoked before a spell lands, but after the target is selected.
	OnBeforeSpellHit OnBeforeSpellHit

	// Invoked when a spell hits, after results are calculated. Results can be modified by changing
	// properties of result.
	OnSpellHit OnSpellHit

	// Invoked when a spell is fully resisted.
	OnSpellMiss OnSpellMiss

	// Invoked when this Aura expires.
	OnExpire OnExpire
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

	// Set to true if this aura tracker is tracking target debuffs, instead of player buffs.
	useDebuffIDs bool

	// Whether finalize() has been called for this object.
	finalized bool

	// Maps MagicIDs to sim duration at which CD is done. Using array for perf.
	cooldowns []time.Duration

	// Maps MagicIDs to aura for that ID. Using array for perf.
	auras []Aura

	// IDs of Auras that are active, in no particular order
	activeAuraIDs []AuraID

	// IDs of Auras that have a non-nil OnCast function set.
	onCastIDs []AuraID

	// IDs of Auras that have a non-nil OnCastComplete function set.
	onCastCompleteIDs []AuraID

	// IDs of Auras that have a non-nil OnBeforeSpellHit function set.
	onBeforeSpellHitIDs []AuraID

	// IDs of Auras that have a non-nil OnSpellHit function set.
	onSpellHitIDs []AuraID

	// IDs of Auras that have a non-nil OnSpellMiss function set.
	onSpellMissIDs []AuraID
}

func newAuraTracker(useDebuffIDs bool) auraTracker {
	return auraTracker{
		permanentAuras:      []PermanentAura{},
		activeAuraIDs:       make([]AuraID, 0, 16),
		onCastIDs:           make([]AuraID, 0, 16),
		onCastCompleteIDs:   make([]AuraID, 0, 16),
		onBeforeSpellHitIDs: make([]AuraID, 0, 16),
		onSpellHitIDs:       make([]AuraID, 0, 16),
		onSpellMissIDs:      make([]AuraID, 0, 16),
		useDebuffIDs:        useDebuffIDs,
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
	if at.useDebuffIDs {
		at.auras = make([]Aura, numDebuffIDs)
	} else {
		at.auras = make([]Aura, numAuraIDs)
	}

	at.cooldowns = make([]time.Duration, numCooldownIDs)
	at.activeAuraIDs = at.activeAuraIDs[:0]
	at.onCastIDs = at.onCastIDs[:0]
	at.onCastCompleteIDs = at.onCastCompleteIDs[:0]
	at.onBeforeSpellHitIDs = at.onBeforeSpellHitIDs[:0]
	at.onSpellHitIDs = at.onSpellHitIDs[:0]
	at.onSpellMissIDs = at.onSpellMissIDs[:0]

	for _, permAura := range at.permanentAuras {
		aura := permAura(sim)
		aura.Expires = NeverExpires
		at.AddAura(sim, aura)
	}
}

func (at *auraTracker) advance(sim *Simulation) {
	// Go in reverse order so we can safely delete while looping
	for i := len(at.activeAuraIDs) - 1; i >= 0; i-- {
		id := at.activeAuraIDs[i]
		if at.auras[id].Expires != 0 && at.auras[id].Expires <= sim.CurrentTime {
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

	if newAura.OnCast != nil {
		at.auras[newAura.ID].onCastIndex = int32(len(at.onCastIDs))
		at.onCastIDs = append(at.onCastIDs, newAura.ID)
	}

	if newAura.OnCastComplete != nil {
		at.auras[newAura.ID].onCastCompleteIndex = int32(len(at.onCastCompleteIDs))
		at.onCastCompleteIDs = append(at.onCastCompleteIDs, newAura.ID)
	}

	if newAura.OnBeforeSpellHit != nil {
		at.auras[newAura.ID].onBeforeSpellHitIndex = int32(len(at.onBeforeSpellHitIDs))
		at.onBeforeSpellHitIDs = append(at.onBeforeSpellHitIDs, newAura.ID)
	}

	if newAura.OnSpellHit != nil {
		at.auras[newAura.ID].onSpellHitIndex = int32(len(at.onSpellHitIDs))
		at.onSpellHitIDs = append(at.onSpellHitIDs, newAura.ID)
	}

	if newAura.OnSpellMiss != nil {
		at.auras[newAura.ID].onSpellMissIndex = int32(len(at.onSpellMissIDs))
		at.onSpellMissIDs = append(at.onSpellMissIDs, newAura.ID)
	}

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
	at.activeAuraIDs = removeBySwappingToBack(at.activeAuraIDs, removeActiveIndex)
	if removeActiveIndex < int32(len(at.activeAuraIDs)) {
		at.auras[at.activeAuraIDs[removeActiveIndex]].activeIndex = removeActiveIndex
	}

	if at.auras[id].OnCast != nil {
		removeOnCastIndex := at.auras[id].onCastIndex
		at.onCastIDs = removeBySwappingToBack(at.onCastIDs, removeOnCastIndex)
		if removeOnCastIndex < int32(len(at.onCastIDs)) {
			at.auras[at.onCastIDs[removeOnCastIndex]].onCastIndex = removeOnCastIndex
		}
	}

	if at.auras[id].OnCastComplete != nil {
		removeOnCastCompleteIndex := at.auras[id].onCastCompleteIndex
		at.onCastCompleteIDs = removeBySwappingToBack(at.onCastCompleteIDs, removeOnCastCompleteIndex)
		if removeOnCastCompleteIndex < int32(len(at.onCastCompleteIDs)) {
			at.auras[at.onCastCompleteIDs[removeOnCastCompleteIndex]].onCastCompleteIndex = removeOnCastCompleteIndex
		}
	}

	if at.auras[id].OnBeforeSpellHit != nil {
		removeOnBeforeSpellHitIndex := at.auras[id].onBeforeSpellHitIndex
		at.onBeforeSpellHitIDs = removeBySwappingToBack(at.onBeforeSpellHitIDs, removeOnBeforeSpellHitIndex)
		if removeOnBeforeSpellHitIndex < int32(len(at.onBeforeSpellHitIDs)) {
			at.auras[at.onBeforeSpellHitIDs[removeOnBeforeSpellHitIndex]].onBeforeSpellHitIndex = removeOnBeforeSpellHitIndex
		}
	}

	if at.auras[id].OnSpellHit != nil {
		removeOnSpellHitIndex := at.auras[id].onSpellHitIndex
		at.onSpellHitIDs = removeBySwappingToBack(at.onSpellHitIDs, removeOnSpellHitIndex)
		if removeOnSpellHitIndex < int32(len(at.onSpellHitIDs)) {
			at.auras[at.onSpellHitIDs[removeOnSpellHitIndex]].onSpellHitIndex = removeOnSpellHitIndex
		}
	}

	if at.auras[id].OnSpellMiss != nil {
		removeOnSpellMissIndex := at.auras[id].onSpellMissIndex
		at.onSpellMissIDs = removeBySwappingToBack(at.onSpellMissIDs, removeOnSpellMissIndex)
		if removeOnSpellMissIndex < int32(len(at.onSpellMissIDs)) {
			at.auras[at.onSpellMissIDs[removeOnSpellMissIndex]].onSpellMissIndex = removeOnSpellMissIndex
		}
	}

	at.auras[id] = Aura{}
}

// Constant-time removal from slice by swapping with the last element before removing.
func removeBySwappingToBack(arr []AuraID, removeIdx int32) []AuraID {
	arr[removeIdx] = arr[len(arr)-1]
	return arr[:len(arr)-1]
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
func (at *auraTracker) OnCast(sim *Simulation, cast *Cast) {
	for _, id := range at.onCastIDs {
		at.auras[id].OnCast(sim, cast)
	}
}

// Invokes the OnCastComplete event for all tracked Auras.
func (at *auraTracker) OnCastComplete(sim *Simulation, cast *Cast) {
	for _, id := range at.onCastCompleteIDs {
		at.auras[id].OnCastComplete(sim, cast)
	}
}

// Invokes the OnBeforeSpellHit event for all tracked Auras.
func (at *auraTracker) OnBeforeSpellHit(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
	for _, id := range at.onBeforeSpellHitIDs {
		at.auras[id].OnBeforeSpellHit(sim, spellCast, spellEffect)
	}
}

// Invokes the OnSpellMiss event for all tracked Auras.
func (at *auraTracker) OnSpellMiss(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
	for _, id := range at.onSpellMissIDs {
		at.auras[id].OnSpellMiss(sim, spellCast, spellEffect)
	}
}

// Invokes the OnSpellHit event for all tracked Auras.
func (at *auraTracker) OnSpellHit(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
	for _, id := range at.onSpellHitIDs {
		at.auras[id].OnSpellHit(sim, spellCast, spellEffect)
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
