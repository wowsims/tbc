package rotations

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

// Interface for shared rotations. These can be embedded in an Agent to provide
// the Reset() and Act() handling.
type Rotation interface {
	// Identical to Agent.Reset().
	Reset(*core.Simulation)

	// Identical to Agent.Act().
	Act(*core.Simulation) time.Duration
}
