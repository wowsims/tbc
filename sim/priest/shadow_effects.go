package priest

import (
	"github.com/wowsims/tbc/sim/core"
)

func (priest *Priest) ApplyMisery(sim *core.Simulation, target *core.Unit) {
	if priest.MiseryAura != nil {
		priest.MiseryAura.Activate(sim)
	}
}

func (priest *Priest) ApplyShadowWeaving(sim *core.Simulation, target *core.Unit) {
	if priest.ShadowWeavingAura == nil {
		return
	}

	if priest.Talents.ShadowWeaving < 5 && sim.RandomFloat("Shadow Weaving") > 0.2*float64(priest.Talents.ShadowWeaving) {
		return
	}

	priest.ShadowWeavingAura.Activate(sim)
	if priest.ShadowWeavingAura.IsActive() {
		priest.ShadowWeavingAura.AddStack(sim)
	}
}

func (priest *Priest) ApplyShadowOnHitEffects() {
	// This is a combined aura for all priest major on hit effects.
	//  Shadow Weaving, Vampiric Touch, and Misery
	priest.RegisterAura(core.Aura{
		Label:    "Priest Shadow Effects",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnPeriodicDamageDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Damage > 0 && priest.VampiricTouchDot.IsActive() {
				amount := spellEffect.Damage * 0.05
				for _, partyMember := range priest.Party.Players {
					partyMember.GetCharacter().AddMana(sim, amount, priest.VampiricTouch.ActionID, false)
				}
				for _, petAgent := range priest.Party.Pets {
					pet := petAgent.GetPet()
					if pet.IsEnabled() {
						pet.Character.AddMana(sim, amount, priest.VampiricTouch.ActionID, false)
					}
				}
			}
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() {
				return
			}
			priest.ApplyShadowWeaving(sim, spellEffect.Target)
			if spellEffect.Damage > 0 && priest.VampiricTouchDot.IsActive() {
				amount := spellEffect.Damage * 0.05
				for _, partyMember := range priest.Party.Players {
					partyMember.GetCharacter().AddMana(sim, amount, priest.VampiricTouch.ActionID, false)
				}
				for _, petAgent := range priest.Party.Pets {
					pet := petAgent.GetPet()
					if pet.IsEnabled() {
						pet.Character.AddMana(sim, amount, priest.VampiricTouch.ActionID, false)
					}
				}
			}

			if spell == priest.ShadowWordPain || spell == priest.VampiricTouch || spell.ActionID.SpellID == priest.MindFlay[1].ActionID.SpellID {
				priest.ApplyMisery(sim, spellEffect.Target)
			}
		},
	})
}
