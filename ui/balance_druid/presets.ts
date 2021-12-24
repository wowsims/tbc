import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { Player } from '/tbc/core/player.js';

import { BalanceDruid, BalanceDruid_Rotation as BalanceDruidRotation, DruidTalents as DruidTalents, BalanceDruid_Options as BalanceDruidOptions } from '/tbc/core/proto/druid.js';
import { BalanceDruid_Rotation_PrimarySpell as PrimarySpell } from '/tbc/core/proto/druid.js';

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
	data: '510022312503135231351--520033',
};

export const DefaultRotation = BalanceDruidRotation.create({
	primarySpell: PrimarySpell.Adaptive,
	faerieFire: true,
});

export const DefaultOptions = BalanceDruidOptions.create({
});

export const DefaultConsumes = Consumes.create({
	defaultPotion: Potions.SuperManaPotion,
	flaskOfBlindingLight: true,
	brilliantWizardOil: true,
	blackenedBasilisk: true,
});

export const P1_ALLIANCE_PRESET = {
	name: 'P1 Alliance Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<any>) => player.getFaction() == Faction.Alliance,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 29093, // Antlers of Malorne
				enchant: Enchants.GLYPH_OF_POWER,
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.CHAOTIC_SKYFIRE_DIAMOND,
				],
			}),
			ItemSpec.create({
				id: 28762, // Adornment of Stolen Souls
			}),
			ItemSpec.create({
				id: 29095, // Pauldrons of Malorne
				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
				gems: [
					Gems.GLOWING_NIGHTSEYE,
					Gems.VEILED_NOBLE_TOPAZ,
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
					Gems.GLOWING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 24250, // Bracers of Havok
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
					Gems.VEILED_NOBLE_TOPAZ,
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
					Gems.RUNED_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 28753, // Ring of Recurrence
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 29287, // Voilet Signet of the Archmage
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 29370, // Icon of the Silver Crescent
			}),
			ItemSpec.create({
				id: 27683, // Quagmirran's Eye
			}),
			ItemSpec.create({
				id: 28770, // Nathrezim Mindblade
				enchant: Enchants.SUNFIRE,
			}),
			ItemSpec.create({
				id: 29271, // Talisman of Kalecgos
			}),
			ItemSpec.create({
				id: 27518, // Ivory Idol of the Moongodddess
			}),
		],
	}),
};

export const P1_HORDE_PRESET = {
	name: 'P1 Horde Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<any>) => player.getFaction() == Faction.Horde,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 29093, // Antlers of Malorne
				enchant: Enchants.GLYPH_OF_POWER,
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.CHAOTIC_SKYFIRE_DIAMOND,
				],
			}),
			ItemSpec.create({
				id: 28762, // Adornment of Stolen Souls
			}),
			ItemSpec.create({
				id: 29095, // Pauldrons of Malorne
				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
				gems: [
					Gems.GLOWING_NIGHTSEYE,
					Gems.POTENT_NOBLE_TOPAZ,
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
					Gems.POTENT_NOBLE_TOPAZ,
					Gems.GLOWING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 24250, // Bracers of Havok
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
					Gems.VEILED_NOBLE_TOPAZ,
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
					Gems.POTENT_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 28753, // Ring of Recurrence
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 28793, // Band of Crimson Fury
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 29370, // Icon of the Silver Crescent
			}),
			ItemSpec.create({
				id: 27683, // Quagmirran's Eye
			}),
			ItemSpec.create({
				id: 28770, // Nathrezim Mindblade
				enchant: Enchants.SUNFIRE,
			}),
			ItemSpec.create({
				id: 29271, // Talisman of Kalecgos
			}),
			ItemSpec.create({
				id: 27518, // Ivory Idol of the Moongodddess
			}),
		],
	}),
};

export const P2_ALLIANCE_PRESET = {
	name: 'P2 Alliance Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<any>) => player.getFaction() == Faction.Alliance,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 30233, // Nordrassil Headpiece
				enchant: Enchants.GLYPH_OF_POWER,
				gems: [
					Gems.POTENT_NOBLE_TOPAZ,
					Gems.CHAOTIC_SKYFIRE_DIAMOND,
				],
			}),
			ItemSpec.create({
				id: 30015, // The Sun King's Talisman
			}),
			ItemSpec.create({
				id: 30235, // Nordrassil Wrath-Mantle
				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
				gems: [
					Gems.GLOWING_NIGHTSEYE,
					Gems.POTENT_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 28797, // Brute Cloak of the Ogre-Magi
				enchant: Enchants.SUBTLETY,
			}),
			ItemSpec.create({
				id: 30231, // Nordrassil Chestpiece
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
				id: 30232, // Nordrassil Gauntlets
				enchant: Enchants.GLOVES_SPELLPOWER,
				gems: [
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
				id: 24262, // Spellstrike Pants
				enchant: Enchants.RUNIC_SPELLTHREAD,
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 30067, // Velvet Boots of the Guardian
				enchant: Enchants.BOARS_SPEED,
			}),
			ItemSpec.create({
				id: 28753, // Ring of Recurrence
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 29302, // Band of Eternity
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 29370, // Icon of the Silver Crescent
			}),
			ItemSpec.create({
				id: 27683, // Quagmirran's Eye
			}),
			ItemSpec.create({
				id: 29988, // The Nexus Key
				enchant: Enchants.SUNFIRE,
			}),
			ItemSpec.create({
				id: 32387, // Idol of the Raven Goddess
			}),
		],
	}),
};

export const P2_HORDE_PRESET = {
	name: 'P2 Horde Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<any>) => player.getFaction() == Faction.Horde,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 30233, // Nordrassil Headpiece
				enchant: Enchants.GLYPH_OF_POWER,
				gems: [
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.CHAOTIC_SKYFIRE_DIAMOND,
				],
			}),
			ItemSpec.create({
				id: 30015, // The Sun King's Talisman
			}),
			ItemSpec.create({
				id: 30235, // Nordrassil Wrath-Mantle
				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
				gems: [
					Gems.GLOWING_NIGHTSEYE,
					Gems.VEILED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 28797, // Brute Cloak of the Ogre-Magi
				enchant: Enchants.SUBTLETY,
			}),
			ItemSpec.create({
				id: 30231, // Nordrassil Chestpiece
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
				id: 30232, // Nordrassil Gauntlets
				enchant: Enchants.GLOVES_SPELLPOWER,
				gems: [
				],
			}),
			ItemSpec.create({
				id: 30038, // Belt of Blasting
				gems: [
					Gems.GLOWING_NIGHTSEYE,
					Gems.VEILED_NOBLE_TOPAZ,
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
				id: 30067, // Velvet Boots of the Guardian
				enchant: Enchants.BOARS_SPEED,
			}),
			ItemSpec.create({
				id: 28753, // Ring of Recurrence
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 29302, // Band of Eternity
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 29370, // Icon of the Silver Crescent
			}),
			ItemSpec.create({
				id: 27683, // Quagmirran's Eye
			}),
			ItemSpec.create({
				id: 29988, // The Nexus Key
				enchant: Enchants.SUNFIRE,
			}),
			ItemSpec.create({
				id: 32387, // Idol of the Raven Goddess
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
				id: 31040, // Thunderheart Headguard
				enchant: Enchants.GLYPH_OF_POWER,
				gems: [
					Gems.POTENT_PYRESTONE,
					Gems.CHAOTIC_SKYFIRE_DIAMOND,
				],
			}),
			ItemSpec.create({
				id: 30015, // The Sun King's Talisman
			}),
			ItemSpec.create({
				id: 31049, // Thunderheart Shoulderpads
				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
				gems: [
					Gems.GLOWING_SHADOWSONG_AMETHYST,
					Gems.POTENT_PYRESTONE,
				],
			}),
			ItemSpec.create({
				id: 32331, // Cloak of the Illidari Council
				enchant: Enchants.SUBTLETY,
			}),
			ItemSpec.create({
				id: 31043, // Thunderheart Vest
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.RUNED_CRIMSON_SPINEL,
					Gems.RUNED_CRIMSON_SPINEL,
					Gems.RUNED_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 32586, // Bracers of Nimble Thought
				enchant: Enchants.WRIST_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 31035, // Thunderheart Handguards
				enchant: Enchants.GLOVES_SPELLPOWER,
				gems: [
					Gems.POTENT_PYRESTONE,
				],
			}),
			ItemSpec.create({
				id: 30914, // Belt of the Crescent Moon
			}),
			ItemSpec.create({
				id: 30916, // Leggings of Channeled Elements
				enchant: Enchants.RUNIC_SPELLTHREAD,
				gems: [
					Gems.RUNED_CRIMSON_SPINEL,
					Gems.RUNED_CRIMSON_SPINEL,
					Gems.RUNED_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 32352, // Naturewarden's Treads
				enchant: Enchants.BOARS_SPEED,
				gems: [
					Gems.POTENT_PYRESTONE,
					Gems.GLOWING_SHADOWSONG_AMETHYST,
				],
			}),
			ItemSpec.create({
				id: 32527, // Ring of Ancient Knowledge
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 29305, // Band of the Eternal Sage
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 32486, // Ashtongue Talisman of Equilibrium
			}),
			ItemSpec.create({
				id: 32483, // The Skull Of Guldan
			}),
			ItemSpec.create({
				id: 32374, // Zhar'doom, Greatstaff of the Devourer
				enchant: Enchants.SUNFIRE,
			}),
			ItemSpec.create({
				id: 32387, // Idol of the Raven Goddess
			}),
		],
	}),
};
