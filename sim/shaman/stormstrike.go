package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

var StormstrikeCD = core.NewCooldownID()
var StormstrikeActionID = core.ActionID{SpellID: 17364, CooldownID: StormstrikeCD}

func (shaman *Shaman) stormstrikeDebuffAura(target *core.Target) *core.Aura {
	return target.GetOrRegisterAura(core.Aura{
		Label:     "Stormstrike",
		ActionID:  StormstrikeActionID,
		Duration:  time.Second * 12,
		MaxStacks: 2,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			target.PseudoStats.NatureDamageTakenMultiplier *= 1.2
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			target.PseudoStats.NatureDamageTakenMultiplier /= 1.2
		},
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spell.SpellSchool != core.SpellSchoolNature {
				return
			}
			if !spellEffect.Landed() || spellEffect.Damage == 0 {
				return
			}

			aura.RemoveStack(sim)
		},
	})
}

func (shaman *Shaman) newStormstrikeHitSpell(isMH bool) *core.Spell {
	effect := core.SpellEffect{
		DamageMultiplier: 1,
		ThreatMultiplier: core.TernaryFloat64(shaman.Talents.SpiritWeapons, 0.7, 1),
		OutcomeApplier:   core.OutcomeFuncMeleeSpecialCritOnly(shaman.DefaultMeleeCritMultiplier()),
	}

	flatDamageBonus := core.TernaryFloat64(ItemSetCycloneHarness.CharacterHasSetBonus(&shaman.Character, 4), 30, 0)
	if isMH {
		effect.ProcMask = core.ProcMaskMeleeMHSpecial
		effect.BaseDamage = core.BaseDamageConfigMeleeWeapon(core.MainHand, false, flatDamageBonus, 1, true)
	} else {
		effect.ProcMask = core.ProcMaskMeleeOHSpecial
		effect.BaseDamage = core.BaseDamageConfigMeleeWeapon(core.OffHand, false, flatDamageBonus, 1, true)
	}

	return shaman.RegisterSpell(core.SpellConfig{
		ActionID:    StormstrikeActionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

		ApplyEffects: core.ApplyEffectFuncDirectDamage(effect),
	})
}

func (shaman *Shaman) registerStormstrikeSpell(sim *core.Simulation) {
	mhHit := shaman.newStormstrikeHitSpell(true)
	ohHit := shaman.newStormstrikeHitSpell(false)

	baseCost := 237.0
	if shaman.Equip[items.ItemSlotRanged].ID == StormfuryTotem {
		baseCost -= 22
	}

	ssDebuffAura := shaman.stormstrikeDebuffAura(sim.GetPrimaryTarget())

	var skyshatterAura *core.Aura
	if ItemSetSkyshatterHarness.CharacterHasSetBonus(&shaman.Character, 4) {
		skyshatterAura = shaman.NewTemporaryStatsAura("Skyshatter 4pc AP Bonus", core.ActionID{SpellID: 38432}, stats.Stats{stats.AttackPower: 70}, time.Second*12)
	}

	shaman.Stormstrike = shaman.RegisterSpell(core.SpellConfig{
		ActionID:    StormstrikeActionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  core.GCDDefault,
			},
			IgnoreHaste: true,
			Cooldown:    time.Second * 10,
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ThreatMultiplier: 1,
			OutcomeApplier:   core.OutcomeFuncMeleeSpecialHit(),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}

				ssDebuffAura.Activate(sim)
				ssDebuffAura.SetStacks(sim, 2)

				if skyshatterAura != nil {
					skyshatterAura.Activate(sim)
				}

				mhHit.Cast(sim, spellEffect.Target)
				ohHit.Cast(sim, spellEffect.Target)
				shaman.Stormstrike.Casts -= 2
				shaman.Stormstrike.Hits--
			},
		}),
	})
}
