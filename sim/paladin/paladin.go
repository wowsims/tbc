package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func RegisterRetributionPaladin() {
	core.RegisterAgentFactory(
		proto.Player_RetributionPaladin{},
		func(character core.Character, options proto.Player) core.Agent {
			return NewPaladin(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_RetributionPaladin) // I don't really understand this line
			if !ok {
				panic("Invalid spec value for Paladin!")
			}
			player.Spec = playerSpec
		},
	)
}


type Paladin struct {
	core.Character
	Talents   proto.PaladinTalents
	Options proto.RetributionPaladin_Options
	Rotation proto.RetributionPaladin_Rotation

}

func (paladin *Paladin) GetCharacter() *core.Character {
	return &paladin.Character
}

func (paladin *Paladin) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
}

func (paladin *Paladin) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
}

func (paladin *Paladin) Init(sim *core.Simulation) {
}

func (paladin *Paladin) Reset(sim *core.Simulation) {}

func (enh *Paladin) OnGCDReady(sim *core.Simulation) {
}

func NewPaladin(character core.Character, options proto.Player) *Paladin {
	paladinOptions := options.GetRetributionPaladin()

	paladin := &Paladin{
		Character:    character,
		Talents:      *paladinOptions.Talents,
		Options:      *paladinOptions.Options,
		Rotation: 	  *paladinOptions.Rotation,
	}

	paladin.EnableManaBar()
	paladin.EnableAutoAttacks(paladin, core.AutoAttackOptions{
		MainHand:       paladin.WeaponFromMainHand(paladin.DefaultMeleeCritMultiplier()),
		OffHand:        paladin.WeaponFromOffHand(paladin.DefaultMeleeCritMultiplier()),
		AutoSwingMelee: true,
	})

	return paladin
}

// Idk how to generate these for other races so I'll start with blood elf only
// I probably also did this wrong for belfs so revist - just trying to get this to compile to start
func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceBloodElf, Class: proto.Class_ClassPaladin}] = stats.Stats{
		stats.Strength:  123,
		stats.Agility:   79,
		stats.Stamina:   118,
		stats.Intellect: 87,
		stats.Spirit:    88,
		stats.Mana:      3978, // pretty sure I need to subtract mana from the int stat

		stats.AttackPower:       120,
	}
}

func (paladin *Paladin) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // makes the bot wait forever and do nothing.
}

type Agent interface {
	GetPaladin() *Paladin
}
