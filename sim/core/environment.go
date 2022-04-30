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

type Environment struct {
	State EnvironmentState

	Raid      *Raid
	Encounter Encounter

	BaseDuration      time.Duration // base duration
	DurationVariation time.Duration // variation per duration
	Duration          time.Duration // Duration of current iteration
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
	}
	for _, target := range env.Encounter.Targets {
		target.Env = env
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
	env.Encounter.finalize()
	env.Raid.finalize()

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

func (env *Environment) GetPrimaryTarget() *Target {
	return env.GetTarget(0)
}
