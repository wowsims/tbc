package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Keep these (and their functions) in alphabetical order.
func init() {
	// Proc effects. Keep these in order by item ID.
	core.AddItemEffect(28602, ApplyRobeOfTheElderScribes)
	core.AddItemEffect(29305, ApplyEternalSage)
}

var RobeOfTheElderScribeAuraID = core.NewAuraID()
var PowerOfArcanagosAuraID = core.NewAuraID()

func ApplyRobeOfTheElderScribes(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		// Gives a chance when your harmful spells land to increase the damage of your spells and effects by up to 130 for 10 sec. (Proc chance: 20%, 50s cooldown)
		icd := core.NewICD()
		const spellBonus = 130.0
		const dur = time.Second * 10
		const icdDur = time.Second * 50
		const proc = 0.2

		return core.Aura{
			ID:   RobeOfTheElderScribeAuraID,
			Name: "Robes of the Elder Scibe",
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !icd.IsOnCD(sim) && sim.RandomFloat("Robe of the Elder Scribe") < proc {
					icd = core.InternalCD(sim.CurrentTime + icdDur)
					character.AddAuraWithTemporaryStats(sim, PowerOfArcanagosAuraID, 0, "Power of Arcanagos", stats.SpellPower, spellBonus, dur)
				}
			},
		}
	})
}

var EternalSageItemAuraID = core.NewAuraID()
var BandoftheEternalSageAuraID = core.NewAuraID()

func ApplyEternalSage(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		// Your offensive spells have a chance on hit to increase your spell damage by 95 for 10 secs.
		icd := core.NewICD()
		const spellBonus = 95.0
		const dur = time.Second * 10
		const icdDur = time.Second * 60
		const proc = 0.1

		return core.Aura{
			ID:   EternalSageItemAuraID,
			Name: "Band of the Enternal Sage Passive",
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !icd.IsOnCD(sim) && sim.RandomFloat("Band of the Eternal Sage") < proc {
					icd = core.InternalCD(sim.CurrentTime + icdDur)
					character.AddAuraWithTemporaryStats(sim, BandoftheEternalSageAuraID, 0, "Band of the Eternal Sage", stats.SpellPower, spellBonus, dur)
				}
			},
		}
	})
}
