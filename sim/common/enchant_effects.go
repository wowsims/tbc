package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	// Keep these in order by item ID.
	core.AddItemEffect(22559, ApplyMongooseEffect)
}

var MongooseAuraID = core.NewAuraID()

var LightningSpeedMHAuraID = core.NewAuraID()
var LightningSpeedOHAuraID = core.NewAuraID()

// ApplyMongooseEffect will be applied twice if there is two weapons with this enchant.
//   However it will automatically overwrite one of them so it should be ok.
//   A single application of the aura will handle both mh and oh procs.
func ApplyMongooseEffect(agent core.Agent) {
	character := agent.GetCharacter()
	ppmm := character.AutoAttacks.NewPPMManager(1.0)
	mh := character.Equip[proto.ItemSlot_ItemSlotMainHand].Enchant.ID == 22559
	oh := character.Equip[proto.ItemSlot_ItemSlotOffHand].Enchant.ID == 22559
	if !mh && !oh {
		return
	}
	if !mh {
		ppmm.SetProcChance(false, 0)
	}
	if !oh {
		ppmm.SetProcChance(true, 0)
	}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		buffs := character.ApplyStatDependencies(stats.Stats{stats.Agility: 120})
		buffs[stats.Mana] = 0 // mana is weird
		unbuffs := buffs.Multiply(-1)
		haste := 2 * core.HasteRatingPerHastePercent

		applyLightningSpeed := func(sim *core.Simulation, character *core.Character, isMH bool) {
			// https://tbc.wowhead.com/spell=28093/lightning-speed
			character.AddStats(buffs)
			character.AddMeleeHaste(sim, haste)
			var tag int32
			var auraID core.AuraID
			if isMH {
				tag = 1
				auraID = LightningSpeedMHAuraID
			} else {
				tag = 2
				auraID = LightningSpeedOHAuraID
			}
			character.AddAura(sim, core.Aura{
				ID:       auraID,
				ActionID: core.ActionID{SpellID: 28093, Tag: tag},
				Expires:  sim.CurrentTime + (time.Second * 15),
				OnExpire: func(sim *core.Simulation) {
					character.AddStats(unbuffs)
					character.AddMeleeHaste(sim, -haste)
				},
			})
		}

		return core.Aura{
			ID: MongooseAuraID,
			OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				// TODO: Does mongoose apply to the hit that procs it? Otherwise this should be OnMeleeAttack
				if hitEffect.IsWeaponHit() {
					isMH := hitEffect.IsMH()
					if ppmm.Proc(sim, isMH, "mongoose") {
						applyLightningSpeed(sim, character, isMH)
					}
				}
			},
		}
	})
}
