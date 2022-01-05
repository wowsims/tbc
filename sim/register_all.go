package sim

import (
	_ "github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/druid/balance"
	"github.com/wowsims/tbc/sim/mage"
	"github.com/wowsims/tbc/sim/priest/shadow"
	"github.com/wowsims/tbc/sim/shaman/elemental"
	"github.com/wowsims/tbc/sim/shaman/enhancement"
)

var registered = false

func RegisterAll() {
	if registered {
		return
	}
	registered = true

	balance.RegisterBalanceDruid()
	elemental.RegisterElementalShaman()
	enhancement.RegisterEnhancementShaman()
	mage.RegisterMage()
	shadow.RegisterShadowPriest()
}
