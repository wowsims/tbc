package core

import (
	"log"
	"time"

	"github.com/wowsims/tbc/items"
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

	// Armor
	28602: {Activate: ActivateElderScribes, ActivateCD: NeverExpires},
}

type ItemSet struct {
	Name    string
	Items   map[string]bool
	Bonuses map[int]ItemActivation // maps item count to activations
}

func AddItemSet(set ItemSet) {
	// TODO: validate the set doesnt exist already probably?
	sets = append(sets, set)
}

// cache for mapping item to set for fast resetting of sim.
var itemSetLookup = map[int32]*ItemSet{}

var sets = []ItemSet{
	{
		Name:  "Netherstrike",
		Items: map[string]bool{"Netherstrike Breastplate": true, "Netherstrike Bracers": true, "Netherstrike Belt": true},
		Bonuses: map[int]ItemActivation{3: func(sim *Simulation, agent Agent) Aura {
			agent.GetCharacter().Stats[stats.SpellPower] += 23
			return Aura{ID: MagicIDNetherstrike}
		}},
	},
	{
		Name:  "The Twin Stars",
		Items: map[string]bool{"Charlotte's Ivy": true, "Lola's Eve": true},
		Bonuses: map[int]ItemActivation{2: func(sim *Simulation, agent Agent) Aura {
			agent.GetCharacter().Stats[stats.SpellPower] += 15
			return Aura{ID: MagicIDNetherstrike}
		}},
	},
	{
		Name:  "Tidefury",
		Items: map[string]bool{"Tidefury Helm": true, "Tidefury Shoulderguards": true, "Tidefury Chestpiece": true, "Tidefury Kilt": true, "Tidefury Gauntlets": true},
		Bonuses: map[int]ItemActivation{
			2: func(sim *Simulation, agent Agent) Aura {
				return Aura{ID: MagicIDTidefury}
			},
			4: func(sim *Simulation, agent Agent) Aura {
				// TODO: should we even allow for unchecking water shield?
				// if sim.Options.Buffs.WaterShield {
				agent.GetCharacter().Stats[stats.MP5] += 3
				// }
				return Aura{ID: MagicIDTidefury}
			},
		},
	},
	{
		Name:    "Spellstrike",
		Items:   map[string]bool{"Spellstrike Hood": true, "Spellstrike Pants": true},
		Bonuses: map[int]ItemActivation{2: ActivateSpellstrike},
	},
	{
		Name:  "Mana Etched",
		Items: map[string]bool{"Mana-Etched Crown": true, "Mana-Etched Spaulders": true, "Mana-Etched Vestments": true, "Mana-Etched Gloves": true, "Mana-Etched Pantaloons": true},
		Bonuses: map[int]ItemActivation{4: ActivateManaEtched, 2: func(sim *Simulation, agent Agent) Aura {
			agent.GetCharacter().Stats[stats.SpellHit] += 35
			return Aura{ID: MagicIDManaEtchedHit}
		}},
	},
	{
		Name:  "Cyclone Regalia",
		Items: map[string]bool{"Cyclone Faceguard": true, "Cyclone Shoulderguards": true, "Cyclone Chestguard": true, "Cyclone Handguards": true, "Cyclone Legguards": true},
		Bonuses: map[int]ItemActivation{4: ActivateCycloneManaReduce, 2: func(sim *Simulation, agent Agent) Aura {
			// if sim.Options.Totems.WrathOfAir {

			// FUTURE: Only one ele shaman in the party can use this at a time.
			//   not a big deal now but will need to be fixed to support full raid sim.
			agent.GetCharacter().Party.AddStats(stats.Stats{stats.SpellPower: 20})
			// }
			return Aura{ID: MagicIDCyclone2pc}
		}},
	},
	{
		Name:  "Windhawk",
		Items: map[string]bool{"Windhawk Hauberk": true, "Windhawk Belt": true, "Windhawk Bracers": true},
		Bonuses: map[int]ItemActivation{3: func(sim *Simulation, agent Agent) Aura {
			agent.GetCharacter().Stats[stats.MP5] += 8
			return Aura{ID: MagicIDWindhawk}
		}},
	},
	{
		Name:    "Cataclysm Regalia",
		Items:   map[string]bool{"Cataclysm Headpiece": true, "Cataclysm Shoulderpads": true, "Cataclysm Chestpiece": true, "Cataclysm Handgrips": true, "Cataclysm Leggings": true},
		Bonuses: map[int]ItemActivation{4: ActivateCataclysmLBDiscount},
	},
	{
		Name: "Skyshatter Regalia",
		Items: map[string]bool{
			"Skyshatter Headguard":   true,
			"Skyshatter Mantle":      true,
			"Skyshatter Breastplate": true,
			"Skyshatter Gauntlets":   true,
			"Skyshatter Legguards":   true,
			"Skyshatter Cord":        true,
			"Skyshatter Treads":      true,
			"Skyshatter Bands":       true,
		},
		Bonuses: map[int]ItemActivation{2: func(sim *Simulation, agent Agent) Aura {
			agent.GetCharacter().Stats[stats.MP5] += 15
			agent.GetCharacter().Stats[stats.SpellCrit] += 35
			agent.GetCharacter().Stats[stats.SpellPower] += 45
			return Aura{ID: MagicIDSkyshatter2pc}
		}, 4: ActivateSkyshatterImpLB},
	},
}

func init() {
	// pre-cache item to set lookup for faster sim resetting.
	for _, v := range items.Items {
		setFound := false
		for setIdx, set := range sets {
			if set.Items[v.Name] {
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
