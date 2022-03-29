package rogue

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

var ShivActionID = core.ActionID{SpellID: 5938}

func (rogue *Rogue) newShivTemplate(_ *core.Simulation) core.SimpleSpellTemplate {
	rogue.shivEnergyCost = 20
	if rogue.GetOHWeapon() != nil {
		rogue.shivEnergyCost = 20 + 10*rogue.GetOHWeapon().SwingSpeed
	}

	ability := rogue.newAbility(ShivActionID, rogue.shivEnergyCost, SpellFlagBuilder|core.SpellExtrasCannotBeDodged, core.ProcMaskMeleeOHSpecial)
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			rogue.AddComboPoints(sim, 1, ShivActionID)

			switch rogue.Consumes.OffHandImbue {
			case proto.WeaponImbue_WeaponImbueRogueDeadlyPoison:
				rogue.procDeadlyPoison(sim, spellEffect)
			case proto.WeaponImbue_WeaponImbueRogueInstantPoison:
				rogue.procInstantPoison(sim, spellEffect)
			}
		}
	}
	ability.Effect.BaseDamage = core.BaseDamageConfigMeleeWeapon(core.OffHand, true, 0, 1+0.1*float64(rogue.Talents.DualWieldSpecialization), true)

	if rogue.Talents.SurpriseAttacks {
		ability.Effect.DamageMultiplier += 0.1
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (rogue *Rogue) NewShiv(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	sh := &rogue.shiv
	rogue.shivTemplate.Apply(sh)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	sh.Effect.Target = target
	sh.Init(sim)

	return sh
}
