package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var StormstrikeCD = core.NewCooldownID()
var StormstrikeDebuffID = core.NewDebuffID()
var StormstrikeActionID = core.ActionID{SpellID: 17364, CooldownID: StormstrikeCD}
var SkyshatterAPBonusAuraID = core.NewAuraID()

func (shaman *Shaman) newStormstrikeTemplate(sim *core.Simulation) core.MeleeAbilittyTemplate {

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
			Name:     "Stormstrike",
			Cooldown: time.Second * 10,
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 237,
			},
			CritMultiplier:  2.0,
			ResetSwingTimer: true,
			Character:       &shaman.Character,
		},
		WeaponDamageInput: core.WeaponDamageInput{
			MainHand: 1.0,
			Offhand:  1.0,
		},
		AbilityEffect: core.AbilityEffect{
			DamageMultiplier:       1.0,
			StaticDamageMultiplier: 1.0,
			IgnoreDualWieldPenalty: true,
		},
		OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
			ssDebuffAura.Stacks = 2
			target.ReplaceAura(sim, ssDebuffAura)
			if hasSkyshatter4p {
				shaman.Character.AddAuraWithTemporaryStats(sim, SkyshatterAPBonusAuraID, core.ActionID{SpellID: 38432}, stats.SpellPower, 70, skyshatterDur)
			}
		},
	}

	if ItemSetCycloneHarness.CharacterHasSetBonus(&shaman.Character, 4) {
		ss.WeaponDamageInput.MainHandFlat += 30
		ss.WeaponDamageInput.OffhandFlat += 30
	}

	// Add weapon % bonus to stormstrike weapons
	ss.MainHand *= 1 + 0.02*float64(shaman.Talents.WeaponMastery)
	ss.Offhand *= 1 + 0.02*float64(shaman.Talents.WeaponMastery)

	return core.NewMeleeAbilittyTemplate(ss)
}

func (shaman *Shaman) NewStormstrike(sim *core.Simulation, target *core.Target) *core.ActiveMeleeAbility {
	ss := &shaman.stormstrikeSpell
	shaman.stormstrikeTemplate.Apply(ss)
	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ss.Target = target
	return ss
}
