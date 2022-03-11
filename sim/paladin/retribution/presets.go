package retribution

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var BasicRaidBuffs = &proto.RaidBuffs{}
var BasicPartyBuffs = &proto.PartyBuffs{}
var BasicIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings: true,
}

// Need to create a standard Retribution talents
var StandardTalents = &proto.PaladinTalents{
	Benediction:               5,
	ImprovedSealOfTheCrusader: 3,
	ImprovedJudgement:         2,
	Conviction:                5,
	SealOfCommand:             true,
	Crusade:                   3,
	SanctityAura:              true,
	ImprovedSanctityAura:      2,
	Vengeance:                 5,
	SanctifiedJudgement:       2,
	SanctifiedSeals:           0,
	Fanaticism:                5,
	CrusaderStrike:            true,
	Precision:                 3,
	DivineStrength:            5,
}

var PlayerOptionsBasic = &proto.Player_RetributionPaladin{
	RetributionPaladin: &proto.RetributionPaladin{
		Talents:  StandardTalents,
		Options:  retPalOptions,
		Rotation: retPalRotation,
	},
}

var retPalRotation = &proto.RetributionPaladin_Rotation{}

var retPalOptions = &proto.RetributionPaladin_Options{}

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
	Flask:           proto.Flask_FlaskOfRelentlessAssault,
	DefaultPotion:   proto.Potions_HastePotion,
	SuperSapper:     true,
	GoblinSapper:    true,
	FillerExplosive: proto.Explosive_ExplosiveFelIronBomb,
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
		Name:    "Cataclysm's Edge",
		Enchant: "Weapon - Mongoose",
	},
	{
		Name: "Libram of Avengement",
	},
})
