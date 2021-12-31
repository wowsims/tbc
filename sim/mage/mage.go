package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func RegisterMage() {
	core.RegisterAgentFactory(
		proto.Player_Mage{},
		func(character core.Character, options proto.Player) core.Agent {
			return NewMage(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_Mage)
			if !ok {
				panic("Invalid spec value for Mage!")
			}
			player.Spec = playerSpec
		},
	)
}

type Mage struct {
	core.Character
	Talents proto.MageTalents

	RotationType   proto.Mage_Rotation_Type
	ArcaneRotation proto.Mage_Rotation_ArcaneRotation
	FireRotation   proto.Mage_Rotation_FireRotation
	FrostRotation  proto.Mage_Rotation_FrostRotation

	// Cached value for a few talents.
	spellDamageMultiplier float64

	// cached cast stuff
	arcaneBlastSpell        core.SimpleSpell
	arcaneBlastCastTemplate core.SimpleSpellTemplate

	igniteSpell        core.SimpleSpell
	igniteCastTemplate core.SimpleSpellTemplate

	fireballSpell        core.SimpleSpell
	fireballCastTemplate core.SimpleSpellTemplate

	fireballDotSpell        core.SimpleSpell
	fireballDotCastTemplate core.SimpleSpellTemplate

	frostboltSpell        core.SimpleSpell
	frostboltCastTemplate core.SimpleSpellTemplate

	scorchSpell        core.SimpleSpell
	scorchCastTemplate core.SimpleSpellTemplate

	wintersChillSpell        core.SimpleSpell
	wintersChillCastTemplate core.SimpleSpellTemplate
}

func (mage *Mage) GetCharacter() *core.Character {
	return &mage.Character
}

func (mage *Mage) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
	raidBuffs.ArcaneBrilliance = true
}
func (mage *Mage) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
}

func (mage *Mage) Init(sim *core.Simulation) {
	mage.arcaneBlastCastTemplate = mage.newArcaneBlastTemplate(sim)
	mage.igniteCastTemplate = mage.newIgniteTemplate(sim)
	mage.fireballCastTemplate = mage.newFireballTemplate(sim)
	mage.fireballDotCastTemplate = mage.newFireballDotTemplate(sim)
	mage.frostboltCastTemplate = mage.newFrostboltTemplate(sim)
	mage.scorchCastTemplate = mage.newScorchTemplate(sim)
	mage.wintersChillCastTemplate = mage.newWintersChillTemplate(sim)
}

func (mage *Mage) Reset(newsim *core.Simulation) {
}

func (mage *Mage) Advance(sim *core.Simulation, elapsedTime time.Duration) {
	mage.Character.RegenMana(sim, elapsedTime)
}

func NewMage(character core.Character, options proto.Player) *Mage {
	mageOptions := options.GetMage()

	mage := &Mage{
		Character:    character,
		Talents:      *mageOptions.Talents,
		RotationType: mageOptions.Rotation.Type,

		spellDamageMultiplier: 1.0,
	}

	if mage.RotationType == proto.Mage_Rotation_Arcane && mageOptions.Rotation.Arcane != nil {
		mage.ArcaneRotation = *mageOptions.Rotation.Arcane
	} else if mage.RotationType == proto.Mage_Rotation_Fire && mageOptions.Rotation.Fire != nil {
		mage.FireRotation = *mageOptions.Rotation.Fire
	} else if mage.RotationType == proto.Mage_Rotation_Frost && mageOptions.Rotation.Frost != nil {
		mage.FrostRotation = *mageOptions.Rotation.Frost
	}

	mage.Character.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Intellect,
		ModifiedStat: stats.SpellCrit,
		Modifier: func(intellect float64, spellCrit float64) float64 {
			return spellCrit + (intellect/81)*core.SpellCritRatingPerCritChance
		},
	})

	mage.Character.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Strength,
		ModifiedStat: stats.AttackPower,
		Modifier: func(strength float64, attackPower float64) float64 {
			return attackPower + strength*2
		},
	})

	if mageOptions.Options.Armor == proto.Mage_Options_MageArmor {
		mage.PseudoStats.SpiritRegenRateCasting += 0.3
	} else if mageOptions.Options.Armor == proto.Mage_Options_MoltenArmor {
		mage.AddStat(stats.SpellCrit, 3*core.SpellCritRatingPerCritChance)
	}

	mage.registerEvocationCD()
	mage.applyTalents()

	return mage
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceBloodElf, Class: proto.Class_ClassMage}] = stats.Stats{
		stats.Strength:  30,
		stats.Agility:   41,
		stats.Stamina:   50,
		stats.Intellect: 155,
		stats.Spirit:    144,
		stats.Mana:      2241,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 0.926,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDraenei, Class: proto.Class_ClassMage}] = stats.Stats{
		stats.Strength:  34,
		stats.Agility:   36,
		stats.Stamina:   50,
		stats.Intellect: 152,
		stats.Spirit:    147,
		stats.Mana:      2241,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 0.933,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceGnome, Class: proto.Class_ClassMage}] = stats.Stats{
		stats.Strength:  28,
		stats.Agility:   42,
		stats.Stamina:   50,
		stats.Intellect: 154.3, // Gnomes start with 162 int, we assume this include racial so / 1.05
		stats.Spirit:    145,
		stats.Mana:      2241,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.029,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceHuman, Class: proto.Class_ClassMage}] = stats.Stats{
		stats.Strength:  33,
		stats.Agility:   39,
		stats.Stamina:   51,
		stats.Intellect: 151,
		stats.Spirit:    159,
		stats.Mana:      2241,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 0.926,
	}
	trollStats := stats.Stats{
		stats.Strength:  34,
		stats.Agility:   41,
		stats.Stamina:   52,
		stats.Intellect: 147,
		stats.Spirit:    146,
		stats.Mana:      2241,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 0.935,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll10, Class: proto.Class_ClassMage}] = trollStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll30, Class: proto.Class_ClassMage}] = trollStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceUndead, Class: proto.Class_ClassMage}] = stats.Stats{
		stats.Strength:  32,
		stats.Agility:   37,
		stats.Stamina:   52,
		stats.Intellect: 149,
		stats.Spirit:    150,
		stats.Mana:      2241,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 0.930,
	}
}

// Agent is a generic way to access underlying mage on any of the agents.
type Agent interface {
	GetMage() *Mage
}
