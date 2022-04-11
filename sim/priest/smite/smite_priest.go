package smite

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/priest"
	"time"
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

	if smiteOptions.Options.PowerInfusionTarget != nil {
		selfBuffs.PowerInfusionTarget = *smiteOptions.Options.PowerInfusionTarget
	} else {
		selfBuffs.PowerInfusionTarget.TargetIndex = -1
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

	// Calculate higher SW:P uptime if using HF
	swpRemaining := spriest.ShadowWordPainDot.RemainingDuration(sim)

	castSpeed := spriest.CastSpeed()

	// smite cast time, talent assumed
	smiteCastTime := time.Duration(float64(time.Millisecond*2000) / castSpeed)

	// holy fire cast time
	hfCastTime := time.Duration(float64(time.Millisecond*3000) / castSpeed)

	var spell *core.Spell
	// Always attempt to keep SW:P up if its down
	if !spriest.ShadowWordPainDot.IsActive() {
		spell = spriest.ShadowWordPain
		// Favor star shards for NE if off cooldown first
	} else if spriest.rotation.UseStarshards && spriest.GetRemainingCD(priest.SSCooldownID, sim.CurrentTime) == 0 {
		spell = spriest.Starshards
		// Allow for undead to use devouring plague off CD
	} else if spriest.rotation.UseDevPlague && spriest.GetRemainingCD(priest.DevouringPlagueCooldownID, sim.CurrentTime) == 0 {
		spell = spriest.DevouringPlague
		// If setting enabled, throw mind blast into our rotation off CD
	} else if spriest.rotation.UseMindBlast && spriest.Character.GetRemainingCD(priest.MBCooldownID, sim.CurrentTime) == 0 {
		spell = spriest.MindBlast
		// If setting enabled, cast Shadow Word: Death on cooldown
	} else if spriest.rotation.UseShadowWordDeath && spriest.Character.GetRemainingCD(priest.SWDCooldownID, sim.CurrentTime) == 0 {
		spell = spriest.ShadowWordDeath
		// Consider HF if SWP will fall off after 1 smite but before 2 smites from now finishes
		//	and swp falls off after hf finishes (assumption never worth clipping)
	} else if spriest.rotation.RotationType == proto.SmitePriest_Rotation_HolyFireWeave && swpRemaining > smiteCastTime && swpRemaining < hfCastTime {
		spell = spriest.HolyFire
		// Base filler spell is smite
	} else {
		spell = spriest.Smite
	}

	if success := spell.Cast(sim, sim.GetPrimaryTarget()); !success {
		spriest.WaitForMana(sim, spell.CurCast.Cost)
	}
}
