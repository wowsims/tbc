import { Consumes } from '/tbc/core/proto/common.js';
import { Drums } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { Player } from '/tbc/core/player.js';

import { EnhancementShaman, EnhancementShaman_Rotation as EnhancementShamanRotation, EnhancementShaman_Options as EnhancementShamanOptions } from '/tbc/core/proto/shaman.js';
import {
	AirTotem,
	EarthTotem,
	FireTotem,
	WaterTotem,
	EnhancementShaman_Rotation_PrimaryShock as PrimaryShock,
	ShamanTotems,
	ShamanWeaponImbue,
} from '/tbc/core/proto/shaman.js';

import * as Enchants from '/tbc/core/constants/enchants.js';
import * as Gems from '/tbc/core/proto_utils/gems.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

// Default talents. Uses the wowhead calculator format, make the talents on
// https://tbc.wowhead.com/talent-calc and copy the numbers in the url.
export const StandardTalents = {
	name: 'Ele Sub',
	data: '250030501-500500212502133531151',
};

export const RestoSubspecTalents = {
	name: 'Resto Sub',
	data: '-502520210502133531151-05005301',
};

export const DefaultRotation = EnhancementShamanRotation.create({
	totems: ShamanTotems.create({
		earth: EarthTotem.StrengthOfEarthTotem,
		air: AirTotem.GraceOfAirTotem,
		fire: FireTotem.SearingTotem,
		water: WaterTotem.ManaSpringTotem,
		twistWindfury: true,
		windfuryRank: 5,
	}),
	primaryShock: PrimaryShock.Frost,
	weaveFlameShock: true,
});

export const DefaultOptions = EnhancementShamanOptions.create({
	waterShield: true,
	bloodlust: true,
	delayOffhandSwings: true,
	mainHandImbue: ShamanWeaponImbue.ImbueWindfury,
	offHandImbue: ShamanWeaponImbue.ImbueWindfury,
});

export const DefaultConsumes = Consumes.create({
	drums: Drums.DrumsOfBattle,
	defaultPotion: Potions.SuperManaPotion,
	flaskOfBlindingLight: true,
	brilliantWizardOil: true,
	blackenedBasilisk: true,
});

export const P1_PRESET = {
	name: 'P1 Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 29040, // Cyclone Helm
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
					Gems.BOLD_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 29381, // Choker of Vile Intent
			}),
			ItemSpec.create({
				id: 29043, // Cyclone Shoulderplates
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					Gems.BOLD_LIVING_RUBY,
					Gems.INSCRIBED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 24259, // Vengeance Wrap
				gems: [
					Gems.BOLD_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 29038, // Cyclone Breastplate
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.BOLD_LIVING_RUBY,
					Gems.BOLD_LIVING_RUBY,
					Gems.INSCRIBED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 25697, // Felstalker Bracers
				enchant: Enchants.WRIST_BRAWN,
				gems: [
					Gems.BOLD_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 29039, // Cyclone Gauntlets
				enchant: Enchants.GLOVES_STRENGTH,
			}),
			ItemSpec.create({
				id: 28656, // Girdle of the Prowlder
			}),
			ItemSpec.create({
				id: 30534, // Wyrmscale Greaves
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
				gems: [
					Gems.SOVEREIGN_NIGHTSEYE,
					Gems.SOVEREIGN_NIGHTSEYE,
					Gems.INSCRIBED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 28746, // Fiend Slayer Boots
				gems: [
					Gems.BOLD_LIVING_RUBY,
					Gems.BOLD_LIVING_RUBY,
				],
				enchant: Enchants.FEET_CATS_SWIFTNESS,
			}),
			ItemSpec.create({
				id: 28757, // Ring of a Thousand Marks
			}),
			ItemSpec.create({
				id: 29283, // Violet Signet of the Master Assassin
			}),
			ItemSpec.create({
				id: 28830, // Dragonspine Trophy
			}),
			ItemSpec.create({
				id: 29383, // Bloodlust Brooch
			}),
			ItemSpec.create({
				id: 28767, // The Decapitator
				enchant: Enchants.MONGOOSE,
			}),
			ItemSpec.create({
				id: 27872, // Harvester of Souls
				enchant: Enchants.MONGOOSE,
			}),
			ItemSpec.create({
				id: 27815, // Totem of the Astral Winds
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
				id: 30190, // Cataclysm Helm
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
					Gems.INSCRIBED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 30017, // Telonicus's Pendant of Mayhem
			}),
			ItemSpec.create({
				id: 30055, // Shoulderpads of the Stranger
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					Gems.BOLD_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 29994, // Thalassian Wildercloak
			}),
			ItemSpec.create({
				id: 30185, // Cataclysm Chestplate
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.BOLD_LIVING_RUBY,
					Gems.SOVEREIGN_NIGHTSEYE,
					Gems.INSCRIBED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 30091, // True-Aim Stalker Bands
				enchant: Enchants.WRIST_BRAWN,
				gems: [
					Gems.BOLD_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 30189, // Cataclysm Gauntlets
				enchant: Enchants.GLOVES_STRENGTH,
			}),
			ItemSpec.create({
				id: 30106, // Belt of One-Hundred Deaths
				gems: [
					Gems.BOLD_LIVING_RUBY,
					Gems.SOVEREIGN_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 30192, // Cataclysm Legplates
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
				gems: [
					Gems.BOLD_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 30039, // Boots of Utter Darkness
				enchant: Enchants.FEET_CATS_SWIFTNESS,
			}),
			ItemSpec.create({
				id: 29997, // Band of the Ranger-General
			}),
			ItemSpec.create({
				id: 30052, // Ring of Lethality
			}),
			ItemSpec.create({
				id: 28830, // Dragonspine Trophy
			}),
			ItemSpec.create({
				id: 29383, // Bloodlust Brooch
			}),
			ItemSpec.create({
				id: 32944, // Talon of the Phoenix
				enchant: Enchants.MONGOOSE,
			}),
			ItemSpec.create({
				id: 29996, // Rod of the Sun King
				enchant: Enchants.MONGOOSE,
			}),
			ItemSpec.create({
				id: 27815, // Totem of the Astral Winds
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
				id: 32235, // Cursed Vision of Sargeras
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
					Gems.BOLD_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 32260, // Choker of Endless Nightmares
			}),
			ItemSpec.create({
				id: 32575, // Shoulders of Lightning Reflexes
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
			}),
			ItemSpec.create({
				id: 32323, // Shadowmoon Destroyer's Drape
			}),
			ItemSpec.create({
				id: 30905, // Midnight Chestguard
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.SOVEREIGN_NIGHTSEYE,
					Gems.BOLD_LIVING_RUBY,
					Gems.INSCRIBED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 32574, // Bindings of Lightning Reflexes
				enchant: Enchants.WRIST_BRAWN,
			}),
			ItemSpec.create({
				id: 32234, // Fists of Mukoa
				enchant: Enchants.GLOVES_STRENGTH,
			}),
			ItemSpec.create({
				id: 30106, // Belt of One-Hundred Deaths
				gems: [
					Gems.BOLD_LIVING_RUBY,
					Gems.BOLD_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 30900, // Bow-Stitched Leggings
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
				gems: [
					Gems.BOLD_LIVING_RUBY,
					Gems.INSCRIBED_NOBLE_TOPAZ,
					Gems.SOVEREIGN_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 32510, // Softstep Boots of Tracking
				enchant: Enchants.FEET_CATS_SWIFTNESS,
			}),
			ItemSpec.create({
				id: 29997, // Band of the Ranger-General
			}),
			ItemSpec.create({
				id: 32497, // Stormrage Signet Ring
			}),
			ItemSpec.create({
				id: 28830, // Dragonspine Trophy
			}),
			ItemSpec.create({
				id: 29383, // Bloodlust Brooch
			}),
			ItemSpec.create({
				id: 32236, // Rising Tide
				enchant: Enchants.MONGOOSE,
			}),
			ItemSpec.create({
				id: 29924, // Netherbane
				enchant: Enchants.MONGOOSE,
			}),
			ItemSpec.create({
				id: 27815, // Totem of the Astral Winds
			}),
		],
	}),
};
