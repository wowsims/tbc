package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var CrusaderStrikeCD = core.NewCooldownID()
var CrusaderStrikeActionID = core.ActionID{SpellID: 35395, CooldownID: CrusaderStrikeCD}

// Do some research on the spell fields to make sure I'm doing this right
// Need to add in judgement debuff refreshing feature at some point
func (paladin *Paladin) registerCrusaderStrikeSpell(sim *core.Simulation) {
	cs := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    CrusaderStrikeActionID,
				Character:   &paladin.Character,
				SpellSchool: core.SpellSchoolPhysical,
				GCD:         core.GCDDefault,
				Cooldown:    time.Second * 6,
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 236,
				},
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategorySpecial,
			CritRollCategory:    core.CritRollCategoryPhysical,
			CritMultiplier:      paladin.DefaultMeleeCritMultiplier(),
			IsPhantom:           true,
			ProcMask:            core.ProcMaskMeleeMHSpecial,
			DamageMultiplier:    1, // Need to review to make sure I set these properly
			ThreatMultiplier:    1,
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}
			},
			// maybe this isn't the one that should be set to 1.1
			BaseDamage: core.BaseDamageConfigMeleeWeapon(core.MainHand, false, 0, 1.1, true),
		},
	}

	paladin.CrusaderStrike = paladin.RegisterSpell(core.SpellConfig{
		Template:   cs,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
