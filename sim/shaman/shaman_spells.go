package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const (
	MagicIDLB12 int32 = 1 + iota
	MagicIDCL6
)

// All Spells
// FUTURE: Downrank Penalty == (spellrankavailbetobetrained+11)/70
//    Might not even be worth calculating because I don't think there is much case for ever downranking.
var spells = []core.Spell{
	{ID: MagicIDLB12, Name: "LB12", Coeff: 0.794, CastTime: time.Millisecond * 2500, MinDmg: 571, MaxDmg: 652, Mana: 300, DamageType: stats.NatureSpellPower},
	{ID: MagicIDCL6, Name: "CL6", Coeff: 0.651, CastTime: time.Millisecond * 2000, Cooldown: time.Second * 6, MinDmg: 734, MaxDmg: 838, Mana: 760, DamageType: stats.NatureSpellPower},
	// {ID: MagicIDES8,  Name: "ES8",  Coeff: 0.3858, CastTime: time.Millisecond * 1500, Cooldown: time.Second * 6, MinDmg: 658, MaxDmg: 692, Mana: 535, DamageType: stats.NatureSpellPower},
	// {ID: MagicIDFrS5, Name: "FrS5", Coeff: 0.3858, CastTime: time.Millisecond * 1500, Cooldown: time.Second * 6, MinDmg: 640, MaxDmg: 676, Mana: 525, DamageType: StatFrostSpellPower},
	// {ID: MagicIDFlS7, Name: "FlS7", Coeff: 0.15, CastTime: time.Millisecond * 1500, Cooldown: time.Second * 6, MinDmg: 377, MaxDmg: 420, Mana: 500, DotDmg: 100, DotDur: time.Second * 6, DamageType: StatFireSpellPower},
}

// Spell lookup map to make lookups faster.
var Spells = map[int32]*core.Spell{}

func init() {
	for _, sp := range spells {
		// Turns out to increase efficiency go 'range' will actually only allocate a single struct and mutate.
		// If we want to create a pointer we need to clone the struct.
		sp2 := sp
		sp2.DmgDiff = sp2.MaxDmg - sp2.MinDmg
		spp := &sp2
		Spells[sp.ID] = spp
	}
}
