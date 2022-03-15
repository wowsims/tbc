package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	// Proc effects. Keep these in order by item ID.
	core.AddItemEffect(11815, ApplyHandOfJustice)
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
	var KissOfTheSpiderAuraID = core.NewAuraID()
	var KissOfTheSpiderCooldownID = core.NewCooldownID()
	core.AddItemEffect(22954, core.MakeTemporaryStatsOnUseCDRegistration(
		KissOfTheSpiderAuraID,
		stats.Stats{stats.MeleeHaste: 200},
		time.Second*15,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 22954},
			CooldownID:       KissOfTheSpiderCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var SlayersCrestAuraID = core.NewAuraID()
	var SlayersCrestCooldownID = core.NewCooldownID()
	core.AddItemEffect(23041, core.MakeTemporaryStatsOnUseCDRegistration(
		SlayersCrestAuraID,
		stats.Stats{stats.AttackPower: 260, stats.RangedAttackPower: 260},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 23041},
			CooldownID:       SlayersCrestCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var FigurineNightseyePantherAuraID = core.NewAuraID()
	var FigurineNightseyePantherCooldownID = core.NewCooldownID()
	core.AddItemEffect(24128, core.MakeTemporaryStatsOnUseCDRegistration(
		FigurineNightseyePantherAuraID,
		stats.Stats{stats.AttackPower: 320, stats.RangedAttackPower: 320},
		time.Second*12,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 24128},
			CooldownID:       FigurineNightseyePantherCooldownID,
			Cooldown:         time.Minute * 3,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var BladefistsBreadthAuraID = core.NewAuraID()
	var BladefistsBreadthCooldownID = core.NewCooldownID()
	core.AddItemEffect(28041, core.MakeTemporaryStatsOnUseCDRegistration(
		BladefistsBreadthAuraID,
		stats.Stats{stats.AttackPower: 200, stats.RangedAttackPower: 200},
		time.Second*15,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 28041},
			CooldownID:       BladefistsBreadthCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var IconOfUnyieldingCourageAuraID = core.NewAuraID()
	var IconOfUnyieldingCourageCooldownID = core.NewCooldownID()
	core.AddItemEffect(28121, core.MakeTemporaryStatsOnUseCDRegistration(
		IconOfUnyieldingCourageAuraID,
		stats.Stats{stats.ArmorPenetration: 600},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 28121},
			CooldownID:       IconOfUnyieldingCourageCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var AbacusViolentOddsAuraID = core.NewAuraID()
	var AbacusViolentOddsCooldownID = core.NewCooldownID()
	core.AddItemEffect(28288, core.MakeTemporaryStatsOnUseCDRegistration(
		AbacusViolentOddsAuraID,
		stats.Stats{stats.MeleeHaste: 260},
		time.Second*10,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 28288},
			CooldownID:       AbacusViolentOddsCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var BloodlustBroochAuraID = core.NewAuraID()
	var BloodlustBroochCooldownID = core.NewCooldownID()
	core.AddItemEffect(29383, core.MakeTemporaryStatsOnUseCDRegistration(
		BloodlustBroochAuraID,
		stats.Stats{stats.AttackPower: 278, stats.RangedAttackPower: 278},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 29383},
			CooldownID:       BloodlustBroochCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var CoreOfArkelosAuraID = core.NewAuraID()
	var CoreOfArkelosCooldownID = core.NewCooldownID()
	core.AddItemEffect(29776, core.MakeTemporaryStatsOnUseCDRegistration(
		CoreOfArkelosAuraID,
		stats.Stats{stats.AttackPower: 200, stats.RangedAttackPower: 200},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 29776},
			CooldownID:       CoreOfArkelosCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	core.AddItemEffect(32654, ApplyCrystalforgedTrinket)

	var BadgeOfTenacityAuraID = core.NewAuraID()
	var BadgeOfTenacityCooldownID = core.NewCooldownID()
	core.AddItemEffect(32658, core.MakeTemporaryStatsOnUseCDRegistration(
		BadgeOfTenacityAuraID,
		stats.Stats{stats.Agility: 150},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 32658},
			CooldownID:       BadgeOfTenacityCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var BerserkersCallAuraID = core.NewAuraID()
	var BerserkersCallCooldownID = core.NewCooldownID()
	core.AddItemEffect(33831, core.MakeTemporaryStatsOnUseCDRegistration(
		BerserkersCallAuraID,
		stats.Stats{stats.AttackPower: 360, stats.RangedAttackPower: 360},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 33831},
			CooldownID:       BerserkersCallCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var FigurineShadowsongPantherAuraID = core.NewAuraID()
	var FigurineShadowsongPantherCooldownID = core.NewCooldownID()
	core.AddItemEffect(35702, core.MakeTemporaryStatsOnUseCDRegistration(
		FigurineShadowsongPantherAuraID,
		stats.Stats{stats.AttackPower: 320, stats.RangedAttackPower: 320},
		time.Second*15,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 35702},
			CooldownID:       FigurineShadowsongPantherCooldownID,
			Cooldown:         time.Second * 90,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var EmptyDirebrewMugAuraID = core.NewAuraID()
	var EmptyDirebrewMugCooldownID = core.NewCooldownID()
	core.AddItemEffect(38287, core.MakeTemporaryStatsOnUseCDRegistration(
		EmptyDirebrewMugAuraID,
		stats.Stats{stats.AttackPower: 278, stats.RangedAttackPower: 278},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 38287},
			CooldownID:       EmptyDirebrewMugCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))
}

var HandOfJusticeAuraID = core.NewAuraID()

func ApplyHandOfJustice(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		procChance := 0.013333
		var icd core.InternalCD
		icdDur := time.Second * 2

		mhAttack := character.AutoAttacks.MHAuto
		mhAttack.ActionID = core.ActionID{ItemID: 11815}
		cachedAttack := core.SimpleSpell{}

		return core.Aura{
			ID: HandOfJusticeAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				// https://tbc.wowhead.com/spell=15600/hand-of-justice, proc mask = 20.
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellCast.IsPhantom {
					return
				}

				if icd.IsOnCD(sim) {
					return
				}

				if sim.RandomFloat("HandOfJustice") > procChance {
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

var CrystalforgedTrinketCooldownID = core.NewCooldownID()

func ApplyCrystalforgedTrinket(agent core.Agent) {
	agent.GetCharacter().PseudoStats.BonusDamage += 7
	core.RegisterTemporaryStatsOnUseCD(
		agent,
		core.NewAuraID(),
		stats.Stats{stats.AttackPower: 216, stats.RangedAttackPower: 216},
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
					OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						if !spellEffect.Landed() {
							return
						}
						if !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
							return
						}

						if !ppmm.Proc(sim, spellEffect.IsMH(), spellCast.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged), "Badge of the Swarmguard") {
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
			OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
				if spellEffect.Target.MobType == proto.MobType_MobTypeDemon || spellEffect.Target.MobType == proto.MobType_MobTypeUndead {
					spellEffect.BonusAttackPower += 150
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
		applyStatAura := character.NewTemporaryStatsAuraApplier(RageOfUnravellerAuraID, core.ActionID{ItemID: 28034}, stats.Stats{stats.AttackPower: statBonus, stats.RangedAttackPower: statBonus}, dur)

		icd := core.NewICD()
		const icdDur = time.Second * 50

		return core.Aura{
			ID: HourglassUnravellerAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Outcome.Matches(core.OutcomeCrit) || spellCast.IsPhantom {
					return
				}
				if !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if sim.RandomFloat("Hourglass of the Unraveller") > procChance {
					return
				}

				icd = core.InternalCD(sim.CurrentTime + icdDur)
				applyStatAura(sim)
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
					ActionID:            core.ActionID{ItemID: 28579},
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
					MinBaseDamage: 222,
					MaxBaseDamage: 332,
				},
			},
		})

		return core.Aura{
			ID: RomulosPoisonVialAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				// mask 340
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellCast.IsPhantom {
					return
				}
				if !ppmm.Proc(sim, spellEffect.IsMH(), spellCast.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged), "RomulosPoisonVial") {
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

var DragonspineTrophyAuraID = core.NewAuraID()
var MeleeHasteAuraID = core.NewAuraID()

func ApplyDragonspineTrophy(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		icd := core.NewICD()
		const hasteBonus = 325.0
		const dur = time.Second * 10
		const icdDur = time.Second * 20
		applyStatAura := character.NewTemporaryStatsAuraApplier(MeleeHasteAuraID, core.ActionID{ItemID: 28830}, stats.Stats{stats.MeleeHaste: hasteBonus}, dur)

		ppmm := character.AutoAttacks.NewPPMManager(1.0)
		return core.Aura{
			ID: DragonspineTrophyAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				// mask: 340
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellCast.IsPhantom {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if !ppmm.Proc(sim, spellEffect.IsMH(), spellCast.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged), "dragonspine") {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)

				applyStatAura(sim)
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
		applyStatAura := character.NewTemporaryStatsAuraApplier(TsunamiTalismanProcAuraID, core.ActionID{ItemID: 30627}, stats.Stats{stats.AttackPower: apBonus, stats.RangedAttackPower: apBonus}, dur)

		icd := core.NewICD()
		const icdDur = time.Second * 45

		return core.Aura{
			ID: TsunamiTalismanAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Outcome.Matches(core.OutcomeCrit) || spellCast.IsPhantom {
					return
				}
				if !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if sim.RandomFloat("Tsunami Talisman") > procChance {
					return
				}

				icd = core.InternalCD(sim.CurrentTime + icdDur)
				applyStatAura(sim)
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
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				// mask 340
				if !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellCast.IsPhantom {
					return
				}

				if spellEffect.Outcome.Matches(core.OutcomeCrit) {
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
		applyStatAura := character.NewTemporaryStatsAuraApplier(MadnessOfTheBetrayerProcAuraID, core.ActionID{ItemID: 32505}, stats.Stats{stats.ArmorPenetration: arPenBonus}, dur)

		return core.Aura{
			ID: MadnessOfTheBetrayerAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				// mask 340
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellCast.IsPhantom {
					return
				}
				if !ppmm.Proc(sim, spellEffect.IsMH(), spellCast.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged), "Madness of the Betrayer") {
					return
				}

				applyStatAura(sim)
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
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				// mask 340
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellCast.IsPhantom {
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
					OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellCast.IsPhantom {
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
		applyStatAura := character.NewTemporaryStatsAuraApplier(ShardOfContemptProcAuraID, core.ActionID{ItemID: 34472}, stats.Stats{stats.AttackPower: apBonus, stats.RangedAttackPower: apBonus}, dur)

		icd := core.NewICD()
		const icdDur = time.Second * 45

		return core.Aura{
			ID: ShardOfContemptAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellCast.IsPhantom {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if sim.RandomFloat("Shard of Contempt") > procChance {
					return
				}

				icd = core.InternalCD(sim.CurrentTime + icdDur)
				applyStatAura(sim)
			},
		}
	})
}
