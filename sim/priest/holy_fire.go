package priest

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDHolyFire int32 = 25384

var HolyFireActionID = core.ActionID{SpellID: SpellIDHolyFire}

func (priest *Priest) registerHolyFireSpell(sim *core.Simulation) {
	baseCost := 290.0

	priest.HolyFire = priest.RegisterSpell(core.SpellConfig{
		ActionID:    HolyFireActionID,
		SpellSchool: core.SpellSchoolHoly,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost:     baseCost,
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond*3500 - time.Millisecond*100*time.Duration(priest.Talents.DivineFury),
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:             core.ProcMaskSpellDamage,
			BonusSpellCritRating: float64(priest.Talents.HolySpecialization) * 1 * core.SpellCritRatingPerCritChance,
			DamageMultiplier:     1 + 0.05*float64(priest.Talents.SearingLight),
			ThreatMultiplier:     1 - 0.04*float64(priest.Talents.SilentResolve),
			BaseDamage:           core.BaseDamageConfigMagic(426, 537, 0.8571),
			OutcomeApplier:       priest.OutcomeFuncMagicHitAndCrit(priest.DefaultSpellCritMultiplier()),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					priest.HolyFireDot.Apply(sim)
				}
			},
		}),
	})

	target := sim.GetPrimaryTarget()
	priest.HolyFireDot = core.NewDot(core.Dot{
		Spell: priest.HolyFire,
		Aura: target.RegisterAura(core.Aura{
			Label:    "HolyFire-" + strconv.Itoa(int(priest.Index)),
			ActionID: HolyFireActionID,
		}),
		NumberOfTicks: 5,
		TickLength:    time.Second * 2,
		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1 + 0.05*float64(priest.Talents.SearingLight),
			ThreatMultiplier: 1 - 0.04*float64(priest.Talents.SilentResolve),

			BaseDamage:     core.BaseDamageConfigMagicNoRoll(33, 0.17),
			OutcomeApplier: priest.OutcomeFuncTick(),
			IsPeriodic:     true,
		}),
	})
}
