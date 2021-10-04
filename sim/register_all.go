package sim

import (
	"github.com/wowsims/tbc/sim/shaman"
)

var registered = false
func RegisterAll() {
	if registered {
		return
	}
	registered = true

	shaman.RegisterElementalShaman()
}
