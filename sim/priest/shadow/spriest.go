package shadow

import (
	"time"

	"github.com/wowsims/tbc/sim/common"
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

	selfBuffs := priest.SelfBuffs{}

	basePriest := priest.NewPriest(character, selfBuffs, *shadowOptions.Talents)
	spriest := &ShadowPriest{
		Priest:          basePriest,
		primaryRotation: *shadowOptions.Rotation,
	}

	if basePriest.Talents.ShadowWeaving > 0 {
		const dur = time.Second * 15
		const misDur = time.Second * 24

		// This is a combined aura for all spriest major on hit effects.
		//  Shadow Weaving, Vampiric Touch, and Misery
		spriest.Character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
			return core.Aura{
				ID:   ShadowWeaverAuraID,
				Name: "Shadow Weaver",
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					if spellEffect.Damage > 0 && spriest.VTSpell.DotInput.IsTicking(sim) {
						s := stats.Stats{stats.Mana: spellEffect.Damage * 0.05}
						if sim.Log != nil {
							sim.Log("VT Regenerated %0.1f mana.\n", s[stats.Mana])
						}
						spriest.Party.AddStats(s)
					}

					if spriest.swStacks < 5 {
						spriest.swStacks++
						if sim.Log != nil {
							sim.Log("(%d) Shadow Weaving: stacks on target %0.0f\n", spriest.ID, spriest.swStacks)
						}
					}
					// Just keep replacing it with new expire time.
					spellEffect.Target.ReplaceAura(sim, core.Aura{
						ID:      ShadowWeavingDebuffID,
						Name:    "Shadow Weaving",
						Expires: sim.CurrentTime + dur,
						OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
							spellEffect.DamageMultiplier *= 1 + 0.02*spriest.swStacks
						},
						OnExpire: func(sim *core.Simulation) {
							spriest.swStacks = 0
						},
					})

					if spellCast.ActionID.SpellID == priest.SpellIDSWP || spellCast.ActionID.SpellID == priest.SpellIDVT || spellCast.ActionID.SpellID == priest.SpellIDMF {
						spellEffect.Target.ReplaceAura(sim, core.Aura{
							ID:      core.MiseryDebuffID,
							Expires: sim.CurrentTime + misDur,
							Name:    "Misery",
							OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
								spellEffect.DamageMultiplier *= 1.05
							},
						})
					}
				},
			}
		})
	}
	// spriest.PickRotations(*shadowOptions.Rotation, isr)

	return spriest
}

type ShadowPriest struct {
	priest.Priest

	swStacks float64

	primaryRotation proto.ShadowPriest_Rotation
	// These are only used when primary spell is set to 'Adaptive'. When the mana
	// tracker tells us we have extra mana to spare, use surplusRotation instead of
	// primaryRotation.
	useSurplusRotation bool
	surplusRotation    proto.ShadowPriest_Rotation
	manaTracker        common.ManaSpendingRateTracker
}

func (spriest *ShadowPriest) GetPriest() *priest.Priest {
	return &spriest.Priest
}

func (spriest *ShadowPriest) Reset(sim *core.Simulation) {
	if spriest.useSurplusRotation {
		spriest.manaTracker.Reset()
	}
	spriest.Priest.Reset(sim)
	spriest.swStacks = 0
}

func (spriest *ShadowPriest) Act(sim *core.Simulation) time.Duration {
	if spriest.useSurplusRotation {
		spriest.manaTracker.Update(sim, spriest.GetCharacter())
		projectedManaCost := spriest.manaTracker.ProjectedManaCost(sim, spriest.GetCharacter())

		// If we have enough mana to burn, use the surplus rotation.
		if projectedManaCost < spriest.CurrentMana() {
			return spriest.actRotation(sim, spriest.surplusRotation)
		} else {
			return spriest.actRotation(sim, spriest.primaryRotation)
		}
	} else {
		return spriest.actRotation(sim, spriest.primaryRotation)
	}
}

func (spriest *ShadowPriest) actRotation(sim *core.Simulation, rotation proto.ShadowPriest_Rotation) time.Duration {
	// Activate shared druid behaviors
	target := sim.GetPrimaryTarget()
	var spell *core.SimpleSpell
	var wait time.Duration

	if spriest.Talents.VampiricTouch && !spriest.VTSpell.DotInput.IsTicking(sim) {
		spell = spriest.NewVT(sim, target)
	} else if !spriest.SWPSpell.DotInput.IsTicking(sim) {
		spell = spriest.NewSWP(sim, target)
	} else if rotation.UseDevPlague && spriest.Race == proto.Race_RaceUndead {
		// TODO: add dev plague
		panic("not implemented")
	} else if spriest.Talents.MindFlay {
		mbcd := spriest.Character.GetRemainingCD(priest.MBCooldownID, sim.CurrentTime)
		swdcd := spriest.Character.GetRemainingCD(priest.SWDCooldownID, sim.CurrentTime)

		if mbcd == 0 {
			if spriest.Talents.InnerFocus && spriest.GetRemainingCD(priest.InnerFocusCooldownID, sim.CurrentTime) == 0 {
				priest.ApplyInnerFocus(sim, &spriest.Priest)
			}

			spell = spriest.NewMindBlast(sim, target)
		} else if rotation.UseSwd && swdcd == 0 {
			spell = spriest.NewSWD(sim, target)
		} else {
			spell = spriest.NewMindFlay(sim, target)
			mfLength := spriest.MindFlaySpell.DotInput.TickLength * time.Duration(spriest.MindFlaySpell.DotInput.NumberOfTicks)
			wait = mfLength

			switch rotation.Type {
			case proto.ShadowPriest_Rotation_Basic:
				// basic rotation will simply wait the whole length
			case proto.ShadowPriest_Rotation_ClipAlways:
				// Prio MindBlast
				minWait := core.MinDuration(mbcd, swdcd)
				if minWait < mfLength && minWait > 1 {
					spell.DotInput.NumberOfTicks = int((mfLength - minWait).Seconds()) // cut fractional seconds off
					wait = minWait
				} else if minWait < 1 {
					return sim.CurrentTime + minWait // just wait until its off CD.. dont cast a spell for no reason
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
		return sim.CurrentTime + wait
	}

	return sim.CurrentTime + core.MaxDuration(
		spriest.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime),
		spell.CastTime)
}
