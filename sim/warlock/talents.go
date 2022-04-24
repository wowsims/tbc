package warlock

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
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
		bonus := (0.01) * float64(warlock.Talents.FelIntellect)
		// Adding a second 3% bonus int->mana dependency
		warlock.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.Mana,
			Modifier: func(intellect float64, mana float64) float64 {
				return mana + intellect*(15*bonus)
			},
		})
	}

	warlock.PseudoStats.BonusCritRating += float64(warlock.Talents.DemonicTactics) * 1 * core.SpellCritRatingPerCritChance

	//  TODO: fel stamina increases max health (might be useful for warlock tanking sim)

	if !warlock.Options.SacrificeSummon && warlock.Options.Summon != proto.Warlock_Options_NoSummon {
		warlock.AddStats(stats.Stats{
			stats.MeleeCrit: float64(warlock.Talents.DemonicTactics) * 5 * core.MeleeCritRatingPerCritChance,
			stats.SpellCrit: float64(warlock.Talents.DemonicTactics) * 5 * core.SpellCritRatingPerCritChance,
		})

		if warlock.Talents.MasterDemonologist > 0 {
			switch warlock.Options.Summon {
			case proto.Warlock_Options_Imp:
				warlock.PseudoStats.ThreatMultiplier *= 0.96 * float64(warlock.Talents.MasterDemonologist)
			case proto.Warlock_Options_Succubus:
				warlock.PseudoStats.DamageDealtMultiplier *= 1.0 + 0.02*float64(warlock.Talents.MasterDemonologist)
			case proto.Warlock_Options_Felgaurd:
				warlock.PseudoStats.DamageDealtMultiplier *= 1.0 + 0.01*float64(warlock.Talents.MasterDemonologist)
				// 		Felguard - Increases all damage caused by 1% and all resistances by .1 per level.
				// 		Voidwalker - Reduces physical damage taken by 2%.
				// 		Felhunter - Increases all resistances by .2 per level.
			}
		}

		// Create the pet
		warlock.NewWarlockPet()

		// Extract stats for demonic knowledge
		petChar := warlock.Pets[0].GetCharacter()
		bonus := (petChar.GetStat(stats.Stamina) + petChar.GetStat(stats.Intellect)) * (0.04 * float64(warlock.Talents.DemonicKnowledge))
		warlock.AddStat(stats.SpellPower, bonus)
	}

	// demonic tactics, applies even without pet out

	warlock.setupNightfall()
}

func (warlock *Warlock) setupNightfall() {
	if warlock.Talents.Nightfall == 0 {
		return
	}

	warlock.NightfallProcAura = warlock.RegisterAura(core.Aura{
		Label:    "Nightfall Shadow Trance",
		ActionID: core.ActionID{SpellID: 17941},
		Duration: core.NeverExpires,
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			// Check for an instant cast shadowbolt to disable aura
			if spell != warlock.Shadowbolt || spell.CurCast.CastTime != 0 {
				return
			}
			aura.Deactivate(sim)
		},
	})

	warlock.RegisterAura(core.Aura{
		Label: "Nightfall",
		// ActionID: core.ActionID{SpellID: 18095},
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnPeriodicDamage: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spell != warlock.Corruption { // TODO: also works on drain life...
				return
			}
			if sim.RandomFloat("nightfall") > 0.04 {
				return
			}
			warlock.NightfallProcAura.Activate(sim)
		},
	})
}

func (warlock *Warlock) applyNightfall(cast *core.Cast) {
	if warlock.NightfallProcAura.IsActive() {
		cast.CastTime = 0
	}
}
