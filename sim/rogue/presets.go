package rogue

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var CombatTalents = &proto.RogueTalents{
	Malice:              5,
	Ruthlessness:        3,
	Murder:              2,
	RelentlessStrikes:   true,
	ImprovedExposeArmor: 2,
	Lethality:           5,
	VilePoisons:         2,

	ImprovedSinisterStrike:  2,
	ImprovedSliceAndDice:    3,
	Precision:               5,
	DualWieldSpecialization: 5,
	BladeFlurry:             true,
	SwordSpecialization:     5,
	WeaponExpertise:         2,
	Aggression:              3,
	Vitality:                2,
	AdrenalineRush:          true,
	CombatPotency:           5,
	SurpriseAttacks:         true,
}

var MutilateTalents = &proto.RogueTalents{
	Malice:              5,
	Ruthlessness:        3,
	Murder:              2,
	PuncturingWounds:    3,
	RelentlessStrikes:   true,
	ImprovedExposeArmor: 2,
	Lethality:           5,
	ImprovedPoisons:     5,
	ColdBlood:           true,
	QuickRecovery:       2,
	SealFate:            5,
	Vigor:               true,
	FindWeakness:        5,
	Mutilate:            true,

	ImprovedSinisterStrike:  2,
	ImprovedSliceAndDice:    3,
	Precision:               5,
	DualWieldSpecialization: 5,
}

var HemoTalents = &proto.RogueTalents{
	ImprovedSinisterStrike:  2,
	ImprovedSliceAndDice:    3,
	Precision:               5,
	DualWieldSpecialization: 5,
	BladeFlurry:             true,
	SwordSpecialization:     5,
	WeaponExpertise:         2,
	Aggression:              3,
	Vitality:                2,
	AdrenalineRush:          true,
	CombatPotency:           5,

	SerratedBlades: 3,
	Hemorrhage:     true,
}

var PlayerOptionsBasic = &proto.Player_Rogue{
	Rogue: &proto.Rogue{
		Talents:  CombatTalents,
		Options:  basicOptions,
		Rotation: basicRotation,
	},
}

var PlayerOptionsMutilate = &proto.Player_Rogue{
	Rogue: &proto.Rogue{
		Talents:  MutilateTalents,
		Options:  basicOptions,
		Rotation: basicRotation,
	},
}

var PlayerOptionsHemo = &proto.Player_Rogue{
	Rogue: &proto.Rogue{
		Talents:  HemoTalents,
		Options:  basicOptions,
		Rotation: basicRotation,
	},
}

var basicRotation = &proto.Rogue_Rotation{
	Builder:             proto.Rogue_Rotation_Auto,
	MaintainExposeArmor: true,
	UseRupture:          true,
	UseShiv:             true,

	MinComboPointsForDamageFinisher: 3,
}

var basicOptions = &proto.Rogue_Options{
	ArResetsTicks: true,
}

var FullRaidBuffs = &proto.RaidBuffs{
	GiftOfTheWild: proto.TristateEffect_TristateEffectImproved,
}
var FullPartyBuffs = &proto.PartyBuffs{
	Bloodlust: 1,
	Drums:     proto.Drums_DrumsOfBattle,

	BattleShout:       proto.TristateEffect_TristateEffectImproved,
	LeaderOfThePack:   proto.TristateEffect_TristateEffectImproved,
	GraceOfAirTotem:   proto.TristateEffect_TristateEffectRegular,
	WindfuryTotemRank: 5,
}
var FullIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings: true,
	BlessingOfMight: proto.TristateEffect_TristateEffectImproved,
}

var FullConsumes = &proto.Consumes{
	MainHandImbue:   proto.WeaponImbue_WeaponImbueRogueInstantPoison,
	OffHandImbue:    proto.WeaponImbue_WeaponImbueRogueDeadlyPoison,
	Flask:           proto.Flask_FlaskOfRelentlessAssault,
	DefaultPotion:   proto.Potions_HastePotion,
	DefaultConjured: proto.Conjured_ConjuredRogueThistleTea,
	SuperSapper:     true,
	FillerExplosive: proto.Explosive_ExplosiveGnomishFlameTurret,
}

var FullDebuffs = &proto.Debuffs{
	BloodFrenzy:               true,
	Mangle:                    true,
	SunderArmor:               true,
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
		Name:    "Netherblade Facemask",
		Enchant: "Glyph of Ferocity",
		Gems: []string{
			"Relentless Earthstorm Diamond",
			"Glinting Noble Topaz",
		},
	},
	{
		Name: "Choker of Vile Intent",
	},
	{
		Name:    "Wastewalker Shoulderpads",
		Enchant: "Greater Inscription of Vengeance",
		Gems: []string{
			"Glinting Noble Topaz",
			"Shifting Nightseye",
		},
	},
	{
		Name: "Drape of the Dark Reavers",
	},
	{
		Name:    "Netherblade Chestpiece",
		Enchant: "Chest - Exceptional Stats",
		Gems: []string{
			"Glinting Noble Topaz",
			"Glinting Noble Topaz",
			"Shifting Nightseye",
		},
	},
	{
		Name:    "Nightfall Wristguards",
		Enchant: "Bracer - Assault",
	},
	{
		Name:    "Wastewalker Gloves",
		Enchant: "Gloves - Major Agility",
		Gems: []string{
			"Glinting Noble Topaz",
			"Glinting Noble Topaz",
		},
	},
	{
		Name: "Girdle of the Deathdealer",
		Gems: []string{},
	},
	{
		Name:    "Skulker's Greaves",
		Enchant: "Nethercobra Leg Armor",
		Gems: []string{
			"Delicate Living Ruby",
			"Glinting Noble Topaz",
			"Glinting Noble Topaz",
		},
	},
	{
		Name:    "Edgewalker Longboots",
		Enchant: "Enchant Boots - Cat's Swiftness",
		Gems: []string{
			"Glinting Noble Topaz",
			"Glinting Noble Topaz",
		},
	},
	{
		Name: "Ring of a Thousand Marks",
	},
	{
		Name: "Garona's Signet Ring",
	},
	{
		Name: "Dragonspine Trophy",
	},
	{
		Name: "Bloodlust Brooch",
	},
	{
		Name:    "Spiteblade",
		Enchant: "Weapon - Mongoose",
	},
	{
		Name:    "Latro's Shifting Sword",
		Enchant: "Weapon - Mongoose",
	},
	{
		Name:    "Sunfury Bow of the Phoenix",
		Enchant: "Stabilized Eternium Scope",
	},
})
