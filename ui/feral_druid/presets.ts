import { Consumes } from '/tbc/core/proto/common.js';
import { BattleElixir } from '/tbc/core/proto/common.js';
import { Food } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { WeaponImbue } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { Player } from '/tbc/core/player.js';

import { FeralDruid, FeralDruid_Rotation as FeralDruidRotation, DruidTalents as DruidTalents, FeralDruid_Options as FeralDruidOptions } from '/tbc/core/proto/druid.js';
import { FeralDruid_Rotation_FinishingMove as FinishingMove } from '/tbc/core/proto/druid.js';

import * as Enchants from '/tbc/core/constants/enchants.js';
import * as Gems from '/tbc/core/proto_utils/gems.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

// Default talents. Uses the wowhead calculator format, make the talents on
// https://tbc.wowhead.com/talent-calc and copy the numbers in the url.
export const StandardTalents = {
	name: 'Standard',
	data: '-503032132322105301251-05503301',
};

export const DefaultRotation = FeralDruidRotation.create({
	finishingMove: FinishingMove.Rip,
	mangleTrick: true,
	biteweave: true,
	mangleBot: false,
	ripCp: 4,
	biteCp: 4,
	rakeTrick: false,
	biteTrick: false,
});

export const DefaultOptions = FeralDruidOptions.create({
});

export const DefaultConsumes = Consumes.create({
	battleElixir: BattleElixir.ElixirOfMajorAgility,
	food: Food.FoodGrilledMudfish,
	mainHandImbue: WeaponImbue.WeaponImbueAdamantiteWeightstone,
	defaultPotion: Potions.SuperManaPotion,
});

export const P3_PRESET = {
	name: 'P3 Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 8345, // Wolfshead Helm
				enchant: Enchants.GLYPH_OF_FEROCITY,
			}),
			ItemSpec.create({
				id: 24114, // Braided Eternium Chain
			}),
			ItemSpec.create({
				id: 31048, // Thunderheart Pauldrons
				enchant: Enchants.MIGHT_OF_THE_SCOURGE,
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
				id: 31042, // Thunderheart Chestguard
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
					Gems.DELICATE_CRIMSON_SPINEL,
					Gems.DELICATE_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 33881, // Vindicator's Dragonhide Bracers
				enchant: Enchants.WRIST_BRAWN,
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 31034, // Thunderheart Gauntlets
				enchant: Enchants.GLOVES_MAJOR_AGILITY,
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 30106, // Belt of One-Hundred Deaths
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
					Gems.DELICATE_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 31044, // Thunderheart Leggings
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 32366, // Shadowmaster's Boots
				enchant: Enchants.FEET_CATS_SWIFTNESS,
				gems: [
					Gems.DELICATE_CRIMSON_SPINEL,
					Gems.DELICATE_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 29301, // Band of the Eternal Champion
				enchant: Enchants.RING_STATS,
			}),
			ItemSpec.create({
				id: 32335, // Unstoppable Aggressor's Ring
				enchant: Enchants.RING_STATS,
			}),
			ItemSpec.create({
				id: 30627, // Tsunami Talisman
			}),
			ItemSpec.create({
				id: 29383, // Bloodlust Brooch
			}),
			ItemSpec.create({
				id: 33716, // Vengeful Gladiator's Staff
				enchant: Enchants.WEAPON_2H_MAJOR_AGILITY,
			}),
			ItemSpec.create({
				id: 32387, // Idol of the Raven Goddess
			}),
		],
	}),
};
