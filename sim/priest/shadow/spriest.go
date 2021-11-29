package shadow

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
	"github.com/wowsims/tbc/sim/priest"
)

func RegisterShadowPriest() {
	core.RegisterAgentFactory(proto.PlayerOptions_ShadowPriest{}, func(character core.Character, options proto.PlayerOptions, isr proto.IndividualSimRequest) core.Agent {
		return NewShadowPriest(character, options, isr)
	})
}

var ShadowWeavingDebuffID = core.NewDebuffID()
var ShadowWeaverAuraID = core.NewAuraID()

func NewShadowPriest(character core.Character, options proto.PlayerOptions, isr proto.IndividualSimRequest) *ShadowPriest {
	shadowOptions := options.GetShadowPriest()

	// Only undead can do Dev Plague
	if shadowOptions.Rotation.UseDevPlague == true && options.Race != proto.Race_RaceUndead {
		shadowOptions.Rotation.UseDevPlague = false
	}

	selfBuffs := priest.SelfBuffs{
		UseShadowfiend: shadowOptions.Options.UseShadowfiend,
	}

	basePriest := priest.New(character, selfBuffs, *shadowOptions.Talents)
	spriest := &ShadowPriest{
		Priest:   basePriest,
		rotation: *shadowOptions.Rotation,
	}

	if basePriest.Talents.ShadowWeaving > 0 {
		const dur = time.Second * 15
		const misDur = time.Second * 24

		// TODO: use Aura.Stacks and add a function to increment stacks.
		//  This is required to make this work correctly for a raid sim.
		swAura := core.Aura{
			ID:   ShadowWeavingDebuffID,
			Name: "Shadow Weaving",
			OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				spellEffect.DamageMultiplier *= 1 + 0.02*spriest.swStacks
			},
			OnPeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage *float64) {
				*tickDamage *= (1 + 0.02*spriest.swStacks)
			},
			OnExpire: func(sim *core.Simulation) {
				spriest.swStacks = 0
			},
		}

		addShadowWeaving := func(sim *core.Simulation, spellEffect *core.SpellEffect) {
			if spriest.swStacks < 5 {
				spriest.swStacks++
				if sim.Log != nil {
					sim.Log("(%d) Shadow Weaving: stacks on target %0.0f\n", spriest.ID, spriest.swStacks)
				}
			}
			// Just keep replacing it with new expire time.
			swAura.Expires = sim.CurrentTime + dur
			spellEffect.Target.ReplaceAura(sim, swAura)
		}
		// This is a combined aura for all spriest major on hit effects.
		//  Shadow Weaving, Vampiric Touch, and Misery
		spriest.Character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
			return core.Aura{
				ID:   ShadowWeaverAuraID,
				Name: "Shadow Weaver",
				OnPeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage *float64) {
					if *tickDamage > 0 && spriest.VTSpell.DotInput.IsTicking(sim) {
						s := stats.Stats{stats.Mana: *tickDamage * 0.05}
						if sim.Log != nil {
							sim.Log("VT Regenerated %0f mana.\n", s[stats.Mana])
						}
						spriest.Party.AddStats(s)
					}
				},
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					addShadowWeaving(sim, spellEffect)
					if spellEffect.Damage > 0 && spriest.VTSpell.DotInput.IsTicking(sim) {
						s := stats.Stats{stats.Mana: spellEffect.Damage * 0.05}
						if sim.Log != nil {
							sim.Log("VT Regenerated %0.1f mana.\n", s[stats.Mana])
						}
						spriest.Party.AddStats(s)
					}

					if spellCast.ActionID.SpellID == priest.SpellIDSWP || spellCast.ActionID.SpellID == priest.SpellIDVT || spellCast.ActionID.SpellID == priest.SpellIDMF {
						spellEffect.Target.ReplaceAura(sim, core.Aura{
							ID:      core.MiseryDebuffID,
							Expires: sim.CurrentTime + misDur,
							Name:    "Misery",
							OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
								spellEffect.DamageMultiplier *= 1.05
							},
							OnPeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage *float64) {
								*tickDamage *= 1.05
							},
						})
					}
				},
			}
		})
	}

	return spriest
}

type ShadowPriest struct {
	priest.Priest

	swStacks float64

	rotation proto.ShadowPriest_Rotation
}

func (spriest *ShadowPriest) GetPriest() *priest.Priest {
	return &spriest.Priest
}

func (spriest *ShadowPriest) Reset(sim *core.Simulation) {
	spriest.Priest.Reset(sim)
	spriest.swStacks = 0
}

func (spriest *ShadowPriest) Act(sim *core.Simulation) time.Duration {
	// Activate shared behaviors
	target := sim.GetPrimaryTarget()
	var spell *core.SimpleSpell
	var wait time.Duration

	// need to add something here to track spell cool downs and remaining dot durations
	// calculate dps of each possible spell as (DMG)/(remaining_cd + gcd)
	// find maximum dps spell to cast next
	// if max dps spell's wait time > gcd, then find the second highest dps spell to cast
	// if max dps spell's wait time > 0.9 seconds (TBD), then check to see if adding a mf filler is more dps
	// option 1: (Spell_dmg)/(wait_time + gcd) ; option 2: (Spell_dmg + mf_dmg)/(gcd + gcd)
	// if option 2 is chosen, then enter into the mf routine to optimize sequence

	timeForDots := sim.Duration-sim.CurrentTime > time.Second*12
	if spriest.UseShadowfiend &&
		spriest.CurrentMana()/spriest.MaxMana() < 0.5 &&
		spriest.GetRemainingCD(priest.ShadowfiendCD, sim.CurrentTime) == 0 {
		spell = spriest.NewShadowfiend(sim, target)
	} else if spriest.Talents.VampiricTouch && timeForDots && !spriest.VTSpell.DotInput.IsTicking(sim) {
		spell = spriest.NewVT(sim, target)
	} else if timeForDots && !spriest.SWPSpell.DotInput.IsTicking(sim) {
		spell = spriest.NewSWP(sim, target)
	} else if spriest.rotation.UseDevPlague && timeForDots && spriest.GetRemainingCD(priest.DPCooldownID, sim.CurrentTime) == 0 {
		spell = spriest.NewDevouringPlague(sim, target)
	} else if spriest.Talents.MindFlay {
		mbcd := spriest.Character.GetRemainingCD(priest.MBCooldownID, sim.CurrentTime)
		swdcd := spriest.Character.GetRemainingCD(priest.SWDCooldownID, sim.CurrentTime)

		if mbcd == 0 {
			if spriest.Talents.InnerFocus && spriest.GetRemainingCD(priest.InnerFocusCooldownID, sim.CurrentTime) == 0 {
				priest.ApplyInnerFocus(sim, &spriest.Priest)
			}
			spell = spriest.NewMindBlast(sim, target)
		} else if swdcd == 0 {
			spell = spriest.NewSWD(sim, target)
		} else {
			spell = spriest.NewMindFlay(sim, target)
			mfLength := spriest.MindFlaySpell.DotInput.TickLength * time.Duration(spriest.MindFlaySpell.DotInput.NumberOfTicks)
			wait = mfLength

			switch spriest.rotation.RotationType {
			case proto.ShadowPriest_Rotation_Basic:
				// basic rotation will simply wait the whole length
			case proto.ShadowPriest_Rotation_IntelligentClipping:

				minWait := core.MinDuration(mbcd, swdcd) + 1
				if minWait < mfLength && minWait > (spriest.MindFlaySpell.DotInput.TickLength/2) {
					numTicks := int(minWait/spriest.MindFlaySpell.DotInput.TickLength) + 1
					if minWait == spriest.MindFlaySpell.DotInput.TickLength {
						numTicks = 1
					}
					if numTicks < 4 {
						spriest.MindFlaySpell.DotInput.NumberOfTicks = numTicks
						spell.ActionID.Tag = int32(numTicks)
					}
					wait = spriest.MindFlaySpell.DotInput.TickLength * time.Duration(spriest.MindFlaySpell.DotInput.NumberOfTicks)
				} else if minWait < spriest.MindFlaySpell.DotInput.TickLength {
					// just wait until its off CD.. dont cast a spell for no reason
					spell.Cancel() // turn off 'in use'
					return sim.CurrentTime + core.MaxDuration(
						spriest.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime),
						minWait)
				}
			case proto.ShadowPriest_Rotation_ClipAlways:
				// Prio MindBlast/SWD

				// TODO: also account for dots falling off.
				minWait := core.MinDuration(mbcd, swdcd) + 1
				if minWait < mfLength && minWait > spriest.MindFlaySpell.DotInput.TickLength {
					numTicks := int(float64(core.GCDDefault)/float64(spriest.MindFlaySpell.DotInput.TickLength)) + 1
					if numTicks < 4 {
						spriest.MindFlaySpell.DotInput.NumberOfTicks = numTicks
						spell.ActionID.Tag = int32(numTicks)
					}
					wait = spriest.MindFlaySpell.DotInput.TickLength * time.Duration(spriest.MindFlaySpell.DotInput.NumberOfTicks)
				} else if minWait < spriest.MindFlaySpell.DotInput.TickLength {
					// just wait until its off CD.. dont cast a spell for no reason
					spell.Cancel() // turn off 'in use'
					return sim.CurrentTime + core.MaxDuration(
						spriest.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime),
						minWait)
				}
			default:
				panic("not implemented")
			}
		}
	} else {
		// what do you even do... i guess just sit around
		mbcd := spriest.Character.GetRemainingCD(priest.MBCooldownID, sim.CurrentTime)
		swdcd := spriest.Character.GetRemainingCD(priest.SWDCooldownID, sim.CurrentTime)
		wait = core.MinDuration(mbcd, swdcd)
	}

	actionSuccessful := spell.Cast(sim)

	if !actionSuccessful {
		regenTime := spriest.TimeUntilManaRegen(spell.GetManaCost())
		if sim.Log != nil {
			sim.Log("Not enough mana, regenerating for %s.\n", regenTime)
		}
		return sim.CurrentTime + regenTime
	}
	if wait != 0 {
		return sim.CurrentTime + core.MaxDuration(spriest.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime), wait)
	}

	return sim.CurrentTime + core.MaxDuration(
		spriest.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime),
		spell.CastTime)
}
