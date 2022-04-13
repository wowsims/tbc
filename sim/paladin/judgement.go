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
	var loaAura *core.Aura
	if paladin.Equip[proto.ItemSlot_ItemSlotRanged].ID == 27484 {
		loaAura = paladin.NewTemporaryStatsAura(
			"Libram of Avengement",
			LibramOfAvengementActionID,
			stats.Stats{stats.MeleeCrit: 53},
			time.Second*5)
	}

	effect := core.SpellEffect{
		ProcMask: core.ProcMaskMeleeOrRangedSpecial,

		BonusCritRating:  3 * core.MeleeCritRatingPerCritChance * float64(paladin.Talents.Fanaticism),
		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		BaseDamage:     core.BaseDamageConfigMagic(295, 325, 0.429),
		OutcomeApplier: core.OutcomeFuncMeleeSpecialHitAndCrit(paladin.DefaultMeleeCritMultiplier()),

		OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			paladin.sanctifiedJudgement(sim, paladin.SealOfBlood.DefaultCast.Cost)
			paladin.SealOfBloodAura.Deactivate(sim)
			if loaAura != nil {
				loaAura.Activate(sim)
			}
		},
	}
	paladin.applyTwoHandedWeaponSpecializationToSpell(&effect)

	paladin.JudgementOfBlood = paladin.RegisterSpell(core.SpellConfig{
		ActionID:    JudgementOfBloodActionID,
		SpellSchool: core.SpellSchoolHoly,

		ResourceType: stats.Mana,
		BaseCost:     JudgementManaCost,

		Cast: core.CastConfig{
			DefaultCast: core.NewCast{
				Cost: JudgementManaCost * (1 - 0.03*float64(paladin.Talents.Benediction)),
			},
			Cooldown: JudgementCDTime - (time.Second * time.Duration(paladin.Talents.ImprovedJudgement)),
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(effect),
	})
}

func (paladin *Paladin) CanJudgementOfBlood(sim *core.Simulation) bool {
	return paladin.canJudgement(sim) && paladin.CurrentSeal == paladin.SealOfBloodAura
}

var JudgementOfTheCrusaderActionID = core.ActionID{SpellID: 27159, CooldownID: JudgementCD}

func (paladin *Paladin) registerJudgementOfTheCrusaderSpell(sim *core.Simulation) {
	paladin.JudgementOfTheCrusaderAura = core.JudgementOfTheCrusaderAura(sim.GetPrimaryTarget(), paladin.Talents.ImprovedSealOfTheCrusader)

	paladin.JudgementOfTheCrusader = paladin.RegisterSpell(core.SpellConfig{
		ActionID:    JudgementOfTheCrusaderActionID,
		SpellSchool: core.SpellSchoolHoly,

		ResourceType: stats.Mana,
		BaseCost:     JudgementManaCost,

		Cast: core.CastConfig{
			DefaultCast: core.NewCast{
				Cost: JudgementManaCost * (1 - 0.03*float64(paladin.Talents.Benediction)),
			},
			Cooldown: JudgementCDTime - (time.Second * time.Duration(paladin.Talents.ImprovedJudgement)),
			OnCastComplete: func(sim *core.Simulation, spell *core.Spell) {
				paladin.sanctifiedJudgement(sim, paladin.SealOfTheCrusader.DefaultCast.Cost)
				paladin.SealOfTheCrusaderAura.Deactivate(sim)
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			OutcomeApplier: core.OutcomeFuncAlwaysHit(),

			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}
				paladin.JudgementOfTheCrusaderAura.Activate(sim)
				paladin.CurrentJudgement = paladin.JudgementOfTheCrusaderAura
			},
		}),
	})
}

func (paladin *Paladin) CanJudgementOfTheCrusader(sim *core.Simulation) bool {
	return paladin.canJudgement(sim) && paladin.CurrentSeal == paladin.SealOfTheCrusaderAura
}

var JudgementOfWisdomActionID = core.ActionID{SpellID: 27164, CooldownID: JudgementCD}

func (paladin *Paladin) registerJudgementOfWisdomSpell(sim *core.Simulation) {
	paladin.JudgementOfWisdomAura = core.JudgementOfWisdomAura(sim.GetPrimaryTarget())

	paladin.JudgementOfWisdom = paladin.RegisterSpell(core.SpellConfig{
		ActionID:    JudgementOfWisdomActionID,
		SpellSchool: core.SpellSchoolHoly,

		ResourceType: stats.Mana,
		BaseCost:     JudgementManaCost,

		Cast: core.CastConfig{
			DefaultCast: core.NewCast{
				Cost: JudgementManaCost * (1 - 0.03*float64(paladin.Talents.Benediction)),
			},
			Cooldown: JudgementCDTime - (time.Second * time.Duration(paladin.Talents.ImprovedJudgement)),
			OnCastComplete: func(sim *core.Simulation, spell *core.Spell) {
				paladin.sanctifiedJudgement(sim, paladin.SealOfWisdom.DefaultCast.Cost)
				paladin.SealOfWisdomAura.Deactivate(sim)
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			OutcomeApplier: core.OutcomeFuncMagicHit(),

			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}
				paladin.JudgementOfWisdomAura.Activate(sim)
				paladin.CurrentJudgement = paladin.JudgementOfWisdomAura
			},
		}),
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
