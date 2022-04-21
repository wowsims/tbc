import { Consumes } from '/tbc/core/proto/common.js';
import { Drums } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Flask } from '/tbc/core/proto/common.js';
import { Food } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { WeaponImbue } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { ElementalShaman_Rotation as ElementalShamanRotation, ElementalShaman_Options as ElementalShamanOptions } from '/tbc/core/proto/shaman.js';
import { ElementalShaman_Rotation_RotationType as RotationType } from '/tbc/core/proto/shaman.js';
import { AirTotem, EarthTotem, FireTotem, WaterTotem, ShamanTotems, } from '/tbc/core/proto/shaman.js';
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
    data: '55003105100213351051--05105301005',
};
export const DefaultRotation = ElementalShamanRotation.create({
    totems: ShamanTotems.create({
        earth: EarthTotem.TremorTotem,
        air: AirTotem.WrathOfAirTotem,
        fire: FireTotem.TotemOfWrath,
        water: WaterTotem.ManaSpringTotem,
    }),
    type: RotationType.Adaptive,
});
export const DefaultOptions = ElementalShamanOptions.create({
    waterShield: true,
    bloodlust: true,
});
export const DefaultConsumes = Consumes.create({
    drums: Drums.DrumsOfBattle,
    defaultPotion: Potions.SuperManaPotion,
    flask: Flask.FlaskOfBlindingLight,
    food: Food.FoodBlackenedBasilisk,
    mainHandImbue: WeaponImbue.WeaponImbueBrilliantWizardOil,
});
export const P1_PRESET = {
    name: 'P1 Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 29035,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.CHAOTIC_SKYFIRE_DIAMOND,
                    Gems.POTENT_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28762, // Adornment of Stolen Souls
            }),
            ItemSpec.create({
                id: 29037,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.POTENT_NOBLE_TOPAZ,
                    Gems.POTENT_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28797,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 29519,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 29521,
                enchant: Enchants.WRIST_SPELLPOWER,
                gems: [
                    Gems.POTENT_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28780,
                enchant: Enchants.GLOVES_SPELLPOWER,
                gems: [
                    Gems.POTENT_NOBLE_TOPAZ,
                    Gems.GLOWING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 29520,
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
                id: 28517,
                enchant: Enchants.BOARS_SPEED,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 30667,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 28753,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29370, // Icon of the Silver Crescent
            }),
            ItemSpec.create({
                id: 28785, // Lightning Capacitor
            }),
            ItemSpec.create({
                id: 28770,
                enchant: Enchants.WEAPON_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29273, // Khadgar's Knapsack
            }),
            ItemSpec.create({
                id: 28248, // Totem of the Void
            }),
        ],
    }),
};
export const P2_PRESET = {
    name: 'P2 Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 29035,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.CHAOTIC_SKYFIRE_DIAMOND,
                    Gems.POTENT_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 30015, // The Sun King's Talisman
            }),
            ItemSpec.create({
                id: 29037,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.POTENT_NOBLE_TOPAZ,
                    Gems.POTENT_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28797,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 30169,
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
                id: 28780,
                enchant: Enchants.GLOVES_SPELLPOWER,
                gems: [
                    Gems.POTENT_NOBLE_TOPAZ,
                    Gems.GLOWING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 30038,
                gems: [
                    Gems.GLOWING_NIGHTSEYE,
                    Gems.POTENT_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 30172,
                enchant: Enchants.RUNIC_SPELLTHREAD,
                gems: [
                    Gems.POTENT_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 30067,
                enchant: Enchants.BOARS_SPEED,
            }),
            ItemSpec.create({
                id: 30667,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 30109,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29370, // Icon of the Silver Crescent
            }),
            ItemSpec.create({
                id: 28785, // Lightning Capacitor
            }),
            ItemSpec.create({
                id: 29988,
                enchant: Enchants.WEAPON_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 28248, // Totem of the Void
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
                id: 31014,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.CHAOTIC_SKYFIRE_DIAMOND,
                    Gems.GLOWING_SHADOWSONG_AMETHYST,
                ],
            }),
            ItemSpec.create({
                id: 30015, // The Sun King's Talisman
            }),
            ItemSpec.create({
                id: 31023,
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
                id: 31017,
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
                id: 31008,
                enchant: Enchants.GLOVES_SPELLPOWER,
                gems: [
                    Gems.POTENT_PYRESTONE,
                ],
            }),
            ItemSpec.create({
                id: 32276, // Flashfire Girdle
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
                    Gems.RUNED_CRIMSON_SPINEL,
                    Gems.RUNED_CRIMSON_SPINEL,
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
                id: 32483, // The Skull Of Guldan
            }),
            ItemSpec.create({
                id: 28785, // Lightning Capacitor
            }),
            ItemSpec.create({
                id: 32374,
                enchant: Enchants.WEAPON_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 32330, // Totem of Ancestral Guidance
            }),
        ],
    }),
};
export const P4_PRESET = {
    name: 'P4 Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 31014,
			"enchant": 29191,
			"gems": [
				34220,
				32215
			]
		},
		{
			"id": 33281
		},
		{
			"id": 31023,
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
			"id": 31017,
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
			"id": 31008,
			"enchant": 28272,
			"gems": [
				32218
			]
		},
		{
			"id": 32276
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
				32196,
				32196
			]
		},
		{
			"id": 32527,
			"enchant": 22536
		},
		{
			"id": 32527,
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
			"enchant": 22555
		},
		{},
		{
			"id": 32330
		}
	]}`),
};
export const P5_ALLIANCE_PRESET = {
    name: 'P5 Alliance Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getFaction() == Faction.Alliance,
    gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 34332,
			"enchant": 29191,
			"gems": [
				35761,
				34220
			]
		},
		{
			"id": 34204
		},
		{
			"id": 31023,
			"enchant": 23545,
			"gems": [
				32215,
				35761
			]
		},
		{
			"id": 34242,
			"enchant": 33150,
			"gems": [
				35760
			]
		},
		{
			"id": 34396,
			"enchant": 24003,
			"gems": [
				35760,
				35761,
				35761
			]
		},
		{
			"id": 34437,
			"enchant": 22534,
			"gems": [
				35761
			]
		},
		{
			"id": 34350,
			"enchant": 28272,
			"gems": [
				35760,
				32215
			]
		},
		{
			"id": 34542,
			"gems": [
				35761
			]
		},
		{
			"id": 34186,
			"enchant": 24274,
			"gems": [
				35761,
				35760,
				35760
			]
		},
		{
			"id": 34566,
			"enchant": 35297,
			"gems": [
				35760
			]
		},
		{
			"id": 34230,
			"enchant": 22536
		},
		{
			"id": 32527,
			"enchant": 22536
		},
		{
			"id": 34429
		},
		{
			"id": 33829
		},
		{
			"id": 34336,
			"enchant": 22555
		},
		{
			"id": 34179
		},
		{
			"id": 32330
		}
	]}`),
};
export const P5_HORDE_PRESET = {
    name: 'P5 Horde Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getFaction() == Faction.Horde,
    gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 34332,
			"enchant": 29191,
			"gems": [
				35761,
				34220
			]
		},
		{
			"id": 34204
		},
		{
			"id": 31023,
			"enchant": 23545,
			"gems": [
				32215,
				35761
			]
		},
		{
			"id": 34242,
			"enchant": 33150,
			"gems": [
				35760
			]
		},
		{
			"id": 34396,
			"enchant": 24003,
			"gems": [
				35760,
				35761,
				35761
			]
		},
		{
			"id": 34437,
			"enchant": 22534,
			"gems": [
				35761
			]
		},
		{
			"id": 34350,
			"enchant": 28272,
			"gems": [
				35760,
				32215
			]
		},
		{
			"id": 34542,
			"gems": [
				35761
			]
		},
		{
			"id": 34186,
			"enchant": 24274,
			"gems": [
				35761,
				35760,
				35760
			]
		},
		{
			"id": 34566,
			"enchant": 35297,
			"gems": [
				35760
			]
		},
		{
			"id": 34230,
			"enchant": 22536
		},
		{
			"id": 32527,
			"enchant": 22536
		},
		{
			"id": 34429
		},
		{
			"id": 32483
		},
		{
			"id": 34336,
			"enchant": 22555
		},
		{
			"id": 34179
		},
		{
			"id": 32330
		}
	]}`),
};
