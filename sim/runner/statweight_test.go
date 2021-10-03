package runner

import (
	"testing"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
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
		PlayerOptions: &playerOptionsAdaptive,
		CustomStats: stats.Stats{},
	}

	tests := []struct {
		name   string
		params IndividualParams
		want   StatWeightsResult
	}{
		{name: "First Test", params: params, want: StatWeightsResult{
			EpValues: stats.Stats{stats.Intellect: 0.23, stats.SpellPower: 1, stats.SpellHit: 1.90, stats.SpellCrit: 0.65},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcStatWeight(tt.params, []stats.Stat{stats.SpellPower, stats.SpellHit, stats.Intellect, stats.SpellCrit}, stats.SpellPower); !statsEqual(got.EpValues, tt.want.EpValues) {
				t.Errorf("CalcStatWeight() = %v, want %v", got.EpValues, tt.want.EpValues)
			}
		})
	}
}

func statsEqual(got stats.Stats, want stats.Stats) bool {
	const tolerance = 0.05
	for i := range got {
		if got[i] < want[i]-tolerance || got[i] > want[i]+tolerance {
			return false
		}
	}

	return true
}
