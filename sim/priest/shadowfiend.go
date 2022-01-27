package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

// https://web.archive.org/web/20071201221602/http://www.shadowpriest.com/viewtopic.php?t=7616

const SpellIDShadowfiend int32 = 34433

var ShadowfiendCD = core.NewCooldownID()
var ShadowfiendActionID = core.ActionID{SpellID: SpellIDShadowfiend, CooldownID: ShadowfiendCD}

func (priest *Priest) registerShadowfiendCD() {
	if !priest.UseShadowfiend {
		return
	}

	priest.AddMajorCooldown(core.MajorCooldown{
		ActionID:   ShadowfiendActionID,
		CooldownID: ShadowfiendCD,
		Cooldown:   time.Minute * 5,
		UsesGCD:    true,
		Type:       core.CooldownTypeMana,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			if character.CurrentMana() < 575 {
				return false
			}

			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			if character.CurrentManaPercent() >= 0.5 {
				return false
			}

			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				priest.NewShadowfiend(sim, sim.GetPrimaryTarget()).Cast(sim)

				// All MCDs that use the GCD and have a non-zero cast time must call this.
				priest.UpdateMajorCooldowns()
			}
		},
	})
}

func (priest *Priest) newShadowfiendTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		CritMultiplier: 1.5,
		SpellSchool:    stats.ShadowSpellPower,
		Character:      &priest.Character,
		BaseManaCost:   575,
		ManaCost:       575,
		CastTime:       0,
		Cooldown:       time.Minute * 5,
		ActionID:       ShadowfiendActionID,
	}

	// Dmg over 15 sec = shadow_dmg*.6 + 1191
	// just simulate 10 1.5s long ticks
	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
		},
		DotInput: core.DotDamageInput{
			NumberOfTicks:        10,
			TickLength:           time.Millisecond * 1500,
			TickBaseDamage:       1191 / 10,
			TickSpellCoefficient: 0.06,
			OnPeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage float64) {
				// TODO: This should also do something with ExpectedBonusMana
				priest.AddMana(sim, tickDamage*2.5, ShadowfiendActionID, false)
			},
		},
	}

	priest.applyTalentsToShadowSpell(&baseCast, &effect)

	if ItemSetIncarnate.CharacterHasSetBonus(&priest.Character, 2) { // Increases duration by 3s
		effect.DotInput.NumberOfTicks += 2
	}

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		Effect: effect,
	})
}

func (priest *Priest) NewShadowfiend(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	mf := &priest.ShadowfiendSpell

	priest.shadowfiendTemplate.Apply(mf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mf.Effect.Target = target
	mf.Init(sim)

	return mf
}
