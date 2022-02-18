package enhancement

import (
	"time"

	"github.com/wowsims/tbc/sim/common"
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

	selfBuffs := shaman.SelfBuffs{
		Bloodlust:        enhOptions.Options.Bloodlust,
		WaterShield:      enhOptions.Options.WaterShield,
		SnapshotSOET42Pc: enhOptions.Options.SnapshotT4_2Pc,
	}

	totems := proto.ShamanTotems{}
	if enhOptions.Rotation.Totems != nil {
		totems = *enhOptions.Rotation.Totems
	}
	enh := &EnhancementShaman{
		Shaman:   shaman.NewShaman(character, *enhOptions.Talents, totems, selfBuffs),
		Rotation: *enhOptions.Rotation,
	}
	// Enable Auto Attacks for this spec
	enh.EnableAutoAttacks(enh, core.AutoAttackOptions{
		MainHand:       enh.WeaponFromMainHand(enh.DefaultMeleeCritMultiplier()),
		OffHand:        enh.WeaponFromOffHand(enh.DefaultMeleeCritMultiplier()),
		AutoSwingMelee: true,
		DelayOHSwings:  enhOptions.Options.DelayOffhandSwings,
	})

	if !enh.HasMHWeapon() {
		enhOptions.Options.MainHandImbue = proto.ShamanWeaponImbue_ImbueNone
	}
	if !enh.HasOHWeapon() {
		enhOptions.Options.OffHandImbue = proto.ShamanWeaponImbue_ImbueNone
	}
	enh.ApplyWindfuryImbue(
		enhOptions.Options.MainHandImbue == proto.ShamanWeaponImbue_ImbueWindfury,
		enhOptions.Options.OffHandImbue == proto.ShamanWeaponImbue_ImbueWindfury)
	enh.ApplyFlametongueImbue(
		enhOptions.Options.MainHandImbue == proto.ShamanWeaponImbue_ImbueFlametongue,
		enhOptions.Options.OffHandImbue == proto.ShamanWeaponImbue_ImbueFlametongue)
	enh.ApplyFrostbrandImbue(
		enhOptions.Options.MainHandImbue == proto.ShamanWeaponImbue_ImbueFrostbrand,
		enhOptions.Options.OffHandImbue == proto.ShamanWeaponImbue_ImbueFrostbrand)
	enh.ApplyRockbiterImbue(
		enhOptions.Options.MainHandImbue == proto.ShamanWeaponImbue_ImbueRockbiter,
		enhOptions.Options.OffHandImbue == proto.ShamanWeaponImbue_ImbueRockbiter)

	if enhOptions.Options.MainHandImbue != proto.ShamanWeaponImbue_ImbueNone {
		enh.HasMHWeaponImbue = true
	}

	return enh
}

type EnhancementShaman struct {
	*shaman.Shaman

	Rotation proto.EnhancementShaman_Rotation

	scheduler common.GCDScheduler
}

func (enh *EnhancementShaman) GetShaman() *shaman.Shaman {
	return enh.Shaman
}

func (enh *EnhancementShaman) Init(sim *core.Simulation) {
	enh.Shaman.Init(sim)

	// Fill the GCD schedule based on our settings.
	maxDuration := sim.GetMaxDuration()

	var curTime time.Duration

	ssAction := common.ScheduledAbility{
		Duration: core.GCDDefault,
		TryCast: func(sim *core.Simulation) bool {
			ss := enh.NewStormstrike(sim, sim.GetPrimaryTarget())
			success := ss.Attack(sim)
			if !success {
				enh.WaitForMana(sim, ss.Cost.Value)
			}
			return success
		},
	}
	curTime = 0
	for curTime <= maxDuration {
		ability := ssAction
		ability.DesiredCastAt = curTime
		castAt := enh.scheduler.Schedule(ability)
		curTime = castAt + time.Second*10
	}

	shockCD := enh.ShockCD()
	shockAction := common.ScheduledAbility{
		Duration: core.GCDDefault,
		TryCast: func(sim *core.Simulation) bool {
			var shock *core.SimpleSpell
			target := sim.GetPrimaryTarget()
			if enh.Rotation.WeaveFlameShock && !enh.FlameShockSpell.IsInUse() {
				shock = enh.NewFlameShock(sim, target)
			} else if enh.Rotation.PrimaryShock == proto.EnhancementShaman_Rotation_Earth {
				shock = enh.NewEarthShock(sim, target)
			} else if enh.Rotation.PrimaryShock == proto.EnhancementShaman_Rotation_Frost {
				shock = enh.NewFrostShock(sim, target)
			}

			success := shock.Cast(sim)
			if !success {
				enh.WaitForMana(sim, shock.ManaCost)
			}
			return success
		},
	}
	if enh.Rotation.PrimaryShock != proto.EnhancementShaman_Rotation_None {
		curTime = 0
		for curTime <= maxDuration {
			ability := shockAction
			ability.DesiredCastAt = curTime
			ability.MinCastAt = curTime
			ability.MaxCastAt = curTime + time.Second*10
			castAt := enh.scheduler.Schedule(ability)
			curTime = castAt + shockCD
		}
	} else if enh.Rotation.WeaveFlameShock {
		// Flame shock but no regular shock, so only use it once every 12s.
		curTime = 0
		for curTime <= maxDuration {
			ability := shockAction
			ability.DesiredCastAt = curTime
			ability.MinCastAt = curTime
			ability.MaxCastAt = curTime + time.Second*10
			castAt := enh.scheduler.Schedule(ability)
			curTime = castAt + time.Second*12
		}
	}
}

func (enh *EnhancementShaman) Reset(sim *core.Simulation) {
	enh.Shaman.Reset(sim)
	enh.scheduler.Reset(sim, enh.GetCharacter())
}

func (enh *EnhancementShaman) OnGCDReady(sim *core.Simulation) {
	enh.tryUseGCD(sim)
}

func (enh *EnhancementShaman) OnManaTick(sim *core.Simulation) {
	if enh.IsWaitingForMana() && !enh.DoneWaitingForMana(sim) {
		// Do nothing, just need to check so metrics get updated.
	}
}

func (enh *EnhancementShaman) tryUseGCD(sim *core.Simulation) {

	target := sim.GetPrimaryTarget()

	if enh.Talents.Stormstrike && !enh.IsOnCD(shaman.StormstrikeCD, sim.CurrentTime) {
		ss := enh.NewStormstrike(sim, target)
		if success := ss.Attack(sim); !success {
			enh.WaitForMana(sim, ss.Cost.Value)
		}
		return
	} else if !enh.IsOnCD(shaman.ShockCooldownID, sim.CurrentTime) {
		var shock *core.SimpleSpell
		if enh.Rotation.WeaveFlameShock && !enh.FlameShockSpell.IsInUse() {
			shock = enh.NewFlameShock(sim, target)
		} else if enh.Rotation.PrimaryShock == proto.EnhancementShaman_Rotation_Earth {
			shock = enh.NewEarthShock(sim, target)
		} else if enh.Rotation.PrimaryShock == proto.EnhancementShaman_Rotation_Frost {
			shock = enh.NewFrostShock(sim, target)
		}

		if shock != nil {
			if success := shock.Cast(sim); !success {
				enh.WaitForMana(sim, shock.ManaCost)
			}
			return
		}
	}

	// Redrop totems when needed.
	if enh.TryDropTotems(sim) {
		return
	}

	// We didn't try to cast anything. Just wait for next auto or CD.
	nextEventAt := enh.NextTotemAt(sim)
	if enh.Talents.Stormstrike {
		nextEventAt = core.MinDuration(nextEventAt, enh.CDReadyAt(shaman.StormstrikeCD))
	}
	if enh.Rotation.PrimaryShock != proto.EnhancementShaman_Rotation_None {
		nextEventAt = core.MinDuration(nextEventAt, enh.CDReadyAt(shaman.ShockCooldownID))
	}
	enh.WaitUntil(sim, nextEventAt)
}
