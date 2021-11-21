package shadow

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var BasicRaidBuffs = &proto.RaidBuffs{}
var BasicPartyBuffs = &proto.PartyBuffs{
	Bloodlust: 1,
}
var BasicIndividualBuffs = &proto.IndividualBuffs{}

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
	// TODO: Inner Focus
}

var FullRaidBuffs = &proto.RaidBuffs{
	ArcaneBrilliance: true,
	GiftOfTheWild:    proto.TristateEffect_TristateEffectImproved,
}
var FullPartyBuffs = &proto.PartyBuffs{
	MoonkinAura: proto.TristateEffect_TristateEffectRegular,
}
var FullIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings:  true,
	BlessingOfWisdom: proto.TristateEffect_TristateEffectImproved,
}

var BasicConsumes = &proto.Consumes{
	DefaultPotion: proto.Potions_SuperManaPotion,
}
var FullConsumes = &proto.Consumes{
	FlaskOfPureDeath:   true,
	BrilliantWizardOil: true,
	BlackenedBasilisk:  true,
	DefaultPotion:      proto.Potions_SuperManaPotion,
	NumStartingPotions: 1,
	DarkRune:           true,
}

var NoDebuffTarget = &proto.Target{
	Debuffs: &proto.Debuffs{},
}

var FullDebuffTarget = &proto.Target{
	Debuffs: &proto.Debuffs{
		JudgementOfWisdom: true,
		CurseOfElements:   proto.TristateEffect_TristateEffectImproved,
	},
}

var PlayerOptionsBasic = &proto.PlayerOptions{
	Spec: &proto.PlayerOptions_ShadowPriest{
		ShadowPriest: &proto.ShadowPriest{
			Talents: StandardTalents,
			Options: &proto.ShadowPriest_Options{},
			Rotation: &proto.ShadowPriest_Rotation{
				Type:   proto.ShadowPriest_Rotation_Basic,
				UseSwd: true,
			},
		},
	},
}

var PreRaidGear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{})

// TODO: fill out a p1 gear for spriest.
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
		Name: "Ruby Drape of the Mysticant",
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
		Name:    "Soul-Eater's Handwraps",
		Enchant: "Gloves - Major Spellpower",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name: "Girdle of Ruination",
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
		Name: "Frozen Shadoweave Boots",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name: "Band of the Inevitable",
	},
	{
		Name: "Ring of the Fallen God",
	},
	{
		Name: "Quagmirran's Eye",
	},
	{
		Name: "Icon of the Silver Crescent",
	},
	{
		Name:    "Nathrezim Mindblade",
		Enchant: "Weapon - Major Spellpower",
	},
	{
		Name: "Orb of the Soul-Eater",
	},
})
