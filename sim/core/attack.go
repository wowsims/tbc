package core

import (
	"fmt"
	"time"

	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

// ReplaceMHSwing is called right before an auto attack fires
//  If it returns nil, the attack takes place as normal. If it returns an ability,
//  that ability is used in place of the attack.
//  This allows for abilities that convert a white attack into yellow attack.
type ReplaceMHSwing func(sim *Simulation) *Spell

// Represents a generic weapon. Pets / unarmed / various other cases dont use
// actual weapon items so this is an abstraction of a Weapon.
type Weapon struct {
	BaseDamageMin        float64
	BaseDamageMax        float64
	SwingSpeed           float64
	NormalizedSwingSpeed float64
	SwingDuration        time.Duration // Duration between 2 swings.
	CritMultiplier       float64
}

func newWeaponFromUnarmed(critMultiplier float64) Weapon {
	// These numbers are probably wrong but nobody cares.
	return Weapon{
		BaseDamageMin:        0,
		BaseDamageMax:        0,
		SwingSpeed:           1,
		NormalizedSwingSpeed: 1,
		SwingDuration:        time.Second,
		CritMultiplier:       critMultiplier,
	}
}

func newWeaponFromItem(item items.Item, critMultiplier float64) Weapon {
	normalizedWeaponSpeed := 2.4
	if item.WeaponType == proto.WeaponType_WeaponTypeDagger {
		normalizedWeaponSpeed = 1.7
	} else if item.HandType == proto.HandType_HandTypeTwoHand {
		normalizedWeaponSpeed = 3.3
	} else if item.RangedWeaponType != proto.RangedWeaponType_RangedWeaponTypeUnknown {
		normalizedWeaponSpeed = 2.8
	}

	return Weapon{
		BaseDamageMin:        item.WeaponDamageMin,
		BaseDamageMax:        item.WeaponDamageMax,
		SwingSpeed:           item.SwingSpeed,
		NormalizedSwingSpeed: normalizedWeaponSpeed,
		SwingDuration:        time.Duration(item.SwingSpeed * float64(time.Second)),
		CritMultiplier:       critMultiplier,
	}
}

// Returns weapon stats using the main hand equipped weapon.
func (character *Character) WeaponFromMainHand(critMultiplier float64) Weapon {
	if weapon := character.GetMHWeapon(); weapon != nil {
		return newWeaponFromItem(*weapon, critMultiplier)
	} else {
		return newWeaponFromUnarmed(critMultiplier)
	}
}

// Returns weapon stats using the off hand equipped weapon.
func (character *Character) WeaponFromOffHand(critMultiplier float64) Weapon {
	if weapon := character.GetOHWeapon(); weapon != nil {
		return newWeaponFromItem(*weapon, critMultiplier)
	} else {
		return Weapon{}
	}
}

// Returns weapon stats using the off hand equipped weapon.
func (character *Character) WeaponFromRanged(critMultiplier float64) Weapon {
	if weapon := character.GetRangedWeapon(); weapon != nil {
		return newWeaponFromItem(*weapon, critMultiplier)
	} else {
		return Weapon{}
	}
}

func (weapon Weapon) BaseDamage(sim *Simulation) float64 {
	return weapon.BaseDamageMin + (weapon.BaseDamageMax-weapon.BaseDamageMin)*sim.RandomFloat("Weapon Base Damage")
}

func (weapon Weapon) calculateWeaponDamage(sim *Simulation, attackPower float64) float64 {
	return weapon.BaseDamage(sim) + (weapon.SwingSpeed*attackPower)/MeleeAttackRatingPerDamage
}

func (weapon Weapon) calculateNormalizedWeaponDamage(sim *Simulation, attackPower float64) float64 {
	return weapon.BaseDamage(sim) + (weapon.NormalizedSwingSpeed*attackPower)/MeleeAttackRatingPerDamage
}

type MeleeDamageCalculator func(attackPower float64, bonusWeaponDamage float64) float64

// Returns whether this hit effect is associated with the main-hand weapon.
func (ahe *SpellEffect) IsMH() bool {
	const mhmask = ProcMaskMeleeMH
	return ahe.ProcMask.Matches(mhmask)
}

// Returns whether this hit effect is associated with the off-hand weapon.
func (ahe *SpellEffect) IsOH() bool {
	return ahe.ProcMask.Matches(ProcMaskMeleeOH)
}

// Returns whether this hit effect is associated with either melee weapon.
func (ahe *SpellEffect) IsMelee() bool {
	return ahe.ProcMask.Matches(ProcMaskMelee)
}

type AutoAttacks struct {
	// initialized
	agent     Agent
	character *Character
	MH        Weapon
	OH        Weapon
	Ranged    Weapon

	IsDualWielding bool

	// If true, core engine will handle calling SwingMelee(). Set to false to manually manage
	// swings, for example for hunter melee weaving.
	AutoSwingMelee bool

	// If true, core engine will handle calling SwingRanged(). Unless you're a hunter, don't
	// use this.
	AutoSwingRanged bool

	// Set this to true to use the OH delay macro, mostly used by enhance shamans.
	// This will intentionally delay OH swings to that they always fall within the
	// 0.5s window following a MH swing.
	DelayOHSwings bool

	MainhandSwingAt time.Duration
	OffhandSwingAt  time.Duration
	RangedSwingAt   time.Duration

	MHEffect     SpellEffect
	OHEffect     SpellEffect
	RangedEffect SpellEffect

	MHAuto     *Spell
	OHAuto     *Spell
	RangedAuto *Spell

	RangedSwingInProgress bool

	ReplaceMHSwing ReplaceMHSwing

	// The time at which the last MH swing occurred.
	previousMHSwingAt time.Duration

	// PendingAction which handles auto attacks.
	autoSwingAction    *PendingAction
	autoSwingCancelled bool
}

// Options for initializing auto attacks.
type AutoAttackOptions struct {
	MainHand       Weapon
	OffHand        Weapon
	Ranged         Weapon
	AutoSwingMelee bool // If true, core engine will handle calling SwingMelee() for you.
	DelayOHSwings  bool
	ReplaceMHSwing ReplaceMHSwing
}

func (character *Character) EnableAutoAttacks(agent Agent, options AutoAttackOptions) {
	character.AutoAttacks = AutoAttacks{
		agent:          agent,
		character:      character,
		MH:             options.MainHand,
		OH:             options.OffHand,
		Ranged:         options.Ranged,
		AutoSwingMelee: options.AutoSwingMelee,
		DelayOHSwings:  options.DelayOHSwings,
		ReplaceMHSwing: options.ReplaceMHSwing,
		IsDualWielding: options.MainHand.SwingSpeed != 0 && options.OffHand.SwingSpeed != 0,
		MHEffect: SpellEffect{
			ProcMask:         ProcMaskMeleeMHAuto,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       BaseDamageConfigMeleeWeapon(MainHand, false, 0, 1, true),
			OutcomeApplier:   OutcomeFuncMeleeWhite(options.MainHand.CritMultiplier),
		},
		OHEffect: SpellEffect{
			ProcMask:         ProcMaskMeleeOHAuto,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       BaseDamageConfigMeleeWeapon(OffHand, false, 0, 1, true),
			OutcomeApplier:   OutcomeFuncMeleeWhite(options.OffHand.CritMultiplier),
		},
		RangedEffect: SpellEffect{
			ProcMask:         ProcMaskRangedAuto,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       BaseDamageConfigRangedWeapon(0),
			OutcomeApplier:   OutcomeFuncRangedHitAndCrit(options.Ranged.CritMultiplier),
		},
	}
}

func (aa *AutoAttacks) IsEnabled() bool {
	return aa.MH.SwingSpeed != 0
}

// Empty handler so Agents don't have to provide one if they have no logic to add.
func (character *Character) OnAutoAttack(sim *Simulation, spell *Spell) {}

func (aa *AutoAttacks) reset(sim *Simulation) {
	if !aa.IsEnabled() {
		return
	}

	aa.MHAuto = aa.character.GetOrRegisterSpell(SpellConfig{
		ActionID:    ActionID{OtherID: proto.OtherAction_OtherActionAttack, Tag: 1},
		SpellSchool: SpellSchoolPhysical,
		SpellExtras: SpellExtrasMeleeMetrics,

		ApplyEffects: ApplyEffectFuncDirectDamage(aa.MHEffect),
	})

	aa.OHAuto = aa.character.GetOrRegisterSpell(SpellConfig{
		ActionID:    ActionID{OtherID: proto.OtherAction_OtherActionAttack, Tag: 2},
		SpellSchool: SpellSchoolPhysical,
		SpellExtras: SpellExtrasMeleeMetrics,

		ApplyEffects: ApplyEffectFuncDirectDamage(aa.OHEffect),
	})

	aa.RangedAuto = aa.character.GetOrRegisterSpell(SpellConfig{
		ActionID:    ActionID{OtherID: proto.OtherAction_OtherActionShoot},
		SpellSchool: SpellSchoolPhysical,
		SpellExtras: SpellExtrasMeleeMetrics,

		Cast: CastConfig{
			DefaultCast: Cast{
				CastTime: 1, // Dummy non-zero value so the optimization doesnt remove the cast time.
			},
			ModifyCast: func(_ *Simulation, _ *Spell, cast *Cast) {
				cast.CastTime = aa.RangedSwingWindup()
			},
			IgnoreHaste: true,
			AfterCast: func(sim *Simulation, spell *Spell) {
				aa.RangedSwingInProgress = false
				aa.agent.OnAutoAttack(sim, aa.RangedAuto)
			},
		},

		ApplyEffects: ApplyEffectFuncDirectDamage(aa.RangedEffect),
	})

	aa.MainhandSwingAt = 0
	aa.OffhandSwingAt = 0

	// Apply random delay of 0 - 50% swing time, to one of the weapons if dual wielding
	if aa.IsDualWielding {
		// Set a fake value for previousMHSwing so that offhand swing delay works
		// properly at the start.
		aa.previousMHSwingAt = time.Second * -1

		delay := time.Duration(sim.RandomFloat("SwingResetDelay") * float64(aa.MH.SwingDuration/2))
		isMHDelay := sim.RandomFloat("SwingResetWeapon") < 0.5

		if isMHDelay {
			aa.MainhandSwingAt = delay
		} else {
			aa.OffhandSwingAt = delay
		}
	}

	aa.autoSwingAction = nil
	aa.autoSwingCancelled = false
	aa.resetAutoSwing(sim)

	aa.RangedSwingAt = 0
	aa.RangedSwingInProgress = false
}

func (aa *AutoAttacks) resetAutoSwing(sim *Simulation) {
	if aa.autoSwingCancelled || !aa.AutoSwingMelee {
		return
	}

	if aa.autoSwingAction != nil {
		aa.autoSwingAction.Cancel(sim)
	}

	pa := sim.pendingActionPool.Get()
	pa.Priority = ActionPriorityAuto

	pa.OnAction = func(sim *Simulation) {
		aa.SwingMelee(sim, sim.GetPrimaryTarget())
		pa.NextActionAt = aa.NextAttackAt()

		// Cancelled means we made a new one because of a swing speed change.
		if !pa.cancelled {
			sim.AddPendingAction(pa)
		} else {
			sim.pendingActionPool.Put(pa)
		}
	}
	pa.NextActionAt = aa.NextAttackAt()

	aa.autoSwingAction = pa
	sim.AddPendingAction(pa)
}

// Stops the auto swing action for the rest of the iteration. Used for pets
// after being disabled.
func (aa *AutoAttacks) CancelAutoSwing(sim *Simulation) {
	if aa.autoSwingAction != nil {
		aa.autoSwingAction.Cancel(sim)
		aa.autoSwingAction = nil
		aa.autoSwingCancelled = true
	}
}

// Renables the auto swing action for the iteration
func (aa *AutoAttacks) EnableAutoSwing(sim *Simulation) {
	// Already enabled so nothing to do
	if aa.autoSwingAction != nil {
		return
	}

	if aa.MainhandSwingAt < sim.CurrentTime {
		aa.MainhandSwingAt = sim.CurrentTime
	}
	if aa.OffhandSwingAt < sim.CurrentTime {
		aa.OffhandSwingAt = sim.CurrentTime
	}
	if aa.RangedSwingAt < sim.CurrentTime {
		if aa.RangedSwingInProgress {
			panic("Ranged swing already in progress!")
		}
		aa.RangedSwingAt = sim.CurrentTime
	}

	aa.autoSwingCancelled = false
	aa.resetAutoSwing(sim)
}

// The amount of time between two MH swings.
func (aa *AutoAttacks) MainhandSwingSpeed() time.Duration {
	return time.Duration(float64(aa.MH.SwingDuration) / aa.character.SwingSpeed())
}

// The amount of time between two OH swings.
func (aa *AutoAttacks) OffhandSwingSpeed() time.Duration {
	return time.Duration(float64(aa.OH.SwingDuration) / aa.character.SwingSpeed())
}

// The amount of time between two ranged swings.
func (aa *AutoAttacks) RangedSwingSpeed() time.Duration {
	return time.Duration(float64(aa.Ranged.SwingDuration) / aa.character.RangedSwingSpeed())
}

// Ranged swings have a 0.5s 'windup' time before they can fire, affected by haste.
// This function computes the amount of windup time based on the current haste.
func (aa *AutoAttacks) RangedSwingWindup() time.Duration {
	return time.Duration(float64(time.Millisecond*500) / aa.character.RangedSwingSpeed())
}

// Time between a ranged auto finishes casting and the next one becomes available.
func (aa *AutoAttacks) RangedSwingGap() time.Duration {
	return time.Duration(float64(aa.Ranged.SwingDuration-time.Millisecond*500) / aa.character.RangedSwingSpeed())
}

// Returns the amount of time available before ranged auto will be clipped.
func (aa *AutoAttacks) TimeBeforeClippingRanged(sim *Simulation) time.Duration {
	return aa.RangedSwingAt - aa.RangedSwingWindup() - sim.CurrentTime
}

// SwingMelee will check any swing timers if they are up, and if so, swing!
func (aa *AutoAttacks) SwingMelee(sim *Simulation, target *Target) {
	aa.TrySwingMH(sim, target)
	aa.TrySwingOH(sim, target)
}

func (aa *AutoAttacks) SwingRanged(sim *Simulation, target *Target) {
	aa.TrySwingRanged(sim, target)
}

// Performs an autoattack using the main hand weapon, if the MH CD is ready.
func (aa *AutoAttacks) TrySwingMH(sim *Simulation, target *Target) {
	if aa.MainhandSwingAt > sim.CurrentTime {
		return
	}

	attackSpell := aa.MHAuto
	if aa.ReplaceMHSwing != nil {
		// Allow MH swing to be overridden for abilities like Heroic Strike.
		replacementSpell := aa.ReplaceMHSwing(sim)
		if replacementSpell != nil {
			attackSpell = replacementSpell
		}
	}

	attackSpell.Cast(sim, target)
	aa.MainhandSwingAt = sim.CurrentTime + aa.MainhandSwingSpeed()
	aa.previousMHSwingAt = sim.CurrentTime
	aa.agent.OnAutoAttack(sim, attackSpell)
}

// Performs an autoattack using the main hand weapon, if the OH CD is ready.
func (aa *AutoAttacks) TrySwingOH(sim *Simulation, target *Target) {
	if !aa.IsDualWielding || aa.OffhandSwingAt > sim.CurrentTime {
		return
	}

	if aa.DelayOHSwings && (sim.CurrentTime-aa.previousMHSwingAt) > time.Millisecond*500 {
		// Delay the OH swing for later, so it follows the MH swing.
		aa.OffhandSwingAt = aa.MainhandSwingAt + time.Millisecond*100
		if sim.Log != nil {
			aa.character.Log(sim, "Delaying OH swing by %s", aa.OffhandSwingAt-sim.CurrentTime)
		}
		return
	}

	aa.OHAuto.Cast(sim, target)
	aa.OffhandSwingAt = sim.CurrentTime + aa.OffhandSwingSpeed()
	aa.agent.OnAutoAttack(sim, aa.OHAuto)
}

// Performs an autoattack using the ranged weapon, if the ranged CD is ready.
func (aa *AutoAttacks) TrySwingRanged(sim *Simulation, target *Target) {
	if aa.RangedSwingAt > sim.CurrentTime {
		return
	}

	aa.RangedAuto.Cast(sim, target)
	aa.RangedSwingAt = sim.CurrentTime + aa.RangedSwingSpeed()
	aa.RangedSwingInProgress = true

	// It's important that we update the GCD timer AFTER starting the ranged auto.
	// Otherwise the hardcast action won't be created separately.
	nextGCD := sim.CurrentTime + aa.RangedAuto.CurCast.CastTime
	if nextGCD > aa.character.NextGCDAt() {
		aa.character.SetGCDTimer(sim, nextGCD)
	}
}

func (aa *AutoAttacks) ModifySwingTime(sim *Simulation, amount float64) {
	if !aa.IsEnabled() {
		return
	}

	mhSwingTime := aa.MainhandSwingAt - sim.CurrentTime
	if mhSwingTime > 1 { // If its 1 we end up rounding down to 0 and causing a panic.
		aa.MainhandSwingAt = sim.CurrentTime + time.Duration(float64(mhSwingTime)/amount)
	}

	if aa.OH.SwingSpeed != 0 {
		ohSwingTime := aa.OffhandSwingAt - sim.CurrentTime
		if ohSwingTime > 1 {
			newTime := time.Duration(float64(ohSwingTime) / amount)
			if newTime > 0 {
				aa.OffhandSwingAt = sim.CurrentTime + newTime
			}
		}
	}

	aa.resetAutoSwing(sim)
}

// Delays all swing timers until the specified time.
func (aa *AutoAttacks) DelayAllUntil(sim *Simulation, readyAt time.Duration) {
	autoChanged := false

	if readyAt > aa.MainhandSwingAt {
		aa.MainhandSwingAt = readyAt
		if aa.AutoSwingMelee {
			autoChanged = true
		}
	}
	if readyAt > aa.OffhandSwingAt {
		aa.OffhandSwingAt = readyAt
		if aa.AutoSwingMelee {
			autoChanged = true
		}
	}
	if readyAt > aa.RangedSwingAt {
		if aa.RangedSwingInProgress {
			panic("Ranged swing already in progress!")
		}
		aa.RangedSwingAt = readyAt
	}

	if autoChanged {
		aa.resetAutoSwing(sim)
	}
}

func (aa *AutoAttacks) DelayRangedUntil(sim *Simulation, readyAt time.Duration) {
	if aa.RangedSwingInProgress {
		panic("Ranged swing already in progress!")
	}
	if readyAt > aa.RangedSwingAt {
		aa.RangedSwingAt = readyAt
	}
}

// Returns the time at which the next attack will occur.
func (aa *AutoAttacks) NextAttackAt() time.Duration {
	nextAttack := aa.MainhandSwingAt
	if aa.OH.SwingSpeed != 0 {
		nextAttack = MinDuration(nextAttack, aa.OffhandSwingAt)
	}
	return nextAttack
}

// Returns the time at which all melee swings will be ready.
func (aa *AutoAttacks) MeleeSwingsReadyAt() time.Duration {
	return MaxDuration(aa.MainhandSwingAt, aa.OffhandSwingAt)
}

// Returns true if all melee weapons are ready for a swing.
func (aa *AutoAttacks) MeleeSwingsReady(sim *Simulation) bool {
	return aa.MainhandSwingAt <= sim.CurrentTime &&
		(aa.OH.SwingSpeed == 0 || aa.OffhandSwingAt <= sim.CurrentTime)
}

// Returns the time at which the next event will occur, considering both autos and the gcd.
func (aa *AutoAttacks) NextEventAt(sim *Simulation) time.Duration {
	if aa.NextAttackAt() == sim.CurrentTime {
		panic(fmt.Sprintf("Returned 0 from next attack at %s, mh: %s, oh: %s", sim.CurrentTime, aa.MainhandSwingAt, aa.OffhandSwingAt))
	}
	return MinDuration(
		sim.CurrentTime+aa.character.GetRemainingCD(GCDCooldownID, sim.CurrentTime),
		aa.NextAttackAt())
}

type PPMManager struct {
	mhProcChance     float64
	ohProcChance     float64
	rangedProcChance float64
}

// For manually overriding proc chance.
func (ppmm *PPMManager) SetProcChance(isMH bool, newChance float64) {
	if isMH {
		ppmm.mhProcChance = newChance
	} else {
		ppmm.ohProcChance = newChance
	}
}
func (ppmm *PPMManager) SetRangedChance(newChance float64) {
	ppmm.rangedProcChance = newChance
}

// Returns whether the effect procced.
func (ppmm *PPMManager) Proc(sim *Simulation, isMH bool, isRanged bool, label string) bool {
	if isMH {
		return ppmm.ProcMH(sim, label)
	} else if !isRanged {
		return ppmm.ProcOH(sim, label)
	} else {
		return ppmm.ProcRanged(sim, label)
	}
}

// Returns whether the effect procced, assuming MH.
func (ppmm *PPMManager) ProcMH(sim *Simulation, label string) bool {
	return ppmm.mhProcChance > 0 && sim.RandomFloat(label) < ppmm.mhProcChance
}

// Returns whether the effect procced, assuming OH.
func (ppmm *PPMManager) ProcOH(sim *Simulation, label string) bool {
	return ppmm.ohProcChance > 0 && sim.RandomFloat(label) < ppmm.ohProcChance
}

// Returns whether the effect procced, assuming Ranged.
func (ppmm *PPMManager) ProcRanged(sim *Simulation, label string) bool {
	return ppmm.rangedProcChance > 0 && sim.RandomFloat(label) < ppmm.rangedProcChance
}

// PPMToChance converts a character proc-per-minute into mh/oh proc chances
func (aa *AutoAttacks) NewPPMManager(ppm float64) PPMManager {
	if aa.MH.SwingSpeed == 0 {
		// Means this character didn't enable autoattacks.
		return PPMManager{
			mhProcChance:     0,
			ohProcChance:     0,
			rangedProcChance: 0,
		}
	}

	return PPMManager{
		mhProcChance:     (aa.MH.SwingSpeed * ppm) / 60.0,
		ohProcChance:     (aa.OH.SwingSpeed * ppm) / 60.0,
		rangedProcChance: (aa.Ranged.SwingSpeed * ppm) / 60.0,
	}
}
