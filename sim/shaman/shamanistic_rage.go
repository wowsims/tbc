package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ShamanisticRageCD = core.NewCooldownID()
var ShamanisticRageActionID = core.ActionID{SpellID: 30823, CooldownID: ShamanisticRageCD}

func (shaman *Shaman) registerShamanisticRageCD() {
	if !shaman.Talents.ShamanisticRage {
		return
	}

	const cd = time.Minute * 2

	ppmm := shaman.AutoAttacks.NewPPMManager(15)
	srAura := shaman.RegisterAura(core.Aura{
		Label:    "Shamanistic Rage",
		ActionID: ShamanisticRageActionID,
		Duration: time.Second * 15,
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			// proc mask: 20
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}
			if !ppmm.Proc(sim, spellEffect.IsMH(), false, "shamanistic rage") {
				return
			}
			mana := shaman.GetStat(stats.AttackPower) * 0.3
			shaman.AddMana(sim, mana, ShamanisticRageActionID, true)
		},
	})

	spell := shaman.RegisterSpell(core.SpellConfig{
		ActionID: ShamanisticRageActionID,
		Cast: core.CastConfig{
			Cooldown:         cd,
			DisableCallbacks: true,
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
