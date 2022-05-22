package druid

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

type DruidForm uint8

const (
	Humanoid = 1 << iota
	Bear
	Cat
	Moonkin
)

func (form DruidForm) Matches(other DruidForm) bool {
	return (form & other) != 0
}

func (druid *Druid) registerPowershiftSpell() {
	actionID := core.ActionID{SpellID: 768}
	baseCost := druid.BaseMana() * 0.35

	// Assumes 5/5 Furor.
	finalEnergy := 40.0
	if druid.Equip[items.ItemSlotHead].ID == 8345 { // Wolfshead Helm
		finalEnergy += 20.0
	}

	druid.Powershift = druid.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellExtras: core.SpellExtrasNoOnCastComplete,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost * (1 - 0.1*float64(druid.Talents.NaturalShapeshifter)),
				GCD:  core.GCDDefault,
			},
			IgnoreHaste: true,
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			druid.AddEnergy(sim, finalEnergy-druid.CurrentEnergy(), spell.ActionID)
			druid.Form = Cat
		},
	})
}
