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
	numOom      int32
	oomAtSum    float64
	dpsAtOomSum float64
	actions     map[ActionKey]ActionMetrics
	resources   []ResourceMetric
}

type ResourceMetric struct {
	seconds  float64
	resource float64
}

// Metrics for the current iteration, for 1 agent. Keep this as a separate
// struct so its easy to clear.
type CharacterIterationMetrics struct {
	ManaSpent   float64
	DamageAtOOM float64
	OOMAt       time.Duration
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

func (cm *CharacterMetrics) addCastInternal(actionID ActionID, manaCost float64) {
	cm.ManaSpent += manaCost

	actionKey := NewActionKey(actionID)
	actionMetrics, ok := cm.actions[actionKey]

	if !ok {
		actionMetrics.ActionID = actionID
	}

	actionMetrics.Casts++

	cm.actions[actionKey] = actionMetrics
}

func (cm *CharacterMetrics) AddInstantCast(actionID ActionID) {
	cm.addCastInternal(actionID, 0)
}

// Adds the results of a cast to the aggregated metrics.
func (cm *CharacterMetrics) AddCast(cast *Cast) {
	manaCost := cast.ManaCost
	if cast.IgnoreManaCost {
		manaCost = 0
	}

	cm.addCastInternal(cast.ActionID, manaCost)
}

// Adds the results of an action to the aggregated metrics.
func (cm *CharacterMetrics) AddSpellCast(spellCast *SpellCast) {
	if !spellCast.IgnoreManaCost {
		cm.ManaSpent += spellCast.ManaCost
	}

	actionID := spellCast.ActionID
	actionKey := NewActionKey(actionID)
	actionMetrics, ok := cm.actions[actionKey]

	if !ok {
		actionMetrics.ActionID = actionID
	}

	actionMetrics.Casts++
	actionMetrics.Hits += spellCast.Hits
	actionMetrics.Misses += spellCast.Misses
	actionMetrics.Crits += spellCast.Crits
	actionMetrics.Damage += spellCast.TotalDamage
	cm.TotalDamage += spellCast.TotalDamage

	cm.actions[actionKey] = actionMetrics
}

func (cm *CharacterMetrics) MarkOOM(sim *Simulation, character *Character) {
	if cm.OOMAt == 0 {
		if sim.Log != nil {
			character.Log(sim, "Went OOM!")
		}
		cm.DamageAtOOM = cm.TotalDamage
		cm.OOMAt = sim.CurrentTime
	}
}

// This should be called when a Sim iteration is complete.
func (cm *CharacterMetrics) doneIteration(encounterDurationSeconds float64) {
	cm.DpsMetrics.doneIteration(encounterDurationSeconds)

	if cm.OOMAt > 0 {
		cm.numOom++
		cm.oomAtSum += float64(cm.OOMAt)
		cm.dpsAtOomSum += float64(cm.DamageAtOOM) / float64(cm.OOMAt.Seconds())
	}

	// Clear the iteration metrics
	cm.CharacterIterationMetrics = CharacterIterationMetrics{}
}

func (cm *CharacterMetrics) ToProto(numIterations int32) *proto.PlayerMetrics {
	protoMetrics := &proto.PlayerMetrics{
		Dps:    cm.DpsMetrics.ToProto(numIterations),
		NumOom: cm.numOom,
	}

	if cm.numOom > 0 {
		protoMetrics.OomAtAvg = cm.oomAtSum / float64(cm.numOom)
		protoMetrics.DpsAtOomAvg = cm.dpsAtOomSum / float64(cm.numOom)
	}

	resources := make([]*proto.ResourceMetrics, len(cm.resources))
	for i, v := range cm.resources {
		resources[i] = &proto.ResourceMetrics{
			Seconds:  v.seconds,
			Resource: v.resource,
		}
	}
	protoMetrics.Resources = resources

	for _, action := range cm.actions {
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
