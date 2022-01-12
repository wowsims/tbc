package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	// Proc effects. Keep these in order by item ID.
	core.AddItemEffect(28034, ApplyHourglassUnraveller)
	core.AddItemEffect(28579, ApplyRomulosPoisonVial)
	core.AddItemEffect(28830, ApplyDragonspineTrophy)
	core.AddItemEffect(32505, ApplyMadnessOfTheBetrayer)

	// Activatable effects. Keep these in order by item ID.
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

var RomulosPoisonVialAuraID = core.NewAuraID()

func ApplyRomulosPoisonVial(agent core.Agent) {
	character := agent.GetCharacter()
	spellObj := core.SimpleSpell{}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		ppmm := character.AutoAttacks.NewPPMManager(1.0)

		castTemplate := core.NewSimpleSpellTemplate(core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:        core.ActionID{ItemID: 28579},
					Character:       character,
					IgnoreCooldowns: true,
					IgnoreManaCost:  true,
					SpellSchool:     stats.NatureSpellPower,
					CritMultiplier:  1.5,
				},
			},
			SpellHitEffect: core.SpellHitEffect{
				SpellEffect: core.SpellEffect{
					DamageMultiplier:       1,
					StaticDamageMultiplier: 1,
				},
				DirectInput: core.DirectDamageInput{
					MinBaseDamage: 222,
					MaxBaseDamage: 332,
				},
			},
		})

		return core.Aura{
			ID: RomulosPoisonVialAuraID,
			OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
				if result == core.MeleeHitTypeMiss || result == core.MeleeHitTypeDodge || result == core.MeleeHitTypeParry {
					return
				}
				if !ppmm.Proc(sim, isOH, "RomulosPoisonVial") {
					return
				}

				castAction := &spellObj
				castTemplate.Apply(castAction)
				castAction.Target = target
				castAction.Init(sim)
				castAction.Cast(sim)
			},
		}
	})
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

var MadnessOfTheBetrayerAuraID = core.NewAuraID()
var MadnessOfTheBetrayerProcAuraID = core.NewAuraID()

func ApplyMadnessOfTheBetrayer(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const arPenBonus = 300
		const dur = time.Second * 10
		ppmm := character.AutoAttacks.NewPPMManager(1.0)

		return core.Aura{
			ID: MadnessOfTheBetrayerAuraID,
			OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
				if result == core.MeleeHitTypeMiss || result == core.MeleeHitTypeDodge || result == core.MeleeHitTypeParry {
					return
				}
				if !ppmm.Proc(sim, isOH, "Madness of the Betrayer") {
					return
				}

				character.AddAuraWithTemporaryStats(sim, MadnessOfTheBetrayerProcAuraID, core.ActionID{ItemID: 32505}, stats.ArmorPenetration, arPenBonus, dur)
			},
		}
	})
}
