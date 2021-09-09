package core

import (
	"strconv"
	"time"
)

type Party struct {
	Players []PlayerAgent
}

type Raid struct {
	Parties []*Party
}

// Player is a data structure to hold all the shared values that all
// class logic shares.
//  All players have stats, equipment, auras, etc
type Player struct {
	Consumes Consumes // pretty sure most classes have consumes to care about.

	InitialStats Stats
	Stats        Stats

	Equip       Equipment // Current Gear
	EquipSpec   EquipmentSpec
	ActiveEquip []*Item // cache of gear that can activate.

	*AuraTracker
}

func (p *Player) HasteBonus() float64 {
	return 1 + (p.Stats[StatSpellHaste] / 1576)
}

func (p *Player) GetPlayer() *Player {
	return p
}

func NewPlayer(equip EquipmentSpec, consumes Consumes) *Player {
	// TODO: configure player here.

	// equip := NewEquipmentSet(equipSpec)
	// initialStats := CalculateTotalStats(options, equip)

	// for i, eq := range equip {
	// 	if eq.Activate != nil {
	// 		player.activeEquip = append(player.activeEquip, &equip[i])
	// 	}
	// 	for _, g := range eq.Gems {
	// 		if g.Activate != nil {
	// 			player.activeEquip = append(player.activeEquip, &equip[i])
	// 		}
	// 	}
	// }

	return &Player{
		AuraTracker: &AuraTracker{},
	}
}

// Advance moves time forward counting down auras, CDs, mana regen, etc
func (p *Player) Advance(sim *Simulation, elapsedTime time.Duration, newTime time.Duration) {
	// MP5 regen
	p.Stats[StatMana] += p.manaRegenPerSecond() * elapsedTime.Seconds()
	if p.Stats[StatMana] > p.InitialStats[StatMana] {
		p.Stats[StatMana] = p.InitialStats[StatMana]
	}
	p.AuraTracker.Advance(sim, p, newTime)
}

func (p *Player) Reset() {
	// sim.Stats = sim.InitialStats
	// sim.cache.destructionPotion = false
	// sim.cache.bloodlustCasts = 0
	// sim.CurrentTime = 0.0
	// sim.CurrentMana = sim.Stats[StatMana]
	// sim.CDs = [MagicIDLen]time.Duration{}
	// sim.auras = [MagicIDLen]Aura{}
	// if len(sim.activeAuraIDs) > 0 {
	// 	sim.activeAuraIDs = sim.activeAuraIDs[len(sim.activeAuraIDs)-1:] // chop off end of activeids slice, faster than making a new one
	// }
	// sim.metrics = SimMetrics{
	// 	Casts: make([]*Cast, 0, 1000),
	// }

	// if sim.Debug != nil {
	// 	sim.Debug("SIM RESET\n")
	// 	sim.Debug("----------------------\n")
	// }

	// // Activate all talents
	// if sim.Options.Talents.LightningOverload > 0 {
	// 	sim.addAura(AuraLightningOverload(sim.Options.Talents.LightningOverload))
	// }

	// // Chain lightning bounces
	// if sim.Options.Encounter.NumClTargets > 1 {
	// 	sim.addAura(ActivateChainLightningBounce(sim))
	// }

	// // Judgement of Wisdom
	// if sim.Options.Buffs.JudgementOfWisdom {
	// 	sim.addAura(AuraJudgementOfWisdom())
	// }

	// // Activate all permanent item effects.
	// for _, item := range sim.activeEquip {
	// 	if item.Activate != nil && item.ActivateCD == NeverExpires {
	// 		sim.addAura(item.Activate(sim))
	// 	}
	// 	for _, g := range item.Gems {
	// 		if g.Activate != nil {
	// 			sim.addAura(g.Activate(sim))
	// 		}
	// 	}
	// }

	// sim.ActivateSets()
	// // Currently no temp hit increase effects in the sim... so lets cache this!
	// sim.cache.spellHit = 0.83 + sim.Stats[StatSpellHit]/1260.0 // 12.6 hit == 1% hit
}

// AddAura on player is a simple wrapper around AuraTracker so the
// consumer doesn't need to pass player back into itself.
func (p *Player) AddAura(sim *Simulation, a Aura) {
	p.AuraTracker.AddAura(sim, p, a)
}

// Returns rate of mana regen, as mana / second
func (p *Player) manaRegenPerSecond() float64 {
	return p.Stats[StatMP5] / 5.0
}

// Pops any on-use trinkets / gear
func (p *Player) TryActivateEquipment() {
	// TODO: Fix this.

	// 	for _, item := range sim.activeEquip {
	// 		if item.Activate == nil || item.ActivateCD == NeverExpires { // ignore non-activatable, and always active items.
	// 			continue
	// 		}
	// 		if sim.isOnCD(item.CoolID) {
	// 			continue
	// 		}
	// 		if item.Slot == EquipTrinket && sim.isOnCD(MagicIDAllTrinket) {
	// 			continue
	// 		}
	// 		sim.addAura(item.Activate(sim))
	// 		sim.setCD(item.CoolID, item.ActivateCD)
	// 		if item.Slot == EquipTrinket {
	// 			sim.setCD(MagicIDAllTrinket, time.Second*30)
	// 		}
	// 	}
}

// Activates set bonuses, returning the list of active bonuses.
func (p *Player) ActivateSets(sim *Simulation) []string {
	active := []string{}
	// Activate Set Bonuses
	setItemCount := map[string]int{}
	for _, i := range p.Equip {
		set := itemSetLookup[i.ID]
		if set != nil {
			setItemCount[set.Name]++
			if bonus, ok := set.Bonuses[setItemCount[set.Name]]; ok {
				active = append(active, set.Name+" ("+strconv.Itoa(setItemCount[set.Name])+"pc)")
				p.AddAura(sim, bonus(sim, p))
			}
		}
	}

	return active
}
