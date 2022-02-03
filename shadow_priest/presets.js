import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Flask } from '/tbc/core/proto/common.js';
import { Food } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { WeaponImbue } from '/tbc/core/proto/common.js';
import { ShadowPriest_Rotation as Rotation, ShadowPriest_Options as Options, ShadowPriest_Rotation_RotationType } from '/tbc/core/proto/priest.js';
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
    data: '500230013--503250510240103051451',
};
export const DefaultRotation = Rotation.create({
    rotationType: ShadowPriest_Rotation_RotationType.Ideal,
});
export const DefaultOptions = Options.create({
    useShadowfiend: true,
});
export const DefaultConsumes = Consumes.create({
    flask: Flask.FlaskOfPureDeath,
    food: Food.FoodBlackenedBasilisk,
    mainHandImbue: WeaponImbue.WeaponImbueSuperiorWizardOil,
    defaultPotion: Potions.SuperManaPotion,
});
export const P1_PRESET = {
    name: 'P1 Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 24266,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.RUNED_ORNATE_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 30666, // Ritssyn's Lost Pendant
            }),
            ItemSpec.create({
                id: 21869,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 28570,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 21871,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 24250,
                enchant: Enchants.WRIST_SPELLPOWER,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 28507,
                enchant: Enchants.GLOVES_SPELLPOWER,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 28799,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
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
                id: 21870,
                enchant: Enchants.BOARS_SPEED,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 29352,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 28793,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 28789, // Eye of Magtheridon
            }),
            ItemSpec.create({
                id: 29370, // Icon of the Silver Crescent
            }),
            ItemSpec.create({
                id: 29350, // The Black Stalk
            }),
            ItemSpec.create({
                id: 28770,
                enchant: Enchants.SOULFROST,
            }),
            ItemSpec.create({
                id: 29272, // Orb of the Soul-Eater
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
                id: 32494,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.MYSTICAL_SKYFIRE_DIAMOND,
                    Gems.GLOWING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 30666 // ritssyns-lost-pendant
            }),
            ItemSpec.create({
                id: 30163,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 29992,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 30107,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: -19,
                enchant: Enchants.WRIST_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 28780,
                enchant: Enchants.GLOVES_SPELLPOWER,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 30038,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 29972,
                enchant: Enchants.RUNIC_SPELLTHREAD,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 21870,
                enchant: Enchants.BOARS_SPEED,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 30109,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29922,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29370, // Icon of the Silver Crescent
            }),
            ItemSpec.create({
                id: 38290, // dark-iron-smoking-pipe
            }),
            ItemSpec.create({
                id: 29982, // wand of forgotten star
            }),
            ItemSpec.create({
                id: 28770,
                enchant: Enchants.SOULFROST,
            }),
            ItemSpec.create({
                id: 29272, // orb-of-the-soul-eater
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
                id: 31064,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.MYSTICAL_SKYFIRE_DIAMOND,
                    Gems.GLOWING_SHADOWSONG_AMETHYST,
                ],
            }),
            ItemSpec.create({
                id: 30666, // Ritssyn's Lost Pendant
            }),
            ItemSpec.create({
                id: 31070,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.RUNED_CRIMSON_SPINEL,
                    Gems.RUNED_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 32590,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 31065,
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
                id: 31061,
                enchant: Enchants.GLOVES_SPELLPOWER,
                gems: [
                    Gems.RUNED_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 32256, // Waistwrap of Infinity
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
                id: 32239,
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
                id: 32527,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 32483, // The Skull Of Guldan
            }),
            ItemSpec.create({
                id: 29370, // Icon of the Silver Crescent
            }),
            ItemSpec.create({
                id: 29982, // Wand of Forgotten Star
            }),
            ItemSpec.create({
                id: 32374,
                enchant: Enchants.SOULFROST,
            }),
        ],
    }),
};
