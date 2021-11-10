package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func NewShaman(character core.Character, talents proto.ShamanTalents, selfBuffs SelfBuffs) Shaman {
	shaman := Shaman{
		Character: character,
		Talents:   talents,
		SelfBuffs: selfBuffs,

		convectionBonus: 0.02 * float64(talents.Convection),
		concussionBonus: 1 + 0.01*float64(talents.Concussion),
	}

	// Add Shaman stat dependencies
	shaman.AddStatDependency(stats.StatDependency{
		SourceStat: stats.Intellect,
		ModifiedStat: stats.SpellCrit,
		Modifier: func(intellect float64, spellCrit float64) float64 {
			return spellCrit + (intellect / 78.1) * core.SpellCritRatingPerCritChance
		},
	})

	if shaman.Talents.UnrelentingStorm > 0 {
		coeff := 0.02 * float64(shaman.Talents.UnrelentingStorm)
		shaman.AddStatDependency(stats.StatDependency{
			SourceStat: stats.Intellect,
			ModifiedStat: stats.MP5,
			Modifier: func(intellect float64, mp5 float64) float64 {
				return mp5 + intellect * coeff
			},
		})
	}

	if selfBuffs.WaterShield {
		shaman.AddStat(stats.MP5, 50)
	}

	shaman.registerElementalMasteryCD()

	return shaman
}

// Which buffs this shaman is using.
type SelfBuffs struct {
	Bloodlust    bool
	WaterShield  bool
	TotemOfWrath bool
	WrathOfAir   bool
	ManaSpring   bool
}

// Shaman represents a shaman character.
type Shaman struct {
	core.Character

	Talents   proto.ShamanTalents
	SelfBuffs SelfBuffs

	ElementalFocusStacks byte

	// cache
	convectionBonus float64
	concussionBonus float64

	// "object pool" for shaman spells.
	electricSpell   core.DirectCastAction
	electricSpellLO core.DirectCastAction
}

// Implemented by each Shaman spec.
type ShamanAgent interface {
	core.Agent

	// The Shaman controlled by this Agent.
	GetShaman() *Shaman
}

func (shaman *Shaman) GetCharacter() *core.Character {
	return &shaman.Character
}

func (shaman *Shaman) AddRaidBuffs(buffs *proto.Buffs) {
}
func (shaman *Shaman) AddPartyBuffs(buffs *proto.Buffs) {
	if shaman.SelfBuffs.Bloodlust {
		buffs.Bloodlust += 1
	}

	if shaman.SelfBuffs.TotemOfWrath {
		buffs.TotemOfWrath += 1
	}

	if shaman.SelfBuffs.ManaSpring {
		buffs.ManaSpringTotem = core.MaxTristate(buffs.ManaSpringTotem, proto.TristateEffect_TristateEffectRegular)
	}

	if shaman.SelfBuffs.WrathOfAir {
		woaValue := proto.TristateEffect_TristateEffectRegular
		if ItemSetCycloneRegalia.CharacterHasSetBonus(shaman.GetCharacter(), 2) {
			woaValue = proto.TristateEffect_TristateEffectImproved
		}
		buffs.WrathOfAirTotem = core.MaxTristate(buffs.WrathOfAirTotem, woaValue)
	}
}

func (shaman *Shaman) Reset(sim *core.Simulation) {
	shaman.Character.Reset(sim)
}

func (shaman *Shaman) Advance(sim *core.Simulation, elapsedTime time.Duration) {
	shaman.Character.RegenManaMP5Only(sim, elapsedTime)
	shaman.Character.Advance(sim, elapsedTime)
}

var ElementalMasteryAuraID = core.NewAuraID()
var ElementalMasteryCooldownID = core.NewCooldownID()
func (shaman *Shaman) registerElementalMasteryCD() {
	if !shaman.Talents.ElementalMastery {
		return
	}

	shaman.AddMajorCooldown(core.MajorCooldown{
		CooldownID: ElementalMasteryCooldownID,
		Cooldown: time.Minute*3,
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) bool {
				character.AddAura(sim, core.Aura{
					ID:      ElementalMasteryAuraID,
					Name:    "Elemental Mastery",
					Expires: core.NeverExpires,
					OnCast: func(sim *core.Simulation, cast *core.Cast) {
						cast.ManaCost = 0
						cast.GuaranteedCrit = true
					},
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						// Remove the buff and put skill on CD
						character.SetCD(ElementalMasteryCooldownID, sim.CurrentTime+time.Minute*3)
						character.RemoveAura(sim, ElementalMasteryAuraID)
						character.UpdateMajorCooldowns(sim)
					},
				})
				return true
			}
		},
	})
}

func init() {
	core.BaseStats[core.BaseStatsKey{ Race: proto.Race_RaceDraenei, Class: proto.Class_ClassShaman }] = stats.Stats{
		stats.Strength:  103,
		stats.Agility:   61,
		stats.Stamina:   113,
		stats.Intellect: 109,
		stats.Spirit:    122,
		stats.Mana:      2678,
		stats.SpellCrit: 48.576,
	}
	core.BaseStats[core.BaseStatsKey{ Race: proto.Race_RaceOrc, Class: proto.Class_ClassShaman }] = stats.Stats{
		stats.Intellect: 104,
		stats.Mana:      2678,
		stats.Spirit:    135,
		stats.SpellCrit: 48.576,
	}
	core.BaseStats[core.BaseStatsKey{ Race: proto.Race_RaceTauren, Class: proto.Class_ClassShaman }] = stats.Stats{
		stats.Intellect: 104,
		stats.Mana:      2678,
		stats.Spirit:    135,
		stats.SpellCrit: 48.576,
	}

	trollStats := stats.Stats{
		stats.Intellect: 104,
		stats.Mana:      2678,
		stats.Spirit:    135,
		stats.SpellCrit: 48.576,
	}
	core.BaseStats[core.BaseStatsKey{ Race: proto.Race_RaceTroll10, Class: proto.Class_ClassShaman }] = trollStats
	core.BaseStats[core.BaseStatsKey{ Race: proto.Race_RaceTroll30, Class: proto.Class_ClassShaman }] = trollStats
}
