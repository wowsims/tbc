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

	// Activatable effects. Keep these in order by item ID.
	AddSimpleStatItemActiveEffect(23046, stats.Stats{stats.SpellPower: 130}, time.Second*20, time.Minute*2, core.OffensiveTrinketSharedCooldownID)  // Restrained Essence of Sapphiron
	AddSimpleStatItemActiveEffect(24126, stats.Stats{stats.SpellPower: 150}, time.Second*20, time.Minute*5, core.OffensiveTrinketSharedCooldownID)  // Living Ruby Serpent
	AddSimpleStatItemActiveEffect(29132, stats.Stats{stats.SpellPower: 150}, time.Second*15, time.Second*90, core.OffensiveTrinketSharedCooldownID) // Scryer's Bloodgem
	AddSimpleStatItemActiveEffect(29179, stats.Stats{stats.SpellPower: 150}, time.Second*15, time.Second*90, core.OffensiveTrinketSharedCooldownID) // Xiri's Gift
	AddSimpleStatItemActiveEffect(29370, stats.Stats{stats.SpellPower: 155}, time.Second*20, time.Minute*2, core.OffensiveTrinketSharedCooldownID)  // Icon of the Silver Crescent
	AddSimpleStatItemActiveEffect(29376, stats.Stats{stats.SpellPower: 99}, time.Second*20, time.Minute*2, core.DefensiveTrinketSharedCooldownID)   // Essence of the Marytr
	AddSimpleStatItemActiveEffect(32483, stats.Stats{stats.SpellHaste: 175}, time.Second*20, time.Minute*2, core.OffensiveTrinketSharedCooldownID)  // Skull of Gul'dan
	AddSimpleStatItemActiveEffect(33829, stats.Stats{stats.SpellPower: 211}, time.Second*20, time.Minute*2, core.OffensiveTrinketSharedCooldownID)  // Hex Shrunken Head
	AddSimpleStatItemActiveEffect(34429, stats.Stats{stats.SpellPower: 320}, time.Second*15, time.Second*90, core.OffensiveTrinketSharedCooldownID) // Shifting Naaru Sliver
	AddSimpleStatItemActiveEffect(38290, stats.Stats{stats.SpellPower: 155}, time.Second*20, time.Minute*2, core.OffensiveTrinketSharedCooldownID)  // Dark Iron Smoking Pipe

	// Even though these item effects are handled elsewhere, add them so they are
	// detected for automatic testing.
	core.AddItemEffect(core.AlchStoneItemID, func(core.Agent) {})
}

func ApplyMarkOfTheChampionCaster(agent core.Agent) {
	character := agent.GetCharacter()
	character.RegisterResetEffect(func(sim *core.Simulation) {
		if sim.GetPrimaryTarget().MobType == proto.MobType_MobTypeDemon || sim.GetPrimaryTarget().MobType == proto.MobType_MobTypeUndead {
			character.PseudoStats.MobTypeSpellPower += 85
		}
	})
}

func ApplyQuagmirransEye(agent core.Agent) {
	character := agent.GetCharacter()
	procAura := character.NewTemporaryStatsAura("Fungal Frenzy", core.ActionID{ItemID: 27683}, stats.Stats{stats.SpellHaste: 320}, time.Second*6)

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		icd := core.NewICD()
		const icdDur = time.Second * 45

		return character.GetOrRegisterAura(core.Aura{
			Label: "Quagmirran's Eye",
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
				if icd.IsOnCD(sim) || sim.RandomFloat("Quagmirran's Eye") > 0.1 {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				procAura.Activate(sim)
			},
		})
	})
}

func ApplyShiffarsNexusHorn(agent core.Agent) {
	character := agent.GetCharacter()
	procAura := character.NewTemporaryStatsAura("Call of the Nexus", core.ActionID{ItemID: 28418}, stats.Stats{stats.SpellPower: 225}, time.Second*10)

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		icd := core.NewICD()
		const dur = time.Second * 45

		return character.GetOrRegisterAura(core.Aura{
			Label: "Shiffar's Nexus Horn",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}
				if icd.IsOnCD(sim) || !spellEffect.Outcome.Matches(core.OutcomeCrit) || spellEffect.IsPhantom {
					return
				}
				if sim.RandomFloat("Shiffar's Nexus-Horn") > 0.2 {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + dur)
				procAura.Activate(sim)
			},
		})
	})
}

func ApplyEyeOfMagtheridon(agent core.Agent) {
	character := agent.GetCharacter()
	procAura := character.NewTemporaryStatsAura("Recurring Power", core.ActionID{ItemID: 28789}, stats.Stats{stats.SpellPower: 170}, time.Second*10)

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			Label: "Eye of Magtheridon",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}
				if !spellEffect.Outcome.Matches(core.OutcomeMiss) {
					return
				}
				procAura.Activate(sim)
			},
		})
	})
}

func ApplySextantOfUnstableCurrents(agent core.Agent) {
	character := agent.GetCharacter()
	procAura := character.NewTemporaryStatsAura("Unstable Currents", core.ActionID{ItemID: 30626}, stats.Stats{stats.SpellPower: 190}, time.Second*15)

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		icd := core.NewICD()
		const icdDur = time.Second * 45

		return character.GetOrRegisterAura(core.Aura{
			Label: "Sextant of Unstable Currents",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}
				if !spellEffect.Outcome.Matches(core.OutcomeCrit) || icd.IsOnCD(sim) || spellEffect.IsPhantom {
					return
				}
				if sim.RandomFloat("Sextant of Unstable Currents") > 0.2 {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				procAura.Activate(sim)
			},
		})
	})
}

func ApplyDarkmoonCardCrusade(agent core.Agent) {
	character := agent.GetCharacter()

	apAura := character.RegisterAura(core.Aura{
		Label:     "DMC Crusade AP",
		ActionID:  core.ActionID{ItemID: 31856, Tag: 1},
		Duration:  time.Second * 10,
		MaxStacks: 20,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			character.AddStat(stats.AttackPower, 6*float64(newStacks-oldStacks))
			character.AddStat(stats.RangedAttackPower, 6*float64(newStacks-oldStacks))
		},
	})
	spAura := character.RegisterAura(core.Aura{
		Label:     "DMC Crusade SP",
		ActionID:  core.ActionID{ItemID: 31856, Tag: 2},
		Duration:  time.Second * 10,
		MaxStacks: 10,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			character.AddStat(stats.SpellPower, 8*float64(newStacks-oldStacks))
		},
	})

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			Label: "DMC Crusade",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					if spellEffect.IsPhantom {
						return
					}
					apAura.Activate(sim)
					apAura.AddStack(sim)
					apAura.Refresh(sim)
				} else {
					if !spellEffect.Landed() {
						return
					}
					spAura.Activate(sim)
					spAura.AddStack(sim)
					spAura.Refresh(sim)
				}
			},
		})
	})
}
