package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	// Proc effects. Keep these in order by item ID.
	core.AddItemEffect(12632, ApplyStormGauntlets)
	core.AddItemEffect(17111, ApplyBlazefuryMedallion)
	core.AddItemEffect(17112, ApplyEmpyreanDemolisher)
	core.AddItemEffect(23541, ApplyKhoriumChampion)
	core.AddItemEffect(24114, ApplyBraidedEterniumChain)
	core.AddItemEffect(27901, ApplyBlackoutTruncheon)
	core.AddItemEffect(28429, ApplyLionheartChampion)
	core.AddItemEffect(28430, ApplyLionheartExecutioner)
	core.AddItemEffect(28437, ApplyDrakefistHammer)
	core.AddItemEffect(28438, ApplyDragonmaw)
	core.AddItemEffect(28439, ApplyDragonstrike)
	core.AddItemEffect(28573, ApplyDespair)
	core.AddItemEffect(28767, ApplyTheDecapitator)
	core.AddItemEffect(28774, ApplyGlaiveOfThePit)
	core.AddItemEffect(29301, ApplyBandOfTheEternalChampion)
	core.AddItemEffect(29348, ApplyTheBladefist)
	core.AddItemEffect(29962, ApplyHeartrazor)
	core.AddItemEffect(29996, ApplyRodOfTheSunKing)
	core.AddItemEffect(30090, ApplyWorldBreaker)
	core.AddItemEffect(30311, ApplyWarpSlicer)
	core.AddItemEffect(30316, ApplyDevastation)
	core.AddItemEffect(31193, ApplyBladeOfUnquenchedThirst)
	core.AddItemEffect(31318, ApplySingingCrystalAxe)
	core.AddItemEffect(31332, ApplyBlinkstrike)
	core.AddItemEffect(31331, ApplyTheNightBlade)
	core.AddItemEffect(32262, ApplySyphonOfTheNathrezim)
	core.AddItemEffect(33122, ApplyCloakOfDarkness)
	core.AddItemEffect(34679, ApplyShatteredSunPendantofMight)

	AddSimpleStatItemEffect(28484, stats.Stats{stats.Health: 1500, stats.Strength: 150}, time.Second*15, time.Minute*30) // Bulwark of Kings
	AddSimpleStatItemEffect(28485, stats.Stats{stats.Health: 1500, stats.Strength: 150}, time.Second*15, time.Minute*30) // Bulwark of Ancient Kings
}

func ApplyShatteredSunPendantofMight(agent core.Agent) {
	character := agent.GetCharacter()
	const proc = 0.15

	var aldorAura *core.Aura
	var scryerSpell *core.Spell

	if character.ShattFaction == proto.ShattrathFaction_ShattrathFactionAldor {
		aldorAura = character.NewTemporaryStatsAura("Light's Strength", core.ActionID{SpellID: 45480}, stats.Stats{stats.AttackPower: 200}, time.Second*10)
	} else if character.ShattFaction == proto.ShattrathFaction_ShattrathFactionScryer {
		scryerSpell = character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 45428},
			SpellSchool: core.SpellSchoolArcane,
			ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
				ProcMask:         core.ProcMaskEmpty,
				IsPhantom:        true,
				DamageMultiplier: 1,
				ThreatMultiplier: 1,
				BaseDamage:       core.BaseDamageConfigRoll(333, 367),
				// TODO: validate this is a melee hit roll
				OutcomeApplier: character.OutcomeFuncMeleeSpecialHitAndCrit(character.DefaultMeleeCritMultiplier()),
			}),
		})
	}

	// Gives a chance when your harmful spells land to increase the damage of your spells and effects by up to 130 for 10 sec. (Proc chance: 20%, 50s cooldown)
	icd := core.Cooldown{
		Timer:    character.NewTimer(),
		Duration: time.Second * 45,
	}

	character.RegisterAura(core.Aura{
		Label:    "Shattered Sun Pendant of Acumen",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
				return
			}
			if !spellEffect.Landed() {
				return
			}
			if !icd.IsReady(sim) || sim.RandomFloat("pendant of acumen") > proc { // can't activate if on CD or didn't proc
				return
			}
			icd.Use(sim)

			if character.ShattFaction == proto.ShattrathFaction_ShattrathFactionAldor {
				aldorAura.Activate(sim)
			} else if character.ShattFaction == proto.ShattrathFaction_ShattrathFactionScryer {
				scryerSpell.Cast(sim, spellEffect.Target)
			}
		},
	})
}

func ApplyStormGauntlets(agent core.Agent) {
	character := agent.GetCharacter()

	procSpell := character.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{ItemID: 12632},
		SpellSchool: core.SpellSchoolNature,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			IsPhantom:        true,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			BaseDamage:     core.BaseDamageConfigFlat(3),
			OutcomeApplier: character.OutcomeFuncMagicHitAndCrit(character.DefaultSpellCritMultiplier()),
		}),
	})

	character.RegisterAura(core.Aura{
		Label:    "Storm Gauntlets",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			// https://tbc.wowhead.com/spell=16615/add-lightning-dam-weap-03, proc mask = 20.
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellEffect.IsPhantom {
				return
			}

			procSpell.Cast(sim, spellEffect.Target)
		},
	})
}

func ApplyBlazefuryMedallion(agent core.Agent) {
	character := agent.GetCharacter()

	procSpell := character.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{ItemID: 17111},
		SpellSchool: core.SpellSchoolFire,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			IsPhantom:        true,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			BaseDamage:     core.BaseDamageConfigFlat(2),
			OutcomeApplier: character.OutcomeFuncMagicHitAndCrit(character.DefaultSpellCritMultiplier()),
		}),
	})

	character.RegisterAura(core.Aura{
		Label:    "Blazefury Medallion",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			// https://tbc.wowhead.com/spell=7711/add-fire-dam-weap-02, proc mask = 20.
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellEffect.IsPhantom {
				return
			}

			procSpell.Cast(sim, spellEffect.Target)
		},
	})
}

func ApplyEmpyreanDemolisher(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(17112)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)
	const procChance = 2.8 / 60.0

	procAura := character.NewTemporaryStatsAura("Empyrean Demolisher Proc", core.ActionID{ItemID: 17112}, stats.Stats{stats.MeleeHaste: 212}, time.Second*10)

	character.GetOrRegisterAura(core.Aura{
		Label:    "Empyrean Demolisher",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
				return
			}
			if sim.RandomFloat("EmpyreanDemolisher") > procChance {
				return
			}

			procAura.Activate(sim)
		},
	})
}

func ApplyKhoriumChampion(agent core.Agent) {
	character := agent.GetCharacter()

	const procChance = 0.5 * 3.3 / 60.0
	procAura := character.NewTemporaryStatsAura("Khorium Champion Proc", core.ActionID{ItemID: 23541}, stats.Stats{stats.Strength: 120}, time.Second*30)

	character.GetOrRegisterAura(core.Aura{
		Label:    "Khorium Champion",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			// https://tbc.wowhead.com/spell=16916/strength-of-the-champion, proc mask = 0. Handled in-game via script.
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}
			if sim.RandomFloat("KhoriumChampion") > procChance {
				return
			}

			procAura.Activate(sim)
		},
	})
}

func ApplyBraidedEterniumChain(agent core.Agent) {
	agent.GetCharacter().PseudoStats.BonusDamage += 5
}

func ApplyBlackoutTruncheon(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(27901)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	const procChance = 1.5 * 0.8 / 60.0
	procAura := character.NewTemporaryStatsAura("Blackout Truncheon Proc", core.ActionID{ItemID: 27901}, stats.Stats{stats.MeleeHaste: 132}, time.Second*10)

	character.GetOrRegisterAura(core.Aura{
		Label:    "Blackout Truncheon",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			// https://tbc.wowhead.com/spell=33489/blinding-speed, proc mask = 0. Handled in-game via script.
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
				return
			}
			if sim.RandomFloat("BlackoutTruncheon") > procChance {
				return
			}

			procAura.Activate(sim)
		},
	})
}

func ApplyLionheartChampion(agent core.Agent) {
	character := agent.GetCharacter()

	const procChance = 3.6 / 60.0
	procAura := character.NewTemporaryStatsAura("Lionheart Champion Proc", core.ActionID{ItemID: 28429}, stats.Stats{stats.Strength: 100}, time.Second*10)

	character.GetOrRegisterAura(core.Aura{
		Label:    "Lionheart Champion",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			// https://tbc.wowhead.com/spell=34513/lionheart, proc mask = 0. Handled in-game via script.
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}
			if sim.RandomFloat("LionheartChampion") > procChance {
				return
			}

			procAura.Activate(sim)
		},
	})
}

func ApplyLionheartExecutioner(agent core.Agent) {
	character := agent.GetCharacter()

	const procChance = 3.6 / 60.0
	procAura := character.NewTemporaryStatsAura("Lionheart Executioner Proc", core.ActionID{ItemID: 28430}, stats.Stats{stats.Strength: 100}, time.Second*10)

	character.GetOrRegisterAura(core.Aura{
		Label:    "Lionheart Executioner",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}
			if sim.RandomFloat("LionheartExecutioner") > procChance {
				return
			}

			procAura.Activate(sim)
		},
	})
}

func ApplyDrakefistHammer(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(28437)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	const procChance = 2.7 / 60.0
	procAura := character.NewTemporaryStatsAura("Drakefist Hammer Proc", core.ActionID{ItemID: 28437}, stats.Stats{stats.MeleeHaste: 212}, time.Second*10)

	character.GetOrRegisterAura(core.Aura{
		Label:    "Drakefist Hammer",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
				return
			}
			if sim.RandomFloat("DrakefistHammer") > procChance {
				return
			}

			procAura.Activate(sim)
		},
	})
}

func ApplyDragonmaw(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(28438)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	const procChance = 2.7 / 60.0
	procAura := character.NewTemporaryStatsAura("Dragonmaw Proc", core.ActionID{ItemID: 28438}, stats.Stats{stats.MeleeHaste: 212}, time.Second*10)

	character.GetOrRegisterAura(core.Aura{
		Label:    "Dragonmaw",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
				return
			}
			if sim.RandomFloat("Dragonmaw") > procChance {
				return
			}

			procAura.Activate(sim)
		},
	})
}

func ApplyDragonstrike(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(28439)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	const procChance = 2.7 / 60.0
	procAura := character.NewTemporaryStatsAura("Dragonstrike Proc", core.ActionID{ItemID: 28439}, stats.Stats{stats.MeleeHaste: 212}, time.Second*10)

	character.GetOrRegisterAura(core.Aura{
		Label:    "Dragonstrike",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
				return
			}
			if sim.RandomFloat("Dragonstrike") > procChance {
				return
			}

			procAura.Activate(sim)
		},
	})
}

func ApplyDespair(agent core.Agent) {
	character := agent.GetCharacter()
	actionID := core.ActionID{SpellID: 34580}

	procSpell := character.GetOrRegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasIgnoreResists,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			// TODO: This should be removed once we have an attack mask.
			//  This is only set here to correctly calculate damage.
			ProcMask:         core.ProcMaskMeleeMHSpecial,
			IsPhantom:        true,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			BaseDamage:     core.BaseDamageConfigFlat(600),
			OutcomeApplier: character.OutcomeFuncMeleeSpecialHitAndCrit(character.DefaultMeleeCritMultiplier()),
		}),
	})

	const procChance = 0.5 * 3.5 / 60.0
	character.GetOrRegisterAura(core.Aura{
		Label:    "Despair",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			// ProcMask: 20
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}
			if sim.RandomFloat("Despair") > procChance {
				return
			}

			procSpell.Cast(sim, spellEffect.Target)
		},
	})
}

func ApplyTheDecapitator(agent core.Agent) {
	character := agent.GetCharacter()
	actionID := core.ActionID{ItemID: 28767}

	spell := character.GetOrRegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasIgnoreResists | core.SpellExtrasNoOnCastComplete,

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    character.NewTimer(),
				Duration: time.Minute * 3,
			},
			SharedCD: core.Cooldown{
				Timer:    character.GetOffensiveTrinketCD(),
				Duration: time.Second * 10,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:         core.ProcMaskMeleeMHSpecial,
			IsPhantom:        true,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			BaseDamage:     core.BaseDamageConfigRoll(513, 567),
			OutcomeApplier: character.OutcomeFuncMeleeSpecialHitAndCrit(character.DefaultMeleeCritMultiplier()),
		}),
	})

	character.AddMajorCooldown(core.MajorCooldown{
		Spell:    spell,
		Priority: core.CooldownPriorityLow, // Use low prio so other actives get used first.
		Type:     core.CooldownTypeDPS,
	})
}

func ApplyGlaiveOfThePit(agent core.Agent) {
	character := agent.GetCharacter()

	procSpell := character.GetOrRegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 34696},
		SpellSchool: core.SpellSchoolShadow,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			IsPhantom:        true,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ProcMask:         core.ProcMaskEmpty,
			BaseDamage:       core.BaseDamageConfigRoll(285, 315),
			OutcomeApplier:   character.OutcomeFuncMagicHitAndCrit(character.DefaultSpellCritMultiplier()),
		}),
	})

	const hasteBonus = 212.0
	const procChance = 3.7 / 60.0

	character.GetOrRegisterAura(core.Aura{
		Label:    "Glaive of the Pit",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}
			if sim.RandomFloat("GlaiveOfThePit") > procChance {
				return
			}

			procSpell.Cast(sim, spellEffect.Target)
		},
	})
}

func ApplyBandOfTheEternalChampion(agent core.Agent) {
	character := agent.GetCharacter()

	procAura := character.NewTemporaryStatsAura("Band of the Eternal Champion Proc", core.ActionID{ItemID: 29301}, stats.Stats{stats.AttackPower: 160, stats.RangedAttackPower: 160}, time.Second*10)
	ppmm := character.AutoAttacks.NewPPMManager(1.0)

	icd := core.Cooldown{
		Timer:    character.NewTimer(),
		Duration: time.Second * 60,
	}

	character.GetOrRegisterAura(core.Aura{
		Label:    "Band of the Eternal Champion",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			// mask 340
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellEffect.IsPhantom {
				return
			}
			if !icd.IsReady(sim) {
				return
			}
			if !ppmm.Proc(sim, spellEffect.IsMH(), spellEffect.ProcMask.Matches(core.ProcMaskRanged), "Band of the Eternal Champion") {
				return
			}

			icd.Use(sim)
			procAura.Activate(sim)
		},
	})
}

func ApplyTheBladefist(agent core.Agent) {
	character := agent.GetCharacter()

	const procChance = 2.7 / 60.0
	procAura := character.NewTemporaryStatsAura("The Bladefist Proc", core.ActionID{ItemID: 29348}, stats.Stats{stats.MeleeHaste: 180}, time.Second*10)

	character.GetOrRegisterAura(core.Aura{
		Label:    "The Bladefist",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeMH) {
				return
			}
			if sim.RandomFloat("The Bladefist") > procChance {
				return
			}

			procAura.Activate(sim)
		},
	})
}

func ApplyHeartrazor(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(29962)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	procAura := character.NewTemporaryStatsAura("Heartrazor Proc", core.ActionID{ItemID: 29962}, stats.Stats{stats.AttackPower: 270, stats.RangedAttackPower: 270}, time.Second*10)
	ppmm := character.AutoAttacks.NewPPMManager(1.0)

	character.GetOrRegisterAura(core.Aura{
		Label:    "Heartrazor",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) || spellEffect.IsPhantom {
				return
			}

			if !ppmm.Proc(sim, spellEffect.IsMH(), false, "Heartrazor") {
				return
			}

			procAura.Activate(sim)
		},
	})
}

func ApplyRodOfTheSunKing(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(29996)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	const procChance = 2.7 / 60.0
	actionID := core.ActionID{ItemID: 29996}

	character.GetOrRegisterAura(core.Aura{
		Label:    "Rod of the Sun King",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
				return
			}

			if spell.Unit.HasRageBar() {
				if sim.RandomFloat("Rod of the Sun King") > procChance {
					return
				}
				spell.Unit.AddRage(sim, 5, actionID)
			} else if spell.Unit.HasEnergyBar() {
				if sim.RandomFloat("Rod of the Sun King") > procChance {
					return
				}
				spell.Unit.AddEnergy(sim, 10, actionID)
			}
		},
	})
}

func ApplyWorldBreaker(agent core.Agent) {
	character := agent.GetCharacter()

	const procChance = 3.7 / 60.0
	procAura := character.NewTemporaryStatsAura("World Breaker Proc", core.ActionID{ItemID: 30090}, stats.Stats{stats.MeleeCrit: 900}, time.Second*4)

	character.RegisterAura(core.Aura{
		Label:    "World Breaker",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
				procAura.Deactivate(sim)
				return
			}
			if sim.RandomFloat("World Breaker") > procChance {
				procAura.Deactivate(sim)
				return
			}

			procAura.Activate(sim)
		},
	})
}

func ApplyWarpSlicer(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(30311)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	const bonus = 1.2
	const inverseBonus = 1 / 1.2
	const procChance = 0.5

	procAura := character.GetOrRegisterAura(core.Aura{
		Label:    "Warp Slicer Proc",
		ActionID: core.ActionID{ItemID: 30311},
		Duration: time.Second * 30,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			character.MultiplyMeleeSpeed(sim, bonus)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			character.MultiplyMeleeSpeed(sim, inverseBonus)
		},
	})

	character.GetOrRegisterAura(core.Aura{
		Label:    "Warp Slicer",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
				return
			}
			if sim.RandomFloat("WarpSlicer") > procChance {
				return
			}

			procAura.Activate(sim)
		},
	})
}

func ApplyDevastation(agent core.Agent) {
	character := agent.GetCharacter()

	const bonus = 1.2
	const inverseBonus = 1 / 1.2
	const procChance = 0.5

	procAura := character.GetOrRegisterAura(core.Aura{
		Label:    "Devastation Proc",
		ActionID: core.ActionID{ItemID: 30316},
		Duration: time.Second * 30,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			character.MultiplyMeleeSpeed(sim, bonus)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			character.MultiplyMeleeSpeed(sim, inverseBonus)
		},
	})

	character.GetOrRegisterAura(core.Aura{
		Label:    "Devastation",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}
			if sim.RandomFloat("Devastation") > procChance {
				return
			}

			procAura.Activate(sim)
		},
	})
}

func ApplyBladeOfUnquenchedThirst(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(31193)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	procSpell := character.GetOrRegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{ItemID: 31193},
		SpellSchool: core.SpellSchoolShadow,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			IsPhantom:        true,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ProcMask:         core.ProcMaskEmpty,
			BaseDamage:       core.BaseDamageConfigMagic(48, 54, 1),
			OutcomeApplier:   character.OutcomeFuncMagicHitAndCrit(character.DefaultSpellCritMultiplier()),
		}),
	})

	const procChance = 0.02
	character.GetOrRegisterAura(core.Aura{
		Label:    "Blade of Unquenched Thirst",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
				return
			}
			if sim.RandomFloat("BladeOfUnquenchedThirst") > procChance {
				return
			}

			procSpell.Cast(sim, spellEffect.Target)
		},
	})
}

func ApplySingingCrystalAxe(agent core.Agent) {
	character := agent.GetCharacter()

	const procChance = 3.5 / 60.0
	procAura := character.NewTemporaryStatsAura("Singing Crystal Axe Proc", core.ActionID{ItemID: 31318}, stats.Stats{stats.MeleeHaste: 400}, time.Second*10)

	character.GetOrRegisterAura(core.Aura{
		Label:    "Singing Crystal Axe",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}
			if sim.RandomFloat("SingingCrystalAxe") > procChance {
				return
			}

			procAura.Activate(sim)
		},
	})
}

func ApplyBlinkstrike(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(31332)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	ppmm := character.AutoAttacks.NewPPMManager(1.0)
	if !mh {
		ppmm.SetProcChance(true, 0)
	}
	if !oh {
		ppmm.SetProcChance(false, 0)
	}

	var blinkstrikeSpell *core.Spell
	icd := core.Cooldown{
		Timer:    character.NewTimer(),
		Duration: time.Millisecond,
	}

	character.GetOrRegisterAura(core.Aura{
		Label:    "Blinkstrike",
		Duration: core.NeverExpires,
		OnInit: func(aura *core.Aura, sim *core.Simulation) {
			blinkstrikeSpell = character.GetOrRegisterSpell(core.SpellConfig{
				ActionID:     core.ActionID{ItemID: 31332},
				SpellSchool:  core.SpellSchoolPhysical,
				SpellExtras:  core.SpellExtrasMeleeMetrics | core.SpellExtrasNoOnCastComplete,
				ApplyEffects: core.ApplyEffectFuncDirectDamage(character.AutoAttacks.MHEffect),
			})
		},
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) || spellEffect.IsPhantom {
				return
			}

			if !icd.IsReady(sim) {
				return
			}

			if !ppmm.Proc(sim, spellEffect.IsMH(), false, "Blinkstrike") {
				return
			}
			icd.Use(sim)

			aura.Unit.AutoAttacks.MaybeReplaceMHSwing(sim, blinkstrikeSpell).Cast(sim, spellEffect.Target)
		},
	})
}

func ApplyTheNightBlade(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(31331)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	procAura := character.GetOrRegisterAura(core.Aura{
		Label:     "The Night Blade Proc",
		ActionID:  core.ActionID{ItemID: 31331},
		Duration:  time.Second * 10,
		MaxStacks: 3,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			character.AddStatDynamic(sim, stats.ArmorPenetration, 435*float64(newStacks-oldStacks))
		},
	})

	const procChance = 2 * 1.8 / 60.0
	character.GetOrRegisterAura(core.Aura{
		Label:    "The Night Blade",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
				return
			}
			if sim.RandomFloat("The Night Blade") > procChance {
				return
			}

			procAura.Activate(sim)
			procAura.AddStack(sim)
		},
	})
}

func ApplySyphonOfTheNathrezim(agent core.Agent) {
	character := agent.GetCharacter()
	ppmm := character.AutoAttacks.NewPPMManager(1.0)
	mh, oh := character.GetWeaponHands(32262)
	if !mh {
		ppmm.SetProcChance(true, 0)
	}
	if !oh {
		ppmm.SetProcChance(false, 0)
	}

	procSpell := character.GetOrRegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 40291},
		SpellSchool: core.SpellSchoolShadow,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			IsPhantom:        true,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ProcMask:         core.ProcMaskEmpty,
			BaseDamage:       core.BaseDamageConfigFlat(20),
			OutcomeApplier:   character.OutcomeFuncMagicHitAndCrit(character.DefaultSpellCritMultiplier()),
		}),
	})

	procAura := character.GetOrRegisterAura(core.Aura{
		Label:    "Siphon Essence",
		ActionID: core.ActionID{SpellID: 40291},
		Duration: time.Second * 6,
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellEffect.IsPhantom {
				return
			}

			procSpell.Cast(sim, spellEffect.Target)
		},
	})

	character.GetOrRegisterAura(core.Aura{
		Label:    "Syphon of the Nathrezim",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}

			if ppmm.Proc(sim, spellEffect.IsMH(), false, "Syphon Of The Nathrezim") {
				procAura.Activate(sim)
			}
		},
	})
}

func ApplyCloakOfDarkness(agent core.Agent) {
	character := agent.GetCharacter()

	if character.Class != proto.Class_ClassHunter {
		// For non-hunters just give direct crit so it shows on the stats panel.
		character.AddStats(stats.Stats{
			stats.MeleeCrit: 24,
		})
	} else {
		character.PseudoStats.BonusMeleeCritRating += 24
	}
}
