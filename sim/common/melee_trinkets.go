package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	// Proc effects. Keep these in order by item ID.
	core.AddItemEffect(28830, ApplyDragonspineTrophy)

	// Activatable effects. Keep these in order by item ID.
	var BloodlustBroochCooldownID = core.NewCooldownID()
	core.AddItemEffect(29383, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		35166,
		"Lust for Battle",
		stats.AttackPower,
		278,
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 29383},
			CooldownID:       BloodlustBroochCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))
}

var DragonspineTrophyAuraID = core.NewAuraID()
var MeleeHasteAuraID = core.NewAuraID()

func ApplyDragonspineTrophy(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		icd := core.NewICD()
		const hasteBonus = 325.0
		const dur = time.Second * 10
		const icdDur = time.Second * 20

		procChance := character.Equip[proto.ItemSlot_ItemSlotMainHand].SwingSpeed / 60.0
		ohProcChance := character.Equip[proto.ItemSlot_ItemSlotOffHand].SwingSpeed / 60.0
		return core.Aura{
			ID:   DragonspineTrophyAuraID,
			Name: "Dragonspine Trophy",
			OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
				if result == core.MeleeHitTypeMiss {
					return
				}
				// Do blocked/parried/dodged attacks proc DST?
				if icd.IsOnCD(sim) {
					return // dont activate
				}
				if !isOH {
					if sim.RandomFloat("dragonspine") > procChance {
						return // didn't proc
					}
				} else {
					if sim.RandomFloat("dragonspine") > ohProcChance {
						return // didn't proc
					}
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				character.AddAuraWithTemporaryStats(sim, MeleeHasteAuraID, 34775, "Haste", stats.MeleeHaste, hasteBonus, dur)
			},
		}
	})
}
