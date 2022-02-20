package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var SliceAndDiceActionID = core.ActionID{SpellID: 6774}
var SliceAndDiceAuraID = core.NewAuraID()

const SliceAndDiceEnergyCost = 25.0

func (rogue *Rogue) initSliceAndDice(sim *core.Simulation) {
	sliceAndDiceDurations := []time.Duration{
		0,
		time.Second * 9,
		time.Second * 12,
		time.Second * 15,
		time.Second * 18,
		time.Second * 21,
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
	rogue.applySliceAndDiceAura = func(numPoints int32) {
		aura := sliceAndDiceAura
		aura.Expires = sim.CurrentTime + sliceAndDiceDurations[numPoints]
		if rogue.HasAura(SliceAndDiceAuraID) {
			rogue.ReplaceAura(sim, aura)
		} else {
			rogue.MultiplyMeleeSpeed(sim, hasteBonus)
			rogue.AddAura(sim, aura)
		}
	}
}

func (rogue *Rogue) CastSliceAndDice(sim *core.Simulation) {
	if rogue.comboPoints == 0 {
		panic("SliceAndDice requires combo points!")
	}

	actionID := SliceAndDiceActionID
	actionID.Tag = rogue.comboPoints

	rogue.SpendEnergy(sim, SliceAndDiceEnergyCost, actionID)
	rogue.SetGCDTimer(sim, sim.CurrentTime+time.Second*1)
	rogue.Metrics.AddInstantCast(actionID)

	rogue.applySliceAndDiceAura(rogue.comboPoints)
	rogue.SpendComboPoints(sim)
}
