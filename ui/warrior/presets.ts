import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { Player } from '/tbc/core/player.js';

import { Warrior, Warrior_Rotation as WarriorRotation, WarriorTalents as WarriorTalents, Warrior_Options as WarriorOptions } from '/tbc/core/proto/warrior.js';
import { Warrior_Rotation_Type as RotationType, Warrior_Rotation_ArmsSlamRotation as ArmsSlamRotation, Warrior_Rotation_ArmsDWRotation as ArmsDWRotation, Warrior_Rotation_FuryRotation as FuryRotation } from '/tbc/core/proto/warrior.js';
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
		useWwExec: true,
		useHsExec: true,
		hsRageTresh: 60,
		useCleave: true,
		useOverpower: false,
		useHamstring: false,
		sunderGlobal: 0,
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
	armsSiam: ArmsSlamRotation.create({
		useSlamExec: true,
		slamLatency: 150,
		useMsExec: true,
		useWwExec: true,
		useHsExec: true,
		hsRageTresh: 70,
		useCleave: true,
		useOverpower: false,
		useHamstring: false,
		sunderGlobal: 0,
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
		useWwExec: true,
		useHsExec: true,
		hsRageTresh: 60,
		useCleave: true,
		useOverpower: false,
		useHamstring: false,
		sunderGlobal: 0,
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

export const P1_FURY_PRESET = {
	name: 'P1 Fury Preset',
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

export const P1_ARMSSLAM_PRESET = {
	name: 'P1 Arms Slam Preset',
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

export const P1_ARMSDW_PRESET = {
	name: 'P1 Arms DW Preset',
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
