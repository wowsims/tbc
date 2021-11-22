package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func NewShaman(character core.Character, talents proto.ShamanTalents, selfBuffs SelfBuffs) Shaman {
	shaman := Shaman{
		Character: character,
		Talents:   talents,
		SelfBuffs: selfBuffs,
	}

	if shaman.Talents.NaturesGuidance > 0 {
		shaman.AddStat(stats.SpellHit, float64(shaman.Talents.NaturesGuidance)*1*core.SpellHitRatingPerHitChance)
	}

	// Add Shaman stat dependencies
	shaman.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Intellect,
		ModifiedStat: stats.SpellCrit,
		Modifier: func(intellect float64, spellCrit float64) float64 {
			return spellCrit + (intellect/78.1)*core.SpellCritRatingPerCritChance
		},
	})

	shaman.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Strength,
		ModifiedStat: stats.AttackPower,
		Modifier: func(strength float64, attackPower float64) float64 {
			return attackPower + strength*2
		},
	})

	shaman.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.MeleeCrit,
		Modifier: func(agility float64, meleeCrit float64) float64 {
			return meleeCrit + (agility/33)*core.MeleeCritRatingPerCritChance
		},
	})

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

	if selfBuffs.WaterShield {
		shaman.AddStat(stats.MP5, 50)
	}

	shaman.registerElementalMasteryCD()
	shaman.registerNaturesSwiftnessCD()

	return shaman
}

// Which buffs this shaman is using.
type SelfBuffs struct {
	Bloodlust    bool
	WaterShield  bool
	TotemOfWrath bool
	WrathOfAir   bool
	ManaSpring   bool

	NextTotemDrops [4]time.Duration // track when to drop totems
}

// Indexes into NextTotemDrops for self buffs
const (
	AirTotem int = iota
	EarthTotem
	FireTotem
	WaterTotem
)

// Shaman represents a shaman character.
type Shaman struct {
	core.Character

	Talents   proto.ShamanTalents
	SelfBuffs SelfBuffs

	ElementalFocusStacks byte

	// "object pool" for shaman spells that are currently being cast.
	lightningBoltSpell   core.SimpleSpell
	lightningBoltSpellLO core.SimpleSpell

	chainLightningSpell    core.MultiTargetDirectDamageSpell
	chainLightningSpellLOs []core.MultiTargetDirectDamageSpell

	// Precomputed templated cast generator for quickly resetting cast fields.
	lightningBoltCastTemplate   core.SimpleSpellTemplate
	lightningBoltLOCastTemplate core.SimpleSpellTemplate

	chainLightningCastTemplate    core.MultiTargetDirectDamageSpellTemplate
	chainLightningLOCastTemplates []core.MultiTargetDirectDamageSpellTemplate
}

// Implemented by each Shaman spec.
type ShamanAgent interface {
	core.Agent

	// The Shaman controlled by this Agent.
	GetShaman() *Shaman
}

func (shaman *Shaman) GetCharacter() *core.Character {
	return &shaman.Character
}

func (shaman *Shaman) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
}
func (shaman *Shaman) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
	if shaman.SelfBuffs.Bloodlust {
		partyBuffs.Bloodlust += 1
	}

	if shaman.Talents.TotemOfWrath && shaman.SelfBuffs.TotemOfWrath {
		partyBuffs.TotemOfWrath += 1
	}

	if shaman.SelfBuffs.ManaSpring {
		partyBuffs.ManaSpringTotem = core.MaxTristate(partyBuffs.ManaSpringTotem, proto.TristateEffect_TristateEffectRegular)
		if shaman.Talents.RestorativeTotems == 5 {
			partyBuffs.ManaSpringTotem = proto.TristateEffect_TristateEffectImproved
		}
	}

	if shaman.SelfBuffs.WrathOfAir {
		woaValue := proto.TristateEffect_TristateEffectRegular
		if ItemSetCycloneRegalia.CharacterHasSetBonus(shaman.GetCharacter(), 2) {
			woaValue = proto.TristateEffect_TristateEffectImproved
		}
		partyBuffs.WrathOfAirTotem = core.MaxTristate(partyBuffs.WrathOfAirTotem, woaValue)
	}
}

func (shaman *Shaman) Init(sim *core.Simulation) {
	// Precompute all the spell templates.
	shaman.lightningBoltCastTemplate = shaman.newLightningBoltTemplate(sim, false)
	shaman.lightningBoltLOCastTemplate = shaman.newLightningBoltTemplate(sim, true)

	shaman.chainLightningCastTemplate = shaman.newChainLightningTemplate(sim, false)

	numHits := core.MinInt32(3, sim.GetNumTargets())
	shaman.chainLightningSpellLOs = make([]core.MultiTargetDirectDamageSpell, numHits)
	shaman.chainLightningLOCastTemplates = []core.MultiTargetDirectDamageSpellTemplate{}
	for i := int32(0); i < numHits; i++ {
		shaman.chainLightningLOCastTemplates = append(shaman.chainLightningLOCastTemplates, shaman.newChainLightningTemplate(sim, true))
	}
}

func (shaman *Shaman) Reset(sim *core.Simulation) {
	// Check to see if we are casting a totem to set its expire time.
	for i := range shaman.SelfBuffs.NextTotemDrops {
		shaman.SelfBuffs.NextTotemDrops[i] = core.NeverExpires
		switch i {
		case AirTotem:
			if shaman.SelfBuffs.WrathOfAir {
				shaman.SelfBuffs.NextTotemDrops[i] = time.Second * 120 // 2 min until drop totems
			}
		case FireTotem:
			if shaman.SelfBuffs.TotemOfWrath {
				shaman.SelfBuffs.NextTotemDrops[i] = time.Second * 120 // 2 min until drop totems
			}
		case WaterTotem:
			if shaman.SelfBuffs.ManaSpring {
				shaman.SelfBuffs.NextTotemDrops[i] = time.Second * 120 // 2 min until drop totems
			}
		}
	}

	// Reset all spells so any pending casts are cleaned up
	shaman.lightningBoltSpell = core.SimpleSpell{}
	shaman.lightningBoltSpellLO = core.SimpleSpell{}
	shaman.chainLightningSpell = core.MultiTargetDirectDamageSpell{}

	numHits := core.MinInt32(3, sim.GetNumTargets())
	shaman.chainLightningSpellLOs = make([]core.MultiTargetDirectDamageSpell, numHits)
}

func (shaman *Shaman) Advance(sim *core.Simulation, elapsedTime time.Duration) {
	// Shaman should never be outside the 5s window, use combat regen
	shaman.Character.RegenManaCasting(sim, elapsedTime)
}

var ElementalMasteryAuraID = core.NewAuraID()
var ElementalMasteryCooldownID = core.NewCooldownID()

func (shaman *Shaman) registerElementalMasteryCD() {
	if !shaman.Talents.ElementalMastery {
		return
	}

	shaman.AddMajorCooldown(core.MajorCooldown{
		CooldownID: ElementalMasteryCooldownID,
		Cooldown:   time.Minute * 3,
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) bool {
				character.Metrics.AddInstantCast(core.ActionID{SpellID: 16166})

				character.AddAura(sim, core.Aura{
					ID:      ElementalMasteryAuraID,
					Name:    "Elemental Mastery",
					Expires: core.NeverExpires,
					OnCast: func(sim *core.Simulation, cast *core.Cast) {
						cast.ManaCost = 0
						cast.GuaranteedCrit = true
					},
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						// Remove the buff and put skill on CD
						character.SetCD(ElementalMasteryCooldownID, sim.CurrentTime+time.Minute*3)
						character.RemoveAura(sim, ElementalMasteryAuraID)
						character.UpdateMajorCooldowns(sim)
					},
				})
				return true
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

	shaman.AddMajorCooldown(core.MajorCooldown{
		CooldownID: NaturesSwiftnessCooldownID,
		Cooldown:   time.Minute * 3,
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) bool {
				// Don't use NS unless we're casting a full-length lightning bolt, which is
				// the only spell shamans have with a cast longer than GCD.

				if character.HasTemporarySpellCastSpeedIncrease() {
					return false
				}

				character.AddAura(sim, core.Aura{
					ID:      NaturesSwiftnessAuraID,
					Name:    "Nature's Swiftness",
					Expires: core.NeverExpires,
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
						character.UpdateMajorCooldowns(sim)
						character.Metrics.AddInstantCast(core.ActionID{SpellID: 16188})
					},
				})
				return true
			}
		},
	})
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDraenei, Class: proto.Class_ClassShaman}] = stats.Stats{
		stats.Strength:    103,
		stats.Agility:     61,
		stats.Stamina:     113,
		stats.Intellect:   109,
		stats.Spirit:      122,
		stats.Mana:        2678,
		stats.SpellCrit:   47.89,
		stats.AttackPower: 120,
		stats.MeleeCrit:   50.16,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceOrc, Class: proto.Class_ClassShaman}] = stats.Stats{
		stats.Strength:    105,
		stats.Agility:     61,
		stats.Stamina:     116,
		stats.Intellect:   105,
		stats.Spirit:      123,
		stats.Mana:        2678,
		stats.SpellCrit:   47.89,
		stats.AttackPower: 120,
		stats.MeleeCrit:   50.16,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTauren, Class: proto.Class_ClassShaman}] = stats.Stats{
		stats.Strength:    107,
		stats.Agility:     59,
		stats.Stamina:     116,
		stats.Intellect:   103,
		stats.Spirit:      122,
		stats.Mana:        2678,
		stats.SpellCrit:   47.89,
		stats.AttackPower: 120,
		stats.MeleeCrit:   49.72,
	}

	trollStats := stats.Stats{
		stats.Strength:    103,
		stats.Agility:     66,
		stats.Stamina:     115,
		stats.Intellect:   104,
		stats.Spirit:      121,
		stats.Mana:        2678,
		stats.SpellCrit:   47.89,
		stats.AttackPower: 120,
		stats.MeleeCrit:   51.23,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll10, Class: proto.Class_ClassShaman}] = trollStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll30, Class: proto.Class_ClassShaman}] = trollStats
}
