package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func init() {
	core.AddItemEffect(core.ItemIDTheLightningCapacitor, ApplyTheLightningCapacitor)
}

var TheLightningCapacitorAuraID = core.NewAuraID()

func ApplyTheLightningCapacitor(agent core.Agent) {
	character := agent.GetCharacter()

	tlcSpell := character.RegisterSpell(core.SpellConfig{
		Template: core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID: core.ActionID{
						ItemID: core.ItemIDTheLightningCapacitor,
					},
					Character:   character,
					SpellSchool: core.SpellSchoolNature,
				},
			},
			Effect: core.SpellEffect{
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				CritRollCategory:    core.CritRollCategoryMagical,
				CritMultiplier:      character.DefaultSpellCritMultiplier(),
				IsPhantom:           true, // TODO: replace with ProcMask
				DamageMultiplier:    1,
				ThreatMultiplier:    1,
				BaseDamage:          core.BaseDamageConfigRoll(694, 807),
			},
		},
		ModifyCast: core.ModifyCastAssignTarget,
	})

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		charges := 0

		const icdDur = time.Millisecond * 2500
		icd := core.NewICD()

		return core.Aura{
			ID: TheLightningCapacitorAuraID,
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if icd.IsOnCD(sim) {
					return
				}

				if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}

				if !spellEffect.Outcome.Matches(core.OutcomeCrit) {
					return
				}

				charges++
				if charges >= 3 {
					icd = core.InternalCD(sim.CurrentTime + icdDur)
					tlcSpell.Cast(sim, spellEffect.Target)
					charges = 0
				}
			},
		}
	})
}
