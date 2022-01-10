package main

import (
	"github.com/wowsims/tbc/sim/core/proto"
)

type ItemDeclaration struct {
	ID int

	// Override fields, in case wowhead is wrong.
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
	Phase int

	Filter bool // If true, this item will be omitted from the sim.
}
type GemData struct {
	Declaration GemDeclaration
	Response    WowheadItemResponse
}

// Allows manual overriding for Gem fields in case WowHead is wrong.
var GemDeclarationOverrides = []GemDeclaration{
	{ID: 35315, Filter: true},
	{ID: 35316, Filter: true},
	{ID: 35318, Filter: true},
	{ID: 35759, Phase: 5},
	{ID: 35760, Phase: 5},

	// Meta gems.
	{ID: 25890},
	{ID: 25893},
	{ID: 25894},
	{ID: 25895},
	{ID: 25896},
	{ID: 25897},
	{ID: 25898},
	{ID: 25899},
	{ID: 25901},
	{ID: 28556},
	{ID: 28557},
	{ID: 32409},
	{ID: 32410},
	{ID: 32640},
	{ID: 32641},
	{ID: 34220},
	{ID: 35501},
	{ID: 35503},
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

	{ID: 18582, Filter: true},
	{ID: 18583, Filter: true},
	{ID: 18584, Filter: true},
	{ID: 24265, Filter: true},
	{ID: 24525, Filter: true},
}
