package core

import (
	"log"
)

// Function for applying permanent effects to an Agent.
//
// Passing Character instead of Agent would work for almost all cases,
// but there are occasionally class-specific item effects.
type ApplyEffect func(Agent)

var itemEffects = map[int32]ApplyEffect{}

func HasItemEffect(id int32) bool {
	_, ok := itemEffects[id]
	return ok
}

// Registers an ApplyEffect function which will be called before the Sim
// starts, for any Agent that is wearing the item.
func AddItemEffect(id int32, itemEffect ApplyEffect) {
	if HasItemEffect(id) {
		log.Fatalf("Cannot add multiple effects for one item: %d, %#v", id, itemEffect)
	}
	itemEffects[id] = itemEffect
}
