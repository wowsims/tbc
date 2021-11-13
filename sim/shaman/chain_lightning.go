package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

const SpellIDCL6 int32 = 25442

var ChainLightningCooldownID = core.NewCooldownID()

// newChainLightningTemplate returns a cast generator for Chain Lightning with as many fields precomputed as possible.
func (shaman *Shaman) newChainLightningTemplate(sim *core.Simulation, isLightningOverload bool) core.MultiTargetDirectDamageSpellTemplate {
	spellTemplate := core.MultiTargetDirectDamageSpell{
		SpellCast: shaman.newElectricSpellCast(
			"Chain Lightning",
			core.ActionID{
				SpellID:    SpellIDCL6,
				CooldownID: ChainLightningCooldownID,
			},
			760.0,
			time.Millisecond*2000,
			isLightningOverload),
	}
	spellTemplate.Cooldown = time.Second * 6

	effect := shaman.newElectricSpellEffect(734, 838, 0.651, isLightningOverload)

	if !isLightningOverload && shaman.Talents.LightningOverload > 0 {
		lightningOverloadChance := float64(shaman.Talents.LightningOverload) * 0.04 / 3
		effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			if shaman.Talents.ElementalFocus && spellEffect.Crit {
				shaman.ElementalFocusStacks = 2
			}

			if sim.RandomFloat("CL Lightning Overload") < lightningOverloadChance {
				overloadAction := shaman.NewChainLightning(sim, spellEffect.Target, true)
				overloadAction.Act(sim)
			}
		}
	} else {
		effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			if shaman.Talents.ElementalFocus && spellEffect.Crit {
				shaman.ElementalFocusStacks = 2
			}
		}
	}

	numHits := core.MinInt32(3, sim.GetNumTargets())
	effects := make([]core.DirectDamageSpellEffect, 0, numHits)
	effects = append(effects, effect)
	for i := int32(1); i < numHits; i++ {
		bounceEffect := effects[i-1] // Makes a copy

		if shaman.HasAura(Tidefury2PcAuraID) {
			bounceEffect.DamageMultiplier *= 0.83
		} else {
			bounceEffect.DamageMultiplier *= 0.7
		}

		effects = append(effects, bounceEffect)
	}
	spellTemplate.Effects = effects

	return core.NewMultiTargetDirectDamageSpellTemplate(spellTemplate)
}

func (shaman *Shaman) NewChainLightning(sim *core.Simulation, target *core.Target, isLightningOverload bool) *core.MultiTargetDirectDamageSpell {
	var cl *core.MultiTargetDirectDamageSpell

	// Initialize cast from precomputed template.
	if isLightningOverload {
		cl = &shaman.chainLightningSpellLO
		shaman.chainLightningLOCastTemplate.Apply(cl)
	} else {
		cl = &shaman.chainLightningSpell
		shaman.chainLightningCastTemplate.Apply(cl)
	}

	// Set dynamic fields, i.e. the stuff we couldn't precompute.

	// Set the targets.
	cl.Effects[0].Target = target
	curTargetIndex := target.Index
	numHits := core.MinInt32(3, sim.GetNumTargets())
	for i := int32(1); i < numHits; i++ {
		// Pick targets by index, incrementally.
		newTargetIndex := (curTargetIndex + 1) % sim.GetNumTargets()
		cl.Effects[i].Target = sim.GetTarget(newTargetIndex)
		curTargetIndex = newTargetIndex
	}

	shaman.applyElectricSpellCastInitModifiers(&cl.SpellCast)

	cl.Init(sim)

	return cl
}
