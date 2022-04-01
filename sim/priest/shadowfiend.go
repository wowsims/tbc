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
				priest.Shadowfiend.Cast(sim, sim.GetPrimaryTarget())

				// All MCDs that use the GCD and have a non-zero cast time must call this.
				priest.UpdateMajorCooldowns()
			}
		},
	})
}

func (priest *Priest) registerShadowfiendSpell(sim *core.Simulation) {
	cost := core.ResourceCost{Type: stats.Mana, Value: priest.BaseMana() * 0.06}
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    ShadowfiendActionID,
				Character:   &priest.Character,
				SpellSchool: core.SpellSchoolShadow,
				BaseCost:    cost,
				Cost:        cost,
				CastTime:    0,
				GCD:         core.GCDDefault,
				Cooldown:    time.Minute * 5,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      priest.DefaultSpellCritMultiplier(),
			// Dmg over 15 sec = shadow_dmg*.6 + 1191
			// just simulate 10 1.5s long ticks
			DamageMultiplier: 1,
			DotInput: core.DotDamageInput{
				NumberOfTicks:  10,
				TickLength:     time.Millisecond * 1500,
				TickBaseDamage: core.DotSnapshotFuncMagic(1191/10, 0.06),
				OnPeriodicDamage: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect, tickDamage float64) {
					// TODO: This should also do something with ExpectedBonusMana
					priest.AddMana(sim, tickDamage*2.5, ShadowfiendActionID, false)
				},
			},
		},
	}

	priest.applyTalentsToShadowSpell(&template.SpellCast.Cast, &template.Effect)

	if ItemSetIncarnate.CharacterHasSetBonus(&priest.Character, 2) { // Increases duration by 3s
		template.Effect.DotInput.NumberOfTicks += 2
	}

	priest.Shadowfiend = priest.RegisterSpell(core.SpellConfig{
		Template:   template,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
