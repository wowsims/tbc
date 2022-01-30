package hunter

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var RotationCheckAuraID = core.NewAuraID()

func (hunter *Hunter) applyRotationAura() {
	hunter.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: RotationCheckAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.IsRanged() || !hitEffect.IsWhiteHit || hunter.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
					return
				}
				hunter.tryUseGCD(sim)
			},
		}
	})
}

func (hunter *Hunter) tryUseGCD(sim *core.Simulation) {
	if sim.CurrentTime == 0 && hunter.Rotation.PrecastAimedShot {
		hunter.NewAimedShot(sim, sim.GetPrimaryTarget()).Attack(sim)
		return
	}

	if hunter.IsWaitingForMana() {
		return
	}

	target := sim.GetPrimaryTarget()
	currentMana := hunter.CurrentManaPercent()

	if hunter.aspectOfTheViper && currentMana > hunter.Rotation.ViperStopManaPercent {
		aspect := hunter.NewAspectOfTheHawk(sim)
		aspect.StartCast(sim)
	} else if !hunter.aspectOfTheViper && currentMana < hunter.Rotation.ViperStartManaPercent {
		aspect := hunter.NewAspectOfTheViper(sim)
		aspect.StartCast(sim)
	} else if hunter.Rotation.MaintainScorpidSting && !target.HasAura(ScorpidStingDebuffID) {
		ss := hunter.NewScorpidSting(sim, target)
		if success := ss.Attack(sim); !success {
			hunter.WaitForMana(sim, ss.Cost.Value)
		}
	} else if hunter.Rotation.UseMultiShot && !hunter.IsOnCD(MultiShotCooldownID, sim.CurrentTime) {
		ms := hunter.NewMultiShot(sim)
		if success := ms.Attack(sim); !success {
			hunter.WaitForMana(sim, ms.Cost.Value)
		}
	} else if hunter.Rotation.UseArcaneShot && !hunter.IsOnCD(ArcaneShotCooldownID, sim.CurrentTime) {
		as := hunter.NewArcaneShot(sim, target)
		if success := as.Attack(sim); !success {
			hunter.WaitForMana(sim, as.Cost.Value)
		}
	} else {
		cast := hunter.NewSteadyShot(sim, target)
		if success := cast.StartCast(sim); !success {
			hunter.WaitForMana(sim, cast.GetManaCost())
		}
	}
}

func (hunter *Hunter) OnGCDReady(sim *core.Simulation) {
}

func (hunter *Hunter) OnManaTick(sim *core.Simulation) {
	if hunter.aspectOfTheViper {
		// https://wowpedia.fandom.com/wiki/Aspect_of_the_Viper?oldid=1458832
		percentMana := core.MaxFloat(0.2, core.MinFloat(0.9, hunter.CurrentManaPercent()))
		scaling := 22.0/35.0*(0.9-percentMana) + 0.11
		bonusPer5Seconds := hunter.GetStat(stats.Intellect)*scaling + 0.35*70
		manaGain := bonusPer5Seconds * 2 / 5
		hunter.AddMana(sim, manaGain, AspectOfTheViperActionID, false)
	}

	if hunter.IsWaitingForMana() && hunter.DoneWaitingForMana(sim) {
		// Don't do anything, just need to check because DoneWaitingForMana
		// updates mana metrics.
	}
}
