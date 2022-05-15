package encounters

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	const bossPrefix = "Black Temple/Bosses"

	core.AddPresetTarget(core.PresetTarget{
		PathPrefix: bossPrefix,
		Config: proto.Target{
			Id:      22917,
			Name:    "Illidan Stormrage",
			Level:   73,
			MobType: proto.MobType_MobTypeDemon,

			Stats: stats.Stats{
				stats.Health:      8_300_000,
				stats.Armor:       7684,
				stats.AttackPower: 320,
				stats.BlockValue:  54,
			}.ToFloatArray(),

			SpellSchool:   proto.SpellSchool_SpellSchoolPhysical,
			SwingSpeed:    1.5,
			MinBaseDamage: 16486.65,
			CanCrush:      false,
			ParryHaste:    true,
			DualWield:     true,
		},
	})

	core.AddPresetEncounter("Illidan Stormrage", []string{
		bossPrefix + "/Illidan Stormrage",
	})
}
