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
	return paladin.CurrentSeal != nil && paladin.CurrentSeal.IsActive() && paladin.JudgementOfWisdom.IsReady(sim)
}

func (paladin *Paladin) registerJudgementOfBloodSpell(cdTimer *core.Timer) {
	var loaAura *core.Aura
	if paladin.Equip[proto.ItemSlot_ItemSlotRanged].ID == 27484 {
		loaAura = paladin.NewTemporaryStatsAura(
			"Libram of Avengement",
			core.ActionID{SpellID: 34260},
			stats.Stats{stats.MeleeCrit: 53, stats.SpellCrit: 53},
			time.Second*5)
	}

	effect := core.SpellEffect{
		ProcMask: core.ProcMaskMeleeOrRangedSpecial,

		BonusCritRating:  3 * core.MeleeCritRatingPerCritChance * float64(paladin.Talents.Fanaticism),
		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		BaseDamage:     core.BaseDamageConfigMagic(295, 325, 0.429),
		OutcomeApplier: paladin.OutcomeFuncMeleeSpecialCritOnly(paladin.DefaultMeleeCritMultiplier()),

		OnSpellHitDealt: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			paladin.sanctifiedJudgement(sim, paladin.SealOfBlood.DefaultCast.Cost)
			paladin.SealOfBloodAura.Deactivate(sim)
			if loaAura != nil {
				loaAura.Activate(sim)
			}

			// Add mana from Spiritual Attunement
			// 33% of damage is self-inflicted, 10% of self-inflicted damage is returned as mana
			paladin.AddMana(sim, spellEffect.Damage*0.33*0.1, core.ActionID{SpellID: 33776}, false)
		},
	}
	paladin.applyTwoHandedWeaponSpecializationToSpell(&effect)

	baseCost := core.TernaryFloat64(ItemSetCrystalforgeBattlegear.CharacterHasSetBonus(&paladin.Character, 2), JudgementManaCost-35, JudgementManaCost)
	paladin.JudgementOfBlood = paladin.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 31898},
		SpellSchool: core.SpellSchoolHoly,

		ResourceType: stats.Mana,
		BaseCost:     JudgementManaCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost - JudgementManaCost*(0.03*float64(paladin.Talents.Benediction)),
			},
			CD: core.Cooldown{
				Timer:    cdTimer,
				Duration: JudgementCDTime - (time.Second * time.Duration(paladin.Talents.ImprovedJudgement)),
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(effect),
	})
}

func (paladin *Paladin) CanJudgementOfBlood(sim *core.Simulation) bool {
	return paladin.canJudgement(sim) && paladin.CurrentSeal == paladin.SealOfBloodAura
}

func (paladin *Paladin) registerJudgementOfTheCrusaderSpell(cdTimer *core.Timer) {
	percentBonus := 1.0
	if ItemSetJusticarBattlegear.CharacterHasSetBonus(&paladin.Character, 2) {
		percentBonus = 1.15
	}
	flatBonus := 0.0
	if paladin.Equip[proto.ItemSlot_ItemSlotRanged].ID == 23203 {
		flatBonus += 33.0
	} else if paladin.Equip[proto.ItemSlot_ItemSlotRanged].ID == 27949 {
		flatBonus += 47.0
	}
	paladin.JudgementOfTheCrusaderAura = core.JudgementOfTheCrusaderAura(paladin.Env.GetPrimaryTarget(), paladin.Talents.ImprovedSealOfTheCrusader, flatBonus, percentBonus)

	baseCost := core.TernaryFloat64(ItemSetCrystalforgeBattlegear.CharacterHasSetBonus(&paladin.Character, 2), JudgementManaCost-35, JudgementManaCost)
	paladin.JudgementOfTheCrusader = paladin.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 27159},
		SpellSchool: core.SpellSchoolHoly,

		ResourceType: stats.Mana,
		BaseCost:     JudgementManaCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost - JudgementManaCost*(0.03*float64(paladin.Talents.Benediction)),
			},
			CD: core.Cooldown{
				Timer:    cdTimer,
				Duration: JudgementCDTime - (time.Second * time.Duration(paladin.Talents.ImprovedJudgement)),
			},
			OnCastComplete: func(sim *core.Simulation, spell *core.Spell) {
				paladin.sanctifiedJudgement(sim, paladin.SealOfTheCrusader.DefaultCast.Cost)
				paladin.SealOfTheCrusaderAura.Deactivate(sim)
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			IsPhantom:      true,
			ProcMask:       core.ProcMaskEmpty,
			OutcomeApplier: paladin.OutcomeFuncAlwaysHit(),

			OnSpellHitDealt: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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

func (paladin *Paladin) registerJudgementOfWisdomSpell(cdTimer *core.Timer) {
	paladin.JudgementOfWisdomAura = core.JudgementOfWisdomAura(paladin.Env.GetPrimaryTarget())

	baseCost := core.TernaryFloat64(ItemSetCrystalforgeBattlegear.CharacterHasSetBonus(&paladin.Character, 2), JudgementManaCost-35, JudgementManaCost)
	paladin.JudgementOfWisdom = paladin.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 27164},
		SpellSchool: core.SpellSchoolHoly,

		ResourceType: stats.Mana,
		BaseCost:     JudgementManaCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost - JudgementManaCost*(0.03*float64(paladin.Talents.Benediction)),
			},
			CD: core.Cooldown{
				Timer:    cdTimer,
				Duration: JudgementCDTime - (time.Second * time.Duration(paladin.Talents.ImprovedJudgement)),
			},
			OnCastComplete: func(sim *core.Simulation, spell *core.Spell) {
				paladin.sanctifiedJudgement(sim, paladin.SealOfWisdom.DefaultCast.Cost)
				paladin.SealOfWisdomAura.Deactivate(sim)
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			OutcomeApplier: paladin.OutcomeFuncMagicHit(),

			OnSpellHitDealt: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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

// Defines judgement refresh behavior from attacks
// Returns extra mana if a different pally applied Judgement of Wisdom
func (paladin *Paladin) setupJudgementRefresh() {
	const mana = 74 / 2
	paladin.RegisterAura(core.Aura{
		Label:    "Refresh Judgement",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Landed() && spellEffect.ProcMask.Matches(core.ProcMaskMeleeWhiteHit) {
				if paladin.CurrentJudgement != nil && paladin.CurrentJudgement.IsActive() {
					// Refresh the judgement
					paladin.CurrentJudgement.Refresh(sim)

					// Check if current judgement is not JoW and also that JoW is on the target
					if paladin.CurrentJudgement != paladin.JudgementOfWisdomAura && paladin.JudgementOfWisdomAura.IsActive() {
						// Just trigger a second JoW
						paladin.AddMana(sim, mana, core.ActionID{SpellID: 27164}, false)
					}
				}
			}
		},
	})
}

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
		paladin.AddMana(sim, 0.8*mana, core.ActionID{SpellID: 31930}, false)
	}
}

func (paladin *Paladin) registerJudgements() {
	cdTimer := paladin.NewTimer()
	paladin.registerJudgementOfBloodSpell(cdTimer)
	paladin.registerJudgementOfTheCrusaderSpell(cdTimer)
	paladin.registerJudgementOfWisdomSpell(cdTimer)
}
