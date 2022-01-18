package core

import ()

const MaxRage = 100.0

var RageBarAuraID = NewAuraID()

type rageBar struct {
	character *Character

	startingRage float64
	currentRage  float64
}

func (character *Character) EnableRageBar(startingRage float64) {
	character.AddPermanentAura(func(sim *Simulation) Aura {
		return Aura{
			ID: RageBarAuraID,
			OnMeleeAttack: func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect) {
				if !hitEffect.IsWhiteHit {
					return
				}

				damage := hitEffect.Damage

				// TODO: Put the actual equation here.
				generatedRage := damage

				character.AddRage(sim, generatedRage, ability.ActionID)
			},
		}
	})
	character.rageBar = rageBar{
		character:    character,
		startingRage: MaxFloat(0, MinFloat(startingRage, MaxRage)),
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

	if sim.Log != nil {
		rb.character.Log(sim, "Gained %0.3f rage from %s (%0.3f --> %0.3f).", amount, actionID, rb.currentRage, newRage)
	}

	rb.currentRage = newRage
}

func (rb *rageBar) SpendRage(sim *Simulation, amount float64, actionID ActionID) {
	if amount < 0 {
		panic("Trying to spend negative rage!")
	}

	newRage := rb.currentRage - amount

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
