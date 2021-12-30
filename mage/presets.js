import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Mage_Rotation as MageRotation, Mage_Options as MageOptions } from '/tbc/core/proto/mage.js';
import { Mage_Rotation_Type as RotationType, Mage_Rotation_FireRotation as FireRotation } from '/tbc/core/proto/mage.js';
import * as Enchants from '/tbc/core/constants/enchants.js';
import * as Gems from '/tbc/core/proto_utils/gems.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';
// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.
// Default talents. Uses the wowhead calculator format, make the talents on
// https://tbc.wowhead.com/talent-calc and copy the numbers in the url.
export const ArcaneTalents = {
    name: 'Arcane',
    data: '2500250300030150330125--053500031003001',
};
export const FireTalents = {
    name: 'Fire',
    data: '2-505202012303331053125-043500001',
};
export const FrostTalents = {
    name: 'Frost',
    data: '2500250300030150330125--053500031003001',
};
export const DefaultFireRotation = MageRotation.create({
    type: RotationType.Fire,
    fire: FireRotation.create({}),
});
export const DefaultFireOptions = MageOptions.create({});
export const DefaultFireConsumes = Consumes.create({
    defaultPotion: Potions.SuperManaPotion,
    flaskOfPureDeath: true,
    brilliantWizardOil: true,
    blackenedBasilisk: true,
});
export const P1_FIRE_PRESET = {
    name: 'P1 Fire Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 29076,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.CHAOTIC_SKYFIRE_DIAMOND,
                    Gems.GLOWING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 28134, // Brooch of Heightened Potential
            }),
            ItemSpec.create({
                id: 29079,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.RUNED_LIVING_RUBY,
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
                    Gems.VEILED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28411,
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
                    Gems.GLOWING_NIGHTSEYE,
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
                    Gems.VEILED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28793,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29172,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29370, // Icon of the Silver Crescent
            }),
            ItemSpec.create({
                id: 27683, // Quagmirran's Eye
            }),
            ItemSpec.create({
                id: 28802,
                enchant: Enchants.SUNFIRE,
            }),
            ItemSpec.create({
                id: 29270, // Flametongue Seal
            }),
            ItemSpec.create({
                id: 28673, // Tirisfal Wand of Ascendancy
            }),
        ],
    }),
};
