package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ExposeArmorActionID = core.ActionID{SpellID: 26866, Tag: 5}
var ExposeArmorEnergyCost = 25.0

func (rogue *Rogue) newExposeArmorTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	finishingMoveEffects := rogue.makeFinishingMoveEffectApplier(sim)
	refundAmount := 0.4 * float64(rogue.Talents.QuickRecovery)

	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            ExposeArmorActionID,
				Character:           &rogue.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryNone,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 time.Second * 1,
				Cost: core.ResourceCost{
					Type:  stats.Energy,
					Value: ExposeArmorEnergyCost,
				},
				SpellExtras: SpellFlagFinisher,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					if spellEffect.Landed() {
						spellEffect.Target.AddAura(sim, core.ExposeArmorAura(sim, spellEffect.Target, rogue.Talents.ImprovedExposeArmor))
						numPoints := rogue.ComboPoints()
						rogue.SpendComboPoints(sim, spellCast.ActionID)
						finishingMoveEffects(sim, numPoints)
					} else {
						if refundAmount > 0 {
							rogue.AddEnergy(sim, spellCast.Cost.Value*refundAmount, core.ActionID{SpellID: 31245})
						}
					}
				},
			},
		},
	}

	if rogue.Talents.SurpriseAttacks {
		ability.SpellExtras |= core.SpellExtrasCannotBeDodged
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (rogue *Rogue) NewExposeArmor(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	if rogue.ComboPoints() != 5 {
		panic("Expose Armor requires 5 combo points!")
	}

	ea := &rogue.exposeArmor
	rogue.exposeArmorTemplate.Apply(ea)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ea.Effect.Target = target

	if rogue.deathmantle4pcProc {
		ea.Cost.Value = 0
		rogue.deathmantle4pcProc = false
	}

	return ea
}

func (rogue *Rogue) MaintainingExpose(target *core.Target) bool {
	permaEA := target.AuraExpiresAt(core.ExposeArmorDebuffID) == core.NeverExpires
	return rogue.Rotation.MaintainExposeArmor &&
		!permaEA &&
		(rogue.Talents.ImprovedExposeArmor == 2 || !target.HasAura(core.SunderArmorDebuffID))
}
