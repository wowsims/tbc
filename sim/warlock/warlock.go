package warlock

import "github.com/wowsims/tbc/sim/core"

type Warlock struct {
	*core.Character

	malediction int // bonus level of coe
}

func (warlock *Warlock) GetCharacter() *core.Character {
	return warlock.Character
}

func (warlock *Warlock) ChooseAction(sim *core.Simulation) core.AgentAction {
	return core.NewWaitAction(sim, warlock, core.NeverExpires) // makes the bot wait forever and do nothing.
}

func (warlock *Warlock) OnActionAccepted(*core.Simulation, core.AgentAction) {
}

func (warlock *Warlock) BuffUp(sim *core.Simulation) {
	sim.AddAura(sim, CurseOfElementsAura(warlock.malediction))
}

func (warlock *Warlock) Reset(sim *core.Simulation) {}

func CurseOfElementsAura(malediction int) core.Aura {
	multiplier := 1.10 + 0.1*float64(malediction)
	return core.Aura{
		ID:      core.MagicIDCurseOfElements,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, cast *core.DirectCastAction, result *core.DirectCastDamageResult) {
			result.Damage *= multiplier
		},
	}
}
