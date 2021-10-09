package core

import (
	"time"
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
