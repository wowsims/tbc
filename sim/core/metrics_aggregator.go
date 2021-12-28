package core

import (
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

// A unique number based on an ActionID.
// This works by making item IDs negative to avoid collisions, and assumes
// there are no collisions with OtherID. Tag adds decimals.
// Actual key values dont matter, just need something unique and fast to compute.
type ActionKey float64

func NewActionKey(actionID ActionID) ActionKey {
	return ActionKey(float64((int32(actionID.OtherID) + actionID.SpellID - actionID.ItemID)) + (float64(actionID.Tag) / 256))
}

type DpsMetrics struct {
	// Values for the current iteration. These are cleared after each iteration.
	TotalDamage float64

	// Aggregate values. These are updated after each iteration.
	sum        float64
	sumSquared float64
	max        float64
	hist       map[int32]int32 // rounded DPS to count
}

// This should be called when a Sim iteration is complete.
func (dpsMetrics *DpsMetrics) doneIteration(encounterDurationSeconds float64) {
	dps := dpsMetrics.TotalDamage / encounterDurationSeconds

	dpsMetrics.sum += dps
	dpsMetrics.sumSquared += dps * dps
	dpsMetrics.max = MaxFloat(dpsMetrics.max, dps)

	dpsRounded := int32(math.Round(dps/10) * 10)
	dpsMetrics.hist[dpsRounded]++

	// Clear the iteration metrics
	dpsMetrics.TotalDamage = 0
}

func (dpsMetrics *DpsMetrics) ToProto(numIterations int32) *proto.DpsMetrics {
	dpsAvg := dpsMetrics.sum / float64(numIterations)

	return &proto.DpsMetrics{
		Avg:   dpsAvg,
		Stdev: math.Sqrt((dpsMetrics.sumSquared / float64(numIterations)) - (dpsAvg * dpsAvg)),
		Max:   dpsMetrics.max,
		Hist:  dpsMetrics.hist,
	}
}

func NewDpsMetrics() DpsMetrics {
	return DpsMetrics{
		hist: make(map[int32]int32),
	}
}

type CharacterMetrics struct {
	DpsMetrics
	CharacterIterationMetrics

	// Aggregate values. These are updated after each iteration.
	oomTimeSum float64
	actions    map[ActionKey]ActionMetrics
}

// Metrics for the current iteration, for 1 agent. Keep this as a separate
// struct so its easy to clear.
type CharacterIterationMetrics struct {
	WentOOM bool // Whether the agent has hit OOM at least once in this iteration.

	ManaSpent       float64
	ManaGained      float64
	BonusManaGained float64 // Only includes amount from mana pots / runes / innervates.

	OOMTime time.Duration // time spent not casting and waiting for regen.
}

type ActionMetrics struct {
	ActionID ActionID

	Casts  int32
	Hits   int32
	Crits  int32
	Misses int32

	Damage float64
}

func (actionMetrics *ActionMetrics) ToProto() *proto.ActionMetrics {
	return &proto.ActionMetrics{
		Id: actionMetrics.ActionID.ToProto(),

		Casts:  actionMetrics.Casts,
		Hits:   actionMetrics.Hits,
		Crits:  actionMetrics.Crits,
		Misses: actionMetrics.Misses,
		Damage: actionMetrics.Damage,
	}
}

func NewCharacterMetrics() CharacterMetrics {
	return CharacterMetrics{
		DpsMetrics: NewDpsMetrics(),
		actions:    make(map[ActionKey]ActionMetrics),
	}
}

func (characterMetrics *CharacterMetrics) addCastInternal(actionID ActionID) {
	actionKey := NewActionKey(actionID)
	actionMetrics, ok := characterMetrics.actions[actionKey]

	if !ok {
		actionMetrics.ActionID = actionID
	}

	actionMetrics.Casts++

	characterMetrics.actions[actionKey] = actionMetrics
}

func (characterMetrics *CharacterMetrics) AddInstantCast(actionID ActionID) {
	characterMetrics.addCastInternal(actionID)
}

// Adds the results of a cast to the aggregated metrics.
func (characterMetrics *CharacterMetrics) AddCast(cast *Cast) {
	characterMetrics.addCastInternal(cast.ActionID)
}

// Adds the results of an action to the aggregated metrics.
func (characterMetrics *CharacterMetrics) AddSpellCast(spellCast *SpellCast) {
	actionID := spellCast.ActionID
	actionKey := NewActionKey(actionID)
	actionMetrics, ok := characterMetrics.actions[actionKey]

	if !ok {
		actionMetrics.ActionID = actionID
	}

	actionMetrics.Casts++
	actionMetrics.Hits += spellCast.Hits
	actionMetrics.Misses += spellCast.Misses
	actionMetrics.Crits += spellCast.Crits
	actionMetrics.Damage += spellCast.TotalDamage
	characterMetrics.TotalDamage += spellCast.TotalDamage

	characterMetrics.actions[actionKey] = actionMetrics
}

// Adds the results of a melee action to the aggregated metrics.
func (characterMetrics *CharacterMetrics) AddAutoAttack(itemID int32, result MeleeHitType, dmg float64, isOH bool) {
	var tag int32 = 10
	if isOH {
		tag = 11
	}
	actionID := ActionID{ItemID: itemID, Tag: tag}
	actionKey := NewActionKey(actionID)
	actionMetrics, ok := characterMetrics.actions[actionKey]
	if !ok {
		actionMetrics.ActionID = actionID
	}
	actionMetrics.Casts++
	if result == MeleeHitTypeBlock || result == MeleeHitTypeMiss || result == MeleeHitTypeParry || result == MeleeHitTypeDodge {
		actionMetrics.Misses++
	} else {
		actionMetrics.Hits++
		if result == MeleeHitTypeCrit {
			actionMetrics.Crits++
		}
	}
	actionMetrics.Damage += dmg
	characterMetrics.TotalDamage += dmg
	characterMetrics.actions[actionKey] = actionMetrics
}

// Adds the results of a melee action to the aggregated metrics.
func (characterMetrics *CharacterMetrics) AddMeleeAbility(ability *ActiveMeleeAbility) {
	actionID := ability.ActionID
	actionKey := NewActionKey(actionID)
	actionMetrics, ok := characterMetrics.actions[actionKey]

	if !ok {
		actionMetrics.ActionID = actionID
	}

	actionMetrics.Casts++
	actionMetrics.Hits += ability.Hits
	actionMetrics.Misses += ability.Misses
	actionMetrics.Crits += ability.Crits
	actionMetrics.Damage += ability.TotalDamage
	characterMetrics.TotalDamage += ability.TotalDamage

	characterMetrics.actions[actionKey] = actionMetrics
}

func (characterMetrics *CharacterMetrics) MarkOOM(sim *Simulation, character *Character, dur time.Duration) {
	characterMetrics.CharacterIterationMetrics.OOMTime += dur
	characterMetrics.CharacterIterationMetrics.WentOOM = true
}

// This should be called when a Sim iteration is complete.
func (characterMetrics *CharacterMetrics) doneIteration(encounterDurationSeconds float64) {
	characterMetrics.DpsMetrics.doneIteration(encounterDurationSeconds)
	characterMetrics.oomTimeSum += float64(characterMetrics.OOMTime.Seconds())

	// Clear the iteration metrics
	characterMetrics.CharacterIterationMetrics = CharacterIterationMetrics{}
}

func (characterMetrics *CharacterMetrics) ToProto(numIterations int32) *proto.PlayerMetrics {
	protoMetrics := &proto.PlayerMetrics{
		Dps:           characterMetrics.DpsMetrics.ToProto(numIterations),
		SecondsOomAvg: characterMetrics.oomTimeSum / float64(numIterations),
	}

	for _, action := range characterMetrics.actions {
		protoMetrics.Actions = append(protoMetrics.Actions, action.ToProto())
	}

	return protoMetrics
}

type AuraMetrics struct {
	ID int32

	// Metrics for the current iteration.
	Uptime time.Duration

	// Aggregate values. These are updated after each iteration.
	uptimeSum        time.Duration
	uptimeSumSquared time.Duration
}

// This should be called when a Sim iteration is complete.
func (auraMetrics *AuraMetrics) doneIteration() {
	auraMetrics.uptimeSum += auraMetrics.Uptime
	auraMetrics.uptimeSumSquared += auraMetrics.Uptime * auraMetrics.Uptime

	auraMetrics.Uptime = 0
}

func (auraMetrics *AuraMetrics) ToProto(numIterations int32) *proto.AuraMetrics {
	uptimeAvg := auraMetrics.uptimeSum.Seconds() / float64(numIterations)

	return &proto.AuraMetrics{
		Id: auraMetrics.ID,

		UptimeSecondsAvg:   uptimeAvg,
		UptimeSecondsStdev: math.Sqrt((auraMetrics.uptimeSumSquared.Seconds() / float64(numIterations)) - (uptimeAvg * uptimeAvg)),
	}
}
