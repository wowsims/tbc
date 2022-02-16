package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	// Keep these in order by item ID.
	core.AddItemEffect(16250, ApplyWeaponSuperiorStriking)
	core.AddItemEffect(16252, ApplyCrusader)
	core.AddItemEffect(18283, ApplyBiznicksScope)
	core.AddItemEffect(22535, ApplyRingStriking)
	core.AddItemEffect(22559, ApplyMongoose)
	core.AddItemEffect(23765, ApplyKhoriumScope)
	core.AddItemEffect(23766, ApplyStabilizedEterniumScope)
	core.AddItemEffect(33150, ApplyBackSubtlety)
	core.AddItemEffect(33153, ApplyGlovesThreat)
	core.AddItemEffect(33307, ApplyExecutioner)
}

func ApplyWeaponSuperiorStriking(agent core.Agent) {
	agent.GetCharacter().PseudoStats.BonusMeleeDamage += 5
	// Melee only, no ranged bonus.
}

var CrusaderAuraID = core.NewAuraID()

var CrusaderStrengthMHAuraID = core.NewAuraID()
var CrusaderStrengthOHAuraID = core.NewAuraID()

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
		ppmm.SetProcChance(false, 0)
	}
	if !oh {
		ppmm.SetProcChance(true, 0)
	}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		// -4 str per level over 60
		buffs := character.ApplyStatDependencies(stats.Stats{stats.Strength: 60})
		buffs[stats.Mana] = 0 // mana is wierd
		unbuffs := buffs.Multiply(-1)

		applyCrusaderStrength := func(sim *core.Simulation, character *core.Character, isMH bool) {
			character.AddStats(buffs)
			var tag int32
			var auraID core.AuraID
			if isMH {
				tag = 1
				auraID = CrusaderStrengthMHAuraID
			} else {
				tag = 2
				auraID = CrusaderStrengthOHAuraID
			}
			character.AddAura(sim, core.Aura{
				ID:       auraID,
				ActionID: core.ActionID{ItemID: 16252, Tag: tag},
				Expires:  sim.CurrentTime + (time.Second * 15),
				OnExpire: func(sim *core.Simulation) {
					character.AddStats(unbuffs)
				},
			})
		}

		return core.Aura{
			ID: CrusaderAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}

				isMH := hitEffect.IsMH()
				if ppmm.Proc(sim, isMH, false, "Crusader") {
					applyCrusaderStrength(sim, character, isMH)
				}
			},
		}
	})
}

var BiznicksScopeAuraID = core.NewAuraID()

func ApplyBiznicksScope(agent core.Agent) {
	character := agent.GetCharacter()
	if character.Class != proto.Class_ClassHunter {
		return
	}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: BiznicksScopeAuraID,
			OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if hitEffect.IsRanged() {
					hitEffect.BonusHitRating += 30
				}
			},
		}
	})
}

func ApplyRingStriking(agent core.Agent) {
	agent.GetCharacter().PseudoStats.BonusMeleeDamage += 2
	agent.GetCharacter().PseudoStats.BonusRangedDamage += 2
}

var MongooseAuraID = core.NewAuraID()

var LightningSpeedMHAuraID = core.NewAuraID()
var LightningSpeedOHAuraID = core.NewAuraID()

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
		ppmm.SetProcChance(false, 0)
	}
	if !oh {
		ppmm.SetProcChance(true, 0)
	}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		buffs := character.ApplyStatDependencies(stats.Stats{stats.Agility: 120})
		buffs[stats.Mana] = 0 // mana is weird
		unbuffs := buffs.Multiply(-1)
		haste := 2 * core.HasteRatingPerHastePercent

		applyLightningSpeed := func(sim *core.Simulation, character *core.Character, isMH bool) {
			// https://tbc.wowhead.com/spell=28093/lightning-speed
			character.AddStats(buffs)
			character.AddMeleeHaste(sim, haste)
			var tag int32
			var auraID core.AuraID
			if isMH {
				tag = 1
				auraID = LightningSpeedMHAuraID
			} else {
				tag = 2
				auraID = LightningSpeedOHAuraID
			}
			character.AddAura(sim, core.Aura{
				ID:       auraID,
				ActionID: core.ActionID{SpellID: 28093, Tag: tag},
				Expires:  sim.CurrentTime + (time.Second * 15),
				OnExpire: func(sim *core.Simulation) {
					character.AddStats(unbuffs)
					character.AddMeleeHaste(sim, -haste)
				},
			})
		}

		return core.Aura{
			ID: MongooseAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}

				isMH := hitEffect.IsMH()
				if ppmm.Proc(sim, isMH, false, "mongoose") {
					applyLightningSpeed(sim, character, isMH)
				}
			},
		}
	})
}

func ApplyKhoriumScope(agent core.Agent) {
	agent.GetCharacter().PseudoStats.BonusRangedDamage += 12
}

var StabilizedEterniumScopeAuraID = core.NewAuraID()

func ApplyStabilizedEterniumScope(agent core.Agent) {
	character := agent.GetCharacter()
	if character.Class != proto.Class_ClassHunter {
		return
	}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: StabilizedEterniumScopeAuraID,
			OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if hitEffect.IsRanged() {
					hitEffect.BonusCritRating += 28
				}
			},
		}
	})
}

func ApplyBackSubtlety(agent core.Agent) {
	character := agent.GetCharacter()
	character.PseudoStats.ThreatMultiplier *= 0.98
}
func ApplyGlovesThreat(agent core.Agent) {
	character := agent.GetCharacter()
	character.PseudoStats.ThreatMultiplier *= 1.02
}

var ExecutionerAuraID = core.NewAuraID()
var ExecutionerProcAuraID = core.NewAuraID()

func ApplyExecutioner(agent core.Agent) {
	character := agent.GetCharacter()
	ppmm := character.AutoAttacks.NewPPMManager(1.0)
	mh := character.Equip[proto.ItemSlot_ItemSlotMainHand].Enchant.ID == 33307
	oh := character.Equip[proto.ItemSlot_ItemSlotOffHand].Enchant.ID == 33307
	if !mh && !oh {
		return
	}
	if !mh {
		ppmm.SetProcChance(false, 0)
	}
	if !oh {
		ppmm.SetProcChance(true, 0)
	}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const arPenBonus = 840.0
		const dur = time.Second * 15
		return core.Aura{
			ID: ExecutionerAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}

				if ppmm.Proc(sim, hitEffect.IsMH(), false, "Executioner") {
					character.AddAuraWithTemporaryStats(sim, ExecutionerProcAuraID, core.ActionID{SpellID: 42976}, stats.ArmorPenetration, arPenBonus, dur)
				}
			},
		}
	})
}
