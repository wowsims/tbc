package shaman

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDEarthShock int32 = 25454
const SpellIDFlameShock int32 = 25457
const SpellIDFrostShock int32 = 25464

var ShockCooldownID = core.NewCooldownID() // shared CD for all shocks

func (shaman *Shaman) ShockCD() time.Duration {
	return time.Second*6 - time.Millisecond*200*time.Duration(shaman.Talents.Reverberation)
}

// Shared logic for all shocks.
func (shaman *Shaman) newShockSpellConfig(sim *core.Simulation, spellID int32, spellSchool core.SpellSchool, baseManaCost float64) (core.SpellConfig, core.SpellEffect) {
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
	}

	spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.Convection) * 0.02
	spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.MentalQuickness) * 0.02
	// TODO: confirm this is how it reduces mana cost.
	if ItemSetSkyshatterHarness.CharacterHasSetBonus(&shaman.Character, 2) {
		spell.Cost.Value -= spell.BaseCost.Value * 0.1
	}

	effect := core.SpellEffect{
		BonusSpellHitRating: float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance,
		BonusSpellPower: 0 +
			core.TernaryFloat64(shaman.Equip[items.ItemSlotRanged].ID == TotemOfRage, 30, 0) +
			core.TernaryFloat64(shaman.Equip[items.ItemSlotRanged].ID == TotemOfImpact, 46, 0),
		DamageMultiplier: 1 * (1 + 0.01*float64(shaman.Talents.Concussion)),
		ThreatMultiplier: 1 - (0.1/3)*float64(shaman.Talents.ElementalPrecision),
	}

	if shaman.Talents.ElementalFocus {
		spell.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
			if shaman.ElementalFocusStacks > 0 {
				shaman.ElementalFocusStacks--
			}
		}
		effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Outcome.Matches(core.OutcomeCrit) {
				shaman.ElementalFocusStacks = 2
			}
		}
	}

	return core.SpellConfig{
		Template: spell,
		ModifyCast: func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
			spell := &instance.SpellCast
			if shaman.ElementalFocusStacks > 0 {
				// Reduces mana cost by 40%
				spell.Cost.Value -= spell.BaseCost.Value * 0.4
			}
			if shaman.ShamanisticFocusAura != nil && shaman.ShamanisticFocusAura.IsActive() {
				spell.Cost.Value -= spell.BaseCost.Value * 0.6
			}
			if shaman.ElementalMasteryAura != nil && shaman.ElementalMasteryAura.IsActive() {
				spell.Cost.Value = 0
			}
		},
	}, effect
}

func (shaman *Shaman) registerEarthShockSpell(sim *core.Simulation) {
	config, effect := shaman.newShockSpellConfig(sim, SpellIDEarthShock, core.SpellSchoolNature, 535.0)
	config.Template.SpellExtras |= core.SpellExtrasBinary

	effect.BaseDamage = core.BaseDamageConfigMagic(661, 696, 0.386)
	effect.OutcomeApplier = core.OutcomeFuncMagicHitAndCrit(shaman.ElementalCritMultiplier())
	config.ApplyEffects = core.ApplyEffectFuncDirectDamage(effect)

	shaman.EarthShock = shaman.RegisterSpell(config)
}

func (shaman *Shaman) registerFlameShockSpell(sim *core.Simulation) {
	config, effect := shaman.newShockSpellConfig(sim, SpellIDFlameShock, core.SpellSchoolFire, 500.0)

	effect.BaseDamage = core.BaseDamageConfigMagic(377, 377, 0.214)
	effect.OutcomeApplier = core.OutcomeFuncMagicHitAndCrit(shaman.ElementalCritMultiplier())
	if effect.OnSpellHit == nil {
		effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Landed() {
				shaman.FlameShockDot.Apply(sim)
			}
		}
	} else {
		oldSpellHit := effect.OnSpellHit
		effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			oldSpellHit(sim, spell, spellEffect)
			if spellEffect.Landed() {
				shaman.FlameShockDot.Apply(sim)
			}
		}
	}

	config.ApplyEffects = core.ApplyEffectFuncDirectDamage(effect)
	shaman.FlameShock = shaman.RegisterSpell(config)

	target := sim.GetPrimaryTarget()
	shaman.FlameShockDot = core.NewDot(core.Dot{
		Spell: shaman.FlameShock,
		Aura: target.RegisterAura(&core.Aura{
			Label:    "FlameShock-" + strconv.Itoa(int(shaman.Index)),
			ActionID: core.ActionID{SpellID: SpellIDFlameShock},
		}),
		NumberOfTicks: 4,
		TickLength:    time.Second * 3,
		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1 * (1 + 0.01*float64(shaman.Talents.Concussion)),
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(420/4, 0.1),
			OutcomeApplier:   core.OutcomeFuncTick(),
			IsPeriodic:       true,
		}),
	})
}

func (shaman *Shaman) registerFrostShockSpell(sim *core.Simulation) {
	config, effect := shaman.newShockSpellConfig(sim, SpellIDFrostShock, core.SpellSchoolFrost, 525.0)
	config.Template.SpellExtras |= core.SpellExtrasBinary

	effect.ThreatMultiplier *= 2
	effect.BaseDamage = core.BaseDamageConfigMagic(647, 683, 0.386)
	effect.OutcomeApplier = core.OutcomeFuncMagicHitAndCrit(shaman.ElementalCritMultiplier())
	config.ApplyEffects = core.ApplyEffectFuncDirectDamage(effect)

	shaman.FrostShock = shaman.RegisterSpell(config)
}
