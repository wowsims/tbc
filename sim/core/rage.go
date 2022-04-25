package core

import (
	"github.com/wowsims/tbc/sim/core/proto"
)

const MaxRage = 100.0

const RageFactor = 3.75 / 274.7

// OnRageGain is called any time rage is increased.
type OnRageGain func(sim *Simulation)

type rageBar struct {
	unit *Unit

	startingRage float64
	currentRage  float64

	onRageGain OnRageGain
}

func (unit *Unit) EnableRageBar(startingRage float64, rageMultiplier float64, onRageGain OnRageGain) {
	unit.RegisterAura(Aura{
		Label:    "RageBar",
		Duration: NeverExpires,
		OnReset: func(aura *Aura, sim *Simulation) {
			aura.Activate(sim)
		},
		OnSpellHit: func(aura *Aura, sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
			if spellEffect.Outcome.Matches(OutcomeMiss) {
				return
			}
			if !spellEffect.ProcMask.Matches(ProcMaskWhiteHit) {
				return
			}

			// Need separate check to exclude auto replacers (e.g. Heroic Strike and Cleave).
			if spellEffect.ProcMask.Matches(ProcMaskMeleeMHSpecial) {
				return
			}

			var HitFactor float64
			var BaseSwingSpeed float64

			if spellEffect.IsMH() {
				HitFactor = 3.5 / 2
				BaseSwingSpeed = unit.AutoAttacks.MH.SwingSpeed
			} else {
				HitFactor = 1.75 / 2
				BaseSwingSpeed = unit.AutoAttacks.OH.SwingSpeed
			}

			if spellEffect.Outcome.Matches(OutcomeCrit) {
				HitFactor *= 2
			}

			damage := spellEffect.Damage
			if spellEffect.Outcome.Matches(OutcomeDodge | OutcomeParry) {
				// Rage is still generated for dodges/parries, based on the damage it WOULD have done.
				damage = spellEffect.PreoutcomeDamage
			}

			generatedRage := damage*RageFactor + HitFactor*BaseSwingSpeed*rageMultiplier

			unit.AddRage(sim, generatedRage, spell.ActionID)
		},
	})

	unit.rageBar = rageBar{
		unit:         unit,
		startingRage: MaxFloat(0, MinFloat(startingRage, MaxRage)),
		onRageGain:   onRageGain,
	}
}

func (unit *Unit) HasRageBar() bool {
	return unit.rageBar.unit != nil
}

func (rb *rageBar) CurrentRage() float64 {
	return rb.currentRage
}

func (rb *rageBar) AddRage(sim *Simulation, amount float64, actionID ActionID) {
	if amount < 0 {
		panic("Trying to add negative rage!")
	}

	newRage := MinFloat(rb.currentRage+amount, MaxRage)
	rb.unit.Metrics.AddResourceEvent(actionID, proto.ResourceType_ResourceTypeRage, amount, newRage-rb.currentRage)

	if sim.Log != nil {
		rb.unit.Log(sim, "Gained %0.3f rage from %s (%0.3f --> %0.3f).", amount, actionID, rb.currentRage, newRage)
	}

	rb.currentRage = newRage
	rb.onRageGain(sim)
}

func (rb *rageBar) SpendRage(sim *Simulation, amount float64, actionID ActionID) {
	if amount < 0 {
		panic("Trying to spend negative rage!")
	}

	newRage := rb.currentRage - amount
	rb.unit.Metrics.AddResourceEvent(actionID, proto.ResourceType_ResourceTypeRage, -amount, -amount)

	if sim.Log != nil {
		rb.unit.Log(sim, "Spent %0.3f rage from %s (%0.3f --> %0.3f).", amount, actionID, rb.currentRage, newRage)
	}

	rb.currentRage = newRage
}

func (rb *rageBar) reset(sim *Simulation) {
	if rb.unit == nil {
		return
	}

	rb.currentRage = rb.startingRage
}
