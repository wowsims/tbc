package sim

import (
	"log"
	"testing"

	"github.com/wowsims/tbc/sim/api"
)

var basicSpec = &api.PlayerOptions_ElementalShaman{
	ElementalShaman: &api.ElementalShaman{
		Agent: &api.ElementalShaman_Agent{
			Type: api.ElementalShaman_Agent_Adaptive,
		},
		Talents: &api.ShamanTalents{
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
		Options: &api.ElementalShaman_Options{
			WaterShield: true,
		},
	},
}

var basicConsumes = &api.Consumes{
	FlaskOfBlindingLight: true,
	BlackenedBasilisk:    true,
	BrilliantWizardOil:   true,
	SuperManaPotion:      true,
	DarkRune:             true,
}

var basicBuffs = &api.Buffs{
	ArcaneBrilliance: true,
	BlessingOfKings:  true,
	Bloodlust:        1,
	MoonkinAura:      api.TristateEffect_TristateEffectRegular,
	ManaSpringTotem:  api.TristateEffect_TristateEffectRegular,
	TotemOfWrath:     1,
	WrathOfAirTotem:  api.TristateEffect_TristateEffectRegular,
}

var p1Equip = &api.EquipmentSpec{
	Items: []*api.ItemSpec{
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
	req := &api.IndividualSimRequest{
		Player: &api.Player{
			Options: &api.PlayerOptions{
				Race:     api.Race_RaceTroll10,
				Spec:     basicSpec,
				Consumes: basicConsumes,
			},
			Equipment: p1Equip,
		},
		Buffs: basicBuffs,
		Encounter: &api.Encounter{
			Duration:   120,
			NumTargets: 1,
		},
		Iterations: 1,
		RandomSeed: 1,
		Debug:      false,
	}

	res := RunSimulation(req)
	log.Printf("LOGS:\n%s\n", res.Logs)

	// TODO: validate something that wont break if we change core logic.
}

func TestGearList(t *testing.T) {
	glr := &api.GearListRequest{Spec: api.Spec_SpecElementalShaman}
	res := GetGearList(glr)

	// Print first item
	log.Printf("GEAR: %#v", res.Items[0])
}

func TestComputeStat(t *testing.T) {
	req := &api.ComputeStatsRequest{
		Player: &api.Player{
			Options: &api.PlayerOptions{
				Race:     api.Race_RaceTroll10,
				Spec:     basicSpec,
				Consumes: basicConsumes,
			},
			Equipment: p1Equip,
		},
		Buffs: basicBuffs,
	}

	ComputeStats(req)
}
