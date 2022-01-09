package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDFrostShock int32 = 25464

var ShockCooldownID = core.NewCooldownID() // shared CD for all shocks

func (shaman *Shaman) newFrostShockTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseManaCost := 525.0
	baseCast := core.Cast{
		CritMultiplier: 1.5,
		SpellSchool:    stats.FrostSpellPower,
		Character:      &shaman.Character,
		BaseManaCost:   baseManaCost,
		ManaCost:       baseManaCost,
		Cooldown:       time.Second * 6,
		ActionID: core.ActionID{
			SpellID:    SpellIDFrostShock,
			CooldownID: ShockCooldownID,
		},
		Binary: true,
	}
	baseCast.ManaCost -= baseCast.BaseManaCost * float64(shaman.Talents.Convection) * 0.02
	baseCast.ManaCost -= baseCast.BaseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02
	baseCast.Cooldown -= time.Millisecond * 200 * time.Duration(shaman.Talents.Reverberation)

	// TODO: confirm this is how it reduces mana cost.
	if ItemSetSkyshatterHarness.CharacterHasSetBonus(&shaman.Character, 2) {
		baseCast.ManaCost -= baseCast.BaseManaCost * 0.1
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
		},
		DirectInput: core.DirectDamageInput{
			MinBaseDamage:    647,
			MaxBaseDamage:    683,
			SpellCoefficient: 0.386,
		},
	}

	if shaman.Talents.ElementalFury {
		baseCast.CritMultiplier = 2
	}

	effect.SpellEffect.DamageMultiplier *= 1 + 0.01*float64(shaman.Talents.Concussion)
	effect.SpellEffect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance

	if shaman.Equip[items.ItemSlotRanged].ID == TotemOfImpact {
		effect.SpellEffect.BonusSpellPower += 46
	}

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		SpellHitEffect: effect,
	})
}

func (shaman *Shaman) NewFrostShock(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	shock := &shaman.shockSpell
	shaman.frostShockTemplate.Apply(shock)
	if shaman.Focused {
		shaman.applyFocusedEffect(shock)
	}
	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	shock.Target = target
	shock.Init(sim)

	return shock
}

const SpellIDEarthShock int32 = 25454

func (shaman *Shaman) newEarthShockTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseManaCost := 535.0
	baseCast := core.Cast{
		CritMultiplier: 1.5,
		SpellSchool:    stats.NatureSpellPower,
		Character:      &shaman.Character,
		BaseManaCost:   baseManaCost,
		ManaCost:       baseManaCost,
		Cooldown:       time.Second * 6,
		ActionID: core.ActionID{
			SpellID:    SpellIDEarthShock,
			CooldownID: ShockCooldownID,
		},
		Binary: true,
	}
	baseCast.ManaCost -= baseCast.BaseManaCost * float64(shaman.Talents.Convection) * 0.02
	baseCast.ManaCost -= baseCast.BaseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02
	baseCast.Cooldown -= time.Millisecond * 200 * time.Duration(shaman.Talents.Reverberation)

	// TODO: confirm this is how it reduces mana cost.
	if ItemSetSkyshatterHarness.CharacterHasSetBonus(&shaman.Character, 2) {
		baseCast.ManaCost -= baseCast.BaseManaCost * 0.1
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
		},
		DirectInput: core.DirectDamageInput{
			MinBaseDamage:    661,
			MaxBaseDamage:    696,
			SpellCoefficient: 0.386,
		},
	}

	if shaman.Talents.ElementalFury {
		baseCast.CritMultiplier = 2
	}

	effect.SpellEffect.DamageMultiplier *= 1 + 0.01*float64(shaman.Talents.Concussion)
	effect.SpellEffect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance

	if shaman.Equip[items.ItemSlotRanged].ID == TotemOfImpact {
		effect.SpellEffect.BonusSpellPower += 46
	}

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		SpellHitEffect: effect,
	})
}

func (shaman *Shaman) NewEarthShock(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	shock := &shaman.shockSpell
	shaman.earthShockTemplate.Apply(shock)
	if shaman.Focused {
		shaman.applyFocusedEffect(shock)
	}
	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	shock.Target = target
	shock.Init(sim)

	return shock
}
