package common

import (
	"strconv"
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
	core.AddItemEffect(32654, ApplyCrystalforgedTrinket)
	core.AddItemEffect(34427, ApplyBlackenedNaaruSliver)
	core.AddItemEffect(34472, ApplyShardOfContempt)

	// Battlemasters trinkets
	sharedBattlemasterCooldownID := core.NewCooldownID()
	addBattlemasterEffect := func(itemID int32) {
		core.AddItemEffect(itemID, core.MakeTemporaryStatsOnUseCDRegistration(
			"BattlemasterTrinket-"+strconv.Itoa(int(itemID)),
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

	// Activatable effects. Keep these in order by item ID.
	AddSimpleStatItemActiveEffect(22954, stats.Stats{stats.MeleeHaste: 200}, time.Second*15, time.Minute*2, core.OffensiveTrinketSharedCooldownID)                                                      // Kiss of the Spider
	AddSimpleStatItemActiveEffect(23041, stats.Stats{stats.AttackPower: 260, stats.RangedAttackPower: 260}, time.Second*20, time.Minute*2, core.OffensiveTrinketSharedCooldownID)                       // Slayer's Crest
	AddSimpleStatItemActiveEffect(24128, stats.Stats{stats.AttackPower: 320, stats.RangedAttackPower: 320}, time.Second*12, time.Minute*3, core.OffensiveTrinketSharedCooldownID)                       // Figurine Nightseye Panther
	AddSimpleStatItemActiveEffect(27891, stats.Stats{stats.Armor: 1280}, time.Second*20, time.Minute*2, core.DefensiveTrinketSharedCooldownID)                                                          // Adamantine Figurine
	AddSimpleStatItemActiveEffect(28041, stats.Stats{stats.AttackPower: 200, stats.RangedAttackPower: 200}, time.Second*15, time.Minute*2, core.OffensiveTrinketSharedCooldownID)                       // Bladefists Breadth
	AddSimpleStatItemActiveEffect(28121, stats.Stats{stats.ArmorPenetration: 600}, time.Second*20, time.Minute*2, core.OffensiveTrinketSharedCooldownID)                                                // Icon of Unyielding Courage
	AddSimpleStatItemActiveEffect(28288, stats.Stats{stats.MeleeHaste: 260}, time.Second*10, time.Minute*2, core.OffensiveTrinketSharedCooldownID)                                                      // Abacus of Violent Odds
	AddSimpleStatItemActiveEffect(28528, stats.Stats{stats.Dodge: 300}, time.Second*10, time.Minute*2, core.DefensiveTrinketSharedCooldownID)                                                           // Moroes Lucky Pocket Watch
	AddSimpleStatItemActiveEffect(29383, stats.Stats{stats.AttackPower: 278, stats.RangedAttackPower: 278}, time.Second*20, time.Minute*2, core.OffensiveTrinketSharedCooldownID)                       // Bloodlust Brooch
	AddSimpleStatItemActiveEffect(29387, stats.Stats{stats.BlockValue: 200}, time.Second*20, time.Minute*2, core.DefensiveTrinketSharedCooldownID)                                                      // Gnomeregan Auto-Blocker 600
	AddSimpleStatItemActiveEffect(29776, stats.Stats{stats.AttackPower: 200, stats.RangedAttackPower: 200}, time.Second*20, time.Minute*2, core.OffensiveTrinketSharedCooldownID)                       // Core of Arkelos
	AddSimpleStatItemActiveEffect(30300, stats.Stats{stats.Block: 125}, time.Second*15, time.Second*90, core.DefensiveTrinketSharedCooldownID)                                                          // Dabiris Enigma
	AddSimpleStatItemActiveEffect(30629, stats.Stats{stats.Defense: 165, stats.AttackPower: -330, stats.RangedAttackPower: -330}, time.Second*15, time.Minute*3, core.DefensiveTrinketSharedCooldownID) // Scarab of Displacement
	AddSimpleStatItemActiveEffect(32501, stats.Stats{stats.Health: 1750}, time.Second*20, time.Minute*3, core.DefensiveTrinketSharedCooldownID)                                                         // Shadowmoon Insignia
	AddSimpleStatItemActiveEffect(32534, stats.Stats{stats.Health: 1250}, time.Second*15, time.Minute*5, core.DefensiveTrinketSharedCooldownID)                                                         // Brooch of the Immortal King
	AddSimpleStatItemActiveEffect(32658, stats.Stats{stats.Agility: 150}, time.Second*20, time.Minute*2, core.OffensiveTrinketSharedCooldownID)                                                         // Badge of Tenacity
	AddSimpleStatItemActiveEffect(33830, stats.Stats{stats.Armor: 2500}, time.Second*20, time.Minute*2, core.DefensiveTrinketSharedCooldownID)                                                          // Ancient Aqir Artifact
	AddSimpleStatItemActiveEffect(33831, stats.Stats{stats.AttackPower: 360, stats.RangedAttackPower: 360}, time.Second*20, time.Minute*2, core.OffensiveTrinketSharedCooldownID)                       // Berserkers Call
	AddSimpleStatItemActiveEffect(35702, stats.Stats{stats.AttackPower: 320, stats.RangedAttackPower: 320}, time.Second*15, time.Second*90, core.OffensiveTrinketSharedCooldownID)                      // Figurine Shadowsong Panther
	AddSimpleStatItemActiveEffect(38287, stats.Stats{stats.AttackPower: 278, stats.RangedAttackPower: 278}, time.Second*20, time.Minute*2, core.OffensiveTrinketSharedCooldownID)                       // Empty Direbrew Mug
	AddSimpleStatItemActiveEffect(38289, stats.Stats{stats.BlockValue: 200}, time.Second*20, time.Minute*2, core.DefensiveTrinketSharedCooldownID)                                                      // Coren's Lucky Coin
}

func ApplyHandOfJustice(agent core.Agent) {
	character := agent.GetCharacter()
	if !character.AutoAttacks.IsEnabled() {
		return
	}

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		procChance := 0.013333
		var icd core.InternalCD
		icdDur := time.Second * 2

		handOfJusticeSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:     core.ActionID{ItemID: 11815},
			SpellSchool:  core.SpellSchoolPhysical,
			SpellExtras:  core.SpellExtrasMeleeMetrics,
			ApplyEffects: core.ApplyEffectFuncDirectDamage(character.AutoAttacks.MHEffect),
		})

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Hand of Justice",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
		})
	})
}

var CrystalforgedTrinketCooldownID = core.NewCooldownID()

func ApplyCrystalforgedTrinket(agent core.Agent) {
	agent.GetCharacter().PseudoStats.BonusDamage += 7
	core.RegisterTemporaryStatsOnUseCD(
		agent,
		"Crystalforged Trinket",
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

var BadgeOfTheSwarmguardCooldownID = core.NewCooldownID()
var BadgeOfTheSwarmguardActionID = core.ActionID{ItemID: 21670}

func ApplyBadgeOfTheSwarmguard(agent core.Agent) {
	character := agent.GetCharacter()

	procAura := character.GetOrRegisterAura(&core.Aura{
		Label:     "Badge of the Swarmguard Proc",
		ActionID:  core.ActionID{SpellID: 26481},
		Duration:  core.NeverExpires,
		MaxStacks: 6,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			character.AddStat(stats.ArmorPenetration, 200*float64(newStacks-oldStacks))
		},
	})

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
			ppmm := character.AutoAttacks.NewPPMManager(10.0)
			activeAura := character.GetOrRegisterAura(&core.Aura{
				Label:    "Badge of the Swarmguard",
				ActionID: BadgeOfTheSwarmguardActionID,
				Duration: time.Second * 30,
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					procAura.Deactivate(sim)
				},
				OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
					if !spellEffect.Landed() {
						return
					}
					if !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
						return
					}

					if !ppmm.Proc(sim, spellEffect.IsMH(), spellEffect.ProcMask.Matches(core.ProcMaskRanged), "Badge of the Swarmguard") {
						return
					}

					procAura.Activate(sim)
					procAura.AddStack(sim)
				},
			})

			return func(sim *core.Simulation, character *core.Character) {
				activeAura.Activate(sim)
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

func ApplyHourglassUnraveller(agent core.Agent) {
	character := agent.GetCharacter()

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		procAura := character.NewTemporaryStatsAura("Rage of the Unraveller", core.ActionID{ItemID: 28034}, stats.Stats{stats.AttackPower: 300, stats.RangedAttackPower: 300}, time.Second*10)
		const procChance = 0.1
		icd := core.NewICD()
		const icdDur = time.Second * 50

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Hourglass of the Unraveller",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
				procAura.Activate(sim)
			},
		})
	})
}

func ApplyRomulosPoisonVial(agent core.Agent) {
	character := agent.GetCharacter()

	procSpell := character.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{ItemID: 28579},
		SpellSchool: core.SpellSchoolNature,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			IsPhantom:        true,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			BaseDamage:     core.BaseDamageConfigRoll(222, 332),
			OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(character.DefaultSpellCritMultiplier()),
		}),
	})

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		ppmm := character.AutoAttacks.NewPPMManager(1.0)

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Romulos Poison Vial",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				// mask 340
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellEffect.IsPhantom {
					return
				}
				if !ppmm.Proc(sim, spellEffect.IsMH(), spellEffect.ProcMask.Matches(core.ProcMaskRanged), "RomulosPoisonVial") {
					return
				}

				procSpell.Cast(sim, spellEffect.Target)
			},
		})
	})
}

func ApplyDragonspineTrophy(agent core.Agent) {
	character := agent.GetCharacter()

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		procAura := character.NewTemporaryStatsAura("Dragonspine Trophy Proc", core.ActionID{ItemID: 28830}, stats.Stats{stats.MeleeHaste: 325}, time.Second*10)
		icd := core.NewICD()
		const icdDur = time.Second * 20
		ppmm := character.AutoAttacks.NewPPMManager(1.0)

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Dragonspine Trophy",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				// mask: 340
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellEffect.IsPhantom {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if !ppmm.Proc(sim, spellEffect.IsMH(), spellEffect.ProcMask.Matches(core.ProcMaskRanged), "dragonspine") {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)

				procAura.Activate(sim)
			},
		})
	})
}

func ApplyTsunamiTalisman(agent core.Agent) {
	character := agent.GetCharacter()

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		procAura := character.NewTemporaryStatsAura("Tsunami Talisman Proc", core.ActionID{ItemID: 30627}, stats.Stats{stats.AttackPower: 340, stats.RangedAttackPower: 340}, time.Second*10)

		icd := core.NewICD()
		const icdDur = time.Second * 45
		const procChance = 0.1

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Tsunami Talisman",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
				procAura.Activate(sim)
			},
		})
	})
}

func ApplyDarkmoonCardWrath(agent core.Agent) {
	character := agent.GetCharacter()

	procAura := character.GetOrRegisterAura(&core.Aura{
		Label:     "DMC Wrath Proc",
		ActionID:  core.ActionID{ItemID: 31857},
		Duration:  time.Second * 10,
		MaxStacks: 1000,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			character.AddStat(stats.MeleeCrit, 17*float64(newStacks-oldStacks))
			character.AddStat(stats.SpellCrit, 17*float64(newStacks-oldStacks))
		},
	})

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return character.GetOrRegisterAura(&core.Aura{
			Label: "DMC Wrath",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				// mask 340
				if !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellEffect.IsPhantom {
					return
				}

				if spellEffect.Outcome.Matches(core.OutcomeCrit) {
					procAura.Deactivate(sim)
				} else {
					procAura.Activate(sim)
					procAura.AddStack(sim)
				}
			},
		})
	})
}

func ApplyMadnessOfTheBetrayer(agent core.Agent) {
	character := agent.GetCharacter()

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		procAura := character.NewTemporaryStatsAura("Madness of the Betrayer Proc", core.ActionID{ItemID: 32505}, stats.Stats{stats.ArmorPenetration: 300}, time.Second*10)
		ppmm := character.AutoAttacks.NewPPMManager(1.0)

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Madness of the Betrayer",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				// mask 340
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellEffect.IsPhantom {
					return
				}
				if !ppmm.Proc(sim, spellEffect.IsMH(), spellEffect.ProcMask.Matches(core.ProcMaskRanged), "Madness of the Betrayer") {
					return
				}

				procAura.Activate(sim)
			},
		})
	})
}

func ApplyBlackenedNaaruSliver(agent core.Agent) {
	character := agent.GetCharacter()

	procAura := character.GetOrRegisterAura(&core.Aura{
		Label:     "Blackened Naaru Sliver Proc",
		ActionID:  core.ActionID{ItemID: 34427},
		Duration:  time.Second * 20,
		MaxStacks: 10,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			character.AddStat(stats.AttackPower, 44*float64(newStacks-oldStacks))
			character.AddStat(stats.RangedAttackPower, 44*float64(newStacks-oldStacks))
		},
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) || spellEffect.IsPhantom {
				return
			}
			aura.AddStack(sim)
		},
	})

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		const procChance = 0.1

		icd := core.NewICD()
		const icdDur = time.Second * 45

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Blackened Naaru Sliver",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
				procAura.Activate(sim)
			},
		})
	})
}

func ApplyShardOfContempt(agent core.Agent) {
	character := agent.GetCharacter()

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		procAura := character.NewTemporaryStatsAura("Shard of Contempt Proc", core.ActionID{ItemID: 34472}, stats.Stats{stats.AttackPower: 230, stats.RangedAttackPower: 230}, time.Second*20)
		icd := core.NewICD()
		const icdDur = time.Second * 45
		const procChance = 0.1

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Shard of Contempt",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
				procAura.Activate(sim)
			},
		})
	})
}
