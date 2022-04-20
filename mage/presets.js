import { Conjured } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Flask } from '/tbc/core/proto/common.js';
import { Food } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { WeaponImbue } from '/tbc/core/proto/common.js';
import { Mage_Rotation as MageRotation, Mage_Options as MageOptions } from '/tbc/core/proto/mage.js';
import { Mage_Rotation_Type as RotationType, Mage_Rotation_ArcaneRotation as ArcaneRotation, Mage_Rotation_FireRotation as FireRotation, Mage_Rotation_FrostRotation as FrostRotation } from '/tbc/core/proto/mage.js';
import { Mage_Rotation_FireRotation_PrimarySpell as PrimaryFireSpell } from '/tbc/core/proto/mage.js';
import { Mage_Rotation_ArcaneRotation_Filler as ArcaneFiller } from '/tbc/core/proto/mage.js';
import { Mage_Options_ArmorType as ArmorType } from '/tbc/core/proto/mage.js';
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
export const DeepFrostTalents = {
    name: 'Deep Frost',
    data: '230015031003--0535000310230012241551',
};
export const DefaultFireRotation = MageRotation.create({
    type: RotationType.Fire,
    fire: FireRotation.create({
        primarySpell: PrimaryFireSpell.Fireball,
        maintainImprovedScorch: true,
    }),
});
export const DefaultFireOptions = MageOptions.create({
    armor: ArmorType.MageArmor,
});
export const DefaultFireConsumes = Consumes.create({
    flask: Flask.FlaskOfPureDeath,
    food: Food.FoodBlackenedBasilisk,
    mainHandImbue: WeaponImbue.WeaponImbueBrilliantWizardOil,
    defaultPotion: Potions.SuperManaPotion,
    defaultConjured: Conjured.ConjuredFlameCap,
});
export const DefaultFrostRotation = MageRotation.create({
    type: RotationType.Frost,
    frost: FrostRotation.create({
        waterElementalDisobeyChance: 0.1,
    }),
});
export const DefaultFrostOptions = MageOptions.create({
    armor: ArmorType.MageArmor,
});
export const DefaultFrostConsumes = Consumes.create({
    defaultPotion: Potions.SuperManaPotion,
    defaultConjured: Conjured.ConjuredMageManaEmerald,
    flask: Flask.FlaskOfPureDeath,
    food: Food.FoodBlackenedBasilisk,
    mainHandImbue: WeaponImbue.WeaponImbueBrilliantWizardOil,
});
export const DefaultArcaneRotation = MageRotation.create({
    type: RotationType.Arcane,
    arcane: ArcaneRotation.create({
        filler: ArcaneFiller.Frostbolt,
        arcaneBlastsBetweenFillers: 3,
        startRegenRotationPercent: 0.2,
        stopRegenRotationPercent: 0.5,
    }),
});
export const DefaultArcaneOptions = MageOptions.create({
    armor: ArmorType.MageArmor,
});
export const DefaultArcaneConsumes = Consumes.create({
    defaultPotion: Potions.SuperManaPotion,
    defaultConjured: Conjured.ConjuredMageManaEmerald,
    flask: Flask.FlaskOfBlindingLight,
    food: Food.FoodBlackenedBasilisk,
    mainHandImbue: WeaponImbue.WeaponImbueBrilliantWizardOil,
});
export const P1_ARCANE_PRESET = {
    name: 'P1 Arcane Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getRotation().type == RotationType.Arcane,
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
                id: 28762, // Adornment of Stolen Souls
            }),
            ItemSpec.create({
                id: 29079,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.BRILLIANT_DAWNSTONE,
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
                    Gems.BRILLIANT_DAWNSTONE,
                    Gems.VEILED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 21846,
                gems: [
                    Gems.BRILLIANT_DAWNSTONE,
                    Gems.GLOWING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 29078,
                enchant: Enchants.RUNIC_SPELLTHREAD,
            }),
            ItemSpec.create({
                id: 28517,
                enchant: Enchants.BOARS_SPEED,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.BRILLIANT_DAWNSTONE,
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
                id: 28785, // Lightning Capacitor
            }),
            ItemSpec.create({
                id: 28770,
                enchant: Enchants.SUNFIRE,
            }),
            ItemSpec.create({
                id: 29271, // Talisman of Kalecgos
            }),
            ItemSpec.create({
                id: 28783, // Eredar Wand of Obliteration
            }),
        ],
    }),
};
export const P1_FIRE_PRESET = {
    name: 'P1 Fire Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getRotation().type == RotationType.Fire,
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
export const P1_FROST_PRESET = {
    name: 'P1 Frost Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getRotation().type == RotationType.Frost,
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
                id: 28762, // Adornment of Stolen Souls
            }),
            ItemSpec.create({
                id: 29079,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.VEILED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28766,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 21871,
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
                id: 28780,
                enchant: Enchants.GLOVES_SPELLPOWER,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.GLOWING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 24256,
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
                id: 21870,
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
                enchant: Enchants.SOULFROST,
            }),
            ItemSpec.create({
                id: 29269, // Sapphiron's Wing Bone
            }),
            ItemSpec.create({
                id: 28673, // Tirisfal Wand of Ascendancy
            }),
        ],
    }),
};
export const P2_ARCANE_PRESET = {
    name: 'P2 Arcane Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getRotation().type == RotationType.Arcane,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 30206,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.CHAOTIC_SKYFIRE_DIAMOND,
                    Gems.BRILLIANT_DAWNSTONE,
                ],
            }),
            ItemSpec.create({
                id: 30015, // The Sun King's Talisman
            }),
            ItemSpec.create({
                id: 30210,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.BRILLIANT_DAWNSTONE,
                    Gems.GLOWING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 29992,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 30196,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.BRILLIANT_DAWNSTONE,
                    Gems.BRILLIANT_DAWNSTONE,
                    Gems.GLOWING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 29918,
                enchant: Enchants.WRIST_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29987,
                enchant: Enchants.GLOVES_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 30038,
                gems: [
                    Gems.BRILLIANT_DAWNSTONE,
                    Gems.GLOWING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 30207,
                enchant: Enchants.RUNIC_SPELLTHREAD,
                gems: [
                    Gems.BRILLIANT_DAWNSTONE,
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
                id: 29287,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 29370, // Icon of the Silver Crescent
            }),
            ItemSpec.create({
                id: 30720, // Serpent-Coil Braid
            }),
            ItemSpec.create({
                id: 29988,
                enchant: Enchants.SUNFIRE,
            }),
            ItemSpec.create({
                id: 28783, // Eredar Wand of Obliteration
            }),
        ],
    }),
};
export const P2_FIRE_PRESET = {
    name: 'P2 Fire Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getRotation().type == RotationType.Fire,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 32494,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.CHAOTIC_SKYFIRE_DIAMOND,
                    Gems.GLOWING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 30015, // The Sun King's Talisman
            }),
            ItemSpec.create({
                id: 30024,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
            }),
            ItemSpec.create({
                id: 28766,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 30107,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.GLOWING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 29918,
                enchant: Enchants.WRIST_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 21847,
                enchant: Enchants.GLOVES_SPELLPOWER,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.RUNED_LIVING_RUBY,
                ],
            }),
            ItemSpec.create({
                id: 30038,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.GLOWING_NIGHTSEYE,
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
                id: 30037,
                enchant: Enchants.BOARS_SPEED,
            }),
            ItemSpec.create({
                id: 28753,
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
                id: 27683, // Quagmirran's Eye
            }),
            ItemSpec.create({
                id: 30095,
                enchant: Enchants.SUNFIRE,
            }),
            ItemSpec.create({
                id: 29270, // Flametongue Seal
            }),
            ItemSpec.create({
                id: 29982, // Wand of the Forgotten Star
            }),
        ],
    }),
};
export const P2_FROST_PRESET = {
    name: 'P2 Frost Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getRotation().type == RotationType.Frost,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 30206,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.CHAOTIC_SKYFIRE_DIAMOND,
                    Gems.GLOWING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 30015, // The Sun King's Talisman
            }),
            ItemSpec.create({
                id: 30210,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.GLOWING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 28766,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 30107,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.GLOWING_NIGHTSEYE,
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
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.GLOWING_NIGHTSEYE,
                ],
            }),
            ItemSpec.create({
                id: 30038,
                gems: [
                    Gems.VEILED_NOBLE_TOPAZ,
                    Gems.GLOWING_NIGHTSEYE,
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
                id: 21870,
                enchant: Enchants.BOARS_SPEED,
                gems: [
                    Gems.RUNED_LIVING_RUBY,
                    Gems.VEILED_NOBLE_TOPAZ,
                ],
            }),
            ItemSpec.create({
                id: 28753,
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
                id: 27683, // Quagmirran's Eye
            }),
            ItemSpec.create({
                id: 30095,
                enchant: Enchants.SOULFROST,
            }),
            ItemSpec.create({
                id: 29269, // Sapphiron's Wing Bone
            }),
            ItemSpec.create({
                id: 29982, // Wand of the Forgotten Star
            }),
        ],
    }),
};
export const P3_ARCANE_PRESET = {
    name: 'P3 Arcane Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getRotation().type == RotationType.Arcane,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 30206,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.CHAOTIC_SKYFIRE_DIAMOND,
                    Gems.BRILLIANT_LIONSEYE,
                ],
            }),
            ItemSpec.create({
                id: 30015, // The Sun King's Talisman
            }),
            ItemSpec.create({
                id: 30210,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.BRILLIANT_LIONSEYE,
                    Gems.GLOWING_SHADOWSONG_AMETHYST,
                ],
            }),
            ItemSpec.create({
                id: 32331,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 30196,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.BRILLIANT_LIONSEYE,
                    Gems.BRILLIANT_LIONSEYE,
                    Gems.GLOWING_SHADOWSONG_AMETHYST,
                ],
            }),
            ItemSpec.create({
                id: 30870,
                enchant: Enchants.WRIST_SPELLPOWER,
                gems: [
                    Gems.BRILLIANT_LIONSEYE,
                ],
            }),
            ItemSpec.create({
                id: 30205,
                enchant: Enchants.GLOVES_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 30888,
                gems: [
                    Gems.BRILLIANT_LIONSEYE,
                    Gems.BRILLIANT_LIONSEYE,
                ],
            }),
            ItemSpec.create({
                id: 31058,
                enchant: Enchants.RUNIC_SPELLTHREAD,
                gems: [
                    Gems.BRILLIANT_LIONSEYE,
                ],
            }),
            ItemSpec.create({
                id: 32239,
                enchant: Enchants.BOARS_SPEED,
                gems: [
                    Gems.BRILLIANT_LIONSEYE,
                    Gems.BRILLIANT_LIONSEYE,
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
                id: 32483, // The Skull of Gul'dan
            }),
            ItemSpec.create({
                id: 30720, // Serpent-Coil Braid
            }),
            ItemSpec.create({
                id: 32374,
                enchant: Enchants.SUNFIRE,
            }),
            ItemSpec.create({
                id: 28783, // Eredar Wand of Obliteration
            }),
        ],
    }),
};
export const P3_FIRE_PRESET = {
    name: 'P3 Fire Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getRotation().type == RotationType.Fire,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 31056,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.CHAOTIC_SKYFIRE_DIAMOND,
                    Gems.RUNED_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 32589, // Hellfire-Encased Pendant
            }),
            ItemSpec.create({
                id: 31059,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.VEILED_PYRESTONE,
                    Gems.GLOWING_SHADOWSONG_AMETHYST,
                ],
            }),
            ItemSpec.create({
                id: 32331,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 31057,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.VEILED_PYRESTONE,
                    Gems.VEILED_PYRESTONE,
                    Gems.GLOWING_SHADOWSONG_AMETHYST,
                ],
            }),
            ItemSpec.create({
                id: 32586,
                enchant: Enchants.WRIST_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 31055,
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
                id: 29305,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 32483, // The Skull of Gul'dan
            }),
            ItemSpec.create({
                id: 27683, // Quagmirran's Eye
            }),
            ItemSpec.create({
                id: 30910,
                enchant: Enchants.SUNFIRE,
            }),
            ItemSpec.create({
                id: 30872, // Chronicle of Dark Secrets
            }),
            ItemSpec.create({
                id: 29982, // Wand of the Forgotten Star
            }),
        ],
    }),
};
export const P3_FROST_PRESET = {
    name: 'P3 Frost Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getRotation().type == RotationType.Frost,
    gear: EquipmentSpec.create({
        items: [
            ItemSpec.create({
                id: 31056,
                enchant: Enchants.GLYPH_OF_POWER,
                gems: [
                    Gems.CHAOTIC_SKYFIRE_DIAMOND,
                    Gems.RUNED_CRIMSON_SPINEL,
                ],
            }),
            ItemSpec.create({
                id: 32349, // Translucent Spellthread Necklace
            }),
            ItemSpec.create({
                id: 31059,
                enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
                gems: [
                    Gems.VEILED_PYRESTONE,
                    Gems.GLOWING_SHADOWSONG_AMETHYST,
                ],
            }),
            ItemSpec.create({
                id: 32331,
                enchant: Enchants.SUBTLETY,
            }),
            ItemSpec.create({
                id: 31057,
                enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
                gems: [
                    Gems.POTENT_PYRESTONE,
                    Gems.POTENT_PYRESTONE,
                    Gems.GLOWING_SHADOWSONG_AMETHYST,
                ],
            }),
            ItemSpec.create({
                id: 32586,
                enchant: Enchants.WRIST_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 31055,
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
                id: 29305,
                enchant: Enchants.RING_SPELLPOWER,
            }),
            ItemSpec.create({
                id: 32483, // The Skull of Gul'dan
            }),
            ItemSpec.create({
                id: 27683, // Quagmirran's Eye
            }),
            ItemSpec.create({
                id: 30910,
                enchant: Enchants.SUNFIRE,
            }),
            ItemSpec.create({
                id: 30872, // Chronicle of Dark Secrets
            }),
            ItemSpec.create({
                id: 29982, // Wand of the Forgotten Star
            }),
        ],
    }),
};
export const P4_ARCANE_PRESET = {
    name: 'P4 Arcane Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getRotation().type == RotationType.Arcane,
    gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 30206,
			"enchant": 29191,
			"gems": [
				34220,
				32204
			]
		},
		{
			"id": 33281
		},
		{
			"id": 30210,
			"enchant": 28886,
			"gems": [
				32204,
				32215
			]
		},
		{
			"id": 32331,
			"enchant": 33150
		},
		{
			"id": 30196,
			"enchant": 24003,
			"gems": [
				32204,
				32204,
				32215
			]
		},
		{
			"id": 30870,
			"enchant": 22534,
			"gems": [
				32204
			]
		},
		{
			"id": 30205,
			"enchant": 28272
		},
		{
			"id": 30888,
			"gems": [
				32204,
				32204
			]
		},
		{
			"id": 31058,
			"enchant": 24274,
			"gems": [
				32204
			]
		},
		{
			"id": 32239,
			"enchant": 35297,
			"gems": [
				32204,
				32204
			]
		},
		{
			"id": 32527,
			"enchant": 22536
		},
		{
			"id": 29305,
			"enchant": 22536
		},
		{
			"id": 32483
		},
		{
			"id": 30720
		},
		{
			"id": 32374,
			"enchant": 22560
		},
		{
			"id": 33192,
			"gems": [
				32204
			]
		}
	]}`),
};
export const P4_FIRE_PRESET = {
    name: 'P4 Fire Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getRotation().type == RotationType.Fire,
    gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 31056,
			"enchant": 29191,
			"gems": [
				34220,
				32196
			]
		},
		{
			"id": 32589
		},
		{
			"id": 31059,
			"enchant": 28886,
			"gems": [
				32221,
				32215
			]
		},
		{
			"id": 32331,
			"enchant": 33150
		},
		{
			"id": 31057,
			"enchant": 24003,
			"gems": [
				32221,
				32221,
				32215
			]
		},
		{
			"id": 32586,
			"enchant": 22534
		},
		{
			"id": 31055,
			"enchant": 28272,
			"gems": [
				32196
			]
		},
		{
			"id": 32256
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
			"id": 32239,
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
			"id": 30910,
			"enchant": 22560
		},
		{
			"id": 30872
		},
		{
			"id": 29982
		}
	]}`),
};
export const P4_FROST_PRESET = {
    name: 'P4 Frost Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getRotation().type == RotationType.Frost,
    gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 31056,
			"enchant": 29191,
			"gems": [
				34220,
				32196
			]
		},
		{
			"id": 33281
		},
		{
			"id": 31059,
			"enchant": 28886,
			"gems": [
				32221,
				32215
			]
		},
		{
			"id": 32524,
			"enchant": 33150
		},
		{
			"id": 31057,
			"enchant": 24003,
			"gems": [
				32221,
				32221,
				32215
			]
		},
		{
			"id": 32586,
			"enchant": 22534
		},
		{
			"id": 31055,
			"enchant": 28272,
			"gems": [
				32196
			]
		},
		{
			"id": 32256
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
			"id": 32239,
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
			"id": 30910,
			"enchant": 22560
		},
		{
			"id": 30872
		},
		{
			"id": 29982
		}
	]}`),
};
export const P5_ARCANE_PRESET = {
    name: 'P5 Arcane Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getRotation().type == RotationType.Arcane,
    gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 34405,
			"enchant": 29191,
			"gems": [
				34220,
				32196
			]
		},
		{
			"id": 33281
		},
		{
			"id": 30210,
			"enchant": 28886,
			"gems": [
				32204,
				32215
			]
		},
		{
			"id": 34242,
			"enchant": 33150,
			"gems": [
				32204
			]
		},
		{
			"id": 34399,
			"enchant": 24003,
			"gems": [
				32204,
				32215,
				32204
			]
		},
		{
			"id": 34447,
			"enchant": 22534,
			"gems": [
				32204
			]
		},
		{
			"id": 30205,
			"enchant": 28272
		},
		{
			"id": 34557,
			"gems": [
				32204
			]
		},
		{
			"id": 34386,
			"enchant": 24274,
			"gems": [
				32196,
				32204,
				32204
			]
		},
		{
			"id": 34574,
			"enchant": 35297,
			"gems": [
				32204
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
			"id": 30720
		},
		{
			"id": 34336,
			"enchant": 22560
		},
		{
			"id": 34179
		},
		{
			"id": 34347,
			"gems": [
				32204
			]
		}
	]}`),
};
export const P5_FIRE_PRESET = {
    name: 'P5 Fire Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getRotation().type == RotationType.Fire,
    gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 34340,
			"enchant": 29191,
			"gems": [
				34220,
				32215
			]
		},
		{
			"id": 34204
		},
		{
			"id": 34210,
			"enchant": 28886,
			"gems": [
				32196,
				35760
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
			"id": 34232,
			"enchant": 24003,
			"gems": [
				32196,
				35760,
				35760
			]
		},
		{
			"id": 34447,
			"enchant": 22534,
			"gems": [
				32215
			]
		},
		{
			"id": 34344,
			"enchant": 28272,
			"gems": [
				35760,
				32196
			]
		},
		{
			"id": 34557,
			"gems": [
				32221
			]
		},
		{
			"id": 34181,
			"enchant": 24274,
			"gems": [
				32196,
				32196,
				32221
			]
		},
		{
			"id": 34574,
			"enchant": 35297,
			"gems": [
				32221
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
			"enchant": 22560
		},
		{
			"id": 34179
		},
		{
			"id": 34347,
			"gems": [
				35760
			]
		}
	]}`),
};
export const P5_FROST_PRESET = {
    name: 'P5 Frost Preset',
    tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
    enableWhen: (player) => player.getRotation().type == RotationType.Frost,
    gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 34340,
			"enchant": 29191,
			"gems": [
				34220,
				32215
			]
		},
		{
			"id": 34204
		},
		{
			"id": 34210,
			"enchant": 28886,
			"gems": [
				32196,
				35760
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
			"id": 34232,
			"enchant": 24003,
			"gems": [
				32196,
				35760,
				35760
			]
		},
		{
			"id": 34447,
			"enchant": 22534,
			"gems": [
				32215
			]
		},
		{
			"id": 34344,
			"enchant": 28272,
			"gems": [
				35760,
				32196
			]
		},
		{
			"id": 34557,
			"gems": [
				32221
			]
		},
		{
			"id": 34181,
			"enchant": 24274,
			"gems": [
				32196,
				32196,
				32221
			]
		},
		{
			"id": 34574,
			"enchant": 35297,
			"gems": [
				32221
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
			"enchant": 22560
		},
		{
			"id": 34179
		},
		{
			"id": 34347,
			"gems": [
				35760
			]
		}
	]}`),
};
