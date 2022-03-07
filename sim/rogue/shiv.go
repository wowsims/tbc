package rogue

import (
	"github.com/wowsims/tbc/sim/core/proto"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ShivActionID = core.ActionID{SpellID: 5938}

func (rogue *Rogue) newShivTemplate(_ *core.Simulation) core.SimpleSpellTemplate {
	shivEnergyCost := 20 + 10*rogue.GetOHWeapon().SwingSpeed

	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            ShivActionID,
				Character:           &rogue.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 time.Second,
				IgnoreHaste:         true,
				Cost: core.ResourceCost{
					Type:  stats.Energy,
					Value: shivEnergyCost,
				},
				CritMultiplier: rogue.critMultiplier(false, true),
				SpellExtras:    SpellFlagBuilder | core.SpellExtrasCannotBeDodged,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               core.ProcMaskMeleeOHSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					if spellEffect.Landed() {
						rogue.AddComboPoints(sim, 1, ShivActionID)

						switch rogue.Consumes.OffHandImbue {
						case proto.WeaponImbue_WeaponImbueRogueDeadlyPoison:
							rogue.procDeadlyPoison(sim, spellEffect)
						case proto.WeaponImbue_WeaponImbueRogueInstantPoison:
							rogue.procInstantPoison(sim, spellEffect)
						}
					}
				},
			},
			WeaponInput: core.WeaponDamageInput{
				Normalized:       true,
				DamageMultiplier: 1 + 0.1*float64(rogue.Talents.DualWieldSpecialization),
			},
		},
	}

	if rogue.Talents.SurpriseAttacks {
		ability.Effect.StaticDamageMultiplier += 0.1
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (rogue *Rogue) NewShiv(_ *core.Simulation, target *core.Target) *core.SimpleSpell {
	sh := &rogue.shiv
	rogue.shivTemplate.Apply(sh)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	sh.Effect.Target = target

	return sh
}
