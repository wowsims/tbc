package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func NewShaman(character core.Character, talents proto.ShamanTalents, selfBuffs SelfBuffs, rotation Rotation) *Shaman {
	if selfBuffs.WaterShield {
		character.InitialStats[stats.MP5] += 50
	}

	return &Shaman{
		Character: character,
		rotation:  rotation,
		Talents:   talents,
		SelfBuffs: selfBuffs,

		convectionBonus: 0.02 * float64(talents.Convection),
		concussionBonus: 1 + 0.01*float64(talents.Concussion),
	}
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
	// Returns the action this Agent would like to take next.
	ChooseAction(*Shaman, *core.Simulation) core.AgentAction

	// This will be invoked if the chosen action is actually executed, so the Agent can update its state.
	OnActionAccepted(*Shaman, *core.Simulation, core.AgentAction)

	// Returns this Agent to its initial state.
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

func (shaman *Shaman) BuffUp(sim *core.Simulation) {
	shaman.Stats[stats.MP5] += shaman.Stats[stats.Intellect] * (0.02 * float64(shaman.Talents.UnrelentingStorm))
}

func (shaman *Shaman) ChooseAction(sim *core.Simulation) core.AgentAction {
	// Before casting, activate shaman powers!
	TryActivateBloodlust(sim, shaman)
	if shaman.Talents.ElementalMastery {
		TryActivateEleMastery(sim, shaman)
	}

	return shaman.rotation.ChooseAction(shaman, sim)
}
func (shaman *Shaman) OnActionAccepted(sim *core.Simulation, action core.AgentAction) {
	shaman.rotation.OnActionAccepted(shaman, sim, action)
}
func (shaman *Shaman) Reset(newsim *core.Simulation) {
	shaman.rotation.Reset(shaman, newsim)
}

func TryActivateBloodlust(sim *core.Simulation, shaman *Shaman) {
	if shaman.IsOnCD(core.MagicIDBloodlust, sim.CurrentTime) {
		return
	}

	dur := time.Second * 40 // assumes that multiple BLs are different shaman.
	shaman.SetCD(core.MagicIDBloodlust, time.Minute*10+sim.CurrentTime)

	shaman.Party.AddAura(sim, core.Aura{
		ID:      core.MagicIDBloodlust,
		Expires: sim.CurrentTime + dur,
		OnCast: func(sim *core.Simulation, cast core.DirectCastAction, input *core.DirectCastInput) {
			input.CastTime = (input.CastTime * 10) / 13 // 30% faster
		},
	})
}

func TryActivateEleMastery(sim *core.Simulation, shaman *Shaman) {
	if shaman.IsOnCD(core.MagicIDEleMastery, sim.CurrentTime) {
		return
	}

	shaman.AddAura(sim, core.Aura{
		ID:      core.MagicIDEleMastery,
		Expires: core.NeverExpires,
		OnCast: func(sim *core.Simulation, cast core.DirectCastAction, input *core.DirectCastInput) {
			input.ManaCost = 0
			input.GuaranteedCrit = true
		},
		OnCastComplete: func(sim *core.Simulation, cast core.DirectCastAction) {
			// Remove the buff and put skill on CD
			shaman.SetCD(core.MagicIDEleMastery, time.Second*180+sim.CurrentTime)
			shaman.RemoveAura(sim, core.MagicIDEleMastery)
		},
	})
}

func init() {
	core.BaseStats[core.BaseStatsKey{ Race: core.RaceBonusTypeDraenei, Class: proto.Class_ClassShaman }] = stats.Stats{
		stats.Strength:  103,
		stats.Agility:   61,
		stats.Stamina:   113,
		stats.Intellect: 109,
		stats.Spirit:    122,
		stats.Mana:      2678,
		stats.SpellCrit: 48.576,
	}
	core.BaseStats[core.BaseStatsKey{ Race: core.RaceBonusTypeOrc, Class: proto.Class_ClassShaman }] = stats.Stats{
		stats.Intellect: 104,
		stats.Mana:      2678,
		stats.Spirit:    135,
		stats.SpellCrit: 48.576,
	}
	core.BaseStats[core.BaseStatsKey{ Race: core.RaceBonusTypeTauren, Class: proto.Class_ClassShaman }] = stats.Stats{
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
	core.BaseStats[core.BaseStatsKey{ Race: core.RaceBonusTypeTroll10, Class: proto.Class_ClassShaman }] = trollStats
	core.BaseStats[core.BaseStatsKey{ Race: core.RaceBonusTypeTroll30, Class: proto.Class_ClassShaman }] = trollStats
}
