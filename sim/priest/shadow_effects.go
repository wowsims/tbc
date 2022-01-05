package priest

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (priest *Priest) applyTalentsToShadowSpell(cast *core.Cast, effect *core.SpellHitEffect) {
	if cast.ActionID.SpellID == SpellIDShadowWordDeath || cast.ActionID.SpellID == SpellIDMindBlast {
		effect.BonusSpellCritRating += float64(priest.Talents.ShadowPower) * 3 * core.SpellCritRatingPerCritChance
	}
	if cast.ActionID.SpellID == SpellIDMindFlay || cast.ActionID.SpellID == SpellIDMindBlast {
		cast.ManaCost -= cast.BaseManaCost * float64(priest.Talents.FocusedMind) * 0.05
	}
	if cast.SpellSchool == stats.ShadowSpellPower {
		effect.StaticDamageMultiplier *= 1 + float64(priest.Talents.Darkness)*0.02

		if priest.Talents.Shadowform {
			effect.StaticDamageMultiplier *= 1.15
		}

		// shadow focus gives 2% hit per level
		effect.BonusSpellHitRating += float64(priest.Talents.ShadowFocus) * 2 * core.SpellHitRatingPerHitChance
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
			ID:   ShadowWeaverAuraID,
			Name: "Shadow Weaver",
			OnPeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage float64) {
				if tickDamage > 0 && priest.VTSpell.DotInput.IsTicking(sim) {
					amount := tickDamage * 0.05
					for _, partyMember := range priest.Party.Players {
						partyMember.GetCharacter().AddMana(sim, amount, "Vampiric Touch", false)
					}
				}
			},
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				priest.ApplyShadowWeaving(sim, spellEffect.Target)
				if spellEffect.Damage > 0 && priest.VTSpell.DotInput.IsTicking(sim) {
					amount := spellEffect.Damage * 0.05
					for _, partyMember := range priest.Party.Players {
						partyMember.GetCharacter().AddMana(sim, amount, "Vampiric Touch", false)
					}
				}

				if spellCast.ActionID.SpellID == SpellIDShadowWordPain || spellCast.ActionID.SpellID == SpellIDVampiricTouch || spellCast.ActionID.SpellID == SpellIDMindFlay {
					priest.ApplyMisery(sim, spellEffect.Target)
				}
			},
		}
	})
}
