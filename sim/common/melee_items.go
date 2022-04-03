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
}

func ApplyStormGauntlets(agent core.Agent) {
	character := agent.GetCharacter()

	procSpell := character.RegisterSpell(core.SpellConfig{
		Template: core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:    core.ActionID{ItemID: 12632},
					Character:   character,
					SpellSchool: core.SpellSchoolNature,
				},
			},
			Effect: core.SpellEffect{
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				CritRollCategory:    core.CritRollCategoryMagical,
				CritMultiplier:      character.DefaultSpellCritMultiplier(),
				IsPhantom:           true,
				DamageMultiplier:    1,
				ThreatMultiplier:    1,
				BaseDamage:          core.BaseDamageConfigFlat(3),
			},
		},
		ModifyCast: core.ModifyCastAssignTarget,
	})

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return character.GetOrRegisterAura(&core.Aura{
			Label: "Storm Gauntlets",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				// https://tbc.wowhead.com/spell=16615/add-lightning-dam-weap-03, proc mask = 20.
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellEffect.IsPhantom {
					return
				}

				procSpell.Cast(sim, spellEffect.Target)
			},
		})
	})
}

func ApplyBlazefuryMedallion(agent core.Agent) {
	character := agent.GetCharacter()

	procSpell := character.RegisterSpell(core.SpellConfig{
		Template: core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:    core.ActionID{ItemID: 17111},
					Character:   character,
					SpellSchool: core.SpellSchoolFire,
				},
			},
			Effect: core.SpellEffect{
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				CritRollCategory:    core.CritRollCategoryMagical,
				CritMultiplier:      character.DefaultSpellCritMultiplier(),
				IsPhantom:           true,
				DamageMultiplier:    1,
				ThreatMultiplier:    1,
				BaseDamage:          core.BaseDamageConfigFlat(2),
			},
		},
		ModifyCast: core.ModifyCastAssignTarget,
	})

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return character.GetOrRegisterAura(&core.Aura{
			Label: "Blazefury Medallion",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				// https://tbc.wowhead.com/spell=7711/add-fire-dam-weap-02, proc mask = 20.
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellEffect.IsPhantom {
					return
				}

				procSpell.Cast(sim, spellEffect.Target)
			},
		})
	})
}

func ApplyEmpyreanDemolisher(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(17112)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)
	const procChance = 2.8 / 60.0

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		procAura := character.NewTemporaryStatsAura("Empyrean Demolisher Proc", core.ActionID{ItemID: 17112}, stats.Stats{stats.MeleeHaste: 212}, time.Second*10)

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Empyrean Demolisher",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("EmpyreanDemolisher") > procChance {
					return
				}

				procAura.Activate(sim)
			},
		})
	})
}

func ApplyKhoriumChampion(agent core.Agent) {
	character := agent.GetCharacter()

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		const procChance = 0.5 * 3.3 / 60.0
		procAura := character.NewTemporaryStatsAura("Khorium Champion Proc", core.ActionID{ItemID: 23541}, stats.Stats{stats.Strength: 120}, time.Second*30)

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Khorium Champion",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
	})
}

func ApplyBraidedEterniumChain(agent core.Agent) {
	agent.GetCharacter().PseudoStats.BonusDamage += 5
}

func ApplyBlackoutTruncheon(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(27901)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		const procChance = 1.5 * 0.8 / 60.0
		procAura := character.NewTemporaryStatsAura("Blackout Truncheon Proc", core.ActionID{ItemID: 27901}, stats.Stats{stats.MeleeHaste: 132}, time.Second*10)

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Blackout Truncheon",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
	})
}

func ApplyLionheartChampion(agent core.Agent) {
	character := agent.GetCharacter()

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		const procChance = 3.6 / 60.0
		procAura := character.NewTemporaryStatsAura("Lionheart Champion Proc", core.ActionID{ItemID: 28429}, stats.Stats{stats.Strength: 100}, time.Second*10)

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Lionheart Champion",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
	})
}

func ApplyLionheartExecutioner(agent core.Agent) {
	character := agent.GetCharacter()

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		const procChance = 3.6 / 60.0
		procAura := character.NewTemporaryStatsAura("Lionheart Executioner Proc", core.ActionID{ItemID: 28430}, stats.Stats{stats.Strength: 100}, time.Second*10)

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Lionheart Executioner",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("LionheartExecutioner") > procChance {
					return
				}

				procAura.Activate(sim)
			},
		})
	})
}

func ApplyDrakefistHammer(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(28437)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		const procChance = 2.7 / 60.0
		procAura := character.NewTemporaryStatsAura("Drakefist Hammer Proc", core.ActionID{ItemID: 28437}, stats.Stats{stats.MeleeHaste: 212}, time.Second*10)

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Drakefist Hammer",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("DrakefistHammer") > procChance {
					return
				}

				procAura.Activate(sim)
			},
		})
	})
}

func ApplyDragonmaw(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(28438)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		const procChance = 2.7 / 60.0
		procAura := character.NewTemporaryStatsAura("Dragonmaw Proc", core.ActionID{ItemID: 28438}, stats.Stats{stats.MeleeHaste: 212}, time.Second*10)

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Dragonmaw",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("Dragonmaw") > procChance {
					return
				}

				procAura.Activate(sim)
			},
		})
	})
}

func ApplyDragonstrike(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(28439)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		const procChance = 2.7 / 60.0
		procAura := character.NewTemporaryStatsAura("Dragonstrike Proc", core.ActionID{ItemID: 28439}, stats.Stats{stats.MeleeHaste: 212}, time.Second*10)

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Dragonstrike",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("Dragonstrike") > procChance {
					return
				}

				procAura.Activate(sim)
			},
		})
	})
}

func ApplyDespair(agent core.Agent) {
	character := agent.GetCharacter()
	actionID := core.ActionID{SpellID: 34580}

	procSpell := character.RegisterSpell(core.SpellConfig{
		Template: core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:    actionID,
					Character:   character,
					SpellSchool: core.SpellSchoolPhysical,
					SpellExtras: core.SpellExtrasIgnoreResists,
				},
			},
			Effect: core.SpellEffect{
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				CritMultiplier:      character.DefaultMeleeCritMultiplier(),
				IsPhantom:           true,
				// TODO: This should be removed once we have an attack mask.
				//  This is only set here to correctly calculate damage.
				ProcMask:         core.ProcMaskMeleeMHSpecial,
				DamageMultiplier: 1,
				ThreatMultiplier: 1,
				BaseDamage:       core.BaseDamageConfigFlat(600),
			},
		},
		ModifyCast: core.ModifyCastAssignTarget,
	})

	const procChance = 0.5 * 3.5 / 60.0
	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return character.GetOrRegisterAura(&core.Aura{
			Label: "Despair",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
	})
}

var TheDecapitatorCooldownID = core.NewCooldownID()

func ApplyTheDecapitator(agent core.Agent) {
	character := agent.GetCharacter()
	actionID := core.ActionID{ItemID: 28767}

	spell := character.RegisterSpell(core.SpellConfig{
		Template: core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:    actionID,
					Character:   character,
					SpellSchool: core.SpellSchoolPhysical,
					SpellExtras: core.SpellExtrasIgnoreResists,
				},
			},
			Effect: core.SpellEffect{
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				CritMultiplier:      character.DefaultMeleeCritMultiplier(),
				IsPhantom:           true,
				ProcMask:            core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:    1,
				ThreatMultiplier:    1,
				BaseDamage:          core.BaseDamageConfigRoll(513, 567),
			},
		},
		ModifyCast: core.ModifyCastAssignTarget,
	})

	character.AddMajorCooldown(core.MajorCooldown{
		ActionID:         actionID,
		CooldownID:       TheDecapitatorCooldownID,
		Cooldown:         time.Minute * 3,
		SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		SharedCooldown:   time.Second * 10,
		Priority:         core.CooldownPriorityLow, // Use low prio so other actives get used first.
		Type:             core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				spell.Cast(sim, sim.GetPrimaryTarget())

				character.SetCD(TheDecapitatorCooldownID, sim.CurrentTime+time.Minute*3)
				character.SetCD(core.OffensiveTrinketSharedCooldownID, sim.CurrentTime+time.Second*10)
			}
		},
	})
}

func ApplyGlaiveOfThePit(agent core.Agent) {
	character := agent.GetCharacter()

	procSpell := character.RegisterSpell(core.SpellConfig{
		Template: core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:    core.ActionID{SpellID: 34696},
					Character:   character,
					SpellSchool: core.SpellSchoolShadow,
				},
			},
			Effect: core.SpellEffect{
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				CritRollCategory:    core.CritRollCategoryMagical,
				CritMultiplier:      character.DefaultSpellCritMultiplier(),
				IsPhantom:           true,
				DamageMultiplier:    1,
				ThreatMultiplier:    1,
				BaseDamage:          core.BaseDamageConfigRoll(285, 315),
			},
		},
		ModifyCast: core.ModifyCastAssignTarget,
	})

	const hasteBonus = 212.0
	const procChance = 3.7 / 60.0

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return character.GetOrRegisterAura(&core.Aura{
			Label: "Glaive of the Pit",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("GlaiveOfThePit") > procChance {
					return
				}

				procSpell.Cast(sim, spellEffect.Target)
			},
		})
	})
}

func ApplyBandOfTheEternalChampion(agent core.Agent) {
	character := agent.GetCharacter()

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		procAura := character.NewTemporaryStatsAura("Band of the Eternal Champion Proc", core.ActionID{ItemID: 29301}, stats.Stats{stats.AttackPower: 160, stats.RangedAttackPower: 160}, time.Second*10)

		ppmm := character.AutoAttacks.NewPPMManager(1.0)
		icd := core.NewICD()
		const icdDur = time.Second * 60

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Band of the Eternal Champion",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				// mask 340
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellEffect.IsPhantom {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if !ppmm.Proc(sim, spellEffect.IsMH(), spellEffect.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged), "Band of the Eternal Champion") {
					return
				}

				icd = core.InternalCD(sim.CurrentTime + icdDur)
				procAura.Activate(sim)
			},
		})
	})
}

func ApplyTheBladefist(agent core.Agent) {
	character := agent.GetCharacter()

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		const procChance = 2.7 / 60.0
		procAura := character.NewTemporaryStatsAura("The Bladefist Proc", core.ActionID{ItemID: 29348}, stats.Stats{stats.MeleeHaste: 180}, time.Second*10)

		return character.GetOrRegisterAura(&core.Aura{
			Label: "The Bladefist",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeMH) {
					return
				}
				if sim.RandomFloat("The Bladefist") > procChance {
					return
				}

				procAura.Activate(sim)
			},
		})
	})
}

func ApplyHeartrazor(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(29962)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		procAura := character.NewTemporaryStatsAura("Heartrazor Proc", core.ActionID{ItemID: 29962}, stats.Stats{stats.AttackPower: 270, stats.RangedAttackPower: 270}, time.Second*10)

		ppmm := character.AutoAttacks.NewPPMManager(1.0)

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Heartrazor",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) || spellEffect.IsPhantom {
					return
				}

				if !ppmm.Proc(sim, spellEffect.IsMH(), false, "Heartrazor") {
					return
				}

				procAura.Activate(sim)
			},
		})
	})
}

func ApplyRodOfTheSunKing(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(29996)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	const procChance = 2.7 / 60.0
	actionID := core.ActionID{ItemID: 29996}

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return character.GetOrRegisterAura(&core.Aura{
			Label: "Rod of the Sun King",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}

				if spell.Character.HasRageBar() {
					if sim.RandomFloat("Rod of the Sun King") > procChance {
						return
					}
					spell.Character.AddRage(sim, 5, actionID)
				} else if spell.Character.HasEnergyBar() {
					if sim.RandomFloat("Rod of the Sun King") > procChance {
						return
					}
					spell.Character.AddEnergy(sim, 10, actionID)
				}
			},
		})
	})
}

func ApplyWorldBreaker(agent core.Agent) {
	character := agent.GetCharacter()

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		const procChance = 3.7 / 60.0
		procAura := character.NewTemporaryStatsAura("World Breaker Proc", core.ActionID{ItemID: 30090}, stats.Stats{stats.MeleeCrit: 900}, time.Second*4)

		return character.GetOrRegisterAura(&core.Aura{
			Label: "World Breaker",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
	})
}

func ApplyWarpSlicer(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(30311)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	const bonus = 1.2
	const inverseBonus = 1 / 1.2
	const procChance = 0.5

	procAura := character.GetOrRegisterAura(&core.Aura{
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

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return character.GetOrRegisterAura(&core.Aura{
			Label: "Warp Slicer",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("WarpSlicer") > procChance {
					return
				}

				procAura.Activate(sim)
			},
		})
	})
}

func ApplyDevastation(agent core.Agent) {
	character := agent.GetCharacter()

	const bonus = 1.2
	const inverseBonus = 1 / 1.2
	const procChance = 0.5

	procAura := character.GetOrRegisterAura(&core.Aura{
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

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return character.GetOrRegisterAura(&core.Aura{
			Label: "Devastation",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("Devastation") > procChance {
					return
				}

				procAura.Activate(sim)
			},
		})
	})
}

func ApplyBladeOfUnquenchedThirst(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(31193)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	procSpell := character.RegisterSpell(core.SpellConfig{
		Template: core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:    core.ActionID{ItemID: 31193},
					Character:   character,
					SpellSchool: core.SpellSchoolShadow,
				},
			},
			Effect: core.SpellEffect{
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				CritRollCategory:    core.CritRollCategoryMagical,
				CritMultiplier:      character.DefaultSpellCritMultiplier(),
				IsPhantom:           true,
				DamageMultiplier:    1,
				ThreatMultiplier:    1,
				BaseDamage:          core.BaseDamageConfigMagic(48, 54, 1),
			},
		},
		ModifyCast: core.ModifyCastAssignTarget,
	})

	const procChance = 0.02
	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return character.GetOrRegisterAura(&core.Aura{
			Label: "Blade of Unquenched Thirst",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("BladeOfUnquenchedThirst") > procChance {
					return
				}

				procSpell.Cast(sim, spellEffect.Target)
			},
		})
	})
}

func ApplySingingCrystalAxe(agent core.Agent) {
	character := agent.GetCharacter()

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		const procChance = 3.5 / 60.0
		procAura := character.NewTemporaryStatsAura("Singing Crystal Axe Proc", core.ActionID{ItemID: 31318}, stats.Stats{stats.MeleeHaste: 400}, time.Second*10)

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Singing Crystal Axe",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("SingingCrystalAxe") > procChance {
					return
				}

				procAura.Activate(sim)
			},
		})
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

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		var icd core.InternalCD
		icdDur := time.Millisecond * 1

		template := character.AutoAttacks.MHAuto.Template
		template.ActionID = core.ActionID{ItemID: 31332}
		blinkstrikeSpell := character.GetOrRegisterSpell(core.SpellConfig{
			Template:   template,
			ModifyCast: core.ModifyCastAssignTarget,
		})

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Blinkstrike",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) || spellEffect.IsPhantom {
					return
				}

				if icd.IsOnCD(sim) {
					return
				}

				if !ppmm.Proc(sim, spellEffect.IsMH(), false, "Blinkstrike") {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)

				blinkstrikeSpell.Cast(sim, spellEffect.Target)
			},
		})
	})
}

func ApplyTheNightBlade(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(31331)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	procAura := character.GetOrRegisterAura(&core.Aura{
		Label:     "The Night Blade Proc",
		ActionID:  core.ActionID{ItemID: 31331},
		Duration:  time.Second * 10,
		MaxStacks: 3,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			character.AddStat(stats.ArmorPenetration, 435*float64(newStacks-oldStacks))
		},
	})

	const procChance = 2 * 1.8 / 60.0
	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return character.GetOrRegisterAura(&core.Aura{
			Label: "The Night Blade",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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

	procSpell := character.RegisterSpell(core.SpellConfig{
		Template: core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:    core.ActionID{SpellID: 40291},
					Character:   character,
					SpellSchool: core.SpellSchoolShadow,
				},
			},
			Effect: core.SpellEffect{
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				CritRollCategory:    core.CritRollCategoryMagical,
				CritMultiplier:      character.DefaultSpellCritMultiplier(),
				IsPhantom:           true,
				DamageMultiplier:    1,
				ThreatMultiplier:    1,
				BaseDamage:          core.BaseDamageConfigFlat(20),
			},
		},
		ModifyCast: core.ModifyCastAssignTarget,
	})

	procAura := character.GetOrRegisterAura(&core.Aura{
		Label:    "Siphon Essence",
		ActionID: core.ActionID{SpellID: 40291},
		Duration: time.Second * 6,
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellEffect.IsPhantom {
				return
			}

			procSpell.Cast(sim, spellEffect.Target)
		},
	})

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return character.GetOrRegisterAura(&core.Aura{
			Label: "Syphon of the Nathrezim",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}

				if ppmm.Proc(sim, spellEffect.IsMH(), false, "Syphon Of The Nathrezim") {
					procAura.Activate(sim)
				}
			},
		})
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
