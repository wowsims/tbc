package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var ExposeArmorActionID = core.ActionID{SpellID: 26866, Tag: 5}
var ExposeArmorEnergyCost = 25.0

func (rogue *Rogue) newExposeArmorTemplate(_ *core.Simulation) core.SimpleSpellTemplate {
	refundAmount := 0.4 * float64(rogue.Talents.QuickRecovery)

	ability := rogue.newAbility(ExposeArmorActionID, ExposeArmorEnergyCost, SpellFlagFinisher, core.ProcMaskMeleeMHSpecial)
	ability.SpellCast.Cast.CritRollCategory = core.CritRollCategoryNone
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			spellEffect.Target.AddAura(sim, core.ExposeArmorAura(sim, spellEffect.Target, rogue.Talents.ImprovedExposeArmor))
			rogue.ApplyFinisher(sim, spellCast.ActionID)
			if sim.GetRemainingDuration() <= time.Second*30 {
				rogue.doneEA = true
			}
		} else {
			if refundAmount > 0 {
				rogue.AddEnergy(sim, spellCast.Cost.Value*refundAmount, core.ActionID{SpellID: 31245})
			}
		}
	}

	if rogue.Talents.SurpriseAttacks {
		ability.SpellExtras |= core.SpellExtrasCannotBeDodged
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (rogue *Rogue) NewExposeArmor(_ *core.Simulation, target *core.Target) *core.SimpleSpell {
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
	return !rogue.doneEA && (rogue.Talents.ImprovedExposeArmor == 2 || !target.HasAura(core.SunderArmorDebuffID))
}
