package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Encounter struct {
	Duration   float64
	Targets    []*Target
}

func NewEncounter(options proto.Encounter) Encounter {
	encounter := Encounter{
		Duration: options.Duration,
		Targets: []*Target{},
	}

	for _, targetOptions := range options.Targets {
		target := NewTarget(*targetOptions)
		encounter.Targets = append(encounter.Targets, target)
	}

	return encounter
}

func (encounter *Encounter) Finalize() {
	for _, target := range encounter.Targets {
		target.Finalize()
	}
}

// Target is an enemy that can be the target of attacks/spells.
// Currently targets are basically just lvl 73 target dummies.
type Target struct {
	armor int32

	// Provides aura tracking behavior. Targets need auras to handle debuffs.
	auraTracker

	// Whether Finalize() has been called yet for this Character.
	// All fields above this may not be altered once finalized is set.
	finalized bool
}

func NewTarget(options proto.Target) *Target {
	target := &Target{
		armor: options.Armor,
		auraTracker: newAuraTracker(),
	}
	// TODO: Do something with this
	target.auraTracker.playerID = -1

	if options.Debuffs != nil {
		applyDebuffEffects(target, *options.Debuffs)
	}

	return target
}

func applyDebuffEffects(target *Target, debuffs proto.Debuffs) {
	if debuffs.Misery {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return miseryAura()
		})
	}

	if debuffs.JudgementOfWisdom {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return auraJudgementOfWisdom()
		})
	}

	if debuffs.ImprovedSealOfTheCrusader {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return Aura{
				ID: MagicIDImprovedSealOfTheCrusader,
				OnBeforeSpellHit: func(sim *Simulation, hitInput *DirectCastDamageInput) {
					hitInput.BonusCrit += 3
					// FUTURE: melee crit bonus, research actual value
				},
			}
		})
	}
}

func miseryAura() Aura {
	return Aura{
		ID:      MagicIDMisery,
		OnSpellHit: func(sim *Simulation, cast DirectCastAction, result *DirectCastDamageResult) {
			result.Damage *= 1.05
		},
	}
}

func auraJudgementOfWisdom() Aura {
	const mana = 74 / 2 // 50% proc
	return Aura{
		ID:      MagicIDJoW,
		OnSpellHit: func(sim *Simulation, cast DirectCastAction, result *DirectCastDamageResult) {
			if cast.GetActionID().ItemID == ItemIDTheLightningCapacitor {
				return // TLC cant proc JoW
			}

			character := cast.GetCharacter()
			// Only apply to agents that have mana.
			if character.MaxMana() > 0 {
				character.AddStat(stats.Mana, mana)
				if sim.Log != nil {
					sim.Log("(%d) +Judgement Of Wisdom: 37 mana (74 @ 50%% proc)\n", character.ID)
				}
			}
		},
	}
}

func (target *Target) Finalize() {
	if target.finalized {
		return
	}
	target.finalized = true

	target.auraTracker.finalize()
}

func (target *Target) Reset(sim *Simulation) {
	target.auraTracker.reset(sim)
}

func (target *Target) Advance(sim *Simulation, elapsedTime time.Duration) {
	target.auraTracker.advance(sim, elapsedTime)
}
