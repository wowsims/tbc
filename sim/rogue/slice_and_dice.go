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
	durationBonus := time.Duration(0)
	if ItemSetNetherblade.CharacterHasSetBonus(&rogue.Character, 2) {
		durationBonus = time.Second * 3
	}
	sliceAndDiceDurations := []time.Duration{
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
		OnExpire: func(sim *core.Simulation) {
			rogue.MultiplyMeleeSpeed(sim, inverseHasteBonus)
		},
	}

	finishingMoveEffects := rogue.makeFinishingMoveEffectApplier(sim)

	rogue.castSliceAndDice = func() {
		if rogue.comboPoints == 0 {
			panic("SliceAndDice requires combo points!")
		}

		actionID := SliceAndDiceActionID
		actionID.Tag = rogue.comboPoints

		if rogue.deathmantle4pcProc {
			rogue.deathmantle4pcProc = false
		} else {
			rogue.SpendEnergy(sim, SliceAndDiceEnergyCost, actionID)
		}
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
