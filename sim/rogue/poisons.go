package rogue

import (
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

var DeadlyPoisonAuraID = core.NewAuraID()
var DeadlyPoisonDebuffID = core.NewDebuffID()

func (rogue *Rogue) registerDeadlyPoisonSpell(_ *core.Simulation) {
	cast := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: 27186},
				Character:   rogue.GetCharacter(),
				SpellSchool: core.SpellSchoolNature,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			IsPhantom:           true,
			DamageMultiplier:    1 + 0.04*float64(rogue.Talents.VilePoisons),
			ThreatMultiplier:    1,
			BonusSpellHitRating: 5 * core.SpellHitRatingPerHitChance * float64(rogue.Talents.MasterPoisoner),
			DotInput: core.DotDamageInput{
				NumberOfTicks:  4,
				TickLength:     time.Second * 3,
				TickBaseDamage: core.DotSnapshotFuncMagic(180/4, 0),
				DebuffID:       DeadlyPoisonDebuffID,
			},
		},
	}
	rogue.DeadlyPoison = rogue.RegisterSpell(core.SpellConfig{
		Template:   cast,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

func (rogue *Rogue) registerDeadlyPoisonRefreshSpell(_ *core.Simulation) {
	cast := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: 27186},
				Character:   rogue.GetCharacter(),
				SpellSchool: core.SpellSchoolNature,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryNone,
			IsPhantom:           true,
			DamageMultiplier:    1 + 0.04*float64(rogue.Talents.VilePoisons),
			ThreatMultiplier:    1,
			BonusSpellHitRating: 5 * core.SpellHitRatingPerHitChance * float64(rogue.Talents.MasterPoisoner),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}

				const tickDamagePerStack = 180.0 / 4.0
				rogue.deadlyPoisonStacks = core.MinInt(rogue.deadlyPoisonStacks+1, 5)
				rogue.DeadlyPoison.Instance.Effect.DotInput.SetTickDamage(tickDamagePerStack * float64(rogue.deadlyPoisonStacks))
				rogue.DeadlyPoison.Instance.Effect.DotInput.RefreshDot(sim)
			},
		},
	}
	rogue.DeadlyPoisonRefresh = rogue.RegisterSpell(core.SpellConfig{
		Template:   cast,
		ModifyCast: core.ModifyCastAssignTarget,
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

	rogue.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: DeadlyPoisonAuraID,
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) || spellEffect.IsPhantom {
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
	if rogue.DeadlyPoison.Instance.IsInUse() {
		rogue.DeadlyPoisonRefresh.Cast(sim, spellEffect.Target)
	} else {
		rogue.DeadlyPoison.Cast(sim, spellEffect.Target)
		rogue.deadlyPoisonStacks = 1
	}
}

var InstantPoisonAuraID = core.NewAuraID()

func (rogue *Rogue) registerInstantPoisonSpell(_ *core.Simulation) {
	cast := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: 26891},
				Character:   rogue.GetCharacter(),
				SpellSchool: core.SpellSchoolNature,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      rogue.DefaultSpellCritMultiplier(),
			IsPhantom:           true,
			DamageMultiplier:    1 + 0.04*float64(rogue.Talents.VilePoisons),
			ThreatMultiplier:    1,
			BonusSpellHitRating: 5 * core.SpellHitRatingPerHitChance * float64(rogue.Talents.MasterPoisoner),
			BaseDamage:          core.BaseDamageConfigRoll(146, 194),
		},
	}

	rogue.InstantPoison = rogue.RegisterSpell(core.SpellConfig{
		Template:   cast,
		ModifyCast: core.ModifyCastAssignTarget,
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

	rogue.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: InstantPoisonAuraID,
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(procMask) || spellEffect.IsPhantom {
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
	rogue.InstantPoison.Cast(sim, spellEffect.Target)
}
