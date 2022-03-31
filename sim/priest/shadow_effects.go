package priest

import (
	"github.com/wowsims/tbc/sim/core"
)

func (priest *Priest) ApplyMisery(sim *core.Simulation, target *core.Target) {
	if priest.Talents.Misery >= target.NumStacks(core.MiseryDebuffID) {
		target.AddAura(sim, core.MiseryAura(target, priest.Talents.Misery))
	}
}

func (priest *Priest) ApplyShadowWeaving(sim *core.Simulation, target *core.Target) {
	if priest.Talents.ShadowWeaving == 0 {
		return
	}

	if priest.Talents.ShadowWeaving < 5 && sim.RandomFloat("Shadow Weaving") > 0.2*float64(priest.Talents.ShadowWeaving) {
		return
	}

	curStacks := target.NumStacks(core.ShadowWeavingDebuffID)
	newStacks := core.MinInt32(curStacks+1, 5)

	if sim.Log != nil && curStacks != newStacks {
		priest.Log(sim, "Applied Shadow Weaving stack, %d --> %d", curStacks, newStacks)
	}

	target.AddAura(sim, core.ShadowWeavingAura(target, newStacks))
}

var ShadowWeaverAuraID = core.NewAuraID()

func (priest *Priest) ApplyShadowOnHitEffects() {
	// This is a combined aura for all priest major on hit effects.
	//  Shadow Weaving, Vampiric Touch, and Misery
	priest.Character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: ShadowWeaverAuraID,
			OnPeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage float64) {
				if tickDamage > 0 && priest.CurVTSpell.Instance.Effect.DotInput.IsTicking(sim) {
					amount := tickDamage * 0.05
					for _, partyMember := range priest.Party.Players {
						partyMember.GetCharacter().AddMana(sim, amount, VampiricTouchActionID, false)
					}
					for _, petAgent := range priest.Party.Pets {
						pet := petAgent.GetPet()
						if pet.IsEnabled() {
							pet.Character.AddMana(sim, amount, VampiricTouchActionID, false)
						}
					}
				}
			},
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}
				priest.ApplyShadowWeaving(sim, spellEffect.Target)
				if spellEffect.Damage > 0 && priest.CurVTSpell.Instance.Effect.DotInput.IsTicking(sim) {
					amount := spellEffect.Damage * 0.05
					for _, partyMember := range priest.Party.Players {
						partyMember.GetCharacter().AddMana(sim, amount, VampiricTouchActionID, false)
					}
					for _, petAgent := range priest.Party.Pets {
						pet := petAgent.GetPet()
						if pet.IsEnabled() {
							pet.Character.AddMana(sim, amount, VampiricTouchActionID, false)
						}
					}
				}

				if spellCast.ActionID.SpellID == SpellIDShadowWordPain || spellCast.ActionID.SpellID == VampiricTouchActionID.SpellID || spellCast.ActionID.SpellID == SpellIDMindFlay {
					priest.ApplyMisery(sim, spellEffect.Target)
				}
			},
		}
	})
}
