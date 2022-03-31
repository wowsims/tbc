package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDSearingTotem int32 = 25533

func (shaman *Shaman) registerSearingTotemSpell(sim *core.Simulation) {
	cost := core.ResourceCost{Type: stats.Mana, Value: 205}
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: SpellIDSearingTotem},
				Character:   &shaman.Character,
				SpellSchool: core.SpellSchoolFire,
				BaseCost:    cost,
				Cost:        cost,
				GCD:         time.Second,
				SpellExtras: SpellFlagTotem | core.SpellExtrasAlwaysHits,
			},
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      shaman.DefaultSpellCritMultiplier(),
			IsPhantom:           true,
		},
		Effect: core.SpellEffect{
			DamageMultiplier: 1,
			DotInput: core.DotDamageInput{
				// These are the real tick values, but searing totem doesn't start its next
				// cast until the previous missile hits the target. We don't have an option
				// for target distance yet so just pretend the tick rate is lower.
				//NumberOfTicks:        30,
				//TickLength:           time.Second * 2,
				NumberOfTicks: 24,
				TickLength:    time.Second * 60 / 24,

				TickBaseDamage:      core.DotSnapshotFuncMagic(58, 0.167),
				TicksCanMissAndCrit: true,
			},
		},
	}
	spell.Effect.DamageMultiplier *= 1 + float64(shaman.Talents.CallOfFlame)*0.05
	spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.TotemicFocus) * 0.05
	spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.MentalQuickness) * 0.02
	if shaman.Talents.ElementalFury {
		spell.CritMultiplier = shaman.SpellCritMultiplier(1, 1)
	}
	spell.Effect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance

	spell.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.NextTotemDrops[FireTotem] = sim.CurrentTime + time.Second*60
		shaman.tryTwistFireNova(sim)
	}

	shaman.SearingTotem = shaman.RegisterSpell(core.SpellConfig{
		Template:   spell,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

const SpellIDMagmaTotem int32 = 25552

func (shaman *Shaman) registerMagmaTotemSpell(sim *core.Simulation) {
	cost := core.ResourceCost{Type: stats.Mana, Value: 800}
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: SpellIDMagmaTotem},
				Character:   &shaman.Character,
				SpellSchool: core.SpellSchoolFire,
				BaseCost:    cost,
				Cost:        cost,
				GCD:         time.Second,
				SpellExtras: SpellFlagTotem | core.SpellExtrasAlwaysHits,
			},
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      shaman.DefaultSpellCritMultiplier(),
			IsPhantom:           true,
		},
		AOECap: 1600,
	}
	spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.TotemicFocus) * 0.05
	spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.MentalQuickness) * 0.02
	if shaman.Talents.ElementalFury {
		spell.CritMultiplier = shaman.SpellCritMultiplier(1, 1)
	}

	baseEffect := core.SpellEffect{
		DamageMultiplier: 1,
		DotInput: core.DotDamageInput{
			NumberOfTicks:       10,
			TickLength:          time.Second * 2,
			TickBaseDamage:      core.DotSnapshotFuncMagic(97, 0.067),
			TicksCanMissAndCrit: true,
		},
	}
	baseEffect.DamageMultiplier *= 1 + float64(shaman.Talents.CallOfFlame)*0.05
	baseEffect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance

	spell.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.NextTotemDrops[FireTotem] = sim.CurrentTime + time.Second*20
		shaman.tryTwistFireNova(sim)
	}

	numHits := sim.GetNumTargets()
	effects := make([]core.SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	spell.Effects = effects

	shaman.MagmaTotem = shaman.RegisterSpell(core.SpellConfig{
		Template: spell,
	})
}

const SpellIDNovaTotem int32 = 25537

var CooldownIDNovaTotem = core.NewCooldownID()

func (shaman *Shaman) FireNovaTickLength() time.Duration {
	return time.Second * time.Duration(4-shaman.Talents.ImprovedFireTotems)
}

// This is probably not worth simming since no other spell in the game does this and AM isn't
// even a popular choice for arcane mages.
func (shaman *Shaman) registerNovaTotemSpell(sim *core.Simulation) {
	cost := core.ResourceCost{Type: stats.Mana, Value: 765}
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID:    SpellIDNovaTotem,
					CooldownID: CooldownIDNovaTotem,
				},
				Character:   &shaman.Character,
				SpellSchool: core.SpellSchoolFire,
				BaseCost:    cost,
				Cost:        cost,
				GCD:         time.Second,
				Cooldown:    time.Second * 15,
				SpellExtras: SpellFlagTotem | core.SpellExtrasAlwaysHits,
			},
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      shaman.DefaultSpellCritMultiplier(),
			IsPhantom:           true,
		},
		AOECap: 9975,
	}
	spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.TotemicFocus) * 0.05
	spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.MentalQuickness) * 0.02
	if shaman.Talents.ElementalFury {
		spell.CritMultiplier = shaman.SpellCritMultiplier(1, 1)
	}

	baseEffect := core.SpellEffect{
		DamageMultiplier: 1,
		DotInput: core.DotDamageInput{
			NumberOfTicks:       1,
			TickLength:          shaman.FireNovaTickLength(),
			TickBaseDamage:      core.DotSnapshotFuncMagic(692, 0.214),
			TicksCanMissAndCrit: true,
		},
	}
	baseEffect.DamageMultiplier *= 1 + float64(shaman.Talents.CallOfFlame)*0.05
	baseEffect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance

	tickLength := baseEffect.DotInput.TickLength
	spell.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.NextTotemDrops[FireTotem] = sim.CurrentTime + tickLength
		shaman.tryTwistFireNova(sim)
	}

	numHits := sim.GetNumTargets()
	effects := make([]core.SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	spell.Effects = effects

	shaman.FireNovaTotem = shaman.RegisterSpell(core.SpellConfig{
		Template: spell,
	})
}
