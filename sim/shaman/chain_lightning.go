package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

const SpellIDCL6 int32 = 25442

var ChainLightningCooldownID = core.NewCooldownID()

func (shaman *Shaman) newChainLightningSpell(sim *core.Simulation, isLightningOverload bool) *core.SimpleSpellTemplate {
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

	makeOnSpellHit := func(hitIndex int32) core.OnSpellHit {
		if !isLightningOverload && shaman.Talents.LightningOverload > 0 {
			lightningOverloadChance := float64(shaman.Talents.LightningOverload) * 0.04 / 3
			return func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}
				if shaman.Talents.ElementalFocus && spellEffect.Outcome.Matches(core.OutcomeCrit) {
					shaman.ElementalFocusStacks = 2
				}

				if sim.RandomFloat("CL Lightning Overload") > lightningOverloadChance {
					return
				}
				if sim.Log != nil {
					sim.Log("LO #%d", hitIndex)
				}
				shaman.ChainLightningLOs[hitIndex].Cast(sim, spellEffect.Target)
			}
		} else {
			return func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if shaman.Talents.ElementalFocus && spellEffect.Outcome.Matches(core.OutcomeCrit) {
					shaman.ElementalFocusStacks = 2
				}
			}
		}
	}

	hasTidefury := ItemSetTidefury.CharacterHasSetBonus(&shaman.Character, 2)
	numHits := core.MinInt32(3, sim.GetNumTargets())
	effects := make([]core.SpellEffect, 0, numHits)

	effect.OnSpellHit = makeOnSpellHit(0)
	effects = append(effects, effect)

	for i := int32(1); i < numHits; i++ {
		bounceEffect := effects[i-1] // Makes a copy of the previous bounce
		if hasTidefury {
			bounceEffect.DamageMultiplier *= 0.83
		} else {
			bounceEffect.DamageMultiplier *= 0.7
		}
		bounceEffect.OnSpellHit = makeOnSpellHit(i)

		effects = append(effects, bounceEffect)
	}
	spellTemplate.Effects = effects

	return shaman.RegisterSpell(core.SpellConfig{
		Template: spellTemplate,
		ModifyCast: func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
			// Set the targets.
			instance.Effects[0].Target = target
			curTargetIndex := target.Index
			numHits := core.MinInt32(3, sim.GetNumTargets())
			for i := int32(1); i < numHits; i++ {
				// Pick targets by index, incrementally.
				newTargetIndex := (curTargetIndex + 1) % sim.GetNumTargets()
				instance.Effects[i].Target = sim.GetTarget(newTargetIndex)
				curTargetIndex = newTargetIndex
			}

			shaman.applyElectricSpellCastInitModifiers(&instance.SpellCast)
		},
	})
}
