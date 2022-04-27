package warlock

import (
	"github.com/wowsims/tbc/sim/core/proto"
)

var defaultDestroTalents = &proto.WarlockTalents{
	// destro
	ImprovedShadowBolt: 5,
	Bane:               5,
	Devastation:        5,
	Shadowburn:         true,
	DestructiveReach:   2,
	ImprovedImmolate:   5,
	Ruin:               true,
	Emberstorm:         5,
	Backlash:           3,
	Conflagrate:        true,
	ShadowAndFlame:     5,
	//demo
	DemonicEmbrace:     5,
	ImprovedVoidwalker: 1,
	FelIntellect:       3,
	// FelDomination: true,
	// MasterSummoner: 2,
	FelStamina:       3,
	DemonicAegis:     3,
	DemonicSacrifice: true,
}

var defaultDestroRotation = &proto.Warlock_Rotation{
	PrimarySpell: proto.Warlock_Rotation_Shadowbolt,
	Immolate:     true,
}

var defaultDestroOptions = &proto.Warlock_Options{
	Armor:           proto.Warlock_Options_FelArmor,
	Summon:          proto.Warlock_Options_Succubus,
	SacrificeSummon: true,
}

var DefaultDestroWarlock = &proto.Player_Warlock{
	Warlock: &proto.Warlock{
		Talents:  defaultDestroTalents,
		Options:  defaultDestroOptions,
		Rotation: defaultDestroRotation,
	},
}

var FullRaidBuffs = &proto.RaidBuffs{
	ArcaneBrilliance: true,
	GiftOfTheWild:    proto.TristateEffect_TristateEffectImproved,
	DivineSpirit:     proto.TristateEffect_TristateEffectImproved,
}

var FullPartyBuffs = &proto.PartyBuffs{
	Bloodlust:       1,
	Drums:           proto.Drums_DrumsOfBattle,
	ManaSpringTotem: proto.TristateEffect_TristateEffectRegular,
	WrathOfAirTotem: proto.TristateEffect_TristateEffectRegular,
	TotemOfWrath:    1,
}

var FullIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings:     true,
	BlessingOfSalvation: true,
}

var FullConsumes = &proto.Consumes{
	Flask:           proto.Flask_FlaskOfPureDeath,
	DefaultPotion:   proto.Potions_DestructionPotion,
	DefaultConjured: proto.Conjured_ConjuredDarkRune,
	Food:            proto.Food_FoodBlackenedBasilisk,
}

var NoDebuffTarget = &proto.Target{
	Debuffs: &proto.Debuffs{},
	Armor:   7700,
}

var FullDebuffs = &proto.Debuffs{
	JudgementOfWisdom:           true,
	Misery:                      true,
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

var Phase4Gear = &proto.EquipmentSpec{
	Items: []*proto.ItemSpec{
		{
			Id:      31051,
			Enchant: 29191,
			Gems: []int32{
				32218,
				34220,
			},
		},
		{
			Id: 33281,
		},
		{
			Id:      31054,
			Enchant: 28886,
			Gems: []int32{
				32215,
				32218,
			},
		},
		{
			Id:      32524,
			Enchant: 33150,
		},
		{
			Id:      30107,
			Enchant: 24003,
			Gems: []int32{
				32196,
				32196,
				32196,
			},
		},
		{
			Id:      32586,
			Enchant: 22534,
		},
		{
			Id:      31050,
			Enchant: 28272,
			Gems: []int32{
				32196,
			},
		},
		{
			Id: 30888,
			Gems: []int32{
				32196,
				32196,
			},
		},
		{
			Id:      31053,
			Enchant: 24274,
			Gems: []int32{
				32196,
			},
		},
		{
			Id:      32239,
			Enchant: 35297,
			Gems: []int32{
				32218,
				32215,
			},
		},
		{
			Id:      32527,
			Enchant: 22536,
		},
		{
			Id:      33497,
			Enchant: 22536,
		},
		{
			Id: 32483,
		},
		{
			Id: 33829,
		},
		{
			Id:      32374,
			Enchant: 22560,
		},
		{
			Id: 33192,
			Gems: []int32{
				32215,
			},
		},
	},
}
