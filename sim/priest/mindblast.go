package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDMB int32 = 25375

var MBCooldownID = core.NewCooldownID()

func (priest *Priest) newMindBlastTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		Name:           "Mind Blast",
		CritMultiplier: 1.5,
		SpellSchool:    stats.ShadowSpellPower,
		Character:      &priest.Character,
		BaseManaCost:   450,
		ManaCost:       450,
		CastTime:       time.Millisecond * 1500,
		Cooldown:       time.Second * 8,
		ActionID: core.ActionID{
			SpellID:    SpellIDMB,
			CooldownID: MBCooldownID,
		},
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
		},
		DirectInput: core.DirectDamageSpellInput{
			MinBaseDamage:    711,
			MaxBaseDamage:    752,
			SpellCoefficient: 0.429,
		},
	}

	priest.applyTalentsToShadowSpell(&baseCast, &effect)

	baseCast.Cooldown -= time.Millisecond * 500 * time.Duration(priest.Talents.ImprovedMindBlast)

	if ItemSetAbsolution.CharacterHasSetBonus(&priest.Character, 4) { // Absolution 4p adds 10% damage
		effect.StaticDamageMultiplier *= 0.1
	}

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		SpellHitEffect: effect,
	})
}

func (priest *Priest) NewMindBlast(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	mf := &priest.mindblastSpell

	priest.mindblastCastTemplate.Apply(mf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mf.Target = target
	mf.Init(sim)

	return mf
}
