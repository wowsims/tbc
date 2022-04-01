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

func (shaman *Shaman) stormstrikeDebuffAura(target *core.Target) core.Aura {
	ssDebuffAura := core.Aura{
		ID:       StormstrikeDebuffID,
		ActionID: StormstrikeActionID,
		Duration: time.Second * 12,
		Stacks:   2,
		OnGain: func(sim *core.Simulation) {
			target.PseudoStats.NatureDamageTakenMultiplier *= 1.2
		},
		OnExpire: func(sim *core.Simulation) {
			target.PseudoStats.NatureDamageTakenMultiplier /= 1.2
		},
	}
	ssDebuffAura.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellCast.SpellSchool != core.SpellSchoolNature {
			return
		}
		if !spellEffect.Landed() || spellEffect.Damage == 0 {
			return
		}

		stacks := spellEffect.Target.NumStacks(StormstrikeDebuffID) - 1
		if stacks == 0 {
			spellEffect.Target.RemoveAura(sim, StormstrikeDebuffID)
		} else {
			ssDebuffAura.Stacks = stacks
			spellEffect.Target.ReplaceAura(sim, ssDebuffAura)
		}
	}
	return ssDebuffAura
}

func (shaman *Shaman) newStormstrikeHitSpell(isMH bool) *core.SimpleSpellTemplate {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    StormstrikeActionID,
				Character:   &shaman.Character,
				SpellSchool: core.SpellSchoolPhysical,
				SpellExtras: core.SpellExtrasAlwaysHits,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategorySpecial,
			CritRollCategory:    core.CritRollCategoryPhysical,
			CritMultiplier:      shaman.DefaultMeleeCritMultiplier(),
			DamageMultiplier:    1,
			ThreatMultiplier:    1,
		},
	}

	if shaman.Talents.SpiritWeapons {
		template.Effect.ThreatMultiplier *= 0.7
	}

	flatDamageBonus := 0.0
	if ItemSetCycloneHarness.CharacterHasSetBonus(&shaman.Character, 4) {
		flatDamageBonus += 30
	}

	if isMH {
		template.Effect.ProcMask = core.ProcMaskMeleeMHSpecial
		template.Effect.BaseDamage = core.BaseDamageConfigMeleeWeapon(core.MainHand, false, flatDamageBonus, 1, true)
	} else {
		template.Effect.ProcMask = core.ProcMaskMeleeOHSpecial
		template.Effect.BaseDamage = core.BaseDamageConfigMeleeWeapon(core.OffHand, false, flatDamageBonus, 1, true)
	}

	return shaman.RegisterSpell(core.SpellConfig{
		Template:   template,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

func (shaman *Shaman) registerStormstrikeSpell(sim *core.Simulation) {
	ssDebuffAura := shaman.stormstrikeDebuffAura(sim.GetPrimaryTarget())

	hasSkyshatter4p := ItemSetSkyshatterHarness.CharacterHasSetBonus(&shaman.Character, 4)
	skyshatterAuraApplier := shaman.NewTemporaryStatsAuraApplier(SkyshatterAPBonusAuraID, core.ActionID{SpellID: 38432}, stats.Stats{stats.AttackPower: 70}, time.Second*12)

	mhHit := shaman.newStormstrikeHitSpell(true)
	ohHit := shaman.newStormstrikeHitSpell(false)

	ss := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    StormstrikeActionID,
				Character:   &shaman.Character,
				SpellSchool: core.SpellSchoolPhysical,
				GCD:         core.GCDDefault,
				IgnoreHaste: true,
				Cooldown:    time.Second * 10,
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 237,
				},
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategorySpecial,
			CritRollCategory:    core.CritRollCategoryNone,
			ThreatMultiplier:    1,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}

				ssDebuffAura.Stacks = 2
				spellEffect.Target.ReplaceAura(sim, ssDebuffAura)
				if hasSkyshatter4p {
					skyshatterAuraApplier(sim)
				}

				mhHit.Cast(sim, spellEffect.Target)
				ohHit.Cast(sim, spellEffect.Target)
				shaman.Stormstrike.Casts -= 2
				shaman.Stormstrike.Hits--
			},
		},
	}

	if shaman.Equip[items.ItemSlotRanged].ID == StormfuryTotem {
		ss.Cost.Value -= 22
	}

	shaman.Stormstrike = shaman.RegisterSpell(core.SpellConfig{
		Template:   ss,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
