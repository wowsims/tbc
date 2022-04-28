package balance

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var StandardTalents = &proto.DruidTalents{
	StarlightWrath:        5,
	FocusedStarlight:      2,
	ImprovedMoonfire:      2,
	Brambles:              3,
	InsectSwarm:           true,
	Vengeance:             5,
	LunarGuidance:         3,
	NaturesGrace:          true,
	Moonglow:              3,
	Moonfury:              5,
	BalanceOfPower:        2,
	Dreamstate:            3,
	MoonkinForm:           true,
	ImprovedFaerieFire:    3,
	WrathOfCenarius:       5,
	ForceOfNature:         true,
	ImprovedMarkOfTheWild: 5,
	Furor:                 2,
	NaturalShapeshifter:   3,
	Intensity:             3,
}

var FullRaidBuffs = &proto.RaidBuffs{
	ArcaneBrilliance: true,
	GiftOfTheWild:    proto.TristateEffect_TristateEffectImproved,
}
var FullPartyBuffs = &proto.PartyBuffs{
	MoonkinAura: proto.TristateEffect_TristateEffectRegular,
	Drums:       proto.Drums_DrumsOfBattle,
}
var FullIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings:  true,
	BlessingOfWisdom: proto.TristateEffect_TristateEffectImproved,
	ShadowPriestDps:  500,
}

var FullConsumes = &proto.Consumes{
	Flask:              proto.Flask_FlaskOfBlindingLight,
	Food:               proto.Food_FoodBlackenedBasilisk,
	DefaultPotion:      proto.Potions_SuperManaPotion,
	StartingPotion:     proto.Potions_DestructionPotion,
	MainHandImbue:      proto.WeaponImbue_WeaponImbueBrilliantWizardOil,
	NumStartingPotions: 1,
	DefaultConjured:    proto.Conjured_ConjuredDarkRune,
}

var FullDebuffs = &proto.Debuffs{
	JudgementOfWisdom: true,
	Misery:            true,
	CurseOfElements:   proto.TristateEffect_TristateEffectImproved,
}

var FullDebuffTarget = &proto.Target{
	Debuffs: FullDebuffs,
}

var PlayerOptionsAdaptive = &proto.Player_BalanceDruid{
	BalanceDruid: &proto.BalanceDruid{
		Talents: StandardTalents,
		Options: &proto.BalanceDruid_Options{
			InnervateTarget: &proto.RaidTarget{TargetIndex: 0}, // self innervate
		},
		Rotation: &proto.BalanceDruid_Rotation{
			PrimarySpell: proto.BalanceDruid_Rotation_Adaptive,
			FaerieFire:   true,
		},
	},
}

var PlayerOptionsStarfire = &proto.Player_BalanceDruid{
	BalanceDruid: &proto.BalanceDruid{
		Talents: StandardTalents,
		Options: &proto.BalanceDruid_Options{
			InnervateTarget: &proto.RaidTarget{TargetIndex: 0}, // self innervate
		},
		Rotation: &proto.BalanceDruid_Rotation{
			PrimarySpell: proto.BalanceDruid_Rotation_Starfire,
			Moonfire:     true,
			FaerieFire:   true,
		},
	},
}

var PlayerOptionsWrath = &proto.Player_BalanceDruid{
	BalanceDruid: &proto.BalanceDruid{
		Talents: StandardTalents,
		Options: &proto.BalanceDruid_Options{
			InnervateTarget: &proto.RaidTarget{TargetIndex: 0}, // self innervate
		},
		Rotation: &proto.BalanceDruid_Rotation{
			PrimarySpell: proto.BalanceDruid_Rotation_Wrath,
			Moonfire:     true,
		},
	},
}

var PlayerOptionsAOE = &proto.Player_BalanceDruid{
	BalanceDruid: &proto.BalanceDruid{
		Talents: StandardTalents,
		Options: &proto.BalanceDruid_Options{
			InnervateTarget: &proto.RaidTarget{TargetIndex: 0}, // self innervate
		},
		Rotation: &proto.BalanceDruid_Rotation{
			PrimarySpell: proto.BalanceDruid_Rotation_Starfire,
			Moonfire:     true,
			FaerieFire:   true,
			Hurricane:    true,
		},
	},
}

var P1Gear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{
	{
		Name:    "Antlers of Malorne",
		Enchant: "Glyph of Power",
		Gems: []string{
			"Potent Noble Topaz",
			"Chaotic Skyfire Diamond",
		},
	},
	{
		Name: "Adornment of Stolen Souls",
	},
	{
		Name:    "Pauldrons of Malorne",
		Enchant: "Greater Inscription of Discipline",
		Gems: []string{
			"Glowing Nightseye",
			"Potent Noble Topaz",
		},
	},
	{
		Name: "Ruby Drape of the Mysticant",
	},
	{
		Name:    "Spellfire Robe",
		Enchant: "Chest - Exceptional Stats",
		Gems: []string{
			"Potent Noble Topaz",
			"Glowing Nightseye",
		},
	},
	{
		Name:    "Bracers of Havok",
		Enchant: "Bracer - Spellpower",
		Gems: []string{
			"Potent Noble Topaz",
		},
	},
	{
		Name:    "Spellfire Gloves",
		Enchant: "Gloves - Major Spellpower",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name: "Spellfire Belt",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name:    "Spellstrike Pants",
		Enchant: "Runic Spellthread",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name: "Boots of Foretelling",
	},
	{
		Name:    "Violet Signet of the Archmage",
		Enchant: "Ring - Spellpower",
	},
	{
		Name:    "Ring of Recurrence",
		Enchant: "Ring - Spellpower",
	},
	{
		Name: "Quagmirran's Eye",
	},
	{
		Name: "Icon of the Silver Crescent",
	},
	{
		Name: "Ivory Idol of the Moongoddess",
	},
	{
		Name:    "Nathrezim Mindblade",
		Enchant: "Weapon - Major Spellpower",
	},
	{
		Name: "Talisman of Kalecgos",
	},
})

var P2Gear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{
	{
		Name:    "Nordrassil Headpiece",
		Enchant: "Glyph of Power",
		Gems: []string{
			"Potent Noble Topaz",
			"Chaotic Skyfire Diamond",
		},
	},
	{
		Name: "The Sun King's Talisman",
	},
	{
		Name:    "Nordrassil Wrath-Mantle",
		Enchant: "Greater Inscription of Discipline",
		Gems: []string{
			"Glowing Nightseye",
			"Potent Noble Topaz",
		},
	},
	{
		Name: "Brute Cloak of the Ogre-Magi",
	},
	{
		Name:    "Nordrassil Chestpiece",
		Enchant: "Chest - Exceptional Stats",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name:    "Mindstorm Wristbands",
		Enchant: "Bracer - Spellpower",
	},
	{
		Name:    "Spellfire Gloves",
		Enchant: "Gloves - Major Spellpower",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name: "Belt of Blasting",
		Gems: []string{
			"Glowing Nightseye",
			"Potent Noble Topaz",
		},
	},
	{
		Name:    "Nordrassil Wrath-Kilt",
		Enchant: "Runic Spellthread",
		Gems: []string{
			"Runed Living Ruby",
		},
	},
	{
		Name: "Boots of Blasting",
	},
	{
		Name:    "Band of Eternity",
		Enchant: "Ring - Spellpower",
	},
	{
		Name:    "Ring of Recurrence",
		Enchant: "Ring - Spellpower",
	},
	{
		Name: "Quagmirran's Eye",
	},
	{
		Name: "Icon of the Silver Crescent",
	},
	{
		Name: "Idol of the Raven Goddess",
	},
	{
		Name:    "The Nexus Key",
		Enchant: "Weapon - Major Spellpower",
	},
})
