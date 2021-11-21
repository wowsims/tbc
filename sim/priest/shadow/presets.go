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

var StandardTalents = &proto.PriestTalents{}

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
	FlaskOfBlindingLight: true,
	BrilliantWizardOil:   true,
	BlackenedBasilisk:    true,
	DefaultPotion:        proto.Potions_SuperManaPotion,
	StartingPotion:       proto.Potions_DestructionPotion,
	NumStartingPotions:   1,
	DarkRune:             true,
	Drums:                proto.Drums_DrumsOfBattle,
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

var PlayerOptionsAdaptive = &proto.PlayerOptions{
	Spec: &proto.PlayerOptions_ShadowPriest{
		ShadowPriest: &proto.ShadowPriest{
			Talents:  StandardTalents,
			Options:  &proto.ShadowPriest_Options{},
			Rotation: &proto.ShadowPriest_Rotation{},
		},
	},
}

var PreRaidGear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{})

// TODO: fill out a p1 gear for spriest.
var P1Gear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{
	{
		Name:    "Spellstrike Hood", // TODO: helm
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
		Name:    "Frozen Shadoweave Shoulders", // TODO: shoulders
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
		Name:    "Frozen Shadoweave Robe", // TODO: chest
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
		Name:    "", // TODO: gloves
		Enchant: "Gloves - Major Spellpower",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name: "", // TODO: waist
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	{
		Name:    "", // TODO: legs
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
		Name:    "Nathrezim Mindblade",
		Enchant: "Weapon - Major Spellpower",
	},
	{
		Name: "Talisman of Kalecgos",
	},
})
