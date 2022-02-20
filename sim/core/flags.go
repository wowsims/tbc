package core

type ProcMask uint32

// Returns whether there is any overlap between the given masks.
func (pm ProcMask) Matches(other ProcMask) bool {
	return (pm & other) != 0
}

// Actual Blizzard flag values:
// 1  1        = Triggered by script
// 2  2        = Triggers on kill
// 3  4        = Melee auto attack
// 4  8        = On take melee auto attack
// 5  16       = Melee special attack / melee damage
// 6  32       = On take melee special attack
// 7  64       = Ranged auto attack
// 8  128      = On take ranged auto attack
// 9  256      = Ranged special attack / ranged damage
// 10 512      = On take ranged special attack
// 11 1024     = ???? On use combo points? Shapeshift? Change stance? Gain buff? Some rogue stuff
// 12 2048     = ???? Rogue related? Script?
// 13 4096     = ???? Stealth related? Script? On gain/lose stealth? Also possibly on stance change
// 14 8192     = On spell hit on you
// 15 16384    = Cast heal
// 16 32768    = On get healed
// 17 65536    = Deal spell damage
// 18 131072   = On take spell damage
// 19 262144   = Deal periodic damage
// 20 524288   = On take periodic damage
// 21 1048576  = On take any damage
// 22 2097152  = On Apply debuff
// 23 4194304  = ???? On have debuff applied to you? really bizarre mask
// 24 8388608  = On offhand attack
// 25 16777216 = What the fuck?

// Single-bit masks. These don't need to match Blizzard's values.
const (
	ProcMaskEmpty       ProcMask = 0
	ProcMaskMeleeMHAuto ProcMask = 1 << iota
	ProcMaskMeleeOHAuto
	ProcMaskMeleeMHSpecial
	ProcMaskMeleeOHSpecial
	ProcMaskRangedAuto
	ProcMaskRangedSpecial
	ProcMaskSpellDamage
	ProcMaskPeriodicDamage
)

const (
	ProcMaskMeleeMH = ProcMaskMeleeMHAuto | ProcMaskMeleeMHSpecial
	ProcMaskMeleeOH = ProcMaskMeleeOHAuto | ProcMaskMeleeOHSpecial
	// Equivalent to in-game mask of 4.
	ProcMaskMeleeWhiteHit = ProcMaskMeleeMHAuto | ProcMaskMeleeOHAuto
	// Equivalent to in-game mask of 68.
	ProcMaskWhiteHit = ProcMaskMeleeMHAuto | ProcMaskMeleeOHAuto | ProcMaskRangedAuto
	// Equivalent to in-game mask of 16.
	ProcMaskMeleeSpecial = ProcMaskMeleeMHSpecial | ProcMaskMeleeOHSpecial
	// Equivalent to in-game mask of 272.
	ProcMaskMeleeOrRangedSpecial = ProcMaskMeleeSpecial | ProcMaskRangedSpecial
	// Equivalent to in-game mask of 20.
	ProcMaskMelee = ProcMaskMeleeWhiteHit | ProcMaskMeleeSpecial
	// Equivalent to in-game mask of 320.
	ProcMaskRanged = ProcMaskRangedAuto | ProcMaskRangedSpecial
	// Equivalent to in-game mask of 340.
	ProcMaskMeleeOrRanged = ProcMaskMelee | ProcMaskRanged

	ProcMaskTwoRoll = ProcMaskRanged | ProcMaskMeleeSpecial
)

func GetMeleeProcMaskForHands(mh bool, oh bool) ProcMask {
	mask := ProcMaskEmpty
	if mh {
		mask |= ProcMaskMeleeMH
	}
	if oh {
		mask |= ProcMaskMeleeOH
	}
	return mask
}

// Possible outcomes of any hit/damage roll.
type HitOutcome uint16

// Returns whether there is any overlap between the given masks.
func (ho HitOutcome) Matches(other HitOutcome) bool {
	return (ho & other) != 0
}

// Single-bit outcomes.
const (
	OutcomeEmpty HitOutcome = 0
	OutcomeMiss  HitOutcome = 1 << iota
	OutcomeHit
	OutcomeCrit
	OutcomeDodge
	OutcomeGlance
	OutcomeParry
	OutcomeBlock
	OutcomePartial1_4 // 1/4 of the spell was resisted.
	OutcomePartial2_4 // 2/4 of the spell was resisted.
	OutcomePartial3_4 // 3/4 of the spell was resisted.
)

const (
	OutcomePartial = OutcomePartial1_4 | OutcomePartial2_4 | OutcomePartial3_4
	OutcomeLanded  = OutcomeHit | OutcomeCrit | OutcomeGlance | OutcomeBlock | OutcomePartial
)

func (ho HitOutcome) String() string {
	if ho.Matches(OutcomeMiss) {
		return "Miss"
	} else if ho.Matches(OutcomeDodge) {
		return "Dodge"
	} else if ho.Matches(OutcomeParry) {
		return "Parry"
	} else if ho.Matches(OutcomeGlance) {
		return "Glance"
	} else if ho.Matches(OutcomeBlock) {
		return "Block"
	} else if ho.Matches(OutcomeCrit) {
		return "Crit" + ho.PartialResistString()
	} else if ho.Matches(OutcomeHit) {
		return "Hit" + ho.PartialResistString()
	} else {
		return "Empty"
	}
}

func (ho HitOutcome) PartialResistString() string {
	if ho.Matches(OutcomePartial1_4) {
		return " (25% Resist)"
	} else if ho.Matches(OutcomePartial2_4) {
		return " (50% Resist)"
	} else if ho.Matches(OutcomePartial3_4) {
		return " (75% Resist)"
	} else {
		return ""
	}
}
