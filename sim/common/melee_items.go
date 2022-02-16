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
	core.AddItemEffect(23541, ApplyKhoriumChampion)
	core.AddItemEffect(24114, ApplyBraidedEterniumChain)
	core.AddItemEffect(27901, ApplyBlackoutTruncheon)
	core.AddItemEffect(28429, ApplyLionheartChampion)
	core.AddItemEffect(28430, ApplyLionheartExecutioner)
	core.AddItemEffect(28437, ApplyDrakefistHammer)
	core.AddItemEffect(28438, ApplyDragonmaw)
	core.AddItemEffect(28439, ApplyDragonstrike)
	core.AddItemEffect(-23, ApplyDragonstrike)
	core.AddItemEffect(28573, ApplyDespair)
	core.AddItemEffect(28767, ApplyTheDecapitator)
	core.AddItemEffect(28774, ApplyGlaiveOfThePit)
	core.AddItemEffect(29301, ApplyBandOfTheEternalChampion)
	core.AddItemEffect(29348, ApplyTheBladefist)
	core.AddItemEffect(29996, ApplyRodOfTheSunKing)
	core.AddItemEffect(30090, ApplyWorldBreaker)
	core.AddItemEffect(30311, ApplyWarpSlicer)
	core.AddItemEffect(30316, ApplyDevastation)
	core.AddItemEffect(31318, ApplySingingCrystalAxe)
	core.AddItemEffect(31331, ApplyTheNightBlade)
	core.AddItemEffect(32262, ApplySyphonOfTheNathrezim)
	core.AddItemEffect(33122, ApplyCloakOfDarkness)

	// TODO:
	// blinkstrike
}

var StormGauntletsAuraID = core.NewAuraID()

func ApplyStormGauntlets(agent core.Agent) {
	character := agent.GetCharacter()
	spellObj := core.SimpleSpell{}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		castTemplate := core.NewSimpleSpellTemplate(core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:       core.ActionID{ItemID: 12632},
					Character:      character,
					IgnoreManaCost: true,
					IsPhantom:      true,
					SpellSchool:    stats.NatureSpellPower,
					CritMultiplier: character.DefaultSpellCritMultiplier(),
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
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				// https://tbc.wowhead.com/spell=16615/add-lightning-dam-weap-03, proc mask = 20.
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) || ability.IsPhantom {
					return
				}

				castAction := &spellObj
				castTemplate.Apply(castAction)
				castAction.Effect.Target = hitEffect.Target
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
					ActionID:       core.ActionID{ItemID: 17111},
					Character:      character,
					IgnoreManaCost: true,
					IsPhantom:      true,
					SpellSchool:    stats.FireSpellPower,
					CritMultiplier: character.DefaultSpellCritMultiplier(),
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
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				// https://tbc.wowhead.com/spell=7711/add-fire-dam-weap-02, proc mask = 20.
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) || ability.IsPhantom {
					return
				}

				castAction := &spellObj
				castTemplate.Apply(castAction)
				castAction.Effect.Target = hitEffect.Target
				castAction.Init(sim)
				castAction.Cast(sim)
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

		return core.Aura{
			ID: KhoriumChampionAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				// https://tbc.wowhead.com/spell=16916/strength-of-the-champion, proc mask = 0. Handled in-game via script.
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) || ability.IsPhantom {
					return
				}
				if sim.RandomFloat("KhoriumChampion") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, KhoriumChampionProcAuraID, core.ActionID{ItemID: 23541}, stats.Strength, strBonus, dur)
			},
		}
	})
}

func ApplyBraidedEterniumChain(agent core.Agent) {
	agent.GetCharacter().PseudoStats.BonusMeleeDamage += 5
	agent.GetCharacter().PseudoStats.BonusRangedDamage += 5
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

		return core.Aura{
			ID: BlackoutTruncheonAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				// https://tbc.wowhead.com/spell=33489/blinding-speed, proc mask = 0. Handled in-game via script.
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(procMask) || ability.IsPhantom {
					return
				}
				if sim.RandomFloat("BlackoutTruncheon") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, BlackoutTruncheonProcAuraID, core.ActionID{ItemID: 27901}, stats.MeleeHaste, hasteBonus, dur)
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

		return core.Aura{
			ID: LionheartChampionAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				// https://tbc.wowhead.com/spell=34513/lionheart, proc mask = 0. Handled in-game via script.
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("LionheartChampion") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, LionheartChampionProcAuraID, core.ActionID{ItemID: 28429}, stats.Strength, strBonus, dur)
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

		return core.Aura{
			ID: LionheartExecutionerAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("LionheartExecutioner") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, LionheartExecutionerProcAuraID, core.ActionID{ItemID: 28430}, stats.Strength, strBonus, dur)
			},
		}
	})
}

var DrakefistHammerAuraID = core.NewAuraID()
var DrakefistHammerProcAuraID = core.NewAuraID()

func ApplyDrakefistHammer(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(28437)
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 212.0
		const dur = time.Second * 10
		const procChance = 2.7 / 60.0

		return core.Aura{
			ID: DrakefistHammerAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) || !hitEffect.IsEquippedHand(mh, oh) {
					return
				}
				if sim.RandomFloat("DrakefistHammer") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, DrakefistHammerProcAuraID, core.ActionID{ItemID: 28437}, stats.MeleeHaste, hasteBonus, dur)
			},
		}
	})
}

var DragonmawAuraID = core.NewAuraID()
var DragonmawProcAuraID = core.NewAuraID()

func ApplyDragonmaw(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(28438)
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 212.0
		const dur = time.Second * 10
		const procChance = 2.7 / 60.0

		return core.Aura{
			ID: DragonmawAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) || !hitEffect.IsEquippedHand(mh, oh) {
					return
				}
				if sim.RandomFloat("Dragonmaw") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, DragonmawProcAuraID, core.ActionID{ItemID: 28438}, stats.MeleeHaste, hasteBonus, dur)
			},
		}
	})
}

var DragonstrikeAuraID = core.NewAuraID()
var DragonstrikeProcAuraID = core.NewAuraID()

func ApplyDragonstrike(agent core.Agent) {
	character := agent.GetCharacter()
	mh, _ := character.GetWeaponHands(28439)
	_, oh := character.GetWeaponHands(-23)
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 212.0
		const dur = time.Second * 10
		const procChance = 2.7 / 60.0

		return core.Aura{
			ID: DragonstrikeAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) || !hitEffect.IsEquippedHand(mh, oh) {
					return
				}
				if sim.RandomFloat("Dragonstrike") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, DragonstrikeProcAuraID, core.ActionID{ItemID: 28439}, stats.MeleeHaste, hasteBonus, dur)
			},
		}
	})
}

var DespairAuraID = core.NewAuraID()

func ApplyDespair(agent core.Agent) {
	character := agent.GetCharacter()
	actionID := core.ActionID{SpellID: 34580}

	templ := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			ActionID:       actionID,
			Character:      character,
			SpellSchool:    stats.AttackPower,
			CritMultiplier: character.DefaultMeleeCritMultiplier(),
			IgnoreCost:     true,
			IsPhantom:      true,
		},
		Effect: core.AbilityHitEffect{
			AbilityEffect: core.AbilityEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			DirectInput: core.DirectDamageInput{
				FlatDamageBonus: 600,
			},
		},
	}

	abilityTemplate := core.NewMeleeAbilityTemplate(templ)
	cast := core.ActiveMeleeAbility{}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const procChance = 0.5 * 3.5 / 60.0

		return core.Aura{
			ID: DespairAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("Despair") > procChance {
					return
				}

				abilityTemplate.Apply(&cast)
				cast.Effect.Target = sim.GetPrimaryTarget()
				cast.Attack(sim)
			},
		}
	})
}

var TheDecapitatorCooldownID = core.NewCooldownID()

func ApplyTheDecapitator(agent core.Agent) {
	character := agent.GetCharacter()
	actionID := core.ActionID{ItemID: 28767}

	templ := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			ActionID:       actionID,
			Character:      character,
			SpellSchool:    stats.AttackPower,
			CritMultiplier: character.DefaultMeleeCritMultiplier(),
			IgnoreCost:     true,
			IsPhantom:      true,
		},
		Effect: core.AbilityHitEffect{
			AbilityEffect: core.AbilityEffect{
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

	abilityTemplate := core.NewMeleeAbilityTemplate(templ)
	ability := core.ActiveMeleeAbility{}

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
				abilityTemplate.Apply(&ability)
				ability.Effect.Target = sim.GetPrimaryTarget()
				ability.Attack(sim)

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
					ActionID:       core.ActionID{SpellID: 34696},
					Character:      character,
					IgnoreManaCost: true,
					IsPhantom:      true,
					SpellSchool:    stats.ShadowSpellPower,
					CritMultiplier: character.DefaultSpellCritMultiplier(),
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
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("GlaiveOfThePit") > procChance {
					return
				}

				castAction := &spellObj
				castTemplate.Apply(castAction)
				castAction.Effect.Target = hitEffect.Target
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

		icd := core.NewICD()
		const icdDur = time.Second * 60

		return core.Aura{
			ID: BandOfTheEternalChampionAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) || ability.IsPhantom {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if !ppmm.Proc(sim, hitEffect.IsMH(), hitEffect.IsRanged(), "Band of the Eternal Champion") {
					return
				}

				icd = core.InternalCD(sim.CurrentTime + icdDur)
				character.AddAuraWithTemporaryStats(sim, BandOfTheEternalChampionProcAuraID, core.ActionID{ItemID: 29301}, stats.AttackPower, apBonus, dur)
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

		return core.Aura{
			ID: TheBladefistAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) || !hitEffect.IsMH() {
					return
				}
				if sim.RandomFloat("The Bladefist") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, TheBladefistProcAuraID, core.ActionID{ItemID: 29348}, stats.MeleeHaste, hasteBonus, dur)
			},
		}
	})
}

var RodOfTheSunKingAuraID = core.NewAuraID()

func ApplyRodOfTheSunKing(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(29996)
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const procChance = 2.7 / 60.0
		actionID := core.ActionID{ItemID: 29996}

		return core.Aura{
			ID: RodOfTheSunKingAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) || !hitEffect.IsEquippedHand(mh, oh) {
					return
				}

				if ability.Character.HasRageBar() {
					if sim.RandomFloat("Rod of the Sun King") > procChance {
						return
					}
					ability.Character.AddRage(sim, 5, actionID)
				} else if ability.Character.HasEnergyBar() {
					if sim.RandomFloat("Rod of the Sun King") > procChance {
						return
					}
					ability.Character.AddEnergy(sim, 10, actionID)
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

		return core.Aura{
			ID: WorldBreakerAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) {
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

				character.AddAuraWithTemporaryStats(sim, WorldBreakerProcAuraID, core.ActionID{ItemID: 30090}, stats.MeleeCrit, critBonus, dur)
			},
		}
	})
}

var WarpSlicerAuraID = core.NewAuraID()
var WarpSlicerProcAuraID = core.NewAuraID()

func ApplyWarpSlicer(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(30311)
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const bonus = 1.2
		const inverseBonus = 1 / 1.2
		const dur = time.Second * 30
		const procChance = 0.5

		return core.Aura{
			ID: WarpSlicerAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) || !hitEffect.IsEquippedHand(mh, oh) {
					return
				}
				if sim.RandomFloat("WarpSlicer") > procChance {
					return
				}

				character.MultiplyMeleeSpeed(sim, bonus)
				character.AddAura(sim, core.Aura{
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
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("Devastation") > procChance {
					return
				}

				character.MultiplyMeleeSpeed(sim, bonus)
				character.AddAura(sim, core.Aura{
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

var SingingCrystalAxeAuraID = core.NewAuraID()
var SingingCrystalAxeProcAuraID = core.NewAuraID()

func ApplySingingCrystalAxe(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 400.0
		const dur = time.Second * 10
		const procChance = 3.5 / 60.0

		return core.Aura{
			ID: SingingCrystalAxeAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("SingingCrystalAxe") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, SingingCrystalAxeProcAuraID, core.ActionID{ItemID: 31318}, stats.MeleeHaste, hasteBonus, dur)
			},
		}
	})
}

var TheNightBladeAuraID = core.NewAuraID()
var TheNightBladeProcAuraID = core.NewAuraID()

func ApplyTheNightBlade(agent core.Agent) {
	character := agent.GetCharacter()
	mh, oh := character.GetWeaponHands(31331)
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const arPenBonus = 435.0
		const dur = time.Second * 10
		const procChance = 2 * 1.8 / 60.0

		return core.Aura{
			ID: TheNightBladeAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) || !hitEffect.IsEquippedHand(mh, oh) {
					return
				}
				if sim.RandomFloat("The Night Blade") > procChance {
					return
				}

				stacks := character.NumStacks(TheNightBladeProcAuraID) + 1
				newBonus := arPenBonus * float64(stacks)
				character.AddAura(sim, core.Aura{
					ID:       TheNightBladeProcAuraID,
					ActionID: core.ActionID{ItemID: 31331},
					Expires:  sim.CurrentTime + dur,
					Stacks:   stacks,
					OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
						hitEffect.BonusArmorPenetration += newBonus
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
		ppmm.SetProcChance(false, 0)
	}
	if !oh {
		ppmm.SetProcChance(true, 0)
	}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		castTemplate := core.NewSimpleSpellTemplate(core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:       core.ActionID{SpellID: 40291},
					Character:      character,
					IgnoreManaCost: true,
					IsPhantom:      true,
					SpellSchool:    stats.ShadowSpellPower,
					CritMultiplier: character.DefaultSpellCritMultiplier(),
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
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) || ability.IsPhantom {
					return
				}

				castAction := &spellObj
				castTemplate.Apply(castAction)
				castAction.Effect.Target = hitEffect.Target
				castAction.Init(sim)
				castAction.Cast(sim)
			},
		}

		return core.Aura{
			ID: SyphonOfTheNathrezimAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}

				if ppmm.Proc(sim, hitEffect.IsMH(), false, "Syphon Of The Nathrezim") {
					aura := procAura
					aura.Expires = sim.CurrentTime + time.Second*6
					character.AddAura(sim, aura)
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
			OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.IsRanged() {
					hitEffect.BonusCritRating += 24
				}
			},
		}
	})
}
