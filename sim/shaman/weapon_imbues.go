package shaman

import (
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var TotemOfTheAstralWinds int32 = 27815

var WFImbueAuraID = core.NewAuraID()

func (shaman *Shaman) ApplyWindfuryImbue(mh bool, oh bool) {
	if !mh && !oh {
		return
	}

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
			ActionID: core.ActionID{
				SpellID: 25505,
			},
			CritMultiplier:  2.0,
			Character:       &shaman.Character,
			IgnoreCooldowns: true,
		},
		MainHit: core.AbilityHitEffect{
			AbilityEffect: core.AbilityEffect{
				DamageMultiplier:       1.0,
				StaticDamageMultiplier: 1.0,
				BonusAttackPower:       apBonus,
			},
			WeaponInput: core.WeaponDamageInput{
				DamageMultiplier: 1.0,
			},
		},
	}
	if shaman.Talents.ElementalWeapons > 0 {
		wftempl.MainHit.WeaponInput.DamageMultiplier *= 1 + math.Round(float64(shaman.Talents.ElementalWeapons)*13.33)/100
	}

	wfTemplate := core.NewMeleeAbilityTemplate(wftempl)

	wfAtk := core.ActiveMeleeAbility{}
	shaman.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		var icd core.InternalCD
		const icdDur = time.Second * 3

		return core.Aura{
			ID: WFImbueAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					return
				}

				isMHHit := hitEffect.IsMH()
				if (isMHHit && !mh) || (!isMHHit && !oh) {
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
					wfAtk.MainHit.BonusAttackPower += apBonus

					// Set so only the proc'd hand attacks
					wfAtk.MainHit.WeaponInput.IsOH = !isMHHit

					wfAtk.MainHit.Target = hitEffect.Target
					wfAtk.Attack(sim)
				}
			},
		}
	})
}

var FTImbueAuraID = core.NewAuraID()

func (shaman *Shaman) ApplyFlametongueImbue(mh bool, oh bool) {
	if !mh && !oh {
		return
	}

	ftTmpl := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:        core.ActionID{ItemID: 25489},
				Character:       &shaman.Character,
				IgnoreCooldowns: true,
				IgnoreManaCost:  true,
				IsPhantom:       true,
				SpellSchool:     stats.FireSpellPower,
				CritMultiplier:  1.5,
			},
		},
		SpellHitEffect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
			},
			DirectInput: core.DirectDamageInput{
				SpellCoefficient: 0.1,
			},
		},
	}
	ftTmpl.SpellHitEffect.SpellEffect.StaticDamageMultiplier *= 1 + 0.05*float64(shaman.Talents.ElementalWeapons)

	mhTmpl := ftTmpl
	ohTmpl := ftTmpl

	if weapon := shaman.Equip[proto.ItemSlot_ItemSlotMainHand]; weapon.ID != 0 {
		baseDamage := weapon.SwingSpeed * 35.0
		mhTmpl.SpellHitEffect.DirectInput.MinBaseDamage = baseDamage
		mhTmpl.SpellHitEffect.DirectInput.MaxBaseDamage = baseDamage
	}
	if weapon := shaman.Equip[proto.ItemSlot_ItemSlotOffHand]; weapon.ID != 0 {
		baseDamage := weapon.SwingSpeed * 35.0
		ohTmpl.SpellHitEffect.DirectInput.MinBaseDamage = baseDamage
		ohTmpl.SpellHitEffect.DirectInput.MaxBaseDamage = baseDamage
	}

	mhTemplate := core.NewSimpleSpellTemplate(mhTmpl)
	ohTemplate := core.NewSimpleSpellTemplate(ohTmpl)
	ftSpell := core.SimpleSpell{}

	shaman.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: FTImbueAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					return
				}

				isMHHit := hitEffect.IsMH()
				if (isMHHit && !mh) || (!isMHHit && !oh) {
					return // cant proc if not enchanted
				}

				if isMHHit {
					mhTemplate.Apply(&ftSpell)
				} else {
					ohTemplate.Apply(&ftSpell)
				}
				ftSpell.Target = hitEffect.Target
				ftSpell.Init(sim)
				ftSpell.Cast(sim)
			},
		}
	})
}

var FBImbueAuraID = core.NewAuraID()

func (shaman *Shaman) ApplyFrostbrandImbue(mh bool, oh bool) {
	if !mh && !oh {
		return
	}

	fbTmpl := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:        core.ActionID{ItemID: 25500},
				Character:       &shaman.Character,
				IgnoreCooldowns: true,
				IgnoreManaCost:  true,
				IsPhantom:       true,
				SpellSchool:     stats.FrostSpellPower,
				CritMultiplier:  1.5,
			},
		},
		SpellHitEffect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage:    246,
				MaxBaseDamage:    246,
				SpellCoefficient: 0.1,
			},
		},
	}
	fbTmpl.SpellHitEffect.SpellEffect.StaticDamageMultiplier *= 1 + 0.05*float64(shaman.Talents.ElementalWeapons)

	fbTemplate := core.NewSimpleSpellTemplate(fbTmpl)
	fbSpell := core.SimpleSpell{}

	shaman.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		ppmm := shaman.AutoAttacks.NewPPMManager(9.0)
		return core.Aura{
			ID: FBImbueAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					return
				}

				isMHHit := hitEffect.IsMH()
				if (isMHHit && !mh) || (!isMHHit && !oh) {
					return // cant proc if not enchanted
				}

				if !ppmm.Proc(sim, isMHHit, "Frostbrand Weapon") {
					return
				}

				fbTemplate.Apply(&fbSpell)
				fbSpell.Target = hitEffect.Target
				fbSpell.Init(sim)
				fbSpell.Cast(sim)
			},
		}
	})
}

var RBImbueAuraID = core.NewAuraID()

func (shaman *Shaman) ApplyRockbiterImbue(mh bool, oh bool) {
	if !mh && !oh {
		return
	}

	mhBonus := 0.0
	ohBonus := 0.0
	if weapon := shaman.Equip[proto.ItemSlot_ItemSlotMainHand]; weapon.ID != 0 {
		mhBonus = 62.0 * weapon.SwingSpeed
	}
	if weapon := shaman.Equip[proto.ItemSlot_ItemSlotOffHand]; weapon.ID != 0 {
		ohBonus = 62.0 * weapon.SwingSpeed
	}

	shaman.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: RBImbueAuraID,
			OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.IsWeaponHit() {
					return
				}

				if hitEffect.IsMH() {
					if mh {
						hitEffect.BonusWeaponDamage += mhBonus
					}
				} else if oh {
					hitEffect.BonusWeaponDamage += ohBonus
				}
			},
		}
	})
}
