package enhancement

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var BasicRaidBuffs = &proto.RaidBuffs{}
var BasicPartyBuffs = &proto.PartyBuffs{
	Bloodlust: 1,
}
var BasicIndividualBuffs = &proto.IndividualBuffs{}

var StandardTalents = &proto.ShamanTalents{
	ThunderingStrikes:       5,
	EnhancingTotems:         2,
	Flurry:                  5,
	DualWieldSpecialization: 3,
	Stormstrike:             true,
	ElementalWeapons:        3,
	WeaponMastery:           5,
	UnleashedRage:           5,
	ShamanisticFocus:        true,
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
var FullPartyBuffs = &proto.PartyBuffs{}
var FullIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings:  true,
	BlessingOfWisdom: proto.TristateEffect_TristateEffectImproved,
	ShadowPriestDps:  500,
}

var FullConsumes = &proto.Consumes{
	Drums: proto.Drums_DrumsOfBattle,
}

var NoDebuffTarget = &proto.Target{
	Debuffs: &proto.Debuffs{},
	Armor:   6700,
}

var FullDebuffTarget = &proto.Target{
	Debuffs: &proto.Debuffs{
		ImprovedSealOfTheCrusader: true,
		JudgementOfWisdom:         true,
		Misery:                    true,
	},
}

var PreRaidGear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{
	items.ItemStringSpec{
		Name:    "Tidefury Helm",
		Enchant: "Glyph of Power",
		Gems: []string{
			"Runed Living Ruby",
			"Insightful Earthstorm Diamond",
		},
	},
	items.ItemStringSpec{
		Name:    "Brooch of Heightened Potential",
		Enchant: "Zandalar Signet of Mojo",
	},
	items.ItemStringSpec{
		Name: "Tidefury Shoulderguards",
	},
	items.ItemStringSpec{
		Name: "Cloak of the Black Void",
	},
	items.ItemStringSpec{
		Name: "Tidefury Chestpiece",
	},
	items.ItemStringSpec{
		Name: "Shattrath Wraps",
	},
	items.ItemStringSpec{
		Name: "Tidefury Gauntlets",
	},
	items.ItemStringSpec{
		Name: "Moonrage Girdle",
	},
	items.ItemStringSpec{
		Name:    "Tidefury Kilt",
		Enchant: "Mystic Spellthread",
	},
	items.ItemStringSpec{
		Name: "Earthbreaker's Greaves",
	},
	items.ItemStringSpec{
		Name: "Seal of the Exorcist",
	},
	items.ItemStringSpec{
		Name: "Spectral Band of Innervation",
	},
	items.ItemStringSpec{
		Name: "Xi'ri's Gift",
	},
	items.ItemStringSpec{
		Name: "Quagmirran's Eye",
	},
	items.ItemStringSpec{
		Name: "Totem of the Void",
	},
	items.ItemStringSpec{
		Name: "Gladiator's Cleaver",
	},
	items.ItemStringSpec{
		Name: "Gladiator's Cleaver",
	},
})
