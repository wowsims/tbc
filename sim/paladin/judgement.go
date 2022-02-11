package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var JudgementCD = core.NewCooldownID()
var JudgementOfBloodActionID = core.ActionID{SpellID: 31898, CooldownID: JudgementCD}

func (paladin *Paladin) newJudgementOfBloodTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	job := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:       JudgementOfBloodActionID,
				Character:      &paladin.Character,
				SpellSchool:    stats.HolySpellPower,
				BaseManaCost:   147,
				ManaCost:       147,
				Cooldown: 	    time.Second * 10,
				CritMultiplier: paladin.SpellCritMultiplier(1, 0.25), // no idea what the math is for judgment crits
				OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
					currentSeal := &paladin.currentSeal
					currentSeal.Expires = sim.CurrentTime
					paladin.RemoveAura(sim, currentSeal.ID)
				},
			},
		},
		// need to do some research on the effects and inputs
		// unsure if seal of blood scales with spell damage, weapon damage, both?
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage:    295,
				MaxBaseDamage:    325,
				SpellCoefficient: 1,   
			},
		},
	}

	return core.NewSimpleSpellTemplate(job)
}

func (paladin *Paladin) NewJudgementOfBlood(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	job := &paladin.judgementOfBloodSpell
	paladin.judgementOfBloodTemplate.Apply(job)

	job.Effect.Target = target
	job.Init(sim)

	return job
}

var JudgementOfTheCrusaderActionID = core.ActionID{SpellID: 27159, CooldownID: JudgementCD}

func (paladin *Paladin) newJudgementOfTheCrusaderTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	jotc := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:       JudgementOfTheCrusaderActionID,
				Character:      &paladin.Character,
				SpellSchool:    stats.HolySpellPower,
				BaseManaCost:   147,
				ManaCost:       147,
				Cooldown: 	    time.Second * 10,
				OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
					currentSeal := &paladin.currentSeal
					currentSeal.Expires = sim.CurrentTime
					paladin.RemoveAura(sim, currentSeal.ID)
				},
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				IgnoreHitCheck: true,
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					// Need to modify implementation of JudgementOfTheCrusader debuff in core
					spellEffect.Target.AddAura(sim, core.ImprovedSealOfTheCrusaderAura())
				},
			},
		},
	}

	return core.NewSimpleSpellTemplate(jotc)
}

func (paladin *Paladin) NewJudgementOfTheCrusader(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	jotc := &paladin.judgementOfTheCrusaderSpell
	paladin.judgementOfTheCrusaderTemplate.Apply(jotc)

	jotc.Effect.Target = target
	jotc.Init(sim)

	return jotc
}

func (paladin *Paladin) NewJudgement(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	currentSeal := &paladin.currentSeal

	// No seal has even been active, so we can't cast judgement
	if(currentSeal == nil) { return nil }

	// Most recent seal has expired so we can't cast judgement
	if(currentSeal.Expires <= sim.CurrentTime) { return nil }

	// Switch on Seal IDs to return the right judgement
	switch currentSeal.ID {
	case SealOfBloodAuraID:
		return paladin.NewJudgementOfBlood(sim, target)
	case SealOfTheCrusaderAuraID:
		return paladin.NewJudgementOfTheCrusader(sim, target)
	default:
		return nil // case if for some reason judgement type isn't programmed yet
	}
}
