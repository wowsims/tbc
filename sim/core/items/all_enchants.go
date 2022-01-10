package items

import (
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var Enchants = []Enchant{
	{ID: 29191, EffectID: 3002, Name: "Glyph of Power", Quality: proto.ItemQuality_ItemQualityUncommon, Bonus: stats.Stats{stats.SpellPower: 22, stats.SpellHit: 14}, ItemType: proto.ItemType_ItemTypeHead},
	{ID: 28909, EffectID: 2995, Name: "Greater Inscription of the Orb", Quality: proto.ItemQuality_ItemQualityRare, Bonus: stats.Stats{stats.SpellPower: 12, stats.SpellCrit: 15}, ItemType: proto.ItemType_ItemTypeShoulder},
	{ID: 28886, EffectID: 2982, Name: "Greater Inscription of Discipline", Quality: proto.ItemQuality_ItemQualityRare, Bonus: stats.Stats{stats.SpellPower: 18, stats.SpellCrit: 10}, ItemType: proto.ItemType_ItemTypeShoulder},
	{ID: 20076, EffectID: 2605, Name: "Zandalar Signet of Mojo", Quality: proto.ItemQuality_ItemQualityRare, Bonus: stats.Stats{stats.SpellPower: 18}, ItemType: proto.ItemType_ItemTypeShoulder},
	{ID: 23545, EffectID: 2721, Name: "Power of the Scourge", Quality: proto.ItemQuality_ItemQualityEpic, Bonus: stats.Stats{stats.SpellPower: 15, stats.SpellCrit: 14}, ItemType: proto.ItemType_ItemTypeShoulder},
	{ID: 22546, EffectID: 2660, Name: "Chest - Exceptional Mana", Quality: proto.ItemQuality_ItemQualityCommon, Bonus: stats.Stats{stats.Mana: 150}, ItemType: proto.ItemType_ItemTypeChest},
	{ID: 24003, EffectID: 2661, Name: "Chest - Exceptional Stats", Quality: proto.ItemQuality_ItemQualityCommon, Bonus: stats.Stats{stats.Stamina: 6, stats.Intellect: 6, stats.Spirit: 6, stats.Strength: 6, stats.Agility: 6}, ItemType: proto.ItemType_ItemTypeChest},
	{ID: 33990, EffectID: 1144, Name: "Chest - Major Spirit", IsSpellID: true, Quality: proto.ItemQuality_ItemQualityCommon, Bonus: stats.Stats{stats.Spirit: 15}, ItemType: proto.ItemType_ItemTypeChest},
	{ID: 33991, EffectID: 3150, Name: "Chest - Restore Mana Prime", IsSpellID: true, Quality: proto.ItemQuality_ItemQualityCommon, Bonus: stats.Stats{stats.MP5: 6}, ItemType: proto.ItemType_ItemTypeChest},
	{ID: 22534, EffectID: 2650, Name: "Bracer - Spellpower", Quality: proto.ItemQuality_ItemQualityUncommon, Bonus: stats.Stats{stats.SpellPower: 15}, ItemType: proto.ItemType_ItemTypeWrist},
	{ID: 28271, EffectID: 2935, Name: "Gloves - Spell Strike", Quality: proto.ItemQuality_ItemQualityUncommon, Bonus: stats.Stats{stats.SpellHit: 15}, ItemType: proto.ItemType_ItemTypeHands},
	{ID: 28272, EffectID: 2937, Name: "Gloves - Major Spellpower", Quality: proto.ItemQuality_ItemQualityUncommon, Bonus: stats.Stats{stats.SpellPower: 20}, ItemType: proto.ItemType_ItemTypeHands},
	{ID: 24274, EffectID: 2748, Name: "Runic Spellthread", Quality: proto.ItemQuality_ItemQualityEpic, Bonus: stats.Stats{stats.SpellPower: 35, stats.Stamina: 20}, ItemType: proto.ItemType_ItemTypeLegs},
	{ID: 24273, EffectID: 2747, Name: "Mystic Spellthread", Quality: proto.ItemQuality_ItemQualityRare, Bonus: stats.Stats{stats.SpellPower: 25, stats.Stamina: 15}, ItemType: proto.ItemType_ItemTypeLegs},
	{ID: 22551, EffectID: 2666, Name: "Weapon - Major Intellect", Quality: proto.ItemQuality_ItemQualityUncommon, Bonus: stats.Stats{stats.Intellect: 30}, ItemType: proto.ItemType_ItemTypeWeapon},
	{ID: 22555, EffectID: 2669, Name: "Weapon - Major Spellpower", Quality: proto.ItemQuality_ItemQualityUncommon, Bonus: stats.Stats{stats.SpellPower: 40}, ItemType: proto.ItemType_ItemTypeWeapon},
	{ID: 22560, EffectID: 2671, Name: "Sunfire", Quality: proto.ItemQuality_ItemQualityRare, Bonus: stats.Stats{stats.ArcaneSpellPower: 50, stats.FireSpellPower: 50}, ItemType: proto.ItemType_ItemTypeWeapon},
	{ID: 22561, EffectID: 2672, Name: "Soulfrost", Quality: proto.ItemQuality_ItemQualityRare, Bonus: stats.Stats{stats.FrostSpellPower: 50, stats.ShadowSpellPower: 50}, ItemType: proto.ItemType_ItemTypeWeapon},
	{ID: 22536, EffectID: 2928, Name: "Ring - Spellpower", Quality: proto.ItemQuality_ItemQualityCommon, Bonus: stats.Stats{stats.SpellPower: 12}, ItemType: proto.ItemType_ItemTypeFinger},
	{ID: 22539, EffectID: 2654, Name: "Shield - Intellect", Quality: proto.ItemQuality_ItemQualityCommon, Bonus: stats.Stats{stats.Intellect: 12}, ItemType: proto.ItemType_ItemTypeWeapon, EnchantType: proto.EnchantType_EnchantTypeShield},
	{ID: 35297, EffectID: 2940, Name: "Enchant Boots - Boar's Speed", Quality: proto.ItemQuality_ItemQualityRare, Bonus: stats.Stats{stats.Stamina: 9}, ItemType: proto.ItemType_ItemTypeFeet},
	{ID: 35298, EffectID: 2656, Name: "Enchant Boots - Vitality", Quality: proto.ItemQuality_ItemQualityUncommon, Bonus: stats.Stats{stats.MP5: 4}, ItemType: proto.ItemType_ItemTypeFeet},
	{ID: 33150, EffectID: 2621, Name: "Enchant Cloak - Subtlety", Quality: proto.ItemQuality_ItemQualityRare, Bonus: stats.Stats{}, ItemType: proto.ItemType_ItemTypeBack},

	{ID: 29192, EffectID: 3003, Name: "Glyph of Ferocity", Quality: proto.ItemQuality_ItemQualityUncommon, Bonus: stats.Stats{stats.AttackPower: 34, stats.MeleeHit: 16}, ItemType: proto.ItemType_ItemTypeHead},
	{ID: 28888, EffectID: 2986, Name: "Greater Inscription of Vengeance", Quality: proto.ItemQuality_ItemQualityRare, Bonus: stats.Stats{stats.AttackPower: 30, stats.MeleeCrit: 10}, ItemType: proto.ItemType_ItemTypeShoulder},
	{ID: 29535, EffectID: 3012, Name: "Nethercobra Leg Armor", Quality: proto.ItemQuality_ItemQualityEpic, Bonus: stats.Stats{stats.AttackPower: 50, stats.MeleeCrit: 12}, ItemType: proto.ItemType_ItemTypeLegs},
	{ID: 27899, EffectID: 2647, Name: "Bracer - Brawn", IsSpellID: true, Quality: proto.ItemQuality_ItemQualityUncommon, Bonus: stats.Stats{stats.Strength: 12}, ItemType: proto.ItemType_ItemTypeWrist},
	{ID: 33995, EffectID: 684, Name: "Gloves - Major Strength", IsSpellID: true, Quality: proto.ItemQuality_ItemQualityUncommon, Bonus: stats.Stats{stats.Strength: 15}, ItemType: proto.ItemType_ItemTypeHands},
	{ID: 28279, EffectID: 2939, Name: "Enchant Boots - Cat's Swiftness", Quality: proto.ItemQuality_ItemQualityRare, Bonus: stats.Stats{stats.Agility: 6}, ItemType: proto.ItemType_ItemTypeFeet},
	{ID: 22559, EffectID: 2673, Name: "Weapon - Mongoose", Quality: proto.ItemQuality_ItemQualityRare, Bonus: stats.Stats{}, ItemType: proto.ItemType_ItemTypeWeapon},
}
