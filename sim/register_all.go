package sim

import (
	_ "github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/druid/balance"
	"github.com/wowsims/tbc/sim/hunter"
	"github.com/wowsims/tbc/sim/mage"
	"github.com/wowsims/tbc/sim/paladin/retribution"
	"github.com/wowsims/tbc/sim/priest/shadow"
	"github.com/wowsims/tbc/sim/priest/smite"
	"github.com/wowsims/tbc/sim/rogue"
	"github.com/wowsims/tbc/sim/shaman/elemental"
	"github.com/wowsims/tbc/sim/shaman/enhancement"
	dpsWarrior "github.com/wowsims/tbc/sim/warrior/dps"
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
	hunter.RegisterHunter()
	mage.RegisterMage()
	shadow.RegisterShadowPriest()
	rogue.RegisterRogue()
	dpsWarrior.RegisterDpsWarrior()
	retribution.RegisterRetributionPaladin()
	smite.RegisterSmitePriest()
}
