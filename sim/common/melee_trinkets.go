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

	// Adamantine Figurine
	core.AddItemEffect(27891, core.MakeTemporaryStatsOnUseCDRegistration(
		core.NewAuraID(),
		stats.Stats{stats.Armor: 1280},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 27891},
			CooldownID:       core.NewCooldownID(),
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.DefensiveTrinketSharedCooldownID,
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

	// Moroes Lucky Pocket Watch
	core.AddItemEffect(28528, core.MakeTemporaryStatsOnUseCDRegistration(
		core.NewAuraID(),
		stats.Stats{stats.Dodge: 300},
		time.Second*10,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 28528},
			CooldownID:       core.NewCooldownID(),
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.DefensiveTrinketSharedCooldownID,
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

	// Gnomeregan Auto-Blocker 600
	core.AddItemEffect(29387, core.MakeTemporaryStatsOnUseCDRegistration(
		core.NewAuraID(),
		stats.Stats{stats.BlockValue: 200},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 29387},
			CooldownID:       core.NewCooldownID(),
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.DefensiveTrinketSharedCooldownID,
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

	// Dabiris Enigma
	core.AddItemEffect(30300, core.MakeTemporaryStatsOnUseCDRegistration(
		core.NewAuraID(),
		stats.Stats{stats.Block: 125},
		time.Second*15,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 30300},
			CooldownID:       core.NewCooldownID(),
			Cooldown:         time.Second * 90,
			SharedCooldownID: core.DefensiveTrinketSharedCooldownID,
		},
	))

	// Scarab of Displacement
	core.AddItemEffect(30629, core.MakeTemporaryStatsOnUseCDRegistration(
		core.NewAuraID(),
		stats.Stats{stats.Defense: 165, stats.AttackPower: -330, stats.RangedAttackPower: -330},
		time.Second*15,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 30629},
			CooldownID:       core.NewCooldownID(),
			Cooldown:         time.Minute * 3,
			SharedCooldownID: core.DefensiveTrinketSharedCooldownID,
		},
	))

	// Shadowmoon Insignia
	core.AddItemEffect(32501, core.MakeTemporaryStatsOnUseCDRegistration(
		core.NewAuraID(),
		stats.Stats{stats.Health: 1750},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 32501},
			CooldownID:       core.NewCooldownID(),
			Cooldown:         time.Minute * 3,
			SharedCooldownID: core.DefensiveTrinketSharedCooldownID,
		},
	))

	// Brooch of the Immortal King
	core.AddItemEffect(32534, core.MakeTemporaryStatsOnUseCDRegistration(
		core.NewAuraID(),
		stats.Stats{stats.Health: 1250},
		time.Second*15,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 32534},
			CooldownID:       core.NewCooldownID(),
			Cooldown:         time.Minute * 5,
			SharedCooldownID: core.DefensiveTrinketSharedCooldownID,
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

	// Ancient Aqir Artifact
	core.AddItemEffect(33830, core.MakeTemporaryStatsOnUseCDRegistration(
		core.NewAuraID(),
		stats.Stats{stats.Armor: 2500},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 33830},
			CooldownID:       core.NewCooldownID(),
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.DefensiveTrinketSharedCooldownID,
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

	// Battlemasters trinkets
	sharedBattlemasterCooldownID := core.NewCooldownID()
	addBattlemasterEffect := func(itemID int32) {
		core.AddItemEffect(itemID, core.MakeTemporaryStatsOnUseCDRegistration(
			core.NewAuraID(),
			stats.Stats{stats.Health: 1750},
			time.Second*15,
			core.MajorCooldown{
				ActionID:         core.ActionID{ItemID: itemID},
				CooldownID:       sharedBattlemasterCooldownID,
				Cooldown:         time.Minute * 3,
				SharedCooldownID: core.DefensiveTrinketSharedCooldownID,
			},
		))
	}
	addBattlemasterEffect(33832)
	addBattlemasterEffect(34049)
	addBattlemasterEffect(34050)
	addBattlemasterEffect(34162)
	addBattlemasterEffect(34163)

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

	// Coren's Lucky Coin
	core.AddItemEffect(38289, core.MakeTemporaryStatsOnUseCDRegistration(
		core.NewAuraID(),
		stats.Stats{stats.BlockValue: 200},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 38289},
			CooldownID:       core.NewCooldownID(),
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.DefensiveTrinketSharedCooldownID,
		},
	))
}

var HandOfJusticeAuraID = core.NewAuraID()

func ApplyHandOfJustice(agent core.Agent) {
	character := agent.GetCharacter()
	if !character.AutoAttacks.IsEnabled() {
		return
	}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		procChance := 0.013333
		var icd core.InternalCD
		icdDur := time.Second * 2

		template := character.AutoAttacks.MHAuto.Template
		template.ActionID = core.ActionID{ItemID: 11815}
		handOfJusticeSpell := character.GetOrRegisterSpell(core.SpellConfig{
			Template:   template,
			ModifyCast: core.ModifyCastAssignTarget,
		})

		return core.Aura{
			ID: HandOfJusticeAuraID,
			OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
				// https://tbc.wowhead.com/spell=15600/hand-of-justice, proc mask = 20.
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellEffect.IsPhantom {
					return
				}

				if icd.IsOnCD(sim) {
					return
				}

				if sim.RandomFloat("HandOfJustice") > procChance {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)

				handOfJusticeSpell.Cast(sim, spellEffect.Target)
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
				ppmm := character.AutoAttacks.NewPPMManager(10.0)
				stacks := 0

				character.AddAura(sim, core.Aura{
					ID:       BadgeOfTheSwarmguardProcAuraID,
					ActionID: BadgeOfTheSwarmguardActionID,
					Duration: time.Second * 30,
					OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
						if !spellEffect.Landed() {
							return
						}
						if !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
							return
						}

						if !ppmm.Proc(sim, spellEffect.IsMH(), spellEffect.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged), "Badge of the Swarmguard") {
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

func ApplyMarkOfTheChampionMelee(agent core.Agent) {
	character := agent.GetCharacter()
	character.RegisterResetEffect(func(sim *core.Simulation) {
		if sim.GetPrimaryTarget().MobType == proto.MobType_MobTypeDemon || sim.GetPrimaryTarget().MobType == proto.MobType_MobTypeUndead {
			character.PseudoStats.MobTypeAttackPower += 150
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
			OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
				if !spellEffect.Outcome.Matches(core.OutcomeCrit) || spellEffect.IsPhantom {
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

	procSpell := character.RegisterSpell(core.SpellConfig{
		Template: core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:    core.ActionID{ItemID: 28579},
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
				BaseDamage:          core.BaseDamageConfigRoll(222, 332),
			},
		},
		ModifyCast: core.ModifyCastAssignTarget,
	})

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		ppmm := character.AutoAttacks.NewPPMManager(1.0)

		return core.Aura{
			ID: RomulosPoisonVialAuraID,
			OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
				// mask 340
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellEffect.IsPhantom {
					return
				}
				if !ppmm.Proc(sim, spellEffect.IsMH(), spellEffect.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged), "RomulosPoisonVial") {
					return
				}

				procSpell.Cast(sim, spellEffect.Target)
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
			OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
				// mask: 340
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellEffect.IsPhantom {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if !ppmm.Proc(sim, spellEffect.IsMH(), spellEffect.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged), "dragonspine") {
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
			OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
				if !spellEffect.Outcome.Matches(core.OutcomeCrit) || spellEffect.IsPhantom {
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
			OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
				// mask 340
				if !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellEffect.IsPhantom {
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
			OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
				// mask 340
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellEffect.IsPhantom {
					return
				}
				if !ppmm.Proc(sim, spellEffect.IsMH(), spellEffect.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged), "Madness of the Betrayer") {
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
			OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
				// mask 340
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellEffect.IsPhantom {
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
				stacks := 0

				character.AddAura(sim, core.Aura{
					ID:       BlackenedNaaruSliverProcAuraID,
					ActionID: core.ActionID{ItemID: 34427},
					Duration: time.Second * 20,
					OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
						if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellEffect.IsPhantom {
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
			OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellEffect.IsPhantom {
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
