package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func (mage *Mage) Act(sim *core.Simulation) time.Duration {
	var spell *core.SimpleSpell
	if mage.RotationType == proto.Mage_Rotation_Arcane {
		spell = mage.doArcaneRotation(sim)
	} else if mage.RotationType == proto.Mage_Rotation_Fire {
		spell = mage.doFireRotation(sim)
	} else {
		spell = mage.doFrostRotation(sim)
	}

	actionSuccessful := spell.Cast(sim)

	if !actionSuccessful {
		regenTime := mage.TimeUntilManaRegen(spell.GetManaCost())
		waitAction := core.NewWaitAction(sim, mage.GetCharacter(), regenTime, core.WaitReasonOOM)
		waitAction.Cast(sim)
		return sim.CurrentTime + waitAction.GetDuration()
	}

	return sim.CurrentTime + core.MaxDuration(
		mage.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime),
		spell.CastTime)
}

func (mage *Mage) doArcaneRotation(sim *core.Simulation) *core.SimpleSpell {
	target := sim.GetPrimaryTarget()
	spell := mage.NewArcaneBlast(sim, target)
	return spell
}

func (mage *Mage) doFireRotation(sim *core.Simulation) *core.SimpleSpell {
	target := sim.GetPrimaryTarget()

	if mage.FireRotation.MaintainImprovedScorch && (target.NumStacks(core.ImprovedScorchDebuffID) < 5 || target.RemainingAuraDuration(sim, core.ImprovedScorchDebuffID) < time.Millisecond*5500) {
		return mage.NewScorch(sim, target)
	}

	if mage.FireRotation.PrimarySpell == proto.Mage_Rotation_FireRotation_Fireball {
		return mage.NewFireball(sim, target)
	} else {
		return mage.NewScorch(sim, target)
	}
}

func (mage *Mage) doFrostRotation(sim *core.Simulation) *core.SimpleSpell {
	target := sim.GetPrimaryTarget()
	spell := mage.NewFrostbolt(sim, target)
	return spell
}
