package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	// Proc effects. Keep these in order by item ID.
	core.AddItemEffect(23207, ApplyMarkOfTheChampion)
	core.AddItemEffect(27683, ApplyQuagmirransEye)
	core.AddItemEffect(28418, ApplyShiffarsNexusHorn)
	core.AddItemEffect(28789, ApplyEyeOfMagtheridon)
	core.AddItemEffect(30626, ApplySextantOfUnstableCurrents)
	core.AddItemEffect(31856, ApplyDarkmoonCardCrusade)

	// Activatable effects. Keep these in order by item ID.
	core.AddItemEffect(23046, core.MakeTemporaryStatsOnUseCDRegistration(
		core.MagicIDSpellPower,
		stats.SpellPower,
		130,
		time.Second*20,
		core.MajorCooldown{
			CooldownID: core.MagicIDEssSappTrink,
			Cooldown: time.Minute * 2,
			SharedCooldownID: core.MagicIDAtkTrinket,
		},
	))
	core.AddItemEffect(24126, core.MakeTemporaryStatsOnUseCDRegistration(
		core.MagicIDSpellPower,
		stats.SpellPower,
		150,
		time.Second*20,
		core.MajorCooldown{
			CooldownID: core.MagicIDRubySerpentTrink,
			Cooldown: time.Minute * 5,
			SharedCooldownID: core.MagicIDAtkTrinket,
		},
	))
	core.AddItemEffect(29132, core.MakeTemporaryStatsOnUseCDRegistration(
		core.MagicIDSpellPower,
		stats.SpellPower,
		150,
		time.Second*15,
		core.MajorCooldown{
			CooldownID: core.MagicIDScryerTrink,
			Cooldown: time.Second * 90,
			SharedCooldownID: core.MagicIDAtkTrinket,
		},
	))
	core.AddItemEffect(29179, core.MakeTemporaryStatsOnUseCDRegistration(
		core.MagicIDSpellPower,
		stats.SpellPower,
		150,
		time.Second*15,
		core.MajorCooldown{
			CooldownID: core.MagicIDXiriTrink,
			Cooldown: time.Second * 90,
			SharedCooldownID: core.MagicIDAtkTrinket,
		},
	))
	core.AddItemEffect(29370, core.MakeTemporaryStatsOnUseCDRegistration(
		core.MagicIDBlessingSilverCrescent,
		stats.SpellPower,
		155,
		time.Second*20,
		core.MajorCooldown{
			CooldownID: core.MagicIDISCTrink,
			Cooldown: time.Minute * 2,
			SharedCooldownID: core.MagicIDAtkTrinket,
		},
	))
	core.AddItemEffect(29376, core.MakeTemporaryStatsOnUseCDRegistration(
		core.MagicIDEssMartyrTrink,
		stats.SpellPower,
		99,
		time.Second*20,
		core.MajorCooldown{
			CooldownID: core.MagicIDEssMartyrTrink,
			Cooldown: time.Minute * 2,
			SharedCooldownID: core.MagicIDHealTrinket,
		},
	))
	core.AddItemEffect(32483, core.MakeTemporaryStatsOnUseCDRegistration(
		core.MagicIDSkullGuldan,
		stats.SpellHaste,
		175,
		time.Second*20,
		core.MajorCooldown{
			CooldownID: core.MagicIDSkullGuldanTrink,
			Cooldown: time.Minute * 2,
			SharedCooldownID: core.MagicIDAtkTrinket,
		},
	))
	core.AddItemEffect(33829, core.MakeTemporaryStatsOnUseCDRegistration(
		core.MagicIDHexShunkHead,
		stats.SpellPower,
		211,
		time.Second*20,
		core.MajorCooldown{
			CooldownID: core.MagicIDHexTrink,
			Cooldown: time.Minute * 2,
			SharedCooldownID: core.MagicIDAtkTrinket,
		},
	))
	core.AddItemEffect(34429, core.MakeTemporaryStatsOnUseCDRegistration(
		core.MagicIDShiftingNaaru,
		stats.SpellPower,
		320,
		time.Second*15,
		core.MajorCooldown{
			CooldownID: core.MagicIDShiftingNaaruTrink,
			Cooldown: time.Second * 90,
			SharedCooldownID: core.MagicIDAtkTrinket,
		},
	))
	core.AddItemEffect(38290, core.MakeTemporaryStatsOnUseCDRegistration(
		core.MagicIDDarkIronPipeweed,
		stats.SpellPower,
		155,
		time.Second*20,
		core.MajorCooldown{
			CooldownID: core.MagicIDDITrink,
			Cooldown: time.Minute * 2,
			SharedCooldownID: core.MagicIDAtkTrinket,
		},
	))
}

func ApplyMarkOfTheChampion(agent core.Agent) {
	agent.GetCharacter().AddStat(stats.SpellPower, 85)
}

func ApplyQuagmirransEye(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 320.0
		const dur = time.Second * 45
		icd := core.NewICD()

		return core.Aura{
			ID:      core.MagicIDQuagsEye,
			OnCastComplete: func(sim *core.Simulation, cast core.DirectCastAction) {
				if !icd.IsOnCD(sim) && sim.RandomFloat("quags") < 0.1 {
					icd = core.InternalCD(sim.CurrentTime + dur)
					character.AddAuraWithTemporaryStats(sim, core.MagicIDFungalFrenzy, stats.SpellHaste, hasteBonus, time.Second*6)
				}
			},
		}
	})
}

func ApplyShiffarsNexusHorn(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		icd := core.NewICD()
		const spellBonus = 225.0
		const dur = time.Second * 45

		return core.Aura{
			ID:      core.MagicIDNexusHorn,
			OnSpellHit: func(sim *core.Simulation, cast core.DirectCastAction, result *core.DirectCastDamageResult) {
				if cast.GetActionID().ItemID == core.ItemIDTheLightningCapacitor {
					return // TLC can't proc Sextant
				}
				if !icd.IsOnCD(sim) && result.Crit && sim.RandomFloat("unmarked") < 0.2 {
					icd = core.InternalCD(sim.CurrentTime + dur)
					character.AddAuraWithTemporaryStats(sim, core.MagicIDCallOfTheNexus, stats.SpellPower, spellBonus, time.Second*10)
				}
			},
		}
	})
}

func ApplyEyeOfMagtheridon(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const spellBonus = 170.0
		const dur = time.Second * 10

		return core.Aura{
			ID:      core.MagicIDEyeOfMag,
			OnSpellMiss: func(sim *core.Simulation, cast core.DirectCastAction) {
				character.AddAuraWithTemporaryStats(sim, core.MagicIDRecurringPower, stats.SpellPower, spellBonus, dur)
			},
		}
	})
}

func ApplySextantOfUnstableCurrents(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		icd := core.NewICD()
		const spellBonus = 190.0
		const dur = time.Second * 15
		const icdDur = time.Second * 45

		return core.Aura{
			ID:      core.MagicIDSextant,
			OnSpellHit: func(sim *core.Simulation, cast core.DirectCastAction, result *core.DirectCastDamageResult) {
				if cast.GetActionID().ItemID == core.ItemIDTheLightningCapacitor {
					return // TLC can't proc Sextant
				}
				if result.Crit && !icd.IsOnCD(sim) && sim.RandomFloat("unmarked") < 0.2 {
					icd = core.InternalCD(sim.CurrentTime + icdDur)
					character.AddAuraWithTemporaryStats(sim, core.MagicIDUnstableCurrents, stats.SpellPower, spellBonus, dur)
				}
			},
		}
	})
}

func ApplyDarkmoonCardCrusade(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const spellBonus = 8.0
		stacks := 0

		return core.Aura{
			ID:      core.MagicIDDCC,
			OnCastComplete: func(sim *core.Simulation, cast core.DirectCastAction) {
				if stacks < 10 {
					stacks++
					character.AddStat(stats.SpellPower, spellBonus)
				}
				// Removal aura will refresh with new total spellpower based on stacks.
				//  This will remove the old stack removal buff.
				character.AddAura(sim, core.Aura{
					ID:      core.MagicIDDCCBonus,
					Expires: sim.CurrentTime + time.Second*10,
					OnExpire: func(sim *core.Simulation) {
						character.AddStat(stats.SpellPower, -spellBonus * float64(stacks))
						stacks = 0
					},
				})
			},
		}
	})
}
