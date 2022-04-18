package shaman

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDSearingTotem int32 = 25533

var SearingTotemActionID = core.ActionID{SpellID: SpellIDSearingTotem}

func (shaman *Shaman) registerSearingTotemSpell(sim *core.Simulation) {
	baseCost := 205.0

	shaman.SearingTotem = shaman.RegisterSpell(core.SpellConfig{
		ActionID:    SearingTotemActionID,
		SpellSchool: core.SpellSchoolFire,
		SpellExtras: SpellFlagTotem,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost -
					baseCost*float64(shaman.Talents.TotemicFocus)*0.05 -
					baseCost*float64(shaman.Talents.MentalQuickness)*0.02,
				GCD: time.Second,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, _ *core.Spell) {
			shaman.SearingTotemDot.Apply(sim)
			// +1 needed because of rounding issues with Searing totem tick time.
			shaman.NextTotemDrops[FireTotem] = sim.CurrentTime + time.Second*60 + 1
			shaman.tryTwistFireNova(sim)
		},
	})

	target := sim.GetPrimaryTarget()
	shaman.SearingTotemDot = core.NewDot(core.Dot{
		Spell: shaman.SearingTotem,
		Aura: target.RegisterAura(core.Aura{
			Label:    "SearingTotem-" + strconv.Itoa(int(shaman.Index)),
			ActionID: SearingTotemActionID,
		}),
		// These are the real tick values, but searing totem doesn't start its next
		// cast until the previous missile hits the target. We don't have an option
		// for target distance yet so just pretend the tick rate is lower.
		//NumberOfTicks:        30,
		//TickLength:           time.Second * 2,
		NumberOfTicks: 24,
		TickLength:    time.Second * 60 / 24,
		TickEffects: core.TickFuncApplyEffects(core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			BonusSpellHitRating: float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance,
			DamageMultiplier:    1 + float64(shaman.Talents.CallOfFlame)*0.05,
			IsPhantom:           true,
			BaseDamage:          core.BaseDamageConfigMagic(50, 66, 0.167),
			OutcomeApplier:      core.OutcomeFuncMagicHitAndCrit(shaman.ElementalCritMultiplier()),
		})),
	})
}

const SpellIDMagmaTotem int32 = 25552

var MagmaTotemActionID = core.ActionID{SpellID: SpellIDMagmaTotem}

func (shaman *Shaman) registerMagmaTotemSpell(sim *core.Simulation) {
	//AOECap: 1600,
	baseCost := 800.0

	shaman.MagmaTotem = shaman.RegisterSpell(core.SpellConfig{
		ActionID:    MagmaTotemActionID,
		SpellSchool: core.SpellSchoolFire,
		SpellExtras: SpellFlagTotem,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost -
					baseCost*float64(shaman.Talents.TotemicFocus)*0.05 -
					baseCost*float64(shaman.Talents.MentalQuickness)*0.02,
				GCD: time.Second,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, _ *core.Spell) {
			shaman.MagmaTotemDot.Apply(sim)
			shaman.NextTotemDrops[FireTotem] = sim.CurrentTime + time.Second*20 + 1
			shaman.tryTwistFireNova(sim)
		},
	})

	target := sim.GetPrimaryTarget()
	shaman.MagmaTotemDot = core.NewDot(core.Dot{
		Spell: shaman.MagmaTotem,
		Aura: target.RegisterAura(core.Aura{
			Label:    "MagmaTotem-" + strconv.Itoa(int(shaman.Index)),
			ActionID: MagmaTotemActionID,
		}),
		NumberOfTicks: 10,
		TickLength:    time.Second * 2,
		TickEffects: core.TickFuncApplyEffects(core.ApplyEffectFuncAOEDamage(sim, core.SpellEffect{
			BonusSpellHitRating: float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance,
			DamageMultiplier:    1 + float64(shaman.Talents.CallOfFlame)*0.05,
			IsPhantom:           true,
			BaseDamage:          core.BaseDamageConfigMagicNoRoll(97, 0.067),
			OutcomeApplier:      core.OutcomeFuncMagicHitAndCrit(shaman.ElementalCritMultiplier()),
		})),
	})
}

const SpellIDNovaTotem int32 = 25537

var FireNovaTotemActionID = core.ActionID{SpellID: SpellIDNovaTotem}

func (shaman *Shaman) FireNovaTickLength() time.Duration {
	return time.Second * time.Duration(4-shaman.Talents.ImprovedFireTotems)
}

// This is probably not worth simming since no other spell in the game does this and AM isn't
// even a popular choice for arcane mages.
func (shaman *Shaman) registerNovaTotemSpell(sim *core.Simulation) {
	//AOECap: 9975,
	baseCost := 765.0

	tickLength := shaman.FireNovaTickLength()
	shaman.FireNovaTotem = shaman.RegisterSpell(core.SpellConfig{
		ActionID:    FireNovaTotemActionID,
		SpellSchool: core.SpellSchoolFire,
		SpellExtras: SpellFlagTotem,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost -
					baseCost*float64(shaman.Talents.TotemicFocus)*0.05 -
					baseCost*float64(shaman.Talents.MentalQuickness)*0.02,
				GCD: time.Second,
			},
			CD: core.Cooldown{
				Timer:    shaman.NewTimer(),
				Duration: time.Second * 15,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, _ *core.Spell) {
			shaman.FireNovaTotemDot.Apply(sim)
			shaman.NextTotemDrops[FireTotem] = sim.CurrentTime + tickLength + 1
			shaman.tryTwistFireNova(sim)
		},
	})

	target := sim.GetPrimaryTarget()
	shaman.FireNovaTotemDot = core.NewDot(core.Dot{
		Spell: shaman.FireNovaTotem,
		Aura: target.RegisterAura(core.Aura{
			Label:    "FireNovaTotem-" + strconv.Itoa(int(shaman.Index)),
			ActionID: FireNovaTotemActionID,
		}),
		NumberOfTicks: 1,
		TickLength:    tickLength,
		TickEffects: core.TickFuncApplyEffects(core.ApplyEffectFuncAOEDamage(sim, core.SpellEffect{
			BonusSpellHitRating: float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance,
			DamageMultiplier:    1 + float64(shaman.Talents.CallOfFlame)*0.05,
			IsPhantom:           true,
			BaseDamage:          core.BaseDamageConfigMagic(654, 730, 0.214),
			OutcomeApplier:      core.OutcomeFuncMagicHitAndCrit(shaman.ElementalCritMultiplier()),
		})),
	})
}
