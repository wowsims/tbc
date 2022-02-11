package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const TwistWindow = time.Millisecond * 400

var SealOfBloodAuraID = core.NewAuraID()
var SealOfBloodCastActionID = core.ActionID{SpellID: 31892}
var SealOfBloodProcActionID = core.ActionID{SpellID: 31893}

// Handles the cast, gcd, deducts the mana cost
func (paladin *Paladin) setupSealOfBlood() {
	// The proc behaviour
	sobProc := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			ActionID:       SealOfBloodProcActionID,
			Character:      &paladin.Character,
			SpellSchool:    stats.HolySpellPower,
			CritMultiplier: paladin.DefaultMeleeCritMultiplier(),
			IsPhantom:      true,
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

	sobTemplate := core.NewMeleeAbilityTemplate(sobProc)
	sobAtk := core.ActiveMeleeAbility{}

	// Define the aura
	sobAura := core.Aura{
		ID:       SealOfBloodAuraID,
		ActionID: SealOfBloodProcActionID,

		OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			if !hitEffect.Landed() || !hitEffect.IsWeaponHit() || ability.IsPhantom {
				return
			}

			sobTemplate.Apply(&sobAtk)
			sobAtk.Effect.Target = hitEffect.Target
			sobAtk.Attack(sim)
		},
	}

	sob := core.SimpleCast{
		Cast: core.Cast{
			ActionID:  SealOfBloodCastActionID,
			Character: paladin.GetCharacter(),
			ManaCost:  210,
			GCD:       core.GCDDefault,
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			sobAura.Expires = sim.CurrentTime + time.Second * 30
			paladin.UpdateSeal(sim, sobAura)
		},
	}

	paladin.sealOfBlood = sob
}

func (paladin *Paladin) NewSealOfBlood(sim *core.Simulation) *core.SimpleCast {
	sob := &paladin.sealOfBlood
	sob.Init(sim)
	return sob
}

var SealOfCommandAuraID = core.NewAuraID()
var SealOfCommandCastActionID = core.ActionID{SpellID: 20375}
var SealOfCommandProcActionID = core.ActionID{SpellID: 20424}

func (paladin *Paladin) setupSealOfCommand() {
	socProc := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			ActionID:       SealOfCommandProcActionID,
			Character:      &paladin.Character,
			SpellSchool:    stats.HolySpellPower,
			CritMultiplier: paladin.DefaultMeleeCritMultiplier(),
		},
		Effect: core.AbilityHitEffect{
			AbilityEffect: core.AbilityEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
				IgnoreArmor:            true,
			},
			WeaponInput: core.WeaponDamageInput{
				DamageMultiplier: 0.70, // should deal 70% weapon deamage
			},
		},
	}

	socTemplate := core.NewMeleeAbilityTemplate(socProc)
	socAtk := core.ActiveMeleeAbility{}

	ppmm := paladin.AutoAttacks.NewPPMManager(7.0)

	// I might not be implementing the ICD correctly here, should debug later
	var icd core.InternalCD
	const icdDur = time.Second * 1

	socAura := core.Aura{
		ID:       SealOfCommandAuraID,
		ActionID: SealOfCommandProcActionID,

		OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			if !hitEffect.Landed() || !hitEffect.IsWeaponHit() || ability.IsPhantom {
				return
			}

			if icd.IsOnCD(sim) {
				return
			}

			if !ppmm.Proc(sim, true, false, "Frostbrand Weapon") {
				return
			}

			icd = core.InternalCD(sim.CurrentTime + icdDur)

			socTemplate.Apply(&socAtk)
			socAtk.Effect.Target = hitEffect.Target
			socAtk.Attack(sim)
		},
	}

	soc := core.SimpleCast{
		Cast: core.Cast{
			ActionID:  SealOfCommandCastActionID,
			Character: paladin.GetCharacter(),
			ManaCost:  65,
			GCD:       core.GCDDefault,
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			socAura.Expires = sim.CurrentTime + time.Second * 30
			paladin.UpdateSeal(sim, socAura)
		},
	}

	paladin.sealOfCommand = soc
}


func (paladin *Paladin) NewSealOfCommand(sim *core.Simulation) *core.SimpleCast {
	soc := &paladin.sealOfCommand
	soc.Init(sim)
	return soc
}

func (paladin *Paladin) UpdateSeal(sim *core.Simulation, newSeal core.Aura) {
	oldSeal := paladin.currentSeal

	// For Seal of Command, reduce duration to 0.4 seconds
	if oldSeal.ID == SealOfCommandAuraID {
		// Technically the current expiration could be shorter than 0.4 seconds
		// TO-DO: Lookup behavior when seal of command is twisted at shorter than 0.4 seconds duration
		oldSeal.Expires = sim.CurrentTime + TwistWindow 
		paladin.ReplaceAura(sim, oldSeal)

		// This is a hack to get the sim to process and log the SoC aura expiring at the right time
		if (sim.Options.Iterations == 1) {
			sim.AddPendingAction(&core.PendingAction{
				NextActionAt:sim.CurrentTime + TwistWindow, 
				OnAction: func(_ *core.Simulation) {},
			})
		}
	} else {
		paladin.RemoveAura(sim, oldSeal.ID)
	}

	paladin.currentSeal = newSeal
	paladin.AddAura(sim, newSeal)
}
