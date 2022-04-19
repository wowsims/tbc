package warrior

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var SweepingStrikesActionID = core.ActionID{SpellID: 12328}

func (warrior *Warrior) applySweepingStrikes() {
	if !warrior.Talents.SweepingStrikes {
		return
	}

	cost := 30.0
	sweepingStrikes := warrior.RegisterSpell(core.SpellConfig{
		ActionID:    SweepingStrikesActionID,
		SpellSchool: core.SpellSchoolPhysical,

		ResourceType: stats.Rage,
		BaseCost:     cost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: cost,
			},
			CD: core.Cooldown{
				Timer:    warrior.NewTimer(),
				Duration: time.Second * 30,
			},
		},
	})

	warrior.RegisterAura(core.Aura{
		Label:    "Sweeping Strikes",
		ActionID: SweepingStrikesActionID,
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.SetStacks(sim, 10)
		},
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Damage == 0 || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}

			// TODO: If the triggering spell is Execute and 2nd target health > 20%, do a normalized MH hit instead.
			damage := spellEffect.Damage * reverseArmorMit

			targetIndex := spellEffect.Target.Index + 1
			if targetIndex >= sim.NumTargets() {
				targetIndex = 0
			}
			target := sim.GetTarget(targetIndex)

			sweepingStrikes.doDirectDamage(sim, target, damage)

			aura.RemoveStack(sim)
		},
	})

	warrior.AddMajorCooldown(core.MajorCooldown{
		Spell: sweepingStrikes,
		Type:  core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return warrior.CurrentRage() >= cost
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
	})
}
