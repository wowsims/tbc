package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var RuptureActionID = core.ActionID{SpellID: 26867}
var RuptureDebuffID = core.NewDebuffID()
var RuptureEnergyCost = 25.0

func (rogue *Rogue) newRuptureTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	finishingMoveEffects := rogue.makeFinishingMoveEffectApplier(sim)
	refundAmount := 0.4 * float64(rogue.Talents.QuickRecovery)

	ability := rogue.newAbility(RuptureActionID, RuptureEnergyCost, SpellFlagFinisher|core.SpellExtrasBinary, core.ProcMaskMeleeMHSpecial)
	ability.SpellCast.Cast.CritRollCategory = core.CritRollCategoryNone
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			numPoints := rogue.ComboPoints()
			rogue.SpendComboPoints(sim, spellCast.ActionID)
			finishingMoveEffects(sim, numPoints)
		} else {
			if refundAmount > 0 {
				rogue.AddEnergy(sim, spellCast.Cost.Value*refundAmount, core.ActionID{SpellID: 31245})
			}
		}
	}
	ability.Effect.DotInput = core.DotDamageInput{
		NumberOfTicks:  0, // Set dynamically.
		TickLength:     time.Second * 2,
		TickBaseDamage: 0, // Set dynamically.
		DebuffID:       DeadlyPoisonDebuffID,
	}

	ability.Effect.StaticDamageMultiplier += 0.1 * float64(rogue.Talents.SerratedBlades)
	if rogue.Talents.SurpriseAttacks {
		ability.SpellExtras |= core.SpellExtrasCannotBeDodged
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (rogue *Rogue) NewRupture(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	comboPoints := rogue.ComboPoints()
	if comboPoints == 0 {
		panic("Rupture requires combo points!")
	}

	rp := &rogue.rupture
	rogue.ruptureTemplate.Apply(rp)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	rp.ActionID.Tag = comboPoints
	rp.Effect.Target = target
	rp.Effect.DotInput.NumberOfTicks = int(comboPoints) + 3
	rp.Effect.DotInput.TickBaseDamage = 70 + float64(comboPoints)*11

	if rogue.deathmantle4pcProc {
		rp.Cost.Value = 0
		rogue.deathmantle4pcProc = false
	}

	rp.Init(sim)
	return rp
}
