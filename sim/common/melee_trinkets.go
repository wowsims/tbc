package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	// Proc effects. Keep these in order by item ID.
	core.AddItemEffect(21670, ApplyBadgeOfTheSwarmguard)
	core.AddItemEffect(23206, ApplyMarkOfTheChampionMelee)
	core.AddItemEffect(28034, ApplyHourglassUnraveller)
	core.AddItemEffect(28579, ApplyRomulosPoisonVial)
	core.AddItemEffect(28830, ApplyDragonspineTrophy)
	core.AddItemEffect(30627, ApplyTsunamiTalisman)
	core.AddItemEffect(31857, ApplyDarkmoonCardWrath)
	core.AddItemEffect(32505, ApplyMadnessOfTheBetrayer)
	core.AddItemEffect(34427, ApplyBlackenedNaaruSliver)
	core.AddItemEffect(34472, ApplyShardOfContempt)

	// Activatable effects. Keep these in order by item ID.
	var KissOfTheSpiderCooldownID = core.NewCooldownID()
	core.AddItemEffect(22954, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		stats.MeleeHaste,
		200,
		time.Second*15,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 22954},
			CooldownID:       KissOfTheSpiderCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var SlayersCrestCooldownID = core.NewCooldownID()
	core.AddItemEffect(23041, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		stats.AttackPower,
		260,
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 23041},
			CooldownID:       SlayersCrestCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var FigurineNightseyePantherCooldownID = core.NewCooldownID()
	core.AddItemEffect(24128, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		stats.AttackPower,
		320,
		time.Second*12,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 24128},
			CooldownID:       FigurineNightseyePantherCooldownID,
			Cooldown:         time.Minute * 3,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var BladefistsBreadthCooldownID = core.NewCooldownID()
	core.AddItemEffect(28041, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		stats.AttackPower,
		200,
		time.Second*15,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 28041},
			CooldownID:       BladefistsBreadthCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var IconOfUnyieldingCourageCooldownID = core.NewCooldownID()
	core.AddItemEffect(28121, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		stats.ArmorPenetration,
		600,
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 28121},
			CooldownID:       IconOfUnyieldingCourageCooldownID,
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

	var CoreOfArkelosCooldownID = core.NewCooldownID()
	core.AddItemEffect(29776, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		stats.AttackPower,
		200,
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 29776},
			CooldownID:       CoreOfArkelosCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	core.AddItemEffect(32654, ApplyCrystalforgedTrinket)

	var BadgeOfTenacityCooldownID = core.NewCooldownID()
	core.AddItemEffect(32658, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		stats.Agility,
		150,
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 32658},
			CooldownID:       BadgeOfTenacityCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var BerserkersCallCooldownID = core.NewCooldownID()
	core.AddItemEffect(33831, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		stats.AttackPower,
		360,
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 33831},
			CooldownID:       BerserkersCallCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var FigurineShadowsongPantherCooldownID = core.NewCooldownID()
	core.AddItemEffect(35702, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		stats.AttackPower,
		320,
		time.Second*15,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 35702},
			CooldownID:       FigurineShadowsongPantherCooldownID,
			Cooldown:         time.Second * 90,
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

var CrystalforgedTrinketCooldownID = core.NewCooldownID()

func ApplyCrystalforgedTrinket(agent core.Agent) {
	agent.GetCharacter().PseudoStats.BonusMeleeDamage += 7
	agent.GetCharacter().PseudoStats.BonusRangedDamage += 7
	core.RegisterTemporaryStatsOnUseCD(
		agent,
		core.OffensiveTrinketActiveAuraID,
		stats.AttackPower,
		216,
		time.Second*10,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 32654},
			CooldownID:       CrystalforgedTrinketCooldownID,
			Cooldown:         time.Minute * 1,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	)
}

var BadgeOfTheSwarmguardAuraID = core.NewAuraID()
var BadgeOfTheSwarmguardProcAuraID = core.NewAuraID()
var BadgeOfTheSwarmguardCooldownID = core.NewCooldownID()
var BadgeOfTheSwarmguardActionID = core.ActionID{ItemID: 21670}

func ApplyBadgeOfTheSwarmguard(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddMajorCooldown(core.MajorCooldown{
		ActionID:         BadgeOfTheSwarmguardActionID,
		CooldownID:       BadgeOfTheSwarmguardCooldownID,
		Cooldown:         time.Minute * 3,
		SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		Type:             core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				const arPenBonus = 200.0
				const dur = time.Second * 30
				ppmm := character.AutoAttacks.NewPPMManager(10.0)
				stacks := 0

				character.AddAura(sim, core.Aura{
					ID:       BadgeOfTheSwarmguardProcAuraID,
					ActionID: BadgeOfTheSwarmguardActionID,
					Expires:  sim.CurrentTime + dur,
					OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
						if !hitEffect.Landed() {
							return
						}

						if !ppmm.Proc(sim, hitEffect.IsMH(), hitEffect.IsRanged(), "Badge of the Swarmguard") {
							return
						}

						if stacks < 6 {
							character.AddStat(stats.ArmorPenetration, arPenBonus)
							stacks++
						}
					},
					OnExpire: func(sim *core.Simulation) {
						character.AddStat(stats.ArmorPenetration, -arPenBonus*float64(stacks))
						stacks = 0
					},
				})
				character.SetCD(BadgeOfTheSwarmguardCooldownID, sim.CurrentTime+time.Minute*3)
			}
		},
	})
}

var MarkOfTheChampionMeleeAuraID = core.NewAuraID()

func ApplyMarkOfTheChampionMelee(agent core.Agent) {
	agent.GetCharacter().AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: MarkOfTheChampionMeleeAuraID,
			OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if hitEffect.Target.MobType == proto.MobType_MobTypeDemon || hitEffect.Target.MobType == proto.MobType_MobTypeUndead {
					hitEffect.BonusAttackPower += 150
				}
			},
		}
	})
}

var HourglassUnravellerAuraID = core.NewAuraID()
var RageOfUnravellerAuraID = core.NewAuraID()

func ApplyHourglassUnraveller(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const statBonus = 300.0
		const dur = time.Second * 10
		const procChance = 0.1

		icd := core.NewICD()
		const icdDur = time.Second * 50

		return core.Aura{
			ID: HourglassUnravellerAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Outcome.Matches(core.OutcomeCrit) || ability.IsPhantom {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if sim.RandomFloat("Hourglass of the Unraveller") > procChance {
					return
				}

				icd = core.InternalCD(sim.CurrentTime + icdDur)
				character.AddAuraWithTemporaryStats(sim, RageOfUnravellerAuraID, core.ActionID{ItemID: 28034}, stats.AttackPower, statBonus, dur)
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
					ActionID:       core.ActionID{ItemID: 28579},
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
					MinBaseDamage: 222,
					MaxBaseDamage: 332,
				},
			},
		})

		return core.Aura{
			ID: RomulosPoisonVialAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				// mask 340
				if !hitEffect.IsWeaponHit() || !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || ability.IsPhantom {
					return
				}
				if !ppmm.Proc(sim, hitEffect.IsMH(), hitEffect.IsRanged(), "RomulosPoisonVial") {
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
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				// mask: 340
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() || !hitEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || ability.IsPhantom {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if !ppmm.Proc(sim, hitEffect.IsMH(), hitEffect.IsRanged(), "dragonspine") {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)

				character.AddAuraWithTemporaryStats(sim, MeleeHasteAuraID, core.ActionID{ItemID: 28830}, stats.MeleeHaste, hasteBonus, dur)
			},
		}
	})
}

var TsunamiTalismanAuraID = core.NewAuraID()
var TsunamiTalismanProcAuraID = core.NewAuraID()

func ApplyTsunamiTalisman(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const apBonus = 340
		const dur = time.Second * 10
		const procChance = 0.1

		icd := core.NewICD()
		const icdDur = time.Second * 45

		return core.Aura{
			ID: TsunamiTalismanAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Outcome.Matches(core.OutcomeCrit) || ability.IsPhantom {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if sim.RandomFloat("Tsunami Talisman") > procChance {
					return
				}

				icd = core.InternalCD(sim.CurrentTime + icdDur)
				character.AddAuraWithTemporaryStats(sim, TsunamiTalismanProcAuraID, core.ActionID{ItemID: 30627}, stats.AttackPower, apBonus, dur)
			},
		}
	})
}

var DarkmoonCardWrathAuraID = core.NewAuraID()

func ApplyDarkmoonCardWrath(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const critBonus = 17.0
		stacks := 0

		return core.Aura{
			ID: DarkmoonCardWrathAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				// mask 340
				if !hitEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || ability.IsPhantom || !hitEffect.IsWeaponHit() {
					return
				}

				if hitEffect.Outcome.Matches(core.OutcomeCrit) {
					removeAmount := -1 * critBonus * float64(stacks)
					character.AddStat(stats.MeleeCrit, removeAmount)
					character.AddStat(stats.SpellCrit, removeAmount)
					stacks = 0
				} else {
					character.AddStat(stats.MeleeCrit, critBonus)
					character.AddStat(stats.SpellCrit, critBonus)
					stacks++
				}
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
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				// mask 340
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() || !hitEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || ability.IsPhantom {
					return
				}
				if !ppmm.Proc(sim, hitEffect.IsMH(), hitEffect.IsRanged(), "Madness of the Betrayer") {
					return
				}

				character.AddAuraWithTemporaryStats(sim, MadnessOfTheBetrayerProcAuraID, core.ActionID{ItemID: 32505}, stats.ArmorPenetration, arPenBonus, dur)
			},
		}
	})
}

var BlackenedNaaruSliverAuraID = core.NewAuraID()
var BlackenedNaaruSliverProcAuraID = core.NewAuraID()

func ApplyBlackenedNaaruSliver(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const procChance = 0.1

		icd := core.NewICD()
		const icdDur = time.Second * 45

		return core.Aura{
			ID: BlackenedNaaruSliverAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				// 340
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() || !hitEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || ability.IsPhantom {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if sim.RandomFloat("Blackened Naaru Sliver") > procChance {
					return
				}

				icd = core.InternalCD(sim.CurrentTime + icdDur)

				const apBonus = 44.0
				const dur = time.Second * 20
				stacks := 0

				character.AddAura(sim, core.Aura{
					ID:       BlackenedNaaruSliverProcAuraID,
					ActionID: core.ActionID{ItemID: 34427},
					Expires:  sim.CurrentTime + dur,
					OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
						if !hitEffect.Landed() || !hitEffect.IsWeaponHit() || !hitEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || ability.IsPhantom {
							return
						}

						if stacks < 10 {
							character.AddStat(stats.AttackPower, apBonus)
							character.AddStat(stats.RangedAttackPower, apBonus)
							stacks++
						}
					},
					OnExpire: func(sim *core.Simulation) {
						character.AddStat(stats.AttackPower, -apBonus*float64(stacks))
						character.AddStat(stats.RangedAttackPower, -apBonus*float64(stacks))
						stacks = 0
					},
				})
			},
		}
	})
}

var ShardOfContemptAuraID = core.NewAuraID()
var ShardOfContemptProcAuraID = core.NewAuraID()

func ApplyShardOfContempt(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const apBonus = 230
		const dur = time.Second * 20
		const procChance = 0.1

		icd := core.NewICD()
		const icdDur = time.Second * 45

		return core.Aura{
			ID: ShardOfContemptAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.IsWeaponHit() || !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || ability.IsPhantom {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if sim.RandomFloat("Shard of Contempt") > procChance {
					return
				}

				icd = core.InternalCD(sim.CurrentTime + icdDur)
				character.AddAuraWithTemporaryStats(sim, ShardOfContemptProcAuraID, core.ActionID{ItemID: 34472}, stats.AttackPower, apBonus, dur)
			},
		}
	})
}
