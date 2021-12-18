package items

import (
	"fmt"
	"log"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ByName = map[string]Item{}
var ByID = map[int32]Item{}
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

	// Add hard-coded items. Wowhead doesn't seem to have tooltips for random enchant items.
	// Use negative IDs to avoid collisions with real item IDs.
	Items = append(Items, []Item{
		{Name: "Glider's Boots of Nature's Wrath", WowheadID: 30681, ID: -1, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeFeet, ArmorType: proto.ArmorType_ArmorTypeLeather, Phase: 1, Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Armor: 250, stats.NatureSpellPower: 78}},
		{Name: "Glider's Foot-Wraps of Arcane Wrath", WowheadID: 30680, ID: -2, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeFeet, ArmorType: proto.ArmorType_ArmorTypeCloth, Phase: 1, Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Armor: 134, stats.ArcaneSpellPower: 78}},
		{Name: "Glider's Foot-Wraps of Fiery Wrath", WowheadID: 30680, ID: -3, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeFeet, ArmorType: proto.ArmorType_ArmorTypeCloth, Phase: 1, Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Armor: 134, stats.FireSpellPower: 78}},
		{Name: "Glider's Foot-Wraps of Frozen Wrath", WowheadID: 30680, ID: -4, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeFeet, ArmorType: proto.ArmorType_ArmorTypeCloth, Phase: 1, Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Armor: 134, stats.FrostSpellPower: 78}},
		{Name: "Glider's Foot-Wraps of Shadow Wrath", WowheadID: 30680, ID: -5, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeFeet, ArmorType: proto.ArmorType_ArmorTypeCloth, Phase: 1, Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Armor: 134, stats.ShadowSpellPower: 78}},
		{Name: "Lurker's Cord of Arcane Wrath", WowheadID: 30675, ID: -6, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeWaist, ArmorType: proto.ArmorType_ArmorTypeCloth, Phase: 1, Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Armor: 109, stats.ArcaneSpellPower: 78}},
		{Name: "Lurker's Cord of Fiery Wrath", WowheadID: 30675, ID: -7, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeWaist, ArmorType: proto.ArmorType_ArmorTypeCloth, Phase: 1, Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Armor: 109, stats.FireSpellPower: 78}},
		{Name: "Lurker's Cord of Frozen Wrath", WowheadID: 30675, ID: -8, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeWaist, ArmorType: proto.ArmorType_ArmorTypeCloth, Phase: 1, Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Armor: 109, stats.FrostSpellPower: 78}},
		{Name: "Lurker's Cord of Shadow Wrath", WowheadID: 30675, ID: -9, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeWaist, ArmorType: proto.ArmorType_ArmorTypeCloth, Phase: 1, Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Armor: 109, stats.ShadowSpellPower: 78}},
		{Name: "Lurker's Grasp of Nature's Wrath", WowheadID: 30676, ID: -10, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeWaist, ArmorType: proto.ArmorType_ArmorTypeLeather, Phase: 1, Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Armor: 205, stats.NatureSpellPower: 78}},
		{Name: "Ravager's Cuffs of Arcane Wrath", WowheadID: 30684, ID: -11, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeWrist, ArmorType: proto.ArmorType_ArmorTypeCloth, Phase: 1, Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Armor: 85, stats.ArcaneSpellPower: 58}},
		{Name: "Ravager's Cuffs of Fiery Wrath", WowheadID: 30684, ID: -12, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeWrist, ArmorType: proto.ArmorType_ArmorTypeCloth, Phase: 1, Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Armor: 85, stats.FireSpellPower: 58}},
		{Name: "Ravager's Cuffs of Frozen Wrath", WowheadID: 30684, ID: -13, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeWrist, ArmorType: proto.ArmorType_ArmorTypeCloth, Phase: 1, Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Armor: 85, stats.FrostSpellPower: 58}},
		{Name: "Ravager's Cuffs of Shadow Wrath", WowheadID: 30684, ID: -14, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeWrist, ArmorType: proto.ArmorType_ArmorTypeCloth, Phase: 1, Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Armor: 85, stats.ShadowSpellPower: 58}},
		{Name: "Ravager's Wrist-Wraps of Nature's Wrath", WowheadID: 30685, ID: -15, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeWrist, ArmorType: proto.ArmorType_ArmorTypeLeather, Phase: 1, Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Armor: 159, stats.NatureSpellPower: 58}},
		{Name: "Flawless Wand of Shadow Wrath", WowheadID: 25295, ID: -16, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeRanged, RangedWeaponType: proto.RangedWeaponType_RangedWeaponTypeWand, Phase: 1, Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.ShadowSpellPower: 25}},
		{Name: "Amber Cape of Shadow Wrath", WowheadID: 25043, ID: -17, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeBack, Phase: 1, Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.ShadowSpellPower: 45}},
		{Name: "Illidari Cape of Shadow Wrath", WowheadID: 31201, ID: -18, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeBack, Phase: 1, Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.ShadowSpellPower: 47}},
		{Name: "Elementalist Bracelets of Shadow Wrath", WowheadID: 24692, ID: -19, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeWrist, Phase: 1, Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.ShadowSpellPower: 45}},
		{Name: "Amber Cape of Shadow Wrath", WowheadID: 25043, ID: -20, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeBack, Phase: 1, Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.ShadowSpellPower: 45}},
		{Name: "Elementalist Gloves of Shadow Wrath", WowheadID: 24688, ID: -21, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeHands, Phase: 1, Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.ShadowSpellPower: 60}},
		{Name: "Nethersteel-Lined Handwraps of Shadow Wrath", WowheadID: 31166, ID: -22, Categories: []proto.ItemCategory{proto.ItemCategory_ItemCategoryCaster}, Type: proto.ItemType_ItemTypeHands, Phase: 1, Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.ShadowSpellPower: 62}},
	}...)

	for _, v := range Items {
		if _, ok := ByID[v.ID]; ok {
			fmt.Printf("Found dup item: %s\n", v.Name)
			panic("no dupes allowed")
		}
		ByName[v.Name] = v
		ByID[v.ID] = v
	}
}

type Item struct {
	ID        int32
	WowheadID int32
	Type      proto.ItemType
	ArmorType proto.ArmorType
	// Weapon Stats
	WeaponType       proto.WeaponType
	HandType         proto.HandType
	RangedWeaponType proto.RangedWeaponType
	WeaponDamageMin  float64
	WeaponDamageMax  float64
	SwingSpeed       float64

	// Used by the UI to filter which items are shown.
	Categories     []proto.ItemCategory
	ClassAllowlist []proto.Class

	Name       string
	SourceZone string
	SourceDrop string
	Stats      stats.Stats // Stats applied to wearer
	Phase      byte
	Quality    proto.ItemQuality
	Unique     bool

	GemSockets  []proto.GemColor
	SocketBonus stats.Stats

	// Modified for each instance of the item.
	Gems    []Gem
	Enchant Enchant
}

func (item Item) ToProto() *proto.Item {
	return &proto.Item{
		Id:               item.ID,
		WowheadId:        item.WowheadID,
		Name:             item.Name,
		Categories:       item.Categories[:],
		ClassAllowlist:   item.ClassAllowlist[:],
		Type:             proto.ItemType(item.Type),
		ArmorType:        proto.ArmorType(item.ArmorType),
		WeaponType:       proto.WeaponType(item.WeaponType),
		HandType:         proto.HandType(item.HandType),
		RangedWeaponType: proto.RangedWeaponType(item.RangedWeaponType),
		Stats:            item.Stats[:],
		Phase:            int32(item.Phase),
		Quality:          item.Quality,
		Unique:           item.Unique,
		GemSockets:       item.GemSockets,
		SocketBonus:      item.SocketBonus[:],
	}
}

type Enchant struct {
	ID         int32 // ID of the enchant item.
	EffectID   int32 // Used by UI to apply effect to tooltip
	Name       string
	Quality    proto.ItemQuality
	Bonus      stats.Stats
	ItemType   proto.ItemType // which slot does the enchant go on.
	HandType   proto.HandType // If ItemType is weapon, check hand type / weapon type
	WeaponType proto.WeaponType
}

func (enchant Enchant) ToProto() *proto.Enchant {
	return &proto.Enchant{
		Id:       enchant.ID,
		EffectId: enchant.EffectID,
		Name:     enchant.Name,
		Type:     enchant.ItemType,
		Stats:    enchant.Bonus[:],
		Quality:  enchant.Quality,
	}
}

type Gem struct {
	ID      int32
	Name    string
	Stats   stats.Stats // flat stats gem adds
	Color   proto.GemColor
	Phase   byte
	Quality proto.ItemQuality
	Unique  bool
	// Requirements  // Validate the gem can be used... later
}

func (gem Gem) ToProto() *proto.Gem {
	return &proto.Gem{
		Id:      gem.ID,
		Name:    gem.Name,
		Stats:   gem.Stats[:],
		Color:   gem.Color,
		Phase:   int32(gem.Phase),
		Quality: gem.Quality,
		Unique:  gem.Unique,
	}
}

type ItemSpec struct {
	ID      int32
	Enchant int32
	Gems    []int32
}

type Equipment [proto.ItemSlot_ItemSlotRanged + 1]Item

// Structs used for looking up items/gems/enchants
type EquipmentSpec [proto.ItemSlot_ItemSlotRanged + 1]ItemSpec

func ProtoToEquipmentSpec(es proto.EquipmentSpec) EquipmentSpec {
	coreEquip := EquipmentSpec{}

	for i, item := range es.Items {
		spec := ItemSpec{
			ID: item.Id,
		}
		spec.Gems = item.Gems
		spec.Enchant = item.Enchant
		coreEquip[i] = spec
	}

	return coreEquip
}

func NewEquipmentSet(equipSpec EquipmentSpec) Equipment {
	equipment := Equipment{}

	for _, itemSpec := range equipSpec {
		item := Item{}
		if foundItem, ok := ByID[itemSpec.ID]; ok {
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

		if item.Type == proto.ItemType_ItemTypeFinger {
			if equipment[ItemSlotFinger1].Name == "" {
				equipment[ItemSlotFinger1] = item
			} else {
				equipment[ItemSlotFinger2] = item
			}
		} else if item.Type == proto.ItemType_ItemTypeTrinket {
			if equipment[ItemSlotTrinket1].Name == "" {
				equipment[ItemSlotTrinket1] = item
			} else {
				equipment[ItemSlotTrinket2] = item
			}
		} else if item.Type == proto.ItemType_ItemTypeWeapon {
			if item.WeaponType == proto.WeaponType_WeaponTypeShield && equipment[ItemSlotMainHand].HandType != proto.HandType_HandTypeTwoHand {
				equipment[ItemSlotOffHand] = item
			} else if item.HandType == proto.HandType_HandTypeMainHand || item.HandType == proto.HandType_HandTypeUnknown {
				equipment[ItemSlotMainHand] = item
			} else if item.HandType == proto.HandType_HandTypeTwoHand {
				equipment[ItemSlotMainHand] = item
				equipment[ItemSlotOffHand] = Item{} // clear offhand
			} else if item.HandType == proto.HandType_HandTypeOffHand && equipment[ItemSlotMainHand].HandType != proto.HandType_HandTypeTwoHand {
				equipment[ItemSlotOffHand] = item
			} else if item.HandType == proto.HandType_HandTypeOneHand {
				if equipment[ItemSlotMainHand].ID == 0 {
					equipment[ItemSlotMainHand] = item
				} else if equipment[ItemSlotOffHand].ID == 0 {
					equipment[ItemSlotOffHand] = item
				}
			}
		} else {
			equipment[ItemTypeToSlot(item.Type)] = item
		}
	}
	return equipment
}

func ProtoToEquipment(es proto.EquipmentSpec) Equipment {
	return NewEquipmentSet(ProtoToEquipmentSpec(es))
}

// Like ItemSpec, but uses names for reference instead of ID.
type ItemStringSpec struct {
	Name    string
	Enchant string
	Gems    []string
}

func EquipmentSpecFromStrings(itemStringSpecs []ItemStringSpec) *proto.EquipmentSpec {
	eq := &proto.EquipmentSpec{
		Items: make([]*proto.ItemSpec, len(itemStringSpecs)),
	}

	for i, itemStringSpec := range itemStringSpecs {
		item := ByName[itemStringSpec.Name]
		if item.ID == 0 {
			log.Fatalf("Item not found: %#v", itemStringSpec)
		}
		itemSpec := &proto.ItemSpec{
			Id: item.ID,
		}

		if itemStringSpec.Enchant != "" {
			enchant := EnchantsByName[itemStringSpec.Enchant]
			if enchant.ID == 0 {
				log.Fatalf("Enchant not found: %s", itemStringSpec.Enchant)
			}
			itemSpec.Enchant = enchant.ID
		}

		for _, gemName := range itemStringSpec.Gems {
			gem := GemsByName[gemName]
			if gem.ID == 0 {
				log.Fatalf("Gem not found: %s", gemName)
			}
			itemSpec.Gems = append(itemSpec.Gems, gem.ID)
		}

		eq.Items[i] = itemSpec
	}
	return eq
}

func (equipment Equipment) Clone() Equipment {
	newEquipment := Equipment{}
	for idx, item := range equipment {
		newItem := item
		newEquipment[idx] = newItem
	}
	return newEquipment
}

func (equipment Equipment) Stats() stats.Stats {
	equipStats := stats.Stats{}
	for _, item := range equipment {
		equipStats = equipStats.Add(item.Stats)
		equipStats = equipStats.Add(item.Enchant.Bonus)

		for _, gem := range item.Gems {
			equipStats = equipStats.Add(gem.Stats)
		}

		// Check socket bonus
		if len(item.GemSockets) > 0 && len(item.GemSockets) == len(item.Gems) {
			allMatch := true
			for gemIndex, gem := range item.Gems {
				if !ColorIntersects(gem.Color, item.GemSockets[gemIndex]) {
					allMatch = false
					break
				}
			}

			if allMatch {
				equipStats = equipStats.Add(item.SocketBonus)
			}
		}
	}
	return equipStats
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

func ItemTypeToSlot(it proto.ItemType) ItemSlot {
	switch it {
	case proto.ItemType_ItemTypeHead:
		return ItemSlotHead
	case proto.ItemType_ItemTypeNeck:
		return ItemSlotNeck
	case proto.ItemType_ItemTypeShoulder:
		return ItemSlotShoulder
	case proto.ItemType_ItemTypeBack:
		return ItemSlotBack
	case proto.ItemType_ItemTypeChest:
		return ItemSlotChest
	case proto.ItemType_ItemTypeWrist:
		return ItemSlotWrist
	case proto.ItemType_ItemTypeHands:
		return ItemSlotHands
	case proto.ItemType_ItemTypeWaist:
		return ItemSlotWaist
	case proto.ItemType_ItemTypeLegs:
		return ItemSlotLegs
	case proto.ItemType_ItemTypeFeet:
		return ItemSlotFeet
	case proto.ItemType_ItemTypeFinger:
		return ItemSlotFinger1
	case proto.ItemType_ItemTypeTrinket:
		return ItemSlotTrinket1
	case proto.ItemType_ItemTypeWeapon:
		return ItemSlotMainHand
	case proto.ItemType_ItemTypeRanged:
		return ItemSlotRanged
	}

	return 255
}

func ColorIntersects(g proto.GemColor, o proto.GemColor) bool {
	if g == o {
		return true
	}
	if g == proto.GemColor_GemColorPrismatic || o == proto.GemColor_GemColorPrismatic {
		return true
	}
	if g == proto.GemColor_GemColorMeta {
		return false // meta gems o nothing.
	}
	if g == proto.GemColor_GemColorRed {
		return o == proto.GemColor_GemColorOrange || o == proto.GemColor_GemColorPurple
	}
	if g == proto.GemColor_GemColorBlue {
		return o == proto.GemColor_GemColorGreen || o == proto.GemColor_GemColorPurple
	}
	if g == proto.GemColor_GemColorYellow {
		return o == proto.GemColor_GemColorGreen || o == proto.GemColor_GemColorOrange
	}
	if g == proto.GemColor_GemColorOrange {
		return o == proto.GemColor_GemColorYellow || o == proto.GemColor_GemColorRed
	}
	if g == proto.GemColor_GemColorGreen {
		return o == proto.GemColor_GemColorYellow || o == proto.GemColor_GemColorBlue
	}
	if g == proto.GemColor_GemColorPurple {
		return o == proto.GemColor_GemColorBlue || o == proto.GemColor_GemColorRed
	}

	return false // dunno what else could be.
}
