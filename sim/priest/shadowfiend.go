package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

// https://web.archive.org/web/20071201221602/http://www.shadowpriest.com/viewtopic.php?t=7616

const SpellIDShadowfiend int32 = 34433

var ShadowfiendCD = core.NewCooldownID()

func (priest *Priest) newShadowfiendTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		Name:           "Shadowfiend",
		CritMultiplier: 1.5,
		SpellSchool:    stats.ShadowSpellPower,
		Character:      &priest.Character,
		BaseManaCost:   575,
		ManaCost:       575,
		CastTime:       0,
		ActionID: core.ActionID{
			SpellID:    SpellIDShadowfiend,
			CooldownID: ShadowfiendCD,
		},
	}

	// Dmg over 15 sec = shadow_dmg*.6 + 1191
	// just simulate 10 1.5s long ticks
	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
		},
		DotInput: core.DotDamageInput{
			NumberOfTicks:        10,
			TickLength:           time.Millisecond * 1500,
			TickBaseDamage:       1191 / 10,
			TickSpellCoefficient: 0.06,
			OnPeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage *float64) {
				s := stats.Stats{stats.Mana: *tickDamage * 2.5}
				if sim.Log != nil {
					sim.Log("Shadowfiend Regenerated %0f mana.\n", s[stats.Mana])
				}
				priest.AddStats(s)
			},
		},
	}

	effect.DotInput.NumberOfTicks += int(priest.Talents.ImprovedShadowWordPain) // extra tick per point
	priest.applyTalentsToShadowSpell(&baseCast, &effect)

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		SpellHitEffect: effect,
	})
}

func (priest *Priest) NewShadowfiend(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	mf := &priest.ShadowfiendSpell

	priest.shadowfiendTemplate.Apply(mf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mf.Target = target
	mf.Init(sim)

	return mf
}
