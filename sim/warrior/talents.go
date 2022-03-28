package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (warrior *Warrior) ApplyTalents() {
	warrior.registerBloodrageCD()

	warrior.AddStat(stats.Parry, core.ParryRatingPerParryChance*1*float64(warrior.Talents.Deflection))
	warrior.AddStat(stats.MeleeCrit, core.MeleeCritRatingPerCritChance*1*float64(warrior.Talents.Cruelty))
	// TODO WeaponMastery (reduces dodge only)
	warrior.AddStat(stats.MeleeHit, core.MeleeHitRatingPerHitChance*1*float64(warrior.Talents.Precision))
	warrior.AddStat(stats.Defense, core.DefenseRatingPerDefense*4*float64(warrior.Talents.Anticipation))
	warrior.AddStat(stats.Block, core.BlockRatingPerBlockChance*1*float64(warrior.Talents.ShieldSpecialization))
	warrior.AddStat(stats.Armor, warrior.Equip.Stats()[stats.Armor]*0.02*float64(warrior.Talents.Toughness))
	warrior.AddStat(stats.Expertise, core.ExpertisePerQuarterPercentReduction*2*float64(warrior.Talents.Defiance))

	if warrior.Talents.ImprovedBerserkerStance > 0 {
		bonus := 1 + 0.02*float64(warrior.Talents.ImprovedBerserkerStance)
		warrior.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.AttackPower,
			ModifiedStat: stats.AttackPower,
			Modifier: func(ap float64, _ float64) float64 {
				return ap * bonus
			},
		})
	}

	if warrior.Talents.ShieldMastery > 0 {
		bonus := 1 + 0.1*float64(warrior.Talents.ShieldMastery)
		warrior.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.BlockValue,
			ModifiedStat: stats.BlockValue,
			Modifier: func(bv float64, _ float64) float64 {
				return bv * bonus
			},
		})
	}

	if warrior.Talents.Vitality > 0 {
		stamBonus := 1 + 0.01*float64(warrior.Talents.Vitality)
		warrior.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Stamina,
			ModifiedStat: stats.Stamina,
			Modifier: func(stamina float64, _ float64) float64 {
				return stamina * stamBonus
			},
		})
		strBonus := 1 + 0.02*float64(warrior.Talents.Vitality)
		warrior.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Strength,
			ModifiedStat: stats.Strength,
			Modifier: func(strength float64, _ float64) float64 {
				return strength * strBonus
			},
		})
	}

	warrior.applyOneHandedWeaponSpecialization()
}

// Unlike most other applyXXX() talent appliers, this is meant to be called from Reset().
func (warrior *Warrior) applyAngerManagement(sim *core.Simulation) {
	if !warrior.Talents.AngerManagement {
		return
	}

	core.StartPeriodicAction(sim, core.PeriodicActionOptions{
		Period: time.Second * 3,
		OnAction: func(sim *core.Simulation) {
			warrior.AddRage(sim, 1, core.ActionID{SpellID: 12296})
		},
	})
}

func (warrior *Warrior) applyOneHandedWeaponSpecialization() {
	if warrior.Talents.OneHandedWeaponSpecialization == 0 {
		return
	}
	if warrior.Equip[proto.ItemSlot_ItemSlotMainHand].HandType == proto.HandType_HandTypeTwoHand {
		return
	}

	warrior.PseudoStats.PhysicalDamageDealtMultiplier *= 1 + 0.02*float64(warrior.Talents.OneHandedWeaponSpecialization)
}
