package core

import (
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// TODO: Make this into an object like rageBar or energyBar.
func (character *Character) EnableManaBar() {
	character.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Intellect,
		ModifiedStat: stats.Mana,
		Modifier: func(intellect float64, mana float64) float64 {
			// Assumes all characters have >= 20 intellect.
			// See https://wowwiki-archive.fandom.com/wiki/Base_mana.
			return mana + (20 + 15*(intellect-20))
		},
	})
}

func (character *Character) HasManaBar() bool {
	return character.MaxMana() > 0
}

// Empty handler so Agents don't have to provide one if they have no logic to add.
func (character *Character) OnManaTick(sim *Simulation) {}

func (character *Character) BaseMana() float64 {
	return character.GetBaseStats()[stats.Mana]
}
func (character *Character) MaxMana() float64 {
	return character.GetInitialStat(stats.Mana)
}
func (character *Character) CurrentMana() float64 {
	return character.stats[stats.Mana]
}
func (character *Character) CurrentManaPercent() float64 {
	return character.CurrentMana() / character.MaxMana()
}

func (character *Character) AddMana(sim *Simulation, amount float64, actionID ActionID, isBonusMana bool) {
	if amount < 0 {
		panic("Trying to add negative mana!")
	}

	oldMana := character.CurrentMana()
	newMana := MinFloat(oldMana+amount, character.MaxMana())
	character.Metrics.AddResourceEvent(actionID, proto.ResourceType_ResourceTypeMana, amount, newMana-oldMana)

	if sim.Log != nil {
		character.Log(sim, "Gained %0.3f mana from %s (%0.3f --> %0.3f).", amount, actionID, oldMana, newMana)
	}

	character.stats[stats.Mana] = newMana
	character.Metrics.ManaGained += newMana - oldMana
	if isBonusMana {
		character.Metrics.BonusManaGained += newMana - oldMana
	}
}

func (character *Character) SpendMana(sim *Simulation, amount float64, actionID ActionID) {
	if amount < 0 {
		panic("Trying to spend negative mana!")
	}

	newMana := character.CurrentMana() - amount
	character.Metrics.AddResourceEvent(actionID, proto.ResourceType_ResourceTypeMana, -amount, -amount)

	if sim.Log != nil {
		character.Log(sim, "Spent %0.3f mana from %s (%0.3f --> %0.3f).", amount, actionID, character.CurrentMana(), newMana)
	}

	character.stats[stats.Mana] = newMana
	character.Metrics.ManaSpent += amount
}

// Returns the rate of mana regen per second from mp5.
func (character *Character) MP5ManaRegenPerSecond() float64 {
	return character.stats[stats.MP5] / 5.0
}

// Returns the rate of mana regen per second from spirit.
func (character *Character) SpiritManaRegenPerSecond() float64 {
	return 0.001 + character.stats[stats.Spirit]*math.Sqrt(character.stats[stats.Intellect])*0.009327
}

// Returns the rate of mana regen per second, assuming this character is
// considered to be casting.
func (character *Character) ManaRegenPerSecondWhileCasting() float64 {
	regenRate := character.MP5ManaRegenPerSecond()

	spiritRegenRate := 0.0
	if character.PseudoStats.SpiritRegenRateCasting != 0 || character.PseudoStats.ForceFullSpiritRegen {
		spiritRegenRate = character.SpiritManaRegenPerSecond() * character.PseudoStats.SpiritRegenMultiplier
		if !character.PseudoStats.ForceFullSpiritRegen {
			spiritRegenRate *= character.PseudoStats.SpiritRegenRateCasting
		}
	}
	regenRate += spiritRegenRate

	return regenRate
}

// Returns the rate of mana regen per second, assuming this character is
// considered to be not casting.
func (character *Character) ManaRegenPerSecondWhileNotCasting() float64 {
	regenRate := character.MP5ManaRegenPerSecond()

	regenRate += character.SpiritManaRegenPerSecond() * character.PseudoStats.SpiritRegenMultiplier

	return regenRate
}

func (character *Character) UpdateManaRegenRates() {
	character.manaTickWhileCasting = character.ManaRegenPerSecondWhileCasting() * 2
	character.manaTickWhileNotCasting = character.ManaRegenPerSecondWhileNotCasting() * 2
}

// Applies 1 'tick' of mana regen, which worth 2s of regeneration based on mp5/int/spirit/etc.
func (character *Character) ManaTick(sim *Simulation) {
	if sim.CurrentTime < character.PseudoStats.FiveSecondRuleRefreshTime {
		regen := character.manaTickWhileCasting
		character.AddMana(sim, regen, ActionID{OtherID: proto.OtherAction_OtherActionManaRegen, Tag: 1}, false)
	} else {
		regen := character.manaTickWhileNotCasting
		character.AddMana(sim, regen, ActionID{OtherID: proto.OtherAction_OtherActionManaRegen, Tag: 2}, false)
	}
}

// Returns the amount of time this Character would need to wait in order to reach
// the desired amount of mana, via mana regen.
//
// Assumes that desiredMana > currentMana. Calculation assumes the Character
// will not take any actions during this period that would reset the 5-second rule.
func (character *Character) TimeUntilManaRegen(desiredMana float64) time.Duration {
	// +1 at the end is to deal with floating point math rounding errors.
	manaNeeded := desiredMana - character.CurrentMana()
	regenTime := NeverExpires

	regenWhileCasting := character.ManaRegenPerSecondWhileCasting()
	if regenWhileCasting != 0 {
		regenTime = DurationFromSeconds(manaNeeded/regenWhileCasting) + 1
	}

	// TODO: this needs to have access to the sim to see current time vs character.PseudoStats.FiveSecondRule.
	//  it is possible that we have been waiting.
	//  In practice this function is always used right after a previous cast so no big deal for now.
	if regenTime > time.Second*5 {
		regenTime = time.Second * 5
		manaNeeded -= regenWhileCasting * 5
		// now we move into spirit based regen.
		regenTime += DurationFromSeconds(manaNeeded / character.ManaRegenPerSecondWhileNotCasting())
	}

	return regenTime
}

func (sim *Simulation) initManaTickAction() {
	var playersWithManaBars []Agent
	var petsWithManaBars []PetAgent

	for _, party := range sim.Raid.Parties {
		for _, player := range party.Players {
			character := player.GetCharacter()
			if character.HasManaBar() {
				playersWithManaBars = append(playersWithManaBars, player)
			}

			for _, petAgent := range character.Pets {
				pet := petAgent.GetPet()
				if pet.HasManaBar() {
					petsWithManaBars = append(petsWithManaBars, petAgent)
				}
			}
		}
	}

	if len(playersWithManaBars) == 0 && len(petsWithManaBars) == 0 {
		return
	}

	interval := time.Second * 2
	pa := &PendingAction{
		Name:     "Mana Tick",
		Priority: ActionPriorityRegen,

		NextActionAt: interval,
	}
	pa.OnAction = func(sim *Simulation) {
		for _, player := range playersWithManaBars {
			player.GetCharacter().ManaTick(sim)
			player.OnManaTick(sim)
		}
		for _, petAgent := range petsWithManaBars {
			pet := petAgent.GetPet()
			if pet.IsEnabled() {
				pet.ManaTick(sim)
				petAgent.OnManaTick(sim)
			}
		}

		pa.NextActionAt = sim.CurrentTime + interval
		sim.AddPendingAction(pa)
	}
	sim.AddPendingAction(pa)
}
