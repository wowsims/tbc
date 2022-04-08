package shaman

import (
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

var TotemOfTheAstralWinds int32 = 27815

func (shaman *Shaman) newWindfuryImbueSpell(isMH bool) *core.Spell {
	apBonus := 475.0
	if shaman.Equip[proto.ItemSlot_ItemSlotRanged].ID == TotemOfTheAstralWinds {
		apBonus += 80
	}

	wftempl := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: 25505},
				Character:   &shaman.Character,
				SpellSchool: core.SpellSchoolPhysical,
			},
		},
	}

	baseEffect := core.SpellEffect{
		BonusAttackPower: apBonus,
		IsPhantom:        true,
		DamageMultiplier: 1.0,
		ThreatMultiplier: core.TernaryFloat64(shaman.Talents.SpiritWeapons, 0.7, 1),
		OutcomeApplier:   core.OutcomeFuncMeleeSpecialHitAndCrit(shaman.DefaultMeleeCritMultiplier()),
	}

	weaponDamageMultiplier := 1 + math.Round(float64(shaman.Talents.ElementalWeapons)*13.33)/100
	if isMH {
		wftempl.ActionID.Tag = 1
		baseEffect.ProcMask = core.ProcMaskMeleeMHSpecial
		baseEffect.BaseDamage = core.BaseDamageConfigMeleeWeapon(core.MainHand, false, 0, weaponDamageMultiplier, true)
	} else {
		wftempl.ActionID.Tag = 2
		baseEffect.ProcMask = core.ProcMaskMeleeOHSpecial
		baseEffect.BaseDamage = core.BaseDamageConfigMeleeWeapon(core.OffHand, false, 0, weaponDamageMultiplier, true)

		// For whatever reason, OH penalty does not apply to the bonus AP from WF OH
		// hits. Implement this by doubling the AP bonus we provide.
		baseEffect.BonusAttackPower += apBonus
	}

	effects := []core.SpellEffect{
		baseEffect,
		baseEffect,
	}

	return shaman.RegisterSpell(core.SpellConfig{
		Template:     wftempl,
		ApplyEffects: core.ApplyEffectFuncDamageMultipleTargeted(effects),
	})
}

func (shaman *Shaman) ApplyWindfuryImbue(mh bool, oh bool) {
	if !mh && !oh {
		return
	}

	var proc = 0.2
	if mh && oh {
		proc = 0.36
	}

	mhSpell := shaman.newWindfuryImbueSpell(true)
	ohSpell := shaman.newWindfuryImbueSpell(false)

	shaman.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		var icd core.InternalCD
		const icdDur = time.Second * 3

		return shaman.GetOrRegisterAura(&core.Aura{
			Label: "Windfury Imbue",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				// ProcMask: 20
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellEffect.IsPhantom {
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

				if isMHHit {
					mhSpell.Cast(sim, spellEffect.Target)
				} else {
					ohSpell.Cast(sim, spellEffect.Target)
				}
			},
		})
	})
}

func (shaman *Shaman) newFlametongueImbueSpell(isMH bool) *core.Spell {
	ftTmpl := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: 25489},
				Character:   &shaman.Character,
				SpellSchool: core.SpellSchoolFire,
			},
		},
	}

	effect := core.SpellEffect{
		IsPhantom:        true,
		DamageMultiplier: 1 + 0.05*float64(shaman.Talents.ElementalWeapons),
		ThreatMultiplier: 1,
		OutcomeApplier:   core.OutcomeFuncMagicHitAndCrit(shaman.DefaultSpellCritMultiplier()),
	}

	if isMH {
		if weapon := shaman.GetMHWeapon(); weapon != nil {
			baseDamage := weapon.SwingSpeed * 35.0
			effect.BaseDamage = core.BaseDamageConfigMagic(baseDamage, baseDamage, 0.1)
		}
	} else {
		if weapon := shaman.GetOHWeapon(); weapon != nil {
			baseDamage := weapon.SwingSpeed * 35.0
			effect.BaseDamage = core.BaseDamageConfigMagic(baseDamage, baseDamage, 0.1)
		}
	}

	return shaman.RegisterSpell(core.SpellConfig{
		Template:     ftTmpl,
		ModifyCast:   core.ModifyCastAssignTarget,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(effect),
	})
}

func (shaman *Shaman) ApplyFlametongueImbue(mh bool, oh bool) {
	if !mh && !oh {
		return
	}

	mhSpell := shaman.newFlametongueImbueSpell(true)
	ohSpell := shaman.newFlametongueImbueSpell(false)

	shaman.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return shaman.GetOrRegisterAura(&core.Aura{
			Label: "Flametongue Imbue",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellEffect.IsPhantom {
					return
				}

				isMHHit := spellEffect.IsMH()
				if (isMHHit && !mh) || (!isMHHit && !oh) {
					return // cant proc if not enchanted
				}

				if isMHHit {
					mhSpell.Cast(sim, spellEffect.Target)
				} else {
					ohSpell.Cast(sim, spellEffect.Target)
				}
			},
		})
	})
}

func (shaman *Shaman) newFrostbrandImbueSpell(isMH bool) *core.Spell {
	fbTmpl := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: 25500},
				Character:   &shaman.Character,
				SpellSchool: core.SpellSchoolFrost,
			},
		},
	}

	effect := core.SpellEffect{
		IsPhantom:        true,
		DamageMultiplier: 1 + 0.05*float64(shaman.Talents.ElementalWeapons),
		ThreatMultiplier: 1,
		BaseDamage:       core.BaseDamageConfigMagic(246, 246, 0.1),
		OutcomeApplier:   core.OutcomeFuncMagicHitAndCrit(shaman.DefaultSpellCritMultiplier()),
	}

	return shaman.RegisterSpell(core.SpellConfig{
		Template:     fbTmpl,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(effect),
	})
}

func (shaman *Shaman) ApplyFrostbrandImbue(mh bool, oh bool) {
	if !mh && !oh {
		return
	}

	mhSpell := shaman.newFrostbrandImbueSpell(true)
	ohSpell := shaman.newFrostbrandImbueSpell(false)

	shaman.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		ppmm := shaman.AutoAttacks.NewPPMManager(9.0)
		return shaman.GetOrRegisterAura(&core.Aura{
			Label: "Frostbrand Imbue",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellEffect.IsPhantom {
					return
				}

				isMHHit := spellEffect.IsMH()
				if (isMHHit && !mh) || (!isMHHit && !oh) {
					return // cant proc if not enchanted
				}

				if !ppmm.Proc(sim, isMHHit, false, "Frostbrand Weapon") {
					return
				}

				if isMHHit {
					mhSpell.Cast(sim, spellEffect.Target)
				} else {
					ohSpell.Cast(sim, spellEffect.Target)
				}
			},
		})
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
