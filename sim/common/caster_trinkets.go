package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Keep these (and their functions) in order by item ID.
func init() {
	core.AddActiveItem(23046, core.ActiveItem{Activate: core.CreateSpellDmgActivate(core.MagicIDSpellPower, 130, time.Second*20), ActivateCD: time.Second * 120, CoolID: core.MagicIDEssSappTrink, SharedID: core.MagicIDAtkTrinket})
	core.AddActiveItem(23207, core.ActiveItem{BuffUp: ActivateMarkOfTheChampion, ActivateCD: core.NeverExpires})
	core.AddActiveItem(24126, core.ActiveItem{Activate: core.CreateSpellDmgActivate(core.MagicIDRubySerpent, 150, time.Second*20), ActivateCD: time.Second * 300, CoolID: core.MagicIDRubySerpentTrink, SharedID: core.MagicIDAtkTrinket})
	core.AddActiveItem(27683, core.ActiveItem{BuffUp: ActivateQuagsEye, ActivateCD: core.NeverExpires, SharedID: core.MagicIDAtkTrinket})
	core.AddActiveItem(28418, core.ActiveItem{BuffUp: ActivateNexusHorn, ActivateCD: core.NeverExpires, SharedID: core.MagicIDAtkTrinket})
	core.AddActiveItem(28789, core.ActiveItem{BuffUp: ActivateEyeOfMag, ActivateCD: core.NeverExpires, SharedID: core.MagicIDAtkTrinket})
	core.AddActiveItem(29132, core.ActiveItem{Activate: core.CreateSpellDmgActivate(core.MagicIDSpellPower, 150, time.Second*15), ActivateCD: time.Second * 90, CoolID: core.MagicIDScryerTrink, SharedID: core.MagicIDAtkTrinket})
	core.AddActiveItem(29179, core.ActiveItem{Activate: core.CreateSpellDmgActivate(core.MagicIDSpellPower, 150, time.Second*15), ActivateCD: time.Second * 90, CoolID: core.MagicIDXiriTrink, SharedID: core.MagicIDAtkTrinket})
	core.AddActiveItem(29370, core.ActiveItem{Activate: core.CreateSpellDmgActivate(core.MagicIDBlessingSilverCrescent, 155, time.Second*20), ActivateCD: time.Second * 120, CoolID: core.MagicIDISCTrink, SharedID: core.MagicIDAtkTrinket})
	core.AddActiveItem(29376, core.ActiveItem{Activate: core.CreateSpellDmgActivate(core.MagicIDSpellPower, 99, time.Second*20), ActivateCD: time.Second * 120, CoolID: core.MagicIDEssMartyrTrink, SharedID: core.MagicIDHealTrinket})
	core.AddActiveItem(30626, core.ActiveItem{BuffUp: ActivateSextant, ActivateCD: core.NeverExpires, SharedID: core.MagicIDAtkTrinket})
	core.AddActiveItem(31856, core.ActiveItem{BuffUp: ActivateDCC, ActivateCD: core.NeverExpires, SharedID: core.MagicIDAtkTrinket})
	core.AddActiveItem(32483, core.ActiveItem{Activate: core.CreateHasteActivate(core.MagicIDSkullGuldan, 175, time.Second*20), ActivateCD: time.Second * 120, CoolID: core.MagicIDSkullGuldanTrink, SharedID: core.MagicIDAtkTrinket})
	core.AddActiveItem(33829, core.ActiveItem{Activate: core.CreateSpellDmgActivate(core.MagicIDHexShunkHead, 211, time.Second*20), ActivateCD: time.Second * 120, CoolID: core.MagicIDHexTrink, SharedID: core.MagicIDAtkTrinket})
	core.AddActiveItem(34429, core.ActiveItem{Activate: core.CreateSpellDmgActivate(core.MagicIDShiftingNaaru, 320, time.Second*15), ActivateCD: time.Second * 90, CoolID: core.MagicIDShiftingNaaruTrink, SharedID: core.MagicIDAtkTrinket})
	core.AddActiveItem(35749, core.ActiveItem{BuffUp: ActivateAlchStone, ActivateCD: core.NeverExpires, SharedID: core.MagicIDAtkTrinket})
	core.AddActiveItem(38290, core.ActiveItem{Activate: core.CreateSpellDmgActivate(core.MagicIDDarkIronPipeweed, 155, time.Second*20), ActivateCD: time.Second * 120, CoolID: core.MagicIDDITrink, SharedID: core.MagicIDAtkTrinket})
}

func ActivateMarkOfTheChampion(sim *core.Simulation, agent core.Agent) {
	agent.GetCharacter().Stats[stats.SpellPower] += 85
}

func ActivateQuagsEye(sim *core.Simulation, agent core.Agent) {
	character := agent.GetCharacter()
	const hasteBonus = 320.0
	const dur = time.Second * 45
	icd := core.NewICD()

	character.AddAura(sim, core.Aura{
		ID:      core.MagicIDQuagsEye,
		Expires: core.NeverExpires,
		OnCastComplete: func(sim *core.Simulation, cast core.DirectCastAction) {
			if !icd.IsOnCD(sim) && sim.Rando.Float64("quags") < 0.1 {
				icd = core.InternalCD(sim.CurrentTime + dur)
				core.AddAuraWithTemporaryStats(sim, character, core.MagicIDFungalFrenzy, stats.SpellHaste, hasteBonus, time.Second*6)
			}
		},
	})
}

func ActivateNexusHorn(sim *core.Simulation, agent core.Agent) {
	character := agent.GetCharacter()
	icd := core.NewICD()
	const spellBonus = 225.0
	const dur = time.Second * 45

	character.AddAura(sim, core.Aura{
		ID:      core.MagicIDNexusHorn,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, cast core.DirectCastAction, result *core.DirectCastDamageResult) {
			if cast.GetActionID().ItemID == core.ItemIDTheLightningCapacitor {
				return // TLC can't proc Sextant
			}
			if !icd.IsOnCD(sim) && result.Crit && sim.Rando.Float64("unmarked") < 0.2 {
				icd = core.InternalCD(sim.CurrentTime + dur)
				core.AddAuraWithTemporaryStats(sim, character, core.MagicIDCallOfTheNexus, stats.SpellPower, spellBonus, time.Second*10)
			}
		},
	})
}

func ActivateEyeOfMag(sim *core.Simulation, agent core.Agent) {
	character := agent.GetCharacter()
	const spellBonus = 170.0
	const dur = time.Second * 10

	character.AddAura(sim, core.Aura{
		ID:      core.MagicIDEyeOfMag,
		Expires: core.NeverExpires,
		OnSpellMiss: func(sim *core.Simulation, cast core.DirectCastAction) {
			core.AddAuraWithTemporaryStats(sim, character, core.MagicIDRecurringPower, stats.SpellPower, spellBonus, dur)
		},
	})
}

func ActivateSextant(sim *core.Simulation, agent core.Agent) {
	character := agent.GetCharacter()
	icd := core.NewICD()
	const spellBonus = 190.0
	const dur = time.Second * 15
	const icdDur = time.Second * 45

	character.AddAura(sim, core.Aura{
		ID:      core.MagicIDSextant,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, cast core.DirectCastAction, result *core.DirectCastDamageResult) {
			if cast.GetActionID().ItemID == core.ItemIDTheLightningCapacitor {
				return // TLC can't proc Sextant
			}
			if result.Crit && !icd.IsOnCD(sim) && sim.Rando.Float64("unmarked") < 0.2 {
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				core.AddAuraWithTemporaryStats(sim, character, core.MagicIDUnstableCurrents, stats.SpellPower, spellBonus, dur)
			}
		},
	})
}

func ActivateDCC(sim *core.Simulation, agent core.Agent) {
	character := agent.GetCharacter()
	const spellBonus = 8.0
	stacks := 0

	character.AddAura(sim, core.Aura{
		ID:      core.MagicIDDCC,
		Expires: core.NeverExpires,
		OnCastComplete: func(sim *core.Simulation, cast core.DirectCastAction) {
			if stacks < 10 {
				stacks++
				character.Stats[stats.SpellPower] += spellBonus
			}
			// Removal aura will refresh with new total spellpower based on stacks.
			//  This will remove the old stack removal buff.
			character.AddAura(sim, core.Aura{
				ID:      core.MagicIDDCCBonus,
				Expires: sim.CurrentTime + time.Second*10,
				OnExpire: func(sim *core.Simulation) {
					character.Stats[stats.SpellPower] -= spellBonus * float64(stacks)
					stacks = 0
				},
			})
		},
	})
}

// ActivateAlchStone adds the alch stone aura that has no effect on casts.
//  The usage for this aura is in the potion usage function.
func ActivateAlchStone(sim *core.Simulation, agent core.Agent) {
	agent.GetCharacter().AddAura(sim, core.Aura{
		ID:      core.MagicIDAlchStone,
		Expires: core.NeverExpires,
	})
}
