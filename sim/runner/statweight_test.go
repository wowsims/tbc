package runner

import (
	"testing"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
	"github.com/wowsims/tbc/sim/shaman"
)

func TestCalcStatWeight(t *testing.T) {
	options := basicOptions
	options.Iterations = 5000
	options.Encounter = shortEncounter

	params := IndividualParams{
		Equip:       gearFromStrings(p1Gear),
		Race:        core.RaceBonusTypeTroll10,
		Consumes:    fullConsumes,
		Buffs:       fullBuffs,
		Options:     options,
		Spec:        shaman.ElementalSpec{Talents: shamTalents, Totems: shamTotems, AgentID: shaman.AgentTypeAdaptive},
		CustomStats: []float64{},
	}

	tests := []struct {
		name   string
		params IndividualParams
		want   StatWeightsResult
	}{
		{name: "First Test", params: params, want: StatWeightsResult{
			EpValues: []float64{stats.Intellect: 0.18, stats.SpellPower: 1, stats.SpellHit: 1.75, stats.SpellCrit: 0.72, stats.Armor: 0}, // armor at the end makes the array the right length....
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcStatWeight(tt.params, []stats.Stat{stats.SpellPower, stats.SpellHit, stats.Intellect, stats.SpellCrit}, stats.SpellPower); !floatArrayEqual(got.EpValues, tt.want.EpValues) {
				t.Errorf("CalcStatWeight() = %v, want %v", got.EpValues, tt.want.EpValues)
			}
		})
	}
}

func floatArrayEqual(got, want []float64) bool {
	if len(got) != len(want) {
		return false
	}
	const tolerance = 0.5
	for i := range got {
		if got[i] < want[i]-tolerance || got[i] > want[i]+tolerance {
			return false
		}
	}

	return true
}
