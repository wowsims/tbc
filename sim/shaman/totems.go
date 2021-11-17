package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

// Totems that shaman will cast.
// TODO: add logic inside these to select each totem based on options on the shaman?

func (shaman *Shaman) NewAirTotem() *core.NoEffectSpell {
	manaCost := 360 * (1 - float64(shaman.Talents.TotemicFocus) * 0.05)

	return &core.NoEffectSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				Name:            "Wrath of Air",
				ActionID:        core.ActionID{SpellID: 3738},
				Character:       shaman.GetCharacter(),
				BaseManaCost:    manaCost,
				ManaCost:        manaCost,
				CastTime:        time.Second * 1,
				IgnoreCooldowns: true, // lets us override the GCD
			},
		},
	}
}

func (shaman *Shaman) NewWaterTotem() *core.NoEffectSpell {
	manaCost := 120 * (1 - float64(shaman.Talents.TotemicFocus) * 0.05)

	return &core.NoEffectSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				Name:            "Mana Stream",
				ActionID:        core.ActionID{SpellID: 25570},
				Character:       shaman.GetCharacter(),
				BaseManaCost:    manaCost,
				ManaCost:        manaCost,
				CastTime:        time.Second * 1,
				IgnoreCooldowns: true, // lets us override the GCD
			},
		},
	}
}

func (shaman *Shaman) NewFireTotem() *core.NoEffectSpell {
	manaCost := 360 * (1 - float64(shaman.Talents.TotemicFocus) * 0.05)

	return &core.NoEffectSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				Name:            "Totem of Wrath",
				ActionID:        core.ActionID{SpellID: 30706},
				Character:       shaman.GetCharacter(),
				BaseManaCost:    manaCost,
				ManaCost:        manaCost,
				CastTime:        time.Second * 1,
				IgnoreCooldowns: true, // lets us override the GCD
			},
		},
	}
}
