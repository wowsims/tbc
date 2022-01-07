package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var SummonWaterElementalCooldownID = core.NewCooldownID()

func (mage *Mage) registerSummonWaterElementalCD() {
	if !mage.Talents.SummonWaterElemental {
		return
	}

	manaCost := 0.0
	actionID := core.ActionID{SpellID: 31687}

	mage.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: SummonWaterElementalCooldownID,
		Cooldown:   time.Minute * 3,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			if mage.waterElemental.IsEnabled() {
				return false
			}
			if character.CurrentMana() < manaCost {
				return false
			}
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				manaCost = mage.BaseMana() * 0.16
				manaCost -= manaCost * float64(mage.Talents.FrostChanneling) * 0.05

				mage.waterElemental.EnableWithTimeout(sim, mage.waterElemental, time.Second*45)

				character.SpendMana(sim, manaCost, actionID)
				character.Metrics.AddInstantCast(actionID)
				character.SetCD(SummonWaterElementalCooldownID, sim.CurrentTime+time.Minute*3)
			}
		},
	})
}

type WaterElemental struct {
	core.Pet

	// Water Ele almost never just stands still and spams like we want, it sometimes
	// does its own thing. This controls how much it does that.
	disobeyChance float64

	waterboltSpell        core.SimpleSpell
	waterboltCastTemplate core.SimpleSpellTemplate
}

func (mage *Mage) NewWaterElemental(disobeyChance float64) *WaterElemental {
	waterElemental := &WaterElemental{
		Pet: core.NewPet(
			"Water Elemental",
			&mage.Character,
			waterElementalBaseStats,
			waterElementalInheritanceCoeffs,
			false,
		),
		disobeyChance: disobeyChance,
	}

	mage.AddPet(waterElemental)

	return waterElemental
}

func (we *WaterElemental) GetPet() *core.Pet {
	return &we.Pet
}

func (we *WaterElemental) Init(sim *core.Simulation) {
	we.waterboltCastTemplate = we.newWaterboltTemplate(sim)
}

func (we *WaterElemental) Reset(newsim *core.Simulation) {
}

func (we *WaterElemental) Advance(sim *core.Simulation, elapsedTime time.Duration) {
	// Water ele probably has mana regen?
}

func (we *WaterElemental) Act(sim *core.Simulation) time.Duration {
	spell := we.NewWaterbolt(sim, sim.GetPrimaryTarget())

	if sim.RandomFloat("Water Elemental Disobey") < we.disobeyChance {
		// Water ele has decided not to cooperate, so just wait for the cast time
		// instead of casting.
		spell.Cancel(sim)
		waitAction := core.NewWaitAction(sim, we.GetCharacter(), spell.GetDuration(), core.WaitReasonNone)
		waitAction.Cast(sim)
		return sim.CurrentTime + waitAction.GetDuration()
	}

	actionSuccessful := spell.Cast(sim)

	if !actionSuccessful {
		// If water ele has gone OOM then there won't be enough time left for meaningful
		// regen to occur before the ele expires. So just murder itself.
		we.Disable(sim)
	}

	return sim.CurrentTime + spell.GetDuration()
}

// These numbers are just rough guesses based on looking at some logs.
var waterElementalBaseStats = stats.Stats{
	stats.Intellect:  100,
	stats.SpellPower: 500,
	stats.Mana:       2000,
}

// These numbers are just rough guesses based on looking at some logs.
var waterElementalInheritanceCoeffs = stats.Stats{
	// Computed based on my lvl 65 mage, need to ask someone with a 70 to check these
	stats.Stamina:   0.2238,
	stats.Intellect: 0.01,

	stats.SpellPower:      0.333,
	stats.FrostSpellPower: 0.333,
	stats.SpellHit:        0,
	stats.SpellCrit:       0,
}

const SpellIDWaterbolt int32 = 31707

func (we *WaterElemental) newWaterboltTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseManaCost := we.BaseMana() * 0.1
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				CritMultiplier: 1.5,
				SpellSchool:    stats.FrostSpellPower,
				Character:      &we.Character,
				BaseManaCost:   baseManaCost,
				ManaCost:       baseManaCost,
				CastTime:       time.Millisecond * 3000,
				ActionID: core.ActionID{
					SpellID: SpellIDWaterbolt,
				},
			},
		},
		SpellHitEffect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage:    256,
				MaxBaseDamage:    328,
				SpellCoefficient: 3.0 / 3.5,
			},
		},
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (we *WaterElemental) NewWaterbolt(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	waterbolt := &we.waterboltSpell
	we.waterboltCastTemplate.Apply(waterbolt)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	waterbolt.Target = target
	waterbolt.Init(sim)

	return waterbolt
}
