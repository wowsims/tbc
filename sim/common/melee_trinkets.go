package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	// Proc effects. Keep these in order by item ID.
	core.AddItemEffect(28830, ApplyDragonspineTrophy)
	core.AddItemEffect(28034, ApplyHourglassUnraveller)

	// Activatable effects. Keep these in order by item ID.
	var BloodlustBroochCooldownID = core.NewCooldownID()
	core.AddItemEffect(29383, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
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

	var AbacusViolentOddsCooldownID = core.NewCooldownID()
	core.AddItemEffect(28288, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		stats.MeleeHaste,
		260,
		time.Second*10,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 28288},
			CooldownID:       AbacusViolentOddsCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var EmptyDirebrewMugCooldownID = core.NewCooldownID()
	core.AddItemEffect(38287, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		stats.AttackPower,
		278,
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 38287},
			CooldownID:       EmptyDirebrewMugCooldownID,
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

		ppmm := character.AutoAttacks.NewPPMManager(1.0)
		return core.Aura{
			ID: DragonspineTrophyAuraID,
			OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
				if result == core.MeleeHitTypeMiss || result == core.MeleeHitTypeDodge || result == core.MeleeHitTypeParry {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if !ppmm.Proc(sim, isOH, "dragonspine") {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)

				character.AddAuraWithTemporaryStats(sim, MeleeHasteAuraID, core.ActionID{ItemID: 28830}, stats.MeleeHaste, hasteBonus, dur)
			},
		}
	})
}

var HourglassUnravellerAuraID = core.NewAuraID()
var RageOfUnravellerAuraID = core.NewAuraID()

func ApplyHourglassUnraveller(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		icd := core.NewICD()
		const statBonus = 300.0
		const dur = time.Second * 10
		const icdDur = time.Second * 50

		ppmm := character.AutoAttacks.NewPPMManager(1.0)

		return core.Aura{
			ID: HourglassUnravellerAuraID,
			OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
				if result != core.MeleeHitTypeCrit {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if !ppmm.Proc(sim, isOH, "hourglass") {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				character.AddAuraWithTemporaryStats(sim, RageOfUnravellerAuraID, core.ActionID{ItemID: 33648}, stats.AttackPower, statBonus, dur)
			},
		}
	})
}
