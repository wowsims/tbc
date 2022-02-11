package main

import (
	"github.com/wowsims/tbc/sim/core/proto"
)

type ItemDeclaration struct {
	ID int

	// Override fields, in case wowhead is wrong.
	Stats          Stats // Only non-zero values will override
	ClassAllowlist []proto.Class
	Phase          int
	Filter         bool // If true, this item will be omitted from the sim.
}
type ItemData struct {
	Declaration ItemDeclaration
	Response    WowheadItemResponse
}

type GemDeclaration struct {
	ID int

	// Override fields, in case wowhead is wrong.
	Stats Stats // Only non-zero values will override
	Phase int

	Filter bool // If true, this item will be omitted from the sim.
}
type GemData struct {
	Declaration GemDeclaration
	Response    WowheadItemResponse
}

// Allows manual overriding for Gem fields in case WowHead is wrong.
var GemDeclarationOverrides = []GemDeclaration{
	{ID: 33131, Stats: Stats{proto.Stat_StatAttackPower: 32, proto.Stat_StatRangedAttackPower: 32}},

	// pvp non-unique gems not in game currently.
	{ID: 35489, Filter: true},
	{ID: 38545, Filter: true},
	{ID: 38546, Filter: true},
	{ID: 38547, Filter: true},
	{ID: 38548, Filter: true},
	{ID: 38549, Filter: true},
	{ID: 38550, Filter: true},

	// BoP version of gems that can be made non-bop
	{ID: 35487, Filter: true},
	{ID: 35488, Filter: true},
	{ID: 35489, Filter: true},

	// Other gems to ignore.
	{ID: 35315, Filter: true},
	{ID: 35316, Filter: true},
	{ID: 35318, Filter: true},

	{ID: 35759, Phase: 5},
	{ID: 35760, Phase: 5},
}

// Allows manual overriding for Item fields in case WowHead is wrong.
var ItemDeclarationOverrides = []ItemDeclaration{
	{ /** Band of Eternity */ ID: 29302, Phase: 2},
	{ /** Destruction Holo-gogs */ ID: 32494, ClassAllowlist: []proto.Class{proto.Class_ClassMage, proto.Class_ClassPriest, proto.Class_ClassWarlock}},
	{ /** Gadgetstorm Goggles */ ID: 32476, ClassAllowlist: []proto.Class{proto.Class_ClassShaman}},
	{ /** Idol of the Raven Goddess */ ID: 32387, Phase: 2},
	{ /** Idol of the Unseen Moon */ ID: 33510, Phase: 4},
	{ /** Magnified Moon Specs */ ID: 32480, ClassAllowlist: []proto.Class{proto.Class_ClassDruid}},
	{ /** Skycall Totem */ ID: 33506, Phase: 4},
	{ /** Vindicator's Band of Triumph */ ID: 33919, Phase: 3},
	{ /** Vindicator's Pendant of Reprieve */ ID: 35317, Phase: 4},
	{ /** Vindicator's Pendant of Subjugation */ ID: 35319, Phase: 4},
	{ /** Twinblade of the Pheonix */ ID: 29993, Stats: Stats{proto.Stat_StatRangedAttackPower: 108}},

	{ID: 17782, Filter: true}, // talisman of the binding shard
	{ID: 17783, Filter: true}, // talisman of the binding fragment
	{ID: 18582, Filter: true},
	{ID: 18583, Filter: true},
	{ID: 18584, Filter: true},
	{ID: 24265, Filter: true},
	{ID: 24525, Filter: true},
	{ID: 32384, Filter: true},
	{ID: 32421, Filter: true},
	{ID: 32422, Filter: true},
	{ID: 33482, Filter: true},
}
