package core

import (
	"math"
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
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

func (encounter *Encounter) doneIteration(sim *Simulation) {
	for i, _ := range encounter.Targets {
		target := encounter.Targets[i]
		target.doneIteration(sim)
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

// Target is an enemy/boss that can be the target of player attacks/spells.
type Target struct {
	Unit

	MobType proto.MobType

	MissChance      float64
	HitSuppression  float64
	CritSuppression float64
	Dodge           float64
	Glance          float64
}

func NewTarget(options proto.Target, targetIndex int32) *Target {
	target := &Target{
		Unit: Unit{
			Type:        EnemyUnit,
			Index:       targetIndex,
			Label:       "Target " + strconv.Itoa(int(targetIndex)+1),
			Level:       options.Level,
			auraTracker: newAuraTracker(),
			stats: stats.Stats{
				stats.Armor: float64(options.Armor),
			},
			PseudoStats: stats.NewPseudoStats(),
			Metrics:     NewCharacterMetrics(),
		},
		MobType: options.MobType,
	}
	if target.Level == 0 {
		target.Level = 73
	}
	if target.GetStat(stats.Armor) == 0 {
		target.AddStat(stats.Armor, 7684)
	}

	const skill = 350.0
	skillDifference := float64(target.Level*5) - skill

	target.MissChance = 0.05 + skillDifference*0.002
	target.HitSuppression = (skillDifference - 10) * 0.002
	target.CritSuppression = (skillDifference * 0.002) + 0.018
	target.Dodge = 0.05 + skillDifference*0.001
	target.Glance = math.Max(0.06+skillDifference*0.012, 0)

	if options.Debuffs != nil {
		applyDebuffEffects(target, *options.Debuffs)
	}

	return target
}

func (target *Target) finalize() {
	target.Unit.finalize()
}

func (target *Target) Reset(sim *Simulation) {
	target.Unit.reset(sim)
}

func (target *Target) Advance(sim *Simulation, elapsedTime time.Duration) {
	target.Unit.advance(sim, elapsedTime)
}

func (target *Target) doneIteration(sim *Simulation) {
	target.Unit.doneIteration(sim)
}

func (target *Target) GetMetricsProto(numIterations int32) *proto.TargetMetrics {
	return &proto.TargetMetrics{
		Auras: target.auraTracker.GetMetricsProto(numIterations),
	}
}
