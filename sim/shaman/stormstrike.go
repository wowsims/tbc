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

func (shaman *Shaman) newStormstrikeTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	ssDebuffAura := shaman.stormstrikeDebuffAura(sim.GetPrimaryTarget())

	hasSkyshatter4p := ItemSetSkyshatterHarness.CharacterHasSetBonus(&shaman.Character, 4)
	skyshatterAuraApplier := shaman.NewTemporaryStatsAuraApplier(SkyshatterAPBonusAuraID, core.ActionID{SpellID: 38432}, stats.Stats{stats.AttackPower: 70}, time.Second*12)

	flatDamageBonus := 0.0
	if ItemSetCycloneHarness.CharacterHasSetBonus(&shaman.Character, 4) {
		flatDamageBonus += 30
	}

	ss := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            StormstrikeActionID,
				Character:           &shaman.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 core.GCDDefault,
				IgnoreHaste:         true,
				Cooldown:            time.Second * 10,
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 237,
				},
				CritMultiplier: shaman.DefaultMeleeCritMultiplier(),
			},
		},
		Effects: []core.SpellEffect{
			{
				ProcMask:         core.ProcMaskMeleeMHSpecial,
				DamageMultiplier: 1,
				ThreatMultiplier: 1,
				BaseDamage:       core.BaseDamageConfigMeleeWeapon(core.MainHand, false, flatDamageBonus, 1, true),
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					if !spellEffect.Landed() {
						return
					}

					ssDebuffAura.Stacks = 2
					spellEffect.Target.ReplaceAura(sim, ssDebuffAura)
					if hasSkyshatter4p {
						skyshatterAuraApplier(sim)
					}
				},
			},
			{
				ProcMask:         core.ProcMaskMeleeOHSpecial,
				DamageMultiplier: 1,
				ThreatMultiplier: 1,
				ReuseMainHitRoll: true,
				BaseDamage:       core.BaseDamageConfigMeleeWeapon(core.OffHand, false, flatDamageBonus, 1, true),
			},
		},
	}

	if shaman.Equip[items.ItemSlotRanged].ID == StormfuryTotem {
		ss.Cost.Value -= 22
	}

	if shaman.Talents.SpiritWeapons {
		ss.Effects[0].ThreatMultiplier *= 0.7
		ss.Effects[1].ThreatMultiplier *= 0.7
	}

	return core.NewSimpleSpellTemplate(ss)
}

func (shaman *Shaman) NewStormstrike(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	ss := &shaman.stormstrikeSpell
	shaman.stormstrikeTemplate.Apply(ss)
	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ss.Effects[0].Target = target
	ss.Effects[1].Target = target
	return ss
}
