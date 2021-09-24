package core

import (
	"fmt"
	"time"
)

// TODO:
//   This is a temporary place to house all item definitions. Next body of work should probably be to use the generated item definitions instead of these.

var Gems = []Gem{
	{ID: 34220, Name: "Chaotic Skyfire Diamond", Quality: ItemQualityRare, Phase: 1, Color: GemColorMeta, Stats: Stats{StatSpellCrit: 12}, Activate: ActivateCSD},
	{ID: 25897, Name: "Bracing Earthstorm Diamond", Quality: ItemQualityRare, Phase: 1, Color: GemColorMeta, Stats: Stats{StatSpellPower: 14}},
	{ID: 32641, Name: "Imbued Unstable Diamond", Quality: ItemQualityRare, Phase: 1, Color: GemColorMeta, Stats: Stats{StatSpellPower: 14}},
	{ID: 35503, Name: "Ember Skyfire Diamond", Quality: ItemQualityRare, Phase: 1, Color: GemColorMeta, Stats: Stats{StatSpellPower: 14}, Activate: ActivateESD},
	{ID: 28557, Name: "Swift Starfire Diamond", Quality: ItemQualityRare, Phase: 1, Color: GemColorMeta, Stats: Stats{StatSpellPower: 12}},
	{ID: 25893, Name: "Mystical Skyfire Diamond", Quality: ItemQualityRare, Phase: 1, Color: GemColorMeta, Stats: Stats{}, Activate: ActivateMSD},
	{ID: 25901, Name: "Insightful Earthstorm Diamond", Quality: ItemQualityRare, Phase: 1, Color: GemColorMeta, Stats: Stats{StatIntellect: 12}, Activate: ActivateIED},
	{ID: 23096, Name: "Runed Blood Garnet", Quality: ItemQualityUncommon, Phase: 1, Color: GemColorRed, Stats: Stats{StatSpellPower: 7}},
	{ID: 24030, Name: "Runed Living Ruby", Quality: ItemQualityRare, Phase: 1, Color: GemColorRed, Stats: Stats{StatSpellPower: 9}},
	{ID: 32196, Name: "Runed Crimson Spinel", Quality: ItemQualityEpic, Phase: 3, Color: GemColorRed, Stats: Stats{StatSpellPower: 12}},
	{ID: 28118, Name: "Runed Ornate Ruby", Quality: ItemQualityEpic, Phase: 1, Color: GemColorRed, Stats: Stats{StatSpellPower: 12}},
	{ID: 33133, Name: "Don Julio's Heart", Quality: ItemQualityEpic, Phase: 1, Color: GemColorRed, Stats: Stats{StatSpellPower: 14}},
	{ID: 23121, Name: "Lustrous Azure Moonstone", Quality: ItemQualityUncommon, Phase: 1, Color: GemColorBlue, Stats: Stats{StatMP5: 2}},
	{ID: 24037, Name: "Lustrous Star of Elune", Quality: ItemQualityRare, Phase: 1, Color: GemColorBlue, Stats: Stats{StatMP5: 3}},
	{ID: 32202, Name: "Lustrous Empyrean Sapphire", Quality: ItemQualityEpic, Phase: 1, Color: GemColorBlue, Stats: Stats{StatMP5: 4}},
	{ID: 23113, Name: "Brilliant Golden Draenite", Quality: ItemQualityUncommon, Phase: 1, Color: GemColorYellow, Stats: Stats{StatIntellect: 6}},
	{ID: 24047, Name: "Brilliant Dawnstone", Quality: ItemQualityRare, Phase: 1, Color: GemColorYellow, Stats: Stats{StatIntellect: 8}},
	{ID: 32204, Name: "Brilliant Lionseye", Quality: ItemQualityEpic, Phase: 3, Color: GemColorYellow, Stats: Stats{StatIntellect: 10}},
	{ID: 23114, Name: "Gleaming Golden Draenite", Quality: ItemQualityUncommon, Phase: 1, Color: GemColorYellow, Stats: Stats{StatSpellCrit: 6}},
	{ID: 24050, Name: "Gleaming Dawnstone", Quality: ItemQualityRare, Phase: 1, Color: GemColorYellow, Stats: Stats{StatSpellCrit: 8}},
	{ID: 32207, Name: "Gleaming Lionseye", Quality: ItemQualityEpic, Phase: 3, Color: GemColorYellow, Stats: Stats{StatSpellCrit: 10}},
	{ID: 30551, Name: "Infused Fire Opal", Quality: ItemQualityEpic, Phase: 1, Color: GemColorOrange, Stats: Stats{StatIntellect: 4, StatSpellPower: 6}},
	{ID: 23101, Name: "Potent Flame Spessarite", Quality: ItemQualityUncommon, Phase: 1, Color: GemColorOrange, Stats: Stats{StatSpellCrit: 3, StatSpellPower: 4}},
	{ID: 24059, Name: "Potent Noble Topaz", Quality: ItemQualityRare, Phase: 1, Color: GemColorOrange, Stats: Stats{StatSpellCrit: 4, StatSpellPower: 5}},
	{ID: 32218, Name: "Potent Pyrestone", Quality: ItemQualityEpic, Phase: 3, Color: GemColorOrange, Stats: Stats{StatSpellCrit: 5, StatSpellPower: 6}},
	{ID: 35760, Name: "Reckless Pyrestone", Quality: ItemQualityEpic, Phase: 3, Color: GemColorOrange, Stats: Stats{StatSpellHaste: 5, StatSpellPower: 6}},
	{ID: 30588, Name: "Potent Fire Opal", Quality: ItemQualityEpic, Phase: 1, Color: GemColorOrange, Stats: Stats{StatSpellPower: 6, StatSpellCrit: 4}},
	{ID: 28123, Name: "Potent Ornate Topaz", Quality: ItemQualityEpic, Phase: 1, Color: GemColorOrange, Stats: Stats{StatSpellPower: 6, StatSpellCrit: 5}},
	{ID: 31866, Name: "Veiled Flame Spessarite", Quality: ItemQualityUncommon, Phase: 1, Color: GemColorOrange, Stats: Stats{StatSpellHit: 3, StatSpellPower: 4}},
	{ID: 31867, Name: "Veiled Noble Topaz", Quality: ItemQualityRare, Phase: 1, Color: GemColorOrange, Stats: Stats{StatSpellHit: 4, StatSpellPower: 5}},
	{ID: 32221, Name: "Shining Fire Opal", Quality: ItemQualityEpic, Phase: 1, Color: GemColorOrange, Stats: Stats{StatSpellHit: 5, StatSpellPower: 6}},
	{ID: 30564, Name: "Veiled Pyrestone", Quality: ItemQualityEpic, Phase: 3, Color: GemColorOrange, Stats: Stats{StatSpellHit: 5, StatSpellPower: 6}},
	{ID: 30560, Name: "Rune Covered Chrysoprase", Quality: ItemQualityEpic, Phase: 1, Color: GemColorGreen, Stats: Stats{StatMP5: 2, StatSpellCrit: 5}},
	{ID: 24065, Name: "Dazzling Talasite", Quality: ItemQualityRare, Phase: 1, Color: GemColorGreen, Stats: Stats{StatMP5: 2, StatIntellect: 4}},
	{ID: 35759, Name: "Forceful Seaspray Emerald", Quality: ItemQualityEpic, Phase: 3, Color: GemColorGreen, Stats: Stats{StatSpellHaste: 5, StatStamina: 7}},
	{ID: 24056, Name: "Glowing Nightseye", Quality: ItemQualityRare, Phase: 1, Color: GemColorPurple, Stats: Stats{StatSpellPower: 5, StatStamina: 6}},
	{ID: 30555, Name: "Glowing Tanzanite", Quality: ItemQualityEpic, Phase: 1, Color: GemColorPurple, Stats: Stats{StatSpellPower: 6, StatStamina: 6}},
	{ID: 32215, Name: "Glowing Shadowsong Amethyst", Quality: ItemQualityEpic, Phase: 3, Color: GemColorPurple, Stats: Stats{StatSpellPower: 6, StatStamina: 7}},
	{ID: 31116, Name: "Infused Amethyst", Quality: ItemQualityEpic, Phase: 1, Color: GemColorPurple, Stats: Stats{StatSpellPower: 6, StatStamina: 6}},
	{ID: 30600, Name: "Fluorescent Tanzanite", Quality: ItemQualityEpic, Phase: 1, Color: GemColorPurple, Stats: Stats{StatSpellPower: 6, StatSpirit: 4}},
	{ID: 30605, Name: "Vivid Chrysoprase", Quality: ItemQualityEpic, Phase: 1, Color: GemColorGreen, Stats: Stats{StatSpellHit: 5, StatStamina: 6}},
}

var Enchants = []Enchant{
	{ID: 29191, EffectID: 3002, Name: "Glyph of Power", Bonus: Stats{StatSpellPower: 22, StatSpellHit: 14}, ItemType: ItemTypeHead},
	{ID: 28909, EffectID: 2995, Name: "Greater Inscription of the Orb", Bonus: Stats{StatSpellPower: 12, StatSpellCrit: 15}, ItemType: ItemTypeShoulder},
	{ID: 28886, EffectID: 2982, Name: "Greater Inscription of Discipline", Bonus: Stats{StatSpellPower: 18, StatSpellCrit: 10}, ItemType: ItemTypeShoulder},
	{ID: 20076, EffectID: 2605, Name: "Zandalar Signet of Mojo", Bonus: Stats{StatSpellPower: 18}, ItemType: ItemTypeShoulder},
	{ID: 23545, EffectID: 2721, Name: "Power of the Scourge", Bonus: Stats{StatSpellPower: 15, StatSpellCrit: 14}, ItemType: ItemTypeShoulder},
	{ID: 27960, EffectID: 2661, Name: "Chest - Exceptional Stats", Bonus: Stats{StatStamina: 6, StatIntellect: 6, StatSpirit: 6}, ItemType: ItemTypeChest},
	{ID: 22534, EffectID: 2650, Name: "Bracer - Spellpower", Bonus: Stats{StatSpellPower: 15}, ItemType: ItemTypeWrist},
	{ID: 28272, EffectID: 2937, Name: "Gloves - Major Spellpower", Bonus: Stats{StatSpellPower: 20}, ItemType: ItemTypeHands},
	{ID: 24274, EffectID: 2748, Name: "Runic Spellthread", Bonus: Stats{StatSpellPower: 35, StatStamina: 20}, ItemType: ItemTypeLegs},
	{ID: 24273, EffectID: 2747, Name: "Mystic Spellthread", Bonus: Stats{StatSpellPower: 25, StatStamina: 15}, ItemType: ItemTypeLegs},
	{ID: 22555, EffectID: 2669, Name: "Weapon - Major Spellpower", Bonus: Stats{StatSpellPower: 40}, ItemType: ItemTypeWeapon},
	{ID: 35445, EffectID: 2928, Name: "Ring - Spellpower", Bonus: Stats{StatSpellPower: 12}, ItemType: ItemTypeFinger},
	{ID: 27945, EffectID: 2654, Name: "Shield - Intellect", Bonus: Stats{StatIntellect: 12}, ItemType: ItemTypeWeapon, WeaponType: WeaponTypeShield},
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
			statsMatch := it.ItemType == v.ItemType
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
	ID               int32
	ItemType         ItemType
	ArmorType        byte
	WeaponType       WeaponType
	HandType         HandType
	RangedWeaponType RangedWeaponType

	Name       string
	SourceZone string
	SourceDrop string
	Stats      Stats // Stats applied to wearer
	Phase      byte
	Quality    ItemQuality

	GemSockets  []GemColor
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
	SharedID   int32          `json:"-"` // ID used for shared item cooldowns (trinkets etc)
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
	ID         int32 // ID of the enchant item.
	EffectID   int32 // Used by UI to apply effect to tooltip
	Name       string
	Bonus      Stats
	ItemType   ItemType // which slot does the enchant go on.
	HandType   HandType // If ItemType is weapon, check hand type / weapon type
	WeaponType WeaponType
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

// ItemActivation needs the state from simulator, party, and player
//  because items can impact all 3. (potions, drums, JC necks, etc)
type ItemActivation func(*Simulation, *Party, PlayerAgent) Aura

type Equipment [ItemSlotRanged + 1]Item

// Structs used for looking up items/gems/enchants
type ItemSpec struct {
	ID      int32
	Enchant int32
	Gems    []int32
}
type EquipmentSpec [ItemSlotRanged + 1]ItemSpec

func NewEquipmentSet(equipSpec EquipmentSpec) Equipment {
	equipment := Equipment{}

	for _, itemSpec := range equipSpec {
		item := Item{}
		if foundItem, ok := ItemsByID[itemSpec.ID]; ok {
			item = foundItem
		} else {
			if itemSpec.ID != 0 {
				panic(fmt.Sprintf("No item with id: %d", itemSpec.ID))
			}
			continue
		}

		if itemSpec.Enchant != 0 {
			if enchant, ok := EnchantsByID[itemSpec.Enchant]; ok {
				item.Enchant = enchant
			} else {
				panic(fmt.Sprintf("No enchant with id: %d", itemSpec.Enchant))
			}
		}

		if len(itemSpec.Gems) > 0 {
			item.Gems = make([]Gem, len(item.GemSockets))
			for gemIdx, gemID := range itemSpec.Gems {
				if gemIdx >= len(item.GemSockets) {
					break // in case we get invalid gem settings.
				}
				if gem, ok := GemsByID[gemID]; ok {
					item.Gems[gemIdx] = gem
				} else {
					if gemID != 0 {
						panic(fmt.Sprintf("No gem with id: %d", gemID))
					}
				}
			}
		}

		if item.ItemType == ItemTypeFinger {
			if equipment[ItemSlotFinger1].Name == "" {
				equipment[ItemSlotFinger1] = item
			} else {
				equipment[ItemSlotFinger2] = item
			}
		} else if item.ItemType == ItemTypeTrinket {
			if equipment[ItemSlotTrinket1].Name == "" {
				equipment[ItemSlotTrinket1] = item
			} else {
				equipment[ItemSlotTrinket2] = item
			}
		} else if item.ItemType == ItemTypeWeapon {
			if item.WeaponType == WeaponTypeShield {
				if equipment[ItemSlotMainHand].HandType != HandTypeTwoHand {
					equipment[ItemSlotOffHand] = item
				}
			} else if item.HandType == HandTypeMainHand || item.HandType == HandTypeUnknown {
				equipment[ItemSlotMainHand] = item
			} else if item.HandType == HandTypeTwoHand {
				equipment[ItemSlotMainHand] = item
				equipment[ItemSlotOffHand] = Item{} // clear offhand
			}
		} else {
			equipment[item.ItemType.Slot()] = item
		}
	}
	return equipment
}

type RangedWeaponType byte

const (
	RangedWeaponTypeUnknown RangedWeaponType = iota
	RangedWeaponTypeBow
	RangedWeaponTypeCrossbow
	RangedWeaponTypeGun
	RangedWeaponTypeLibram
	RangedWeaponTypeRelic
	RangedWeaponTypeThrown
	RangedWeaponTypeTotem
	RangedWeaponTypeWand
)

type HandType byte

const (
	HandTypeUnknown HandType = iota
	HandTypeMainHand
	HandTypeOneHand
	HandTypeOffHand
	HandTypeTwoHand
)

type WeaponType byte

const (
	WeaponTypeUnknown WeaponType = iota
	WeaponTypeAxe
	WeaponTypeDagger
	WeaponTypeFist
	WeaponTypeMace
	WeaponTypePolearm
	WeaponTypeOffHand
	WeaponTypeStaff
	WeaponTypeSword
	WeaponTypeShield
)

type ItemType byte

const (
	ItemTypeUnknown ItemType = iota
	ItemTypeHead
	ItemTypeNeck
	ItemTypeShoulder
	ItemTypeBack
	ItemTypeChest
	ItemTypeWrist
	ItemTypeHands
	ItemTypeWaist
	ItemTypeLegs
	ItemTypeFeet
	ItemTypeFinger
	ItemTypeTrinket
	ItemTypeWeapon
	ItemTypeRanged
)

func (it ItemType) Slot() ItemSlot {
	switch it {
	case ItemTypeHead:
		return ItemSlotHead
	case ItemTypeNeck:
		return ItemSlotNeck
	case ItemTypeShoulder:
		return ItemSlotShoulder
	case ItemTypeBack:
		return ItemSlotBack
	case ItemTypeChest:
		return ItemSlotChest
	case ItemTypeWrist:
		return ItemSlotWrist
	case ItemTypeHands:
		return ItemSlotHands
	case ItemTypeWaist:
		return ItemSlotWaist
	case ItemTypeLegs:
		return ItemSlotLegs
	case ItemTypeFeet:
		return ItemSlotFeet
	case ItemTypeFinger:
		return ItemSlotFinger1
	case ItemTypeTrinket:
		return ItemSlotTrinket1
	case ItemTypeWeapon:
		return ItemSlotMainHand
	case ItemTypeRanged:
		return ItemSlotRanged
	}

	return 255
}

type ItemSlot byte

const (
	ItemSlotHead ItemSlot = iota
	ItemSlotNeck
	ItemSlotShoulder
	ItemSlotBack
	ItemSlotChest
	ItemSlotWrist
	ItemSlotHands
	ItemSlotWaist
	ItemSlotLegs
	ItemSlotFeet
	ItemSlotFinger1
	ItemSlotFinger2
	ItemSlotTrinket1
	ItemSlotTrinket2
	ItemSlotMainHand // can be 1h or 2h
	ItemSlotOffHand
	ItemSlotRanged
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
			if v == 0 {
				continue
			}
			s[k] += v
		}
		if len(item.GemSockets) > 0 {
			isMatched := len(item.Gems) == len(item.GemSockets) && len(item.GemSockets) > 0
			for gi, g := range item.Gems {
				for k, v := range g.Stats {
					s[k] += v
				}
				isMatched = isMatched && g.Color.Intersects(item.GemSockets[gi])
				if !isMatched {
					break
				}
			}
			if isMatched {
				for k, v := range item.SocketBonus {
					if v == 0 {
						continue
					}
					s[k] += v
				}
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
	{ID: 27471, ItemType: ItemTypeHead, Name: "Gladiator's Mail Helm", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Arena Season 1 Reward", SourceDrop: "", Stats: Stats{StatIntellect: 15, StatStamina: 54, StatSpellCrit: 18, StatSpellPower: 37}, GemSockets: []GemColor{0x1, 0x2}},
	{ID: 24266, ItemType: ItemTypeHead, Name: "Spellstrike Hood", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Tailoring BoE", SourceDrop: "", Stats: Stats{StatIntellect: 12, StatStamina: 16, StatSpellCrit: 24, StatSpellHit: 16, StatSpellPower: 46}, GemSockets: []GemColor{0x2, 0x4, 0x3}, SocketBonus: Stats{StatStamina: 6}},
	{ID: 28278, ItemType: ItemTypeHead, Name: "Incanter's Cowl", Phase: 1, Quality: ItemQualityRare, SourceZone: "Mech - Pathaleon the Calculator", SourceDrop: "", Stats: Stats{StatIntellect: 27, StatStamina: 15, StatSpellCrit: 19, StatSpellPower: 29}, GemSockets: []GemColor{0x1, 0x4}},
	{ID: 31330, ItemType: ItemTypeHead, Name: "Lightning Crown", Phase: 1, Quality: ItemQualityEpic, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{0, 0, 43, 0, 66}},
	{ID: 28415, ItemType: ItemTypeHead, Name: "Hood of Oblivion", Phase: 1, Quality: ItemQualityRare, SourceZone: "Arc - Harbinger Skyriss", SourceDrop: "", Stats: Stats{StatIntellect: 32, StatStamina: 27, StatSpellPower: 40}, GemSockets: []GemColor{0x1, 0x3}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 28758, ItemType: ItemTypeHead, Name: "Exorcist's Mail Helm", Phase: 1, Quality: ItemQualityRare, SourceZone: "18 Spirit Shards", SourceDrop: "", Stats: Stats{StatIntellect: 16, StatStamina: 30, StatSpellCrit: 24, StatSpellPower: 29}, GemSockets: []GemColor{0x1}, SocketBonus: Stats{StatSpellCrit: 3}},
	{ID: 28349, ItemType: ItemTypeHead, Name: "Tidefury Helm", Phase: 1, Quality: ItemQualityRare, SourceZone: "Bot - Warp Splinter", SourceDrop: "", Stats: Stats{StatIntellect: 26, StatStamina: 32, StatSpellPower: 32, StatMP5: 6}, GemSockets: []GemColor{0x1, 0x4}, SocketBonus: Stats{StatIntellect: 4}},
	{ID: 29504, ItemType: ItemTypeHead, Name: "Windscale Hood", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Leatherworking BoE", SourceDrop: "", Stats: Stats{StatIntellect: 18, StatStamina: 16, StatSpellCrit: 37, StatSpellPower: 44, StatMP5: 10}},
	{ID: 31107, ItemType: ItemTypeHead, Name: "Shamanistic Helmet of Second Sight", Phase: 1, Quality: ItemQualityRare, SourceZone: "Teron Gorfiend, I am... - SMV Quest", SourceDrop: "", Stats: Stats{StatIntellect: 15, StatStamina: 12, StatSpellCrit: 24, StatSpellPower: 35, StatMP5: 4}, GemSockets: []GemColor{0x4, 0x3, 0x3}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 28193, ItemType: ItemTypeHead, Name: "Mana-Etched Crown", Phase: 1, Quality: ItemQualityRare, SourceZone: "BM - Aeonus", SourceDrop: "", Stats: Stats{StatIntellect: 20, StatStamina: 27, StatSpellPower: 34}, GemSockets: []GemColor{0x1, 0x2}},
	{ID: 28169, ItemType: ItemTypeHead, Name: "Mag'hari Ritualist's Horns", Phase: 1, Quality: ItemQualityRare, SourceZone: "Hero of the Mag'har - Nagrand quest (Horde)", SourceDrop: "", Stats: Stats{StatIntellect: 16, StatStamina: 18, StatSpellCrit: 15, StatSpellHit: 12, StatSpellPower: 50}},
	{ID: 27488, ItemType: ItemTypeHead, Name: "Mage-Collar of the Firestorm", Phase: 1, Quality: ItemQualityRare, SourceZone: "H BF - The Maker", SourceDrop: "", Stats: Stats{StatIntellect: 33, StatStamina: 32, StatSpellCrit: 23, StatSpellPower: 39}},
	{ID: 30297, ItemType: ItemTypeHead, Name: "Circlet of the Starcaller", Phase: 1, Quality: ItemQualityRare, SourceZone: "Dimensius the All-Devouring - NS Quest", SourceDrop: "", Stats: Stats{StatIntellect: 18, StatStamina: 27, StatSpellCrit: 18, StatSpellPower: 47}},
	{ID: 27993, ItemType: ItemTypeHead, Name: "Mask of Inner Fire", Phase: 1, Quality: ItemQualityRare, SourceZone: "BM - Chrono Lord Deja", SourceDrop: "", Stats: Stats{StatIntellect: 33, StatStamina: 30, StatSpellCrit: 22, StatSpellPower: 37}},
	{ID: 30946, ItemType: ItemTypeHead, Name: "Mooncrest Headdress", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "Blast the Infernals! - SMV Quest", SourceDrop: "", Stats: Stats{StatIntellect: 16, StatSpellCrit: 21, StatSpellPower: 44}},
	{ID: 28245, ItemType: ItemTypeNeck, Name: "Pendant of Dominance", Phase: 1, Quality: ItemQualityEpic, SourceZone: "15,300 Honor & 10 EotS Marks", SourceDrop: "", Stats: Stats{StatIntellect: 12, StatStamina: 31, StatSpellPower: 26}, GemSockets: []GemColor{0x4}, SocketBonus: Stats{StatSpellCrit: 2}},
	{ID: 28134, ItemType: ItemTypeNeck, Name: "Brooch of Heightened Potential", Phase: 1, Quality: ItemQualityRare, SourceZone: "SLabs - Blackheart the Inciter", SourceDrop: "", Stats: Stats{StatIntellect: 14, StatStamina: 15, StatSpellCrit: 14, StatSpellHit: 9, StatSpellPower: 22}},
	{ID: 29333, ItemType: ItemTypeNeck, Name: "Torc of the Sethekk Prophet", Phase: 1, Quality: ItemQualityRare, SourceZone: "Brother Against Brother - Auchindoun ", SourceDrop: "", Stats: Stats{StatIntellect: 18, StatSpellCrit: 21, StatSpellPower: 19}},
	{ID: 31692, ItemType: ItemTypeNeck, Name: "Natasha's Ember Necklace", Phase: 1, Quality: ItemQualityRare, SourceZone: "The Hound-Master - BEM Quest", SourceDrop: "", Stats: Stats{StatIntellect: 15, StatSpellCrit: 10, StatSpellPower: 29}},
	{ID: 28254, ItemType: ItemTypeNeck, Name: "Warp Engineer's Prismatic Chain", Phase: 1, Quality: ItemQualityRare, SourceZone: "Mech - Mechano Lord Capacitus", SourceDrop: "", Stats: Stats{StatIntellect: 18, StatStamina: 17, StatSpellCrit: 16, StatSpellPower: 19}},
	{ID: 27758, ItemType: ItemTypeNeck, Name: "Hydra-fang Necklace", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H UB - Ghaz'an", SourceDrop: "", Stats: Stats{StatIntellect: 16, StatStamina: 17, StatSpellHit: 16, StatSpellPower: 19}},
	{ID: 31693, ItemType: ItemTypeNeck, Name: "Natasha's Arcane Filament", Phase: 1, Quality: ItemQualityEpic, SourceZone: "The Hound-Master - BEM Quest", SourceDrop: "", Stats: Stats{StatIntellect: 10, StatStamina: 22, StatSpellPower: 29}},
	{ID: 27464, ItemType: ItemTypeNeck, Name: "Omor's Unyielding Will", Phase: 1, Quality: ItemQualityRare, SourceZone: "H Ramps - Omar the Unscarred", SourceDrop: "", Stats: Stats{StatIntellect: 19, StatStamina: 19, StatSpellPower: 25}},
	{ID: 31338, ItemType: ItemTypeNeck, Name: "Charlotte's Ivy", Phase: 1, Quality: ItemQualityEpic, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{StatIntellect: 19, StatStamina: 18, StatSpellPower: 23}},
	{ID: 27473, ItemType: ItemTypeShoulder, Name: "Gladiator's Mail Spaulders", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Arena Season 1 Reward", SourceDrop: "", Stats: Stats{StatIntellect: 17, StatStamina: 33, StatSpellCrit: 20, StatSpellPower: 22, StatMP5: 6}, GemSockets: []GemColor{0x2, 0x4}},
	{ID: 32078, ItemType: ItemTypeShoulder, Name: "Pauldrons of Wild Magic", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H SP - Quagmirran", SourceDrop: "", Stats: Stats{StatIntellect: 28, StatStamina: 21, StatSpellCrit: 23, StatSpellPower: 33}},
	{ID: 27796, ItemType: ItemTypeShoulder, Name: "Mana-Etched Spaulders", Phase: 1, Quality: ItemQualityRare, SourceZone: "H UB - Quagmirran", SourceDrop: "", Stats: Stats{StatIntellect: 17, StatStamina: 25, StatSpellCrit: 16, StatSpellPower: 20}, GemSockets: []GemColor{0x2, 0x4}},
	{ID: 30925, ItemType: ItemTypeShoulder, Name: "Spaulders of the Torn-heart", Phase: 1, Quality: ItemQualityRare, SourceZone: "The Cipher of Damnation - SMV Quest", SourceDrop: "", Stats: Stats{StatIntellect: 7, StatStamina: 10, StatSpellCrit: 18, StatSpellPower: 40}},
	{ID: 31797, ItemType: ItemTypeShoulder, Name: "Elekk Hide Spaulders", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "The Fallen Exarch - Terokkar Forest Quest", SourceDrop: "", Stats: Stats{StatIntellect: 12, StatSpellCrit: 28, StatSpellPower: 25}},
	{ID: 27778, ItemType: ItemTypeShoulder, Name: "Spaulders of Oblivion", Phase: 1, Quality: ItemQualityRare, SourceZone: "SLabs - Murmur", SourceDrop: "", Stats: Stats{StatIntellect: 17, StatStamina: 25, StatSpellPower: 29}, GemSockets: []GemColor{0x4, 0x3}, SocketBonus: Stats{StatSpellHit: 3}},
	{ID: 27802, ItemType: ItemTypeShoulder, Name: "Tidefury Shoulderguards", Phase: 1, Quality: ItemQualityRare, SourceZone: "SH - O'mrogg", SourceDrop: "", Stats: Stats{StatIntellect: 23, StatStamina: 18, StatSpellPower: 19, StatMP5: 6}, GemSockets: []GemColor{0x2, 0x3}, SocketBonus: Stats{StatSpellPower: 4}},
	{ID: 27994, ItemType: ItemTypeShoulder, Name: "Mantle of Three Terrors", Phase: 1, Quality: ItemQualityRare, SourceZone: "BM - Chrono Lord Deja", SourceDrop: "", Stats: Stats{StatIntellect: 25, StatStamina: 29, StatSpellHit: 12, StatSpellPower: 29}},
	{ID: 25777, ItemType: ItemTypeBack, Name: "Ogre Slayer's Cover", Phase: 1, Quality: ItemQualityRare, SourceZone: "Cho'war the Pillager - Nagrand Quest", SourceDrop: "", Stats: Stats{StatIntellect: 18, StatSpellCrit: 16, StatSpellPower: 20}},
	{ID: 28269, ItemType: ItemTypeBack, Name: "Baba's Cloak of Arcanistry", Phase: 1, Quality: ItemQualityRare, SourceZone: "Mech - Pathaleon the Calculator", SourceDrop: "", Stats: Stats{StatIntellect: 15, StatStamina: 15, StatSpellCrit: 14, StatSpellPower: 22}},
	{ID: 29813, ItemType: ItemTypeBack, Name: "Cloak of Woven Energy", Phase: 1, Quality: ItemQualityRare, SourceZone: "Hitting the Motherlode - Netherstorm Quest", SourceDrop: "", Stats: Stats{StatIntellect: 13, StatStamina: 6, StatSpellCrit: 6, StatSpellPower: 29}},
	{ID: 27981, ItemType: ItemTypeBack, Name: "Sethekk Oracle Cloak", Phase: 1, Quality: ItemQualityRare, SourceZone: "SH - Talon King Ikiss", SourceDrop: "", Stats: Stats{StatIntellect: 18, StatStamina: 18, StatSpellHit: 12, StatSpellPower: 22}},
	{ID: 32541, ItemType: ItemTypeBack, Name: "Terokk's Wisdom", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Terokk - Skettis Summoned Boss", SourceDrop: "", Stats: Stats{StatIntellect: 16, StatStamina: 18, StatSpellPower: 33}},
	{ID: 24252, ItemType: ItemTypeBack, Name: "Cloak of the Black Void", Phase: 1, Quality: ItemQualityRare, SourceZone: "Tailoring BoE", SourceDrop: "", Stats: Stats{StatIntellect: 11, StatSpellPower: 35}},
	{ID: 31140, ItemType: ItemTypeBack, Name: "Cloak of Entropy", Phase: 1, Quality: ItemQualityRare, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{StatIntellect: 11, StatSpellHit: 10, StatSpellPower: 25}},
	{ID: 28379, ItemType: ItemTypeBack, Name: "Sergeant's Heavy Cape", Phase: 1, Quality: ItemQualityEpic, SourceZone: "9,435 Honor & 20 AB Marks", SourceDrop: "", Stats: Stats{StatIntellect: 12, StatStamina: 33, StatSpellPower: 26}},
	{ID: 27469, ItemType: ItemTypeChest, Name: "Gladiator's Mail Armor", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Arena Season 1 Reward", SourceDrop: "", Stats: Stats{StatIntellect: 23, StatStamina: 42, StatSpellCrit: 23, StatSpellPower: 32, StatMP5: 7}, GemSockets: []GemColor{0x2, 0x4, 0x4}, SocketBonus: Stats{StatSpellCrit: 4}},
	{ID: 31340, ItemType: ItemTypeChest, Name: "Will of Edward the Odd", Phase: 1, Quality: ItemQualityEpic, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{StatIntellect: 30, StatSpellCrit: 30, StatSpellPower: 53}},
	{ID: 29129, ItemType: ItemTypeChest, Name: "Anchorite's Robe", Phase: 1, Quality: ItemQualityEpic, SourceZone: "The Aldor - Honored", SourceDrop: "", Stats: Stats{StatIntellect: 38, StatStamina: 16, StatSpellPower: 29, StatMP5: 18}, GemSockets: []GemColor{0x4, 0x4, 0x3}},
	{ID: 28231, ItemType: ItemTypeChest, Name: "Tidefury Chestpiece", Phase: 1, Quality: ItemQualityRare, SourceZone: "Arc - Harbinger Skyriss", SourceDrop: "", Stats: Stats{StatIntellect: 22, StatStamina: 28, StatSpellHit: 10, StatSpellPower: 36, StatMP5: 4}, GemSockets: []GemColor{0x4, 0x4, 0x3}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 29341, ItemType: ItemTypeChest, Name: "Auchenai Anchorite's Robe", Phase: 1, Quality: ItemQualityRare, SourceZone: "Everything Will Be Alright - AC Quest", SourceDrop: "", Stats: Stats{StatIntellect: 24, StatSpellPower: 28, StatSpellHit: 23}, GemSockets: []GemColor{0x2, 0x4, 0x4}, SocketBonus: Stats{StatSpellCrit: 4}},
	{ID: 28191, ItemType: ItemTypeChest, Name: "Mana-Etched Vestments", Phase: 1, Quality: ItemQualityRare, SourceZone: "OHF - Epoch Hunter", SourceDrop: "", Stats: Stats{StatIntellect: 25, StatStamina: 25, StatSpellCrit: 17, StatSpellPower: 29}, GemSockets: []GemColor{0x2, 0x4, 0x3}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 31297, ItemType: ItemTypeChest, Name: "Robe of the Crimson Order", Phase: 1, Quality: ItemQualityEpic, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{StatIntellect: 23, StatSpellHit: 30, StatSpellPower: 50}},
	{ID: 28342, ItemType: ItemTypeChest, Name: "Warp Infused Drape", Phase: 1, Quality: ItemQualityRare, SourceZone: "Bot - Warp Splinter", SourceDrop: "", Stats: Stats{StatIntellect: 28, StatStamina: 27, StatSpellHit: 12, StatSpellPower: 30}, GemSockets: []GemColor{0x2, 0x4, 0x3}},
	{ID: 28232, ItemType: ItemTypeChest, Name: "Robe of Oblivion", Phase: 1, Quality: ItemQualityRare, SourceZone: "SLabs - Murmur", SourceDrop: "", Stats: Stats{StatIntellect: 20, StatStamina: 30, StatSpellPower: 40}, GemSockets: []GemColor{0x2, 0x4, 0x3}},
	{ID: 28229, ItemType: ItemTypeChest, Name: "Incanter's Robe", Phase: 1, Quality: ItemQualityRare, SourceZone: "Bot - Warp Splinter", SourceDrop: "", Stats: Stats{StatIntellect: 22, StatStamina: 24, StatSpellCrit: 8, StatSpellPower: 29}, GemSockets: []GemColor{0x2, 0x4, 0x4}},
	{ID: 27824, ItemType: ItemTypeChest, Name: "Robe of the Great Dark Beyond", Phase: 1, Quality: ItemQualityRare, SourceZone: "MT - Tavarok", SourceDrop: "", Stats: Stats{StatIntellect: 30, StatStamina: 25, StatSpellCrit: 23, StatSpellPower: 39}},
	{ID: 28391, ItemType: ItemTypeChest, Name: "Worldfire Chestguard", Phase: 1, Quality: ItemQualityRare, SourceZone: "Arc - Dalliah the Doomsayer", SourceDrop: "", Stats: Stats{StatIntellect: 32, StatStamina: 33, StatSpellCrit: 22, StatSpellPower: 40}},
	{ID: 28638, ItemType: ItemTypeWrist, Name: "General's Mail Bracers", Phase: 1, Quality: ItemQualityEpic, SourceZone: "7,548 Honor & 20 WSG Marks", SourceDrop: "", Stats: Stats{StatIntellect: 12, StatStamina: 22, StatSpellCrit: 14, StatSpellPower: 20}, GemSockets: []GemColor{0x4}},
	{ID: 27522, ItemType: ItemTypeWrist, Name: "World's End Bracers", Phase: 1, Quality: ItemQualityRare, SourceZone: "H BF - Keli'dan the Breaker", SourceDrop: "", Stats: Stats{StatIntellect: 19, StatStamina: 18, StatSpellCrit: 17, StatSpellPower: 22}},
	{ID: 24250, ItemType: ItemTypeWrist, Name: "Bracers of Havok", Phase: 1, Quality: ItemQualityRare, SourceZone: "Tailoring BoE", SourceDrop: "", Stats: Stats{StatIntellect: 12, StatSpellPower: 30}, GemSockets: []GemColor{0x4}, SocketBonus: Stats{StatSpellCrit: 2}},
	{ID: 27462, ItemType: ItemTypeWrist, Name: "Crimson Bracers of Gloom", Phase: 1, Quality: ItemQualityRare, SourceZone: "H Ramps - Omor the Unscarred", SourceDrop: "", Stats: Stats{StatIntellect: 18, StatStamina: 18, StatSpellHit: 12, StatSpellPower: 22}},
	{ID: 29240, ItemType: ItemTypeWrist, Name: "Bands of Negation", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H MT - Nexus- Prince Shaffar", SourceDrop: "", Stats: Stats{StatIntellect: 22, StatStamina: 25, StatSpellPower: 29}},
	{ID: 27746, ItemType: ItemTypeWrist, Name: "Arcanium Signet Bands", Phase: 1, Quality: ItemQualityRare, SourceZone: "H UB - Hungarfen", SourceDrop: "", Stats: Stats{StatIntellect: 15, StatStamina: 14, StatSpellPower: 30}},
	{ID: 29243, ItemType: ItemTypeWrist, Name: "Wave-Fury Vambraces", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H SV - Warlod Kalithresh", SourceDrop: "", Stats: Stats{StatIntellect: 18, StatStamina: 19, StatSpellPower: 22, StatMP5: 5}},
	{ID: 29955, ItemType: ItemTypeWrist, Name: "Mana Infused Wristguards", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "A Fate Worse Than Death - Netherstorm Quest", SourceDrop: "", Stats: Stats{StatIntellect: 8, StatStamina: 12, StatSpellPower: 25}},
	{ID: 27465, ItemType: ItemTypeHands, Name: "Mana-Etched Gloves", Phase: 1, Quality: ItemQualityRare, SourceZone: "H Ramps - Omor the Unscarred", SourceDrop: "", Stats: Stats{StatIntellect: 17, StatStamina: 25, StatSpellCrit: 16, StatSpellPower: 20}, GemSockets: []GemColor{0x2, 0x4}},
	{ID: 27793, ItemType: ItemTypeHands, Name: "Earth Mantle Handwraps", Phase: 1, Quality: ItemQualityRare, SourceZone: "SV - Mekgineer Steamrigger", SourceDrop: "", Stats: Stats{StatIntellect: 18, StatStamina: 21, StatSpellCrit: 16, StatSpellPower: 19}, GemSockets: []GemColor{0x2, 0x4}, SocketBonus: Stats{StatIntellect: 3}},
	{ID: 31149, ItemType: ItemTypeHands, Name: "Gloves of Pandemonium", Phase: 1, Quality: ItemQualityRare, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{StatIntellect: 15, StatSpellCrit: 22, StatSpellHit: 10, StatSpellPower: 25}},
	{ID: 27470, ItemType: ItemTypeHands, Name: "Gladiator's Mail Gauntlets", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Arena Season 1 Reward", SourceDrop: "", Stats: Stats{StatIntellect: 18, StatStamina: 36, StatSpellCrit: 21, StatSpellPower: 32}},
	{ID: 31280, ItemType: ItemTypeHands, Name: "Thundercaller's Gauntlets", Phase: 1, Quality: ItemQualityRare, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{StatIntellect: 16, StatStamina: 16, StatSpellCrit: 18, StatSpellPower: 35}},
	{ID: 30924, ItemType: ItemTypeHands, Name: "Gloves of the High Magus", Phase: 1, Quality: ItemQualityRare, SourceZone: "News of Victory - SMV Quest", SourceDrop: "", Stats: Stats{StatIntellect: 18, StatStamina: 13, StatSpellCrit: 22, StatSpellPower: 26}},
	{ID: 29317, ItemType: ItemTypeHands, Name: "Tempest's Touch", Phase: 1, Quality: ItemQualityRare, SourceZone: "Return to Andormu - CoT Quest", SourceDrop: "", Stats: Stats{StatIntellect: 20, StatStamina: 10, StatSpellPower: 27}, GemSockets: []GemColor{0x3, 0x3}},
	{ID: 27493, ItemType: ItemTypeHands, Name: "Gloves of the Deadwatcher", Phase: 1, Quality: ItemQualityRare, SourceZone: "H AC - Shirrak the Dead Watcher", SourceDrop: "", Stats: Stats{StatIntellect: 24, StatStamina: 24, StatSpellHit: 18, StatSpellPower: 29}},
	{ID: 27508, ItemType: ItemTypeHands, Name: "Incanter's Gloves", Phase: 1, Quality: ItemQualityRare, SourceZone: "SV - Thespia", SourceDrop: "", Stats: Stats{StatIntellect: 24, StatStamina: 21, StatSpellCrit: 14, StatSpellPower: 29}},
	{ID: 24452, ItemType: ItemTypeHands, Name: "Starlight Gauntlets", Phase: 1, Quality: ItemQualityRare, SourceZone: "N UB - Hungarfen", SourceDrop: "", Stats: Stats{StatIntellect: 21, StatStamina: 10, StatSpellPower: 25}, GemSockets: []GemColor{0x3, 0x3}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 27537, ItemType: ItemTypeHands, Name: "Gloves of Oblivion", Phase: 1, Quality: ItemQualityRare, SourceZone: "SH - Kargath", SourceDrop: "", Stats: Stats{StatIntellect: 21, StatStamina: 33, StatSpellHit: 20, StatSpellPower: 26}},
	{ID: 29784, ItemType: ItemTypeHands, Name: "Harmony's Touch", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "Building a Perimeter - Netherstorm Quest", SourceDrop: "", Stats: Stats{StatStamina: 18, StatSpellCrit: 16, StatSpellPower: 33}},
	{ID: 27743, ItemType: ItemTypeWaist, Name: "Girdle of Living Flame", Phase: 1, Quality: ItemQualityRare, SourceZone: "H UB - Hungarfen", SourceDrop: "", Stats: Stats{StatIntellect: 17, StatStamina: 15, StatSpellHit: 16, StatSpellPower: 29}, GemSockets: []GemColor{0x4, 0x3}, SocketBonus: Stats{StatSpellCrit: 3}},
	{ID: 29244, ItemType: ItemTypeWaist, Name: "Wave-Song Girdle", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H AC - Exarch Maladaar", SourceDrop: "", Stats: Stats{StatIntellect: 25, StatStamina: 25, StatSpellCrit: 23, StatSpellPower: 32}},
	{ID: 31461, ItemType: ItemTypeWaist, Name: "A'dal's Gift", Phase: 1, Quality: ItemQualityRare, SourceZone: "How to Break Into the Arcatraz - Quest", SourceDrop: "", Stats: Stats{StatIntellect: 25, StatSpellCrit: 21, StatSpellPower: 34}},
	{ID: 29257, ItemType: ItemTypeWaist, Name: "Sash of Arcane Visions", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H AC - Exarch Maladaar", SourceDrop: "", Stats: Stats{StatIntellect: 23, StatStamina: 18, StatSpellCrit: 22, StatSpellPower: 28}},
	{ID: 29241, ItemType: ItemTypeWaist, Name: "Belt of Depravity", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H Arc - Harbinger Skyriss", SourceDrop: "", Stats: Stats{StatIntellect: 27, StatStamina: 31, StatSpellHit: 17, StatSpellPower: 34}},
	{ID: 27783, ItemType: ItemTypeWaist, Name: "Moonrage Girdle", Phase: 1, Quality: ItemQualityRare, SourceZone: "SV - Hydromancer Thespia", SourceDrop: "", Stats: Stats{StatIntellect: 22, StatSpellCrit: 20, StatSpellPower: 25}},
	{ID: 27795, ItemType: ItemTypeWaist, Name: "Sash of Serpentra", Phase: 1, Quality: ItemQualityRare, SourceZone: "SV - Warlord Kalithresh", SourceDrop: "", Stats: Stats{StatIntellect: 21, StatStamina: 31, StatSpellHit: 17, StatSpellPower: 25}},
	{ID: 31513, ItemType: ItemTypeWaist, Name: "Blackwhelp Belt", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "Whelps of the Wyrmcult - BEM Quest", SourceDrop: "", Stats: Stats{StatIntellect: 11, StatSpellCrit: 10, StatSpellPower: 32}},
	{ID: 24262, ItemType: ItemTypeLegs, Name: "Spellstrike Pants", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Tailoring BoE", SourceDrop: "", Stats: Stats{StatIntellect: 8, StatStamina: 12, StatSpellCrit: 26, StatSpellHit: 22, StatSpellPower: 46}, GemSockets: []GemColor{0x2, 0x4, 0x3}, SocketBonus: Stats{StatStamina: 6}},
	{ID: 30541, ItemType: ItemTypeLegs, Name: "Stormsong Kilt", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H UB - The Black Stalker", SourceDrop: "", Stats: Stats{StatIntellect: 30, StatStamina: 25, StatSpellCrit: 26, StatSpellPower: 35}, GemSockets: []GemColor{0x2, 0x4, 0x3}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 29141, ItemType: ItemTypeLegs, Name: "Tempest Leggings", Phase: 1, Quality: ItemQualityRare, SourceZone: "The Mag'har - Revered (Horde)", SourceDrop: "", Stats: Stats{StatIntellect: 11, StatSpellCrit: 18, StatSpellPower: 44}, GemSockets: []GemColor{0x2, 0x4, 0x4}, SocketBonus: Stats{StatMP5: 2}},
	{ID: 29142, ItemType: ItemTypeLegs, Name: "Kurenai Kilt", Phase: 1, Quality: ItemQualityRare, SourceZone: "Kurenai - Revered (Ally)", SourceDrop: "", Stats: Stats{StatIntellect: 11, StatSpellCrit: 18, StatSpellPower: 44}, GemSockets: []GemColor{0x2, 0x4, 0x4}, SocketBonus: Stats{StatMP5: 2}},
	{ID: 30531, ItemType: ItemTypeLegs, Name: "Breeches of the Occultist", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H BM - Aeonus", SourceDrop: "", Stats: Stats{StatIntellect: 22, StatStamina: 37, StatSpellCrit: 23, StatSpellPower: 26}, GemSockets: []GemColor{0x4, 0x4, 0x3}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 30709, ItemType: ItemTypeLegs, Name: "Pantaloons of Flaming Wrath", Phase: 1, Quality: ItemQualityRare, SourceZone: "H SH - Blood Guard Porung", SourceDrop: "", Stats: Stats{StatIntellect: 28, StatSpellCrit: 42, StatSpellPower: 33}},
	{ID: 27492, ItemType: ItemTypeLegs, Name: "Moonchild Leggings", Phase: 1, Quality: ItemQualityRare, SourceZone: "H BF - Broggok", SourceDrop: "", Stats: Stats{StatIntellect: 20, StatStamina: 26, StatSpellCrit: 21, StatSpellPower: 23}, GemSockets: []GemColor{0x2, 0x4, 0x4}, SocketBonus: Stats{StatMP5: 2}},
	{ID: 29343, ItemType: ItemTypeLegs, Name: "Haramad's Leggings of the Third Coin", Phase: 1, Quality: ItemQualityRare, SourceZone: "Undercutting the Competition - MT Quest", SourceDrop: "", Stats: Stats{StatIntellect: 29, StatSpellCrit: 16, StatSpellPower: 27}, GemSockets: []GemColor{0x2, 0x4, 0x4}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 27472, ItemType: ItemTypeLegs, Name: "Gladiator's Mail Leggings", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Arena Season 1 Reward", SourceDrop: "", Stats: Stats{StatIntellect: 25, StatStamina: 54, StatSpellCrit: 22, StatSpellPower: 42, StatMP5: 6}},
	{ID: 30532, ItemType: ItemTypeLegs, Name: "Kirin Tor Master's Trousers", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H SLabs - Murmur", SourceDrop: "", Stats: Stats{StatIntellect: 29, StatStamina: 27, StatSpellPower: 36}, GemSockets: []GemColor{0x2, 0x4, 0x3}, SocketBonus: Stats{StatSpellHit: 4}},
	{ID: 28185, ItemType: ItemTypeLegs, Name: "Khadgar's Kilt of Abjuration", Phase: 1, Quality: ItemQualityRare, SourceZone: "BM - Temporus", SourceDrop: "", Stats: Stats{StatIntellect: 22, StatStamina: 20, StatSpellPower: 36}, GemSockets: []GemColor{0x4, 0x3, 0x3}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 27838, ItemType: ItemTypeLegs, Name: "Incanter's Trousers", Phase: 1, Quality: ItemQualityRare, SourceZone: "SH - Talon King Ikiss", SourceDrop: "", Stats: Stats{StatIntellect: 30, StatStamina: 25, StatSpellCrit: 18, StatSpellPower: 42}},
	{ID: 27907, ItemType: ItemTypeLegs, Name: "Mana-Etched Pantaloons", Phase: 1, Quality: ItemQualityRare, SourceZone: "H UB - The Black Stalker", SourceDrop: "", Stats: Stats{StatIntellect: 32, StatStamina: 34, StatSpellCrit: 21, StatSpellPower: 33}},
	{ID: 27909, ItemType: ItemTypeLegs, Name: "Tidefury Kilt", Phase: 1, Quality: ItemQualityRare, SourceZone: "SLabs - Murmur", SourceDrop: "", Stats: Stats{StatIntellect: 31, StatStamina: 39, StatSpellCrit: 19, StatSpellPower: 35}},
	{ID: 28266, ItemType: ItemTypeLegs, Name: "Molten Earth Kilt", Phase: 1, Quality: ItemQualityRare, SourceZone: "Mech - Pathaleon the Calculator", SourceDrop: "", Stats: Stats{StatIntellect: 32, StatStamina: 24, StatSpellPower: 40, StatMP5: 10}},
	{ID: 27948, ItemType: ItemTypeLegs, Name: "Trousers of Oblivion", Phase: 1, Quality: ItemQualityRare, SourceZone: "SH - Talon King Ikiss", SourceDrop: "", Stats: Stats{StatIntellect: 33, StatStamina: 42, StatSpellHit: 12, StatSpellPower: 39}},
	{ID: 29314, ItemType: ItemTypeLegs, Name: "Leggings of the Third Coin", Phase: 1, Quality: ItemQualityRare, SourceZone: "Levixus the Soul Caller - Auchindoun Quest", SourceDrop: "", Stats: Stats{StatIntellect: 26, StatStamina: 34, StatSpellCrit: 12, StatSpellPower: 32, StatMP5: 4}},
	{ID: 28406, ItemType: ItemTypeFeet, Name: "Sigil-Laced Boots", Phase: 1, Quality: ItemQualityRare, SourceZone: "Arc - Harbinger Skyriss", SourceDrop: "", Stats: Stats{StatIntellect: 18, StatStamina: 24, StatSpellCrit: 17, StatSpellPower: 20}, GemSockets: []GemColor{0x2, 0x4}, SocketBonus: Stats{StatIntellect: 3}},
	{ID: 28640, ItemType: ItemTypeFeet, Name: "General's Mail Sabatons", Phase: 1, Quality: ItemQualityEpic, SourceZone: "11,424 Honor & 40 EotS Marks", SourceDrop: "", Stats: Stats{StatIntellect: 23, StatStamina: 34, StatSpellCrit: 24, StatSpellPower: 28}},
	{ID: 27914, ItemType: ItemTypeFeet, Name: "Moonstrider Boots", Phase: 1, Quality: ItemQualityRare, SourceZone: "SH - Darkweaver Syth", SourceDrop: "", Stats: Stats{StatIntellect: 22, StatStamina: 21, StatSpellCrit: 20, StatSpellPower: 25, StatMP5: 6}},
	{ID: 28179, ItemType: ItemTypeFeet, Name: "Shattrath Jumpers", Phase: 1, Quality: ItemQualityRare, SourceZone: "Into the Heart of the Labyrinth - Auch. Quest", SourceDrop: "", Stats: Stats{StatIntellect: 17, StatStamina: 25, StatSpellPower: 29}, GemSockets: []GemColor{0x4, 0x3}, SocketBonus: Stats{StatIntellect: 3}},
	{ID: 29245, ItemType: ItemTypeFeet, Name: "Wave-Crest Striders", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H BF - Keli'dan the Breaker", SourceDrop: "", Stats: Stats{StatIntellect: 26, StatStamina: 28, StatSpellPower: 33, StatMP5: 8}},
	{ID: 27821, ItemType: ItemTypeFeet, Name: "Extravagant Boots of Malice", Phase: 1, Quality: ItemQualityRare, SourceZone: "H MT - Tavarok", SourceDrop: "", Stats: Stats{StatIntellect: 24, StatStamina: 27, StatSpellHit: 14, StatSpellPower: 30}},
	{ID: 27845, ItemType: ItemTypeFeet, Name: "Magma Plume Boots", Phase: 1, Quality: ItemQualityRare, SourceZone: "H AC - Shirrak the Dead Watcher", SourceDrop: "", Stats: Stats{StatIntellect: 26, StatStamina: 24, StatSpellHit: 14, StatSpellPower: 29}},
	{ID: 29808, ItemType: ItemTypeFeet, Name: "Shimmering Azure Boots", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "Securing the Celestial Ridge - NS Quest", SourceDrop: "", Stats: Stats{StatIntellect: 19, StatSpellHit: 16, StatSpellPower: 23, StatMP5: 5}},
	{ID: 29242, ItemType: ItemTypeFeet, Name: "Boots of Blasphemy", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H SP - Quagmirran", SourceDrop: "", Stats: Stats{StatIntellect: 29, StatStamina: 36, StatSpellPower: 36}},
	{ID: 29258, ItemType: ItemTypeFeet, Name: "Boots of Ethereal Manipulation", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H Bot - Warp Splinter", SourceDrop: "", Stats: Stats{StatIntellect: 27, StatStamina: 27, StatSpellPower: 33}},
	{ID: 29313, ItemType: ItemTypeFeet, Name: "Earthbreaker's Greaves", Phase: 1, Quality: ItemQualityRare, SourceZone: "Levixus the Soul Caller - Auchindoun Quest", SourceDrop: "", Stats: Stats{StatIntellect: 20, StatStamina: 27, StatSpellCrit: 8, StatSpellPower: 25, StatMP5: 3}},
	{ID: 30519, ItemType: ItemTypeFeet, Name: "Boots of the Nexus Warden", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "The Flesh Lies... - Netherstorm Quest", SourceDrop: "", Stats: Stats{StatIntellect: 17, StatStamina: 27, StatSpellHit: 18, StatSpellPower: 21}},
	{ID: 28227, ItemType: ItemTypeFinger, Name: "Sparking Arcanite Ring", Phase: 1, Quality: ItemQualityRare, SourceZone: "H OHF - Epoch Hunter", SourceDrop: "", Stats: Stats{StatIntellect: 14, StatStamina: 13, StatSpellCrit: 14, StatSpellHit: 10, StatSpellPower: 22}},
	{ID: 29126, ItemType: ItemTypeFinger, Name: "Seer's Signet", Phase: 1, Quality: ItemQualityEpic, SourceZone: "The Scryers - Exalted", SourceDrop: "", Stats: Stats{StatStamina: 24, StatSpellCrit: 12, StatSpellPower: 34}},
	{ID: 31922, ItemType: ItemTypeFinger, Name: "Ring of Conflict Survival", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H MT - Yor (Summoned Boss)", SourceDrop: "", Stats: Stats{StatStamina: 28, StatSpellCrit: 20, StatSpellPower: 23}},
	{ID: 28394, ItemType: ItemTypeFinger, Name: "Ryngo's Band of Ingenuity", Phase: 1, Quality: ItemQualityRare, SourceZone: "Arc - Wrath-Scryer Soccothrates", SourceDrop: "", Stats: Stats{StatIntellect: 14, StatStamina: 12, StatSpellCrit: 14, StatSpellPower: 25}},
	{ID: 29320, ItemType: ItemTypeFinger, Name: "Band of the Guardian", Phase: 1, Quality: ItemQualityRare, SourceZone: "Hero of the Brood - CoT Quest", SourceDrop: "", Stats: Stats{StatIntellect: 11, StatSpellCrit: 17, StatSpellPower: 23}},
	{ID: 27784, ItemType: ItemTypeFinger, Name: "Scintillating Coral Band", Phase: 1, Quality: ItemQualityRare, SourceZone: "SV - Hydromancer Thespia", SourceDrop: "", Stats: Stats{StatIntellect: 15, StatStamina: 14, StatSpellCrit: 17, StatSpellPower: 21}},
	{ID: 30366, ItemType: ItemTypeFinger, Name: "Manastorm Band", Phase: 1, Quality: ItemQualityRare, SourceZone: "Shutting Down Manaforge Ara - Quest", SourceDrop: "", Stats: Stats{StatIntellect: 15, StatSpellCrit: 10, StatSpellPower: 29}},
	{ID: 29172, ItemType: ItemTypeFinger, Name: "Ashyen's Gift", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Cenarion Expedition - Exalted", SourceDrop: "", Stats: Stats{StatStamina: 30, StatSpellHit: 21, StatSpellPower: 23}},
	{ID: 29352, ItemType: ItemTypeFinger, Name: "Cobalt Band of Tyrigosa", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H MT - Nexus-Prince Shaffar", SourceDrop: "", Stats: Stats{StatIntellect: 17, StatStamina: 19, StatSpellPower: 35}},
	{ID: 28555, ItemType: ItemTypeFinger, Name: "Seal of the Exorcist", Phase: 1, Quality: ItemQualityEpic, SourceZone: "50 Spirit Shards ", SourceDrop: "", Stats: Stats{StatStamina: 24, StatSpellHit: 12, StatSpellPower: 28}},
	{ID: 31339, ItemType: ItemTypeFinger, Name: "Lola's Eve", Phase: 1, Quality: ItemQualityEpic, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{StatIntellect: 14, StatStamina: 15, StatSpellPower: 29}},
	{ID: 31921, ItemType: ItemTypeFinger, Name: "Yor's Collapsing Band", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H MT - Yor (Summoned Boss)", SourceDrop: "", Stats: Stats{StatIntellect: 20, StatSpellPower: 23}},

	// FUTURE: just to get this working, shaman "newcast" checks for totems and spell being cast and applies totem bonuses.
	{ID: 28248, ItemType: ItemTypeRanged, RangedWeaponType: RangedWeaponTypeTotem, Name: "Totem of the Void", Phase: 1, Quality: ItemQualityRare, SourceZone: "Mech - Cache of the Legion", SourceDrop: ""},
	{ID: 23199, ItemType: ItemTypeRanged, RangedWeaponType: RangedWeaponTypeTotem, Name: "Totem of the Storm", Phase: 0, Quality: ItemQualityRare, SourceZone: "Boe World Drop", SourceDrop: ""},
	{ID: 32330, ItemType: ItemTypeRanged, RangedWeaponType: RangedWeaponTypeTotem, Name: "Totem of Ancestral Guidance", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: ""},

	{ID: 27543, ItemType: ItemTypeWeapon, Name: "Starlight Dagger", Phase: 1, Quality: ItemQualityRare, SourceZone: "H SP - Mennu the Betrayer", SourceDrop: "", Stats: Stats{StatIntellect: 15, StatStamina: 15, StatSpellHit: 16, StatSpellPower: 121}},
	{ID: 27868, ItemType: ItemTypeWeapon, Name: "Runesong Dagger", Phase: 1, Quality: ItemQualityRare, SourceZone: "SH - Warbringer O'mrogg", SourceDrop: "", Stats: Stats{StatIntellect: 11, StatStamina: 12, StatSpellCrit: 20, StatSpellPower: 121}},
	{ID: 27741, ItemType: ItemTypeWeapon, Name: "Bleeding Hollow Warhammer", Phase: 1, Quality: ItemQualityRare, SourceZone: "H SP - Quagmirran", SourceDrop: "", Stats: Stats{StatIntellect: 17, StatStamina: 12, StatSpellCrit: 16, StatSpellPower: 121}},
	{ID: 27937, ItemType: ItemTypeWeapon, Name: "Sky Breaker", Phase: 1, Quality: ItemQualityRare, SourceZone: "H AC - Avatar of the Martyred", SourceDrop: "", Stats: Stats{StatIntellect: 20, StatStamina: 13, StatSpellPower: 132}},
	{ID: 28412, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, Name: "Lamp of Peaceful Radiance", Phase: 1, Quality: ItemQualityRare, SourceZone: "Arc - Harbinger Skyriss", SourceDrop: "", Stats: Stats{StatIntellect: 14, StatStamina: 13, StatSpellCrit: 13, StatSpellHit: 12, StatSpellPower: 21}},
	{ID: 28260, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, Name: "Manual of the Nethermancer", Phase: 1, Quality: ItemQualityRare, SourceZone: "Mech - Nethermancer Sepethrea", SourceDrop: "", Stats: Stats{StatIntellect: 15, StatStamina: 12, StatSpellCrit: 19, StatSpellPower: 21}},
	{ID: 31287, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, WeaponType: WeaponTypeShield, Name: "Draenei Honor Guard Shield", Phase: 1, Quality: ItemQualityRare, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{StatIntellect: 16, StatSpellCrit: 21, StatSpellPower: 19}},
	{ID: 28187, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, Name: "Star-Heart Lamp", Phase: 1, Quality: ItemQualityRare, SourceZone: "BM - Temporus", SourceDrop: "", Stats: Stats{StatIntellect: 18, StatStamina: 17, StatSpellHit: 12, StatSpellPower: 22}},
	{ID: 29330, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, Name: "The Saga of Terokk", Phase: 1, Quality: ItemQualityRare, SourceZone: "Terokk's Legacy - Auchindoun Quest", SourceDrop: "", Stats: Stats{StatIntellect: 23, StatSpellPower: 28}},
	{ID: 27910, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, WeaponType: WeaponTypeShield, Name: "Silvermoon Crest Shield", Phase: 1, Quality: ItemQualityRare, SourceZone: "SLabs - Murmur", SourceDrop: "", Stats: Stats{StatIntellect: 20, StatSpellPower: 23, StatMP5: 5}},
	{ID: 30984, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, WeaponType: WeaponTypeShield, Name: "Spellbreaker's Buckler", Phase: 1, Quality: ItemQualityRare, SourceZone: "Akama's Promise - SMV Quest", SourceDrop: "", Stats: Stats{StatIntellect: 10, StatStamina: 22, StatSpellPower: 29}},
	{ID: 27534, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, Name: "Hortus' Seal of Brilliance", Phase: 1, Quality: ItemQualityRare, SourceZone: "SH - Warchief Kargath Bladefist", SourceDrop: "", Stats: Stats{StatIntellect: 20, StatStamina: 18, StatSpellPower: 23}},
	{ID: 29355, ItemType: ItemTypeWeapon, HandType: HandTypeTwoHand, Name: "Terokk's Shadowstaff", Phase: 1, Quality: ItemQualityEpic, SourceZone: "H SH - Talon King Ikiss", SourceDrop: "", Stats: Stats{StatIntellect: 42, StatStamina: 40, StatSpellCrit: 37, StatSpellPower: 168}},
	{ID: 29130, ItemType: ItemTypeWeapon, HandType: HandTypeTwoHand, Name: "Auchenai Staff", Phase: 1, Quality: ItemQualityRare, SourceZone: "The Aldor - Revered", SourceDrop: "", Stats: Stats{StatIntellect: 46, StatSpellCrit: 26, StatSpellHit: 19, StatSpellPower: 121}},
	{ID: 28341, ItemType: ItemTypeWeapon, HandType: HandTypeTwoHand, Name: "Warpstaff of Arcanum", Phase: 1, Quality: ItemQualityRare, SourceZone: "Bot - Warp Splinter", SourceDrop: "", Stats: Stats{StatIntellect: 38, StatStamina: 37, StatSpellCrit: 26, StatSpellHit: 16, StatSpellPower: 121}},
	{ID: 31308, ItemType: ItemTypeWeapon, HandType: HandTypeTwoHand, Name: "The Bringer of Death", Phase: 1, Quality: ItemQualityRare, SourceZone: "BoE World Drop", SourceDrop: "", Stats: Stats{StatIntellect: 31, StatStamina: 32, StatSpellCrit: 42, StatSpellPower: 121}},
	{ID: 28188, ItemType: ItemTypeWeapon, HandType: HandTypeTwoHand, Name: "Bloodfire Greatstaff", Phase: 1, Quality: ItemQualityRare, SourceZone: "BM - Aeonus", SourceDrop: "", Stats: Stats{StatIntellect: 42, StatStamina: 42, StatSpellCrit: 28, StatSpellPower: 121}},
	{ID: 30011, ItemType: ItemTypeWeapon, HandType: HandTypeTwoHand, Name: "Ameer's Impulse Taser", Phase: 1, Quality: ItemQualityRare, SourceZone: "Nexus-King Salhadaar - Netherstorm Quest", SourceDrop: "", Stats: Stats{StatIntellect: 27, StatStamina: 27, StatSpellCrit: 27, StatSpellHit: 17, StatSpellPower: 103}},
	{ID: 27842, ItemType: ItemTypeWeapon, HandType: HandTypeTwoHand, Name: "Grand Scepter of the Nexus-Kings", Phase: 1, Quality: ItemQualityRare, SourceZone: "H MT - Nexus-Prince Shaffar", SourceDrop: "", Stats: Stats{StatIntellect: 43, StatStamina: 45, StatSpellHit: 19, StatSpellPower: 121}},

	{ID: 28346, ItemType: ItemTypeWeapon, Name: "Gladiator's Endgame", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Arena Season 1 Reward", SourceDrop: "", Stats: Stats{StatIntellect: 14, StatStamina: 21, StatSpellPower: 19}},
	{ID: 24557, ItemType: ItemTypeWeapon, HandType: HandTypeTwoHand, Name: "Gladiator's War Staff", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Arena Season 1 Reward", SourceDrop: "", Stats: Stats{StatIntellect: 35, StatStamina: 48, StatSpellCrit: 36, StatSpellHit: 21, StatSpellPower: 199}},

	{ID: 29389, ItemType: ItemTypeRanged, RangedWeaponType: RangedWeaponTypeTotem, Name: "Totem of the Pulsing Earth", Phase: 1, Quality: ItemQualityEpic, SourceZone: "15 Badge of Justice - G'eras", SourceDrop: "", Activate: ActivateTotemOfPulsingEarth, ActivateCD: NeverExpires},
	// {ItemType: ItemTypeRanged, Name: "Totem of Impact", Phase: 1, Quality: ItemQualityRare, SourceZone: "15 Mark of Thrallmar/ Honor Hold", SourceDrop: "", Stats: Stats{     StatMP5:0}},
	// {ItemType: ItemTypeRanged, Name: "Totem of Lightning", Phase: 1, Quality: ItemQualityRare, SourceZone: "Colossal Menace - HFP Quest", SourceDrop: "", Stats: Stats{     StatMP5:0}},

	// source: https://docs.google.com/spreadsheets/d/1T4DEuq0yroEPb-11okC3qjj7aYfCGu2e6nT9LeT30zg/edit#gid=0
	{ID: 28744, ItemType: ItemTypeHead, Name: "Uni-Mind Headdress", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Netherspite", Stats: Stats{StatStamina: 31, StatIntellect: 40, StatSpellPower: 46, StatSpellCrit: 25, StatSpellHit: 19}},
	{ID: 28586, ItemType: ItemTypeHead, Name: "Wicked Witch's Hat", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Opera", Stats: Stats{StatStamina: 37, StatIntellect: 38, StatSpellPower: 43, StatSpellCrit: 32}},
	{ID: 29035, ItemType: ItemTypeHead, Name: "Cyclone Faceguard (Tier 4)", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Prince", Stats: Stats{StatStamina: 30, StatIntellect: 31, StatSpellPower: 39, StatSpellCrit: 25, StatMP5: 8}, GemSockets: []GemColor{GemColorMeta, GemColorYellow}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 30171, ItemType: ItemTypeHead, Name: "Cataclysm Headpiece (Tier 5)", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "Lady Vashj", Stats: Stats{StatStamina: 35, StatIntellect: 28, StatSpellPower: 54, StatSpellCrit: 26, StatSpellHit: 18, StatMP5: 7}, GemSockets: []GemColor{GemColorMeta, GemColorYellow}, SocketBonus: Stats{StatSpellHit: 5}},
	{ID: 29986, ItemType: ItemTypeHead, Name: "Cowl of the Grand Engineer", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Void Reaver", Stats: Stats{StatStamina: 22, StatIntellect: 27, StatSpellPower: 53, StatSpellCrit: 35, StatSpellHit: 16}, GemSockets: []GemColor{GemColorYellow, GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 32476, ItemType: ItemTypeHead, Name: "Gadgetstorm Goggles", Phase: 2, Quality: ItemQualityEpic, SourceZone: "Crafted (Patch 2.1)", SourceDrop: "Engineering (Mail)", Stats: Stats{StatStamina: 28, StatSpellPower: 55, StatSpellCrit: 40, StatSpellHit: 12}, GemSockets: []GemColor{GemColorMeta, GemColorBlue}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 31014, ItemType: ItemTypeHead, Name: "Skyshatter Headguard (Tier 6)", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Archimonde", Stats: Stats{StatStamina: 42, StatIntellect: 37, StatSpellPower: 62, StatSpellCrit: 36, StatMP5: 8}, GemSockets: []GemColor{GemColorMeta, GemColorBlue}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 32525, ItemType: ItemTypeHead, Name: "Cowl of the Illidari High Lord", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Illidan", Stats: Stats{StatStamina: 33, StatIntellect: 31, StatSpellPower: 64, StatSpellCrit: 47, StatSpellHit: 21}, GemSockets: []GemColor{GemColorMeta, GemColorBlue}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 28530, ItemType: ItemTypeNeck, Name: "Brooch of Unquenchable Fury", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Moroes", Stats: Stats{StatStamina: 24, StatIntellect: 21, StatSpellPower: 26, StatSpellHit: 15}},
	{ID: 29368, ItemType: ItemTypeNeck, Name: "Manasurge Pendant", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Shattrah", SourceDrop: "Badges", Stats: Stats{StatStamina: 24, StatIntellect: 22, StatSpellPower: 28}},
	{ID: 30008, ItemType: ItemTypeNeck, Name: "Pendant of the Lost Ages", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "Tidewalker", Stats: Stats{StatStamina: 27, StatIntellect: 17, StatSpellPower: 36}},
	{ID: 28762, ItemType: ItemTypeNeck, Name: "Adornment of Stolen Souls", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Prince", Stats: Stats{StatStamina: 18, StatIntellect: 20, StatSpellPower: 28, StatSpellCrit: 23}},
	{ID: 30015, ItemType: ItemTypeNeck, Name: "The Sun King's Talisman", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Kael Reward", Stats: Stats{StatStamina: 22, StatIntellect: 16, StatSpellPower: 41, StatSpellCrit: 24}},
	{ID: 32349, ItemType: ItemTypeNeck, Name: "Translucent Spellthread Necklace", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "RoS", Stats: Stats{StatSpellPower: 46, StatSpellCrit: 24, StatSpellHit: 15}},
	{ID: 28726, ItemType: ItemTypeShoulder, Name: "Mantle of the Mind Flayer", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Aran", Stats: Stats{StatStamina: 33, StatIntellect: 29, StatSpellPower: 35}},
	{ID: 30024, ItemType: ItemTypeShoulder, Name: "Mantle of the Elven Kings", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Trash", Stats: Stats{StatStamina: 27, StatIntellect: 18, StatSpellPower: 39, StatSpellCrit: 25, StatSpellHit: 18}},
	{ID: 29037, ItemType: ItemTypeShoulder, Name: "Cyclone Shoulderguards (Tier 4)", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Gruul's Lair", SourceDrop: "Maulgar", Stats: Stats{StatStamina: 28, StatIntellect: 26, StatSpellPower: 36, StatSpellCrit: 12}, GemSockets: []GemColor{GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatSpellPower: 4}},
	{ID: 30079, ItemType: ItemTypeShoulder, Name: "Illidari Shoulderpads", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "Tidewalker", Stats: Stats{StatStamina: 34, StatIntellect: 23, StatSpellPower: 39, StatSpellCrit: 16}, GemSockets: []GemColor{GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatSpellPower: 4}},
	{ID: 32338, ItemType: ItemTypeShoulder, Name: "Blood-cursed Shoulderpads", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Bloodboil", Stats: Stats{StatStamina: 25, StatIntellect: 19, StatSpellPower: 55, StatSpellCrit: 25, StatSpellHit: 18}},
	{ID: 30173, ItemType: ItemTypeShoulder, Name: "Cataclysm Shoulderpads (Tier 5)", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "VoidReaver", Stats: Stats{StatStamina: 26, StatIntellect: 19, StatSpellPower: 41, StatSpellCrit: 24, StatMP5: 6}, GemSockets: []GemColor{GemColorBlue, GemColorYellow}, SocketBonus: Stats{StatSpellCrit: 3}},
	{ID: 32587, ItemType: ItemTypeShoulder, Name: "Mantle of Nimble Thought", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Tailoring", Stats: Stats{StatStamina: 37, StatIntellect: 26, StatSpellPower: 44, StatSpellHaste: 38}},
	{ID: 31023, ItemType: ItemTypeShoulder, Name: "Skyshatter Mantle (Tier 6)", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Mother", Stats: Stats{StatStamina: 30, StatIntellect: 31, StatSpellPower: 46, StatSpellCrit: 27, StatSpellHit: 11, StatMP5: 4}, GemSockets: []GemColor{GemColorBlue, GemColorYellow}, SocketBonus: Stats{StatSpellPower: 4}},
	{ID: 30884, ItemType: ItemTypeShoulder, Name: "Hatefury Mantle", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Anetheron", Stats: Stats{StatStamina: 15, StatIntellect: 18, StatSpellPower: 55, StatSpellCrit: 24}, GemSockets: []GemColor{GemColorBlue, GemColorYellow}, SocketBonus: Stats{StatSpellCrit: 3}},
	{ID: 28766, ItemType: ItemTypeBack, Name: "Ruby Drape of the Mysticant", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Prince", Stats: Stats{StatStamina: 22, StatIntellect: 21, StatSpellPower: 30, StatSpellHit: 18}},
	{ID: 28570, ItemType: ItemTypeBack, Name: "Shadow-Cloak of Dalaran", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Moroes", Stats: Stats{StatStamina: 19, StatIntellect: 18, StatSpellPower: 36}},
	{ID: 29369, ItemType: ItemTypeBack, Name: "Shawl of Shifting Probabilities", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Shattrah", SourceDrop: "Badges", Stats: Stats{StatStamina: 18, StatIntellect: 16, StatSpellPower: 21, StatSpellCrit: 22}},
	{ID: 29992, ItemType: ItemTypeBack, Name: "Royal Cloak of the Sunstriders", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Kaelthas", Stats: Stats{StatStamina: 27, StatIntellect: 22, StatSpellPower: 44}},
	{ID: 28797, ItemType: ItemTypeBack, Name: "Brute Cloak of the Ogre-Magi", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Gruul's Lair", SourceDrop: "Maulgar", Stats: Stats{StatStamina: 18, StatIntellect: 20, StatSpellPower: 28, StatSpellCrit: 23}},
	{ID: 30735, ItemType: ItemTypeBack, Name: "Ancient Spellcloak of the Highborne", Phase: 1, Quality: ItemQualityEpic, SourceZone: "WorldBoss", SourceDrop: "Kazzak", Stats: Stats{StatIntellect: 15, StatSpellPower: 36, StatSpellCrit: 19}},
	{ID: 32331, ItemType: ItemTypeBack, Name: "Cloak of the Illidari Council", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Illidari Council", Stats: Stats{StatStamina: 24, StatIntellect: 16, StatSpellPower: 42, StatSpellCrit: 25}},
	{ID: 29033, ItemType: ItemTypeChest, Name: "Cyclone Chestguard (Tier 4)", Phase: 1, Quality: ItemQualityEpic, SourceZone: "GruulsLair", SourceDrop: "Maulgar", Stats: Stats{StatStamina: 33, StatIntellect: 32, StatSpellPower: 39, StatSpellCrit: 20, StatMP5: 8}, GemSockets: []GemColor{GemColorRed, GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellHit: 4}},
	{ID: 29519, ItemType: ItemTypeChest, Name: "Netherstrike Breastplate", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Crafted", SourceDrop: "Leatherworking", Stats: Stats{StatStamina: 34, StatIntellect: 23, StatSpellPower: 37, StatSpellCrit: 32, StatMP5: 8}, GemSockets: []GemColor{GemColorBlue, GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 30056, ItemType: ItemTypeChest, Name: "Robe of Hateful Echoes", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "Hydross", Stats: Stats{StatStamina: 34, StatIntellect: 36, StatSpellPower: 50, StatSpellCrit: 25}, GemSockets: []GemColor{GemColorRed, GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatStamina: 6}},
	{ID: 32327, ItemType: ItemTypeChest, Name: "Robe of the Shadow Council", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Teron", Stats: Stats{StatStamina: 37, StatIntellect: 36, StatSpellPower: 73, StatSpellCrit: 28}},
	{ID: 30913, ItemType: ItemTypeChest, Name: "Robes of Rhonin", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Archimonde", Stats: Stats{StatStamina: 55, StatIntellect: 38, StatSpellPower: 81, StatSpellCrit: 24, StatSpellHit: 27}},
	{ID: 30169, ItemType: ItemTypeChest, Name: "Cataclysm Chestpiece (Tier 5)", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Kaelthas", Stats: Stats{StatStamina: 37, StatIntellect: 28, StatSpellPower: 55, StatSpellCrit: 24, StatMP5: 10}, GemSockets: []GemColor{GemColorBlue, GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 30107, ItemType: ItemTypeChest, Name: "Vestments of the Sea-Witch", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "LadyVashj", Stats: Stats{StatStamina: 28, StatIntellect: 28, StatSpellPower: 57, StatSpellCrit: 31, StatSpellHit: 27}, GemSockets: []GemColor{GemColorYellow, GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 32592, ItemType: ItemTypeChest, Name: "Chestguard of Relentless Storms", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Trash", Stats: Stats{StatStamina: 36, StatIntellect: 30, StatSpellPower: 74, StatSpellCrit: 46}},
	{ID: 31017, ItemType: ItemTypeChest, Name: "Skyshatter Breastplate (Tier 6)", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Illidan", Stats: Stats{StatStamina: 42, StatIntellect: 41, StatSpellPower: 62, StatSpellCrit: 27, StatSpellHit: 17, StatMP5: 7}, GemSockets: []GemColor{GemColorBlue, GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 28515, ItemType: ItemTypeWrist, Name: "Bands of Nefarious Deeds", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Maiden", Stats: Stats{StatStamina: 27, StatIntellect: 22, StatSpellPower: 32}},
	{ID: 32351, ItemType: ItemTypeWrist, Name: "Elunite Empowered Bracers", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "RoS", Stats: Stats{StatStamina: 27, StatIntellect: 22, StatSpellPower: 34, StatSpellHit: 19, StatMP5: 6}},
	{ID: 32270, ItemType: ItemTypeWrist, Name: "Focused Mana Bindings", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Akama", Stats: Stats{StatStamina: 27, StatIntellect: 20, StatSpellPower: 42, StatSpellHit: 19}},
	{ID: 29521, ItemType: ItemTypeWrist, Name: "Netherstrike Bracers", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Crafted", SourceDrop: "Leatherworking", Stats: Stats{StatStamina: 13, StatIntellect: 13, StatSpellPower: 20, StatSpellCrit: 17, StatMP5: 6}, GemSockets: []GemColor{GemColorYellow}, SocketBonus: Stats{StatSpellPower: 2}},
	{ID: 32259, ItemType: ItemTypeWrist, Name: "Bands of the Coming Storm", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Supremus", Stats: Stats{StatStamina: 28, StatIntellect: 28, StatSpellPower: 34, StatSpellCrit: 21}},
	{ID: 29918, ItemType: ItemTypeWrist, Name: "Mindstorm Wristbands", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Alar", Stats: Stats{StatStamina: 13, StatIntellect: 13, StatSpellPower: 36, StatSpellCrit: 23}},
	{ID: 30870, ItemType: ItemTypeWrist, Name: "Cuffs of Devastation", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Winterchill", Stats: Stats{StatStamina: 22, StatIntellect: 20, StatSpellPower: 34, StatSpellCrit: 14}, GemSockets: []GemColor{GemColorYellow}, SocketBonus: Stats{StatStamina: 3}},
	{ID: 29034, ItemType: ItemTypeHands, Name: "Cyclone Handguards (Tier 4)", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Curator", Stats: Stats{StatStamina: 26, StatIntellect: 29, StatSpellPower: 34, StatSpellHit: 19, StatMP5: 6}},
	{ID: 28507, ItemType: ItemTypeHands, Name: "Handwraps of Flowing Thought", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Huntsman", Stats: Stats{StatStamina: 24, StatIntellect: 22, StatSpellPower: 35, StatSpellHit: 14}, GemSockets: []GemColor{GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellHit: 3}},
	{ID: 30170, ItemType: ItemTypeHands, Name: "Cataclysm Handgrips (Tier 5)", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "LeotherastheBlind", Stats: Stats{StatStamina: 25, StatIntellect: 27, StatSpellPower: 41, StatSpellCrit: 19, StatSpellHit: 19, StatMP5: 7}},
	{ID: 29987, ItemType: ItemTypeHands, Name: "Gauntlets of the Sun King", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Kaelthas", Stats: Stats{StatStamina: 28, StatIntellect: 29, StatSpellPower: 42, StatSpellCrit: 28}},
	{ID: 30725, ItemType: ItemTypeHands, Name: "Anger-Spark Gloves", Phase: 1, Quality: ItemQualityEpic, SourceZone: "World Boss", SourceDrop: "Doomwalker", Stats: Stats{StatSpellPower: 30, StatSpellCrit: 25, StatSpellHit: 20}, GemSockets: []GemColor{GemColorRed, GemColorRed}, SocketBonus: Stats{StatSpellCrit: 3}},
	{ID: 28780, ItemType: ItemTypeHands, Name: "Soul-Eater's Handwraps", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Magtheridon's Lair", SourceDrop: "Magtheridon", Stats: Stats{StatStamina: 31, StatIntellect: 24, StatSpellPower: 36, StatSpellCrit: 21}, GemSockets: []GemColor{GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellPower: 4}},
	{ID: 31008, ItemType: ItemTypeHands, Name: "Skyshatter Gauntlets (Tier 6)", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Azgalor", Stats: Stats{StatStamina: 30, StatIntellect: 31, StatSpellPower: 46, StatSpellCrit: 26, StatSpellHit: 19}, GemSockets: []GemColor{GemColorYellow}, SocketBonus: Stats{StatSpellPower: 2}},
	{ID: 28565, ItemType: ItemTypeWaist, Name: "Nethershard Girdle", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Moroes", Stats: Stats{StatStamina: 22, StatIntellect: 30, StatSpellPower: 35}},
	{ID: 28639, ItemType: ItemTypeWaist, Name: "General's Mail Girdle", Phase: 1, Quality: ItemQualityEpic, SourceZone: "PvP", SourceDrop: "PvP", Stats: Stats{StatStamina: 34, StatIntellect: 23, StatSpellPower: 28, StatSpellCrit: 23}},
	{ID: 28654, ItemType: ItemTypeWaist, Name: "Malefic Girdle", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Illhoof", Stats: Stats{StatStamina: 27, StatIntellect: 26, StatSpellPower: 37, StatSpellCrit: 21}},
	{ID: 30044, ItemType: ItemTypeWaist, Name: "Monsoon Belt", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC/TK", SourceDrop: "Leatherworking", Stats: Stats{StatStamina: 23, StatIntellect: 24, StatSpellPower: 39, StatSpellHit: 21}, GemSockets: []GemColor{GemColorBlue, GemColorYellow}, SocketBonus: Stats{StatSpellPower: 4}},
	{ID: 29520, ItemType: ItemTypeWaist, Name: "Netherstrike Belt", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Crafted", SourceDrop: "Leatherworking", Stats: Stats{StatStamina: 10, StatIntellect: 17, StatSpellPower: 30, StatSpellCrit: 16, StatMP5: 9}, GemSockets: []GemColor{GemColorBlue, GemColorYellow}, SocketBonus: Stats{StatSpellCrit: 3}},
	{ID: 28799, ItemType: ItemTypeWaist, Name: "Belt of Divine Inspiration", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Gruul's Lair", SourceDrop: "Maulgar", Stats: Stats{StatStamina: 27, StatIntellect: 26, StatSpellPower: 43}, GemSockets: []GemColor{GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellPower: 4}},
	{ID: 30064, ItemType: ItemTypeWaist, Name: "Cord of Screaming Terrors", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "Lurker", Stats: Stats{StatStamina: 34, StatIntellect: 15, StatSpellPower: 50, StatSpellHit: 24}, GemSockets: []GemColor{GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatStamina: 4}},
	{ID: 24256, ItemType: ItemTypeWaist, Name: "Girdle of Ruination", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Crafted", SourceDrop: "Tailoring", Stats: Stats{StatStamina: 18, StatIntellect: 13, StatSpellPower: 39, StatSpellCrit: 20}, GemSockets: []GemColor{GemColorRed, GemColorYellow}, SocketBonus: Stats{StatStamina: 4}},
	{ID: 30914, ItemType: ItemTypeWaist, Name: "Belt of the Crescent Moon", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Kazrogal", Stats: Stats{StatStamina: 25, StatIntellect: 27, StatSpellPower: 44, StatSpellHaste: 36}},
	{ID: 32256, ItemType: ItemTypeWaist, Name: "Waistwrap of Infinity", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Supremus", Stats: Stats{StatStamina: 31, StatIntellect: 22, StatSpellPower: 56, StatSpellHaste: 32}},
	{ID: 30038, ItemType: ItemTypeWaist, Name: "Belt of Blasting", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC/TK", SourceDrop: "Tailoring", Stats: Stats{StatSpellPower: 50, StatSpellCrit: 30, StatSpellHit: 23}, GemSockets: []GemColor{GemColorBlue, GemColorYellow}, SocketBonus: Stats{StatSpellPower: 4}},
	{ID: 30888, ItemType: ItemTypeWaist, Name: "Anetheron's Noose", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Anetheron", Stats: Stats{StatStamina: 22, StatIntellect: 23, StatSpellPower: 55, StatSpellCrit: 24}, GemSockets: []GemColor{GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellPower: 4}},
	{ID: 32276, ItemType: ItemTypeWaist, Name: "Flashfire Girdle", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Akama", Stats: Stats{StatStamina: 27, StatIntellect: 26, StatSpellPower: 44, StatSpellHaste: 37, StatSpellCrit: 18}},
	{ID: 29036, ItemType: ItemTypeLegs, Name: "Cyclone Legguards (Tier 4)", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Gruul's Lair", SourceDrop: "Gruul", Stats: Stats{StatStamina: 40, StatIntellect: 40, StatSpellPower: 49, StatSpellHit: 20, StatMP5: 8}},
	{ID: 28594, ItemType: ItemTypeLegs, Name: "Trial-Fire Trousers", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Opera", Stats: Stats{StatStamina: 42, StatIntellect: 40, StatSpellPower: 49}, GemSockets: []GemColor{GemColorYellow, GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 29972, ItemType: ItemTypeLegs, Name: "Trousers of the Astromancer", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Solarian", Stats: Stats{StatStamina: 33, StatIntellect: 36, StatSpellPower: 54}, GemSockets: []GemColor{GemColorBlue, GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 30172, ItemType: ItemTypeLegs, Name: "Cataclysm Leggings (Tier 5)", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Karathress", Stats: Stats{StatStamina: 48, StatIntellect: 46, StatSpellPower: 54, StatSpellCrit: 24, StatSpellHit: 14}, GemSockets: []GemColor{GemColorYellow}, SocketBonus: Stats{StatSpellPower: 2}},
	{ID: 32367, ItemType: ItemTypeLegs, Name: "Leggings of Devastation", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Mother", Stats: Stats{StatStamina: 40, StatIntellect: 42, StatSpellPower: 60, StatSpellHit: 26}, GemSockets: []GemColor{GemColorYellow, GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 31020, ItemType: ItemTypeLegs, Name: "Skyshatter Legguards (Tier 6)", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Illidari Council", Stats: Stats{StatStamina: 40, StatIntellect: 42, StatSpellPower: 62, StatSpellCrit: 29, StatSpellHit: 20, StatMP5: 11}, GemSockets: []GemColor{GemColorYellow}, SocketBonus: Stats{StatSpellPower: 2}},
	{ID: 30734, ItemType: ItemTypeLegs, Name: "Leggings of the Seventh Circle", Phase: 1, Quality: ItemQualityEpic, SourceZone: "World Boss", SourceDrop: "Kazzak", Stats: Stats{StatIntellect: 22, StatSpellPower: 50, StatSpellCrit: 25, StatSpellHit: 18}, GemSockets: []GemColor{GemColorRed, GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 30916, ItemType: ItemTypeLegs, Name: "Leggings of Channeled Elements", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Kazrogal", Stats: Stats{StatStamina: 25, StatIntellect: 28, StatSpellPower: 59, StatSpellCrit: 34, StatSpellHit: 18}, GemSockets: []GemColor{GemColorYellow, GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 28670, ItemType: ItemTypeFeet, Name: "Boots of the Infernal Coven", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Aran", Stats: Stats{StatStamina: 27, StatIntellect: 27, StatSpellPower: 34}},
	{ID: 28585, ItemType: ItemTypeFeet, Name: "Ruby Slippers", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Opera", Stats: Stats{StatStamina: 33, StatIntellect: 29, StatSpellPower: 35, StatSpellHit: 16}},
	{ID: 28810, ItemType: ItemTypeFeet, Name: "Windshear Boots", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Gruul's Lair", SourceDrop: "Gruul", Stats: Stats{StatStamina: 37, StatIntellect: 32, StatSpellPower: 39, StatSpellHit: 18}},
	{ID: 30894, ItemType: ItemTypeFeet, Name: "Blue Suede Shoes", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Kazrogal", Stats: Stats{StatStamina: 37, StatIntellect: 32, StatSpellPower: 56, StatSpellHit: 18}},
	{ID: 30037, ItemType: ItemTypeFeet, Name: "Boots of Blasting", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC/TK", SourceDrop: "Tailoring", Stats: Stats{StatStamina: 25, StatIntellect: 25, StatSpellPower: 39, StatSpellCrit: 25, StatSpellHit: 18}},
	{ID: 28517, ItemType: ItemTypeFeet, Name: "Boots of Foretelling", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Maiden", Stats: Stats{StatStamina: 27, StatIntellect: 23, StatSpellPower: 26, StatSpellCrit: 19}, GemSockets: []GemColor{GemColorRed, GemColorYellow}, SocketBonus: Stats{StatIntellect: 3}},
	{ID: 30043, ItemType: ItemTypeFeet, Name: "Hurricane Boots", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC/TK", SourceDrop: "Leatherworking", Stats: Stats{StatStamina: 25, StatIntellect: 26, StatSpellPower: 39, StatSpellCrit: 26, StatMP5: 6}},
	{ID: 30067, ItemType: ItemTypeFeet, Name: "Velvet Boots of the Guardian", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "Lurker", Stats: Stats{StatStamina: 21, StatIntellect: 21, StatSpellPower: 49, StatSpellCrit: 24}},
	{ID: 32242, ItemType: ItemTypeFeet, Name: "Boots of Oceanic Fury", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Najentus", Stats: Stats{StatStamina: 28, StatIntellect: 36, StatSpellPower: 55, StatSpellCrit: 26}},
	{ID: 32352, ItemType: ItemTypeFeet, Name: "Naturewarden's Treads", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "RoS", Stats: Stats{StatStamina: 39, StatIntellect: 18, StatSpellPower: 44, StatSpellCrit: 26, StatMP5: 7}, GemSockets: []GemColor{GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellPower: 4}},
	{ID: 32239, ItemType: ItemTypeFeet, Name: "Slippers of the Seacaller", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Najentus", Stats: Stats{StatStamina: 25, StatIntellect: 18, StatSpellPower: 44, StatSpellCrit: 29}, GemSockets: []GemColor{GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellPower: 4}},
	{ID: 28793, ItemType: ItemTypeFinger, Name: "Band of Crimson Fury", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Magtheridon's Lair", SourceDrop: "MagtheridonQuest", Stats: Stats{StatStamina: 22, StatIntellect: 22, StatSpellPower: 28, StatSpellHit: 16}},
	{ID: 28510, ItemType: ItemTypeFinger, Name: "Spectral Band of Innervation", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Huntsman", Stats: Stats{StatStamina: 22, StatIntellect: 24, StatSpellPower: 29}},
	{ID: 29922, ItemType: ItemTypeFinger, Name: "Band of Al'ar", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Alar", Stats: Stats{StatStamina: 24, StatIntellect: 23, StatSpellPower: 37}},
	{ID: 29367, ItemType: ItemTypeFinger, Name: "Ring of Cryptic Dreams", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Shattrah", SourceDrop: "Badges", Stats: Stats{StatStamina: 16, StatIntellect: 17, StatSpellPower: 23, StatSpellCrit: 20}},
	{ID: 29287, ItemType: ItemTypeFinger, Name: "Violet Signet of the Archmage", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Exalted", Stats: Stats{StatStamina: 24, StatIntellect: 23, StatSpellPower: 29, StatSpellCrit: 17}},
	{ID: 29286, ItemType: ItemTypeFinger, Name: "Violet Signet (R)", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Revered", Stats: Stats{StatStamina: 22, StatIntellect: 22, StatSpellPower: 28, StatSpellCrit: 17}},
	{ID: 29285, ItemType: ItemTypeFinger, Name: "Violet Signet (H)", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Honored", Stats: Stats{StatStamina: 19, StatIntellect: 21, StatSpellPower: 26, StatSpellCrit: 15}},
	{ID: 28753, ItemType: ItemTypeFinger, Name: "Ring of Recurrence", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Chess", Stats: Stats{StatStamina: 15, StatIntellect: 15, StatSpellPower: 32, StatSpellCrit: 19}},
	{ID: 29305, ItemType: ItemTypeFinger, Name: "Band of the Eternal Sage", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Exalted", Stats: Stats{StatStamina: 28, StatIntellect: 25, StatSpellPower: 34, StatSpellCrit: 24}},
	{ID: 30109, ItemType: ItemTypeFinger, Name: "Ring of Endless Coils", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "LadyVashj", Stats: Stats{StatStamina: 31, StatSpellPower: 37, StatSpellCrit: 22}},
	{ID: 30667, ItemType: ItemTypeFinger, Name: "Ring of Unrelenting Storms", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Trash", Stats: Stats{StatIntellect: 15, StatSpellPower: 43, StatSpellCrit: 19}},
	{ID: 32247, ItemType: ItemTypeFinger, Name: "Ring of Captured Storms", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Najentus", Stats: Stats{StatSpellPower: 42, StatSpellCrit: 29, StatSpellHit: 19}},
	{ID: 32527, ItemType: ItemTypeFinger, Name: "Ring of Ancient Knowledge", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Trash", Stats: Stats{StatStamina: 30, StatIntellect: 20, StatSpellPower: 39, StatSpellHaste: 31}},
	{ID: 30832, ItemType: ItemTypeWeapon, Name: "Gavel of Unearthed Secrets", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Shattrah", SourceDrop: "Lower City - Exalted", Stats: Stats{StatStamina: 24, StatIntellect: 16, StatSpellPower: 159, StatSpellCrit: 15}},
	{ID: 23554, ItemType: ItemTypeWeapon, Name: "Eternium Runed Blade", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Crafted", SourceDrop: "Blacksmithing", Stats: Stats{StatIntellect: 19, StatSpellPower: 168, StatSpellCrit: 21}},
	{ID: 28770, ItemType: ItemTypeWeapon, Name: "Nathrezim Mindblade", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Prince", Stats: Stats{StatStamina: 18, StatIntellect: 18, StatSpellPower: 203, StatSpellCrit: 23}},
	{ID: 30723, ItemType: ItemTypeWeapon, Name: "Talon of the Tempest", Phase: 1, Quality: ItemQualityEpic, SourceZone: "World Boss", SourceDrop: "Doomwalker", Stats: Stats{StatIntellect: 10, StatSpellPower: 194, StatSpellCrit: 19, StatSpellHit: 9}, GemSockets: []GemColor{GemColorYellow, GemColorYellow}, SocketBonus: Stats{StatIntellect: 3}},
	{ID: 34009, ItemType: ItemTypeWeapon, Name: "Hammer of Judgement", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Trash", Stats: Stats{StatStamina: 33, StatIntellect: 22, StatSpellPower: 236, StatSpellHit: 22}},
	{ID: 32237, ItemType: ItemTypeWeapon, Name: "The Maelstrom's Fury", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Najentus", Stats: Stats{StatStamina: 33, StatIntellect: 21, StatSpellPower: 236, StatSpellCrit: 22}},
	{ID: 28633, ItemType: ItemTypeWeapon, HandType: HandTypeTwoHand, Name: "Staff of Infinite Mysteries", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Curator", Stats: Stats{StatStamina: 61, StatIntellect: 51, StatSpellPower: 185, StatSpellHit: 23}},
	{ID: 29988, ItemType: ItemTypeWeapon, HandType: HandTypeTwoHand, Name: "The Nexus Key", Phase: 2, Quality: ItemQualityEpic, SourceZone: "TK", SourceDrop: "Kaelthas", Stats: Stats{StatStamina: 76, StatIntellect: 52, StatSpellPower: 236, StatSpellCrit: 51}},
	{ID: 32374, ItemType: ItemTypeWeapon, HandType: HandTypeTwoHand, Name: "Zhar'doom, Greatstaff of the Devourer", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Illidan", Stats: Stats{StatStamina: 70, StatIntellect: 47, StatSpellPower: 259, StatSpellHaste: 55, StatSpellCrit: 36}},
	{ID: 28734, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, Name: "Jewel of Infinite Possibilities", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Netherspite", Stats: Stats{StatStamina: 19, StatIntellect: 18, StatSpellPower: 23, StatSpellHit: 21}},
	{ID: 28611, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, WeaponType: WeaponTypeShield, Name: "Dragonheart Flameshield", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Nightbane", Stats: Stats{StatStamina: 19, StatIntellect: 33, StatSpellPower: 23, StatMP5: 7}},
	{ID: 34011, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, WeaponType: WeaponTypeShield, Name: "Illidari Runeshield", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Trash", Stats: Stats{StatStamina: 45, StatIntellect: 39, StatSpellPower: 34}},
	{ID: 28781, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, Name: "Karaborian Talisman", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Magtheridon's Lair", SourceDrop: "Magtheridon", Stats: Stats{StatStamina: 23, StatIntellect: 23, StatSpellPower: 35}},
	{ID: 29268, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, WeaponType: WeaponTypeShield, Name: "Mazthoril Honor Shield", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Shattrah", SourceDrop: "Badges", Stats: Stats{StatStamina: 16, StatIntellect: 17, StatSpellPower: 23, StatSpellCrit: 21}},
	{ID: 28603, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, Name: "Talisman of Nightbane", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Nightbane", Stats: Stats{StatStamina: 19, StatIntellect: 19, StatSpellPower: 28, StatSpellCrit: 17}},
	{ID: 32361, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, Name: "Blind-Seers Icon", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Akama", Stats: Stats{StatStamina: 25, StatIntellect: 16, StatSpellPower: 42, StatSpellHit: 24}},
	{ID: 29273, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, Name: "Khadgar's Knapsack", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Shattrah", SourceDrop: "Badges", Stats: Stats{StatSpellPower: 49}},
	{ID: 30049, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, Name: "FathomStone", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "Lurker", Stats: Stats{StatStamina: 16, StatIntellect: 12, StatSpellPower: 36, StatSpellCrit: 23}},
	{ID: 30909, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, WeaponType: WeaponTypeShield, Name: "Antonidas's Aegis of Rapt Concentration", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Archimonde", Stats: Stats{StatStamina: 28, StatIntellect: 32, StatSpellPower: 42, StatSpellCrit: 20}},
	{ID: 30872, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, Name: "Chronicle of Dark Secrets", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Hyjal", SourceDrop: "Winterchill", Stats: Stats{StatStamina: 16, StatIntellect: 12, StatSpellPower: 42, StatSpellCrit: 23, StatSpellHit: 17}},
	{ID: 28297, ItemType: ItemTypeWeapon, Name: "Gladiator's Gavel / Gladiator's Spellblade", Phase: 1, Quality: ItemQualityEpic, SourceZone: "PvP", SourceDrop: "PvP", Stats: Stats{StatStamina: 28, StatIntellect: 18, StatSpellPower: 199}},

	// Hand Written
	{ID: 27683, ItemType: ItemTypeTrinket, Name: "Quagmirran's Eye", Phase: 1, Quality: ItemQualityRare, SourceZone: "The Slave Pens", SourceDrop: "Quagmirran", Stats: Stats{StatSpellPower: 37}, Activate: ActivateQuagsEye, ActivateCD: NeverExpires, SharedID: MagicIDAtkTrinket},
	{ID: 29370, ItemType: ItemTypeTrinket, Name: "Icon of the Silver Crescent", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Shattrath", SourceDrop: "G'eras - 41 Badges", Stats: Stats{StatSpellPower: 43}, Activate: createSpellDmgActivate(MagicIDBlessingSilverCrescent, 155, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDISCTrink, SharedID: MagicIDAtkTrinket},
	{ID: 19344, ItemType: ItemTypeTrinket, Name: "Natural Alignment Crystal", Phase: 0, Quality: ItemQualityEpic, SourceZone: "BWL", SourceDrop: "", Stats: Stats{}, Activate: ActivateNAC, ActivateCD: time.Second * 300, CoolID: MagicIDNACTrink, SharedID: MagicIDAtkTrinket},
	{ID: 19379, ItemType: ItemTypeTrinket, Name: "Neltharion's Tear", Phase: 0, Quality: ItemQualityEpic, SourceZone: "BWL", SourceDrop: "Nefarian", Stats: Stats{StatSpellPower: 44, StatSpellHit: 16}, SharedID: MagicIDAtkTrinket},
	{ID: 23046, ItemType: ItemTypeTrinket, Name: "The Restrained Essence of Sapphiron", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "Sapphiron", Stats: Stats{StatSpellPower: 40}, Activate: createSpellDmgActivate(MagicIDSpellPower, 130, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDEssSappTrink, SharedID: MagicIDAtkTrinket},
	{ID: 23207, ItemType: ItemTypeTrinket, Name: "Mark of the Champion", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "KT", Stats: Stats{StatSpellPower: 85}, SharedID: MagicIDAtkTrinket},
	{ID: 29132, ItemType: ItemTypeTrinket, Name: "Scryer's Bloodgem", Phase: 1, Quality: ItemQualityRare, SourceZone: "The Scryers - Revered", SourceDrop: "", Stats: Stats{StatSpellHit: 32}, Activate: createSpellDmgActivate(MagicIDSpellPower, 150, time.Second*15), ActivateCD: time.Second * 90, CoolID: MagicIDScryerTrink, SharedID: MagicIDAtkTrinket},
	{ID: 24126, ItemType: ItemTypeTrinket, Name: "Figurine - Living Ruby Serpent", Phase: 1, Quality: ItemQualityRare, SourceZone: "Jewelcrafting BoP", SourceDrop: "", Stats: Stats{StatIntellect: 23, StatStamina: 33}, Activate: createSpellDmgActivate(MagicIDRubySerpent, 150, time.Second*20), ActivateCD: time.Second * 300, CoolID: MagicIDRubySerpentTrink, SharedID: MagicIDAtkTrinket},
	{ID: 29179, ItemType: ItemTypeTrinket, Name: "Xi'ri's Gift", Phase: 1, Quality: ItemQualityRare, SourceZone: "The Sha'tar - Revered", SourceDrop: "", Stats: Stats{StatSpellCrit: 32}, Activate: createSpellDmgActivate(MagicIDSpellPower, 150, time.Second*15), ActivateCD: time.Second * 90, CoolID: MagicIDXiriTrink, SharedID: MagicIDAtkTrinket},
	{ID: 28418, ItemType: ItemTypeTrinket, Name: "Shiffar's Nexus-Horn", Phase: 1, Quality: ItemQualityRare, SourceZone: "Arc - Harbinger Skyriss", SourceDrop: "", Stats: Stats{StatSpellCrit: 30}, Activate: ActivateNexusHorn, ActivateCD: NeverExpires, SharedID: MagicIDAtkTrinket},
	{ID: 31856, ItemType: ItemTypeTrinket, Name: "Darkmoon Card: Crusade", Phase: 2, Quality: ItemQualityEpic, SourceZone: "Blessings Deck", SourceDrop: "", Activate: ActivateDCC, ActivateCD: NeverExpires, SharedID: MagicIDAtkTrinket},
	{ID: 28785, ItemType: ItemTypeTrinket, Name: "The Lightning Capacitor", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "", Activate: ActivateTLC, ActivateCD: NeverExpires, SharedID: MagicIDAtkTrinket},
	{ID: 28789, ItemType: ItemTypeTrinket, Name: "Eye of Magtheridon", Phase: 1, Quality: ItemQualityEpic, SourceZone: "", SourceDrop: "", Stats: Stats{StatSpellPower: 54}, Activate: ActivateEyeOfMag, ActivateCD: NeverExpires, SharedID: MagicIDAtkTrinket},
	{ID: 30626, ItemType: ItemTypeTrinket, Name: "Sextant of Unstable Currents", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "", Stats: Stats{StatSpellCrit: 40}, Activate: ActivateSextant, ActivateCD: NeverExpires, SharedID: MagicIDAtkTrinket},
	{ID: 34429, ItemType: ItemTypeTrinket, Name: "Shifting Naaru Sliver", Phase: 5, Quality: ItemQualityEpic, SourceZone: "Sunwell", SourceDrop: "", Stats: Stats{StatSpellHaste: 54}, Activate: createSpellDmgActivate(MagicIDShiftingNaaru, 320, time.Second*15), ActivateCD: time.Second * 90, CoolID: MagicIDShiftingNaaruTrink, SharedID: MagicIDAtkTrinket},
	{ID: 32483, ItemType: ItemTypeTrinket, Name: "The Skull of Gul'dan", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Black Temple", SourceDrop: "", Stats: Stats{StatSpellHit: 25, StatSpellPower: 55}, Activate: createHasteActivate(MagicIDSkullGuldan, 175, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDSkullGuldanTrink, SharedID: MagicIDAtkTrinket},
	{ID: 33829, ItemType: ItemTypeTrinket, Name: "Hex Shrunken Head", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "", Stats: Stats{StatSpellPower: 53}, Activate: createSpellDmgActivate(MagicIDHexShunkHead, 211, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDHexTrink, SharedID: MagicIDAtkTrinket},
	{ID: 29376, ItemType: ItemTypeTrinket, Name: "Essence of the Martyr", Phase: 1, Quality: ItemQualityRare, SourceZone: "G'eras", SourceDrop: "Badges", Stats: Stats{StatSpellPower: 28}, Activate: createSpellDmgActivate(MagicIDSpellPower, 99, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDEssMartyrTrink, SharedID: MagicIDHealTrinket},
	{ID: 38290, ItemType: ItemTypeTrinket, Name: "Dark Iron Smoking Pipe", Phase: 2, Quality: ItemQualityEpic, SourceZone: "Brewfest", SourceDrop: "Corin Direbrew", Stats: Stats{StatSpellPower: 43}, Activate: createSpellDmgActivate(MagicIDDarkIronPipeweed, 155, time.Second*20), ActivateCD: time.Second * 120, CoolID: MagicIDDITrink, SharedID: MagicIDAtkTrinket},
	{ID: 30663, ItemType: ItemTypeTrinket, Name: "Fathom-Brooch of the Tidewalker", Phase: 2, Quality: ItemQualityEpic, SourceZone: "SSC", SourceDrop: "Fathom-Lord Karathress", Stats: Stats{}, Activate: ActivateFathomBrooch, ActivateCD: NeverExpires, SharedID: MagicIDAtkTrinket},
	{ID: 35749, ItemType: ItemTypeTrinket, Name: "Sorcerer's Alchemist Stone", Phase: 5, Quality: ItemQualityEpic, SourceZone: "Shattered Sun Offensive", SourceDrop: "Exalted", Stats: Stats{StatSpellPower: 63}, Activate: ActivateAlchStone, ActivateCD: NeverExpires, SharedID: MagicIDAtkTrinket},

	{ID: 24116, ItemType: ItemTypeNeck, Name: "Eye of the Night", Phase: 1, Quality: ItemQualityRare, SourceZone: "Jewelcrafting", SourceDrop: "", Stats: Stats{StatSpellCrit: 26, StatSpellHit: 16}, Activate: ActivateEyeOfNight, ActivateCD: NeverExpires},
	{ID: 24121, ItemType: ItemTypeNeck, Name: "Chain of the Twilight Owl", Phase: 1, Quality: ItemQualityRare, SourceZone: "Jewelcrafting", SourceDrop: "", Stats: Stats{StatIntellect: 19, StatSpellPower: 21}, Activate: ActivateChainTO, ActivateCD: NeverExpires},
	{ID: 31075, ItemType: ItemTypeFinger, Name: "Evoker's Mark of the Redemption", Phase: 1, Quality: ItemQualityRare, SourceZone: "Quest SMV", SourceDrop: "Dissension Amongst the Ranks...", Stats: Stats{StatIntellect: 15, StatSpellPower: 29, StatSpellCrit: 10}},
	{ID: 32664, ItemType: ItemTypeFinger, Name: "Dreamcrystal Band", Phase: 1, Quality: ItemQualityRare, SourceZone: "Blades Edge Moutains", SourceDrop: "50 Apexis Shards", Stats: Stats{StatIntellect: 10, StatSpellPower: 38, StatSpellCrit: 15}},
	{ID: 29522, ItemType: ItemTypeChest, Name: "Windhawk Hauberk", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Leatherworking", SourceDrop: "", Stats: Stats{StatStamina: 28, StatIntellect: 29, StatSpirit: 29, StatSpellPower: 46, StatSpellCrit: 19}, GemSockets: []GemColor{GemColorBlue, GemColorYellow, GemColorBlue}, SocketBonus: Stats{StatSpellPower: 5}},
	{ID: 29524, ItemType: ItemTypeWaist, Name: "Windhawk Belt", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Leatherworking", SourceDrop: "", Stats: Stats{StatStamina: 17, StatIntellect: 19, StatSpirit: 20, StatSpellPower: 37, StatSpellCrit: 12}, GemSockets: []GemColor{GemColorBlue, GemColorYellow}, SocketBonus: Stats{StatSpellPower: 4}},
	{ID: 29523, ItemType: ItemTypeWrist, Name: "Windhawk Bracers", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Leatherworking", SourceDrop: "", Stats: Stats{StatStamina: 22, StatIntellect: 17, StatSpirit: 7, StatSpellPower: 27, StatSpellCrit: 16}, GemSockets: []GemColor{GemColorYellow}, SocketBonus: Stats{StatIntellect: 2}},
	{ID: 27510, ItemType: ItemTypeHands, Name: "Tidefury Gauntlets", Phase: 1, Quality: ItemQualityRare, SourceZone: "", SourceDrop: "", Stats: Stats{StatStamina: 22, StatIntellect: 26, StatSpellPower: 29, StatMP5: 7}},
	{ID: 22730, ItemType: ItemTypeWaist, Name: "Eyestalk Waist Cord", Phase: 0, Quality: ItemQualityEpic, SourceZone: "AQ40", SourceDrop: "C'thun", Stats: Stats{StatStamina: 10, StatIntellect: 9, StatSpellPower: 41, StatSpellCrit: 14}},
	{ID: 23070, ItemType: ItemTypeLegs, Name: "Leggings of Polarity", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "Thaddius", Stats: Stats{StatStamina: 20, StatIntellect: 14, StatSpellPower: 44, StatSpellCrit: 28}},
	{ID: 21709, ItemType: ItemTypeFinger, Name: "Ring of the Fallen God", Phase: 0, Quality: ItemQualityEpic, SourceZone: "AQ40", SourceDrop: "C'thun", Stats: Stats{StatStamina: 5, StatIntellect: 6, StatSpellPower: 37, StatSpellHit: 8}},
	{ID: 23031, ItemType: ItemTypeFinger, Name: "Band of the Inevitable", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "Noth", Stats: Stats{StatSpellPower: 36, StatSpellHit: 8}},
	{ID: 23025, ItemType: ItemTypeFinger, Name: "Seal of the Damned", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "4H", Stats: Stats{StatStamina: 17, StatSpellPower: 21, StatSpellCrit: 14, StatSpellHit: 8}},
	{ID: 23057, ItemType: ItemTypeNeck, Name: "Gem of Trapped Innocents", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "KT", Stats: Stats{StatStamina: 9, StatIntellect: 7, StatSpellPower: 15, StatSpellCrit: 28}},
	{ID: 21608, ItemType: ItemTypeNeck, Name: "Amulet of Vek'nilash", Phase: 0, Quality: ItemQualityEpic, SourceZone: "AQ", SourceDrop: "Twin Emp", Stats: Stats{StatStamina: 9, StatIntellect: 5, StatSpellPower: 27, StatSpellCrit: 14}},
	{ID: 23664, ItemType: ItemTypeShoulder, Name: "Pauldrons of Elemental Fury", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "Trash", Stats: Stats{StatStamina: 19, StatIntellect: 21, StatSpellPower: 26, StatSpellCrit: 14, StatSpellHit: 8}},
	{ID: 23665, ItemType: ItemTypeLegs, Name: "Leggings of Elemental Fury", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "Trash", Stats: Stats{StatStamina: 26, StatIntellect: 27, StatSpellPower: 32, StatSpellCrit: 28}},
	{ID: 23050, ItemType: ItemTypeBack, Name: "Cloak of the Necropolis", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "Sapp", Stats: Stats{StatStamina: 12, StatIntellect: 11, StatSpellPower: 26, StatSpellCrit: 14, StatSpellHit: 8}},
	{ID: 30682, ItemType: ItemTypeFeet, Name: "Glider's Sabatons of Nature's Wrath", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Servants Quarter", Stats: Stats{StatNatureSpellPower: 78}},
	{ID: 30677, ItemType: ItemTypeWaist, Name: "Lurker's Belt of Nature's Wrath", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Servants Quarter", Stats: Stats{StatNatureSpellPower: 78}},
	{ID: 30686, ItemType: ItemTypeWrist, Name: "Ravager's Bands of Nature's Wrath", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Servants Quarter", Stats: Stats{StatNatureSpellPower: 58}},
	// {ItemType: ItemTypeFeet, Name: "Glider's Sabatons of the Invoker", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Servants Quarter", Stats: Stats{StatSpellPower: 33, StatSpellCrit: 28}},
	// {ItemType: ItemTypeWaist, Name: "Lurker's Belt of the Invoker", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Servants Quarter", Stats: Stats{StatSpellPower: 33, StatSpellCrit: 28}},
	// {ItemType: ItemTypeWrist, Name: "Ravager's Bands of the Invoker", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "Servants Quarter", Stats: Stats{StatSpellPower: 25, StatSpellCrit: 21}},
	{ID: 28583, ItemType: ItemTypeHead, Name: "Big Bad Wolf's Head", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Kara", SourceDrop: "The Big Bad Wolf", Stats: Stats{StatStamina: 42, StatIntellect: 40, StatSpellPower: 47, StatSpellCrit: 28}},
	{ID: 32586, ItemType: ItemTypeWrist, Name: "Bracers of Nimble Thought", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Tailoring", Stats: Stats{StatStamina: 27, StatIntellect: 20, StatSpellPower: 34, StatSpellHaste: 28}},
	{ID: 23049, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, Name: "Sapphiron's Left Eye", Phase: 0, Quality: ItemQualityEpic, SourceZone: "Naxx", SourceDrop: "Sapphiron", Stats: Stats{StatStamina: 12, StatIntellect: 8, StatSpellPower: 26, StatSpellCrit: 14, StatSpellHit: 8}},
	{ID: 25778, ItemType: ItemTypeWrist, Name: "Manacles of Rememberance", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "Nagrand", SourceDrop: "Quest", Stats: Stats{StatSpirit: 9, StatIntellect: 10, StatSpellPower: 16, StatSpellCrit: 14}},
	{ID: 28174, ItemType: ItemTypeWrist, Name: "Shattrath Wraps", Phase: 1, Quality: ItemQualityRare, SourceZone: "Auchindoun", SourceDrop: "Quest", Stats: Stats{StatStamina: 15, StatIntellect: 15, StatSpellPower: 21}, GemSockets: []GemColor{GemColorRed}, SocketBonus: Stats{StatStamina: 3}},
	{ID: 31283, ItemType: ItemTypeWaist, Name: "Sash of Sealed Fate", Phase: 1, Quality: ItemQualityRare, SourceZone: "World Drop", SourceDrop: "BoE", Stats: Stats{StatIntellect: 15, StatSpellPower: 35, StatSpellCrit: 23}},
	{ID: 30004, ItemType: ItemTypeFeet, Name: "Landing Boots", Phase: 1, Quality: ItemQualityUncommon, SourceZone: "Netherstorm", SourceDrop: "Quest", Stats: Stats{StatStamina: 12, StatIntellect: 8, StatSpellPower: 35, StatSpellCrit: 16}},
	{ID: 31290, ItemType: ItemTypeFinger, Name: "Band of Dominion", Phase: 1, Quality: ItemQualityRare, SourceZone: "World Drop", SourceDrop: "BoE", Stats: Stats{StatSpellPower: 28, StatSpellCrit: 21}},

	// Sash of Sealed Fate - blue BoE
	// Landing Boots - green from quest
	// Band of Dominion - Blue BoE

	{ID: 34336, ItemType: ItemTypeWeapon, Name: "Sunflare", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "Kil'jaden",
		Stats: Stats{StatStamina: 17, StatIntellect: 20, StatSpellPower: 292, StatSpellHaste: 23, StatSpellCrit: 30}},
	{ID: 34179, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, Name: "Heart of the Pit", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "Brutalis",
		Stats: Stats{StatStamina: 33, StatIntellect: 21, StatSpellPower: 39, StatSpellHaste: 32}},
	{ID: 34350, ItemType: ItemTypeHands, Name: "Gauntlets of the Ancient Shadowmoon", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "Trash",
		Stats:       Stats{StatStamina: 30, StatIntellect: 32, StatSpellPower: 43, StatSpellHaste: 24, StatSpellCrit: 28},
		GemSockets:  []GemColor{GemColorRed, GemColorBlue},
		SocketBonus: Stats{StatSpellCrit: 2},
	},
	{ID: 34542, ItemType: ItemTypeWaist, Name: "Skyshatter Cord", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStamina: 19, StatIntellect: 30, StatSpellPower: 50, StatSpellHaste: 27, StatSpellCrit: 29, StatMP5: 6},
		GemSockets:  []GemColor{GemColorYellow},
		SocketBonus: Stats{StatSpellPower: 2},
	},
	{ID: 34186, ItemType: ItemTypeLegs, Name: "Chain Links of the Tumultuous Storm", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStamina: 48, StatIntellect: 41, StatSpellPower: 71, StatSpellHaste: 30, StatSpellCrit: 35},
		GemSockets:  []GemColor{GemColorYellow, GemColorRed, GemColorRed},
		SocketBonus: Stats{StatSpellCrit: 4},
	},
	{ID: 34566, ItemType: ItemTypeFeet, Name: "Skyshatter Treads", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStamina: 21, StatIntellect: 30, StatSpellPower: 50, StatSpellHaste: 30, StatSpellCrit: 23, StatMP5: 7},
		GemSockets:  []GemColor{GemColorYellow},
		SocketBonus: Stats{StatSpellPower: 2},
	},
	{ID: 34437, ItemType: ItemTypeWrist, Name: "Skyshatter Bands", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStamina: 15, StatIntellect: 23, StatSpellPower: 39, StatSpellHaste: 11, StatSpellCrit: 28},
		GemSockets:  []GemColor{GemColorYellow},
		SocketBonus: Stats{StatSpellPower: 2},
	},
	{ID: 34230, ItemType: ItemTypeFinger, Name: "Ring of Omnipotence", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats: Stats{StatStamina: 21, StatIntellect: 14, StatSpellPower: 40, StatSpellHaste: 31, StatSpellCrit: 22},
	},
	{ID: 34362, ItemType: ItemTypeFinger, Name: "Loop of Forged Power", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats: Stats{StatStamina: 27, StatIntellect: 28, StatSpellPower: 34, StatSpellHaste: 30, StatSpellHit: 19},
	},
	{ID: 34204, ItemType: ItemTypeNeck, Name: "Amulet of Unfettered Magics", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats: Stats{StatStamina: 24, StatIntellect: 17, StatSpellPower: 39, StatSpellHaste: 32, StatSpellHit: 15},
	},
	{ID: 34332, ItemType: ItemTypeHead, Name: "Cowl of Gul'dan", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStamina: 51, StatIntellect: 43, StatSpellPower: 74, StatSpellHaste: 32, StatSpellCrit: 36},
		GemSockets:  []GemColor{GemColorMeta, GemColorYellow},
		SocketBonus: Stats{StatSpellPower: 5},
	},
	{ID: 34242, ItemType: ItemTypeBack, Name: "Tattered Cape of Antonidas", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStamina: 25, StatIntellect: 26, StatSpellPower: 42, StatSpellHaste: 32},
		GemSockets:  []GemColor{GemColorRed},
		SocketBonus: Stats{StatSpellPower: 2},
	},
	{ID: 34396, ItemType: ItemTypeChest, Name: "Garments of Crashing Shores", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStamina: 48, StatIntellect: 41, StatSpellPower: 71, StatSpellHaste: 40, StatSpellCrit: 25},
		GemSockets:  []GemColor{GemColorRed, GemColorYellow, GemColorYellow},
		SocketBonus: Stats{StatSpellPower: 5},
	},
	{ID: 34390, ItemType: ItemTypeShoulder, Name: "Erupting Epaulets", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStamina: 30, StatIntellect: 30, StatSpellPower: 53, StatSpellHaste: 24, StatSpellCrit: 30},
		GemSockets:  []GemColor{GemColorYellow, GemColorRed},
		SocketBonus: Stats{StatSpellPower: 4},
	},
	{ID: 33970, ItemType: ItemTypeShoulder, Name: "Pauldrons of the Furious Elements", Phase: 4, Quality: ItemQualityEpic, SourceZone: "Shattrath - G'eras", SourceDrop: "60 Badges",
		Stats: Stats{StatStamina: 28, StatIntellect: 24, StatSpellPower: 40, StatSpellHaste: 33},
	},
	{ID: 33965, ItemType: ItemTypeChest, Name: "Hauberk of the Furious Elements", Phase: 4, Quality: ItemQualityEpic, SourceZone: "Shattrath - G'eras", SourceDrop: "75 Badges",
		Stats: Stats{StatStamina: 39, StatIntellect: 34, StatSpellPower: 54, StatSpellHaste: 35, StatSpellCrit: 23},
	},
	{ID: 33588, ItemType: ItemTypeWrist, Name: "Runed Spell-Cuffs", Phase: 4, Quality: ItemQualityEpic, SourceZone: "Shattrath - G'eras", SourceDrop: "Badges",
		Stats: Stats{StatStamina: 20, StatIntellect: 18, StatSpellPower: 29, StatSpellHaste: 25},
	},
	{ID: 33537, ItemType: ItemTypeFeet, Name: "Treads of Booming Thunder", Phase: 4, Quality: ItemQualityEpic, SourceZone: "Shattrath - G'eras", SourceDrop: "Badges",
		Stats:       Stats{StatStamina: 21, StatIntellect: 33, StatSpellPower: 40, StatSpellCrit: 14},
		GemSockets:  []GemColor{GemColorRed, GemColorYellow},
		SocketBonus: Stats{StatSpellCrit: 3},
	},
	{ID: 33534, ItemType: ItemTypeHands, Name: "Grips of Nature's Wrath", Phase: 4, Quality: ItemQualityEpic, SourceZone: "Shattrath - G'eras", SourceDrop: "Badges",
		Stats:       Stats{StatStamina: 30, StatIntellect: 27, StatSpellPower: 34, StatSpellCrit: 21},
		GemSockets:  []GemColor{GemColorRed, GemColorYellow},
		SocketBonus: Stats{StatSpellPower: 4},
	},
	{ID: 34359, ItemType: ItemTypeNeck, Name: "Pendant of Sunfire", Phase: 5, Quality: ItemQualityEpic, SourceZone: "Sunwell", SourceDrop: "Jewelcrafting",
		Stats:       Stats{StatStamina: 27, StatIntellect: 19, StatSpellPower: 34, StatSpellHaste: 25, StatSpellCrit: 25},
		GemSockets:  []GemColor{GemColorYellow},
		SocketBonus: Stats{StatSpellPower: 2},
	},

	// {Slot:EquipTrinket, Name:"Arcanist's Stone", Phase: 1, Quality: ItemQualityEpic, SourceZone:"H OHF - Epoch Hunter", SourceDrop:"", Stats:Stats{  StatSpellHit: 25,   StatMP5:0} }
	// {Slot:EquipTrinket, Name:"Vengeance of the Illidari", Phase: 1, Quality: ItemQualityEpic, SourceZone:"Cruel's Intentions/Overlord - HFP Quest", SourceDrop:"", Stats:Stats{ StatSpellCrit: 26,    StatMP5:0} }
	{ID: 33506, ItemType: ItemTypeRanged, RangedWeaponType: RangedWeaponTypeTotem, Name: "Skycall Totem", Phase: 4, Quality: ItemQualityEpic, SourceZone: "Geras", SourceDrop: "20 Badges", Stats: Stats{}, Activate: ActivateSkycall, ActivateCD: NeverExpires},
	{ID: 32086, ItemType: ItemTypeHead, Name: "Storm Master's Helmet", Phase: 1, Quality: ItemQualityRare, SourceZone: "Geras", SourceDrop: "50 Badges", Stats: Stats{StatStamina: 24, StatIntellect: 32, StatSpellCrit: 24, StatSpellPower: 37}, GemSockets: []GemColor{GemColorMeta, GemColorBlue}, SocketBonus: Stats{StatSpellCrit: 4}},
	{ID: 28602, ItemType: ItemTypeChest, Name: "Robe of the Elder Scribes", Phase: 1, Quality: ItemQualityEpic, SourceZone: "Karazhan", SourceDrop: "Nightbane", Stats: Stats{StatStamina: 27, StatIntellect: 29, StatSpirit: 24, StatSpellPower: 32, StatSpellCrit: 24}, Activate: ActivateElderScribes, ActivateCD: NeverExpires},

	{ID: 32963, ItemType: ItemTypeWeapon, Name: "Merciless Gladiator's Gavel / Spellblade", Phase: 2, Quality: ItemQualityEpic, SourceZone: "PvP", SourceDrop: "", Stats: Stats{StatStamina: 27, StatIntellect: 18, StatSpellPower: 225, StatSpellHit: 15}},
	{ID: 32524, ItemType: ItemTypeBack, Name: "Shroud of the Highborne", Phase: 3, Quality: ItemQualityEpic, SourceZone: "Black Temple", SourceDrop: "Illidan Stormrage", Stats: Stats{StatStamina: 24, StatIntellect: 23, StatSpellPower: 23, StatSpellHaste: 32}},
	{ID: 33357, ItemType: ItemTypeFeet, Name: "Footpads of Madness", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "Jan'alai", Stats: Stats{StatStamina: 25, StatIntellect: 22, StatSpellPower: 50, StatSpellHaste: 25}},
	{ID: 33533, ItemType: ItemTypeLegs, Name: "Avalanche Leggings", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "Halazzi",
		Stats:       Stats{StatStamina: 31, StatIntellect: 40, StatSpellPower: 46, StatSpellCrit: 30},
		GemSockets:  []GemColor{GemColorRed, GemColorYellow, GemColorBlue},
		SocketBonus: Stats{StatSpellPower: 5},
	},
	{ID: 33354, ItemType: ItemTypeWeapon, Name: "Wub's Cursed Hexblade", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "", Stats: Stats{StatIntellect: 21, StatSpellPower: 217, StatSpellHit: 13, StatSpellCrit: 20, StatMP5: 6}},
	{ID: 33283, ItemType: ItemTypeWeapon, Name: "Amani Punisher", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "", Stats: Stats{StatStamina: 30, StatIntellect: 21, StatSpellPower: 217, StatSpellHit: 20}},
	{ID: 33466, ItemType: ItemTypeNeck, Name: "Loop of Cursed Bones", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "", Stats: Stats{StatStamina: 19, StatIntellect: 20, StatSpellPower: 32, StatSpellHaste: 27}},
	{ID: 33591, ItemType: ItemTypeBack, Name: "Shadowcaster's Drape", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "", Stats: Stats{StatStamina: 22, StatIntellect: 20, StatSpellPower: 27, StatSpellHaste: 25}},
	{ID: 32817, ItemType: ItemTypeWrist, Name: "Veteran's Mail Bracers", Phase: 2, Quality: ItemQualityEpic, SourceZone: "", SourceDrop: "",
		Stats:       Stats{StatStamina: 25, StatIntellect: 14, StatSpellPower: 22, StatSpellCrit: 17},
		GemSockets:  []GemColor{GemColorYellow},
		SocketBonus: Stats{}, // resil bonus
	},
	{ID: 32792, ItemType: ItemTypeFeet, Name: "Veteran's Mail Sabatons", Phase: 2, Quality: ItemQualityEpic, SourceZone: "", SourceDrop: "", Stats: Stats{StatStamina: 39, StatIntellect: 27, StatSpellPower: 32, StatSpellCrit: 26}},
	{ID: 32328, ItemType: ItemTypeHands, Name: "Botanist's Gloves of Growth", Phase: 3, Quality: ItemQualityEpic, SourceZone: "BT", SourceDrop: "Teron Gorefiend",
		Stats:       Stats{StatStamina: 22, StatIntellect: 21, StatSpellPower: 28, StatSpellHaste: 37},
		GemSockets:  []GemColor{GemColorYellow, GemColorBlue},
		SocketBonus: Stats{StatSpellPower: 3},
	},
	{ID: 33281, ItemType: ItemTypeNeck, Name: "Brooch of Nature's Mercy", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "", Stats: Stats{StatIntellect: 24, StatSpellPower: 25, StatSpellHaste: 33}},
	{ID: 33334, ItemType: ItemTypeWeapon, HandType: HandTypeOffHand, Name: "Fetish of the Primal Gods", Phase: 4, Quality: ItemQualityEpic, SourceZone: "ZA", SourceDrop: "", Stats: Stats{StatStamina: 24, StatIntellect: 17, StatSpellPower: 37, StatSpellHaste: 17}},
	{ID: 34344, ItemType: ItemTypeHands, Name: "Handguards of Defiled Worlds", Phase: 5, Quality: ItemQualityEpic, SourceZone: "SW", SourceDrop: "",
		Stats:       Stats{StatStamina: 33, StatIntellect: 32, StatSpellPower: 47, StatSpellHaste: 36, StatSpellHit: 27},
		GemSockets:  []GemColor{GemColorYellow, GemColorRed},
		SocketBonus: Stats{StatSpellPower: 4},
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
		Bonuses: map[int]ItemActivation{3: func(sim *Simulation, party *Party, player PlayerAgent) Aura {
			player.Stats[StatSpellPower] += 23
			return Aura{ID: MagicIDNetherstrike}
		}},
	},
	{
		Name:  "The Twin Stars",
		Items: map[string]bool{"Charlotte's Ivy": true, "Lola's Eve": true},
		Bonuses: map[int]ItemActivation{2: func(sim *Simulation, party *Party, player PlayerAgent) Aura {
			player.Stats[StatSpellPower] += 15
			return Aura{ID: MagicIDNetherstrike}
		}},
	},
	{
		Name:  "Tidefury",
		Items: map[string]bool{"Tidefury Helm": true, "Tidefury Shoulderguards": true, "Tidefury Chestpiece": true, "Tidefury Kilt": true, "Tidefury Gauntlets": true},
		Bonuses: map[int]ItemActivation{
			2: func(sim *Simulation, party *Party, player PlayerAgent) Aura {
				return Aura{ID: MagicIDTidefury}
			},
			4: func(sim *Simulation, party *Party, player PlayerAgent) Aura {
				// TODO: should we even allow for unchecking water shield?
				// if sim.Options.Buffs.WaterShield {
				player.Stats[StatMP5] += 3
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
		Bonuses: map[int]ItemActivation{4: ActivateManaEtched, 2: func(sim *Simulation, party *Party, player PlayerAgent) Aura {
			player.Stats[StatSpellHit] += 35
			return Aura{ID: MagicIDManaEtchedHit}
		}},
	},
	{
		Name:  "Cyclone Regalia",
		Items: map[string]bool{"Cyclone Faceguard (Tier 4)": true, "Cyclone Shoulderguards (Tier 4)": true, "Cyclone Chestguard (Tier 4)": true, "Cyclone Handguards (Tier 4)": true, "Cyclone Legguards (Tier 4)": true},
		Bonuses: map[int]ItemActivation{4: ActivateCycloneManaReduce, 2: func(sim *Simulation, party *Party, player PlayerAgent) Aura {
			// if sim.Options.Totems.WrathOfAir {

			// FUTURE: Only one ele shaman in the party can use this at a time.
			//   not a big deal now but will need to be fixed to support full raid sim.
			for _, p := range party.Players {
				p.Stats[StatSpellPower] += 20
			}

			// }
			return Aura{ID: MagicIDCyclone2pc}
		}},
	},
	{
		Name:  "Windhawk",
		Items: map[string]bool{"Windhawk Hauberk": true, "Windhawk Belt": true, "Windhawk Bracers": true},
		Bonuses: map[int]ItemActivation{3: func(sim *Simulation, party *Party, player PlayerAgent) Aura {
			// TODO: check if player has water shield on?
			player.Stats[StatMP5] += 8
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
		Bonuses: map[int]ItemActivation{2: func(sim *Simulation, party *Party, player PlayerAgent) Aura {
			player.Stats[StatMP5] += 15
			player.Stats[StatSpellCrit] += 35
			player.Stats[StatSpellPower] += 45
			return Aura{ID: MagicIDSkyshatter2pc}
		}, 4: ActivateSkyshatterImpLB},
	},
}
