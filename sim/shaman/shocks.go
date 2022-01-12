package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDEarthShock int32 = 25454
const SpellIDFlameShock int32 = 25457
const SpellIDFrostShock int32 = 25464

var ShockCooldownID = core.NewCooldownID() // shared CD for all shocks

// Shared logic for all shocks.
func (shaman *Shaman) newShockTemplateSpell(sim *core.Simulation, spellID int32, spellSchool stats.Stat, baseManaCost float64) core.SimpleSpell {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				CritMultiplier: 1.5,
				SpellSchool:    spellSchool,
				Character:      &shaman.Character,
				BaseManaCost:   baseManaCost,
				ManaCost:       baseManaCost,
				Cooldown:       time.Second * 6,
				ActionID: core.ActionID{
					SpellID:    spellID,
					CooldownID: ShockCooldownID,
				},
				Binary: true,
			},
		},
		SpellHitEffect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
			},
		},
	}

	spell.ManaCost -= spell.BaseManaCost * float64(shaman.Talents.Convection) * 0.02
	spell.ManaCost -= spell.BaseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02
	spell.Cooldown -= time.Millisecond * 200 * time.Duration(shaman.Talents.Reverberation)

	// TODO: confirm this is how it reduces mana cost.
	if ItemSetSkyshatterHarness.CharacterHasSetBonus(&shaman.Character, 2) {
		spell.ManaCost -= spell.BaseManaCost * 0.1
	}

	if shaman.Talents.ElementalFury {
		spell.CritMultiplier = 2
	}

	spell.SpellHitEffect.SpellEffect.DamageMultiplier *= 1 + 0.01*float64(shaman.Talents.Concussion)
	spell.SpellHitEffect.SpellEffect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance

	if shaman.Equip[items.ItemSlotRanged].ID == TotemOfRage {
		spell.SpellHitEffect.SpellEffect.BonusSpellPower += 30
	} else if shaman.Equip[items.ItemSlotRanged].ID == TotemOfImpact {
		spell.SpellHitEffect.SpellEffect.BonusSpellPower += 46
	}

	if shaman.Talents.ElementalFocus {
		spell.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
			if shaman.ElementalFocusStacks > 0 {
				shaman.ElementalFocusStacks--
			}
		}
		spell.SpellHitEffect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			if spellEffect.Crit {
				shaman.ElementalFocusStacks = 2
			}
		}
	}

	return spell
}

// Shared shock logic that is dynamic, i.e. can't be precomputed.
func (shaman *Shaman) applyShockInitModifiers(spellCast *core.SpellCast) {
	if shaman.ElementalFocusStacks > 0 {
		// Reduces mana cost by 40%
		spellCast.Cast.ManaCost -= spellCast.Cast.BaseManaCost * 0.4
	}
}

func (shaman *Shaman) newFrostShockTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := shaman.newShockTemplateSpell(sim, SpellIDFrostShock, stats.FrostSpellPower, 525.0)

	spell.SpellHitEffect.DirectInput = core.DirectDamageInput{
		MinBaseDamage:    647,
		MaxBaseDamage:    683,
		SpellCoefficient: 0.386,
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (shaman *Shaman) NewFrostShock(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	shock := &shaman.shockSpell
	shaman.frostShockTemplate.Apply(shock)
	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	shock.Target = target
	shaman.applyShockInitModifiers(&shock.SpellCast)
	shock.Init(sim)

	return shock
}

func (shaman *Shaman) newEarthShockTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := shaman.newShockTemplateSpell(sim, SpellIDEarthShock, stats.NatureSpellPower, 535.0)

	spell.SpellHitEffect.DirectInput = core.DirectDamageInput{
		MinBaseDamage:    661,
		MaxBaseDamage:    696,
		SpellCoefficient: 0.386,
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (shaman *Shaman) NewEarthShock(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	shock := &shaman.shockSpell
	shaman.earthShockTemplate.Apply(shock)
	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	shock.Target = target
	shaman.applyShockInitModifiers(&shock.SpellCast)
	shock.Init(sim)

	return shock
}
