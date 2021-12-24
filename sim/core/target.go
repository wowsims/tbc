package core

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

type Encounter struct {
	Duration float64
	Targets  []*Target
}

func NewEncounter(options proto.Encounter) Encounter {
	encounter := Encounter{
		Duration: options.Duration,
		Targets:  []*Target{},
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

	armor float64

	Level int32 // level of target

	MobType proto.MobType

	// Provides aura tracking behavior. Targets need auras to handle debuffs.
	auraTracker

	// Whether finalize() has been called yet for this Character.
	// All fields above this may not be altered once finalized is set.
	finalized bool

	// For logging.
	Name string
}

func NewTarget(options proto.Target, targetIndex int32) *Target {
	target := &Target{
		Index:       targetIndex,
		armor:       float64(options.Armor),
		MobType:     options.MobType,
		auraTracker: newAuraTracker(true),
		Name:        "Target " + strconv.Itoa(int(targetIndex)+1),
		Level:       73,
	}
	if target.armor == 0 {
		target.armor = 7700
	}
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

	target.auraTracker.finalize()
}

func (target *Target) Reset(sim *Simulation) {
	target.auraTracker.reset(sim)
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

func (target *Target) ArmorDamageReduction() float64 {
	return 0
}
