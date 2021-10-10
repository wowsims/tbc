package sim

import (
	"github.com/wowsims/tbc/sim/shaman/elemental"
)

var registered = false
func RegisterAll() {
	if registered {
		return
	}
	registered = true

	elemental.RegisterElementalShaman()
}
