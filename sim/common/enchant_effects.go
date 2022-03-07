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
		ppmm.SetProcChance(true, 0)
	}
	if !oh {
		ppmm.SetProcChance(false, 0)
	}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		// -4 str per level over 60
		const strBonus = 100.0 - 4.0*float64(core.CharacterLevel-60)
		applyStatAuraMH := character.NewTemporaryStatsAuraApplier(CrusaderStrengthMHAuraID, core.ActionID{ItemID: 16252, Tag: 1}, stats.Stats{stats.Strength: strBonus}, time.Second*15)
		applyStatAuraOH := character.NewTemporaryStatsAuraApplier(CrusaderStrengthOHAuraID, core.ActionID{ItemID: 16252, Tag: 2}, stats.Stats{stats.Strength: strBonus}, time.Second*15)

		return core.Aura{
			ID: CrusaderAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}

				isMH := spellEffect.IsMH()
				if ppmm.Proc(sim, isMH, false, "Crusader") {
					if isMH {
						applyStatAuraMH(sim)
					} else {
						applyStatAuraOH(sim)
					}
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
			OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
				if spellCast.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged) {
					spellEffect.BonusHitRating += 30
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
		ppmm.SetProcChance(true, 0)
	}
	if !oh {
		ppmm.SetProcChance(false, 0)
	}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		agiBonus := 120.0
		hasteBonus := 2 * core.HasteRatingPerHastePercent
		applyStatAuraMH := character.NewTemporaryStatsAuraApplier(LightningSpeedMHAuraID, core.ActionID{SpellID: 28093, Tag: 1}, stats.Stats{stats.Agility: agiBonus, stats.MeleeHaste: hasteBonus}, time.Second*15)
		applyStatAuraOH := character.NewTemporaryStatsAuraApplier(LightningSpeedOHAuraID, core.ActionID{SpellID: 28093, Tag: 2}, stats.Stats{stats.Agility: agiBonus, stats.MeleeHaste: hasteBonus}, time.Second*15)

		return core.Aura{
			ID: MongooseAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}

				isMH := spellEffect.IsMH()
				if ppmm.Proc(sim, isMH, false, "mongoose") {
					if isMH {
						applyStatAuraMH(sim)
					} else {
						applyStatAuraOH(sim)
					}
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
			OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
				if spellCast.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged) {
					spellEffect.BonusCritRating += 28
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
		ppmm.SetProcChance(true, 0)
	}
	if !oh {
		ppmm.SetProcChance(false, 0)
	}

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const arPenBonus = 840.0
		const dur = time.Second * 15
		applyStatAura := character.NewTemporaryStatsAuraApplier(ExecutionerProcAuraID, core.ActionID{SpellID: 42976}, stats.Stats{stats.ArmorPenetration: arPenBonus}, dur)

		return core.Aura{
			ID: ExecutionerAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}

				if ppmm.Proc(sim, spellEffect.IsMH(), false, "Executioner") {
					applyStatAura(sim)
				}
			},
		}
	})
}
