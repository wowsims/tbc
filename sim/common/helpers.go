package common

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func AddSimpleStatItemActiveEffect(itemID int32, bonus stats.Stats, duration time.Duration, cooldown time.Duration, sharedCooldownID core.CooldownID) {
	cooldownID := core.NewCooldownID()

	core.AddItemEffect(itemID, core.MakeTemporaryStatsOnUseCDRegistration(
		"ItemActive-"+strconv.Itoa(int(itemID)),
		bonus,
		duration,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: itemID},
			CooldownID:       cooldownID,
			Cooldown:         cooldown,
			SharedCooldownID: sharedCooldownID,
		},
	))
}
