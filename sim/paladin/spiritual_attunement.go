package paladin

import (
	"github.com/wowsims/tbc/sim/core"
)

func (paladin *Paladin) registerSpiritualAttunement() {
	paladin.RegisterAura(core.Aura{
		Label:    "Spiritual Attunement",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Damage > 0 {
				paladin.AddMana(sim, spellEffect.Damage*0.1, core.ActionID{SpellID: 33776}, false)
			}
		},
	})
}
