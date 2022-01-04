package shaman

import (
	"log"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ShamanisticRageCD = core.NewCooldownID()
var ShamanisticRageAuraID = core.NewAuraID()

const ShamanisticRageSpellID = 30823

func (shaman *Shaman) TryActivateShamanisticRage(sim *core.Simulation) bool {
	if shaman.GetRemainingCD(ShamanisticRageCD, sim.CurrentTime) > 0 {
		return false
	}
	const proc = 0.3
	const dur = time.Second * 15
	const cd = time.Second * 120
	shaman.AddAura(sim, core.Aura{
		ID:      ShamanisticRageAuraID,
		Name:    "Shamanistic Rage",
		SpellID: ShamanisticRageSpellID,
		Expires: sim.CurrentTime + dur,
		OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
			if result == core.MeleeHitTypeMiss {
				return
			}
			if sim.RandomFloat("shamanistic rage") > proc {
				return
			}
			mana := shaman.GetStat(stats.AttackPower) * 0.3
			if mana < 0 {
				log.Printf("Attack power!? %0.1f", shaman.GetStat(stats.AttackPower))
			}
			shaman.AddMana(sim, mana, "shamanistic rage", true)
		},
	})

	shaman.Metrics.AddInstantCast(core.ActionID{SpellID: ShamanisticRageSpellID})
	shaman.SetCD(ShamanisticRageCD, sim.CurrentTime+cd)
	return true
}
