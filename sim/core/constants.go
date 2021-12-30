package core

import (
	"time"
)

const GCDMin = time.Second * 1
const GCDDefault = time.Millisecond * 1500

const HasteRatingPerHastePercent = 15.77

const MeleeCritRatingPerCritChance = 22.08
const MeleeHitRatingPerHitChance = 15.77
const MeleeAttackRatingPerDamage = 14

const ExpertisePerQuarterPercentReduction = 3.94
const ArmorPenPerPercentArmor = 5.92

const SpellCritRatingPerCritChance = 22.08
const SpellHitRatingPerHitChance = 12.62

// IDs for items used in core
const (
	ItemIDAtieshMage            = 22589
	ItemIDAtieshWarlock         = 22630
	ItemIDBraidedEterniumChain  = 24114
	ItemIDChainOfTheTwilightOwl = 24121
	ItemIDEyeOfTheNight         = 24116
	ItemIDJadePendantOfBlasting = 20966
	ItemIDTheLightningCapacitor = 28785
)
