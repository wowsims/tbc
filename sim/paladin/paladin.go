package paladin

import (
	"github.com/wowsims/tbc/items"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func NewBuffBot(sim *core.Simulation, party *core.Party, blessingOfKings, improvedBlessingOfWisdom, improvedSealOfTheCrusader, judgementOfWisdom bool) *Paladin {
	paladin := &Paladin{
		Character: core.NewCharacter(items.EquipmentSpec{}, core.RaceBonusTypeNone, core.Consumes{}, stats.Stats{}),
		blessingOfKings: blessingOfKings,
		improvedBlessingOfWisdom: improvedBlessingOfWisdom,
		improvedSealOfTheCrusader: improvedSealOfTheCrusader,
		useJoW: judgementOfWisdom,
	}
	paladin.Character.Agent = paladin
	return paladin
}

type Paladin struct {
	*core.Character

	blessingOfKings           bool
	improvedBlessingOfWisdom  bool
	improvedSealOfTheCrusader bool
	useJoW                    bool
}

func (paladin *Paladin) GetCharacter() *core.Character {
	return paladin.Character
}

func (p *Paladin) ChooseAction(_ *core.Simulation) core.AgentAction {
	return core.AgentAction{Wait: core.NeverExpires} // makes the bot wait forever and do nothing.
}

func (p *Paladin) OnActionAccepted(*core.Simulation, core.AgentAction) {

}

func (p *Paladin) BuffUp(sim *core.Simulation) {
	if p.improvedBlessingOfWisdom {
		sim.Raid.AddStats(stats.Stats{
			stats.MP5: 42,
		})
	}
	if p.improvedSealOfTheCrusader {
		sim.Raid.AddStats(stats.Stats{
			stats.SpellCrit: 66.24, // 3% crit
		})
		// FUTURE: melee crit bonus, research actual value
	}

	// Judgement of Wisdom
	if p.useJoW {
		sim.AddAura(sim, nil, AuraJudgementOfWisdom()) // no player for global auras
	}
}

func (p *Paladin) Reset(sim *core.Simulation) {

}
func (p *Paladin) OnSpellHit(*core.Simulation, *core.Cast) {}

func AuraJudgementOfWisdom() core.Aura {
	const mana = 74 / 2 // 50% proc
	return core.Aura{
		ID:      core.MagicIDJoW,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, agent core.Agent, c *core.Cast) {
			if c.Spell.ID == core.MagicIDTLCLB {
				return // TLC cant proc JoW
			}
			if sim.Log != nil {
				sim.Log("(%d) +Judgement Of Wisdom: 37 mana (74 @ 50%% proc)\n", agent.GetCharacter().ID)
			}
			// Only apply to agents that have mana.
			if agent.GetCharacter().InitialStats[stats.Mana] > 0 {
				agent.GetCharacter().Stats[stats.Mana] += mana
			}
		},
	}
}
