package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var CrusaderStrikeActionID = core.ActionID{SpellID: 35395}

// Do some research on the spell fields to make sure I'm doing this right
// Need to add in judgement debuff refreshing feature at some point
func (paladin *Paladin) registerCrusaderStrikeSpell(sim *core.Simulation) {
	baseCost := 236.0

	paladin.CrusaderStrike = paladin.RegisterSpell(core.SpellConfig{
		ActionID:    CrusaderStrikeActionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    paladin.NewTimer(),
				Duration: time.Second * 6,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:  core.ProcMaskMeleeMHSpecial,
			IsPhantom: true,

			DamageMultiplier: 1, // Need to review to make sure I set these properly
			ThreatMultiplier: 1,

			// maybe this isn't the one that should be set to 1.1
			BaseDamage:     core.BaseDamageConfigMeleeWeapon(core.MainHand, false, 0, 1.1, true),
			OutcomeApplier: paladin.OutcomeFuncMeleeSpecialHitAndCrit(paladin.DefaultMeleeCritMultiplier()),

			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}
			},
		}),
	})
}
