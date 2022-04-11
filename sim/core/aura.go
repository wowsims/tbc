package core

import (
	"fmt"
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const NeverExpires = time.Duration(math.MaxInt64)

type CooldownID int32

// Reserve the default value so no cooldown uses it.
const UnknownCooldownID = CooldownID(0)

var numCooldownIDs = 1

func NewCooldownID() CooldownID {
	newCooldownID := CooldownID(numCooldownIDs)
	numCooldownIDs++
	return newCooldownID
}

var GCDCooldownID = NewCooldownID()
var OffensiveTrinketSharedCooldownID = NewCooldownID()
var DefensiveTrinketSharedCooldownID = NewCooldownID()

type OnGain func(aura *Aura, sim *Simulation)
type OnExpire func(aura *Aura, sim *Simulation)
type OnStacksChange func(aura *Aura, sim *Simulation, oldStacks int32, newStacks int32)

const Inactive = -1

// Aura lifecycle:
//
// myAura := unit.RegisterAura(myAuraConfig)
// myAura.Activate(sim)
// myAura.SetStacks(sim, 3)
// myAura.Refresh(sim)
// myAura.Deactivate(sim)
type Aura struct {
	// String label for this Aura. Gauranteed to be unique among the Auras for a single Unit.
	Label string

	// For easily grouping auras.
	Tag string

	ActionID ActionID // If set, metrics will be tracked for this aura.

	Duration time.Duration // Duration of aura, upon being applied.

	startTime time.Duration // Time at which the aura was applied.
	expires   time.Duration // Time at which aura will be removed.

	// The unit this aura is attached to.
	Unit *Unit

	active                   bool
	activeIndex              int32 // Position of this aura's index in the activeAuras array.
	onCastCompleteIndex      int32 // Position of this aura's index in the onCastCompleteAuras array.
	onSpellCastCompleteIndex int32 // Position of this aura's index in the onCastCompleteAuras array.
	onSpellHitIndex          int32 // Position of this aura's index in the onSpellHitAuras array.
	onPeriodicDamageIndex    int32 // Position of this aura's index in the onPeriodicDamageAuras array.

	// The number of stacks, or charges, of this aura. If this aura doesn't care
	// about charges, is just 0.
	stacks    int32
	MaxStacks int32

	// If nonzero, activation of this aura will deactivate other auras with the
	// same Tag and equal or lower Priority.
	Priority float64

	// Invoked when a spell cast completes casting, before results are calculated.
	OnCastComplete      OnCastComplete
	OnSpellCastComplete OnSpellCastComplete

	// Invoked when a spell hits, after results are calculated. Results can be modified by changing
	// properties of result.
	OnSpellHit OnSpellHit

	// Invoked when this Aura is added/remvoed. Neither is invoked on refresh.
	OnGain   OnGain
	OnExpire OnExpire

	// Invoked when the number of stacks of this aura changes.
	OnStacksChange OnStacksChange

	// Invoked when a dot tick occurs, after damage is calculated.
	OnPeriodicDamage OnPeriodicDamage

	// Metrics for this aura.
	metrics AuraMetrics
}

func (aura *Aura) IsActive() bool {
	return aura.active
}

func (aura *Aura) Refresh(sim *Simulation) {
	if aura.Duration == NeverExpires {
		aura.expires = NeverExpires
	} else {
		aura.expires = sim.CurrentTime + aura.Duration
		aura.Unit.minExpires = 0
	}
}

func (aura *Aura) GetStacks() int32 {
	return aura.stacks
}

func (aura *Aura) SetStacks(sim *Simulation, newStacks int32) {
	if aura.MaxStacks == 0 {
		panic("MaxStacks required to set Aura stacks: " + aura.Label)
	}
	oldStacks := aura.stacks
	newStacks = MinInt32(newStacks, aura.MaxStacks)

	if oldStacks == newStacks {
		return
	}

	if sim.Log != nil {
		aura.Unit.Log(sim, "%s stacks: %d --> %d", aura.ActionID, oldStacks, newStacks)
	}
	aura.stacks = newStacks
	if aura.OnStacksChange != nil {
		aura.OnStacksChange(aura, sim, oldStacks, newStacks)
	}
	if aura.stacks == 0 {
		aura.Deactivate(sim)
	}
}
func (aura *Aura) AddStack(sim *Simulation) {
	aura.SetStacks(sim, aura.stacks+1)
}
func (aura *Aura) RemoveStack(sim *Simulation) {
	aura.SetStacks(sim, aura.stacks-1)
}

func (aura *Aura) UpdateExpires(newExpires time.Duration) {
	aura.expires = newExpires
}

func (aura *Aura) RemainingDuration(sim *Simulation) time.Duration {
	if aura.expires == NeverExpires {
		return NeverExpires
	} else {
		return aura.expires - sim.CurrentTime
	}
}

func (aura *Aura) ExpiresAt() time.Duration {
	return aura.expires
}

type AuraFactory func(*Simulation) *Aura

// Callbacks for doing something on finalize/reset.
type FinalizeEffect func()
type ResetEffect func(*Simulation)

// Convenience for some common Aura behavior.
type PermanentAura struct {
	AuraFactory AuraFactory

	// By default, permanent auras have their expiration overwritten to never expire.
	// This option disables that behavior, creating an aura which is applied at the
	// beginning of every iteration but expires after a period of time. This is
	// used for some snapshotting effects like Warrior battle shout.
	RespectDuration bool

	// Multiplies uptime for the aura metrics of this aura. This is for buffs coded
	// as permanent but which are actually averaged versions of the real buff.
	UptimeMultiplier float64

	// The aura created by AuraFactory.
	aura *Aura
}

// auraTracker is a centralized implementation of CD and Aura tracking.
//  This is used by all Units.
type auraTracker struct {
	// Effects to invoke when the Unit is finalized.
	finalizeEffects []FinalizeEffect

	// Effects to invoke on every sim reset.
	resetEffects []ResetEffect

	// Auras that never expire and should always be active.
	// These are automatically applied on each Sim reset.
	permanentAuras []PermanentAura

	// Whether finalize() has been called for this object.
	finalized bool

	// Maps MagicIDs to sim duration at which CD is done. Using array for perf.
	cooldowns []time.Duration

	// All registered auras, both active and inactive.
	auras []*Aura

	aurasByTag map[string][]*Aura

	// IDs of Auras that may expire and are currently active, in no particular order.
	activeAuras []*Aura

	// caches the minimum expires time of all active auras; reset to 0 on Activate(), Deactivate(), and Refresh()
	minExpires time.Duration

	// Auras that have a non-nil XXX function set and are currently active.
	onCastCompleteAuras      []*Aura
	onSpellCastCompleteAuras []*Aura
	onSpellHitAuras          []*Aura
	onPeriodicDamageAuras    []*Aura
}

func newAuraTracker() auraTracker {
	return auraTracker{
		finalizeEffects:          []FinalizeEffect{},
		resetEffects:             []ResetEffect{},
		permanentAuras:           []PermanentAura{},
		activeAuras:              make([]*Aura, 0, 16),
		onCastCompleteAuras:      make([]*Aura, 0, 16),
		onSpellCastCompleteAuras: make([]*Aura, 0, 16),
		onSpellHitAuras:          make([]*Aura, 0, 16),
		onPeriodicDamageAuras:    make([]*Aura, 0, 16),
		auras:                    make([]*Aura, 0, 16),
		aurasByTag:               make(map[string][]*Aura),
		cooldowns:                make([]time.Duration, numCooldownIDs),
	}
}

func (at *auraTracker) GetAura(label string) *Aura {
	for _, aura := range at.auras {
		if aura.Label == label {
			return aura
		}
	}
	return nil
}
func (at *auraTracker) HasActiveAura(label string) bool {
	aura := at.GetAura(label)
	return aura != nil && aura.IsActive()
}

func (at *auraTracker) registerAura(unit *Unit, aura *Aura) *Aura {
	if unit == nil {
		panic("Aura unit is required!")
	}
	if aura.Label == "" {
		panic("Aura label is required!")
	}
	if aura.Priority != 0 && aura.Tag == "" {
		panic("Aura.Priority requires Aura.Tag also be set")
	}
	if at.GetAura(aura.Label) != nil {
		panic(fmt.Sprintf("Aura %s already registered!", aura.Label))
	}
	if len(at.auras) > 100 {
		panic(fmt.Sprintf("Over 100 registered auras when registering %s! There is probably an aura being registered every iteration.", aura.Label))
	}

	aura.Unit = unit
	aura.metrics.ID = aura.ActionID

	at.auras = append(at.auras, aura)
	if aura.Tag != "" {
		at.aurasByTag[aura.Tag] = append(at.aurasByTag[aura.Tag], aura)
	}

	return aura
}
func (unit *Unit) RegisterAura(aura *Aura) *Aura {
	return unit.auraTracker.registerAura(unit, aura)
}

func (unit *Unit) GetOrRegisterAura(aura *Aura) *Aura {
	curAura := unit.GetAura(aura.Label)
	if curAura == nil {
		return unit.RegisterAura(aura)
	} else {
		curAura.OnCastComplete = aura.OnCastComplete
		curAura.OnSpellCastComplete = aura.OnSpellCastComplete
		curAura.OnSpellHit = aura.OnSpellHit
		curAura.OnPeriodicDamage = aura.OnPeriodicDamage
		return curAura
	}
}

func (at *auraTracker) GetAurasWithTag(tag string) []*Aura {
	return at.aurasByTag[tag]
}

func (at *auraTracker) HasActiveAuraWithTag(tag string) bool {
	for _, aura := range at.aurasByTag[tag] {
		if aura.active {
			return true
		}
	}
	return false
}

// Registers a callback to this Character which will be invoked on
// every Sim reset.
func (at *auraTracker) RegisterFinalizeEffect(finalizeEffect FinalizeEffect) {
	if at.finalized {
		panic("Finalize effects may not be added once finalized!")
	}

	at.finalizeEffects = append(at.finalizeEffects, finalizeEffect)
}

// Registers a callback to this Character which will be invoked on
// every Sim reset.
func (at *auraTracker) RegisterResetEffect(resetEffect ResetEffect) {
	if at.finalized {
		panic("Reset effects may not be added once finalized!")
	}

	at.resetEffects = append(at.resetEffects, resetEffect)
}

// Registers a permanent aura to this Character which will be re-applied on
// every Sim reset.
func (at *auraTracker) AddPermanentAuraWithOptions(permAura PermanentAura) {
	if at.finalized {
		panic("Permanent auras may not be added once finalized!")
	}

	at.permanentAuras = append(at.permanentAuras, permAura)
}
func (at *auraTracker) AddPermanentAura(factory AuraFactory) {
	at.AddPermanentAuraWithOptions(PermanentAura{
		AuraFactory: factory,
	})
}

func (at *auraTracker) finalize() {
	if at.finalized {
		return
	}
	at.finalized = true

	for _, finalizeEffect := range at.finalizeEffects {
		finalizeEffect()
	}
}

func (at *auraTracker) reset(sim *Simulation) {
	at.cooldowns = make([]time.Duration, numCooldownIDs)
	at.activeAuras = at.activeAuras[:0]
	at.onCastCompleteAuras = at.onCastCompleteAuras[:0]
	at.onSpellCastCompleteAuras = at.onSpellCastCompleteAuras[:0]
	at.onSpellHitAuras = at.onSpellHitAuras[:0]
	at.onPeriodicDamageAuras = at.onPeriodicDamageAuras[:0]

	for _, aura := range at.auras {
		if aura.IsActive() {
			panic("Active aura during reset: " + aura.Label)
		}
		if aura.stacks != 0 {
			panic("Aura nonzero stacks during reset: " + aura.Label)
		}
		aura.metrics.reset()
	}

	for _, resetEffect := range at.resetEffects {
		resetEffect(sim)
	}

	for i := range at.permanentAuras {
		permAura := &at.permanentAuras[i]
		permAura.aura = permAura.AuraFactory(sim)
		aura := permAura.aura
		if !permAura.RespectDuration {
			aura.Duration = NeverExpires
		}
		aura.Activate(sim)
	}
}

func (at *auraTracker) advance(sim *Simulation) {
	if at.minExpires > sim.CurrentTime {
		return
	}

restart:
	minExpires := NeverExpires
	for _, aura := range at.activeAuras {
		if aura.expires <= sim.CurrentTime {
			aura.Deactivate(sim)
			goto restart // activeAuras have changed
		}
		if aura.expires < minExpires {
			minExpires = aura.expires
		}
	}
	at.minExpires = minExpires
}

func (at *auraTracker) doneIteration(sim *Simulation) {
	// Expire all the remaining auras. Need to keep looping because sometimes
	// expiring auras can trigger other auras.
	foundUnexpired := true
	for foundUnexpired {
		foundUnexpired = false
		for _, aura := range at.auras {
			if aura.IsActive() {
				foundUnexpired = true
				aura.Deactivate(sim)
			}
		}
	}

	for _, aura := range at.auras {
		if aura.IsActive() {
			panic("Active aura during doneIter: " + aura.Label)
		}
		if aura.stacks != 0 {
			panic("Aura nonzero stacks during doneIter: " + aura.Label)
		}
		aura.startTime = 0
		aura.expires = 0
	}

	for _, permAura := range at.permanentAuras {
		if permAura.UptimeMultiplier != 0 {
			permAura.aura.metrics.Uptime = time.Duration(float64(permAura.aura.metrics.Uptime) * permAura.UptimeMultiplier)
		}
	}

	// Add metrics for any auras that are still active.
	for _, aura := range at.auras {
		aura.metrics.doneIteration()
	}
}

// Adds a new aura to the simulation. If an aura with the same ID already
// exists it will be replaced with the new one.
func (aura *Aura) Activate(sim *Simulation) {
	if aura.IsActive() {
		if sim.Log != nil && !aura.ActionID.IsEmptyAction() {
			aura.Unit.Log(sim, "Aura refreshed: %s", aura.ActionID)
		}
		aura.Refresh(sim)
		return
	}

	if aura.Duration == 0 {
		panic("Aura with 0 duration")
	}

	// If there is already an active aura stronger than this one, then this one
	// is blocked.
	if aura.Tag != "" {
		for _, otherAura := range aura.Unit.GetAurasWithTag(aura.Tag) {
			if otherAura.Priority > aura.Priority && otherAura.active {
				return
			}
		}
	}

	// Remove weaker versions of the same aura.
	if aura.Priority != 0 {
		for _, otherAura := range aura.Unit.GetAurasWithTag(aura.Tag) {
			if otherAura.Priority <= aura.Priority && otherAura != aura {
				otherAura.Deactivate(sim)
			}
		}
	}

	aura.active = true
	aura.startTime = sim.CurrentTime
	aura.Refresh(sim)

	if aura.Duration != NeverExpires {
		aura.activeIndex = int32(len(aura.Unit.activeAuras))
		aura.Unit.activeAuras = append(aura.Unit.activeAuras, aura)
	}

	if aura.OnCastComplete != nil {
		aura.onCastCompleteIndex = int32(len(aura.Unit.onCastCompleteAuras))
		aura.Unit.onCastCompleteAuras = append(aura.Unit.onCastCompleteAuras, aura)
	}

	if aura.OnSpellCastComplete != nil {
		aura.onSpellCastCompleteIndex = int32(len(aura.Unit.onSpellCastCompleteAuras))
		aura.Unit.onSpellCastCompleteAuras = append(aura.Unit.onSpellCastCompleteAuras, aura)
	}

	if aura.OnSpellHit != nil {
		aura.onSpellHitIndex = int32(len(aura.Unit.onSpellHitAuras))
		aura.Unit.onSpellHitAuras = append(aura.Unit.onSpellHitAuras, aura)
	}

	if aura.OnPeriodicDamage != nil {
		aura.onPeriodicDamageIndex = int32(len(aura.Unit.onPeriodicDamageAuras))
		aura.Unit.onPeriodicDamageAuras = append(aura.Unit.onPeriodicDamageAuras, aura)
	}

	if sim.Log != nil && !aura.ActionID.IsEmptyAction() {
		aura.Unit.Log(sim, "Aura gained: %s", aura.ActionID)
	}

	if aura.OnGain != nil {
		aura.OnGain(aura, sim)
	}
}

// Moves an Aura to the front of the list of active Auras, so its callbacks are invoked first.
func (aura *Aura) Prioritize() {
	if aura.onCastCompleteIndex > 0 {
		otherAura := aura.Unit.onCastCompleteAuras[0]
		aura.Unit.onCastCompleteAuras[0] = aura
		aura.Unit.onCastCompleteAuras[len(aura.Unit.onCastCompleteAuras)-1] = otherAura
		otherAura.onCastCompleteIndex = aura.onCastCompleteIndex
		aura.onCastCompleteIndex = 0
	}

	if aura.onSpellCastCompleteIndex > 0 {
		otherAura := aura.Unit.onSpellCastCompleteAuras[0]
		aura.Unit.onSpellCastCompleteAuras[0] = aura
		aura.Unit.onSpellCastCompleteAuras[len(aura.Unit.onSpellCastCompleteAuras)-1] = otherAura
		otherAura.onSpellCastCompleteIndex = aura.onSpellCastCompleteIndex
		aura.onSpellCastCompleteIndex = 0
	}

	if aura.onSpellHitIndex > 0 {
		otherAura := aura.Unit.onSpellHitAuras[0]
		aura.Unit.onSpellHitAuras[0] = aura
		aura.Unit.onSpellHitAuras[len(aura.Unit.onSpellHitAuras)-1] = otherAura
		otherAura.onSpellHitIndex = aura.onSpellHitIndex
		aura.onSpellHitIndex = 0
	}

	if aura.onPeriodicDamageIndex > 0 {
		otherAura := aura.Unit.onPeriodicDamageAuras[0]
		aura.Unit.onPeriodicDamageAuras[0] = aura
		aura.Unit.onPeriodicDamageAuras[len(aura.Unit.onPeriodicDamageAuras)-1] = otherAura
		otherAura.onPeriodicDamageIndex = aura.onPeriodicDamageIndex
		aura.onPeriodicDamageIndex = 0
	}
}

// Remove an aura by its ID
func (aura *Aura) Deactivate(sim *Simulation) {
	if !aura.active {
		return
	}
	aura.active = false

	if aura.stacks != 0 {
		aura.SetStacks(sim, 0)
	}
	if aura.OnExpire != nil {
		aura.OnExpire(aura, sim)
	}

	if !aura.ActionID.IsEmptyAction() {
		if sim.CurrentTime > aura.expires {
			aura.metrics.Uptime += aura.expires - aura.startTime
		} else {
			aura.metrics.Uptime += sim.CurrentTime - aura.startTime
		}
	}

	if sim.Log != nil && !aura.ActionID.IsEmptyAction() {
		aura.Unit.Log(sim, "Aura faded: %s", aura.ActionID)
	}

	if aura.Duration != NeverExpires {
		removeActiveIndex := aura.activeIndex
		aura.Unit.activeAuras = removeBySwappingToBack(aura.Unit.activeAuras, removeActiveIndex)
		if removeActiveIndex < int32(len(aura.Unit.activeAuras)) {
			aura.Unit.activeAuras[removeActiveIndex].activeIndex = removeActiveIndex
		}
		aura.activeIndex = Inactive

		aura.Unit.minExpires = 0
	}

	if aura.OnCastComplete != nil {
		removeOnCastCompleteIndex := aura.onCastCompleteIndex
		aura.Unit.onCastCompleteAuras = removeBySwappingToBack(aura.Unit.onCastCompleteAuras, removeOnCastCompleteIndex)
		if removeOnCastCompleteIndex < int32(len(aura.Unit.onCastCompleteAuras)) {
			aura.Unit.onCastCompleteAuras[removeOnCastCompleteIndex].onCastCompleteIndex = removeOnCastCompleteIndex
		}
		aura.onCastCompleteIndex = Inactive
	}

	if aura.OnSpellCastComplete != nil {
		removeOnSpellCastCompleteIndex := aura.onSpellCastCompleteIndex
		aura.Unit.onSpellCastCompleteAuras = removeBySwappingToBack(aura.Unit.onSpellCastCompleteAuras, removeOnSpellCastCompleteIndex)
		if removeOnSpellCastCompleteIndex < int32(len(aura.Unit.onSpellCastCompleteAuras)) {
			aura.Unit.onSpellCastCompleteAuras[removeOnSpellCastCompleteIndex].onSpellCastCompleteIndex = removeOnSpellCastCompleteIndex
		}
		aura.onSpellCastCompleteIndex = Inactive
	}

	if aura.OnSpellHit != nil {
		removeOnSpellHitIndex := aura.onSpellHitIndex
		aura.Unit.onSpellHitAuras = removeBySwappingToBack(aura.Unit.onSpellHitAuras, removeOnSpellHitIndex)
		if removeOnSpellHitIndex < int32(len(aura.Unit.onSpellHitAuras)) {
			aura.Unit.onSpellHitAuras[removeOnSpellHitIndex].onSpellHitIndex = removeOnSpellHitIndex
		}
		aura.onSpellHitIndex = Inactive
	}

	if aura.OnPeriodicDamage != nil {
		removeOnPeriodicDamage := aura.onPeriodicDamageIndex
		aura.Unit.onPeriodicDamageAuras = removeBySwappingToBack(aura.Unit.onPeriodicDamageAuras, removeOnPeriodicDamage)
		if removeOnPeriodicDamage < int32(len(aura.Unit.onPeriodicDamageAuras)) {
			aura.Unit.onPeriodicDamageAuras[removeOnPeriodicDamage].onPeriodicDamageIndex = removeOnPeriodicDamage
		}
		aura.onPeriodicDamageIndex = Inactive
	}
}

// Constant-time removal from slice by swapping with the last element before removing.
func removeBySwappingToBack(arr []*Aura, removeIdx int32) []*Aura {
	arr[removeIdx] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

func (at *auraTracker) IsOnCD(id CooldownID, currentTime time.Duration) bool {
	return at.cooldowns[id] > currentTime
}

func (at *auraTracker) CDReadyAt(id CooldownID) time.Duration {
	return at.cooldowns[id]
}

func (at *auraTracker) GetRemainingCD(id CooldownID, currentTime time.Duration) time.Duration {
	return MaxDuration(0, at.cooldowns[id]-currentTime)
}

func (at *auraTracker) SetCD(id CooldownID, newCD time.Duration) {
	if id == 0 {
		panic("Trying to set CD with ID 0!")
	}
	at.cooldowns[id] = newCD
}

// Invokes the OnCastComplete event for all tracked Auras.
func (at *auraTracker) OnCastComplete(sim *Simulation, cast *Cast) {
	for _, aura := range at.onCastCompleteAuras {
		aura.OnCastComplete(aura, sim, cast)
	}
}

// Invokes the OnSpellCastComplete event for all tracked Auras.
func (at *auraTracker) OnSpellCastComplete(sim *Simulation, spell *Spell) {
	for _, aura := range at.onSpellCastCompleteAuras {
		aura.OnSpellCastComplete(aura, sim, spell)
	}
}

// Invokes the OnSpellHit event for all tracked Auras.
func (at *auraTracker) OnSpellHit(sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
	for _, aura := range at.onSpellHitAuras {
		aura.OnSpellHit(aura, sim, spell, spellEffect)
	}
}

// Invokes the OnPeriodicDamage
//   As a debuff when target is being hit by dot.
//   As a buff when caster's dots are ticking.
func (at *auraTracker) OnPeriodicDamage(sim *Simulation, spell *Spell, spellEffect *SpellEffect, tickDamage float64) {
	for _, aura := range at.onPeriodicDamageAuras {
		aura.OnPeriodicDamage(sim, spell, spellEffect, tickDamage)
	}
}

func (at *auraTracker) GetMetricsProto(numIterations int32) []*proto.AuraMetrics {
	metrics := make([]*proto.AuraMetrics, 0, len(at.auras))

	for _, aura := range at.auras {
		if !aura.metrics.ID.IsEmptyAction() {
			metrics = append(metrics, aura.metrics.ToProto(numIterations))
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

// Helper for the common case of making an aura that adds stats.
func (character *Character) NewTemporaryStatsAura(auraLabel string, actionID ActionID, tempStats stats.Stats, duration time.Duration) *Aura {
	return character.NewTemporaryStatsAuraWrapped(auraLabel, actionID, tempStats, duration, nil)
}

// Alternative that allows modifying the Aura config.
func (character *Character) NewTemporaryStatsAuraWrapped(auraLabel string, actionID ActionID, tempStats stats.Stats, duration time.Duration, modConfig func(*Aura)) *Aura {
	buffs := character.ApplyStatDependencies(tempStats)
	unbuffs := buffs.Multiply(-1)

	config := &Aura{
		Label:    auraLabel,
		ActionID: actionID,
		Duration: duration,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.AddStatsDynamic(sim, buffs)
			if sim.Log != nil {
				character.Log(sim, "Gained %s from %s.", buffs.FlatString(), actionID)
			}
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			if sim.Log != nil {
				character.Log(sim, "Lost %s from fading %s.", buffs.FlatString(), actionID)
			}
			character.AddStatsDynamic(sim, unbuffs)
		},
	}

	if modConfig != nil {
		modConfig(config)
	}

	return character.GetOrRegisterAura(config)
}
