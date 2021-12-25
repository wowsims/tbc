package shadow

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var StandardTalents = &proto.PriestTalents{
	ImprovedShadowWordPain: 2,
	ImprovedMindBlast:      5,
	ShadowFocus:            5,
	MindFlay:               true,
	ShadowWeaving:          5,
	VampiricEmbrace:        true,
	FocusedMind:            3,
	Darkness:               5,
	Shadowform:             true,
	Misery:                 5,
	VampiricTouch:          true,
	InnerFocus:             true,
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
	FlaskOfPureDeath:   true,
	BrilliantWizardOil: true,
	BlackenedBasilisk:  true,
	DefaultPotion:      proto.Potions_SuperManaPotion,
	NumStartingPotions: 1,
	DarkRune:           true,
}

var FullDebuffs = &proto.Debuffs{
	JudgementOfWisdom: true,
	CurseOfElements:   proto.TristateEffect_TristateEffectImproved,
}

var FullDebuffTarget = &proto.Target{
	Debuffs: FullDebuffs,
}

var PlayerOptionsBasic = &proto.Player_ShadowPriest{
	ShadowPriest: &proto.ShadowPriest{
		Talents: StandardTalents,
		Options: &proto.ShadowPriest_Options{
			UseShadowfiend: true,
		},
		Rotation: &proto.ShadowPriest_Rotation{
			RotationType: proto.ShadowPriest_Rotation_Basic,
		},
	},
}
var PlayerOptionsClipping = &proto.Player_ShadowPriest{
	ShadowPriest: &proto.ShadowPriest{
		Talents: StandardTalents,
		Options: &proto.ShadowPriest_Options{
			UseShadowfiend: true,
		},
		Rotation: &proto.ShadowPriest_Rotation{
			RotationType: proto.ShadowPriest_Rotation_Clipping,
			PrecastVt:    true,
		},
	},
}
var PlayerOptionsIdeal = &proto.Player_ShadowPriest{
	ShadowPriest: &proto.ShadowPriest{
		Talents: StandardTalents,
		Options: &proto.ShadowPriest_Options{
			UseShadowfiend: true,
		},
		Rotation: &proto.ShadowPriest_Rotation{
			RotationType: proto.ShadowPriest_Rotation_Ideal,
			PrecastVt:    true,
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
