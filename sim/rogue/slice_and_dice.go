package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var SliceAndDiceActionID = core.ActionID{SpellID: 6774}

const SliceAndDiceEnergyCost = 25.0

func (rogue *Rogue) registerSliceAndDice(sim *core.Simulation) {
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

	rogue.SliceAndDiceAura = rogue.RegisterAura(&core.Aura{
		Label:    "Slice and Dice",
		ActionID: SliceAndDiceActionID,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			rogue.MultiplyMeleeSpeed(sim, hasteBonus)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			rogue.MultiplyMeleeSpeed(sim, inverseHasteBonus)
		},
	})

	baseCost := SliceAndDiceEnergyCost
	rogue.SliceAndDice = rogue.RegisterSpell(core.SpellConfig{
		ActionID:    SliceAndDiceActionID,
		SpellExtras: SpellFlagFinisher,

		ResourceType: stats.Energy,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  time.Second,
			},
			ModifyCast:  rogue.applyDeathmantle,
			IgnoreHaste: true,
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, spell *core.Spell) {
			numPoints := rogue.ComboPoints()
			if numPoints == 0 {
				panic("SliceAndDice requires combo points!")
			}
			rogue.SliceAndDiceAura.Duration = rogue.sliceAndDiceDurations[numPoints]
			rogue.SliceAndDiceAura.Activate(sim)

			rogue.ApplyFinisher(sim, spell.ActionID)

			if rogue.SliceAndDiceAura.Duration >= sim.GetRemainingDuration() {
				rogue.doneSND = true
			}
		},
	})
}
