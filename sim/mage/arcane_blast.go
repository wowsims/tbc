package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDArcaneBlast int32 = 30451
const ArcaneBlastBaseManaCost = 195.0
const ArcaneBlastBaseCastTime = time.Millisecond * 2500

func (mage *Mage) newArcaneBlastTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	abAura := core.Aura{
		ID:       ArcaneBlastAuraID,
		ActionID: core.ActionID{SpellID: 36032},
		Duration: time.Second * 8,
		Stacks:   0,
		OnExpire: func(sim *core.Simulation) {
			// Reset the mana cost on expiration.
			if mage.arcaneBlastSpell.IsInUse() {
				mage.arcaneBlastSpell.Cost.Value = core.MaxFloat(0, mage.arcaneBlastSpell.Cost.Value-3.0*ArcaneBlastBaseManaCost*0.75)
				mage.arcaneBlastSpell.ActionID.Tag = 1
			}
		},
	}

	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: SpellIDArcaneBlast},
				Character:           &mage.Character,
				CritRollCategory:    core.CritRollCategoryMagical,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolArcane,
				SpellExtras:         SpellFlagMage,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: ArcaneBlastBaseManaCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: ArcaneBlastBaseManaCost,
				},
				CastTime:       ArcaneBlastBaseCastTime,
				GCD:            core.GCDDefault,
				CritMultiplier: mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower)),
				OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
					abAura.Stacks = core.MinInt32(3, mage.NumStacks(ArcaneBlastAuraID)+1)
					cast.Character.ReplaceAura(sim, abAura)
				},
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: mage.spellDamageMultiplier,
				ThreatMultiplier:       1 - 0.2*float64(mage.Talents.ArcaneSubtlety),
			},
			BaseDamage: core.BaseDamageFuncMagic(668, 772, 2.5/3.5),
		},
	}

	spell.Effect.BonusSpellHitRating += float64(mage.Talents.ArcaneFocus) * 2 * core.SpellHitRatingPerHitChance
	spell.Effect.BonusSpellCritRating += float64(mage.Talents.ArcaneImpact) * 2 * core.SpellCritRatingPerCritChance

	if mage.hasTristfal {
		spell.Effect.StaticDamageMultiplier *= 1.2
		spell.Cost.Value += 0.2 * ArcaneBlastBaseManaCost
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (mage *Mage) NewArcaneBlast(sim *core.Simulation, target *core.Target) (*core.SimpleSpell, int32) {
	// Initialize cast from precomputed template.
	arcaneBlast := &mage.arcaneBlastSpell
	mage.arcaneBlastCastTemplate.Apply(arcaneBlast)

	numStacks := mage.NumStacks(ArcaneBlastAuraID)
	arcaneBlast.CastTime -= time.Duration(numStacks) * time.Second / 3
	arcaneBlast.Cost.Value += float64(numStacks) * ArcaneBlastBaseManaCost * 0.75
	arcaneBlast.ActionID.Tag = numStacks + 1

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	arcaneBlast.Effect.Target = target
	arcaneBlast.Init(sim)

	return arcaneBlast, numStacks
}

var ArcaneBlastAuraID = core.NewAuraID()

// Whether Arcane Blast stacks will fall off before a new blast could finish casting.
func (mage *Mage) willDropArcaneBlastStacks(sim *core.Simulation, curArcaneBlast *core.SimpleSpell, numStacks int32) bool {
	remainingBuffTime := mage.RemainingAuraDuration(sim, ArcaneBlastAuraID)

	return numStacks == 0 || remainingBuffTime < curArcaneBlast.CastTime
}

// Determines whether we can spam arcane blast for the remainder of the encounter.
func (mage *Mage) canBlast(sim *core.Simulation, curArcaneBlast *core.SimpleSpell, numStacks int32, willDropStacks bool) bool {
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
	projectedRemainingMana -= curArcaneBlast.Cost.Value
	remainingDuration -= curArcaneBlast.CastTime
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
