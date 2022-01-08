package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (mage *Mage) applyTalents() {
	mage.applyArcaneConcentration()
	mage.applyIgnite()
	mage.applyMasterOfElements()
	mage.applyWintersChill()
	mage.applyMoltenFury()
	mage.registerPresenceOfMindCD()
	mage.registerCombustionCD()
	mage.registerIcyVeinsCD()
	mage.registerColdSnapCD()
	mage.registerSummonWaterElementalCD()

	if mage.Talents.ArcaneMeditation > 0 {
		mage.PseudoStats.SpiritRegenRateCasting += float64(mage.Talents.ArcaneMeditation) * 0.1
	}

	if mage.Talents.ArcaneMind > 0 {
		mage.Character.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.Intellect,
			Modifier: func(intellect float64, _ float64) float64 {
				return intellect * (1.0 + 0.03*float64(mage.Talents.ArcaneMind))
			},
		})
	}

	if mage.Talents.MindMastery > 0 {
		mage.Character.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.SpellPower,
			Modifier: func(intellect float64, spellPower float64) float64 {
				return spellPower + intellect*0.05*float64(mage.Talents.MindMastery)
			},
		})
	}

	if mage.Talents.ArcaneInstability > 0 {
		mage.AddStat(stats.SpellCrit, float64(mage.Talents.ArcaneInstability)*1*core.SpellCritRatingPerCritChance)
		mage.spellDamageMultiplier += float64(mage.Talents.ArcaneInstability) * 0.01
	}

	if mage.Talents.PlayingWithFire > 0 {
		mage.spellDamageMultiplier += float64(mage.Talents.PlayingWithFire) * 0.01
	}
}

var ArcaneConcentrationAuraID = core.NewAuraID()
var ClearcastingAuraID = core.NewAuraID()

func (mage *Mage) applyArcaneConcentration() {
	if mage.Talents.ArcaneConcentration == 0 {
		return
	}

	mage.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		procChance := 0.02 * float64(mage.Talents.ArcaneConcentration)
		bonusCrit := float64(mage.Talents.ArcanePotency) * 10 * core.SpellCritRatingPerCritChance
		icd := core.NewICD()
		const icdDur = time.Second * 1

		return core.Aura{
			ID: ArcaneConcentrationAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellCast.IsSpellAction(SpellIDArcaneMissiles) {
					// AM shouldn't proc clearcasting.
					return
				}

				// TODO: This should only get 1 roll for each aoe cast.
				if icd.IsOnCD(sim) || sim.RandomFloat("Arcane Concentration") > procChance {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)

				mage.AddAura(sim, core.Aura{
					ID:       ClearcastingAuraID,
					ActionID: core.ActionID{SpellID: 12536},
					OnCast: func(sim *core.Simulation, cast *core.Cast) {
						cast.ManaCost = 0
						cast.BonusCritRating += bonusCrit
					},
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						mage.RemoveAura(sim, ClearcastingAuraID)
					},
				})
			},
		}
	})
}

var PresenceOfMindCooldownID = core.NewCooldownID()

func (mage *Mage) registerPresenceOfMindCD() {
	if !mage.Talents.PresenceOfMind {
		return
	}

	cooldown := time.Minute * 3
	if ItemSetAldorRegalia.CharacterHasSetBonus(&mage.Character, 4) {
		cooldown -= time.Second * 24
	}

	var spell *core.SimpleSpell
	actionID := core.ActionID{SpellID: 12043}

	mage.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: PresenceOfMindCooldownID,
		Cooldown:   cooldown,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			if character.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
				return false
			}

			target := sim.GetPrimaryTarget()
			if mage.Talents.Pyroblast {
				spell = mage.NewPyroblast(sim, target)
			} else if mage.RotationType == proto.Mage_Rotation_Fire {
				spell = mage.NewFireball(sim, target)
			} else {
				spell = mage.NewFrostbolt(sim, target)
			}

			if character.CurrentMana() < spell.ManaCost {
				spell.Cancel(sim)
				return false
			}

			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				spell.CastTime = 0
				spell.Cast(sim)
				character.Metrics.AddInstantCast(actionID)
				character.SetCD(PresenceOfMindCooldownID, sim.CurrentTime+cooldown)
			}
		},
	})
}

var ArcanePowerAuraID = core.NewAuraID()
var ArcanePowerCooldownID = core.NewCooldownID()

func (mage *Mage) registerArcanePowerCD() {
	if !mage.Talents.ArcanePower {
		return
	}
	actionID := core.ActionID{SpellID: 12042}

	mage.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: ArcanePowerCooldownID,
		Cooldown:   time.Minute * 3,
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
					ID:       ArcanePowerAuraID,
					ActionID: actionID,
					Expires:  sim.CurrentTime + time.Second*15,
					OnCast: func(sim *core.Simulation, cast *core.Cast) {
						cast.ManaCost *= 1.3
					},
					OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						spellEffect.DamageMultiplier *= 1.3
					},
					OnBeforePeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage *float64) {
						*tickDamage *= 1.3
					},
				})
				mage.SetCD(ArcanePowerCooldownID, sim.CurrentTime+time.Minute*3)
			}
		},
	})
}

var MasterOfElementsAuraID = core.NewAuraID()

func (mage *Mage) applyMasterOfElements() {
	if mage.Talents.MasterOfElements == 0 {
		return
	}

	refundCoeff := 0.1 * float64(mage.Talents.MasterOfElements)

	mage.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: MasterOfElementsAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellEffect.Crit {
					mage.AddMana(sim, spellCast.BaseManaCost*refundCoeff, core.ActionID{SpellID: 29076}, false)
				}
			},
		}
	})
}

var CombustionAuraID = core.NewAuraID()
var CombustionCooldownID = core.NewCooldownID()

func (mage *Mage) registerCombustionCD() {
	if !mage.Talents.Combustion {
		return
	}
	actionID := core.ActionID{SpellID: 11129}

	mage.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: CombustionCooldownID,
		Cooldown:   time.Minute * 3,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			if character.HasAura(CombustionAuraID) {
				return false
			}
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				character.Metrics.AddInstantCast(actionID)

				numHits := 0
				numCrits := 0

				character.AddAura(sim, core.Aura{
					ID:       CombustionAuraID,
					ActionID: actionID,
					Expires:  core.NeverExpires,
					OnCast: func(sim *core.Simulation, cast *core.Cast) {
						cast.BonusCritRating += float64(numHits) * 10 * core.SpellCritRatingPerCritChance
					},
					OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						numHits++
						if spellEffect.Crit {
							numCrits++
							if numCrits == 3 {
								character.RemoveAuraOnNextAdvance(sim, CombustionAuraID)
								character.SetCD(CombustionCooldownID, sim.CurrentTime+time.Minute*3)
								character.UpdateMajorCooldowns()
							}
						}
					},
				})
			}
		},
	})
}

var IcyVeinsAuraID = core.NewAuraID()
var IcyVeinsCooldownID = core.NewCooldownID()

func (mage *Mage) registerIcyVeinsCD() {
	if !mage.Talents.IcyVeins {
		return
	}

	manaCost := 0.0
	actionID := core.ActionID{SpellID: 12472}

	mage.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: IcyVeinsCooldownID,
		Cooldown:   time.Minute * 3,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Need to check for icy veins already active in case Cold Snap is used right after.
			if character.HasAura(IcyVeinsAuraID) {
				return false
			}

			if character.CurrentMana() < manaCost {
				return false
			}

			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			const bonus = 1.2
			const inverseBonus = 1 / bonus
			manaCost = mage.BaseMana() * 0.03

			return func(sim *core.Simulation, character *core.Character) {
				character.PseudoStats.CastSpeedMultiplier *= bonus

				character.AddAura(sim, core.Aura{
					ID:       IcyVeinsAuraID,
					ActionID: actionID,
					Expires:  sim.CurrentTime + time.Second*20,
					OnExpire: func(sim *core.Simulation) {
						character.PseudoStats.CastSpeedMultiplier *= inverseBonus
					},
				})
				character.SpendMana(sim, manaCost, actionID)
				character.Metrics.AddInstantCast(actionID)
				character.SetCD(IcyVeinsCooldownID, sim.CurrentTime+time.Minute*3)
			}
		},
	})
}

var ColdSnapCooldownID = core.NewCooldownID()

func (mage *Mage) registerColdSnapCD() {
	if !mage.Talents.ColdSnap {
		return
	}

	cooldown := time.Duration(float64(time.Minute*8) * (1.0 - float64(mage.Talents.IceFloes)*0.1))
	actionID := core.ActionID{SpellID: 11958}

	mage.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: ColdSnapCooldownID,
		Cooldown:   cooldown,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Don't use if there are no cooldowns to reset.
			if !character.IsOnCD(IcyVeinsCooldownID, sim.CurrentTime) && !character.IsOnCD(SummonWaterElementalCooldownID, sim.CurrentTime) {
				return false
			}
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Ideally wait for both water ele and icy veins so we can reset both.
			if mage.Talents.IcyVeins && !character.IsOnCD(IcyVeinsCooldownID, sim.CurrentTime) {
				return false
			}
			if mage.Talents.SummonWaterElemental && !character.IsOnCD(SummonWaterElementalCooldownID, sim.CurrentTime) {
				return false
			}

			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				character.SetCD(IcyVeinsCooldownID, 0)
				character.SetCD(SummonWaterElementalCooldownID, 0)

				character.Metrics.AddInstantCast(actionID)
				character.SetCD(ColdSnapCooldownID, sim.CurrentTime+cooldown)
			}
		},
	})
}

var MoltenFuryAuraID = core.NewAuraID()

func (mage *Mage) applyMoltenFury() {
	if mage.Talents.MoltenFury == 0 {
		return
	}

	multiplier := 1.0 + 0.1*float64(mage.Talents.MoltenFury)

	mage.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: MoltenFuryAuraID,
			OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if sim.IsExecutePhase() {
					spellEffect.DamageMultiplier *= multiplier
				}
			},
			OnBeforePeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage *float64) {
				if sim.IsExecutePhase() {
					*tickDamage *= multiplier
				}
			},
		}
	})
}
