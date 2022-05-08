package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (shaman *Shaman) registerShamanisticRageCD() {
	if !shaman.Talents.ShamanisticRage {
		return
	}

	actionID := core.ActionID{SpellID: 30823}
	ppmm := shaman.AutoAttacks.NewPPMManager(15)
	srAura := shaman.RegisterAura(core.Aura{
		Label:    "Shamanistic Rage",
		ActionID: actionID,
		Duration: time.Second * 15,
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			// proc mask: 20
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}
			if !ppmm.Proc(sim, spellEffect.IsMH(), false, "shamanistic rage") {
				return
			}
			mana := shaman.GetStat(stats.AttackPower) * 0.3
			shaman.AddMana(sim, mana, actionID, true)
		},
	})

	spell := shaman.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellExtras: core.SpellExtrasNoOnCastComplete,
		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    shaman.NewTimer(),
				Duration: time.Minute * 2,
			},
		},
		ApplyEffects: func(sim *core.Simulation, _ *core.Target, _ *core.Spell) {
			srAura.Activate(sim)
		},
	})

	shaman.AddMajorCooldown(core.MajorCooldown{
		Spell: spell,
		Type:  core.CooldownTypeMana,
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			const manaReserve = 1000 // If mana goes under 1000 we will need more soon. Pop shamanistic rage.
			if character.CurrentMana() > manaReserve {
				return false
			}

			return true
		},
	})
}
