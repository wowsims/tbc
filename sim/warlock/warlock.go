package warlock

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Warlock struct {
	core.Character
	Talents  proto.WarlockTalents
	Options  proto.Warlock_Options
	Rotation proto.Warlock_Rotation

	Shadowbolt     *core.Spell
	Incinerate     *core.Spell
	Immolate       *core.Spell
	ImmolateDot    *core.Dot
	UnstableAff    *core.Spell
	UnstableAffDot *core.Dot
	Corruption     *core.Spell
	CorruptionDot  *core.Dot
	SiphonLife     *core.Spell
	SiphonLifeDot  *core.Dot

	LifeTap *core.Spell

	CurseOfElements     *core.Spell
	CurseOfElementsAura *core.Aura

	NightfallProcAura *core.Aura

	DoingRegen bool
}

func (warlock *Warlock) GetCharacter() *core.Character {
	return &warlock.Character
}

func (warlock *Warlock) Init(sim *core.Simulation) {
	warlock.registerIncinerateSpell(sim)
	warlock.registerShadowboltSpell(sim)
	warlock.registerImmolateSpell(sim)
	warlock.registerCorruptionSpell(sim)
	warlock.registerCurseOfElementsSpell(sim)
	warlock.registerLifeTapSpell(sim)
	if warlock.Talents.UnstableAffliction {
		warlock.registerUnstableAffSpell(sim)
	}
	if warlock.Talents.SiphonLife {
		warlock.registerSiphonLifeSpell(sim)
	}
}

func (warlock *Warlock) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
}

func (warlock *Warlock) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
}

func (warlock *Warlock) Reset(sim *core.Simulation) {

}

func NewWarlock(character core.Character, options proto.Player) *Warlock {
	warlockOptions := options.GetWarlock()

	warlock := &Warlock{
		Character: character,
		Talents:   *warlockOptions.Talents,
		Options:   *warlockOptions.Options,
		Rotation:  *warlockOptions.Rotation,
		// manaTracker:           common.NewManaSpendingRateTracker(),
	}
	warlock.EnableManaBar()

	warlock.Character.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Intellect,
		ModifiedStat: stats.SpellCrit,
		Modifier: func(intellect float64, spellCrit float64) float64 {
			return spellCrit + (intellect/81.92)*core.SpellCritRatingPerCritChance
		},
	})

	warlock.Character.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Strength,
		ModifiedStat: stats.AttackPower,
		Modifier: func(strength float64, attackPower float64) float64 {
			return attackPower + strength*2
		},
	})

	if warlock.Options.Armor == proto.Warlock_Options_FelArmor {
		amount := 100.0
		amount *= 1 + float64(warlock.Talents.DemonicAegis)*0.1
		warlock.AddStat(stats.SpellPower, amount)
	}

	if warlock.Talents.DemonicSacrifice && warlock.Options.SacrificeSummon {
		switch warlock.Options.Summon {
		case proto.Warlock_Options_Succubus:
			warlock.PseudoStats.ShadowDamageDealtMultiplier *= 1.15
		case proto.Warlock_Options_Imp:
			warlock.PseudoStats.FireDamageDealtMultiplier *= 1.15
		}
	}

	return warlock
}

func RegisterWarlock() {
	core.RegisterAgentFactory(
		proto.Player_Warlock{},
		proto.Spec_SpecWarlock,
		func(character core.Character, options proto.Player) core.Agent {
			return NewWarlock(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_Warlock)
			if !ok {
				panic("Invalid spec value for Warlock!")
			}
			player.Spec = playerSpec
		},
	)
}

func init() {
	// hardcoded to bloodelf for now
	baseStats := stats.Stats{
		stats.Health:      2005,
		stats.Strength:    48,
		stats.Agility:     60,
		stats.Stamina:     75,
		stats.Intellect:   137,
		stats.Spirit:      130,
		stats.Mana:        2335,
		stats.SpellCrit:   1.697 * core.SpellCritRatingPerCritChance,
		stats.AttackPower: 86,
		// Not sure how stats modify the crit chance.
		// stats.MeleeCrit:   4.43 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceBloodElf, Class: proto.Class_ClassWarlock}] = baseStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceOrc, Class: proto.Class_ClassWarlock}] = baseStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceUndead, Class: proto.Class_ClassWarlock}] = baseStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceHuman, Class: proto.Class_ClassWarlock}] = baseStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceGnome, Class: proto.Class_ClassWarlock}] = baseStats
}
