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
func (paladin *Paladin) newJudgementOfBloodTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	job := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            JudgementOfBloodActionID,
				Character:           &paladin.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolHoly,
				SpellExtras:         core.SpellExtrasAlwaysHits,
				CritMultiplier:      paladin.DefaultMeleeCritMultiplier(),
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               core.ProcMaskMeleeOrRangedSpecial,
				DamageMultiplier:       1, // Need to review to make sure I set these properly
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					paladin.sanctifiedJudgement(sim, paladin.sealOfBlood.Cost.Value)
					paladin.RemoveAura(sim, SealOfBloodAuraID)
					paladin.currentSeal = core.Aura{}
				},
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage:    295,
				MaxBaseDamage:    325,
				SpellCoefficient: 0.429,
			},
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

	return core.NewSimpleSpellTemplate(job)
}

func (paladin *Paladin) NewJudgementOfBlood(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
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
				ActionID:            JudgementOfTheCrusaderActionID,
				Character:           &paladin.Character,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolHoly,
				SpellExtras:         core.SpellExtrasAlwaysHits,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: JudgementManaCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: JudgementManaCost,
				},
				OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
					paladin.sanctifiedJudgement(sim, paladin.sealOfTheCrusader.Cost.Value)
					paladin.RemoveAura(sim, SealOfTheCrusaderAuraID)
					paladin.currentSeal = core.Aura{}
				},
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					if !spellEffect.Landed() {
						return
					}
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

var JudgementOfWisdomActionID = core.ActionID{SpellID: 27164, CooldownID: JudgementCD}

func (paladin *Paladin) newJudgementOfWisdomTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	jow := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            JudgementOfWisdomActionID,
				Character:           &paladin.Character,
				CritRollCategory:    core.CritRollCategoryMagical,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolHoly,
				SpellExtras:         core.SpellExtrasAlwaysHits,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: JudgementManaCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: JudgementManaCost,
				},
				OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
					paladin.sanctifiedJudgement(sim, paladin.sealOfWisdom.Cost.Value)
					paladin.RemoveAura(sim, SealOfWisdomAuraID)
					paladin.currentSeal = core.Aura{}
				},
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					if !spellEffect.Landed() {
						return
					}
					aura := core.JudgementOfWisdomAura(sim)
					spellEffect.Target.AddAura(sim, aura)
					paladin.currentJudgement = aura
				},
			},
		},
	}

	// Reduce mana cost if we have Benediction Talent
	jow.Cost.Value = JudgementManaCost * (1 - 0.03*float64(paladin.Talents.Benediction))

	// Reduce CD if we have Improved Judgement Talent
	jow.Cooldown = JudgementCDTime - (time.Second * time.Duration(paladin.Talents.ImprovedJudgement))

	return core.NewSimpleSpellTemplate(jow)
}

func (paladin *Paladin) NewJudgementOfWisdom(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// No seal has even been active, so we can't cast judgement
	if paladin.currentSeal.ID != SealOfWisdomAuraID {
		return nil
	}

	// Most recent seal has expired so we can't cast judgement
	if paladin.currentSeal.Expires <= sim.CurrentTime {
		return nil
	}

	jow := &paladin.judgementOfWisdomSpell
	paladin.judgementOfWisdomTemplate.Apply(jow)

	jow.Effect.Target = target
	jow.Init(sim)

	return jow
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
