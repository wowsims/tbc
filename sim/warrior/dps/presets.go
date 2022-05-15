package dps

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var PlayerOptionsArmsSlam = &proto.Player_Warrior{
	Warrior: &proto.Warrior{
		Talents:  ArmsSlamTalents,
		Options:  warriorOptions,
		Rotation: armsSlamRotation,
	},
}

var PlayerOptionsFury = &proto.Player_Warrior{
	Warrior: &proto.Warrior{
		Talents:  FuryTalents,
		Options:  warriorOptions,
		Rotation: warriorRotation,
	},
}

var ArmsSlamTalents = &proto.WarriorTalents{
	ImprovedHeroicStrike:          3,
	Deflection:                    2,
	ImprovedThunderClap:           3,
	AngerManagement:               true,
	DeepWounds:                    3,
	TwoHandedWeaponSpecialization: 5,
	Impale:                        2,
	DeathWish:                     true,
	SwordSpecialization:           5,
	ImprovedDisciplines:           2,
	BloodFrenzy:                   2,
	MortalStrike:                  true,

	Cruelty:                   5,
	ImprovedDemoralizingShout: 5,
	CommandingPresence:        5,
	ImprovedSlam:              2,
	SweepingStrikes:           true,
	WeaponMastery:             2,
	Flurry:                    3,
}

var FuryTalents = &proto.WarriorTalents{
	ImprovedHeroicStrike: 3,
	AngerManagement:      true,
	DeepWounds:           3,
	Impale:               2,

	Cruelty:                 5,
	UnbridledWrath:          5,
	CommandingPresence:      5,
	DualWieldSpecialization: 5,
	SweepingStrikes:         true,
	WeaponMastery:           2,
	Flurry:                  5,
	Precision:               3,
	Bloodthirst:             true,
	ImprovedWhirlwind:       1,
	ImprovedBerserkerStance: 5,
	Rampage:                 true,
}

var armsSlamRotation = &proto.Warrior_Rotation{
	UseOverpower: true,
	UseHamstring: true,
	UseSlam:      true,

	HsRageThreshold:        70,
	HamstringRageThreshold: 75,
	OverpowerRageThreshold: 20,
	SlamLatency:            100,
	SlamGcdDelay:           400,
	SlamMsWwDelay:          2000,

	UseSlamDuringExecute: true,
	UseWwDuringExecute:   true,
	UseMsDuringExecute:   true,
	UseHsDuringExecute:   true,

	MaintainDemoShout:   true,
	MaintainThunderClap: true,
}

var warriorRotation = &proto.Warrior_Rotation{
	UseOverpower: true,
	UseHamstring: true,

	HsRageThreshold:        70,
	HamstringRageThreshold: 75,
	OverpowerRageThreshold: 20,
	RampageCdThreshold:     5,

	UseHsDuringExecute: true,
	UseWwDuringExecute: true,
	UseBtDuringExecute: true,
}

var warriorOptions = &proto.Warrior_Options{
	StartingRage:         50,
	UseRecklessness:      true,
	Shout:                proto.WarriorShout_WarriorShoutBattle,
	PrecastShout:         false,
	PrecastShoutT2:       false,
	PrecastShoutSapphire: false,
}

var FullRaidBuffs = &proto.RaidBuffs{
	ArcaneBrilliance: true,
	GiftOfTheWild:    proto.TristateEffect_TristateEffectImproved,
}
var FullPartyBuffs = &proto.PartyBuffs{
	BattleShout:     proto.TristateEffect_TristateEffectImproved,
	LeaderOfThePack: proto.TristateEffect_TristateEffectImproved,
}
var FullIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings:  true,
	BlessingOfWisdom: proto.TristateEffect_TristateEffectImproved,
	BlessingOfMight:  proto.TristateEffect_TristateEffectImproved,
}

var FullConsumes = &proto.Consumes{
	Drums: proto.Drums_DrumsOfBattle,
}

var FullDebuffs = &proto.Debuffs{
	BloodFrenzy:               true,
	FaerieFire:                proto.TristateEffect_TristateEffectImproved,
	ImprovedSealOfTheCrusader: true,
	JudgementOfWisdom:         true,
	Misery:                    true,
}

var FuryP1Gear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{
	{
		Name:    "Warbringer Battle-Helm",
		Enchant: "Glyph of Ferocity",
		Gems: []string{
			"Relentless Earthstorm Diamond",
			"Smooth Dawnstone",
		},
	},
	{
		Name: "Choker of Vile Intent",
	},
	{
		Name:    "Warbringer Shoulderplates",
		Enchant: "Greater Inscription of Vengeance",
		Gems: []string{
			"Smooth Dawnstone",
			"Jagged Talasite",
		},
	},
	{
		Name: "Vengeance Wrap",
	},
	{
		Name:    "Warbringer Breastplate",
		Enchant: "Chest - Exceptional Stats",
		Gems: []string{
			"Smooth Dawnstone",
			"Smooth Dawnstone",
			"Smooth Dawnstone",
		},
	},
	{
		Name:    "Bladespire Warbands",
		Enchant: "Bracer - Brawn",
		Gems: []string{
			"Jagged Talasite",
			"Inscribed Noble Topaz",
		},
	},
	{
		Name:    "Gauntlets of Martial Perfection",
		Enchant: "Gloves - Major Strength",
		Gems: []string{
			"Jagged Talasite",
			"Smooth Dawnstone",
		},
	},
	{
		Name: "Girdle of the Endless Pit",
		Gems: []string{
			"Inscribed Noble Topaz",
			"Jagged Talasite",
		},
	},
	{
		Name:    "Skulker's Greaves",
		Enchant: "Nethercobra Leg Armor",
		Gems: []string{
			"Smooth Dawnstone",
			"Smooth Dawnstone",
			"Smooth Dawnstone",
		},
	},
	{
		Name:    "Ironstriders of Urgency",
		Enchant: "Enchant Boots - Cat's Swiftness",
		Gems: []string{
			"Inscribed Noble Topaz",
			"Smooth Dawnstone",
		},
	},
	{
		Name: "Ring of a Thousand Marks",
	},
	{
		Name: "Shapeshifter's Signet",
	},
	{
		Name: "Dragonspine Trophy",
	},
	{
		Name: "Bloodlust Brooch",
	},
	{
		Name:    "Dragonmaw",
		Enchant: "Weapon - Mongoose",
	},
	{
		Name:    "Spiteblade",
		Enchant: "Weapon - Mongoose",
	},
	{
		Name: "Mama's Insurance",
	},
})
