package dps

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/warrior"
)

func (war *DpsWarrior) OnGCDReady(sim *core.Simulation) {
	war.doRotation(sim)
}

func (war *DpsWarrior) OnAutoAttack(sim *core.Simulation, spell *core.Spell) {
	if war.doSlamNext && war.castSlamAt == 0 && spell != war.AutoAttacks.OHAuto {
		slamAt := sim.CurrentTime + war.slamLatency
		if slamAt >= war.GCD.ReadyAt() && war.CanSlam() {
			war.castSlamAt = slamAt
			war.WaitUntil(sim, slamAt) // Pause GCD until slam time
		}
	}
}

func (war *DpsWarrior) doRotation(sim *core.Simulation) {
	if !war.StanceMatches(warrior.BerserkerStance) && war.BerserkerStance.IsReady(sim) {
		war.BerserkerStance.Cast(sim, nil)
	}
	if war.shouldSunder(sim) {
		war.SunderArmor.Cast(sim, sim.GetPrimaryTarget())
		war.tryQueueHsCleave(sim)
		return
	}

	if war.doSlamNext {
		if war.castSlamAt == 0 && war.AutoAttacks.MainhandSwingAt < sim.CurrentTime+time.Millisecond*100 {
			return
		} else {
			if sim.CurrentTime > war.castSlamAt {
				war.castSlamAt = 0
				war.doSlamNext = false
			} else if war.CanSlam() {
				war.Slam.Cast(sim, sim.GetPrimaryTarget())
				war.castSlamAt = 0
				war.doSlamNext = false
				war.tryQueueHsCleave(sim)
				return
			}
		}
	}

	if sim.IsExecutePhase() {
		war.executeRotation(sim)
	} else {
		war.normalRotation(sim)
	}

	if war.Rotation.UseSlam {
		war.doSlamNext = true
	}
}

func (war *DpsWarrior) normalRotation(sim *core.Simulation) {
	if war.GCD.IsReady(sim) {
		if war.ShouldRampage(sim) {
			war.Rampage.Cast(sim, nil)
		} else if war.Rotation.PrioritizeWw && war.CanWhirlwind(sim) {
			war.Whirlwind.Cast(sim, sim.GetPrimaryTarget())
		} else if war.CanBloodthirst(sim) {
			war.Bloodthirst.Cast(sim, sim.GetPrimaryTarget())
		} else if war.CanMortalStrike(sim) {
			war.MortalStrike.Cast(sim, sim.GetPrimaryTarget())
		} else if !war.Rotation.PrioritizeWw && war.CanWhirlwind(sim) {
			war.Whirlwind.Cast(sim, sim.GetPrimaryTarget())
		} else if war.Rotation.UseOverpower && war.CurrentRage() < war.Rotation.OverpowerRageThreshold && war.ShouldOverpower(sim) {
			if !war.StanceMatches(warrior.BattleStance) {
				if !war.BattleStance.IsReady(sim) {
					return
				}
				war.BattleStance.Cast(sim, nil)
			}
			war.Overpower.Cast(sim, sim.GetPrimaryTarget())
		} else if war.ShouldBerserkerRage(sim) {
			war.BerserkerRage.Cast(sim, nil)
		} else if war.tryMaintainDebuffs(sim) {
			// Do nothing, already cast
		} else if war.Rotation.UseHamstring && war.CurrentRage() >= war.Rotation.HamstringRageThreshold && war.ShouldHamstring(sim) {
			war.Hamstring.Cast(sim, sim.GetPrimaryTarget())
		}
	}

	war.tryQueueHsCleave(sim)
}

func (war *DpsWarrior) executeRotation(sim *core.Simulation) {
	if war.GCD.IsReady(sim) {
		if war.ShouldRampage(sim) {
			war.Rampage.Cast(sim, nil)
		} else if war.Rotation.PrioritizeWw && war.Rotation.UseWwDuringExecute && war.CanWhirlwind(sim) {
			war.Whirlwind.Cast(sim, sim.GetPrimaryTarget())
		} else if war.Rotation.UseBtDuringExecute && war.CanBloodthirst(sim) {
			war.Bloodthirst.Cast(sim, sim.GetPrimaryTarget())
		} else if war.Rotation.UseMsDuringExecute && war.CanMortalStrike(sim) {
			war.MortalStrike.Cast(sim, sim.GetPrimaryTarget())
		} else if !war.Rotation.PrioritizeWw && war.Rotation.UseWwDuringExecute && war.CanWhirlwind(sim) {
			war.Whirlwind.Cast(sim, sim.GetPrimaryTarget())
		} else if war.tryMaintainDebuffs(sim) {
			// Do nothing, already cast
		} else if war.CanExecute() {
			war.Execute.Cast(sim, sim.GetPrimaryTarget())
		} else if war.ShouldBerserkerRage(sim) {
			war.BerserkerRage.Cast(sim, nil)
		}
	}

	if war.Rotation.UseHsDuringExecute {
		war.tryQueueHsCleave(sim)
	}
}

func (war *DpsWarrior) shouldSunder(sim *core.Simulation) bool {
	if war.Rotation.SunderArmor == proto.Warrior_Rotation_SunderArmorOnce && !war.castFirstSunder && war.CanSunderArmor(sim) {
		war.castFirstSunder = true
		return true
	} else if war.Rotation.SunderArmor == proto.Warrior_Rotation_SunderArmorMaintain &&
		!war.CanSunderArmor(sim) &&
		(war.SunderArmorAura.GetStacks() < 5 || war.SunderArmorAura.RemainingDuration(sim) < time.Second*3) {
		return true
	}
	return false
}

// Returns whether any ability was cast.
func (war *DpsWarrior) tryMaintainDebuffs(sim *core.Simulation) bool {
	if war.ShouldShout(sim) {
		war.Shout.Cast(sim, nil)
		return true
	} else if war.Rotation.MaintainDemoShout && war.DemoralizingShoutAura.RemainingDuration(sim) < time.Second*2 && war.CanDemoralizingShout(sim) {
		war.DemoralizingShout.Cast(sim, sim.GetPrimaryTarget())
		return true
	} else if war.Rotation.MaintainThunderClap && war.ThunderClapAura.RemainingDuration(sim) < time.Second*2 && war.CanThunderClapIgnoreStance(sim) {
		if !war.StanceMatches(warrior.BattleStance) {
			if !war.BattleStance.IsReady(sim) {
				return false
			}
			war.BattleStance.Cast(sim, nil)
		}
		war.ThunderClap.Cast(sim, sim.GetPrimaryTarget())
		return true
	}
	return false
}

func (war *DpsWarrior) tryQueueHsCleave(sim *core.Simulation) {
	if war.CurrentRage() >= float64(war.Rotation.HsRageThreshold) {
		if war.Rotation.UseCleave {
			if war.CanCleave(sim) {
				war.QueueCleave(sim)
			}
		} else {
			if war.CanHeroicStrike(sim) {
				war.QueueHeroicStrike(sim)
			}
		}
	}
}
