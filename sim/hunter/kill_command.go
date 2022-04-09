package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var KillCommandCooldownID = core.NewCooldownID()
var KillCommandActionID = core.ActionID{SpellID: 34026, CooldownID: KillCommandCooldownID}

func (hunter *Hunter) applyKillCommand() {
	if hunter.pet == nil {
		return
	}

	hunter.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return hunter.GetOrRegisterAura(&core.Aura{
			Label: "Kill Command Trigger",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Outcome.Matches(core.OutcomeCrit) {
					hunter.killCommandEnabledUntil = sim.CurrentTime + time.Second*5
					hunter.TryKillCommand(sim, sim.GetPrimaryTarget())
				}
			},
		})
	})
}

func (hunter *Hunter) registerKillCommandSpell(sim *core.Simulation) {
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
				OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
					hunter.killCommandEnabledUntil = 0
					hunter.pet.KillCommand.Cast(sim, sim.GetPrimaryTarget())
				},
			},
		},
		Effect: core.SpellEffect{
			ThreatMultiplier: 1,
		},
	}

	hunter.KillCommand = hunter.RegisterSpell(core.SpellConfig{
		Template:   spell,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

func (hp *HunterPet) registerKillCommandSpell(sim *core.Simulation) {
	var beastLordProcAura *core.Aura
	if ItemSetBeastLord.CharacterHasSetBonus(&hp.hunterOwner.Character, 4) {
		beastLordProcAura = hp.hunterOwner.NewTemporaryStatsAura("Beast Lord Proc", core.ActionID{SpellID: 37483}, stats.Stats{stats.ArmorPenetration: 600}, time.Second*15)
	}

	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: 34027},
				Character:   &hp.Character,
				SpellSchool: core.SpellSchoolPhysical,
				SpellExtras: core.SpellExtrasMeleeMetrics,
			},
		},
	}

	hp.KillCommand = hp.RegisterSpell(core.SpellConfig{
		Template: ama,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:         core.ProcMaskMeleeMHSpecial,
			BonusCritRating:  float64(hp.hunterOwner.Talents.FocusedFire) * 10 * core.MeleeCritRatingPerCritChance,
			DamageMultiplier: hp.config.DamageMultiplier,
			ThreatMultiplier: 1,

			BaseDamage:     core.BaseDamageConfigMeleeWeapon(core.MainHand, false, 127, 1, true),
			OutcomeApplier: core.OutcomeFuncMeleeSpecialHitAndCrit(2),

			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if beastLordProcAura != nil {
					beastLordProcAura.Activate(sim)
				}
			},
		}),
	})
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
		hunter.KillCommand.Cast(sim, target)
	}
}
