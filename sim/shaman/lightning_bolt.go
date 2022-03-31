package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
)

const SpellIDLB12 int32 = 25449

// newLightningBoltTemplate returns a cast generator for Lightning Bolt with as many fields precomputed as possible.
func (shaman *Shaman) newLightningBoltSpell(sim *core.Simulation, isLightningOverload bool) *core.SimpleSpellTemplate {
	baseManaCost := 300.0
	if shaman.Equip[items.ItemSlotRanged].ID == TotemOfThePulsingEarth {
		baseManaCost -= 27.0
	}

	spellTemplate := core.SimpleSpell{
		SpellCast: shaman.newElectricSpellCast(
			core.ActionID{
				SpellID: SpellIDLB12,
			},
			baseManaCost,
			time.Millisecond*2500,
			isLightningOverload),
		Effect: shaman.newElectricSpellEffect(571, 652, 0.794, isLightningOverload),
	}

	if !isLightningOverload && shaman.Talents.LightningOverload > 0 {
		lightningOverloadChance := float64(shaman.Talents.LightningOverload) * 0.04
		spellTemplate.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() {
				return
			}
			if shaman.Talents.ElementalFocus && spellEffect.Outcome.Matches(core.OutcomeCrit) {
				shaman.ElementalFocusStacks = 2
			}

			if sim.RandomFloat("LB Lightning Overload") > lightningOverloadChance {
				return
			}
			shaman.LightningBoltLO.Cast(sim, spellEffect.Target)
		}
	} else {
		spellTemplate.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			if shaman.Talents.ElementalFocus && spellEffect.Outcome.Matches(core.OutcomeCrit) {
				shaman.ElementalFocusStacks = 2
			}
		}
	}

	if ItemSetSkyshatterRegalia.CharacterHasSetBonus(&shaman.Character, 4) {
		spellTemplate.Effect.DamageMultiplier *= 1.05
	}

	return shaman.RegisterSpell(core.SpellConfig{
		Template: spellTemplate,
		ModifyCast: func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
			instance.Effect.Target = target
			shaman.applyElectricSpellCastInitModifiers(&instance.SpellCast)
			if shaman.HasAura(NaturesSwiftnessAuraID) {
				instance.CastTime = 0
			}
		},
	})
}
