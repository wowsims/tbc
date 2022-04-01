package hunter

import (
	"log"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddItemEffect(30448, ApplyTalonOfAlar)
	core.AddItemEffect(30892, ApplyBeasttamersShoulders)
	core.AddItemEffect(32336, ApplyBlackBowOfTheBetrayer)
	core.AddItemEffect(32487, ApplyAshtongueTalismanOfSwiftness)

	core.AddItemSet(&ItemSetBeastLord)
	core.AddItemSet(&ItemSetDemonStalker)
	core.AddItemSet(&ItemSetRiftStalker)
	core.AddItemSet(&ItemSetGronnstalker)
}

var BeastLord4PcAuraID = core.NewAuraID()
var ItemSetBeastLord = core.ItemSet{
	Name: "Beast Lord Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
		},
		4: func(agent core.Agent) {
			// Handled in kill_command.go
		},
	},
}

var ItemSetDemonStalker = core.ItemSet{
	Name: "Demon Stalker Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
		},
		4: func(agent core.Agent) {
			// Handled in multi_shot.go
		},
	},
}

var ItemSetRiftStalker = core.ItemSet{
	Name: "Rift Stalker Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
		},
		4: func(agent core.Agent) {
			// Handled in steady_shot.go
		},
	},
}

var ItemSetGronnstalker = core.ItemSet{
	Name: "Gronnstalker's Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Handled in rotation.go
		},
		4: func(agent core.Agent) {
			// Handled in steady_shot.go
		},
	},
}

var TalonOfAlarAuraID = core.NewAuraID()
var TalonOfAlarProcAuraID = core.NewAuraID()

func ApplyTalonOfAlar(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		procAura := core.Aura{
			ID:       TalonOfAlarProcAuraID,
			ActionID: core.ActionID{ItemID: 30448},
			// Add 1 in case we use arcane shot exactly off CD.
			Duration: time.Second*6 + 1,
		}

		return core.Aura{
			ID: TalonOfAlarAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellCast.SameAction(ArcaneShotActionID) {
					return
				}

				character.AddAura(sim, procAura)
			},
		}
	})
}

func (hunter *Hunter) talonOfAlarDamageMod(baseDamageConfig core.BaseDamageConfig) core.BaseDamageConfig {
	if hunter.HasTrinketEquipped(30448) {
		return core.WrapBaseDamageConfig(baseDamageConfig, func(oldCalculator core.BaseDamageCalculator) core.BaseDamageCalculator {
			return func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.SimpleSpellTemplate) float64 {
				normalDamage := oldCalculator(sim, hitEffect, spell)
				if hunter.HasAura(TalonOfAlarProcAuraID) {
					return normalDamage + 40
				} else {
					return normalDamage
				}
			}
		})
	} else {
		return baseDamageConfig
	}
}

func ApplyBeasttamersShoulders(agent core.Agent) {
	hunterAgent, ok := agent.(Agent)
	if !ok {
		log.Fatalf("Non-hunter attempted to activate hunter item effect.")
	}
	hunter := hunterAgent.GetHunter()

	hunter.pet.PseudoStats.DamageDealtMultiplier *= 1.03
	hunter.pet.AddStat(stats.MeleeCrit, core.MeleeCritRatingPerCritChance*2)
}

var BlackBowOfTheBetrayerAuraID = core.NewAuraID()

func ApplyBlackBowOfTheBetrayer(agent core.Agent) {
	character := agent.GetCharacter()
	const manaGain = 8.0
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: BlackBowOfTheBetrayerAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged) {
					return
				}
				character.AddMana(sim, manaGain, core.ActionID{SpellID: 46939}, false)
			},
		}
	})
}

var AshtongueTalismanOfSwiftnessAuraID = core.NewAuraID()
var AshtongueTalismanOfSwiftnessProcAuraID = core.NewAuraID()

func ApplyAshtongueTalismanOfSwiftness(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		applyStatAura := character.NewTemporaryStatsAuraApplier(AshtongueTalismanOfSwiftnessProcAuraID, core.ActionID{ItemID: 32487}, stats.Stats{stats.AttackPower: 275, stats.RangedAttackPower: 275}, time.Second*8)
		const procChance = 0.15

		return core.Aura{
			ID: AshtongueTalismanOfSwiftnessAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellCast.SameAction(SteadyShotActionID) {
					return
				}
				if sim.RandomFloat("Ashtongue Talisman of Swiftness") > procChance {
					return
				}
				applyStatAura(sim)
			},
		}
	})
}
