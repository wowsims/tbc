import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { RetributionPaladin_Rotation as RetributionPaladinRotation, RetributionPaladin_Options as RetributionPaladinOptions } from '/tbc/core/proto/paladin.js';
import * as Enchants from '/tbc/core/constants/enchants.js';
import * as Gems from '/tbc/core/proto_utils/gems.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';
// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.
// Default talents. Uses the wowhead calculator format, make the talents on
// https://tbc.wowhead.com/talent-calc and copy the numbers in the url.
export const RetributionPaladinTalents = {
    name: 'Retribution Paladin',
    data: '5-503201-0523005130033125231051',
};
export const DefaultRotation = RetributionPaladinRotation.create({});
export const DefaultOptions = RetributionPaladinOptions.create({});
export const DefaultConsumes = Consumes.create({
    defaultPotion: Potions.HastePotion,
    flaskOfRelentlessAssault: true,
    roastedClefthoof: true,
});
export const P2_PRESET = {
    name: 'P2 Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => true,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 32461,
                enchant: Enchants.GLYPH_OF_FEROCITY,
                gems: [
                    Gems.RELENTLESS_EARTHSTORM_DIAMOND,
                    Gems.SOVEREIGN_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 30022, // Pendant of the Perilous
            }),
            ItemSpec.create({
                id: 30055,
                enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
                gems: [
                    Gems.BOLD_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 30098,
                enchant: Enchants.CLOAK_GREATER_AGILITY,
            }),
            ItemSpec.create({
                id: 30129,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.BOLD_LIVING_RUBY,
                    Gems.INSCRIBED_NOBLE_TOPAZ,
                    Gems.INSCRIBED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28795,
                enchant: Enchants.WRIST_BRAWN,
                gems: [
                    Gems.SOVEREIGN_NIGHTSEYE,
                    Gems.BOLD_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 29947,
                enchant: Enchants.GLOVES_STRENGTH,
            }),
            ItemSpec.create({
                id: 30106,
                gems: [
                    Gems.BOLD_LIVING_RUBY,
                    Gems.SOVEREIGN_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 30257,
                enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
            }),
            ItemSpec.create({
                id: 33482, // Cobra-Lash Boots
            }),
            ItemSpec.create({
                id: 30061, // Ring of a Thousand Marks
            }),
            ItemSpec.create({
                id: 30834, // Shapeshifter's Signet
            }),
            ItemSpec.create({
                id: 29383, // Bloodlust Brooch
            }),
            ItemSpec.create({
                id: 28830, // Dragonspine Trophy
            }),
            ItemSpec.create({
                id: 28430,
                enchant: Enchants.MONGOOSE,
            }),
            ItemSpec.create({
                id: 27484, // Libram of Avengement
            }),
        ],
    }),
};
