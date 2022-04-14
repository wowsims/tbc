package rogue

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

// Returns whether any Deadly Poisons are being used.
func (rogue *Rogue) applyPoisons() {
	hasWFTotem := rogue.HasWFTotem
	rogue.applyDeadlyPoison(hasWFTotem)
	rogue.applyInstantPoison(hasWFTotem)
}

var DeadlyPoisonActionID = core.ActionID{SpellID: 27186}

func (rogue *Rogue) registerDeadlyPoisonSpell(sim *core.Simulation) {
	rogue.DeadlyPoison = rogue.RegisterSpell(core.SpellConfig{
		ActionID:    DeadlyPoisonActionID,
		SpellSchool: core.SpellSchoolNature,

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			BonusSpellHitRating: 5 * core.SpellHitRatingPerHitChance * float64(rogue.Talents.MasterPoisoner),
			ThreatMultiplier:    1,
			IsPhantom:           true,
			OutcomeApplier:      core.OutcomeFuncMagicHit(),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					if rogue.DeadlyPoisonDot.IsActive() {
						rogue.DeadlyPoisonDot.Refresh(sim)
						rogue.DeadlyPoisonDot.AddStack(sim)
					} else {
						rogue.DeadlyPoisonDot.Apply(sim)
						rogue.DeadlyPoisonDot.SetStacks(sim, 1)
					}
				}
			},
		}),
	})

	target := sim.GetPrimaryTarget()
	dotAura := target.RegisterAura(core.Aura{
		Label:     "DeadlyPoison-" + strconv.Itoa(int(rogue.Index)),
		ActionID:  DeadlyPoisonActionID,
		MaxStacks: 5,
		Duration:  time.Second * 12,
	})
	rogue.DeadlyPoisonDot = core.NewDot(core.Dot{
		Spell:         rogue.DeadlyPoison,
		Aura:          dotAura,
		NumberOfTicks: 4,
		TickLength:    time.Second * 3,
		TickEffects: core.TickFuncApplyEffects(core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			DamageMultiplier: 1 + 0.04*float64(rogue.Talents.VilePoisons),
			ThreatMultiplier: 1,
			IsPeriodic:       true,
			IsPhantom:        true,
			BaseDamage:       core.MultiplyByStacks(core.BaseDamageConfigFlat(180/4), dotAura),
			OutcomeApplier:   core.OutcomeFuncTick(),
		})),
	})
}

func (rogue *Rogue) applyDeadlyPoison(hasWFTotem bool) {
	procMask := core.GetMeleeProcMaskForHands(
		!hasWFTotem && rogue.Consumes.MainHandImbue == proto.WeaponImbue_WeaponImbueRogueDeadlyPoison,
		rogue.Consumes.OffHandImbue == proto.WeaponImbue_WeaponImbueRogueDeadlyPoison)

	if procMask == core.ProcMaskEmpty {
		return
	}

	procChance := 0.3 + 0.02*float64(rogue.Talents.ImprovedPoisons)

	rogue.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return rogue.GetOrRegisterAura(core.Aura{
			Label: "Deadly Poison",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) || spellEffect.IsPhantom {
					return
				}
				if sim.RandomFloat("Deadly Poison") > procChance {
					return
				}

				rogue.DeadlyPoison.Cast(sim, spellEffect.Target)
			},
		})
	})
}

func (rogue *Rogue) registerInstantPoisonSpell(_ *core.Simulation) {
	rogue.InstantPoison = rogue.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 26891},
		SpellSchool: core.SpellSchoolNature,

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			IsPhantom:           true,
			DamageMultiplier:    1 + 0.04*float64(rogue.Talents.VilePoisons),
			ThreatMultiplier:    1,
			BonusSpellHitRating: 5 * core.SpellHitRatingPerHitChance * float64(rogue.Talents.MasterPoisoner),
			BaseDamage:          core.BaseDamageConfigRoll(146, 194),
			OutcomeApplier:      core.OutcomeFuncMagicHitAndCrit(rogue.DefaultSpellCritMultiplier()),
		}),
	})
}

func (rogue *Rogue) applyInstantPoison(hasWFTotem bool) {
	procMask := core.GetMeleeProcMaskForHands(
		!hasWFTotem && rogue.Consumes.MainHandImbue == proto.WeaponImbue_WeaponImbueRogueInstantPoison,
		rogue.Consumes.OffHandImbue == proto.WeaponImbue_WeaponImbueRogueInstantPoison)

	if procMask == core.ProcMaskEmpty {
		return
	}

	procChance := 0.2 + 0.02*float64(rogue.Talents.ImprovedPoisons)

	rogue.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return rogue.GetOrRegisterAura(core.Aura{
			Label: "Instant Poison",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) || spellEffect.IsPhantom {
					return
				}
				if sim.RandomFloat("Instant Poison") > procChance {
					return
				}

				rogue.procInstantPoison(sim, spellEffect)
			},
		})
	})
}

func (rogue *Rogue) procInstantPoison(sim *core.Simulation, spellEffect *core.SpellEffect) {
	rogue.InstantPoison.Cast(sim, spellEffect.Target)
}
