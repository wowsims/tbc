package druid

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (druid *Druid) registerLacerateSpell() {
	actionID := core.ActionID{SpellID: 33745}

	cost := 15.0 - float64(druid.Talents.ShreddingAttacks)
	refundAmount := cost * 0.8

	druid.Lacerate = druid.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

		ResourceType: stats.Rage,
		BaseCost:     cost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: cost,
				GCD:  core.GCDDefault,
			},
			IgnoreHaste: true,
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			OutcomeApplier:   druid.OutcomeFuncMeleeSpecialHit(),
			OnSpellHitDealt: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					if druid.LacerateDot.IsActive() {
						druid.LacerateDot.Refresh(sim)
						druid.LacerateDot.AddStack(sim)
					} else {
						druid.LacerateDot.Apply(sim)
						druid.LacerateDot.SetStacks(sim, 1)
					}
				} else {
					druid.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
				}
			},
		}),
	})

	tickDamage := 155.0 / 5
	if ItemSetNordrassilHarness.CharacterHasSetBonus(&druid.Character, 4) {
		tickDamage += 15
	}
	if druid.Equip[items.ItemSlotRanged].ID == 27744 { // Idol of Ursoc
		tickDamage += 8
	}

	target := druid.CurrentTarget
	dotAura := target.RegisterAura(core.Aura{
		Label:     "Lacerate-" + strconv.Itoa(int(druid.Index)),
		ActionID:  actionID,
		MaxStacks: 5,
		Duration:  time.Second * 15,
	})
	druid.LacerateDot = core.NewDot(core.Dot{
		Spell:         druid.Lacerate,
		Aura:          dotAura,
		NumberOfTicks: 5,
		TickLength:    time.Second * 3,
		TickEffects: core.TickFuncApplyEffects(core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			IsPeriodic:       true,
			IsPhantom:        true,
			BaseDamage:       core.MultiplyByStacks(core.BaseDamageConfigFlat(tickDamage), dotAura),
			OutcomeApplier:   druid.OutcomeFuncTick(),
		})),
	})
}

func (druid *Druid) CanLacerate(sim *core.Simulation) bool {
	return druid.CurrentRage() >= druid.Lacerate.DefaultCast.Cost
}
