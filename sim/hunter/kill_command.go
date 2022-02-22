package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var KillCommandCooldownID = core.NewCooldownID()
var KillCommandActionID = core.ActionID{SpellID: 34026, CooldownID: KillCommandCooldownID}

var KillCommandAuraID = core.NewAuraID()

func (hunter *Hunter) applyKillCommand() {
	if hunter.pet == nil {
		return
	}

	hunter.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: KillCommandAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if hitEffect.Outcome.Matches(core.OutcomeCrit) {
					hunter.killCommandEnabledUntil = sim.CurrentTime + time.Second*5
					hunter.TryKillCommand(sim, sim.GetPrimaryTarget())
				}
			},
		}
	})
}

// ActiveMeleeAbility doesn't support cast times, so we wrap it in a SimpleCast.
func (hunter *Hunter) newKillCommandTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:  KillCommandActionID,
				Character: hunter.GetCharacter(),
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 75,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 75,
				},
				Cooldown: time.Second * 5,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ThreatMultiplier: 1,
				IgnoreHitCheck:   true,
			},
		},
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (hp *HunterPet) newKillCommandTemplate(sim *core.Simulation) core.MeleeAbilityTemplate {
	hasBeastLord4Pc := ItemSetBeastLord.CharacterHasSetBonus(&hp.hunterOwner.Character, 4)
	beastLordApplier := hp.hunterOwner.NewTempStatAuraApplier(sim, BeastLord4PcAuraID, core.ActionID{SpellID: 37483}, stats.ArmorPenetration, 600, time.Second*15)

	ama := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			ActionID:            core.ActionID{SpellID: 34027},
			Character:           &hp.Character,
			OutcomeRollCategory: core.OutcomeRollCategorySpecial,
			CritRollCategory:    core.CritRollCategoryPhysical,
			SpellSchool:         core.SpellSchoolPhysical,
			CritMultiplier:      2,
		},
		Effect: core.AbilityHitEffect{
			AbilityEffect: core.AbilityEffect{
				ProcMask:               core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			WeaponInput: core.WeaponDamageInput{
				DamageMultiplier: 1,
				FlatDamageBonus:  127,
			},
		},
		OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			if hasBeastLord4Pc {
				beastLordApplier(sim)
			}
		},
	}

	ama.Effect.DamageMultiplier *= hp.config.DamageMultiplier
	ama.Effect.BonusCritRating += float64(hp.hunterOwner.Talents.FocusedFire) * 10 * core.MeleeCritRatingPerCritChance

	return core.NewMeleeAbilityTemplate(ama)
}

func (hunter *Hunter) NewKillCommand(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	killCommand := &hunter.killCommand
	hunter.killCommandTemplate.Apply(killCommand)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	killCommand.Effect.Target = target
	killCommand.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		hunter.killCommandEnabledUntil = 0

		kc := &hunter.pet.killCommand
		hunter.pet.killCommandTemplate.Apply(kc)
		kc.Effect.Target = target
		kc.Attack(sim)
	}

	killCommand.Init(sim)
	return killCommand
}

func (hunter *Hunter) TryKillCommand(sim *core.Simulation, target *core.Target) {
	if !hunter.pet.IsEnabled() {
		return
	}

	if hunter.killCommandEnabledUntil < sim.CurrentTime || hunter.killCommandBlocked {
		return
	}

	if hunter.CurrentMana() < 75 {
		return
	}

	if !hunter.IsOnCD(KillCommandCooldownID, sim.CurrentTime) {
		kc := hunter.NewKillCommand(sim, target)
		kc.Cast(sim)
	}
}
