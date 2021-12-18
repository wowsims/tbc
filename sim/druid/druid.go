package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Druid struct {
	core.Character
	SelfBuffs
	Talents proto.DruidTalents

	innervateCD  time.Duration
	NaturesGrace bool // when true next spellcast is 0.5s faster
	RebirthUsed  bool

	// cached cast stuff
	starfireSpell         core.SimpleSpell
	starfire8CastTemplate core.SimpleSpellTemplate
	starfire6CastTemplate core.SimpleSpellTemplate

	MoonfireSpell        core.SimpleSpell
	moonfireCastTemplate core.SimpleSpellTemplate

	wrathSpell        core.SimpleSpell
	wrathCastTemplate core.SimpleSpellTemplate

	InsectSwarmSpell        core.SimpleSpell
	insectSwarmCastTemplate core.SimpleSpellTemplate

	FaerieFireSpell        core.SimpleSpell
	faerieFireCastTemplate core.SimpleSpellTemplate

	malorne4p bool // cached since we need to check on every innervate

	// Used for accounting for bonus mana expected from future innervates.
	RemainingInnervateUsages int
	ExpectedManaPerInnervate float64
}

type SelfBuffs struct {
	Omen      bool
	Innervate bool
}

func (druid *Druid) GetCharacter() *core.Character {
	return &druid.Character
}

func (druid *Druid) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
	raidBuffs.GiftOfTheWild = core.MaxTristate(raidBuffs.GiftOfTheWild, proto.TristateEffect_TristateEffectRegular)
	if druid.Talents.ImprovedMarkOfTheWild == 5 { // probably could work on actually calculating the fraction effect later if we care.
		raidBuffs.GiftOfTheWild = proto.TristateEffect_TristateEffectImproved
	}
}

const ravenGoddessItemID = 32387

func (druid *Druid) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
	if druid.Talents.MoonkinForm { // assume if you have moonkin talent you are using it.
		partyBuffs.MoonkinAura = core.MaxTristate(partyBuffs.MoonkinAura, proto.TristateEffect_TristateEffectRegular)
		for _, e := range druid.Equip {
			if e.ID == ravenGoddessItemID {
				partyBuffs.MoonkinAura = proto.TristateEffect_TristateEffectImproved
				break
			}
		}
	}
}

func (druid *Druid) Init(sim *core.Simulation) {
	druid.starfire8CastTemplate = druid.newStarfireTemplate(sim, 8)
	druid.starfire6CastTemplate = druid.newStarfireTemplate(sim, 6)
	druid.moonfireCastTemplate = druid.newMoonfireTemplate(sim)
	druid.wrathCastTemplate = druid.newWrathTemplate(sim)
	druid.insectSwarmCastTemplate = druid.newInsectSwarmTemplate(sim)
	druid.faerieFireCastTemplate = druid.newFaerieFireTemplate(sim)

	if druid.SelfBuffs.Innervate {
		druid.ExpectedManaPerInnervate = druid.SpiritManaRegenPerSecond() * 5 * 20
	}
}

func (druid *Druid) Reset(sim *core.Simulation) {
	// Cleanup and pending dots and casts
	druid.MoonfireSpell = core.SimpleSpell{}
	druid.InsectSwarmSpell = core.SimpleSpell{}
	druid.FaerieFireSpell = core.SimpleSpell{}
	druid.starfireSpell = core.SimpleSpell{}
	druid.wrathSpell = core.SimpleSpell{}
	druid.RebirthUsed = false

	innervateCD := time.Minute * 6
	if druid.malorne4p {
		innervateCD -= time.Second * 48
	}

	if druid.SelfBuffs.Innervate {
		druid.RemainingInnervateUsages = int(1 + (core.MaxDuration(0, sim.Duration))/innervateCD)
		druid.ExpectedBonusMana += druid.ExpectedManaPerInnervate * float64(druid.RemainingInnervateUsages)
	}
}

func (druid *Druid) Advance(sim *core.Simulation, elapsedTime time.Duration) {
	// druid should never be outside the 5s window, use combat regen.
	druid.Character.RegenManaCasting(sim, elapsedTime)
}

func (druid *Druid) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // does nothing
}

func (druid *Druid) applyOnHitTalents(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
	if druid.Talents.NaturesGrace && spellEffect.Crit {
		druid.NaturesGrace = true
	}
}

func (druid *Druid) applyNaturesGrace(spellCast *core.SpellCast) {
	if druid.NaturesGrace {
		spellCast.CastTime -= time.Millisecond * 500
		// This applies on cast complete, removing the effect.
		//  if it crits, during 'onspellhit' then it will be reapplied (see func above)
		spellCast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
			druid.NaturesGrace = false
		}
	}
}

func New(char core.Character, selfBuffs SelfBuffs, talents proto.DruidTalents) Druid {

	char.AddStat(stats.SpellHit, float64(talents.BalanceOfPower)*2*core.SpellHitRatingPerHitChance)

	char.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Intellect,
		ModifiedStat: stats.SpellCrit,
		Modifier: func(intellect float64, spellCrit float64) float64 {
			return spellCrit + (intellect/79.4)*core.SpellCritRatingPerCritChance
		},
	})

	if talents.LunarGuidance > 0 {
		bonus := (0.25 / 3) * float64(talents.LunarGuidance)
		char.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.SpellPower,
			Modifier: func(intellect float64, spellPower float64) float64 {
				return spellPower + intellect*bonus
			},
		})
	}

	if talents.Dreamstate > 0 {
		bonus := (0.1 / 3) * float64(talents.Dreamstate)
		char.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.MP5,
			Modifier: func(intellect float64, mp5 float64) float64 {
				return mp5 + intellect*bonus
			},
		})
	}

	if talents.Intensity > 0 {
		char.PseudoStats.SpiritRegenRateCasting = float64(talents.Intensity) * 0.1
	}

	if talents.HeartOfTheWild > 0 {
		bonus := 0.04 * float64(talents.HeartOfTheWild)
		char.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.Intellect,
			Modifier: func(intellect float64, _ float64) float64 {
				return intellect + intellect*bonus
			},
		})
	}

	if talents.SurvivalOfTheFittest > 0 {
		bonus := 0.01 * float64(talents.SurvivalOfTheFittest)
		char.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Stamina,
			ModifiedStat: stats.Stamina,
			Modifier: func(stat float64, _ float64) float64 {
				return stat + stat*bonus
			},
		})
		char.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Strength,
			ModifiedStat: stats.Strength,
			Modifier: func(stat float64, _ float64) float64 {
				return stat + stat*bonus
			},
		})
		char.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Agility,
			ModifiedStat: stats.Agility,
			Modifier: func(stat float64, _ float64) float64 {
				return stat + stat*bonus
			},
		})
		char.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.Intellect,
			Modifier: func(stat float64, _ float64) float64 {
				return stat + stat*bonus
			},
		})
		char.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.Intellect,
			Modifier: func(stat float64, _ float64) float64 {
				return stat + stat*bonus
			},
		})
		char.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Spirit,
			ModifiedStat: stats.Spirit,
			Modifier: func(stat float64, _ float64) float64 {
				return stat + stat*bonus
			},
		})
	}

	if talents.LivingSpirit > 0 {
		bonus := 0.05 * float64(talents.LivingSpirit)
		char.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Spirit,
			ModifiedStat: stats.Spirit,
			Modifier: func(spirit float64, _ float64) float64 {
				return spirit + spirit*bonus
			},
		})
	}

	if talents.NaturalPerfection > 0 {
		char.AddStat(stats.SpellCrit, float64(talents.NaturalPerfection)*1*core.SpellCritRatingPerCritChance)
	}

	druid := Druid{
		Character:   char,
		SelfBuffs:   selfBuffs,
		Talents:     talents,
		malorne4p:   ItemSetMalorne.CharacterHasSetBonus(&char, 4),
		RebirthUsed: false,
	}

	druid.registerNaturesSwiftnessCD()

	return druid
}

var NaturesSwiftnessAuraID = core.NewAuraID()
var NaturesSwiftnessCooldownID = core.NewCooldownID()

func (druid *Druid) registerNaturesSwiftnessCD() {
	if !druid.Talents.NaturesSwiftness {
		return
	}

	druid.AddMajorCooldown(core.MajorCooldown{
		CooldownID: NaturesSwiftnessCooldownID,
		Cooldown:   time.Minute * 3,
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) bool {
				// Don't use NS unless we're casting a full-length starfire or wrath.

				if character.HasTemporarySpellCastSpeedIncrease() {
					return false
				}

				character.AddAura(sim, core.Aura{
					ID:      NaturesSwiftnessAuraID,
					Name:    "Nature's Swiftness",
					Expires: core.NeverExpires,
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
						character.UpdateMajorCooldowns(sim)
						character.Metrics.AddInstantCast(core.ActionID{SpellID: 17116})
					},
				})
				return true
			}
		},
	})
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTauren, Class: proto.Class_ClassDruid}] = stats.Stats{
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2370,
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceNightElf, Class: proto.Class_ClassDruid}] = stats.Stats{
		stats.Strength:  73,
		stats.Agility:   75,
		stats.Stamina:   82,
		stats.Intellect: 120,
		stats.Spirit:    133,
		stats.Mana:      2370,
		stats.SpellCrit: 40.60, // 3.35% chance to crit shown on naked character screen
		// 4254 health shown on naked character
	}
}

// Agent is a generic way to access underlying druid on any of the agents (for example balance druid.)
type Agent interface {
	GetDruid() *Druid
}
