package warlock

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (warlock *Warlock) ApplyTalents() {
	// demonic embrace
	if warlock.Talents.DemonicEmbrace > 0 {
		bonus := 1 + (0.03)*float64(warlock.Talents.DemonicEmbrace)
		negative := 1 - (0.01)*float64(warlock.Talents.DemonicEmbrace)
		warlock.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Stamina,
			ModifiedStat: stats.Stamina,
			Modifier: func(in float64, _ float64) float64 {
				return in * bonus
			},
		})
		warlock.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Spirit,
			ModifiedStat: stats.Spirit,
			Modifier: func(in float64, _ float64) float64 {
				return in * negative
			},
		})
	}
	// fel intellect
	if warlock.Talents.FelIntellect > 0 {
		bonus := 1 + (0.01)*float64(warlock.Talents.FelIntellect)
		warlock.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Mana,
			ModifiedStat: stats.Mana,
			Modifier: func(in float64, out float64) float64 {
				return in * bonus
			},
		})
	}
	//  TODO: fel stamina increases max health (might be useful for warlock tanking sim)

	// if pet is out:
	// master demonologist -
	// 		Grants both the Warlock and the summoned demon an effect as long as that demon is active.
	// 		Imp - Reduces threat caused by 4%.
	// 		Voidwalker - Reduces physical damage taken by 2%.
	// 		Succubus/Incubus - Increases all damage caused by 2%.
	// 		Felhunter - Increases all resistances by .2 per level.
	// 		Felguard - Increases all damage caused by 1% and all resistances by .1 per level.
	// demonic knowledge -
	// 		Increases your spell damage by an amount equal to 4%(per point of talent) of
	//		the total of your active demon's Stamina plus Intellect.

	// demonic tactics, applies even without pet out
	warlock.PseudoStats.BonusCritRating += float64(warlock.Talents.DemonicTactics) * 1 * core.SpellCritRatingPerCritChance
}
