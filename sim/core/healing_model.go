package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ChanceOfDeathAuraLabel = "Chance of Death"

func (character *Character) trackChanceOfDeath(healingModel *proto.HealingModel) {
	if healingModel == nil {
		return
	}

	isTanking := false
	for _, target := range character.Env.Encounter.Targets {
		if target.CurrentTarget == &character.Unit {
			isTanking = true
		}
	}
	if !isTanking {
		return
	}

	character.RegisterAura(Aura{
		Label:    ChanceOfDeathAuraLabel,
		Duration: NeverExpires,
		OnReset: func(aura *Aura, sim *Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitTaken: func(aura *Aura, sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
			if spellEffect.Damage > 0 {
				aura.Unit.CurrentHealth -= spellEffect.Damage
				if sim.Log != nil {
					character.Log(sim, "Damage received for %0.02f, current health: %0.01f", spellEffect.Damage, character.CurrentHealth)
				}
				if aura.Unit.CurrentHealth <= 0 && !aura.Unit.Metrics.Died {
					aura.Unit.Metrics.Died = true
					if sim.Log != nil {
						character.Log(sim, "Dead")
					}
				}
			}
		},
		OnPeriodicDamageTaken: func(aura *Aura, sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
			if spellEffect.Damage > 0 {
				aura.Unit.CurrentHealth -= spellEffect.Damage
				if sim.Log != nil {
					character.Log(sim, "Damage received for %0.02f, current health: %0.01f", spellEffect.Damage, character.CurrentHealth)
				}
				if aura.Unit.CurrentHealth <= 0 && !aura.Unit.Metrics.Died {
					aura.Unit.Metrics.Died = true
					if sim.Log != nil {
						character.Log(sim, "Dead")
					}
				}
			}
		},
	})

	if healingModel.Hps != 0 {
		character.applyHealingModel(*healingModel)
	}
}

func (character *Character) applyHealingModel(healingModel proto.HealingModel) {
	cadence := DurationFromSeconds(healingModel.CadenceSeconds)
	if cadence == 0 {
		cadence = time.Millisecond * 2500
	}
	healPerTick := healingModel.Hps * (float64(cadence) / float64(time.Second))

	character.RegisterResetEffect(func(sim *Simulation) {
		StartPeriodicAction(sim, PeriodicActionOptions{
			Period: cadence,
			OnAction: func(sim *Simulation) {
				character.CurrentHealth = MinFloat(character.CurrentHealth+healPerTick, character.GetStat(stats.Health))
				if sim.Log != nil {
					character.Log(sim, "Heal for %0.02f, current health: %0.01f", healPerTick, character.CurrentHealth)
				}
			},
		})
	})
}

func (character *Character) GetPresimOptions(playerConfig proto.Player) *PresimOptions {
	healingModel := playerConfig.HealingModel
	if healingModel == nil || healingModel.Hps != 0 {
		// If Hps is not 0, then we don't need to run the presim.
		return nil
	}

	return &PresimOptions{
		SetPresimPlayerOptions: func(player *proto.Player) {
			player.HealingModel = nil
		},

		OnPresimResult: func(presimResult proto.UnitMetrics, iterations int32, duration time.Duration) bool {
			character.applyHealingModel(proto.HealingModel{
				Hps:            presimResult.Dtps.Avg * 1.25,
				CadenceSeconds: healingModel.CadenceSeconds,
			})
			return true
		},
	}
}
