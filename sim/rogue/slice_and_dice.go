package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var SliceAndDiceActionID = core.ActionID{SpellID: 6774}
var SliceAndDiceAuraID = core.NewAuraID()

const SliceAndDiceEnergyCost = 25.0

func (rogue *Rogue) initSliceAndDice(sim *core.Simulation) {
	durationMultiplier := 1.0 + 0.15*float64(rogue.Talents.ImprovedSliceAndDice)
	sliceAndDiceDurations := []time.Duration{
		0,
		time.Duration(float64(time.Second*9) * durationMultiplier),
		time.Duration(float64(time.Second*12) * durationMultiplier),
		time.Duration(float64(time.Second*15) * durationMultiplier),
		time.Duration(float64(time.Second*18) * durationMultiplier),
		time.Duration(float64(time.Second*21) * durationMultiplier),
	}

	hasteBonus := 1.3
	inverseHasteBonus := 1.0 / 1.3
	sliceAndDiceAura := core.Aura{
		ID:       SliceAndDiceAuraID,
		ActionID: SliceAndDiceActionID,
		OnExpire: func(sim *core.Simulation) {
			rogue.MultiplyMeleeSpeed(sim, inverseHasteBonus)
		},
	}

	finishingMoveEffects := rogue.makeFinishingMoveEffectApplier()

	rogue.castSliceAndDice = func() {
		if rogue.comboPoints == 0 {
			panic("SliceAndDice requires combo points!")
		}

		actionID := SliceAndDiceActionID
		actionID.Tag = rogue.comboPoints

		rogue.SpendEnergy(sim, SliceAndDiceEnergyCost, actionID)
		rogue.SetGCDTimer(sim, sim.CurrentTime+time.Second*1)
		rogue.Metrics.AddInstantCast(actionID)

		numPoints := rogue.comboPoints
		aura := sliceAndDiceAura
		aura.Expires = sim.CurrentTime + sliceAndDiceDurations[numPoints]
		if rogue.HasAura(SliceAndDiceAuraID) {
			rogue.ReplaceAura(sim, aura)
		} else {
			rogue.MultiplyMeleeSpeed(sim, hasteBonus)
			rogue.AddAura(sim, aura)
		}

		rogue.SpendComboPoints(sim)
		finishingMoveEffects(sim, numPoints)
	}
}
