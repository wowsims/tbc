package core

import (
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
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
	SpellID     int32         // In-game spell ID. If set, metrics will be tracked for this aura using this ID.
	Name        string        // Label used for logging.
	Expires     time.Duration // Time at which aura will be removed.
	activeIndex int32         // Position of this aura's index in the sim.activeAuraIDs array.

	startTime time.Duration // Time at which the aura was applied.

	onCastIndex                 int32 // Position of this aura's index in the sim.onCastIDs array.
	onCastCompleteIndex         int32 // Position of this aura's index in the sim.onCastCompleteIDs array.
	onBeforeSpellHitIndex       int32 // Position of this aura's index in the sim.onBeforeSpellHitIDs array.
	onSpellHitIndex             int32 // Position of this aura's index in the sim.onSpellHitIDs array.
	onSpellMissIndex            int32 // Position of this aura's index in the sim.onSpellMissIDs array.
	onBeforePeriodicDamageIndex int32 // Position of this aura's index in the sim.onBeforePeriodicDamageIDs array.
	onPeriodicDamageIndex       int32 // Position of this aura's index in the sim.onPeriodicDamageIDs array.
	onBeforeSwingHitIndex       int32 // Position of this aura's index in the sim.onBeforeSwingHit array.
	OnMeleeAttackIndex          int32 // Position of this aura's index in the sim.OnMeleeAttack array.
	OnBeforeMeleeIndex          int32 // Position of this aura's index in the sim.OnBeforeMelee array.

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

	// Invoked when a dot tick occurs, before damage is calculated.
	OnBeforePeriodicDamage OnBeforePeriodicDamage

	// Invoked when a dot tick occurs, after damage is calculated.
	OnPeriodicDamage OnPeriodicDamage

	// Invoked before an auto attack swing happens.
	OnBeforeSwingHit OnBeforeSwingHit

	// Invoked after a melee hit has occured (could be auto or skill).
	OnMeleeAttack OnMeleeAttack

	// Invoked before melee of any kind (swing or ability)
	OnBeforeMelee OnBeforeMelee
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

	// Callback to format aura-related logs.
	logFn func(string, ...interface{})

	// Set to true if this aura tracker is tracking target debuffs, instead of player buffs.
	useDebuffIDs bool

	// Whether finalize() has been called for this object.
	finalized bool

	// Maps MagicIDs to sim duration at which CD is done. Using array for perf.
	cooldowns []time.Duration

	// Maps MagicIDs to aura for that ID. Using array for perf.
	auras []Aura

	// IDs of Auras that are active, in no particular order.
	activeAuraIDs []AuraID

	// IDs of Auras that have a non-nil XXX function set.
	onCastIDs                 []AuraID
	onCastCompleteIDs         []AuraID
	onBeforeSpellHitIDs       []AuraID
	onSpellHitIDs             []AuraID
	onSpellMissIDs            []AuraID
	onBeforePeriodicDamageIDs []AuraID
	onPeriodicDamageIDs       []AuraID
	onBeforeSwingHitIDs       []AuraID
	onMeleeAttackIDs          []AuraID
	onBeforeMeleeIDs          []AuraID

	// Metrics for each aura.
	metrics []AuraMetrics
}

func newAuraTracker(useDebuffIDs bool) auraTracker {
	numAura := numAuraIDs
	if useDebuffIDs {
		numAura = numDebuffIDs
	}
	return auraTracker{
		permanentAuras:            []PermanentAura{},
		activeAuraIDs:             make([]AuraID, 0, 16),
		onCastIDs:                 make([]AuraID, 0, 16),
		onCastCompleteIDs:         make([]AuraID, 0, 16),
		onBeforeSpellHitIDs:       make([]AuraID, 0, 16),
		onSpellHitIDs:             make([]AuraID, 0, 16),
		onSpellMissIDs:            make([]AuraID, 0, 16),
		onBeforePeriodicDamageIDs: make([]AuraID, 0, 16),
		onPeriodicDamageIDs:       make([]AuraID, 0, 16),
		onBeforeSwingHitIDs:       make([]AuraID, 0, 16),
		onMeleeAttackIDs:          make([]AuraID, 0, 16),
		onBeforeMeleeIDs:          make([]AuraID, 0, 16),
		auras:                     make([]Aura, numAura),
		cooldowns:                 make([]time.Duration, numCooldownIDs),
		useDebuffIDs:              useDebuffIDs,
		metrics:                   make([]AuraMetrics, numAura),
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
	at.onBeforePeriodicDamageIDs = at.onBeforePeriodicDamageIDs[:0]
	at.onPeriodicDamageIDs = at.onPeriodicDamageIDs[:0]
	at.onBeforeSwingHitIDs = at.onBeforeSwingHitIDs[:0]
	at.onMeleeAttackIDs = at.onMeleeAttackIDs[:0]
	at.onBeforeMeleeIDs = at.onBeforeMeleeIDs[:0]

	for _, permAura := range at.permanentAuras {
		aura := permAura(sim)
		aura.Expires = NeverExpires
		at.AddAura(sim, aura)
	}
}

func (at *auraTracker) advance(sim *Simulation) {
	for _, id := range at.activeAuraIDs {
		if at.auras[id].Expires != 0 && at.auras[id].Expires <= sim.CurrentTime {
			at.RemoveAura(sim, id)
		}
	}
}

func (at *auraTracker) doneIteration(simDuration time.Duration) {
	// Add metrics for any auras that are still active.
	for _, aura := range at.auras {
		if aura.SpellID != 0 {
			at.AddAuraUptime(aura.ID, aura.SpellID, simDuration-aura.startTime)
		}
	}

	for i, _ := range at.metrics {
		auraMetric := &at.metrics[i]
		auraMetric.doneIteration()
	}
}

// ReplaceAura is like AddAura but an existing aura will not be removed/readded.
// This means that 'OnExpire' will not fire off on the old aura.
func (at *auraTracker) ReplaceAura(sim *Simulation, newAura Aura) {
	if at.HasAura(newAura.ID) {
		old := at.auras[newAura.ID]

		// private cached state has to be copied over
		newAura.activeIndex = old.activeIndex
		newAura.onCastIndex = old.onCastIndex
		newAura.onCastCompleteIndex = old.onCastCompleteIndex
		newAura.onBeforeSpellHitIndex = old.onBeforeSpellHitIndex
		newAura.onSpellHitIndex = old.onSpellHitIndex
		newAura.onSpellMissIndex = old.onSpellMissIndex
		newAura.onBeforePeriodicDamageIndex = old.onBeforePeriodicDamageIndex
		newAura.onPeriodicDamageIndex = old.onPeriodicDamageIndex
		newAura.onBeforeSwingHitIndex = old.onBeforeSwingHitIndex
		newAura.OnMeleeAttackIndex = old.OnMeleeAttackIndex
		newAura.OnBeforeMeleeIndex = old.OnBeforeMeleeIndex
		newAura.startTime = old.startTime

		at.auras[newAura.ID] = newAura
		return
	}

	at.AddAura(sim, newAura)
}

// Adds a new aura to the simulation. If an aura with the same ID already
// exists it will be replaced with the new one.
func (at *auraTracker) AddAura(sim *Simulation, newAura Aura) {
	if at.HasAura(newAura.ID) {
		at.RemoveAura(sim, newAura.ID)
	}

	newAura.startTime = sim.CurrentTime

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

	if newAura.OnBeforePeriodicDamage != nil {
		at.auras[newAura.ID].onBeforePeriodicDamageIndex = int32(len(at.onBeforePeriodicDamageIDs))
		at.onBeforePeriodicDamageIDs = append(at.onBeforePeriodicDamageIDs, newAura.ID)
	}

	if newAura.OnPeriodicDamage != nil {
		at.auras[newAura.ID].onPeriodicDamageIndex = int32(len(at.onPeriodicDamageIDs))
		at.onPeriodicDamageIDs = append(at.onPeriodicDamageIDs, newAura.ID)
	}

	if newAura.OnBeforeSwingHit != nil {
		at.auras[newAura.ID].onBeforeSwingHitIndex = int32(len(at.onBeforeSwingHitIDs))
		at.onBeforeSwingHitIDs = append(at.onBeforeSwingHitIDs, newAura.ID)
	}

	if newAura.OnMeleeAttack != nil {
		at.auras[newAura.ID].OnMeleeAttackIndex = int32(len(at.onMeleeAttackIDs))
		at.onMeleeAttackIDs = append(at.onMeleeAttackIDs, newAura.ID)
	}

	if newAura.OnBeforeMelee != nil {
		at.auras[newAura.ID].OnBeforeMeleeIndex = int32(len(at.onBeforeMeleeIDs))
		at.onBeforeMeleeIDs = append(at.onBeforeMeleeIDs, newAura.ID)
	}

	if sim.Log != nil {
		at.logFn("Aura gained: %s", newAura.Name)
	}
}

// Remove an aura by its ID
func (at *auraTracker) RemoveAura(sim *Simulation, id AuraID) {
	if at.auras[id].OnExpire != nil {
		at.auras[id].OnExpire(sim)
	}

	if at.auras[id].SpellID != 0 {
		at.AddAuraUptime(id, at.auras[id].SpellID, sim.CurrentTime-at.auras[id].startTime)
	}

	if sim.Log != nil {
		at.logFn("Aura faded: %s", at.auras[id].Name)
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

	if at.auras[id].OnBeforePeriodicDamage != nil {
		removeOnBeforePeriodicDamage := at.auras[id].onBeforePeriodicDamageIndex
		at.onBeforePeriodicDamageIDs = removeBySwappingToBack(at.onBeforePeriodicDamageIDs, removeOnBeforePeriodicDamage)
		if removeOnBeforePeriodicDamage < int32(len(at.onBeforePeriodicDamageIDs)) {
			at.auras[at.onBeforePeriodicDamageIDs[removeOnBeforePeriodicDamage]].onBeforePeriodicDamageIndex = removeOnBeforePeriodicDamage
		}
	}

	if at.auras[id].OnPeriodicDamage != nil {
		removeOnPeriodicDamage := at.auras[id].onPeriodicDamageIndex
		at.onPeriodicDamageIDs = removeBySwappingToBack(at.onPeriodicDamageIDs, removeOnPeriodicDamage)
		if removeOnPeriodicDamage < int32(len(at.onPeriodicDamageIDs)) {
			at.auras[at.onPeriodicDamageIDs[removeOnPeriodicDamage]].onPeriodicDamageIndex = removeOnPeriodicDamage
		}
	}

	if at.auras[id].OnBeforeSwingHit != nil {
		removeOnBeforeSwingHit := at.auras[id].onBeforeSwingHitIndex
		at.onBeforeSwingHitIDs = removeBySwappingToBack(at.onBeforeSwingHitIDs, removeOnBeforeSwingHit)
		if removeOnBeforeSwingHit < int32(len(at.onBeforeSwingHitIDs)) {
			at.auras[at.onBeforeSwingHitIDs[removeOnBeforeSwingHit]].onBeforeSwingHitIndex = removeOnBeforeSwingHit
		}
	}
	if at.auras[id].OnMeleeAttack != nil {
		removeOnMeleeAttack := at.auras[id].OnMeleeAttackIndex
		at.onMeleeAttackIDs = removeBySwappingToBack(at.onMeleeAttackIDs, removeOnMeleeAttack)
		if removeOnMeleeAttack < int32(len(at.onMeleeAttackIDs)) {
			at.auras[at.onMeleeAttackIDs[removeOnMeleeAttack]].OnMeleeAttackIndex = removeOnMeleeAttack
		}
	}
	if at.auras[id].OnBeforeMelee != nil {
		removeOnBeforeMelee := at.auras[id].OnBeforeMeleeIndex
		at.onBeforeMeleeIDs = removeBySwappingToBack(at.onBeforeMeleeIDs, removeOnBeforeMelee)
		if removeOnBeforeMelee < int32(len(at.onBeforeMeleeIDs)) {
			at.auras[at.onBeforeMeleeIDs[removeOnBeforeMelee]].OnBeforeMeleeIndex = removeOnBeforeMelee
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

func (at *auraTracker) NumStacks(id AuraID) int32 {
	if at.HasAura(id) {
		return at.auras[id].Stacks
	} else {
		return 0
	}
}

func (at *auraTracker) RemainingAuraDuration(sim *Simulation, id AuraID) time.Duration {
	if at.HasAura(id) {
		expires := at.auras[id].Expires
		if expires == NeverExpires {
			return NeverExpires
		} else {
			return expires - sim.CurrentTime
		}
	} else {
		return 0
	}
}

func (at *auraTracker) IsOnCD(id CooldownID, currentTime time.Duration) bool {
	return at.cooldowns[id] > currentTime
}

func (at *auraTracker) GetRemainingCD(id CooldownID, currentTime time.Duration) time.Duration {
	return MaxDuration(0, at.cooldowns[id]-currentTime)
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

// Invokes the OnBeforePeriodicDamage
//   As a debuff when target is being hit by dot.
//   As a buff when caster's dots are ticking.
func (at *auraTracker) OnBeforePeriodicDamage(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage *float64) {
	for _, id := range at.onBeforePeriodicDamageIDs {
		at.auras[id].OnBeforePeriodicDamage(sim, spellCast, spellEffect, tickDamage)
	}
}

// Invokes the OnPeriodicDamage
//   As a debuff when target is being hit by dot.
//   As a buff when caster's dots are ticking.
func (at *auraTracker) OnPeriodicDamage(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage float64) {
	for _, id := range at.onPeriodicDamageIDs {
		at.auras[id].OnPeriodicDamage(sim, spellCast, spellEffect, tickDamage)
	}
}

func (at *auraTracker) OnBeforeSwingHit(sim *Simulation, isOH bool) bool {
	doSwing := true
	for _, id := range at.onBeforeSwingHitIDs {
		doSwing = at.auras[id].OnBeforeSwingHit(sim, isOH) && doSwing
	}
	return doSwing
}

func (at *auraTracker) OnMeleeAttack(sim *Simulation, target *Target, result MeleeHitType, ability *ActiveMeleeAbility, isOH bool) {
	for _, id := range at.onMeleeAttackIDs {
		at.auras[id].OnMeleeAttack(sim, target, result, ability, isOH)
	}
}

func (at *auraTracker) OnBeforeMelee(sim *Simulation, ability *ActiveMeleeAbility, isOH bool) {
	for _, id := range at.onBeforeMeleeIDs {
		at.auras[id].OnBeforeMelee(sim, ability, isOH)
	}
}

func (at *auraTracker) AddAuraUptime(auraID AuraID, spellID int32, uptime time.Duration) {
	metrics := &at.metrics[auraID]

	metrics.ID = spellID
	metrics.Uptime += uptime
}

func (at *auraTracker) GetMetricsProto(numIterations int32) []*proto.AuraMetrics {
	metrics := make([]*proto.AuraMetrics, 0, len(at.metrics))

	for _, auraMetric := range at.metrics {
		if auraMetric.ID != 0 {
			metrics = append(metrics, auraMetric.ToProto(numIterations))
		}
	}

	return metrics
}

// Stored value is the time at which the ICD will be off CD
type InternalCD time.Duration

func (icd InternalCD) IsOnCD(sim *Simulation) bool {
	return time.Duration(icd) > sim.CurrentTime
}

func (icd InternalCD) GetRemainingCD(sim *Simulation) time.Duration {
	return MaxDuration(0, time.Duration(icd)-sim.CurrentTime)
}

func NewICD() InternalCD {
	return InternalCD(0)
}

// Helper for the common case of adding an Aura that gives a temporary stat boost.
func (character *Character) AddAuraWithTemporaryStats(sim *Simulation, auraID AuraID, spellID int32, auraName string, stat stats.Stat, amount float64, duration time.Duration) {
	if sim.Log != nil {
		character.Log(sim, "Gained %0.0f %s from %s.", amount, stat.StatName(), auraName)
	}
	if stat == stats.MeleeHaste {
		character.AddMeleeHaste(sim, amount)
	} else {
		character.AddStat(stat, amount)
	}

	character.AddAura(sim, Aura{
		ID:      auraID,
		SpellID: spellID,
		Name:    auraName,
		Expires: sim.CurrentTime + duration,
		OnExpire: func(sim *Simulation) {
			if sim.Log != nil {
				character.Log(sim, "Lost %0.0f %s from fading %s.", amount, stat.StatName(), auraName)
			}
			if stat == stats.MeleeHaste {
				character.AddMeleeHaste(sim, -amount)
			} else {
				character.AddStat(stat, -amount)
			}
		},
	})
}
