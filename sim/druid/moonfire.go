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
		CritMultiplier: 1.5,
		SpellSchool:    stats.ArcaneSpellPower,
		Character:      &druid.Character,
		BaseManaCost:   495,
		ManaCost:       495,
		CastTime:       0,
		GCD:            core.GCDDefault,
		ActionID: core.ActionID{
			SpellID: SpellIDMoonfire,
		},
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

	if ItemSetThunderheart.CharacterHasSetBonus(&druid.Character, 2) { // Thunderheart 2p adds 1 extra tick to moonfire
		effect.DotInput.NumberOfTicks += 1
	}

	// Moonfire only talents
	effect.StaticDamageMultiplier *= 1 + 0.05*float64(druid.Talents.ImprovedMoonfire)
	effect.BonusSpellCritRating += float64(druid.Talents.ImprovedMoonfire) * 5 * core.SpellCritRatingPerCritChance

	// TODO: Shared talents
	baseCast.ManaCost -= baseCast.BaseManaCost * float64(druid.Talents.Moonglow) * 0.03
	effect.StaticDamageMultiplier *= 1 + 0.02*float64(druid.Talents.Moonfury)

	// Convert to percent, multiply by percent increase, convert back to multiplier by adding 1
	baseCast.CritMultiplier = (baseCast.CritMultiplier-1)*(1+float64(druid.Talents.Vengeance)*0.2) + 1

	// moonfire can proc the on hit but doesn't consume charges (i think)
	effect.OnSpellHit = druid.applyOnHitTalents

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		Effect: effect,
	})
}

// TODO: This might behave weird if we have a moonfire still ticking when we cast one.
//   We could do a check and if the spell is still ticking allocate a new SingleHitSpell?
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
