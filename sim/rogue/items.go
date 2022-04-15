package rogue

import (
	"log"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddItemEffect(30450, ApplyWarpSpringCoil)
	core.AddItemEffect(32492, ApplyAshtongueTalismanOfLethality)

	core.AddItemSet(&ItemSetAssassination)
	core.AddItemSet(&ItemSetNetherblade)
	core.AddItemSet(&ItemSetDeathmantle)
	core.AddItemSet(&ItemSetSlayers)
}

var ItemSetAssassination = core.ItemSet{
	Name: "Assassination Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
		},
		4: func(agent core.Agent) {
			// Your Eviscerate and Envenom abilities cost 10 less energy.
			// Handled in eviscerate.go.
		},
	},
}

var ItemSetNetherblade = core.ItemSet{
	Name: "Netherblade",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Increases the duration of your Slice and Dice ability by 3 sec.
			// Handled in slice_and_dice.go.
		},
		4: func(agent core.Agent) {
			// Your finishing moves have a 15% chance to grant you an extra combo point.
			// Handled in talents.go.
		},
	},
}

var ItemSetDeathmantle = core.ItemSet{
	Name: "Deathmantle",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Your Eviscerate and Envenom abilities cause 40 extra damage per combo point.
			// Handled in eviscerate.go.
		},
		4: func(agent core.Agent) {
			// Your attacks have a chance to make your next finishing move cost no energy.
			rogueAgent, ok := agent.(RogueAgent)
			if !ok {
				log.Fatalf("Non-rogue attempted to activate rogue t5 4p bonus.")
			}
			rogue := rogueAgent.GetRogue()

			ppmm := rogue.AutoAttacks.NewPPMManager(1.0)

			rogue.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
				return rogue.GetOrRegisterAura(core.Aura{
					Label: "Deathmantle 4pc",
					OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
						if !spellEffect.Landed() {
							return
						}

						// https://tbc.wowhead.com/spell=37170/free-finisher-chance, proc mask = 20.
						if !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
							return
						}

						if !ppmm.Proc(sim, spellEffect.IsMH(), false, "Deathmantle 4pc") {
							return
						}

						rogue.deathmantle4pcProc = true
					},
				})
			})
		},
	},
}

func (rogue *Rogue) applyDeathmantle(_ *core.Simulation, _ *core.Spell, cast *core.Cast) {
	//instance.ActionID.Tag = rogue.ComboPoints()
	if rogue.deathmantle4pcProc {
		cast.Cost = 0
		rogue.deathmantle4pcProc = false
	}
}

var ItemSetSlayers = core.ItemSet{
	Name: "Slayer's Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Increases the haste from your Slice and Dice ability by 5%.
			// Handled in slice_and_dice.go.
		},
		4: func(agent core.Agent) {
			// Increases the damage dealt by your Backstab, Sinister Strike, Mutilate, and Hemorrhage abilities by 6%.
			// Handled in the corresponding ability files.
		},
	},
}

func ApplyWarpSpringCoil(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		procAura := character.NewTemporaryStatsAura("Warp Spring Coil Proc", core.ActionID{ItemID: 30450}, stats.Stats{stats.ArmorPenetration: 1000}, time.Second*15)
		const procChance = 0.25
		const icdDur = time.Second * 30
		icd := core.NewICD()

		return character.GetOrRegisterAura(core.Aura{
			Label: "Warp Spring Coil",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}

				// https://tbc.wowhead.com/spell=37173/armor-penetration, proc mask = 16.
				if !spellEffect.ProcMask.Matches(core.ProcMaskMeleeSpecial) {
					return
				}

				if icd.IsOnCD(sim) {
					return
				}

				if sim.RandomFloat("WarpSpringCoil") > procChance {
					return
				}

				icd = core.InternalCD(sim.CurrentTime + icdDur)
				procAura.Activate(sim)
			},
		})
	})
}

func ApplyAshtongueTalismanOfLethality(agent core.Agent) {
	rogueAgent, ok := agent.(RogueAgent)
	if !ok {
		log.Fatalf("Non-rogue attempted to activate Ashtongue Talisman of Lethality.")
	}
	rogue := rogueAgent.GetRogue()

	rogue.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		procAura := rogue.NewTemporaryStatsAura("Ashtongue Talisman Proc", core.ActionID{ItemID: 32492}, stats.Stats{stats.MeleeCrit: 145}, time.Second*10)
		numPoints := int32(0)

		return rogue.GetOrRegisterAura(core.Aura{
			Label: "Ashtongue Talisman",
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
				if !spell.SpellExtras.Matches(SpellFlagFinisher) {
					return
				}

				// Need to store the points because they get spent before OnSpellHit is called.
				numPoints = rogue.ComboPoints()

				if spell.SameActionIgnoreTag(SliceAndDiceActionID) {
					// SND won't call OnSpellHit so we have to add the effect now.
					if numPoints == 5 || sim.RandomFloat("AshtongueTalismanOfLethality") < 0.2*float64(numPoints) {
						procAura.Activate(sim)
					}
				}
			},
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spell.SpellExtras.Matches(SpellFlagFinisher) {
					return
				}

				if numPoints == 5 || sim.RandomFloat("AshtongueTalismanOfLethality") < 0.2*float64(numPoints) {
					procAura.Activate(sim)
				}
			},
		})
	})
}
