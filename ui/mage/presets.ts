import { Conjured } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Flask } from '/tbc/core/proto/common.js';
import { Food } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { WeaponImbue } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { Player } from '/tbc/core/player.js';

import { Mage, Mage_Rotation as MageRotation, MageTalents as MageTalents, Mage_Options as MageOptions } from '/tbc/core/proto/mage.js';
import { Mage_Rotation_Type as RotationType, Mage_Rotation_ArcaneRotation as ArcaneRotation, Mage_Rotation_FireRotation as FireRotation, Mage_Rotation_FrostRotation as FrostRotation } from '/tbc/core/proto/mage.js';
import { Mage_Rotation_FireRotation_PrimarySpell as PrimaryFireSpell } from '/tbc/core/proto/mage.js';
import { Mage_Rotation_ArcaneRotation_Filler as ArcaneFiller } from '/tbc/core/proto/mage.js';
import { Mage_Options_ArmorType as ArmorType } from '/tbc/core/proto/mage.js';

import * as Enchants from '/tbc/core/constants/enchants.js';
import * as Gems from '/tbc/core/proto_utils/gems.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

// Default talents. Uses the wowhead calculator format, make the talents on
// https://tbc.wowhead.com/talent-calc and copy the numbers in the url.
export const ArcaneTalents = {
	name: 'Arcane',
	data: '2500250300030150330125--053500031003001',
};
export const FireTalents = {
	name: 'Fire',
	data: '2-505202012303331053125-043500001',
};
export const FrostTalents = {
	name: 'Frost',
	data: '2500250300030150330125--053500031003001',
};
export const DeepFrostTalents = {
	name: 'Deep Frost',
	data: '230015031003--0535000310230012241551',
};

export const DefaultFireRotation = MageRotation.create({
	type: RotationType.Fire,
	fire: FireRotation.create({
		primarySpell: PrimaryFireSpell.Fireball,
		maintainImprovedScorch: true,
	}),
});

export const DefaultFireOptions = MageOptions.create({
	armor: ArmorType.MageArmor,
});

export const DefaultFireConsumes = Consumes.create({
	flask: Flask.FlaskOfPureDeath,
	food: Food.FoodBlackenedBasilisk,
	mainHandImbue: WeaponImbue.WeaponImbueBrilliantWizardOil,
	defaultPotion: Potions.SuperManaPotion,
	defaultConjured: Conjured.ConjuredFlameCap,
});

export const DefaultFrostRotation = MageRotation.create({
	type: RotationType.Frost,
	frost: FrostRotation.create({
		waterElementalDisobeyChance: 0.1,
	}),
});

export const DefaultFrostOptions = MageOptions.create({
	armor: ArmorType.MageArmor,
	useManaEmeralds: true,
});

export const DefaultFrostConsumes = Consumes.create({
	defaultPotion: Potions.SuperManaPotion,
	flask: Flask.FlaskOfPureDeath,
	food: Food.FoodBlackenedBasilisk,
	mainHandImbue: WeaponImbue.WeaponImbueBrilliantWizardOil,
});

export const DefaultArcaneRotation = MageRotation.create({
	type: RotationType.Arcane,
	arcane: ArcaneRotation.create({
		filler: ArcaneFiller.Frostbolt,
		arcaneBlastsBetweenFillers: 3,
		startRegenRotationPercent: 0.2,
		stopRegenRotationPercent: 0.5,
	}),
});

export const DefaultArcaneOptions = MageOptions.create({
	armor: ArmorType.MageArmor,
	useManaEmeralds: true,
});

export const DefaultArcaneConsumes = Consumes.create({
	defaultPotion: Potions.SuperManaPotion,
	flask: Flask.FlaskOfBlindingLight,
	food: Food.FoodBlackenedBasilisk,
	mainHandImbue: WeaponImbue.WeaponImbueBrilliantWizardOil,
});

export const P1_ARCANE_PRESET = {
	name: 'P1 Arcane Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Arcane,
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
				id: 28762, // Adornment of Stolen Souls
			}),
			ItemSpec.create({
				id: 29079, // Pauldrons of the Aldor
				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
				gems: [
					Gems.BRILLIANT_DAWNSTONE,
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
					Gems.BRILLIANT_DAWNSTONE,
					Gems.VEILED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 21846, // Spellfire Belt
				gems: [
					Gems.BRILLIANT_DAWNSTONE,
					Gems.GLOWING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 29078, // Legwraps of the Aldor
				enchant: Enchants.RUNIC_SPELLTHREAD,
			}),
			ItemSpec.create({
				id: 28517, // Boots of Foretelling
				enchant: Enchants.BOARS_SPEED,
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.BRILLIANT_DAWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 28753, // Ring of Recurrence
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 29287, // Violet Signet of the Archmage
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
				enchant: Enchants.SUNFIRE,
			}),
			ItemSpec.create({
				id: 29271, // Talisman of Kalecgos
			}),
			ItemSpec.create({
				id: 28783, // Eredar Wand of Obliteration
			}),
		],
	}),
};

export const P1_FIRE_PRESET = {
	name: 'P1 Fire Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Fire,
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

export const P1_FROST_PRESET = {
	name: 'P1 Frost Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Frost,
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
				id: 28762, // Adornment of Stolen Souls
			}),
			ItemSpec.create({
				id: 29079, // Pauldrons of the Aldor
				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
				gems: [
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.VEILED_NOBLE_TOPAZ,
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
				id: 28780, // Soul-Eaters's Handwraps
				enchant: Enchants.GLOVES_SPELLPOWER,
				gems: [
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.GLOWING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 24256, // Girdle of Ruination
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
				id: 21870, // Frozen Shadoweave Boots
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
				enchant: Enchants.SOULFROST,
			}),
			ItemSpec.create({
				id: 29269, // Sapphiron's Wing Bone
			}),
			ItemSpec.create({
				id: 28673, // Tirisfal Wand of Ascendancy
			}),
		],
	}),
};

export const P2_ARCANE_PRESET = {
	name: 'P2 Arcane Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Arcane,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 30206, // Cowl of Tirisfal
				enchant: Enchants.GLYPH_OF_POWER,
				gems: [
					Gems.CHAOTIC_SKYFIRE_DIAMOND,
					Gems.BRILLIANT_DAWNSTONE,
				],
			}),
			ItemSpec.create({
				id: 30015, // The Sun King's Talisman
			}),
			ItemSpec.create({
				id: 30210, // Mantle of Tirisfal
				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
				gems: [
					Gems.BRILLIANT_DAWNSTONE,
					Gems.GLOWING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 29992, // Royal Cloak of the Sunstriders
				enchant: Enchants.SUBTLETY,
			}),
			ItemSpec.create({
				id: 30196, // Robes of Tirisfal
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.BRILLIANT_DAWNSTONE,
					Gems.BRILLIANT_DAWNSTONE,
					Gems.GLOWING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 29918, // Mindstorm Wristbands
				enchant: Enchants.WRIST_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 29987, // Gauntlets of the Sun King
				enchant: Enchants.GLOVES_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 30038, // Belt of Blasting
				gems: [
					Gems.BRILLIANT_DAWNSTONE,
					Gems.GLOWING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 30207, // Leggings of Tirisfal
				enchant: Enchants.RUNIC_SPELLTHREAD,
				gems: [
					Gems.BRILLIANT_DAWNSTONE,
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
				id: 29287, // Violet Signet of the Archmage
				enchant: Enchants.RING_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 29370, // Icon of the Silver Crescent
			}),
			ItemSpec.create({
				id: 30720, // Serpent-Coil Braid
			}),
			ItemSpec.create({
				id: 29988, // The Nexus Key
				enchant: Enchants.SUNFIRE,
			}),
			ItemSpec.create({
				id: 28783, // Eredar Wand of Obliteration
			}),
		],
	}),
};

export const P2_FIRE_PRESET = {
	name: 'P2 Fire Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Fire,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 32494, // Destruction Holo-gogs
				enchant: Enchants.GLYPH_OF_POWER,
				gems: [
					Gems.CHAOTIC_SKYFIRE_DIAMOND,
					Gems.GLOWING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 30015, // The Sun King's Talisman
			}),
			ItemSpec.create({
				id: 30024, // Mantle of the Elven Kings
				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
			}),
			ItemSpec.create({
				id: 28766, // Ruby Drape of the Mysticant
				enchant: Enchants.SUBTLETY,
			}),
			ItemSpec.create({
				id: 30107, // Vestments of the Sea-Witch
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.GLOWING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 29918, // Mindstorm Wristbands
				enchant: Enchants.WRIST_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 21847, // Spellfire Gloves
				enchant: Enchants.GLOVES_SPELLPOWER,
				gems: [
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.RUNED_LIVING_RUBY,
				],
			}),
			ItemSpec.create({
				id: 30038, // Belt of Blasting
				gems: [
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.GLOWING_NIGHTSEYE,
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
				id: 30037, // Boots of Blasting
				enchant: Enchants.BOARS_SPEED,
			}),
			ItemSpec.create({
				id: 28753, // Ring of Recurrence
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
				id: 27683, // Quagmirran's Eye
			}),
			ItemSpec.create({
				id: 30095, // Fang of the Leviathan
				enchant: Enchants.SUNFIRE,
			}),
			ItemSpec.create({
				id: 29270, // Flametongue Seal
			}),
			ItemSpec.create({
				id: 29982, // Wand of the Forgotten Star
			}),
		],
	}),
};

export const P2_FROST_PRESET = {
	name: 'P2 Frost Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Frost,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 30206, // Cowl of Tirisfal
				enchant: Enchants.GLYPH_OF_POWER,
				gems: [
					Gems.CHAOTIC_SKYFIRE_DIAMOND,
					Gems.GLOWING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 30015, // The Sun King's Talisman
			}),
			ItemSpec.create({
				id: 30210, // Mantle of Tirisfal
				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
				gems: [
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.GLOWING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 28766, // Ruby Drape of the Mysticant
				enchant: Enchants.SUBTLETY,
			}),
			ItemSpec.create({
				id: 30107, // Vestments of the Sea-Witch
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.GLOWING_NIGHTSEYE,
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
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.GLOWING_NIGHTSEYE,
				],
			}),
			ItemSpec.create({
				id: 30038, // Belt of Blasting
				gems: [
					Gems.VEILED_NOBLE_TOPAZ,
					Gems.GLOWING_NIGHTSEYE,
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
				id: 21870, // Frozen Shadoweave Boots
				enchant: Enchants.BOARS_SPEED,
				gems: [
					Gems.RUNED_LIVING_RUBY,
					Gems.VEILED_NOBLE_TOPAZ,
				],
			}),
			ItemSpec.create({
				id: 28753, // Ring of Recurrence
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
				id: 27683, // Quagmirran's Eye
			}),
			ItemSpec.create({
				id: 30095, // Fang of the Leviathan
				enchant: Enchants.SOULFROST,
			}),
			ItemSpec.create({
				id: 29269, // Sapphiron's Wing Bone
			}),
			ItemSpec.create({
				id: 29982, // Wand of the Forgotten Star
			}),
		],
	}),
};

export const P3_ARCANE_PRESET = {
	name: 'P3 Arcane Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Arcane,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 30206, // Cowl of Tirisfal
				enchant: Enchants.GLYPH_OF_POWER,
				gems: [
					Gems.CHAOTIC_SKYFIRE_DIAMOND,
					Gems.BRILLIANT_LIONSEYE,
				],
			}),
			ItemSpec.create({
				id: 30015, // The Sun King's Talisman
			}),
			ItemSpec.create({
				id: 30210, // Mantle of Tirisfal
				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
				gems: [
					Gems.BRILLIANT_LIONSEYE,
					Gems.GLOWING_SHADOWSONG_AMETHYST,
				],
			}),
			ItemSpec.create({
				id: 32331, // Cloak of the Illidari Council
				enchant: Enchants.SUBTLETY,
			}),
			ItemSpec.create({
				id: 30196, // Robes of Tirisfal
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.BRILLIANT_LIONSEYE,
					Gems.BRILLIANT_LIONSEYE,
					Gems.GLOWING_SHADOWSONG_AMETHYST,
				],
			}),
			ItemSpec.create({
				id: 30870, // Cuffs of Devastation
				enchant: Enchants.WRIST_SPELLPOWER,
				gems: [
					Gems.BRILLIANT_LIONSEYE,
				],
			}),
			ItemSpec.create({
				id: 30205, // Gloves of Tirisfal
				enchant: Enchants.GLOVES_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 30888, // Anetheron's Noose
				gems: [
					Gems.BRILLIANT_LIONSEYE,
					Gems.BRILLIANT_LIONSEYE,
				],
			}),
			ItemSpec.create({
				id: 31058, // Leggings of the Tempest
				enchant: Enchants.RUNIC_SPELLTHREAD,
				gems: [
					Gems.BRILLIANT_LIONSEYE,
				],
			}),
			ItemSpec.create({
				id: 32239, // Slippers of the Seacaller
				enchant: Enchants.BOARS_SPEED,
				gems: [
					Gems.BRILLIANT_LIONSEYE,
					Gems.BRILLIANT_LIONSEYE,
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
				id: 32483, // The Skull of Gul'dan
			}),
			ItemSpec.create({
				id: 30720, // Serpent-Coil Braid
			}),
			ItemSpec.create({
				id: 32374, // Zhar'doom, Greatstaff of the Devourer
				enchant: Enchants.SUNFIRE,
			}),
			ItemSpec.create({
				id: 28783, // Eredar Wand of Obliteration
			}),
		],
	}),
};

export const P3_FIRE_PRESET = {
	name: 'P3 Fire Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Fire,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 31056, // Cowl of the Tempest
				enchant: Enchants.GLYPH_OF_POWER,
				gems: [
					Gems.CHAOTIC_SKYFIRE_DIAMOND,
					Gems.RUNED_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 32589, // Hellfire-Encased Pendant
			}),
			ItemSpec.create({
				id: 31059, // Mantle of the Tempest
				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
				gems: [
					Gems.VEILED_PYRESTONE,
					Gems.GLOWING_SHADOWSONG_AMETHYST,
				],
			}),
			ItemSpec.create({
				id: 32331, // Cloak of the Illidari Council
				enchant: Enchants.SUBTLETY,
			}),
			ItemSpec.create({
				id: 31057, // Robes of the Tempest
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.VEILED_PYRESTONE,
					Gems.VEILED_PYRESTONE,
					Gems.GLOWING_SHADOWSONG_AMETHYST,
				],
			}),
			ItemSpec.create({
				id: 32586, // Bracers of Nimble Thought
				enchant: Enchants.WRIST_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 31055, // Gloves of the Tempest
				enchant: Enchants.GLOVES_SPELLPOWER,
				gems: [
					Gems.RUNED_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 32256, // Waistwrap of Infinity
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
				id: 32239, // Slippers of the Seacaller
				enchant: Enchants.BOARS_SPEED,
				gems: [
					Gems.RUNED_CRIMSON_SPINEL,
					Gems.RUNED_CRIMSON_SPINEL,
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
				id: 32483, // The Skull of Gul'dan
			}),
			ItemSpec.create({
				id: 27683, // Quagmirran's Eye
			}),
			ItemSpec.create({
				id: 30910, // Tempest of Chaos
				enchant: Enchants.SUNFIRE,
			}),
			ItemSpec.create({
				id: 30872, // Chronicle of Dark Secrets
			}),
			ItemSpec.create({
				id: 29982, // Wand of the Forgotten Star
			}),
		],
	}),
};

export const P3_FROST_PRESET = {
	name: 'P3 Frost Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Frost,
	gear: EquipmentSpec.create({
		items: [
			ItemSpec.create({
				id: 31056, // Cowl of the Tempest
				enchant: Enchants.GLYPH_OF_POWER,
				gems: [
					Gems.CHAOTIC_SKYFIRE_DIAMOND,
					Gems.RUNED_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 32349, // Translucent Spellthread Necklace
			}),
			ItemSpec.create({
				id: 31059, // Mantle of the Tempest
				enchant: Enchants.GREATER_INSCRIPTION_OF_DISCIPLINE,
				gems: [
					Gems.VEILED_PYRESTONE,
					Gems.GLOWING_SHADOWSONG_AMETHYST,
				],
			}),
			ItemSpec.create({
				id: 32331, // Cloak of the Illidari Council
				enchant: Enchants.SUBTLETY,
			}),
			ItemSpec.create({
				id: 31057, // Robes of the Tempest
				enchant: Enchants.CHEST_EXCEPTIONAL_STATS,
				gems: [
					Gems.POTENT_PYRESTONE,
					Gems.POTENT_PYRESTONE,
					Gems.GLOWING_SHADOWSONG_AMETHYST,
				],
			}),
			ItemSpec.create({
				id: 32586, // Bracers of Nimble Thought
				enchant: Enchants.WRIST_SPELLPOWER,
			}),
			ItemSpec.create({
				id: 31055, // Gloves of the Tempest
				enchant: Enchants.GLOVES_SPELLPOWER,
				gems: [
					Gems.RUNED_CRIMSON_SPINEL,
				],
			}),
			ItemSpec.create({
				id: 32256, // Waistwrap of Infinity
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
				id: 32239, // Slippers of the Seacaller
				enchant: Enchants.BOARS_SPEED,
				gems: [
					Gems.RUNED_CRIMSON_SPINEL,
					Gems.RUNED_CRIMSON_SPINEL,
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
				id: 32483, // The Skull of Gul'dan
			}),
			ItemSpec.create({
				id: 27683, // Quagmirran's Eye
			}),
			ItemSpec.create({
				id: 30910, // Tempest of Chaos
				enchant: Enchants.SUNFIRE,
			}),
			ItemSpec.create({
				id: 30872, // Chronicle of Dark Secrets
			}),
			ItemSpec.create({
				id: 29982, // Wand of the Forgotten Star
			}),
		],
	}),
};
