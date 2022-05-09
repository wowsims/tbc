package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	// Keep these in order by item ID.
	core.AddWeaponEffect(16250, ApplyWeaponSuperiorStriking)
	core.AddWeaponEffect(22552, ApplyWeaponMajorStriking) // TODO: add to frontend, probably replacing Superior Striking
	core.AddItemEffect(16252, ApplyCrusader)
	core.AddItemEffect(18283, ApplyBiznicksScope)
	core.AddItemEffect(22535, ApplyRingStriking)
	core.AddItemEffect(22559, ApplyMongoose)
	core.AddWeaponEffect(23765, ApplyKhoriumScope)
	core.AddItemEffect(23766, ApplyStabilizedEterniumScope)
	core.AddItemEffect(33150, ApplyBackSubtlety)
	core.AddItemEffect(33153, ApplyGlovesThreat)
	core.AddItemEffect(33307, ApplyExecutioner)
}

// TODO: Crusader, Mongoose, and Executioner could also be modelled as AddWeaponEffect instead
func ApplyWeaponSuperiorStriking(agent core.Agent, slot proto.ItemSlot) {
	w := &agent.GetCharacter().AutoAttacks.MH
	if slot == proto.ItemSlot_ItemSlotOffHand {
		w = &agent.GetCharacter().AutoAttacks.OH
	}
	w.BaseDamageMin += 5
	w.BaseDamageMax += 5
}

func ApplyWeaponMajorStriking(agent core.Agent, slot proto.ItemSlot) {
	w := &agent.GetCharacter().AutoAttacks.MH
	if slot == proto.ItemSlot_ItemSlotOffHand {
		w = &agent.GetCharacter().AutoAttacks.OH
	}
	w.BaseDamageMin += 7
	w.BaseDamageMax += 7
}

// ApplyCrusaderEffect will be applied twice if there is two weapons with this enchant.
//   However it will automatically overwrite one of them so it should be ok.
//   A single application of the aura will handle both mh and oh procs.
func ApplyCrusader(agent core.Agent) {
	character := agent.GetCharacter()
	ppmm := character.AutoAttacks.NewPPMManager(1.0)
	mh := character.Equip[proto.ItemSlot_ItemSlotMainHand].Enchant.ID == 16252
	oh := character.Equip[proto.ItemSlot_ItemSlotOffHand].Enchant.ID == 16252
	if !mh && !oh {
		return
	}
	if !mh {
		ppmm.SetProcChance(true, 0)
	}
	if !oh {
		ppmm.SetProcChance(false, 0)
	}

	// -4 str per level over 60
	const strBonus = 100.0 - 4.0*float64(core.CharacterLevel-60)
	mhAura := character.NewTemporaryStatsAura("Crusader Enchant MH", core.ActionID{SpellID: 20007, Tag: 1}, stats.Stats{stats.Strength: strBonus}, time.Second*15)
	ohAura := character.NewTemporaryStatsAura("Crusader Enchant OH", core.ActionID{SpellID: 20007, Tag: 2}, stats.Stats{stats.Strength: strBonus}, time.Second*15)

	character.GetOrRegisterAura(core.Aura{
		Label:    "Crusader Enchant",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}

			isMH := spellEffect.IsMH()
			if ppmm.Proc(sim, isMH, false, "Crusader") {
				if isMH {
					mhAura.Activate(sim)
				} else {
					ohAura.Activate(sim)
				}
			}
		},
	})
}

func ApplyBiznicksScope(agent core.Agent) {
	character := agent.GetCharacter()
	character.PseudoStats.BonusRangedHitRating += 30
}

func ApplyRingStriking(agent core.Agent) {
	agent.GetCharacter().PseudoStats.BonusDamage += 2
}

func newLightningSpeedAura(character *core.Character, auraLabel string, actionID core.ActionID) *core.Aura {
	return character.NewTemporaryStatsAuraWrapped(auraLabel, actionID, stats.Stats{stats.Agility: 120}, time.Second*15, func(aura *core.Aura) {
		oldOnGain := aura.OnGain
		oldOnExpire := aura.OnExpire
		aura.OnGain = func(aura *core.Aura, sim *core.Simulation) {
			oldOnGain(aura, sim)
			character.MultiplyMeleeSpeed(sim, 1.02)
		}
		aura.OnExpire = func(aura *core.Aura, sim *core.Simulation) {
			oldOnExpire(aura, sim)
			character.MultiplyMeleeSpeed(sim, 1/1.02)
		}
	})
}

// ApplyMongooseEffect will be applied twice if there is two weapons with this enchant.
//   However it will automatically overwrite one of them so it should be ok.
//   A single application of the aura will handle both mh and oh procs.
func ApplyMongoose(agent core.Agent) {
	character := agent.GetCharacter()
	ppmm := character.AutoAttacks.NewPPMManager(1.0)
	mh := character.Equip[proto.ItemSlot_ItemSlotMainHand].Enchant.ID == 22559
	oh := character.Equip[proto.ItemSlot_ItemSlotOffHand].Enchant.ID == 22559
	if !mh && !oh {
		return
	}
	if !mh {
		ppmm.SetProcChance(true, 0)
	}
	if !oh {
		ppmm.SetProcChance(false, 0)
	}

	mhAura := newLightningSpeedAura(character, "Lightning Speed MH", core.ActionID{SpellID: 28093, Tag: 1})
	ohAura := newLightningSpeedAura(character, "Lightning Speed OH", core.ActionID{SpellID: 28093, Tag: 2})

	character.GetOrRegisterAura(core.Aura{
		Label:    "Mongoose Enchant",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}

			isMH := spellEffect.IsMH()
			if ppmm.Proc(sim, isMH, false, "mongoose") {
				if isMH {
					mhAura.Activate(sim)
				} else {
					ohAura.Activate(sim)
				}
			}
		},
	})
}

func ApplyKhoriumScope(agent core.Agent, _ proto.ItemSlot) {
	w := &agent.GetCharacter().AutoAttacks.Ranged
	w.BaseDamageMin += 12
	w.BaseDamageMax += 12
}

func ApplyStabilizedEterniumScope(agent core.Agent) {
	character := agent.GetCharacter()
	character.PseudoStats.BonusRangedCritRating += 28
}

func ApplyBackSubtlety(agent core.Agent) {
	character := agent.GetCharacter()
	character.PseudoStats.ThreatMultiplier *= 0.98
}
func ApplyGlovesThreat(agent core.Agent) {
	character := agent.GetCharacter()
	character.PseudoStats.ThreatMultiplier *= 1.02
}

func ApplyExecutioner(agent core.Agent) {
	character := agent.GetCharacter()
	ppmm := character.AutoAttacks.NewPPMManager(1.0)
	mh := character.Equip[proto.ItemSlot_ItemSlotMainHand].Enchant.ID == 33307
	oh := character.Equip[proto.ItemSlot_ItemSlotOffHand].Enchant.ID == 33307
	if !mh && !oh {
		return
	}
	if !mh {
		ppmm.SetProcChance(true, 0)
	}
	if !oh {
		ppmm.SetProcChance(false, 0)
	}

	procAura := character.NewTemporaryStatsAura("Executioner Proc", core.ActionID{SpellID: 42976}, stats.Stats{stats.ArmorPenetration: 840}, time.Second*15)

	character.GetOrRegisterAura(core.Aura{
		Label:    "Executioner",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}

			if ppmm.Proc(sim, spellEffect.IsMH(), false, "Executioner") {
				procAura.Activate(sim)
			}
		},
	})
}
