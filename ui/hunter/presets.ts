import { Consumes } from '/tbc/core/proto/common.js';
import { Drums } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Flask } from '/tbc/core/proto/common.js';
import { Food } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { WeaponImbue } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { Player } from '/tbc/core/player.js';

import {
	Hunter,
	Hunter_Rotation as HunterRotation,
	Hunter_Options as HunterOptions,
	Hunter_Options_Ammo as Ammo,
	Hunter_Options_QuiverBonus as QuiverBonus,
	Hunter_Options_PetType as PetType,
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
	data: '502030015150122431051-0505201205',
};

export const MarksmanTalents = {
	name: 'Marksman',
	data: '51200200502-0551201205013253135',
};

export const SurvivalTalents = {
	name: 'Survival',
	data: '502-0550201205-333200022003223005103',
};

export const DefaultRotation = HunterRotation.create({
	useMultiShot: true,
	useArcaneShot: true,
	viperStartManaPercent: 0.2,
	viperStopManaPercent: 0.4,

	meleeWeave: true,
	useRaptorStrike: true,
	timeToWeaveMs: 500,
	percentWeaved: 0.8,
});

export const DefaultOptions = HunterOptions.create({
	quiverBonus: QuiverBonus.Speed15,
	ammo: Ammo.AdamantiteStinger,
	petType: PetType.Ravager,
});

export const DefaultConsumes = Consumes.create({
	defaultPotion: Potions.HastePotion,
	flask: Flask.FlaskOfRelentlessAssault,
	food: Food.FoodGrilledMudfish,
	mainHandImbue: WeaponImbue.WeaponImbueAdamantiteSharpeningStone,
	offHandImbue: WeaponImbue.WeaponImbueAdamantiteSharpeningStone,
});

export const P1_BM_PRESET = {
	name: 'P1 BM Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<any>) => player.getTalentTree() != 2,
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
				enchant: Enchants.WEAPON_2H_MAJOR_AGILITY,
			}),
			ItemSpec.create({
				id: 28772, // Sunfury Bow of the Pheonix
				enchant: Enchants.STABILIZED_ETERNIUM_SCOPE,
			}),
		],
	}),
};

export const P2_BM_PRESET = {
	name: 'P2 BM Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<any>) => player.getTalentTree() != 2,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 30141, // Rift Stalker Helm
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
				],
			}),
			ItemSpec.create({
				id: 30017, // Telonicus's Pendant of Mayhem
			}),
			ItemSpec.create({
				id: 30143, // Rift Stalker Mantle
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 29994, // Thalassian Wildercloak
				enchant: Enchants.CLOAK_GREATER_AGILITY,
			}),
			ItemSpec.create({
				id: 30139, // Rift Stalker Hauberk
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.SHIFTING_NIGHTSEYE,
					Gems.WICKED_NOBLE_TOPAZ,
					Gems.WICKED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 29966, // Vambraces of Ending
				enchant: Enchants.WRIST_ASSAULT,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 30140, // Rift Stalker Gauntlets
				enchant: Enchants.GLOVES_MAJOR_AGILITY,
			}),
			ItemSpec.create({
				id: 30040, // Belt of Deep Shadow
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 29995, // Leggings of Murderous Intent
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
			}),
			ItemSpec.create({
				id: 30104, // Cobra-Lash Boots
				gems: [
					Gems.SHIFTING_NIGHTSEYE,
					Gems.DELICATE_LIVING_RUBY,
				],
				enchant: Enchants.FEET_CATS_SWIFTNESS,
			}),
			ItemSpec.create({
				id: 29997, // Band of the Ranger-General
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
				id: 29993, // Twinblade of the Phoenix
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
				],
				enchant: Enchants.WEAPON_2H_MAJOR_AGILITY,
			}),
			ItemSpec.create({
				id: 30105, // Serpent Spine Longbow
				enchant: Enchants.STABILIZED_ETERNIUM_SCOPE,
			}),
		],
	}),
};

export const P3_BM_PRESET = {
	name: 'P3 BM Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<any>) => player.getTalentTree() != 2,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 32235, // Cursed Vision of Sargeras
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
					Gems.DELICATE_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 32591, // Choker of Serrated Blades
			}),
			ItemSpec.create({
				id: 31006, // Gronnstalker's Spaulders
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
					Gems.DELICATE_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 32323, // Shadowmoon Destroyer's Drape
				enchant: Enchants.CLOAK_GREATER_AGILITY,
			}),
			ItemSpec.create({
				id: 31004, // Gronnstalker's Chestguard
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
					Gems.JAGGED_SEASPRAY_EMERALD,
					Gems.JAGGED_SEASPRAY_EMERALD,
				],
			}),
			ItemSpec.create({
				id: 32324, // Insidious Bands
				enchant: Enchants.WRIST_ASSAULT,
				gems: [
					Gems.WICKED_PYRESTONE,
				],
			}),
			ItemSpec.create({
				id: 31001, // Gronnstalker's Gloves
				enchant: Enchants.GLOVES_MAJOR_AGILITY,
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 30879, // Don Alejandro's Money Belt
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
					Gems.DELICATE_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 31005, // Gronnstalkers Leggings
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 32366, // Shadowmaster's Boots
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
					Gems.WICKED_PYRESTONE,
				],
				enchant: Enchants.FEET_CATS_SWIFTNESS,
			}),
			ItemSpec.create({
				id: 29997, // Band of the Ranger-General
			}),
			ItemSpec.create({
				id: 29301, // Band of the Eternal Champion
			}),
			ItemSpec.create({
				id: 28830, // Dragonspine Trophy
			}),
			ItemSpec.create({
				id: 32505, // Madness of the Betrayer
			}),
			ItemSpec.create({
				id: 30901, // Boundless Agony
				enchant: Enchants.WEAPON_GREATER_AGILITY,
			}),
			ItemSpec.create({
				id: 30881, // Blade of Infamy
				enchant: Enchants.WEAPON_GREATER_AGILITY,
			}),
			ItemSpec.create({
				id: 30906, // Bristleblitz Striker
				enchant: Enchants.STABILIZED_ETERNIUM_SCOPE,
			}),
		],
	}),
};

export const P1_SV_PRESET = {
	name: 'P1 SV Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<any>) => player.getTalentTree() == 2,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 28275, // Beast Lord Helm
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.GLINTING_NOBLE_TOPAZ,
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
				],
			}),
			ItemSpec.create({
				id: 28343, // Jagged Bark Pendant
			}),
			ItemSpec.create({
				id: 27801, // Beast Lord Mantle
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 29382, // Blood Knight War Cloak
				enchant: Enchants.CLOAK_GREATER_AGILITY,
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
				id: 25697, // Felstalker Bracers
				enchant: Enchants.WRIST_ASSAULT,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
				],
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
				id: 28750, // Girdle of Treachery
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 28741, // Skulker's Greaves
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 28545, // Edgewalker Longboots
				enchant: Enchants.FEET_DEXTERITY,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 31277, // Pathfinder's Band
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
				id: 27846, // Claw of the Watcher
				enchant: Enchants.WEAPON_AGILITY,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 28572, // Blade of the Unrequited
				enchant: Enchants.WEAPON_AGILITY,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.GLINTING_NOBLE_TOPAZ,
					Gems.SHIFTING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 28772, // Sunfury Bow of the Pheonix
				enchant: Enchants.STABILIZED_ETERNIUM_SCOPE,
			}),
		],
	}),
};

export const P2_SV_PRESET = {
	name: 'P2 SV Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<any>) => player.getTalentTree() == 2,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 30141, // Rift Stalker Helm
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
				],
			}),
			ItemSpec.create({
				id: 30017, // Telonicus's Pendant of Mayhem
			}),
			ItemSpec.create({
				id: 30143, // Rift Stalker Mantle
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 29994, // Thalassian Wildercloak
				enchant: Enchants.CLOAK_GREATER_AGILITY,
			}),
			ItemSpec.create({
				id: 30054, // Ranger-General's Chestguard
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 29966, // Vambraces of Ending
				enchant: Enchants.WRIST_ASSAULT,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 28506, // Gloves of Dexterous Manipulation
				enchant: Enchants.GLOVES_MAJOR_AGILITY,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 30040, // Belt of Deep Shadow
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 29985, // Void Reaver Greaves
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
				gems: [
					Gems.DELICATE_LIVING_RUBY,
					Gems.GLINTING_NOBLE_TOPAZ,
					Gems.SHIFTING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 30104, // Cobra-Lash Boots
				enchant: Enchants.FEET_DEXTERITY,
				gems: [
					Gems.JAGGED_TALASITE,
					Gems.DELICATE_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 29298, // Band of Eternity
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
				id: 29924, // Netherbane
				enchant: Enchants.WEAPON_AGILITY,
			}),
			ItemSpec.create({
				id: 29948, // Claw of the Phoenix
				enchant: Enchants.WEAPON_AGILITY,
			}),
			ItemSpec.create({
				id: 30105, // Serpent Spine Longbow
				enchant: Enchants.STABILIZED_ETERNIUM_SCOPE,
			}),
		],
	}),
};

export const P3_SV_PRESET = {
	name: 'P3 SV Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<any>) => player.getTalentTree() == 2,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 31003, // Gonnstalker's Helmt
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
				],
			}),
			ItemSpec.create({
				id: 30017, // Telonicus's Pendant of Mayhem
			}),
			ItemSpec.create({
				id: 31006, // Gronnstalker's Spaulders
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
					Gems.DELICATE_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 29994, // Thalassian Wildercloak
				enchant: Enchants.CLOAK_GREATER_AGILITY,
			}),
			ItemSpec.create({
				id: 31004, // Gronnstalker's Chestguard
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
					Gems.JAGGED_SEASPRAY_EMERALD,
					Gems.JAGGED_SEASPRAY_EMERALD,
				],
			}),
			ItemSpec.create({
				id: 32324, // Insidious Bands
				enchant: Enchants.WRIST_ASSAULT,
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 31001, // Gronnstalker's Gloves
				enchant: Enchants.GLOVES_MAJOR_AGILITY,
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 30879, // Don Alejandro's Money Belt
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
					Gems.DELICATE_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 30900, // Bow-Stitched Leggings
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
					Gems.DELICATE_CRIMSON_SPINEL,
					Gems.DELICATE_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 32366, // Shadowmaster's Boots
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
					Gems.DELICATE_CRIMSON_SPINEL,
				],
				enchant: Enchants.FEET_DEXTERITY,
			}),
			ItemSpec.create({
				id: 28791, // Ring of the Recalcitrant
			}),
			ItemSpec.create({
				id: 29301, // Band of the Eternal Champion
			}),
			ItemSpec.create({
				id: 28830, // Dragonspine Trophy
			}),
			ItemSpec.create({
				id: 32505, // Madness of the Betrayer
			}),
			ItemSpec.create({
				id: 30881, // Blade of Infamy
				enchant: Enchants.WEAPON_GREATER_AGILITY,
			}),
			ItemSpec.create({
				id: 30881, // Blade of Infamy
				enchant: Enchants.WEAPON_GREATER_AGILITY,
			}),
			ItemSpec.create({
				id: 32336, // Black Bow of the Betrayer
				enchant: Enchants.STABILIZED_ETERNIUM_SCOPE,
			}),
		],
	}),
};
