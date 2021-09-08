package core

import (
	"fmt"
	"time"
)

var Gems = []Gem{
	{ID: 34220, Name: "Chaotic Skyfire Diamond", Quality: ItemQualityRare, Phase: 1, Color: GemColorMeta, Stats: Stats{StatSpellCrit: 12}, Activate: ActivateCSD},
	{ID: 25897, Name: "Bracing Earthstorm Diamond", Quality: ItemQualityRare, Phase: 1, Color: GemColorMeta, Stats: Stats{StatSpellDmg: 14}},
	{ID: 32641, Name: "Imbued Unstable Diamond", Quality: ItemQualityRare, Phase: 1, Color: GemColorMeta, Stats: Stats{StatSpellDmg: 14}},
	{ID: 35503, Name: "Ember Skyfire Diamond", Quality: ItemQualityRare, Phase: 1, Color: GemColorMeta, Stats: Stats{StatSpellDmg: 14}, Activate: ActivateESD},
	{ID: 28557, Name: "Swift Starfire Diamond", Quality: ItemQualityRare, Phase: 1, Color: GemColorMeta, Stats: Stats{StatSpellDmg: 12}},
	{ID: 25893, Name: "Mystical Skyfire Diamond", Quality: ItemQualityRare, Phase: 1, Color: GemColorMeta, Stats: Stats{}, Activate: ActivateMSD},
	{ID: 25901, Name: "Insightful Earthstorm Diamond", Quality: ItemQualityRare, Phase: 1, Color: GemColorMeta, Stats: Stats{StatInt: 12}, Activate: ActivateIED},
	{ID: 23096, Name: "Runed Blood Garnet", Quality: ItemQualityUncommon, Phase: 1, Color: GemColorRed, Stats: Stats{StatSpellDmg: 7}},
	{ID: 24030, Name: "Runed Living Ruby", Quality: ItemQualityRare, Phase: 1, Color: GemColorRed, Stats: Stats{StatSpellDmg: 9}},
	{ID: 32196, Name: "Runed Crimson Spinel", Quality: ItemQualityEpic, Phase: 3, Color: GemColorRed, Stats: Stats{StatSpellDmg: 12}},
	{ID: 28118, Name: "Runed Ornate Ruby", Quality: ItemQualityEpic, Phase: 1, Color: GemColorRed, Stats: Stats{StatSpellDmg: 12}},
	{ID: 33133, Name: "Don Julio's Heart", Quality: ItemQualityEpic, Phase: 1, Color: GemColorRed, Stats: Stats{StatSpellDmg: 14}},
	{ID: 23121, Name: "Lustrous Azure Moonstone", Quality: ItemQualityUncommon, Phase: 1, Color: GemColorBlue, Stats: Stats{StatMP5: 2}},
	{ID: 24037, Name: "Lustrous Star of Elune", Quality: ItemQualityRare, Phase: 1, Color: GemColorBlue, Stats: Stats{StatMP5: 3}},
	{ID: 32202, Name: "Lustrous Empyrean Sapphire", Quality: ItemQualityEpic, Phase: 1, Color: GemColorBlue, Stats: Stats{StatMP5: 4}},
	{ID: 23113, Name: "Brilliant Golden Draenite", Quality: ItemQualityUncommon, Phase: 1, Color: GemColorYellow, Stats: Stats{StatInt: 6}},
	{ID: 24047, Name: "Brilliant Dawnstone", Quality: ItemQualityRare, Phase: 1, Color: GemColorYellow, Stats: Stats{StatInt: 8}},
	{ID: 32204, Name: "Brilliant Lionseye", Quality: ItemQualityEpic, Phase: 3, Color: GemColorYellow, Stats: Stats{StatInt: 10}},
	{ID: 23114, Name: "Gleaming Golden Draenite", Quality: ItemQualityUncommon, Phase: 1, Color: GemColorYellow, Stats: Stats{StatSpellCrit: 6}},
	{ID: 24050, Name: "Gleaming Dawnstone", Quality: ItemQualityRare, Phase: 1, Color: GemColorYellow, Stats: Stats{StatSpellCrit: 8}},
	{ID: 32207, Name: "Gleaming Lionseye", Quality: ItemQualityEpic, Phase: 3, Color: GemColorYellow, Stats: Stats{StatSpellCrit: 10}},
	{ID: 30551, Name: "Infused Fire Opal", Quality: ItemQualityEpic, Phase: 1, Color: GemColorOrange, Stats: Stats{StatInt: 4, StatSpellDmg: 6}},
	{ID: 23101, Name: "Potent Flame Spessarite", Quality: ItemQualityUncommon, Phase: 1, Color: GemColorOrange, Stats: Stats{StatSpellCrit: 3, StatSpellDmg: 4}},
	{ID: 24059, Name: "Potent Noble Topaz", Quality: ItemQualityRare, Phase: 1, Color: GemColorOrange, Stats: Stats{StatSpellCrit: 4, StatSpellDmg: 5}},
	{ID: 32218, Name: "Potent Pyrestone", Quality: ItemQualityEpic, Phase: 3, Color: GemColorOrange, Stats: Stats{StatSpellCrit: 5, StatSpellDmg: 6}},
	{ID: 35760, Name: "Reckless Pyrestone", Quality: ItemQualityEpic, Phase: 3, Color: GemColorOrange, Stats: Stats{StatHaste: 5, StatSpellDmg: 6}},
	{ID: 30588, Name: "Potent Fire Opal", Quality: ItemQualityEpic, Phase: 1, Color: GemColorOrange, Stats: Stats{StatSpellDmg: 6, StatSpellCrit: 4}},
	{ID: 28123, Name: "Potent Ornate Topaz", Quality: ItemQualityEpic, Phase: 1, Color: GemColorOrange, Stats: Stats{StatSpellDmg: 6, StatSpellCrit: 5}},
	{ID: 31866, Name: "Veiled Flame Spessarite", Quality: ItemQualityUncommon, Phase: 1, Color: GemColorOrange, Stats: Stats{StatSpellHit: 3, StatSpellDmg: 4}},
	{ID: 31867, Name: "Veiled Noble Topaz", Quality: ItemQualityRare, Phase: 1, Color: GemColorOrange, Stats: Stats{StatSpellHit: 4, StatSpellDmg: 5}},
	{ID: 32221, Name: "Shining Fire Opal", Quality: ItemQualityEpic, Phase: 1, Color: GemColorOrange, Stats: Stats{StatSpellHit: 5, StatSpellDmg: 6}},
	{ID: 30564, Name: "Veiled Pyrestone", Quality: ItemQualityEpic, Phase: 3, Color: GemColorOrange, Stats: Stats{StatSpellHit: 5, StatSpellDmg: 6}},
	{ID: 30560, Name: "Rune Covered Chrysoprase", Quality: ItemQualityEpic, Phase: 1, Color: GemColorGreen, Stats: Stats{StatMP5: 2, StatSpellCrit: 5}},
	{ID: 24065, Name: "Dazzling Talasite", Quality: ItemQualityRare, Phase: 1, Color: GemColorGreen, Stats: Stats{StatMP5: 2, StatInt: 4}},
	{ID: 35759, Name: "Forceful Seaspray Emerald", Quality: ItemQualityEpic, Phase: 3, Color: GemColorGreen, Stats: Stats{StatHaste: 5, StatStm: 7}},
	{ID: 24056, Name: "Glowing Nightseye", Quality: ItemQualityRare, Phase: 1, Color: GemColorPurple, Stats: Stats{StatSpellDmg: 5, StatStm: 6}},
	{ID: 30555, Name: "Glowing Tanzanite", Quality: ItemQualityEpic, Phase: 1, Color: GemColorPurple, Stats: Stats{StatSpellDmg: 6, StatStm: 6}},
	{ID: 32215, Name: "Glowing Shadowsong Amethyst", Quality: ItemQualityEpic, Phase: 3, Color: GemColorPurple, Stats: Stats{StatSpellDmg: 6, StatStm: 7}},
	{ID: 31116, Name: "Infused Amethyst", Quality: ItemQualityEpic, Phase: 1, Color: GemColorPurple, Stats: Stats{StatSpellDmg: 6, StatStm: 6}},
	{ID: 30600, Name: "Fluorescent Tanzanite", Quality: ItemQualityEpic, Phase: 1, Color: GemColorPurple, Stats: Stats{StatSpellDmg: 6, StatSpirit: 4}},
	{ID: 30605, Name: "Vivid Chrysoprase", Quality: ItemQualityEpic, Phase: 1, Color: GemColorGreen, Stats: Stats{StatSpellHit: 5, StatStm: 6}},
}

var Enchants = []Enchant{
	{ID: 29191, EffectID: 3002, Name: "Glyph of Power", Bonus: Stats{StatSpellDmg: 22, StatSpellHit: 14}, Slot: EquipHead},
	{ID: 28909, EffectID: 2995, Name: "Greater Inscription of the Orb", Bonus: Stats{StatSpellDmg: 12, StatSpellCrit: 15}, Slot: EquipShoulder},
	{ID: 28886, EffectID: 2982, Name: "Greater Inscription of Discipline", Bonus: Stats{StatSpellDmg: 18, StatSpellCrit: 10}, Slot: EquipShoulder},
	{ID: 24421, EffectID: 2605, Name: "Zandalar Signet of Mojo", Bonus: Stats{StatSpellDmg: 18}, Slot: EquipUnknown}, // This should no longer show up in the UI.
	{ID: 20076, EffectID: 2605, Name: "Zandalar Signet of Mojo", Bonus: Stats{StatSpellDmg: 18}, Slot: EquipShoulder},
	{ID: 23545, EffectID: 2721, Name: "Power of the Scourge", Bonus: Stats{StatSpellDmg: 15, StatSpellCrit: 14}, Slot: EquipShoulder},
	{ID: 27960, EffectID: 2661, Name: "Chest - Exceptional Stats", Bonus: Stats{StatStm: 6, StatInt: 6, StatSpirit: 6}, Slot: EquipChest},
	{ID: 27917, EffectID: 2650, Name: "Bracer - Spellpower", Bonus: Stats{StatSpellDmg: 15}, Slot: EquipUnknown}, // This should no longer show up in the UI.
	{ID: 22534, EffectID: 2650, Name: "Bracer - Spellpower", Bonus: Stats{StatSpellDmg: 15}, Slot: EquipWrist},
	{ID: 33997, EffectID: 2937, Name: "Gloves - Major Spellpower", Bonus: Stats{StatSpellDmg: 20}, Slot: EquipUnknown}, // This should no longer show up in the UI.
	{ID: 28272, EffectID: 2937, Name: "Gloves - Major Spellpower", Bonus: Stats{StatSpellDmg: 20}, Slot: EquipHands},
	{ID: 24274, EffectID: 2748, Name: "Runic Spellthread", Bonus: Stats{StatSpellDmg: 35, StatStm: 20}, Slot: EquipLegs},
	{ID: 24273, EffectID: 2747, Name: "Mystic Spellthread", Bonus: Stats{StatSpellDmg: 25, StatStm: 15}, Slot: EquipLegs},
	{ID: 27975, EffectID: 2669, Name: "Weapon - Major Spellpower", Bonus: Stats{StatSpellDmg: 40}, Slot: EquipUnknown}, // This should no longer show up in the UI.
	{ID: 22555, EffectID: 2669, Name: "Weapon - Major Spellpower", Bonus: Stats{StatSpellDmg: 40}, Slot: EquipWeapon},
	{ID: 35445, EffectID: 2928, Name: "Ring - Spellpower", Bonus: Stats{StatSpellDmg: 12}, Slot: EquipFinger},
	{ID: 27945, EffectID: 2654, Name: "Shield - Intellect", Bonus: Stats{StatInt: 12}, Slot: EquipOffhand},
}

var ItemsByName = map[string]Item{}
var ItemsByID = map[int32]Item{}
var GemsByName = map[string]Gem{}
var GemsByID = map[int32]Gem{}
var EnchantsByName = map[string]Enchant{}
var EnchantsByID = map[int32]Enchant{}

func init() {
	for _, v := range Enchants {
		EnchantsByName[v.Name] = v
		EnchantsByID[v.ID] = v
	}
	for _, v := range Gems {
		GemsByName[v.Name] = v
		GemsByID[v.ID] = v
	}
	for _, v := range items {
		if _, ok := ItemsByID[v.ID]; ok {
			fmt.Printf("Found dup item: %s\n", v.Name)
			panic("no dupes allowed")
		}
		if it, ok := ItemsByName[v.Name]; ok {
			fmt.Printf("Found dup item: %s\n", v.Name)
			statsMatch := it.Slot == v.Slot
			for i, v := range v.Stats {
				if len(it.Stats) <= i {
					break
				}
				if it.Stats[i] != v {
					statsMatch = false
				}
			}
			if !statsMatch {
				// log.Printf("Mismatched slot/stats: \n\tMoreItem: \n%#v\n\t FirstSet: \n%#v", it, v)
			}
		} else {
			ItemsByName[v.Name] = v
			ItemsByID[v.ID] = v
		}
		// pre-cache item to set lookup for faster sim resetting.
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

type Item struct {
	ID         int32
	Slot       byte
	SubSlot    byte `json:"subSlot,omitempty"`
	Name       string
	SourceZone string
	SourceDrop string
	Stats      Stats // Stats applied to wearer
	Phase      byte
	Quality    ItemQuality

	GemSlots    []GemColor
	SocketBonus Stats

	// Modified for each instance of the item.
	Gems    []Gem
	Enchant Enchant

	// For simplicity all items that produce an aura are 'activatable'.
	// Since we activate all items on CD, this works fine for stuff like Quags Eye.
	// TODO: is this the best design for this?
	Activate   ItemActivation `json:"-"` // Activatable Ability, produces an aura
	ActivateCD time.Duration  `json:"-"` // cooldown on activation
	CoolID     int32          `json:"-"` // ID used for cooldown
}

type ItemQuality byte

const (
	ItemQualityJunk      ItemQuality = iota // anything less than green
	ItemQualityUncommon                     // green
	ItemQualityRare                         // blue
	ItemQualityEpic                         // purple
	ItemQualityLegendary                    // orange
)

type Enchant struct {
	ID       int32 // ID of the enchant item.
	EffectID int32 // Used by UI to apply effect to tooltip
	Name     string
	Bonus    Stats
	Slot     byte // which slot does the enchant go on.
}

type Gem struct {
	ID       int32
	Name     string
	Stats    Stats          // flat stats gem adds
	Activate ItemActivation `json:"-"` // Meta gems activate an aura on player when socketed. Assumes all gems are 'always active'
	Color    GemColor
	Phase    byte
	Quality  ItemQuality
	// Requirements  // Validate the gem can be used... later
}

type GemColor byte

const (
	GemColorUnknown GemColor = iota
	GemColorMeta
	GemColorRed
	GemColorBlue
	GemColorYellow
	GemColorGreen
	GemColorOrange
	GemColorPurple
	GemColorPrismatic
)

func (gm GemColor) Intersects(o GemColor) bool {
	if gm == o {
		return true
	}
	if gm == GemColorPrismatic || o == GemColorPrismatic {
		return true
	}
	if gm == GemColorMeta {
		return false // meta gems o nothing.
	}
	if gm == GemColorRed {
		return o == GemColorOrange || o == GemColorPurple
	}
	if gm == GemColorBlue {
		return o == GemColorGreen || o == GemColorPurple
	}
	if gm == GemColorYellow {
		return o == GemColorGreen || o == GemColorOrange
	}
	if gm == GemColorOrange {
		return o == GemColorYellow || o == GemColorRed
	}
	if gm == GemColorGreen {
		return o == GemColorYellow || o == GemColorBlue
	}
	if gm == GemColorPurple {
		return o == GemColorBlue || o == GemColorRed
	}

	return false // dunno wtf this is.
}

type ItemActivation func(*Simulation) Aura

type Equipment [EquipTotem + 1]Item

// Structs used for looking up items/gems/enchants
type ItemSpec struct {
	// Only name or ID needs to be set, not both
	Name string
	ID   int32

	Enchant EnchantSpec
	Gems    []GemSpec
}
type GemSpec struct {
	// Only name or ID needs to be set, not both
	Name string
	ID   int32
}
type EnchantSpec struct {
	// Only name or ID needs to be set, not both
	Name string
	ID   int32
}
type EquipmentSpec [EquipTotem + 1]ItemSpec

func NewEquipmentSet(equipSpec EquipmentSpec) Equipment {
	equipment := Equipment{}

	for _, itemSpec := range equipSpec {
		item := Item{}
		if foundItem, ok := ItemsByName[itemSpec.Name]; ok {
			item = foundItem
		} else if foundItem, ok := ItemsByID[itemSpec.ID]; ok {
			item = foundItem
		} else {
			if itemSpec.Name != "" {
				panic("No item with name: " + itemSpec.Name)
			} else if itemSpec.ID != 0 {
				panic(fmt.Sprintf("No item with id: %d", itemSpec.ID))
			}
			continue
		}

		if itemSpec.Enchant.Name != "" {
			if enchant, ok := EnchantsByName[itemSpec.Enchant.Name]; ok {
				item.Enchant = enchant
			} else {
				panic("No enchant with name: " + itemSpec.Enchant.Name)
			}
		} else if itemSpec.Enchant.ID != 0 {
			if enchant, ok := EnchantsByID[itemSpec.Enchant.ID]; ok {
				item.Enchant = enchant
			} else {
				panic(fmt.Sprintf("No enchant with id: %d", itemSpec.Enchant.ID))
			}
		}

		if len(itemSpec.Gems) > 0 {
			item.Gems = make([]Gem, len(item.GemSlots))
			for gemIdx, gemSpec := range itemSpec.Gems {
				if gemIdx >= len(item.GemSlots) {
					break // in case we get invalid gem settings.
				}
				if gem, ok := GemsByName[gemSpec.Name]; ok {
					item.Gems[gemIdx] = gem
				} else if gem, ok := GemsByID[gemSpec.ID]; ok {
					item.Gems[gemIdx] = gem
				} else {
					if gemSpec.Name != "" {
						panic("No gem with name: " + gemSpec.Name)
					} else if gemSpec.ID != 0 {
						panic(fmt.Sprintf("No gem with id: %d", gemSpec.ID))
					}
				}
			}
		}

		if item.Slot == EquipFinger {
			if equipment[EquipFinger1].Name == "" {
				equipment[EquipFinger1] = item
			} else {
				equipment[EquipFinger2] = item
			}
		} else if item.Slot == EquipTrinket {
			if equipment[EquipTrinket1].Name == "" {
				equipment[EquipTrinket1] = item
			} else {
				equipment[EquipTrinket2] = item
			}
		} else {
			equipment[item.Slot] = item
		}
	}
	return equipment
}

// subslot consts
const (
	SubslotUnknown byte = iota
	SubslotShield
	SubslotTwoHand
)

// slot consts
const (
	EquipUnknown byte = iota
	EquipHead
	EquipNeck
	EquipShoulder
	EquipBack
	EquipChest
	EquipWrist
	EquipHands
	EquipWaist
	EquipLegs
	EquipFeet
	EquipFinger  // generic finger item
	EquipFinger1 // specific slot in equipment array
	EquipFinger2
	EquipTrinket  // generic trinket
	EquipTrinket1 // specific trinket slot in equip array
	EquipTrinket2
	EquipWeapon // holds 1 or 2h
	EquipOffhand
	EquipTotem
)

func (e Equipment) Clone() Equipment {
	ne := Equipment{}
	for i, v := range e {
		vc := v
		ne[i] = vc
	}
	return ne
}

func (e Equipment) Stats() Stats {
	s := Stats{}
	for _, item := range e {
		for k, v := range item.Stats {
			s[k] += v
		}
		isMatched := len(item.Gems) == len(item.GemSlots) && len(item.GemSlots) > 0
		for gi, g := range item.Gems {
			for k, v := range g.Stats {
				s[k] += v
			}
			isMatched = isMatched && g.Color.Intersects(item.GemSlots[gi])
			if !isMatched {
			}
		}
		if len(item.GemSlots) > 0 {
		}
		if isMatched {
			for k, v := range item.SocketBonus {
				if v == 0 {
					continue
				}
				s[k] += v
			}
		}
		for k, v := range item.Enchant.Bonus {
			s[k] += v
		}
	}
	return s
}

// Hopefully get access to:
// https://docs.google.com/spreadsheets/d/1XkLW3o9VrYg8VT84tCoINq-KxP9EA876RdhIzO-PcQk/edit#gid=1056257705

var items = []Item{
	// source: https://docs.google.com/spreadsheets/d/1X-XO9N1_MPIq-UIpTN13LrhXRoho9fe26YEEM48QmPk/edit#gid=2035379487
	{ID: 27471, Slot: EquipHead, Name: "Gladiator's Mail Helm", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Arena Season 1 Reward", SourceDrop: "", Stats: Stats{15, 54, 18, 0, 37}, GemSlots: []GemColor{0x1, 0x2}, SocketBonus: Stats{}},
	{ID: 24266, Slot: EquipHead, Name: "Spellstrike Hood", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Tailoring BoE", SourceDrop: "", Stats: Stats{12, 16, 24, 16, 46}, GemSlots: []GemColor{0x2, 0x4, 0x3}, SocketBonus: Stats{StatStm: 6}},
	{ID: 28278, Slot: EquipHead, Name: "Incanter's Cowl", Phase: 1, Quality: ItemQualityRare, SourceZone: "Mech - Pathaleon the Calculator", SourceDrop: "", Stats: Stats{27, 15, 19, 0, 29}, GemSlots: []GemColor{0x1, 0x4}, SocketBonus: Stats{}},
	{ID: 31330, Slot: EquipHead, Name: "Lightning Crown", Phase: 1, Quality: ItemQualityEpic, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{0, 0, 43, 0, 66}},
	{ID: 28415, Slot: EquipHead, Name: "Hood of Oblivion", Phase: 1, Quality: ItemQualityRare, SourceZone: "Arc - Harbinger Skyriss", SourceDrop: "", Stats: Stats{32, 27, 0, 0, 40}, GemSlots: []GemColor{0x1, 0x3}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 28758, Slot: EquipHead, Name: "Exorcist's Mail Helm", Phase: 1, Quality: ItemQualityRare, SourceZone: "18 Spirit Shards", SourceDrop: "", Stats: Stats{16, 30, 24, 0, 29}, GemSlots: []GemColor{0x1}, SocketBonus: Stats{StatSpellCrit: 3}},
	{ID: 28349, Slot: EquipHead, Name: "Tidefury Helm", Phase: 1, Quality: ItemQualityRare, SourceZone: "Bot - Warp Splinter", SourceDrop: "", Stats: Stats{26, 32, 0, 0, 32, 0, 6}, GemSlots: []GemColor{0x1, 0x4}, SocketBonus: Stats{StatInt: 4}},
	{ID: 29504, Slot: EquipHead, Name: "Windscale Hood", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Leatherworking BoE", SourceDrop: "", Stats: Stats{18, 16, 37, 0, 44, 0, 10}},
	{ID: 31107, Slot: EquipHead, Name: "Shamanistic Helmet of Second Sight", Phase: 1, Quality: ItemQualityRare, SourceZone: "Teron Gorfiend, I am... - SMV Quest", SourceDrop: "", Stats: Stats{15, 12, 24, 0, 35, 0, 4}, GemSlots: []GemColor{0x4, 0x3, 0x3}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 28193, Slot: EquipHead, Name: "Mana-Etched Crown", Phase: 1, Quality: ItemQualityRare, SourceZone: "BM - Aeonus", SourceDrop: "", Stats: Stats{20, 27, 0, 0, 34}, GemSlots: []GemColor{0x1, 0x2}, SocketBonus: Stats{}},
	{ID: 28169, Slot: EquipHead, Name: "Mag'hari Ritualist's Horns", Phase: 1, Quality: ItemQualityRare, SourceZone: "Hero of the Mag'har - Nagrand quest (Horde)", SourceDrop: "", Stats: Stats{16, 18, 15, 12, 50}},
	{ID: 27488, Slot: EquipHead, Name: "Mage-Collar of the Firestorm", Phase: 1, Quality: ItemQualityRare, SourceZone: "H BF - The Maker", SourceDrop: "", Stats: Stats{33, 32, 23, 0, 39}},
	{ID: 30297, Slot: EquipHead, Name: "Circlet of the Starcaller", Phase: 1, Quality: ItemQualityRare, SourceZone: "Dimensius the All-Devouring - NS Quest", SourceDrop: "", Stats: Stats{18, 27, 18, 0, 47}},
	{ID: 27993, Slot: EquipHead, Name: "Mask of Inner Fire", Phase: 1, Quality: ItemQualityRare, SourceZone: "BM - Chrono Lord Deja", SourceDrop: "", Stats: Stats{33, 30, 22, 0, 37}},
	{ID: 30946, Slot: EquipHead, Name: "Mooncrest Headdress", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "Blast the Infernals! - SMV Quest", SourceDrop: "", Stats: Stats{16, 0, 21, 0, 44}},
	{ID: 28245, Slot: EquipNeck, Name: "Pendant of Dominance", Phase: 1, Quality: ItemQualityEpic, SourceZone: "15,300 Honor & 10 EotS Marks", SourceDrop: "", Stats: Stats{StatInt: 12, StatStm: 31, StatSpellDmg: 26}, GemSlots: []GemColor{0x4}, SocketBonus: Stats{StatSpellCrit: 2}},
	{ID: 28134, Slot: EquipNeck, Name: "Brooch of Heightened Potential", Phase: 1, Quality: ItemQualityRare, SourceZone: "SLabs - Blackheart the Inciter", SourceDrop: "", Stats: Stats{14, 15, 14, 9, 22}},
	{ID: 29333, Slot: EquipNeck, Name: "Torc of the Sethekk Prophet", Phase: 1, Quality: ItemQualityRare, SourceZone: "Brother Against Brother - Auchindoun ", SourceDrop: "", Stats: Stats{18, 0, 21, 0, 19}},
	{ID: 31692, Slot: EquipNeck, Name: "Natasha's Ember Necklace", Phase: 1, Quality: ItemQualityRare, SourceZone: "The Hound-Master - BEM Quest", SourceDrop: "", Stats: Stats{15, 0, 10, 0, 29}},
	{ID: 28254, Slot: EquipNeck, Name: "Warp Engineer's Prismatic Chain", Phase: 1, Quality: ItemQualityRare, SourceZone: "Mech - Mechano Lord Capacitus", SourceDrop: "", Stats: Stats{18, 17, 16, 0, 19}},
	{ID: 27758, Slot: EquipNeck, Name: "Hydra-fang Necklace", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H UB - Ghaz'an", SourceDrop: "", Stats: Stats{16, 17, 0, 16, 19}},
	{ID: 31693, Slot: EquipNeck, Name: "Natasha's Arcane Filament", Phase: 1, Quality: ItemQualityEpic, SourceZone: "The Hound-Master - BEM Quest", SourceDrop: "", Stats: Stats{10, 22, 0, 0, 29}},
	{ID: 27464, Slot: EquipNeck, Name: "Omor's Unyielding Will", Phase: 1, Quality: ItemQualityRare, SourceZone: "H Ramps - Omar the Unscarred", SourceDrop: "", Stats: Stats{19, 19, 0, 0, 25}},
	{ID: 31338, Slot: EquipNeck, Name: "Charlotte's Ivy", Phase: 1, Quality: ItemQualityEpic, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{19, 18, 0, 0, 23}},
	{ID: 27473, Slot: EquipShoulder, Name: "Gladiator's Mail Spaulders", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Arena Season 1 Reward", SourceDrop: "", Stats: Stats{17, 33, 20, 0, 22, 0, 6}, GemSlots: []GemColor{0x2, 0x4}, SocketBonus: Stats{}},
	{ID: 32078, Slot: EquipShoulder, Name: "Pauldrons of Wild Magic", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H SP - Quagmirran", SourceDrop: "", Stats: Stats{28, 21, 23, 0, 33}},
	{ID: 27796, Slot: EquipShoulder, Name: "Mana-Etched Spaulders", Phase: 1, Quality: ItemQualityRare, SourceZone: "H UB - Quagmirran", SourceDrop: "", Stats: Stats{17, 25, 16, 0, 20}, GemSlots: []GemColor{0x2, 0x4}, SocketBonus: Stats{}},
	{ID: 30925, Slot: EquipShoulder, Name: "Spaulders of the Torn-heart", Phase: 1, Quality: ItemQualityRare, SourceZone: "The Cipher of Damnation - SMV Quest", SourceDrop: "", Stats: Stats{7, 10, 18, 0, 40}},
	{ID: 31797, Slot: EquipShoulder, Name: "Elekk Hide Spaulders", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "The Fallen Exarch - Terokkar Forest Quest", SourceDrop: "", Stats: Stats{12, 0, 28, 0, 25}},
	{ID: 27778, Slot: EquipShoulder, Name: "Spaulders of Oblivion", Phase: 1, Quality: ItemQualityRare, SourceZone: "SLabs - Murmur", SourceDrop: "", Stats: Stats{17, 25, 0, 0, 29}, GemSlots: []GemColor{0x4, 0x3}, SocketBonus: Stats{StatSpellHit: 3}},
	{ID: 27802, Slot: EquipShoulder, Name: "Tidefury Shoulderguards", Phase: 1, Quality: ItemQualityRare, SourceZone: "SH - O'mrogg", SourceDrop: "", Stats: Stats{23, 18, 0, 0, 19, 0, 6}, GemSlots: []GemColor{0x2, 0x3}, SocketBonus: Stats{StatSpellDmg: 4}},
	{ID: 27994, Slot: EquipShoulder, Name: "Mantle of Three Terrors", Phase: 1, Quality: ItemQualityRare, SourceZone: "BM - Chrono Lord Deja", SourceDrop: "", Stats: Stats{25, 29, 0, 12, 29}},
	{ID: 25777, Slot: EquipBack, Name: "Ogre Slayer's Cover", Phase: 1, Quality: ItemQualityRare, SourceZone: "Cho'war the Pillager - Nagrand Quest", SourceDrop: "", Stats: Stats{18, 0, 16, 0, 20}},
	{ID: 28269, Slot: EquipBack, Name: "Baba's Cloak of Arcanistry", Phase: 1, Quality: ItemQualityRare, SourceZone: "Mech - Pathaleon the Calculator", SourceDrop: "", Stats: Stats{15, 15, 14, 0, 22}},
	{ID: 29813, Slot: EquipBack, Name: "Cloak of Woven Energy", Phase: 1, Quality: ItemQualityRare, SourceZone: "Hitting the Motherlode - Netherstorm Quest", SourceDrop: "", Stats: Stats{13, 6, 6, 0, 29}},
	{ID: 27981, Slot: EquipBack, Name: "Sethekk Oracle Cloak", Phase: 1, Quality: ItemQualityRare, SourceZone: "SH - Talon King Ikiss", SourceDrop: "", Stats: Stats{18, 18, 0, 12, 22}},
	{ID: 32541, Slot: EquipBack, Name: "Terokk's Wisdom", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Terokk - Skettis Summoned Boss", SourceDrop: "", Stats: Stats{16, 18, 0, 0, 33}},
	{ID: 24252, Slot: EquipBack, Name: "Cloak of the Black Void", Phase: 1, Quality: ItemQualityRare, SourceZone: "Tailoring BoE", SourceDrop: "", Stats: Stats{11, 0, 0, 0, 35}},
	{ID: 31140, Slot: EquipBack, Name: "Cloak of Entropy", Phase: 1, Quality: ItemQualityRare, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{11, 0, 0, 10, 25}},
	{ID: 28379, Slot: EquipBack, Name: "Sergeant's Heavy Cape", Phase: 1, Quality: ItemQualityEpic, SourceZone: "9,435 Honor & 20 AB Marks", SourceDrop: "", Stats: Stats{12, 33, 0, 0, 26}},
	{ID: 27469, Slot: EquipChest, Name: "Gladiator's Mail Armor", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Arena Season 1 Reward", SourceDrop: "", Stats: Stats{23, 42, 23, 0, 32, 0, 7}, GemSlots: []GemColor{0x2, 0x4, 0x4}, SocketBonus: Stats{StatSpellCrit: 4}},
	{ID: 31340, Slot: EquipChest, Name: "Will of Edward the Odd", Phase: 1, Quality: ItemQualityEpic, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{30, 0, 30, 0, 53}},
	{ID: 29129, Slot: EquipChest, Name: "Anchorite's Robe", Phase: 1, Quality: ItemQualityEpic, SourceZone: "The Aldor - Honored", SourceDrop: "", Stats: Stats{38, 16, 0, 0, 29, 0, 18}, GemSlots: []GemColor{0x4, 0x4, 0x3}},
	{ID: 28231, Slot: EquipChest, Name: "Tidefury Chestpiece", Phase: 1, Quality: ItemQualityRare, SourceZone: "Arc - Harbinger Skyriss", SourceDrop: "", Stats: Stats{22, 28, 0, 10, 36, 0, 4}, GemSlots: []GemColor{0x4, 0x4, 0x3}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 29341, Slot: EquipChest, Name: "Auchenai Anchorite's Robe", Phase: 1, Quality: ItemQualityRare, SourceZone: "Everything Will Be Alright - AC Quest", SourceDrop: "", Stats: Stats{StatInt: 24, StatSpellDmg: 28, StatSpellHit: 23}, GemSlots: []GemColor{0x2, 0x4, 0x4}, SocketBonus: Stats{StatSpellCrit: 4}},
	{ID: 28191, Slot: EquipChest, Name: "Mana-Etched Vestments", Phase: 1, Quality: ItemQualityRare, SourceZone: "OHF - Epoch Hunter", SourceDrop: "", Stats: Stats{25, 25, 17, 0, 29, 0, 0}, GemSlots: []GemColor{0x2, 0x4, 0x3}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 31297, Slot: EquipChest, Name: "Robe of the Crimson Order", Phase: 1, Quality: ItemQualityEpic, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{23, 0, 0, 30, 50, 0, 0}},
	{ID: 28342, Slot: EquipChest, Name: "Warp Infused Drape", Phase: 1, Quality: ItemQualityRare, SourceZone: "Bot - Warp Splinter", SourceDrop: "", Stats: Stats{28, 27, 0, 12, 30, 0, 0}, GemSlots: []GemColor{0x2, 0x4, 0x3}, SocketBonus: Stats{}},
	{ID: 28232, Slot: EquipChest, Name: "Robe of Oblivion", Phase: 1, Quality: ItemQualityRare, SourceZone: "SLabs - Murmur", SourceDrop: "", Stats: Stats{20, 30, 0, 0, 40, 0, 0}, GemSlots: []GemColor{0x2, 0x4, 0x3}, SocketBonus: Stats{}},
	{ID: 28229, Slot: EquipChest, Name: "Incanter's Robe", Phase: 1, Quality: ItemQualityRare, SourceZone: "Bot - Warp Splinter", SourceDrop: "", Stats: Stats{22, 24, 8, 0, 29, 0, 0}, GemSlots: []GemColor{0x2, 0x4, 0x4}, SocketBonus: Stats{}},
	{ID: 27824, Slot: EquipChest, Name: "Robe of the Great Dark Beyond", Phase: 1, Quality: ItemQualityRare, SourceZone: "MT - Tavarok", SourceDrop: "", Stats: Stats{30, 25, 23, 0, 39, 0, 0}},
	{ID: 28391, Slot: EquipChest, Name: "Worldfire Chestguard", Phase: 1, Quality: ItemQualityRare, SourceZone: "Arc - Dalliah the Doomsayer", SourceDrop: "", Stats: Stats{32, 33, 22, 0, 40, 0, 0}},
	{ID: 28638, Slot: EquipWrist, Name: "General's Mail Bracers", Phase: 1, Quality: ItemQualityEpic, SourceZone: "7,548 Honor & 20 WSG Marks", SourceDrop: "", Stats: Stats{12, 22, 14, 0, 20, 0, 0}, GemSlots: []GemColor{0x4}, SocketBonus: Stats{}},
	{ID: 27522, Slot: EquipWrist, Name: "World's End Bracers", Phase: 1, Quality: ItemQualityRare, SourceZone: "H BF - Keli'dan the Breaker", SourceDrop: "", Stats: Stats{19, 18, 17, 0, 22, 0, 0}},
	{ID: 24250, Slot: EquipWrist, Name: "Bracers of Havok", Phase: 1, Quality: ItemQualityRare, SourceZone: "Tailoring BoE", SourceDrop: "", Stats: Stats{12, 0, 0, 0, 30, 0, 0}, GemSlots: []GemColor{0x4}, SocketBonus: Stats{StatSpellCrit: 2}},
	{ID: 27462, Slot: EquipWrist, Name: "Crimson Bracers of Gloom", Phase: 1, Quality: ItemQualityRare, SourceZone: "H Ramps - Omor the Unscarred", SourceDrop: "", Stats: Stats{18, 18, 0, 12, 22, 0, 0}},
	{ID: 29240, Slot: EquipWrist, Name: "Bands of Negation", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H MT - Nexus- Prince Shaffar", SourceDrop: "", Stats: Stats{22, 25, 0, 0, 29, 0, 0}},
	{ID: 27746, Slot: EquipWrist, Name: "Arcanium Signet Bands", Phase: 1, Quality: ItemQualityRare, SourceZone: "H UB - Hungarfen", SourceDrop: "", Stats: Stats{15, 14, 0, 0, 30, 0, 0}},
	{ID: 29243, Slot: EquipWrist, Name: "Wave-Fury Vambraces", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H SV - Warlod Kalithresh", SourceDrop: "", Stats: Stats{18, 19, 0, 0, 22, 0, 5}},
	{ID: 29955, Slot: EquipWrist, Name: "Mana Infused Wristguards", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "A Fate Worse Than Death - Netherstorm Quest", SourceDrop: "", Stats: Stats{8, 12, 0, 0, 25, 0, 0}},
	{ID: 27465, Slot: EquipHands, Name: "Mana-Etched Gloves", Phase: 1, Quality: ItemQualityRare, SourceZone: "H Ramps - Omor the Unscarred", SourceDrop: "", Stats: Stats{17, 25, 16, 0, 20, 0, 0}, GemSlots: []GemColor{0x2, 0x4}, SocketBonus: Stats{}},
	{ID: 27793, Slot: EquipHands, Name: "Earth Mantle Handwraps", Phase: 1, Quality: ItemQualityRare, SourceZone: "SV - Mekgineer Steamrigger", SourceDrop: "", Stats: Stats{18, 21, 16, 0, 19, 0, 0}, GemSlots: []GemColor{0x2, 0x4}, SocketBonus: Stats{StatInt: 3}},
	{ID: 31149, Slot: EquipHands, Name: "Gloves of Pandemonium", Phase: 1, Quality: ItemQualityRare, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{15, 0, 22, 10, 25, 0, 0}},
	{ID: 27470, Slot: EquipHands, Name: "Gladiator's Mail Gauntlets", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Arena Season 1 Reward", SourceDrop: "", Stats: Stats{18, 36, 21, 0, 32, 0, 0}},
	{ID: 31280, Slot: EquipHands, Name: "Thundercaller's Gauntlets", Phase: 1, Quality: ItemQualityRare, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{16, 16, 18, 0, 35, 0, 0}},
	{ID: 30924, Slot: EquipHands, Name: "Gloves of the High Magus", Phase: 1, Quality: ItemQualityRare, SourceZone: "News of Victory - SMV Quest", SourceDrop: "", Stats: Stats{18, 13, 22, 0, 26, 0, 0}},
	{ID: 29317, Slot: EquipHands, Name: "Tempest's Touch", Phase: 1, Quality: ItemQualityRare, SourceZone: "Return to Andormu - CoT Quest", SourceDrop: "", Stats: Stats{20, 10, 0, 0, 27, 0, 0}, GemSlots: []GemColor{0x3, 0x3}, SocketBonus: Stats{}},
	{ID: 27493, Slot: EquipHands, Name: "Gloves of the Deadwatcher", Phase: 1, Quality: ItemQualityRare, SourceZone: "H AC - Shirrak the Dead Watcher", SourceDrop: "", Stats: Stats{24, 24, 0, 18, 29, 0, 0}},
	{ID: 27508, Slot: EquipHands, Name: "Incanter's Gloves", Phase: 1, Quality: ItemQualityRare, SourceZone: "SV - Thespia", SourceDrop: "", Stats: Stats{24, 21, 14, 0, 29, 0, 0}},
	{ID: 24452, Slot: EquipHands, Name: "Starlight Gauntlets", Phase: 1, Quality: ItemQualityRare, SourceZone: "N UB - Hungarfen", SourceDrop: "", Stats: Stats{21, 10, 0, 0, 25, 0, 0}, GemSlots: []GemColor{0x3, 0x3}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 27537, Slot: EquipHands, Name: "Gloves of Oblivion", Phase: 1, Quality: ItemQualityRare, SourceZone: "SH - Kargath", SourceDrop: "", Stats: Stats{21, 33, 0, 20, 26, 0, 0}},
	{ID: 29784, Slot: EquipHands, Name: "Harmony's Touch", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "Building a Perimeter - Netherstorm Quest", SourceDrop: "", Stats: Stats{0, 18, 16, 0, 33, 0, 0}},
	{ID: 27743, Slot: EquipWaist, Name: "Girdle of Living Flame", Phase: 1, Quality: ItemQualityRare, SourceZone: "H UB - Hungarfen", SourceDrop: "", Stats: Stats{17, 15, 0, 16, 29, 0, 0}, GemSlots: []GemColor{0x4, 0x3}, SocketBonus: Stats{StatSpellCrit: 3}},
	{ID: 29244, Slot: EquipWaist, Name: "Wave-Song Girdle", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H AC - Exarch Maladaar", SourceDrop: "", Stats: Stats{25, 25, 23, 0, 32, 0, 0}},
	{ID: 31461, Slot: EquipWaist, Name: "A'dal's Gift", Phase: 1, Quality: ItemQualityRare, SourceZone: "How to Break Into the Arcatraz - Quest", SourceDrop: "", Stats: Stats{25, 0, 21, 0, 34, 0, 0}},
	{ID: 29257, Slot: EquipWaist, Name: "Sash of Arcane Visions", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H AC - Exarch Maladaar", SourceDrop: "", Stats: Stats{23, 18, 22, 0, 28, 0, 0}},
	{ID: 29241, Slot: EquipWaist, Name: "Belt of Depravity", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H Arc - Harbinger Skyriss", SourceDrop: "", Stats: Stats{27, 31, 0, 17, 34, 0, 0}},
	{ID: 27783, Slot: EquipWaist, Name: "Moonrage Girdle", Phase: 1, Quality: ItemQualityRare, SourceZone: "SV - Hydromancer Thespia", SourceDrop: "", Stats: Stats{22, 0, 20, 0, 25, 0, 0}},
	{ID: 27795, Slot: EquipWaist, Name: "Sash of Serpentra", Phase: 1, Quality: ItemQualityRare, SourceZone: "SV - Warlord Kalithresh", SourceDrop: "", Stats: Stats{21, 31, 0, 17, 25, 0, 0}},
	{ID: 31513, Slot: EquipWaist, Name: "Blackwhelp Belt", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "Whelps of the Wyrmcult - BEM Quest", SourceDrop: "", Stats: Stats{11, 0, 10, 0, 32, 0, 0}},
	{ID: 24262, Slot: EquipLegs, Name: "Spellstrike Pants", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Tailoring BoE", SourceDrop: "", Stats: Stats{8, 12, 26, 22, 46, 0, 0}, GemSlots: []GemColor{0x2, 0x4, 0x3}, SocketBonus: Stats{StatStm: 6}},
	{ID: 30541, Slot: EquipLegs, Name: "Stormsong Kilt", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H UB - The Black Stalker", SourceDrop: "", Stats: Stats{30, 25, 26, 0, 35, 0, 0}, GemSlots: []GemColor{0x2, 0x4, 0x3}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 29141, Slot: EquipLegs, Name: "Tempest Leggings", Phase: 1, Quality: ItemQualityRare, SourceZone: "The Mag'har - Revered (Horde)", SourceDrop: "", Stats: Stats{11, 0, 18, 0, 44, 0, 0}, GemSlots: []GemColor{0x2, 0x4, 0x4}, SocketBonus: Stats{StatMP5: 2}},
	{ID: 29142, Slot: EquipLegs, Name: "Kurenai Kilt", Phase: 1, Quality: ItemQualityRare, SourceZone: "Kurenai - Revered (Ally)", SourceDrop: "", Stats: Stats{11, 0, 18, 0, 44, 0, 0}, GemSlots: []GemColor{0x2, 0x4, 0x4}, SocketBonus: Stats{StatMP5: 2}},
	{ID: 30531, Slot: EquipLegs, Name: "Breeches of the Occultist", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H BM - Aeonus", SourceDrop: "", Stats: Stats{22, 37, 23, 0, 26, 0, 0}, GemSlots: []GemColor{0x4, 0x4, 0x3}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 30709, Slot: EquipLegs, Name: "Pantaloons of Flaming Wrath", Phase: 1, Quality: ItemQualityRare, SourceZone: "H SH - Blood Guard Porung", SourceDrop: "", Stats: Stats{28, 0, 42, 0, 33, 0, 0}},
	{ID: 27492, Slot: EquipLegs, Name: "Moonchild Leggings", Phase: 1, Quality: ItemQualityRare, SourceZone: "H BF - Broggok", SourceDrop: "", Stats: Stats{20, 26, 21, 0, 23, 0, 0}, GemSlots: []GemColor{0x2, 0x4, 0x4}, SocketBonus: Stats{StatMP5: 2}},
	{ID: 29343, Slot: EquipLegs, Name: "Haramad's Leggings of the Third Coin", Phase: 1, Quality: ItemQualityRare, SourceZone: "Undercutting the Competition - MT Quest", SourceDrop: "", Stats: Stats{29, 0, 16, 0, 27, 0, 0}, GemSlots: []GemColor{0x2, 0x4, 0x4}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 27472, Slot: EquipLegs, Name: "Gladiator's Mail Leggings", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Arena Season 1 Reward", SourceDrop: "", Stats: Stats{25, 54, 22, 0, 42, 0, 6}},
	{ID: 30532, Slot: EquipLegs, Name: "Kirin Tor Master's Trousers", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H SLabs - Murmur", SourceDrop: "", Stats: Stats{29, 27, 0, 0, 36, 0, 0}, GemSlots: []GemColor{0x2, 0x4, 0x3}, SocketBonus: Stats{StatSpellHit: 4}},
	{ID: 28185, Slot: EquipLegs, Name: "Khadgar's Kilt of Abjuration", Phase: 1, Quality: ItemQualityRare, SourceZone: "BM - Temporus", SourceDrop: "", Stats: Stats{22, 20, 0, 0, 36, 0, 0}, GemSlots: []GemColor{0x4, 0x3, 0x3}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 27838, Slot: EquipLegs, Name: "Incanter's Trousers", Phase: 1, Quality: ItemQualityRare, SourceZone: "SH - Talon King Ikiss", SourceDrop: "", Stats: Stats{30, 25, 18, 0, 42, 0, 0}},
	{ID: 27907, Slot: EquipLegs, Name: "Mana-Etched Pantaloons", Phase: 1, Quality: ItemQualityRare, SourceZone: "H UB - The Black Stalker", SourceDrop: "", Stats: Stats{32, 34, 21, 0, 33, 0, 0}},
	{ID: 27909, Slot: EquipLegs, Name: "Tidefury Kilt", Phase: 1, Quality: ItemQualityRare, SourceZone: "SLabs - Murmur", SourceDrop: "", Stats: Stats{31, 39, 19, 0, 35, 0, 0}},
	{ID: 28266, Slot: EquipLegs, Name: "Molten Earth Kilt", Phase: 1, Quality: ItemQualityRare, SourceZone: "Mech - Pathaleon the Calculator", SourceDrop: "", Stats: Stats{32, 24, 0, 0, 40, 0, 10}},
	{ID: 27948, Slot: EquipLegs, Name: "Trousers of Oblivion", Phase: 1, Quality: ItemQualityRare, SourceZone: "SH - Talon King Ikiss", SourceDrop: "", Stats: Stats{33, 42, 0, 12, 39, 0, 0}},
	{ID: 29314, Slot: EquipLegs, Name: "Leggings of the Third Coin", Phase: 1, Quality: ItemQualityRare, SourceZone: "Levixus the Soul Caller - Auchindoun Quest", SourceDrop: "", Stats: Stats{26, 34, 12, 0, 32, 0, 4}},
	{ID: 28406, Slot: EquipFeet, Name: "Sigil-Laced Boots", Phase: 1, Quality: ItemQualityRare, SourceZone: "Arc - Harbinger Skyriss", SourceDrop: "", Stats: Stats{18, 24, 17, 0, 20, 0, 0}, GemSlots: []GemColor{0x2, 0x4}, SocketBonus: Stats{StatInt: 3}},
	{ID: 28640, Slot: EquipFeet, Name: "General's Mail Sabatons", Phase: 1, Quality: ItemQualityEpic, SourceZone: "11,424 Honor & 40 EotS Marks", SourceDrop: "", Stats: Stats{23, 34, 24, 0, 28, 0, 0}},
	{ID: 27914, Slot: EquipFeet, Name: "Moonstrider Boots", Phase: 1, Quality: ItemQualityRare, SourceZone: "SH - Darkweaver Syth", SourceDrop: "", Stats: Stats{22, 21, 20, 0, 25, 0, 6}},
	{ID: 28179, Slot: EquipFeet, Name: "Shattrath Jumpers", Phase: 1, Quality: ItemQualityRare, SourceZone: "Into the Heart of the Labyrinth - Auch. Quest", SourceDrop: "", Stats: Stats{17, 25, 0, 0, 29, 0, 0}, GemSlots: []GemColor{0x4, 0x3}, SocketBonus: Stats{StatInt: 3}},
	{ID: 29245, Slot: EquipFeet, Name: "Wave-Crest Striders", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H BF - Keli'dan the Breaker", SourceDrop: "", Stats: Stats{26, 28, 0, 0, 33, 0, 8}},
	{ID: 27821, Slot: EquipFeet, Name: "Extravagant Boots of Malice", Phase: 1, Quality: ItemQualityRare, SourceZone: "H MT - Tavarok", SourceDrop: "", Stats: Stats{24, 27, 0, 14, 30, 0, 0}},
	{ID: 27845, Slot: EquipFeet, Name: "Magma Plume Boots", Phase: 1, Quality: ItemQualityRare, SourceZone: "H AC - Shirrak the Dead Watcher", SourceDrop: "", Stats: Stats{26, 24, 0, 14, 29, 0, 0}},
	{ID: 29808, Slot: EquipFeet, Name: "Shimmering Azure Boots", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "Securing the Celestial Ridge - NS Quest", SourceDrop: "", Stats: Stats{19, 0, 0, 16, 23, 0, 5}},
	{ID: 29242, Slot: EquipFeet, Name: "Boots of Blasphemy", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H SP - Quagmirran", SourceDrop: "", Stats: Stats{29, 36, 0, 0, 36, 0, 0}},
	{ID: 29258, Slot: EquipFeet, Name: "Boots of Ethereal Manipulation", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H Bot - Warp Splinter", SourceDrop: "", Stats: Stats{27, 27, 0, 0, 33, 0, 0}},
	{ID: 29313, Slot: EquipFeet, Name: "Earthbreaker's Greaves", Phase: 1, Quality: ItemQualityRare, SourceZone: "Levixus the Soul Caller - Auchindoun Quest", SourceDrop: "", Stats: Stats{20, 27, 8, 0, 25, 0, 3}},
	{ID: 30519, Slot: EquipFeet, Name: "Boots of the Nexus Warden", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "The Flesh Lies... - Netherstorm Quest", SourceDrop: "", Stats: Stats{17, 27, 0, 18, 21, 0, 0}},
	{ID: 28227, Slot: EquipFinger, Name: "Sparking Arcanite Ring", Phase: 1, Quality: ItemQualityRare, SourceZone: "H OHF - Epoch Hunter", SourceDrop: "", Stats: Stats{14, 13, 14, 10, 22, 0, 0}},
	{ID: 29126, Slot: EquipFinger, Name: "Seer's Signet", Phase: 1, Quality: ItemQualityEpic, SourceZone: "The Scryers - Exalted", SourceDrop: "", Stats: Stats{0, 24, 12, 0, 34, 0, 0}},
	{ID: 31922, Slot: EquipFinger, Name: "Ring of Conflict Survival", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H MT - Yor (Summoned Boss)", SourceDrop: "", Stats: Stats{0, 28, 20, 0, 23, 0, 0}},
	{ID: 28394, Slot: EquipFinger, Name: "Ryngo's Band of Ingenuity", Phase: 1, Quality: ItemQualityRare, SourceZone: "Arc - Wrath-Scryer Soccothrates", SourceDrop: "", Stats: Stats{14, 12, 14, 0, 25, 0, 0}},
	{ID: 29320, Slot: EquipFinger, Name: "Band of the Guardian", Phase: 1, Quality: ItemQualityRare, SourceZone: "Hero of the Brood - CoT Quest", SourceDrop: "", Stats: Stats{11, 0, 17, 0, 23, 0, 0}},
	{ID: 27784, Slot: EquipFinger, Name: "Scintillating Coral Band", Phase: 1, Quality: ItemQualityRare, SourceZone: "SV - Hydromancer Thespia", SourceDrop: "", Stats: Stats{15, 14, 17, 0, 21, 0, 0}},
	{ID: 30366, Slot: EquipFinger, Name: "Manastorm Band", Phase: 1, Quality: ItemQualityRare, SourceZone: "Shutting Down Manaforge Ara - Quest", SourceDrop: "", Stats: Stats{15, 0, 10, 0, 29, 0, 0}},
	{ID: 29172, Slot: EquipFinger, Name: "Ashyen's Gift", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Cenarion Expedition - Exalted", SourceDrop: "", Stats: Stats{0, 30, 0, 21, 23, 0, 0}},
	{ID: 29352, Slot: EquipFinger, Name: "Cobalt Band of Tyrigosa", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H MT - Nexus-Prince Shaffar", SourceDrop: "", Stats: Stats{17, 19, 0, 0, 35, 0, 0}},
	{ID: 28555, Slot: EquipFinger, Name: "Seal of the Exorcist", Phase: 1, Quality: ItemQualityEpic, SourceZone: "50 Spirit Shards ", SourceDrop: "", Stats: Stats{0, 24, 0, 12, 28, 0, 0}},
	{ID: 31339, Slot: EquipFinger, Name: "Lola's Eve", Phase: 1, Quality: ItemQualityEpic, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{14, 15, 0, 0, 29, 0, 0}},
	{ID: 31921, Slot: EquipFinger, Name: "Yor's Collapsing Band", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H MT - Yor (Summoned Boss)", SourceDrop: "", Stats: Stats{20, 0, 0, 0, 23, 0, 0}},
	{ID: 28248, Slot: EquipTotem, Name: "Totem of the Void", Phase: 1, Quality: ItemQualityRare, SourceZone: "Mech - Cache of the Legion", SourceDrop: "", Stats: Stats{StatSpellDmg: 55}},
	{ID: 23199, Slot: EquipTotem, Name: "Totem of the Storm", Phase: 0, Quality: ItemQualityRare, SourceZone: "Boe World Drop", SourceDrop: "", Stats: Stats{StatSpellDmg: 33}},
	{ID: 27543, Slot: EquipWeapon, Name: "Starlight Dagger", Phase: 1, Quality: ItemQualityRare, SourceZone: "H SP - Mennu the Betrayer", SourceDrop: "", Stats: Stats{15, 15, 0, 16, 121, 0, 0}},
	{ID: 27868, Slot: EquipWeapon, Name: "Runesong Dagger", Phase: 1, Quality: ItemQualityRare, SourceZone: "SH - Warbringer O'mrogg", SourceDrop: "", Stats: Stats{11, 12, 20, 0, 121, 0, 0}},
	{ID: 27741, Slot: EquipWeapon, Name: "Bleeding Hollow Warhammer", Phase: 1, Quality: ItemQualityRare, SourceZone: "H SP - Quagmirran", SourceDrop: "", Stats: Stats{17, 12, 16, 0, 121, 0, 0}},
	{ID: 27937, Slot: EquipWeapon, Name: "Sky Breaker", Phase: 1, Quality: ItemQualityRare, SourceZone: "H AC - Avatar of the Martyred", SourceDrop: "", Stats: Stats{20, 13, 0, 0, 132, 0, 0}},
	{ID: 28412, Slot: EquipOffhand, Name: "Lamp of Peaceful Radiance", Phase: 1, Quality: ItemQualityRare, SourceZone: "Arc - Harbinger Skyriss", SourceDrop: "", Stats: Stats{14, 13, 13, 12, 21, 0, 0}},
	{ID: 28260, Slot: EquipOffhand, Name: "Manual of the Nethermancer", Phase: 1, Quality: ItemQualityRare, SourceZone: "Mech - Nethermancer Sepethrea", SourceDrop: "", Stats: Stats{15, 12, 19, 0, 21, 0, 0}},
	{ID: 31287, Slot: EquipOffhand, SubSlot: SubslotShield, Name: "Draenei Honor Guard Shield", Phase: 1, Quality: ItemQualityRare, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{16, 0, 21, 0, 19, 0, 0}},
	{ID: 28187, Slot: EquipOffhand, Name: "Star-Heart Lamp", Phase: 1, Quality: ItemQualityRare, SourceZone: "BM - Temporus", SourceDrop: "", Stats: Stats{18, 17, 0, 12, 22, 0, 0}},
	{ID: 29330, Slot: EquipOffhand, Name: "The Saga of Terokk", Phase: 1, Quality: ItemQualityRare, SourceZone: "Terokk's Legacy - Auchindoun Quest", SourceDrop: "", Stats: Stats{23, 0, 0, 0, 28, 0, 0}},
	{ID: 27910, Slot: EquipOffhand, SubSlot: SubslotShield, Name: "Silvermoon Crest Shield", Phase: 1, Quality: ItemQualityRare, SourceZone: "SLabs - Murmur", SourceDrop: "", Stats: Stats{20, 0, 0, 0, 23, 0, 5}},
	{ID: 30984, Slot: EquipOffhand, SubSlot: SubslotShield, Name: "Spellbreaker's Buckler", Phase: 1, Quality: ItemQualityRare, SourceZone: "Akama's Promise - SMV Quest", SourceDrop: "", Stats: Stats{10, 22, 0, 0, 29, 0, 0}},
	{ID: 27534, Slot: EquipOffhand, Name: "Hortus' Seal of Brilliance", Phase: 1, Quality: ItemQualityRare, SourceZone: "SH - Warchief Kargath Bladefist", SourceDrop: "", Stats: Stats{20, 18, 0, 0, 23, 0, 0}},
	{ID: 29355, Slot: EquipWeapon, SubSlot: SubslotTwoHand, Name: "Terokk's Shadowstaff", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H SH - Talon King Ikiss", SourceDrop: "", Stats: Stats{42, 40, 37, 0, 168, 0, 0}},
	{ID: 29130, Slot: EquipWeapon, SubSlot: SubslotTwoHand, Name: "Auchenai Staff", Phase: 1, Quality: ItemQualityRare, SourceZone: "The Aldor - Revered", SourceDrop: "", Stats: Stats{46, 0, 26, 19, 121, 0, 0}},
	{ID: 28341, Slot: EquipWeapon, SubSlot: SubslotTwoHand, Name: "Warpstaff of Arcanum", Phase: 1, Quality: ItemQualityRare, SourceZone: "Bot - Warp Splinter", SourceDrop: "", Stats: Stats{38, 37, 26, 16, 121, 0, 0}},
	{ID: 31308, Slot: EquipWeapon, SubSlot: SubslotTwoHand, Name: "The Bringer of Death", Phase: 1, Quality: ItemQualityRare, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{31, 32, 42, 0, 121, 0, 0}},
	{ID: 28188, Slot: EquipWeapon, SubSlot: SubslotTwoHand, Name: "Bloodfire Greatstaff", Phase: 1, Quality: ItemQualityRare, SourceZone: "BM - Aeonus", SourceDrop: "", Stats: Stats{42, 42, 28, 0, 121, 0, 0}},
	{ID: 30011, Slot: EquipWeapon, SubSlot: SubslotTwoHand, Name: "Ameer's Impulse Taser", Phase: 1, Quality: ItemQualityRare, SourceZone: "Nexus-King Salhadaar - Netherstorm Quest", SourceDrop: "", Stats: Stats{27, 27, 27, 17, 103, 0, 0}},
	{ID: 27842, Slot: EquipWeapon, SubSlot: SubslotTwoHand, Name: "Grand Scepter of the Nexus-Kings", Phase: 1, Quality: ItemQualityRare, SourceZone: "H MT - Nexus-Prince Shaffar", SourceDrop: "", Stats: Stats{43, 45, 0, 19, 121, 0, 0}},

	{ID: 28346, Slot: EquipOffhand, Name: "Gladiator's Endgame", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Arena Season 1 Reward", SourceDrop: "", Stats: Stats{14, 21, 0, 0, 19, 0, 0}},
	{ID: 24557, Slot: EquipWeapon, SubSlot: SubslotTwoHand, Name: "Gladiator's War Staff", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Arena Season 1 Reward", SourceDrop: "", Stats: Stats{35, 48, 36, 21, 199, 0, 0}},

	{ID: 29389, Slot: EquipTotem, Name: "Totem of the Pulsing Earth", Phase: 1, Quality: ItemQualityEpic, SourceZone: "15 Badge of Justice - G'eras", SourceDrop: "", Stats: Stats{0, 0, 0, 0, 0, 0, 0}, Activate: ActivateTotemOfPulsingEarth, ActivateCD: neverExpires},
	// {Slot: EquipTotem, Name: "Totem of Impact", Phase: 1, Quality: ItemQualityRare, SourceZone: "15 Mark of Thrallmar/ Honor Hold", SourceDrop: "", Stats: Stats{0, 0, 0, 0, 0, 0, 0}},
	// {Slot: EquipTotem, Name: "Totem of Lightning", Phase: 1, Quality: ItemQualityRare, SourceZone: "Colossal Menace - HFP Quest", SourceDrop: "", Stats: Stats{0, 0, 0, 0, 0, 0, 0}},

	// source: https://docs.google.com/spreadsheets/d/1T4DEuq0yroEPb-11okC3qjj7aYfCGu2e6nT9LeT30zg/edit#gid=0
	{ID: 28744, Slot: EquipHead, Name: "Uni-Mind Headdress", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Netherspite", Stats: Stats{StatStm: 31, StatInt: 40, StatSpellDmg: 46, StatSpellCrit: 25, StatSpellHit: 19}},
	{ID: 28586, Slot: EquipHead, Name: "Wicked Witch's Hat", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Opera", Stats: Stats{StatStm: 37, StatInt: 38, StatSpellDmg: 43, StatSpellCrit: 32}},
	{ID: 29035, Slot: EquipHead, Name: "Cyclone Faceguard (Tier 4)", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Prince", Stats: Stats{StatStm: 30, StatInt: 31, StatSpellDmg: 39, StatSpellCrit: 25, StatMP5: 8}, GemSlots: []GemColor{GemColorMeta, GemColorYellow}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 30171, Slot: EquipHead, Name: "Cataclysm Headpiece (Tier 5)", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "Lady Vashj", Stats: Stats{StatStm: 35, StatInt: 28, StatSpellDmg: 54, StatSpellCrit: 26, StatSpellHit: 18, StatMP5: 7}, GemSlots: []GemColor{GemColorMeta, GemColorYellow}, SocketBonus: Stats{StatSpellHit: 5}},
	{ID: 29986, Slot: EquipHead, Name: "Cowl of the Grand Engineer", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Void Reaver", Stats: Stats{StatStm: 22, StatInt: 27, StatSpellDmg: 53, StatSpellCrit: 35, StatSpellHit: 16}, GemSlots: []GemColor{GemColorYellow, GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 32476, Slot: EquipHead, Name: "Gadgetstorm Goggles", Phase: 2, Quality: ItemQualityEpic, SourceZone: "Crafted (Patch 2.1)", SourceDrop: "Engineering (Mail)", Stats: Stats{StatStm: 28, StatInt: 0, StatSpellDmg: 55, StatSpellCrit: 40, StatSpellHit: 12}, GemSlots: []GemColor{GemColorMeta, GemColorBlue}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 31014, Slot: EquipHead, Name: "Skyshatter Headguard (Tier 6)", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Archimonde", Stats: Stats{StatStm: 42, StatInt: 37, StatSpellDmg: 62, StatSpellCrit: 36, StatMP5: 8}, GemSlots: []GemColor{GemColorMeta, GemColorBlue}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 32525, Slot: EquipHead, Name: "Cowl of the Illidari High Lord", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Illidan", Stats: Stats{StatStm: 33, StatInt: 31, StatSpellDmg: 64, StatSpellCrit: 47, StatSpellHit: 21}, GemSlots: []GemColor{GemColorMeta, GemColorBlue}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 28530, Slot: EquipNeck, Name: "Brooch of Unquenchable Fury", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Moroes", Stats: Stats{StatStm: 24, StatInt: 21, StatSpellDmg: 26, StatSpellHit: 15}},
	{ID: 29368, Slot: EquipNeck, Name: "Manasurge Pendant", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Shattrah", SourceDrop: "Badges", Stats: Stats{StatStm: 24, StatInt: 22, StatSpellDmg: 28}},
	{ID: 30008, Slot: EquipNeck, Name: "Pendant of the Lost Ages", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "Tidewalker", Stats: Stats{StatStm: 27, StatInt: 17, StatSpellDmg: 36}},
	{ID: 28762, Slot: EquipNeck, Name: "Adornment of Stolen Souls", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Prince", Stats: Stats{StatStm: 18, StatInt: 20, StatSpellDmg: 28, StatSpellCrit: 23}},
	{ID: 30015, Slot: EquipNeck, Name: "The Sun King's Talisman", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Kael Reward", Stats: Stats{StatStm: 22, StatInt: 16, StatSpellDmg: 41, StatSpellCrit: 24}},
	{ID: 32349, Slot: EquipNeck, Name: "Translucent Spellthread Necklace", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "RoS", Stats: Stats{StatStm: 0, StatInt: 0, StatSpellDmg: 46, StatSpellCrit: 24, StatSpellHit: 15}},
	{ID: 28726, Slot: EquipShoulder, Name: "Mantle of the Mind Flayer", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Aran", Stats: Stats{StatStm: 33, StatInt: 29, StatSpellDmg: 35}},
	{ID: 30024, Slot: EquipShoulder, Name: "Mantle of the Elven Kings", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Trash", Stats: Stats{StatStm: 27, StatInt: 18, StatSpellDmg: 39, StatSpellCrit: 25, StatSpellHit: 18}},
	{ID: 29037, Slot: EquipShoulder, Name: "Cyclone Shoulderguards (Tier 4)", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Gruul's Lair", SourceDrop: "Maulgar", Stats: Stats{StatStm: 28, StatInt: 26, StatSpellDmg: 36, StatSpellCrit: 12}, GemSlots: []GemColor{GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatSpellDmg: 4}},
	{ID: 30079, Slot: EquipShoulder, Name: "Illidari Shoulderpads", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "Tidewalker", Stats: Stats{StatStm: 34, StatInt: 23, StatSpellDmg: 39, StatSpellCrit: 16}, GemSlots: []GemColor{GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatSpellDmg: 4}},
	{ID: 32338, Slot: EquipShoulder, Name: "Blood-cursed Shoulderpads", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Bloodboil", Stats: Stats{StatStm: 25, StatInt: 19, StatSpellDmg: 55, StatSpellCrit: 25, StatSpellHit: 18}},
	{ID: 30173, Slot: EquipShoulder, Name: "Cataclysm Shoulderpads (Tier 5)", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "VoidReaver", Stats: Stats{StatStm: 26, StatInt: 19, StatSpellDmg: 41, StatSpellCrit: 24, StatMP5: 6}, GemSlots: []GemColor{GemColorBlue, GemColorYellow}, SocketBonus: Stats{StatSpellCrit: 3}},
	{ID: 32587, Slot: EquipShoulder, Name: "Mantle of Nimble Thought", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Tailoring", Stats: Stats{StatStm: 37, StatInt: 26, StatSpellDmg: 44, StatHaste: 38}},
	{ID: 31023, Slot: EquipShoulder, Name: "Skyshatter Mantle (Tier 6)", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Mother", Stats: Stats{StatStm: 30, StatInt: 31, StatSpellDmg: 46, StatSpellCrit: 27, StatSpellHit: 11, StatMP5: 4}, GemSlots: []GemColor{GemColorBlue, GemColorYellow}, SocketBonus: Stats{StatSpellDmg: 4}},
	{ID: 30884, Slot: EquipShoulder, Name: "Hatefury Mantle", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Anetheron", Stats: Stats{StatStm: 15, StatInt: 18, StatSpellDmg: 55, StatSpellCrit: 24}, GemSlots: []GemColor{GemColorBlue, GemColorYellow}, SocketBonus: Stats{StatSpellCrit: 3}},
	{ID: 28766, Slot: EquipBack, Name: "Ruby Drape of the Mysticant", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Prince", Stats: Stats{StatStm: 22, StatInt: 21, StatSpellDmg: 30, StatSpellHit: 18}},
	{ID: 28570, Slot: EquipBack, Name: "Shadow-Cloak of Dalaran", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Moroes", Stats: Stats{StatStm: 19, StatInt: 18, StatSpellDmg: 36}},
	{ID: 29369, Slot: EquipBack, Name: "Shawl of Shifting Probabilities", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Shattrah", SourceDrop: "Badges", Stats: Stats{StatStm: 18, StatInt: 16, StatSpellDmg: 21, StatSpellCrit: 22}},
	{ID: 29992, Slot: EquipBack, Name: "Royal Cloak of the Sunstriders", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Kaelthas", Stats: Stats{StatStm: 27, StatInt: 22, StatSpellDmg: 44}},
	{ID: 28797, Slot: EquipBack, Name: "Brute Cloak of the Ogre-Magi", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Gruul's Lair", SourceDrop: "Maulgar", Stats: Stats{StatStm: 18, StatInt: 20, StatSpellDmg: 28, StatSpellCrit: 23}},
	{ID: 30735, Slot: EquipBack, Name: "Ancient Spellcloak of the Highborne", Phase: 1, Quality: ItemQualityEpic, SourceZone: "WorldBoss", SourceDrop: "Kazzak", Stats: Stats{StatStm: 0, StatInt: 15, StatSpellDmg: 36, StatSpellCrit: 19}},
	{ID: 32331, Slot: EquipBack, Name: "Cloak of the Illidari Council", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Illidari Council", Stats: Stats{StatStm: 24, StatInt: 16, StatSpellDmg: 42, StatSpellCrit: 25}},
	{ID: 29033, Slot: EquipChest, Name: "Cyclone Chestguard (Tier 4)", Phase: 1, Quality: ItemQualityEpic, SourceZone: "GruulsLair", SourceDrop: "Maulgar", Stats: Stats{StatStm: 33, StatInt: 32, StatSpellDmg: 39, StatSpellCrit: 20, StatMP5: 8}, GemSlots: []GemColor{GemColorRed, GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellHit: 4}},
	{ID: 29519, Slot: EquipChest, Name: "Netherstrike Breastplate", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Crafted", SourceDrop: "Leatherworking", Stats: Stats{StatStm: 34, StatInt: 23, StatSpellDmg: 37, StatSpellCrit: 32, StatMP5: 8}, GemSlots: []GemColor{GemColorBlue, GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 30056, Slot: EquipChest, Name: "Robe of Hateful Echoes", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "Hydross", Stats: Stats{StatStm: 34, StatInt: 36, StatSpellDmg: 50, StatSpellCrit: 25}, GemSlots: []GemColor{GemColorRed, GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatStm: 6}},
	{ID: 32327, Slot: EquipChest, Name: "Robe of the Shadow Council", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Teron", Stats: Stats{StatStm: 37, StatInt: 36, StatSpellDmg: 73, StatSpellCrit: 28}},
	{ID: 30913, Slot: EquipChest, Name: "Robes of Rhonin", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Archimonde", Stats: Stats{StatStm: 55, StatInt: 38, StatSpellDmg: 81, StatSpellCrit: 24, StatSpellHit: 27}},
	{ID: 30169, Slot: EquipChest, Name: "Cataclysm Chestpiece (Tier 5)", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Kaelthas", Stats: Stats{StatStm: 37, StatInt: 28, StatSpellDmg: 55, StatSpellCrit: 24, StatMP5: 10}, GemSlots: []GemColor{GemColorBlue, GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 30107, Slot: EquipChest, Name: "Vestments of the Sea-Witch", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "LadyVashj", Stats: Stats{StatStm: 28, StatInt: 28, StatSpellDmg: 57, StatSpellCrit: 31, StatSpellHit: 27}, GemSlots: []GemColor{GemColorYellow, GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 32592, Slot: EquipChest, Name: "Chestguard of Relentless Storms", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Trash", Stats: Stats{StatStm: 36, StatInt: 30, StatSpellDmg: 74, StatSpellCrit: 46}},
	{ID: 31017, Slot: EquipChest, Name: "Skyshatter Breastplate (Tier 6)", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Illidan", Stats: Stats{StatStm: 42, StatInt: 41, StatSpellDmg: 62, StatSpellCrit: 27, StatSpellHit: 17, StatMP5: 7}, GemSlots: []GemColor{GemColorBlue, GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 28515, Slot: EquipWrist, Name: "Bands of Nefarious Deeds", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Maiden", Stats: Stats{StatStm: 27, StatInt: 22, StatSpellDmg: 32}},
	{ID: 32351, Slot: EquipWrist, Name: "Elunite Empowered Bracers", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "RoS", Stats: Stats{StatStm: 27, StatInt: 22, StatSpellDmg: 34, StatSpellHit: 19, StatMP5: 6}},
	{ID: 32270, Slot: EquipWrist, Name: "Focused Mana Bindings", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Akama", Stats: Stats{StatStm: 27, StatInt: 20, StatSpellDmg: 42, StatSpellHit: 19}},
	{ID: 29521, Slot: EquipWrist, Name: "Netherstrike Bracers", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Crafted", SourceDrop: "Leatherworking", Stats: Stats{StatStm: 13, StatInt: 13, StatSpellDmg: 20, StatSpellCrit: 17, StatMP5: 6}, GemSlots: []GemColor{GemColorYellow}, SocketBonus: Stats{StatSpellDmg: 2}},
	{ID: 32259, Slot: EquipWrist, Name: "Bands of the Coming Storm", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Supremus", Stats: Stats{StatStm: 28, StatInt: 28, StatSpellDmg: 34, StatSpellCrit: 21}},
	{ID: 29918, Slot: EquipWrist, Name: "Mindstorm Wristbands", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Alar", Stats: Stats{StatStm: 13, StatInt: 13, StatSpellDmg: 36, StatSpellCrit: 23}},
	{ID: 30870, Slot: EquipWrist, Name: "Cuffs of Devastation", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Winterchill", Stats: Stats{StatStm: 22, StatInt: 20, StatSpellDmg: 34, StatSpellCrit: 14}, GemSlots: []GemColor{GemColorYellow}, SocketBonus: Stats{StatStm: 3}},
	{ID: 29034, Slot: EquipHands, Name: "Cyclone Handguards (Tier 4)", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Curator", Stats: Stats{StatStm: 26, StatInt: 29, StatSpellDmg: 34, StatSpellHit: 19, StatMP5: 6}},
	{ID: 28507, Slot: EquipHands, Name: "Handwraps of Flowing Thought", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Huntsman", Stats: Stats{StatStm: 24, StatInt: 22, StatSpellDmg: 35, StatSpellHit: 14}, GemSlots: []GemColor{GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellHit: 3}},
	{ID: 30170, Slot: EquipHands, Name: "Cataclysm Handgrips (Tier 5)", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "LeotherastheBlind", Stats: Stats{StatStm: 25, StatInt: 27, StatSpellDmg: 41, StatSpellCrit: 19, StatSpellHit: 19, StatMP5: 7}},
	{ID: 29987, Slot: EquipHands, Name: "Gauntlets of the Sun King", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Kaelthas", Stats: Stats{StatStm: 28, StatInt: 29, StatSpellDmg: 42, StatSpellCrit: 28}},
	{ID: 30725, Slot: EquipHands, Name: "Anger-Spark Gloves", Phase: 1, Quality: ItemQualityEpic, SourceZone: "World Boss", SourceDrop: "Doomwalker", Stats: Stats{StatStm: 0, StatInt: 0, StatSpellDmg: 30, StatSpellCrit: 25, StatSpellHit: 20}, GemSlots: []GemColor{GemColorRed, GemColorRed}, SocketBonus: Stats{StatSpellCrit: 3}},
	{ID: 28780, Slot: EquipHands, Name: "Soul-Eater's Handwraps", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Magtheridon's Lair", SourceDrop: "Magtheridon", Stats: Stats{StatStm: 31, StatInt: 24, StatSpellDmg: 36, StatSpellCrit: 21}, GemSlots: []GemColor{GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellDmg: 4}},
	{ID: 31008, Slot: EquipHands, Name: "Skyshatter Gauntlets (Tier 6)", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Azgalor", Stats: Stats{StatStm: 30, StatInt: 31, StatSpellDmg: 46, StatSpellCrit: 26, StatSpellHit: 19}, GemSlots: []GemColor{GemColorYellow}, SocketBonus: Stats{StatSpellDmg: 2}},
	{ID: 28565, Slot: EquipWaist, Name: "Nethershard Girdle", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Moroes", Stats: Stats{StatStm: 22, StatInt: 30, StatSpellDmg: 35}},
	{ID: 28639, Slot: EquipWaist, Name: "General's Mail Girdle", Phase: 1, Quality: ItemQualityEpic, SourceZone: "PvP", SourceDrop: "PvP", Stats: Stats{StatStm: 34, StatInt: 23, StatSpellDmg: 28, StatSpellCrit: 23}},
	{ID: 28654, Slot: EquipWaist, Name: "Malefic Girdle", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Illhoof", Stats: Stats{StatStm: 27, StatInt: 26, StatSpellDmg: 37, StatSpellCrit: 21}},
	{ID: 30044, Slot: EquipWaist, Name: "Monsoon Belt", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC/TK", SourceDrop: "Leatherworking", Stats: Stats{StatStm: 23, StatInt: 24, StatSpellDmg: 39, StatSpellHit: 21}, GemSlots: []GemColor{GemColorBlue, GemColorYellow}, SocketBonus: Stats{StatSpellDmg: 4}},
	{ID: 29520, Slot: EquipWaist, Name: "Netherstrike Belt", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Crafted", SourceDrop: "Leatherworking", Stats: Stats{StatStm: 10, StatInt: 17, StatSpellDmg: 30, StatSpellCrit: 16, StatMP5: 9}, GemSlots: []GemColor{GemColorBlue, GemColorYellow}, SocketBonus: Stats{StatSpellCrit: 3}},
	{ID: 28799, Slot: EquipWaist, Name: "Belt of Divine Inspiration", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Gruul's Lair", SourceDrop: "Maulgar", Stats: Stats{StatStm: 27, StatInt: 26, StatSpellDmg: 43}, GemSlots: []GemColor{GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellDmg: 4}},
	{ID: 30064, Slot: EquipWaist, Name: "Cord of Screaming Terrors", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "Lurker", Stats: Stats{StatStm: 34, StatInt: 15, StatSpellDmg: 50, StatSpellHit: 24}, GemSlots: []GemColor{GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatStm: 4}},
	{ID: 24256, Slot: EquipWaist, Name: "Girdle of Ruination", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Crafted", SourceDrop: "Tailoring", Stats: Stats{StatStm: 18, StatInt: 13, StatSpellDmg: 39, StatSpellCrit: 20}, GemSlots: []GemColor{GemColorRed, GemColorYellow}, SocketBonus: Stats{StatStm: 4}},
	{ID: 30914, Slot: EquipWaist, Name: "Belt of the Crescent Moon", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Kazrogal", Stats: Stats{StatStm: 25, StatInt: 27, StatSpellDmg: 44, StatHaste: 36}},
	{ID: 32256, Slot: EquipWaist, Name: "Waistwrap of Infinity", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Supremus", Stats: Stats{StatStm: 31, StatInt: 22, StatSpellDmg: 56, StatHaste: 32}},
	{ID: 30038, Slot: EquipWaist, Name: "Belt of Blasting", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC/TK", SourceDrop: "Tailoring", Stats: Stats{StatStm: 0, StatInt: 0, StatSpellDmg: 50, StatSpellCrit: 30, StatSpellHit: 23}, GemSlots: []GemColor{GemColorBlue, GemColorYellow}, SocketBonus: Stats{StatSpellDmg: 4}},
	{ID: 30888, Slot: EquipWaist, Name: "Anetheron's Noose", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Anetheron", Stats: Stats{StatStm: 22, StatInt: 23, StatSpellDmg: 55, StatSpellCrit: 24}, GemSlots: []GemColor{GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellDmg: 4}},
	{ID: 32276, Slot: EquipWaist, Name: "Flashfire Girdle", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Akama", Stats: Stats{StatStm: 27, StatInt: 26, StatSpellDmg: 44, StatHaste: 37, StatSpellCrit: 18}},
	{ID: 29036, Slot: EquipLegs, Name: "Cyclone Legguards (Tier 4)", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Gruul's Lair", SourceDrop: "Gruul", Stats: Stats{StatStm: 40, StatInt: 40, StatSpellDmg: 49, StatSpellHit: 20, StatMP5: 8}},
	{ID: 28594, Slot: EquipLegs, Name: "Trial-Fire Trousers", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Opera", Stats: Stats{StatStm: 42, StatInt: 40, StatSpellDmg: 49}, GemSlots: []GemColor{GemColorYellow, GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 29972, Slot: EquipLegs, Name: "Trousers of the Astromancer", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Solarian", Stats: Stats{StatStm: 33, StatInt: 36, StatSpellDmg: 54}, GemSlots: []GemColor{GemColorBlue, GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 30172, Slot: EquipLegs, Name: "Cataclysm Leggings (Tier 5)", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Karathress", Stats: Stats{StatStm: 48, StatInt: 46, StatSpellDmg: 54, StatSpellCrit: 24, StatSpellHit: 14}, GemSlots: []GemColor{GemColorYellow}, SocketBonus: Stats{StatSpellDmg: 2}},
	{ID: 32367, Slot: EquipLegs, Name: "Leggings of Devastation", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Mother", Stats: Stats{StatStm: 40, StatInt: 42, StatSpellDmg: 60, StatSpellHit: 26}, GemSlots: []GemColor{GemColorYellow, GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 31020, Slot: EquipLegs, Name: "Skyshatter Legguards (Tier 6)", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Illidari Council", Stats: Stats{StatStm: 40, StatInt: 42, StatSpellDmg: 62, StatSpellCrit: 29, StatSpellHit: 20, StatMP5: 11}, GemSlots: []GemColor{GemColorYellow}, SocketBonus: Stats{StatSpellDmg: 2}},
	{ID: 30734, Slot: EquipLegs, Name: "Leggings of the Seventh Circle", Phase: 1, Quality: ItemQualityEpic, SourceZone: "World Boss", SourceDrop: "Kazzak", Stats: Stats{StatStm: 0, StatInt: 22, StatSpellDmg: 50, StatSpellCrit: 25, StatSpellHit: 18}, GemSlots: []GemColor{GemColorRed, GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 30916, Slot: EquipLegs, Name: "Leggings of Channeled Elements", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Kazrogal", Stats: Stats{StatStm: 25, StatInt: 28, StatSpellDmg: 59, StatSpellCrit: 34, StatSpellHit: 18}, GemSlots: []GemColor{GemColorYellow, GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 28670, Slot: EquipFeet, Name: "Boots of the Infernal Coven", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Aran", Stats: Stats{StatStm: 27, StatInt: 27, StatSpellDmg: 34}},
	{ID: 28585, Slot: EquipFeet, Name: "Ruby Slippers", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Opera", Stats: Stats{StatStm: 33, StatInt: 29, StatSpellDmg: 35, StatSpellHit: 16}},
	{ID: 28810, Slot: EquipFeet, Name: "Windshear Boots", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Gruul's Lair", SourceDrop: "Gruul", Stats: Stats{StatStm: 37, StatInt: 32, StatSpellDmg: 39, StatSpellHit: 18}},
	{ID: 30894, Slot: EquipFeet, Name: "Blue Suede Shoes", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Kazrogal", Stats: Stats{StatStm: 37, StatInt: 32, StatSpellDmg: 56, StatSpellHit: 18}},
	{ID: 30037, Slot: EquipFeet, Name: "Boots of Blasting", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC/TK", SourceDrop: "Tailoring", Stats: Stats{StatStm: 25, StatInt: 25, StatSpellDmg: 39, StatSpellCrit: 25, StatSpellHit: 18}},
	{ID: 28517, Slot: EquipFeet, Name: "Boots of Foretelling", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Maiden", Stats: Stats{StatStm: 27, StatInt: 23, StatSpellDmg: 26, StatSpellCrit: 19}, GemSlots: []GemColor{GemColorRed, GemColorYellow}, SocketBonus: Stats{StatInt: 3}},
	{ID: 30043, Slot: EquipFeet, Name: "Hurricane Boots", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC/TK", SourceDrop: "Leatherworking", Stats: Stats{StatStm: 25, StatInt: 26, StatSpellDmg: 39, StatSpellCrit: 26, StatMP5: 6}},
	{ID: 30067, Slot: EquipFeet, Name: "Velvet Boots of the Guardian", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "Lurker", Stats: Stats{StatStm: 21, StatInt: 21, StatSpellDmg: 49, StatSpellCrit: 24}},
	{ID: 32242, Slot: EquipFeet, Name: "Boots of Oceanic Fury", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Najentus", Stats: Stats{StatStm: 28, StatInt: 36, StatSpellDmg: 55, StatSpellCrit: 26}},
	{ID: 32352, Slot: EquipFeet, Name: "Naturewarden's Treads", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "RoS", Stats: Stats{StatStm: 39, StatInt: 18, StatSpellDmg: 44, StatSpellCrit: 26, StatMP5: 7}, GemSlots: []GemColor{GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellDmg: 4}},
	{ID: 32239, Slot: EquipFeet, Name: "Slippers of the Seacaller", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Najentus", Stats: Stats{StatStm: 25, StatInt: 18, StatSpellDmg: 44, StatSpellCrit: 29}, GemSlots: []GemColor{GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellDmg: 4}},
	{ID: 28793, Slot: EquipFinger, Name: "Band of Crimson Fury", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Magtheridon's Lair", SourceDrop: "MagtheridonQuest", Stats: Stats{StatStm: 22, StatInt: 22, StatSpellDmg: 28, StatSpellHit: 16}},
	{ID: 28510, Slot: EquipFinger, Name: "Spectral Band of Innervation", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Huntsman", Stats: Stats{StatStm: 22, StatInt: 24, StatSpellDmg: 29}},
	{ID: 29922, Slot: EquipFinger, Name: "Band of Al'ar", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Alar", Stats: Stats{StatStm: 24, StatInt: 23, StatSpellDmg: 37}},
	{ID: 29367, Slot: EquipFinger, Name: "Ring of Cryptic Dreams", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Shattrah", SourceDrop: "Badges", Stats: Stats{StatStm: 16, StatInt: 17, StatSpellDmg: 23, StatSpellCrit: 20}},
	{ID: 29287, Slot: EquipFinger, Name: "Violet Signet of the Archmage", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Exalted", Stats: Stats{StatStm: 24, StatInt: 23, StatSpellDmg: 29, StatSpellCrit: 17}},
	{ID: 29286, Slot: EquipFinger, Name: "Violet Signet (R)", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Revered", Stats: Stats{StatStm: 22, StatInt: 22, StatSpellDmg: 28, StatSpellCrit: 17}},
	{ID: 29285, Slot: EquipFinger, Name: "Violet Signet (H)", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Honored", Stats: Stats{StatStm: 19, StatInt: 21, StatSpellDmg: 26, StatSpellCrit: 15}},
	{ID: 28753, Slot: EquipFinger, Name: "Ring of Recurrence", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Chess", Stats: Stats{StatStm: 15, StatInt: 15, StatSpellDmg: 32, StatSpellCrit: 19}},
	{ID: 29305, Slot: EquipFinger, Name: "Band of the Eternal Sage", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Exalted", Stats: Stats{StatStm: 28, StatInt: 25, StatSpellDmg: 34, StatSpellCrit: 24}},
	{ID: 30109, Slot: EquipFinger, Name: "Ring of Endless Coils", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "LadyVashj", Stats: Stats{StatStm: 31, StatInt: 0, StatSpellDmg: 37, StatSpellCrit: 22}},
	{ID: 30667, Slot: EquipFinger, Name: "Ring of Unrelenting Storms", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Trash", Stats: Stats{StatStm: 0, StatInt: 15, StatSpellDmg: 43, StatSpellCrit: 19}},
	{ID: 32247, Slot: EquipFinger, Name: "Ring of Captured Storms", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Najentus", Stats: Stats{StatStm: 0, StatInt: 0, StatSpellDmg: 42, StatSpellCrit: 29, StatSpellHit: 19}},
	{ID: 32527, Slot: EquipFinger, Name: "Ring of Ancient Knowledge", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Trash", Stats: Stats{StatStm: 30, StatInt: 20, StatSpellDmg: 39, StatHaste: 31}},
	{ID: 30832, Slot: EquipWeapon, Name: "Gavel of Unearthed Secrets", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Shattrah", SourceDrop: "Lower City - Exalted", Stats: Stats{StatStm: 24, StatInt: 16, StatSpellDmg: 159, StatSpellCrit: 15}},
	{ID: 23554, Slot: EquipWeapon, Name: "Eternium Runed Blade", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Crafted", SourceDrop: "Blacksmithing", Stats: Stats{StatStm: 0, StatInt: 19, StatSpellDmg: 168, StatSpellCrit: 21}},
	{ID: 28770, Slot: EquipWeapon, Name: "Nathrezim Mindblade", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Prince", Stats: Stats{StatStm: 18, StatInt: 18, StatSpellDmg: 203, StatSpellCrit: 23}},
	{ID: 30723, Slot: EquipWeapon, Name: "Talon of the Tempest", Phase: 1, Quality: ItemQualityEpic, SourceZone: "World Boss", SourceDrop: "Doomwalker", Stats: Stats{StatStm: 0, StatInt: 10, StatSpellDmg: 194, StatSpellCrit: 19, StatSpellHit: 9}, GemSlots: []GemColor{GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatInt: 3}},
	{ID: 34009, Slot: EquipWeapon, Name: "Hammer of Judgement", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Trash", Stats: Stats{StatStm: 33, StatInt: 22, StatSpellDmg: 236, StatSpellHit: 22}},
	{ID: 32237, Slot: EquipWeapon, Name: "The Maelstrom's Fury", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Najentus", Stats: Stats{StatStm: 33, StatInt: 21, StatSpellDmg: 236, StatSpellCrit: 22}},
	{ID: 28633, Slot: EquipWeapon, SubSlot: SubslotTwoHand, Name: "Staff of Infinite Mysteries", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Curator", Stats: Stats{StatStm: 61, StatInt: 51, StatSpellDmg: 185, StatSpellHit: 23}},
	{ID: 29988, Slot: EquipWeapon, SubSlot: SubslotTwoHand, Name: "The Nexus Key", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Kaelthas", Stats: Stats{StatStm: 76, StatInt: 52, StatSpellDmg: 236, StatSpellCrit: 51}},
	{ID: 32374, Slot: EquipWeapon, SubSlot: SubslotTwoHand, Name: "Zhar'doom, Greatstaff of the Devourer", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Illidan", Stats: Stats{StatStm: 70, StatInt: 47, StatSpellDmg: 259, StatHaste: 55, StatSpellCrit: 36}},
	{ID: 28734, Slot: EquipOffhand, Name: "Jewel of Infinite Possibilities", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Netherspite", Stats: Stats{StatStm: 19, StatInt: 18, StatSpellDmg: 23, StatSpellHit: 21}},
	{ID: 28611, Slot: EquipOffhand, SubSlot: SubslotShield, Name: "Dragonheart Flameshield", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Nightbane", Stats: Stats{StatStm: 19, StatInt: 33, StatSpellDmg: 23, StatMP5: 7}},
	{ID: 34011, Slot: EquipOffhand, SubSlot: SubslotShield, Name: "Illidari Runeshield", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Trash", Stats: Stats{StatStm: 45, StatInt: 39, StatSpellDmg: 34}},
	{ID: 28781, Slot: EquipOffhand, Name: "Karaborian Talisman", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Magtheridon's Lair", SourceDrop: "Magtheridon", Stats: Stats{StatStm: 23, StatInt: 23, StatSpellDmg: 35}},
	{ID: 29268, Slot: EquipOffhand, SubSlot: SubslotShield, Name: "Mazthoril Honor Shield", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Shattrah", SourceDrop: "Badges", Stats: Stats{StatStm: 16, StatInt: 17, StatSpellDmg: 23, StatSpellCrit: 21}},
	{ID: 28603, Slot: EquipOffhand, Name: "Talisman of Nightbane", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Nightbane", Stats: Stats{StatStm: 19, StatInt: 19, StatSpellDmg: 28, StatSpellCrit: 17}},
	{ID: 32361, Slot: EquipOffhand, Name: "Blind-Seers Icon", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Akama", Stats: Stats{StatStm: 25, StatInt: 16, StatSpellDmg: 42, StatSpellHit: 24}},
	{ID: 29273, Slot: EquipOffhand, Name: "Khadgar's Knapsack", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Shattrah", SourceDrop: "Badges", Stats: Stats{StatSpellDmg: 49}},
	{ID: 30049, Slot: EquipOffhand, Name: "FathomStone", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "Lurker", Stats: Stats{StatStm: 16, StatInt: 12, StatSpellDmg: 36, StatSpellCrit: 23}},
	{ID: 30909, Slot: EquipOffhand, SubSlot: SubslotShield, Name: "Antonidas's Aegis of Rapt Concentration", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Archimonde", Stats: Stats{StatStm: 28, StatInt: 32, StatSpellDmg: 42, StatSpellCrit: 20}},
	{ID: 30872, Slot: EquipOffhand, Name: "Chronicle of Dark Secrets", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Winterchill", Stats: Stats{StatStm: 16, StatInt: 12, StatSpellDmg: 42, StatSpellCrit: 23, StatSpellHit: 17}},
	{ID: 28297, Slot: EquipWeapon, Name: "Gladiator's Gavel / Gladiator's Spellblade", Phase: 1, Quality: ItemQualityEpic, SourceZone: "PvP", SourceDrop: "PvP", Stats: Stats{StatStm: 28, StatInt: 18, StatSpellDmg: 199}},

	// Hand Written
	{ID: 27683, Slot: EquipTrinket, Name: "Quagmirran's Eye", Phase: 1, Quality: ItemQualityRare, SourceZone: "The Slave Pens", SourceDrop: "Quagmirran", Stats: Stats{StatSpellDmg: 37}, Activate: ActivateQuagsEye, ActivateCD: neverExpires},
	{ID: 29370, Slot: EquipTrinket, Name: "Icon of the Silver Crescent", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Shattrath", SourceDrop: "G'eras - 41 Badges", Stats: Stats{StatSpellDmg: 43}, Activate: createSpellDmgActivate(MagicIDBlessingSilverCrescent, 155, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDISCTrink},
	{ID: 19344, Slot: EquipTrinket, Name: "Natural Alignment Crystal", Phase: 0, Quality: ItemQualityEpic, SourceZone: "BWL", SourceDrop: "", Stats: Stats{}, Activate: ActivateNAC, ActivateCD: time.Second * 300, CoolID: MagicIDNACTrink},
	{ID: 19379, Slot: EquipTrinket, Name: "Neltharion's Tear", Phase: 0, Quality: ItemQualityEpic, SourceZone: "BWL", SourceDrop: "Nefarian", Stats: Stats{StatSpellDmg: 44, StatSpellHit: 16}},
	{ID: 23046, Slot: EquipTrinket, Name: "The Restrained Essence of Sapphiron", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "Sapphiron", Stats: Stats{StatSpellDmg: 40}, Activate: createSpellDmgActivate(MagicIDSpellPower, 130, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDEssSappTrink},
	{ID: 23207, Slot: EquipTrinket, Name: "Mark of the Champion", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "KT", Stats: Stats{StatSpellDmg: 85}},
	{ID: 29132, Slot: EquipTrinket, Name: "Scryer's Bloodgem", Phase: 1, Quality: ItemQualityRare, SourceZone: "The Scryers - Revered", SourceDrop: "", Stats: Stats{0, 0, 0, 32, 0, 0, 0}, Activate: createSpellDmgActivate(MagicIDSpellPower, 150, time.Second*15), ActivateCD: time.Second * 90, CoolID: MagicIDScryerTrink},
	{ID: 24126, Slot: EquipTrinket, Name: "Figurine - Living Ruby Serpent", Phase: 1, Quality: ItemQualityRare, SourceZone: "Jewelcarfting BoP", SourceDrop: "", Stats: Stats{23, 33, 0, 0, 0, 0, 0}, Activate: createSpellDmgActivate(MagicIDRubySerpent, 150, time.Second*20), ActivateCD: time.Second * 300, CoolID: MagicIDRubySerpentTrink},
	{ID: 29179, Slot: EquipTrinket, Name: "Xi'ri's Gift", Phase: 1, Quality: ItemQualityRare, SourceZone: "The Sha'tar - Revered", SourceDrop: "", Stats: Stats{0, 0, 32, 0, 0, 0, 0}, Activate: createSpellDmgActivate(MagicIDSpellPower, 150, time.Second*15), ActivateCD: time.Second * 90, CoolID: MagicIDXiriTrink},
	{ID: 28418, Slot: EquipTrinket, Name: "Shiffar's Nexus-Horn", Phase: 1, Quality: ItemQualityRare, SourceZone: "Arc - Harbinger Skyriss", SourceDrop: "", Stats: Stats{0, 0, 30, 0, 0, 0, 0}, Activate: ActivateNexusHorn, ActivateCD: neverExpires},
	{ID: 31856, Slot: EquipTrinket, Name: "Darkmoon Card: Crusade", Phase: 2, Quality: ItemQualityEpic, SourceZone: "Blessings Deck", SourceDrop: "", Activate: ActivateDCC, ActivateCD: neverExpires},
	{ID: 28785, Slot: EquipTrinket, Name: "The Lightning Capacitor", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "", Activate: ActivateTLC, ActivateCD: neverExpires},
	{ID: 28789, Slot: EquipTrinket, Name: "Eye of Magtheridon", Phase: 1, Quality: ItemQualityEpic, SourceZone: "", SourceDrop: "", Stats: Stats{StatSpellDmg: 54}, Activate: ActivateEyeOfMag, ActivateCD: neverExpires},
	{ID: 30626, Slot: EquipTrinket, Name: "Sextant of Unstable Currents", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "", Stats: Stats{StatSpellCrit: 40}, Activate: ActivateSextant, ActivateCD: neverExpires},
	{ID: 34429, Slot: EquipTrinket, Name: "Shifting Naaru Sliver", Phase: 5, Quality: ItemQualityEpic, SourceZone: "Sunwell", SourceDrop: "", Stats: Stats{StatHaste: 54}, Activate: createSpellDmgActivate(MagicIDShiftingNaaru, 320, time.Second*15), ActivateCD: time.Second * 90, CoolID: MagicIDShiftingNaaruTrink},
	{ID: 32483, Slot: EquipTrinket, Name: "The Skull of Gul'dan", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Black Temple", SourceDrop: "", Stats: Stats{StatSpellHit: 25, StatSpellDmg: 55}, Activate: createHasteActivate(MagicIDSkullGuldan, 175, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDSkullGuldanTrink},
	{ID: 33829, Slot: EquipTrinket, Name: "Hex Shrunken Head", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "", Stats: Stats{StatSpellDmg: 53}, Activate: createSpellDmgActivate(MagicIDHexShunkHead, 211, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDHexTrink},
	{ID: 29376, Slot: EquipTrinket, Name: "Essence of the Martyr", Phase: 1, Quality: ItemQualityRare, SourceZone: "G'eras", SourceDrop: "Badges", Stats: Stats{StatSpellDmg: 28}, Activate: createSpellDmgActivate(MagicIDSpellPower, 99, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDEssMartyrTrink},
	{ID: 38290, Slot: EquipTrinket, Name: "Dark Iron Smoking Pipe", Phase: 2, Quality: ItemQualityEpic, SourceZone: "Brewfest", SourceDrop: "Corin Direbrew", Stats: Stats{StatSpellDmg: 43}, Activate: createSpellDmgActivate(MagicIDDarkIronPipeweed, 155, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDDITrink},
	{ID: 30663, Slot: EquipTrinket, Name: "Fathom-Brooch of the Tidewalker", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "Fathom-Lord Karathress", Stats: Stats{}, Activate: ActivateFathomBrooch, ActivateCD: neverExpires},
	{ID: 35749, Slot: EquipTrinket, Name: "Sorcerer's Alchemist Stone", Phase: 5, Quality: ItemQualityEpic, SourceZone: "Shattered Sun Offensive", SourceDrop: "Exalted", Stats: Stats{StatSpellDmg: 63}, Activate: ActivateAlchStone, ActivateCD: neverExpires},

	{ID: 24116, Slot: EquipNeck, Name: "Eye of the Night", Phase: 1, Quality: ItemQualityRare, SourceZone: "Jewelcrafting", SourceDrop: "", Stats: Stats{StatSpellCrit: 26, StatSpellHit: 16, StatSpellPen: 15}, Activate: func(sim *Simulation) Aura {
		if sim.Options.Buffs.EyeOfNight {
			return Aura{} // if we already have buff from party member, no need to activate this
		}
		activate := createSpellDmgActivate(MagicIDEyeOfTheNight, 34, time.Minute*30)
		return activate(sim)
	}, ActivateCD: neverExpires, CoolID: MagicIDEyeOfTheNightTrink},
	{ID: 24121, Slot: EquipNeck, Name: "Chain of the Twilight Owl", Phase: 1, Quality: ItemQualityRare, SourceZone: "Jewelcrafting", SourceDrop: "", Stats: Stats{StatStm: 0, StatInt: 19, StatSpellDmg: 21}, Activate: ActivateChainTO, ActivateCD: neverExpires, CoolID: MagicIDChainTOTrink},
	{ID: 31075, Slot: EquipFinger, Name: "Evoker's Mark of the Redemption", Phase: 1, Quality: ItemQualityRare, SourceZone: "Quest SMV", SourceDrop: "Dissension Amongst the Ranks...", Stats: Stats{StatInt: 15, StatSpellDmg: 29, StatSpellCrit: 10}},
	{ID: 32664, Slot: EquipFinger, Name: "Dreamcrystal Band", Phase: 1, Quality: ItemQualityRare, SourceZone: "Blades Edge Moutains", SourceDrop: "50 Apexis Shards", Stats: Stats{StatInt: 10, StatSpellDmg: 38, StatSpellCrit: 15}},
	{ID: 29522, Slot: EquipChest, Name: "Windhawk Hauberk", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Leatherworking", SourceDrop: "", Stats: Stats{StatStm: 28, StatInt: 29, StatSpirit: 29, StatSpellDmg: 46, StatSpellCrit: 19}, GemSlots: []GemColor{GemColorBlue, GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellDmg: 5}},
	{ID: 29524, Slot: EquipWaist, Name: "Windhawk Belt", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Leatherworking", SourceDrop: "", Stats: Stats{StatStm: 17, StatInt: 19, StatSpirit: 20, StatSpellDmg: 37, StatSpellCrit: 12}, GemSlots: []GemColor{GemColorBlue, GemColorYellow}, SocketBonus: Stats{StatSpellDmg: 4}},
	{ID: 29523, Slot: EquipWrist, Name: "Windhawk Bracers", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Leatherworking", SourceDrop: "", Stats: Stats{StatStm: 22, StatInt: 17, StatSpirit: 7, StatSpellDmg: 27, StatSpellCrit: 16}, GemSlots: []GemColor{GemColorYellow}, SocketBonus: Stats{StatInt: 2}},
	{ID: 27510, Slot: EquipHands, Name: "Tidefury Gauntlets", Phase: 1, Quality: ItemQualityRare, SourceZone: "", SourceDrop: "", Stats: Stats{StatStm: 22, StatInt: 26, StatSpellDmg: 29, StatMP5: 7}},
	{ID: 22730, Slot: EquipWaist, Name: "Eyestalk Waist Cord", Phase: 0, Quality: ItemQualityEpic, SourceZone: "AQ40", SourceDrop: "C'thun", Stats: Stats{StatStm: 10, StatInt: 9, StatSpellDmg: 41, StatSpellCrit: 14}},
	{ID: 23070, Slot: EquipLegs, Name: "Leggings of Polarity", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "Thaddius", Stats: Stats{StatStm: 20, StatInt: 14, StatSpellDmg: 44, StatSpellCrit: 28}},
	{ID: 21709, Slot: EquipFinger, Name: "Ring of the Fallen God", Phase: 0, Quality: ItemQualityEpic, SourceZone: "AQ40", SourceDrop: "C'thun", Stats: Stats{StatStm: 5, StatInt: 6, StatSpellDmg: 37, StatSpellHit: 8}},
	{ID: 23031, Slot: EquipFinger, Name: "Band of the Inevitable", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "Noth", Stats: Stats{StatSpellDmg: 36, StatSpellHit: 8}},
	{ID: 23025, Slot: EquipFinger, Name: "Seal of the Damned", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "4H", Stats: Stats{StatStm: 17, StatSpellDmg: 21, StatSpellCrit: 14, StatSpellHit: 8}},
	{ID: 23057, Slot: EquipNeck, Name: "Gem of Trapped Innocents", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "KT", Stats: Stats{StatStm: 9, StatInt: 7, StatSpellDmg: 15, StatSpellCrit: 28}},
	{ID: 21608, Slot: EquipNeck, Name: "Amulet of Vek'nilash", Phase: 0, Quality: ItemQualityEpic, SourceZone: "AQ", SourceDrop: "Twin Emp", Stats: Stats{StatStm: 9, StatInt: 5, StatSpellDmg: 27, StatSpellCrit: 14}},
	{ID: 23664, Slot: EquipShoulder, Name: "Pauldrons of Elemental Fury", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "Trash", Stats: Stats{StatStm: 19, StatInt: 21, StatSpellDmg: 26, StatSpellCrit: 14, StatSpellHit: 8}},
	{ID: 23665, Slot: EquipLegs, Name: "Leggings of Elemental Fury", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "Trash", Stats: Stats{StatStm: 26, StatInt: 27, StatSpellDmg: 32, StatSpellCrit: 28}},
	{ID: 23050, Slot: EquipBack, Name: "Cloak of the Necropolis", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "Sapp", Stats: Stats{StatStm: 12, StatInt: 11, StatSpellDmg: 26, StatSpellCrit: 14, StatSpellHit: 8}},
	{ID: 30682, Slot: EquipFeet, Name: "Glider's Sabatons of Nature's Wrath", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Servants Quarter", Stats: Stats{StatSpellDmg: 78}},
	{ID: 30677, Slot: EquipWaist, Name: "Lurker's Belt of Nature's Wrath", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Servants Quarter", Stats: Stats{StatSpellDmg: 78}},
	{ID: 30686, Slot: EquipWrist, Name: "Ravager's Bands of Nature's Wrath", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Servants Quarter", Stats: Stats{StatSpellDmg: 58}},
	// {Slot: EquipFeet, Name: "Glider's Sabatons of the Invoker", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Servants Quarter", Stats: Stats{StatSpellDmg: 33, StatSpellCrit: 28}},
	// {Slot: EquipWaist, Name: "Lurker's Belt of the Invoker", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Servants Quarter", Stats: Stats{StatSpellDmg: 33, StatSpellCrit: 28}},
	// {Slot: EquipWrist, Name: "Ravager's Bands of the Invoker", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Servants Quarter", Stats: Stats{StatSpellDmg: 25, StatSpellCrit: 21}},
	{ID: 28583, Slot: EquipHead, Name: "Big Bad Wolf's Head", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "The Big Bad Wolf", Stats: Stats{StatStm: 42, StatInt: 40, StatSpellDmg: 47, StatSpellCrit: 28}},
	{ID: 32586, Slot: EquipWrist, Name: "Bracers of Nimble Thought", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Tailoring", Stats: Stats{StatStm: 27, StatInt: 20, StatSpellDmg: 34, StatHaste: 28}},
	{ID: 23049, Slot: EquipOffhand, Name: "Sapphiron's Left Eye", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "Sapphiron", Stats: Stats{StatStm: 12, StatInt: 8, StatSpellDmg: 26, StatSpellCrit: 14, StatSpellHit: 8}},
	{ID: 25778, Slot: EquipWrist, Name: "Manacles of Rememberance", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "Nagrand", SourceDrop: "Quest", Stats: Stats{StatSpirit: 9, StatInt: 10, StatSpellDmg: 16, StatSpellCrit: 14}},
	{ID: 28174, Slot: EquipWrist, Name: "Shattrath Wraps", Phase: 1, Quality: ItemQualityRare, SourceZone: "Auchindoun", SourceDrop: "Quest", Stats: Stats{StatStm: 15, StatInt: 15, StatSpellDmg: 21}, GemSlots: []GemColor{GemColorRed}, SocketBonus: Stats{StatStm: 3}},
	{ID: 31283, Slot: EquipWaist, Name: "Sash of Sealed Fate", Phase: 1, Quality: ItemQualityRare, SourceZone: "World Drop", SourceDrop: "BoE", Stats: Stats{StatInt: 15, StatSpellDmg: 35, StatSpellCrit: 23}},
	{ID: 30004, Slot: EquipFeet, Name: "Landing Boots", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "Netherstorm", SourceDrop: "Quest", Stats: Stats{StatStm: 12, StatInt: 8, StatSpellDmg: 35, StatSpellCrit: 16}},
	{ID: 31290, Slot: EquipFinger, Name: "Band of Dominion", Phase: 1, Quality: ItemQualityRare, SourceZone: "World Drop", SourceDrop: "BoE", Stats: Stats{StatSpellDmg: 28, StatSpellCrit: 21}},

	// Sash of Sealed Fate - blue BoE
	// Landing Boots - green from quest
	// Band of Dominion - Blue BoE

	{ID: 34336, Slot: EquipWeapon, Name: "Sunflare", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "Kil'jaden",
		Stats: Stats{StatStm: 17, StatInt: 20, StatSpellDmg: 292, StatHaste: 23, StatSpellCrit: 30}},
	{ID: 34179, Slot: EquipOffhand, Name: "Heart of the Pit", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "Brutalis",
		Stats: Stats{StatStm: 33, StatInt: 21, StatSpellDmg: 39, StatHaste: 32}},
	{ID: 34350, Slot: EquipHands, Name: "Gauntlets of the Ancient Shadowmoon", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "Trash",
		Stats:       Stats{StatStm: 30, StatInt: 32, StatSpellDmg: 43, StatHaste: 24, StatSpellCrit: 28},
		GemSlots:    []GemColor{GemColorRed, GemColorBlue},
		SocketBonus: Stats{StatSpellCrit: 2},
	},
	{ID: 34542, Slot: EquipWaist, Name: "Skyshatter Cord", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStm: 19, StatInt: 30, StatSpellDmg: 50, StatHaste: 27, StatSpellCrit: 29, StatMP5: 6},
		GemSlots:    []GemColor{GemColorYellow},
		SocketBonus: Stats{StatSpellDmg: 2},
	},
	{ID: 34186, Slot: EquipLegs, Name: "Chain Links of the Tumultuous Storm", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStm: 48, StatInt: 41, StatSpellDmg: 71, StatHaste: 30, StatSpellCrit: 35},
		GemSlots:    []GemColor{GemColorYellow, GemColorRed, GemColorRed},
		SocketBonus: Stats{StatSpellCrit: 4},
	},
	{ID: 34566, Slot: EquipFeet, Name: "Skyshatter Treads", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStm: 21, StatInt: 30, StatSpellDmg: 50, StatHaste: 30, StatSpellCrit: 23, StatMP5: 7},
		GemSlots:    []GemColor{GemColorYellow},
		SocketBonus: Stats{StatSpellDmg: 2},
	},
	{ID: 34437, Slot: EquipWrist, Name: "Skyshatter Bands", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStm: 15, StatInt: 23, StatSpellDmg: 39, StatHaste: 11, StatSpellCrit: 28},
		GemSlots:    []GemColor{GemColorYellow},
		SocketBonus: Stats{StatSpellDmg: 2},
	},
	{ID: 34230, Slot: EquipFinger, Name: "Ring of Omnipotence", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats: Stats{StatStm: 21, StatInt: 14, StatSpellDmg: 40, StatHaste: 31, StatSpellCrit: 22},
	},
	{ID: 34362, Slot: EquipFinger, Name: "Loop of Forged Power", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats: Stats{StatStm: 27, StatInt: 28, StatSpellDmg: 34, StatHaste: 30, StatSpellHit: 19},
	},
	{ID: 34204, Slot: EquipNeck, Name: "Amulet of Unfettered Magics", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats: Stats{StatStm: 24, StatInt: 17, StatSpellDmg: 39, StatHaste: 32, StatSpellHit: 15},
	},
	{ID: 34332, Slot: EquipHead, Name: "Cowl of Gul'dan", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStm: 51, StatInt: 43, StatSpellDmg: 74, StatHaste: 32, StatSpellCrit: 36},
		GemSlots:    []GemColor{GemColorMeta, GemColorYellow},
		SocketBonus: Stats{StatSpellDmg: 5},
	},
	{ID: 34242, Slot: EquipBack, Name: "Tattered Cape of Antonidas", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStm: 25, StatInt: 26, StatSpellDmg: 42, StatHaste: 32},
		GemSlots:    []GemColor{GemColorRed},
		SocketBonus: Stats{StatSpellDmg: 2},
	},
	{ID: 34396, Slot: EquipChest, Name: "Garments of Crashing Shores", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStm: 48, StatInt: 41, StatSpellDmg: 71, StatHaste: 40, StatSpellCrit: 25},
		GemSlots:    []GemColor{GemColorRed, GemColorYellow, GemColorYellow},
		SocketBonus: Stats{StatSpellDmg: 5},
	},
	{ID: 34390, Slot: EquipShoulder, Name: "Erupting Epaulets", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStm: 30, StatInt: 30, StatSpellDmg: 53, StatHaste: 24, StatSpellCrit: 30},
		GemSlots:    []GemColor{GemColorYellow, GemColorRed},
		SocketBonus: Stats{StatSpellDmg: 4},
	},
	{ID: 33970, Slot: EquipShoulder, Name: "Pauldrons of the Furious Elements", Phase: 4, Quality: ItemQualityEpic, SourceZone: "Shattrath - G'eras", SourceDrop: "60 Badges",
		Stats: Stats{StatStm: 28, StatInt: 24, StatSpellDmg: 40, StatHaste: 33},
	},
	{ID: 33965, Slot: EquipChest, Name: "Hauberk of the Furious Elements", Phase: 4, Quality: ItemQualityEpic, SourceZone: "Shattrath - G'eras", SourceDrop: "75 Badges",
		Stats: Stats{StatStm: 39, StatInt: 34, StatSpellDmg: 54, StatHaste: 35, StatSpellCrit: 23},
	},
	{ID: 33588, Slot: EquipWrist, Name: "Runed Spell-Cuffs", Phase: 4, Quality: ItemQualityEpic, SourceZone: "Shattrath - G'eras", SourceDrop: "Badges",
		Stats: Stats{StatStm: 20, StatInt: 18, StatSpellDmg: 29, StatHaste: 25},
	},
	{ID: 33537, Slot: EquipFeet, Name: "Treads of Booming Thunder", Phase: 4, Quality: ItemQualityEpic, SourceZone: "Shattrath - G'eras", SourceDrop: "Badges",
		Stats:       Stats{StatStm: 21, StatInt: 33, StatSpellDmg: 40, StatSpellCrit: 14},
		GemSlots:    []GemColor{GemColorRed, GemColorYellow},
		SocketBonus: Stats{StatSpellCrit: 3},
	},
	{ID: 33534, Slot: EquipHands, Name: "Grips of Nature's Wrath", Phase: 4, Quality: ItemQualityEpic, SourceZone: "Shattrath - G'eras", SourceDrop: "Badges",
		Stats:       Stats{StatStm: 30, StatInt: 27, StatSpellDmg: 34, StatSpellCrit: 21},
		GemSlots:    []GemColor{GemColorRed, GemColorYellow},
		SocketBonus: Stats{StatSpellDmg: 4},
	},
	{ID: 34359, Slot: EquipNeck, Name: "Pendant of Sunfire", Phase: 5, Quality: ItemQualityEpic, SourceZone: "Sunwell", SourceDrop: "Jewelcrafting",
		Stats:       Stats{StatStm: 27, StatInt: 19, StatSpellDmg: 34, StatHaste: 25, StatSpellCrit: 25},
		GemSlots:    []GemColor{GemColorYellow},
		SocketBonus: Stats{StatSpellDmg: 2},
	},

	// {Slot:EquipTrinket, Name:"Arcanist's Stone", Phase: 1, Quality: ItemQualityEpic, SourceZone:"H OHF - Epoch Hunter", SourceDrop:"", Stats:Stats{0, 0, 0, 25, 0, 0, 0} }
	// {Slot:EquipTrinket, Name:"Vengeance of the Illidari", Phase: 1, Quality: ItemQualityEpic, SourceZone:"Cruel's Intentions/Overlord - HFP Quest", SourceDrop:"", Stats:Stats{0, 0, 26, 0, 0, 0, 0} }
	{ID: 32330, Slot: EquipTotem, Name: "Totem of Ancestral Guidance", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "", Stats: Stats{StatSpellDmg: 85}},
	{ID: 33506, Slot: EquipTotem, Name: "Skycall Totem", Phase: 4, Quality: ItemQualityEpic, SourceZone: "Geras", SourceDrop: "20 Badges", Stats: Stats{}, Activate: ActivateSkycall, ActivateCD: neverExpires},
	{ID: 32086, Slot: EquipHead, Name: "Storm Master's Helmet", Phase: 1, Quality: ItemQualityRare, SourceZone: "Geras", SourceDrop: "50 Badges", Stats: Stats{StatStm: 24, StatInt: 32, StatSpellCrit: 24, StatSpellDmg: 37}, GemSlots: []GemColor{GemColorMeta, GemColorBlue}, SocketBonus: Stats{StatSpellCrit: 4}},
	{ID: 28602, Slot: EquipChest, Name: "Robe of the Elder Scribes", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Karazhan", SourceDrop: "Nightbane", Stats: Stats{StatStm: 27, StatInt: 29, StatSpirit: 24, StatSpellDmg: 32, StatSpellCrit: 24}, Activate: ActivateElderScribes, ActivateCD: neverExpires},

	{ID: 32963, Slot: EquipWeapon, Name: "Merciless Gladiator's Gavel / Spellblade", Phase: 2, Quality: ItemQualityEpic, SourceZone: "PvP", SourceDrop: "", Stats: Stats{StatStm: 27, StatInt: 18, StatSpellDmg: 225, StatSpellHit: 15}},
	{ID: 32524, Slot: EquipBack, Name: "Shroud of the Highborne", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Black Temple", SourceDrop: "Illidan Stormrage", Stats: Stats{StatStm: 24, StatInt: 23, StatSpellDmg: 23, StatHaste: 32}},
	{ID: 33357, Slot: EquipFeet, Name: "Footpads of Madness", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "Jan'alai", Stats: Stats{StatStm: 25, StatInt: 22, StatSpellDmg: 50, StatHaste: 25}},
	{ID: 33533, Slot: EquipLegs, Name: "Avalanche Leggings", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "Halazzi",
		Stats:       Stats{StatStm: 31, StatInt: 40, StatSpellDmg: 46, StatSpellCrit: 30},
		GemSlots:    []GemColor{GemColorRed, GemColorYellow, GemColorBlue},
		SocketBonus: Stats{StatSpellDmg: 5},
	},
	{ID: 33354, Slot: EquipWeapon, Name: "Wub's Cursed Hexblade", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "", Stats: Stats{StatInt: 21, StatSpellDmg: 217, StatSpellHit: 13, StatSpellCrit: 20, StatMP5: 6}},
	{ID: 33283, Slot: EquipWeapon, Name: "Amani Punisher", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "", Stats: Stats{StatStm: 30, StatInt: 21, StatSpellDmg: 217, StatSpellHit: 20}},
	{ID: 33466, Slot: EquipNeck, Name: "Loop of Cursed Bones", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "", Stats: Stats{StatStm: 19, StatInt: 20, StatSpellDmg: 32, StatHaste: 27}},
	{ID: 33591, Slot: EquipBack, Name: "Shadowcaster's Drape", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "", Stats: Stats{StatStm: 22, StatInt: 20, StatSpellDmg: 27, StatHaste: 25}},
	{ID: 32817, Slot: EquipWrist, Name: "Veteran's Mail Bracers", Phase: 2, Quality: ItemQualityEpic, SourceZone: "", SourceDrop: "",
		Stats:       Stats{StatStm: 25, StatInt: 14, StatSpellDmg: 22, StatSpellCrit: 17},
		GemSlots:    []GemColor{GemColorYellow},
		SocketBonus: Stats{}, // resil bonus
	},
	{ID: 32792, Slot: EquipFeet, Name: "Veteran's Mail Sabatons", Phase: 2, Quality: ItemQualityEpic, SourceZone: "", SourceDrop: "", Stats: Stats{StatStm: 39, StatInt: 27, StatSpellDmg: 32, StatSpellCrit: 26}},
	{ID: 32328, Slot: EquipHands, Name: "Botanist's Gloves of Growth", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Teron Gorefiend",
		Stats:       Stats{StatStm: 22, StatInt: 21, StatSpellDmg: 28, StatHaste: 37},
		GemSlots:    []GemColor{GemColorYellow, GemColorBlue},
		SocketBonus: Stats{StatSpellDmg: 3},
	},
	{ID: 33281, Slot: EquipNeck, Name: "Brooch of Nature's Mercy", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "", Stats: Stats{StatInt: 24, StatSpellDmg: 25, StatHaste: 33}},
	{ID: 33334, Slot: EquipOffhand, Name: "Fetish of the Primal Gods", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "", Stats: Stats{StatStm: 24, StatInt: 17, StatSpellDmg: 37, StatHaste: 17}},
	{ID: 34344, Slot: EquipHands, Name: "Handguards of Defiled Worlds", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStm: 33, StatInt: 32, StatSpellDmg: 47, StatHaste: 36, StatSpellHit: 27},
		GemSlots:    []GemColor{GemColorYellow, GemColorRed},
		SocketBonus: Stats{StatSpellDmg: 4},
	},
}

type ItemSet struct {
	Name    string
	Items   map[string]bool
	Bonuses map[int]ItemActivation // maps item count to activations
}

// cache for mapping item to set for fast resetting of sim.
var itemSetLookup = map[int32]*ItemSet{}

var sets = []ItemSet{
	{
		Name:  "Netherstrike",
		Items: map[string]bool{"Netherstrike Breastplate": true, "Netherstrike Bracers": true, "Netherstrike Belt": true},
		Bonuses: map[int]ItemActivation{3: func(sim *Simulation) Aura {
			sim.Stats[StatSpellDmg] += 23
			return Aura{ID: MagicIDNetherstrike}
		}},
	},
	{
		Name:  "The Twin Stars",
		Items: map[string]bool{"Charlotte's Ivy": true, "Lola's Eve": true},
		Bonuses: map[int]ItemActivation{2: func(sim *Simulation) Aura {
			sim.Stats[StatSpellDmg] += 15
			return Aura{ID: MagicIDNetherstrike}
		}},
	},
	{
		Name:  "Tidefury",
		Items: map[string]bool{"Tidefury Helm": true, "Tidefury Shoulderguards": true, "Tidefury Chestpiece": true, "Tidefury Kilt": true, "Tidefury Gauntlets": true},
		Bonuses: map[int]ItemActivation{
			2: func(sim *Simulation) Aura {
				return Aura{ID: MagicIDTidefury}
			},
			4: func(sim *Simulation) Aura {
				if sim.Options.Buffs.WaterShield {
					sim.Stats[StatMP5] += 3
				}
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
		Bonuses: map[int]ItemActivation{4: ActivateManaEtched, 2: func(sim *Simulation) Aura {
			sim.Stats[StatSpellHit] += 35
			return Aura{ID: MagicIDManaEtchedHit}
		}},
	},
	{
		Name:  "Cyclone Regalia",
		Items: map[string]bool{"Cyclone Faceguard (Tier 4)": true, "Cyclone Shoulderguards (Tier 4)": true, "Cyclone Chestguard (Tier 4)": true, "Cyclone Handguards (Tier 4)": true, "Cyclone Legguards (Tier 4)": true},
		Bonuses: map[int]ItemActivation{4: ActivateCycloneManaReduce, 2: func(sim *Simulation) Aura {
			if !sim.Options.Totems.Cyclone2PC && sim.Options.Totems.WrathOfAir {
				sim.Stats[StatSpellDmg] += 20 // only activate if we don't already have it from party/
			}
			return Aura{ID: MagicIDCyclone2pc}
		}},
	},
	{
		Name:  "Windhawk",
		Items: map[string]bool{"Windhawk Hauberk": true, "Windhawk Belt": true, "Windhawk Bracers": true},
		Bonuses: map[int]ItemActivation{3: func(sim *Simulation) Aura {
			if sim.Options.Buffs.WaterShield {
				sim.Stats[StatMP5] += 8
			}
			return Aura{ID: MagicIDWindhawk}
		}},
	},
	{
		Name:    "Cataclysm Regalia",
		Items:   map[string]bool{"Cataclysm Headpiece (Tier 5)": true, "Cataclysm Shoulderpads (Tier 5)": true, "Cataclysm Chestpiece (Tier 5)": true, "Cataclysm Handgrips (Tier 5)": true, "Cataclysm Leggings (Tier 5)": true},
		Bonuses: map[int]ItemActivation{4: ActivateCataclysmLBDiscount},
	},
	{
		Name: "Skyshatter Regalia",
		Items: map[string]bool{
			"Skyshatter Headguard (Tier 6)":   true,
			"Skyshatter Mantle (Tier 6)":      true,
			"Skyshatter Breastplate (Tier 6)": true,
			"Skyshatter Gauntlets (Tier 6)":   true,
			"Skyshatter Legguards (Tier 6)":   true,
			"Skyshatter Cord":                 true,
			"Skyshatter Treads":               true,
			"Skyshatter Bands":                true,
		},
		Bonuses: map[int]ItemActivation{2: func(sim *Simulation) Aura {
			sim.Stats[StatMP5] += 15
			sim.Stats[StatSpellCrit] += 35
			sim.Stats[StatSpellDmg] += 45
			return Aura{ID: MagicIDSkyshatter2pc}
		}, 4: ActivateSkyshatterImpLB},
	},
}
