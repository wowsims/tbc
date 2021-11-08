package items

import (
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// TODO: Create Generator for Gems and Enchants

var Gems = []Gem{
	{ID: 34220, Name: "Chaotic Skyfire Diamond", Quality: proto.ItemQuality_ItemQualityRare, Phase: 1, Color: proto.GemColor_GemColorMeta, Stats: stats.Stats{stats.SpellCrit: 12}},
	{ID: 25897, Name: "Bracing Earthstorm Diamond", Quality: proto.ItemQuality_ItemQualityRare, Phase: 1, Color: proto.GemColor_GemColorMeta, Stats: stats.Stats{stats.SpellPower: 14}},
	{ID: 32641, Name: "Imbued Unstable Diamond", Quality: proto.ItemQuality_ItemQualityRare, Phase: 1, Color: proto.GemColor_GemColorMeta, Stats: stats.Stats{stats.SpellPower: 14}},
	{ID: 35503, Name: "Ember Skyfire Diamond", Quality: proto.ItemQuality_ItemQualityRare, Phase: 1, Color: proto.GemColor_GemColorMeta, Stats: stats.Stats{stats.SpellPower: 14}},
	{ID: 28557, Name: "Swift Starfire Diamond", Quality: proto.ItemQuality_ItemQualityRare, Phase: 1, Color: proto.GemColor_GemColorMeta, Stats: stats.Stats{stats.SpellPower: 12}},
	{ID: 25893, Name: "Mystical Skyfire Diamond", Quality: proto.ItemQuality_ItemQualityRare, Phase: 1, Color: proto.GemColor_GemColorMeta, Stats: stats.Stats{}},
	{ID: 25901, Name: "Insightful Earthstorm Diamond", Quality: proto.ItemQuality_ItemQualityRare, Phase: 1, Color: proto.GemColor_GemColorMeta, Stats: stats.Stats{stats.Intellect: 12}},
	{ID: 23096, Name: "Runed Blood Garnet", Quality: proto.ItemQuality_ItemQualityUncommon, Phase: 1, Color: proto.GemColor_GemColorRed, Stats: stats.Stats{stats.SpellPower: 7}},
	{ID: 24030, Name: "Runed Living Ruby", Quality: proto.ItemQuality_ItemQualityRare, Phase: 1, Color: proto.GemColor_GemColorRed, Stats: stats.Stats{stats.SpellPower: 9}},
	{ID: 32196, Name: "Runed Crimson Spinel", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 3, Color: proto.GemColor_GemColorRed, Stats: stats.Stats{stats.SpellPower: 12}},
	{ID: 28118, Name: "Runed Ornate Ruby", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 1, Color: proto.GemColor_GemColorRed, Stats: stats.Stats{stats.SpellPower: 12}, Unique: true},
	{ID: 33133, Name: "Don Julio's Heart", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 1, Color: proto.GemColor_GemColorRed, Stats: stats.Stats{stats.SpellPower: 14}, Unique: true},
	{ID: 23121, Name: "Lustrous Azure Moonstone", Quality: proto.ItemQuality_ItemQualityUncommon, Phase: 1, Color: proto.GemColor_GemColorBlue, Stats: stats.Stats{stats.MP5: 2}},
	{ID: 24037, Name: "Lustrous Star of Elune", Quality: proto.ItemQuality_ItemQualityRare, Phase: 1, Color: proto.GemColor_GemColorBlue, Stats: stats.Stats{stats.MP5: 3}},
	{ID: 32202, Name: "Lustrous Empyrean Sapphire", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 1, Color: proto.GemColor_GemColorBlue, Stats: stats.Stats{stats.MP5: 4}},
	{ID: 23113, Name: "Brilliant Golden Draenite", Quality: proto.ItemQuality_ItemQualityUncommon, Phase: 1, Color: proto.GemColor_GemColorYellow, Stats: stats.Stats{stats.Intellect: 6}},
	{ID: 24047, Name: "Brilliant Dawnstone", Quality: proto.ItemQuality_ItemQualityRare, Phase: 1, Color: proto.GemColor_GemColorYellow, Stats: stats.Stats{stats.Intellect: 8}},
	{ID: 32204, Name: "Brilliant Lionseye", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 3, Color: proto.GemColor_GemColorYellow, Stats: stats.Stats{stats.Intellect: 10}},
	{ID: 23114, Name: "Gleaming Golden Draenite", Quality: proto.ItemQuality_ItemQualityUncommon, Phase: 1, Color: proto.GemColor_GemColorYellow, Stats: stats.Stats{stats.SpellCrit: 6}},
	{ID: 24050, Name: "Gleaming Dawnstone", Quality: proto.ItemQuality_ItemQualityRare, Phase: 1, Color: proto.GemColor_GemColorYellow, Stats: stats.Stats{stats.SpellCrit: 8}},
	{ID: 32207, Name: "Gleaming Lionseye", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 3, Color: proto.GemColor_GemColorYellow, Stats: stats.Stats{stats.SpellCrit: 10}},
	{ID: 30551, Name: "Infused Fire Opal", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 1, Color: proto.GemColor_GemColorOrange, Stats: stats.Stats{stats.Intellect: 4, stats.SpellPower: 6}, Unique: true},
	{ID: 23101, Name: "Potent Flame Spessarite", Quality: proto.ItemQuality_ItemQualityUncommon, Phase: 1, Color: proto.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellCrit: 3, stats.SpellPower: 4}},
	{ID: 24059, Name: "Potent Noble Topaz", Quality: proto.ItemQuality_ItemQualityRare, Phase: 1, Color: proto.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellCrit: 4, stats.SpellPower: 5}},
	{ID: 32218, Name: "Potent Pyrestone", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 3, Color: proto.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellCrit: 5, stats.SpellPower: 6}},
	{ID: 35760, Name: "Reckless Pyrestone", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 3, Color: proto.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellHaste: 5, stats.SpellPower: 6}},
	{ID: 30588, Name: "Potent Fire Opal", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 1, Color: proto.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellPower: 6, stats.SpellCrit: 4}, Unique: true},
	{ID: 28123, Name: "Potent Ornate Topaz", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 1, Color: proto.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellPower: 6, stats.SpellCrit: 5}, Unique: true},
	{ID: 31866, Name: "Veiled Flame Spessarite", Quality: proto.ItemQuality_ItemQualityUncommon, Phase: 1, Color: proto.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellHit: 3, stats.SpellPower: 4}},
	{ID: 31867, Name: "Veiled Noble Topaz", Quality: proto.ItemQuality_ItemQualityRare, Phase: 1, Color: proto.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellHit: 4, stats.SpellPower: 5}},
	{ID: 32221, Name: "Shining Fire Opal", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 1, Color: proto.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellHit: 5, stats.SpellPower: 6}},
	{ID: 30564, Name: "Veiled Pyrestone", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 3, Color: proto.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellHit: 5, stats.SpellPower: 6}},
	{ID: 30560, Name: "Rune Covered Chrysoprase", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 1, Color: proto.GemColor_GemColorGreen, Stats: stats.Stats{stats.MP5: 2, stats.SpellCrit: 5}, Unique: true},
	{ID: 24065, Name: "Dazzling Talasite", Quality: proto.ItemQuality_ItemQualityRare, Phase: 1, Color: proto.GemColor_GemColorGreen, Stats: stats.Stats{stats.MP5: 2, stats.Intellect: 4}},
	{ID: 35759, Name: "Forceful Seaspray Emerald", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 3, Color: proto.GemColor_GemColorGreen, Stats: stats.Stats{stats.SpellHaste: 5, stats.Stamina: 7}},
	{ID: 24056, Name: "Glowing Nightseye", Quality: proto.ItemQuality_ItemQualityRare, Phase: 1, Color: proto.GemColor_GemColorPurple, Stats: stats.Stats{stats.SpellPower: 5, stats.Stamina: 6}},
	{ID: 30555, Name: "Glowing Tanzanite", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 1, Color: proto.GemColor_GemColorPurple, Stats: stats.Stats{stats.SpellPower: 6, stats.Stamina: 6}, Unique: true},
	{ID: 32215, Name: "Glowing Shadowsong Amethyst", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 3, Color: proto.GemColor_GemColorPurple, Stats: stats.Stats{stats.SpellPower: 6, stats.Stamina: 7}},
	{ID: 31116, Name: "Infused Amethyst", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 1, Color: proto.GemColor_GemColorPurple, Stats: stats.Stats{stats.SpellPower: 6, stats.Stamina: 6}, Unique: true},
	{ID: 30600, Name: "Fluorescent Tanzanite", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 1, Color: proto.GemColor_GemColorPurple, Stats: stats.Stats{stats.SpellPower: 6, stats.Spirit: 4}, Unique: true},
	{ID: 30605, Name: "Vivid Chrysoprase", Quality: proto.ItemQuality_ItemQualityEpic, Phase: 1, Color: proto.GemColor_GemColorGreen, Stats: stats.Stats{stats.SpellHit: 5, stats.Stamina: 6}, Unique: true},
}

var Enchants = []Enchant{
	{ID: 29191, EffectID: 3002, Name: "Glyph of Power", Quality: proto.ItemQuality_ItemQualityUncommon, Bonus: stats.Stats{stats.SpellPower: 22, stats.SpellHit: 14}, ItemType: proto.ItemType_ItemTypeHead},
	{ID: 28909, EffectID: 2995, Name: "Greater Inscription of the Orb", Quality: proto.ItemQuality_ItemQualityRare, Bonus: stats.Stats{stats.SpellPower: 12, stats.SpellCrit: 15}, ItemType: proto.ItemType_ItemTypeShoulder},
	{ID: 28886, EffectID: 2982, Name: "Greater Inscription of Discipline", Quality: proto.ItemQuality_ItemQualityRare, Bonus: stats.Stats{stats.SpellPower: 18, stats.SpellCrit: 10}, ItemType: proto.ItemType_ItemTypeShoulder},
	{ID: 20076, EffectID: 2605, Name: "Zandalar Signet of Mojo", Quality: proto.ItemQuality_ItemQualityRare, Bonus: stats.Stats{stats.SpellPower: 18}, ItemType: proto.ItemType_ItemTypeShoulder},
	{ID: 23545, EffectID: 2721, Name: "Power of the Scourge", Quality: proto.ItemQuality_ItemQualityEpic, Bonus: stats.Stats{stats.SpellPower: 15, stats.SpellCrit: 14}, ItemType: proto.ItemType_ItemTypeShoulder},
	{ID: 24003, EffectID: 2661, Name: "Chest - Exceptional Stats", Quality: proto.ItemQuality_ItemQualityCommon, Bonus: stats.Stats{stats.Stamina: 6, stats.Intellect: 6, stats.Spirit: 6}, ItemType: proto.ItemType_ItemTypeChest},
	{ID: 22534, EffectID: 2650, Name: "Bracer - Spellpower", Quality: proto.ItemQuality_ItemQualityUncommon, Bonus: stats.Stats{stats.SpellPower: 15}, ItemType: proto.ItemType_ItemTypeWrist},
	{ID: 28272, EffectID: 2937, Name: "Gloves - Major Spellpower", Quality: proto.ItemQuality_ItemQualityUncommon, Bonus: stats.Stats{stats.SpellPower: 20}, ItemType: proto.ItemType_ItemTypeHands},
	{ID: 24274, EffectID: 2748, Name: "Runic Spellthread", Quality: proto.ItemQuality_ItemQualityEpic, Bonus: stats.Stats{stats.SpellPower: 35, stats.Stamina: 20}, ItemType: proto.ItemType_ItemTypeLegs},
	{ID: 24273, EffectID: 2747, Name: "Mystic Spellthread", Quality: proto.ItemQuality_ItemQualityRare, Bonus: stats.Stats{stats.SpellPower: 25, stats.Stamina: 15}, ItemType: proto.ItemType_ItemTypeLegs},
	{ID: 22555, EffectID: 2669, Name: "Weapon - Major Spellpower", Quality: proto.ItemQuality_ItemQualityUncommon, Bonus: stats.Stats{stats.SpellPower: 40}, ItemType: proto.ItemType_ItemTypeWeapon},
	{ID: 22536, EffectID: 2928, Name: "Ring - Spellpower", Quality: proto.ItemQuality_ItemQualityCommon, Bonus: stats.Stats{stats.SpellPower: 12}, ItemType: proto.ItemType_ItemTypeFinger},
	{ID: 22539, EffectID: 2654, Name: "Shield - Intellect", Quality: proto.ItemQuality_ItemQualityCommon, Bonus: stats.Stats{stats.Intellect: 12}, ItemType: proto.ItemType_ItemTypeWeapon, WeaponType: proto.WeaponType_WeaponTypeShield},
}
