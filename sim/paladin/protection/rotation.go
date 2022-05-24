package protection

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func (prot *ProtectionPaladin) OnGCDReady(sim *core.Simulation) {
	if prot.CurrentSeal == prot.SealOfWisdomAura && prot.JudgementOfWisdom.IsReady(sim) {
		prot.JudgementOfWisdom.Cast(sim, prot.CurrentTarget)
		if prot.JudgementOfWisdomAura.IsActive() {
			prot.SealOfRighteousness.Cast(sim, nil)
		} else {
			// Re-cast seal of wisdom if we missed.
			prot.SealOfWisdom.Cast(sim, nil)
		}
		return
	} else if prot.CurrentSeal == prot.SealOfLightAura && prot.JudgementOfLight.IsReady(sim) {
		prot.JudgementOfLight.Cast(sim, prot.CurrentTarget)
		if prot.JudgementOfLightAura.IsActive() {
			prot.SealOfRighteousness.Cast(sim, nil)
		} else {
			// Re-cast seal of wisdom if we missed.
			prot.SealOfLight.Cast(sim, nil)
		}
		return
	}

	if prot.CurrentSeal == nil {
		prot.SealOfRighteousness.Cast(sim, nil)
	} else if prot.Rotation.PrioritizeHolyShield && prot.HolyShield.IsReady(sim) {
		prot.HolyShield.Cast(sim, nil)
	} else if prot.Consecration != nil && prot.Consecration.IsReady(sim) {
		prot.Consecration.Cast(sim, nil)
	} else if prot.JudgementOfRighteousness.IsReady(sim) {
		prot.JudgementOfRighteousness.Cast(sim, prot.CurrentTarget)
		prot.SealOfRighteousness.Cast(sim, nil)
	} else if prot.shouldExorcism(sim) {
		prot.Exorcism.Cast(sim, prot.CurrentTarget)
	} else if !prot.Rotation.PrioritizeHolyShield && prot.HolyShield.IsReady(sim) {
		prot.HolyShield.Cast(sim, nil)
	} else {
		prot.WaitUntil(sim, prot.nextCDAt(sim))
	}
}

func (prot *ProtectionPaladin) nextCDAt(sim *core.Simulation) time.Duration {
	nextCDAt := core.MinDuration(prot.HolyShield.ReadyAt(), prot.JudgementOfRighteousness.ReadyAt())
	if prot.Consecration != nil {
		nextCDAt = core.MinDuration(nextCDAt, prot.Consecration.ReadyAt())
	}
	return nextCDAt
}

func (prot *ProtectionPaladin) shouldExorcism(sim *core.Simulation) bool {
	return prot.Rotation.UseExorcism &&
		prot.CanExorcism(prot.CurrentTarget) &&
		prot.Exorcism.IsReady(sim) &&
		prot.CurrentMana() > prot.MaxMana()*0.4
}
