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

func (character *Character) BaseMana() float64 {
	return character.GetBaseStats()[stats.Mana]
}
func (character *Character) MaxMana() float64 {
	return character.GetInitialStat(stats.Mana)
}
func (character *Character) CurrentMana() float64 {
	return character.GetStat(stats.Mana)
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

// Regenerates mana based on MP5 stat, spirit regen allowed while casting and the elapsed time.
func (character *Character) RegenManaCasting(sim *Simulation, elapsedTime time.Duration) {
	elapsedSeconds := elapsedTime.Seconds()
	manaRegen := character.ManaRegenPerSecondWhileCasting() * elapsedSeconds
	character.AddMana(sim, manaRegen, ActionID{OtherID: proto.OtherAction_OtherActionManaRegen, Tag: int32(elapsedSeconds * 1000)}, false)
}

// Regenerates mana using mp5 and spirit. Will calculate time since last cast and then enable spirit regen if needed.
func (character *Character) RegenMana(sim *Simulation, elapsedTime time.Duration) {
	elapsedSeconds := elapsedTime.Seconds()
	var regen float64
	if sim.CurrentTime-elapsedTime > character.PseudoStats.FiveSecondRuleRefreshTime {
		// Five second rule activated before the advance window started, so use full
		// spirit regen for the full duration.
		regen = character.ManaRegenPerSecondWhileNotCasting() * elapsedSeconds
	} else if sim.CurrentTime > character.PseudoStats.FiveSecondRuleRefreshTime {
		// Five second rule activated sometime in the middle of the advance window,
		// so regen is a combination of casting and not-casting regen.
		notCastingRegenTime := sim.CurrentTime - character.PseudoStats.FiveSecondRuleRefreshTime // how many seconds of full spirit regen
		castingRegenTime := elapsedTime - notCastingRegenTime
		regen = (character.ManaRegenPerSecondWhileNotCasting() * notCastingRegenTime.Seconds()) + (character.ManaRegenPerSecondWhileCasting() * castingRegenTime.Seconds())
	} else {
		regen = character.ManaRegenPerSecondWhileCasting() * elapsedSeconds
	}
	character.AddMana(sim, regen, ActionID{OtherID: proto.OtherAction_OtherActionManaRegen, Tag: int32(elapsedSeconds * 1000)}, false)
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
