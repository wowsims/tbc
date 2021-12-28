package enhancement

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/shaman"
)

func RegisterEnhancementShaman() {
	core.RegisterAgentFactory(
		proto.Player_EnhancementShaman{},
		func(character core.Character, options proto.Player) core.Agent {
			return NewEnhancementShaman(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_EnhancementShaman)
			if !ok {
				panic("Invalid spec value for Enhancement Shaman!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewEnhancementShaman(character core.Character, options proto.Player) *EnhancementShaman {
	enhOptions := options.GetEnhancementShaman()

	selfBuffs := shaman.SelfBuffs{}

	if enhOptions.Rotation.Totems != nil {
		selfBuffs.StrengthOfEarth = enhOptions.Rotation.Totems.Earth == proto.EarthTotem_StrengthOfEarthTotem
		selfBuffs.GraceOfAir = enhOptions.Rotation.Totems.Air == proto.AirTotem_GraceOfAirTotem
		selfBuffs.ManaSpring = enhOptions.Rotation.Totems.Water == proto.WaterTotem_ManaSpringTotem
	}
	enh := &EnhancementShaman{
		Shaman: shaman.NewShaman(character, *enhOptions.Talents, selfBuffs),
	}
	// Enable Auto Attacks for this spec
	enh.EnableAutoAttacks()

	// TODO: de-sync dual weapons swing timers?

	// Modify auto attacks multiplier from weapon mastery.
	enh.AutoAttacks.DamageMultiplier *= 1 + 0.02*float64(enhOptions.Talents.WeaponMastery)
	shaman.ApplyWindfuryImbue(enh.Shaman, true, true)

	return enh
}

type EnhancementShaman struct {
	*shaman.Shaman
}

func (enh *EnhancementShaman) GetShaman() *shaman.Shaman {
	return enh.Shaman
}

func (enh *EnhancementShaman) Reset(sim *core.Simulation) {
	enh.Shaman.Reset(sim)
}

func (enh *EnhancementShaman) Act(sim *core.Simulation) time.Duration {
	// Redrop totems when needed.
	dropTime := enh.TryDropTotems(sim)
	if dropTime > 0 {
		return sim.CurrentTime + enh.AutoAttacks.TimeUntil(sim, nil, nil, dropTime)
	}

	if enh.GetRemainingCD(shaman.StormstrikeCD, sim.CurrentTime) == 0 {
		ss := enh.NewStormstrike(sim, sim.GetPrimaryTarget())
		ss.Attack(sim)
		return sim.CurrentTime + enh.AutoAttacks.TimeUntil(sim, nil, ss, 0)
	} else if enh.GetRemainingCD(shaman.ShockCooldownID, sim.CurrentTime) == 0 {
		shock := enh.NewEarthShock(sim, sim.GetPrimaryTarget())
		shock.Cast(sim)
		return sim.CurrentTime + enh.AutoAttacks.TimeUntil(sim, shock, nil, 0)
	}

	// Do nothing, just swing axes until next CD available
	nextCD := enh.GetRemainingCD(shaman.StormstrikeCD, sim.CurrentTime)
	shockCD := enh.GetRemainingCD(shaman.ShockCooldownID, sim.CurrentTime)
	if shockCD < nextCD {
		nextCD = shockCD
	}
	return sim.CurrentTime + enh.AutoAttacks.TimeUntil(sim, nil, nil, nextCD)
}
