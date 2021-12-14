package shadow

import (
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
	"github.com/wowsims/tbc/sim/priest"
)

func RegisterShadowPriest() {
	core.RegisterAgentFactory(proto.Player_ShadowPriest{}, func(character core.Character, options proto.Player) core.Agent {
		return NewShadowPriest(character, options)
	})
}

var ShadowWeavingDebuffID = core.NewDebuffID()
var ShadowWeaverAuraID = core.NewAuraID()

func NewShadowPriest(character core.Character, options proto.Player) *ShadowPriest {
	shadowOptions := options.GetShadowPriest()

	// Only undead can do Dev Plague
	if shadowOptions.Rotation.UseDevPlague && options.Race != proto.Race_RaceUndead {
		shadowOptions.Rotation.UseDevPlague = false
	}
	// Only nelf can do starshards
	if shadowOptions.Rotation.UseStarshards && options.Race != proto.Race_RaceNightElf {
		shadowOptions.Rotation.UseStarshards = false
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

// TODO: probably do something different instead of making it global?
const (
	mbidx int = iota
	swdidx
	vtidx
	swpidx
)

func (spriest *ShadowPriest) Act(sim *core.Simulation) time.Duration {
	if spriest.rotation.PrecastVt && sim.CurrentTime == 0 {
		spell := spriest.NewVT(sim, sim.GetPrimaryTarget())
		spell.CastTime = 0
		spell.IgnoreCooldowns = true
		spell.Cast(sim)
	}

	// This if block is to handle being able to cast a VT while having another one ticking.
	//  This will swap the casting spell if it is ticking, so the newly cast spell is now the ticking spell.
	//  spriest.NewVT() will always use the priest.VTSpellCasting as the target.
	if spriest.VTSpellCasting.DotInput.IsTicking(sim) {
		if spriest.VTSpell.DotInput.TimeRemaining(sim) > 0 {
			// If we still have VT ticking that isn't allowed... crash immediately so we can fix the logic.
			panic("never should have two copies ticking")
		}
		oldVT := spriest.VTSpell
		spriest.VTSpell = spriest.VTSpellCasting
		spriest.VTSpellCasting = oldVT // will probably have one more tick
	}

	// Activate shared behaviors
	target := sim.GetPrimaryTarget()
	var spell *core.SimpleSpell
	var wait1 time.Duration
	var wait2 time.Duration
	var wait time.Duration

	// calculate how much time a VT cast would take so we can possibly start casting right before the dot is up.
	vtCastTime := time.Duration(float64(time.Millisecond*1500) / spriest.CastSpeed())

	// timeForDots := sim.Duration-sim.CurrentTime > time.Second*12
	// TODO: stop casting dots near the end?

	if spriest.UseShadowfiend &&
		spriest.CurrentMana()/spriest.MaxMana() < 0.5 &&
		spriest.GetRemainingCD(priest.ShadowfiendCD, sim.CurrentTime) == 0 {
		spell = spriest.NewShadowfiend(sim, target)
	} else if spriest.Talents.VampiricTouch && spriest.VTSpell.DotInput.TimeRemaining(sim) < vtCastTime {
		spell = spriest.NewVT(sim, target)
	} else if !spriest.SWPSpell.DotInput.IsTicking(sim) {
		spell = spriest.NewSWP(sim, target)
	} else if spriest.rotation.UseStarshards && spriest.GetRemainingCD(priest.SSCooldownID, sim.CurrentTime) == 0 {
		spell = spriest.NewStarshards(sim, target)
	} else if spriest.rotation.UseDevPlague && spriest.GetRemainingCD(priest.DPCooldownID, sim.CurrentTime) == 0 {
		spell = spriest.NewDevouringPlague(sim, target)
	} else if spriest.Talents.MindFlay {

		allCDs := []time.Duration{
			mbidx:  spriest.Character.GetRemainingCD(priest.MBCooldownID, sim.CurrentTime),
			swdidx: spriest.Character.GetRemainingCD(priest.SWDCooldownID, sim.CurrentTime),
			vtidx:  spriest.VTSpell.DotInput.TimeRemaining(sim) - vtCastTime,
			swpidx: spriest.SWPSpell.DotInput.TimeRemaining(sim),
		}

		if allCDs[mbidx] == 0 {
			if spriest.Talents.InnerFocus && spriest.GetRemainingCD(priest.InnerFocusCooldownID, sim.CurrentTime) == 0 {
				priest.ApplyInnerFocus(sim, &spriest.Priest)
			}
			spell = spriest.NewMindBlast(sim, target)
		} else if allCDs[swdidx] == 0 {
			spell = spriest.NewSWD(sim, target)
		} else {
			spell = spriest.NewMindFlay(sim, target)

			gcd := spell.CalculatedGCD(&spriest.Character)
			switch spriest.rotation.RotationType {
			case proto.ShadowPriest_Rotation_Ideal:
				// PerfectMindflayRotation to modify how many mindflay ticks to perform.
				wait = spriest.IdealMindflayRotation(sim, spell, allCDs, gcd)
			case proto.ShadowPriest_Rotation_Clipping:
				wait = spriest.ClippingMindflayRotation(sim, spell, allCDs, gcd)
			case proto.ShadowPriest_Rotation_Basic:
				// just do MF3, never clipping
				nextCD := core.NeverExpires
				for _, v := range allCDs {
					if v < nextCD {
						nextCD = v
					}
				}
				// But don't start a MF if we can't get a single tick off.
				if nextCD < gcd {
					spell.DotInput.NumberOfTicks = 0
					wait = nextCD + 1
				} else {
					wait = time.Duration(spell.DotInput.NumberOfTicks) * spell.DotInput.TickLength
				}

			}
			if sim.Log != nil {
				sim.Log("<spriest> Selected %d mindflay ticks.\n", spell.DotInput.NumberOfTicks)
			}
			if spell.DotInput.NumberOfTicks == 0 {
				spell.Cancel()
				return sim.CurrentTime + core.MaxDuration(spriest.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime), wait)
			}
			// if our channel is longer than GCD it will have human latency to end it beause you can't queue the next spell.
			if wait > gcd && spriest.rotation.Latency > 0 {
				base := spriest.rotation.Latency * 0.66
				variation := base + sim.RandomFloat("spriest latency")*base // should vary from 0.66 - 1.33 of given latency

				const minimumLatencyMS = 10
				if variation < minimumLatencyMS {
					variation = minimumLatencyMS // no player can go under XXXms response time
				}
				wait += time.Duration(variation) * time.Millisecond
			}
		}
	} else {
		// what do you even do... i guess just sit around
		mbcd := spriest.Character.GetRemainingCD(priest.MBCooldownID, sim.CurrentTime)
		swdcd := spriest.Character.GetRemainingCD(priest.SWDCooldownID, sim.CurrentTime)
		vtidx := spriest.VTSpell.DotInput.TimeRemaining(sim) - vtCastTime
		swpidx := spriest.SWPSpell.DotInput.TimeRemaining(sim)
		wait1 = core.MinDuration(mbcd, swdcd)
		wait2 = core.MinDuration(vtidx, swpidx)
		wait = core.MinDuration(wait1, wait2)
	}

	actionSuccessful := spell.Cast(sim)

	// fmt.Printf("\tCasting: %s, %0.2f\n", spell.Name, spell.CastTime.Seconds())
	if !actionSuccessful {
		regenTime := spriest.TimeUntilManaRegen(spell.GetManaCost())
		if sim.Log != nil {
			sim.Log("<spriest> Not enough mana, regenerating for %s.\n", regenTime)
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

// IdealMindflayRotation will calculate how many ticks should be cast and mutate the cast.
//  It will calculate the DPS difference between MF and other pending CDs and select clipping based on that.
func (spriest *ShadowPriest) IdealMindflayRotation(sim *core.Simulation, spell *core.SimpleSpell, allCDs []time.Duration, gcd time.Duration) time.Duration {
	nextCD := core.NeverExpires
	nextIdx := -1
	for i, v := range allCDs {
		if v < nextCD {
			nextCD = v
			nextIdx = i
		}
	}

	var numTicks int
	var Major_dmg float64
	// Add a millisecond as fudge factor to these checks
	//  sometimes we wait 1ns extra from waiting for a dot tick to finish.
	//  these can cause some CDs to be slightly shorter than they should be.
	if nextCD <= gcd {
		numTicks = int((nextCD + time.Millisecond) / gcd)
	} else {
		numTicks = int((nextCD + time.Millisecond) / spell.DotInput.TickLength)
	}

	if numTicks == 0 {
		spell.DotInput.NumberOfTicks = 0
		crit := spriest.GetStat(stats.SpellCrit) / (core.SpellCritRatingPerCritChance * 100)
		if nextCD <= 0 {
			nextCD = 1 // add a nanosecond to be sure any ticking dot finishes and we don't get sim stuck.
		} else {
			//  calculate the dps gain from casting vs waiting.
			if nextIdx == 0 {
				Major_dmg = (731.5 + spriest.GetStat(stats.SpellPower)*0.429) / (gcd + nextCD).Seconds() * (1 + (0.5 * (crit + 0.15)))
			} else if nextIdx == 1 {
				Major_dmg = (618 + spriest.GetStat(stats.SpellPower)*0.429) / (gcd + nextCD).Seconds() * (1 + (0.5 * (crit + 0.15)))
			} else if nextIdx == 2 {
				Major_dmg = (spriest.VTSpell.DotInput.DamagePerTick() * float64(spriest.VTSpell.DotInput.NumberOfTicks)) / (gcd + nextCD).Seconds()
			} else if nextIdx == 3 {
				Major_dmg = (spriest.SWPSpell.DotInput.DamagePerTick() * float64(spriest.SWPSpell.DotInput.NumberOfTicks)) / (gcd + nextCD).Seconds()
			}

			mfDamage := (528 + 0.57*(spriest.GetStat(stats.SpellPower))) * 0.3333

			dpsPossibleshort := []float64{
				(Major_dmg * float64(nextCD+gcd)) / float64(gcd+nextCD), // dps with no tick and just wait
				0,
				0,
				0,
			}
			dpsPossibleshort[1] = (Major_dmg*(nextCD+gcd).Seconds() + mfDamage) / (gcd + gcd).Seconds()                           // new damage for 1 extra tick
			dpsPossibleshort[2] = (Major_dmg*(nextCD+gcd).Seconds() + 2*mfDamage) / (2*spell.DotInput.TickLength + gcd).Seconds() // new damage for 2 extra tick
			dpsPossibleshort[3] = (Major_dmg*(nextCD+gcd).Seconds() + 3*mfDamage) / (3*spell.DotInput.TickLength + gcd).Seconds() // new damage for 3 extra tick

			// Find the highest possible dps and its index
			highestPossibleIdx := 0
			highestPossibleDmg := 0.0
			if highestPossibleIdx == 0 {
				for i, v := range dpsPossibleshort {
					if v >= highestPossibleDmg {
						highestPossibleIdx = i
						highestPossibleDmg = v
					}
				}
			}

			if highestPossibleIdx == 1 {
				numTicks = numTicks + 1
			} else if highestPossibleIdx == 2 {
				numTicks = numTicks + 2
			} else if highestPossibleIdx == 3 {
				numTicks = numTicks + 3
			} else if highestPossibleIdx == 0 {
				// ind_sp = ind_sp3  // What does this mean??
			}

			var numtickss float64 = float64(numTicks)
			res := math.Mod(numtickss, 3)

			//  Now that the number of optimal ticks has been determined to optimize dps
			//  Now optimize mf2s and mf3s
			if numTicks == 1 {
				spell.DotInput.NumberOfTicks = numTicks
			} else if res < 2 && numTicks < 5 && numTicks != 3 && numTicks != 0 {
				spell.DotInput.NumberOfTicks = 2
			} else if res == 0 && numTicks != 0 {
				//  cast MF3 MF3...MF3
				spell.DotInput.NumberOfTicks = 3
			} else if res == 2 {
				//  cast MF3 MF3...MF2
				if numTicks == 2 {
					spell.DotInput.NumberOfTicks = 2
				} else {
					spell.DotInput.NumberOfTicks = 3
				}
			} else if res < 2 && numTicks > 5 { // % need to optomize between 3 and 2 ticks and not allowing 1 tick
				//  cast MF3 MF3...MF2 MF2
				spell.DotInput.NumberOfTicks = 3
			} else if numTicks == 0 {
				return nextCD
			}
			spell.ActionID.Tag = int32(spell.DotInput.NumberOfTicks)

			return spell.DotInput.TickLength * time.Duration(spell.DotInput.NumberOfTicks)
		}
		return nextCD
	}

	mfTime := time.Duration(numTicks) * spell.DotInput.TickLength
	if mfTime < gcd {
		mfTime = gcd
	}

	cdDiffs := []time.Duration{
		allCDs[0] - mfTime,
		allCDs[1] - mfTime,
		allCDs[2] - mfTime,
		allCDs[3] - mfTime,
	}

	if cdDiffs[0] < 0 {
		cdDiffs[0] = 0
	}
	if cdDiffs[1] < 0 {
		cdDiffs[1] = 0
	}
	if cdDiffs[2] < 0 {
		cdDiffs[2] = 0
	}
	if cdDiffs[3] < 0 {
		cdDiffs[3] = 0
	}

	crit := spriest.GetStat(stats.SpellCrit) / (core.SpellCritRatingPerCritChance * 100)

	spellDamages := []float64{
		mbidx:  (731.5 + spriest.GetStat(stats.SpellPower)*0.429) / (gcd + cdDiffs[mbidx]).Seconds() * (1 + (0.5 * (crit + 0.15))),
		swdidx: (618 + spriest.GetStat(stats.SpellPower)*0.429) / (gcd + cdDiffs[swdidx]).Seconds() * (1 + (0.5 * (crit + 0.15))),
		vtidx:  (spriest.VTSpell.DotInput.DamagePerTick() * float64(spriest.VTSpell.DotInput.NumberOfTicks)) / (gcd + cdDiffs[vtidx]).Seconds(),
		swpidx: (spriest.SWPSpell.DotInput.DamagePerTick() * float64(spriest.SWPSpell.DotInput.NumberOfTicks)) / (gcd + cdDiffs[swpidx]).Seconds(),
	}

	bestIdx := 0
	bestDmg := 0.0
	for i, v := range spellDamages {
		if sim.Log != nil {
			//sim.Log("\tSpellDamages[%d]: %01.f\n", i, v)
			//sim.Log("\tcdDiffs[%d]: %0.1f\n", i, cdDiffs[i].Seconds())
		}
		if v > bestDmg {
			bestIdx = i
			bestDmg = v
		}
	}

	chosenWait := cdDiffs[bestIdx]
	mfDamage := (528 + 0.57*(spriest.GetStat(stats.SpellPower))) * 0.3333

	if nextIdx != bestIdx && chosenWait.Seconds() < 1.49 {
		numTicks = int((allCDs[bestIdx] + time.Millisecond) / spell.DotInput.TickLength)
	}

	if chosenWait > cdDiffs[nextIdx] && cdDiffs[nextIdx].Seconds() < 0.1 {
		chosenWait = cdDiffs[nextIdx]
	}

	finalMFStart := numTicks // Base ticks before adding additional

	highestPossibleIdx := 0
	if (finalMFStart == 2) && (chosenWait < 1000000000 && chosenWait > 999999985) {
		highestPossibleIdx = 1 // if the wait time is equal to an extra mf tick, and there are already 2 ticks, then just add 1
	}
	//sim.Log("CW %d \n", chosenWait)
	dpsPossible := []float64{
		bestDmg, // dps with no tick and just wait
		0,
		0,
		0,
	}

	dpsDuration := float64((chosenWait + gcd).Seconds())
	if highestPossibleIdx == 0 {
		switch finalMFStart {
		case 0:
			// this means that the extra ticks will be relative to starting a new mf cast entirely
			dpsPossible[1] = (bestDmg*dpsDuration + mfDamage) / float64(gcd+gcd)                           // new damage for 1 extra tick
			dpsPossible[2] = (bestDmg*dpsDuration + 2*mfDamage) / float64(2*spell.DotInput.TickLength+gcd) // new damage for 2 extra tick
			dpsPossible[3] = (bestDmg*dpsDuration + 3*mfDamage) / float64(3*spell.DotInput.TickLength+gcd) // new damage for 3 extra tick
		case 1:
			total_check_time := 2 * spell.DotInput.TickLength

			if total_check_time < gcd {
				newDuration := float64((gcd + gcd).Seconds())
				dpsPossible[1] = (bestDmg*dpsDuration + (mfDamage * float64(finalMFStart+1))) / newDuration
			} else {
				newDuration := float64(((total_check_time - gcd) + gcd).Seconds())
				dpsPossible[1] = (bestDmg*dpsDuration + (mfDamage * float64(finalMFStart+1))) / newDuration
			}
			// % check add 2
			total_check_time2 := 2 * spell.DotInput.TickLength
			if total_check_time2 < gcd {
				dpsPossible[2] = (bestDmg*dpsDuration + (mfDamage * float64(finalMFStart+2))) / float64(gcd+gcd)
			} else {
				dpsPossible[2] = (bestDmg*dpsDuration + (mfDamage * float64(finalMFStart+2))) / float64(total_check_time2+gcd)
			}
		case 2:
			// % check add 1
			total_check_time := spell.DotInput.TickLength
			newDuration := float64((total_check_time + gcd).Seconds())
			dpsPossible[1] = (bestDmg*dpsDuration + mfDamage) / newDuration

		default:
			dpsPossible[1] = (bestDmg*dpsDuration + mfDamage) / float64(gcd+gcd)
			if spell.DotInput.TickLength*2 > gcd {
				dpsPossible[2] = (bestDmg*dpsDuration + 2*mfDamage) / float64(2*spell.DotInput.TickLength+gcd)
			} else {
				dpsPossible[2] = (bestDmg*dpsDuration + 2*mfDamage) / float64(gcd+gcd)
			}
			dpsPossible[3] = (bestDmg*dpsDuration + 3*mfDamage) / float64(3*spell.DotInput.TickLength+gcd)
		}
	}

	// Find the highest possible dps and its index
	// highestPossibleIdx := 0
	highestPossibleDmg := 0.0
	if highestPossibleIdx == 0 {
		for i, v := range dpsPossible {
			if sim.Log != nil {
				//sim.Log("\tdpsPossible[%d]: %01.f\n", i, v)
			}
			if v >= highestPossibleDmg {
				highestPossibleIdx = i
				highestPossibleDmg = v
			}
		}
	}

	if highestPossibleIdx == 1 {
		numTicks = numTicks + 1
	} else if highestPossibleIdx == 2 {
		numTicks = numTicks + 2
	} else if highestPossibleIdx == 3 {
		numTicks = numTicks + 3
	} else if highestPossibleIdx == 0 {
		// ind_sp = ind_sp3  // What does this mean??
	}

	var numbticks float64 = float64(numTicks)
	res2 := math.Mod(numbticks, 3)

	//  Now that the number of optimal ticks has been determined to optimize dps
	//  Now optimize mf2s and mf3s
	if numTicks == 1 {
		spell.DotInput.NumberOfTicks = numTicks
	} else if res2 < 2 && numTicks < 5 && numTicks != 3 && numTicks != 0 {
		spell.DotInput.NumberOfTicks = 2
		if sim.Log != nil {
			//sim.Log("Final rotation should be: MF2 MF2\n")
		}
	} else if res2 == 0 && numTicks != 0 {
		//  cast MF3 MF3...MF3
		if sim.Log != nil {
			//sim.Log("Final rotation should be: MF3 MF3 ... MF3\n")
		}
		spell.DotInput.NumberOfTicks = 3
	} else if res2 == 2 {
		//  cast MF3 MF3...MF2
		if sim.Log != nil {
			//sim.Log("Final rotation should be: MF3 MF3 ... MF2\n")
		}
		if numTicks == 2 {
			spell.DotInput.NumberOfTicks = 2
		} else {
			spell.DotInput.NumberOfTicks = 3
		}

	} else if res2 < 2 && numTicks > 5 { // % need to optomize between 3 and 2 ticks and not allowing 1 tick
		//  cast MF3 MF3...MF2 MF2
		if sim.Log != nil {
			//sim.Log("Final rotation should be: MF3 MF3 ... MF2 MF2\n")
		}
		spell.DotInput.NumberOfTicks = 3
	} else if numTicks == 0 {
		return nextCD
	}
	spell.ActionID.Tag = int32(spell.DotInput.NumberOfTicks)

	//  ONE BIG CAVEAT THAT STILL NEEDS WORK.. THIS NEEDS TO BE UPDATED TO INCLUDE HASTE PROCS THAT CAN OCCUR/DROP OFF MID MF SEQUENCE

	return spell.DotInput.TickLength * time.Duration(spell.DotInput.NumberOfTicks)
}

// ClippingMindflayRotation is to be a 'sweaty but not perfect' rotation.
//  it will prioritize casting MB / SWD by clipping.
//  If there is 4s until the next CD it will use a 2xMF2 instead of 3+1
//  This will mutate the input cast to the correct number of ticks.
func (spriest *ShadowPriest) ClippingMindflayRotation(sim *core.Simulation, spell *core.SimpleSpell, allCDs []time.Duration, gcd time.Duration) time.Duration {
	nextCD := core.NeverExpires
	for _, v := range allCDs[:2] {
		if v < nextCD {
			nextCD = v
		}
	}

	if sim.Log != nil {
		sim.Log("<spriest> NextCD: %0.2f\n", nextCD.Seconds())
	}
	// This means a CD is coming up before we could cast a single MF
	if nextCD < gcd {
		spell.DotInput.NumberOfTicks = 0
		if nextCD < 0 {
			nextCD = 0
		}
		return nextCD + 1
	}

	mfTwoTime := 2*spell.DotInput.TickLength + time.Duration(spriest.rotation.Latency)
	mfBaseTime := 3*spell.DotInput.TickLength + time.Duration(spriest.rotation.Latency)
	mfFiveTime := 5*spell.DotInput.TickLength + time.Duration(spriest.rotation.Latency)

	if nextCD >= mfFiveTime {
		spell.DotInput.NumberOfTicks = 3
	} else if nextCD >= mfTwoTime*2 {
		// time for between 4-5 ticks should use 2xMF2
		spell.DotInput.NumberOfTicks = 2
	} else if nextCD >= mfBaseTime {
		spell.DotInput.NumberOfTicks = 3
	} else if nextCD >= mfTwoTime {
		spell.DotInput.NumberOfTicks = 2
	} else {
		// means we can squeeze in a single tick
		spell.DotInput.NumberOfTicks = 1
	}

	spell.ActionID.Tag = int32(spell.DotInput.NumberOfTicks)
	return spell.DotInput.TickLength * time.Duration(spell.DotInput.NumberOfTicks)
}
