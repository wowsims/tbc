package retribution

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var defaultRetTalents = &proto.PaladinTalents{
	Benediction:                   5,
	ImprovedSealOfTheCrusader:     3,
	ImprovedJudgement:             2,
	Conviction:                    5,
	SealOfCommand:                 true,
	Crusade:                       3,
	SanctityAura:                  true,
	TwoHandedWeaponSpecialization: 3,
	ImprovedSanctityAura:          2,
	Vengeance:                     5,
	SanctifiedJudgement:           2,
	SanctifiedSeals:               3,
	Fanaticism:                    5,
	CrusaderStrike:                true,
	Precision:                     3,
	DivineStrength:                5,
}

var defaultRetRotation = &proto.RetributionPaladin_Rotation{
	ConsecrationRank: proto.RetributionPaladin_Rotation_None,
	UseExorcism:      false,
}

var defaultRetOptions = &proto.RetributionPaladin_Options{
	Judgement:             proto.RetributionPaladin_Options_Crusader,
	CrusaderStrikeDelayMs: 1700,
	HasteLeewayMs:         100,
	DamageTakenPerSecond:  0,
}

var DefaultOptions = &proto.Player_RetributionPaladin{
	RetributionPaladin: &proto.RetributionPaladin{
		Talents:  defaultRetTalents,
		Options:  defaultRetOptions,
		Rotation: defaultRetRotation,
	},
}

var FullRaidBuffs = &proto.RaidBuffs{
	ArcaneBrilliance: true,
	GiftOfTheWild:    proto.TristateEffect_TristateEffectImproved,
	DivineSpirit:     proto.TristateEffect_TristateEffectImproved,
}

var FullPartyBuffs = &proto.PartyBuffs{
	Bloodlust:            1,
	Drums:                proto.Drums_DrumsOfBattle,
	BraidedEterniumChain: true,
	ManaSpringTotem:      proto.TristateEffect_TristateEffectRegular,
	StrengthOfEarthTotem: proto.StrengthOfEarthType_EnhancingTotems,
	WindfuryTotemRank:    5,
	BattleShout:          proto.TristateEffect_TristateEffectImproved,
	WindfuryTotemIwt:     2,
}

var FullIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings:     true,
	BlessingOfMight:     proto.TristateEffect_TristateEffectImproved,
	BlessingOfSalvation: true,
	UnleashedRage:       true,
}

var FullConsumes = &proto.Consumes{
	Flask:           proto.Flask_FlaskOfRelentlessAssault,
	DefaultPotion:   proto.Potions_HastePotion,
	DefaultConjured: proto.Conjured_ConjuredDarkRune,
	Food:            proto.Food_FoodRoastedClefthoof,
}

var NoDebuffTarget = &proto.Target{
	Debuffs: &proto.Debuffs{},
	Armor:   7700,
}

var FullDebuffs = &proto.Debuffs{
	JudgementOfWisdom:           true,
	Misery:                      true,
	CurseOfElements:             proto.TristateEffect_TristateEffectImproved,
	IsbUptime:                   1,
	BloodFrenzy:                 true,
	ExposeArmor:                 proto.TristateEffect_TristateEffectImproved,
	FaerieFire:                  proto.TristateEffect_TristateEffectImproved,
	CurseOfRecklessness:         true,
	HuntersMark:                 proto.TristateEffect_TristateEffectImproved,
	ExposeWeaknessUptime:        1,
	ExposeWeaknessHunterAgility: 800,
}

var FullDebuffTarget = &proto.Target{
	Debuffs: FullDebuffs,
	Armor:   7700,
}

var Phase4Gear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{
	{
		Name:    "Cursed Vision of Sargeras",
		Enchant: "Glyph of Ferocity",
		Gems: []string{
			"Relentless Earthstorm Diamond",
			"Bold Crimson Spinel",
		},
	},
	{
		Name: "Pendant of the Perilous",
	},
	{
		Name:    "Shoulderpads of the Stranger",
		Enchant: "Greater Inscription of Vengeance",
		Gems: []string{
			"Bold Crimson Spinel",
		},
	},
	{
		Name:    "Cloak of Fiends",
		Enchant: "Enchant Cloak - Greater Agility",
	},
	{
		Name:    "Midnight Chestguard",
		Enchant: "Chest - Exceptional Stats",
		Gems: []string{
			"Bold Crimson Spinel",
			"Sovereign Shadowsong Amethyst",
			"Inscribed Pyrestone",
		},
	},
	{
		Name:    "Bindings of Lightning Reflexes",
		Enchant: "Bracer - Brawn",
	},
	{
		Name:    "Gloves of the Searing Grip",
		Enchant: "Gloves - Major Strength",
	},
	{
		Name: "Belt of One-Hundred Deaths",
		Gems: []string{
			"Bold Crimson Spinel",
			"Sovereign Shadowsong Amethyst",
		},
	},
	{
		Name:    "Bow-stitched Leggings",
		Enchant: "Nethercobra Leg Armor",
		Gems: []string{
			"Bold Crimson Spinel",
			"Bold Crimson Spinel",
			"Bold Crimson Spinel",
		},
	},
	{
		Name:    "Shadowmaster's Boots",
		Enchant: "Enchant Boots - Dexterity",
		Gems: []string{
			"Bold Crimson Spinel",
			"Inscribed Pyrestone",
		},
	},
	{
		Name: "Shapeshifter's Signet",
	},
	{
		Name: "Band of Devastation",
	},
	{
		Name: "Dragonspine Trophy",
	},
	{
		Name: "Berserker's Call",
	},
	{
		Name:    "Torch of the Damned",
		Enchant: "Weapon - Mongoose",
	},
	{
		Name: "Libram of Avengement",
	},
})
