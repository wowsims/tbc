package protection

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var defaultProtTalents = &proto.PaladinTalents{
	Redoubt:                       5,
	Precision:                     3,
	Toughness:                     5,
	BlessingOfKings:               true,
	ImprovedRighteousFury:         3,
	Anticipation:                  5,
	BlessingOfSanctuary:           true,
	Reckoning:                     4,
	SacredDuty:                    2,
	OneHandedWeaponSpecialization: 5,
	HolyShield:                    true,
	ImprovedHolyShield:            2,
	CombatExpertise:               5,
	AvengersShield:                true,

	Benediction:       5,
	ImprovedJudgement: 2,
	Deflection:        5,
	PursuitOfJustice:  3,
	Crusade:           3,
}

var defaultProtRotation = &proto.ProtectionPaladin_Rotation{
	ConsecrationRank: proto.ProtectionPaladin_Rotation_None,
	UseExorcism:      false,
	UseHammerOfWrath: false,
}

var defaultProtOptions = &proto.ProtectionPaladin_Options{
	PrimaryJudgement: proto.ProtectionPaladin_Options_Righteousness,
	BuffJudgement:    proto.PaladinJudgement_JudgementOfWisdom,
}

var DefaultOptions = &proto.Player_ProtectionPaladin{
	ProtectionPaladin: &proto.ProtectionPaladin{
		Talents:  defaultProtTalents,
		Options:  defaultProtOptions,
		Rotation: defaultProtRotation,
	},
}

var FullRaidBuffs = &proto.RaidBuffs{
	ArcaneBrilliance:   true,
	GiftOfTheWild:      proto.TristateEffect_TristateEffectImproved,
	PowerWordFortitude: proto.TristateEffect_TristateEffectRegular,
}
var FullPartyBuffs = &proto.PartyBuffs{
	MoonkinAura:     proto.TristateEffect_TristateEffectRegular,
	TotemOfWrath:    1,
	WrathOfAirTotem: proto.TristateEffect_TristateEffectImproved,
	ManaSpringTotem: proto.TristateEffect_TristateEffectRegular,
}
var FullIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings: true,
	//BlessingOfSanctuary: true,
	BlessingOfWisdom: proto.TristateEffect_TristateEffectImproved,
	BlessingOfMight:  proto.TristateEffect_TristateEffectImproved,
}

var FullConsumes = &proto.Consumes{
	Flask:              proto.Flask_FlaskOfBlindingLight,
	Food:               proto.Food_FoodBlackenedBasilisk,
	DefaultPotion:      proto.Potions_SuperManaPotion,
	NumStartingPotions: 1,
	DefaultConjured:    proto.Conjured_ConjuredDarkRune,
	MainHandImbue:      proto.WeaponImbue_WeaponImbueSuperiorWizardOil,
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
		Name:    "Faceplate of the Impenetrable",
		Enchant: "Glyph of Power",
		Gems: []string{
			"Powerful Earthstorm Diamond",
			"Veiled Pyrestone",
		},
	},
	{
		Name: "The Darkener's Grasp",
	},
	{
		Name:    "Lightbringer Shoulderguards",
		Enchant: "Greater Inscription of the Knight",
		Gems: []string{
			"Glowing Shadowsong Amethyst",
			"Glowing Shadowsong Amethyst",
		},
	},
	{
		Name:    "Phoenix-Wing Cloak",
		Enchant: "Enchant Cloak - Dodge",
	},
	{
		Name:    "Lightbringer Chestguard",
		Enchant: "Chest - Exceptional Stats",
		Gems: []string{
			"Glowing Shadowsong Amethyst",
			"Veiled Pyrestone",
			"Veiled Pyrestone",
		},
	},
	{
		Name:    "The Seeker's Wristguards",
		Enchant: "Bracer - Spellpower",
	},
	{
		Name:    "Lightbringer Handguards",
		Enchant: "Gloves - Threat",
		Gems: []string{
			"Glowing Shadowsong Amethyst",
		},
	},
	{
		Name: "Girdle of the Protector",
		Gems: []string{
			"Glowing Shadowsong Amethyst",
			"Veiled Pyrestone",
		},
	},
	{
		Name:    "Lightbringer Legguards",
		Enchant: "Runic Spellthread",
		Gems: []string{
			"Glowing Shadowsong Amethyst",
		},
	},
	{
		Name:    "Sabatons of the Righteous Defender",
		Enchant: "Enchant Boots - Fortitude",
		Gems: []string{
			"Glowing Shadowsong Amethyst",
			"Glowing Shadowsong Amethyst",
		},
	},
	{
		Name:    "Band of the Eternal Sage",
		Enchant: "Ring - Spellpower",
	},
	{
		Name:    "Ashyen's Gift",
		Enchant: "Ring - Spellpower",
	},
	{
		Name: "Hex Shrunken Head",
	},
	{
		Name: "Dark Iron Smoking Pipe",
	},
	{
		Name:    "Tempest of Chaos",
		Enchant: "Weapon - Major Spellpower",
	},
	{
		Name:    "Bulwark of Azzinoth",
		Enchant: "Shield - Major Stamina",
	},
	{
		Name: "Libram of Divine Purpose",
	},
})
