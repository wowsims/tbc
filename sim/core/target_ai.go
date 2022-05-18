package core

import (
	"log"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

type TargetAI struct {
	Target *Target
}

func (target *Target) initialize(config *proto.Target) {
	if config == nil || target.CurrentTarget == nil {
		return
	}

	if config.SwingSpeed > 0 {
		aaOptions := AutoAttackOptions{
			MainHand: Weapon{
				BaseDamageMin:  config.MinBaseDamage,
				SwingSpeed:     config.SwingSpeed,
				SwingDuration:  time.Duration(float64(time.Second) * config.SwingSpeed),
				CritMultiplier: 2,
				SpellSchool:    SpellSchoolFromProto(config.SpellSchool),
			},
			AutoSwingMelee: true,
		}
		if config.DualWield {
			aaOptions.OffHand = aaOptions.MainHand
			if !config.DualWieldPenalty {
				target.PseudoStats.DisableDWMissPenalty = true
			}
		}
		target.EnableAutoAttacks(target, aaOptions)
	}

	//target.gcdAction = &PendingAction{
	//	Priority: ActionPriorityGCD,
	//	OnAction: func(sim *Simulation) {
	//	},
	//}
}

// Empty Agent interface functions.
// TODO: Figure out how to get rid of these.
func (target *Target) AddRaidBuffs(raidBuffs *proto.RaidBuffs)    {}
func (target *Target) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {}
func (target *Target) ApplyTalents()                              {}
func (target *Target) GetCharacter() *Character                   { return nil }
func (target *Target) Initialize()                                {}
func (target *Target) OnGCDReady(*Simulation)                     {}
func (target *Target) OnAutoAttack(sim *Simulation, spell *Spell) {}

type PresetTarget struct {
	// String in folder-structure format identifying a category for this unit, e.g. "Black Temple/Bosses".
	PathPrefix string

	Config proto.Target
}

func (pt PresetTarget) Path() string {
	return pt.PathPrefix + "/" + pt.Config.Name
}
func (pt PresetTarget) ToProto() *proto.PresetTarget {
	target := &proto.Target{}
	*target = pt.Config

	return &proto.PresetTarget{
		Path:   pt.Path(),
		Target: target,
	}
}

var presetTargets = []PresetTarget{}
var presetEncounters = []*proto.PresetEncounter{}

func AddPresetTarget(newPreset PresetTarget) {
	for _, preset := range presetTargets {
		if preset.Path() == newPreset.Path() {
			log.Fatalf("Preset Target with path %s already added!", newPreset.Path())
		}
	}
	presetTargets = append(presetTargets, newPreset)
}

func GetPresetTargetWithPath(path string) *PresetTarget {
	for i, _ := range presetTargets {
		preset := &presetTargets[i]
		if preset.Path() == path {
			return preset
		}
	}
	return nil
}

func AddPresetEncounter(name string, targetPaths []string) {
	if len(targetPaths) == 0 {
		log.Fatalf("Encounter must have targets!")
	}

	var path string
	var targetProtos []*proto.PresetTarget

	for i, targetPath := range targetPaths {
		presetTarget := GetPresetTargetWithPath(targetPath)
		if presetTarget == nil {
			log.Fatalf("No preset target with path: %s", targetPath)
		}
		targetProtos = append(targetProtos, presetTarget.ToProto())

		if i == 0 {
			path = presetTarget.PathPrefix + "/" + name
		}
	}

	for _, preset := range presetEncounters {
		if preset.Path == path {
			log.Fatalf("Preset Encounter with path %s already added!", path)
		}
	}

	presetEncounters = append(presetEncounters, &proto.PresetEncounter{
		Path:    path,
		Targets: targetProtos,
	})
}
