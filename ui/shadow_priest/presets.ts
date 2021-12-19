import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { Player } from '/tbc/core/player.js';

import { ShadowPriest, ShadowPriest_Rotation as Rotation, ShadowPriest_Options as Options, ShadowPriest_Rotation_RotationType } from '/tbc/core/proto/priest.js';

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

export const DefaultRotation = Rotation.create({
	rotationType: ShadowPriest_Rotation_RotationType.Ideal,
});

export const DefaultOptions = Options.create({
	useShadowfiend: true,
});

export const DefaultConsumes = Consumes.create({
	defaultPotion: Potions.SuperManaPotion,
	flaskOfPureDeath: true,
	superiorWizardOil: true,
	blackenedBasilisk: true,
});

export const P1_PRESET = {
	name: 'P1 Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 24266, // Spellstrike Hood
				enchant: Enchants.GLYPH_OF_POWER,
				gems: [
					Gems.RUNED_ORNATE_RUBY,
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
				id: 28570, // Shadow-Cloak of Dalaran
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
				id: 24250, // Bracers of Havok
				enchant: Enchants.WRIST_SPELLPOWER,
				gems: [
					Gems.RUNED_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 28507, // Handwraps of Flowing Thought
				enchant: Enchants.GLOVES_SPELLPOWER,
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 28799, // Belt of Divine Inspiration
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
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
				id: 21870, // Frozen Shadoweave Boots
				enchant: Enchants.BOARS_SPEED,
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 29352, // Cobalt Band of Tyrigosa
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 28793, // Band of Crimson Fury
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 28789, // Eye of Magtheridon
			}),
			ItemSpec.create({
				id: 29370, // Icon of the Silver Crescent
			}),
			ItemSpec.create({
				id: 29350, // The Black Stalk
			}),
			ItemSpec.create({
				id: 28770, // Nathrezim Mindblade
				enchant: Enchants.SOULFROST,
			}),
			ItemSpec.create({
				id: 29272, // Orb of the Soul-Eater
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
				id: 32494, // Destruction Holo-gogs
				enchant: Enchants.GLYPH_OF_POWER,
				gems: [
					Gems.MYSTICAL_SKYFIRE_DIAMOND,
					Gems.GLOWING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 30666// ritssyns-lost-pendant
			}),
			ItemSpec.create({
				id: 30163, // wings-of-the-avatar
				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 29992, // Royal Cloak of the Sunstriders
				enchant: Enchants.SUBTLETY,
			}),
			ItemSpec.create({
				id: 30107, // vestments-of-the-sea-witch
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: -19, // Elementalist Bracelets of Shadow Wrath
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
				id: 30038, // Belt of Blasting
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.RUNED_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 29972, // trousers-of-the-astromancer
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
				id: 30109, // ring-of-endless-coils
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 29922, // band-of-alar
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 29370, // Icon of the Silver Crescent
			}),
			ItemSpec.create({
				id: 38290, // dark-iron-smoking-pipe
			}),
			ItemSpec.create({
				id: 29982, // wand of forgotten star
			}),
			ItemSpec.create({
				id: 28770, // Nathrezim Mindblade
				enchant: Enchants.SOULFROST,
			}),
			ItemSpec.create({
				id: 29272, // orb-of-the-soul-eater
			}),
		],
	}),
};
