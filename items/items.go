package items

import (
	"fmt"

	"github.com/wowsims/tbc/sim/api"
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
	for _, v := range Items {
		if _, ok := ByID[v.ID]; ok {
			fmt.Printf("Found dup item: %s\n", v.Name)
			panic("no dupes allowed")
		}
		if it, ok := ByName[v.Name]; ok {
			fmt.Printf("Found dup item: %s\n", v.Name)
			statsMatch := it.Type == v.Type
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
			ByName[v.Name] = v
			ByID[v.ID] = v
		}
	}
}

type Item struct {
	ID               int32
	Type             api.ItemType
	ArmorType        api.ArmorType
	WeaponType       api.WeaponType
	HandType         api.HandType
	RangedWeaponType api.RangedWeaponType

	Name       string
	SourceZone string
	SourceDrop string
	Stats      stats.Stats // Stats applied to wearer
	Phase      byte
	Quality    api.ItemQuality

	GemSockets  []api.GemColor
	SocketBonus stats.Stats

	// Modified for each instance of the item.
	Gems    []Gem
	Enchant Enchant
}

type Enchant struct {
	ID         int32 // ID of the enchant item.
	EffectID   int32 // Used by UI to apply effect to tooltip
	Name       string
	Bonus      stats.Stats
	ItemType   api.ItemType // which slot does the enchant go on.
	HandType   api.HandType // If ItemType is weapon, check hand type / weapon type
	WeaponType api.WeaponType
}

type Gem struct {
	ID      int32
	Name    string
	Stats   stats.Stats // flat stats gem adds
	Color   api.GemColor
	Phase   byte
	Quality api.ItemQuality
	// Requirements  // Validate the gem can be used... later
}

type ItemSpec struct {
	ID      int32
	Enchant int32
	Gems    []int32
}

type Equipment [api.ItemSlot_ItemSlotRanged + 1]Item

// Structs used for looking up items/gems/enchants
type EquipmentSpec [api.ItemSlot_ItemSlotRanged + 1]ItemSpec

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

		if item.Type == api.ItemType_ItemTypeFinger {
			if equipment[ItemSlotFinger1].Name == "" {
				equipment[ItemSlotFinger1] = item
			} else {
				equipment[ItemSlotFinger2] = item
			}
		} else if item.Type == api.ItemType_ItemTypeTrinket {
			if equipment[ItemSlotTrinket1].Name == "" {
				equipment[ItemSlotTrinket1] = item
			} else {
				equipment[ItemSlotTrinket2] = item
			}
		} else if item.Type == api.ItemType_ItemTypeWeapon {
			if item.WeaponType == api.WeaponType_WeaponTypeShield && equipment[ItemSlotMainHand].HandType != api.HandType_HandTypeTwoHand {
				equipment[ItemSlotOffHand] = item
			} else if item.HandType == api.HandType_HandTypeMainHand || item.HandType == api.HandType_HandTypeUnknown {
				equipment[ItemSlotMainHand] = item
			} else if item.HandType == api.HandType_HandTypeTwoHand {
				equipment[ItemSlotMainHand] = item
				equipment[ItemSlotOffHand] = Item{} // clear offhand
			} else if item.HandType == api.HandType_HandTypeOffHand && equipment[ItemSlotMainHand].HandType != api.HandType_HandTypeTwoHand {
				equipment[ItemSlotOffHand] = item
			}
		} else {
			equipment[ItemTypeToSlot(item.Type)] = item
		}
	}
	return equipment
}

func (e Equipment) Clone() Equipment {
	ne := Equipment{}
	for i, v := range e {
		vc := v
		ne[i] = vc
	}
	return ne
}

func (e Equipment) Stats() []float64 {
	s := make([]float64, api.Stat_StatArmor+1) // TODO: perhaps scan the Stat enum and find max value and cache it so if it changes this doesnt break.
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
				isMatched = isMatched && ColorIntersects(g.Color, item.GemSockets[gi])
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

func ItemTypeToSlot(it api.ItemType) ItemSlot {
	switch it {
	case api.ItemType_ItemTypeHead:
		return ItemSlotHead
	case api.ItemType_ItemTypeNeck:
		return ItemSlotNeck
	case api.ItemType_ItemTypeShoulder:
		return ItemSlotShoulder
	case api.ItemType_ItemTypeBack:
		return ItemSlotBack
	case api.ItemType_ItemTypeChest:
		return ItemSlotChest
	case api.ItemType_ItemTypeWrist:
		return ItemSlotWrist
	case api.ItemType_ItemTypeHands:
		return ItemSlotHands
	case api.ItemType_ItemTypeWaist:
		return ItemSlotWaist
	case api.ItemType_ItemTypeLegs:
		return ItemSlotLegs
	case api.ItemType_ItemTypeFeet:
		return ItemSlotFeet
	case api.ItemType_ItemTypeFinger:
		return ItemSlotFinger1
	case api.ItemType_ItemTypeTrinket:
		return ItemSlotTrinket1
	case api.ItemType_ItemTypeWeapon:
		return ItemSlotMainHand
	case api.ItemType_ItemTypeRanged:
		return ItemSlotRanged
	}

	return 255
}

func ColorIntersects(g api.GemColor, o api.GemColor) bool {
	if g == o {
		return true
	}
	if g == api.GemColor_GemColorPrismatic || o == api.GemColor_GemColorPrismatic {
		return true
	}
	if g == api.GemColor_GemColorMeta {
		return false // meta gems o nothing.
	}
	if g == api.GemColor_GemColorRed {
		return o == api.GemColor_GemColorOrange || o == api.GemColor_GemColorPurple
	}
	if g == api.GemColor_GemColorBlue {
		return o == api.GemColor_GemColorGreen || o == api.GemColor_GemColorPurple
	}
	if g == api.GemColor_GemColorYellow {
		return o == api.GemColor_GemColorGreen || o == api.GemColor_GemColorOrange
	}
	if g == api.GemColor_GemColorOrange {
		return o == api.GemColor_GemColorYellow || o == api.GemColor_GemColorRed
	}
	if g == api.GemColor_GemColorGreen {
		return o == api.GemColor_GemColorYellow || o == api.GemColor_GemColorBlue
	}
	if g == api.GemColor_GemColorPurple {
		return o == api.GemColor_GemColorBlue || o == api.GemColor_GemColorRed
	}

	return false // dunno what else could be.
}
