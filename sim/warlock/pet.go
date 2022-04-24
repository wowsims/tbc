package warlock

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type WarlockPet struct {
	core.Pet

	config PetConfig

	owner *Warlock

	// Time when pet should die, as per petUptime.
	deathTime time.Duration

	primaryAbility   PetAbility
	secondaryAbility PetAbility
}

func (warlock *Warlock) NewWarlockPet() *WarlockPet {
	// if warlock.Options.PetUptime <= 0 {
	// 	return nil
	// }
	petConfig := PetConfigs[warlock.Options.Summon]

	wp := &WarlockPet{
		Pet: core.NewPet(
			petConfig.Name,
			&warlock.Character,
			petConfig.Stats,
			petStatInheritance,
			true,
		),
		config: petConfig,
		owner:  warlock,
	}
	wp.AddStats(stats.Stats{
		stats.MeleeCrit: float64(warlock.Talents.DemonicTactics) * 5 * core.MeleeCritRatingPerCritChance,
		stats.SpellCrit: float64(warlock.Talents.DemonicTactics) * 5 * core.SpellCritRatingPerCritChance,
	})
	wp.PseudoStats.DamageDealtMultiplier *= 1.0 + (0.04 * float64(warlock.Talents.UnholyPower))

	wp.EnableManaBar()

	if petConfig.Melee {
		wp.EnableAutoAttacks(wp, core.AutoAttackOptions{
			MainHand: core.Weapon{
				// TODO: validate base weapon damage.
				BaseDamageMin:  176,
				BaseDamageMax:  232,
				SwingSpeed:     2,
				SwingDuration:  time.Second * 2,
				CritMultiplier: 2,
			},
			AutoSwingMelee: true,
		})
	}
	// wp.AutoAttacks.MHEffect.DamageMultiplier *= petConfig.DamageMultiplier
	switch warlock.Options.Summon {
	case proto.Warlock_Options_Felgaurd:
		wp.PseudoStats.DamageDealtMultiplier *= 1.0 + (0.01 * float64(warlock.Talents.MasterDemonologist))
		// Simulates a pre-stacked demonic frenzy
		wp.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.AttackPower,
			ModifiedStat: stats.AttackPower,
			Modifier: func(ap float64, _ float64) float64 {
				return ap * 1.5
			},
		})
	case proto.Warlock_Options_Succubus:
		wp.PseudoStats.DamageDealtMultiplier *= 1.0 + (0.02 * float64(warlock.Talents.MasterDemonologist))
	}

	if warlock.Talents.FelIntellect > 0 {
		intBonus := 1 + (0.05)*float64(warlock.Talents.FelIntellect)
		wp.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.Intellect,
			Modifier: func(in float64, _ float64) float64 {
				return in * intBonus
			},
		})
	}

	wp.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Intellect,
		ModifiedStat: stats.SpellCrit,
		Modifier: func(intellect float64, spellCrit float64) float64 {
			return spellCrit + (intellect/81)*core.SpellCritRatingPerCritChance
		},
	})

	wp.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Intellect,
		ModifiedStat: stats.Mana,
		Modifier: func(intellect float64, mana float64) float64 {
			return mana + intellect*petConfig.ManaIntRatio
		},
	})
	wp.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Strength,
		ModifiedStat: stats.AttackPower,
		Modifier: func(strength float64, attackPower float64) float64 {
			return attackPower + strength*2
		},
	})
	wp.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.MeleeCrit,
		Modifier: func(agility float64, meleeCrit float64) float64 {
			return meleeCrit + (agility/33)*core.MeleeCritRatingPerCritChance
		},
	})

	core.ApplyPetConsumeEffects(&wp.Character, warlock.Consumes)

	warlock.AddPet(wp)

	return wp
}

func (wp *WarlockPet) GetPet() *core.Pet {
	return &wp.Pet
}

func (wp *WarlockPet) Init(sim *core.Simulation) {
	wp.primaryAbility = wp.NewPetAbility(sim, wp.config.PrimaryAbility, true)
	wp.secondaryAbility = wp.NewPetAbility(sim, wp.config.SecondaryAbility, false)
}

func (wp *WarlockPet) Reset(sim *core.Simulation) {
	// if sim.Log != nil {
	// 	wp.Log(sim, "Total Pet stats: %s", wp.GetStats())
	// 	inheritedStats := hunterPetStatInheritance(wp.hunterOwner.GetStats())
	// 	wp.Log(sim, "Inherited Pet stats: %s", inheritedStats)
	// }

	// uptime := core.MinFloat(1, core.MaxFloat(0, wp.owner.Options.PetUptime))
	// wp.deathTime = time.Duration(float64(sim.Duration) * uptime)
}

func (wp *WarlockPet) OnGCDReady(sim *core.Simulation) {
	// if sim.CurrentTime > wp.deathTime {
	// 	wp.Disable(sim)
	// 	return
	// }

	target := sim.GetPrimaryTarget()
	if wp.config.RandomSelection {
		if sim.RandomFloat("Warlock Pet Ability") < 0.5 {
			if !wp.primaryAbility.TryCast(sim, target, wp) {
				wp.secondaryAbility.TryCast(sim, target, wp)
			}
		} else {
			if !wp.secondaryAbility.TryCast(sim, target, wp) {
				wp.primaryAbility.TryCast(sim, target, wp)
			}
		}
		return
	}

	if !wp.primaryAbility.TryCast(sim, target, wp) {
		if wp.secondaryAbility.Type != Unknown {
			wp.secondaryAbility.TryCast(sim, target, wp)
		} else {
			wp.WaitUntil(sim, wp.primaryAbility.CD.ReadyAt())
		}
	}
}

var petStatInheritance = func(ownerStats stats.Stats) stats.Stats {
	return stats.Stats{
		stats.Stamina:          ownerStats[stats.Stamina] * 0.3,
		stats.Intellect:        ownerStats[stats.Intellect] * 0.3,
		stats.Armor:            ownerStats[stats.Armor] * 0.35,
		stats.AttackPower:      (ownerStats[stats.SpellPower] + ownerStats[stats.ShadowSpellPower]) * 0.57,
		stats.SpellPower:       (ownerStats[stats.SpellPower] + ownerStats[stats.ShadowSpellPower]) * 0.15,
		stats.SpellPenetration: ownerStats[stats.SpellPenetration],
		stats.SpellHit:         ownerStats[stats.SpellHit],
		stats.MeleeHit:         ownerStats[stats.MeleeHit],

		// Resists, 40%
	}
}

type PetConfig struct {
	Name string
	// DamageMultiplier float64
	Melee        bool
	Stats        stats.Stats
	ManaIntRatio float64
	// Weapon

	// Randomly select between abilities instead of using a prio.
	RandomSelection bool

	PrimaryAbility   PetAbilityType
	SecondaryAbility PetAbilityType
}

var PetConfigs = map[proto.Warlock_Options_Summon]PetConfig{
	proto.Warlock_Options_Felgaurd: {
		Name: "Felguard",
		// DamageMultiplier: 1,
		Melee:            true,
		PrimaryAbility:   Cleave,
		SecondaryAbility: Intercept,
		ManaIntRatio:     11.5,
		Stats: stats.Stats{
			stats.Strength:  153,
			stats.Agility:   108,
			stats.Intellect: 196,
			stats.Mana:      893,
			stats.Spirit:    122,
			// Add 1.8% because pets aren't affected by that component of crit suppression.
			stats.MeleeCrit: (1.1515 + 1.8) * core.MeleeCritRatingPerCritChance,
		},
	},
	proto.Warlock_Options_Imp: {
		Name: "Imp",
		// DamageMultiplier: 1,
		ManaIntRatio:   4.9,
		Melee:          false,
		PrimaryAbility: Firebolt,
		// TODO: no idea if these stats are correct
		Stats: stats.Stats{
			stats.Strength:  153,
			stats.Agility:   108,
			stats.Intellect: 196,
			stats.Mana:      756,
			stats.Spirit:    122,
			// Add 1.8% because pets aren't affected by that component of crit suppression.
			stats.MeleeCrit: (1.1515 + 1.8) * core.MeleeCritRatingPerCritChance,
		},
	},
	proto.Warlock_Options_Succubus: {
		Name: "Succubus",
		// DamageMultiplier: 1,
		ManaIntRatio:   11.5,
		Melee:          true,
		PrimaryAbility: LashOfPain,
		// TODO: no idea if these stats are correct
		Stats: stats.Stats{
			stats.Strength:  153,
			stats.Agility:   108,
			stats.Intellect: 196,
			stats.Mana:      849,
			stats.Spirit:    122,
			// Add 1.8% because pets aren't affected by that component of crit suppression.
			stats.MeleeCrit: (1.1515 + 1.8) * core.MeleeCritRatingPerCritChance,
		},
	},
}

// Minion 		Health per bonus stamina 	Mana per bonus intellect
// Imp 			~8.4 						~4.9
// Voidwalker 	~11.0 						~11.5
// Sayaad 		~9.1 						~11.5
// Felhunter 	~9.5 						~11.5
// Felguard 	~11.0 						~11.5

// double Pet::GetSpellCritChance(SpellType) { return 0.0125 * GetIntellect() + 0.91 + stats.spell_crit_chance; }

// Spell hit 	Spell hit, physical hit, expertise, being capped will cap your minion for all three stats, see below for details
