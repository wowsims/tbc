package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const JudgementManaCost = 147.0
const JudgementCDTime = time.Second * 10
const JudgementDuration = time.Second * 20

// Shared conditions required to be able to cast any Judgement.
func (paladin *Paladin) canJudgement(sim *core.Simulation) bool {
	return paladin.CurrentSeal != nil && paladin.CurrentSeal.IsActive() && !paladin.IsOnCD(JudgementCD, sim.CurrentTime)
}

var JudgementCD = core.NewCooldownID()
var JudgementOfBloodActionID = core.ActionID{SpellID: 31898, CooldownID: JudgementCD}

var LibramOfAvengementActionID = core.ActionID{SpellID: 34260}

func (paladin *Paladin) registerJudgementOfBloodSpell(sim *core.Simulation) {
	loaIsEquipped := paladin.Equip[proto.ItemSlot_ItemSlotRanged].ID == 27484

	loaAura := paladin.RegisterAura(&core.Aura{
		Label:    "Libram of Avengement",
		ActionID: LibramOfAvengementActionID,
		Duration: time.Second * 5,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.BonusMeleeCritRating += 53
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.BonusMeleeCritRating -= 53
		},
	})

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
			DamageMultiplier:    1,
			ThreatMultiplier:    1,
			BaseDamage:          core.BaseDamageConfigMagic(295, 325, 0.429),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				paladin.sanctifiedJudgement(sim, paladin.sealOfBlood.Cost.Value)
				paladin.SealOfBloodAura.Deactivate(sim)
				if loaIsEquipped {
					loaAura.Activate(sim)
				}
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
	return paladin.canJudgement(sim) && paladin.CurrentSeal == paladin.SealOfBloodAura
}

var JudgementOfTheCrusaderActionID = core.ActionID{SpellID: 27159, CooldownID: JudgementCD}

func (paladin *Paladin) registerJudgementOfTheCrusaderSpell(sim *core.Simulation) {
	paladin.JudgementOfTheCrusaderAura = core.JudgementOfTheCrusaderAura(sim.GetPrimaryTarget(), paladin.Talents.ImprovedSealOfTheCrusader)

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
					paladin.SealOfTheCrusaderAura.Deactivate(sim)
				},
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}
				paladin.JudgementOfTheCrusaderAura.Activate(sim)
				paladin.CurrentJudgement = paladin.JudgementOfTheCrusaderAura
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
	return paladin.canJudgement(sim) && paladin.CurrentSeal == paladin.SealOfTheCrusaderAura
}

var JudgementOfWisdomActionID = core.ActionID{SpellID: 27164, CooldownID: JudgementCD}

func (paladin *Paladin) registerJudgementOfWisdomSpell(sim *core.Simulation) {
	paladin.JudgementOfWisdomAura = core.JudgementOfWisdomAura(sim.GetPrimaryTarget())

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
					paladin.SealOfWisdomAura.Deactivate(sim)
				},
			},
		},
		Effect: core.SpellEffect{
			CritRollCategory:    core.CritRollCategoryMagical,
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}
				paladin.JudgementOfWisdomAura.Activate(sim)
				paladin.CurrentJudgement = paladin.JudgementOfWisdomAura
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
	return paladin.canJudgement(sim) && paladin.CurrentSeal == paladin.SealOfWisdomAura
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
