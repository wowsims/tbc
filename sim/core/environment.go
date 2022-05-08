package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

type EnvironmentState int

const (
	Created EnvironmentState = iota
	Constructed
	Initialized
	Finalized
)

// Callback for doing something after finalization.
type PostFinalizeEffect func()

type Environment struct {
	State EnvironmentState

	Raid      *Raid
	Encounter Encounter

	BaseDuration      time.Duration // base duration
	DurationVariation time.Duration // variation per duration

	// Effects to invoke when the Env is finalized.
	postFinalizeEffects []PostFinalizeEffect
}

func NewEnvironment(raidProto proto.Raid, encounterProto proto.Encounter) *Environment {
	env := &Environment{
		State: Created,
	}

	env.construct(raidProto, encounterProto)
	env.initialize(raidProto, encounterProto)
	env.finalize(raidProto, encounterProto)

	return env
}

// The construction phase.
func (env *Environment) construct(raidProto proto.Raid, encounterProto proto.Encounter) {
	env.Encounter = NewEncounter(encounterProto)
	env.BaseDuration = env.Encounter.Duration
	env.DurationVariation = env.Encounter.DurationVariation
	env.Raid = NewRaid(raidProto)

	env.Raid.updatePlayersAndPets()

	for _, unit := range env.Raid.AllUnits {
		unit.Env = env
		unit.CurrentTarget = &env.Encounter.Targets[0].Unit
	}
	for _, target := range env.Encounter.Targets {
		target.Env = env
		if int32(len(encounterProto.Targets)) > target.Index {
			raidTargetProto := encounterProto.Targets[target.Index].Target
			if raidTargetProto != nil {
				target.CurrentTarget = &env.Raid.GetPlayerFromRaidTarget(*raidTargetProto).GetCharacter().Unit
			}
		}
	}

	env.State = Constructed
}

// The initialization phase.
func (env *Environment) initialize(raidProto proto.Raid, encounterProto proto.Encounter) {
	for _, party := range env.Raid.Parties {
		for _, playerOrPet := range party.PlayersAndPets {
			playerOrPet.GetCharacter().initialize()
		}
	}

	env.Raid.applyCharacterEffects(raidProto)

	for _, party := range env.Raid.Parties {
		for _, playerOrPet := range party.PlayersAndPets {
			playerOrPet.Initialize()
		}
	}

	env.State = Initialized
}

// The finalization phase.
func (env *Environment) finalize(raidProto proto.Raid, encounterProto proto.Encounter) {
	for _, target := range env.Encounter.Targets {
		target.finalize()
	}

	for _, party := range env.Raid.Parties {
		for _, player := range party.Players {
			character := player.GetCharacter()
			character.Finalize()

			for _, petAgent := range character.Pets {
				petAgent.GetPet().Finalize()
			}
		}
	}

	for _, finalizeEffect := range env.postFinalizeEffects {
		finalizeEffect()
	}
	env.postFinalizeEffects = nil

	for _, target := range env.Encounter.Targets {
		target.setupAttackTables()
	}

	env.State = Finalized
}

func (env *Environment) IsFinalized() bool {
	return env.State >= Finalized
}

// The maximum possible duration for any iteration.
func (env *Environment) GetMaxDuration() time.Duration {
	return env.BaseDuration + env.DurationVariation
}

func (env *Environment) GetNumTargets() int32 {
	return int32(len(env.Encounter.Targets))
}

func (env *Environment) GetTarget(index int32) *Target {
	return env.Encounter.Targets[index]
}
func (env *Environment) GetTargetUnit(index int32) *Unit {
	return &env.Encounter.Targets[index].Unit
}
func (env *Environment) NextTarget(target *Unit) *Target {
	return env.Encounter.Targets[target.Index].NextTarget()
}
func (env *Environment) NextTargetUnit(target *Unit) *Unit {
	return &env.NextTarget(target).Unit
}

// Registers a callback to this Character which will be invoked after all Units
// are finalized.
func (env *Environment) RegisterPostFinalizeEffect(postFinalizeEffect PostFinalizeEffect) {
	if env.IsFinalized() {
		panic("Finalize effects may not be added once finalized!")
	}

	env.postFinalizeEffects = append(env.postFinalizeEffects, postFinalizeEffect)
}
