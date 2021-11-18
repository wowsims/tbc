package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

const manaTrackingWindowSeconds = 60
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
}

type manaSnapshot struct {
	time      time.Duration // time this snapshot was taken
	manaSpent float64       // total amount of mana spent up to this time
}

func NewManaSpendingRateTracker() ManaSpendingRateTracker {
	return ManaSpendingRateTracker{}
}

// This needs to be called on sim reset.
func (tracker *ManaSpendingRateTracker) Reset() {
	tracker.manaSnapshots = [manaSnapshotsBufferSize]manaSnapshot{}
	tracker.firstSnapshotIndex = 0
	tracker.numSnapshots = 0
}

func (tracker *ManaSpendingRateTracker) getOldestSnapshot() manaSnapshot {
	return tracker.manaSnapshots[tracker.firstSnapshotIndex]
}

func (tracker *ManaSpendingRateTracker) purgeExpiredSnapshots(sim *core.Simulation) {
	expirationCutoff := sim.CurrentTime - manaTrackingWindow

	curIndex := tracker.firstSnapshotIndex
	for tracker.numSnapshots > 0 && tracker.manaSnapshots[curIndex].time < expirationCutoff {
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

	snapshot := manaSnapshot{
		time:      sim.CurrentTime,
		manaSpent: sim.GetIndividualMetrics(character.ID).ManaSpent,
	}

	nextIndex := (tracker.firstSnapshotIndex + tracker.numSnapshots) % manaSnapshotsBufferSize
	tracker.manaSnapshots[nextIndex] = snapshot
	tracker.numSnapshots++
}

func (tracker *ManaSpendingRateTracker) ManaSpentPerSecond(sim *core.Simulation, character *core.Character) float64 {
	tracker.purgeExpiredSnapshots(sim)
	oldestSnapshot := tracker.getOldestSnapshot()

	manaSpent := sim.GetIndividualMetrics(character.ID).ManaSpent - oldestSnapshot.manaSpent
	timeDelta := sim.CurrentTime - oldestSnapshot.time
	if timeDelta == 0 {
		timeDelta = 1
	}

	return manaSpent / timeDelta.Seconds()
}

// The amount of mana we will need to spend over the remaining sim duration
// at the current rate of mana spending.
func (tracker *ManaSpendingRateTracker) ProjectedManaCost(sim *core.Simulation, character *core.Character) float64 {
	manaSpentPerSecond := tracker.ManaSpentPerSecond(sim, character)

	timeRemaining := sim.Duration - sim.CurrentTime
	return manaSpentPerSecond * timeRemaining.Seconds()
}
