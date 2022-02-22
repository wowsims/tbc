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
		if spellCast.SpellSchool != core.SpellSchoolNature {
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
			ActionID:            StormstrikeActionID,
			Character:           &shaman.Character,
			OutcomeRollCategory: core.OutcomeRollCategorySpecial,
			CritRollCategory:    core.CritRollCategoryPhysical,
			SpellSchool:         core.SpellSchoolPhysical,
			GCD:                 core.GCDDefault,
			Cooldown:            time.Second * 10,
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 237,
			},
			CritMultiplier: shaman.DefaultMeleeCritMultiplier(),
		},
		Effects: []core.AbilityHitEffect{
			{
				AbilityEffect: core.AbilityEffect{
					ProcMask:               core.ProcMaskMeleeMHSpecial,
					DamageMultiplier:       1,
					StaticDamageMultiplier: 1,
					ThreatMultiplier:       1,
				},
				WeaponInput: core.WeaponDamageInput{
					DamageMultiplier: 1,
				},
			},
			{
				AbilityEffect: core.AbilityEffect{
					ProcMask:               core.ProcMaskMeleeOHSpecial,
					DamageMultiplier:       1,
					StaticDamageMultiplier: 1,
					ThreatMultiplier:       1,
					ReuseMainHitRoll:       true,
				},
				WeaponInput: core.WeaponDamageInput{
					DamageMultiplier: 1,
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
				shaman.Character.AddAuraWithTemporaryStats(sim, SkyshatterAPBonusAuraID, core.ActionID{SpellID: 38432}, stats.AttackPower, 70, skyshatterDur)
			}
		},
	}

	if shaman.Equip[items.ItemSlotRanged].ID == StormfuryTotem {
		ss.MeleeAbility.Cost.Value -= 22
	}

	if ItemSetCycloneHarness.CharacterHasSetBonus(&shaman.Character, 4) {
		ss.Effects[0].WeaponInput.FlatDamageBonus += 30
		ss.Effects[1].WeaponInput.FlatDamageBonus += 30
	}

	return core.NewMeleeAbilityTemplate(ss)
}

func (shaman *Shaman) NewStormstrike(sim *core.Simulation, target *core.Target) *core.ActiveMeleeAbility {
	ss := &shaman.stormstrikeSpell
	shaman.stormstrikeTemplate.Apply(ss)
	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ss.Effects[0].Target = target
	ss.Effects[1].Target = target
	return ss
}
