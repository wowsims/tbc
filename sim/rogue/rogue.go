package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func RegisterRogue() {
	core.RegisterAgentFactory(
		proto.Player_Rogue{},
		proto.Spec_SpecRogue,
		func(character core.Character, options proto.Player) core.Agent {
			return NewRogue(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_Rogue)
			if !ok {
				panic("Invalid spec value for Rogue!")
			}
			player.Spec = playerSpec
		},
	)
}

const (
	SpellFlagBuilder  = core.SpellExtrasAgentReserved1
	SpellFlagFinisher = core.SpellExtrasAgentReserved2
)

type Rogue struct {
	core.Character

	Talents  proto.RogueTalents
	Options  proto.Rogue_Options
	Rotation proto.Rogue_Rotation

	// Current rotation plan.
	plan int

	// Cached values for calculating rotation.
	energyPerSecondAvg    float64
	eaBuildTime           time.Duration // Time to build EA following a finisher at ~35 energy
	sliceAndDiceDurations [6]time.Duration

	doneSND bool // Current SND will last for the rest of the iteration
	doneEA  bool // Current EA will last for the rest of the iteration, or not using EA

	deathmantle4pcProc bool
	disabledMCDs       []*core.MajorCooldown

	shivEnergyCost    float64
	builderEnergyCost float64
	newBuilder        func(sim *core.Simulation, target *core.Target) *core.SimpleSpell

	sinisterStrikeTemplate core.SimpleSpellTemplate
	sinisterStrike         core.SimpleSpell

	backstabTemplate core.SimpleSpellTemplate
	backstab         core.SimpleSpell

	hemorrhageTemplate core.SimpleSpellTemplate
	hemorrhage         core.SimpleSpell

	mutilateTemplate core.SimpleSpellTemplate
	mutilate         core.SimpleSpell

	shivTemplate core.SimpleSpellTemplate
	shiv         core.SimpleSpell

	finishingMoveEffectApplier func(sim *core.Simulation, numPoints int32)

	castSliceAndDice func()

	eviscerateEnergyCost float64
	eviscerateTemplate   core.SimpleSpellTemplate
	eviscerate           core.SimpleSpell

	envenomEnergyCost float64
	envenomTemplate   core.SimpleSpellTemplate
	envenom           core.SimpleSpell

	exposeArmorTemplate core.SimpleSpellTemplate
	exposeArmor         core.SimpleSpell

	ruptureTemplate core.SimpleSpellTemplate
	rupture         core.SimpleSpell

	deadlyPoisonStacks   int
	deadlyPoisonTemplate core.SimpleSpellTemplate
	deadlyPoison         core.SimpleSpell

	deadlyPoisonRefreshTemplate core.SimpleSpellTemplate
	deadlyPoisonRefresh         core.SimpleSpell

	instantPoisonTemplate core.SimpleSpellTemplate
	instantPoison         core.SimpleSpell
}

func (rogue *Rogue) GetCharacter() *core.Character {
	return &rogue.Character
}

func (rogue *Rogue) GetRogue() *Rogue {
	return rogue
}

func (rogue *Rogue) AddRaidBuffs(raidBuffs *proto.RaidBuffs)    {}
func (rogue *Rogue) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {}

func (rogue *Rogue) Finalize(raid *core.Raid) {
	// Need to apply poisons now so we can check for WF totem.
	rogue.applyPoisons()
}

func (rogue *Rogue) newAbility(actionID core.ActionID, cost float64, spellExtras core.SpellExtras, procMask core.ProcMask) core.SimpleSpell {
	return core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            actionID,
				Character:           &rogue.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 time.Second,
				IgnoreHaste:         true,
				BaseCost: core.ResourceCost{
					Type:  stats.Energy,
					Value: cost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Energy,
					Value: cost,
				},
				CritMultiplier: rogue.critMultiplier(procMask.Matches(core.ProcMaskMeleeMH), spellExtras.Matches(SpellFlagBuilder)),
				SpellExtras:    spellExtras,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               procMask,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
		},
	}
}

func (rogue *Rogue) ApplyFinisher(sim *core.Simulation, actionID core.ActionID) {
	numPoints := rogue.ComboPoints()
	rogue.SpendComboPoints(sim, actionID)
	rogue.finishingMoveEffectApplier(sim, numPoints)
}

func (rogue *Rogue) Init(sim *core.Simulation) {
	// Precompute all the spell templates.
	rogue.sinisterStrikeTemplate = rogue.newSinisterStrikeTemplate(sim)
	rogue.backstabTemplate = rogue.newBackstabTemplate(sim)
	rogue.hemorrhageTemplate = rogue.newHemorrhageTemplate(sim)
	rogue.mutilateTemplate = rogue.newMutilateTemplate(sim)
	rogue.shivTemplate = rogue.newShivTemplate(sim)

	rogue.finishingMoveEffectApplier = rogue.makeFinishingMoveEffectApplier(sim)

	rogue.initSliceAndDice(sim)
	rogue.eviscerateTemplate = rogue.newEviscerateTemplate(sim)
	rogue.exposeArmorTemplate = rogue.newExposeArmorTemplate(sim)
	rogue.ruptureTemplate = rogue.newRuptureTemplate(sim)
	rogue.deadlyPoisonTemplate = rogue.newDeadlyPoisonTemplate(sim)
	rogue.deadlyPoisonRefreshTemplate = rogue.newDeadlyPoisonRefreshTemplate(sim)
	rogue.instantPoisonTemplate = rogue.newInstantPoisonTemplate(sim)

	rogue.energyPerSecondAvg = core.EnergyPerTick / core.EnergyTickDuration.Seconds()

	// TODO: Currently assumes default combat spec.
	expectedComboPointsAfterFinisher := 0
	expectedEnergyAfterFinisher := 25.0
	comboPointsNeeded := 5 - expectedComboPointsAfterFinisher
	energyForEA := rogue.builderEnergyCost*float64(comboPointsNeeded) + ExposeArmorEnergyCost
	rogue.eaBuildTime = time.Duration(((energyForEA - expectedEnergyAfterFinisher) / rogue.energyPerSecondAvg) * float64(time.Second))
}

func (rogue *Rogue) Reset(sim *core.Simulation) {
	rogue.plan = PlanOpener
	rogue.deathmantle4pcProc = false
	rogue.deadlyPoisonStacks = 0
	rogue.doneSND = false

	permaEA := sim.GetPrimaryTarget().AuraExpiresAt(core.ExposeArmorDebuffID) == core.NeverExpires
	rogue.doneEA = !rogue.Rotation.MaintainExposeArmor || permaEA

	rogue.disabledMCDs = rogue.DisableAllEnabledCooldowns(core.CooldownTypeUnknown)
}

func (rogue *Rogue) critMultiplier(isMH bool, applyLethality bool) float64 {
	primaryModifier := 1.0
	secondaryModifier := 0.0

	isMace := false
	if weapon := rogue.Equip[proto.ItemSlot_ItemSlotMainHand]; isMH && weapon.ID != 0 {
		if weapon.WeaponType == proto.WeaponType_WeaponTypeMace {
			isMace = true
		}
	} else if weapon := rogue.Equip[proto.ItemSlot_ItemSlotOffHand]; !isMH && weapon.ID != 0 {
		if weapon.WeaponType == proto.WeaponType_WeaponTypeMace {
			isMace = true
		}
	}
	if isMace {
		primaryModifier *= 1 + 0.01*float64(rogue.Talents.MaceSpecialization)
	}

	if applyLethality {
		secondaryModifier += 0.06 * float64(rogue.Talents.Lethality)
	}

	return rogue.MeleeCritMultiplier(primaryModifier, secondaryModifier)
}

func NewRogue(character core.Character, options proto.Player) *Rogue {
	rogueOptions := options.GetRogue()

	rogue := &Rogue{
		Character: character,
		Talents:   *rogueOptions.Talents,
		Options:   *rogueOptions.Options,
		Rotation:  *rogueOptions.Rotation,
	}

	daggerMH := rogue.Equip[proto.ItemSlot_ItemSlotMainHand].WeaponType == proto.WeaponType_WeaponTypeDagger
	if rogue.Rotation.Builder == proto.Rogue_Rotation_Unknown {
		rogue.Rotation.Builder = proto.Rogue_Rotation_Auto
	}
	if rogue.Rotation.Builder == proto.Rogue_Rotation_Backstab && !daggerMH {
		rogue.Rotation.Builder = proto.Rogue_Rotation_Auto
	} else if rogue.Rotation.Builder == proto.Rogue_Rotation_Hemorrhage && !rogue.Talents.Hemorrhage {
		rogue.Rotation.Builder = proto.Rogue_Rotation_Auto
	} else if rogue.Rotation.Builder == proto.Rogue_Rotation_Mutilate && !rogue.Talents.Mutilate {
		rogue.Rotation.Builder = proto.Rogue_Rotation_Auto
	}
	if rogue.Rotation.Builder == proto.Rogue_Rotation_Auto {
		if rogue.Talents.Mutilate {
			rogue.Rotation.Builder = proto.Rogue_Rotation_Mutilate
		} else if rogue.Talents.Hemorrhage {
			rogue.Rotation.Builder = proto.Rogue_Rotation_Hemorrhage
		} else if daggerMH {
			rogue.Rotation.Builder = proto.Rogue_Rotation_Backstab
		} else {
			rogue.Rotation.Builder = proto.Rogue_Rotation_SinisterStrike
		}
	}

	var newBuilder func(sim *core.Simulation, target *core.Target) *core.SimpleSpell
	switch rogue.Rotation.Builder {
	case proto.Rogue_Rotation_SinisterStrike:
		rogue.builderEnergyCost = rogue.SinisterStrikeEnergyCost()
		newBuilder = func(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
			return rogue.NewSinisterStrike(sim, target)
		}
	case proto.Rogue_Rotation_Backstab:
		rogue.builderEnergyCost = BackstabEnergyCost
		newBuilder = func(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
			return rogue.NewBackstab(sim, target)
		}
	case proto.Rogue_Rotation_Hemorrhage:
		rogue.builderEnergyCost = HemorrhageEnergyCost
		newBuilder = func(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
			return rogue.NewHemorrhage(sim, target)
		}
	case proto.Rogue_Rotation_Mutilate:
		rogue.builderEnergyCost = MutilateEnergyCost
		newBuilder = func(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
			return rogue.NewMutilate(sim, target)
		}
	}

	if rogue.Rotation.UseShiv && rogue.Consumes.OffHandImbue == proto.WeaponImbue_WeaponImbueRogueDeadlyPoison {
		rogue.newBuilder = func(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
			if rogue.deadlyPoison.Effect.DotInput.IsTicking(sim) && rogue.deadlyPoison.Effect.DotInput.TimeRemaining(sim) < time.Second*2 && rogue.CurrentEnergy() >= rogue.shivEnergyCost {
				return rogue.NewShiv(sim, target)
			} else {
				return newBuilder(sim, target)
			}
		}
	} else {
		rogue.newBuilder = newBuilder
	}

	maxEnergy := 100.0
	if rogue.Talents.Vigor {
		maxEnergy = 110
	}
	rogue.EnableEnergyBar(maxEnergy, func(sim *core.Simulation) {
		if !rogue.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
			rogue.doRotation(sim)
		}
	})
	rogue.EnableAutoAttacks(rogue, core.AutoAttackOptions{
		MainHand:       rogue.WeaponFromMainHand(rogue.critMultiplier(true, false)),
		OffHand:        rogue.WeaponFromOffHand(rogue.critMultiplier(false, false)),
		AutoSwingMelee: true,
	})

	rogue.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Strength,
		ModifiedStat: stats.AttackPower,
		Modifier: func(strength float64, attackPower float64) float64 {
			return attackPower + strength*1
		},
	})

	rogue.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.AttackPower,
		Modifier: func(agility float64, attackPower float64) float64 {
			return attackPower + agility*1
		},
	})

	rogue.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.MeleeCrit,
		Modifier: func(agility float64, meleeCrit float64) float64 {
			return meleeCrit + (agility/40)*core.MeleeCritRatingPerCritChance
		},
	})

	rogue.applyTalents()
	rogue.registerThistleTeaCD()

	return rogue
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceBloodElf, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Strength:  92,
		stats.Agility:   160,
		stats.Stamina:   88,
		stats.Intellect: 43,
		stats.Spirit:    57,

		stats.AttackPower: 120,
		stats.MeleeCrit:   -0.3 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDwarf, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Strength:  97,
		stats.Agility:   154,
		stats.Stamina:   92,
		stats.Intellect: 38,
		stats.Spirit:    57,

		stats.AttackPower: 120,
		stats.MeleeCrit:   -0.3 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceGnome, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Strength:  90,
		stats.Agility:   161,
		stats.Stamina:   88,
		stats.Intellect: 45,
		stats.Spirit:    58,

		stats.AttackPower: 120,
		stats.MeleeCrit:   -0.3 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceHuman, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Strength:  95,
		stats.Agility:   158,
		stats.Stamina:   89,
		stats.Intellect: 39,
		stats.Spirit:    58,

		stats.AttackPower: 120,
		stats.MeleeCrit:   -0.3 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceNightElf, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Strength:  92,
		stats.Agility:   163,
		stats.Stamina:   88,
		stats.Intellect: 39,
		stats.Spirit:    58,

		stats.AttackPower: 120,
		stats.MeleeCrit:   -0.3 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceOrc, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Strength:  98,
		stats.Agility:   155,
		stats.Stamina:   91,
		stats.Intellect: 36,
		stats.Spirit:    61,

		stats.AttackPower: 120,
		stats.MeleeCrit:   -0.3 * core.MeleeCritRatingPerCritChance,
	}
	trollStats := stats.Stats{
		stats.Strength:  96,
		stats.Agility:   160,
		stats.Stamina:   90,
		stats.Intellect: 35,
		stats.Spirit:    59,

		stats.AttackPower: 120,
		stats.MeleeCrit:   -0.3 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll10, Class: proto.Class_ClassRogue}] = trollStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll30, Class: proto.Class_ClassRogue}] = trollStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceUndead, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Strength:  94,
		stats.Agility:   156,
		stats.Stamina:   90,
		stats.Intellect: 37,
		stats.Spirit:    63,

		stats.AttackPower: 120,
		stats.MeleeCrit:   -0.3 * core.MeleeCritRatingPerCritChance,
	}
}

// Agent is a generic way to access underlying rogue on any of the agents.
type RogueAgent interface {
	GetRogue() *Rogue
}
