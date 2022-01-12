package shaman

import (
	"log"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddItemEffect(19344, ApplyNaturalAlignmentCrystal)
	core.AddItemEffect(30663, ApplyFathomBroochOfTheTidewalker)
	core.AddItemEffect(32491, ApplyAshtongueTalismanOfVision)
	core.AddItemEffect(33506, ApplySkycallTotem)

	core.AddItemSet(ItemSetTidefury)
	core.AddItemSet(ItemSetCycloneRegalia)
	core.AddItemSet(ItemSetCataclysmRegalia)
	core.AddItemSet(ItemSetSkyshatterRegalia)

	core.AddItemSet(ItemSetCycloneHarness)
	core.AddItemSet(ItemSetCataclysmHarness)
	core.AddItemSet(ItemSetSkyshatterHarness)

	// Even though these item effects are handled elsewhere, add them so they are
	// detected for automatic testing.
	core.AddItemEffect(TotemOfThePulsingEarth, func(core.Agent) {})
}

var Tidefury2PcAuraID = core.NewAuraID()
var ItemSetTidefury = core.ItemSet{
	Name:  "Tidefury Raiment",
	Items: map[int32]struct{}{28231: {}, 27510: {}, 28349: {}, 27909: {}, 27802: {}},
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

var Cyclone4PcAuraID = core.NewAuraID()
var Cyclone4PcManaRegainAuraID = core.NewAuraID()
var ItemSetCycloneRegalia = core.ItemSet{
	Name:  "Cyclone Regalia",
	Items: map[int32]struct{}{29033: {}, 29035: {}, 29034: {}, 29036: {}, 29037: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Handled in shaman.go
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				return core.Aura{
					ID: Cyclone4PcAuraID,
					OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						if !spellEffect.Crit || sim.RandomFloat("cycl4p") > 0.11 {
							return // if not a crit or didn't proc, don't activate
						}
						character.AddAura(sim, core.Aura{
							ID: Cyclone4PcManaRegainAuraID,
							OnCast: func(sim *core.Simulation, cast *core.Cast) {
								// TODO: how to make sure this goes in before clearcasting?
								cast.ManaCost -= 270
							},
							OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
								character.RemoveAura(sim, Cyclone4PcManaRegainAuraID)
							},
						})
					},
				}
			})
		},
	},
}

var Cataclysm4PcAuraID = core.NewAuraID()
var ItemSetCataclysmRegalia = core.ItemSet{
	Name:  "Cataclysm Regalia",
	Items: map[int32]struct{}{30169: {}, 30170: {}, 30171: {}, 30172: {}, 30173: {}},
	Bonuses: map[int32]core.ApplyEffect{
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				return core.Aura{
					ID: Cataclysm4PcAuraID,
					OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						if !spellEffect.Crit || sim.RandomFloat("cata4p") > 0.25 {
							return
						}
						character.AddMana(sim, 120, core.ActionID{SpellID: 37237}, false)
					},
				}
			})
		},
	},
}

var ItemSetSkyshatterRegalia = core.ItemSet{
	Name:  "Skyshatter Regalia",
	Items: map[int32]struct{}{34437: {}, 31017: {}, 34542: {}, 31008: {}, 31014: {}, 31020: {}, 31023: {}, 34566: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MP5, 15)
			agent.GetCharacter().AddStat(stats.SpellCrit, 35)
			agent.GetCharacter().AddStat(stats.SpellPower, 45)
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

	agent.GetCharacter().AddMajorCooldown(core.MajorCooldown{
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
				character.SetCD(NaturalAlignmentCrystalCooldownID, sim.CurrentTime+cd)
				character.AddStat(stats.SpellPower, sp)
				character.Metrics.AddInstantCast(actionID)

				character.AddAura(sim, core.Aura{
					ID:       core.OffensiveTrinketActiveAuraID,
					ActionID: actionID,
					Expires:  sim.CurrentTime + dur,
					OnCast: func(sim *core.Simulation, cast *core.Cast) {
						cast.ManaCost *= 1.2
					},
					OnExpire: func(sim *core.Simulation) {
						character.AddStat(stats.SpellPower, -sp)
					},
				})
			}
		},
	})
}

// ActivateFathomBrooch adds an aura that has a chance on cast of nature spell
//  to restore 335 mana. 40s ICD
var FathomBroochOfTheTidewalkerAuraID = core.NewAuraID()

func ApplyFathomBroochOfTheTidewalker(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		icd := core.NewICD()
		const icdDur = time.Second * 40

		return core.Aura{
			ID: FathomBroochOfTheTidewalkerAuraID,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				if icd.IsOnCD(sim) {
					return
				}
				if cast.SpellSchool != stats.NatureSpellPower {
					return
				}
				if sim.RandomFloat("Fathom-Brooch of the Tidewalker") > 0.15 {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				character.AddMana(sim, 335, core.ActionID{ItemID: 30663}, false)
			},
		}
	})
}

var AshtongueTalismanOfVisionAuraID = core.NewAuraID()
var AshtongueTalismanOfVisionProcAuraID = core.NewAuraID()

func ApplyAshtongueTalismanOfVision(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const apBonus = 275
		const dur = time.Second * 10
		const procChance = 0.5

		return core.Aura{
			ID: AshtongueTalismanOfVisionAuraID,
			OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
				if !ability.ActionID.SameAction(StormstrikeActionID) {
					return
				}
				if sim.RandomFloat("Ashtongue Talisman of Vision") > procChance {
					return
				}
				character.AddAuraWithTemporaryStats(sim, AshtongueTalismanOfVisionProcAuraID, core.ActionID{ItemID: 32491}, stats.AttackPower, apBonus, dur)
			},
		}
	})
}

var SkycallTotemAuraID = core.NewAuraID()
var EnergizedAuraID = core.NewAuraID()

func ApplySkycallTotem(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 101
		const dur = time.Second * 10

		return core.Aura{
			ID:      SkycallTotemAuraID,
			Expires: core.NeverExpires,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				if cast.ActionID.SpellID != SpellIDLB12 || sim.RandomFloat("Skycall Totem") > 0.15 {
					return
				}
				character.AddAuraWithTemporaryStats(sim, EnergizedAuraID, core.ActionID{ItemID: 33506}, stats.SpellHaste, hasteBonus, dur)
			},
		}
	})
}

// Cyclone Harness
// (2) Set : Your Strength of Earth Totem ability grants an additional 12 strength.
// (4) Set : Your Stormstrike ability does an additional 30 damage per weapon.

var ItemSetCycloneHarness = core.ItemSet{
	Name:  "Cyclone Harness",
	Items: map[int32]struct{}{29038: {}, 29039: {}, 29040: {}, 29043: {}, 29042: {}},
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
	Name:  "Cataclysm Harness",
	Items: map[int32]struct{}{30185: {}, 30189: {}, 30190: {}, 30192: {}, 30194: {}},
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
	Name:  "Skyshatter Harness",
	Items: map[int32]struct{}{31018: {}, 31011: {}, 31015: {}, 31021: {}, 31024: {}, 34567: {}, 34439: {}, 34545: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// implemented in shocks.go
		},
		4: func(agent core.Agent) {
			// implemented in stormstrike.go
		},
	},
}
