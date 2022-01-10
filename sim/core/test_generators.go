package core

import (
	"fmt"
	"strings"

	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	googleProto "google.golang.org/protobuf/proto"
)

type RaceCombo struct {
	Label string
	Race  proto.Race
}
type GearSetCombo struct {
	Label   string
	GearSet *proto.EquipmentSpec
}
type SpecOptionsCombo struct {
	Label       string
	SpecOptions interface{}
}
type BuffsCombo struct {
	Label    string
	Raid     *proto.RaidBuffs
	Party    *proto.PartyBuffs
	Player   *proto.IndividualBuffs
	Consumes *proto.Consumes
}
type EncounterCombo struct {
	Label     string
	Encounter *proto.Encounter
}
type SettingsCombos struct {
	Class       proto.Class
	Races       []RaceCombo
	GearSets    []GearSetCombo
	SpecOptions []SpecOptionsCombo
	Buffs       []BuffsCombo
	Encounters  []EncounterCombo
	SimOptions  *proto.SimOptions
}

func (combos *SettingsCombos) NumTests() int {
	return len(combos.Races) * len(combos.GearSets) * len(combos.SpecOptions) * len(combos.Buffs) * len(combos.Encounters)
}

func (combos *SettingsCombos) GetTest(testIdx int) (string, *proto.RaidSimRequest) {
	testNameParts := []string{}

	raceIdx := testIdx % len(combos.Races)
	testIdx /= len(combos.Races)
	raceCombo := combos.Races[raceIdx]
	testNameParts = append(testNameParts, raceCombo.Label)

	gearSetIdx := testIdx % len(combos.GearSets)
	testIdx /= len(combos.GearSets)
	gearSetCombo := combos.GearSets[gearSetIdx]
	testNameParts = append(testNameParts, gearSetCombo.Label)

	specOptionsIdx := testIdx % len(combos.SpecOptions)
	testIdx /= len(combos.SpecOptions)
	specOptionsCombo := combos.SpecOptions[specOptionsIdx]
	testNameParts = append(testNameParts, specOptionsCombo.Label)

	buffsIdx := testIdx % len(combos.Buffs)
	testIdx /= len(combos.Buffs)
	buffsCombo := combos.Buffs[buffsIdx]
	testNameParts = append(testNameParts, buffsCombo.Label)

	encounterIdx := testIdx % len(combos.Encounters)
	testIdx /= len(combos.Encounters)
	encounterCombo := combos.Encounters[encounterIdx]
	testNameParts = append(testNameParts, encounterCombo.Label)

	rsr := &proto.RaidSimRequest{
		Raid: SinglePlayerRaidProto(
			WithSpec(&proto.Player{
				Race:      raceCombo.Race,
				Class:     combos.Class,
				Equipment: gearSetCombo.GearSet,
				Consumes:  buffsCombo.Consumes,
				Buffs:     buffsCombo.Player,
				// TODO: Allow cooldowns in tests
				//Cooldowns: &proto.Cooldowns{
				//	Cooldowns: []*proto.Cooldown{
				//		&proto.Cooldown{
				//			Id: &proto.ActionID{
				//				RawId: &proto.ActionID_SpellId{
				//					SpellId: 12043,
				//				},
				//			},
				//			Timings: []float64{
				//				5,
				//			},
				//		},
				//	},
				//},
			}, specOptionsCombo.SpecOptions),
			buffsCombo.Party,
			buffsCombo.Raid),
		Encounter:  encounterCombo.Encounter,
		SimOptions: combos.SimOptions,
	}

	return strings.Join(testNameParts, "-"), rsr
}

// Returns all items that meet the given conditions.
type ItemFilter struct {
	// If set to ClassUnknown, any class is fine.
	Class proto.Class

	// Blank list allows any value. Otherwise item must match 1 value from the list.
	ArmorTypes        []proto.ArmorType
	WeaponTypes       []proto.WeaponType
	HandTypes         []proto.HandType
	RangedWeaponTypes []proto.RangedWeaponType

	// Item IDs to ignore.
	IDBlacklist []int32
}

// Returns whether the given item matches the conditions of this filter.
//
// If equipChecksOnly is true, will only check conditions related to whether
// the item is equippable.
func (filter *ItemFilter) Matches(item items.Item, equipChecksOnly bool) bool {
	if filter.Class != proto.Class_ClassUnknown && len(item.ClassAllowlist) > 0 {
		found := false
		for _, class := range item.ClassAllowlist {
			if class == filter.Class {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if item.Type == proto.ItemType_ItemTypeWeapon {
		if len(filter.WeaponTypes) > 0 {
			found := false
			for _, weaponType := range filter.WeaponTypes {
				if weaponType == item.WeaponType {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}

		if len(filter.HandTypes) > 0 {
			found := false
			for _, handType := range filter.HandTypes {
				if handType == item.HandType {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	} else if item.Type == proto.ItemType_ItemTypeRanged {
		if len(filter.RangedWeaponTypes) > 0 {
			found := false
			for _, rangedWeaponType := range filter.RangedWeaponTypes {
				if rangedWeaponType == item.RangedWeaponType {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	} else {
		if len(filter.ArmorTypes) > 0 {
			found := false
			for _, armorType := range filter.ArmorTypes {
				if armorType == item.ArmorType {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}

	if !equipChecksOnly {
		if !HasItemEffect(item.ID) {
			return false
		}

		if len(filter.IDBlacklist) > 0 {
			for _, itemID := range filter.IDBlacklist {
				if itemID == item.ID {
					return false
				}
			}
		}
	}

	return true
}

func (filter *ItemFilter) FindAllItems() []items.Item {
	filteredItems := []items.Item{}

	for _, item := range items.ByID {
		if filter.Matches(item, false) {
			filteredItems = append(filteredItems, item)
		}
	}

	return filteredItems
}

func (filter *ItemFilter) FindAllSets() []ItemSet {
	filteredSets := []ItemSet{}

	for _, set := range GetAllItemSets() {
		firstItem := items.ByID[set.ItemIDs()[0]]
		if filter.Matches(firstItem, true) {
			filteredSets = append(filteredSets, set)
		}
	}

	return filteredSets
}

func (filter *ItemFilter) FindAllMetaGems() []items.Gem {
	filteredGems := []items.Gem{}

	for _, gem := range items.GemsByID {
		if gem.Color == proto.GemColor_GemColorMeta {
			filteredGems = append(filteredGems, gem)
		}
	}

	return filteredGems
}

type ItemsTestGenerator struct {
	// Fields describing the base API request.
	Player     *proto.Player
	PartyBuffs *proto.PartyBuffs
	RaidBuffs  *proto.RaidBuffs
	Encounter  *proto.Encounter
	SimOptions *proto.SimOptions

	// Some fields are populated automatically.
	ItemFilter ItemFilter

	initialized bool

	items []items.Item
	sets  []ItemSet

	metagems []items.Gem

	metaSocketIdx int
}

func (generator *ItemsTestGenerator) init() {
	if generator.initialized {
		return
	}
	generator.initialized = true

	generator.ItemFilter.Class = generator.Player.Class
	if generator.ItemFilter.IDBlacklist == nil {
		generator.ItemFilter.IDBlacklist = []int32{}
	}
	for _, itemSpec := range generator.Player.Equipment.Items {
		generator.ItemFilter.IDBlacklist = append(generator.ItemFilter.IDBlacklist, itemSpec.Id)
	}

	generator.items = generator.ItemFilter.FindAllItems()
	generator.sets = generator.ItemFilter.FindAllSets()

	baseEquipment := items.ProtoToEquipment(*generator.Player.Equipment)
	metaSocketIdx := -1
	for i, socketColor := range baseEquipment[proto.ItemSlot_ItemSlotHead].GemSockets {
		if socketColor == proto.GemColor_GemColorMeta {
			metaSocketIdx = i
			break
		}
	}
	if metaSocketIdx == -1 {
		panic("Please use a base head item with a meta socket so we can test meta effects!")
	}
	generator.metagems = generator.ItemFilter.FindAllMetaGems()
}

func (generator *ItemsTestGenerator) NumTests() int {
	generator.init()
	return len(generator.items) + len(generator.sets) + len(generator.metagems)
}

func (generator *ItemsTestGenerator) GetTest(testIdx int) (string, *proto.RaidSimRequest) {
	generator.init()
	label := ""

	playerCopy := googleProto.Clone(generator.Player).(*proto.Player)
	equipment := items.ProtoToEquipment(*playerCopy.Equipment)
	if testIdx < len(generator.items) {
		testItem := generator.items[testIdx]
		equipment.EquipItem(generator.items[testIdx])
		label = fmt.Sprintf("%s-%d", strings.ReplaceAll(testItem.Name, " ", ""), testItem.ID)
	} else if testIdx < len(generator.items)+len(generator.sets) {
		testSet := generator.sets[testIdx-len(generator.items)]
		for _, itemID := range testSet.ItemIDs() {
			setItem := items.ByID[itemID]
			equipment.EquipItem(setItem)
		}
		label = strings.ReplaceAll(testSet.Name, " ", "")
	} else {
		testMetaGem := generator.metagems[testIdx-len(generator.items)-len(generator.sets)]
		equipment[proto.ItemSlot_ItemSlotHead].Gems[generator.metaSocketIdx] = testMetaGem
		label = strings.ReplaceAll(testMetaGem.Name, " ", "")
	}
	playerCopy.Equipment = equipment.ToEquipmentSpecProto()

	rsr := &proto.RaidSimRequest{
		Raid: SinglePlayerRaidProto(
			playerCopy,
			generator.PartyBuffs,
			generator.RaidBuffs),
		Encounter:  generator.Encounter,
		SimOptions: generator.SimOptions,
	}

	return label, rsr
}
