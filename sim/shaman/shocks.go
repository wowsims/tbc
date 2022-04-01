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
var FlameShockDebuffID = core.NewDebuffID()

func (shaman *Shaman) ShockCD() time.Duration {
	return time.Second*6 - time.Millisecond*200*time.Duration(shaman.Talents.Reverberation)
}

// Shared logic for all shocks.
func (shaman *Shaman) newShockSpellConfig(sim *core.Simulation, spellID int32, spellSchool core.SpellSchool, baseManaCost float64) core.SpellConfig {
	cost := core.ResourceCost{Type: stats.Mana, Value: baseManaCost}
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID:    spellID,
					CooldownID: ShockCooldownID,
				},
				Character:   &shaman.Character,
				SpellSchool: spellSchool,
				BaseCost:    cost,
				Cost:        cost,
				GCD:         core.GCDDefault,
				Cooldown:    shaman.ShockCD(),
				SpellExtras: SpellFlagShock,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      shaman.DefaultSpellCritMultiplier(),
			DamageMultiplier:    1,
			ThreatMultiplier:    1,
		},
	}

	spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.Convection) * 0.02
	spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.MentalQuickness) * 0.02

	if shaman.Talents.ElementalFury {
		spell.Effect.CritMultiplier = shaman.SpellCritMultiplier(1, 1)
	}
	spell.Effect.ThreatMultiplier *= 1 - (0.1/3)*float64(shaman.Talents.ElementalPrecision)
	spell.Effect.DamageMultiplier *= 1 + 0.01*float64(shaman.Talents.Concussion)
	spell.Effect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance
	spell.Effect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance

	// TODO: confirm this is how it reduces mana cost.
	if ItemSetSkyshatterHarness.CharacterHasSetBonus(&shaman.Character, 2) {
		spell.Cost.Value -= spell.BaseCost.Value * 0.1
	}

	if shaman.Equip[items.ItemSlotRanged].ID == TotemOfRage {
		spell.Effect.BonusSpellPower += 30
	} else if shaman.Equip[items.ItemSlotRanged].ID == TotemOfImpact {
		spell.Effect.BonusSpellPower += 46
	}

	if shaman.Talents.ElementalFocus {
		spell.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
			if shaman.ElementalFocusStacks > 0 {
				shaman.ElementalFocusStacks--
			}
		}
		spell.Effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Outcome.Matches(core.OutcomeCrit) {
				shaman.ElementalFocusStacks = 2
			}
		}
	}

	return core.SpellConfig{
		Template: spell,
		ModifyCast: func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
			instance.Effect.Target = target

			spell := &instance.SpellCast
			if shaman.ElementalFocusStacks > 0 {
				// Reduces mana cost by 40%
				spell.Cost.Value -= spell.BaseCost.Value * 0.4
			}
			if shaman.HasAura(ShamanisticFocusAuraID) {
				spell.Cost.Value -= spell.BaseCost.Value * 0.6
			}
			if shaman.HasAura(ElementalMasteryAuraID) {
				spell.Cost.Value = 0
			}
		},
	}
}

func (shaman *Shaman) registerEarthShockSpell(sim *core.Simulation) {
	config := shaman.newShockSpellConfig(sim, SpellIDEarthShock, core.SpellSchoolNature, 535.0)
	config.Template.SpellExtras |= core.SpellExtrasBinary
	config.Template.Effect.BaseDamage = core.BaseDamageConfigMagic(661, 696, 0.386)

	shaman.EarthShock = shaman.RegisterSpell(config)
}

func (shaman *Shaman) registerFlameShockSpell(sim *core.Simulation) {
	config := shaman.newShockSpellConfig(sim, SpellIDFlameShock, core.SpellSchoolFire, 500.0)

	config.Template.Effect.BaseDamage = core.BaseDamageConfigMagic(377, 377, 0.214)
	config.Template.Effect.DotInput = core.DotDamageInput{
		NumberOfTicks:  4,
		TickLength:     time.Second * 3,
		TickBaseDamage: core.DotSnapshotFuncMagic(420/4, 0.1),
		DebuffID:       FlameShockDebuffID,
	}

	shaman.FlameShock = shaman.RegisterSpell(config)
}

func (shaman *Shaman) registerFrostShockSpell(sim *core.Simulation) {
	config := shaman.newShockSpellConfig(sim, SpellIDFrostShock, core.SpellSchoolFrost, 525.0)
	config.Template.SpellExtras |= core.SpellExtrasBinary

	config.Template.Effect.BaseDamage = core.BaseDamageConfigMagic(647, 683, 0.386)
	config.Template.Effect.ThreatMultiplier *= 2

	shaman.FrostShock = shaman.RegisterSpell(config)
}
