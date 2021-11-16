package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

// Totems that shaman will cast.
// TODO: add logic inside these to select each totem based on options on the shaman?

func (shaman *Shaman) NewAirTotem() *core.NoEffectSpell {
	return &core.NoEffectSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				Name:            "Wrath of Air",
				ActionID:        core.ActionID{SpellID: 3738}, // just using totem of wrath
				Character:       shaman.GetCharacter(),
				BaseManaCost:    240,
				ManaCost:        240,
				CastTime:        time.Second * 1,
				IgnoreCooldowns: true, // lets us override the GCD
			},
		},
	}
}

func (shaman *Shaman) NewWaterTotem() *core.NoEffectSpell {
	return &core.NoEffectSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				Name:            "Mana Stream",
				ActionID:        core.ActionID{SpellID: 25570}, // just using totem of wrath
				Character:       shaman.GetCharacter(),
				BaseManaCost:    90,
				ManaCost:        90,
				CastTime:        time.Second * 1,
				IgnoreCooldowns: true, // lets us override the GCD
			},
		},
	}
}

func (shaman *Shaman) NewFireTotem() *core.NoEffectSpell {
	return &core.NoEffectSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				Name:            "Totem of Wrath",
				ActionID:        core.ActionID{SpellID: 30706}, // just using totem of wrath
				Character:       shaman.GetCharacter(),
				BaseManaCost:    240,
				ManaCost:        240,
				CastTime:        time.Second * 1,
				IgnoreCooldowns: true, // lets us override the GCD
			},
		},
	}
}
