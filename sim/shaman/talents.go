package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (shaman *Shaman) applyTalents() {
	if shaman.Talents.NaturesGuidance > 0 {
		shaman.AddStat(stats.SpellHit, float64(shaman.Talents.NaturesGuidance)*1*core.SpellHitRatingPerHitChance)
		shaman.AddStat(stats.MeleeHit, float64(shaman.Talents.NaturesGuidance)*1*core.MeleeHitRatingPerHitChance)
	}

	if shaman.Talents.ThunderingStrikes > 0 {
		shaman.AddStat(stats.MeleeCrit, core.MeleeCritRatingPerCritChance*1*float64(shaman.Talents.ThunderingStrikes))
	}

	if shaman.Talents.DualWieldSpecialization > 0 {
		// TODO: Check that player is actually dual wielding
		shaman.AddStat(stats.MeleeHit, core.MeleeHitRatingPerHitChance*2*float64(shaman.Talents.DualWieldSpecialization))
	}

	if shaman.Talents.UnrelentingStorm > 0 {
		coeff := 0.02 * float64(shaman.Talents.UnrelentingStorm)
		shaman.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.MP5,
			Modifier: func(intellect float64, mp5 float64) float64 {
				return mp5 + intellect*coeff
			},
		})
	}

	if shaman.Talents.AncestralKnowledge > 0 {
		coeff := 0.01 * float64(shaman.Talents.AncestralKnowledge)
		shaman.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Mana,
			ModifiedStat: stats.Mana,
			Modifier: func(mana float64, _ float64) float64 {
				return mana + mana*coeff
			},
		})
	}

	if shaman.Talents.MentalQuickness > 0 {
		coeff := 0.1 * float64(shaman.Talents.MentalQuickness)
		shaman.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.AttackPower,
			ModifiedStat: stats.SpellPower,
			Modifier: func(attackPower float64, spellPower float64) float64 {
				return spellPower + attackPower*coeff
			},
		})
	}

	if shaman.Talents.NaturesBlessing > 0 {
		coeff := 0.1 * float64(shaman.Talents.NaturesBlessing)
		shaman.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.SpellPower,
			Modifier: func(intellect float64, spellPower float64) float64 {
				return spellPower + intellect*coeff
			},
		})
	}

	shaman.applyElementalDevastation()
	shaman.applyFlurry()
	shaman.applyShamanisticFocus()
	shaman.applyUnleashedRage()
	shaman.registerElementalMasteryCD()
	shaman.registerNaturesSwiftnessCD()
	shaman.registerShamanisticRageCD()
}

var ElementalDevastationTalentAuraID = core.NewAuraID()
var ElementalDevastationAuraID = core.NewAuraID()

func (shaman *Shaman) applyElementalDevastation() {
	if shaman.Talents.ElementalDevastation == 0 {
		return
	}

	critBonus := 3.0 * float64(shaman.Talents.ElementalDevastation) * core.SpellCritRatingPerCritChance
	const dur = time.Second * 10

	shaman.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: ElementalDevastationTalentAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellEffect.Crit {
					spellCast.Character.AddAuraWithTemporaryStats(sim, ElementalDevastationAuraID, core.ActionID{ItemID: 30160}, stats.MeleeCrit, critBonus, dur)
				}
			},
		}
	})
}

var ElementalMasteryAuraID = core.NewAuraID()
var ElementalMasteryCooldownID = core.NewCooldownID()

func (shaman *Shaman) registerElementalMasteryCD() {
	if !shaman.Talents.ElementalMastery {
		return
	}
	actionID := core.ActionID{SpellID: 16166}

	shaman.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: ElementalMasteryCooldownID,
		Cooldown:   time.Minute * 3,
		Type:       core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				character.Metrics.AddInstantCast(actionID)

				character.AddAura(sim, core.Aura{
					ID:       ElementalMasteryAuraID,
					ActionID: actionID,
					Expires:  core.NeverExpires,
					OnCast: func(sim *core.Simulation, cast *core.Cast) {
						cast.ManaCost = 0
						cast.BonusCritRating = 100.0 * core.SpellCritRatingPerCritChance
					},
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						// Remove the buff and put skill on CD
						character.SetCD(ElementalMasteryCooldownID, sim.CurrentTime+time.Minute*3)
						character.RemoveAura(sim, ElementalMasteryAuraID)
						character.UpdateMajorCooldowns()
					},
				})
			}
		},
	})
}

var NaturesSwiftnessAuraID = core.NewAuraID()
var NaturesSwiftnessCooldownID = core.NewCooldownID()

func (shaman *Shaman) registerNaturesSwiftnessCD() {
	if !shaman.Talents.NaturesSwiftness {
		return
	}
	actionID := core.ActionID{SpellID: 16188}

	shaman.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: NaturesSwiftnessCooldownID,
		Cooldown:   time.Minute * 3,
		Type:       core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Don't use NS unless we're casting a full-length lightning bolt, which is
			// the only spell shamans have with a cast longer than GCD.
			if character.HasTemporarySpellCastSpeedIncrease() {
				return false
			}
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				character.AddAura(sim, core.Aura{
					ID:       NaturesSwiftnessAuraID,
					ActionID: actionID,
					Expires:  core.NeverExpires,
					OnCast: func(sim *core.Simulation, cast *core.Cast) {
						if cast.ActionID.SpellID != SpellIDLB12 {
							return
						}

						cast.CastTime = 0
					},
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						if cast.ActionID.SpellID != SpellIDLB12 {
							return
						}

						// Remove the buff and put skill on CD
						character.SetCD(NaturesSwiftnessCooldownID, sim.CurrentTime+time.Minute*3)
						character.RemoveAura(sim, NaturesSwiftnessAuraID)
						character.UpdateMajorCooldowns()
						character.Metrics.AddInstantCast(actionID)
					},
				})
			}
		},
	})
}

var UnleasedRageTalentAuraID = core.NewAuraID()
var UnleasedRageProcAuraID = core.NewAuraID()

func (shaman *Shaman) applyUnleashedRage() {
	if shaman.Talents.UnleashedRage == 0 {
		return
	}
	level := shaman.Talents.UnleashedRage

	if len(shaman.unleashedRages) == 0 {
		shaman.unleashedRages = make([]core.Aura, 5) //pre-fill up to 5 auras
		for i := range shaman.unleashedRages {
			shaman.unleashedRages[i] = core.Aura{
				ID: UnleasedRageProcAuraID,
			}
		}
	}
	shaman.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		dur := time.Second * 10
		bonus := 0.02 * float64(level)
		return core.Aura{
			ID: UnleasedRageTalentAuraID,
			OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
				if result != core.MeleeHitTypeCrit {
					return
				}
				for i, player := range shaman.GetCharacter().Party.Players {
					char := player.GetCharacter()

					// Set new expire time.
					shaman.unleashedRages[i].Expires = sim.CurrentTime + dur

					if char.HasAura(UnleasedRageProcAuraID) {
						// Renew existing
						char.ReplaceAura(sim, shaman.unleashedRages[i])
						continue
					}

					// Update aura with new OnExpire
					ap := char.GetStat(stats.AttackPower) * bonus
					char.AddStat(stats.AttackPower, ap)
					shaman.unleashedRages[i].OnExpire = func(sim *core.Simulation) {
						char.AddStat(stats.AttackPower, -ap)
					}
					char.AddAura(sim, shaman.unleashedRages[i])
				}
			},
		}
	})
}

var ShamanisticFocusTalentAuraID = core.NewAuraID()
var ShamanisticFocusAuraID = core.NewAuraID()

func (shaman *Shaman) applyShamanisticFocus() {
	if !shaman.Talents.ShamanisticFocus {
		return
	}

	focusedAura := core.Aura{
		ID:       ShamanisticFocusAuraID,
		ActionID: core.ActionID{SpellID: 43338},
		Expires:  core.NeverExpires,
		OnCast: func(sim *core.Simulation, cast *core.Cast) {
			if !cast.IsSpellAction(SpellIDEarthShock) && !cast.IsSpellAction(SpellIDFlameShock) && !cast.IsSpellAction(SpellIDFrostShock) {
				return
			}

			cast.ManaCost -= cast.BaseManaCost * 0.6
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			if cast.IsSpellAction(SpellIDEarthShock) || cast.IsSpellAction(SpellIDFlameShock) || cast.IsSpellAction(SpellIDFrostShock) {
				cast.Character.RemoveAura(sim, ShamanisticFocusAuraID)
			}
		},
	}

	shaman.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: ShamanisticFocusTalentAuraID,
			OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
				if result != core.MeleeHitTypeCrit {
					return
				}
				shaman.AddAura(sim, focusedAura)
			},
		}
	})
}

var FlurryTalentAuraID = core.NewAuraID()
var FlurryProcAuraID = core.NewAuraID()

func (shaman *Shaman) applyFlurry() {
	if shaman.Talents.Flurry == 0 {
		return
	}

	shaman.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		flurryStacks := 0
		icdDur := time.Millisecond * 500
		var icd core.InternalCD

		bonus := 1.05 + 0.05*float64(shaman.Talents.Flurry)
		if ItemSetCataclysmRegalia.CharacterHasSetBonus(&shaman.Character, 4) {
			bonus += 1.05
		}
		inverseBonus := 1 / 1.3

		return core.Aura{
			ID: FlurryTalentAuraID,
			OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
				if result != core.MeleeHitTypeCrit {
					if ability == nil {
						// Remove a stack from auto attacks
						if flurryStacks > 0 && !icd.IsOnCD(sim) {
							icd = core.InternalCD(sim.CurrentTime + icdDur)
							flurryStacks--
							// Remove aura will reset attack speed
							if flurryStacks == 0 {
								shaman.RemoveAura(sim, FlurryProcAuraID)
							}
						}
					}
					return
				}
				if flurryStacks == 0 {
					shaman.MultiplyMeleeSpeed(sim, bonus)
					shaman.AddAura(sim, core.Aura{
						ID:       FlurryProcAuraID,
						ActionID: core.ActionID{SpellID: 16280},
						Expires:  core.NeverExpires,
						OnExpire: func(sim *core.Simulation) {
							shaman.MultiplyMeleeSpeed(sim, inverseBonus)
						},
					})
				}
				icd = 0
				flurryStacks = 3
			},
		}
	})
}
