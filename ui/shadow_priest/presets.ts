import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { Player } from '/tbc/core/player.js';

import * as Enchants from '/tbc/core/constants/enchants.js';
import * as Gems from '/tbc/core/constants/gems.js';
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

export const P1_BIS = {
	name: 'P1 BIS',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 24266, // Spellstrike Hood
				enchant: Enchants.GLYPH_OF_POWER,
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 30666, // Ritssyn's Lost Pendant
			}),
			ItemSpec.create({
				id: 21869, // Frozen Shadoweave Shoulders
				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 28766, // Ruby Drape of the Mysticant
				enchant: Enchants.SUBTLETY,
			}),
			ItemSpec.create({
				id: 21871, // Frozen Shadoweave Robe
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: -14, // Ravager's Cuffs of Shadow Wrath
				enchant: Enchants.WRIST_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 28780, // Soul-Eater's Handwraps
				enchant: Enchants.GLOVES_SPELLPOWER,
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: -9, // Lurker's Cord of Shadow Wrath
			}),
			ItemSpec.create({
				id: 24262, // Spellstrike Pants
				enchant: Enchants.RUNIC_SPELLTHREAD,
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 21870, // Frozen Shadoweave Boots
				enchant: Enchants.BOARS_SPEED,
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 23031, // Band of the Inevitable
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 21709, // Ring of the Fallen God
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 27683, // Quagmirran's Eye
			}),
			ItemSpec.create({
				id: 29370, // Icon of the Silver Crescent
			}),
			ItemSpec.create({
				id: -16, // Flawless Wand of Shadow Wrath
			}),
			ItemSpec.create({
				id: 28770, // Nathrezim Mindblade
				enchant: Enchants.SUNFIRE,
			}),
			ItemSpec.create({
				id: 29272, // Orb of the Soul-Eater
			}),
		],
	}),
};

// export const P2_BIS = {
// 	name: 'P2 BIS',
// 	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
// 	gear: EquipmentSpec.create({
// 		items: [
// 			ItemSpec.create({
// 				id: 30233, // Nordrassil Headpiece
// 				enchant: Enchants.GLYPH_OF_POWER,
// 				gems: [
// 					Gems.VEILED_NOBLE_TOPAZ,
// 					Gems.CHAOTIC_SKYFIRE_DIAMOND,
// 				],
// 			}),
// 			ItemSpec.create({
// 				id: 30015, // The Sun King's Talisman
// 			}),
// 			ItemSpec.create({
// 				id: 30235, // Nordrassil Wrath-Mantle
// 				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
// 				gems: [
// 					Gems.GLOWING_NIGHTSEYE,
// 					Gems.VEILED_NOBLE_TOPAZ,
// 				],
// 			}),
// 			ItemSpec.create({
// 				id: 28797, // Brute Cloak of the Ogre-Magi
// 				enchant: Enchants.SUBTLETY,
// 			}),
// 			ItemSpec.create({
// 				id: 30231, // Nordrassil Chestpiece
// 				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
// 				gems: [
// 					Gems.RUNED_LIVING_RUBY,
// 					Gems.RUNED_LIVING_RUBY,
// 					Gems.RUNED_LIVING_RUBY,
// 				],
// 			}),
// 			ItemSpec.create({
// 				id: 29918, // Mindstorm Wristbands
// 				enchant: Enchants.WRIST_SPELLPOWER,
// 			}),
// 			ItemSpec.create({
// 				id: 30232, // Nordrassil Gauntlets
// 				enchant: Enchants.GLOVES_SPELLPOWER,
// 				gems: [
// 				],
// 			}),
// 			ItemSpec.create({
// 				id: 30038, // Belt of Blasting
// 				gems: [
// 					Gems.GLOWING_NIGHTSEYE,
// 					Gems.VEILED_NOBLE_TOPAZ,
// 				],
// 			}),
// 			ItemSpec.create({
// 				id: 24262, // Spellstrike Pants
// 				enchant: Enchants.RUNIC_SPELLTHREAD,
// 				gems: [
// 					Gems.RUNED_LIVING_RUBY,
// 					Gems.RUNED_LIVING_RUBY,
// 					Gems.RUNED_LIVING_RUBY,
// 				],
// 			}),
// 			ItemSpec.create({
// 				id: 30067, // Velvet Boots of the Guardian
// 				enchant: Enchants.BOARS_SPEED,
// 			}),
// 			ItemSpec.create({
// 				id: 28753, // Ring of Recurrence
// 				enchant: Enchants.RING_SPELLPOWER,
// 			}),
// 			ItemSpec.create({
// 				id: 29302, // Band of Eternity
// 				enchant: Enchants.RING_SPELLPOWER,
// 			}),
// 			ItemSpec.create({
// 				id: 29370, // Icon of the Silver Crescent
// 			}),
// 			ItemSpec.create({
// 				id: 27683, // Quagmirran's Eye
// 			}),
// 			ItemSpec.create({
// 				id: 29988, // The Nexus Key
// 				enchant: Enchants.SUNFIRE,
// 			}),
// 			ItemSpec.create({
// 				id: 32387, // Idol of the Raven Goddess
// 			}),
// 		],
// 	}),
// };
