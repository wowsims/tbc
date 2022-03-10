package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func (rogue *Rogue) applyPoisons() {
	hasWFTotem := rogue.HasWFTotem
	rogue.applyDeadlyPoison(hasWFTotem)
	rogue.applyInstantPoison(hasWFTotem)
}

var DeadlyPoisonAuraID = core.NewAuraID()
var DeadlyPoisonDebuffID = core.NewDebuffID()

func (rogue *Rogue) newDeadlyPoisonTemplate(_ *core.Simulation) core.SimpleSpellTemplate {
	cast := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: 27186},
				Character:           rogue.GetCharacter(),
				IsPhantom:           true,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolNature,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1 + 0.04*float64(rogue.Talents.VilePoisons),
				ThreatMultiplier:       1,
				BonusSpellHitRating:    5 * core.SpellHitRatingPerHitChance * float64(rogue.Talents.MasterPoisoner),
			},
			DotInput: core.DotDamageInput{
				NumberOfTicks:  4,
				TickLength:     time.Second * 3,
				TickBaseDamage: 180 / 4,
				DebuffID:       DeadlyPoisonDebuffID,
			},
		},
	}
	return core.NewSimpleSpellTemplate(cast)
}

func (rogue *Rogue) newDeadlyPoisonRefreshTemplate(_ *core.Simulation) core.SimpleSpellTemplate {
	cast := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: 27186},
				Character:           rogue.GetCharacter(),
				IsPhantom:           true,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				CritRollCategory:    core.CritRollCategoryNone,
				SpellSchool:         core.SpellSchoolNature,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1 + 0.04*float64(rogue.Talents.VilePoisons),
				ThreatMultiplier:       1,
				BonusSpellHitRating:    5 * core.SpellHitRatingPerHitChance * float64(rogue.Talents.MasterPoisoner),
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					if !spellEffect.Landed() {
						return
					}

					const tickDamagePerStack = 180.0 / 4.0
					rogue.deadlyPoisonStacks = core.MinInt(rogue.deadlyPoisonStacks+1, 5)
					rogue.deadlyPoison.Effect.DotInput.SetTickDamage(tickDamagePerStack * float64(rogue.deadlyPoisonStacks))
					rogue.deadlyPoison.Effect.DotInput.RefreshDot()
				},
			},
		},
	}
	return core.NewSimpleSpellTemplate(cast)
}

func (rogue *Rogue) applyDeadlyPoison(hasWFTotem bool) {
	procMask := core.GetMeleeProcMaskForHands(
		!hasWFTotem && rogue.Consumes.MainHandImbue == proto.WeaponImbue_WeaponImbueRogueDeadlyPoison,
		rogue.Consumes.OffHandImbue == proto.WeaponImbue_WeaponImbueRogueDeadlyPoison)

	if procMask == core.ProcMaskEmpty {
		return
	}

	procChance := 0.3 + 0.02*float64(rogue.Talents.ImprovedPoisons)

	rogue.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: DeadlyPoisonAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) || spellCast.IsPhantom {
					return
				}
				if sim.RandomFloat("Deadly Poison") > procChance {
					return
				}

				rogue.procDeadlyPoison(sim, spellEffect)
			},
		}
	})
}

func (rogue *Rogue) procDeadlyPoison(sim *core.Simulation, spellEffect *core.SpellEffect) {
	if rogue.deadlyPoison.IsInUse() {
		dp := &rogue.deadlyPoisonRefresh
		rogue.deadlyPoisonRefreshTemplate.Apply(dp)
		dp.Effect.Target = spellEffect.Target
		dp.Init(sim)
		dp.Cast(sim)
	} else {
		dp := &rogue.deadlyPoison
		rogue.deadlyPoisonTemplate.Apply(dp)
		dp.Effect.Target = spellEffect.Target
		dp.Init(sim)
		dp.Cast(sim)
		rogue.deadlyPoisonStacks = 1
	}

}

var InstantPoisonAuraID = core.NewAuraID()

func (rogue *Rogue) newInstantPoisonTemplate(_ *core.Simulation) core.SimpleSpellTemplate {
	cast := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: 26891},
				Character:           rogue.GetCharacter(),
				IsPhantom:           true,
				CritRollCategory:    core.CritRollCategoryMagical,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolNature,
				CritMultiplier:      rogue.DefaultSpellCritMultiplier(),
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1 + 0.04*float64(rogue.Talents.VilePoisons),
				ThreatMultiplier:       1,
				BonusSpellHitRating:    5 * core.SpellHitRatingPerHitChance * float64(rogue.Talents.MasterPoisoner),
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage: 146,
				MaxBaseDamage: 194,
			},
		},
	}
	return core.NewSimpleSpellTemplate(cast)
}

func (rogue *Rogue) applyInstantPoison(hasWFTotem bool) {
	procMask := core.GetMeleeProcMaskForHands(
		!hasWFTotem && rogue.Consumes.MainHandImbue == proto.WeaponImbue_WeaponImbueRogueInstantPoison,
		rogue.Consumes.OffHandImbue == proto.WeaponImbue_WeaponImbueRogueInstantPoison)

	if procMask == core.ProcMaskEmpty {
		return
	}

	procChance := 0.2 + 0.02*float64(rogue.Talents.ImprovedPoisons)

	rogue.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: InstantPoisonAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) || spellCast.IsPhantom {
					return
				}
				if sim.RandomFloat("Instant Poison") > procChance {
					return
				}

				rogue.procInstantPoison(sim, spellEffect)
			},
		}
	})
}

func (rogue *Rogue) procInstantPoison(sim *core.Simulation, spellEffect *core.SpellEffect) {
	rogue.instantPoisonTemplate.Apply(&rogue.instantPoison)
	rogue.instantPoison.Effect.Target = spellEffect.Target
	rogue.instantPoison.Init(sim)
	rogue.instantPoison.Cast(sim)
}
