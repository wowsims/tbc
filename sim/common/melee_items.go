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
	core.AddItemEffect(-23, ApplyDragonstrikeStacking)
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

var StormGauntletsAuraID = core.NewAuraID()

func ApplyStormGauntlets(agent core.Agent) {
	character := agent.GetCharacter()
	spellObj := core.SimpleSpell{}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		castTemplate := core.NewSimpleSpellTemplate(core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:            core.ActionID{ItemID: 12632},
					Character:           character,
					IsPhantom:           true,
					CritRollCategory:    core.CritRollCategoryMagical,
					OutcomeRollCategory: core.OutcomeRollCategoryMagic,
					SpellSchool:         core.SpellSchoolNature,
					CritMultiplier:      character.DefaultSpellCritMultiplier(),
				},
			},
			Effect: core.SpellHitEffect{
				SpellEffect: core.SpellEffect{
					DamageMultiplier:       1,
					StaticDamageMultiplier: 1,
					ThreatMultiplier:       1,
				},
				DirectInput: core.DirectDamageInput{
					MinBaseDamage: 3,
					MaxBaseDamage: 3,
				},
			},
		})

		return core.Aura{
			ID: StormGauntletsAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				// https://tbc.wowhead.com/spell=16615/add-lightning-dam-weap-03, proc mask = 20.
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellCast.IsPhantom {
					return
				}

				castAction := &spellObj
				castTemplate.Apply(castAction)
				castAction.Effect.Target = spellEffect.Target
				castAction.Init(sim)
				castAction.Cast(sim)
			},
		}
	})
}

var BlazefuryMedallionAuraID = core.NewAuraID()

func ApplyBlazefuryMedallion(agent core.Agent) {
	character := agent.GetCharacter()
	spellObj := core.SimpleSpell{}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		castTemplate := core.NewSimpleSpellTemplate(core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:            core.ActionID{ItemID: 17111},
					Character:           character,
					IsPhantom:           true,
					CritRollCategory:    core.CritRollCategoryMagical,
					OutcomeRollCategory: core.OutcomeRollCategoryMagic,
					SpellSchool:         core.SpellSchoolFire,
					CritMultiplier:      character.DefaultSpellCritMultiplier(),
				},
			},
			Effect: core.SpellHitEffect{
				SpellEffect: core.SpellEffect{
					DamageMultiplier:       1,
					StaticDamageMultiplier: 1,
					ThreatMultiplier:       1,
				},
				DirectInput: core.DirectDamageInput{
					MinBaseDamage: 2,
					MaxBaseDamage: 2,
				},
			},
		})

		return core.Aura{
			ID: BlazefuryMedallionAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				// https://tbc.wowhead.com/spell=7711/add-fire-dam-weap-02, proc mask = 20.
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellCast.IsPhantom {
					return
				}

				castAction := &spellObj
				castTemplate.Apply(castAction)
				castAction.Effect.Target = spellEffect.Target
				castAction.Init(sim)
				castAction.Cast(sim)
			},
		}
	})
}

var EmpyreanDemolisherAuraID = core.NewAuraID()
var EmpyreanDemolisherProcAuraID = core.NewAuraID()

func ApplyEmpyreanDemolisher(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(17112)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 212.0
		const dur = time.Second * 10
		const procChance = 2.8 / 60.0
		applyStatAura := character.NewTemporaryStatsAuraApplier(EmpyreanDemolisherProcAuraID, core.ActionID{ItemID: 17112}, stats.Stats{stats.MeleeHaste: hasteBonus}, dur)

		return core.Aura{
			ID: EmpyreanDemolisherAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("EmpyreanDemolisher") > procChance {
					return
				}

				applyStatAura(sim)
			},
		}
	})
}

var KhoriumChampionAuraID = core.NewAuraID()
var KhoriumChampionProcAuraID = core.NewAuraID()

func ApplyKhoriumChampion(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const strBonus = 120.0
		const dur = time.Second * 30
		const procChance = 0.5 * 3.3 / 60.0
		applyStatAura := character.NewTemporaryStatsAuraApplier(KhoriumChampionProcAuraID, core.ActionID{ItemID: 23541}, stats.Stats{stats.Strength: strBonus}, dur)

		return core.Aura{
			ID: KhoriumChampionAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				// https://tbc.wowhead.com/spell=16916/strength-of-the-champion, proc mask = 0. Handled in-game via script.
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("KhoriumChampion") > procChance {
					return
				}

				applyStatAura(sim)
			},
		}
	})
}

func ApplyBraidedEterniumChain(agent core.Agent) {
	agent.GetCharacter().PseudoStats.BonusDamage += 5
}

var BlackoutTruncheonAuraID = core.NewAuraID()
var BlackoutTruncheonProcAuraID = core.NewAuraID()

func ApplyBlackoutTruncheon(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(27901)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 132.0
		const dur = time.Second * 10
		const procChance = 1.5 * 0.8 / 60.0
		applyStatAura := character.NewTemporaryStatsAuraApplier(BlackoutTruncheonProcAuraID, core.ActionID{ItemID: 27901}, stats.Stats{stats.MeleeHaste: hasteBonus}, dur)

		return core.Aura{
			ID: BlackoutTruncheonAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				// https://tbc.wowhead.com/spell=33489/blinding-speed, proc mask = 0. Handled in-game via script.
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("BlackoutTruncheon") > procChance {
					return
				}

				applyStatAura(sim)
			},
		}
	})
}

var LionheartChampionAuraID = core.NewAuraID()
var LionheartChampionProcAuraID = core.NewAuraID()

func ApplyLionheartChampion(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const strBonus = 100.0
		const dur = time.Second * 10
		const procChance = 3.6 / 60.0
		applyStatAura := character.NewTemporaryStatsAuraApplier(LionheartChampionProcAuraID, core.ActionID{ItemID: 28429}, stats.Stats{stats.Strength: strBonus}, dur)

		return core.Aura{
			ID: LionheartChampionAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				// https://tbc.wowhead.com/spell=34513/lionheart, proc mask = 0. Handled in-game via script.
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("LionheartChampion") > procChance {
					return
				}

				applyStatAura(sim)
			},
		}
	})
}

var LionheartExecutionerAuraID = core.NewAuraID()
var LionheartExecutionerProcAuraID = core.NewAuraID()

func ApplyLionheartExecutioner(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const strBonus = 100.0
		const dur = time.Second * 10
		const procChance = 3.6 / 60.0
		applyStatAura := character.NewTemporaryStatsAuraApplier(LionheartExecutionerProcAuraID, core.ActionID{ItemID: 28430}, stats.Stats{stats.Strength: strBonus}, dur)

		return core.Aura{
			ID: LionheartExecutionerAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("LionheartExecutioner") > procChance {
					return
				}

				applyStatAura(sim)
			},
		}
	})
}

var DrakefistHammerAuraID = core.NewAuraID()
var DrakefistHammerProcAuraID = core.NewAuraID()

func ApplyDrakefistHammer(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(28437)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 212.0
		const dur = time.Second * 10
		const procChance = 2.7 / 60.0
		applyStatAura := character.NewTemporaryStatsAuraApplier(DrakefistHammerProcAuraID, core.ActionID{ItemID: 28437}, stats.Stats{stats.MeleeHaste: hasteBonus}, dur)

		return core.Aura{
			ID: DrakefistHammerAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("DrakefistHammer") > procChance {
					return
				}

				applyStatAura(sim)
			},
		}
	})
}

var DragonmawAuraID = core.NewAuraID()
var DragonmawProcAuraID = core.NewAuraID()

func ApplyDragonmaw(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(28438)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 212.0
		const dur = time.Second * 10
		const procChance = 2.7 / 60.0
		applyStatAura := character.NewTemporaryStatsAuraApplier(DragonmawProcAuraID, core.ActionID{ItemID: 28438}, stats.Stats{stats.MeleeHaste: hasteBonus}, dur)

		return core.Aura{
			ID: DragonmawAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("Dragonmaw") > procChance {
					return
				}

				applyStatAura(sim)
			},
		}
	})
}

var DragonstrikeAuraID = core.NewAuraID()
var DragonstrikeProcMHAuraID = core.NewAuraID()
var DragonstrikeProcOHAuraID = core.NewAuraID()

func ApplyDragonstrike(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(28439)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 212.0
		const dur = time.Second * 10
		const procChance = 2.7 / 60.0
		applyStatAura := character.NewTemporaryStatsAuraApplier(DragonstrikeProcMHAuraID, core.ActionID{ItemID: 28439}, stats.Stats{stats.MeleeHaste: hasteBonus}, dur)

		return core.Aura{
			ID: DragonstrikeAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("Dragonstrike") > procChance {
					return
				}

				applyStatAura(sim)
			},
		}
	})
}

func ApplyDragonstrikeStacking(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(-23)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 212.0
		const dur = time.Second * 10
		const procChance = 2.7 / 60.0
		applyStatMHAura := character.NewTemporaryStatsAuraApplier(DragonstrikeProcMHAuraID, core.ActionID{ItemID: 28439, Tag: 1}, stats.Stats{stats.MeleeHaste: hasteBonus}, dur)
		applyStatOHAura := character.NewTemporaryStatsAuraApplier(DragonstrikeProcOHAuraID, core.ActionID{ItemID: 28439, Tag: 2}, stats.Stats{stats.MeleeHaste: hasteBonus}, dur)

		return core.Aura{
			ID: DragonstrikeAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("Dragonstrike") > procChance {
					return
				}

				if spellEffect.IsMH() {
					applyStatMHAura(sim)
				} else {
					applyStatOHAura(sim)
				}
			},
		}
	})
}

var DespairAuraID = core.NewAuraID()

func ApplyDespair(agent core.Agent) {
	character := agent.GetCharacter()
	actionID := core.ActionID{SpellID: 34580}

	templ := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            actionID,
				Character:           character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				SpellExtras:         core.SpellExtrasIgnoreResists,
				CritMultiplier:      character.DefaultMeleeCritMultiplier(),
				IsPhantom:           true,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				// TODO: This should be removed once we have an attack mask.
				//  This is only set here to correctly calculate damage.
				ProcMask:               core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			DirectInput: core.DirectDamageInput{
				FlatDamageBonus: 600,
			},
		},
	}

	spellCastTemplate := core.NewSimpleSpellTemplate(templ)
	cast := core.SimpleSpell{}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const procChance = 0.5 * 3.5 / 60.0

		return core.Aura{
			ID: DespairAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				// ProcMask: 20
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("Despair") > procChance {
					return
				}

				spellCastTemplate.Apply(&cast)
				cast.Effect.Target = sim.GetPrimaryTarget()
				cast.Cast(sim)
			},
		}
	})
}

var TheDecapitatorCooldownID = core.NewCooldownID()

func ApplyTheDecapitator(agent core.Agent) {
	character := agent.GetCharacter()
	actionID := core.ActionID{ItemID: 28767}

	templ := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            actionID,
				Character:           character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				SpellExtras:         core.SpellExtrasIgnoreResists,
				CritMultiplier:      character.DefaultMeleeCritMultiplier(),
				IsPhantom:           true,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage: 513,
				MaxBaseDamage: 567,
			},
		},
	}

	spellCastTemplate := core.NewSimpleSpellTemplate(templ)
	spellCast := core.SimpleSpell{}

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
				spellCastTemplate.Apply(&spellCast)
				spellCast.Effect.Target = sim.GetPrimaryTarget()
				spellCast.Cast(sim)

				character.SetCD(TheDecapitatorCooldownID, sim.CurrentTime+time.Minute*3)
				character.SetCD(core.OffensiveTrinketSharedCooldownID, sim.CurrentTime+time.Second*10)
			}
		},
	})
}

var GlaiveOfThePitAuraID = core.NewAuraID()
var GlaiveOfThePitProcAuraID = core.NewAuraID()

func ApplyGlaiveOfThePit(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 212.0
		const procChance = 3.7 / 60.0

		spellObj := core.SimpleSpell{}
		castTemplate := core.NewSimpleSpellTemplate(core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:            core.ActionID{SpellID: 34696},
					Character:           character,
					IsPhantom:           true,
					CritRollCategory:    core.CritRollCategoryMagical,
					OutcomeRollCategory: core.OutcomeRollCategoryMagic,
					SpellSchool:         core.SpellSchoolShadow,
					CritMultiplier:      character.DefaultSpellCritMultiplier(),
				},
			},
			Effect: core.SpellHitEffect{
				SpellEffect: core.SpellEffect{
					DamageMultiplier:       1,
					StaticDamageMultiplier: 1,
					ThreatMultiplier:       1,
				},
				DirectInput: core.DirectDamageInput{
					MinBaseDamage: 285,
					MaxBaseDamage: 315,
				},
			},
		})

		return core.Aura{
			ID: GlaiveOfThePitAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("GlaiveOfThePit") > procChance {
					return
				}

				castAction := &spellObj
				castTemplate.Apply(castAction)
				castAction.Effect.Target = spellEffect.Target
				castAction.Init(sim)
				castAction.Cast(sim)
			},
		}
	})
}

var BandOfTheEternalChampionAuraID = core.NewAuraID()
var BandOfTheEternalChampionProcAuraID = core.NewAuraID()

func ApplyBandOfTheEternalChampion(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const apBonus = 160
		const dur = time.Second * 10
		ppmm := character.AutoAttacks.NewPPMManager(1.0)
		applyStatAura := character.NewTemporaryStatsAuraApplier(BandOfTheEternalChampionProcAuraID, core.ActionID{ItemID: 29301}, stats.Stats{stats.AttackPower: apBonus, stats.RangedAttackPower: apBonus}, dur)

		icd := core.NewICD()
		const icdDur = time.Second * 60

		return core.Aura{
			ID: BandOfTheEternalChampionAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				// mask 340
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellCast.IsPhantom {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if !ppmm.Proc(sim, spellEffect.IsMH(), spellCast.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged), "Band of the Eternal Champion") {
					return
				}

				icd = core.InternalCD(sim.CurrentTime + icdDur)
				applyStatAura(sim)
			},
		}
	})
}

var TheBladefistAuraID = core.NewAuraID()
var TheBladefistProcAuraID = core.NewAuraID()

func ApplyTheBladefist(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 180.0
		const dur = time.Second * 10
		const procChance = 2.7 / 60.0
		applyStatAura := character.NewTemporaryStatsAuraApplier(TheBladefistProcAuraID, core.ActionID{ItemID: 29348}, stats.Stats{stats.MeleeHaste: hasteBonus}, dur)

		return core.Aura{
			ID: TheBladefistAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeMH) {
					return
				}
				if sim.RandomFloat("The Bladefist") > procChance {
					return
				}

				applyStatAura(sim)
			},
		}
	})
}

var HeartrazorAuraID = core.NewAuraID()
var HeartrazorProcAuraID = core.NewAuraID()

func ApplyHeartrazor(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(29962)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const apBonus = 270
		const dur = time.Second * 10
		applyStatAura := character.NewTemporaryStatsAuraApplier(HeartrazorProcAuraID, core.ActionID{ItemID: 29962}, stats.Stats{stats.AttackPower: apBonus, stats.RangedAttackPower: apBonus}, dur)

		ppmm := character.AutoAttacks.NewPPMManager(1.0)

		return core.Aura{
			ID: HeartrazorAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) || spellCast.IsPhantom {
					return
				}

				if !ppmm.Proc(sim, spellEffect.IsMH(), false, "Heartrazor") {
					return
				}

				applyStatAura(sim)
			},
		}
	})
}

var RodOfTheSunKingAuraID = core.NewAuraID()

func ApplyRodOfTheSunKing(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(29996)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const procChance = 2.7 / 60.0
		actionID := core.ActionID{ItemID: 29996}

		return core.Aura{
			ID: RodOfTheSunKingAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}

				if spellCast.Character.HasRageBar() {
					if sim.RandomFloat("Rod of the Sun King") > procChance {
						return
					}
					spellCast.Character.AddRage(sim, 5, actionID)
				} else if spellCast.Character.HasEnergyBar() {
					if sim.RandomFloat("Rod of the Sun King") > procChance {
						return
					}
					spellCast.Character.AddEnergy(sim, 10, actionID)
				}
			},
		}
	})
}

var WorldBreakerAuraID = core.NewAuraID()
var WorldBreakerProcAuraID = core.NewAuraID()

func ApplyWorldBreaker(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const critBonus = 900.0
		const dur = time.Second * 4
		const procChance = 3.7 / 60.0
		applyStatAura := character.NewTemporaryStatsAuraApplier(WorldBreakerProcAuraID, core.ActionID{ItemID: 30090}, stats.Stats{stats.MeleeCrit: critBonus}, dur)

		return core.Aura{
			ID: WorldBreakerAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					if character.HasAura(WorldBreakerProcAuraID) {
						character.RemoveAura(sim, WorldBreakerProcAuraID)
					}
					return
				}
				if sim.RandomFloat("World Breaker") > procChance {
					if character.HasAura(WorldBreakerProcAuraID) {
						character.RemoveAura(sim, WorldBreakerProcAuraID)
					}
					return
				}

				applyStatAura(sim)
			},
		}
	})
}

var WarpSlicerAuraID = core.NewAuraID()
var WarpSlicerProcAuraID = core.NewAuraID()

func ApplyWarpSlicer(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(30311)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const bonus = 1.2
		const inverseBonus = 1 / 1.2
		const dur = time.Second * 30
		const procChance = 0.5

		return core.Aura{
			ID: WarpSlicerAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("WarpSlicer") > procChance {
					return
				}

				character.MultiplyMeleeSpeed(sim, bonus)
				character.ReplaceAura(sim, core.Aura{
					ID:       WarpSlicerProcAuraID,
					ActionID: core.ActionID{ItemID: 30311},
					Expires:  sim.CurrentTime + dur,
					OnExpire: func(sim *core.Simulation) {
						character.MultiplyMeleeSpeed(sim, inverseBonus)
					},
				})
			},
		}
	})
}

var DevastationAuraID = core.NewAuraID()
var DevastationProcAuraID = core.NewAuraID()

func ApplyDevastation(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const bonus = 1.2
		const inverseBonus = 1 / 1.2
		const dur = time.Second * 30
		const procChance = 0.5

		return core.Aura{
			ID: DevastationAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("Devastation") > procChance {
					return
				}

				if !character.HasAura(DevastationProcAuraID) {
					character.MultiplyMeleeSpeed(sim, bonus)
				}

				character.ReplaceAura(sim, core.Aura{
					ID:       DevastationProcAuraID,
					ActionID: core.ActionID{ItemID: 30316},
					Expires:  sim.CurrentTime + dur,
					OnExpire: func(sim *core.Simulation) {
						character.MultiplyMeleeSpeed(sim, inverseBonus)
					},
				})
			},
		}
	})
}

var BladeOfUnquenchedThirstAuraID = core.NewAuraID()

func ApplyBladeOfUnquenchedThirst(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(31193)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const procChance = 0.02

		castTemplate := core.NewSimpleSpellTemplate(core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:            core.ActionID{ItemID: 31193},
					Character:           character,
					IsPhantom:           true,
					CritRollCategory:    core.CritRollCategoryMagical,
					OutcomeRollCategory: core.OutcomeRollCategoryMagic,
					SpellSchool:         core.SpellSchoolShadow,
					CritMultiplier:      character.DefaultSpellCritMultiplier(),
				},
			},
			Effect: core.SpellHitEffect{
				SpellEffect: core.SpellEffect{
					DamageMultiplier:       1,
					StaticDamageMultiplier: 1,
					ThreatMultiplier:       1,
				},
				DirectInput: core.DirectDamageInput{
					MinBaseDamage:    48,
					MaxBaseDamage:    54,
					SpellCoefficient: 1,
				},
			},
		})
		spellObj := core.SimpleSpell{}

		return core.Aura{
			ID: BladeOfUnquenchedThirstAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("BladeOfUnquenchedThirst") > procChance {
					return
				}

				castAction := &spellObj
				castTemplate.Apply(castAction)
				castAction.Effect.Target = spellEffect.Target
				castAction.Init(sim)
				castAction.Cast(sim)
			},
		}
	})
}

var SingingCrystalAxeAuraID = core.NewAuraID()
var SingingCrystalAxeProcAuraID = core.NewAuraID()

func ApplySingingCrystalAxe(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 400.0
		const dur = time.Second * 10
		const procChance = 3.5 / 60.0
		applyStatAura := character.NewTemporaryStatsAuraApplier(SingingCrystalAxeProcAuraID, core.ActionID{ItemID: 31318}, stats.Stats{stats.MeleeHaste: hasteBonus}, dur)

		return core.Aura{
			ID: SingingCrystalAxeAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("SingingCrystalAxe") > procChance {
					return
				}

				applyStatAura(sim)
			},
		}
	})
}

var BlinkstrikeAuraID = core.NewAuraID()

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

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		var icd core.InternalCD
		icdDur := time.Millisecond * 1

		mhAttack := character.AutoAttacks.MHAuto
		mhAttack.ActionID = core.ActionID{ItemID: 31332}
		cachedAttack := core.SimpleSpell{}

		return core.Aura{
			ID: BlinkstrikeAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) || spellCast.IsPhantom {
					return
				}

				if icd.IsOnCD(sim) {
					return
				}

				if !ppmm.Proc(sim, spellEffect.IsMH(), false, "Blinkstrike") {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)

				cachedAttack = mhAttack
				cachedAttack.Effect.Target = spellEffect.Target
				cachedAttack.Cast(sim)
			},
		}
	})
}

var TheNightBladeAuraID = core.NewAuraID()
var TheNightBladeProcAuraID = core.NewAuraID()

func ApplyTheNightBlade(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(31331)
	procMask := core.GetMeleeProcMaskForHands(mh, oh)
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const arPenBonus = 435.0
		const dur = time.Second * 10
		const procChance = 2 * 1.8 / 60.0

		return core.Aura{
			ID: TheNightBladeAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("The Night Blade") > procChance {
					return
				}

				stacks := character.NumStacks(TheNightBladeProcAuraID) + 1
				newBonus := arPenBonus * float64(stacks)
				character.ReplaceAura(sim, core.Aura{
					ID:       TheNightBladeProcAuraID,
					ActionID: core.ActionID{ItemID: 31331},
					Expires:  sim.CurrentTime + dur,
					Stacks:   stacks,
					OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
						spellEffect.BonusArmorPenetration += newBonus
					},
				})
			},
		}
	})
}

var SyphonOfTheNathrezimAuraID = core.NewAuraID()

var SiphonEssenceAuraID = core.NewAuraID()

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

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		castTemplate := core.NewSimpleSpellTemplate(core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:            core.ActionID{SpellID: 40291},
					Character:           character,
					IsPhantom:           true,
					CritRollCategory:    core.CritRollCategoryMagical,
					OutcomeRollCategory: core.OutcomeRollCategoryMagic,
					SpellSchool:         core.SpellSchoolShadow,
					CritMultiplier:      character.DefaultSpellCritMultiplier(),
				},
			},
			Effect: core.SpellHitEffect{
				SpellEffect: core.SpellEffect{
					DamageMultiplier:       1,
					StaticDamageMultiplier: 1,
					ThreatMultiplier:       1,
				},
				DirectInput: core.DirectDamageInput{
					MinBaseDamage: 20,
					MaxBaseDamage: 20,
				},
			},
		})
		spellObj := core.SimpleSpell{}

		procAura := core.Aura{
			ID:       SiphonEssenceAuraID,
			ActionID: core.ActionID{SpellID: 40291},
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellCast.IsPhantom {
					return
				}

				castAction := &spellObj
				castTemplate.Apply(castAction)
				castAction.Effect.Target = spellEffect.Target
				castAction.Init(sim)
				castAction.Cast(sim)
			},
		}

		return core.Aura{
			ID: SyphonOfTheNathrezimAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}

				if ppmm.Proc(sim, spellEffect.IsMH(), false, "Syphon Of The Nathrezim") {
					aura := procAura
					aura.Expires = sim.CurrentTime + time.Second*6
					character.ReplaceAura(sim, aura)
				}
			},
		}
	})
}

var CloakOfDarknessAuraID = core.NewAuraID()

func ApplyCloakOfDarkness(agent core.Agent) {
	character := agent.GetCharacter()

	// The melee distinction only matters for hunters, so for others we can skip
	// having a separate aura.
	if character.Class != proto.Class_ClassHunter {
		character.AddStat(stats.MeleeCrit, 24)
		return
	}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: CloakOfDarknessAuraID,
			OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
				if !spellCast.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged) {
					spellEffect.BonusCritRating += 24
				}
			},
		}
	})
}
