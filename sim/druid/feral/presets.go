package feral

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var StandardTalents = &proto.DruidTalents{
	Ferocity:                5,
	SharpenedClaws:          3,
	ShreddingAttacks:        2,
	PredatoryStrikes:        3,
	PrimalFury:              2,
	SavageFury:              2,
	FaerieFire:              true,
	HeartOfTheWild:          5,
	SurvivalOfTheFittest:    3,
	LeaderOfThePack:         true,
	ImprovedLeaderOfThePack: 2,
	PredatoryInstincts:      5,
	Mangle:                  true,
	Furor:                   5,
	Naturalist:              5,
	NaturalShapeshifter:     3,
	Intensity:               3,
	OmenOfClarity:           true,
}

var FullRaidBuffs = &proto.RaidBuffs{
	ArcaneBrilliance: true,
	GiftOfTheWild:    proto.TristateEffect_TristateEffectImproved,
}
var FullPartyBuffs = &proto.PartyBuffs{
	Bloodlust:                  1,
	Drums:                      proto.Drums_DrumsOfBattle,
	BattleShout:                proto.TristateEffect_TristateEffectImproved,
	LeaderOfThePack:            proto.TristateEffect_TristateEffectImproved,
	GraceOfAirTotem:            proto.TristateEffect_TristateEffectImproved,
	ManaSpringTotem:            proto.TristateEffect_TristateEffectRegular,
	BraidedEterniumChain:       true,
	StrengthOfEarthTotem:       proto.StrengthOfEarthType_EnhancingTotems,
	SnapshotBsSolarianSapphire: true,
	SanctityAura:               proto.TristateEffect_TristateEffectImproved,
}
var FullIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings: true,
	BlessingOfMight: proto.TristateEffect_TristateEffectImproved,
	UnleashedRage:   true,
}

var FullConsumes = &proto.Consumes{
	BattleElixir:     proto.BattleElixir_ElixirOfMajorAgility,
	Food:             proto.Food_FoodGrilledMudfish,
	DefaultPotion:    proto.Potions_HastePotion,
	MainHandImbue:    proto.WeaponImbue_WeaponImbueAdamantiteWeightstone,
	DefaultConjured:  proto.Conjured_ConjuredDarkRune,
	ScrollOfStrength: 5,
	ScrollOfAgility:  5,
}

var FullDebuffs = &proto.Debuffs{
	JudgementOfWisdom:           true,
	ImprovedSealOfTheCrusader:   true,
	BloodFrenzy:                 true,
	ExposeArmor:                 proto.TristateEffect_TristateEffectImproved,
	FaerieFire:                  proto.TristateEffect_TristateEffectImproved,
	SunderArmor:                 true,
	CurseOfRecklessness:         true,
	HuntersMark:                 proto.TristateEffect_TristateEffectImproved,
	ExposeWeaknessUptime:        1.0,
	ExposeWeaknessHunterAgility: 1000,
}

var FullDebuffTarget = &proto.Target{
	Debuffs: FullDebuffs,
	Armor:   7684,
}

var PlayerOptionsBiteweave = &proto.Player_FeralDruid{
	FeralDruid: &proto.FeralDruid{
		Talents: StandardTalents,
		Options: &proto.FeralDruid_Options{
			InnervateTarget: &proto.RaidTarget{TargetIndex: -1}, // no Innervate
			LatencyMs:       100,
		},
		Rotation: &proto.FeralDruid_Rotation{
			FinishingMove: proto.FeralDruid_Rotation_Rip,
			MangleTrick:   true,
			Biteweave:     true,
			MangleBot:     false,
			RipCp:         5,
			BiteCp:        5,
			RakeTrick:     false,
			Ripweave:      false,
		},
	},
}

var P1Gear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{
	{
		Name:    "Wolfshead Helm",
		Enchant: "Glyph of Ferocity",
	},
	{
		Name: "Braided Eternium Chain",
	},
	{
		Name:    "Mantle of Malorne",
		Enchant: "Might of the Scourge",
		Gems: []string{
			"Delicate Living Ruby",
			"Delicate Living Ruby",
		},
	},
	{
		Name: "Vengeance Wrap",
		Enchant: "Enchant Cloak - Greater Agility",
		Gems: []string{
			"Delicate Living Ruby",
		},
	},
	{
		Name:    "Breastplate of Malorne",
		Enchant: "Chest - Exceptional Stats",
		Gems: []string{
			"Delicate Living Ruby",
			"Delicate Living Ruby",
			"Delicate Living Ruby",
		},Ruby
	},
	{
		Name:    "Nightfall Wristguards",
		Enchant: "Bracer - Brawn",
	},
	{
		Name:    "Gloves of Dexterous Manipulation",
		Enchant: "Gloves - Major Agility",
		Gems: []string{
			"Delicate Living Ruby",
			"Delicate Living Ruby",
		},
	},
	{
		Name: "Girdle of Treachery",
		Gems: []string{
			"Delicate Living Ruby",
			"Delicate Living Ruby",
		},
	},
	{
		Name:    "Skulker's Greaves",
		Enchant: "Nethercobra Leg Armor",
		Gems: []string{
			"Delicate Living Ruby",
			"Delicate Living Ruby",
			"Delicate Living Ruby",
		},
	},
	{
		Name: "Edgewalker Longboots",
		Enchant: "Enchant Boots - Cat's Swiftness",
		Gems: []string{
			"Delicate Living Ruby",
			"Delicate Living Ruby",
		},
	},
	{
		Name:    "Ring of the Recalcitrant",
		Enchant: "Ring - Striking",
	},
	{
		Name:    "Shapeshifter's Signet",
		Enchant: "Ring - Striking",
	},
	{
		Name: "Dragonspine Trophy",
	},
	{
		Name: "Bloodlust Brooch",
	},
	{
		Name: "Everbloom Idol",
	},
	{
		Name:    "Gladiator's Maul",
		Enchant: "2H Weapon - Major Agility",
	},
})
