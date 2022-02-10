package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var SealOfBloodAuraID = core.NewAuraID()
var SealOfBloodCastActionID = core.ActionID{SpellID: 31892}
var SealOfBloodProcActionID = core.ActionID{SpellID: 31893}

// Handles the cast, gcd, deducts the mana cost
func (paladin *Paladin) NewSealOfBlood(sim *core.Simulation) *core.SimpleCast {
	sob := &core.SimpleCast{
		Cast: core.Cast{
			ActionID:     SealOfBloodCastActionID,
			Character:    paladin.GetCharacter(),
			ManaCost:     210,
			GCD:          core.GCDDefault,
		},
		OnCastComplete: paladin.ApplySealOfBlood,
	}

	sob.Init(sim)

	return sob
}

/* This also needs to deal damage to the paladin or something of similar effect so we can
 * model the mana regeneration from spiritual attunement.
 * 
 * Defines the seal on-hit affects and applies the aura
 * Created seperate function because it might be useful to call this on sim init rather than cast mid-fight
 */
func (paladin *Paladin) ApplySealOfBlood(sim *core.Simulation, cast *core.Cast) {
	sobtempl := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			ActionID:    SealOfBloodProcActionID,
			Character:   &paladin.Character,
			SpellSchool: stats.HolySpellPower,
			CritMultiplier: paladin.DefaultMeleeCritMultiplier(),
			IsPhantom: true,
		},
		Effect: core.AbilityHitEffect{
			AbilityEffect: core.AbilityEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
				IgnoreArmor:            true,
			},
			WeaponInput: core.WeaponDamageInput{
				DamageMultiplier: 0.35, // should deal 35% weapon deamage
			},
		},
	}

	sobTemplate := core.NewMeleeAbilityTemplate(sobtempl)
	sobAtk := core.ActiveMeleeAbility{}
	sobAura := core.Aura{
		ID: SealOfBloodAuraID,
		ActionID: SealOfBloodProcActionID,

		Expires: sim.CurrentTime + time.Second * 30,

		OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			if !hitEffect.Landed() || !hitEffect.IsWeaponHit() || ability.IsPhantom {
				return
			}

			sobTemplate.Apply(&sobAtk)
			sobAtk.Effect.Target = hitEffect.Target
			sobAtk.Attack(sim)
		},
	}

	paladin.AddAura(sim, sobAura)
}

