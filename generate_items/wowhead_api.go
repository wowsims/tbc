package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

type Stats [28]float64

type WowheadItemResponse struct {
	Name    string `json:"name"`
	Quality int    `json:"quality"`
	Icon    string `json:"icon"`
	Tooltip string `json:"tooltip"`
}

func GetRegexIntValue(srcStr string, pattern *regexp.Regexp, matchIdx int) int {
	match := pattern.FindStringSubmatch(srcStr)
	if match == nil {
		return 0
	}

	val, err := strconv.Atoi(match[matchIdx])
	if err != nil {
		return 0
	}

	return val
}
func GetBestRegexIntValue(srcStr string, patterns []*regexp.Regexp, matchIdx int) int {
	best := 0
	for _, pattern := range patterns {
		newVal := GetRegexIntValue(srcStr, pattern, matchIdx)
		if newVal > best {
			best = newVal
		}
	}
	return best
}

func (item WowheadItemResponse) GetTooltipRegexValue(pattern *regexp.Regexp, matchIdx int) int {
	return GetRegexIntValue(item.Tooltip, pattern, matchIdx)
}

func (item WowheadItemResponse) GetIntValue(pattern *regexp.Regexp) int {
	return item.GetTooltipRegexValue(pattern, 1)
}

var armorRegex = regexp.MustCompile("<!--amr-->([0-9]+) Armor")
var agilityRegex = regexp.MustCompile("<!--stat3-->\\+([0-9]+) Agility")
var strengthRegex = regexp.MustCompile("<!--stat4-->\\+([0-9]+) Strength")
var intellectRegex = regexp.MustCompile("<!--stat5-->\\+([0-9]+) Intellect")
var spiritRegex = regexp.MustCompile("<!--stat6-->\\+([0-9]+) Spirit")
var staminaRegex = regexp.MustCompile("<!--stat7-->\\+([0-9]+) Stamina")
var spellPowerRegex = regexp.MustCompile("Increases damage and healing done by magical spells and effects by up to ([0-9]+)\\.")
var healingPowerRegex = regexp.MustCompile("Increases healing done by up to ([0-9]+) and damage done by up to ([0-9]+) for all magical spells and effects\\.")
var arcaneSpellPowerRegex = regexp.MustCompile("Increases damage done by Arcane spells and effects by up to ([0-9]+)\\.")
var fireSpellPowerRegex = regexp.MustCompile("Increases damage done by Fire spells and effects by up to ([0-9]+)\\.")
var frostSpellPowerRegex = regexp.MustCompile("Increases damage done by Frost spells and effects by up to ([0-9]+)\\.")
var holySpellPowerRegex = regexp.MustCompile("Increases damage done by Holy spells and effects by up to ([0-9]+)\\.")
var natureSpellPowerRegex = regexp.MustCompile("Increases damage done by Nature spells and effects by up to ([0-9]+)\\.")
var shadowSpellPowerRegex = regexp.MustCompile("Increases damage done by Shadow spells and effects by up to ([0-9]+)\\.")
var spellHitRegex = regexp.MustCompile("Improves spell hit rating by <!--rtg18-->([0-9]+)\\.")
var spellCritRegex = regexp.MustCompile("Improves spell critical strike rating by <!--rtg21-->([0-9]+)\\.")
var spellHasteRegex = regexp.MustCompile("Improves spell haste rating by <!--rtg30-->([0-9]+)\\.")
var spellPenetrationRegex = regexp.MustCompile("Improves your spell penetration by ([0-9]+)\\.")
var mp5Regex = regexp.MustCompile("Restores ([0-9]+) mana per 5 sec\\.")
var attackPowerRegex = regexp.MustCompile("Increases attack power by ([0-9]+)\\.")
var meleeHitRegex = regexp.MustCompile("Increases your hit rating by ([0-9]+)\\.")
var meleeCritRegex = regexp.MustCompile("Increases your critical strike rating by ([0-9]+)\\.")
var meleeHasteRegex = regexp.MustCompile("Improves haste rating by <!--rtg36-->([0-9]+)\\.")
var armorPenetrationRegex = regexp.MustCompile("Your attacks ignore ([0-9]+) of your opponent's armor\\.")
var expertiseRegex = regexp.MustCompile("Increases your expertise rating by <!--rtg37-->([0-9]+)\\.")

func (item WowheadItemResponse) GetStats() Stats {
	spellPower := item.GetIntValue(spellPowerRegex)
	healingPowerFromHealing := item.GetTooltipRegexValue(healingPowerRegex, 1)
	spellPowerFromHealing := item.GetTooltipRegexValue(healingPowerRegex, 2)

	// Some items have both (e.g. Windhawk Bracers)
	spellPower = spellPower + spellPowerFromHealing
	healingPower := spellPower + healingPowerFromHealing

	return Stats{
		proto.Stat_StatArmor:            float64(item.GetIntValue(armorRegex)),
		proto.Stat_StatStrength:         float64(item.GetIntValue(strengthRegex)),
		proto.Stat_StatAgility:          float64(item.GetIntValue(agilityRegex)),
		proto.Stat_StatStamina:          float64(item.GetIntValue(staminaRegex)),
		proto.Stat_StatIntellect:        float64(item.GetIntValue(intellectRegex)),
		proto.Stat_StatSpirit:           float64(item.GetIntValue(spiritRegex)),
		proto.Stat_StatSpellPower:       float64(spellPower),
		proto.Stat_StatHealingPower:     float64(healingPower),
		proto.Stat_StatArcaneSpellPower: float64(item.GetIntValue(arcaneSpellPowerRegex)),
		proto.Stat_StatFireSpellPower:   float64(item.GetIntValue(fireSpellPowerRegex)),
		proto.Stat_StatFrostSpellPower:  float64(item.GetIntValue(frostSpellPowerRegex)),
		proto.Stat_StatHolySpellPower:   float64(item.GetIntValue(holySpellPowerRegex)),
		proto.Stat_StatNatureSpellPower: float64(item.GetIntValue(natureSpellPowerRegex)),
		proto.Stat_StatShadowSpellPower: float64(item.GetIntValue(shadowSpellPowerRegex)),
		proto.Stat_StatSpellHit:         float64(item.GetIntValue(spellHitRegex)),
		proto.Stat_StatSpellCrit:        float64(item.GetIntValue(spellCritRegex)),
		proto.Stat_StatSpellHaste:       float64(item.GetIntValue(spellHasteRegex)),
		proto.Stat_StatSpellPenetration: float64(item.GetIntValue(spellPenetrationRegex)),
		proto.Stat_StatMP5:              float64(item.GetIntValue(mp5Regex)),
		proto.Stat_StatAttackPower:      float64(item.GetIntValue(attackPowerRegex)),
		proto.Stat_StatMeleeHit:         float64(item.GetIntValue(meleeHitRegex)),
		proto.Stat_StatMeleeCrit:        float64(item.GetIntValue(meleeCritRegex)),
		proto.Stat_StatMeleeHaste:       float64(item.GetIntValue(meleeHasteRegex)),
		proto.Stat_StatArmorPenetration: float64(item.GetIntValue(armorPenetrationRegex)),
		proto.Stat_StatExpertise:        float64(item.GetIntValue(expertiseRegex)),
	}
}

// Detects class-locked items, e.g. tier sets and pvp gear.
var classPatterns = map[proto.Class]*regexp.Regexp{
	proto.Class_ClassWarrior: regexp.MustCompile("<a href=\\\"/class=1\\\" class=\\\"c1\\\">Warrior</a>"),
	proto.Class_ClassPaladin: regexp.MustCompile("<a href=\\\"/class=2\\\" class=\\\"c2\\\">Paladin</a>"),
	proto.Class_ClassHunter:  regexp.MustCompile("<a href=\\\"/class=3\\\" class=\\\"c3\\\">Hunter</a>"),
	proto.Class_ClassRogue:   regexp.MustCompile("<a href=\\\"/class=4\\\" class=\\\"c4\\\">Rogue</a>"),
	proto.Class_ClassPriest:  regexp.MustCompile("<a href=\\\"/class=5\\\" class=\\\"c5\\\">Priest</a>"),
	proto.Class_ClassShaman:  regexp.MustCompile("<a href=\\\"/class=7\\\" class=\\\"c7\\\">Shaman</a>"),
	proto.Class_ClassMage:    regexp.MustCompile("<a href=\\\"/class=8\\\" class=\\\"c8\\\">Mage</a>"),
	proto.Class_ClassWarlock: regexp.MustCompile("<a href=\\\"/class=9\\\" class=\\\"c9\\\">Warlock</a>"),
	proto.Class_ClassDruid:   regexp.MustCompile("<a href=\\\"/class=11\\\" class=\\\"c11\\\">Druid</a>"),
}

func (item WowheadItemResponse) GetClassAllowlist() []proto.Class {
	var allowlist []proto.Class

	for class, pattern := range classPatterns {
		if pattern.MatchString(item.Tooltip) {
			allowlist = append(allowlist, class)
		}
	}

	return allowlist
}

var phaseRegex = regexp.MustCompile("Phase ([0-9])")
func (item WowheadItemResponse) GetPhase() int {
	return item.GetIntValue(phaseRegex)
}

var uniqueRegex = regexp.MustCompile("Unique")
func (item WowheadItemResponse) GetUnique() bool {
	return uniqueRegex.MatchString(item.Tooltip)
}

var itemTypePatterns = map[proto.ItemType]*regexp.Regexp{
	proto.ItemType_ItemTypeHead:     regexp.MustCompile("<td>Head</td>"),
	proto.ItemType_ItemTypeNeck:     regexp.MustCompile("<td>Neck</td>"),
	proto.ItemType_ItemTypeShoulder: regexp.MustCompile("<td>Shoulder</td>"),
	proto.ItemType_ItemTypeBack:     regexp.MustCompile("<td>Back</td>"),
	proto.ItemType_ItemTypeChest:    regexp.MustCompile("<td>Chest</td>"),
	proto.ItemType_ItemTypeWrist:    regexp.MustCompile("<td>Wrist</td>"),
	proto.ItemType_ItemTypeHands:    regexp.MustCompile("<td>Hands</td>"),
	proto.ItemType_ItemTypeWaist:    regexp.MustCompile("<td>Waist</td>"),
	proto.ItemType_ItemTypeLegs:     regexp.MustCompile("<td>Legs</td>"),
	proto.ItemType_ItemTypeFeet:     regexp.MustCompile("<td>Feet</td>"),
	proto.ItemType_ItemTypeFinger:   regexp.MustCompile("<td>Finger</td>"),
	proto.ItemType_ItemTypeTrinket:  regexp.MustCompile("<td>Trinket</td>"),
	proto.ItemType_ItemTypeWeapon:   regexp.MustCompile("<td>((Main Hand)|(Two-Hand)|(One-Hand)|(Off Hand)|(Held In Off-hand))</td>"),
	proto.ItemType_ItemTypeRanged:   regexp.MustCompile("<td>(Ranged|Thrown|Relic)</td>"),
}

func (item WowheadItemResponse) GetItemType() proto.ItemType {
	for itemType, pattern := range itemTypePatterns {
		if pattern.MatchString(item.Tooltip) {
			return itemType
		}
	}
	panic("Could not find item type from tooltip: " + item.Tooltip)
}

var armorTypePatterns = map[proto.ArmorType]*regexp.Regexp{
	proto.ArmorType_ArmorTypeCloth:   regexp.MustCompile("<span class=\\\"q1\\\">Cloth</span>"),
	proto.ArmorType_ArmorTypeLeather: regexp.MustCompile("<span class=\\\"q1\\\">Leather</span>"),
	proto.ArmorType_ArmorTypeMail:    regexp.MustCompile("<span class=\\\"q1\\\">Mail</span>"),
	proto.ArmorType_ArmorTypePlate:   regexp.MustCompile("<span class=\\\"q1\\\">Plate</span>"),
}

func (item WowheadItemResponse) GetArmorType() proto.ArmorType {
	for armorType, pattern := range armorTypePatterns {
		if pattern.MatchString(item.Tooltip) {
			return armorType
		}
	}
	return proto.ArmorType_ArmorTypeUnknown
}

var weaponTypePatterns = map[proto.WeaponType]*regexp.Regexp{
	proto.WeaponType_WeaponTypeAxe:     regexp.MustCompile("<span class=\\\"q1\\\">Axe</span>"),
	proto.WeaponType_WeaponTypeDagger:  regexp.MustCompile("<span class=\\\"q1\\\">Dagger</span>"),
	proto.WeaponType_WeaponTypeFist:    regexp.MustCompile("<span class=\\\"q1\\\">Fist Weapon</span>"),
	proto.WeaponType_WeaponTypeMace:    regexp.MustCompile("<span class=\\\"q1\\\">Mace</span>"),
	proto.WeaponType_WeaponTypeOffHand: regexp.MustCompile("<td>Held In Off-hand</td>"),
	proto.WeaponType_WeaponTypePolearm: regexp.MustCompile("<span class=\\\"q1\\\">Polearm</span>"),
	proto.WeaponType_WeaponTypeShield:  regexp.MustCompile("<span class=\\\"q1\\\">Shield</span>"),
	proto.WeaponType_WeaponTypeStaff:   regexp.MustCompile("<span class=\\\"q1\\\">Staff</span>"),
	proto.WeaponType_WeaponTypeSword:   regexp.MustCompile("<span class=\\\"q1\\\">Sword</span>"),
}

func (item WowheadItemResponse) GetWeaponType() proto.WeaponType {
	for weaponType, pattern := range weaponTypePatterns {
		if pattern.MatchString(item.Tooltip) {
			return weaponType
		}
	}
	return proto.WeaponType_WeaponTypeUnknown
}

var handTypePatterns = map[proto.HandType]*regexp.Regexp{
	proto.HandType_HandTypeMainHand: regexp.MustCompile("<td>Main Hand</td>"),
	proto.HandType_HandTypeOneHand:  regexp.MustCompile("<td>One-Hand</td>"),
	proto.HandType_HandTypeOffHand:  regexp.MustCompile("<td>((Off Hand)|(Held In Off-hand))</td>"),
	proto.HandType_HandTypeTwoHand:  regexp.MustCompile("<td>Two-Hand</td>"),
}

func (item WowheadItemResponse) GetHandType() proto.HandType {
	for handType, pattern := range handTypePatterns {
		if pattern.MatchString(item.Tooltip) {
			return handType
		}
	}
	return proto.HandType_HandTypeUnknown
}

var rangedWeaponTypePatterns = map[proto.RangedWeaponType]*regexp.Regexp{
	proto.RangedWeaponType_RangedWeaponTypeBow:      regexp.MustCompile("<span class=\\\"q1\\\">Bow</span>"),
	proto.RangedWeaponType_RangedWeaponTypeCrossbow: regexp.MustCompile("<span class=\\\"q1\\\">Crossbow</span>"),
	proto.RangedWeaponType_RangedWeaponTypeGun:      regexp.MustCompile("<span class=\\\"q1\\\">Gun</span>"),
	proto.RangedWeaponType_RangedWeaponTypeIdol:     regexp.MustCompile("<span class=\\\"q1\\\">Idol</span>"),
	proto.RangedWeaponType_RangedWeaponTypeLibram:   regexp.MustCompile("<span class=\\\"q1\\\">Libram</span>"),
	proto.RangedWeaponType_RangedWeaponTypeThrown:   regexp.MustCompile("<span class=\\\"q1\\\">Thrown</span>"),
	proto.RangedWeaponType_RangedWeaponTypeTotem:    regexp.MustCompile("<span class=\\\"q1\\\">Totem</span>"),
	proto.RangedWeaponType_RangedWeaponTypeWand:     regexp.MustCompile("<span class=\\\"q1\\\">Wand</span>"),
}

func (item WowheadItemResponse) GetRangedWeaponType() proto.RangedWeaponType {
	for rangedWeaponType, pattern := range rangedWeaponTypePatterns {
		if pattern.MatchString(item.Tooltip) {
			return rangedWeaponType
		}
	}
	return proto.RangedWeaponType_RangedWeaponTypeUnknown
}

var gemColorsRegex, _ = regexp.Compile("(Meta|Yellow|Blue|Red) Socket")

func (item WowheadItemResponse) GetGemSockets() []proto.GemColor {
	matches := gemColorsRegex.FindAllStringSubmatch(item.Tooltip, -1)
	if matches == nil {
		return []proto.GemColor{}
	}

	numSockets := len(matches)
	gemColors := make([]proto.GemColor, numSockets)
	for socketIdx, match := range matches {
		gemColorName := "GemColor" + match[1]
		gemColors[socketIdx] = proto.GemColor(proto.GemColor_value[gemColorName])
	}
	return gemColors
}

var socketBonusRegex = regexp.MustCompile("<span class=\\\"q0\\\">Socket Bonus: (.*?)</span>")
var strengthSocketBonusRegexes = []*regexp.Regexp{regexp.MustCompile("\\+([0-9]+) Strength")}
var agilitySocketBonusRegexes = []*regexp.Regexp{regexp.MustCompile("\\+([0-9]+) Agility")}
var staminaSocketBonusRegexes = []*regexp.Regexp{regexp.MustCompile("\\+([0-9]+) Stamina")}
var intellectSocketBonusRegexes = []*regexp.Regexp{regexp.MustCompile("\\+([0-9]+) Intellect")}
var spiritSocketBonusRegexes = []*regexp.Regexp{regexp.MustCompile("\\+([0-9]+) Spirit")}
var spellPowerSocketBonusRegexes = []*regexp.Regexp{regexp.MustCompile("\\+([0-9]+) Spell Damage and Healing")}
var healingPowerSocketBonusRegexes = []*regexp.Regexp{
	regexp.MustCompile("\\+([0-9]+) Healing and \\+([0-9]+) Spell Damage"),
	regexp.MustCompile("\\+([0-9]+) Healing \\+([0-9]+) Spell Damage"),
}
var spellHitSocketBonusRegexes = []*regexp.Regexp{regexp.MustCompile("\\+([0-9]+) Spell Hit Rating")}
var spellCritSocketBonusRegexes = []*regexp.Regexp{regexp.MustCompile("\\+([0-9]+) Spell Critical Strike Rating")}
var mp5SocketBonusRegexes = []*regexp.Regexp{
	regexp.MustCompile("([0-9]+) Mana per 5 sec"),
	regexp.MustCompile("([0-9]+) mana per 5 sec"),
}
var attackPowerSocketBonusRegexes = []*regexp.Regexp{regexp.MustCompile("\\+([0-9]+) Attack Power")}
var meleeHitSocketBonusRegexes = []*regexp.Regexp{regexp.MustCompile("\\+([0-9]+) Hit Rating")}
var meleeCritSocketBonusRegexes = []*regexp.Regexp{regexp.MustCompile("\\+([0-9]+) Critical Strike Rating")}

func (item WowheadItemResponse) GetSocketBonus() Stats {
	match := socketBonusRegex.FindStringSubmatch(item.Tooltip)
	if match == nil {
		return Stats{}
	}

	bonusStr := match[1]
	//fmt.Printf("\n%s\n", bonusStr)

	stats := Stats{
		proto.Stat_StatStrength:    float64(GetBestRegexIntValue(bonusStr, strengthSocketBonusRegexes, 1)),
		proto.Stat_StatAgility:     float64(GetBestRegexIntValue(bonusStr, agilitySocketBonusRegexes, 1)),
		proto.Stat_StatStamina:     float64(GetBestRegexIntValue(bonusStr, staminaSocketBonusRegexes, 1)),
		proto.Stat_StatIntellect:   float64(GetBestRegexIntValue(bonusStr, intellectSocketBonusRegexes, 1)),
		proto.Stat_StatSpirit:      float64(GetBestRegexIntValue(bonusStr, spiritSocketBonusRegexes, 1)),
		proto.Stat_StatSpellHit:    float64(GetBestRegexIntValue(bonusStr, spellHitSocketBonusRegexes, 1)),
		proto.Stat_StatSpellCrit:   float64(GetBestRegexIntValue(bonusStr, spellCritSocketBonusRegexes, 1)),
		proto.Stat_StatMP5:         float64(GetBestRegexIntValue(bonusStr, mp5SocketBonusRegexes, 1)),
		proto.Stat_StatAttackPower: float64(GetBestRegexIntValue(bonusStr, attackPowerSocketBonusRegexes, 1)),
		proto.Stat_StatMeleeHit:    float64(GetBestRegexIntValue(bonusStr, meleeHitSocketBonusRegexes, 1)),
		proto.Stat_StatMeleeCrit:   float64(GetBestRegexIntValue(bonusStr, meleeCritSocketBonusRegexes, 1)),
	}

	spellPower := GetBestRegexIntValue(bonusStr, spellPowerSocketBonusRegexes, 1)
	healingPower := GetBestRegexIntValue(bonusStr, healingPowerSocketBonusRegexes, 1)
	spellPowerFromHealing := GetBestRegexIntValue(bonusStr, healingPowerSocketBonusRegexes, 2)

	stats[proto.Stat_StatSpellPower] = math.Max(float64(spellPower), float64(spellPowerFromHealing))
	stats[proto.Stat_StatHealingPower] = math.Max(float64(spellPower), float64(healingPower))

	return stats
}

func getWowheadItemResponse(itemId int) WowheadItemResponse {
	url := fmt.Sprintf("https://tbc.wowhead.com/tooltip/item/%d", itemId)

	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	result, err := httpClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer result.Body.Close()

	resultBody, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf(string(resultBody))
	itemResponse := WowheadItemResponse{}
	err = json.Unmarshal(resultBody, &itemResponse)
	if err != nil {
		log.Fatal(err)
	}

	return itemResponse
}
