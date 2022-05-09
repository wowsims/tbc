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

	actionID := core.ActionID{SpellID: 25505}

	baseEffect := core.SpellEffect{
		BonusAttackPower: apBonus,
		IsPhantom:        true,
		DamageMultiplier: 1.0,
		ThreatMultiplier: core.TernaryFloat64(shaman.Talents.SpiritWeapons, 0.7, 1),
		OutcomeApplier:   shaman.OutcomeFuncMeleeSpecialHitAndCrit(shaman.DefaultMeleeCritMultiplier()),
	}

	weaponDamageMultiplier := 1 + math.Round(float64(shaman.Talents.ElementalWeapons)*13.33)/100
	if isMH {
		actionID.Tag = 1
		baseEffect.ProcMask = core.ProcMaskMeleeMHSpecial
		baseEffect.BaseDamage = core.BaseDamageConfigMeleeWeapon(core.MainHand, false, 0, weaponDamageMultiplier, true)
	} else {
		actionID.Tag = 2
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
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

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

	icd := core.Cooldown{
		Timer:    shaman.NewTimer(),
		Duration: time.Second * 3,
	}

	shaman.RegisterAura(core.Aura{
		Label:    "Windfury Imbue",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			// ProcMask: 20
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellEffect.IsPhantom {
				return
			}

			isMHHit := spellEffect.IsMH()
			if (!mh && isMHHit) || (!oh && !isMHHit) {
				return // cant proc if not enchanted
			}
			if !icd.IsReady(sim) {
				return
			}
			if sim.RandomFloat("Windfury Imbue") > proc {
				return
			}
			icd.Use(sim)

			if isMHHit {
				mhSpell.Cast(sim, spellEffect.Target)
			} else {
				ohSpell.Cast(sim, spellEffect.Target)
			}
		},
	})
}

func (shaman *Shaman) newFlametongueImbueSpell(isMH bool) *core.Spell {
	effect := core.SpellEffect{
		IsPhantom:           true,
		BonusSpellHitRating: float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance,

		DamageMultiplier: 1 + 0.05*float64(shaman.Talents.ElementalWeapons),
		ThreatMultiplier: 1,
		OutcomeApplier:   shaman.OutcomeFuncMagicHitAndCrit(shaman.DefaultSpellCritMultiplier()),
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
		ActionID:     core.ActionID{SpellID: 25489},
		SpellSchool:  core.SpellSchoolFire,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(effect),
	})
}

func (shaman *Shaman) ApplyFlametongueImbue(mh bool, oh bool) {
	if !mh && !oh {
		return
	}

	mhSpell := shaman.newFlametongueImbueSpell(true)
	ohSpell := shaman.newFlametongueImbueSpell(false)

	shaman.RegisterAura(core.Aura{
		Label:    "Flametongue Imbue",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
}

func (shaman *Shaman) newFrostbrandImbueSpell(isMH bool) *core.Spell {
	return shaman.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 25500},
		SpellSchool: core.SpellSchoolFrost,

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			IsPhantom:           true,
			BonusSpellHitRating: float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance,

			DamageMultiplier: 1 + 0.05*float64(shaman.Talents.ElementalWeapons),
			ThreatMultiplier: 1,

			BaseDamage:     core.BaseDamageConfigMagic(246, 246, 0.1),
			OutcomeApplier: shaman.OutcomeFuncMagicHitAndCrit(shaman.DefaultSpellCritMultiplier()),
		}),
	})
}

func (shaman *Shaman) ApplyFrostbrandImbue(mh bool, oh bool) {
	if !mh && !oh {
		return
	}

	mhSpell := shaman.newFrostbrandImbueSpell(true)
	ohSpell := shaman.newFrostbrandImbueSpell(false)
	ppmm := shaman.AutoAttacks.NewPPMManager(9.0)

	shaman.RegisterAura(core.Aura{
		Label:    "Frostbrand Imbue",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
