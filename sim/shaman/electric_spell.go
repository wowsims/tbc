package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Totem Item IDs
const (
	TotemOfAncestralGuidance = 32330
	TotemOfStorms            = 23199
	TotemOfThePulsingEarth   = 29389
	TotemOfTheVoid           = 28248
)

const (
	CastTagLightningOverload int32 = 1 // This could be value or bitflag if we ended up needing multiple flags at the same time.
)

// Shared logic for Lightning Bolt and Chain Lightning
type ElectricSpell struct {
	Shaman *Shaman
	IsLightningOverload bool

	name         string
	baseManaCost float64
	baseCastTime time.Duration
}

func (spell ElectricSpell) GetActionID() core.ActionID {
	return core.ActionID{
		SpellID: SpellIDLB12,
	}
}

func (spell ElectricSpell) GetName() string {
	if spell.IsLightningOverload {
		return spell.name + " (LO)"
	} else {
		return spell.name
	}
}

func (spell ElectricSpell) GetTag() int32 {
	if spell.IsLightningOverload {
		return CastTagLightningOverload
	} else {
		return 0
	}
}

func (spell ElectricSpell) GetAgent() core.Agent {
	return spell.Shaman
}

func (spell ElectricSpell) GetBaseManaCost() float64 {
	return spell.baseManaCost
}

func (spell ElectricSpell) GetSpellSchool() stats.Stat {
	return stats.NatureSpellPower
}

func (spell ElectricSpell) ApplyCastInputModifiers(input *core.DirectCastInput) {
}

func (spell ElectricSpell) GetCastInput(sim *core.Simulation, cast core.DirectCastAction) core.DirectCastInput {
	input := core.DirectCastInput{
		ManaCost: spell.baseManaCost,
		CastTime: spell.baseCastTime,
		CritMultiplier: 1.5,
	}

	if spell.IsLightningOverload {
		input.CastTime = 0
		input.ManaCost = 0
		input.IgnoreCooldowns = true
		input.IgnoreManaCost = true
	} else if spell.Shaman.Talents.LightningMastery > 0 {
		// Convection applies against the base cost of the spell.
		input.ManaCost -= spell.GetBaseManaCost() * spell.Shaman.convectionBonus

		input.CastTime -= time.Millisecond * 100 * time.Duration(spell.Shaman.Talents.LightningMastery)
	}

	if spell.Shaman.ElementalFocusStacks > 0 {
		// TODO: This should subtract 40% of base cost
		input.ManaCost *= .6 // reduced by 40%
	}

	if spell.Shaman.Talents.ElementalFury {
		input.CritMultiplier = 2
	}

	return input
}

func (spell ElectricSpell) ApplyHitInputModifiers(hitInput *core.DirectCastDamageInput) {
	hitInput.DamageMultiplier *= spell.Shaman.concussionBonus
	if spell.IsLightningOverload {
		hitInput.DamageMultiplier *= 0.5
	}

	hitInput.BonusHit += float64(spell.Shaman.Talents.ElementalPrecision) * 0.02
	hitInput.BonusHit += float64(spell.Shaman.Talents.NaturesGuidance) * 0.01
	hitInput.BonusCrit += float64(spell.Shaman.Talents.TidalMastery) * 0.01
	hitInput.BonusCrit += float64(spell.Shaman.Talents.CallOfThunder) * 0.01

	if spell.Shaman.Equip[items.ItemSlotRanged].ID == TotemOfStorms {
		hitInput.BonusSpellPower += 33
	} else if spell.Shaman.Equip[items.ItemSlotRanged].ID == TotemOfTheVoid {
		hitInput.BonusSpellPower += 55
	} else if spell.Shaman.Equip[items.ItemSlotRanged].ID == TotemOfAncestralGuidance {
		hitInput.BonusSpellPower += 85
	}
}

func (spell ElectricSpell) OnCastComplete(sim *core.Simulation, cast core.DirectCastAction) {
	if spell.Shaman.Talents.ElementalFocus && !spell.IsLightningOverload && spell.Shaman.ElementalFocusStacks > 0 {
		spell.Shaman.ElementalFocusStacks--
	}
}
func (spell ElectricSpell) OnElectricSpellHit(sim *core.Simulation, cast core.DirectCastAction, result *core.DirectCastDamageResult) {
	if result.Crit {
		spell.Shaman.ElementalFocusStacks = 2
	}
}
func (spell ElectricSpell) OnSpellMiss(sim *core.Simulation, cast core.DirectCastAction) {
}
