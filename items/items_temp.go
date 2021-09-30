package items

import (
	"github.com/wowsims/tbc/sim/api"
	"github.com/wowsims/tbc/sim/core/stats"
)

// TODO: Create Generator for Gems and Enchants

var Gems = []Gem{
	{ID: 34220, Name: "Chaotic Skyfire Diamond", Quality: api.ItemQuality_ItemQualityRare, Phase: 1, Color: api.GemColor_GemColorMeta, Stats: stats.Stats{stats.SpellCrit: 12}}, // Activate: ActivateCSD
	{ID: 25897, Name: "Bracing Earthstorm Diamond", Quality: api.ItemQuality_ItemQualityRare, Phase: 1, Color: api.GemColor_GemColorMeta, Stats: stats.Stats{stats.SpellPower: 14}},
	{ID: 32641, Name: "Imbued Unstable Diamond", Quality: api.ItemQuality_ItemQualityRare, Phase: 1, Color: api.GemColor_GemColorMeta, Stats: stats.Stats{stats.SpellPower: 14}},
	{ID: 35503, Name: "Ember Skyfire Diamond", Quality: api.ItemQuality_ItemQualityRare, Phase: 1, Color: api.GemColor_GemColorMeta, Stats: stats.Stats{stats.SpellPower: 14}}, // Activate: ActivateESD
	{ID: 28557, Name: "Swift Starfire Diamond", Quality: api.ItemQuality_ItemQualityRare, Phase: 1, Color: api.GemColor_GemColorMeta, Stats: stats.Stats{stats.SpellPower: 12}},
	{ID: 25893, Name: "Mystical Skyfire Diamond", Quality: api.ItemQuality_ItemQualityRare, Phase: 1, Color: api.GemColor_GemColorMeta, Stats: stats.Stats{}},                         // Activate: ActivateMSD
	{ID: 25901, Name: "Insightful Earthstorm Diamond", Quality: api.ItemQuality_ItemQualityRare, Phase: 1, Color: api.GemColor_GemColorMeta, Stats: stats.Stats{stats.Intellect: 12}}, // Activate: ActivateIED
	{ID: 23096, Name: "Runed Blood Garnet", Quality: api.ItemQuality_ItemQualityUncommon, Phase: 1, Color: api.GemColor_GemColorRed, Stats: stats.Stats{stats.SpellPower: 7}},
	{ID: 24030, Name: "Runed Living Ruby", Quality: api.ItemQuality_ItemQualityRare, Phase: 1, Color: api.GemColor_GemColorRed, Stats: stats.Stats{stats.SpellPower: 9}},
	{ID: 32196, Name: "Runed Crimson Spinel", Quality: api.ItemQuality_ItemQualityEpic, Phase: 3, Color: api.GemColor_GemColorRed, Stats: stats.Stats{stats.SpellPower: 12}},
	{ID: 28118, Name: "Runed Ornate Ruby", Quality: api.ItemQuality_ItemQualityEpic, Phase: 1, Color: api.GemColor_GemColorRed, Stats: stats.Stats{stats.SpellPower: 12}},
	{ID: 33133, Name: "Don Julio's Heart", Quality: api.ItemQuality_ItemQualityEpic, Phase: 1, Color: api.GemColor_GemColorRed, Stats: stats.Stats{stats.SpellPower: 14}},
	{ID: 23121, Name: "Lustrous Azure Moonstone", Quality: api.ItemQuality_ItemQualityUncommon, Phase: 1, Color: api.GemColor_GemColorBlue, Stats: stats.Stats{stats.MP5: 2}},
	{ID: 24037, Name: "Lustrous Star of Elune", Quality: api.ItemQuality_ItemQualityRare, Phase: 1, Color: api.GemColor_GemColorBlue, Stats: stats.Stats{stats.MP5: 3}},
	{ID: 32202, Name: "Lustrous Empyrean Sapphire", Quality: api.ItemQuality_ItemQualityEpic, Phase: 1, Color: api.GemColor_GemColorBlue, Stats: stats.Stats{stats.MP5: 4}},
	{ID: 23113, Name: "Brilliant Golden Draenite", Quality: api.ItemQuality_ItemQualityUncommon, Phase: 1, Color: api.GemColor_GemColorYellow, Stats: stats.Stats{stats.Intellect: 6}},
	{ID: 24047, Name: "Brilliant Dawnstone", Quality: api.ItemQuality_ItemQualityRare, Phase: 1, Color: api.GemColor_GemColorYellow, Stats: stats.Stats{stats.Intellect: 8}},
	{ID: 32204, Name: "Brilliant Lionseye", Quality: api.ItemQuality_ItemQualityEpic, Phase: 3, Color: api.GemColor_GemColorYellow, Stats: stats.Stats{stats.Intellect: 10}},
	{ID: 23114, Name: "Gleaming Golden Draenite", Quality: api.ItemQuality_ItemQualityUncommon, Phase: 1, Color: api.GemColor_GemColorYellow, Stats: stats.Stats{stats.SpellCrit: 6}},
	{ID: 24050, Name: "Gleaming Dawnstone", Quality: api.ItemQuality_ItemQualityRare, Phase: 1, Color: api.GemColor_GemColorYellow, Stats: stats.Stats{stats.SpellCrit: 8}},
	{ID: 32207, Name: "Gleaming Lionseye", Quality: api.ItemQuality_ItemQualityEpic, Phase: 3, Color: api.GemColor_GemColorYellow, Stats: stats.Stats{stats.SpellCrit: 10}},
	{ID: 30551, Name: "Infused Fire Opal", Quality: api.ItemQuality_ItemQualityEpic, Phase: 1, Color: api.GemColor_GemColorOrange, Stats: stats.Stats{stats.Intellect: 4, stats.SpellPower: 6}},
	{ID: 23101, Name: "Potent Flame Spessarite", Quality: api.ItemQuality_ItemQualityUncommon, Phase: 1, Color: api.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellCrit: 3, stats.SpellPower: 4}},
	{ID: 24059, Name: "Potent Noble Topaz", Quality: api.ItemQuality_ItemQualityRare, Phase: 1, Color: api.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellCrit: 4, stats.SpellPower: 5}},
	{ID: 32218, Name: "Potent Pyrestone", Quality: api.ItemQuality_ItemQualityEpic, Phase: 3, Color: api.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellCrit: 5, stats.SpellPower: 6}},
	{ID: 35760, Name: "Reckless Pyrestone", Quality: api.ItemQuality_ItemQualityEpic, Phase: 3, Color: api.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellHaste: 5, stats.SpellPower: 6}},
	{ID: 30588, Name: "Potent Fire Opal", Quality: api.ItemQuality_ItemQualityEpic, Phase: 1, Color: api.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellPower: 6, stats.SpellCrit: 4}},
	{ID: 28123, Name: "Potent Ornate Topaz", Quality: api.ItemQuality_ItemQualityEpic, Phase: 1, Color: api.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellPower: 6, stats.SpellCrit: 5}},
	{ID: 31866, Name: "Veiled Flame Spessarite", Quality: api.ItemQuality_ItemQualityUncommon, Phase: 1, Color: api.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellHit: 3, stats.SpellPower: 4}},
	{ID: 31867, Name: "Veiled Noble Topaz", Quality: api.ItemQuality_ItemQualityRare, Phase: 1, Color: api.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellHit: 4, stats.SpellPower: 5}},
	{ID: 32221, Name: "Shining Fire Opal", Quality: api.ItemQuality_ItemQualityEpic, Phase: 1, Color: api.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellHit: 5, stats.SpellPower: 6}},
	{ID: 30564, Name: "Veiled Pyrestone", Quality: api.ItemQuality_ItemQualityEpic, Phase: 3, Color: api.GemColor_GemColorOrange, Stats: stats.Stats{stats.SpellHit: 5, stats.SpellPower: 6}},
	{ID: 30560, Name: "Rune Covered Chrysoprase", Quality: api.ItemQuality_ItemQualityEpic, Phase: 1, Color: api.GemColor_GemColorGreen, Stats: stats.Stats{stats.MP5: 2, stats.SpellCrit: 5}},
	{ID: 24065, Name: "Dazzling Talasite", Quality: api.ItemQuality_ItemQualityRare, Phase: 1, Color: api.GemColor_GemColorGreen, Stats: stats.Stats{stats.MP5: 2, stats.Intellect: 4}},
	{ID: 35759, Name: "Forceful Seaspray Emerald", Quality: api.ItemQuality_ItemQualityEpic, Phase: 3, Color: api.GemColor_GemColorGreen, Stats: stats.Stats{stats.SpellHaste: 5, stats.Stamina: 7}},
	{ID: 24056, Name: "Glowing Nightseye", Quality: api.ItemQuality_ItemQualityRare, Phase: 1, Color: api.GemColor_GemColorPurple, Stats: stats.Stats{stats.SpellPower: 5, stats.Stamina: 6}},
	{ID: 30555, Name: "Glowing Tanzanite", Quality: api.ItemQuality_ItemQualityEpic, Phase: 1, Color: api.GemColor_GemColorPurple, Stats: stats.Stats{stats.SpellPower: 6, stats.Stamina: 6}},
	{ID: 32215, Name: "Glowing Shadowsong Amethyst", Quality: api.ItemQuality_ItemQualityEpic, Phase: 3, Color: api.GemColor_GemColorPurple, Stats: stats.Stats{stats.SpellPower: 6, stats.Stamina: 7}},
	{ID: 31116, Name: "Infused Amethyst", Quality: api.ItemQuality_ItemQualityEpic, Phase: 1, Color: api.GemColor_GemColorPurple, Stats: stats.Stats{stats.SpellPower: 6, stats.Stamina: 6}},
	{ID: 30600, Name: "Fluorescent Tanzanite", Quality: api.ItemQuality_ItemQualityEpic, Phase: 1, Color: api.GemColor_GemColorPurple, Stats: stats.Stats{stats.SpellPower: 6, stats.Spirit: 4}},
	{ID: 30605, Name: "Vivid Chrysoprase", Quality: api.ItemQuality_ItemQualityEpic, Phase: 1, Color: api.GemColor_GemColorGreen, Stats: stats.Stats{stats.SpellHit: 5, stats.Stamina: 6}},
}

var Enchants = []Enchant{
	{ID: 29191, EffectID: 3002, Name: "Glyph of Power", Bonus: stats.Stats{stats.SpellPower: 22, stats.SpellHit: 14}, ItemType: api.ItemType_ItemTypeHead},
	{ID: 28909, EffectID: 2995, Name: "Greater Inscription of the Orb", Bonus: stats.Stats{stats.SpellPower: 12, stats.SpellCrit: 15}, ItemType: api.ItemType_ItemTypeShoulder},
	{ID: 28886, EffectID: 2982, Name: "Greater Inscription of Discipline", Bonus: stats.Stats{stats.SpellPower: 18, stats.SpellCrit: 10}, ItemType: api.ItemType_ItemTypeShoulder},
	{ID: 20076, EffectID: 2605, Name: "Zandalar Signet of Mojo", Bonus: stats.Stats{stats.SpellPower: 18}, ItemType: api.ItemType_ItemTypeShoulder},
	{ID: 23545, EffectID: 2721, Name: "Power of the Scourge", Bonus: stats.Stats{stats.SpellPower: 15, stats.SpellCrit: 14}, ItemType: api.ItemType_ItemTypeShoulder},
	{ID: 27960, EffectID: 2661, Name: "Chest - Exceptional Stats", Bonus: stats.Stats{stats.Stamina: 6, stats.Intellect: 6, stats.Spirit: 6}, ItemType: api.ItemType_ItemTypeChest},
	{ID: 22534, EffectID: 2650, Name: "Bracer - Spellpower", Bonus: stats.Stats{stats.SpellPower: 15}, ItemType: api.ItemType_ItemTypeWrist},
	{ID: 28272, EffectID: 2937, Name: "Gloves - Major Spellpower", Bonus: stats.Stats{stats.SpellPower: 20}, ItemType: api.ItemType_ItemTypeHands},
	{ID: 24274, EffectID: 2748, Name: "Runic Spellthread", Bonus: stats.Stats{stats.SpellPower: 35, stats.Stamina: 20}, ItemType: api.ItemType_ItemTypeLegs},
	{ID: 24273, EffectID: 2747, Name: "Mystic Spellthread", Bonus: stats.Stats{stats.SpellPower: 25, stats.Stamina: 15}, ItemType: api.ItemType_ItemTypeLegs},
	{ID: 22555, EffectID: 2669, Name: "Weapon - Major Spellpower", Bonus: stats.Stats{stats.SpellPower: 40}, ItemType: api.ItemType_ItemTypeWeapon},
	{ID: 35445, EffectID: 2928, Name: "Ring - Spellpower", Bonus: stats.Stats{stats.SpellPower: 12}, ItemType: api.ItemType_ItemTypeFinger},
	{ID: 27945, EffectID: 2654, Name: "Shield - Intellect", Bonus: stats.Stats{stats.Intellect: 12}, ItemType: api.ItemType_ItemTypeWeapon, WeaponType: api.WeaponType_WeaponTypeShield},
}
