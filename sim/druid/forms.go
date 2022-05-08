package druid

import (
	"time"

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

func (druid *Druid) setForm(form DruidForm, sim *core.Simulation) {
	if form == druid.form {
		return
	}

	druid.form = form

	druid.manageCooldownsEnabled(sim)
}

func (druid *Druid) GetForm() DruidForm {
	return druid.form
}

func (druid *Druid) InForm(form DruidForm) bool {
	return druid.form.Matches(form)
}

func (druid *Druid) PowerShiftCat(sim *core.Simulation) {

	if !druid.GCD.IsReady(sim) {
		panic("Trying to powershift during gcd")
	}
	// Go out of form
	druid.setForm(Humanoid, sim)

	druid.energyBeforeShift = druid.CurrentEnergy()

	// Try use cds
	druid.TryUseCooldowns(sim)

	// If cds did not trigger gcd (innervate), powershift
	if druid.GCD.IsReady(sim) {
		druid.CatForm.Cast(sim, nil)
	} else {
		druid.AutoAttacks.CancelAutoSwing(sim)
	}
}

func (druid *Druid) registerCatFormSpell() {
	actionID := core.ActionID{SpellID: 768}
	baseCost := druid.BaseMana() * 0.35

	finalEnergy := 0.0
	furorProcChance := 0.2 * float64(druid.Talents.Furor)
	if druid.Equip[items.ItemSlotHead].ID == 8345 { // Wolfshead Helm
		finalEnergy += 20.0
	}

	druid.CatForm = druid.RegisterSpell(core.SpellConfig{
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
			energyDelta := finalEnergy - druid.energyBeforeShift
			if furorProcChance == 1 || (furorProcChance > 0 && sim.RandomFloat("Furor") < furorProcChance) {
				energyDelta += 40.0
			}
			if energyDelta > 0 {
				druid.AddEnergy(sim, energyDelta, spell.ActionID)
			} else if energyDelta < 0 {
				druid.SpendEnergy(sim, -energyDelta, spell.ActionID)
			}
			druid.setForm(Cat, sim)
			druid.AutoAttacks.EnableAutoSwing(sim)
		},
	})
}

func (druid *Druid) PowerShiftBear(sim *core.Simulation) {

	if !druid.GCD.IsReady(sim) {
		panic("Trying to powershift during gcd")
	}
	// Go out of form
	druid.setForm(Humanoid, sim)

	// Try use cds
	druid.TryUseCooldowns(sim)

	// If cds did not trigger gcd (innervate), powershift
	if druid.GCD.IsReady(sim) {
		druid.BearForm.Cast(sim, nil)
	} else {
		druid.AutoAttacks.CancelAutoSwing(sim)
	}
}

func (druid *Druid) registerBearFormSpell() {
	actionID := core.ActionID{SpellID: 9634}
	baseCost := druid.BaseMana() * 0.35

	furorProcChance := 0.2 * float64(druid.Talents.Furor)
	finalRage := 0.0
	if druid.Equip[items.ItemSlotHead].ID == 8345 { // Wolfshead Helm
		finalRage += 5.0
	}

	druid.BearForm = druid.RegisterSpell(core.SpellConfig{
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
			dRage := finalRage - druid.CurrentEnergy()
			if furorProcChance == 1 || (furorProcChance > 0 && sim.RandomFloat("Furor") < furorProcChance) {
				finalRage += 10.0
			}
			if dRage > 0 {
				druid.AddRage(sim, dRage, spell.ActionID)
			} else if dRage < 0 {
				druid.SpendRage(sim, -dRage, spell.ActionID)
			}
			druid.setForm(Bear, sim)
			druid.AutoAttacks.EnableAutoSwing(sim)
		},
	})
}

const cooldownDelayThresHold = time.Second * 10

// Disable cooldowns not usable in form and/or delay others
func (druid *Druid) manageCooldownsEnabled(sim *core.Simulation) {

	if druid.StartingForm.Matches(Cat | Bear) {

		druid.EnableAllCooldowns(druid.disabledMCDs)
		druid.disabledMCDs = nil

		if druid.InForm(Cat | Bear) {
			// Check if any dps cooldown that requires shifting is ready soon
			// disable all cooldowns if that is the case
			nonUsableDpsMCDReadySoon := false
			for _, cd := range druid.GetMajorCooldowns() {
				if cd.TimeToReady(sim) < cooldownDelayThresHold && cd.IsEnabled() && !cd.Type.Matches(core.CooldownTypeUsableShapeShifted) && cd.Type.Matches(core.CooldownTypeDPS) {
					nonUsableDpsMCDReadySoon = true
					break
				}
			}
			for _, cd := range druid.GetMajorCooldowns() {
				if cd.IsEnabled() && (nonUsableDpsMCDReadySoon || !cd.Type.Matches(core.CooldownTypeUsableShapeShifted)) {
					druid.DisableMajorCooldown(cd.Spell.ActionID)
					druid.disabledMCDs = append(druid.disabledMCDs, cd)
				}
			}
		} else {
			// Disable cooldown that can be used in form, but incurs a gcd, so we dont get stuck out of form when we dont need to (Greater Drums)
			for _, cd := range druid.GetMajorCooldowns() {
				if cd.Type.Matches(core.CooldownTypeUsableShapeShifted) && cd.Spell.DefaultCast.GCD > 0 {
					druid.DisableMajorCooldown(cd.Spell.ActionID)
					druid.disabledMCDs = append(druid.disabledMCDs, cd)
				}
			}
		}
	}
}
