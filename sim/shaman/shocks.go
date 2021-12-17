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
		Name:           "Frost Shock",
		CritMultiplier: 1.5,
		SpellSchool:    stats.FrostSpellPower,
		Character:      &shaman.Character,
		BaseManaCost:   baseManaCost,
		ManaCost:       baseManaCost,
		Cooldown:       time.Second * 8,
		ActionID: core.ActionID{
			SpellID:    SpellIDFrostShock,
			CooldownID: ShockCooldownID,
		},
		Binary: true,
	}
	baseCast.ManaCost -= baseCast.BaseManaCost * float64(shaman.Talents.Convection) * 0.02
	baseCast.ManaCost -= baseCast.BaseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02
	baseCast.Cooldown -= time.Millisecond * 200 * time.Duration(shaman.Talents.Reverberation)

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
