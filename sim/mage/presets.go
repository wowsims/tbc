package mage

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var FireTalents = &proto.MageTalents{
	ArcaneSubtlety: 2,

	ImprovedFireball:  5,
	Ignite:            5,
	Incineration:      2,
	Pyroblast:         true,
	ImprovedScorch:    3,
	MasterOfElements:  3,
	PlayingWithFire:   3,
	CriticalMass:      3,
	BlastWave:         true,
	FirePower:         5,
	Pyromaniac:        3,
	Combustion:        true,
	MoltenFury:        2,
	EmpoweredFireball: 5,

	ImprovedFrostbolt:  4,
	ElementalPrecision: 3,
	IceShards:          5,
	IcyVeins:           true,
}

var FrostTalents = &proto.MageTalents{
	ArcaneFocus:         5,
	ArcaneConcentration: 5,
	ArcaneImpact:        3,
	ArcaneMeditation:    3,

	ImprovedFrostbolt:    5,
	ElementalPrecision:   3,
	IceShards:            5,
	IcyVeins:             true,
	PiercingIce:          5,
	FrostChanneling:      5,
	ColdSnap:             true,
	ImprovedConeOfCold:   2,
	IceFloes:             2,
	WintersChill:         4,
	ArcticWinds:          5,
	EmpoweredFrostbolt:   5,
	SummonWaterElemental: true,
}

var fireMageOptions = &proto.Mage_Options{
	Armor: proto.Mage_Options_MageArmor,
}
var PlayerOptionsFire = &proto.Player_Mage{
	Mage: &proto.Mage{
		Talents: FireTalents,
		Options: fireMageOptions,
		Rotation: &proto.Mage_Rotation{
			Type: proto.Mage_Rotation_Fire,
			Fire: &proto.Mage_Rotation_FireRotation{
				PrimarySpell:           proto.Mage_Rotation_FireRotation_Fireball,
				MaintainImprovedScorch: true,
				WeaveFireBlast:         true,
			},
		},
	},
}

var frostMageOptions = &proto.Mage_Options{
	Armor: proto.Mage_Options_MageArmor,
}
var PlayerOptionsFrost = &proto.Player_Mage{
	Mage: &proto.Mage{
		Talents: FrostTalents,
		Options: frostMageOptions,
		Rotation: &proto.Mage_Rotation{
			Type:  proto.Mage_Rotation_Frost,
			Frost: &proto.Mage_Rotation_FrostRotation{},
		},
	},
}

var FullRaidBuffs = &proto.RaidBuffs{
	GiftOfTheWild: proto.TristateEffect_TristateEffectImproved,
}
var FullFirePartyBuffs = &proto.PartyBuffs{
	Drums:           proto.Drums_DrumsOfBattle,
	Bloodlust:       1,
	MoonkinAura:     proto.TristateEffect_TristateEffectRegular,
	ManaSpringTotem: proto.TristateEffect_TristateEffectRegular,
	TotemOfWrath:    1,
	WrathOfAirTotem: proto.TristateEffect_TristateEffectRegular,
}
var FullFrostPartyBuffs = FullFirePartyBuffs
var FullIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings:  true,
	BlessingOfWisdom: proto.TristateEffect_TristateEffectImproved,
}

var FullFireConsumes = &proto.Consumes{
	FlaskOfPureDeath:   true,
	BrilliantWizardOil: true,
	BlackenedBasilisk:  true,
	DefaultPotion:      proto.Potions_SuperManaPotion,
	DarkRune:           true,
}
var FullFrostConsumes = FullFireConsumes

var FullDebuffs = &proto.Debuffs{
	CurseOfElements:           proto.TristateEffect_TristateEffectImproved,
	ImprovedSealOfTheCrusader: true,
	JudgementOfWisdom:         true,
	Misery:                    true,
}

var FullDebuffTarget = &proto.Target{
	Debuffs: FullDebuffs,
}

var P1FireGear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{
	items.ItemStringSpec{
		Name:    "Collar of the Aldor",
		Enchant: "Glyph of Power",
		Gems: []string{
			"Chaotic Skyfire Diamond",
			"Glowing Nightseye",
		},
	},
	items.ItemStringSpec{
		Name: "Brooch of Heightened Potential",
	},
	items.ItemStringSpec{
		Name:    "Pauldrons of the Aldor",
		Enchant: "Greater Inscription of Discipline",
		Gems: []string{
			"Veiled Noble Topaz",
			"Runed Living Ruby",
		},
	},
	items.ItemStringSpec{
		Name: "Ruby Drape of the Mysticant",
	},
	items.ItemStringSpec{
		Name:    "Spellfire Robe",
		Enchant: "Chest - Exceptional Stats",
		Gems: []string{
			"Veiled Noble Topaz",
			"Veiled Noble Topaz",
			"Veiled Noble Topaz",
		},
	},
	items.ItemStringSpec{
		Name:    "General's Silk Cuffs",
		Enchant: "Bracer - Spellpower",
		Gems: []string{
			"Veiled Noble Topaz",
		},
	},
	items.ItemStringSpec{
		Name:    "Spellfire Gloves",
		Enchant: "Gloves - Major Spellpower",
		Gems: []string{
			"Veiled Noble Topaz",
			"Glowing Nightseye",
		},
	},
	items.ItemStringSpec{
		Name: "Spellfire Belt",
		Gems: []string{
			"Veiled Noble Topaz",
			"Veiled Noble Topaz",
		},
	},
	items.ItemStringSpec{
		Name:    "Spellstrike Pants",
		Enchant: "Runic Spellthread",
		Gems: []string{
			"Veiled Noble Topaz",
			"Veiled Noble Topaz",
			"Veiled Noble Topaz",
		},
	},
	items.ItemStringSpec{
		Name: "Boots of Foretelling",
		Gems: []string{
			"Veiled Noble Topaz",
			"Veiled Noble Topaz",
		},
	},
	items.ItemStringSpec{
		Name:    "Band of Crimson Fury",
		Enchant: "Ring - Spellpower",
	},
	items.ItemStringSpec{
		Name:    "Ashyen's Gift",
		Enchant: "Ring - Spellpower",
	},
	items.ItemStringSpec{
		Name: "Quagmirran's Eye",
	},
	items.ItemStringSpec{
		Name: "Icon of the Silver Crescent",
	},
	items.ItemStringSpec{
		Name: "Tirisfal Wand of Ascendancy",
	},
	items.ItemStringSpec{
		Name:    "Bloodmaw Magus-Blade",
		Enchant: "Sunfire",
	},
	items.ItemStringSpec{
		Name: "Flametongue Seal",
	},
})
var P1FrostGear = P1FireGear
