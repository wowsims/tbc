package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

func MinInt32(a int32, b int32) int32 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxInt32(a int32, b int32) int32 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinFloat(a float64, b float64) float64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxFloat(a float64, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinDuration(a time.Duration, b time.Duration) time.Duration {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxDuration(a time.Duration, b time.Duration) time.Duration {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinTristate(a proto.TristateEffect, b proto.TristateEffect) proto.TristateEffect {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxTristate(a proto.TristateEffect, b proto.TristateEffect) proto.TristateEffect {
	if a > b {
		return a
	} else {
		return b
	}
}

func DurationFromSeconds(numSeconds float64) time.Duration {
	return time.Duration(float64(time.Second) * numSeconds)
}
