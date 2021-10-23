package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func NewShaman(character core.Character, talents proto.ShamanTalents, selfBuffs SelfBuffs, rotation Rotation) *Shaman {
	shaman := &Shaman{
		Character: character,
		rotation:  rotation,
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

// Picks which attacks / abilities the Shaman does.
type Rotation interface {
	// Returns the action this rotation would like to take next.
	ChooseAction(*Shaman, *core.Simulation) core.AgentAction

	// This will be invoked right before the chosen action is actually executed, so the rotation can update its state.
	// Note that the action may be different from the action chosen by this rotation.
	OnActionAccepted(*Shaman, *core.Simulation, core.AgentAction)

	// Returns this rotation to its initial state. Called before each Sim iteration.
	Reset(*Shaman, *core.Simulation)
}

// Shaman represents a shaman character.
type Shaman struct {
	core.Character

	rotation Rotation

	Talents   proto.ShamanTalents
	SelfBuffs SelfBuffs

	ElementalFocusStacks byte

	// cache
	convectionBonus float64
	concussionBonus float64
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

func (shaman *Shaman) Reset(newsim *core.Simulation) {
	shaman.rotation.Reset(shaman, newsim)
}

func (shaman *Shaman) Act(sim *core.Simulation) time.Duration {
	newAction := shaman.rotation.ChooseAction(shaman, sim)

	actionSuccessful := newAction.Act(sim)
	if actionSuccessful {
		shaman.rotation.OnActionAccepted(shaman, sim, newAction)
		return sim.CurrentTime + core.MaxDuration(
				shaman.GetRemainingCD(core.MagicIDGCD, sim.CurrentTime),
				newAction.GetDuration())
	} else {
		// Only way for a shaman spell to fail is due to mana cost.
		// Wait until we have enough mana to cast.
		// TODO: This logic should be in ele shaman code, because enhance react differently to going oom.
		regenTime := shaman.TimeUntilManaRegen(newAction.GetManaCost())
		newAction = core.NewWaitAction(sim, shaman.GetCharacter(), regenTime)
		shaman.rotation.OnActionAccepted(shaman, sim, newAction)
		return sim.CurrentTime + regenTime
	}
}

func (shaman *Shaman) registerElementalMasteryCD() {
	if !shaman.Talents.ElementalMastery {
		return
	}

	shaman.AddMajorCooldown(core.MajorCooldown{
		CooldownID: core.MagicIDEleMastery,
		Cooldown: time.Minute*3,
		TryActivate: func(sim *core.Simulation, character *core.Character) bool {
			character.AddAura(sim, core.Aura{
				ID:      core.MagicIDEleMastery,
				Expires: core.NeverExpires,
				OnCast: func(sim *core.Simulation, cast core.DirectCastAction, input *core.DirectCastInput) {
					input.ManaCost = 0
					input.GuaranteedCrit = true
				},
				OnCastComplete: func(sim *core.Simulation, cast core.DirectCastAction) {
					// Remove the buff and put skill on CD
					character.SetCD(core.MagicIDEleMastery, time.Minute*3+sim.CurrentTime)
					character.RemoveAura(sim, core.MagicIDEleMastery)
				},
			})
			return true
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
