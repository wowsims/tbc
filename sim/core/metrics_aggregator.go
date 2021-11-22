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


type CharacterMetrics struct {
	CharacterIterationMetrics

	// Aggregate values. These are updated after each iteration.
	dpsSum        float64
	dpsSumSquared float64
	dpsMax        float64
	dpsHist       map[int32]int32 // rounded DPS to count
	numOom        int32
	oomAtSum      float64
	dpsAtOomSum   float64
	actions       map[ActionKey]ActionMetrics
}

// Metrics for the current iteration, for 1 agent. Keep this as a separate
// struct so its easy to clear.
type CharacterIterationMetrics struct {
	TotalDamage float64
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
		dpsHist: make(map[int32]int32),
		actions: make(map[ActionKey]ActionMetrics),
	}
}

func (characterMetrics *CharacterMetrics) addCastInternal(actionID ActionID, manaCost float64) {
	characterMetrics.ManaSpent += manaCost

	actionKey := NewActionKey(actionID)
	actionMetrics, ok := characterMetrics.actions[actionKey]

	if !ok {
		actionMetrics.ActionID = actionID
	}

	actionMetrics.Casts++

	characterMetrics.actions[actionKey] = actionMetrics
}

func (characterMetrics *CharacterMetrics) AddInstantCast(actionID ActionID) {
	characterMetrics.addCastInternal(actionID, 0)
}

// Adds the results of a cast to the aggregated metrics.
func (characterMetrics *CharacterMetrics) AddCast(cast *Cast) {
	manaCost := cast.ManaCost
	if cast.IgnoreManaCost {
		manaCost = 0
	}

	characterMetrics.addCastInternal(cast.ActionID, manaCost)
}

// Adds the results of an action to the aggregated metrics.
func (characterMetrics *CharacterMetrics) AddSpellCast(spellCast *SpellCast) {
	if !spellCast.IgnoreManaCost {
		characterMetrics.ManaSpent += spellCast.ManaCost
	}

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

func (characterMetrics *CharacterMetrics) MarkOOM(sim *Simulation, character *Character) {
	if characterMetrics.OOMAt == 0 {
		if sim.Log != nil {
			sim.Log("(%d) Went OOM!\n", character.ID)
		}
		characterMetrics.DamageAtOOM = characterMetrics.TotalDamage
		characterMetrics.OOMAt = sim.CurrentTime
	}
}

// This should be called when a Sim iteration is complete.
func (characterMetrics *CharacterMetrics) doneIteration(encounterDurationSeconds float64) {
	dps := characterMetrics.TotalDamage / encounterDurationSeconds
	// log.Printf("total: %0.1f, dur: %0.1f, dps: %0.1f", metrics.TotalDamage, encounterDurationSeconds, dps)

	characterMetrics.dpsSum += dps
	characterMetrics.dpsSumSquared += dps * dps
	characterMetrics.dpsMax = MaxFloat(characterMetrics.dpsMax, dps)

	dpsRounded := int32(math.Round(dps/10) * 10)
	characterMetrics.dpsHist[dpsRounded]++

	if characterMetrics.OOMAt > 0 {
		characterMetrics.numOom++
		characterMetrics.oomAtSum += float64(characterMetrics.OOMAt)
		characterMetrics.dpsAtOomSum += float64(characterMetrics.DamageAtOOM) / float64(characterMetrics.OOMAt.Seconds())
	}

	// Clear the iteration metrics
	characterMetrics.CharacterIterationMetrics = CharacterIterationMetrics{}
}

func (characterMetrics *CharacterMetrics) ToProto(numIterations int32) *proto.PlayerMetrics {
	dpsAvg := characterMetrics.dpsSum / float64(numIterations)

	protoMetrics := &proto.PlayerMetrics{
		DpsAvg: dpsAvg,
		DpsStdev: math.Sqrt((characterMetrics.dpsSumSquared / float64(numIterations)) - (dpsAvg * dpsAvg)),
		DpsMax: characterMetrics.dpsMax,
		DpsHist: characterMetrics.dpsHist,

		NumOom: characterMetrics.numOom,
	}

	if characterMetrics.numOom > 0 {
		protoMetrics.OomAtAvg = characterMetrics.oomAtSum / float64(characterMetrics.numOom)
		protoMetrics.DpsAtOomAvg = characterMetrics.dpsAtOomSum / float64(characterMetrics.numOom)
	}

	for _, action := range characterMetrics.actions {
		protoMetrics.Actions = append(protoMetrics.Actions, action.ToProto())
	}

	return protoMetrics
}
