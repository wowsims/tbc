import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';

import * as Enchants from '/tbc/core/constants/enchants.js';
import * as Gems from '/tbc/core/constants/gems.js';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

export const StandardTalents = '55030105100213351051--05105301005';

export const PRERAID_GEAR = EquipmentSpec.create({
	items: [
		ItemSpec.create({
			id: 28349, // Tidefury Helm
			enchant: Enchants.GLYPH_OF_POWER,
			gems: [
				Gems.POTENT_NOBLE_TOPAZ,
				Gems.CHAOTIC_SKYFIRE_DIAMOND,
			],
		}),
		ItemSpec.create({
			id: 28134, // Brooch of Heightened Potential
		}),
		ItemSpec.create({
			id: 27802, // Tidefury Shoulderguards
			enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
			gems: [
				Gems.RUNED_LIVING_RUBY,
				Gems.GLOWING_NIGHTSEYE,
			],
		}),
		ItemSpec.create({
			id: 28269, // Baba's Cloak of Arcanistry
		}),
		ItemSpec.create({
			id: 28231, //Tidefury Chestpiece 
			enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
			gems: [
				Gems.GLOWING_NIGHTSEYE,
				Gems.POTENT_NOBLE_TOPAZ,
				Gems.POTENT_NOBLE_TOPAZ,
			],
		}),
		ItemSpec.create({
			id: 28174, // Shattrath Wraps
			enchant: Enchants.WRIST_SPELLPOWER,
			gems: [
				Gems.RUNED_LIVING_RUBY,
			],
		}),
		ItemSpec.create({
			id: 27510, // Tidefury Gauntlets
			enchant: Enchants.GLOVES_SPELLPOWER,
		}),
		ItemSpec.create({
			id: 27783, // Moonrage Girdle
		}),
		ItemSpec.create({
			id: 27909, // Tidefury Kilt
			enchant: Enchants.RUNIC_SPELLTHREAD,
		}),
		ItemSpec.create({
			id: 29313, // Earthbreaker's Greaves
		}),
		ItemSpec.create({
			id: 28555, // Seal of the Exorcist
			enchant: Enchants.RING_SPELLPOWER,
		}),
		ItemSpec.create({
			id: 28510, // Spectral Band of Innervation
			enchant: Enchants.RING_SPELLPOWER,
		}),
		ItemSpec.create({
			id: 29370, // Icon of the Silver Crescent
		}),
		ItemSpec.create({
			id: 27683, // Quagmirran's Eye
		}),
		ItemSpec.create({
			id: 30832, // Gavel of Unearthed Secrets
			enchant: Enchants.WEAPON_SPELLPOWER,
		}),
		ItemSpec.create({
			id: 29268, // Mazthoril Honor Shield
		}),
		ItemSpec.create({
			id: 28248, // Totem of the Void
		}),
	],
});

export const P1_BIS = EquipmentSpec.create({
	items: [
		ItemSpec.create({
			id: 29035, // Cyclone Faceguard
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
			id: 29037, // Cyclone Shoulderguards
			enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
			gems: [
				Gems.POTENT_NOBLE_TOPAZ,
				Gems.POTENT_NOBLE_TOPAZ,
			],
		}),
		ItemSpec.create({
			id: 28797, // Brute Cloak of the Ogre-Magi
		}),
		ItemSpec.create({
			id: 29519, // Netherstrike Breastplate
			enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
			gems: [
				Gems.RUNED_LIVING_RUBY,
				Gems.RUNED_LIVING_RUBY,
				Gems.RUNED_LIVING_RUBY,
			],
		}),
		ItemSpec.create({
			id: 29521, // Netherstrike Bracers
			enchant: Enchants.WRIST_SPELLPOWER,
			gems: [
				Gems.POTENT_NOBLE_TOPAZ,
			],
		}),
		ItemSpec.create({
			id: 28780, // Soul-Eaters's Handwraps
			enchant: Enchants.GLOVES_SPELLPOWER,
			gems: [
				Gems.POTENT_NOBLE_TOPAZ,
				Gems.GLOWING_NIGHTSEYE,
			],
		}),
		ItemSpec.create({
			id: 29520, // Netherstrike Belt
			gems: [
				Gems.GLOWING_NIGHTSEYE,
				Gems.POTENT_NOBLE_TOPAZ,
			],
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
			id: 28517, // Boots of Foretelling
			gems: [
				Gems.RUNED_LIVING_RUBY,
				Gems.RUNED_LIVING_RUBY,
			],
		}),
		ItemSpec.create({
			id: 30667, // Ring of Unrelenting Storms
			enchant: Enchants.RING_SPELLPOWER,
		}),
		ItemSpec.create({
			id: 28753, // Ring of Recurrence
			enchant: Enchants.RING_SPELLPOWER,
		}),
		ItemSpec.create({
			id: 29370, // Icon of the Silver Crescent
		}),
		ItemSpec.create({
			id: 28785, // Lightning Capacitor
		}),
		ItemSpec.create({
			id: 28770, // Nathrezim Mindblade
			enchant: Enchants.WEAPON_SPELLPOWER,
		}),
		ItemSpec.create({
			id: 29273, // Khadgar's Knapsack
		}),
		ItemSpec.create({
			id: 28248, // Totem of the Void
		}),
	],
});

export const P2_BIS = EquipmentSpec.create({
	items: [
		ItemSpec.create({
			id: 29035, // Cyclone Faceguard
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
			id: 29037, // Cyclone Shoulderguards
			enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
			gems: [
				Gems.POTENT_NOBLE_TOPAZ,
				Gems.POTENT_NOBLE_TOPAZ,
			],
		}),
		ItemSpec.create({
			id: 28797, // Brute Cloak of the Ogre-Magi
		}),
		ItemSpec.create({
			id: 30169, // Cataclysm Chestpiece
			enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
			gems: [
				Gems.RUNED_LIVING_RUBY,
				Gems.RUNED_LIVING_RUBY,
				Gems.RUNED_LIVING_RUBY,
			],
		}),
		ItemSpec.create({
			id: 29918, // Mindstorm Wristbands
			enchant: Enchants.WRIST_SPELLPOWER,
		}),
		ItemSpec.create({
			id: 28780, // Soul-Eaters's Handwraps
			enchant: Enchants.GLOVES_SPELLPOWER,
			gems: [
				Gems.POTENT_NOBLE_TOPAZ,
				Gems.GLOWING_NIGHTSEYE,
			],
		}),
		ItemSpec.create({
			id: 30038, // Belt of Blasting
			gems: [
				Gems.GLOWING_NIGHTSEYE,
				Gems.POTENT_NOBLE_TOPAZ,
			],
		}),
		ItemSpec.create({
			id: 30172, // Cataclysm Leggings
			enchant: Enchants.RUNIC_SPELLTHREAD,
			gems: [
				Gems.POTENT_NOBLE_TOPAZ,
			],
		}),
		ItemSpec.create({
			id: 30067, // Velvet Boots of the Guardian
		}),
		ItemSpec.create({
			id: 30667, // Ring of Unrelenting Storms
			enchant: Enchants.RING_SPELLPOWER,
		}),
		ItemSpec.create({
			id: 30109, // Ring of Endless Coils
			enchant: Enchants.RING_SPELLPOWER,
		}),
		ItemSpec.create({
			id: 29370, // Icon of the Silver Crescent
		}),
		ItemSpec.create({
			id: 28785, // Lightning Capacitor
		}),
		ItemSpec.create({
			id: 29988, // The Nexus Key
			enchant: Enchants.WEAPON_SPELLPOWER,
		}),
		ItemSpec.create({
			id: 28248, // Totem of the Void
		}),
	],
});
