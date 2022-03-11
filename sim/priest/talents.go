package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (priest *Priest) applyTalents() {
	if priest.Talents.Meditation > 0 {
		priest.PseudoStats.SpiritRegenRateCasting = float64(priest.Talents.Meditation) * 0.1
	}

	if priest.Talents.SpiritualGuidance > 0 {
		bonus := (0.25 / 5) * float64(priest.Talents.SpiritualGuidance)
		priest.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Spirit,
			ModifiedStat: stats.SpellPower,
			Modifier: func(spirit float64, spellPower float64) float64 {
				return spellPower + spirit*bonus
			},
		})
	}

	if priest.Talents.MentalStrength > 0 {
		coeff := 0.02 * float64(priest.Talents.MentalStrength)
		priest.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Mana,
			ModifiedStat: stats.Mana,
			Modifier: func(mana float64, _ float64) float64 {
				return mana + mana*coeff
			},
		})
	}

	if priest.Talents.ForceOfWill > 0 {
		coeff := 0.01 * float64(priest.Talents.ForceOfWill)
		priest.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.SpellPower,
			ModifiedStat: stats.SpellPower,
			Modifier: func(spellPower float64, _ float64) float64 {
				return spellPower + spellPower*coeff
			},
		})
	}
	
	if priest.Talents.Enlightenment > 0 {
		coeff := 0.01 * float64(priest.Talents.Enlightenment)
		priest.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.Intellect,
			Modifier: func(intellect float64, _ float64) float64 {
				return intellect + intellect*coeff
			},
		})
	}
	
	if priest.Talents.Enlightenment > 0 {
		coeff := 0.01 * float64(priest.Talents.Enlightenment)
		priest.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Stamina,
			ModifiedStat: stats.Stamina,
			Modifier: func(intellect float64, _ float64) float64 {
				return intellect + intellect*coeff
			},
		})
	}
	
	if priest.Talents.Enlightenment > 0 {
		coeff := 0.01 * float64(priest.Talents.Enlightenment)
		priest.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Spirit,
			ModifiedStat: stats.Spirit,
			Modifier: func(intellect float64, _ float64) float64 {
				return intellect + intellect*coeff
			},
		})
	}
}

func (priest *Priest) applyOnHitTalents(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {

	roll := sim.RandomFloat("SurgeOfLight")
	if priest.Talents.SurgeOfLight == 2 && spellEffect.Outcome.Matches(core.OutcomeCrit) && roll > 0.5 {
		priest.SurgeOfLight = true
	}
}

func (priest *Priest) applySurgeOfLight(spellCast *core.SpellCast) {
	if priest.SurgeOfLight {
		spellCast.CastTime = 0
		spellCast.Cost.Value = 0
		// This applies on cast complete, removing the effect.
		//  if it crits, during 'onspellhit' then it will be reapplied (see func above)
		spellCast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
			priest.SurgeOfLight = false
		}
	}
}

func (priest *Priest) applyTalentsToHolySpell(cast *core.Cast, effect *core.SpellHitEffect) {
	effect.ThreatMultiplier *= 1 - 0.04*float64(priest.Talents.SilentResolve)
	if cast.ActionID.SpellID == SpellIDSmite || cast.ActionID.SpellID == SpellIDHolyFire {
		effect.BonusSpellCritRating += float64(priest.Talents.HolySpecialization) * 1 * core.SpellCritRatingPerCritChance
	}

	effect.BonusSpellCritRating += float64(priest.Talents.ForceOfWill) * 1 * core.SpellCritRatingPerCritChance
}

func (priest *Priest) applyTalentsToShadowSpell(cast *core.Cast, effect *core.SpellHitEffect) {
	effect.ThreatMultiplier *= 1 - 0.08*float64(priest.Talents.ShadowAffinity)
	if cast.ActionID.SpellID == SpellIDShadowWordDeath || cast.ActionID.SpellID == SpellIDMindBlast {
		effect.BonusSpellCritRating += float64(priest.Talents.ShadowPower) * 3 * core.SpellCritRatingPerCritChance
	}
	if cast.ActionID.SpellID == SpellIDMindFlay || cast.ActionID.SpellID == SpellIDMindBlast {
		cast.Cost.Value -= cast.BaseCost.Value * float64(priest.Talents.FocusedMind) * 0.05
	}
	if cast.SpellSchool == core.SpellSchoolShadow {
		effect.StaticDamageMultiplier *= 1 + float64(priest.Talents.Darkness)*0.02

		if priest.Talents.Shadowform {
			effect.StaticDamageMultiplier *= 1.15
		}

		// shadow focus gives 2% hit per level
		effect.BonusSpellHitRating += float64(priest.Talents.ShadowFocus) * 2 * core.SpellHitRatingPerHitChance

		// TODO should add more instant cast spells here
		if cast.ActionID.SpellID == SpellIDShadowWordPain {
			cast.Cost.Value -= cast.BaseCost.Value * float64(priest.Talents.MentalAgility) * 0.02
		}

		effect.BonusSpellCritRating += float64(priest.Talents.ForceOfWill) * 1 * core.SpellCritRatingPerCritChance
	}
}

var InnerFocusAuraID = core.NewAuraID()
var InnerFocusCooldownID = core.NewCooldownID()

func ApplyInnerFocus(sim *core.Simulation, priest *Priest) bool {
	actionID := core.ActionID{SpellID: 14751}
	priest.Metrics.AddInstantCast(actionID)
	priest.Character.AddAura(sim, core.Aura{
		ID:       InnerFocusAuraID,
		ActionID: actionID,
		Expires:  core.NeverExpires,
		OnCast: func(sim *core.Simulation, cast *core.Cast) {
			cast.Cost.Value = 0
		},
		OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
			spellEffect.BonusSpellCritRating += 25 * core.SpellCritRatingPerCritChance
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			// Remove the buff and put skill on CD
			priest.SetCD(InnerFocusCooldownID, sim.CurrentTime+time.Minute*3)
			priest.RemoveAura(sim, InnerFocusAuraID)
		},
	})
	return true
}
