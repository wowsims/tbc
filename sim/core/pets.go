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

	FlameCannon *Spell
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
	gft.registerFlameCannonSpell(sim)
}

func (gft *GnomishFlameTurret) Reset(sim *Simulation) {
}

func (gft *GnomishFlameTurret) OnGCDReady(sim *Simulation) {
	gft.FlameCannon.Cast(sim, sim.GetPrimaryTarget())
}

const SpellIDFlameCannon int32 = 30527

func (gft *GnomishFlameTurret) registerFlameCannonSpell(sim *Simulation) {
	spell := SimpleSpell{
		SpellCast: SpellCast{
			Cast: Cast{
				ActionID:    ActionID{SpellID: SpellIDFlameCannon},
				Character:   &gft.Character,
				SpellSchool: SpellSchoolFire,
				// Pretty sure this works the same way as Searing Totem, where the next shot
				// fires once the previous missile has hit the target. Just give some static
				// value for now.
				GCD:         time.Millisecond * 800,
				IgnoreHaste: true,
			},
		},
	}

	gft.FlameCannon = gft.RegisterSpell(SpellConfig{
		Template: spell,
		ApplyEffects: ApplyEffectFuncDirectDamage(SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			BaseDamage:     BaseDamageConfigRoll(31, 36),
			OutcomeApplier: OutcomeFuncMagicHitAndCrit(gft.DefaultSpellCritMultiplier()),
		}),
	})
}
