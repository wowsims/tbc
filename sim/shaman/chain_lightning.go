package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

const SpellIDCL6 int32 = 25442

var ChainLightningCooldownID = core.NewCooldownID()

// newChainLightningTemplate returns a cast generator for Chain Lightning with as many fields precomputed as possible.
func (shaman *Shaman) newChainLightningTemplate(sim *core.Simulation, isLightningOverload bool) core.SimpleSpellTemplate {
	spellTemplate := core.SimpleSpell{
		SpellCast: shaman.newElectricSpellCast(
			core.ActionID{
				SpellID:    SpellIDCL6,
				CooldownID: ChainLightningCooldownID,
			},
			760.0,
			time.Millisecond*2000,
			isLightningOverload),
	}
	if !isLightningOverload {
		spellTemplate.Cooldown = time.Second * 6
	}

	effect := shaman.newElectricSpellEffect(734, 838, 0.651, isLightningOverload)

	if !isLightningOverload && shaman.Talents.LightningOverload > 0 {
		lightningOverloadChance := float64(shaman.Talents.LightningOverload) * 0.04 / 3
		effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			if shaman.Talents.ElementalFocus && spellEffect.Crit {
				shaman.ElementalFocusStacks = 2
			}

			if sim.RandomFloat("CL Lightning Overload") > lightningOverloadChance {
				return
			}
			overloadAction := shaman.NewChainLightning(sim, spellEffect.Target, true)
			overloadAction.Cast(sim)
		}
	} else {
		effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			if shaman.Talents.ElementalFocus && spellEffect.Crit {
				shaman.ElementalFocusStacks = 2
			}
		}
	}

	hasTidefury := ItemSetTidefury.CharacterHasSetBonus(&shaman.Character, 2)
	numHits := core.MinInt32(3, sim.GetNumTargets())
	effects := make([]core.SpellHitEffect, 0, numHits)
	effects = append(effects, effect)
	for i := int32(1); i < numHits; i++ {
		bounceEffect := effects[i-1] // Makes a copy of the previous bounce
		if hasTidefury {
			bounceEffect.DamageMultiplier *= 0.83
		} else {
			bounceEffect.DamageMultiplier *= 0.7
		}

		effects = append(effects, bounceEffect)
	}
	spellTemplate.Effects = effects

	return core.NewSimpleSpellTemplate(spellTemplate)
}

func (shaman *Shaman) getFirstAvailableCLLOObjectIndex() int {
	for i, _ := range shaman.chainLightningSpellLOs {
		if !shaman.chainLightningSpellLOs[i].IsInUse() {
			return i
		}
	}
	panic("All chain lightning LO objects in use!")
}

func (shaman *Shaman) NewChainLightning(sim *core.Simulation, target *core.Target, isLightningOverload bool) *core.SimpleSpell {
	var cl *core.SimpleSpell

	// Initialize cast from precomputed template.
	if isLightningOverload {
		objIndex := shaman.getFirstAvailableCLLOObjectIndex()
		cl = &shaman.chainLightningSpellLOs[objIndex]
		shaman.chainLightningLOCastTemplates[objIndex].Apply(cl)
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
