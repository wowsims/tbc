package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const TwistWindow = time.Millisecond * 400
const SealDuration = time.Second * 30

var SealOfBloodCastActionID = core.ActionID{SpellID: 31892}
var SealOfBloodProcActionID = core.ActionID{SpellID: 31893}

// Handles the cast, gcd, deducts the mana cost
func (paladin *Paladin) setupSealOfBlood() {
	effect := core.SpellEffect{
		IsPhantom:        true,
		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		// should deal 35% weapon deamage
		BaseDamage:     core.BaseDamageConfigMeleeWeapon(core.MainHand, false, 0, 0.35, false),
		OutcomeApplier: core.OutcomeFuncMeleeSpecialHitAndCrit(paladin.DefaultMeleeCritMultiplier()),
	}

	// Apply 2 Handed Weapon Specialization talent
	paladin.applyTwoHandedWeaponSpecializationToSpell(&effect)

	sobProc := paladin.RegisterSpell(core.SpellConfig{
		ActionID:     SealOfBloodProcActionID,
		SpellSchool:  core.SpellSchoolHoly,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(effect),
	})

	// Define the aura
	paladin.SealOfBloodAura = paladin.RegisterAura(&core.Aura{
		Label:    "Seal of Blood",
		Tag:      "Seal",
		ActionID: SealOfBloodProcActionID,
		Duration: SealDuration,

		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellEffect.IsPhantom {
				return
			}
			sobProc.Cast(sim, spellEffect.Target)
		},
	})

	baseCost := 210.0
	paladin.SealOfBlood = paladin.RegisterSpell(core.SpellConfig{
		ActionID: SealOfBloodCastActionID,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost * (1 - 0.03*float64(paladin.Talents.Benediction)),
				GCD:  core.GCDDefault,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, _ *core.Spell) {
			paladin.UpdateSeal(sim, paladin.SealOfBloodAura)
		},
	})
}

var SealOfCommandCastActionID = core.ActionID{SpellID: 20375}
var SealOfCommandProcActionID = core.ActionID{SpellID: 20424}

func (paladin *Paladin) SetupSealOfCommand() {
	effect := core.SpellEffect{
		ProcMask:         core.ProcMaskMeleeMHSpecial,
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		OutcomeApplier:   core.OutcomeFuncMeleeSpecialHitAndCrit(paladin.DefaultMeleeCritMultiplier()),
	}
	paladin.applyTwoHandedWeaponSpecializationToSpell(&effect)

	weaponBaseDamage := core.BaseDamageFuncMeleeWeapon(core.MainHand, false, 0, 0.7, false)
	effect.BaseDamage = core.BaseDamageConfig{
		Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
			return weaponBaseDamage(sim, hitEffect, spell) + 0.29*hitEffect.SpellPower(spell.Character, spell)
		},
		TargetSpellCoefficient: 0.29,
	}

	socProc := paladin.RegisterSpell(core.SpellConfig{
		ActionID:     SealOfCommandProcActionID,
		SpellSchool:  core.SpellSchoolHoly,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(effect),
	})

	ppmm := paladin.AutoAttacks.NewPPMManager(7.0)
	const icdDur = time.Second * 1

	paladin.SealOfCommandAura = paladin.RegisterAura(&core.Aura{
		Label:    "Seal of Command",
		Tag:      "Seal",
		ActionID: SealOfCommandProcActionID,
		Duration: SealDuration,
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMeleeWhiteHit) || spellEffect.IsPhantom {
				return
			}

			if paladin.sealOfCommandICD.IsOnCD(sim) {
				return
			}

			if !ppmm.Proc(sim, true, false, "seal of command") {
				return
			}

			paladin.sealOfCommandICD = core.InternalCD(sim.CurrentTime + icdDur)

			socProc.Cast(sim, spellEffect.Target)
		},
	})

	baseCost := 65.0
	paladin.SealOfCommand = paladin.RegisterSpell(core.SpellConfig{
		ActionID: SealOfCommandCastActionID,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost * (1 - 0.03*float64(paladin.Talents.Benediction)),
				GCD:  core.GCDDefault,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, _ *core.Spell) {
			paladin.UpdateSeal(sim, paladin.SealOfCommandAura)
		},
	})
}

var SealOfTheCrusaderActionID = core.ActionID{SpellID: 27158}

// TODO: Make a universal setup seals function

// Seal of the crusader has a bunch of effects that we realistically don't care about (bonus AP, faster swing speed)
// For now, we'll just use it as a setup to casting Judgement of the Crusader
func (paladin *Paladin) setupSealOfTheCrusader() {
	paladin.SealOfTheCrusaderAura = paladin.RegisterAura(&core.Aura{
		Label:    "Seal of the Crusader",
		Tag:      "Seal",
		ActionID: SealOfTheCrusaderActionID,
		Duration: SealDuration,
	})

	baseCost := 210.0
	paladin.SealOfTheCrusader = paladin.RegisterSpell(core.SpellConfig{
		ActionID: SealOfTheCrusaderActionID,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost * (1 - 0.03*float64(paladin.Talents.Benediction)),
				GCD:  core.GCDDefault,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, _ *core.Spell) {
			paladin.UpdateSeal(sim, paladin.SealOfTheCrusaderAura)
		},
	})
}

// Didn't encode all the functionality of seal of wisdom
// For now, we'll just use it as a setup to casting Judgement of Wisdom
var SealOfWisdomActionID = core.ActionID{SpellID: 27166}

func (paladin *Paladin) setupSealOfWisdom() {
	paladin.SealOfWisdomAura = paladin.RegisterAura(&core.Aura{
		Label:    "Seal of Wisdom",
		Tag:      "Seal",
		ActionID: SealOfWisdomActionID,
		Duration: SealDuration,
	})

	baseCost := 270.0
	paladin.SealOfWisdom = paladin.RegisterSpell(core.SpellConfig{
		ActionID: SealOfWisdomActionID,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost * (1 - 0.03*float64(paladin.Talents.Benediction)),
				GCD:  core.GCDDefault,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, _ *core.Spell) {
			paladin.UpdateSeal(sim, paladin.SealOfWisdomAura)
		},
	})
}

func (paladin *Paladin) UpdateSeal(sim *core.Simulation, newSeal *core.Aura) {
	if paladin.CurrentSeal == paladin.SealOfCommandAura {
		// Technically the current expiration could be shorter than 0.4 seconds
		// TO-DO: Lookup behavior when seal of command is twisted at shorter than 0.4 seconds duration
		expiresAt := sim.CurrentTime + TwistWindow
		paladin.CurrentSeal.UpdateExpires(expiresAt)

		// This is a hack to get the sim to process and log the SoC aura expiring at the right time
		if sim.Options.Iterations == 1 {
			sim.AddPendingAction(&core.PendingAction{
				NextActionAt: expiresAt,
				OnAction:     func(_ *core.Simulation) {},
			})
		}
	} else if paladin.CurrentSeal != nil {
		paladin.CurrentSeal.Deactivate(sim)
	}

	paladin.CurrentSeal = newSeal
	newSeal.Activate(sim)
}
