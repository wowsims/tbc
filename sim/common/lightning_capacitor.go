package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func init() {
	core.AddItemEffect(core.ItemIDTheLightningCapacitor, ApplyTheLightningCapacitor)
}

func ApplyTheLightningCapacitor(agent core.Agent) {
	character := agent.GetCharacter()

	tlcSpell := character.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{ItemID: core.ItemIDTheLightningCapacitor},
		SpellSchool: core.SpellSchoolNature,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			IsPhantom:        true, // TODO: replace with ProcMask
			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			BaseDamage:     core.BaseDamageConfigRoll(694, 807),
			OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(character.DefaultSpellCritMultiplier()),
		}),
	})

	var charges int
	icd := core.Cooldown{
		Timer:    character.NewTimer(),
		Duration: time.Millisecond * 2500,
	}

	character.RegisterAura(core.Aura{
		Label:    "Lightning Capacitor",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			charges = 0
			aura.Activate(sim)
		},
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !icd.IsReady(sim) {
				return
			}

			if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
				return
			}

			if !spellEffect.Outcome.Matches(core.OutcomeCrit) {
				return
			}

			charges++
			if charges >= 3 {
				icd.Use(sim)
				tlcSpell.Cast(sim, spellEffect.Target)
				charges = 0
			}
		},
	})
}
