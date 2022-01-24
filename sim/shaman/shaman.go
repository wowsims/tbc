package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func NewShaman(character core.Character, talents proto.ShamanTalents, totems proto.ShamanTotems, selfBuffs SelfBuffs) *Shaman {
	if totems.WindfuryTotemRank == 0 {
		// If rank is 0, disable windfury options.
		totems.TwistWindfury = false
		if totems.Air == proto.AirTotem_WindfuryTotem {
			totems.Air = proto.AirTotem_NoAirTotem
		}
	}
	if totems.Air == proto.AirTotem_WindfuryTotem {
		// No need to twist windfury if its already the default totem.
		totems.TwistWindfury = false
	} else if totems.Air == proto.AirTotem_NoAirTotem && totems.TwistWindfury {
		// If twisting windfury without a default air totem, make windfury the default instead.
		totems.Air = proto.AirTotem_WindfuryTotem
		totems.TwistWindfury = false
	}
	if totems.Fire == proto.FireTotem_TotemOfWrath && !talents.TotemOfWrath {
		totems.Fire = proto.FireTotem_NoFireTotem
	}
	if totems.Air != proto.AirTotem_WrathOfAirTotem && selfBuffs.SnapshotT42Pc {
		selfBuffs.SnapshotT42Pc = false
	}

	shaman := &Shaman{
		Character: character,
		Talents:   talents,
		Totems:    totems,
		SelfBuffs: selfBuffs,
	}
	shaman.EnableManaBar()

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
			return meleeCrit + (agility/25)*core.MeleeCritRatingPerCritChance
		},
	})

	if selfBuffs.WaterShield {
		shaman.AddStat(stats.MP5, 50)
	}

	shaman.registerBloodlustCD()
	shaman.applyTalents()

	return shaman
}

// Which buffs this shaman is using.
type SelfBuffs struct {
	Bloodlust     bool
	WaterShield   bool
	SnapshotT42Pc bool
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

	Totems proto.ShamanTotems

	// The type of totem which should be dropped next and time to drop it, for
	// each totem type (earth, air, fire, water).
	NextTotemDropType [4]int32
	NextTotemDrops    [4]time.Duration

	ElementalFocusStacks byte

	WaitingForMana float64

	// "object pool" for shaman spells that are currently being cast.
	lightningBoltSpell   core.SimpleSpell
	lightningBoltSpellLO core.SimpleSpell

	chainLightningSpell    core.SimpleSpell
	chainLightningSpellLOs []core.SimpleSpell

	// Precomputed templated cast generator for quickly resetting cast fields.
	lightningBoltCastTemplate   core.SimpleSpellTemplate
	lightningBoltLOCastTemplate core.SimpleSpellTemplate

	chainLightningCastTemplate    core.SimpleSpellTemplate
	chainLightningLOCastTemplates []core.SimpleSpellTemplate

	stormstrikeTemplate core.MeleeAbilityTemplate
	stormstrikeSpell    core.ActiveMeleeAbility

	// Shocks
	shockSpell         core.SimpleSpell
	earthShockTemplate core.SimpleSpellTemplate
	frostShockTemplate core.SimpleSpellTemplate

	// Flame shock needs a separate spell object because of the dot.
	FlameShockSpell    core.SimpleSpell
	flameShockTemplate core.SimpleSpellTemplate

	strengthOfEarthTotemTemplate core.SimpleCast
	tremorTotemTemplate          core.SimpleCast
	graceOfAirTotemTemplate      core.SimpleCast
	wrathOfAirTotemTemplate      core.SimpleCast
	tranquilAirTotemTemplate     core.SimpleCast
	windfuryTotemTemplate        core.SimpleCast
	totemOfWrathTemplate         core.SimpleCast
	manaSpringTotemTemplate      core.SimpleCast
	totemSpell                   core.SimpleCast

	searingTotemTemplate core.SimpleSpellTemplate
	magmaTotemTemplate   core.SimpleSpellTemplate
	novaTotemTemplate    core.SimpleSpellTemplate
	FireTotemSpell       core.SimpleSpell
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
	if shaman.Totems.Fire == proto.FireTotem_TotemOfWrath {
		partyBuffs.TotemOfWrath += 1
	}

	switch shaman.Totems.Water {
	case proto.WaterTotem_ManaSpringTotem:
		partyBuffs.ManaSpringTotem = core.MaxTristate(partyBuffs.ManaSpringTotem, proto.TristateEffect_TristateEffectRegular)
		if shaman.Talents.RestorativeTotems == 5 {
			partyBuffs.ManaSpringTotem = proto.TristateEffect_TristateEffectImproved
		}
	}

	switch shaman.Totems.Air {
	case proto.AirTotem_WrathOfAirTotem:
		woaValue := proto.TristateEffect_TristateEffectRegular
		if ItemSetCycloneRegalia.CharacterHasSetBonus(shaman.GetCharacter(), 2) {
			woaValue = proto.TristateEffect_TristateEffectImproved
		} else if shaman.SelfBuffs.SnapshotT42Pc {
			partyBuffs.SnapshotImprovedWrathOfAirTotem = true
		}
		partyBuffs.WrathOfAirTotem = core.MaxTristate(partyBuffs.WrathOfAirTotem, woaValue)
	case proto.AirTotem_GraceOfAirTotem:
		value := proto.TristateEffect_TristateEffectRegular
		if shaman.Talents.EnhancingTotems == 2 {
			value = proto.TristateEffect_TristateEffectImproved
		}
		partyBuffs.GraceOfAirTotem = core.MaxTristate(partyBuffs.GraceOfAirTotem, value)
	case proto.AirTotem_WindfuryTotem:
		if shaman.Totems.WindfuryTotemRank > partyBuffs.WindfuryTotemRank {
			partyBuffs.WindfuryTotemRank = shaman.Totems.WindfuryTotemRank
			partyBuffs.WindfuryTotemIwt = shaman.Talents.ImprovedWeaponTotems
		} else if shaman.Totems.WindfuryTotemRank == partyBuffs.WindfuryTotemRank {
			partyBuffs.WindfuryTotemIwt = core.MaxInt32(partyBuffs.WindfuryTotemIwt, shaman.Talents.ImprovedWeaponTotems)
		}
	case proto.AirTotem_TranquilAirTotem:
		partyBuffs.TranquilAirTotem = true
	}

	switch shaman.Totems.Earth {
	case proto.EarthTotem_StrengthOfEarthTotem:
		value := proto.StrengthOfEarthType_Basic
		if shaman.Talents.EnhancingTotems == 2 {
			value = proto.StrengthOfEarthType_EnhancingTotems
		}
		if ItemSetCycloneHarness.CharacterHasSetBonus(&shaman.Character, 2) {
			if value == proto.StrengthOfEarthType_EnhancingTotems {
				value = proto.StrengthOfEarthType_EnhancingAndCyclone
			} else {
				value = proto.StrengthOfEarthType_CycloneBonus
			}
		}
		if value > partyBuffs.StrengthOfEarthTotem {
			partyBuffs.StrengthOfEarthTotem = value
		}
	}
}

func (shaman *Shaman) Init(sim *core.Simulation) {
	// Precompute all the spell templates.
	shaman.stormstrikeTemplate = shaman.newStormstrikeTemplate(sim)
	shaman.lightningBoltCastTemplate = shaman.newLightningBoltTemplate(sim, false)
	shaman.lightningBoltLOCastTemplate = shaman.newLightningBoltTemplate(sim, true)

	shaman.chainLightningCastTemplate = shaman.newChainLightningTemplate(sim, false)

	numHits := core.MinInt32(3, sim.GetNumTargets())
	shaman.chainLightningSpellLOs = make([]core.SimpleSpell, numHits)
	shaman.chainLightningLOCastTemplates = []core.SimpleSpellTemplate{}
	for i := int32(0); i < numHits; i++ {
		shaman.chainLightningLOCastTemplates = append(shaman.chainLightningLOCastTemplates, shaman.newChainLightningTemplate(sim, true))
	}
	shaman.earthShockTemplate = shaman.newEarthShockTemplate(sim)
	shaman.flameShockTemplate = shaman.newFlameShockTemplate(sim)
	shaman.frostShockTemplate = shaman.newFrostShockTemplate(sim)

	shaman.strengthOfEarthTotemTemplate = shaman.newStrengthOfEarthTotemTemplate(sim)
	shaman.tremorTotemTemplate = shaman.newTremorTotemTemplate(sim)
	shaman.wrathOfAirTotemTemplate = shaman.newWrathOfAirTotemTemplate(sim)
	shaman.tranquilAirTotemTemplate = shaman.newTranquilAirTotemTemplate(sim)
	shaman.graceOfAirTotemTemplate = shaman.newGraceOfAirTotemTemplate(sim)
	shaman.windfuryTotemTemplate = shaman.newWindfuryTotemTemplate(sim, shaman.Totems.WindfuryTotemRank)
	shaman.manaSpringTotemTemplate = shaman.newManaSpringTotemTemplate(sim)
	shaman.totemOfWrathTemplate = shaman.newTotemOfWrathTemplate(sim)

	shaman.searingTotemTemplate = shaman.newSearingTotemTemplate(sim)
	shaman.magmaTotemTemplate = shaman.newMagmaTotemTemplate(sim)
	shaman.novaTotemTemplate = shaman.newNovaTotemTemplate(sim)
}

func (shaman *Shaman) Reset(sim *core.Simulation) {
	// Check to see if we are casting a totem to set its expire time.
	for i := range shaman.NextTotemDrops {
		shaman.NextTotemDrops[i] = core.NeverExpires
		switch i {
		case AirTotem:
			if shaman.Totems.Air != proto.AirTotem_NoAirTotem {
				shaman.NextTotemDrops[i] = time.Second * 120 // 2 min until drop totems
				shaman.NextTotemDropType[i] = int32(shaman.Totems.Air)
			}
			if shaman.Totems.TwistWindfury {
				shaman.NextTotemDropType[i] = int32(proto.AirTotem_WindfuryTotem)
				shaman.NextTotemDrops[i] = time.Second * 10 // gotta recast windfury after 10s
			}
		case EarthTotem:
			if shaman.Totems.Earth != proto.EarthTotem_NoEarthTotem {
				shaman.NextTotemDrops[i] = time.Second * 120 // 2 min until drop totems
				shaman.NextTotemDropType[i] = int32(shaman.Totems.Earth)
			}
		case FireTotem:
			shaman.NextTotemDropType[i] = int32(shaman.Totems.Fire)
			if shaman.Totems.TwistFireNova {
				shaman.NextTotemDropType[FireTotem] = int32(proto.FireTotem_FireNovaTotem) // start by dropping nova, then alternating.
			}
			if shaman.NextTotemDropType[i] != int32(proto.FireTotem_NoFireTotem) {
				shaman.NextTotemDrops[i] = time.Second * 120 // 2 min until drop totems
				if shaman.Totems.Fire != proto.FireTotem_TotemOfWrath {
					shaman.NextTotemDrops[i] = 0 // attack totems we drop immediately
				}
			}
		case WaterTotem:
			if shaman.Totems.Water == proto.WaterTotem_ManaSpringTotem {
				shaman.NextTotemDrops[i] = time.Second * 120 // 2 min until drop totems
			}
		}
	}

	shaman.ElementalFocusStacks = 0
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDraenei, Class: proto.Class_ClassShaman}] = stats.Stats{
		stats.Strength:    103,
		stats.Agility:     61,
		stats.Stamina:     113,
		stats.Intellect:   109,
		stats.Spirit:      122,
		stats.Mana:        2958,
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
		stats.Mana:        2958,
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
		stats.Mana:        2958,
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
		stats.Mana:        2958,
		stats.SpellCrit:   47.89,
		stats.AttackPower: 120,
		stats.MeleeCrit:   51.23,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll10, Class: proto.Class_ClassShaman}] = trollStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll30, Class: proto.Class_ClassShaman}] = trollStats
}
