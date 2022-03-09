package smite

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/priest"
)

func RegisterSmitePriest() {
	core.RegisterAgentFactory(
		proto.Player_SmitePriest{},
		proto.Spec_SpecSmitePriest,
		func(character core.Character, options proto.Player) core.Agent {
			return NewSmitePriest(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_SmitePriest)
			if !ok {
				panic("Invalid spec value for Smite Priest!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewSmitePriest(character core.Character, options proto.Player) *SmitePriest {
	smiteOptions := options.GetSmitePriest()

	// Only undead can do Dev Plague
	if smiteOptions.Rotation.UseDevPlague && options.Race != proto.Race_RaceUndead {
		smiteOptions.Rotation.UseDevPlague = false
	}
	// Only nelf can do starshards
	if smiteOptions.Rotation.UseStarshards && options.Race != proto.Race_RaceNightElf {
		smiteOptions.Rotation.UseStarshards = false
	}

	selfBuffs := priest.SelfBuffs{
		UseShadowfiend: smiteOptions.Options.UseShadowfiend,
	}

	basePriest := priest.New(character, selfBuffs, *smiteOptions.Talents)
	
	spriest := &SmitePriest{
		Priest:   basePriest,
		rotation: *smiteOptions.Rotation,
	}

	return spriest
}

type SmitePriest struct {
	*priest.Priest

	rotation proto.SmitePriest_Rotation
}

func (spriest *SmitePriest) GetPriest() *priest.Priest {
	return spriest.Priest
}

func (spriest *SmitePriest) Reset(sim *core.Simulation) {
	spriest.Priest.Reset(sim)
}

// TODO: probably do something different instead of making it global?
const (
	mbidx int = iota
	swdidx
	vtidx
	swpidx
)

func (spriest *SmitePriest) OnGCDReady(sim *core.Simulation) {
	spriest.tryUseGCD(sim)
}

func (spriest *SmitePriest) OnManaTick(sim *core.Simulation) {
	if spriest.FinishedWaitingForManaAndGCDReady(sim) {
		spriest.tryUseGCD(sim)
	}
}

func (spriest *SmitePriest) tryUseGCD(sim *core.Simulation) {

	target := sim.GetPrimaryTarget()
	var spell *core.SimpleSpell

	if !spriest.SWPSpell.Effect.DotInput.IsTicking(sim) {
		spell = spriest.NewShadowWordPain(sim, target)
	} else {
		spell = spriest.NewSmite(sim, target)
	}
	
	if success := spell.Cast(sim); !success {
		spriest.WaitForMana(sim, spell.GetManaCost())
	}
}

