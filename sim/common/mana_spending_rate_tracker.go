package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

const manaTrackingWindowSeconds = 30
const manaTrackingWindow = time.Second * manaTrackingWindowSeconds

// 2 * (# of seconds) should be plenty of slots
const manaSnapshotsBufferSize = manaTrackingWindowSeconds * 2

// Tracks how fast mana is being spent. This is used by some specs to decide
// whether to use more mana-efficient or higher-dps spells.
type ManaSpendingRateTracker struct {
	// Circular array buffer for recent mana snapshots, within a time window
	manaSnapshots      [manaSnapshotsBufferSize]manaSnapshot
	numSnapshots       int32
	firstSnapshotIndex int32

	manaSpentDuringWindow float64

	previousManaSpent float64
	previousCastSpeed float64
}

type manaSnapshot struct {
	time      time.Duration // time this snapshot was taken
	manaSpent float64       // total amount of mana spent up to this time

	manaSpentDelta float64
}

func NewManaSpendingRateTracker() ManaSpendingRateTracker {
	return ManaSpendingRateTracker{}
}

// This needs to be called on sim reset.
func (tracker *ManaSpendingRateTracker) Reset() {
	tracker.manaSnapshots = [manaSnapshotsBufferSize]manaSnapshot{}
	tracker.firstSnapshotIndex = 0
	tracker.numSnapshots = 0
	tracker.manaSpentDuringWindow = 0
	tracker.previousManaSpent = 0
	tracker.previousCastSpeed = 1
}

func (tracker *ManaSpendingRateTracker) getOldestSnapshot() manaSnapshot {
	return tracker.manaSnapshots[tracker.firstSnapshotIndex]
}

func (tracker *ManaSpendingRateTracker) purgeExpiredSnapshots(sim *core.Simulation) {
	expirationCutoff := sim.CurrentTime - manaTrackingWindow

	curIndex := tracker.firstSnapshotIndex
	for tracker.numSnapshots > 0 && tracker.manaSnapshots[curIndex].time < expirationCutoff {
		tracker.manaSpentDuringWindow -= tracker.manaSnapshots[curIndex].manaSpentDelta
		curIndex = (curIndex + 1) % manaSnapshotsBufferSize
		tracker.numSnapshots--
	}
	tracker.firstSnapshotIndex = curIndex
}

// This needs to be called at regular intervals to update the tracker's data.
func (tracker *ManaSpendingRateTracker) Update(sim *core.Simulation, character *core.Character) {
	if tracker.numSnapshots >= manaSnapshotsBufferSize {
		panic("Mana tracker snapshot buffer is full")
	}

	// Scale down mana spent so we don't get bad estimates from lust/drums/etc.
	manaSpentCoefficient := character.InitialCastSpeed() / tracker.previousCastSpeed

	snapshot := manaSnapshot{
		time:           sim.CurrentTime,
		manaSpent:      character.Metrics.ManaSpent,
		manaSpentDelta: (character.Metrics.ManaSpent - tracker.previousManaSpent) * manaSpentCoefficient,
	}

	nextIndex := (tracker.firstSnapshotIndex + tracker.numSnapshots) % manaSnapshotsBufferSize
	tracker.previousCastSpeed = character.CastSpeed()
	tracker.previousManaSpent = snapshot.manaSpent
	tracker.manaSpentDuringWindow += snapshot.manaSpentDelta
	tracker.manaSnapshots[nextIndex] = snapshot
	tracker.numSnapshots++
}

func (tracker *ManaSpendingRateTracker) ManaSpentPerSecond(sim *core.Simulation, character *core.Character) float64 {
	tracker.purgeExpiredSnapshots(sim)
	oldestSnapshot := tracker.getOldestSnapshot()

	manaSpent := tracker.manaSpentDuringWindow
	timeDelta := sim.CurrentTime - oldestSnapshot.time
	if timeDelta == 0 {
		return 0
	}

	return manaSpent / timeDelta.Seconds()
}

// The amount of mana we will need to spend over the remaining sim duration
// at the current rate of mana spending.
func (tracker *ManaSpendingRateTracker) ProjectedManaCost(sim *core.Simulation, character *core.Character) float64 {
	manaSpentPerSecond := tracker.ManaSpentPerSecond(sim, character)

	timeRemaining := sim.Duration - sim.CurrentTime
	projectedManaCost := manaSpentPerSecond * timeRemaining.Seconds()

	//if sim.Log != nil {
	//	remainingManaPool := character.ExpectedRemainingManaPool(sim)
	//	sim.Log("Mana spent: %0.02f, Projected: %0.02f, total: %0.02f (%0.02f + %0.02f)\n", character.Metrics.ManaSpent, projectedManaCost, remainingManaPool, character.CurrentMana(), character.ExpectedBonusMana)
	//}

	return projectedManaCost
}
