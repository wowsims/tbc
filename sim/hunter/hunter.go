package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const ThoridalTheStarsFuryItemID = 34334

func RegisterHunter() {
	core.RegisterAgentFactory(
		proto.Player_Hunter{},
		func(character core.Character, options proto.Player) core.Agent {
			return NewHunter(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_Hunter)
			if !ok {
				panic("Invalid spec value for Hunter!")
			}
			player.Spec = playerSpec
		},
	)
}

type Hunter struct {
	core.Character

	Talents  proto.HunterTalents
	Options  proto.Hunter_Options
	Rotation proto.Hunter_Rotation

	pet *HunterPet

	AmmoDPS         float64
	AmmoDamageBonus float64

	aspectOfTheViper bool // False indicates aspect of the hawk.

	hasGronnstalker2Pc bool

	killCommandEnabledUntil time.Duration       // Time that KC enablement expires.
	killCommandBlocked      bool                // True while Steady Shot is casting, to prevent KC.
	killCommandAction       *core.PendingAction // Action to use KC when its comes off CD.

	timeToWeave time.Duration

	raptorStrikeCost float64 // Cached mana cost of raptor strike.

	nextAction   int
	nextActionAt time.Duration

	// Expected single-cast damage values calculated by the presim, used for adaptive logic.
	avgShootDmg  float64
	avgWeaveDmg  float64
	avgSteadyDmg float64
	avgMultiDmg  float64
	avgArcaneDmg float64

	// Cached values for adaptive rotation calcs.
	rangedSwingSpeed   float64
	rangedWindup       float64
	shootDPS           float64
	weaveDPS           float64
	steadyDPS          float64
	steadyShotCastTime float64
	multiShotCastTime  float64
	useMultiForCatchup bool

	aimedShotTemplate core.MeleeAbilityTemplate
	aimedShot         core.ActiveMeleeAbility

	arcaneShotTemplate core.MeleeAbilityTemplate
	arcaneShot         core.ActiveMeleeAbility

	aspectOfTheHawkTemplate  core.SimpleCast
	aspectOfTheViperTemplate core.SimpleCast

	killCommandTemplate core.SimpleCast

	multiShotCastTemplate core.SimpleCast
	multiShotCast         core.SimpleCast

	multiShotAbilityTemplate core.MeleeAbilityTemplate
	multiShotAbility         core.ActiveMeleeAbility

	raptorStrikeTemplate core.MeleeAbilityTemplate
	raptorStrike         core.ActiveMeleeAbility

	scorpidStingTemplate core.MeleeAbilityTemplate
	scorpidSting         core.ActiveMeleeAbility

	serpentStingTemplate core.MeleeAbilityTemplate
	serpentSting         core.ActiveMeleeAbility

	serpentStingDotTemplate core.SimpleSpellTemplate
	serpentStingDot         core.SimpleSpell

	steadyShotCastTemplate core.SimpleCast
	steadyShotCast         core.SimpleCast

	steadyShotAbilityTemplate core.MeleeAbilityTemplate
	steadyShotAbility         core.ActiveMeleeAbility

	fakeHardcast core.Cast
}

func (hunter *Hunter) GetCharacter() *core.Character {
	return &hunter.Character
}

func (hunter *Hunter) GetHunter() *Hunter {
	return hunter
}

func (hunter *Hunter) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
}
func (hunter *Hunter) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
	if hunter.Talents.TrueshotAura {
		partyBuffs.TrueshotAura = true
	}
}

func (hunter *Hunter) Init(sim *core.Simulation) {
	// Update auto crit multipliers now that we have the targets.
	hunter.AutoAttacks.MH.CritMultiplier = hunter.critMultiplier(false, sim.GetPrimaryTarget())
	hunter.AutoAttacks.OH.CritMultiplier = hunter.critMultiplier(false, sim.GetPrimaryTarget())
	hunter.AutoAttacks.Ranged.CritMultiplier = hunter.critMultiplier(true, sim.GetPrimaryTarget())

	// Precompute all the spell templates.
	hunter.aimedShotTemplate = hunter.newAimedShotTemplate(sim)
	hunter.arcaneShotTemplate = hunter.newArcaneShotTemplate(sim)
	hunter.aspectOfTheHawkTemplate = hunter.newAspectOfTheHawkTemplate(sim)
	hunter.aspectOfTheViperTemplate = hunter.newAspectOfTheViperTemplate(sim)
	hunter.killCommandTemplate = hunter.newKillCommandTemplate(sim)
	hunter.multiShotCastTemplate = hunter.newMultiShotCastTemplate(sim)
	hunter.multiShotAbilityTemplate = hunter.newMultiShotAbilityTemplate(sim)
	hunter.raptorStrikeTemplate = hunter.newRaptorStrikeTemplate(sim)
	hunter.scorpidStingTemplate = hunter.newScorpidStingTemplate(sim)
	hunter.serpentStingTemplate = hunter.newSerpentStingTemplate(sim)
	hunter.serpentStingDotTemplate = hunter.newSerpentStingDotTemplate(sim)
	hunter.steadyShotCastTemplate = hunter.newSteadyShotCastTemplate(sim)
	hunter.steadyShotAbilityTemplate = hunter.newSteadyShotAbilityTemplate(sim)

	hunter.fakeHardcast = core.Cast{
		Character:   &hunter.Character,
		IgnoreHaste: true,
		CastTime:    hunter.timeToWeave,
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			hunter.rotation(sim, false)
		},
	}
}

func (hunter *Hunter) Reset(sim *core.Simulation) {
	hunter.aspectOfTheViper = false
	hunter.killCommandEnabledUntil = 0
	hunter.killCommandBlocked = false
	hunter.killCommandAction.NextActionAt = 0
	hunter.nextAction = OptionNone
	hunter.nextActionAt = 0
	hunter.rangedSwingSpeed = 0

	target := sim.GetPrimaryTarget()
	impHuntersMark := hunter.Talents.ImprovedHuntersMark
	if !target.HasAura(core.HuntersMarkDebuffID) || target.NumStacks(core.HuntersMarkDebuffID) < impHuntersMark {
		target.AddAura(sim, core.HuntersMarkAura(impHuntersMark, false))
	}
}

func NewHunter(character core.Character, options proto.Player) *Hunter {
	hunterOptions := options.GetHunter()

	hunter := &Hunter{
		Character: character,
		Talents:   *hunterOptions.Talents,
		Options:   *hunterOptions.Options,
		Rotation:  *hunterOptions.Rotation,

		timeToWeave: time.Millisecond * time.Duration(hunterOptions.Rotation.TimeToWeaveMs),
	}
	hunter.hasGronnstalker2Pc = ItemSetGronnstalker.CharacterHasSetBonus(&hunter.Character, 2)

	hunter.PseudoStats.RangedSpeedMultiplier = 1
	hunter.EnableManaBar()
	hunter.EnableAutoAttacks(hunter, core.AutoAttackOptions{
		// We don't know crit multiplier until later when we see the target so just
		// use 0 for now.
		MainHand: hunter.WeaponFromMainHand(0),
		OffHand:  hunter.WeaponFromOffHand(0),
		Ranged:   hunter.WeaponFromRanged(0),
		ReplaceMHSwing: func(sim *core.Simulation) *core.ActiveMeleeAbility {
			return hunter.TryRaptorStrike(sim)
		},
	})

	hunter.pet = hunter.NewHunterPet()

	hunter.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Intellect,
		ModifiedStat: stats.SpellCrit,
		Modifier: func(intellect float64, spellCrit float64) float64 {
			return spellCrit + (intellect/55)*core.SpellCritRatingPerCritChance
		},
	})

	hunter.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Strength,
		ModifiedStat: stats.AttackPower,
		Modifier: func(strength float64, attackPower float64) float64 {
			return attackPower + strength*1
		},
	})

	hunter.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.RangedAttackPower,
		Modifier: func(agility float64, rap float64) float64 {
			return rap + agility*1
		},
	})

	hunter.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.MeleeCrit,
		Modifier: func(agility float64, meleeCrit float64) float64 {
			return meleeCrit + (agility/40)*core.MeleeCritRatingPerCritChance
		},
	})

	if hunter.HasRangedWeapon() && hunter.GetRangedWeapon().ID != ThoridalTheStarsFuryItemID {
		switch hunter.Options.Ammo {
		case proto.Hunter_Options_TimelessArrow:
			hunter.AmmoDPS = 53
		case proto.Hunter_Options_MysteriousArrow:
			hunter.AmmoDPS = 46.5
		case proto.Hunter_Options_AdamantiteStinger:
			hunter.AmmoDPS = 43
		case proto.Hunter_Options_WardensArrow:
			hunter.AmmoDPS = 37
		case proto.Hunter_Options_HalaaniRazorshaft:
			hunter.AmmoDPS = 34
		case proto.Hunter_Options_BlackflightArrow:
			hunter.AmmoDPS = 32
		}
		hunter.AmmoDamageBonus = hunter.AmmoDPS * hunter.AutoAttacks.Ranged.SwingSpeed
		hunter.AutoAttacks.RangedAuto.Effect.WeaponInput.FlatDamageBonus += hunter.AmmoDamageBonus
	}

	switch hunter.Options.QuiverBonus {
	case proto.Hunter_Options_Speed10:
		hunter.PseudoStats.RangedSpeedMultiplier *= 1.1
	case proto.Hunter_Options_Speed11:
		hunter.PseudoStats.RangedSpeedMultiplier *= 1.11
	case proto.Hunter_Options_Speed12:
		hunter.PseudoStats.RangedSpeedMultiplier *= 1.12
	case proto.Hunter_Options_Speed13:
		hunter.PseudoStats.RangedSpeedMultiplier *= 1.13
	case proto.Hunter_Options_Speed14:
		hunter.PseudoStats.RangedSpeedMultiplier *= 1.14
	case proto.Hunter_Options_Speed15:
		hunter.PseudoStats.RangedSpeedMultiplier *= 1.15
	}

	hunter.applyTalents()
	hunter.registerRapidFireCD()
	hunter.applyAspectOfTheHawk()
	hunter.applyKillCommand()

	return hunter
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceBloodElf, Class: proto.Class_ClassHunter}] = stats.Stats{
		stats.Strength:  61,
		stats.Agility:   153,
		stats.Stamina:   106,
		stats.Intellect: 81,
		stats.Spirit:    82,
		stats.Mana:      3383,

		stats.AttackPower:       120,
		stats.RangedAttackPower: 130,
		stats.MeleeCrit:         -1.53 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDraenei, Class: proto.Class_ClassHunter}] = stats.Stats{
		stats.Strength:  65,
		stats.Agility:   148,
		stats.Stamina:   107,
		stats.Intellect: 78,
		stats.Spirit:    85,
		stats.Mana:      3383,

		stats.AttackPower:       120,
		stats.RangedAttackPower: 130,
		stats.MeleeCrit:         -1.53 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDwarf, Class: proto.Class_ClassHunter}] = stats.Stats{
		stats.Strength:  66,
		stats.Agility:   147,
		stats.Stamina:   111,
		stats.Intellect: 76,
		stats.Spirit:    82,
		stats.Mana:      3383,

		stats.AttackPower:       120,
		stats.RangedAttackPower: 130,
		stats.MeleeCrit:         -1.53 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceNightElf, Class: proto.Class_ClassHunter}] = stats.Stats{
		stats.Strength:  61,
		stats.Agility:   156,
		stats.Stamina:   107,
		stats.Intellect: 77,
		stats.Spirit:    83,
		stats.Mana:      3383,

		stats.AttackPower:       120,
		stats.RangedAttackPower: 130,
		stats.MeleeCrit:         -1.53 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceOrc, Class: proto.Class_ClassHunter}] = stats.Stats{
		stats.Strength:  67,
		stats.Agility:   148,
		stats.Stamina:   110,
		stats.Intellect: 74,
		stats.Spirit:    86,
		stats.Mana:      3383,

		stats.AttackPower:       120,
		stats.RangedAttackPower: 130,
		stats.MeleeCrit:         -1.53 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTauren, Class: proto.Class_ClassHunter}] = stats.Stats{
		stats.Strength:  69,
		stats.Agility:   146,
		stats.Stamina:   110,
		stats.Intellect: 72,
		stats.Spirit:    85,
		stats.Mana:      3383,

		stats.AttackPower:       120,
		stats.RangedAttackPower: 130,
		stats.MeleeCrit:         -1.53 * core.MeleeCritRatingPerCritChance,
	}
	trollStats := stats.Stats{
		stats.Strength:  65,
		stats.Agility:   153,
		stats.Stamina:   109,
		stats.Intellect: 73,
		stats.Spirit:    84,
		stats.Mana:      3383,

		stats.AttackPower:       120,
		stats.RangedAttackPower: 130,
		stats.MeleeCrit:         -1.53 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll10, Class: proto.Class_ClassHunter}] = trollStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll30, Class: proto.Class_ClassHunter}] = trollStats
}

// Agent is a generic way to access underlying hunter on any of the agents.
type Agent interface {
	GetHunter() *Hunter
}
