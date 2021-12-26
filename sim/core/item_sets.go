package core

import (
	"fmt"
)

type ItemSet struct {
	Name string

	// IDs of items that are part of this set. map[key]struct{} is roughly a set in go.
	Items map[int32]struct{}

	// Maps set piece requirement to an ApplyEffect function that will be called
	// before the Sim starts.
	//
	// The function should apply any benefits provided by the set bonus.
	Bonuses map[int32]ApplyEffect
}

func (set ItemSet) ItemIDs() []int32 {
	ids := []int32{}
	for id, _ := range set.Items {
		ids = append(ids, id)
	}
	return ids
}

func (set ItemSet) ItemIsInSet(itemID int32) bool {
	_, ok := set.Items[itemID]
	return ok
}

func (set ItemSet) CharacterHasSetBonus(character *Character, numItems int32) bool {
	if _, ok := set.Bonuses[numItems]; !ok {
		panic(fmt.Sprintf("Item set %s does not have a bonus with %d pieces.", set.Name, numItems))
	}

	count := int32(0)
	for _, item := range character.Equip {
		if set.ItemIsInSet(item.ID) {
			count++
		}
	}

	return count >= numItems
}

var sets = []ItemSet{}

func GetAllItemSets() []ItemSet {
	// Defensive copy to prevent modifications.
	tmp := make([]ItemSet, len(sets))
	copy(tmp, sets)
	return tmp
}

// cache for mapping item to set for fast resetting of sim.
var itemSetLookup = map[int32]*ItemSet{}

func AddItemSet(set ItemSet) {
	// TODO: validate the set doesnt exist already?

	setIdx := len(sets)
	sets = append(sets, set)
	for itemID := range set.Items {
		itemSetLookup[itemID] = &sets[setIdx]
	}
}

type ActiveSetBonus struct {
	// Name of the set.
	Name string

	// Number of pieces required for this bonus.
	NumPieces int32

	// Function for applying the effects of this set bonus.
	BonusEffect ApplyEffect
}

// Returns a list describing all active set bonuses.
func (character *Character) GetActiveSetBonuses() []ActiveSetBonus {
	activeBonuses := []ActiveSetBonus{}
	setItemCount := map[string]int32{}

	for _, i := range character.Equip {
		set := itemSetLookup[i.ID]
		if set != nil {
			setItemCount[set.Name]++
			if setBonusFunc, ok := set.Bonuses[setItemCount[set.Name]]; ok {
				activeBonuses = append(activeBonuses, ActiveSetBonus{
					Name:        set.Name,
					NumPieces:   setItemCount[set.Name],
					BonusEffect: setBonusFunc,
				})
			}
		}
	}

	return activeBonuses
}

// Apply effects from item set bonuses.
func (character *Character) applyItemSetBonusEffects(agent Agent) {
	activeSetBonuses := character.GetActiveSetBonuses()

	for _, activeSetBonus := range activeSetBonuses {
		activeSetBonus.BonusEffect(agent)
	}
}

// Returns the names of all active set bonuses.
func (character *Character) GetActiveSetBonusNames() []string {
	activeSetBonuses := character.GetActiveSetBonuses()
	names := []string{}

	for _, activeSetBonus := range activeSetBonuses {
		names = append(names, fmt.Sprintf("%s (%dpc)", activeSetBonus.Name, activeSetBonus.NumPieces))
	}

	return names
}
