package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const JudgementManaCost = 147.0

var JudgementCD = core.NewCooldownID()
var JudgementOfBloodActionID = core.ActionID{SpellID: 31898, CooldownID: JudgementCD}

func (paladin *Paladin) newJudgementOfBloodTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	job := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:       JudgementOfBloodActionID,
				Character:      &paladin.Character,
				SpellSchool:    stats.HolySpellPower,
				Cooldown:       time.Second * 10,
				CritMultiplier: paladin.SpellCritMultiplier(1, 0.25), // no idea what the math is for judgment crits
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

	// Reduce mana cost if we have Benediction Talent
	job.ManaCost = JudgementManaCost * (1 - 0.03*float64(paladin.Talents.Benediction))
	job.BaseManaCost = JudgementManaCost * (1 - 0.03*float64(paladin.Talents.Benediction)) // Is it necessary to define both these lines?

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
				ActionID:     JudgementOfTheCrusaderActionID,
				Character:    &paladin.Character,
				SpellSchool:  stats.HolySpellPower,
				BaseManaCost: JudgementManaCost,
				ManaCost:     JudgementManaCost,
				Cooldown:     time.Second * 10,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				IgnoreHitCheck: true,
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					aura := core.JudgementOfTheCrusaderAura(sim, float64(paladin.Talents.ImprovedSealOfTheCrusader))
					spellEffect.Target.AddAura(sim, aura)
					paladin.currentJudgement = aura
				},
			},
		},
	}

	// Reduce mana cost if we have Benediction Talent
	jotc.ManaCost = JudgementManaCost * (1 - 0.03*float64(paladin.Talents.Benediction))

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
	// No seal has even been active, so we can't cast judgement
	if paladin.currentSeal.ID == 0 {
		return nil
	}

	// Most recent seal has expired so we can't cast judgement
	if paladin.currentSeal.Expires <= sim.CurrentTime {
		return nil
	}

	// Switch on Seal IDs to return the right judgement
	var spell *core.SimpleSpell
	switch paladin.currentSeal.ID {
	case SealOfBloodAuraID:
		spell = paladin.NewJudgementOfBlood(sim, target)
	case SealOfTheCrusaderAuraID:
		spell = paladin.NewJudgementOfTheCrusader(sim, target)
	default:
		return nil // case if for some reason judgement type isn't programmed yet
	}

	spell.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		paladin.RemoveAura(sim, paladin.currentSeal.ID)
		paladin.currentSeal = core.Aura{}
	}

	return spell
}
