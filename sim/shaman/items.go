package shaman

import (
	"log"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddItemEffect(19344, ApplyNaturalAlignmentCrystal)
	core.AddItemEffect(30663, ApplyFathomBroochOfTheTidewalker)
	core.AddItemEffect(32491, ApplyAshtongueTalismanOfVision)
	core.AddItemEffect(33506, ApplySkycallTotem)
	core.AddItemEffect(33507, ApplyStonebreakersTotem)

	core.AddItemSet(&ItemSetTidefury)
	core.AddItemSet(&ItemSetCycloneRegalia)
	core.AddItemSet(&ItemSetCataclysmRegalia)
	core.AddItemSet(&ItemSetSkyshatterRegalia)

	core.AddItemSet(&ItemSetCycloneHarness)
	core.AddItemSet(&ItemSetCataclysmHarness)
	core.AddItemSet(&ItemSetSkyshatterHarness)

	// Even though these item effects are handled elsewhere, add them so they are
	// detected for automatic testing.
	core.AddItemEffect(TotemOfThePulsingEarth, func(core.Agent) {})
}

var ItemSetTidefury = core.ItemSet{
	Name: "Tidefury Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Handled in chain_lightning.go
		},
		4: func(agent core.Agent) {
			shamanAgent, ok := agent.(ShamanAgent)
			if !ok {
				log.Fatalf("Non-shaman attempted to activate shaman cyclone set bonus.")
			}
			shaman := shamanAgent.GetShaman()

			if shaman.SelfBuffs.WaterShield {
				shaman.AddStat(stats.MP5, 3)
			}
		},
	},
}

var ItemSetCycloneRegalia = core.ItemSet{
	Name: "Cyclone Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Handled in shaman.go
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()

			procAura := character.GetOrRegisterAura(&core.Aura{
				Label:    "Cyclone Regalia 4pc Proc",
				Duration: time.Second * 15,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.PseudoStats.CostReduction += 270
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.PseudoStats.CostReduction -= 270
				},
				OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
					aura.Deactivate(sim)
				},
			})

			character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
				return character.GetOrRegisterAura(&core.Aura{
					Label: "Cyclone Regalia 4pc",
					OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
						if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
							return
						}
						if !spellEffect.Outcome.Matches(core.OutcomeCrit) || sim.RandomFloat("cycl4p") > 0.11 {
							return // if not a crit or didn't proc, don't activate
						}
						procAura.Activate(sim)
					},
				})
			})
		},
	},
}

var ItemSetCataclysmRegalia = core.ItemSet{
	Name: "Cataclysm Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
				return character.GetOrRegisterAura(&core.Aura{
					Label: "Cataclysm Regalia 4pc",
					OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
						if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
							return
						}
						if !spellEffect.Outcome.Matches(core.OutcomeCrit) || sim.RandomFloat("cata4p") > 0.25 {
							return
						}
						character.AddMana(sim, 120, core.ActionID{SpellID: 37237}, false)
					},
				})
			})
		},
	},
}

var ItemSetSkyshatterRegalia = core.ItemSet{
	Name: "Skyshatter Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			shamanAgent, ok := agent.(ShamanAgent)
			if !ok {
				log.Fatalf("Non-shaman attempted to activate shaman t6 2p bonus.")
			}
			shaman := shamanAgent.GetShaman()

			if shaman.Totems.Air == proto.AirTotem_NoAirTotem ||
				shaman.Totems.Water == proto.WaterTotem_NoWaterTotem ||
				shaman.Totems.Earth == proto.EarthTotem_NoEarthTotem ||
				shaman.Totems.Fire == proto.FireTotem_NoFireTotem {
				return
			}

			shaman.AddStat(stats.MP5, 15)
			shaman.AddStat(stats.SpellCrit, 35)
			shaman.AddStat(stats.SpellPower, 45)
		},
		4: func(agent core.Agent) {
			// Increases damage done by Lightning Bolt by 5%.
			// Implemented in lightning_bolt.go.
		},
	},
}

var NaturalAlignmentCrystalCooldownID = core.NewCooldownID()

func ApplyNaturalAlignmentCrystal(agent core.Agent) {
	const sp = 250
	const dur = time.Second * 20
	const cd = time.Minute * 5
	actionID := core.ActionID{ItemID: 19344}

	character := agent.GetCharacter()
	activeAura := character.GetOrRegisterAura(&core.Aura{
		Label:    "Natural Alignment Crystal",
		ActionID: actionID,
		Duration: dur,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			character.AddStat(stats.SpellPower, sp)
			character.PseudoStats.CostMultiplier *= 1.2
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			character.AddStat(stats.SpellPower, -sp)
			character.PseudoStats.CostMultiplier /= 1.2
		},
	})

	character.AddMajorCooldown(core.MajorCooldown{
		ActionID:         actionID,
		CooldownID:       NaturalAlignmentCrystalCooldownID,
		Cooldown:         cd,
		SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		SharedCooldown:   dur,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				activeAura.Activate(sim)
				character.SetCD(NaturalAlignmentCrystalCooldownID, sim.CurrentTime+cd)
				character.Metrics.AddInstantCast(actionID)
			}
		},
	})
}

// ActivateFathomBrooch adds an aura that has a chance on cast of nature spell
//  to restore 335 mana. 40s ICD
func ApplyFathomBroochOfTheTidewalker(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		icd := core.NewICD()
		const icdDur = time.Second * 40

		return character.GetOrRegisterAura(&core.Aura{
			Label: "Fathom Brooch of the Tidewalker",
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
				if icd.IsOnCD(sim) {
					return
				}
				if spell.SpellSchool != core.SpellSchoolNature {
					return
				}
				if sim.RandomFloat("Fathom-Brooch of the Tidewalker") > 0.15 {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				character.AddMana(sim, 335, core.ActionID{ItemID: 30663}, false)
			},
		})
	})
}

func ApplyAshtongueTalismanOfVision(agent core.Agent) {
	shamanAgent, ok := agent.(ShamanAgent)
	if !ok {
		log.Fatalf("Non-shaman attempted to activate Ashtongue Talisman of Vision.")
	}
	shaman := shamanAgent.GetShaman()

	shaman.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		procAura := shaman.NewTemporaryStatsAura("Ashtongue Talisman of Vision Proc", core.ActionID{ItemID: 32491}, stats.Stats{stats.AttackPower: 275}, time.Second*10)
		return shaman.GetOrRegisterAura(&core.Aura{
			Label: "Ashtongue Talisman of Vision",
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
				// Note that shaman.Stormstrike is the first 'fake' SS hit.
				if spell != shaman.Stormstrike {
					return
				}
				if sim.RandomFloat("Ashtongue Talisman of Vision") > 0.5 {
					return
				}
				procAura.Activate(sim)
			},
		})
	})
}

func ApplySkycallTotem(agent core.Agent) {
	character := agent.GetCharacter()

	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		procAura := character.NewTemporaryStatsAura("Skycall Totem Proc", core.ActionID{ItemID: 33506}, stats.Stats{stats.SpellHaste: 101}, time.Second*10)
		return character.GetOrRegisterAura(&core.Aura{
			Label: "Skycall Totem",
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
				if spell.ActionID.SpellID == SpellIDLB12 && sim.RandomFloat("Skycall Totem") < 0.15 {
					procAura.Activate(sim)
				}
			},
		})
	})
}

func ApplyStonebreakersTotem(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		procAura := character.NewTemporaryStatsAura("Stonebreakers Totem Proc", core.ActionID{ItemID: 33507}, stats.Stats{stats.AttackPower: 110}, time.Second*10)
		icd := core.NewICD()
		const icdDur = time.Second * 10
		const procChance = 0.5

		return character.GetOrRegisterAura(&core.Aura{
			Label:    "Stonebreakers Totem",
			Duration: core.NeverExpires,
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}

				if !spell.SpellExtras.Matches(SpellFlagShock) {
					return
				}

				if icd.IsOnCD(sim) {
					return
				}

				if sim.RandomFloat("Stonebreakers Totem") > procChance {
					return
				}

				icd = core.InternalCD(sim.CurrentTime + icdDur)
				procAura.Activate(sim)
			},
		})
	})
}

// Cyclone Harness
// (2) Set : Your Strength of Earth Totem ability grants an additional 12 strength.
// (4) Set : Your Stormstrike ability does an additional 30 damage per weapon.

var ItemSetCycloneHarness = core.ItemSet{
	Name: "Cyclone Harness",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// shaman.go
		},
		4: func(agent core.Agent) {
			// stormstrike.go
		},
	},
}

// Cataclysm Harness
// (2) Set : Your melee attacks have a chance to reduce the cast time of your next Lesser Healing Wave by 1.5 sec. (Proc chance: 2%)
// (4) Set : You gain 5% additional haste from your Flurry ability.

var ItemSetCataclysmHarness = core.ItemSet{
	Name: "Cataclysm Harness",
	Bonuses: map[int32]core.ApplyEffect{
		4: func(agent core.Agent) {
			// shaman.go
		},
	},
}

// Skyshatter Harness
// 2 pieces: Your Earth Shock, Flame Shock, and Frost Shock abilities cost 10% less mana.
// 4 pieces: Whenever you use Stormstrike, you gain 70 attack power for 12 sec.

var ItemSetSkyshatterHarness = core.ItemSet{
	Name: "Skyshatter Harness",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// implemented in shocks.go
		},
		4: func(agent core.Agent) {
			// implemented in stormstrike.go
		},
	},
}
