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
		const (
			mbidx int = iota
			swdidx
			vtidx
			swpidx
		)
		allCDs := []time.Duration{
			mbidx:  spriest.Character.GetRemainingCD(priest.MBCooldownID, sim.CurrentTime),
			swdidx: spriest.Character.GetRemainingCD(priest.SWDCooldownID, sim.CurrentTime),
			vtidx:  spriest.VTSpell.DotInput.TimeRemaining(sim),
			swpidx: spriest.SWPSpell.DotInput.TimeRemaining(sim),
		}

		if allCDs[0] == 0 {
			if spriest.Talents.InnerFocus && spriest.GetRemainingCD(priest.InnerFocusCooldownID, sim.CurrentTime) == 0 {
				priest.ApplyInnerFocus(sim, &spriest.Priest)
			}
			spell = spriest.NewMindBlast(sim, target)
		} else if allCDs[1] == 0 {
			spell = spriest.NewSWD(sim, target)
		} else {
			spell = spriest.NewMindFlay(sim, target)
			mfLength := spell.DotInput.TickLength * time.Duration(spriest.MindFlaySpell.DotInput.NumberOfTicks)
			wait = mfLength

			var nextCD time.Duration
			for _, v := range allCDs {
				if v < nextCD {
					nextCD = v
				}
			}

			var numTicks int
			if nextCD <= core.GCDDefault {
				numTicks = int(float64(nextCD) / float64(core.GCDDefault))
			} else {
				numTicks = int(float64(nextCD) / float64(spell.DotInput.TickLength))
			}
			mfTime := time.Duration(numTicks) * spell.DotInput.TickLength

			cdDiffs := []time.Duration{
				allCDs[0] - mfTime,
				allCDs[1] - mfTime,
				allCDs[2] - mfTime,
				allCDs[3] - mfTime,
			}

			crit := spriest.GetStat(stats.SpellCrit) / (core.SpellCritRatingPerCritChance * 100)

			spellDamages := []float64{
				mbidx:  (731.5 + spriest.GetStat(stats.SpellPower)*0.429/(core.GCDDefault+cdDiffs[mbidx]).Seconds()) * (1 + (0.5 * (crit + 0.1))),
				swdidx: (618 + spriest.GetStat(stats.SpellPower)*0.429/(core.GCDDefault+cdDiffs[swdidx]).Seconds()) * (1 + (0.5 * (crit + 0.1))),
				vtidx:  (spriest.VTSpell.DotInput.DamagePerTick() * float64(spriest.VTSpell.DotInput.NumberOfTicks)) / (core.GCDDefault + cdDiffs[vtidx]).Seconds(),
				swpidx: (spriest.SWPSpell.DotInput.DamagePerTick() * float64(spriest.SWPSpell.DotInput.NumberOfTicks)) / (core.GCDDefault + cdDiffs[swpidx]).Seconds(),
			}

			bestIdx := 0
			bestDmg := 0.0
			for i, v := range spellDamages {
				if v > bestDmg {
					bestIdx = i
					bestDmg = v
				}
			}
			chosenWait := cdDiffs[bestIdx]
			mfDamage := (528 + 0.57*(spriest.GetStat(stats.SpellPower))) * 0.3333
			finalMFStart := numTicks % 3 // (how many ticks are left over after mf3s are casted repeatedly

			dpsPossible := []float64{
				(M_CD_DMG * (chosenWait + core.GCDDefault)) / (core.GCDDefault + chosenWait), // dps with no tick and just wait
				0,
				0,
				0,
			}
			if numTicks == 0 {
				// what do we do
			}

			switch finalMFStart {
			case 0:
			// this means that the extra ticks will be relative to starting a new mf cast entirely
			//    total_dps__poss1 = (M_CD_DMG*(chosen_wait + gcd) + MF_DMG)./(gcd + gcd) \\ new damage for 1 extra tick
			//    total_dps__poss2 = (M_CD_DMG*(chosen_wait + gcd) + 2*MF_DMG)./(2*spriest.MindFlaySpell.DotInput.TickLength  + gcd) \\ new damage for 2 extra tick
			//    total_dps__poss3 = (M_CD_DMG*(chosen_wait + gcd) + 3*MF_DMG)./(3*spriest.MindFlaySpell.DotInput.TickLength  + gcd) \\ new damage for 2 extra tick

			case 1:
			//      total_check_time = Final_MF_start.*spriest.MindFlaySpell.DotInput.TickLength + spriest.MindFlaySpell.DotInput.TickLength;
			//      if total_check_time < gcd
			//  	  total_dps__poss1 = (M_CD_DMG*(chosen_wait + gcd) + (MF_DMG * (Final_MF_start + 1))) / (gcd + gcd)
			//  	else
			//  	  total_dps__poss1 = (M_CD_DMG*(chosen_wait + gcd) + (MF_DMG * (Final_MF_start + 1))) / (total_check_time + gcd)
			//  	end
			//	% check add 2
			//	    total_check_time2 = Final_MF_start.*spriest.MindFlaySpell.DotInput.TickLength + 2 * spriest.MindFlaySpell.DotInput.TickLength;
			//	    if total_check_time2 < gcd
			//	    total_dps__poss2 =  (M_CD_DMG*(chosen_wait + gcd) + (MF_DMG * (Final_MF_start + 2))) / (gcd + gcd)
			//   	else
			//	    total_dps__poss2 =  (M_CD_DMG*(chosen_wait + gcd) + (MF_DMG * (Final_MF_start + 2))) / (total_check_time2 + gcd)
			//  	end

			default:
				//      total_dps__poss1 = (M_CD_DMG*(chosen_wait + gcd) + MF_DMG) / (gcd + gcd)
				//     if mf_tick_speed * 2 > gcd
				//      total_dps__poss2 = (M_CD_DMG*(chosen_wait + gcd) + 2*MF_DMG) / (2*spriest.MindFlaySpell.DotInput.TickLength  + gcd)
				//     else
				//      total_dps__poss2 = (M_CD_DMG*(chosen_wait + gcd) + 2*MF_DMG) / (gcd  + gcd)
				//     end
				//      total_dps__poss3 = (M_CD_DMG*(chosen_wait + gcd) + 3*MF_DMG) / (3*spriest.MindFlaySpell.DotInput.TickLength  + gcd)

			}

			//
			// 	spell_adps = [total_dps__poss0,total_dps__poss1,total_dps__poss2,total_dps__poss3];
			//
			//  [highest_adps,ind_mfa] = max(abs(spell_adps));
			//  if ind_sp2 == 2
			//   number_of_ticks = number_of_ticks + 1;
			//  elseif ind_sp2 == 3
			//   number_of_ticks = number_of_ticks + 2;
			//  elseif ind_sp2 == 4
			//   number_of_ticks = number_of_ticks + 3;
			//  elseif ind_sp2 == 1
			//   ind_sp = ind_sp3;
			//  end
			//  Now that the number of optimal ticks has been determined to optimize dps
			//  Now optimize mf2s and mf3s
			//  if number_of_ticks == 1;
			//  cast MF1
			//  else if  rem(number_of_ticks,3) < 2 && number_of_ticks < 5 && current_Time + gcd < fight_length && number_of_ticks ~= 3
			//  cast MF2 MF2
			//  else if  rem(number_of_ticks,3) == 0 && current_Time + gcd < fight_length
			//  cast MF3 MF3...MF3
			//  else if  rem(number_of_ticks,3) == 2 && current_Time + gcd < fight_length
			//  cast MF3 MF3...MF2
			//  else if  rem(number_of_ticks,3) < 2 && number_of_ticks > 5 && time + gcd < fight_length % need to optomize between 3 and 2 ticks and not allowing 1 tick
			//  cast MF3 MF3...MF2 MF2
			//  end
			//  ONE BIG CAVEAT THAT STILL NEEDS WORK.. THIS NEEDS TO BE UPDATED TO INCLUDE HASTE PROCS THAT CAN OCCUR/DROP OFF MID MF SEQUENCE

			// TODO: Get rid of this stuff
			switch spriest.rotation.RotationType {
			case proto.ShadowPriest_Rotation_Basic:
				// basic rotation will simply wait the whole length
			case proto.ShadowPriest_Rotation_IntelligentClipping:

				minWait := core.MinDuration(mbcd, swdcd) + 1
				if minWait < mfLength && minWait > (spell.DotInput.TickLength/2) {
					numTicks := int(minWait/spell.DotInput.TickLength) + 1
					if minWait == spell.DotInput.TickLength {
						numTicks = 1
					}
					if numTicks < 4 {
						spriest.MindFlaySpell.DotInput.NumberOfTicks = numTicks
						spell.ActionID.Tag = int32(numTicks)
					}
					wait = spell.DotInput.TickLength * time.Duration(spriest.MindFlaySpell.DotInput.NumberOfTicks)
				} else if minWait < spell.DotInput.TickLength {
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
				if minWait < mfLength && minWait > spell.DotInput.TickLength {
					numTicks := int(float64(core.GCDDefault)/float64(spell.DotInput.TickLength)) + 1
					if numTicks < 4 {
						spriest.MindFlaySpell.DotInput.NumberOfTicks = numTicks
						spell.ActionID.Tag = int32(numTicks)
					}
					wait = spell.DotInput.TickLength * time.Duration(spriest.MindFlaySpell.DotInput.NumberOfTicks)
				} else if minWait < spell.DotInput.TickLength {
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
