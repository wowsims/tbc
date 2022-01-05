package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

// TODO: magma and nova totems need to apply to all targets probably instead of just the primary target.

const SpellIDSearingTotem int32 = 25533

// This is probably not worth simming since no other spell in the game does this and AM isn't
// even a popular choice for arcane mages.
func (shaman *Shaman) newSearingTotemTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				Name:           "Searing Totem",
				CritMultiplier: 1.5,
				SpellSchool:    stats.FireSpellPower,
				Character:      &shaman.Character,
				BaseManaCost:   205,
				ManaCost:       205,
				ActionID: core.ActionID{
					SpellID: SpellIDSearingTotem,
				},
				GCDCooldown: time.Second,
			},
		},
		SpellHitEffect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				IgnoreHitCheck:         true,
			},
			DotInput: core.DotDamageInput{
				NumberOfTicks:        30,
				TickLength:           time.Second * 2,
				TickBaseDamage:       58,
				TickSpellCoefficient: 0.167,
				TicksCanMissAndCrit:  true,
			},
		},
	}

	spell.DamageMultiplier *= 1 + float64(shaman.Talents.CallOfFlame)*0.05

	return core.NewSimpleSpellTemplate(spell)
}

func (shaman *Shaman) NewSearingTotem(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	searingTotem := &shaman.FireTotemSpell
	shaman.searingTotemTemplate.Apply(searingTotem)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	searingTotem.Target = target
	searingTotem.Init(sim)

	return searingTotem
}

const SpellIDMagmaTotem int32 = 25552

// This is probably not worth simming since no other spell in the game does this and AM isn't
// even a popular choice for arcane mages.
func (shaman *Shaman) newMagmaTotemTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				Name:           "Magma Totem",
				CritMultiplier: 1.5,
				SpellSchool:    stats.FireSpellPower,
				Character:      &shaman.Character,
				BaseManaCost:   800,
				ManaCost:       800,
				ActionID: core.ActionID{
					SpellID: SpellIDMagmaTotem,
				},
				GCDCooldown: time.Second,
			},
		},
		SpellHitEffect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				IgnoreHitCheck:         true,
			},
			DotInput: core.DotDamageInput{
				NumberOfTicks:        10,
				TickLength:           time.Second * 2,
				TickBaseDamage:       97,
				TickSpellCoefficient: 0.067,
				TicksCanMissAndCrit:  true,
			},
		},
	}

	spell.DamageMultiplier *= 1 + float64(shaman.Talents.CallOfFlame)*0.05

	return core.NewSimpleSpellTemplate(spell)
}

func (shaman *Shaman) NewMagmaTotem(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	magmaTotem := &shaman.FireTotemSpell
	shaman.magmaTotemTemplate.Apply(magmaTotem)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	magmaTotem.Target = target
	magmaTotem.Init(sim)

	return magmaTotem
}

const SpellIDNovaTotem int32 = 25537

var CooldownIDNovaTotem = core.NewCooldownID()

// This is probably not worth simming since no other spell in the game does this and AM isn't
// even a popular choice for arcane mages.
func (shaman *Shaman) newNovaTotemTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				Name:           "Fire Nova Totem",
				CritMultiplier: 1.5,
				SpellSchool:    stats.FireSpellPower,
				Character:      &shaman.Character,
				BaseManaCost:   800,
				ManaCost:       800,
				ActionID: core.ActionID{
					SpellID:    SpellIDNovaTotem,
					CooldownID: CooldownIDNovaTotem,
				},
				Cooldown:    time.Second * 15,
				GCDCooldown: time.Second,
			},
		},
		SpellHitEffect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				IgnoreHitCheck:         true,
			},
			DotInput: core.DotDamageInput{
				NumberOfTicks:        1,
				TickLength:           time.Second * 5,
				TickBaseDamage:       692,
				TickSpellCoefficient: 0.214,
				TicksCanMissAndCrit:  true,
			},
		},
	}

	spell.DamageMultiplier *= 1 + float64(shaman.Talents.CallOfFlame)*0.05
	spell.DotInput.TickLength -= time.Duration(shaman.Talents.ImprovedFireTotems) * time.Second

	return core.NewSimpleSpellTemplate(spell)
}

func (shaman *Shaman) NewNovaTotem(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	novaTotem := &shaman.FireTotemSpell
	shaman.novaTotemTemplate.Apply(novaTotem)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	novaTotem.Target = target
	novaTotem.Init(sim)

	return novaTotem
}
