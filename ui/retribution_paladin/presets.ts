import { Conjured, Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';

import { RetributionPaladin_Rotation as RetributionPaladinRotation, RetributionPaladin_Options as RetributionPaladinOptions } from '/tbc/core/proto/paladin.js';
import { RetributionPaladin_Rotation_ConsecrateRank as ConsecrateRank,  RetributionPaladin_Options_Judgement as Judgement } from '/tbc/core/proto/paladin.js';

import * as Enchants from '/tbc/core/constants/enchants.js';
import * as Gems from '/tbc/core/proto_utils/gems.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

// Default talents. Uses the wowhead calculator format, make the talents on
// https://tbc.wowhead.com/talent-calc and copy the numbers in the url.
export const RetKingsPaladinTalents = {
	name: 'Ret w/ Kings',
	data: '5-503201-0523005130033125231051',
};


export const RetNoKingsPaladinTalents = {
	name: 'Ret w/out Kings',
	data: '52-503-0523005130033125331051',
};

export const DefaultRotation = RetributionPaladinRotation.create({
	consecrateRank: ConsecrateRank.None,
	exorcism: false,
});

export const DefaultOptions = RetributionPaladinOptions.create({
	judgement: Judgement.Crusader,
	csDelay: 1700,
	hasteLeeway: 100,
	damageTaken: 0,
});

export const DefaultConsumes = Consumes.create({
	defaultPotion: Potions.HastePotion,
	defaultConjured: Conjured.ConjuredDarkRune,
	flaskOfRelentlessAssault: true,
	roastedClefthoof: true,
});

// Maybe use this later if I can figure out the interactive tooltips from tippy
const RET_BIS_DISCLAIMER = "<p>Please reference <a target=\"_blank\" href=\"https://docs.google.com/spreadsheets/d/1SxO6abYm4k7XRaP1MsxhaqYoukgyZ-cbWDE3ujadjx4/\">Baranor's TBC BiS Lists</a> for more detailed gearing options and information.</p>"

export const PRE_RAID_PRESET = {
	name: 'Pre-Raid Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<Spec.SpecRetributionPaladin>) => true,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 32087, // Mask of the Deceiver
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.INSCRIBED_NOBLE_TOPAZ,
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
				],
			}),
			ItemSpec.create({
				id: 29119, // Haramad's bargain
			}),
			ItemSpec.create({
				id: 33173, // Ragesteel Shoulders
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					Gems.INSCRIBED_NOBLE_TOPAZ,
					Gems.INSCRIBED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 24259, // Vengeance Wrap
				enchant: Enchants.CLOAK_GREATER_AGILITY,
				gems: [
					Gems.BOLD_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 23522, // Ragesteel Breastplate
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
			}),
			ItemSpec.create({
				id: 23537, // Black Felsteel Warbands
				enchant: Enchants.WRIST_BRAWN,

			}),
			ItemSpec.create({
				id: 30341, // Fleshhandler's Gauntlets
				enchant: Enchants.GLOVES_STRENGTH,
			}),
			ItemSpec.create({
				id: 27985, // Deathforge Girdle
				gems: [
					Gems.BOLD_LIVING_RUBY,
					Gems.SOVEREIGN_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 30257, // Shattrath Leggings
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
			}),
			ItemSpec.create({
				id: 28176, // Sha'tari Wrought Greaves
				enchant: Enchants.FEET_DEXTERITY,
				gems: [
					Gems.BOLD_LIVING_RUBY,
					Gems.SOVEREIGN_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 29177, // A'dal's Command
			}),
			ItemSpec.create({
				id: 30834, // Shapeshifter's Signet
			}),
			ItemSpec.create({
				id: 29383, // Bloodlust Brooch
			}),
			ItemSpec.create({
				id: 28288, // Abacus of Violent Odds
			}),
			ItemSpec.create({
				id: 28429, // Lionheart Champion
				enchant: Enchants.MONGOOSE,
			}),
			ItemSpec.create({
				id: 27484, // Libram of Avengement
			}),
		],
	}),
};

export const P1_PRESET = {
	name: 'P1 Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<Spec.SpecRetributionPaladin>) => true,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 29073, // Justicar Crown
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.BOLD_LIVING_RUBY,
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
				],
			}),
			ItemSpec.create({
				id: 28745, // Mithril Chain of Heroism
			}),
			ItemSpec.create({
				id: 29075, // Justicar Shoulders
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					Gems.INSCRIBED_NOBLE_TOPAZ,
					Gems.BOLD_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 24259, // Vengeance Wrap
				enchant: Enchants.CLOAK_GREATER_AGILITY,
				gems: [
					Gems.BOLD_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 29071, // Justicar Breastplate
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.BOLD_LIVING_RUBY,
					Gems.BOLD_LIVING_RUBY,
					Gems.BOLD_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 28795, // Bladespire Warbands
				enchant: Enchants.WRIST_BRAWN,
				gems: [
					Gems.SOVEREIGN_NIGHTSEYE,
					Gems.BOLD_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 30644, // Grips of Deftness
				enchant: Enchants.GLOVES_STRENGTH,
			}),
			ItemSpec.create({
				id: 28779, // Girdle of the Endless Pit
				gems: [
					Gems.BOLD_LIVING_RUBY,
					Gems.SOVEREIGN_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 30257, // Shattrath Leggings
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
			}),
			ItemSpec.create({
				id: 28608, // Ironstriders of Urgency
				enchant: Enchants.FEET_DEXTERITY,
				gems: [
					Gems.BOLD_LIVING_RUBY,
					Gems.INSCRIBED_NOBLE_TOPAZ,
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
				id: 28429, // Lionheart Champion
				enchant: Enchants.MONGOOSE,
			}),
			ItemSpec.create({
				id: 27484, // Libram of Avengement
			}),
		],
	}),
};

export const P2_PRESET = {
	name: 'P2 Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<Spec.SpecRetributionPaladin>) => true,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 32461, // Furios Gizmatic Goggles
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
					Gems.SOVEREIGN_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 30022, // Pendant of the Perilous
			}),
			ItemSpec.create({
				id: 30055, // Shoulderpads of the Stranger
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					Gems.BOLD_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 30098, // Razor-Scale Battlecloak
				enchant: Enchants.CLOAK_GREATER_AGILITY,
			}),
			ItemSpec.create({
				id: 30129, // Crystalforge Breastplate
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.BOLD_LIVING_RUBY,
					Gems.INSCRIBED_NOBLE_TOPAZ,
					Gems.INSCRIBED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 28795, // Bladespire Warbands
				enchant: Enchants.WRIST_BRAWN,
				gems: [
					Gems.SOVEREIGN_NIGHTSEYE,
					Gems.BOLD_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 29947, // Gloves of the Searing Grip
				enchant: Enchants.GLOVES_STRENGTH,
			}),
			ItemSpec.create({
				id: 30106, // Belt of 100 Deaths
				gems: [
					Gems.BOLD_LIVING_RUBY,
					Gems.SOVEREIGN_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 30257, // Shattrath Leggings
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
			}),
			ItemSpec.create({
				id: 30104, // Cobra-Lash Boots
				enchant: Enchants.FEET_DEXTERITY,
				gems: [
					Gems.SOVEREIGN_NIGHTSEYE,
					Gems.BOLD_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 30061, // Ancestral Ring of Conquest
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
				id: 28430, // Lionheart Executioner
				enchant: Enchants.MONGOOSE,
			}),
			ItemSpec.create({
				id: 27484, // Libram of Avengement
			}),
		],
	}),
};

export const P3_PRESET = {
	name: 'P3 Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<Spec.SpecRetributionPaladin>) => true,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 32235, // Cursed vision of Sarg
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
					Gems.BOLD_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 30022, // Pendant of the Perilous
			}),
			ItemSpec.create({
				id: 30055, // Shoulderpads of the Stranger
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					Gems.BOLD_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 33122, // Cloak of Darkness
				enchant: Enchants.CLOAK_GREATER_AGILITY,
				gems: [
					Gems.BOLD_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 30905, // Midnight Chestguard
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.SOVEREIGN_SHADOWSONG_AMETHYST,
					Gems.BOLD_CRIMSON_SPINEL,
					Gems.INSCRIBED_PYRESTONE,
				],
			}),
			ItemSpec.create({
				id: 32574, // Bindings of Lightning Reflexes
				enchant: Enchants.WRIST_BRAWN,
			}),
			ItemSpec.create({
				id: 29947, // Gloves of the Searing Grip
				enchant: Enchants.GLOVES_STRENGTH,
			}),
			ItemSpec.create({
				id: 30106, // Belt of 100 Deaths
				gems: [
					Gems.BOLD_CRIMSON_SPINEL,
					Gems.SOVEREIGN_SHADOWSONG_AMETHYST,
				],
			}),
			ItemSpec.create({
				id: 30900, // Bow-stitched legs
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
				gems: [
					Gems.BOLD_CRIMSON_SPINEL,
					Gems.BOLD_CRIMSON_SPINEL,
					Gems.BOLD_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 32366, // Shadow Master boots
				enchant: Enchants.FEET_DEXTERITY,
				gems: [
					Gems.BOLD_CRIMSON_SPINEL,
					Gems.INSCRIBED_PYRESTONE,
				],
			}),
			ItemSpec.create({
				id: 32526, // Band of Devastation
			}),
			ItemSpec.create({
				id: 30834, // Shapeshifter's Signet
			}),
			ItemSpec.create({
				id: 23206, // Mark of the Champion
			}),
			ItemSpec.create({
				id: 28830, // Dragonspine Trophy
			}),
			ItemSpec.create({
				id: 32332, // Torch of the damned
				enchant: Enchants.MONGOOSE,
			}),
			ItemSpec.create({
				id: 27484, // Libram of Avengement
			}),
		],
	}),
};

export const P4_PRESET = {
	name: 'P4 Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<Spec.SpecRetributionPaladin>) => true,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 32235, // Cursed vision of Sarg
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
					Gems.BOLD_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 30022, // Pendant of the Perilous
			}),
			ItemSpec.create({
				id: 30055, // Shoulderpads of the Stranger
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					Gems.BOLD_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 33590, // Cloak of Fiends
				enchant: Enchants.CLOAK_GREATER_AGILITY,
			}),
			ItemSpec.create({
				id: 30905, // Midnight Chestguard
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.SOVEREIGN_SHADOWSONG_AMETHYST,
					Gems.BOLD_CRIMSON_SPINEL,
					Gems.INSCRIBED_PYRESTONE,
				],
			}),
			ItemSpec.create({
				id: 32574, // Bindings of Lightning Reflexes
				enchant: Enchants.WRIST_BRAWN,
			}),
			ItemSpec.create({
				id: 29947, // Gloves of the Searing Grip
				enchant: Enchants.GLOVES_STRENGTH,
			}),
			ItemSpec.create({
				id: 30106, // Belt of 100 Deaths
				gems: [
					Gems.BOLD_CRIMSON_SPINEL,
					Gems.SOVEREIGN_SHADOWSONG_AMETHYST,
				],
			}),
			ItemSpec.create({
				id: 30900, // Bow-stitched legs
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
				gems: [
					Gems.BOLD_CRIMSON_SPINEL,
					Gems.BOLD_CRIMSON_SPINEL,
					Gems.BOLD_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 32366, // Shadow Master boots
				enchant: Enchants.FEET_DEXTERITY,
				gems: [
					Gems.BOLD_CRIMSON_SPINEL,
					Gems.INSCRIBED_PYRESTONE,
				],
			}),
			ItemSpec.create({
				id: 32526, // Band of Devastation
			}),
			ItemSpec.create({
				id: 30834, // Shapeshifter's Signet
			}),
			ItemSpec.create({
				id: 33831, // Breserker's Call
			}),
			ItemSpec.create({
				id: 28830, // Dragonspine Trophy
			}),
			ItemSpec.create({
				id: 32332, // Torch of the damned
				enchant: Enchants.MONGOOSE,
			}),
			ItemSpec.create({
				id: 27484, // Libram of Avengement
			}),
		],
	}),
};

export const P5_PRESET = {
	name: 'P5 Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<Spec.SpecRetributionPaladin>) => true,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 34244, // Duplicitous Guise
				enchant: Enchants.GLYPH_OF_FEROCITY,
				gems: [
					Gems.RELENTLESS_EARTHSTORM_DIAMOND,
					Gems.BOLD_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 34177, // Clutch of Demise
			}),
			ItemSpec.create({
				id: 34388, // Pauldrons of Breserking
				enchant: Enchants.GREATER_INSCRIPTION_OF_VENGEANCE,
				gems: [
					Gems.BOLD_CRIMSON_SPINEL,
					Gems.INSCRIBED_PYRESTONE,
				],
			}),
			ItemSpec.create({
				id: 34241, // Cloak of Fiends
				enchant: Enchants.CLOAK_GREATER_AGILITY,
				gems: [
					Gems.BOLD_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 34397, // Bladed Chaos Tunic
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [	
					Gems.SOVEREIGN_SHADOWSONG_AMETHYST,
					Gems.INSCRIBED_PYRESTONE,
					Gems.BOLD_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 34431, // Lightbringer Bands
				enchant: Enchants.WRIST_BRAWN,
				gems: [	
					Gems.BOLD_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 34343, // Thal Ranger Gauntlets
				enchant: Enchants.GLOVES_STRENGTH,
				gems: [	
					Gems.BOLD_CRIMSON_SPINEL,
					Gems.INSCRIBED_PYRESTONE,
				],
			}),
			ItemSpec.create({
				id: 34485, // Lightbringer Girdle
				gems: [
					Gems.BOLD_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 34180, // Fel Fury Legplates
				enchant: Enchants.NETHERCOBRA_LEG_ARMOR,
				gems: [	
					Gems.SOVEREIGN_SHADOWSONG_AMETHYST,
					Gems.BOLD_CRIMSON_SPINEL,
					Gems.INSCRIBED_PYRESTONE,
				],
			}),
			ItemSpec.create({
				id: 34561, // Lightbringer Boots
				enchant: Enchants.FEET_DEXTERITY,
				gems: [
					Gems.BOLD_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 34361, // Hard Khorium Band	
			}),
			ItemSpec.create({
				id: 34189, // Band of Ruinous Delight
			}),
			ItemSpec.create({
				id: 34427, // Blackened Naaru Silver
			}),
			ItemSpec.create({
				id: 34472, // Shard of Contempt
			}),
			ItemSpec.create({
				id: 34247, // Apolyon, Soul-Render
				enchant: Enchants.MONGOOSE,
				gems: [
					Gems.BOLD_CRIMSON_SPINEL,
					Gems.BOLD_CRIMSON_SPINEL,
					Gems.BOLD_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 27484, // Libram of Avengement
			}),
		],
	}),
};