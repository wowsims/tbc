package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
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

// Totem Item IDs
const (
	TotemOfTheVoid           = 28248
	TotemOfStorms            = 23199
	TotemOfAncestralGuidance = 32330
)

// NewCastAction is how a shaman creates a new spell
//  FUTURE: Decide if we need separate functions for elemental and enhancement?
func NewCastAction(shaman *Shaman, sim *core.Simulation, sp *core.Spell) core.AgentAction {
	cast := core.NewCast(sim, shaman, sp)

	itsElectric := sp.ActionID.SpellID == SpellIDCL6 || sp.ActionID.SpellID == SpellIDLB12

	if shaman.Talents.ElementalPrecision > 0 {
		// FUTURE: This only impacts "frost fire and nature" spells.
		//  We know it doesnt impact TLC.
		//  Are there any other spells that a shaman can cast?
		cast.BonusHit += float64(shaman.Talents.ElementalPrecision) * 0.02
	}
	if shaman.Talents.NaturesGuidance > 0 {
		cast.BonusHit += float64(shaman.Talents.NaturesGuidance) * 0.01
	}
	if shaman.Talents.TidalMastery > 0 {
		cast.BonusCrit += float64(shaman.Talents.TidalMastery) * 0.01
	}

	if itsElectric {
		// TODO: Should we change these to be full auras?
		//   Doesnt seem needed since they can only be used by shaman right here.
		if shaman.Equip[items.ItemSlotRanged].ID == TotemOfTheVoid {
			cast.BonusSpellPower += 55
		} else if shaman.Equip[items.ItemSlotRanged].ID == TotemOfStorms {
			cast.BonusSpellPower += 33
		} else if shaman.Equip[items.ItemSlotRanged].ID == TotemOfAncestralGuidance {
			cast.BonusSpellPower += 85
		}
		if shaman.Talents.CallOfThunder > 0 { // only applies to CL and LB
			cast.BonusCrit += float64(shaman.Talents.CallOfThunder) * 0.01
		}
		if sp.ActionID.SpellID == SpellIDCL6 && sim.Options.Encounter.NumTargets > 1 {
			cast.DoItNow = ChainCastHandler(shaman)
		}
		if shaman.Talents.LightningMastery > 0 {
			cast.CastTime -= time.Millisecond * 100 * time.Duration(shaman.Talents.LightningMastery)
		}
	}
	cast.CastTime = time.Duration(float64(cast.CastTime) / shaman.HasteBonus())

	// Apply any on cast effects.
	for _, id := range shaman.ActiveAuraIDs {
		if shaman.Auras[id].OnCast != nil {
			shaman.Auras[id].OnCast(sim, cast)
		}
	}
	if itsElectric { // TODO: Add ElementalFury talent
		// This is written this way to deal with making CSD dmg increase correct.
		// The 'OnCast' auras include CSD
		cast.CritDamageMultipier *= 2 // This handles the 'Elemental Fury' talent which increases the crit bonus.
		cast.CritDamageMultipier -= 1 // reduce to multiplier instead of percent.

		// Convection applies against the base cost of the spell.
		cast.ManaCost -= sp.Mana * shaman.convectionBonus
	}

	return core.AgentAction{
		Wait: 0,
		Cast: cast,
	}
}

// ChainCast is how to cast chain lightning.
func ChainCastHandler(shaman *Shaman) func(sim *core.Simulation, cast *core.Cast) {
	return func(sim *core.Simulation, cast *core.Cast) {
		core.DirectCast(sim, cast) // Start with a normal direct cast to start.

		// Now chain
		dmgCoeff := 1.0
		if cast.Tag == CastTagLightningOverload {
			dmgCoeff = 0.5 // since we are overriding the 'Effect' from LO, we will do it here.
		}
		for i := 1; i < sim.Options.Encounter.NumTargets; i++ {
			if shaman.HasAura(core.MagicIDTidefury) {
				dmgCoeff *= 0.83
			} else {
				dmgCoeff *= 0.7
			}
			cloneCoeff := dmgCoeff
			clone := &core.Cast{
				Caster:              cast.Caster,
				Tag:                 cast.Tag, // pass along lightning overload
				Spell:               cast.Spell,
				BonusCrit:           cast.BonusCrit,
				BonusHit:            cast.BonusHit,
				BonusSpellPower:     cast.BonusSpellPower,
				CritDamageMultipier: cast.CritDamageMultipier,
				Effect:              func(sim *core.Simulation, cast *core.Cast) { cast.DidDmg *= cloneCoeff },
				DoItNow:             ChainCastHandler(shaman), // so that LO will call ChainCast instead of DirectCast.
			}
			core.DirectCast(sim, clone)
		}
	}
}
