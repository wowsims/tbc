package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

var StormstrikeCD = core.NewCooldownID()
var StormstrikeDebuffID = core.NewDebuffID()
var StormstrikeActionID = core.ActionID{SpellID: 17364, CooldownID: StormstrikeCD}
var SkyshatterAPBonusAuraID = core.NewAuraID()

func (shaman *Shaman) newStormstrikeTemplate(sim *core.Simulation) core.MeleeAbilityTemplate {

	ssDebuffAura := core.Aura{
		ID:       StormstrikeDebuffID,
		ActionID: StormstrikeActionID,
		Stacks:   2,
	}
	ssDebuffAura.OnBeforeSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellCast.SpellSchool != stats.NatureSpellPower {
			return
		}
		spellEffect.DamageMultiplier *= 1.2
		stacks := spellEffect.Target.NumStacks(StormstrikeDebuffID) - 1
		if stacks == 0 {
			spellEffect.Target.RemoveAura(sim, StormstrikeDebuffID)
		} else {
			ssDebuffAura.Stacks = stacks
			spellEffect.Target.ReplaceAura(sim, ssDebuffAura)
		}
	}

	hasSkyshatter4p := ItemSetSkyshatterHarness.CharacterHasSetBonus(&shaman.Character, 4)
	const skyshatterDur = time.Second * 12
	ss := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			// ID for the action.
			ActionID: StormstrikeActionID,
			Cooldown: time.Second * 10,
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 237,
			},
			CritMultiplier:  2.0,
			ResetSwingTimer: true,
			Character:       &shaman.Character,
		},
		MainHit: core.AbilityHitEffect{
			AbilityEffect: core.AbilityEffect{
				DamageMultiplier:       1.0,
				StaticDamageMultiplier: 1.0,
			},
			WeaponInput: core.WeaponDamageInput{
				IsMH:             true,
				DamageMultiplier: 1.0,
			},
		},
		AdditionalHits: []core.AbilityHitEffect{
			core.AbilityHitEffect{
				AbilityEffect: core.AbilityEffect{
					DamageMultiplier:       1.0,
					StaticDamageMultiplier: 1.0,
				},
				WeaponInput: core.WeaponDamageInput{
					IsMH:             false,
					DamageMultiplier: 1.0,
				},
			},
		},
		OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			if !hitEffect.Landed() {
				return
			}

			ssDebuffAura.Stacks = 2
			hitEffect.Target.ReplaceAura(sim, ssDebuffAura)
			if hasSkyshatter4p {
				shaman.Character.AddAuraWithTemporaryStats(sim, SkyshatterAPBonusAuraID, core.ActionID{SpellID: 38432}, stats.SpellPower, 70, skyshatterDur)
			}
		},
	}

	if shaman.Equip[items.ItemSlotRanged].ID == StormfuryTotem {
		ss.MeleeAbility.Cost.Value -= 22
	}

	if ItemSetCycloneHarness.CharacterHasSetBonus(&shaman.Character, 4) {
		ss.MainHit.WeaponInput.FlatDamageBonus += 30
		ss.AdditionalHits[0].WeaponInput.FlatDamageBonus += 30
	}

	// Add weapon % bonus to stormstrike weapons
	ss.MainHit.WeaponInput.DamageMultiplier *= 1 + 0.02*float64(shaman.Talents.WeaponMastery)
	ss.AdditionalHits[0].WeaponInput.DamageMultiplier *= 1 + 0.02*float64(shaman.Talents.WeaponMastery)

	return core.NewMeleeAbilityTemplate(ss)
}

func (shaman *Shaman) NewStormstrike(sim *core.Simulation, target *core.Target) *core.ActiveMeleeAbility {
	ss := &shaman.stormstrikeSpell
	shaman.stormstrikeTemplate.Apply(ss)
	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ss.MainHit.Target = target
	ss.AdditionalHits[0].Target = target
	return ss
}
