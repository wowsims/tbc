package core

import (
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
		target.EnableAutoAttacks(target, AutoAttackOptions{
			MainHand: Weapon{
				BaseDamageMin:  config.MinBaseDamage,
				SwingSpeed:     config.SwingSpeed,
				SwingDuration:  time.Duration(float64(time.Second) * config.SwingSpeed),
				CritMultiplier: 2,
			},
			AutoSwingMelee: true,
		})
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
