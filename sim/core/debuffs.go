package core

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func applyDebuffEffects(target *Unit, debuffs proto.Debuffs) {
	if debuffs.Misery {
		MakePermanent(MiseryAura(target, 5))
	}

	if debuffs.ShadowWeaving {
		MakePermanent(ShadowWeavingAura(target, 5))
	}

	if debuffs.JudgementOfWisdom {
		MakePermanent(JudgementOfWisdomAura(target))
	}

	if debuffs.ImprovedSealOfTheCrusader {
		MakePermanent(JudgementOfTheCrusaderAura(target, 3, 0, 1.0))
	}

	if debuffs.CurseOfElements != proto.TristateEffect_TristateEffectMissing {
		MakePermanent(CurseOfElementsAura(target, GetTristateValueInt32(debuffs.CurseOfElements, 0, 3)))
	}

	if debuffs.IsbUptime > 0.0 {
		uptime := MinFloat(1.0, debuffs.IsbUptime)
		isbAura := MakePermanent(ImprovedShadowBoltAura(target, 5, uptime))
		if uptime != 1.0 {
			isbAura.OnDoneIteration = func(aura *Aura, _ *Simulation) {
				aura.metrics.Uptime = time.Duration(float64(aura.metrics.Uptime) * uptime)
			}
		}
	}

	if debuffs.ImprovedScorch {
		MakePermanent(ImprovedScorchAura(target, 5))
	}

	if debuffs.WintersChill {
		MakePermanent(WintersChillAura(target, 5))
	}

	if debuffs.BloodFrenzy {
		MakePermanent(BloodFrenzyAura(target, 2))
	}

	if debuffs.GiftOfArthas {
		MakePermanent(GiftOfArthasAura(target))
	}

	if debuffs.Mangle {
		MakePermanent(MangleAura(target))
	}

	if debuffs.ExposeArmor != proto.TristateEffect_TristateEffectMissing {
		ScheduledExposeArmorAura(target, GetTristateValueInt32(debuffs.ExposeArmor, 0, 2))
	}

	if debuffs.SunderArmor {
		ScheduledSunderArmorAura(target)
	}

	if debuffs.FaerieFire != proto.TristateEffect_TristateEffectMissing {
		MakePermanent(FaerieFireAura(target, GetTristateValueInt32(debuffs.FaerieFire, 0, 3)))
	}

	if debuffs.CurseOfRecklessness {
		MakePermanent(CurseOfRecklessnessAura(target))
	}

	if debuffs.ExposeWeaknessUptime > 0 && debuffs.ExposeWeaknessHunterAgility > 0 {
		uptime := MinFloat(1.0, debuffs.ExposeWeaknessUptime)
		ewAura := MakePermanent(ExposeWeaknessAura(target, debuffs.ExposeWeaknessHunterAgility, uptime))
		if uptime != 1.0 {
			ewAura.OnDoneIteration = func(aura *Aura, _ *Simulation) {
				aura.metrics.Uptime = time.Duration(float64(aura.metrics.Uptime) * uptime)
			}
		}
	}

	if debuffs.HuntersMark != proto.TristateEffect_TristateEffectMissing {
		if debuffs.HuntersMark == proto.TristateEffect_TristateEffectImproved {
			MakePermanent(HuntersMarkAura(target, 5, true))
		} else {
			MakePermanent(HuntersMarkAura(target, 0, true))
		}
	}
}

func MiseryAura(target *Unit, numPoints int32) *Aura {
	multiplier := 1.0 + 0.01*float64(numPoints)

	return target.GetOrRegisterAura(Aura{
		Label:    "Misery-" + strconv.Itoa(int(numPoints)),
		Tag:      "Misery",
		ActionID: ActionID{SpellID: 33195},
		Duration: time.Second * 24,
		Priority: float64(numPoints),
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.ArcaneDamageTakenMultiplier *= multiplier
			aura.Unit.PseudoStats.FireDamageTakenMultiplier *= multiplier
			aura.Unit.PseudoStats.FrostDamageTakenMultiplier *= multiplier
			aura.Unit.PseudoStats.HolyDamageTakenMultiplier *= multiplier
			aura.Unit.PseudoStats.NatureDamageTakenMultiplier *= multiplier
			aura.Unit.PseudoStats.ShadowDamageTakenMultiplier *= multiplier
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.ArcaneDamageTakenMultiplier /= multiplier
			aura.Unit.PseudoStats.FireDamageTakenMultiplier /= multiplier
			aura.Unit.PseudoStats.FrostDamageTakenMultiplier /= multiplier
			aura.Unit.PseudoStats.HolyDamageTakenMultiplier /= multiplier
			aura.Unit.PseudoStats.NatureDamageTakenMultiplier /= multiplier
			aura.Unit.PseudoStats.ShadowDamageTakenMultiplier /= multiplier
		},
	})
}

func ShadowWeavingAura(target *Unit, startingStacks int32) *Aura {
	return target.GetOrRegisterAura(Aura{
		Label:     "Shadow Weaving",
		ActionID:  ActionID{SpellID: 15334},
		Duration:  time.Second * 15,
		MaxStacks: 5,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.SetStacks(sim, startingStacks)
		},
		OnStacksChange: func(aura *Aura, sim *Simulation, oldStacks int32, newStacks int32) {
			aura.Unit.PseudoStats.ShadowDamageTakenMultiplier /= 1.0 + 0.02*float64(oldStacks)
			aura.Unit.PseudoStats.ShadowDamageTakenMultiplier *= 1.0 + 0.02*float64(newStacks)
		},
	})
}

func JudgementOfWisdomAura(target *Unit) *Aura {
	const mana = 74 / 2 // 50% proc
	actionID := ActionID{SpellID: 27164}

	return target.GetOrRegisterAura(Aura{
		Label:    "Judgement of Wisdom",
		ActionID: actionID,
		Duration: time.Second * 20,
		OnSpellHitTaken: func(aura *Aura, sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
			// TODO: This check is purely to maintain behavior during refactoring. Should be removed when possible.
			if !spellEffect.ProcMask.Matches(ProcMaskMeleeOrRanged) && !spellEffect.Landed() {
				return
			}
			if spellEffect.IsPhantom {
				return // Phantom spells (Romulo's, Lightning Capacitor, etc) don't proc JoW.
			}

			unit := spell.Unit
			if unit.HasManaBar() {
				unit.AddMana(sim, mana, actionID, false)
			}

			if spell.ActionID.SpellID == 35395 {
				aura.Refresh(sim)
			}
		},
	})
}

func JudgementOfTheCrusaderAura(target *Unit, level int32, flatBonus float64, percentBonus float64) *Aura {
	bonusCrit := float64(level) * SpellCritRatingPerCritChance

	totalSP := 219*percentBonus + flatBonus
	return target.GetOrRegisterAura(Aura{
		Label:    "Judgement of the Crusader-" + strconv.Itoa(int(level)),
		Tag:      "Judgement of the Crusader",
		ActionID: ActionID{SpellID: 27159},
		Duration: time.Second * 20,
		Priority: float64(level),
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.BonusHolyDamageTaken += totalSP
			aura.Unit.PseudoStats.BonusCritRating += bonusCrit
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.BonusHolyDamageTaken -= totalSP
			aura.Unit.PseudoStats.BonusCritRating -= bonusCrit
		},
		OnSpellHitTaken: func(aura *Aura, sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
			if spell.ActionID.SpellID == 35395 {
				aura.Refresh(sim)
			}
		},
	})
}

func CurseOfElementsAura(target *Unit, points int32) *Aura {
	multiplier := 1.1 + 0.01*float64(points)

	return target.GetOrRegisterAura(Aura{
		Label:    "Curse of Elements-" + strconv.Itoa(int(points)),
		Tag:      "Curse of Elements",
		ActionID: ActionID{SpellID: 27228},
		Priority: float64(points),
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.ArcaneDamageTakenMultiplier *= multiplier
			aura.Unit.PseudoStats.FireDamageTakenMultiplier *= multiplier
			aura.Unit.PseudoStats.FrostDamageTakenMultiplier *= multiplier
			aura.Unit.PseudoStats.ShadowDamageTakenMultiplier *= multiplier
			aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.ArcaneResistance: -88, stats.FireResistance: -88, stats.FrostResistance: -88, stats.ShadowResistance: -88})
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.ArcaneDamageTakenMultiplier /= multiplier
			aura.Unit.PseudoStats.FireDamageTakenMultiplier /= multiplier
			aura.Unit.PseudoStats.FrostDamageTakenMultiplier /= multiplier
			aura.Unit.PseudoStats.ShadowDamageTakenMultiplier /= multiplier
			aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.ArcaneResistance: 88, stats.FireResistance: 88, stats.FrostResistance: 88, stats.ShadowResistance: 88})
		},
	})
}

// If uptime is 0, makes an aura for a real warlock. Otherwise makes a group approximation.
func ImprovedShadowBoltAura(target *Unit, points int32, uptime float64) *Aura {
	bonus := 0.04 * float64(points)
	multiplier := 1 + bonus
	if uptime != 0 {
		multiplier = 1 + bonus*uptime
	}

	config := Aura{
		Label:     "ImprovedShadowBolt-" + strconv.Itoa(int(points)),
		Tag:       "ImprovedShadowBolt",
		ActionID:  ActionID{SpellID: 17803},
		Duration:  time.Second * 12,
		Priority:  float64(points),
		MaxStacks: 4,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.ShadowDamageTakenMultiplier *= multiplier
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.ShadowDamageTakenMultiplier /= multiplier
		},
	}

	if uptime == 0 {
		config.OnSpellHitTaken = func(aura *Aura, sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
			if spell.SpellSchool != SpellSchoolShadow {
				return
			}
			if !spellEffect.Landed() || spellEffect.Damage == 0 || spellEffect.IsPhantom || spellEffect.ProcMask == 0 {
				return
			}
			aura.RemoveStack(sim)
		}
	}

	return target.GetOrRegisterAura(config)
}

var BloodFrenzyActionID = ActionID{SpellID: 29859}

func BloodFrenzyAura(target *Unit, points int32) *Aura {
	multiplier := 1 + 0.02*float64(points)
	return target.GetOrRegisterAura(Aura{
		Label:    "Blood Frenzy-" + strconv.Itoa(int(points)),
		Tag:      "Blood Frenzy",
		ActionID: BloodFrenzyActionID,
		// No fixed duration, lasts as long as the bleed that activates it.
		Priority: float64(points),
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.PhysicalDamageTakenMultiplier *= multiplier
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.PhysicalDamageTakenMultiplier /= multiplier
		},
	})
}

func GiftOfArthasAura(target *Unit) *Aura {
	return target.GetOrRegisterAura(Aura{
		Label:    "Gift of Arthas",
		ActionID: ActionID{SpellID: 11374},
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.BonusPhysicalDamageTaken += 8
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.BonusPhysicalDamageTaken -= 8
		},
	})
}

func MangleAura(target *Unit) *Aura {
	return target.GetOrRegisterAura(Aura{
		Label:    "Mangle",
		ActionID: ActionID{SpellID: 33876},
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.PeriodicPhysicalDamageTakenMultiplier *= 1.3
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.PeriodicPhysicalDamageTakenMultiplier /= 1.3
		},
	})
}

var ImprovedScorchAuraLabel = "Improved Scorch"

func ImprovedScorchAura(target *Unit, startingStacks int32) *Aura {
	return target.GetOrRegisterAura(Aura{
		Label:     ImprovedScorchAuraLabel,
		ActionID:  ActionID{SpellID: 12873},
		Duration:  time.Second * 30,
		MaxStacks: 5,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.SetStacks(sim, startingStacks)
		},
		OnStacksChange: func(aura *Aura, sim *Simulation, oldStacks int32, newStacks int32) {
			aura.Unit.PseudoStats.FireDamageTakenMultiplier /= 1.0 + 0.03*float64(oldStacks)
			aura.Unit.PseudoStats.FireDamageTakenMultiplier *= 1.0 + 0.03*float64(newStacks)
		},
	})
}

var WintersChillAuraLabel = "Winter's Chill"

func WintersChillAura(target *Unit, startingStacks int32) *Aura {
	return target.GetOrRegisterAura(Aura{
		Label:     WintersChillAuraLabel,
		ActionID:  ActionID{SpellID: 28595},
		Duration:  time.Second * 15,
		MaxStacks: 5,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.SetStacks(sim, startingStacks)
		},
		OnStacksChange: func(aura *Aura, sim *Simulation, oldStacks int32, newStacks int32) {
			aura.Unit.PseudoStats.BonusFrostCritRating += 2 * SpellCritRatingPerCritChance * float64(newStacks-oldStacks)
		},
	})
}

func FaerieFireAura(target *Unit, level int32) *Aura {
	const armorReduction = 610

	return target.GetOrRegisterAura(Aura{
		Label:    "Faerie Fire-" + strconv.Itoa(int(level)),
		Tag:      "Faerie Fire",
		ActionID: ActionID{SpellID: 26993},
		Duration: time.Second * 40,
		Priority: float64(level),
		OnGain: func(aura *Aura, sim *Simulation) {
			target.AddStatDynamic(sim, stats.Armor, -armorReduction)
			target.PseudoStats.BonusMeleeHitRating += float64(level) * MeleeHitRatingPerHitChance
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			target.AddStatDynamic(sim, stats.Armor, armorReduction)
			target.PseudoStats.BonusMeleeHitRating -= float64(level) * MeleeHitRatingPerHitChance
		},
	})
}

var SunderArmorAuraLabel = "Sunder Armor"
var SunderExposeAuraTag = "SunderExpose"

func SunderArmorAura(target *Unit, startingStacks int32) *Aura {
	armorReductionPerStack := 520.0

	return target.GetOrRegisterAura(Aura{
		Label:     SunderArmorAuraLabel,
		Tag:       SunderExposeAuraTag,
		ActionID:  ActionID{SpellID: 25225},
		Duration:  time.Second * 30,
		MaxStacks: 5,
		Priority:  armorReductionPerStack * 5,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.SetStacks(sim, startingStacks)
		},
		OnStacksChange: func(aura *Aura, sim *Simulation, oldStacks int32, newStacks int32) {
			target.AddStatDynamic(sim, stats.Armor, float64(oldStacks-newStacks)*armorReductionPerStack)
		},
	})
}

func ScheduledSunderArmorAura(target *Unit) *Aura {
	aura := SunderArmorAura(target, 1)
	aura.Duration = NeverExpires
	aura.OnReset = func(aura *Aura, sim *Simulation) {
		aura.Activate(sim)
		StartPeriodicAction(sim, PeriodicActionOptions{
			Period:   time.Duration(1.5 * float64(time.Second)),
			NumTicks: 4,
			Priority: ActionPriorityDOT, // High prio so it comes before actual warrior sunders.
			OnAction: func(sim *Simulation) {
				if aura.IsActive() {
					aura.AddStack(sim)
				}
			},
		})
	}
	return aura
}

func ExposeArmorAura(target *Unit, talentPoints int32) *Aura {
	armorReduction := 2050.0 * (1.0 + 0.25*float64(talentPoints))

	return target.GetOrRegisterAura(Aura{
		Label:    "ExposeArmor-" + strconv.Itoa(int(talentPoints)),
		Tag:      SunderExposeAuraTag,
		ActionID: ActionID{SpellID: 26866},
		Duration: time.Second * 30,
		Priority: armorReduction,
		OnGain: func(aura *Aura, sim *Simulation) {
			target.AddStatDynamic(sim, stats.Armor, -armorReduction)
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			target.AddStatDynamic(sim, stats.Armor, armorReduction)
		},
	})
}

func ScheduledExposeArmorAura(target *Unit, talentPoints int32) *Aura {
	aura := ExposeArmorAura(target, talentPoints)
	aura.Duration = NeverExpires
	aura.OnReset = func(aura *Aura, sim *Simulation) {
		StartPeriodicAction(sim, PeriodicActionOptions{
			Period:   time.Duration(10.0 * float64(time.Second)),
			NumTicks: 1,
			OnAction: func(sim *Simulation) {
				aura.Activate(sim)
			},
		})
	}
	return aura
}

func CurseOfRecklessnessAura(target *Unit) *Aura {
	armorReduction := 800.0

	return target.GetOrRegisterAura(Aura{
		Label:    "Curse of Recklessness",
		ActionID: ActionID{SpellID: 27226},
		Duration: time.Minute * 2,
		OnGain: func(aura *Aura, sim *Simulation) {
			target.AddStatDynamic(sim, stats.Armor, -armorReduction)
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			target.AddStatDynamic(sim, stats.Armor, armorReduction)
		},
	})
}

// Multiplier is for accomodating uptime %. For a real hunter, always pass 1.0
func ExposeWeaknessAura(target *Unit, hunterAgility float64, multiplier float64) *Aura {
	apBonus := hunterAgility * 0.25 * multiplier

	return target.GetOrRegisterAura(Aura{
		Label:    "ExposeWeakness-" + strconv.Itoa(int(hunterAgility)),
		Tag:      "ExposeWeakness",
		ActionID: ActionID{SpellID: 34503},
		Duration: time.Second * 7,
		Priority: apBonus,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.BonusMeleeAttackPower += aura.Priority
			aura.Unit.PseudoStats.BonusRangedAttackPower += aura.Priority
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.BonusMeleeAttackPower -= aura.Priority
			aura.Unit.PseudoStats.BonusRangedAttackPower -= aura.Priority
		},
	})
}

func HuntersMarkAura(target *Unit, points int32, fullyStacked bool) *Aura {
	const baseRangedBonus = 110.0
	const bonusPerStack = 11.0
	const maxStacks = 30
	meleeBonus := baseRangedBonus * 0.2 * float64(points)

	startingStacks := int32(0)
	if fullyStacked {
		startingStacks = maxStacks
	}

	priority := float64(points)
	if fullyStacked {
		// Add a half point so that permanent versions always win.
		priority += 0.5
	}

	return target.GetOrRegisterAura(Aura{
		Label:     "HuntersMark-" + strconv.Itoa(int(points)),
		Tag:       "HuntersMark",
		ActionID:  ActionID{SpellID: 14325},
		Duration:  NeverExpires,
		MaxStacks: 30,
		Priority:  priority,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.BonusMeleeAttackPower += meleeBonus
			aura.Unit.PseudoStats.BonusRangedAttackPower += baseRangedBonus
			aura.SetStacks(sim, startingStacks)
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.BonusMeleeAttackPower -= meleeBonus
			aura.Unit.PseudoStats.BonusRangedAttackPower -= baseRangedBonus
		},
		OnStacksChange: func(aura *Aura, sim *Simulation, oldStacks int32, newStacks int32) {
			aura.Unit.PseudoStats.BonusRangedAttackPower += bonusPerStack * float64(newStacks-oldStacks)
		},
		OnSpellHitTaken: func(aura *Aura, sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
			if spellEffect.ProcMask.Matches(ProcMaskRanged) && spellEffect.Landed() {
				aura.AddStack(sim)
			}
		},
	})
}

func DemoralizingShoutAura(target *Unit, boomingVoicePts int32, impDemoShoutPts int32) *Aura {
	duration := time.Duration(float64(time.Second*30) * (1 + 0.1*float64(boomingVoicePts)))

	return target.GetOrRegisterAura(Aura{
		Label:    "DemoralizingShout-" + strconv.Itoa(int(impDemoShoutPts)),
		Tag:      "DemoralizingShout",
		ActionID: ActionID{SpellID: 25203},
		Duration: duration,
		Priority: float64(impDemoShoutPts),
		OnGain: func(aura *Aura, sim *Simulation) {
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
		},
	})
}

func ThunderClapAura(target *Unit, impThunderClapPts int32) *Aura {
	return target.GetOrRegisterAura(Aura{
		Label:    "ThunderClap-" + strconv.Itoa(int(impThunderClapPts)),
		Tag:      "ThunderClap",
		ActionID: ActionID{SpellID: 25264},
		Duration: time.Second * 30,
		Priority: float64(impThunderClapPts),
		OnGain: func(aura *Aura, sim *Simulation) {
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
		},
	})
}

func ScorpidStingAura(target *Unit) *Aura {
	return target.GetOrRegisterAura(Aura{
		Label:    "Scorpid Sting",
		ActionID: ActionID{SpellID: 3043},
		Duration: time.Second * 20,
		OnGain: func(aura *Aura, sim *Simulation) {
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
		},
	})
}
