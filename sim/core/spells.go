package core

import (
	"math"
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
	DamageType Stat
	Coeff      float64

	CastType
	DotDmg float64
	DotDur time.Duration
}

type CastType byte

const (
	CastTypeNormal CastType = iota
	CastTypeChain
	CastTypeAOE
	CastTypeChannel
)

// All Spells
// FUTURE: Downrank Penalty == (spellrankavailbetobetrained+11)/70
//    Might not even be worth calculating because I don't think there is much case for ever downranking.
var spells = []Spell{
	// {ID: MagicIDLB4,  Name: "LB4",  Coeff: 0.795,  CastTime: time.Millisecond * 2500, MinDmg: 88, MaxDmg: 100, Mana: 50, DamageType: StatNatureSpellPower},
	// {ID: MagicIDLB10, Name: "LB10", Coeff: 0.795,  CastTime: time.Millisecond * 2500, MinDmg: 428, MaxDmg: 477, Mana: 265, DamageType: StatNatureSpellPower},
	{ID: MagicIDLB12, Name: "LB12", Coeff: 0.794, CastTime: time.Millisecond * 2500, MinDmg: 571, MaxDmg: 652, Mana: 300, DamageType: StatNatureSpellPower},
	// {ID: MagicIDCL4,  Name: "CL4",  Coeff: 0.643,  CastTime: time.Millisecond * 2000, Cooldown: time.Second * 6, MinDmg: 505, MaxDmg: 564, Mana: 605, DamageType: StatNatureSpellPower},
	{ID: MagicIDCL6, Name: "CL6", Coeff: 0.651, CastTime: time.Millisecond * 2000, Cooldown: time.Second * 6, MinDmg: 734, MaxDmg: 838, Mana: 760, DamageType: StatNatureSpellPower},
	// {ID: MagicIDES8,  Name: "ES8",  Coeff: 0.3858, CastTime: time.Millisecond * 1500, Cooldown: time.Second * 6, MinDmg: 658, MaxDmg: 692, Mana: 535, DamageType: StatNatureSpellPower},
	// {ID: MagicIDFrS5, Name: "FrS5", Coeff: 0.3858, CastTime: time.Millisecond * 1500, Cooldown: time.Second * 6, MinDmg: 640, MaxDmg: 676, Mana: 525, DamageType: StatFrostSpellPower},
	// {ID: MagicIDFlS7, Name: "FlS7", Coeff: 0.15, CastTime: time.Millisecond * 1500, Cooldown: time.Second * 6, MinDmg: 377, MaxDmg: 420, Mana: 500, DotDmg: 100, DotDur: time.Second * 6, DamageType: StatFireSpellPower},
	{ID: MagicIDTLCLB, Name: "TLCLB", Coeff: 0.0, MinDmg: 694, MaxDmg: 807, Mana: 0, DamageType: StatNatureSpellPower},
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

	Dmg       float64 // Bonus Dmg for only this spell
	Hit       float64 // Direct % bonus... 0.1 == 10%
	Crit      float64 // Direct % bonus... 0.1 == 10%
	CritBonus float64 // Multiplier to critical dmg bonus.

	// Actual spell to call to activate this spell.
	//  currently named after arnold's "come on, do it now"
	DoItNow func(sim *Simulation, p PlayerAgent, cast *Cast)

	// Calculated Values
	DidHit  bool
	DidCrit bool
	DidDmg  float64

	Effect AuraEffect // effects applied ONLY to this cast.
}

// NewCast constructs a Cast from the current simulation and selected spell.
//  OnCast mechanics are applied at this time (anything that modifies the cast before its cast, usually just mana cost stuff)
func NewCast(sim *Simulation, sp *Spell) *Cast {
	cast := sim.cache.NewCast()
	cast.Spell = sp
	cast.ManaCost = float64(sp.Mana)
	cast.CritBonus = 1.5
	cast.CastTime = sp.CastTime
	cast.DoItNow = DirectCast
	return cast
}

// Cast will actually cast and treat all casts as having no 'flight time'.
// This will activate any auras around casting, calculate hit/crit and add to sim metrics.
func DirectCast(sim *Simulation, p PlayerAgent, cast *Cast) {
	if sim.Debug != nil {
		sim.Debug("(%d) Current Mana %0.0f, Cast Cost: %0.0f\n", p.ID, p.Stats[StatMana], cast.ManaCost)
	}
	if cast.ManaCost > 0 {
		p.Stats[StatMana] -= cast.ManaCost
		sim.Metrics.IndividualMetrics[p.ID].ManaSpent += cast.ManaCost
	}

	for _, id := range p.ActiveAuraIDs {
		if p.Auras[id].OnCastComplete != nil {
			p.Auras[id].OnCastComplete(sim, p, cast)
		}
	}
	for _, id := range sim.ActiveAuraIDs {
		if sim.Auras[id].OnCastComplete != nil {
			sim.Auras[id].OnCastComplete(sim, p, cast)
		}
	}

	hit := 0.83 + p.Stats[StatSpellHit]/1260.0 + cast.Hit // 12.6 hit == 1% hit
	hit = math.Min(hit, 0.99)                             // can't get away from the 1% miss

	dbgCast := cast.Spell.Name
	if sim.Debug != nil {
		sim.Debug("(%d) Completed Cast (%0.2f hit chance) (%s)\n", p.ID, hit, dbgCast)
	}
	if sim.Rando.Float64("cast hit") < hit {
		sp := p.Stats[StatSpellPower] + p.Stats[cast.Spell.DamageType] + cast.Dmg
		baseDmg := (sim.Rando.Float64("cast dmg") * cast.Spell.DmgDiff)
		bonus := (sp * cast.Spell.Coeff)
		dmg := baseDmg + cast.Spell.MinDmg + bonus
		// if sim.Debug != nil {
		// 	sim.Debug("base dmg: %0.1f, bonus: %0.1f, total: %0.1f\n", baseDmg, bonus, dmg)
		// }
		if cast.DidDmg != 0 { // use the pre-set dmg
			dmg = cast.DidDmg
		}
		cast.DidHit = true

		crit := (p.Stats[StatSpellCrit] / 2208.0) + cast.Crit // 22.08 crit == 1% crit
		if sim.Rando.Float64("cast crit") < crit {
			cast.DidCrit = true
			dmg *= cast.CritBonus
			if sim.Debug != nil {
				dbgCast += " crit"
			}
		} else if sim.Debug != nil {
			dbgCast += " hit"
		}

		// Average Resistance (AR) = (Target's Resistance / (Caster's Level * 5)) * 0.75
		// P(x) = 50% - 250%*|x - AR| <- where X is %resisted
		// Using these stats:
		//    13.6% chance of
		//  FUTURE: handle boss resists for fights/classes that are actually impacted by that.
		resVal := sim.Rando.Float64("cast resist")
		if resVal < 0.18 { // 13% chance for 25% resist, 4% for 50%, 1% for 75%
			if sim.Debug != nil {
				dbgCast += " (partial resist: "
			}
			if resVal < 0.01 {
				dmg *= .25
				if sim.Debug != nil {
					dbgCast += "75%)"
				}
			} else if resVal < 0.05 {
				dmg *= .5
				if sim.Debug != nil {
					dbgCast += "50%)"
				}
			} else {
				dmg *= .75
				if sim.Debug != nil {
					dbgCast += "25%)"
				}
			}
		}
		cast.DidDmg = dmg
		// Apply any effects specific to this cast.
		if cast.Effect != nil {
			cast.Effect(sim, p, cast)
		}
		// Apply any on spell hit effects.
		for _, id := range p.ActiveAuraIDs {
			if p.Auras[id].OnSpellHit != nil {
				p.Auras[id].OnSpellHit(sim, p, cast)
			}
		}
		for _, id := range sim.ActiveAuraIDs {
			if sim.Auras[id].OnSpellHit != nil {
				sim.Auras[id].OnSpellHit(sim, p, cast)
			}
		}
		p.OnSpellHit(sim, p, cast)
		// if sim.Debug != nil {
		// 	sim.Debug("FINAL DMG: %0.1f\n", cast.DidDmg)
		// }
	} else {
		if sim.Debug != nil {
			dbgCast += " miss"
		}
		cast.DidDmg = 0
		cast.DidCrit = false
		cast.DidHit = false
		for _, id := range p.ActiveAuraIDs {
			if p.Auras[id].OnSpellMiss != nil {
				p.Auras[id].OnSpellMiss(sim, p, cast)
			}
		}
		for _, id := range sim.ActiveAuraIDs {
			if sim.Auras[id].OnSpellMiss != nil {
				sim.Auras[id].OnSpellMiss(sim, p, cast)
			}
		}
	}

	if cast.Spell.Cooldown > 0 {
		p.SetCD(cast.Spell.ID, cast.Spell.Cooldown+sim.CurrentTime)
	}

	if sim.Debug != nil {
		sim.Debug("(%d) %s: %0.0f\n", p.ID, dbgCast, cast.DidDmg)
	}

	sim.Metrics.Casts = append(sim.Metrics.Casts, cast)
	sim.Metrics.TotalDamage += cast.DidDmg
	// if sim.Debug != nil {
	// 	sim.Debug("Total Dmg: %0.1f\n", sim.Metrics.TotalDamage)
	// }
}
