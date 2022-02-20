package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDMoonfire int32 = 26988

var MoonfireDebuffID = core.NewDebuffID()

func (druid *Druid) newMoonfireTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		ActionID:    core.ActionID{SpellID: SpellIDMoonfire},
		Character:   &druid.Character,
		SpellSchool: stats.ArcaneSpellPower,
		BaseCost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 495,
		},
		Cost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 495,
		},
		GCD:            core.GCDDefault,
		CritMultiplier: druid.SpellCritMultiplier(1, 0.2*float64(druid.Talents.Vengeance)),
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
			ThreatMultiplier:       1,
		},
		DirectInput: core.DirectDamageInput{
			MinBaseDamage:    305,
			MaxBaseDamage:    357,
			SpellCoefficient: 0.15,
		},
		DotInput: core.DotDamageInput{
			NumberOfTicks:        4,
			TickLength:           time.Second * 3,
			TickBaseDamage:       600 / 4,
			TickSpellCoefficient: 0.13,
			DebuffID:             MoonfireDebuffID,
		},
	}

	baseCast.Cost.Value -= baseCast.BaseCost.Value * 0.03 * float64(druid.Talents.Moonglow)

	effect.StaticDamageMultiplier *= 1 + 0.05*float64(druid.Talents.ImprovedMoonfire)
	effect.StaticDamageMultiplier *= 1 + 0.02*float64(druid.Talents.Moonfury)
	effect.BonusSpellCritRating += float64(druid.Talents.ImprovedMoonfire) * 5 * core.SpellCritRatingPerCritChance
	if ItemSetThunderheart.CharacterHasSetBonus(&druid.Character, 2) { // Thunderheart 2p adds 1 extra tick to moonfire
		effect.DotInput.NumberOfTicks += 1
	}

	// moonfire can proc the on hit but doesn't consume charges (i think)
	effect.OnSpellHit = druid.applyOnHitTalents

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		Effect: effect,
	})
}

func (druid *Druid) NewMoonfire(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	sf := &druid.MoonfireSpell
	druid.moonfireCastTemplate.Apply(sf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	sf.Effect.Target = target
	sf.Init(sim)

	return sf
}

func (druid *Druid) ShouldCastMoonfire(sim *core.Simulation, target *core.Target, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.Moonfire && !druid.MoonfireSpell.Effect.DotInput.IsTicking(sim)
}
