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
	procChance, ohProcChance := core.PPMToChance(character, 1.0)
	mh := character.Equip[proto.ItemSlot_ItemSlotMainHand].Enchant.ID == 22559
	oh := character.Equip[proto.ItemSlot_ItemSlotOffHand].Enchant.ID == 22559

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		buffs := character.ApplyStatDependencies(stats.Stats{stats.Agility: 120})
		buffs[stats.Mana] = 0 // mana is weird
		unbuffs := buffs.Multiply(-1)
		haste := 2 * core.HasteRatingPerHastePercent
		applyLightningSpeed := func(sim *core.Simulation, character *core.Character, id core.AuraID) {
			// https://tbc.wowhead.com/spell=28093/lightning-speed
			character.AddStats(buffs)
			character.AddMeleeHaste(sim, haste)
			var nameStr string
			if id == LightningSpeedMHAuraID {
				nameStr = "Lightning Speed MH"
			} else {
				nameStr = "Lightning Speed OH"
			}
			character.AddAura(sim, core.Aura{
				ID:      id,
				SpellID: 28093,
				Name:    nameStr,
				Expires: sim.CurrentTime + (time.Second * 15),
				OnExpire: func(sim *core.Simulation) {
					character.AddStats(unbuffs)
					character.AddMeleeHaste(sim, -haste)
				},
			})
		}
		return core.Aura{
			ID:   MongooseAuraID,
			Name: "Mongoose Enchant",
			OnBeforeMelee: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, isOH bool) {
				if mh && !isOH && sim.RandomFloat("mongoose") < procChance {
					applyLightningSpeed(sim, character, LightningSpeedMHAuraID)
				} else if oh && isOH && sim.RandomFloat("mongoose") < ohProcChance {
					applyLightningSpeed(sim, character, LightningSpeedOHAuraID)
				}
			},
		}
	})
}
