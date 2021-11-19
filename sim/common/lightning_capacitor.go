package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddItemEffect(core.ItemIDTheLightningCapacitor, ApplyTheLightningCapacitor)
}

var TheLightningCapacitorAuraID = core.NewAuraID()

func ApplyTheLightningCapacitor(agent core.Agent) {
	spellObj := core.SimpleSpell{}

	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		castTemplate := newLightningCapacitorCastTemplate(sim, character)
		charges := 0

		const icdDur = time.Millisecond * 2500
		icd := core.NewICD()

		return core.Aura{
			ID:   TheLightningCapacitorAuraID,
			Name: "The Lightning Capacitor",
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if icd.IsOnCD(sim) {
					return
				}

				if !spellEffect.Crit {
					return
				}

				charges++
				if charges >= 3 {
					icd = core.InternalCD(sim.CurrentTime + icdDur)
					charges = 0

					castAction := &spellObj
					castTemplate.Apply(castAction)
					castAction.Target = spellEffect.Target
					castAction.Init(sim)
					castAction.Act(sim)
				}
			},
		}
	})
}

// Returns a cast object for a Lightning Capacitor cast with as many fields precomputed as possible.
func newLightningCapacitorCastTemplate(sim *core.Simulation, character *core.Character) core.SimpleSpellTemplate {
	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				Name: "Lightning Capacitor",
				ActionID: core.ActionID{
					ItemID: core.ItemIDTheLightningCapacitor,
				},
				Character:       character,
				IgnoreCooldowns: true,
				IgnoreManaCost:  true,
				SpellSchool:     stats.NatureSpellPower,
				CritMultiplier:  1.5,
				OnCastComplete:  func(sim *core.Simulation, cast *core.Cast) {},
			},
		},
		SpellHitEffect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier: 1,

				OnSpellHit:  func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {},
				OnSpellMiss: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {},
			},
			DirectInput: core.DirectDamageSpellInput{
				MinBaseDamage: 694,
				MaxBaseDamage: 807,
			},
		},
	})
}
