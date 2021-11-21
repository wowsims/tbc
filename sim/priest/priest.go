package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Priest struct {
	core.Character
	SelfBuffs
	Talents proto.PriestTalents

	// cached cast stuff
	mindflaySpell        core.SimpleSpell
	mindflayCastTemplate core.SimpleSpellTemplate

	mindblastSpell        core.SimpleSpell
	mindblastCastTemplate core.SimpleSpellTemplate

	swpSpell        core.SimpleSpell
	swpCastTemplate core.SimpleSpellTemplate

	vtSpell        core.SimpleSpell
	vtCastTemplate core.SimpleSpellTemplate
}

type SelfBuffs struct {
}

func (priest *Priest) GetCharacter() *core.Character {
	return &priest.Character
}

func (priest *Priest) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
	raidBuffs.DivineSpirit = proto.TristateEffect_TristateEffectRegular
}

func (priest *Priest) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {

}

func (priest *Priest) Init(sim *core.Simulation) {
	// druid.starfire8CastTemplate = druid.newStarfireTemplate(sim, 8)
}

func (priest *Priest) Reset(newsim *core.Simulation) {
	// Cleanup and pending dots and casts
	// druid.starfireSpell = core.SimpleSpell{}
	priest.Character.Reset(newsim)
}

func (priest *Priest) Advance(sim *core.Simulation, elapsedTime time.Duration) {
	// druid should never be outside the 5s window, use combat regen.
	priest.Character.RegenManaCasting(sim, elapsedTime)
	priest.Character.Advance(sim, elapsedTime)
}

func (priest *Priest) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // does nothing
}

func NewPriest(char core.Character, selfBuffs SelfBuffs, talents proto.PriestTalents) Priest {

	// char.AddStat(stats.SpellHit, float64(talents.BalanceOfPower)*2*core.SpellHitRatingPerHitChance)

	char.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Intellect,
		ModifiedStat: stats.SpellCrit,
		Modifier: func(intellect float64, spellCrit float64) float64 {
			return spellCrit + (intellect/79.4)*core.SpellCritRatingPerCritChance
		},
	})

	// if talents.LunarGuidance > 0 {
	// 	bonus := (0.25 / 3) * float64(talents.LunarGuidance)
	// 	char.AddStatDependency(stats.StatDependency{
	// 		SourceStat:   stats.Intellect,
	// 		ModifiedStat: stats.SpellPower,
	// 		Modifier: func(intellect float64, spellPower float64) float64 {
	// 			return spellPower + intellect*bonus
	// 		},
	// 	})
	// }
	priest := Priest{
		Character: char,
		SelfBuffs: selfBuffs,
		Talents:   talents,
	}

	return priest
}

// TODO: Get Priest base stats
func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceHuman, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2090,  // 3815 mana shown on naked character
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDwarf, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2090,  // 3815 mana shown on naked character
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceNightElf, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2090,  // 3815 mana shown on naked character
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDraenei, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2090,  // 3815 mana shown on naked character
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceUndead, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2090,  // 3815 mana shown on naked character
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
	troll := stats.Stats{
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2090,  // 3815 mana shown on naked character
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll10, Class: proto.Class_ClassPriest}] = troll
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll30, Class: proto.Class_ClassPriest}] = troll
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceBloodElf, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2090,  // 3815 mana shown on naked character
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
}

// Agent is a generic way to access underlying priest on any of the agents
type Agent interface {
	GetPriest() *Priest
}
