package common

import (
	"log"
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
	switch slot {
	case proto.ItemSlot_ItemSlotMainHand:
		if w := &agent.GetCharacter().AutoAttacks.MH; w.SwingSpeed > 0 {
			w.BaseDamageMin += 5
			w.BaseDamageMax += 5
		}
	case proto.ItemSlot_ItemSlotOffHand:
		if w := &agent.GetCharacter().AutoAttacks.OH; w.SwingSpeed > 0 {
			w.BaseDamageMin += 5
			w.BaseDamageMax += 5
		}
	default:
		log.Fatalf("Cannot add Superior Striking to %s", slot)
	}
}

func ApplyWeaponMajorStriking(agent core.Agent, slot proto.ItemSlot) {
	switch slot {
	case proto.ItemSlot_ItemSlotMainHand:
		if w := &agent.GetCharacter().AutoAttacks.MH; w.SwingSpeed > 0 {
			w.BaseDamageMin += 7
			w.BaseDamageMax += 7
		}
	case proto.ItemSlot_ItemSlotOffHand:
		if w := &agent.GetCharacter().AutoAttacks.OH; w.SwingSpeed > 0 {
			w.BaseDamageMin += 7
			w.BaseDamageMax += 7
		}
	default:
		log.Fatalf("Cannot add Major Striking to %s", slot)
	}
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
		applyStatAuraMH := character.NewTemporaryStatsAuraApplier(CrusaderStrengthMHAuraID, core.ActionID{SpellID: 20007, Tag: 1}, stats.Stats{stats.Strength: strBonus}, time.Second*15)
		applyStatAuraOH := character.NewTemporaryStatsAuraApplier(CrusaderStrengthOHAuraID, core.ActionID{SpellID: 20007, Tag: 2}, stats.Stats{stats.Strength: strBonus}, time.Second*15)

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
	agent.GetCharacter().PseudoStats.BonusDamage += 2
}

var MongooseAuraID = core.NewAuraID()

var LightningSpeedMHAuraID = core.NewAuraID()
var LightningSpeedOHAuraID = core.NewAuraID()

func newLightningSpeedApplier(character *core.Character, auraID core.AuraID, actionID core.ActionID) func(sim *core.Simulation) {
	factory := newLightningSpeedAuraFactory(character, auraID, actionID)

	return func(sim *core.Simulation) {
		character.ReplaceAura(sim, factory(sim))
	}
}

func newLightningSpeedAuraFactory(character *core.Character, auraID core.AuraID, actionID core.ActionID) func(sim *core.Simulation) core.Aura {
	buffs := character.ApplyStatDependencies(stats.Stats{stats.Agility: 120})
	unbuffs := buffs.Multiply(-1)

	aura := core.Aura{
		ID:       auraID,
		ActionID: actionID,
		OnExpire: func(sim *core.Simulation) {
			character.AddStatsDynamic(sim, unbuffs)
			character.MultiplyMeleeSpeed(sim, 1/1.02)
			if sim.Log != nil {
				character.Log(sim, "Lost %s from fading %s", buffs.FlatString(), actionID)
			}
		},
	}

	return func(sim *core.Simulation) core.Aura {
		if !character.HasAura(auraID) {
			character.AddStatsDynamic(sim, buffs)
			character.MultiplyMeleeSpeed(sim, 1.02)
			if sim.Log != nil {
				character.Log(sim, "Gained %s from %s", buffs.FlatString(), actionID)
			}
		}
		aura.Expires = sim.CurrentTime + 15*time.Second
		return aura
	}
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

	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		applyStatAuraMH := newLightningSpeedApplier(character, LightningSpeedMHAuraID, core.ActionID{SpellID: 28093, Tag: 1})
		applyStatAuraOH := newLightningSpeedApplier(character, LightningSpeedOHAuraID, core.ActionID{SpellID: 28093, Tag: 2})

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

func ApplyKhoriumScope(agent core.Agent, slot proto.ItemSlot) {
	switch slot {
	case proto.ItemSlot_ItemSlotRanged:
		if w := &agent.GetCharacter().AutoAttacks.Ranged; w.SwingSpeed > 0 {
			w.BaseDamageMin += 12
			w.BaseDamageMax += 12
		}
	default:
		log.Fatalf("Cannot add Superior Striking to %s", slot)
	}
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
