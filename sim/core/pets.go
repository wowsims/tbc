package core

// Pets used by core effects/buffs/consumes.

import (
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

var GnomishFlameTurretActionID = ActionID{ItemID: 23841}

func (character *Character) newGnomishFlameTurretCaster() func(sim *Simulation) {
	gft := character.NewGnomishFlameTurret()

	castTemplate := SimpleCast{
		Cast: Cast{
			ActionID:  GnomishFlameTurretActionID,
			Character: character,
			OnCastComplete: func(sim *Simulation, cast *Cast) {
				gft.EnableWithTimeout(sim, gft, time.Second*45)
			},
		},
	}

	return func(sim *Simulation) {
		cast := castTemplate
		cast.Init(sim)
		cast.StartCast(sim)
	}
}

type GnomishFlameTurret struct {
	Pet

	flameCannon         SimpleSpell
	flameCannonTemplate SimpleSpellTemplate
}

func (character *Character) NewGnomishFlameTurret() *GnomishFlameTurret {
	gft := &GnomishFlameTurret{
		Pet: NewPet(
			"Gnomish Flame Turret",
			character,
			stats.Stats{
				stats.SpellCrit: 1 * SpellCritRatingPerCritChance,
			},
			func(ownerStats stats.Stats) stats.Stats {
				return stats.Stats{}
			},
			false,
		),
	}

	character.AddPet(gft)

	return gft
}

func (gft *GnomishFlameTurret) GetPet() *Pet {
	return &gft.Pet
}

func (gft *GnomishFlameTurret) Init(sim *Simulation) {
	gft.flameCannonTemplate = gft.newFlameCannonTemplate(sim)
}

func (gft *GnomishFlameTurret) Reset(sim *Simulation) {
}

func (gft *GnomishFlameTurret) OnGCDReady(sim *Simulation) {
	gft.NewFlameCannon(sim, sim.GetPrimaryTarget()).Cast(sim)
}

const SpellIDFlameCannon int32 = 30527

func (gft *GnomishFlameTurret) newFlameCannonTemplate(sim *Simulation) SimpleSpellTemplate {
	spell := SimpleSpell{
		SpellCast: SpellCast{
			Cast: Cast{
				ActionID:            ActionID{SpellID: SpellIDFlameCannon},
				Character:           &gft.Character,
				CritRollCategory:    CritRollCategoryMagical,
				OutcomeRollCategory: OutcomeRollCategoryMagic,
				SpellSchool:         SpellSchoolFire,
				// Pretty sure this works the same way as Searing Totem, where the next shot
				// fires once the previous missile has hit the target. Just give some static
				// value for now.
				GCD:            time.Millisecond * 800,
				CritMultiplier: gft.DefaultSpellCritMultiplier(),
				IgnoreHaste:    true,
			},
		},
		Effect: SpellHitEffect{
			SpellEffect: SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			DirectInput: DirectDamageInput{
				MinBaseDamage: 31,
				MaxBaseDamage: 36,
			},
		},
	}

	return NewSimpleSpellTemplate(spell)
}

func (gft *GnomishFlameTurret) NewFlameCannon(sim *Simulation, target *Target) *SimpleSpell {
	// Initialize cast from precomputed template.
	fc := &gft.flameCannon
	gft.flameCannonTemplate.Apply(fc)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	fc.Effect.Target = target
	fc.Init(sim)

	return fc
}
