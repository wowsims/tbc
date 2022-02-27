import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Flask } from '/tbc/core/proto/common.js';
import { Food } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { WeaponImbue } from '/tbc/core/proto/common.js';
import { Rogue_Rotation as RogueRotation, Rogue_Options as RogueOptions, } from '/tbc/core/proto/rogue.js';
import * as Enchants from '/tbc/core/constants/enchants.js';
import * as Gems from '/tbc/core/proto_utils/gems.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';
// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.
// Default talents. Uses the wowhead calculator format, make the talents on
// https://tbc.wowhead.com/talent-calc and copy the numbers in the url.
export const CombatTalents = {
    name: 'Combat',
    data: '0053201252-023305102005015002321051',
};
export const CombatMaceTalents = {
    name: 'Combat Maces',
    data: '005320123-023305002005515002321051',
};
export const DefaultRotation = RogueRotation.create({
    maintainExposeArmor: true,
    useRupture: true,
});
export const DefaultOptions = RogueOptions.create({
    useThistleTea: true,
});
export const DefaultConsumes = Consumes.create({
    defaultPotion: Potions.HastePotion,
    flask: Flask.FlaskOfRelentlessAssault,
    food: Food.FoodGrilledMudfish,
    mainHandImbue: WeaponImbue.WeaponImbueAdamantiteSharpeningStone,
    offHandImbue: WeaponImbue.WeaponImbueRogueDeadlyPoison,
});
export const P1_PRESET = {
    name: 'P1 Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 29044,
                enchant: Enchants.GLYPH_OF_FEROCITY,
                gems: [
                    Gems.RELENTLESS_EARTHSTORM_DIAMOND,
                    Gems.GLINTING_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 29381, // Choker of Vile Intent
            }),
            ItemSpec.create({
                id: 27797,
                enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
                gems: [
                    Gems.GLINTING_NOBLE_TOPAZ,
                    Gems.SHIFTING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 28672,
                enchant: Enchants.CLOAK_GREATER_AGILITY,
            }),
            ItemSpec.create({
                id: 29045,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.GLINTING_NOBLE_TOPAZ,
                    Gems.GLINTING_NOBLE_TOPAZ,
                    Gems.SHIFTING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 29246,
                enchant: Enchants.WRIST_ASSAULT,
            }),
            ItemSpec.create({
                id: 27531,
                enchant: Enchants.GLOVES_MAJOR_AGILITY,
                gems: [
                    Gems.GLINTING_NOBLE_TOPAZ,
                    Gems.GLINTING_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 29247, // Girdle of the Deathdealer
            }),
            ItemSpec.create({
                id: 28741,
                enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
                gems: [
                    Gems.DELICATE_LIVING_RUBY,
                    Gems.GLINTING_NOBLE_TOPAZ,
                    Gems.GLINTING_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28545,
                gems: [
                    Gems.GLINTING_NOBLE_TOPAZ,
                    Gems.GLINTING_NOBLE_TOPAZ,
                ],
                enchant: Enchants.FEET_CATS_SWIFTNESS,
            }),
            ItemSpec.create({
                id: 28757, // Ring of a Thousand Marks
            }),
            ItemSpec.create({
                id: 28649, // Garona's Signet Ring
            }),
            ItemSpec.create({
                id: 28830, // Dragonspine Trophy
            }),
            ItemSpec.create({
                id: 29383, // Bloodlust Brooch
            }),
            ItemSpec.create({
                id: 28729,
                enchant: Enchants.MONGOOSE,
            }),
            ItemSpec.create({
                id: 28189,
                enchant: Enchants.MONGOOSE,
            }),
            ItemSpec.create({
                id: 28772,
                enchant: Enchants.STABILIZED_ETERNIUM_SCOPE,
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
                id: 30146,
                enchant: Enchants.GLYPH_OF_FEROCITY,
                gems: [
                    Gems.RELENTLESS_EARTHSTORM_DIAMOND,
                    Gems.GLINTING_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 29381, // Choker of Vile Intent
            }),
            ItemSpec.create({
                id: 30149,
                enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
                gems: [
                    Gems.GLINTING_NOBLE_TOPAZ,
                    Gems.GLINTING_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28672,
                enchant: Enchants.CLOAK_GREATER_AGILITY,
            }),
            ItemSpec.create({
                id: 30101,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.GLINTING_NOBLE_TOPAZ,
                    Gems.GLINTING_NOBLE_TOPAZ,
                    Gems.SHIFTING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 29966,
                enchant: Enchants.WRIST_ASSAULT,
                gems: [
                    Gems.GLINTING_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 30145,
                enchant: Enchants.GLOVES_MAJOR_AGILITY,
            }),
            ItemSpec.create({
                id: 30106,
                gems: [
                    Gems.GLINTING_NOBLE_TOPAZ,
                    Gems.SHIFTING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 30148,
                enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
                gems: [
                    Gems.GLINTING_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28545,
                gems: [
                    Gems.GLINTING_NOBLE_TOPAZ,
                    Gems.RIGID_DAWNSTONE,
                ],
                enchant: Enchants.FEET_CATS_SWIFTNESS,
            }),
            ItemSpec.create({
                id: 29997, // Band of the Ranger-General
            }),
            ItemSpec.create({
                id: 30052, // Ring of Lethality
            }),
            ItemSpec.create({
                id: 28830, // Dragonspine Trophy
            }),
            ItemSpec.create({
                id: 30450, // Warp-Spring Coil
            }),
            ItemSpec.create({
                id: 30082,
                enchant: Enchants.MONGOOSE,
            }),
            ItemSpec.create({
                id: 32027,
                enchant: Enchants.MONGOOSE,
            }),
            ItemSpec.create({
                id: 29949,
                enchant: Enchants.STABILIZED_ETERNIUM_SCOPE,
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
                id: 32235,
                enchant: Enchants.GLYPH_OF_FEROCITY,
                gems: [
                    Gems.RELENTLESS_EARTHSTORM_DIAMOND,
                    Gems.RIGID_LIONSEYE,
                ],
            }),
            ItemSpec.create({
                id: 32260, // Choker of Endless Nightmares
            }),
            ItemSpec.create({
                id: 31030,
                enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
                gems: [
                    Gems.RIGID_LIONSEYE,
                    Gems.RIGID_LIONSEYE,
                ],
            }),
            ItemSpec.create({
                id: 32323,
                enchant: Enchants.CLOAK_GREATER_AGILITY,
            }),
            ItemSpec.create({
                id: 31028,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.RIGID_LIONSEYE,
                    Gems.RIGID_LIONSEYE,
                    Gems.SHIFTING_SHADOWSONG_AMETHYST,
                ],
            }),
            ItemSpec.create({
                id: 32324,
                enchant: Enchants.WRIST_ASSAULT,
                gems: [
                    Gems.RIGID_LIONSEYE,
                ],
            }),
            ItemSpec.create({
                id: 31026,
                enchant: Enchants.GLOVES_MAJOR_AGILITY,
                gems: [
                    Gems.RIGID_LIONSEYE,
                ],
            }),
            ItemSpec.create({
                id: 30106,
                gems: [
                    Gems.GLINTING_PYRESTONE,
                    Gems.SHIFTING_SHADOWSONG_AMETHYST,
                ],
            }),
            ItemSpec.create({
                id: 31029,
                enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
                gems: [
                    Gems.RIGID_LIONSEYE,
                ],
            }),
            ItemSpec.create({
                id: 32366,
                gems: [
                    Gems.GLINTING_PYRESTONE,
                    Gems.RIGID_LIONSEYE,
                ],
                enchant: Enchants.FEET_CATS_SWIFTNESS,
            }),
            ItemSpec.create({
                id: 32497, // Stormrage Signet Ring
            }),
            ItemSpec.create({
                id: 29301, // Band of the Eternal Champion
            }),
            ItemSpec.create({
                id: 28830, // Dragonspine Trophy
            }),
            ItemSpec.create({
                id: 30450, // Warp-Spring Coil
            }),
            ItemSpec.create({
                id: 30881,
                enchant: Enchants.MONGOOSE,
            }),
            ItemSpec.create({
                id: 32369,
                enchant: Enchants.MONGOOSE,
            }),
            ItemSpec.create({
                id: 29949,
                enchant: Enchants.STABILIZED_ETERNIUM_SCOPE,
            }),
        ],
    }),
};
