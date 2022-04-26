import { Consumes } from '/tbc/core/proto/common.js';
import { Flask } from '/tbc/core/proto/common.js';
import { Food } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { WeaponImbue } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { BalanceDruid_Rotation as BalanceDruidRotation, BalanceDruid_Options as BalanceDruidOptions } from '/tbc/core/proto/druid.js';
import { BalanceDruid_Rotation_PrimarySpell as PrimarySpell } from '/tbc/core/proto/druid.js';
import * as Enchants from '/tbc/core/constants/enchants.js';
import * as Gems from '/tbc/core/proto_utils/gems.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';
// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.
// Default talents. Uses the wowhead calculator format, make the talents on
// https://tbc.wowhead.com/talent-calc and copy the numbers in the url.
export const StandardTalents = {
    name: 'Standard',
    data: '510022312503135231351--520033',
};
export const DefaultRotation = BalanceDruidRotation.create({
    primarySpell: PrimarySpell.Adaptive,
    faerieFire: true,
});
export const DefaultOptions = BalanceDruidOptions.create({});
export const DefaultConsumes = Consumes.create({
    flask: Flask.FlaskOfBlindingLight,
    food: Food.FoodBlackenedBasilisk,
    mainHandImbue: WeaponImbue.WeaponImbueBrilliantWizardOil,
    defaultPotion: Potions.SuperManaPotion,
});
export const P1_ALLIANCE_PRESET = {
    name: 'P1 Alliance Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getFaction() == Faction.Alliance,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 29093,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.CHAOTIC_SKYFIRE_DIAMOND,
                ],
            }),
            ItemSpec.create({
                id: 28762, // Adornment of Stolen Souls
            }),
            ItemSpec.create({
                id: 29095,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.GLOWING_NIGHTSEYE,
                    Gems.VEILED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28766,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 21848,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.GLOWING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 24250,
                enchant: Enchants.WRIST_SPELLPOWER,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 21847,
                enchant: Enchants.GLOVES_SPELLPOWER,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.VEILED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 21846,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.VEILED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 24262,
                enchant: Enchants.RUNIC_SPELLTHREAD,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.VEILED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28517,
                enchant: Enchants.BOARS_SPEED,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 28753,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29287,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29370, // Icon of the Silver Crescent
            }),
            ItemSpec.create({
                id: 27683, // Quagmirran's Eye
            }),
            ItemSpec.create({
                id: 28770,
                enchant: Enchants.SUNFIRE,
            }),
            ItemSpec.create({
                id: 29271, // Talisman of Kalecgos
            }),
            ItemSpec.create({
                id: 27518, // Ivory Idol of the Moongodddess
            }),
        ],
    }),
};
export const P1_HORDE_PRESET = {
    name: 'P1 Horde Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getFaction() == Faction.Horde,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 29093,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.CHAOTIC_SKYFIRE_DIAMOND,
                ],
            }),
            ItemSpec.create({
                id: 28762, // Adornment of Stolen Souls
            }),
            ItemSpec.create({
                id: 29095,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.GLOWING_NIGHTSEYE,
                    Gems.POTENT_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28766,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 21848,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.POTENT_NOBLE_TOPAZ,
                    Gems.GLOWING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 24250,
                enchant: Enchants.WRIST_SPELLPOWER,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 21847,
                enchant: Enchants.GLOVES_SPELLPOWER,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.VEILED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 21846,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.VEILED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 24262,
                enchant: Enchants.RUNIC_SPELLTHREAD,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.VEILED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28517,
                enchant: Enchants.BOARS_SPEED,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.POTENT_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28753,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 28793,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29370, // Icon of the Silver Crescent
            }),
            ItemSpec.create({
                id: 27683, // Quagmirran's Eye
            }),
            ItemSpec.create({
                id: 28770,
                enchant: Enchants.SUNFIRE,
            }),
            ItemSpec.create({
                id: 29271, // Talisman of Kalecgos
            }),
            ItemSpec.create({
                id: 27518, // Ivory Idol of the Moongodddess
            }),
        ],
    }),
};
export const P2_ALLIANCE_PRESET = {
    name: 'P2 Alliance Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getFaction() == Faction.Alliance,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 30233,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.POTENT_NOBLE_TOPAZ,
                    Gems.CHAOTIC_SKYFIRE_DIAMOND,
                ],
            }),
            ItemSpec.create({
                id: 30015, // The Sun King's Talisman
            }),
            ItemSpec.create({
                id: 30235,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.GLOWING_NIGHTSEYE,
                    Gems.POTENT_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28797,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 30231,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 29918,
                enchant: Enchants.WRIST_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 30232,
                enchant: Enchants.GLOVES_SPELLPOWER,
                gems: [],
            }),
            ItemSpec.create({
                id: 30038,
                gems: [
                    Gems.GLOWING_NIGHTSEYE,
                    Gems.POTENT_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 24262,
                enchant: Enchants.RUNIC_SPELLTHREAD,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 30067,
                enchant: Enchants.BOARS_SPEED,
            }),
            ItemSpec.create({
                id: 28753,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29302,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29370, // Icon of the Silver Crescent
            }),
            ItemSpec.create({
                id: 27683, // Quagmirran's Eye
            }),
            ItemSpec.create({
                id: 29988,
                enchant: Enchants.SUNFIRE,
            }),
            ItemSpec.create({
                id: 32387, // Idol of the Raven Goddess
            }),
        ],
    }),
};
export const P2_HORDE_PRESET = {
    name: 'P2 Horde Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getFaction() == Faction.Horde,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 30233,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.CHAOTIC_SKYFIRE_DIAMOND,
                ],
            }),
            ItemSpec.create({
                id: 30015, // The Sun King's Talisman
            }),
            ItemSpec.create({
                id: 30235,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.GLOWING_NIGHTSEYE,
                    Gems.VEILED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28797,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 30231,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 29918,
                enchant: Enchants.WRIST_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 30232,
                enchant: Enchants.GLOVES_SPELLPOWER,
                gems: [],
            }),
            ItemSpec.create({
                id: 30038,
                gems: [
                    Gems.GLOWING_NIGHTSEYE,
                    Gems.VEILED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 24262,
                enchant: Enchants.RUNIC_SPELLTHREAD,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 30067,
                enchant: Enchants.BOARS_SPEED,
            }),
            ItemSpec.create({
                id: 28753,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29302,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29370, // Icon of the Silver Crescent
            }),
            ItemSpec.create({
                id: 27683, // Quagmirran's Eye
            }),
            ItemSpec.create({
                id: 29988,
                enchant: Enchants.SUNFIRE,
            }),
            ItemSpec.create({
                id: 32387, // Idol of the Raven Goddess
            }),
        ],
    }),
};
export const P3_PRESET = {
    name: 'P3 Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 31040,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.POTENT_PYRESTONE,
                    Gems.CHAOTIC_SKYFIRE_DIAMOND,
                ],
            }),
            ItemSpec.create({
                id: 30015, // The Sun King's Talisman
            }),
            ItemSpec.create({
                id: 31049,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.GLOWING_SHADOWSONG_AMETHYST,
                    Gems.POTENT_PYRESTONE,
                ],
            }),
            ItemSpec.create({
                id: 32331,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 31043,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.RUNED_CRIMSON_SPINEL,
                    Gems.RUNED_CRIMSON_SPINEL,
                    Gems.RUNED_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 32586,
                enchant: Enchants.WRIST_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 31035,
                enchant: Enchants.GLOVES_SPELLPOWER,
                gems: [
                    Gems.POTENT_PYRESTONE,
                ],
            }),
            ItemSpec.create({
                id: 30914, // Belt of the Crescent Moon
            }),
            ItemSpec.create({
                id: 30916,
                enchant: Enchants.RUNIC_SPELLTHREAD,
                gems: [
                    Gems.RUNED_CRIMSON_SPINEL,
                    Gems.RUNED_CRIMSON_SPINEL,
                    Gems.RUNED_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 32352,
                enchant: Enchants.BOARS_SPEED,
                gems: [
                    Gems.POTENT_PYRESTONE,
                    Gems.GLOWING_SHADOWSONG_AMETHYST,
                ],
            }),
            ItemSpec.create({
                id: 32527,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29305,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 32486, // Ashtongue Talisman of Equilibrium
            }),
            ItemSpec.create({
                id: 32483, // The Skull Of Guldan
            }),
            ItemSpec.create({
                id: 32374,
                enchant: Enchants.SUNFIRE,
            }),
            ItemSpec.create({
                id: 32387, // Idol of the Raven Goddess
            }),
        ],
    }),
};
export const P4_PRESET = {
    name: 'P4 Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 31040,
			"enchant": 29191,
			"gems": [
				32218,
				34220
			]
		},
		{
			"id": 33281
		},
		{
			"id": 31049,
			"enchant": 28886,
			"gems": [
				32215,
				32218
			]
		},
		{
			"id": 32331,
			"enchant": 33150
		},
		{
			"id": 31043,
			"enchant": 24003,
			"gems": [
				32196,
				32196,
				32196
			]
		},
		{
			"id": 32586,
			"enchant": 22534
		},
		{
			"id": 31035,
			"enchant": 28272,
			"gems": [
				32218
			]
		},
		{
			"id": 30914
		},
		{
			"id": 30916,
			"enchant": 24274,
			"gems": [
				32196,
				32196,
				32196
			]
		},
		{
			"id": 32352,
			"enchant": 35297,
			"gems": [
				32218,
				32215
			]
		},
		{
			"id": 32527,
			"enchant": 22536
		},
		{
			"id": 33497,
			"enchant": 22536
		},
		{
			"id": 32483
		},
		{
			"id": 33829
		},
		{
			"id": 32374,
			"enchant": 22560
		},
		{
			"id": 32387
		}
	]}`),
};
export const P5_PRESET = {
    name: 'P5 Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 34403,
			"enchant": 29191,
			"gems": [
				34220,
				32196
			]
		},
		{
			"id": 34204
		},
		{
			"id": 34391,
			"enchant": 28886,
			"gems": [
				32221,
				32196
			]
		},
		{
			"id": 34242,
			"enchant": 33150,
			"gems": [
				32196
			]
		},
		{
			"id": 31043,
			"enchant": 24003,
			"gems": [
				32215,
				32215,
				32221
			]
		},
		{
			"id": 34446,
			"enchant": 22534,
			"gems": [
				35760
			]
		},
		{
			"id": 34407,
			"enchant": 28272,
			"gems": [
				32196,
				35760
			]
		},
		{
			"id": 34555,
			"gems": [
				32196
			]
		},
		{
			"id": 34169,
			"enchant": 24274,
			"gems": [
				32196,
				32196,
				35760
			]
		},
		{
			"id": 34572,
			"enchant": 35297,
			"gems": [
				32196
			]
		},
		{
			"id": 34230,
			"enchant": 22536
		},
		{
			"id": 34362,
			"enchant": 22536
		},
		{
			"id": 32483
		},
		{
			"id": 34429
		},
		{
			"id": 34336,
			"enchant": 22560
		},
		{
			"id": 34179
		},
		{
			"id": 32387
		}
	]}`),
};
