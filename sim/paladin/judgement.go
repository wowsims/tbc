package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const JudgementManaCost = 147.0
const JudgementCDTime = time.Second * 10
const JudgementDuration = time.Second * 20

// Shared conditions required to be able to cast any Judgement.
func (paladin *Paladin) canJudgement(sim *core.Simulation) bool {
	return paladin.currentSealExpires > sim.CurrentTime && !paladin.IsOnCD(JudgementCD, sim.CurrentTime)
}

var JudgementCD = core.NewCooldownID()
var JudgementOfBloodActionID = core.ActionID{SpellID: 31898, CooldownID: JudgementCD}

// refactored Judgement of Blood as an ActiveMeleeAbility which is most similar to actual behavior with a typical ret build
// but still has a few differences (differences are: does not scale off spell power, cannot be partially resisted, can be missed or dodged)
func (paladin *Paladin) registerJudgementOfBloodSpell(sim *core.Simulation) {
	job := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    JudgementOfBloodActionID,
				Character:   &paladin.Character,
				SpellSchool: core.SpellSchoolHoly,
				SpellExtras: core.SpellExtrasAlwaysHits,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategorySpecial,
			CritMultiplier:      paladin.DefaultMeleeCritMultiplier(),
			CritRollCategory:    core.CritRollCategoryPhysical,
			ProcMask:            core.ProcMaskMeleeOrRangedSpecial,
			DamageMultiplier:    1, // Need to review to make sure I set these properly
			ThreatMultiplier:    1,
			BaseDamage:          core.BaseDamageConfigMagic(295, 325, 0.429),
			OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
				paladin.sanctifiedJudgement(sim, paladin.sealOfBlood.Cost.Value)
				paladin.RemoveAura(sim, SealOfBloodAuraID)
				paladin.currentSealID = 0
				paladin.currentSealExpires = 0
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

	paladin.JudgementOfBlood = paladin.RegisterSpell(core.SpellConfig{
		Template:   job,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

func (paladin *Paladin) CanJudgementOfBlood(sim *core.Simulation) bool {
	return paladin.canJudgement(sim) && paladin.currentSealID == SealOfBloodAuraID
}

var JudgementOfTheCrusaderActionID = core.ActionID{SpellID: 27159, CooldownID: JudgementCD}

func (paladin *Paladin) registerJudgementOfTheCrusaderSpell(sim *core.Simulation) {
	jotc := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    JudgementOfTheCrusaderActionID,
				Character:   &paladin.Character,
				SpellSchool: core.SpellSchoolHoly,
				SpellExtras: core.SpellExtrasAlwaysHits,
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
					paladin.currentSealID = 0
					paladin.currentSealExpires = 0
				},
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}
				aura := core.JudgementOfTheCrusaderAura(spellEffect.Target, float64(paladin.Talents.ImprovedSealOfTheCrusader))
				spellEffect.Target.AddAura(sim, aura)
				paladin.currentJudgementID = aura.ID
				paladin.currentJudgementExpires = sim.CurrentTime + JudgementDuration
			},
		},
	}

	// Reduce mana cost if we have Benediction Talent
	jotc.Cost.Value = JudgementManaCost * (1 - 0.03*float64(paladin.Talents.Benediction))

	// Reduce CD if we have Improved Judgement Talent
	jotc.Cooldown = JudgementCDTime - (time.Second * time.Duration(paladin.Talents.ImprovedJudgement))

	paladin.JudgementOfTheCrusader = paladin.RegisterSpell(core.SpellConfig{
		Template:   jotc,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

func (paladin *Paladin) CanJudgementOfTheCrusader(sim *core.Simulation) bool {
	return paladin.canJudgement(sim) && paladin.currentSealID == SealOfTheCrusaderAuraID
}

var JudgementOfWisdomActionID = core.ActionID{SpellID: 27164, CooldownID: JudgementCD}

func (paladin *Paladin) registerJudgementOfWisdomSpell(sim *core.Simulation) {
	jow := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    JudgementOfWisdomActionID,
				Character:   &paladin.Character,
				SpellSchool: core.SpellSchoolHoly,
				SpellExtras: core.SpellExtrasAlwaysHits,
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
					paladin.currentSealID = 0
					paladin.currentSealExpires = 0
				},
			},
		},
		Effect: core.SpellEffect{
			CritRollCategory:    core.CritRollCategoryMagical,
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}
				aura := core.JudgementOfWisdomAura()
				spellEffect.Target.AddAura(sim, aura)
				paladin.currentJudgementID = aura.ID
				paladin.currentJudgementExpires = sim.CurrentTime + JudgementDuration
			},
		},
	}

	// Reduce mana cost if we have Benediction Talent
	jow.Cost.Value = JudgementManaCost * (1 - 0.03*float64(paladin.Talents.Benediction))

	// Reduce CD if we have Improved Judgement Talent
	jow.Cooldown = JudgementCDTime - (time.Second * time.Duration(paladin.Talents.ImprovedJudgement))

	paladin.JudgementOfWisdom = paladin.RegisterSpell(core.SpellConfig{
		Template:   jow,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

func (paladin *Paladin) CanJudgementOfWisdom(sim *core.Simulation) bool {
	return paladin.canJudgement(sim) && paladin.currentSealID == SealOfWisdomAuraID
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
