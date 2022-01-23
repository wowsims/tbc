package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (druid *Druid) applyTalents() {
	druid.registerNaturesSwiftnessCD()

	druid.AddStat(stats.SpellHit, float64(druid.Talents.BalanceOfPower)*2*core.SpellHitRatingPerHitChance)

	if druid.Talents.LunarGuidance > 0 {
		bonus := (0.25 / 3) * float64(druid.Talents.LunarGuidance)
		druid.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.SpellPower,
			Modifier: func(intellect float64, spellPower float64) float64 {
				return spellPower + intellect*bonus
			},
		})
	}

	if druid.Talents.Dreamstate > 0 {
		bonus := (0.1 / 3) * float64(druid.Talents.Dreamstate)
		druid.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.MP5,
			Modifier: func(intellect float64, mp5 float64) float64 {
				return mp5 + intellect*bonus
			},
		})
	}

	if druid.Talents.Intensity > 0 {
		druid.PseudoStats.SpiritRegenRateCasting = float64(druid.Talents.Intensity) * 0.1
	}

	if druid.Talents.Subtlety > 0 {
		druid.PseudoStats.ThreatMultiplier *= 1 - 0.04*float64(druid.Talents.Subtlety)
	}

	if druid.Talents.HeartOfTheWild > 0 {
		bonus := 0.04 * float64(druid.Talents.HeartOfTheWild)
		druid.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.Intellect,
			Modifier: func(intellect float64, _ float64) float64 {
				return intellect + intellect*bonus
			},
		})
	}

	if druid.Talents.SurvivalOfTheFittest > 0 {
		bonus := 0.01 * float64(druid.Talents.SurvivalOfTheFittest)
		druid.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Stamina,
			ModifiedStat: stats.Stamina,
			Modifier: func(stat float64, _ float64) float64 {
				return stat + stat*bonus
			},
		})
		druid.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Strength,
			ModifiedStat: stats.Strength,
			Modifier: func(stat float64, _ float64) float64 {
				return stat + stat*bonus
			},
		})
		druid.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Agility,
			ModifiedStat: stats.Agility,
			Modifier: func(stat float64, _ float64) float64 {
				return stat + stat*bonus
			},
		})
		druid.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.Intellect,
			Modifier: func(stat float64, _ float64) float64 {
				return stat + stat*bonus
			},
		})
		druid.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.Intellect,
			Modifier: func(stat float64, _ float64) float64 {
				return stat + stat*bonus
			},
		})
		druid.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Spirit,
			ModifiedStat: stats.Spirit,
			Modifier: func(stat float64, _ float64) float64 {
				return stat + stat*bonus
			},
		})
	}

	if druid.Talents.LivingSpirit > 0 {
		bonus := 0.05 * float64(druid.Talents.LivingSpirit)
		druid.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Spirit,
			ModifiedStat: stats.Spirit,
			Modifier: func(spirit float64, _ float64) float64 {
				return spirit + spirit*bonus
			},
		})
	}

	if druid.Talents.NaturalPerfection > 0 {
		druid.AddStat(stats.SpellCrit, float64(druid.Talents.NaturalPerfection)*1*core.SpellCritRatingPerCritChance)
	}
}

var NaturesSwiftnessAuraID = core.NewAuraID()
var NaturesSwiftnessCooldownID = core.NewCooldownID()

func (druid *Druid) registerNaturesSwiftnessCD() {
	if !druid.Talents.NaturesSwiftness {
		return
	}
	actionID := core.ActionID{SpellID: 17116}

	druid.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: NaturesSwiftnessCooldownID,
		Cooldown:   time.Minute * 3,
		Type:       core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Don't use NS unless we're casting a full-length starfire or wrath.
			if character.HasTemporarySpellCastSpeedIncrease() {
				return false
			}
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				character.AddAura(sim, core.Aura{
					ID:       NaturesSwiftnessAuraID,
					ActionID: actionID,
					Expires:  core.NeverExpires,
					OnCast: func(sim *core.Simulation, cast *core.Cast) {
						if cast.ActionID.SpellID != SpellIDWrath && cast.ActionID.SpellID != SpellIDSF8 && cast.ActionID.SpellID != SpellIDSF6 {
							return
						}

						cast.CastTime = 0
					},
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						if cast.ActionID.SpellID != SpellIDWrath && cast.ActionID.SpellID != SpellIDSF8 && cast.ActionID.SpellID != SpellIDSF6 {
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
