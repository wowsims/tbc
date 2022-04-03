package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func AddSimpleStatItemActiveEffect(itemID int32, bonus stats.Stats, duration time.Duration, cooldown time.Duration, sharedCooldownID core.CooldownID) {
	auraID := core.NewAuraID()
	cooldownID := core.NewCooldownID()

	core.AddItemEffect(itemID, core.MakeTemporaryStatsOnUseCDRegistration(
		auraID,
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
