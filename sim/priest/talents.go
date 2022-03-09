package priest

import (
	"time"
	//"math/rand"

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
			SourceStat:	stats.SpellPower,
			ModifiedStat:	stats.SpellPower,
			Modifier: func(spellPower float64, _ float64) float64 {
				return spellPower + spellPower*coeff
			},
		})
	}
}

func (priest *Priest) applyOnHitTalents(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
	if spellEffect.Outcome.Matches(core.OutcomeCrit) {
		priest.SurgeOfLight = true
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
