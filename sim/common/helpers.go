package common

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func AddSimpleStatItemActiveEffect(itemID int32, bonus stats.Stats, duration time.Duration, cooldown time.Duration, sharedCDFunc func(*core.Character) core.Cooldown) {
	core.AddItemEffect(itemID, core.MakeTemporaryStatsOnUseCDRegistration(
		"ItemActive-"+strconv.Itoa(int(itemID)),
		bonus,
		duration,
		core.SpellConfig{
			ActionID: core.ActionID{ItemID: itemID},
		},
		func(character *core.Character) core.Cooldown {
			return core.Cooldown{
				Timer:    character.NewTimer(),
				Duration: cooldown,
			}
		},
		sharedCDFunc,
	))
}

func AddSimpleStatOffensiveTrinketEffect(itemID int32, bonus stats.Stats, duration time.Duration, cooldown time.Duration) {
	AddSimpleStatItemActiveEffect(itemID, bonus, duration, cooldown, func(character *core.Character) core.Cooldown {
		return core.Cooldown{
			Timer:    character.GetOffensiveTrinketCD(),
			Duration: duration,
		}
	})
}

func AddSimpleStatDefensiveTrinketEffect(itemID int32, bonus stats.Stats, duration time.Duration, cooldown time.Duration) {
	AddSimpleStatItemActiveEffect(itemID, bonus, duration, cooldown, func(character *core.Character) core.Cooldown {
		return core.Cooldown{
			Timer:    character.GetDefensiveTrinketCD(),
			Duration: duration,
		}
	})
}
