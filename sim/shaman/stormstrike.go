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

func (shaman *Shaman) newStormstrikeTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	ssDebuffAura := core.Aura{
		ID:       StormstrikeDebuffID,
		ActionID: StormstrikeActionID,
		Stacks:   2,
	}
	ssDebuffAura.OnBeforeSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
		if spellCast.SpellSchool != core.SpellSchoolNature {
			return
		}
		spellEffect.DamageMultiplier *= 1.2
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

	hasSkyshatter4p := ItemSetSkyshatterHarness.CharacterHasSetBonus(&shaman.Character, 4)
	skyshatterAuraApplier := shaman.NewTemporaryStatsAuraApplier(SkyshatterAPBonusAuraID, core.ActionID{SpellID: 38432}, stats.Stats{stats.AttackPower: 70}, time.Second*12)

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
		Effects: []core.SpellHitEffect{
			{
				SpellEffect: core.SpellEffect{
					ProcMask:               core.ProcMaskMeleeMHSpecial,
					DamageMultiplier:       1,
					StaticDamageMultiplier: 1,
					ThreatMultiplier:       1,
					OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						if !spellEffect.Landed() {
							return
						}

						ssDebuffAura.Stacks = 2
						ssDebuffAura.Expires = sim.CurrentTime + 12*time.Second
						spellEffect.Target.ReplaceAura(sim, ssDebuffAura)
						if hasSkyshatter4p {
							skyshatterAuraApplier(sim)
						}
					},
				},
				WeaponInput: core.WeaponDamageInput{
					DamageMultiplier: 1,
				},
			},
			{
				SpellEffect: core.SpellEffect{
					ProcMask:               core.ProcMaskMeleeOHSpecial,
					DamageMultiplier:       1,
					StaticDamageMultiplier: 1,
					ThreatMultiplier:       1,
					ReuseMainHitRoll:       true,
				},
				WeaponInput: core.WeaponDamageInput{
					Offhand:          true,
					DamageMultiplier: 1,
				},
			},
		},
	}

	if shaman.Equip[items.ItemSlotRanged].ID == StormfuryTotem {
		ss.Cost.Value -= 22
	}

	if ItemSetCycloneHarness.CharacterHasSetBonus(&shaman.Character, 4) {
		ss.Effects[0].WeaponInput.FlatDamageBonus += 30
		ss.Effects[1].WeaponInput.FlatDamageBonus += 30
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
