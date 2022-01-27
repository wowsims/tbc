import { Consumes } from '/tbc/core/proto/common.js';
import { Drums } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { Player } from '/tbc/core/player.js';

import {
	Hunter,
	Hunter_Rotation as HunterRotation,
	Hunter_Options as HunterOptions,
	Hunter_Options_Ammo as Ammo,
	Hunter_Options_QuiverBonus as QuiverBonus,
} from '/tbc/core/proto/hunter.js';

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
	data: '502030015150121531051-0505201205',
};

export const DefaultRotation = HunterRotation.create({
	adaptive: true,
	useMultiShot: true,
	meleeWeave: false,
});

export const DefaultOptions = HunterOptions.create({
	quiverBonus: QuiverBonus.Speed15,
	ammo: Ammo.AdamantineStinger,
	latencyMs: 30,
});

export const DefaultConsumes = Consumes.create({
	defaultPotion: Potions.HastePotion,
	flaskOfRelentlessAssault: true,
});

export const P1_BM_PRESET = {
	name: 'P1 BM Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 28275, // Beast Lord Helm
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
				id: 27801, // Beast Lord Mantle
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					Gems.WICKED_NOBLE_TOPAZ,
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 24259, // Vengeance Wrap
				enchant: Enchants.CLOAK_GREATER_AGILITY,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 28228, // Beast Lord Cuirass
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
					Gems.SHIFTING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 29246, // Nightfall Wristguards
				enchant: Enchants.WRIST_ASSAULT,
			}),
			ItemSpec.create({
				id: 27474, // Beast Lord Handguards
				enchant: Enchants.GLOVES_MAJOR_AGILITY,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 28828, // Gronn-Stitched Girdle
				gems: [
					Gems.SHIFTING_NIGHTSEYE,
					Gems.WICKED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 30739, // Scaled Greaves of the Marksman
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 28545, // Edgewalker Longboots
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
				id: 28435, // Mooncleaver
				enchant: Enchants.WEAPON_MAJOR_AGILITY,
			}),
			ItemSpec.create({
				id: 28772, // Sunfury Bow of the Pheonix
				enchant: Enchants.STABILIZED_ETERNIUM_SCOPE,
			}),
		],
	}),
};
