package smite

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var StandardTalents = &proto.PriestTalents{
	InnerFocus:             true,
	Meditation:             3,
	SilentResolve:          1,
	MentalAgility:          5,
	MentalStrength:         5,
	DivineSpirit:           true,
	ImprovedDivineSpirit:   2,
	ForceOfWill:            5,
	PowerInfusion:          true,
	HolySpecialization:     5,
	DivineFury:             5,
	SearingLight:           2,
	SpiritualGuidance:      5,
	SurgeOfLight:           2,
}

var FullRaidBuffs = &proto.RaidBuffs{
	ArcaneBrilliance: true,
	GiftOfTheWild:    proto.TristateEffect_TristateEffectImproved,
}
var FullPartyBuffs = &proto.PartyBuffs{
	MoonkinAura:     proto.TristateEffect_TristateEffectRegular,
	TotemOfWrath:    1,
	WrathOfAirTotem: proto.TristateEffect_TristateEffectImproved,
	ManaSpringTotem: proto.TristateEffect_TristateEffectRegular,
}
var FullIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings:  true,
	BlessingOfWisdom: proto.TristateEffect_TristateEffectImproved,
}

var FullConsumes = &proto.Consumes{
	Flask:              proto.Flask_FlaskOfBlindingLight,
	Food:               proto.Food_FoodBlackenedBasilisk,
	DefaultPotion:      proto.Potions_SuperManaPotion,
	NumStartingPotions: 1,
	DefaultConjured:    proto.Conjured_ConjuredDarkRune,
	MainHandImbue:      proto.WeaponImbue_WeaponImbueBrilliantWizardOil,
}

var FullDebuffs = &proto.Debuffs{
	JudgementOfWisdom: true,
	CurseOfElements:   proto.TristateEffect_TristateEffectImproved,
}

var FullDebuffTarget = &proto.Target{
	Debuffs: FullDebuffs,
}

var PlayerOptionsBasic = &proto.Player_SmitePriest{
	SmitePriest: &proto.SmitePriest{
		Talents: StandardTalents,
		Options: &proto.SmitePriest_Options{
			UseShadowfiend: true,
		},
		Rotation: &proto.SmitePriest_Rotation{
			RotationType: proto.SmitePriest_Rotation_Basic,
		},
	},
}


var P1Gear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{
	{
		Name:    "Spellstrike Hood",
		Enchant: "Glyph of Power",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name: "Ritssyn's Lost Pendant",
	},
	{
		Name:    "Frozen Shadoweave Shoulders",
		Enchant: "Greater Inscription of Discipline",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name: "Shadow-Cloak of Dalaran",
	},
	{
		Name:    "Frozen Shadoweave Robe",
		Enchant: "Chest - Exceptional Stats",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name:    "Bracers of Havok",
		Enchant: "Bracer - Spellpower",
	},
	{
		Name:    "Handwraps of Flowing Thought",
		Enchant: "Gloves - Major Spellpower",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name: "Belt of Divine Inspiration",
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
		Name: "Frozen Shadoweave Boots",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name: "Cobalt Band of Tyrigosa",
	},
	{
		Name: "Band of Crimson Fury",
	},
	{
		Name: "Eye of Magtheridon",
	},
	{
		Name: "Icon of the Silver Crescent",
	},
	{
		Name: "The Black Stalk",
	},
	{
		Name:    "Nathrezim Mindblade",
		Enchant: "Weapon - Major Spellpower",
	},
	{
		Name: "Orb of the Soul-Eater",
	},
})

var P3Gear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{
	{
		Name:    "Hood of Absolution",
		Enchant: "Glyph of Power",
		Gems: []string{
			"Mystical Skyfire Diamond",
			"Glowing Nightseye",
		},
	},
	{
		Name: "Ritssyn's Lost Pendant",
	},
	{
		Name:    "Shoulderpads of Absolution",
		Enchant: "Greater Inscription of Discipline",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name: "Nethervoid Cloak",
	},
	{
		Name:    "Shroud of Absolution",
		Enchant: "Chest - Exceptional Stats",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name:    "Bracers of Nimble Thought",
		Enchant: "Bracer - Spellpower",
	},
	{
		Name:    "Handguards of Absolution",
		Enchant: "Gloves - Major Spellpower",
		Gems: []string{
			"Runed Living Ruby",
		},
	},
	{
		Name: "Waistwrap of Infinity",
	},
	{
		Name:    "Leggings of Channeled Elements",
		Enchant: "Runic Spellthread",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name: "Slippers of the Seacaller",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name: "Ring of Ancient Knowledge",
	},
	{
		Name: "Ring of Ancient Knowledge",
	},
	{
		Name: "The Skull of Gul'dan",
	},
	{
		Name: "Icon of the Silver Crescent",
	},
	{
		Name: "Wand of the Forgotten Star",
	},
	{
		Name:    "The Maelstrom's Fury",
		Enchant: "Weapon - Major Spellpower",
	},
	{
		Name: "Orb of the Soul-Eater",
	},
})
