package enhancement

import (
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
		Bloodlust:   enhOptions.Options.Bloodlust,
		WaterShield: enhOptions.Options.WaterShield,
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
		MainHand:       enh.WeaponFromMainHand(),
		OffHand:        enh.WeaponFromOffHand(),
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
}

func (enh *EnhancementShaman) GetShaman() *shaman.Shaman {
	return enh.Shaman
}

func (enh *EnhancementShaman) Reset(sim *core.Simulation) {
	enh.Shaman.Reset(sim)
}

func (enh *EnhancementShaman) OnGCDReady(sim *core.Simulation) {
	enh.tryUseGCD(sim)
}

func (enh *EnhancementShaman) OnManaTick(sim *core.Simulation) {
	if enh.FinishedWaitingForManaAndGCDReady(sim) {
		enh.tryUseGCD(sim)
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
