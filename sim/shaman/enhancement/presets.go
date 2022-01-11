package enhancement

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var BasicRaidBuffs = &proto.RaidBuffs{}
var BasicPartyBuffs = &proto.PartyBuffs{
	Bloodlust: 1,
}
var BasicIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings: true,
}

var StandardTalents = &proto.ShamanTalents{
	Convection:         2,
	Concussion:         5,
	CallOfFlame:        3,
	ElementalFocus:     true,
	Reverberation:      5,
	ImprovedFireTotems: 1,

	AncestralKnowledge:      5,
	ThunderingStrikes:       5,
	EnhancingTotems:         2,
	ShamanisticFocus:        true,
	Flurry:                  5,
	ImprovedWeaponTotems:    1,
	ElementalWeapons:        3,
	MentalQuickness:         3,
	WeaponMastery:           5,
	DualWieldSpecialization: 3,
	Stormstrike:             true,
	UnleashedRage:           5,
	ShamanisticRage:         true,
}

var PlayerOptionsBasic = &proto.Player_EnhancementShaman{
	EnhancementShaman: &proto.EnhancementShaman{
		Talents:  StandardTalents,
		Options:  enhShamOptions,
		Rotation: enhShamRotation,
	},
}

var enhShamRotation = &proto.EnhancementShaman_Rotation{
	Totems: &proto.ShamanTotems{
		Earth: proto.EarthTotem_StrengthOfEarthTotem,
		Air:   proto.AirTotem_GraceOfAirTotem,
		Water: proto.WaterTotem_ManaSpringTotem,
		Fire:  proto.FireTotem_NoFireTotem, // TODO: deal with fire totem later... can fire totems just be a DoT?
	},
}

var enhShamOptions = &proto.EnhancementShaman_Options{
	WaterShield: true,
	Bloodlust:   true,
}

var FullRaidBuffs = &proto.RaidBuffs{
	ArcaneBrilliance: true,
	GiftOfTheWild:    proto.TristateEffect_TristateEffectImproved,
}
var FullPartyBuffs = &proto.PartyBuffs{
	FerociousInspiration: 2,
	BattleShout:          proto.TristateEffect_TristateEffectImproved,
	LeaderOfThePack:      proto.TristateEffect_TristateEffectImproved,
	SanctityAura:         proto.TristateEffect_TristateEffectImproved,
	TrueshotAura:         true,
}
var FullIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings:  true,
	BlessingOfWisdom: proto.TristateEffect_TristateEffectImproved,
	BlessingOfMight:  proto.TristateEffect_TristateEffectImproved,
}

var FullConsumes = &proto.Consumes{
	Drums: proto.Drums_DrumsOfBattle,
}

var NoDebuffTarget = &proto.Target{
	Debuffs: &proto.Debuffs{},
	Armor:   6700,
}

var FullDebuffs = &proto.Debuffs{
	BloodFrenzy:                 true,
	ExposeArmor:                 proto.TristateEffect_TristateEffectImproved,
	FaerieFire:                  proto.TristateEffect_TristateEffectImproved,
	ImprovedSealOfTheCrusader:   true,
	JudgementOfWisdom:           true,
	Misery:                      true,
	ExposeWeaknessUptime:        0.8,
	ExposeWeaknessHunterAgility: 800,
}

var FullDebuffTarget = &proto.Target{
	Debuffs: FullDebuffs,
	Armor:   7700,
}

var Phase2Gear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{
	{
		Name:    "Cataclysm Helm",
		Enchant: "Glyph of Ferocity",
		Gems: []string{
			"Relentless Earthstorm Diamond",
			"Bold Living Ruby",
		},
	},
	{
		Name: "Telonicus's Pendant of Mayhem",
	},
	{
		Name:    "Shoulderpads of the Stranger",
		Enchant: "Greater Inscription of Vengeance",
		Gems: []string{
			"Bold Living Ruby",
		},
	},
	{
		Name: "Thalassian Wildercloak",
	},
	{
		Name:    "Cataclysm Chestplate",
		Enchant: "Chest - Exceptional Stats",
		Gems: []string{
			"Bold Living Ruby",
			"Sovereign Nightseye",
			"Inscribed Noble Topaz",
		},
	},
	{
		Name:    "True-Aim Stalker Bands",
		Enchant: "Bracer - Brawn",
		Gems: []string{
			"Bold Living Ruby",
		},
	},
	{
		Name:    "Cataclysm Gauntlets",
		Enchant: "Gloves - Major Strength",
	},
	{
		Name: "Belt of One-Hundred Deaths",
		Gems: []string{
			"Bold Living Ruby",
			"Sovereign Nightseye",
		},
	},
	{
		Name:    "Cataclysm Legplates",
		Enchant: "Nethercobra Leg Armor",
		Gems: []string{
			"Bold Living Ruby",
		},
	},
	{
		Name:    "Boots of Utter Darkness",
		Enchant: "Enchant Boots - Cat's Swiftness",
	},
	{
		Name: "Ring of Lethality",
	},
	{
		Name: "Band of the Ranger-General",
	},
	{
		Name: "Dragonspine Trophy",
	},
	{
		Name: "Bloodlust Brooch",
	},
	{
		Name:    "Talon of the Phoenix",
		Enchant: "Weapon - Mongoose",
	},
	{
		Name:    "Rod of the Sun King",
		Enchant: "Weapon - Mongoose",
	},
	{
		Name: "Totem of the Astral Winds",
	},
})

var PreRaidGear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{
	items.ItemStringSpec{
		Name: "Gladiator's Cleaver",
	},
})
