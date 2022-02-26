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

	core.AddItemSet(ItemSetBeastLord)
	core.AddItemSet(ItemSetDemonStalker)
	core.AddItemSet(ItemSetRiftStalker)
	core.AddItemSet(ItemSetGronnstalker)
}

var BeastLord4PcAuraID = core.NewAuraID()
var ItemSetBeastLord = core.ItemSet{
	Name:  "Beast Lord Armor",
	Items: map[int32]struct{}{28228: {}, 27474: {}, 28275: {}, 27874: {}, 27801: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
		},
		4: func(agent core.Agent) {
			// Handled in kill_command.go
		},
	},
}

var ItemSetDemonStalker = core.ItemSet{
	Name:  "Demon Stalker Armor",
	Items: map[int32]struct{}{29081: {}, 29082: {}, 29083: {}, 29084: {}, 29085: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
		},
		4: func(agent core.Agent) {
			// Handled in multi_shot.go
		},
	},
}

var ItemSetRiftStalker = core.ItemSet{
	Name:  "Rift Stalker Armor",
	Items: map[int32]struct{}{30139: {}, 30140: {}, 30141: {}, 30142: {}, 30143: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
		},
		4: func(agent core.Agent) {
			// Handled in steady_shot.go
		},
	},
}

var ItemSetGronnstalker = core.ItemSet{
	Name:  "Gronnstalker's Armor",
	Items: map[int32]struct{}{31001: {}, 31003: {}, 31004: {}, 31005: {}, 31006: {}, 34443: {}, 34549: {}, 34570: {}},
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
			ID:       TalonOfAlarAuraID,
			ActionID: core.ActionID{ItemID: 30448},
			OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.SimpleSpell, hitEffect *core.SpellHitEffect) {
				if !ability.SameAction(SteadyShotActionID) &&
					!ability.SameAction(MultiShotActionID) &&
					!ability.SameAction(ArcaneShotActionID) &&
					!ability.SameAction(AimedShotActionID) {
					return
				}
				hitEffect.DirectInput.FlatDamageBonus += 40
			},
		}

		return core.Aura{
			ID: TalonOfAlarAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.SimpleSpell, hitEffect *core.SpellEffect) {
				if !ability.SameAction(ArcaneShotActionID) {
					return
				}

				aura := procAura
				// Add 1 in case we use arcane shot exactly off CD.
				aura.Expires = sim.CurrentTime + time.Second*6 + 1

				character.AddAuraOnNextAdvance(sim, aura)
			},
		}
	})
}

func ApplyBeasttamersShoulders(agent core.Agent) {
	hunterAgent, ok := agent.(Agent)
	if !ok {
		log.Fatalf("Non-hunter attempted to activate hunter item effect.")
	}
	hunter := hunterAgent.GetHunter()

	hunter.pet.damageMultiplier *= 1.03
	hunter.pet.AddStat(stats.MeleeCrit, core.MeleeCritRatingPerCritChance*2)
}

var BlackBowOfTheBetrayerAuraID = core.NewAuraID()

func ApplyBlackBowOfTheBetrayer(agent core.Agent) {
	character := agent.GetCharacter()
	const manaGain = 8.0
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: BlackBowOfTheBetrayerAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.SimpleSpell, hitEffect *core.SpellEffect) {
				if !hitEffect.Landed() || !ability.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged) {
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
		statApplier := character.NewTempStatAuraApplier(sim, AshtongueTalismanOfSwiftnessProcAuraID, core.ActionID{ItemID: 32487}, stats.AttackPower, 275, time.Second*8)
		const procChance = 0.15

		return core.Aura{
			ID: AshtongueTalismanOfSwiftnessAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.SimpleSpell, hitEffect *core.SpellEffect) {
				if !ability.SameAction(SteadyShotActionID) {
					return
				}
				if sim.RandomFloat("Ashtongue Talisman of Swiftness") > procChance {
					return
				}
				statApplier(sim)
			},
		}
	})
}
