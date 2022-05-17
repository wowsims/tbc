package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (paladin *Paladin) registerHolyShieldSpell() {
	actionID := core.ActionID{SpellID: 27179}

	numCharges := 4 + 2*paladin.Talents.ImprovedHolyShield

	procSpell := paladin.RegisterSpell(core.SpellConfig{
		ActionID:    actionID.WithTag(1),
		SpellSchool: core.SpellSchoolHoly,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:         core.ProcMaskEmpty,
			IsPhantom:        true,
			DamageMultiplier: 1 + 0.1*float64(paladin.Talents.ImprovedHolyShield),
			ThreatMultiplier: 1.35,

			BaseDamage:     core.BaseDamageConfigMagicNoRoll(155, 0.05),
			OutcomeApplier: paladin.OutcomeFuncAlwaysHit(),
		}),
	})

	holyShieldAura := paladin.RegisterAura(core.Aura{
		Label:     "Holy Shield",
		ActionID:  actionID,
		Duration:  time.Second * 10,
		MaxStacks: numCharges,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			paladin.AddStatDynamic(sim, stats.Block, 30*core.BlockRatingPerBlockChance)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			paladin.AddStatDynamic(sim, stats.Block, -30*core.BlockRatingPerBlockChance)
		},
		OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Outcome.Matches(core.OutcomeBlock) {
				procSpell.Cast(sim, spell.Unit)
				aura.RemoveStack(sim)
			}
		},
	})

	baseCost := 280.0

	paladin.HolyShield = paladin.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolHoly,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    paladin.NewTimer(),
				Duration: time.Second * 10,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			holyShieldAura.Activate(sim)
			holyShieldAura.SetStacks(sim, numCharges)
		},
	})
}
