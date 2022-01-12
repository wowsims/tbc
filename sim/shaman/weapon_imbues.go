package shaman

import (
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

var TotemOfTheAstralWinds int32 = 27815

var WFImbueAuraID = core.NewAuraID()

func ApplyWindfuryImbue(shaman *Shaman, mh bool, oh bool) {
	var proc = 0.2
	if mh && oh {
		proc = 0.36
	}
	apBonus := 475.0

	if shaman.Equip[proto.ItemSlot_ItemSlotRanged].ID == TotemOfTheAstralWinds {
		apBonus += 80
	}

	wftempl := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			Name: "Windfury Attack",
			ActionID: core.ActionID{
				SpellID: 25505,
			},
			CritMultiplier:  2.0,
			Character:       &shaman.Character,
			IgnoreCooldowns: true,
		},
		WeaponDamageInput: core.WeaponDamageInput{
			MainHand: 1.0,
			Offhand:  1.0,
		},
		AbilityEffect: core.AbilityEffect{
			DamageMultiplier:       1.0,
			StaticDamageMultiplier: 1.0,
			BonusAttackPower:       apBonus,
			IgnoreDualWieldPenalty: true,
		},
	}
	if shaman.Talents.ElementalWeapons > 0 {
		wftempl.MainHand *= 1 + math.Round(float64(shaman.Talents.ElementalWeapons)*13.33)/100
		wftempl.Offhand *= 1 + math.Round(float64(shaman.Talents.ElementalWeapons)*13.33)/100
	}

	wfTemplate := core.NewMeleeAbilityTemplate(wftempl)

	wfAtk := core.ActiveMeleeAbility{}
	shaman.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		var icd core.InternalCD
		const icdDur = time.Second * 3

		return core.Aura{
			ID: WFImbueAuraID,
			OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
				if (!mh && !isOH) || (isOH && !oh) {
					return // cant proc if not enchanted
				}
				if icd.IsOnCD(sim) {
					return
				}
				if sim.RandomFloat("wf imbue") > proc {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				for i := 0; i < 2; i++ {
					wfTemplate.Apply(&wfAtk)
					wfAtk.BonusAttackPower += apBonus
					// Set so only the proc'd hand attacks
					if isOH {
						wfAtk.MainHand = 0
					} else {
						wfAtk.Offhand = 0
					}
					wfAtk.Target = target
					wfAtk.Attack(sim)
				}
			},
		}
	})
}
