package paladin

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func NewBuffBot(sim *core.Simulation, party *core.Party, blessingOfKings, improvedBlessingOfWisdom, improvedSealOfTheCrusader, judgementOfWisdom bool) *Paladin {

	for _, raidParty := range sim.Raid.Parties {
		for _, pl := range raidParty.Players {
			if improvedBlessingOfWisdom {
				pl.Stats[stats.MP5] += 42
				pl.InitialStats[stats.MP5] += 42
			}
			if improvedSealOfTheCrusader {
				pl.Stats[stats.SpellCrit] += 66.24        // 3% crit
				pl.InitialStats[stats.SpellCrit] += 66.24 // 3% crit
				// FUTURE: melee crit bonus, research actual value
			}
		}
	}

	return &Paladin{
		useJoW: judgementOfWisdom,
	}
}

type Paladin struct {
	useJoW bool
	core.Agent
}

func (p *Paladin) ChooseAction(_ *core.Simulation, party *core.Party) core.AgentAction {
	return core.AgentAction{Wait: core.NeverExpires} // makes the bot wait forever and do nothing.
}

func (p *Paladin) OnActionAccepted(*core.Simulation, core.AgentAction) {

}

func (p *Paladin) BuffUp(sim *core.Simulation, party *core.Party) {
	// Judgement of Wisdom
	if p.useJoW {
		sim.AddAura(sim, core.PlayerAgent{}, AuraJudgementOfWisdom()) // no player for global auras
	}
}

func (p *Paladin) Reset(sim *core.Simulation) {

}
func (p *Paladin) OnSpellHit(*core.Simulation, core.PlayerAgent, *core.Cast) {}

func AuraJudgementOfWisdom() core.Aura {
	const mana = 74 / 2 // 50% proc
	return core.Aura{
		ID:      core.MagicIDJoW,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, player core.PlayerAgent, c *core.Cast) {
			if c.Spell.ID == core.MagicIDTLCLB {
				return // TLC cant proc JoW
			}
			if sim.Debug != nil {
				sim.Debug("(%d) +Judgement Of Wisdom: 37 mana (74 @ 50%% proc)\n", player.ID)
			}
			// Only apply to players that have mana.
			if player.InitialStats[stats.Mana] > 0 {
				player.Stats[stats.Mana] += mana
			}
		},
	}
}
