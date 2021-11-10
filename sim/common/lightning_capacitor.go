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
	spellObj := core.DirectCastAction{}

	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		charges := 0

		const icdDur = time.Millisecond * 2500
		icd := core.NewICD()

		return core.Aura{
			ID:      TheLightningCapacitorAuraID,
			Name:    "The Lightning Capacitor",
			OnSpellHit: func(sim *core.Simulation, cast *core.Cast, result *core.DirectCastDamageResult) {
				if icd.IsOnCD(sim) {
					return
				}

				if !result.Crit {
					return
				}

				charges++
				if charges >= 3 {
					icd = core.InternalCD(sim.CurrentTime + icdDur)
					charges = 0
					castAction := NewLightningCapacitorCast(sim, character, sim.GetPrimaryTarget(), &spellObj)
					castAction.Act(sim)
				}
			},
		}
	})
}

func NewLightningCapacitorCast(sim *core.Simulation, character *core.Character, target *core.Target, spellObj *core.DirectCastAction) *core.DirectCastAction {
	spell := spellObj
	*spell = core.DirectCastAction{
		Cast: core.Cast{
			Name: "Lightning Capacitor",
			ActionID: core.ActionID{
				ItemID: core.ItemIDTheLightningCapacitor,
			},
			Character: character,
			SpellSchool: stats.NatureSpellPower,
			IgnoreCooldowns: true,
			IgnoreManaCost: true,
			CritMultiplier: 1.5,
		},
		HitInputs: []core.DirectCastDamageInput{
			core.DirectCastDamageInput{
				Target: target,
				MinBaseDamage: 694,
				MaxBaseDamage: 807,
				DamageMultiplier: 1,
			},
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {},
		OnSpellHit: func(sim *core.Simulation, cast *core.Cast, result *core.DirectCastDamageResult) {},
		OnSpellMiss: func(sim *core.Simulation, cast *core.Cast) {},
	}

	spell.Init(sim)
	return spell
}
