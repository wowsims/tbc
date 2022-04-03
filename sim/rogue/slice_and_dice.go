package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var SliceAndDiceActionID = core.ActionID{SpellID: 6774}
var SliceAndDiceAuraID = core.NewAuraID()

const SliceAndDiceEnergyCost = 25.0

func (rogue *Rogue) initSliceAndDice(sim *core.Simulation) {
	durationMultiplier := 1.0 + 0.15*float64(rogue.Talents.ImprovedSliceAndDice)
	durationBonus := time.Duration(0)
	if ItemSetNetherblade.CharacterHasSetBonus(&rogue.Character, 2) {
		durationBonus = time.Second * 3
	}
	rogue.sliceAndDiceDurations = [6]time.Duration{
		0,
		time.Duration(float64(time.Second*9+durationBonus) * durationMultiplier),
		time.Duration(float64(time.Second*12+durationBonus) * durationMultiplier),
		time.Duration(float64(time.Second*15+durationBonus) * durationMultiplier),
		time.Duration(float64(time.Second*18+durationBonus) * durationMultiplier),
		time.Duration(float64(time.Second*21+durationBonus) * durationMultiplier),
	}

	hasteBonus := 1.3
	if ItemSetSlayers.CharacterHasSetBonus(&rogue.Character, 2) {
		hasteBonus += 0.05
	}
	inverseHasteBonus := 1.0 / hasteBonus
	sliceAndDiceAura := core.Aura{
		ID:       SliceAndDiceAuraID,
		ActionID: SliceAndDiceActionID,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			rogue.MultiplyMeleeSpeed(sim, hasteBonus)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			rogue.MultiplyMeleeSpeed(sim, inverseHasteBonus)
		},
	}

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:  SliceAndDiceActionID,
			Character: rogue.GetCharacter(),
			Cost: core.ResourceCost{
				Type:  stats.Energy,
				Value: SliceAndDiceEnergyCost,
			},
			GCD:         time.Second,
			IgnoreHaste: true,
			SpellExtras: SpellFlagFinisher,
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, cast *core.Cast) {
				numPoints := rogue.ComboPoints()
				aura := sliceAndDiceAura
				aura.Duration = rogue.sliceAndDiceDurations[numPoints]
				rogue.ReplaceAura(sim, aura)

				rogue.ApplyFinisher(sim, cast.ActionID)

				if aura.Duration >= sim.GetRemainingDuration() {
					rogue.doneSND = true
				}
			},
		},
	}

	var cast core.SimpleCast

	rogue.castSliceAndDice = func() {
		comboPoints := rogue.ComboPoints()
		if comboPoints == 0 {
			panic("SliceAndDice requires combo points!")
		}

		cast = template
		cast.ActionID.Tag = comboPoints

		if rogue.deathmantle4pcProc {
			cast.Cost.Value = 0
			rogue.deathmantle4pcProc = false
		}

		cast.StartCast(sim)
	}
}
