package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDEarthShock int32 = 25454
const SpellIDFlameShock int32 = 25457
const SpellIDFrostShock int32 = 25464

var ShockCooldownID = core.NewCooldownID() // shared CD for all shocks
var FlameShockDebuffID = core.NewDebuffID()

func (shaman *Shaman) ShockCD() time.Duration {
	return time.Second*6 - time.Millisecond*200*time.Duration(shaman.Talents.Reverberation)
}

// Shared logic for all shocks.
func (shaman *Shaman) newShockTemplateSpell(sim *core.Simulation, spellID int32, spellSchool core.SpellSchool, baseManaCost float64) core.SimpleSpell {
	cost := core.ResourceCost{Type: stats.Mana, Value: baseManaCost}
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID:    spellID,
					CooldownID: ShockCooldownID,
				},
				Character:           &shaman.Character,
				CritRollCategory:    core.CritRollCategoryMagical,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         spellSchool,
				SpellExtras:         SpellFlagShock,
				BaseCost:            cost,
				Cost:                cost,
				GCD:                 core.GCDDefault,
				Cooldown:            shaman.ShockCD(),
				CritMultiplier:      shaman.DefaultSpellCritMultiplier(),
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
		},
	}

	spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.Convection) * 0.02
	spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.MentalQuickness) * 0.02
	if shaman.Talents.ElementalFury {
		spell.CritMultiplier = shaman.SpellCritMultiplier(1, 1)
	}

	spell.Effect.ThreatMultiplier *= 1 - (0.1/3)*float64(shaman.Talents.ElementalPrecision)
	spell.Effect.DamageMultiplier *= 1 + 0.01*float64(shaman.Talents.Concussion)
	spell.Effect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance
	spell.Effect.SpellEffect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance

	// TODO: confirm this is how it reduces mana cost.
	if ItemSetSkyshatterHarness.CharacterHasSetBonus(&shaman.Character, 2) {
		spell.Cost.Value -= spell.BaseCost.Value * 0.1
	}

	if shaman.Equip[items.ItemSlotRanged].ID == TotemOfRage {
		spell.Effect.BonusSpellPower += 30
	} else if shaman.Equip[items.ItemSlotRanged].ID == TotemOfImpact {
		spell.Effect.BonusSpellPower += 46
	}

	if shaman.Talents.ElementalFocus {
		spell.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
			if shaman.ElementalFocusStacks > 0 {
				shaman.ElementalFocusStacks--
			}
		}
		spell.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			if spellEffect.Outcome.Matches(core.OutcomeCrit) {
				shaman.ElementalFocusStacks = 2
			}
		}
	}

	return spell
}

// Shared shock logic that is dynamic, i.e. can't be precomputed.
func (shaman *Shaman) applyShockInitModifiers(spellCast *core.SpellCast) {
	if shaman.ElementalFocusStacks > 0 {
		// Reduces mana cost by 40%
		spellCast.Cost.Value -= spellCast.BaseCost.Value * 0.4
	}
}

func (shaman *Shaman) newEarthShockTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := shaman.newShockTemplateSpell(sim, SpellIDEarthShock, core.SpellSchoolNature, 535.0)
	spell.SpellExtras |= core.SpellExtrasBinary

	spell.Effect.BaseDamage = core.BaseDamageFuncMagic(661, 696, 0.386)

	return core.NewSimpleSpellTemplate(spell)
}

func (shaman *Shaman) NewEarthShock(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	shock := &shaman.shockSpell
	shaman.earthShockTemplate.Apply(shock)
	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	shock.Effect.Target = target
	shaman.applyShockInitModifiers(&shock.SpellCast)
	shock.Init(sim)

	return shock
}

func (shaman *Shaman) newFlameShockTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := shaman.newShockTemplateSpell(sim, SpellIDFlameShock, core.SpellSchoolFire, 500.0)

	spell.Effect.BaseDamage = core.BaseDamageFuncMagic(377, 377, 0.214)
	spell.Effect.DotInput = core.DotDamageInput{
		NumberOfTicks:  4,
		TickLength:     time.Second * 3,
		TickBaseDamage: core.DotSnapshotFuncMagic(420/4, 0.1),
		DebuffID:       FlameShockDebuffID,
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (shaman *Shaman) NewFlameShock(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	if shaman.FlameShockSpell.IsInUse() {
		// Cancel old cast, i.e. overwrite the dot.
		shaman.FlameShockSpell.Cancel(sim)
	}

	shock := &shaman.FlameShockSpell
	shaman.flameShockTemplate.Apply(shock)
	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	shock.Effect.Target = target
	shaman.applyShockInitModifiers(&shock.SpellCast)
	shock.Init(sim)

	return shock
}

func (shaman *Shaman) newFrostShockTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := shaman.newShockTemplateSpell(sim, SpellIDFrostShock, core.SpellSchoolFrost, 525.0)
	spell.SpellExtras |= core.SpellExtrasBinary

	spell.Effect.BaseDamage = core.BaseDamageFuncMagic(647, 683, 0.386)
	spell.Effect.ThreatMultiplier *= 2

	return core.NewSimpleSpellTemplate(spell)
}

func (shaman *Shaman) NewFrostShock(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	shock := &shaman.shockSpell
	shaman.frostShockTemplate.Apply(shock)
	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	shock.Effect.Target = target
	shaman.applyShockInitModifiers(&shock.SpellCast)
	shock.Init(sim)

	return shock
}
