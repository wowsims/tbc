package priest

import (
	"github.com/wowsims/tbc/sim/core"
)

func (priest *Priest) ApplyMisery(sim *core.Simulation, target *core.Target) {
	if priest.Talents.Misery >= target.NumStacks(core.MiseryDebuffID) {
		target.ReplaceAura(sim, core.MiseryAura(sim, priest.Talents.Misery))
	}
}

func (priest *Priest) ApplyShadowWeaving(sim *core.Simulation, target *core.Target) {
	if priest.Talents.ShadowWeaving == 0 {
		return
	}

	if priest.Talents.ShadowWeaving < 5 && sim.RandomFloat("Shadow Weaving") > 0.2*float64(priest.Talents.ShadowWeaving) {
		return
	}

	curStacks := target.NumStacks(core.ShadowWeavingDebuffID)
	newStacks := core.MinInt32(curStacks+1, 5)

	if sim.Log != nil && curStacks != newStacks {
		priest.Log(sim, "Applied Shadow Weaving stack, %d --> %d", curStacks, newStacks)
	}

	target.ReplaceAura(sim, core.ShadowWeavingAura(sim, newStacks))
}
