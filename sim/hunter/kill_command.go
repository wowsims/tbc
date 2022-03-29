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
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellEffect.Outcome.Matches(core.OutcomeCrit) {
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
		Effect: core.SpellEffect{
			ThreatMultiplier: 1,
		},
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (hp *HunterPet) newKillCommandTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	hasBeastLord4Pc := ItemSetBeastLord.CharacterHasSetBonus(&hp.hunterOwner.Character, 4)
	beastLordStatApplier := hp.hunterOwner.NewTemporaryStatsAuraApplier(BeastLord4PcAuraID, core.ActionID{SpellID: 37483}, stats.Stats{stats.ArmorPenetration: 600}, time.Second*15)

	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: 34027},
				Character:           &hp.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				CritMultiplier:      2,
			},
		},
		Effect: core.SpellEffect{
			ProcMask:         core.ProcMaskMeleeMHSpecial,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMeleeWeapon(core.MainHand, false, 127, 1, true),
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if hasBeastLord4Pc {
					beastLordStatApplier(sim)
				}
			},
		},
	}

	ama.Effect.DamageMultiplier *= hp.config.DamageMultiplier
	ama.Effect.BonusCritRating += float64(hp.hunterOwner.Talents.FocusedFire) * 10 * core.MeleeCritRatingPerCritChance

	return core.NewSimpleSpellTemplate(ama)
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
		kc.Init(sim)
		kc.Cast(sim)
	}

	killCommand.Init(sim)
	return killCommand
}

func (hunter *Hunter) TryKillCommand(sim *core.Simulation, target *core.Target) {
	if hunter.pet == nil || !hunter.pet.IsEnabled() {
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
