package protection

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var ImpaleProtTalents = &proto.WarriorTalents{
	ImprovedHeroicStrike: 3,
	Deflection:           5,
	ImprovedThunderClap:  3,
	AngerManagement:      true,
	DeepWounds:           3,
	Impale:               2,

	Cruelty: 3,

	Anticipation:                  5,
	ShieldSpecialization:          5,
	Toughness:                     5,
	ImprovedShieldBlock:           true,
	Defiance:                      3,
	ImprovedSunderArmor:           3,
	ShieldMastery:                 1,
	OneHandedWeaponSpecialization: 5,
	ShieldSlam:                    true,
	FocusedRage:                   3,
	Vitality:                      5,
	Devastate:                     true,
}

var PlayerOptionsBasic = &proto.Player_ProtectionWarrior{
	ProtectionWarrior: &proto.ProtectionWarrior{
		Talents:  ImpaleProtTalents,
		Options:  warriorOptions,
		Rotation: warriorRotation,
	},
}

var warriorRotation = &proto.ProtectionWarrior_Rotation{}

var warriorOptions = &proto.ProtectionWarrior_Options{
	StartingRage:    0,
	PrecastT2:       false,
	PrecastSapphire: false,
}

var FullRaidBuffs = &proto.RaidBuffs{
	ArcaneBrilliance: true,
	GiftOfTheWild:    proto.TristateEffect_TristateEffectImproved,
}
var FullPartyBuffs = &proto.PartyBuffs{
	Bloodlust:            1,
	Drums:                proto.Drums_DrumsOfBattle,
	BattleShout:          proto.TristateEffect_TristateEffectImproved,
	LeaderOfThePack:      proto.TristateEffect_TristateEffectImproved,
	GraceOfAirTotem:      proto.TristateEffect_TristateEffectRegular,
	StrengthOfEarthTotem: proto.StrengthOfEarthType_EnhancingTotems,
	WindfuryTotemRank:    5,
	WindfuryTotemIwt:     2,
}
var FullIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings: true,
	BlessingOfMight: proto.TristateEffect_TristateEffectImproved,
	UnleashedRage:   true,
}

var FullConsumes = &proto.Consumes{}

var FullDebuffs = &proto.Debuffs{
	BloodFrenzy:               true,
	FaerieFire:                proto.TristateEffect_TristateEffectImproved,
	ImprovedSealOfTheCrusader: true,
	Misery:                    true,
}

var FullDebuffTarget = &proto.Target{
	Debuffs: FullDebuffs,
	Armor:   7684,
}

var P1Gear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{
	{
		Name:    "Warbringer Greathelm",
		Enchant: "Glyph of Ferocity",
		Gems: []string{
			"Powerful Earthstorm Diamond",
			"Solid Star of Elune",
		},
	},
	{
		Name: "Pendant of Triumph",
		Gems: []string{
			"Steady Talasite",
		},
	},
	{
		Name:    "Warbringer Shoulderplates",
		Enchant: "Greater Inscription of the Knight",
		Gems: []string{
			"Solid Star of Elune",
			"Solid Star of Elune",
		},
	},
	{
		Name:    "Drape of the Dark Reavers",
		Enchant: "Enchant Cloak - Greater Agility",
	},
	{
		Name:    "Warbringer Chestguard",
		Enchant: "Chest - Exceptional Stats",
		Gems: []string{
			"Solid Star of Elune",
			"Solid Star of Elune",
			"Solid Star of Elune",
		},
	},
	{
		Name:    "Marshal's Plate Bracers",
		Enchant: "Bracer - Fortitude",
		Gems: []string{
			"Steady Talasite",
		},
	},
	{
		Name:    "Grips of Deftness",
		Enchant: "Gloves - Threat",
	},
	{
		Name: "Marshal's Plate Belt",
	},
	{
		Name:    "Wrynn Dynasty Greaves",
		Enchant: "Nethercleft Leg Armor",
		Gems: []string{
			"Solid Star of Elune",
			"Solid Star of Elune",
			"Solid Star of Elune",
		},
	},
	{
		Name:    "Battlescar Boots",
		Enchant: "Enchant Boots - Boar's Speed",
		Gems: []string{
			"Solid Star of Elune",
			"Solid Star of Elune",
		},
	},
	{
		Name: "Violet Signet of the Great Protector",
	},
	{
		Name: "Shapeshifter's Signet",
	},
	{
		Name: "Icon of Unyielding Courage",
	},
	{
		Name: "Gnomeregan Auto-Blocker 600",
	},
	{
		Name:    "King's Defender",
		Enchant: "Weapon - Mongoose",
	},
	{
		Name:    "Aldori Legacy Defender",
		Enchant: "Shield - Major Stamina",
		Gems: []string{
			"Solid Star of Elune",
		},
	},
	{
		Name: "Shuriken of Negation",
	},
})
