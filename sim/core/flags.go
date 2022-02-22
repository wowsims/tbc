package core

import "github.com/wowsims/tbc/sim/core/stats"

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

	// These bits are set by the hit roll
	OutcomeMiss HitOutcome = 1 << iota
	OutcomeHit
	OutcomeDodge
	OutcomeGlance
	OutcomeParry
	OutcomeBlock

	// These bits are set by the crit and damage rolls.
	OutcomeCrit
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

// Other flags
// Ignore Resistance (armor or magical, use school)
// Always Hits

type SpellExtras byte

// Returns whether there is any overlap between the given masks.
func (se SpellExtras) Matches(other SpellExtras) bool {
	return (se & other) != 0
}

const (
	SpellExtrasNone          SpellExtras = 0
	SpellExtrasIgnoreResists SpellExtras = 1 << iota // skip spell resist/armor
	SpelLExtrasIgnoreDodge                           // Ignores dodge in physical hit rolls
	SpellExtrasAlwaysHits                            // Can't miss the hit roll
	SpellExtrasBinary                                // Does not do partial resists and could need a different hit roll.
	SpellExtrasChanneled                             // Spell is channeled
)

// OutcomeRollCategory is the mask for what kind of hit roll to perform
type OutcomeRollCategory byte

// Returns whether there is any overlap between the given masks.
func (at OutcomeRollCategory) Matches(other OutcomeRollCategory) bool {
	return (at & other) != 0
}

const (
	OutcomeRollCategoryNone    OutcomeRollCategory = 0         // No outcome roll needed
	OutcomeRollCategoryWhite   OutcomeRollCategory = 1 << iota // White hit roll rules
	OutcomeRollCategorySpecial                                 // Melee special attack
	OutcomeRollCategoryRanged                                  // Ranged attack roll
	OutcomeRollCategoryMagic                                   // Spell Hit roll
)

type CritRollCategory byte

// Returns whether there is any overlap between the given masks.
func (at CritRollCategory) Matches(other CritRollCategory) bool {
	return (at & other) != 0
}

const (
	CritRollCategoryNone     CritRollCategory = 0         // cannot crit
	CritRollCategoryPhysical CritRollCategory = 1 << iota // uses MeleeCrit for roll
	CritRollCategoryMagical                               // Uses SpellCrit for crit roll
)

type SpellSchool byte

func (ss SpellSchool) Stat() stats.Stat {
	switch ss {
	case SpellSchoolArcane:
		return stats.ArcaneSpellPower
	case SpellSchoolFire:
		return stats.FireSpellPower
	case SpellSchoolFrost:
		return stats.FrostSpellPower
	case SpellSchoolHoly:
		return stats.HolySpellPower
	case SpellSchoolNature:
		return stats.NatureSpellPower
	case SpellSchoolShadow:
		return stats.ShadowSpellPower
	case SpellSchoolPhysical:
		return stats.AttackPower
	}

	return 0
}

// Returns whether there is any overlap between the given masks.
func (ss SpellSchool) Matches(other SpellSchool) bool {
	return (ss & other) != 0
}

const (
	SpellSchoolNone     SpellSchool = 0
	SpellSchoolPhysical SpellSchool = 1 << iota
	SpellSchoolArcane
	SpellSchoolFire
	SpellSchoolFrost
	SpellSchoolHoly
	SpellSchoolNature
	SpellSchoolShadow

	SpellSchoolMagic = SpellSchoolArcane | SpellSchoolFire | SpellSchoolFrost | SpellSchoolHoly | SpellSchoolNature | SpellSchoolShadow
)

/*
outcome roll hit/miss/crit/glance (assigns Outcome mask) -> If Hit, Crit Roll -> damage (applies metrics) -> trigger proc

So in TBC it looks like they just gave it the cannot miss flag even though they also switched its defense type to physical (??)
the damage type is holy, which ignores armor and as it is magic so it can be partially resisted (due to level resistance).
however it also gains the physical bit mask as I explain in a post above

so there is no hit roll, there is a melee crit roll, a spell damage roll, and melee "on hit"

ok so I did some more testing on this.
Judgement of Blood correctly gets the "always hit" (aka cannot miss flag applied to it) --
its only mitigation events are partial resists at the correct rates
however Judgement of Command is broken. even though it has the "always hit" flag it seems to
be ignored because it is procced by an intermediary dummy spell which does not have the "cannot miss" flag applied to it lmao.
for some god forsaken reason Judgement of Command is ALSO a dummy which then casts the correct Judgement of Command
which deals damage, and this dummy can miss, lmao
I got ~16.4% resists in about almost 96 casts which suggests it uses the spell hit check,
which makes sense because its defensetype is set to 1, Magic

arcane shot - ranged hit, spell dmg, procs special ranged
	OutcomeRollRanged, School Arcane, ProcMask - RangedSpecial

judgement of blood - physical hit/crit, spell damage, "cannot miss", procs special melee and ranged
	Damage is (weapon damage + spell power)*0.7*(bonus holy damage against target)+flat bonus damage
	OtherFlagCannotMiss, OutcomeRollSpecial, School Holy (base damage = weapon damage range), Multiplier 70%

judgement of command - spell hit, melee crit, spell damage, procs special melee and ranged
	OutcomeRollSpell, School Holy


moonfire - spell hit, spell dmg, dot dmg, procs spell hit
stormstrike - melee hit, melee dmg, procs special melee
rupture -

wotlk
shadowflame - requires each 'effect' to have its own school.
*/
