package core

import (
	"hash/fnv"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

func MinInt(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

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

func GetTristateValueInt32(effect proto.TristateEffect, regularValue int32, impValue int32) int32 {
	if effect == proto.TristateEffect_TristateEffectRegular {
		return regularValue
	} else if effect == proto.TristateEffect_TristateEffectImproved {
		return impValue
	} else {
		return 0
	}
}

func GetTristateValueFloat(effect proto.TristateEffect, regularValue float64, impValue float64) float64 {
	if effect == proto.TristateEffect_TristateEffectRegular {
		return regularValue
	} else if effect == proto.TristateEffect_TristateEffectImproved {
		return impValue
	} else {
		return 0
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func TernaryFloat64(condition bool, val1 float64, val2 float64) float64 {
	if condition {
		return val1
	} else {
		return val2
	}
}
