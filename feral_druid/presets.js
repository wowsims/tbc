import { Consumes } from '/tbc/core/proto/common.js';
import { BattleElixir } from '/tbc/core/proto/common.js';
import { Food } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Conjured } from '/tbc/core/proto/common.js';
import { WeaponImbue } from '/tbc/core/proto/common.js';
import { FeralDruid_Rotation as FeralDruidRotation, FeralDruid_Options as FeralDruidOptions } from '/tbc/core/proto/druid.js';
import { FeralDruid_Rotation_FinishingMove as FinishingMove } from '/tbc/core/proto/druid.js';
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
    data: '-503032132322105301251-05503301',
};
export const DefaultRotation = FeralDruidRotation.create({
    finishingMove: FinishingMove.Rip,
    mangleTrick: true,
    biteweave: true,
    mangleBot: false,
    ripCp: 5,
    biteCp: 5,
    rakeTrick: false,
    ripweave: false,
});
export const DefaultOptions = FeralDruidOptions.create({});
export const DefaultConsumes = Consumes.create({
    battleElixir: BattleElixir.ElixirOfMajorAgility,
    food: Food.FoodGrilledMudfish,
    mainHandImbue: WeaponImbue.WeaponImbueAdamantiteWeightstone,
    defaultPotion: Potions.HastePotion,
    defaultConjured: Conjured.ConjuredDarkRune,
    scrollOfAgility: 5,
    scrollOfStrength: 5,
});
export const P4_PRESET = {
    name: 'P4 Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 8345,
                enchant: Enchants.GLYPH_OF_FEROCITY,
            }),
            ItemSpec.create({
                id: 24114, // Braided Eternium Chain
            }),
            ItemSpec.create({
                id: 31048,
                enchant: Enchants.MIGHT_OF_THE_SCOURGE,
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
                id: 31042,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                    Gems.DELICATE_CRIMSON_SPINEL,
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 33881,
                enchant: Enchants.WRIST_BRAWN,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 31034,
                enchant: Enchants.GLOVES_MAJOR_AGILITY,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 30106,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 31044,
                enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 32366,
                enchant: Enchants.FEET_CATS_SWIFTNESS,
                gems: [
                    Gems.DELICATE_CRIMSON_SPINEL,
                    Gems.DELICATE_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 29301,
                enchant: Enchants.RING_STATS,
            }),
            ItemSpec.create({
                id: 33496,
                enchant: Enchants.RING_STATS,
            }),
            ItemSpec.create({
                id: 30627, // Tsunami Talisman
            }),
            ItemSpec.create({
                id: 33831, // Berserker's Call
            }),
            ItemSpec.create({
                id: 33716,
                enchant: Enchants.WEAPON_2H_MAJOR_AGILITY,
            }),
            ItemSpec.create({
                id: 32387, // Idol of the Raven Goddess
            }),
        ],
    }),
};
