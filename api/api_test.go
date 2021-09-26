package api

import (
	"log"
	"testing"

	"github.com/wowsims/tbc/api/genapi"
)

var basicSpec = &genapi.PlayerOptions_ElementalShaman{
	ElementalShaman: &genapi.ElementalShaman{
		Agent: &genapi.ElementalShaman_ElementalShamanAgent{
			Type: genapi.ElementalShaman_ElementalShamanAgent_Adaptive,
		},
		Talents: &genapi.ShamanTalents{
			// ElementalDevastation
			ElementalFury:      true,
			Convection:         5,
			Concussion:         5,
			ElementalFocus:     true,
			CallOfThunder:      5,
			UnrelentingStorm:   3,
			ElementalPrecision: 3,
			LightningMastery:   5,
			ElementalMastery:   true,
			LightningOverload:  5,
		},
		Options: &genapi.ElementalShaman_ElementalShamanOptions{
			WaterShield: true,
		},
	},
}

var basicConsumes = &genapi.Consumes{
	FlaskOfBlindingLight: true,
	BlackenedBasilisk:    true,
	BrilliantWizardOil:   true,
	SuperManaPotion:      true,
	DarkRune:             true,
}

var basicBuffs = &genapi.Buffs{
	ArcaneBrilliance: true,
	BlessingOfKings:  true,
	Bloodlust:        1,
	MoonkinAura:      genapi.TristateEffect_TristateEffectRegular,
	ManaSpringTotem:  genapi.TristateEffect_TristateEffectRegular,
	TotemOfWrath:     1,
	WrathOfAirTotem:  genapi.TristateEffect_TristateEffectRegular,
}

var p1Equip = &genapi.EquipmentSpec{
	Items: []*genapi.ItemSpec{
		{Id: 29035, Gems: []int32{34220, 24059}, Enchant: 29191},
		{Id: 28762},
		{Id: 29037, Gems: []int32{24059, 24059}, Enchant: 28909},
		{Id: 28766},
		{Id: 29519},
		{Id: 29521},
		{Id: 28780},
		{Id: 29520},
		{Id: 30541},
		{Id: 28810},
		{Id: 30667},
		{Id: 28753},
		{Id: 28785},
		{Id: 29370},
		{Id: 28248},
		{Id: 28770, Enchant: 22555},
		{Id: 29268},
	},
}

// TestIndividualSim is designed to test the conversion of proto objects to
//  internal objects. This should not be a comprehensive test of the internals of the simulator.
//  It might be worth adding more features to ensure they all convert properly though!
//  Perhaps instead of running a real sim we just test that the output objects from the conversion functions work properly.
func TestIndividualSim(t *testing.T) {
	req := &genapi.IndividualSimRequest{
		Player: &genapi.Player{
			Options: &genapi.PlayerOptions{
				Race:     genapi.Race_RaceTroll10,
				Spec:     basicSpec,
				Consumes: basicConsumes,
			},
			Equipment: p1Equip,
		},
		Buffs: basicBuffs,
		Encounter: &genapi.Encounter{
			Duration:   120,
			NumTargets: 1,
		},
		Iterations: 1,
		RandomSeed: 1,
		Debug:      false,
	}

	res := RunSimulation(req)
	log.Printf("LOGS:\n%s\n", res.Logs)

	// TODO: validate something that wont' break if we change core logic.
}

func TestGearList(t *testing.T) {
	glr := &genapi.GearListRequest{Spec: genapi.Spec_SpecElementalShaman}
	res := GetGearList(glr)

	// Print first item
	log.Printf("GEAR: %#v", res.Items[0])
}

func TestComputeStat(t *testing.T) {
	req := &genapi.ComputeStatsRequest{
		Player: &genapi.Player{
			Options: &genapi.PlayerOptions{
				Race:     genapi.Race_RaceTroll10,
				Spec:     basicSpec,
				Consumes: basicConsumes,
			},
			Equipment: p1Equip,
		},
		Buffs: basicBuffs,
	}

	ComputeStats(req)
}
