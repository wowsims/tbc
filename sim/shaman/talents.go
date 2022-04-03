package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (shaman *Shaman) ApplyTalents() {
	if shaman.Talents.NaturesGuidance > 0 {
		shaman.AddStat(stats.SpellHit, float64(shaman.Talents.NaturesGuidance)*1*core.SpellHitRatingPerHitChance)
		shaman.AddStat(stats.MeleeHit, float64(shaman.Talents.NaturesGuidance)*1*core.MeleeHitRatingPerHitChance)
	}

	if shaman.Talents.ThunderingStrikes > 0 {
		shaman.AddStat(stats.MeleeCrit, core.MeleeCritRatingPerCritChance*1*float64(shaman.Talents.ThunderingStrikes))
	}

	shaman.AddStat(stats.Dodge, core.DodgeRatingPerDodgeChance*1*float64(shaman.Talents.Anticipation))
	shaman.AddStat(stats.Block, core.BlockRatingPerBlockChance*1*float64(shaman.Talents.ShieldSpecialization))
	shaman.AddStat(stats.Armor, shaman.Equip.Stats()[stats.Armor]*0.02*float64(shaman.Talents.Toughness))
	shaman.PseudoStats.PhysicalDamageDealtMultiplier *= 1 + 0.02*float64(shaman.Talents.WeaponMastery)

	if shaman.Talents.ShieldSpecialization > 0 {
		bonus := 1 + 0.05*float64(shaman.Talents.ShieldSpecialization)
		shaman.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.BlockValue,
			ModifiedStat: stats.BlockValue,
			Modifier: func(bv float64, _ float64) float64 {
				return bv * bonus
			},
		})
	}

	if shaman.Talents.DualWieldSpecialization > 0 && shaman.HasOHWeapon() {
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

	if shaman.Talents.SpiritWeapons {
		shaman.AutoAttacks.MHAuto.Template.Effect.ThreatMultiplier *= 0.7
		shaman.AutoAttacks.OHAuto.Template.Effect.ThreatMultiplier *= 0.7
	}

	shaman.applyElementalDevastation()
	shaman.applyFlurry()
	shaman.applyShamanisticFocus()
	shaman.applyUnleashedRage()
	shaman.registerElementalMasteryCD()
	shaman.registerNaturesSwiftnessCD()
	shaman.registerShamanisticRageCD()
}

func (shaman *Shaman) applyElementalDevastation() {
	if shaman.Talents.ElementalDevastation == 0 {
		return
	}

	shaman.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		critBonus := 3.0 * float64(shaman.Talents.ElementalDevastation) * core.SpellCritRatingPerCritChance
		procAura := shaman.NewTemporaryStatsAura("Elemental Devastation Proc", core.ActionID{SpellID: 30160}, stats.Stats{stats.MeleeCrit: critBonus}, time.Second*10)
		return shaman.GetOrRegisterAura(&core.Aura{
			Label: "Elemental Devastation",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}
				if spellEffect.IsPhantom {
					return
				}
				if !spellEffect.Outcome.Matches(core.OutcomeCrit) {
					return
				}
				procAura.Activate(sim)
			},
		})
	})
}

var ElementalMasteryCooldownID = core.NewCooldownID()

func (shaman *Shaman) registerElementalMasteryCD() {
	if !shaman.Talents.ElementalMastery {
		return
	}
	actionID := core.ActionID{SpellID: 16166}

	shaman.ElementalMasteryAura = shaman.RegisterAura(&core.Aura{
		Label:    "Elemental Mastery",
		ActionID: actionID,
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			shaman.AddStat(stats.SpellCrit, 100*core.SpellCritRatingPerCritChance)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			shaman.AddStat(stats.SpellCrit, -100*core.SpellCritRatingPerCritChance)
		},
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spell.SpellExtras.Matches(SpellFlagShock | SpellFlagElectric) {
				return
			}
			// Remove the buff and put skill on CD
			aura.Deactivate(sim)
			shaman.SetCD(ElementalMasteryCooldownID, sim.CurrentTime+time.Minute*3)
			shaman.UpdateMajorCooldowns()
		},
	})

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
				shaman.ElementalMasteryAura.Activate(sim)
				shaman.ElementalMasteryAura.Prioritize()
			}
		},
	})
}

var NaturesSwiftnessCooldownID = core.NewCooldownID()

func (shaman *Shaman) registerNaturesSwiftnessCD() {
	if !shaman.Talents.NaturesSwiftness {
		return
	}
	actionID := core.ActionID{SpellID: 16188}

	shaman.NaturesSwiftnessAura = shaman.RegisterAura(&core.Aura{
		Label:    "Natures Swiftness",
		ActionID: actionID,
		Duration: core.NeverExpires,
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, cast *core.Cast) {
			if cast.ActionID.SpellID != SpellIDLB12 {
				return
			}

			// Remove the buff and put skill on CD
			aura.Deactivate(sim)
			shaman.SetCD(NaturesSwiftnessCooldownID, sim.CurrentTime+time.Minute*3)
			shaman.UpdateMajorCooldowns()
		},
	})

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
				shaman.NaturesSwiftnessAura.Activate(sim)
				character.Metrics.AddInstantCast(actionID)
			}
		},
	})
}

func (shaman *Shaman) applyUnleashedRage() {
	if shaman.Talents.UnleashedRage == 0 {
		return
	}
	level := shaman.Talents.UnleashedRage

	shaman.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		bonusCoeff := 0.02 * float64(level)
		currentAPBonuses := make([]float64, len(shaman.Party.PlayersAndPets))

		urAuras := make([]*core.Aura, len(shaman.Party.PlayersAndPets))
		for i, playerOrPet := range shaman.Party.PlayersAndPets {
			char := playerOrPet.GetCharacter()
			idx := i
			urAuras[i] = char.GetOrRegisterAura(&core.Aura{
				Label:    "Unleahed Rage Proc",
				ActionID: core.ActionID{SpellID: 30811},
				Duration: time.Second * 10,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					buffs := char.ApplyStatDependencies(stats.Stats{stats.AttackPower: currentAPBonuses[idx]})
					char.AddStats(buffs)
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					buffs := char.ApplyStatDependencies(stats.Stats{stats.AttackPower: currentAPBonuses[idx]})
					unbuffs := buffs.Multiply(-1)
					char.AddStats(unbuffs)
				},
			})
		}

		return shaman.GetOrRegisterAura(&core.Aura{
			Label: "Unleashed Rage",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				// proc mask = 20 (melee auto & special)
				if !spellEffect.Outcome.Matches(core.OutcomeCrit) || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}

				for i, playerOrPet := range shaman.Party.PlayersAndPets {
					char := playerOrPet.GetCharacter()
					prevBonus := currentAPBonuses[i]
					newBonus := (char.GetStat(stats.AttackPower) - prevBonus) * bonusCoeff

					if prevBonus != newBonus {
						urAuras[i].Deactivate(sim)
						currentAPBonuses[i] = newBonus
						urAuras[i].Activate(sim)
					} else if newBonus != 0 {
						// If the bonus is the same, we can just refresh.
						urAuras[i].Refresh(sim)
					}
				}
			},
		})
	})
}

func (shaman *Shaman) applyShamanisticFocus() {
	if !shaman.Talents.ShamanisticFocus {
		return
	}

	shaman.ShamanisticFocusAura = shaman.RegisterAura(&core.Aura{
		Label:    "Shamanistic Focus Proc",
		ActionID: core.ActionID{SpellID: 43338},
		Duration: core.NeverExpires,
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, cast *core.Cast) {
			if cast.SpellExtras.Matches(SpellFlagShock) {
				aura.Deactivate(sim)
			}
		},
	})

	shaman.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		return shaman.GetOrRegisterAura(&core.Aura{
			Label: "Shamanistic Focus",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if !spellEffect.Outcome.Matches(core.OutcomeCrit) {
					return
				}
				shaman.ShamanisticFocusAura.Activate(sim)
			},
		})
	})
}

func (shaman *Shaman) applyFlurry() {
	if shaman.Talents.Flurry == 0 {
		return
	}

	bonus := 1.05 + 0.05*float64(shaman.Talents.Flurry)
	if ItemSetCataclysmHarness.CharacterHasSetBonus(&shaman.Character, 4) {
		bonus += 0.05
	}
	inverseBonus := 1 / bonus

	procAura := shaman.RegisterAura(&core.Aura{
		Label:     "Flurry Proc",
		ActionID:  core.ActionID{SpellID: 16280},
		Duration:  core.NeverExpires,
		MaxStacks: 3,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			shaman.MultiplyMeleeSpeed(sim, bonus)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			shaman.MultiplyMeleeSpeed(sim, inverseBonus)
		},
	})

	shaman.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		var icd core.InternalCD
		icdDur := time.Millisecond * 500

		return shaman.GetOrRegisterAura(&core.Aura{
			Label: "Flurry",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}

				if spellEffect.Outcome.Matches(core.OutcomeCrit) {
					procAura.Activate(sim)
					procAura.SetStacks(sim, 3)
					icd = 0 // the "charge protection" ICD isn't up yet
					return
				}

				// Remove a stack.
				if procAura.IsActive() && spellEffect.ProcMask.Matches(core.ProcMaskMeleeWhiteHit) && !icd.IsOnCD(sim) {
					icd = core.InternalCD(sim.CurrentTime + icdDur)
					procAura.RemoveStack(sim)
				}
			},
		})
	})
}
