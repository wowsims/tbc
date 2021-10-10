package core

import (
	"log"
	"time"

	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

// ItemActivation needs the state from simulator, party, and agent
//  because items can impact all 3. (potions, drums, JC necks, etc)
type ItemActivation func(*Simulation, Agent) Aura

type ActiveItem struct {
	Activate   ItemActivation `json:"-"` // Activatable Ability, produces an aura
	ActivateCD time.Duration  `json:"-"` // cooldown on activation
	CoolID     int32          `json:"-"` // ID used for cooldown
	SharedID   int32          `json:"-"` // ID used for shared item cooldowns (trinkets etc)
}

func AddActiveItem(id int32, ai ActiveItem) {
	_, ok := ActiveItemByID[id]
	if ok {
		log.Fatalf("Duplicate active item added: %d, %#v", id, ai)
	}
	ActiveItemByID[id] = ai
}

var ActiveItemByID = map[int32]ActiveItem{
	// Gems
	34220: {Activate: ActivateCSD, ActivateCD: NeverExpires},
	35503: {Activate: ActivateESD, ActivateCD: NeverExpires},
	25893: {Activate: ActivateMSD, ActivateCD: NeverExpires},
	25901: {Activate: ActivateIED, ActivateCD: NeverExpires},

	// Trinkets
	27683: {Activate: ActivateQuagsEye, ActivateCD: NeverExpires, SharedID: MagicIDAtkTrinket},
	29370: {Activate: createSpellDmgActivate(MagicIDBlessingSilverCrescent, 155, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDISCTrink, SharedID: MagicIDAtkTrinket},
	23046: {Activate: createSpellDmgActivate(MagicIDSpellPower, 130, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDEssSappTrink, SharedID: MagicIDAtkTrinket},
	29132: {Activate: createSpellDmgActivate(MagicIDSpellPower, 150, time.Second*15), ActivateCD: time.Second * 90, CoolID: MagicIDScryerTrink, SharedID: MagicIDAtkTrinket},
	24126: {Activate: createSpellDmgActivate(MagicIDRubySerpent, 150, time.Second*20), ActivateCD: time.Second * 300, CoolID: MagicIDRubySerpentTrink, SharedID: MagicIDAtkTrinket},
	29179: {Activate: createSpellDmgActivate(MagicIDSpellPower, 150, time.Second*15), ActivateCD: time.Second * 90, CoolID: MagicIDXiriTrink, SharedID: MagicIDAtkTrinket},
	28418: {Activate: ActivateNexusHorn, ActivateCD: NeverExpires, SharedID: MagicIDAtkTrinket},
	31856: {Activate: ActivateDCC, ActivateCD: NeverExpires, SharedID: MagicIDAtkTrinket},
	28785: {Activate: ActivateTLC, ActivateCD: NeverExpires, SharedID: MagicIDAtkTrinket},
	28789: {Activate: ActivateEyeOfMag, ActivateCD: NeverExpires, SharedID: MagicIDAtkTrinket},
	30626: {Activate: ActivateSextant, ActivateCD: NeverExpires, SharedID: MagicIDAtkTrinket},
	34429: {Activate: createSpellDmgActivate(MagicIDShiftingNaaru, 320, time.Second*15), ActivateCD: time.Second * 90, CoolID: MagicIDShiftingNaaruTrink, SharedID: MagicIDAtkTrinket},
	32483: {Activate: createHasteActivate(MagicIDSkullGuldan, 175, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDSkullGuldanTrink, SharedID: MagicIDAtkTrinket},
	33829: {Activate: createSpellDmgActivate(MagicIDHexShunkHead, 211, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDHexTrink, SharedID: MagicIDAtkTrinket},
	29376: {Activate: createSpellDmgActivate(MagicIDSpellPower, 99, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDEssMartyrTrink, SharedID: MagicIDHealTrinket},
	38290: {Activate: createSpellDmgActivate(MagicIDDarkIronPipeweed, 155, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDDITrink, SharedID: MagicIDAtkTrinket},
	30663: {Activate: ActivateFathomBrooch, ActivateCD: NeverExpires, SharedID: MagicIDAtkTrinket},
	35749: {Activate: ActivateAlchStone, ActivateCD: NeverExpires, SharedID: MagicIDAtkTrinket},
	23207: {Activate: ActivateMarkOfTheChampion, ActivateCD: NeverExpires},

	// Armor
	28602: {Activate: ActivateElderScribes, ActivateCD: NeverExpires},
}

type ItemSet struct {
	Name    string
	Items   map[int32]struct{}     // map[key]struct{} is roughly a set in go.
	Bonuses map[int]ItemActivation // maps item count to activations
}

func AddItemSet(set ItemSet) {
	// TODO: validate the set doesnt exist already?

	setIdx := len(sets)
	sets = append(sets, set)
	for itemID := range set.Items {
		itemSetLookup[itemID] = &sets[setIdx]
	}
}

// cache for mapping item to set for fast resetting of sim.
var itemSetLookup = map[int32]*ItemSet{}

var sets = []ItemSet{
	{
		Name:  "The Twin Stars",
		Items: map[int32]struct{}{31338: {}, 31339: {}},
		Bonuses: map[int]ItemActivation{2: func(sim *Simulation, agent Agent) Aura {
			agent.GetCharacter().Stats[stats.SpellPower] += 15
			return Aura{ID: MagicIDNetherstrike}
		}},
	},
	{
		Name:    "Spellstrike",
		Items:   map[int32]struct{}{24266: {}, 24262: {}},
		Bonuses: map[int]ItemActivation{2: ActivateSpellstrike},
	},
	{
		Name:  "Mana Etched",
		Items: map[int32]struct{}{28193: {}, 27465: {}, 27907: {}, 27796: {}, 28191: {}},
		Bonuses: map[int]ItemActivation{4: ActivateManaEtched, 2: func(sim *Simulation, agent Agent) Aura {
			agent.GetCharacter().Stats[stats.SpellHit] += 35
			return Aura{ID: MagicIDManaEtchedHit}
		}},
	},
	{
		Name:  "Windhawk",
		Items: map[int32]struct{}{29524: {}, 29523: {}, 29522: {}},
		Bonuses: map[int]ItemActivation{3: func(sim *Simulation, agent Agent) Aura {
			agent.GetCharacter().Stats[stats.MP5] += 8
			return Aura{ID: MagicIDWindhawk}
		}},
	},
}

func init() {
	// pre-cache item to set lookup for faster sim resetting.
	for _, v := range items.Items {
		setFound := false
		for setIdx, set := range sets {
			if _, ok := set.Items[v.ID]; ok {
				itemSetLookup[v.ID] = &sets[setIdx]
				setFound = true
				break
			}
		}
		if !setFound {
			itemSetLookup[v.ID] = nil
		}
	}
}
