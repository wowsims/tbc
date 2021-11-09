package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	googleProto "google.golang.org/protobuf/proto"
)

var basicSpec = &proto.PlayerOptions_ElementalShaman{
	ElementalShaman: &proto.ElementalShaman{
		Rotation: &proto.ElementalShaman_Rotation{
			Type: proto.ElementalShaman_Rotation_Adaptive,
		},
		Talents: &proto.ShamanTalents{
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
		Options: &proto.ElementalShaman_Options{
			WaterShield: true,
		},
	},
}

var basicConsumes = &proto.Consumes{
	FlaskOfBlindingLight: true,
	BlackenedBasilisk:    true,
	BrilliantWizardOil:   true,
	DefaultPotion:        proto.Potions_SuperManaPotion,
	DarkRune:             true,
}

var p1Equip = &proto.EquipmentSpec{
	Items: []*proto.ItemSpec{
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

func init() {
	go func() {
		runServer(true, ":3333", false, bufio.NewReader(bytes.NewBuffer([]byte{})))
	}()

	time.Sleep(time.Second) // hack so we have time for server to startup. Probably could repeatedly curl the endpoint until it responds.
}

// TestIndividualSim is just a smoke test to make sure the http server works as expected.
//   Don't modify this test unless the proto defintions change and this no longer compiles.
func TestIndividualSim(t *testing.T) {
	req := &proto.IndividualSimRequest{
		Player: &proto.Player{
			Options: &proto.PlayerOptions{
				Race:     proto.Race_RaceTroll10,
				Class:    proto.Class_ClassShaman,
				Spec:     basicSpec,
				Consumes: basicConsumes,
			},
			Equipment: p1Equip,
		},
		RaidBuffs: &proto.RaidBuffs{},
		PartyBuffs: &proto.PartyBuffs{},
		ApproximationBuffs: &proto.ApproximationBuffs{},
		Encounter: &proto.Encounter{
			Duration: 120,
			Targets: []*proto.Target{
				&proto.Target{},
			},
		},
		SimOptions: &proto.SimOptions{
			Iterations: 5000,
			RandomSeed: 1,
			Debug:      false,
		},
	}

	msgBytes, err := googleProto.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to encode request: %s", err.Error())
	}

	r, err := http.Post("http://localhost:3333/individualSim", "application/x-protobuf", bytes.NewReader(msgBytes))
	if err != nil {
		t.Fatalf("Failed to POST request: %s", err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Failed to read result body: %s", err.Error())
		return
	}

	isr := &proto.IndividualSimResult{}
	if err := googleProto.Unmarshal(body, isr); err != nil {
		t.Fatalf("Failed to parse request: %s", err.Error())
		return
	}

	log.Printf("RESULT: %#v", isr)
}

func TestCalcStatWeight(t *testing.T) {
	req := &proto.IndividualSimRequest{
		Player: &proto.Player{
			Options: &proto.PlayerOptions{
				Race:     proto.Race_RaceTroll10,
				Spec:     basicSpec,
				Consumes: basicConsumes,
			},
			Equipment: p1Equip,
		},
		RaidBuffs: &proto.RaidBuffs{},
		PartyBuffs: &proto.PartyBuffs{},
		ApproximationBuffs: &proto.ApproximationBuffs{},
		Encounter: &proto.Encounter{
			Duration: 120,
			Targets: []*proto.Target{
				&proto.Target{},
			},
		},
		SimOptions: &proto.SimOptions{
			Iterations: 5000,
			RandomSeed: 1,
			Debug:      false,
		},
	}

	msgBytes, err := googleProto.Marshal(&proto.StatWeightsRequest{
		Options:         req,
		StatsToWeigh:    []proto.Stat{proto.Stat_StatSpellPower, proto.Stat_StatSpellCrit, proto.Stat_StatSpellHit},
		EpReferenceStat: proto.Stat_StatSpellPower,
	})
	if err != nil {
		t.Fatalf("Failed to encode request: %s", err.Error())
	}

	r, err := http.Post("http://localhost:3333/statWeights", "application/x-protobuf", bytes.NewReader(msgBytes))
	if err != nil {
		t.Fatalf("Failed to POST request: %s", err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Failed to read result body: %s", err.Error())
		return
	}

	isr := &proto.StatWeightsResult{}
	if err := googleProto.Unmarshal(body, isr); err != nil {
		t.Fatalf("Failed to parse request: %s", err.Error())
		return
	}

	log.Printf("RESULT: %#v", isr)
}
