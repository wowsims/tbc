package paladin

import "github.com/wowsims/tbc/sim/core"

func NewBuffBot(sim *core.Simulation, party *core.Party, blessingOfKings, improvedBlessingOfWisdom, improvedSealOfTheCrusader, judgementOfWisdom bool) *Paladin {

	for _, raidParty := range sim.Raid.Parties {
		for _, pl := range raidParty.Players {
			if improvedBlessingOfWisdom {
				pl.Stats[core.StatMP5] += 42
				pl.InitialStats[core.StatMP5] += 42
			}
			if improvedSealOfTheCrusader {
				pl.Stats[core.StatSpellCrit] += 66.24        // 3% crit
				pl.InitialStats[core.StatSpellCrit] += 66.24 // 3% crit
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
		sim.AddAura(sim, nil, AuraJudgementOfWisdom())
	}
}

func (p *Paladin) Reset(sim *core.Simulation) {

}
func (p *Paladin) OnSpellHit(*core.Simulation, *core.Player, *core.Cast) {}

func AuraJudgementOfWisdom() core.Aura {
	const mana = 74 / 2 // 50% proc
	return core.Aura{
		ID:      core.MagicIDJoW,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, player *core.Player, c *core.Cast) {
			if c.Spell.ID == core.MagicIDTLCLB {
				return // TLC cant proc JoW
			}
			if sim.Debug != nil {
				sim.Debug(" +Judgement Of Wisdom: 37 mana (74 @ 50%% proc)\n")
			}
			// Only apply to players that have mana.
			if player.InitialStats[core.StatMana] > 0 {
				player.Stats[core.StatMana] += mana
			}
		},
	}
}
