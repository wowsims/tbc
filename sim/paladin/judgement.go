package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const JudgementManaCost = 147.0
const JudgementCDTime = time.Second * 10

var JudgementCD = core.NewCooldownID()
var JudgementOfBloodActionID = core.ActionID{SpellID: 31898, CooldownID: JudgementCD}

// refactored Judgement of Blood as an ActiveMeleeAbility which is most similar to actual behavior with a typical ret build
// but still has a few differences (differences are: does not scale off spell power, cannot be partially resisted, can be missed or dodged)
func (paladin *Paladin) newJudgementOfBloodTemplate(sim *core.Simulation) core.MeleeAbilityTemplate {
	job := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			ActionID:       JudgementOfBloodActionID,
			Character:      &paladin.Character,
			SpellSchool:    stats.HolySpellPower,
			CritMultiplier: paladin.DefaultMeleeCritMultiplier(),
			IsPhantom:      true,
		},
		Effect: core.AbilityHitEffect{
			AbilityEffect: core.AbilityEffect{
				DamageMultiplier:       1, // Need to review to make sure I set these properly
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
				IgnoreArmor:            true,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage: 295,
				MaxBaseDamage: 325,
			},
		},
		OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			paladin.sanctifiedJudgement(sim, paladin.sealOfBlood.Cost.Value)
			paladin.RemoveAura(sim, SealOfBloodAuraID)
			paladin.currentSeal = core.Aura{}
		},
	}
	// Reduce mana cost if we have Benediction Talent
	job.Cost = core.ResourceCost{
		Type:  stats.Mana,
		Value: JudgementManaCost * (1 - 0.03*float64(paladin.Talents.Benediction)),
	}

	// Reduce CD if we have Improved Judgement Talent
	job.Cooldown = JudgementCDTime - (time.Second * time.Duration(paladin.Talents.ImprovedJudgement))

	// Increase Judgement Crit Chance if we have Fanaticism talent
	job.Effect.BonusCritRating = 3 * core.MeleeCritRatingPerCritChance * float64(paladin.Talents.Fanaticism)

	return core.NewMeleeAbilityTemplate(job)
}

func (paladin *Paladin) NewJudgementOfBlood(sim *core.Simulation, target *core.Target) *core.ActiveMeleeAbility {
	// No seal has even been active, so we can't cast judgement
	if paladin.currentSeal.ID != SealOfBloodAuraID {
		return nil
	}

	// Most recent seal has expired so we can't cast judgement
	if paladin.currentSeal.Expires <= sim.CurrentTime {
		return nil
	}

	job := &paladin.judgementOfBloodSpell
	paladin.judgementOfBloodTemplate.Apply(job)

	job.Effect.Target = target

	return job
}

var JudgementOfTheCrusaderActionID = core.ActionID{SpellID: 27159, CooldownID: JudgementCD}

func (paladin *Paladin) newJudgementOfTheCrusaderTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	jotc := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    JudgementOfTheCrusaderActionID,
				Character:   &paladin.Character,
				SpellSchool: stats.HolySpellPower,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: JudgementManaCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: JudgementManaCost,
				},
				Cooldown: time.Second * 10,
				OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
					paladin.sanctifiedJudgement(sim, paladin.sealOfTheCrusader.Cost.Value)
					paladin.RemoveAura(sim, SealOfTheCrusaderAuraID)
					paladin.currentSeal = core.Aura{}
				},
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
	jotc.Cost.Value = JudgementManaCost * (1 - 0.03*float64(paladin.Talents.Benediction))

	// Reduce CD if we have Improved Judgement Talent
	jotc.Cooldown = JudgementCDTime - (time.Second * time.Duration(paladin.Talents.ImprovedJudgement))

	return core.NewSimpleSpellTemplate(jotc)
}

func (paladin *Paladin) NewJudgementOfTheCrusader(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// No seal has even been active, so we can't cast judgement
	if paladin.currentSeal.ID != SealOfTheCrusaderAuraID {
		return nil
	}

	// Most recent seal has expired so we can't cast judgement
	if paladin.currentSeal.Expires <= sim.CurrentTime {
		return nil
	}

	jotc := &paladin.judgementOfTheCrusaderSpell
	paladin.judgementOfTheCrusaderTemplate.Apply(jotc)

	jotc.Effect.Target = target
	jotc.Init(sim)

	return jotc
}

var SanctifiedJudgementActionID = core.ActionID{SpellID: 31930}

// Helper function to implement Sanctified Seals talent
func (paladin *Paladin) sanctifiedJudgement(sim *core.Simulation, mana float64) {
	if paladin.Talents.SanctifiedJudgement == 0 {
		return
	}

	var proc float64
	if paladin.Talents.SanctifiedJudgement == 3 {
		proc = 1
	} else {
		proc = 0.33 * float64(paladin.Talents.SanctifiedJudgement)
	}

	if sim.RandomFloat("Sanctified Judgement") < proc {
		paladin.AddMana(sim, 0.8*mana, SanctifiedJudgementActionID, false)
	}
}
