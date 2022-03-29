package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDSearingTotem int32 = 25533

func (shaman *Shaman) newSearingTotemTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	cost := core.ResourceCost{Type: stats.Mana, Value: 205}
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: SpellIDSearingTotem},
				Character:           &shaman.Character,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				CritRollCategory:    core.CritRollCategoryMagical,
				SpellSchool:         core.SpellSchoolFire,
				SpellExtras:         SpellFlagTotem | core.SpellExtrasAlwaysHits,
				BaseCost:            cost,
				Cost:                cost,
				GCD:                 time.Second,
				CritMultiplier:      shaman.DefaultSpellCritMultiplier(),
				IsPhantom:           true,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier: 1,
			},
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
	spell.Effect.SpellEffect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance

	spell.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.NextTotemDrops[FireTotem] = sim.CurrentTime + time.Second*60
		shaman.tryTwistFireNova(sim)
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (shaman *Shaman) NewSearingTotem(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	searingTotem := &shaman.FireTotemSpell
	shaman.searingTotemTemplate.Apply(searingTotem)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	searingTotem.Effect.Target = target
	searingTotem.Init(sim)

	return searingTotem
}

const SpellIDMagmaTotem int32 = 25552

// This is probably not worth simming since no other spell in the game does this and AM isn't
// even a popular choice for arcane mages.
func (shaman *Shaman) newMagmaTotemTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	cost := core.ResourceCost{Type: stats.Mana, Value: 800}
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: SpellIDMagmaTotem},
				Character:           &shaman.Character,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				CritRollCategory:    core.CritRollCategoryMagical,
				SpellSchool:         core.SpellSchoolFire,
				SpellExtras:         SpellFlagTotem | core.SpellExtrasAlwaysHits,
				BaseCost:            cost,
				Cost:                cost,
				GCD:                 time.Second,
				CritMultiplier:      shaman.DefaultSpellCritMultiplier(),
				IsPhantom:           true,
			},
		},
		AOECap: 1600,
	}
	spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.TotemicFocus) * 0.05
	spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.MentalQuickness) * 0.02
	if shaman.Talents.ElementalFury {
		spell.CritMultiplier = shaman.SpellCritMultiplier(1, 1)
	}

	baseEffect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier: 1,
		},
		DotInput: core.DotDamageInput{
			NumberOfTicks:       10,
			TickLength:          time.Second * 2,
			TickBaseDamage:      core.DotSnapshotFuncMagic(97, 0.067),
			TicksCanMissAndCrit: true,
		},
	}
	baseEffect.DamageMultiplier *= 1 + float64(shaman.Talents.CallOfFlame)*0.05
	baseEffect.SpellEffect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance

	spell.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.NextTotemDrops[FireTotem] = sim.CurrentTime + time.Second*20
		shaman.tryTwistFireNova(sim)
	}

	numHits := sim.GetNumTargets()
	effects := make([]core.SpellHitEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	spell.Effects = effects

	return core.NewSimpleSpellTemplate(spell)
}

func (shaman *Shaman) NewMagmaTotem(sim *core.Simulation) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	magmaTotem := &shaman.FireTotemSpell
	shaman.magmaTotemTemplate.Apply(magmaTotem)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	magmaTotem.Init(sim)

	return magmaTotem
}

const SpellIDNovaTotem int32 = 25537

var CooldownIDNovaTotem = core.NewCooldownID()

func (shaman *Shaman) FireNovaTickLength() time.Duration {
	return time.Second * time.Duration(4-shaman.Talents.ImprovedFireTotems)
}

// This is probably not worth simming since no other spell in the game does this and AM isn't
// even a popular choice for arcane mages.
func (shaman *Shaman) newNovaTotemTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	cost := core.ResourceCost{Type: stats.Mana, Value: 765}
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID:    SpellIDNovaTotem,
					CooldownID: CooldownIDNovaTotem,
				},
				Character:           &shaman.Character,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				CritRollCategory:    core.CritRollCategoryMagical,
				SpellSchool:         core.SpellSchoolFire,
				SpellExtras:         SpellFlagTotem | core.SpellExtrasAlwaysHits,
				BaseCost:            cost,
				Cost:                cost,
				GCD:                 time.Second,
				Cooldown:            time.Second * 15,
				CritMultiplier:      shaman.DefaultSpellCritMultiplier(),
				IsPhantom:           true,
			},
		},
		AOECap: 9975,
	}
	spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.TotemicFocus) * 0.05
	spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.MentalQuickness) * 0.02
	if shaman.Talents.ElementalFury {
		spell.CritMultiplier = shaman.SpellCritMultiplier(1, 1)
	}

	baseEffect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier: 1,
		},
		DotInput: core.DotDamageInput{
			NumberOfTicks:       1,
			TickLength:          shaman.FireNovaTickLength(),
			TickBaseDamage:      core.DotSnapshotFuncMagic(692, 0.214),
			TicksCanMissAndCrit: true,
		},
	}
	baseEffect.DamageMultiplier *= 1 + float64(shaman.Talents.CallOfFlame)*0.05
	baseEffect.SpellEffect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance

	tickLength := baseEffect.DotInput.TickLength
	spell.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.NextTotemDrops[FireTotem] = sim.CurrentTime + tickLength
		shaman.tryTwistFireNova(sim)
	}

	numHits := sim.GetNumTargets()
	effects := make([]core.SpellHitEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	spell.Effects = effects

	return core.NewSimpleSpellTemplate(spell)
}

func (shaman *Shaman) NewNovaTotem(sim *core.Simulation) *core.SimpleSpell {
	// If we drop nova while another totem is running, cancel it.
	if shaman.FireTotemSpell.IsInUse() {
		shaman.FireTotemSpell.Cancel(sim)
	}

	// Initialize cast from precomputed template.
	novaTotem := &shaman.FireTotemSpell
	shaman.novaTotemTemplate.Apply(novaTotem)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	novaTotem.Init(sim)

	return novaTotem
}
