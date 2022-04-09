package mage

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDFireball int32 = 27070

var FireballActionID = core.ActionID{SpellID: SpellIDFireball}

func (mage *Mage) registerFireballSpell(sim *core.Simulation) {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    FireballActionID,
				Character:   &mage.Character,
				SpellSchool: core.SpellSchoolFire,
				SpellExtras: SpellFlagMage,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 425,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 425,
				},
				CastTime: time.Millisecond * 3500,
				GCD:      core.GCDDefault,
			},
		},
	}
	spell.CastTime -= time.Millisecond * 100 * time.Duration(mage.Talents.ImprovedFireball)
	spell.Cost.Value -= spell.BaseCost.Value * float64(mage.Talents.Pyromaniac) * 0.01
	spell.Cost.Value *= 1 - float64(mage.Talents.ElementalPrecision)*0.01

	mage.Fireball = mage.RegisterSpell(core.SpellConfig{
		Template: spell,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			BonusSpellHitRating: float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance,

			BonusSpellCritRating: 0 +
				float64(mage.Talents.CriticalMass)*2*core.SpellCritRatingPerCritChance +
				float64(mage.Talents.Pyromaniac)*1*core.SpellCritRatingPerCritChance,

			DamageMultiplier: mage.spellDamageMultiplier *
				(1 + 0.02*float64(mage.Talents.FirePower)) *
				core.TernaryFloat64(ItemSetTempestRegalia.CharacterHasSetBonus(&mage.Character, 4), 1.05, 1),

			ThreatMultiplier: 1 - 0.05*float64(mage.Talents.BurningSoul),

			BaseDamage:     core.BaseDamageConfigMagic(649, 821, 1.0+0.03*float64(mage.Talents.EmpoweredFireball)),
			OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower))),

			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					mage.FireballDot.Apply(sim)
				}
			},
		}),
	})

	target := sim.GetPrimaryTarget()
	mage.FireballDot = core.NewDot(core.Dot{
		Spell: mage.Fireball,
		Aura: target.RegisterAura(&core.Aura{
			Label:    "Fireball-" + strconv.Itoa(int(mage.Index)),
			ActionID: FireballActionID,
		}),
		NumberOfTicks: 4,
		TickLength:    time.Second * 2,
		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: mage.spellDamageMultiplier *
				(1 + 0.02*float64(mage.Talents.FirePower)) *
				core.TernaryFloat64(ItemSetTempestRegalia.CharacterHasSetBonus(&mage.Character, 4), 1.05, 1),

			ThreatMultiplier: 1 - 0.05*float64(mage.Talents.BurningSoul),

			BaseDamage:     core.BaseDamageConfigFlat(84 / 4),
			OutcomeApplier: core.OutcomeFuncTick(),
			IsPeriodic:     true,
		}),
	})
}
