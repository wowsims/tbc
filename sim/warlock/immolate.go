package warlock

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDImmolate9 int32 = 27215

var Immolate9ActionID = core.ActionID{SpellID: SpellIDImmolate9}

func (warlock *Warlock) registerImmolateSpell(sim *core.Simulation) {
	bonusFlatDamage := 0.0
	baseCost := 445.0
	minBaseDamage := 332.0
	maxBaseDamage := 332.0
	spellCoefficient := 0.2

	effect := core.SpellEffect{
		BonusSpellCritRating: float64(warlock.Talents.Devastation)*1*core.SpellCritRatingPerCritChance +
			float64(warlock.Talents.Backlash)*1*core.SpellCritRatingPerCritChance,
		DamageMultiplier: 1 *
			(1 + (0.05 * float64(warlock.Talents.ImprovedImmolate))) *
			(1 + (0.02 * float64(warlock.Talents.Emberstorm))),
		ThreatMultiplier: 1 - 0.05*float64(warlock.Talents.DestructiveReach),
		BaseDamage:       core.BaseDamageConfigMagic(minBaseDamage+bonusFlatDamage, maxBaseDamage+bonusFlatDamage, spellCoefficient+0.04*float64(warlock.Talents.ShadowAndFlame)),
		OutcomeApplier:   core.OutcomeFuncMagicHitAndCrit(warlock.SpellCritMultiplier(1, core.TernaryFloat64(warlock.Talents.Ruin, 0, 1))),
		OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Landed() {
				warlock.ImmolateDot.Apply(sim)
			}
		},
	}

	warlock.Immolate = warlock.RegisterSpell(core.SpellConfig{
		ActionID:    Immolate9ActionID,
		SpellSchool: core.SpellSchoolFire,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost:     baseCost * (1 - 0.01*float64(warlock.Talents.Cataclysm)),
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond*2000 - (time.Millisecond * 100 * time.Duration(warlock.Talents.Bane)),
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(effect),
	})

	target := sim.GetPrimaryTarget()

	// DOT: 615 dmg over 15s (123 every 3 sec, mod 0.13)
	warlock.ImmolateDot = core.NewDot(core.Dot{
		Spell: warlock.Immolate,
		Aura: target.RegisterAura(core.Aura{
			Label:    "immolate-" + strconv.Itoa(int(warlock.Index)),
			ActionID: Immolate9ActionID,
		}),
		NumberOfTicks: 5 + core.TernaryInt(ItemSetVoidheartRaiment.CharacterHasSetBonus(&warlock.Character, 4), 1, 0), // voidheart 4p gives 1 extra tick
		TickLength:    time.Second * 3,
		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(615/5, 0.13),
			OutcomeApplier:   core.OutcomeFuncTick(),
			IsPeriodic:       true,
		}),
	})
}
