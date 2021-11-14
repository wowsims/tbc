package balance

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var BasicRaidBuffs = &proto.RaidBuffs{}
var BasicPartyBuffs = &proto.PartyBuffs{
	Bloodlust: 1,
}
var BasicIndividualBuffs = &proto.IndividualBuffs{}

var StandardTalents = &proto.DruidTalents{}

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
	ShadowPriestDps:  500,
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
	},
}

var PlayerOptionsStarfire = &proto.PlayerOptions{
	Spec: &proto.PlayerOptions_BalanceDruid{
		BalanceDruid: &proto.BalanceDruid{
			Talents: StandardTalents,
			Options: &proto.BalanceDruid_Options{
				// InnervateTarget: proto.,
			},
			Rotation: &proto.BalanceDruid_Rotation{
				PrimarySpell: proto.BalanceDruid_Rotation_Starfire,
				Moonfire:     true,
			},
		},
	},
}

var PreRaidGear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{})

var P1Gear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{
	items.ItemStringSpec{
		Name:    "Antlers of Malorne",
		Enchant: "Glyph of Power",
		Gems: []string{
			"Potent Noble Topaz",
			"Chaotic Skyfire Diamond",
		},
	},
	items.ItemStringSpec{
		Name: "Adornment of Stolen Souls",
	},
	items.ItemStringSpec{
		Name:    "Pauldrons of Malorne",
		Enchant: "Greater Inscription of Discipline",
		Gems: []string{
			"Glowing Nightseye",
			"Potent Noble Topaz",
		},
	},
	items.ItemStringSpec{
		Name: "Ruby Drape of the Mysticant",
	},
	items.ItemStringSpec{
		Name:    "Spellfire Robe",
		Enchant: "Chest - Exceptional Stats",
		Gems: []string{
			"Potent Noble Topaz",
			"Glowing Nightseye",
		},
	},
	items.ItemStringSpec{
		Name:    "Bracers of Havok",
		Enchant: "Bracer - Spellpower",
		Gems: []string{
			"Potent Noble Topaz",
		},
	},
	items.ItemStringSpec{
		Name:    "Spellfire Gloves",
		Enchant: "Gloves - Major Spellpower",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	items.ItemStringSpec{
		Name: "Spellfire Belt",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	items.ItemStringSpec{
		Name:    "Spellstrike Pants",
		Enchant: "Runic Spellthread",
		Gems: []string{
			"Runed Living Ruby",
			"Runed Living Ruby",
			"Runed Living Ruby",
		},
	},
	items.ItemStringSpec{
		Name: "Boots of Foretelling",
	},
	items.ItemStringSpec{
		Name:    "Violet Signet of the Archmage",
		Enchant: "Ring - Spellpower",
	},
	items.ItemStringSpec{
		Name:    "Ring of Recurrence",
		Enchant: "Ring - Spellpower",
	},
	items.ItemStringSpec{
		Name: "Quagmirran's Eye",
	},
	items.ItemStringSpec{
		Name: "Icon of the Silver Crescent",
	},
	items.ItemStringSpec{
		Name: "Ivory Idol of the Moongoddess",
	},
	items.ItemStringSpec{
		Name:    "Nathrezim Mindblade",
		Enchant: "Weapon - Major Spellpower",
	},
	items.ItemStringSpec{
		Name: "Talisman of Kalecgos",
	},
})
