package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const (
	SpellIDLB12 int32 = 25449
	SpellIDCL6  int32 = 25442
)

// Shaman Spells
var spells = []core.Spell{
	{ActionID: core.ActionID{SpellID: SpellIDLB12}, Name: "LB12", Coeff: 0.794, CastTime: time.Millisecond * 2500, MinDmg: 571, MaxDmg: 652, Mana: 300, DamageType: stats.NatureSpellPower},

	// CooldownID is used for any spell that needs to have a CD. This ID should be added manually to core/auras.go for now.
	{ActionID: core.ActionID{SpellID: SpellIDCL6, CooldownID: core.MagicIDChainLightning6}, Name: "CL6", Coeff: 0.651, CastTime: time.Millisecond * 2000, Cooldown: time.Second * 6, MinDmg: 734, MaxDmg: 838, Mana: 760, DamageType: stats.NatureSpellPower},

	// {ID: MagicIDES8,  Name: "ES8",  Coeff: 0.3858, CastTime: time.Millisecond * 1500, Cooldown: time.Second * 6, MinDmg: 658, MaxDmg: 692, Mana: 535, DamageType: stats.NatureSpellPower},
	// {ID: MagicIDFrS5, Name: "FrS5", Coeff: 0.3858, CastTime: time.Millisecond * 1500, Cooldown: time.Second * 6, MinDmg: 640, MaxDmg: 676, Mana: 525, DamageType: StatFrostSpellPower},
	// {ID: MagicIDFlS7, Name: "FlS7", Coeff: 0.15, CastTime: time.Millisecond * 1500, Cooldown: time.Second * 6, MinDmg: 377, MaxDmg: 420, Mana: 500, DotDmg: 100, DotDur: time.Second * 6, DamageType: StatFireSpellPower},
}

// Spell lookup map to make lookups faster.
var Spells = map[int32]*core.Spell{}

func init() {
	for idx, sp := range spells {
		spells[idx].DmgDiff = sp.MaxDmg - sp.MinDmg
		// safe because we know no conflicting IDs in shaman package.
		Spells[sp.ActionID.SpellID] = &spells[idx]
	}
}
