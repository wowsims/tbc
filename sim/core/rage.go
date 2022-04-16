package core

import (
	"github.com/wowsims/tbc/sim/core/proto"
)

const MaxRage = 100.0

const RageFactor = 3.75 / 274.7

// OnRageGain is called any time rage is increased.
type OnRageGain func(sim *Simulation)

type rageBar struct {
	character *Character

	startingRage float64
	currentRage  float64

	onRageGain OnRageGain
}

func (character *Character) EnableRageBar(startingRage float64, onRageGain OnRageGain) {
	character.RegisterAura(Aura{
		Label:    "RageBar",
		Duration: NeverExpires,
		OnReset: func(aura *Aura, sim *Simulation) {
			aura.Activate(sim)
		},
		OnSpellHit: func(aura *Aura, sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
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
				BaseSwingSpeed = character.AutoAttacks.MH.SwingSpeed
			} else {
				HitFactor = 1.75 / 2
				BaseSwingSpeed = character.AutoAttacks.OH.SwingSpeed
			}

			if spellEffect.Outcome.Matches(OutcomeCrit) {
				HitFactor *= 2
			}

			generatedRage := spellEffect.Damage*RageFactor + HitFactor*BaseSwingSpeed

			character.AddRage(sim, generatedRage, spell.ActionID)
		},
	})

	character.rageBar = rageBar{
		character:    character,
		startingRage: MaxFloat(0, MinFloat(startingRage, MaxRage)),
		onRageGain:   onRageGain,
	}
}

func (character *Character) HasRageBar() bool {
	return character.rageBar.character != nil
}

func (rb *rageBar) CurrentRage() float64 {
	return rb.currentRage
}

func (rb *rageBar) AddRage(sim *Simulation, amount float64, actionID ActionID) {
	if amount < 0 {
		panic("Trying to add negative rage!")
	}

	newRage := MinFloat(rb.currentRage+amount, MaxRage)
	rb.character.Metrics.AddResourceEvent(actionID, proto.ResourceType_ResourceTypeRage, amount, newRage-rb.currentRage)

	if sim.Log != nil {
		rb.character.Log(sim, "Gained %0.3f rage from %s (%0.3f --> %0.3f).", amount, actionID, rb.currentRage, newRage)
	}

	rb.currentRage = newRage
	rb.onRageGain(sim)
}

func (rb *rageBar) SpendRage(sim *Simulation, amount float64, actionID ActionID) {
	if amount < 0 {
		panic("Trying to spend negative rage!")
	}

	newRage := rb.currentRage - amount
	rb.character.Metrics.AddResourceEvent(actionID, proto.ResourceType_ResourceTypeRage, -amount, -amount)

	if sim.Log != nil {
		rb.character.Log(sim, "Spent %0.3f rage from %s (%0.3f --> %0.3f).", amount, actionID, rb.currentRage, newRage)
	}

	rb.currentRage = newRage
}

func (rb *rageBar) reset(sim *Simulation) {
	if rb.character == nil {
		return
	}

	rb.currentRage = rb.startingRage
}
