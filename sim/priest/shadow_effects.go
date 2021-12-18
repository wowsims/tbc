package priest

import (
	"github.com/wowsims/tbc/sim/core"
)

func (priest *Priest) ApplyMisery(sim *core.Simulation, target *core.Target) {
	if priest.Talents.Misery == 5 {
		target.ReplaceAura(sim, core.MiseryAura(sim))
	}
}

func (priest *Priest) ApplyShadowWeaving(sim *core.Simulation, target *core.Target) {
	if sim.RandomFloat("Shadow Weaving") > 0.2*float64(priest.Talents.ShadowWeaving) {
		return
	}

	curStacks := target.NumStacks(core.ShadowWeavingDebuffID)
	newStacks := core.MinInt32(curStacks+1, 5)

	if sim.Log != nil && curStacks != newStacks {
		priest.Log(sim, "Applied Shadow Weaving stack, %d --> %d", curStacks, newStacks)
	}

	target.ReplaceAura(sim, core.ShadowWeavingAura(sim, newStacks))
}
