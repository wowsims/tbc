package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (paladin *Paladin) ApplyTalents() {
	paladin.AddStat(stats.MeleeCrit, core.MeleeCritRatingPerCritChance*float64(paladin.Talents.SanctifiedSeals))
	paladin.AddStat(stats.SpellCrit, core.SpellCritRatingPerCritChance*float64(paladin.Talents.SanctifiedSeals))
	paladin.AddStat(stats.MeleeHit, core.MeleeHitRatingPerHitChance*float64(paladin.Talents.Precision))
	paladin.AddStat(stats.SpellHit, core.SpellHitRatingPerHitChance*float64(paladin.Talents.Precision))
	paladin.AddStat(stats.MeleeCrit, core.MeleeCritRatingPerCritChance*float64(paladin.Talents.Conviction))
	paladin.AddStat(stats.Parry, core.ParryRatingPerParryChance*1*float64(paladin.Talents.Deflection))
	paladin.AddStat(stats.Armor, paladin.Equip.Stats()[stats.Armor]*0.02*float64(paladin.Talents.Toughness))
	paladin.AddStat(stats.Defense, core.DefenseRatingPerDefense*4*float64(paladin.Talents.Anticipation))

	spellWardingMultiplier := 1 - 0.02*float64(paladin.Talents.SpellWarding)
	paladin.PseudoStats.ArcaneDamageTakenMultiplier *= spellWardingMultiplier
	paladin.PseudoStats.FireDamageTakenMultiplier *= spellWardingMultiplier
	paladin.PseudoStats.FrostDamageTakenMultiplier *= spellWardingMultiplier
	paladin.PseudoStats.HolyDamageTakenMultiplier *= spellWardingMultiplier
	paladin.PseudoStats.NatureDamageTakenMultiplier *= spellWardingMultiplier
	paladin.PseudoStats.ShadowDamageTakenMultiplier *= spellWardingMultiplier

	if paladin.Talents.DivineStrength > 0 {
		bonus := 1 + 0.02*float64(paladin.Talents.DivineStrength)
		paladin.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Strength,
			ModifiedStat: stats.Strength,
			Modifier: func(str float64, _ float64) float64 {
				return str * bonus
			},
		})
	}
	if paladin.Talents.DivineIntellect > 0 {
		bonus := 1 + 0.02*float64(paladin.Talents.DivineIntellect)
		paladin.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.Intellect,
			Modifier: func(intellect float64, _ float64) float64 {
				return intellect * bonus
			},
		})
	}

	if paladin.Talents.ShieldSpecialization > 0 {
		bonus := 1 + 0.1*float64(paladin.Talents.ShieldSpecialization)
		paladin.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.BlockValue,
			ModifiedStat: stats.BlockValue,
			Modifier: func(bv float64, _ float64) float64 {
				return bv * bonus
			},
		})
	}

	if paladin.Talents.SacredDuty > 0 {
		bonus := 1 + 0.03*float64(paladin.Talents.SacredDuty)
		paladin.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Stamina,
			ModifiedStat: stats.Stamina,
			Modifier: func(stam float64, _ float64) float64 {
				return stam * bonus
			},
		})
	}

	if paladin.Talents.CombatExpertise > 0 {
		paladin.AddStat(stats.Expertise, core.ExpertisePerQuarterPercentReduction*1*float64(paladin.Talents.CombatExpertise))
		bonus := 1 + 0.02*float64(paladin.Talents.CombatExpertise)
		paladin.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Stamina,
			ModifiedStat: stats.Stamina,
			Modifier: func(stam float64, _ float64) float64 {
				return stam * bonus
			},
		})
	}

	paladin.applyCrusade()
	paladin.applyOneHandedWeaponSpecialization()
	paladin.applyTwoHandedWeaponSpecialization()
	paladin.applyVengeance()
}

func (paladin *Paladin) applyCrusade() {
	if paladin.Talents.Crusade == 0 {
		return
	}

	damageMultiplier := 1 + (0.01 * float64(paladin.Talents.Crusade)) // assume multiplicative scaling

	// TO-DO: This doesn't account for multiple targets
	switch paladin.CurrentTarget.MobType {
	case proto.MobType_MobTypeHumanoid, proto.MobType_MobTypeDemon, proto.MobType_MobTypeUndead, proto.MobType_MobTypeElemental:
		paladin.PseudoStats.DamageDealtMultiplier *= damageMultiplier
	}
}

// Affects all physical damage or spells that can be rolled as physical
// It affects white, Windfury, Crusader Strike, Seals, and Judgement of Command / Blood
func (paladin *Paladin) applyTwoHandedWeaponSpecialization() {
	if paladin.Talents.TwoHandedWeaponSpecialization == 0 {
		return
	}

	// This impacts Crusader Strike, Melee Attacks, WF attacks
	// Seals + Judgements need to be implemented separately
	if paladin.GetMHWeapon().HandType == proto.HandType_HandTypeTwoHand {
		paladin.PseudoStats.PhysicalDamageDealtMultiplier *= 1 + (0.02 * float64(paladin.Talents.TwoHandedWeaponSpecialization)) // assume multiplicative scaling
	}
}

func (paladin *Paladin) applyTwoHandedWeaponSpecializationToSpell(spellEffect *core.SpellEffect) {
	if paladin.GetMHWeapon().HandType == proto.HandType_HandTypeTwoHand {
		spellEffect.DamageMultiplier *= 1 + (0.02 * float64(paladin.Talents.TwoHandedWeaponSpecialization))
	}
}

func (paladin *Paladin) applyOneHandedWeaponSpecialization() {
	if paladin.Talents.OneHandedWeaponSpecialization == 0 {
		return
	}
	if paladin.Equip[proto.ItemSlot_ItemSlotMainHand].HandType == proto.HandType_HandTypeTwoHand {
		return
	}

	paladin.PseudoStats.PhysicalDamageDealtMultiplier *= 1 + 0.01*float64(paladin.Talents.OneHandedWeaponSpecialization)
}

// I don't know if the new stack of vengeance applies to the crit that triggered it or not
// Need to check this
func (paladin *Paladin) applyVengeance() {
	if paladin.Talents.Vengeance == 0 {
		return
	}

	bonusPerStack := 0.01 * float64(paladin.Talents.Vengeance)
	procAura := paladin.RegisterAura(core.Aura{
		Label:     "Vengeance Proc",
		ActionID:  core.ActionID{SpellID: 20059},
		Duration:  time.Second * 30,
		MaxStacks: 3,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			aura.Unit.PseudoStats.DamageDealtMultiplier /= 1 + (bonusPerStack * float64(oldStacks))
			aura.Unit.PseudoStats.DamageDealtMultiplier *= 1 + (bonusPerStack * float64(newStacks))
		},
	})

	paladin.RegisterAura(core.Aura{
		Label:    "Vengeance",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Outcome.Matches(core.OutcomeCrit) {
				procAura.Activate(sim)
				procAura.AddStack(sim)
			}
		},
	})
}
