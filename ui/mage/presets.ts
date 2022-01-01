import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { Player } from '/tbc/core/player.js';

import { Mage, Mage_Rotation as MageRotation, MageTalents as MageTalents, Mage_Options as MageOptions } from '/tbc/core/proto/mage.js';
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

export const DefaultFireRotation = MageRotation.create({
	type: RotationType.Fire,
	fire: FireRotation.create({
		primarySpell: PrimaryFireSpell.Fireball,
		maintainImprovedScorch: true,
	}),
});

export const DefaultFireOptions = MageOptions.create({
	armor: ArmorType.MageArmor,
	useManaEmeralds: true,
});

export const DefaultFireConsumes = Consumes.create({
	defaultPotion: Potions.SuperManaPotion,
	flaskOfPureDeath: true,
	brilliantWizardOil: true,
	blackenedBasilisk: true,
});

export const DefaultArcaneRotation = MageRotation.create({
	type: RotationType.Arcane,
	arcane: ArcaneRotation.create({
		filler: ArcaneFiller.Frostbolt,
		arcaneBlastsBetweenFillers: 3,
		startRegenRotationPercent: 0.2,
		stopRegenRotationPercent: 0.3,
	}),
});

export const DefaultArcaneOptions = MageOptions.create({
	armor: ArmorType.MageArmor,
	useManaEmeralds: true,
});

export const DefaultArcaneConsumes = Consumes.create({
	defaultPotion: Potions.SuperManaPotion,
	flaskOfBlindingLight: true,
	brilliantWizardOil: true,
	blackenedBasilisk: true,
});

export const P1_FIRE_PRESET = {
	name: 'P1 Fire Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 29076, // Collar of the Aldor
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
				id: 29079, // Pauldrons of the Aldor
				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
				gems: [
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.RUNED_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 28766, // Ruby Drape of the Mysticant
				enchant: Enchants.SUBTLETY,
			}),
			ItemSpec.create({
				id: 21848, // Spellfire Robe
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.VEILED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 28411, // General's Silk Cuffs
				enchant: Enchants.WRIST_SPELLPOWER,
				gems: [
					Gems.VEILED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 21847, // Spellfire Gloves
				enchant: Enchants.GLOVES_SPELLPOWER,
				gems: [
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.GLOWING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 21846, // Spellfire Belt
				gems: [
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.VEILED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 24262, // Spellstrike Pants
				enchant: Enchants.RUNIC_SPELLTHREAD,
				gems: [
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.VEILED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 28517, // Boots of Foretelling
				enchant: Enchants.BOARS_SPEED,
				gems: [
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.VEILED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 28793, // Band of Crimson Fury
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 29172, // Ashyen's Gift
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 29370, // Icon of the Silver Crescent
			}),
			ItemSpec.create({
				id: 27683, // Quagmirran's Eye
			}),
			ItemSpec.create({
				id: 28802, // Bloodmaw Magus Blade
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
