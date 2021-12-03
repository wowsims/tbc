package sim

import (
	_ "github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/druid/balance"
	"github.com/wowsims/tbc/sim/priest/shadow"
	"github.com/wowsims/tbc/sim/shaman/elemental"
)

var registered = false

func RegisterAll() {
	if registered {
		return
	}
	registered = true

	elemental.RegisterElementalShaman()
	balance.RegisterBalanceDruid()
	shadow.RegisterShadowPriest()
}
