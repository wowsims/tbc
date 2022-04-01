package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDArcaneBlast int32 = 30451
const ArcaneBlastBaseManaCost = 195.0
const ArcaneBlastBaseCastTime = time.Millisecond * 2500

func (mage *Mage) registerArcaneBlastSpell(sim *core.Simulation) {
	abAura := core.Aura{
		ID:       ArcaneBlastAuraID,
		ActionID: core.ActionID{SpellID: 36032},
		Duration: time.Second * 8,
		Stacks:   0,
		OnExpire: func(sim *core.Simulation) {
			// Reset the mana cost on expiration.
			if mage.ArcaneBlast.Instance.IsInUse() {
				mage.ArcaneBlast.Instance.Cost.Value = core.MaxFloat(0, mage.ArcaneBlast.Instance.Cost.Value-3.0*ArcaneBlastBaseManaCost*0.75)
				mage.ArcaneBlast.Instance.ActionID.Tag = 1
			}
		},
	}

	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: SpellIDArcaneBlast},
				Character:   &mage.Character,
				SpellSchool: core.SpellSchoolArcane,
				SpellExtras: SpellFlagMage,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: ArcaneBlastBaseManaCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: ArcaneBlastBaseManaCost,
				},
				CastTime: ArcaneBlastBaseCastTime,
				GCD:      core.GCDDefault,
				OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
					abAura.Stacks = core.MinInt32(3, mage.NumStacks(ArcaneBlastAuraID)+1)
					cast.Character.ReplaceAura(sim, abAura)
				},
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower)),
			DamageMultiplier:    mage.spellDamageMultiplier,
			ThreatMultiplier:    1 - 0.2*float64(mage.Talents.ArcaneSubtlety),
			BaseDamage:          core.BaseDamageConfigMagic(668, 772, 2.5/3.5),
		},
	}

	spell.Effect.BonusSpellHitRating += float64(mage.Talents.ArcaneFocus) * 2 * core.SpellHitRatingPerHitChance
	spell.Effect.BonusSpellCritRating += float64(mage.Talents.ArcaneImpact) * 2 * core.SpellCritRatingPerCritChance

	if mage.hasTristfal {
		spell.Effect.DamageMultiplier *= 1.2
		spell.Cost.Value += 0.2 * ArcaneBlastBaseManaCost
	}

	mage.ArcaneBlast = mage.RegisterSpell(core.SpellConfig{
		Template: spell,
		ModifyCast: func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
			numStacks := mage.NumStacks(ArcaneBlastAuraID)
			instance.CastTime -= time.Duration(numStacks) * time.Second / 3
			instance.Cost.Value += float64(numStacks) * ArcaneBlastBaseManaCost * 0.75
			instance.ActionID.Tag = numStacks + 1

			// Set dynamic fields, i.e. the stuff we couldn't precompute.
			instance.Effect.Target = target
		},
	})
}

func (mage *Mage) ArcaneBlastCastTime(numStacks int32) time.Duration {
	castTime := mage.ArcaneBlast.Template.CastTime
	castTime -= time.Duration(numStacks) * time.Second / 3
	castTime = time.Duration(float64(castTime) / mage.CastSpeed())
	return castTime
}

func (mage *Mage) ArcaneBlastManaCost(numStacks int32) float64 {
	cost := mage.ArcaneBlast.Template.Cost.Value
	cost += float64(numStacks) * ArcaneBlastBaseManaCost * 0.75
	mage.ArcaneBlast.Template.ApplyCostModifiers(&cost)
	return cost
}

var ArcaneBlastAuraID = core.NewAuraID()

// Whether Arcane Blast stacks will fall off before a new blast could finish casting.
func (mage *Mage) willDropArcaneBlastStacks(sim *core.Simulation, castTime time.Duration, numStacks int32) bool {
	remainingBuffTime := mage.RemainingAuraDuration(sim, ArcaneBlastAuraID)

	return numStacks == 0 || remainingBuffTime < castTime
}

// Determines whether we can spam arcane blast for the remainder of the encounter.
func (mage *Mage) canBlast(sim *core.Simulation, curManaCost float64, curCastTime time.Duration, numStacks int32, willDropStacks bool) bool {
	numStacksAfterFirstCast := numStacks + 1
	if willDropStacks {
		numStacksAfterFirstCast = 1
	}

	remainingDuration := sim.GetRemainingDuration()
	projectedRemainingMana := mage.manaTracker.ProjectedRemainingMana(sim, mage.GetCharacter())
	inverseCastSpeed := 1 / mage.CastSpeed()

	extraManaCost := 0.0
	if mage.hasTristfal {
		extraManaCost = 39
	}

	// First cast, which is curArcaneBlast
	projectedRemainingMana -= curManaCost
	remainingDuration -= curCastTime
	if projectedRemainingMana < 0 {
		return false
	} else if remainingDuration < 0 {
		return true
	}

	// Second cast
	if numStacksAfterFirstCast == 1 {
		projectedRemainingMana -= ArcaneBlastBaseManaCost + (1.0 * ArcaneBlastBaseManaCost * 0.75) + extraManaCost
		remainingDuration -= time.Duration(float64(ArcaneBlastBaseCastTime-(1*time.Second/3)) * inverseCastSpeed)
		if projectedRemainingMana < 0 {
			return false
		} else if remainingDuration < 0 {
			return true
		}
	}

	// Third cast
	if numStacksAfterFirstCast < 3 {
		projectedRemainingMana -= ArcaneBlastBaseManaCost + (2.0 * ArcaneBlastBaseManaCost * 0.75) + extraManaCost
		remainingDuration -= time.Duration(float64(ArcaneBlastBaseCastTime-(2*time.Second/3)) * inverseCastSpeed)
		if projectedRemainingMana < 0 {
			return false
		} else if remainingDuration < 0 {
			return true
		}
	}

	// Everything after this will be full stack blasts.
	manaCost := ArcaneBlastBaseManaCost + (3.0 * ArcaneBlastBaseManaCost * 0.75) + extraManaCost
	castTime := time.Duration(float64(ArcaneBlastBaseCastTime-(3*time.Second/3)) * inverseCastSpeed)
	numCasts := remainingDuration / castTime // time.Duration is an integer so we don't need to call math.Floor()
	totalManaCost := manaCost * float64(numCasts)

	clearcastProcChance := 0.02 * float64(mage.Talents.ArcaneConcentration)
	estimatedClearcastProcs := int(float64(numCasts) * clearcastProcChance)
	totalManaCost -= manaCost * float64(estimatedClearcastProcs)

	return totalManaCost < projectedRemainingMana
}
