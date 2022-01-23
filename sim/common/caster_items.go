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
	core.AddItemEffect(34470, ApplyTimbals)
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
			ID: RobeOfTheElderScribeAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if icd.IsOnCD(sim) || sim.RandomFloat("Robe of the Elder Scribe") > proc { // can't activate if on CD or didn't proc
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				character.AddAuraWithTemporaryStats(sim, PowerOfArcanagosAuraID, core.ActionID{ItemID: 28602}, stats.SpellPower, spellBonus, dur)
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
			ID: EternalSageItemAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if icd.IsOnCD(sim) || sim.RandomFloat("Band of the Eternal Sage") > proc { // can't activate if on CD or didn't proc
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				character.AddAuraWithTemporaryStats(sim, BandoftheEternalSageAuraID, core.ActionID{ItemID: 29305}, stats.SpellPower, spellBonus, dur)
			},
		}
	})
}

var AugmentPainAuraID = core.NewAuraID()

func ApplyTimbals(agent core.Agent) {
	timbalsTemplate := core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				CritMultiplier:  1.5,
				SpellSchool:     stats.ShadowSpellPower,
				IgnoreCooldowns: true,
				ActionID: core.ActionID{
					SpellID: 45055,
				},
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage: 285,
				MaxBaseDamage: 475,
			},
		},
	})
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		var shadowBolt = &core.SimpleSpell{}

		// Each time one of your spells deals periodic damage,
		// there is a chance 285 to 475 additional damage will be dealt. (Proc chance: 10%, 15s cooldown)
		icd := core.NewICD()
		const icdDur = time.Second * 15
		const proc = 0.1

		return core.Aura{
			ID: AugmentPainAuraID,
			OnPeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage float64) {
				if icd.IsOnCD(sim) || sim.RandomFloat("timbals") > proc { // can't activate if on CD or didn't proc
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				timbalsTemplate.Apply(shadowBolt)
				// Apply the caster/target from the cast that procd this.
				shadowBolt.Character = spellCast.Character
				shadowBolt.Effect.Target = spellEffect.Target
				shadowBolt.Cast(sim)
			},
		}
	})
}
