import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { Player } from '/tbc/core/player.js';

import { Warrior, Warrior_Rotation as WarriorRotation, WarriorTalents as WarriorTalents, Warrior_Options as WarriorOptions } from '/tbc/core/proto/warrior.js';
import { Warrior_Rotation_Type as RotationType, Warrior_Rotation_ArmsSlamRotation as ArmsSlamRotation, Warrior_Rotation_ArmsDWRotation as ArmsDWRotation, Warrior_Rotation_FuryRotation as FuryRotation, Warrior_Rotation_GeneralRotation as GeneralRotation } from '/tbc/core/proto/warrior.js';
import { Warrior_Rotation_FuryRotation_PrimaryInstant as PrimaryInstant } from '/tbc/core/proto/warrior.js';

import * as Enchants from '/tbc/core/constants/enchants.js';
import * as Gems from '/tbc/core/proto_utils/gems.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

// Default talents. Uses the wowhead calculator format, make the talents on
// https://tbc.wowhead.com/talent-calc and copy the numbers in the url.
export const ArmsSlamTalents = {
	name: 'Arms Slam',
	data: '32003301352010500221-0550000500521203',
};
export const ArmsDWTalents = {
	name: 'Arms DW',
	data: '32003301352010500221-0550000520501203',
};
export const FuryTalents = {
	name: 'Fury',
	data: '32003301302-055000055050120531151',
};

export const DefaultFuryRotation = WarriorRotation.create({
	type: RotationType.Fury,
	fury: FuryRotation.create({
		primaryInstant: PrimaryInstant.Whirlwind,
		useBtExec: true,
		rampageCdTresh: 5,
	}),
});

export const DefaultFuryOptions = WarriorOptions.create({
	startingRage: 0,
	precastSapphire: false,
	precastT2: false,
});

export const DefaultFuryConsumes = Consumes.create({
	defaultPotion: Potions.HastePotion,
	flaskOfRelentlessAssault: true,
	roastedClefthoof: true,
});

export const DefaultArmsSlamRotation = WarriorRotation.create({
	type: RotationType.ArmsSlam,
	armsSlam: ArmsSlamRotation.create({
		useSlamExec: true,
		slamLatency: 150,
		useMsExec: true,
	}),
});

export const DefaultArmsSlamOptions = WarriorOptions.create({
	startingRage: 0,
	precastSapphire: false,
	precastT2: false,
});

export const DefaultArmsSlamConsumes = Consumes.create({
	defaultPotion: Potions.HastePotion,
	flaskOfRelentlessAssault: true,
	roastedClefthoof: true,
});

export const DefaultArmsDWRotation = WarriorRotation.create({
	type: RotationType.ArmsSlam,
	armsDw: ArmsDWRotation.create({
		useMsExec: false,
	}),
});

export const DefaultArmsDWOptions = WarriorOptions.create({
	startingRage: 0,
	precastSapphire: false,
	precastT2: false,
});

export const DefaultArmsDWConsumes = Consumes.create({
	defaultPotion: Potions.HastePotion,
	flaskOfRelentlessAssault: true,
	roastedClefthoof: true,
});

export const DefaultGeneralRotation = WarriorRotation.create({
	type: RotationType.General,
	general: GeneralRotation.create({
		useWwExec: true,
		useHsExec: true,
		useOverpower: false,
		overpowerRageThresh: 25,
		useHamstring: false,
		hamstringRageThresh: 70,
	}),
});

export const P1_FURY_PRESET = {
	name: 'P1 Fury Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 29021, // Warbringer Battle Helm
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
					// Gems.SMOOTH_DAWWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 29381, // Choker of Vile Intent
			}),
			ItemSpec.create({
				id: 29023, // Warbringer Shoulderplates
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					// Gems.SMOOTH_DAWNSTONE,
					// Gems.JAGGED_TALASITE,
				],
			}),
			ItemSpec.create({
				id: 24259, // Vengeance Wrap
				enchant: Enchants.CLOAK_GREATER_AGILITY,
				gems: [
					Gems.INSCRIBED_NOBLE_TOPAZ
				],
			}),
			ItemSpec.create({
				id: 29019, // Warbringer Breastplate
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					// Gems.SMOOTH_DAWNSTONE,
					// Gems.SMOOTH_DAWNSTONE,
					// Gems.SMOOTH_DAWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 28795, // Bladespire Warbands
				enchant: Enchants.WRIST_BRAWN,
				gems: [
					// Gems.JAGGED_TALASITE,
					Gems.INSCRIBED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 21847, // Gauntlets of Martial Perfection
				enchant: Enchants.GLOVES_STRENGTH,
				gems: [
					// Gems.JAGGED_TALASITE,
					// Gems.SMOOTH_DAWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 28779, // Girdle of the Endless Pit
				gems: [
					Gems.INSCRIBED_NOBLE_TOPAZ,
					// Gems.JAGGED_TALASITE,
				],
			}),
			ItemSpec.create({
				id: 28741, // Skulker's Greaves
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
				gems: [
					// Gems.SMOOTH_DAWNSTONE,
					// Gems.SMOOTH_DAWNSTONE,
					// Gems.SMOOTH_DAWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 28608, // Ironstriders of Urgency
				enchant: Enchants.FEET_CATS_SWIFTNESS,
				gems: [
					Gems.INSCRIBED_NOBLE_TOPAZ,
					// Gems.SMOOTH_DAWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 28757, // Ring of a Thousand Marks
			}),
			ItemSpec.create({
				id: 30834, // Shapeshifter's Signet
			}),
			ItemSpec.create({
				id: 29383, // Bloodlust Brooch
			}),
			ItemSpec.create({
				id: 28830, // Dragonspine Trophy
			}),
			ItemSpec.create({
				id: 28438, // Dragonmaw
				enchant: Enchants.MONGOOSE,
			}),
			ItemSpec.create({
				id: 28729, // Spiteblade
				enchant: Enchants.MONGOOSE,
			}),
			ItemSpec.create({
				id: 30279, // Mama's Insurance
			}),
		],
	}),
};

export const P1_ARMSSLAM_PRESET = {
	name: 'P1 Arms Slam Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 29021, // Warbringer Battle Helm
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
					// Gems.SMOOTH_DAWWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 29349, // Adamantine Chain of the Unbroken
			}),
			ItemSpec.create({
				id: 29023, // Warbringer Shoulderplates
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					// Gems.SMOOTH_DAWNSTONE,
					// Gems.JAGGED_TALASITE,
				],
			}),
			ItemSpec.create({
				id: 24259, // Vengeance Wrap
				enchant: Enchants.CLOAK_GREATER_AGILITY,
				gems: [
					Gems.INSCRIBED_NOBLE_TOPAZ
				],
			}),
			ItemSpec.create({
				id: 29019, // Warbringer Breastplate
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					// Gems.SMOOTH_DAWNSTONE,
					// Gems.SMOOTH_DAWNSTONE,
					// Gems.SMOOTH_DAWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 28795, // Bladespire Warbands
				enchant: Enchants.WRIST_BRAWN,
				gems: [
					// Gems.JAGGED_TALASITE,
					Gems.INSCRIBED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 21847, // Gauntlets of Martial Perfection
				enchant: Enchants.GLOVES_STRENGTH,
				gems: [
					// Gems.JAGGED_TALASITE,
					// Gems.SMOOTH_DAWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 28779, // Girdle of the Endless Pit
				gems: [
					Gems.INSCRIBED_NOBLE_TOPAZ,
					// Gems.JAGGED_TALASITE,
				],
			}),
			ItemSpec.create({
				id: 28741, // Skulker's Greaves
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
				gems: [
					// Gems.SMOOTH_DAWNSTONE,
					// Gems.SMOOTH_DAWNSTONE,
					// Gems.SMOOTH_DAWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 28608, // Ironstriders of Urgency
				enchant: Enchants.FEET_CATS_SWIFTNESS,
				gems: [
					Gems.INSCRIBED_NOBLE_TOPAZ,
					// Gems.SMOOTH_DAWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 28757, // Ring of a Thousand Marks
			}),
			ItemSpec.create({
				id: 28730, // Mithril Band of the Unscarred
			}),
			ItemSpec.create({
				id: 29383, // Bloodlust Brooch
			}),
			ItemSpec.create({
				id: 28830, // Dragonspine Trophy
			}),
			ItemSpec.create({
				id: 28429, // Lionheart Champion
				enchant: Enchants.MONGOOSE,
			}),
			ItemSpec.create({
				id: 30279, // Mama's Insurance
			}),
		],
	}),
};

export const P1_ARMSDW_PRESET = {
	name: 'P1 Arms DW Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 29021, // Warbringer Battle Helm
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
					// Gems.SMOOTH_DAWWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 29349, // Adamantine Chain of the Unbroken
			}),
			ItemSpec.create({
				id: 29023, // Warbringer Shoulderplates
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					// Gems.SMOOTH_DAWNSTONE,
					// Gems.JAGGED_TALASITE,
				],
			}),
			ItemSpec.create({
				id: 24259, // Vengeance Wrap
				enchant: Enchants.CLOAK_GREATER_AGILITY,
				gems: [
					Gems.INSCRIBED_NOBLE_TOPAZ
				],
			}),
			ItemSpec.create({
				id: 29019, // Warbringer Breastplate
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					// Gems.SMOOTH_DAWNSTONE,
					// Gems.SMOOTH_DAWNSTONE,
					// Gems.SMOOTH_DAWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 28795, // Bladespire Warbands
				enchant: Enchants.WRIST_BRAWN,
				gems: [
					// Gems.JAGGED_TALASITE,
					Gems.INSCRIBED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 21847, // Gauntlets of Martial Perfection
				enchant: Enchants.GLOVES_STRENGTH,
				gems: [
					// Gems.JAGGED_TALASITE,
					// Gems.SMOOTH_DAWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 28779, // Girdle of the Endless Pit
				gems: [
					Gems.INSCRIBED_NOBLE_TOPAZ,
					// Gems.JAGGED_TALASITE,
				],
			}),
			ItemSpec.create({
				id: 28741, // Skulker's Greaves
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
				gems: [
					// Gems.SMOOTH_DAWNSTONE,
					// Gems.SMOOTH_DAWNSTONE,
					// Gems.SMOOTH_DAWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 28608, // Ironstriders of Urgency
				enchant: Enchants.FEET_CATS_SWIFTNESS,
				gems: [
					Gems.INSCRIBED_NOBLE_TOPAZ,
					// Gems.SMOOTH_DAWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 28757, // Ring of a Thousand Marks
			}),
			ItemSpec.create({
				id: 28730, // Mithril Band of the Unscarred
			}),
			ItemSpec.create({
				id: 29383, // Bloodlust Brooch
			}),
			ItemSpec.create({
				id: 28830, // Dragonspine Trophy
			}),
			ItemSpec.create({
				id: 28729, // Spiteblade
				enchant: Enchants.MONGOOSE,
			}),
			ItemSpec.create({
				id: 28295, // Gladiator's Slicer
				enchant: Enchants.MONGOOSE,
			}),
			ItemSpec.create({
				id: 30279, // Mama's Insurance
			}),
		],
	}),
};
