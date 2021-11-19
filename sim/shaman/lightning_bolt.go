package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
)

const SpellIDLB12 int32 = 25449

// newLightningBoltTemplate returns a cast generator for Lightning Bolt with as many fields precomputed as possible.
func (shaman *Shaman) newLightningBoltTemplate(sim *core.Simulation, isLightningOverload bool) core.SimpleSpellTemplate {
	baseManaCost := 300.0
	if shaman.Equip[items.ItemSlotRanged].ID == TotemOfThePulsingEarth {
		baseManaCost -= 27.0
	}

	spellTemplate := core.SimpleSpell{
		SpellCast: shaman.newElectricSpellCast(
			"Lightning Bolt",
			core.ActionID{
				SpellID: SpellIDLB12,
			},
			baseManaCost,
			time.Millisecond*2500,
			isLightningOverload),
		SpellHitEffect: shaman.newElectricSpellEffect(571, 652, 0.794, isLightningOverload),
	}

	if !isLightningOverload && shaman.Talents.LightningOverload > 0 {
		lightningOverloadChance := float64(shaman.Talents.LightningOverload) * 0.04
		spellTemplate.SpellHitEffect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			if shaman.Talents.ElementalFocus && spellEffect.Crit {
				shaman.ElementalFocusStacks = 2
			}

			if sim.RandomFloat("LB Lightning Overload") < lightningOverloadChance {
				overloadAction := shaman.NewLightningBolt(sim, spellEffect.Target, true)
				overloadAction.Act(sim)
			}
		}
	} else {
		spellTemplate.SpellHitEffect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			if shaman.Talents.ElementalFocus && spellEffect.Crit {
				shaman.ElementalFocusStacks = 2
			}
		}
	}

	return core.NewSimpleSpellTemplate(spellTemplate)
}

func (shaman *Shaman) NewLightningBolt(sim *core.Simulation, target *core.Target, isLightningOverload bool) *core.SimpleSpell {
	var lb *core.SimpleSpell

	// Initialize cast from precomputed template.
	if isLightningOverload {
		lb = &shaman.lightningBoltSpellLO
		shaman.lightningBoltLOCastTemplate.Apply(lb)
	} else {
		lb = &shaman.lightningBoltSpell
		shaman.lightningBoltCastTemplate.Apply(lb)
	}

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	lb.SpellHitEffect.Target = target
	shaman.applyElectricSpellCastInitModifiers(&lb.SpellCast)

	lb.Init(sim)

	return lb
}
