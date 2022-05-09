package warlock

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (warlock *Warlock) registerLifeTapSpell() {
	actionID := core.ActionID{SpellID: 27222}
	mana := 582.0 * (1.0 + 0.1*float64(warlock.Talents.ImprovedLifeTap))
	petRestore := 0.3333 * float64(warlock.Talents.ManaFeed)

	warlock.LifeTap = warlock.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolShadow,
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ThreatMultiplier: 1,
			FlatThreatBonus:  1,
			OutcomeApplier:   warlock.OutcomeFuncAlwaysHit(),
			OnSpellHitDealt: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				// Life tap adds 0.8*sp to mana restore
				// TODO: does AddMana generate threat correctly?
				restore := mana + (warlock.GetStat(stats.SpellPower)+warlock.GetStat(stats.ShadowSpellPower))*0.8
				warlock.AddMana(sim, restore, actionID, true)

				if warlock.Talents.ManaFeed > 0 {
					for _, pet := range warlock.Pets {
						pet.GetPet().AddMana(sim, restore*petRestore, actionID, true)
					}
				}
			},
		}),
	})
}
