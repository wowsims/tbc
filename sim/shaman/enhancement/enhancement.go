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

	// We need to directly manage all GCD-bound CDs ourself.
	if enh.Consumes.Drums == proto.Drums_DrumsOfBattle {
		enh.scheduler.ScheduleMCD(sim, enh.GetCharacter(), core.DrumsOfBattleActionID)
	} else if enh.Consumes.Drums == proto.Drums_DrumsOfRestoration {
		enh.scheduler.ScheduleMCD(sim, enh.GetCharacter(), core.DrumsOfRestorationActionID)
	}
	enh.scheduler.ScheduleMCD(sim, enh.GetCharacter(), enh.BloodlustActionID())

	scheduleTotem := func(duration time.Duration, prioritizeEarlier bool, precast bool, tryCast func(sim *core.Simulation) (bool, float64)) {
		totemAction := common.ScheduledAbility{
			Duration: time.Second * 1,
			TryCast: func(sim *core.Simulation) bool {
				success, manaCost := tryCast(sim)
				if !success {
					enh.WaitForMana(sim, manaCost)
				}
				return success
			},
			PrioritizeEarlierForConflicts: prioritizeEarlier,
		}

		curTime := time.Duration(0)
		if precast {
			curTime = duration
		}
		for curTime <= maxDuration {
			ability := totemAction
			ability.DesiredCastAt = curTime
			if prioritizeEarlier {
				ability.MinCastAt = curTime - time.Second*10
				ability.MaxCastAt = curTime + time.Second*10
			} else {
				ability.MinCastAt = curTime
				ability.MaxCastAt = curTime + time.Second*10
			}
			castAt := enh.scheduler.Schedule(ability)
			if castAt == common.Unresolved {
				panic("No timeslot found for totem")
			}
			curTime = castAt + duration
		}
	}
	scheduleSpellTotem := func(duration time.Duration, castFactory func(sim *core.Simulation) *core.SimpleSpell) {
		scheduleTotem(duration, false, false, func(sim *core.Simulation) (bool, float64) {
			cast := castFactory(sim)
			return cast.Cast(sim), cast.ManaCost
		})
	}
	schedule2MTotem := func(castFactory func(sim *core.Simulation) *core.SimpleCast) {
		scheduleTotem(time.Minute*2, true, true, func(sim *core.Simulation) (bool, float64) {
			cast := castFactory(sim)
			return cast.StartCast(sim), cast.ManaCost
		})
	}

	if enh.Totems.TwistFireNova {
		var defaultCastFactory func(sim *core.Simulation)
		switch enh.Totems.Fire {
		case proto.FireTotem_MagmaTotem:
			defaultCastFactory = func(sim *core.Simulation) {
				if enh.FireTotemSpell.IsInUse() {
					return
				}

				cast := enh.NewMagmaTotem(sim)
				success := cast.Cast(sim)
				if !success {
					enh.WaitForMana(sim, cast.ManaCost)
				}
			}
		case proto.FireTotem_SearingTotem:
			defaultCastFactory = func(sim *core.Simulation) {
				if enh.FireTotemSpell.IsInUse() {
					return
				}

				cast := enh.NewSearingTotem(sim, sim.GetPrimaryTarget())
				success := cast.Cast(sim)
				if !success {
					enh.WaitForMana(sim, cast.ManaCost)
				}
			}
		case proto.FireTotem_TotemOfWrath:
			defaultCastFactory = func(sim *core.Simulation) {
				if enh.NextTotemDrops[shaman.FireTotem] > sim.CurrentTime+time.Second*5 {
					// Skip dropping if we've gone OOM reverted to dropping default only, and have plenty of time left.
					return
				}

				cast := enh.NewTotemOfWrath(sim)
				success := cast.StartCast(sim)
				if !success {
					enh.WaitForMana(sim, cast.ManaCost)
				}
			}
		}

		fntAction := common.ScheduledAbility{
			Duration: time.Second * 1,
			TryCast: func(sim *core.Simulation) bool {
				if enh.Metrics.WentOOM && enh.CurrentManaPercent() < 0.2 {
					return false
				}

				cast := enh.NewNovaTotem(sim)
				success := cast.Cast(sim)
				if !success {
					enh.WaitForMana(sim, cast.ManaCost)
				}
				return success
			},
		}
		defaultAction := common.ScheduledAbility{
			Duration: time.Second * 1,
			TryCast: func(sim *core.Simulation) bool {
				defaultCastFactory(sim)
				return true
			},
		}

		curTime := time.Duration(0)
		nextNovaCD := time.Duration(0)
		defaultNext := false
		for curTime <= maxDuration {
			ability := fntAction
			if defaultNext {
				ability = defaultAction
			}
			ability.DesiredCastAt = curTime
			ability.MinCastAt = curTime
			ability.MaxCastAt = curTime + time.Second*15

			castAt := enh.scheduler.Schedule(ability)

			if defaultNext {
				curTime = nextNovaCD
				defaultNext = false
			} else {
				nextNovaCD = castAt + time.Second*15
				if defaultCastFactory == nil {
					curTime = nextNovaCD
				} else {
					curTime = castAt + enh.FireNovaTickLength()
					defaultNext = true
				}
			}
		}
	} else {
		switch enh.Totems.Fire {
		case proto.FireTotem_MagmaTotem:
			scheduleSpellTotem(time.Second*20, func(sim *core.Simulation) *core.SimpleSpell { return enh.NewMagmaTotem(sim) })
		case proto.FireTotem_SearingTotem:
			scheduleSpellTotem(time.Minute*1, func(sim *core.Simulation) *core.SimpleSpell { return enh.NewSearingTotem(sim, sim.GetPrimaryTarget()) })
		case proto.FireTotem_TotemOfWrath:
			schedule2MTotem(func(sim *core.Simulation) *core.SimpleCast { return enh.NewTotemOfWrath(sim) })
		}
	}

	if enh.Totems.Air != proto.AirTotem_NoAirTotem {
		var defaultCastFactory func(sim *core.Simulation) *core.SimpleCast
		switch enh.Totems.Air {
		case proto.AirTotem_GraceOfAirTotem:
			defaultCastFactory = func(sim *core.Simulation) *core.SimpleCast { return enh.NewGraceOfAirTotem(sim) }
		case proto.AirTotem_TranquilAirTotem:
			defaultCastFactory = func(sim *core.Simulation) *core.SimpleCast { return enh.NewTranquilAirTotem(sim) }
		case proto.AirTotem_WindfuryTotem:
			defaultCastFactory = func(sim *core.Simulation) *core.SimpleCast { return enh.NewWindfuryTotem(sim) }
		case proto.AirTotem_WrathOfAirTotem:
			defaultCastFactory = func(sim *core.Simulation) *core.SimpleCast { return enh.NewWrathOfAirTotem(sim) }
		}

		if enh.Totems.TwistWindfury {
			wfAction := common.ScheduledAbility{
				Duration: time.Second * 1,
				TryCast: func(sim *core.Simulation) bool {
					if enh.Metrics.WentOOM && enh.CurrentManaPercent() < 0.2 {
						return false
					}

					cast := enh.NewWindfuryTotem(sim)
					success := cast.StartCast(sim)
					if !success {
						enh.WaitForMana(sim, cast.ManaCost)
					}
					return success
				},
				PrioritizeEarlierForConflicts: true,
			}
			defaultAction := common.ScheduledAbility{
				Duration: time.Second * 1,
				TryCast: func(sim *core.Simulation) bool {
					if enh.NextTotemDrops[shaman.AirTotem] > sim.CurrentTime+time.Second*10 {
						// Skip dropping if we've gone OOM reverted to dropping default only, and have plenty of time left.
						return true
					}

					cast := defaultCastFactory(sim)
					success := cast.StartCast(sim)
					if !success {
						enh.WaitForMana(sim, cast.ManaCost)
					}
					return success
				},
			}

			curTime := time.Second * 10
			for curTime <= maxDuration {
				ability := wfAction
				ability.DesiredCastAt = curTime
				ability.MinCastAt = curTime - time.Second*8
				ability.MaxCastAt = curTime + time.Second*4
				defaultAbility := defaultAction
				castAt := enh.scheduler.ScheduleGroup(sim, []common.ScheduledAbility{ability, defaultAbility})
				if castAt == common.Unresolved {
					panic("No timeslot found for air totem")
				}
				curTime = castAt + time.Second*10
			}
		} else {
			schedule2MTotem(defaultCastFactory)
		}
	}

	if enh.Totems.Earth != proto.EarthTotem_NoEarthTotem {
		switch enh.Totems.Earth {
		case proto.EarthTotem_StrengthOfEarthTotem:
			schedule2MTotem(func(sim *core.Simulation) *core.SimpleCast { return enh.NewStrengthOfEarthTotem(sim) })
		case proto.EarthTotem_TremorTotem:
			schedule2MTotem(func(sim *core.Simulation) *core.SimpleCast { return enh.NewTremorTotem(sim) })
		}
	}

	if enh.Totems.Water != proto.WaterTotem_NoWaterTotem {
		if enh.Totems.Water == proto.WaterTotem_ManaSpringTotem {
			schedule2MTotem(func(sim *core.Simulation) *core.SimpleCast { return enh.NewManaSpringTotem(sim) })
		}
	}
}

func (enh *EnhancementShaman) Reset(sim *core.Simulation) {
	enh.Shaman.Reset(sim)
	enh.scheduler.Reset(sim, enh.GetCharacter())
}

func (enh *EnhancementShaman) OnGCDReady(sim *core.Simulation) {
	enh.scheduler.DoNextAbility(sim, &enh.Character)
}

func (enh *EnhancementShaman) OnManaTick(sim *core.Simulation) {
	if enh.IsWaitingForMana() && !enh.DoneWaitingForMana(sim) {
		// Do nothing, just need to check so metrics get updated.
	}
}
