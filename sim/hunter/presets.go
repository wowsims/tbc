package hunter

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var BMTalents = &proto.HunterTalents{
	ImprovedAspectOfTheHawk: 5,
	FocusedFire:             2,
	UnleashedFury:           5,
	Ferocity:                5,
	BestialDiscipline:       2,
	AnimalHandler:           1,
	Frenzy:                  5,
	FerociousInspiration:    3,
	BestialWrath:            true,
	SerpentsSwiftness:       5,
	TheBeastWithin:          true,

	LethalShots:    5,
	Efficiency:     5,
	GoForTheThroat: 2,
	AimedShot:      true,
	RapidKilling:   2,
	MortalShots:    5,
}

var PlayerOptionsBasic = &proto.Player_Hunter{
	Hunter: &proto.Hunter{
		Talents:  BMTalents,
		Options:  basicOptions,
		Rotation: basicRotation,
	},
}

var PlayerOptionsFrench = &proto.Player_Hunter{
	Hunter: &proto.Hunter{
		Talents:  BMTalents,
		Options:  windSerpentOptions,
		Rotation: frenchRotation,
	},
}

var PlayerOptionsMeleeWeave = &proto.Player_Hunter{
	Hunter: &proto.Hunter{
		Talents:  BMTalents,
		Options:  windSerpentOptions,
		Rotation: meleeWeaveRotation,
	},
}

var basicRotation = &proto.Hunter_Rotation{
	UseMultiShot:     true,
	UseArcaneShot:    false,
	Sting:            proto.Hunter_Rotation_SerpentSting,
	PrecastAimedShot: true,
	MeleeWeave:       false,

	ViperStartManaPercent: 0.2,
	ViperStopManaPercent:  0.3,
}
var frenchRotation = &proto.Hunter_Rotation{
	UseMultiShot:      true,
	UseArcaneShot:     true,
	Sting:             proto.Hunter_Rotation_SerpentSting,
	PrecastAimedShot:  false,
	MeleeWeave:        false,
	UseFrenchRotation: true,

	ViperStartManaPercent: 0.3,
	ViperStopManaPercent:  0.5,
}
var meleeWeaveRotation = &proto.Hunter_Rotation{
	UseMultiShot:    true,
	UseArcaneShot:   false,
	MeleeWeave:      true,
	UseRaptorStrike: true,
	TimeToWeaveMs:   500,
	PercentWeaved:   0.8,

	ViperStartManaPercent: 0.3,
	ViperStopManaPercent:  0.5,
}

var basicOptions = &proto.Hunter_Options{
	QuiverBonus: proto.Hunter_Options_Speed15,
	Ammo:        proto.Hunter_Options_AdamantiteStinger,
	PetType:     proto.Hunter_Options_Ravager,
	PetUptime:   0.9,
	LatencyMs:   15,
}

var windSerpentOptions = &proto.Hunter_Options{
	QuiverBonus: proto.Hunter_Options_Speed15,
	Ammo:        proto.Hunter_Options_AdamantiteStinger,
	PetType:     proto.Hunter_Options_WindSerpent,
	PetUptime:   0.9,
	LatencyMs:   15,
}

var FullRaidBuffs = &proto.RaidBuffs{
	ArcaneBrilliance: true,
	GiftOfTheWild:    proto.TristateEffect_TristateEffectImproved,
}
var FullPartyBuffs = &proto.PartyBuffs{
	Bloodlust: 1,
	Drums:     proto.Drums_DrumsOfBattle,

	BattleShout:       proto.TristateEffect_TristateEffectImproved,
	LeaderOfThePack:   proto.TristateEffect_TristateEffectImproved,
	ManaSpringTotem:   proto.TristateEffect_TristateEffectRegular,
	GraceOfAirTotem:   proto.TristateEffect_TristateEffectRegular,
	WindfuryTotemRank: 5,
}
var FullIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings:  true,
	BlessingOfWisdom: proto.TristateEffect_TristateEffectImproved,
	BlessingOfMight:  proto.TristateEffect_TristateEffectImproved,
}

var FullConsumes = &proto.Consumes{
	Flask:           proto.Flask_FlaskOfRelentlessAssault,
	DefaultPotion:   proto.Potions_HastePotion,
	DefaultConjured: proto.Conjured_ConjuredDarkRune,
	PetFood:         proto.PetFood_PetFoodKiblersBits,
}

var FullDebuffs = &proto.Debuffs{
	BloodFrenzy:               true,
	FaerieFire:                proto.TristateEffect_TristateEffectImproved,
	ImprovedSealOfTheCrusader: true,
	JudgementOfWisdom:         true,
	Misery:                    true,
}

var FullDebuffTarget = &proto.Target{
	Debuffs: FullDebuffs,
	Armor:   7700,
}

var P1Gear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{
	{
		Name:    "Beast Lord Helm",
		Enchant: "Glyph of Ferocity",
		Gems: []string{
			"Delicate Living Ruby",
			"Relentless Earthstorm Diamond",
		},
	},
	{
		Name: "Choker of Vile Intent",
	},
	{
		Name:    "Beast Lord Mantle",
		Enchant: "Greater Inscription of Vengeance",
		Gems: []string{
			"Delicate Living Ruby",
			"Delicate Living Ruby",
		},
	},
	{
		Name: "Vengeance Wrap",
	},
	{
		Name:    "Beast Lord Cuirass",
		Enchant: "Chest - Exceptional Stats",
		Gems: []string{
			"Delicate Living Ruby",
			"Delicate Living Ruby",
			"Delicate Living Ruby",
		},
	},
	{
		Name:    "Nightfall Wristguards",
		Enchant: "Bracer - Assault",
	},
	{
		Name:    "Beast Lord Handguards",
		Enchant: "Gloves - Major Agility",
		Gems: []string{
			"Delicate Living Ruby",
			"Delicate Living Ruby",
		},
	},
	{
		Name: "Gronn-Stitched Girdle",
		Gems: []string{
			"Delicate Living Ruby",
			"Delicate Living Ruby",
		},
	},
	{
		Name:    "Scaled Greaves of the Marksman",
		Enchant: "Nethercobra Leg Armor",
		Gems: []string{
			"Delicate Living Ruby",
			"Delicate Living Ruby",
			"Delicate Living Ruby",
		},
	},
	{
		Name:    "Edgewalker Longboots",
		Enchant: "Enchant Boots - Cat's Swiftness",
		Gems: []string{
			"Delicate Living Ruby",
			"Delicate Living Ruby",
		},
	},
	{
		Name: "Ring of a Thousand Marks",
	},
	{
		Name: "Ring of the Recalcitrant",
	},
	{
		Name: "Dragonspine Trophy",
	},
	{
		Name: "Bloodlust Brooch",
	},
	{
		Name:    "Mooncleaver",
		Enchant: "2H Weapon - Major Agility",
	},
	{
		Name:    "Sunfury Bow of the Phoenix",
		Enchant: "Stabilized Eternium Scope",
	},
})
