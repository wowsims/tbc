package priest

import (
	"github.com/wowsims/tbc/sim/core"
)

func (priest *Priest) applyTalentsToHolySpell(cast *core.Cast, effect *core.SpellHitEffect) {
	effect.ThreatMultiplier *= 1 - 0.04*float64(priest.Talents.SilentResolve)
	if cast.ActionID.SpellID == SpellIDSmite {
		effect.BonusSpellCritRating += float64(priest.Talents.HolySpecialization) * 1 * core.SpellCritRatingPerCritChance
	}

	effect.BonusSpellCritRating += float64(priest.Talents.ForceOfWill) * 1 * core.SpellCritRatingPerCritChance
}

func (priest *Priest) applyTalentsToShadowSpell(cast *core.Cast, effect *core.SpellHitEffect) {
	effect.ThreatMultiplier *= 1 - 0.08*float64(priest.Talents.ShadowAffinity)
	if cast.ActionID.SpellID == SpellIDShadowWordDeath || cast.ActionID.SpellID == SpellIDMindBlast {
		effect.BonusSpellCritRating += float64(priest.Talents.ShadowPower) * 3 * core.SpellCritRatingPerCritChance
	}
	if cast.ActionID.SpellID == SpellIDMindFlay || cast.ActionID.SpellID == SpellIDMindBlast {
		cast.Cost.Value -= cast.BaseCost.Value * float64(priest.Talents.FocusedMind) * 0.05
	}
	if cast.SpellSchool == core.SpellSchoolShadow {
		effect.StaticDamageMultiplier *= 1 + float64(priest.Talents.Darkness)*0.02

		if priest.Talents.Shadowform {
			effect.StaticDamageMultiplier *= 1.15
		}

		// shadow focus gives 2% hit per level
		effect.BonusSpellHitRating += float64(priest.Talents.ShadowFocus) * 2 * core.SpellHitRatingPerHitChance
		
		// TODO should add more instant cast spells here
		if cast.ActionID.SpellID == SpellIDShadowWordPain {
			cast.Cost.Value -= cast.BaseCost.Value * float64(priest.Talents.MentalAgility) * 0.02
		}

		effect.BonusSpellCritRating += float64(priest.Talents.ForceOfWill) * 1 * core.SpellCritRatingPerCritChance
	}
}

func (priest *Priest) ApplyMisery(sim *core.Simulation, target *core.Target) {
	if priest.Talents.Misery >= target.NumStacks(core.MiseryDebuffID) {
		target.ReplaceAura(sim, core.MiseryAura(sim, priest.Talents.Misery))
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

	target.ReplaceAura(sim, core.ShadowWeavingAura(sim, newStacks))
}

var ShadowWeaverAuraID = core.NewAuraID()

func (priest *Priest) ApplyShadowOnHitEffects() {
	// This is a combined aura for all priest major on hit effects.
	//  Shadow Weaving, Vampiric Touch, and Misery
	priest.Character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: ShadowWeaverAuraID,
			OnPeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage float64) {
				if tickDamage > 0 && priest.VTSpell.Effect.DotInput.IsTicking(sim) {
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
				if spellEffect.Damage > 0 && priest.VTSpell.Effect.DotInput.IsTicking(sim) {
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
