package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	// Proc effects. Keep these in order by item ID.
	core.AddItemEffect(23207, ApplyMarkOfTheChampionCaster)
	core.AddItemEffect(27683, ApplyQuagmirransEye)
	core.AddItemEffect(28418, ApplyShiffarsNexusHorn)
	core.AddItemEffect(28789, ApplyEyeOfMagtheridon)
	core.AddItemEffect(30626, ApplySextantOfUnstableCurrents)
	core.AddItemEffect(31856, ApplyDarkmoonCardCrusade)

	// Offensive trinkets. Keep these in order by item ID.
	AddSimpleStatOffensiveTrinketEffect(23046, stats.Stats{stats.SpellPower: 130}, time.Second*20, time.Minute*2)  // Restrained Essence of Sapphiron
	AddSimpleStatOffensiveTrinketEffect(24126, stats.Stats{stats.SpellPower: 150}, time.Second*20, time.Minute*5)  // Living Ruby Serpent
	AddSimpleStatOffensiveTrinketEffect(29132, stats.Stats{stats.SpellPower: 150}, time.Second*15, time.Second*90) // Scryer's Bloodgem
	AddSimpleStatOffensiveTrinketEffect(29179, stats.Stats{stats.SpellPower: 150}, time.Second*15, time.Second*90) // Xiri's Gift
	AddSimpleStatOffensiveTrinketEffect(29370, stats.Stats{stats.SpellPower: 155}, time.Second*20, time.Minute*2)  // Icon of the Silver Crescent
	AddSimpleStatOffensiveTrinketEffect(32483, stats.Stats{stats.SpellHaste: 175}, time.Second*20, time.Minute*2)  // Skull of Gul'dan
	AddSimpleStatOffensiveTrinketEffect(33829, stats.Stats{stats.SpellPower: 211}, time.Second*20, time.Minute*2)  // Hex Shrunken Head
	AddSimpleStatOffensiveTrinketEffect(34429, stats.Stats{stats.SpellPower: 320}, time.Second*15, time.Second*90) // Shifting Naaru Sliver
	AddSimpleStatOffensiveTrinketEffect(38290, stats.Stats{stats.SpellPower: 155}, time.Second*20, time.Minute*2)  // Dark Iron Smoking Pipe

	// Defensive trinkets. Keep these in order by item ID.
	AddSimpleStatDefensiveTrinketEffect(29376, stats.Stats{stats.SpellPower: 99}, time.Second*20, time.Minute*2) // Essence of the Marytr

	// Even though these item effects are handled elsewhere, add them so they are
	// detected for automatic testing.
	core.AddItemEffect(core.AlchStoneItemID, func(core.Agent) {})
}

func ApplyMarkOfTheChampionCaster(agent core.Agent) {
	character := agent.GetCharacter()
	if character.CurrentTarget.MobType == proto.MobType_MobTypeDemon || character.CurrentTarget.MobType == proto.MobType_MobTypeUndead {
		character.PseudoStats.MobTypeSpellPower += 85
	}
}

func ApplyQuagmirransEye(agent core.Agent) {
	character := agent.GetCharacter()
	procAura := character.NewTemporaryStatsAura("Fungal Frenzy", core.ActionID{ItemID: 27683}, stats.Stats{stats.SpellHaste: 320}, time.Second*6)

	icd := core.Cooldown{
		Timer:    character.NewTimer(),
		Duration: time.Second * 45,
	}

	character.RegisterAura(core.Aura{
		Label:    "Quagmirran's Eye",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			if !icd.IsReady(sim) || sim.RandomFloat("Quagmirran's Eye") > 0.1 {
				return
			}
			icd.Use(sim)
			procAura.Activate(sim)
		},
	})
}

func ApplyShiffarsNexusHorn(agent core.Agent) {
	character := agent.GetCharacter()
	procAura := character.NewTemporaryStatsAura("Call of the Nexus", core.ActionID{ItemID: 28418}, stats.Stats{stats.SpellPower: 225}, time.Second*10)

	icd := core.Cooldown{
		Timer:    character.NewTimer(),
		Duration: time.Second * 45,
	}

	character.RegisterAura(core.Aura{
		Label:    "Shiffar's Nexus Horn",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.ProcMask.Matches(core.ProcMaskSpellDamage) {
				return
			}
			if !icd.IsReady(sim) || !spellEffect.Outcome.Matches(core.OutcomeCrit) {
				return
			}
			if sim.RandomFloat("Shiffar's Nexus-Horn") > 0.2 {
				return
			}
			icd.Use(sim)
			procAura.Activate(sim)
		},
	})
}

func ApplyEyeOfMagtheridon(agent core.Agent) {
	character := agent.GetCharacter()
	procAura := character.NewTemporaryStatsAura("Recurring Power", core.ActionID{ItemID: 28789}, stats.Stats{stats.SpellPower: 170}, time.Second*10)

	character.RegisterAura(core.Aura{
		Label:    "Eye of Magtheridon",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.ProcMask.Matches(core.ProcMaskSpellDamage) {
				return
			}
			if !spellEffect.Outcome.Matches(core.OutcomeMiss) {
				return
			}
			procAura.Activate(sim)
		},
		OnPeriodicDamageDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.ProcMask.Matches(core.ProcMaskPeriodicDamage) {
				return
			}
			if !spellEffect.Outcome.Matches(core.OutcomeMiss) {
				return
			}
			procAura.Activate(sim)
		},
	})
}

func ApplySextantOfUnstableCurrents(agent core.Agent) {
	character := agent.GetCharacter()
	procAura := character.NewTemporaryStatsAura("Unstable Currents", core.ActionID{ItemID: 30626}, stats.Stats{stats.SpellPower: 190}, time.Second*15)

	icd := core.Cooldown{
		Timer:    character.NewTimer(),
		Duration: time.Second * 45,
	}

	character.RegisterAura(core.Aura{
		Label:    "Sextant of Unstable Currents",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.ProcMask.Matches(core.ProcMaskSpellDamage) {
				return
			}
			if !spellEffect.Outcome.Matches(core.OutcomeCrit) || !icd.IsReady(sim) {
				return
			}
			if sim.RandomFloat("Sextant of Unstable Currents") > 0.2 {
				return
			}
			icd.Use(sim)
			procAura.Activate(sim)
		},
	})
}

func ApplyDarkmoonCardCrusade(agent core.Agent) {
	character := agent.GetCharacter()

	var apBonusPerStack stats.Stats
	apAura := character.RegisterAura(core.Aura{
		Label:     "DMC Crusade AP",
		ActionID:  core.ActionID{ItemID: 31856, Tag: 1},
		Duration:  time.Second * 10,
		MaxStacks: 20,
		OnInit: func(aura *core.Aura, sim *core.Simulation) {
			apBonusPerStack = character.ApplyStatDependencies(stats.Stats{stats.AttackPower: 6, stats.RangedAttackPower: 6})
		},
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			character.AddStatsDynamic(sim, apBonusPerStack.Multiply(float64(newStacks-oldStacks)))
		},
	})

	var spBonusPerStack stats.Stats
	spAura := character.RegisterAura(core.Aura{
		Label:     "DMC Crusade SP",
		ActionID:  core.ActionID{ItemID: 31856, Tag: 2},
		Duration:  time.Second * 10,
		MaxStacks: 10,
		OnInit: func(aura *core.Aura, sim *core.Simulation) {
			spBonusPerStack = character.ApplyStatDependencies(stats.Stats{stats.SpellPower: 8})
		},
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			character.AddStatsDynamic(sim, spBonusPerStack.Multiply(float64(newStacks-oldStacks)))
		},
	})

	character.RegisterAura(core.Aura{
		Label:    "DMC Crusade",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
				apAura.Activate(sim)
				apAura.AddStack(sim)
				apAura.Refresh(sim)
			} else if spellEffect.ProcMask.Matches(core.ProcMaskSpellDamage) {
				if !spellEffect.Landed() {
					return
				}
				spAura.Activate(sim)
				spAura.AddStack(sim)
				spAura.Refresh(sim)
			}
		},
	})
}
