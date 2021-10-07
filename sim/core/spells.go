package core

import (
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
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
	DamageType stats.Stat
	Coeff      float64

	DotDmg float64
	DotDur time.Duration
}

const SpellCritRatingPerCritChance = 22.08

// Global Spells (things that any class could cast)
var spells = []Spell{
	{ID: MagicIDTLCLB, Name: "TLCLB", Coeff: 0.0, MinDmg: 694, MaxDmg: 807, Mana: 0, DamageType: stats.NatureSpellPower},
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
	Caster *Agent
	Tag    int32 // Allow any class to create an enum for what tags the cast needs.

	// Pre-hit Mutatable State
	CastTime time.Duration // time to cast the spell
	ManaCost float64

	BonusSpellPower     float64 // Bonus spell power for only this spell
	BonusHit            float64 // Direct % bonus... 0.1 == 10%
	BonusCrit           float64 // Direct % bonus... 0.1 == 10%
	CritDamageMultipier float64 // Multiplier to critical dmg

	// Actual spell to call to activate this spell.
	//  currently named after arnold's "come on, do it now"
	DoItNow func(sim *Simulation, agent Agent, cast *Cast)

	// Calculated Values, can be modified only in 'OnSpellHit'
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
	cast.CritDamageMultipier = 1.5
	cast.CastTime = sp.CastTime
	cast.DoItNow = DirectCast
	return cast
}

// Cast will actually cast and treat all casts as having no 'flight time'.
// This will activate any auras around casting, calculate hit/crit and add to sim metrics.
func DirectCast(sim *Simulation, agent Agent, cast *Cast) {
	character := agent.GetCharacter()

	if sim.Log != nil {
		sim.Log("(%d) Current Mana %0.0f, Cast Cost: %0.0f\n", character.ID, character.Stats[stats.Mana], cast.ManaCost)
	}
	if cast.ManaCost > 0 {
		character.Stats[stats.Mana] -= cast.ManaCost
		sim.Metrics.IndividualMetrics[character.ID].ManaSpent += cast.ManaCost
	}

	for _, id := range character.ActiveAuraIDs {
		if character.Auras[id].OnCastComplete != nil {
			character.Auras[id].OnCastComplete(sim, agent, cast)
		}
	}
	for _, id := range sim.ActiveAuraIDs {
		if sim.Auras[id].OnCastComplete != nil {
			sim.Auras[id].OnCastComplete(sim, agent, cast)
		}
	}

	hit := 0.83 + character.Stats[stats.SpellHit]/1260.0 + cast.BonusHit // 12.6 hit == 1% hit
	hit = math.Min(hit, 0.99)                                            // can't get away from the 1% miss

	dbgCast := cast.Spell.Name
	if sim.Log != nil {
		sim.Log("(%d) Completed Cast (%0.2f hit chance) (%s)\n", character.ID, hit, dbgCast)
	}
	if sim.Rando.Float64("cast hit") < hit {
		sp := character.Stats[stats.SpellPower] + character.Stats[cast.Spell.DamageType] + cast.BonusSpellPower
		baseDmg := (sim.Rando.Float64("cast dmg") * cast.Spell.DmgDiff)
		bonus := (sp * cast.Spell.Coeff)
		dmg := baseDmg + cast.Spell.MinDmg + bonus
		// if sim.Log != nil {
		// 	sim.Log("base dmg: %0.1f, bonus: %0.1f, total: %0.1f\n", baseDmg, bonus, dmg)
		// }
		if cast.DidDmg != 0 { // use the pre-set dmg
			dmg = cast.DidDmg
		}
		cast.DidHit = true

		crit := (character.Stats[stats.SpellCrit] / (SpellCritRatingPerCritChance * 100)) + cast.BonusCrit
		if sim.Rando.Float64("cast crit") < crit {
			cast.DidCrit = true
			dmg *= cast.CritDamageMultipier
			if sim.Log != nil {
				dbgCast += " crit"
			}
		} else if sim.Log != nil {
			dbgCast += " hit"
		}

		// Average Resistance (AR) = (Target's Resistance / (Caster's Level * 5)) * 0.75
		// P(x) = 50% - 250%*|x - AR| <- where X is %resisted
		// Using these stats:
		//    13.6% chance of
		//  FUTURE: handle boss resists for fights/classes that are actually impacted by that.
		resVal := sim.Rando.Float64("cast resist")
		if resVal < 0.18 { // 13% chance for 25% resist, 4% for 50%, 1% for 75%
			if sim.Log != nil {
				dbgCast += " (partial resist: "
			}
			if resVal < 0.01 {
				dmg *= .25
				if sim.Log != nil {
					dbgCast += "75%)"
				}
			} else if resVal < 0.05 {
				dmg *= .5
				if sim.Log != nil {
					dbgCast += "50%)"
				}
			} else {
				dmg *= .75
				if sim.Log != nil {
					dbgCast += "25%)"
				}
			}
		}
		cast.DidDmg = dmg
		// Apply any effects specific to this cast.
		if cast.Effect != nil {
			cast.Effect(sim, agent, cast)
		}
		// Apply any on spell hit effects.
		for _, id := range character.ActiveAuraIDs {
			if character.Auras[id].OnSpellHit != nil {
				character.Auras[id].OnSpellHit(sim, agent, cast)
			}
		}
		for _, id := range sim.ActiveAuraIDs {
			if sim.Auras[id].OnSpellHit != nil {
				sim.Auras[id].OnSpellHit(sim, agent, cast)
			}
		}
		agent.OnSpellHit(sim, cast)
		// if sim.Log != nil {
		// 	sim.Log("FINAL DMG: %0.1f\n", cast.DidDmg)
		// }
	} else {
		if sim.Log != nil {
			dbgCast += " miss"
		}
		cast.DidDmg = 0
		cast.DidCrit = false
		cast.DidHit = false
		for _, id := range character.ActiveAuraIDs {
			if character.Auras[id].OnSpellMiss != nil {
				character.Auras[id].OnSpellMiss(sim, agent, cast)
			}
		}
		for _, id := range sim.ActiveAuraIDs {
			if sim.Auras[id].OnSpellMiss != nil {
				sim.Auras[id].OnSpellMiss(sim, agent, cast)
			}
		}
	}

	if cast.Spell.Cooldown > 0 {
		character.SetCD(cast.Spell.ID, cast.Spell.Cooldown+sim.CurrentTime)
	}

	if sim.Log != nil {
		sim.Log("(%d) %s: %0.0f\n", character.ID, dbgCast, cast.DidDmg)
	}

	sim.Metrics.Casts = append(sim.Metrics.Casts, cast)
	sim.Metrics.TotalDamage += cast.DidDmg
	sim.Metrics.IndividualMetrics[character.ID].TotalDamage += cast.DidDmg
	// if sim.Log != nil {
	// 	sim.Log("Total Dmg: %0.1f\n", sim.Metrics.TotalDamage)
	// }
}
