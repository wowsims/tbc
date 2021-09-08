package core

import (
	"time"
)

// Spell represents a single castable spell. This is all the data needed to begin a cast.
type Spell struct {
	ID         int32
	Name       string
	CastTime   time.Duration
	Cooldown   time.Duration
	Mana       float64
	MinDmg     float64
	MaxDmg     float64
	DmgDiff    float64 // cached for faster dmg calculations
	DamageType DamageType
	Coeff      float64

	DotDmg float64
	DotDur time.Duration
}

// DamageType is currently unused.
type DamageType byte

const (
	DamageTypeUnknown DamageType = iota
	DamageTypeFire
	DamageTypeNature
	DamageTypeFrost

	// Who cares about these fake damage types.
	DamageTypeShadow
	DamageTypeHoly
	DamageTypeArcane
)

// All Spells
// TODO: Downrank Penalty == (spellrankavailbetobetrained+11)/70
//    Might not even be worth calculating because I don't think there is much case for ever downranking.
var spells = []Spell{
	// {ID: MagicIDLB4,  Name: "LB4",  Coeff: 0.795,  CastTime: time.Millisecond * 2500, MinDmg: 88, MaxDmg: 100, Mana: 50, DamageType: DamageTypeNature},
	// {ID: MagicIDLB10, Name: "LB10", Coeff: 0.795,  CastTime: time.Millisecond * 2500, MinDmg: 428, MaxDmg: 477, Mana: 265, DamageType: DamageTypeNature},
	{ID: MagicIDLB12, Name: "LB12", Coeff: 0.794, CastTime: time.Millisecond * 2500, MinDmg: 571, MaxDmg: 652, Mana: 300, DamageType: DamageTypeNature},
	// {ID: MagicIDCL4,  Name: "CL4",  Coeff: 0.643,  CastTime: time.Millisecond * 2000, Cooldown: time.Second * 6, MinDmg: 505, MaxDmg: 564, Mana: 605, DamageType: DamageTypeNature},
	{ID: MagicIDCL6, Name: "CL6", Coeff: 0.651, CastTime: time.Millisecond * 2000, Cooldown: time.Second * 6, MinDmg: 734, MaxDmg: 838, Mana: 760, DamageType: DamageTypeNature},
	// {ID: MagicIDES8,  Name: "ES8",  Coeff: 0.3858, CastTime: time.Millisecond * 1500, Cooldown: time.Second * 6, MinDmg: 658, MaxDmg: 692, Mana: 535, DamageType: DamageTypeNature},
	// {ID: MagicIDFrS5, Name: "FrS5", Coeff: 0.3858, CastTime: time.Millisecond * 1500, Cooldown: time.Second * 6, MinDmg: 640, MaxDmg: 676, Mana: 525, DamageType: DamageTypeFrost},
	// {ID: MagicIDFlS7, Name: "FlS7", Coeff: 0.15, CastTime: time.Millisecond * 1500, Cooldown: time.Second * 6, MinDmg: 377, MaxDmg: 420, Mana: 500, DotDmg: 100, DotDur: time.Second * 6, DamageType: DamageTypeFire},
	{ID: MagicIDTLCLB, Name: "TLCLB", Coeff: 0.0, MinDmg: 694, MaxDmg: 807, Mana: 0, DamageType: DamageTypeNature},
}

// Spell lookup map to make lookups faster.
var Spells = map[int32]*Spell{}

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

type Cast struct {
	Spell  *Spell
	Caster *Player

	// Probably should generalize this... "Tags" map?
	IsLO       bool // stupid hack
	IsClBounce bool // stupider hack

	// Pre-hit Mutatable State
	CastTime time.Duration // time to cast the spell
	ManaCost float64

	Hit       float64 // Direct % bonus... 0.1 == 10%
	Crit      float64 // Direct % bonus... 0.1 == 10%
	CritBonus float64 // Multiplier to critical dmg bonus.

	// Calculated Values
	DidHit  bool
	DidCrit bool
	DidDmg  float64
	CastAt  time.Duration // simulation time the spell cast

	Effect AuraEffect // effects applied ONLY to this cast.
}

// NewCast constructs a Cast from the current simulation and selected spell.
//  OnCast mechanics are applied at this time (anything that modifies the cast before its cast, usually just mana cost stuff)
func NewCast(sim *Simulation, player Player, sp *Spell) *Cast {
	cast := sim.cache.NewCast()
	cast.Spell = sp
	cast.ManaCost = float64(sp.Mana)
	cast.CritBonus = 1.5

	castTime := sp.CastTime

	if sp.ID == MagicIDLB12 || sp.ID == MagicIDCL6 {
		cast.ManaCost *= 1 - (0.02 * float64(sim.Options.Talents.Convection))
		// TODO: Add LightningMaster to talent list (this will never not be selected for an elemental shaman)
		castTime -= time.Millisecond * 500 // Talent Lightning Mastery
	}
	hasteBonus := 1 + (player.Stats[StatSpellHaste] / 1576)
	castTime = time.Duration(float64(castTime) / hasteBonus)
	cast.CastTime = castTime

	// Apply any on cast effects.
	for _, id := range player.activeAuraIDs {
		if player.auras[id].OnCast != nil {
			player.auras[id].OnCast(sim, cast)
		}
	}

	return cast
}
