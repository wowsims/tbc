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
	war.tryQueueSlam(sim)
	war.tryQueueHsCleave(sim)
}

const SlamThreshold = time.Millisecond * 500

func (war *DpsWarrior) tryQueueSlam(sim *core.Simulation) {
	if war.doSlamNext &&
		war.castSlamAt == 0 &&
		(war.AutoAttacks.MainhandSwingAt <= sim.CurrentTime || war.AutoAttacks.MainhandSwingAt == sim.CurrentTime+war.AutoAttacks.MainhandSwingSpeed()) {
		slamAt := sim.CurrentTime + war.slamLatency

		gcdAt := war.GCD.ReadyAt()
		if slamAt < gcdAt {
			if gcdAt-slamAt <= SlamThreshold {
				slamAt = gcdAt
			} else {
				return
			}
		}

		if war.CanSlam() && !war.shouldSunder(sim) {
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
		war.castSlamAt = 0
		war.doSlamNext = false
		war.SunderArmor.Cast(sim, sim.GetPrimaryTarget())
		war.tryQueueHsCleave(sim)
		return
	}

	isExecutePhase := sim.IsExecutePhase()

	if (!isExecutePhase || war.Rotation.UseSlamDuringExecute) && war.doSlamNext && war.castSlamAt != 0 {
		if sim.CurrentTime < war.castSlamAt {
			return
		} else if sim.CurrentTime == war.castSlamAt {
			war.castSlamAt = 0
			war.doSlamNext = false
			if war.CanSlam() {
				war.CastSlam(sim, sim.GetPrimaryTarget())
				war.tryQueueHsCleave(sim)
				return
			}
		} else {
			war.castSlamAt = 0
			war.doSlamNext = false
			return
		}
	}

	// If using a GCD will clip the next slam, only allow high priority spells like BT/MS/WW/debuffs.
	highPrioSpellsOnly := war.Rotation.UseSlam && sim.CurrentTime+core.GCDDefault-SlamThreshold > war.AutoAttacks.MainhandSwingAt+war.slamLatency

	if isExecutePhase {
		war.executeRotation(sim, highPrioSpellsOnly)
	} else {
		war.normalRotation(sim, highPrioSpellsOnly)
	}

	if war.Rotation.UseSlam {
		war.doSlamNext = true
	} else if war.GCD.IsReady(sim) {
		// We didn't cast anything, so wait for the next CD.
		// Note that BT/MS share a CD timer so we don't need to check MS.
		nextCD := core.MinDuration(war.Bloodthirst.CD.ReadyAt(), war.Whirlwind.CD.ReadyAt())
		if nextCD > sim.CurrentTime {
			war.WaitUntil(sim, nextCD)
		}
	}
}

func (war *DpsWarrior) normalRotation(sim *core.Simulation, highPrioSpellsOnly bool) {
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
		} else if !highPrioSpellsOnly {
			if war.tryMaintainDebuffs(sim) {
				// Do nothing, already cast
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
			} else if war.Rotation.UseHamstring && war.CurrentRage() >= war.Rotation.HamstringRageThreshold && war.ShouldHamstring(sim) {
				war.Hamstring.Cast(sim, sim.GetPrimaryTarget())
			}
		}
	}

	war.tryQueueHsCleave(sim)
}

func (war *DpsWarrior) executeRotation(sim *core.Simulation, highPrioSpellsOnly bool) {
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
		} else if !highPrioSpellsOnly {
			if war.tryMaintainDebuffs(sim) {
				// Do nothing, already cast
			} else if war.CanExecute() {
				war.Execute.Cast(sim, sim.GetPrimaryTarget())
			} else if war.ShouldBerserkerRage(sim) {
				war.BerserkerRage.Cast(sim, nil)
			}
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
		war.CanSunderArmor(sim) &&
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
	if war.ShouldQueueHSOrCleave(sim) {
		war.QueueHSOrCleave(sim)
	}
}
