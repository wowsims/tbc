package core

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

type Encounter struct {
	Duration           time.Duration
	DurationVariation  time.Duration
	executePhaseBegins time.Duration
	Targets            []*Target
}

func NewEncounter(options proto.Encounter) Encounter {
	encounter := Encounter{
		Duration:           DurationFromSeconds(options.Duration),
		DurationVariation:  DurationFromSeconds(options.DurationVariation),
		executePhaseBegins: DurationFromSeconds(options.Duration * (1 - options.ExecuteProportion)),
		Targets:            []*Target{},
	}
	if options.DurationVariation == 0 {
		encounter.DurationVariation = 5 * time.Second
	}

	for targetIndex, targetOptions := range options.Targets {
		target := NewTarget(*targetOptions, int32(targetIndex))
		encounter.Targets = append(encounter.Targets, target)
	}

	encounter.finalize()

	return encounter
}

func (encounter *Encounter) finalize() {
	for _, target := range encounter.Targets {
		target.finalize()
	}
}

func (encounter *Encounter) doneIteration(simDuration time.Duration) {
	for i, _ := range encounter.Targets {
		target := encounter.Targets[i]
		target.doneIteration(simDuration)
	}
}

func (encounter *Encounter) GetMetricsProto(numIterations int32) *proto.EncounterMetrics {
	metrics := &proto.EncounterMetrics{
		Targets: make([]*proto.TargetMetrics, len(encounter.Targets)),
	}

	i := 0
	for _, target := range encounter.Targets {
		metrics.Targets[i] = target.GetMetricsProto(numIterations)
		i++
	}

	return metrics
}

// Target is an enemy that can be the target of attacks/spells.
// Currently targets are basically just lvl 73 target dummies.
type Target struct {
	// Index of this target among all the targets. Primary target has index 0,
	// 2nd target has index 1, etc.
	Index int32

	initialArmor         float64 // base armor
	currentArmor         float64 // current armor, can be mutated by spells
	armorDamageReduction float64 // cached armor damage reduction

	Level int32 // level of target

	MobType proto.MobType

	// Provides aura tracking behavior. Targets need auras to handle debuffs.
	auraTracker

	// Whether finalize() has been called yet for this Character.
	// All fields above this may not be altered once finalized is set.
	finalized bool

	// For logging.
	Name string

	// Cached value to handle sunder/expose overriding each other.
	sunderOrExposeArmorReduction float64
}

func NewTarget(options proto.Target, targetIndex int32) *Target {
	target := &Target{
		Index:        targetIndex,
		currentArmor: float64(options.Armor),
		MobType:      options.MobType,
		auraTracker:  newAuraTracker(true),
		Name:         "Target " + strconv.Itoa(int(targetIndex)+1),
		Level:        73,
	}
	if target.currentArmor == 0 {
		target.currentArmor = 7700
	}
	target.calculateReduction()

	if options.Level > 0 {
		target.Level = options.Level
	}

	if options.Debuffs != nil {
		applyDebuffEffects(target, *options.Debuffs)
	}

	return target
}

func (target *Target) Log(sim *Simulation, message string, vals ...interface{}) {
	sim.Log("%s: "+message, append([]interface{}{target.Name}, vals...)...)
}

func (target *Target) finalize() {
	if target.finalized {
		return
	}
	target.finalized = true

	target.initialArmor = target.currentArmor
	target.auraTracker.finalize()
}

func (target *Target) Reset(sim *Simulation) {
	target.currentArmor = target.initialArmor
	target.auraTracker.reset(sim)
	// Reset after removing any auras above
	target.calculateReduction()
}

func (target *Target) Advance(sim *Simulation, elapsedTime time.Duration) {
	target.auraTracker.advance(sim)
}

func (target *Target) doneIteration(simDuration time.Duration) {
	target.auraTracker.doneIteration(simDuration)
}

func (target *Target) GetMetricsProto(numIterations int32) *proto.TargetMetrics {
	return &proto.TargetMetrics{
		Auras: target.auraTracker.GetMetricsProto(numIterations),
	}
}

func (target *Target) calculateReduction() {
	target.armorDamageReduction = target.currentArmor / (target.currentArmor + 10557.5)
}

func (target *Target) AddArmor(value float64) {
	target.currentArmor += value
	target.calculateReduction()
}

// ArmorDamageReduction currently assumes a level 70 attacker
func (target *Target) ArmorDamageReduction(armorPen float64) float64 {
	// TODO: Cache this somehow so we dont have to recalculate every time.
	effectiveArmor := MaxFloat(0, target.currentArmor-armorPen)
	return effectiveArmor / (effectiveArmor + 10557.5)
}
