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

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:  SliceAndDiceActionID,
			Character: rogue.GetCharacter(),
			BaseCost: core.ResourceCost{
				Type:  stats.Energy,
				Value: SliceAndDiceEnergyCost,
			},
			Cost: core.ResourceCost{
				Type:  stats.Energy,
				Value: SliceAndDiceEnergyCost,
			},
			GCD:         time.Second,
			IgnoreHaste: true,
			SpellExtras: SpellFlagFinisher,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				numPoints := rogue.comboPoints
				aura := sliceAndDiceAura
				aura.Expires = sim.CurrentTime + sliceAndDiceDurations[numPoints]
				if rogue.HasAura(SliceAndDiceAuraID) {
					rogue.ReplaceAura(sim, aura)
				} else {
					rogue.MultiplyMeleeSpeed(sim, hasteBonus)
					rogue.AddAura(sim, aura)
				}

				rogue.SpendComboPoints(sim, cast.ActionID)
				finishingMoveEffects(sim, numPoints)
			},
		},
	}

	var cast core.SimpleCast

	rogue.castSliceAndDice = func() {
		if rogue.comboPoints == 0 {
			panic("SliceAndDice requires combo points!")
		}

		cast = template
		cast.ActionID.Tag = rogue.comboPoints

		if rogue.deathmantle4pcProc {
			cast.Cost.Value = 0
			rogue.deathmantle4pcProc = false
		}

		cast.StartCast(sim)
	}
}
