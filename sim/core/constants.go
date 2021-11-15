package core

import (
	"time"
)

const GCDMin = time.Second * 1
const GCDDefault = time.Millisecond * 1500

const HasteRatingPerHastePercent = 15.76
const MeleeCritRatingPerCritChance = 22.1
const MeleeHitRatingPerHitChance = 15.8
const SpellCritRatingPerCritChance = 22.08
const SpellHitRatingPerHitChance = 12.6

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
