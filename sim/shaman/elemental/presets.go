package elemental

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var BasicBuffs = core.Buffs{
	Bloodlust: 1,
}

var StandardTalents = proto.ShamanTalents{
	ElementalFocus:     true,
	LightningMastery:   5,
	LightningOverload:  5,
	ElementalPrecision: 3,
	NaturesGuidance:    3,
	TidalMastery:       5,
	ElementalMastery:   true,
	ElementalFury:      true,
	UnrelentingStorm:   3,
	CallOfThunder:      5,
	Concussion:         5,
	Convection:         5,
}

var eleShamOptionsNoBuffs = &proto.ElementalShaman_Options{
	WaterShield: true,
	// Bloodlust:       true,
	// ManaSpringTotem: true,
	// TotemOfWrath:    true,
	// WrathOfAirTotem: true,
}
var PlayerOptionsAdaptiveNoBuffs = proto.PlayerOptions{
	Spec: &proto.PlayerOptions_ElementalShaman{
		ElementalShaman: &proto.ElementalShaman{
			Talents: &StandardTalents,
			Options: eleShamOptionsNoBuffs,
			Rotation: &proto.ElementalShaman_Rotation{
				Type: proto.ElementalShaman_Rotation_Adaptive,
			},
		},
	},
}

var eleShamOptions = &proto.ElementalShaman_Options{
	WaterShield:     true,
	Bloodlust:       true,
	ManaSpringTotem: true,
	TotemOfWrath:    true,
	WrathOfAirTotem: true,
}
var PlayerOptionsAdaptive = proto.PlayerOptions{
	Spec: &proto.PlayerOptions_ElementalShaman{
		ElementalShaman: &proto.ElementalShaman{
			Talents: &StandardTalents,
			Options: eleShamOptions,
			Rotation: &proto.ElementalShaman_Rotation{
				Type: proto.ElementalShaman_Rotation_Adaptive,
			},
		},
	},
}

var PlayerOptionsLBOnly = proto.PlayerOptions{
	Spec: &proto.PlayerOptions_ElementalShaman{
		ElementalShaman: &proto.ElementalShaman{
			Talents: &StandardTalents,
			Options: eleShamOptions,
			Rotation: &proto.ElementalShaman_Rotation{
				Type: proto.ElementalShaman_Rotation_LBOnly,
			},
		},
	},
}

var PlayerOptionsCLOnClearcast = proto.PlayerOptions{
	Spec: &proto.PlayerOptions_ElementalShaman{
		ElementalShaman: &proto.ElementalShaman{
			Talents: &StandardTalents,
			Options: eleShamOptions,
			Rotation: &proto.ElementalShaman_Rotation{
				Type: proto.ElementalShaman_Rotation_CLOnClearcast,
			},
		},
	},
}

var FullBuffs = core.Buffs{
	ArcaneBrilliance:  true,
	GiftOfTheWild:     proto.TristateEffect_TristateEffectImproved,
	BlessingOfKings:   true,
	BlessingOfWisdom:  proto.TristateEffect_TristateEffectImproved,
	JudgementOfWisdom: true,
	MoonkinAura:       proto.TristateEffect_TristateEffectRegular,
	ShadowPriestDPS:   500,
}

var FullConsumes = core.Consumes{
	FlaskOfBlindingLight: true,
	BrilliantWizardOil:   true,
	BlackenedBasilisk:    true,
	DefaultPotion:        proto.Potions_SuperManaPotion,
	StartingPotion:       proto.Potions_DestructionPotion,
	NumStartingPotions:   1,
	DarkRune:             true,
	Drums:                proto.Drums_DrumsOfBattle,
}

var PreRaidGear = items.EquipmentSpecFromStrings([]string{
	"Tidefury Helm",
	"Brooch of Heightened Potential",
	"Tidefury Shoulderguards",
	"Cloak of the Black Void",
	"Tidefury Chestpiece",
	"Shattrath Wraps",
	"Tidefury Gauntlets",
	"Moonrage Girdle",
	"Tidefury Kilt",
	"Earthbreaker's Greaves",
	"Seal of the Exorcist",
	"Spectral Band of Innervation",
	"Xi'ri's Gift",
	"Quagmirran's Eye",
	"Totem of the Void",
	"Sky Breaker",
	"Silvermoon Crest Shield",
})

var P1Gear = items.EquipmentSpecFromStrings([]string{
	"Cyclone Faceguard",
	"Adornment of Stolen Souls",
	"Cyclone Shoulderguards",
	"Ruby Drape of the Mysticant",
	"Netherstrike Breastplate",
	"Netherstrike Bracers",
	"Soul-Eater's Handwraps",
	"Netherstrike Belt",
	"Stormsong Kilt",
	"Windshear Boots",
	"Ring of Unrelenting Storms",
	"Ring of Recurrence",
	"The Lightning Capacitor",
	"Icon of the Silver Crescent",
	"Totem of the Void",
	"Nathrezim Mindblade",
	"Mazthoril Honor Shield",
})
