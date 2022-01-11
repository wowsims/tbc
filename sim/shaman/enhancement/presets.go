package enhancement

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

var BasicRaidBuffs = &proto.RaidBuffs{}
var BasicPartyBuffs = &proto.PartyBuffs{
	Bloodlust: 1,
}
var BasicIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings: true,
}

var StandardTalents = &proto.ShamanTalents{
	Convection:         2,
	Concussion:         5,
	CallOfFlame:        3,
	ElementalFocus:     true,
	Reverberation:      5,
	ImprovedFireTotems: 1,

	AncestralKnowledge:      5,
	ThunderingStrikes:       5,
	EnhancingTotems:         2,
	ShamanisticFocus:        true,
	Flurry:                  5,
	ImprovedWeaponTotems:    1,
	ElementalWeapons:        3,
	MentalQuickness:         3,
	WeaponMastery:           5,
	DualWieldSpecialization: 3,
	Stormstrike:             true,
	UnleashedRage:           5,
	ShamanisticRage:         true,
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
	Drums: proto.Drums_DrumsOfBattle,
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
	{Name: "Band of the Ranger-General"},
	{Name: "Bloodlust Brooch"},
	{Name: "Boots of Utter Darkness"},
	{Name: "Belt of One-Hundred Deaths"},
	{Name: "Cataclysm Chestplate"},
	{Name: "Cataclysm Gauntlets"},
	{Name: "Cataclysm Helm"},
	{Name: "Cataclysm Legplates"},
	{Name: "Dragonspine Trophy"},
	{Name: "Ring of Lethality"},
	{Name: "Shoulderpads of the Stranger"},
	{Name: "Totem of the Astral Winds"},
	{Name: "True-Aim Stalker Bands"},
	{Name: "Thalassian Wildercloak"},
	{Name: "Telonicus's Pendant of Mayhem"},

	{Name: "Talon of the Phoenix"},
	{Name: "Rod of the Sun King"},
})

var PreRaidGear = items.EquipmentSpecFromStrings([]items.ItemStringSpec{
	items.ItemStringSpec{
		Name: "Gladiator's Cleaver",
	},
})
