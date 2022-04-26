import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Flask } from '/tbc/core/proto/common.js';
import { Food } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { WeaponImbue } from '/tbc/core/proto/common.js';
import { Hunter_Rotation as HunterRotation, Hunter_Rotation_WeaveType as WeaveType, Hunter_Options as HunterOptions, Hunter_Options_Ammo as Ammo, Hunter_Options_QuiverBonus as QuiverBonus, Hunter_Options_PetType as PetType, } from '/tbc/core/proto/hunter.js';
import * as Enchants from '/tbc/core/constants/enchants.js';
import * as Gems from '/tbc/core/proto_utils/gems.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';
// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.
// Default talents. Uses the wowhead calculator format, make the talents on
// https://tbc.wowhead.com/talent-calc and copy the numbers in the url.
export const BeastMasteryTalents = {
    name: 'BM',
    data: '512002015150122431051-0505201205',
};
export const MarksmanTalents = {
    name: 'Marksman',
    data: '51200200502-0551201205013253135',
};
export const SurvivalTalents = {
    name: 'Survival',
    data: '502-0550201205-333200022003223005103',
};
export const DefaultRotation = HunterRotation.create({
    useMultiShot: true,
    useArcaneShot: true,
    viperStartManaPercent: 0.1,
    viperStopManaPercent: 0.3,
    weave: WeaveType.WeaveFull,
    timeToWeaveMs: 500,
    percentWeaved: 0.8,
});
export const DefaultOptions = HunterOptions.create({
    quiverBonus: QuiverBonus.Speed15,
    ammo: Ammo.AdamantiteStinger,
    petType: PetType.Ravager,
    petUptime: 1,
    latencyMs: 30,
});
export const DefaultConsumes = Consumes.create({
    defaultPotion: Potions.HastePotion,
    flask: Flask.FlaskOfRelentlessAssault,
    food: Food.FoodGrilledMudfish,
    mainHandImbue: WeaponImbue.WeaponImbueAdamantiteSharpeningStone,
    offHandImbue: WeaponImbue.WeaponImbueAdamantiteSharpeningStone,
});
export const P1_BM_PRESET = {
    name: 'P1 BM Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getTalentTree() != 2,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 28275,
                enchant: Enchants.GLYPH_OF_FEROCITY,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.RELENTLESS_EARTHSTORM_DIAMOND,
                ],
            }),
            ItemSpec.create({
                id: 29381, // Choker of Vile Intent
            }),
            ItemSpec.create({
                id: 27801,
                enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
                gems: [
                    Gems.WICKED_NOBLE_TOPAZ,
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 24259,
                enchant: Enchants.CLOAK_GREATER_AGILITY,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 28228,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.SHIFTING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 29246,
                enchant: Enchants.WRIST_ASSAULT,
            }),
            ItemSpec.create({
                id: 27474,
                enchant: Enchants.GLOVES_MAJOR_AGILITY,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 28828,
                gems: [
                    Gems.SHIFTING_NIGHTSEYE,
                    Gems.WICKED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 30739,
                enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 28545,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.GLINTING_NOBLE_TOPAZ,
                ],
                enchant: Enchants.FEET_CATS_SWIFTNESS,
            }),
            ItemSpec.create({
                id: 28757, // Ring of a Thousand Marks
            }),
            ItemSpec.create({
                id: 28791, // Ring of the Recalcitrant
            }),
            ItemSpec.create({
                id: 28830, // Dragonspine Trophy
            }),
            ItemSpec.create({
                id: 29383, // Bloodlust Brooch
            }),
            ItemSpec.create({
                id: 28435,
                enchant: Enchants.WEAPON_2H_MAJOR_AGILITY,
            }),
            ItemSpec.create({
                id: 28772,
                enchant: Enchants.STABILIZED_ETERNIUM_SCOPE,
            }),
        ],
    }),
};
export const P2_BM_PRESET = {
    name: 'P2 BM Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getTalentTree() != 2,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 30141,
                enchant: Enchants.GLYPH_OF_FEROCITY,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.RELENTLESS_EARTHSTORM_DIAMOND,
                ],
            }),
            ItemSpec.create({
                id: 30017, // Telonicus's Pendant of Mayhem
            }),
            ItemSpec.create({
                id: 30143,
                enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 29994,
                enchant: Enchants.CLOAK_GREATER_AGILITY,
            }),
            ItemSpec.create({
                id: 30139,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.SHIFTING_NIGHTSEYE,
                    Gems.WICKED_NOBLE_TOPAZ,
                    Gems.WICKED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 29966,
                enchant: Enchants.WRIST_ASSAULT,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 30140,
                enchant: Enchants.GLOVES_MAJOR_AGILITY,
            }),
            ItemSpec.create({
                id: 30040,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 29995,
                enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
            }),
            ItemSpec.create({
                id: 30104,
                gems: [
                    Gems.SHIFTING_NIGHTSEYE,
                    Gems.DELICATE_LIVING_RUBY,
                ],
                enchant: Enchants.FEET_CATS_SWIFTNESS,
            }),
            ItemSpec.create({
                id: 29997, // Band of the Ranger-General
            }),
            ItemSpec.create({
                id: 28791, // Ring of the Recalcitrant
            }),
            ItemSpec.create({
                id: 28830, // Dragonspine Trophy
            }),
            ItemSpec.create({
                id: 29383, // Bloodlust Brooch
            }),
            ItemSpec.create({
                id: 29993,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                ],
                enchant: Enchants.WEAPON_2H_MAJOR_AGILITY,
            }),
            ItemSpec.create({
                id: 30105,
                enchant: Enchants.STABILIZED_ETERNIUM_SCOPE,
            }),
        ],
    }),
};
export const P3_BM_PRESET = {
    name: 'P3 BM Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getTalentTree() != 2,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 32235,
                enchant: Enchants.GLYPH_OF_FEROCITY,
                gems: [
                    Gems.RELENTLESS_EARTHSTORM_DIAMOND,
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 32591, // Choker of Serrated Blades
            }),
            ItemSpec.create({
                id: 31006,
                enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 32323,
                enchant: Enchants.CLOAK_GREATER_AGILITY,
            }),
            ItemSpec.create({
                id: 31004,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                    Gems.JAGGED_SEASPRAY_EMERALD,
                    Gems.JAGGED_SEASPRAY_EMERALD,
                ],
            }),
            ItemSpec.create({
                id: 32324,
                enchant: Enchants.WRIST_ASSAULT,
                gems: [
                    Gems.WICKED_PYRESTONE,
                ],
            }),
            ItemSpec.create({
                id: 31001,
                enchant: Enchants.GLOVES_MAJOR_AGILITY,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 30879,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 31005,
                enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 32366,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                    Gems.WICKED_PYRESTONE,
                ],
                enchant: Enchants.FEET_CATS_SWIFTNESS,
            }),
            ItemSpec.create({
                id: 29997, // Band of the Ranger-General
            }),
            ItemSpec.create({
                id: 29301, // Band of the Eternal Champion
            }),
            ItemSpec.create({
                id: 28830, // Dragonspine Trophy
            }),
            ItemSpec.create({
                id: 32505, // Madness of the Betrayer
            }),
            ItemSpec.create({
                id: 30901,
                enchant: Enchants.WEAPON_GREATER_AGILITY,
            }),
            ItemSpec.create({
                id: 30881,
                enchant: Enchants.WEAPON_GREATER_AGILITY,
            }),
            ItemSpec.create({
                id: 30906,
                enchant: Enchants.STABILIZED_ETERNIUM_SCOPE,
            }),
        ],
    }),
};
export const P1_SV_PRESET = {
    name: 'P1 SV Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getTalentTree() == 2,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 28275,
                enchant: Enchants.GLYPH_OF_FEROCITY,
                gems: [
                    Gems.GLINTING_NOBLE_TOPAZ,
                    Gems.RELENTLESS_EARTHSTORM_DIAMOND,
                ],
            }),
            ItemSpec.create({
                id: 28343, // Jagged Bark Pendant
            }),
            ItemSpec.create({
                id: 27801,
                enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 29382,
                enchant: Enchants.CLOAK_GREATER_AGILITY,
            }),
            ItemSpec.create({
                id: 28228,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.SHIFTING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 25697,
                enchant: Enchants.WRIST_ASSAULT,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 27474,
                enchant: Enchants.GLOVES_MAJOR_AGILITY,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 28750,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 28741,
                enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 28545,
                enchant: Enchants.FEET_DEXTERITY,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 31277, // Pathfinder's Band
            }),
            ItemSpec.create({
                id: 28791, // Ring of the Recalcitrant
            }),
            ItemSpec.create({
                id: 28830, // Dragonspine Trophy
            }),
            ItemSpec.create({
                id: 29383, // Bloodlust Brooch
            }),
            ItemSpec.create({
                id: 27846,
                enchant: Enchants.WEAPON_AGILITY,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 28572,
                enchant: Enchants.WEAPON_AGILITY,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.GLINTING_NOBLE_TOPAZ,
                    Gems.SHIFTING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 28772,
                enchant: Enchants.STABILIZED_ETERNIUM_SCOPE,
            }),
        ],
    }),
};
export const P2_SV_PRESET = {
    name: 'P2 SV Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getTalentTree() == 2,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 30141,
                enchant: Enchants.GLYPH_OF_FEROCITY,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.RELENTLESS_EARTHSTORM_DIAMOND,
                ],
            }),
            ItemSpec.create({
                id: 30017, // Telonicus's Pendant of Mayhem
            }),
            ItemSpec.create({
                id: 30143,
                enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 29994,
                enchant: Enchants.CLOAK_GREATER_AGILITY,
            }),
            ItemSpec.create({
                id: 30054,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 29966,
                enchant: Enchants.WRIST_ASSAULT,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 28506,
                enchant: Enchants.GLOVES_MAJOR_AGILITY,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 30040,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 29985,
                enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.GLINTING_NOBLE_TOPAZ,
                    Gems.SHIFTING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 30104,
                enchant: Enchants.FEET_DEXTERITY,
                gems: [
                    Gems.JAGGED_TALASITE,
                    Gems.DELICATE_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 29298, // Band of Eternity
            }),
            ItemSpec.create({
                id: 28791, // Ring of the Recalcitrant
            }),
            ItemSpec.create({
                id: 28830, // Dragonspine Trophy
            }),
            ItemSpec.create({
                id: 29383, // Bloodlust Brooch
            }),
            ItemSpec.create({
                id: 29924,
                enchant: Enchants.WEAPON_AGILITY,
            }),
            ItemSpec.create({
                id: 29948,
                enchant: Enchants.WEAPON_AGILITY,
            }),
            ItemSpec.create({
                id: 30105,
                enchant: Enchants.STABILIZED_ETERNIUM_SCOPE,
            }),
        ],
    }),
};
export const P3_SV_PRESET = {
    name: 'P3 SV Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getTalentTree() == 2,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 31003,
                enchant: Enchants.GLYPH_OF_FEROCITY,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                    Gems.RELENTLESS_EARTHSTORM_DIAMOND,
                ],
            }),
            ItemSpec.create({
                id: 30017, // Telonicus's Pendant of Mayhem
            }),
            ItemSpec.create({
                id: 31006,
                enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 29994,
                enchant: Enchants.CLOAK_GREATER_AGILITY,
            }),
            ItemSpec.create({
                id: 31004,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                    Gems.JAGGED_SEASPRAY_EMERALD,
                    Gems.JAGGED_SEASPRAY_EMERALD,
                ],
            }),
            ItemSpec.create({
                id: 32324,
                enchant: Enchants.WRIST_ASSAULT,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 31001,
                enchant: Enchants.GLOVES_MAJOR_AGILITY,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 30879,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 30900,
                enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                    Gems.DELICATE_CRIMSON_SPINEL,
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 32366,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
                enchant: Enchants.FEET_DEXTERITY,
            }),
            ItemSpec.create({
                id: 28791, // Ring of the Recalcitrant
            }),
            ItemSpec.create({
                id: 29301, // Band of the Eternal Champion
            }),
            ItemSpec.create({
                id: 28830, // Dragonspine Trophy
            }),
            ItemSpec.create({
                id: 32505, // Madness of the Betrayer
            }),
            ItemSpec.create({
                id: 30881,
                enchant: Enchants.WEAPON_GREATER_AGILITY,
            }),
            ItemSpec.create({
                id: 30881,
                enchant: Enchants.WEAPON_GREATER_AGILITY,
            }),
            ItemSpec.create({
                id: 30906,
                enchant: Enchants.STABILIZED_ETERNIUM_SCOPE,
            }),
        ],
    }),
};
export const P4_BM_PRESET = {
    name: 'P4 BM Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getTalentTree() != 2,
    gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 32235,
			"enchant": 29192,
			"gems": [
				32409,
				32194
			]
		},
		{
			"id": 32260
		},
		{
			"id": 31006,
			"enchant": 28888,
			"gems": [
				32222,
				32212
			]
		},
		{
			"id": 32323,
			"enchant": 34004
		},
		{
			"id": 31004,
			"enchant": 24003,
			"gems": [
				32194,
				32222,
				32226
			]
		},
		{
			"id": 32324,
			"enchant": 34002,
			"gems": [
				32222
			]
		},
		{
			"id": 31001,
			"enchant": 19445,
			"gems": [
				32194
			]
		},
		{
			"id": 32346
		},
		{
			"id": 31005,
			"enchant": 29535,
			"gems": [
				32194
			]
		},
		{
			"id": 32366,
			"enchant": 22544,
			"gems": [
				32194,
				32222
			]
		},
		{
			"id": 29301
		},
		{
			"id": 33496
		},
		{
			"id": 28830
		},
		{
			"id": 33831
		},
		{
			"id": 33389,
			"enchant": 33165
		},
		{
			"id": 33389,
			"enchant": 33165
		},
		{
			"id": 30906,
			"enchant": 23766
		}
	]}`),
};
export const P4_SV_PRESET = {
    name: 'P4 SV Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getTalentTree() == 2,
    gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 32235,
			"enchant": 29192,
			"gems": [
				32409,
				32194
			]
		},
		{
			"id": 30017
		},
		{
			"id": 31006,
			"enchant": 28888,
			"gems": [
				32194,
				32194
			]
		},
		{
			"id": 29994,
			"enchant": 34004
		},
		{
			"id": 31004,
			"enchant": 24003,
			"gems": [
				32194,
				32226,
				32226
			]
		},
		{
			"id": 32324,
			"enchant": 34002,
			"gems": [
				32194
			]
		},
		{
			"id": 31001,
			"enchant": 19445,
			"gems": [
				32194
			]
		},
		{
			"id": 30879,
			"gems": [
				32194,
				32194
			]
		},
		{
			"id": 31005,
			"enchant": 29535,
			"gems": [
				32194
			]
		},
		{
			"id": 32366,
			"enchant": 22544,
			"gems": [
				32194,
				32194
			]
		},
		{
			"id": 33496
		},
		{
			"id": 29301
		},
		{
			"id": 33831
		},
		{
			"id": 28830
		},
		{
			"id": 33389,
			"enchant": 33165
		},
		{
			"id": 33389,
			"enchant": 33165
		},
		{
			"id": 30906,
			"enchant": 23766
		}
	]}`),
};
export const P5_BM_PRESET = {
    name: 'P5 BM Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getTalentTree() != 2,
    gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 34333,
			"enchant": 29192,
			"gems": [
				33131,
				32409
			]
		},
		{
			"id": 34358,
			"gems": [
				32205
			]
		},
		{
			"id": 31006,
			"enchant": 23548,
			"gems": [
				32205,
				32212
			]
		},
		{
			"id": 34241,
			"enchant": 34004,
			"gems": [
				32220
			]
		},
		{
			"id": 34397,
			"enchant": 24003,
			"gems": [
				32212,
				33143,
				32194
			]
		},
		{
			"id": 34443,
			"enchant": 34002,
			"gems": [
				32194
			]
		},
		{
			"id": 34370,
			"enchant": 19445,
			"gems": [
				32194,
				32194
			]
		},
		{
			"id": 34549,
			"gems": [
				32194
			]
		},
		{
			"id": 34188,
			"enchant": 29535,
			"gems": [
				32194,
				32194,
				32194
			]
		},
		{
			"id": 34570,
			"enchant": 22544,
			"gems": [
				32220
			]
		},
		{
			"id": 34189
		},
		{
			"id": 34361
		},
		{
			"id": 34427
		},
		{
			"id": 33831
		},
		{
			"id": 34329,
			"enchant": 33165,
			"gems": [
				32194
			]
		},
		{
			"id": 34329,
			"enchant": 33165,
			"gems": [
				32194
			]
		},
		{
			"id": 34334,
			"enchant": 23766
		}
	]}`),
};
export const P5_SV_PRESET = {
    name: 'P5 SV Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getTalentTree() == 2,
    gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 34333,
			"enchant": 29192,
			"gems": [
				32194,
				32409
			]
		},
		{
			"id": 34177
		},
		{
			"id": 31006,
			"enchant": 23548,
			"gems": [
				32194,
				32194
			]
		},
		{
			"id": 34241,
			"enchant": 34004,
			"gems": [
				32194
			]
		},
		{
			"id": 34397,
			"enchant": 24003,
			"gems": [
				32194,
				32194,
				32194
			]
		},
		{
			"id": 34443,
			"enchant": 34002,
			"gems": [
				32194
			]
		},
		{
			"id": 34343,
			"enchant": 19445,
			"gems": [
				32194,
				32226
			]
		},
		{
			"id": 34549,
			"gems": [
				32194
			]
		},
		{
			"id": 34188,
			"enchant": 29535,
			"gems": [
				32194,
				32194,
				32194
			]
		},
		{
			"id": 34570,
			"enchant": 22544,
			"gems": [
				32226
			]
		},
		{
			"id": 34887
		},
		{
			"id": 34361
		},
		{
			"id": 34427
		},
		{
			"id": 28830
		},
		{
			"id": 34183,
			"enchant": 22556,
			"gems": [
				32194
			]
		},
		{
			"id": 34334,
			"enchant": 23766
		}
	]}`),
};
