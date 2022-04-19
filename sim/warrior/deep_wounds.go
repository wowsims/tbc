package warrior

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var DeepWoundsActionID = core.ActionID{SpellID: 12867}

func (warrior *Warrior) applyDeepWounds() {
	if warrior.Talents.DeepWounds == 0 {
		return
	}

	deepWoundsSpell := warrior.RegisterSpell(core.SpellConfig{
		ActionID:    DeepWoundsActionID,
		SpellSchool: core.SpellSchoolPhysical,
	})

	var dwDots []*core.Dot

	warrior.RegisterAura(core.Aura{
		Label:    "Deep Wounds",
		ActionID: DeepWoundsActionID,
		Duration: core.NeverExpires,
		OnInit: func(aura *core.Aura, sim *core.Simulation) {
			dwDots = nil
			tickDamage := warrior.AutoAttacks.MH.AverageDamage()
			for i := int32(0); i < sim.GetNumTargets(); i++ {
				target := sim.GetTarget(i)
				dotAura := target.RegisterAura(core.Aura{
					Label:    "DeepWounds-" + strconv.Itoa(int(warrior.Index)),
					ActionID: DeepWoundsActionID,
					Duration: time.Second * 12,
				})
				dot := core.NewDot(core.Dot{
					Spell:         deepWoundsSpell,
					Aura:          dotAura,
					NumberOfTicks: 4,
					TickLength:    time.Second * 3,
					TickEffects: core.TickFuncApplyEffects(core.ApplyEffectFuncDirectDamage(core.SpellEffect{
						DamageMultiplier: 0.2 * float64(warrior.Talents.DeepWounds),
						ThreatMultiplier: 1,
						IsPeriodic:       true,
						IsPhantom:        true,
						BaseDamage:       core.BaseDamageConfigFlat(tickDamage),
						OutcomeApplier:   core.OutcomeFuncTick(),
					})),
				})
				dwDots = append(dwDots, dot)
			}
		},
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Outcome.Matches(core.OutcomeCrit) {
				dwDots[spellEffect.Target.Index].Apply(sim)
				warrior.procBloodFrenzy(sim, spellEffect, time.Second*12)
			}
		},
	})
}
