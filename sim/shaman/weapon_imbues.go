package shaman

import (
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
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

	wftempl := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: 25505},
				Character:           &shaman.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				CritMultiplier:      shaman.DefaultMeleeCritMultiplier(),
				IsPhantom:           true,
			},
		},
	}

	baseEffect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier: 1.0,
			ThreatMultiplier: 1.0,
			BonusAttackPower: apBonus,
		},
	}
	if shaman.Talents.SpiritWeapons {
		baseEffect.ThreatMultiplier *= 0.7
	}

	wftempl.Effects = []core.SpellHitEffect{
		baseEffect,
		baseEffect,
	}

	weaponDamageMultiplier := 1 + math.Round(float64(shaman.Talents.ElementalWeapons)*13.33)/100
	mhBaseDamage := core.BaseDamageFuncMeleeWeapon(core.MainHand, false, 0, weaponDamageMultiplier, true)
	ohBaseDamage := core.BaseDamageFuncMeleeWeapon(core.OffHand, false, 0, weaponDamageMultiplier, true)

	wfTemplate := core.NewSimpleSpellTemplate(wftempl)

	wfAtk := core.SimpleSpell{}
	shaman.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		var icd core.InternalCD
		const icdDur = time.Second * 3

		return core.Aura{
			ID: WFImbueAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				// ProcMask: 20
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellCast.IsPhantom {
					return
				}

				isMHHit := spellEffect.IsMH()
				if (!mh && isMHHit) || (!oh && !isMHHit) {
					return // cant proc if not enchanted
				}
				if icd.IsOnCD(sim) {
					return
				}
				if sim.RandomFloat("Windfury Imbue") > proc {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)

				wfTemplate.Apply(&wfAtk)
				// Set so only the proc'd hand attacks
				attackProc := core.ProcMaskMeleeMHSpecial
				if isMHHit {
					wfAtk.ActionID.Tag = 1
				} else {
					wfAtk.ActionID.Tag = 2
					attackProc = core.ProcMaskMeleeOHSpecial
				}
				for i := 0; i < 2; i++ {
					wfAtk.Effects[i].Target = spellEffect.Target
					wfAtk.Effects[i].ProcMask = attackProc
					if isMHHit {
						wfAtk.Effects[i].BaseDamage = mhBaseDamage
					} else {
						// For whatever reason, OH penalty does not apply to the bonus AP from WF OH
						// hits. Implement this by doubling the AP bonus we provide.
						wfAtk.Effects[i].BonusAttackPower += apBonus
						wfAtk.Effects[i].BaseDamage = ohBaseDamage
					}
				}
				wfAtk.Cast(sim)
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
				ActionID:            core.ActionID{SpellID: 25489},
				Character:           &shaman.Character,
				CritRollCategory:    core.CritRollCategoryMagical,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolFire,
				IsPhantom:           true,
				CritMultiplier:      shaman.DefaultSpellCritMultiplier(),
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier: 1,
				ThreatMultiplier: 1,
			},
		},
	}
	ftTmpl.Effect.DamageMultiplier *= 1 + 0.05*float64(shaman.Talents.ElementalWeapons)

	mhTmpl := ftTmpl
	ohTmpl := ftTmpl

	if weapon := shaman.GetMHWeapon(); weapon != nil {
		baseDamage := weapon.SwingSpeed * 35.0
		mhTmpl.Effect.BaseDamage = core.BaseDamageFuncMagic(baseDamage, baseDamage, 0.1)
	}
	if weapon := shaman.GetOHWeapon(); weapon != nil {
		baseDamage := weapon.SwingSpeed * 35.0
		ohTmpl.Effect.BaseDamage = core.BaseDamageFuncMagic(baseDamage, baseDamage, 0.1)
	}

	mhTemplate := core.NewSimpleSpellTemplate(mhTmpl)
	ohTemplate := core.NewSimpleSpellTemplate(ohTmpl)
	ftSpell := core.SimpleSpell{}

	shaman.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: FTImbueAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellCast.IsPhantom {
					return
				}

				isMHHit := spellEffect.IsMH()
				if (isMHHit && !mh) || (!isMHHit && !oh) {
					return // cant proc if not enchanted
				}

				if isMHHit {
					mhTemplate.Apply(&ftSpell)
				} else {
					ohTemplate.Apply(&ftSpell)
				}
				ftSpell.Effect.Target = spellEffect.Target
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
				ActionID:            core.ActionID{SpellID: 25500},
				Character:           &shaman.Character,
				CritRollCategory:    core.CritRollCategoryMagical,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolFrost,
				IsPhantom:           true,
				CritMultiplier:      shaman.DefaultSpellCritMultiplier(),
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier: 1,
				ThreatMultiplier: 1,
			},
			BaseDamage: core.BaseDamageFuncMagic(246, 246, 0.1),
		},
	}
	fbTmpl.Effect.DamageMultiplier *= 1 + 0.05*float64(shaman.Talents.ElementalWeapons)

	fbTemplate := core.NewSimpleSpellTemplate(fbTmpl)
	fbSpell := core.SimpleSpell{}

	shaman.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		ppmm := shaman.AutoAttacks.NewPPMManager(9.0)
		return core.Aura{
			ID: FBImbueAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellCast.IsPhantom {
					return
				}

				isMHHit := spellEffect.IsMH()
				if (isMHHit && !mh) || (!isMHHit && !oh) {
					return // cant proc if not enchanted
				}

				if !ppmm.Proc(sim, isMHHit, false, "Frostbrand Weapon") {
					return
				}

				fbTemplate.Apply(&fbSpell)
				fbSpell.Effect.Target = spellEffect.Target
				fbSpell.Init(sim)
				fbSpell.Cast(sim)
			},
		}
	})
}

func (shaman *Shaman) ApplyRockbiterImbue(mh bool, oh bool) {
	if weapon := shaman.GetMHWeapon(); mh && weapon != nil {
		bonus := 62.0 * weapon.SwingSpeed
		shaman.AutoAttacks.MH.BaseDamageMin += bonus
		shaman.AutoAttacks.MH.BaseDamageMax += bonus
	}
	if weapon := shaman.GetOHWeapon(); oh && weapon != nil {
		bonus := 62.0 * weapon.SwingSpeed
		shaman.AutoAttacks.MH.BaseDamageMin += bonus
		shaman.AutoAttacks.MH.BaseDamageMax += bonus
	}
}
