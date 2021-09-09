package core

import (
	"math"
	"time"
)

// AuraEffects will mutate a cast or simulation state.
type AuraEffect func(sim *Simulation, c *Cast)

const neverExpires = time.Duration(math.MaxInt64)

type Aura struct {
	ID          int32
	Expires     time.Duration // time at which aura will be removed
	activeIndex int32         // Position of this aura's index in the sim.activeAuraIDs array

	// The number of stacks, or charges, of this aura. If this aura doesn't care
	// about charges, is just 0.
	stacks int32

	OnCast         AuraEffect
	OnCastComplete AuraEffect
	OnSpellHit     AuraEffect
	OnSpellMiss    AuraEffect
	OnExpire       AuraEffect
}

func AuraName(a int32) string {
	switch a {
	case MagicIDUnknown:
		return "Unknown"
	case MagicIDClBounce:
		return "Chain Lightning Bounce"
	case MagicIDLOTalent:
		return "Lightning Overload Talent"
	case MagicIDJoW:
		return "Judgement Of Wisdom Aura"
	case MagicIDEleFocus:
		return "Elemental Focus"
	case MagicIDEleMastery:
		return "Elemental Mastery"
	case MagicIDStormcaller:
		return "Stormcaller"
	case MagicIDBlessingSilverCrescent:
		return "Blessing of the Silver Crescent"
	case MagicIDDarkIronPipeweed:
		return "Dark Iron Pipeweed"
	case MagicIDQuagsEye:
		return "Quags Eye"
	case MagicIDFungalFrenzy:
		return "Fungal Frenzy"
	case MagicIDBloodlust:
		return "Bloodlust"
	case MagicIDSkycall:
		return "Skycall"
	case MagicIDEnergized:
		return "Energized"
	case MagicIDNAC:
		return "Nature Alignment Crystal"
	case MagicIDChaoticSkyfire:
		return "Chaotic Skyfire"
	case MagicIDInsightfulEarthstorm:
		return "Insightful Earthstorm"
	case MagicIDMysticSkyfire:
		return "Mystic Skyfire"
	case MagicIDMysticFocus:
		return "Mystic Focus"
	case MagicIDEmberSkyfire:
		return "Ember Skyfire"
	case MagicIDLB12:
		return "LB12"
	case MagicIDCL6:
		return "CL6"
	case MagicIDTLCLB:
		return "TLC-LB"
	case MagicIDISCTrink:
		return "Icon Trinket"
	case MagicIDNACTrink:
		return "NAC Trinket"
	case MagicIDDITrink:
		return "Dark Iron Trinket"
	case MagicIDPotion:
		return "Potion"
	case MagicIDRune:
		return "Rune"
	case MagicIDAllTrinket:
		return "AllTrinket"
	case MagicIDSpellPower:
		return "SpellPower"
	case MagicIDRubySerpent:
		return "RubySerpent"
	case MagicIDCallOfTheNexus:
		return "CallOfTheNexus"
	case MagicIDDCC:
		return "Darkmoon Card Crusade"
	case MagicIDDCCBonus:
		return "Aura of the Crusade"
	case MagicIDScryerTrink:
		return "Scryer Trinket"
	case MagicIDRubySerpentTrink:
		return "Ruby Serpent Trinket"
	case MagicIDXiriTrink:
		return "Xiri Trinket"
	case MagicIDDrums:
		return "Drums of Battle"
	case MagicIDNetherstrike:
		return "Netherstrike Set"
	case MagicIDTwinStars:
		return "Twin Stars Set"
	case MagicIDTidefury:
		return "Tidefury Set"
	case MagicIDSpellstrike:
		return "Spellstrike Set"
	case MagicIDSpellstrikeInfusion:
		return "Spellstrike Infusion"
	case MagicIDManaEtched:
		return "Mana-Etched Set"
	case MagicIDManaEtchedHit:
		return "Mana-EtchedHit"
	case MagicIDManaEtchedInsight:
		return "Mana-EtchedInsight"
	case MagicIDWindhawk:
		return "Windhawk Set Bonus"
	case MagicIDOrcBloodFury:
		return "Orc Blood Fury"
	case MagicIDTrollBerserking:
		return "Troll Berserking"
	case MagicIDEyeOfTheNight:
		return "EyeOfTheNight"
	case MagicIDChainTO:
		return "Chain of the Twilight Owl"
	case MagicIDCyclone2pc:
		return "Cyclone 2pc Bonus"
	case MagicIDCyclone4pc:
		return "Cyclone 4pc Bonus"
	case MagicIDCycloneMana:
		return "Cyclone Mana Cost Reduction"
	case MagicIDTLC:
		return "The Lightning Capacitor Aura"
	case MagicIDDestructionPotion:
		return "Destruction Potion"
	case MagicIDHexShunkHead:
		return "Hex Shunken Head"
	case MagicIDShiftingNaaru:
		return "Shifting Naaru Sliver"
	case MagicIDSkullGuldan:
		return "Skull of Guldan"
	case MagicIDNexusHorn:
		return "Nexus-Horn"
	case MagicIDSextant:
		return "Sextant of Unstable Currents"
	case MagicIDUnstableCurrents:
		return "Unstable Currents"
	case MagicIDEyeOfMag:
		return "Eye Of Mag"
	case MagicIDRecurringPower:
		return "Recurring Power"
	case MagicIDCataclysm4pc:
		return "Cataclysm 4pc Set Bonus"
	case MagicIDSkyshatter2pc:
		return "Skyshatter 2pc Set Bonus"
	case MagicIDSkyshatter4pc:
		return "Skyshatter 4pc Set Bonus"
	case MagicIDTotemOfPulsingEarth:
		return "Totem of Pulsing Earth"
	case MagicIDEssMartyrTrink:
		return "Essence of the Martyr Trinket"
	case MagicIDEssSappTrink:
		return "Restrained Essence of Sapphiron Trinket"
	case MagicIDMisery:
		return "Misery"
	case MagicIDElderScribe:
		return "Robes of the Elder Scribe"
	case MagicIDElderScribeProc:
		return "Power of Arcanagos"
	case MagicIDEyeOfTheNightTrink:
		return "EyeOfTheNight Trinket CD"
	case MagicIDChainTOTrink:
		return "ChainTO Trinket CD"
	case MagicIDHexTrink:
		return "Hex Trinket CD"
	case MagicIDShiftingNaaruTrink:
		return "ShiftingNaaru Trinket CD"
	case MagicIDSkullGuldanTrink:
		return "SkullGuldan Trinket CD"
	case MagicIDRegainMana:
		return "Fathom-Brooch Regain Mana"
	case MagicIDAlchStone:
		return "Alchemist's Stone"
	}

	return "<<TODO: Add Aura name to switch!!>>"
}

// Stored value is the time at which the ICD will be off CD
type InternalCD time.Duration

func (icd InternalCD) isOnCD(sim *Simulation) bool {
	return time.Duration(icd) > sim.CurrentTime
}

func NewICD() InternalCD {
	return InternalCD(0)
}

// List of all magic effects and spells and items and stuff that can go on CD or have an aura.
const (
	MagicIDUnknown int32 = iota
	//Spells
	MagicIDLB12
	MagicIDCL6
	MagicIDTLCLB

	// Auras
	MagicIDClBounce
	MagicIDLOTalent
	MagicIDJoW
	MagicIDEleFocus
	MagicIDEleMastery
	MagicIDStormcaller
	MagicIDBlessingSilverCrescent
	MagicIDDarkIronPipeweed
	MagicIDQuagsEye
	MagicIDFungalFrenzy
	MagicIDBloodlust
	MagicIDSkycall
	MagicIDEnergized
	MagicIDNAC
	MagicIDChaoticSkyfire
	MagicIDInsightfulEarthstorm
	MagicIDMysticSkyfire
	MagicIDMysticFocus
	MagicIDEmberSkyfire
	MagicIDSpellPower
	MagicIDRubySerpent
	MagicIDCallOfTheNexus
	MagicIDDCC
	MagicIDDCCBonus
	MagicIDDrums // drums effect
	MagicIDNetherstrike
	MagicIDTwinStars
	MagicIDTidefury
	MagicIDSpellstrike
	MagicIDSpellstrikeInfusion
	MagicIDManaEtched
	MagicIDManaEtchedHit
	MagicIDManaEtchedInsight
	MagicIDMisery
	MagicIDEyeOfTheNight
	MagicIDChainTO
	MagicIDCyclone2pc
	MagicIDCyclone4pc
	MagicIDCycloneMana // proc from 4pc
	MagicIDWindhawk
	MagicIDOrcBloodFury    // orc racials
	MagicIDTrollBerserking // troll racial
	MagicIDTLC             // aura on equip of TLC, stores charges
	MagicIDDestructionPotion
	MagicIDHexShunkHead
	MagicIDShiftingNaaru
	MagicIDSkullGuldan
	MagicIDNexusHorn
	MagicIDSextant          // Trinket Aura
	MagicIDUnstableCurrents // Sextant Proc Aura
	MagicIDEyeOfMag         // trinket aura
	MagicIDRecurringPower   // eye of mag proc aura
	MagicIDCataclysm4pc     // cyclone 4pc aura
	MagicIDSkyshatter2pc    // skyshatter 2pc aura
	MagicIDSkyshatter4pc    // skyshatter 4pc aura
	MagicIDElderScribe      // elder scribe robe item aura
	MagicIDElderScribeProc  // elder scribe robe temp buff
	MagicIDTotemOfPulsingEarth
	MagicIDRegainMana // effect from fathom brooch
	MagicIDAlchStone

	// Items  (Usually individual trinket CDs)
	MagicIDISCTrink
	MagicIDNACTrink
	MagicIDPotion
	MagicIDRune
	MagicIDAllTrinket
	MagicIDScryerTrink
	MagicIDRubySerpentTrink
	MagicIDXiriTrink
	MagicIDEyeOfTheNightTrink
	MagicIDChainTOTrink
	MagicIDHexTrink
	MagicIDShiftingNaaruTrink
	MagicIDSkullGuldanTrink
	MagicIDEssMartyrTrink
	MagicIDEssSappTrink
	MagicIDDITrink // Dark Iron pipe trinket CD

	// Always at end so we know how many magic IDs there are.
	MagicIDLen
)

func ActivateChainLightningBounce(sim *Simulation) Aura {
	return Aura{
		ID:      MagicIDClBounce,
		Expires: neverExpires,
		OnCastComplete: func(sim *Simulation, c *Cast) {
			if c.Spell.ID != MagicIDCL6 || c.IsClBounce {
				return
			}

			dmgCoeff := 1.0
			if c.IsLO {
				dmgCoeff = 0.5
			}
			for i := 1; i < sim.Options.Encounter.NumTargets; i++ {
				if sim.hasAura(MagicIDTidefury) {
					dmgCoeff *= 0.83
				} else {
					dmgCoeff *= 0.7
				}
				clone := &Cast{
					IsLO:       c.IsLO,
					IsClBounce: true,
					Spell:      c.Spell,
					Crit:       c.Crit,
					CritBonus:  c.CritBonus,
					Effect:     func(sim *Simulation, c *Cast) { c.DidDmg *= dmgCoeff },
				}
				sim.Cast(clone)
			}
		},
	}
}

func AuraLightningOverload(lvl int) Aura {
	chance := 0.04 * float64(lvl)
	return Aura{
		ID:      MagicIDLOTalent,
		Expires: neverExpires,
		OnSpellHit: func(sim *Simulation, c *Cast) {
			if c.Spell.ID != MagicIDLB12 && c.Spell.ID != MagicIDCL6 {
				return
			}
			if c.IsLO {
				return // can't proc LO on LO
			}
			actualChance := chance
			if c.Spell.ID == MagicIDCL6 {
				actualChance /= 3 // 33% chance of regular for CL LO
			}
			if sim.Rando.Float64() < actualChance {
				if sim.Debug != nil {
					sim.Debug(" +Lightning Overload\n")
				}
				clone := sim.cache.NewCast()
				// Don't set IsClBounce even if this is a bounce, so that the clone does a normal CL and bounces
				clone.IsLO = true
				clone.Spell = c.Spell
				clone.CritBonus = c.CritBonus
				clone.Effect = loDmgMod
				sim.Cast(clone)
			}
		},
	}
}

func loDmgMod(sim *Simulation, c *Cast) {
	c.DidDmg /= 2
}

func AuraJudgementOfWisdom() Aura {
	const mana = 74 / 2 // 50% proc
	return Aura{
		ID:      MagicIDJoW,
		Expires: neverExpires,
		OnSpellHit: func(sim *Simulation, c *Cast) {
			if sim.Debug != nil {
				sim.Debug(" +Judgement Of Wisdom: 37 mana (74 @ 50%% proc)\n")
			}
			sim.CurrentMana += mana
		},
	}
}

func elementalFocusOnCast(sim *Simulation, c *Cast) {
	c.ManaCost *= .6 // reduced by 40%
}

func elementalFocusOnCastComplete(sim *Simulation, c *Cast) {
	if c.ManaCost <= 0 {
		return // Don't consume charges from free spells.
	}

	sim.auras[MagicIDEleFocus].stacks--
	if sim.auras[MagicIDEleFocus].stacks == 0 {
		sim.removeAura(MagicIDEleFocus)
	}
}

func AuraElementalFocus(sim *Simulation) Aura {
	return Aura{
		ID:             MagicIDEleFocus,
		Expires:        sim.CurrentTime + time.Second*15,
		stacks:         2,
		OnCast:         elementalFocusOnCast,
		OnCastComplete: elementalFocusOnCastComplete,
	}
}

func TryActivateEleMastery(sim *Simulation) {
	if sim.isOnCD(MagicIDEleMastery) || !sim.Options.Talents.ElementalMastery {
		return
	}

	sim.addAura(Aura{
		ID:      MagicIDEleMastery,
		Expires: neverExpires,
		OnCast: func(sim *Simulation, c *Cast) {
			if c.Spell.ID != MagicIDTLCLB {
				c.ManaCost = 0
			}
		},
		OnCastComplete: func(sim *Simulation, c *Cast) {
			c.Crit += 1.01 // 101% chance of crit
			// Remove the buff and put skill on CD
			sim.setCD(MagicIDEleMastery, time.Second*180)
			sim.removeAura(MagicIDEleMastery)
		},
	})
}

func createHasteActivate(id int32, haste float64, duration time.Duration) ItemActivation {
	// Implemented haste activate as a buff so that the creation of a new cast gets the correct cast time
	return func(sim *Simulation) Aura {
		sim.Stats[StatHaste] += haste
		return Aura{
			ID:      id,
			Expires: sim.CurrentTime + duration,
			OnExpire: func(sim *Simulation, c *Cast) {
				sim.Stats[StatHaste] -= haste
			},
		}
	}
}

// createSpellDmgActivate creates a function for trinket activation that uses +spellpower
//  This is so we don't need a separate function for every spell power trinket.
func createSpellDmgActivate(id int32, sp float64, duration time.Duration) ItemActivation {
	return func(sim *Simulation) Aura {
		sim.Stats[StatSpellDmg] += sp
		return Aura{
			ID:      id,
			Expires: sim.CurrentTime + duration,
			OnExpire: func(sim *Simulation, c *Cast) {
				sim.Stats[StatSpellDmg] -= sp
			},
		}
	}
}

func ActivateQuagsEye(sim *Simulation) Aura {
	const hasteBonus = 320.0
	const dur = time.Second * 45
	icd := NewICD()
	return Aura{
		ID:      MagicIDQuagsEye,
		Expires: neverExpires,
		OnCastComplete: func(sim *Simulation, c *Cast) {
			if !icd.isOnCD(sim) && sim.Rando.Float64() < 0.1 {
				icd = InternalCD(sim.CurrentTime + dur)
				sim.Stats[StatHaste] += hasteBonus
				sim.addAura(AuraStatRemoval(sim.CurrentTime, time.Second*6, hasteBonus, StatHaste, MagicIDFungalFrenzy))
			}
		},
	}
}

func ActivateNexusHorn(sim *Simulation) Aura {
	icd := NewICD()
	const spellBonus = 225.0
	const dur = time.Second * 45
	return Aura{
		ID:      MagicIDNexusHorn,
		Expires: neverExpires,
		OnSpellHit: func(sim *Simulation, c *Cast) {
			if !icd.isOnCD(sim) && c.DidCrit && sim.Rando.Float64() < 0.2 {
				icd = InternalCD(sim.CurrentTime + dur)
				sim.Stats[StatSpellDmg] += spellBonus
				sim.addAura(AuraStatRemoval(sim.CurrentTime, time.Second*10, spellBonus, StatSpellDmg, MagicIDCallOfTheNexus))
			}
		},
	}
}

func ActivateDCC(sim *Simulation) Aura {
	const spellBonus = 8.0
	stacks := 0
	return Aura{
		ID:      MagicIDDCC,
		Expires: neverExpires,
		OnCastComplete: func(sim *Simulation, c *Cast) {
			if stacks < 10 {
				stacks++
				sim.Stats[StatSpellDmg] += spellBonus
			}
			// Removal aura will refresh with new total spellpower based on stacks.
			//  This will remove the old stack removal buff.
			sim.addAura(Aura{
				ID:      MagicIDDCCBonus,
				Expires: sim.CurrentTime + time.Second*10,
				OnExpire: func(sim *Simulation, c *Cast) {
					sim.Stats[StatSpellDmg] -= spellBonus * float64(stacks)
					stacks = 0
				},
			})
		},
	}
}

// AuraStatRemoval creates a general aura for removing any buff stat on expiring.
// This is useful for activations / effects that give temp stats.
func AuraStatRemoval(currentTime time.Duration, duration time.Duration, amount float64, stat Stat, id int32) Aura {
	return Aura{
		ID:      id,
		Expires: currentTime + duration,
		OnExpire: func(sim *Simulation, c *Cast) {
			if sim.Debug != nil {
				sim.Debug(" -%0.0f %s from %s\n", amount, stat.StatName(), AuraName(id))
			}
			sim.Stats[stat] -= amount
		},
	}
}

func TryActivateDrums(sim *Simulation) {
	if sim.Options.NumDrums == 0 || sim.isOnCD(MagicIDDrums) {
		return
	}

	sim.Stats[StatHaste] += 80
	sim.setCD(MagicIDDrums, time.Minute*2)
	sim.addAura(AuraStatRemoval(sim.CurrentTime, time.Second*30, 80, StatHaste, MagicIDDrums))
}

func TryActivateBloodlust(sim *Simulation) {
	if sim.Options.NumBloodlust <= sim.cache.bloodlustCasts || sim.isOnCD(MagicIDBloodlust) {
		return
	}

	dur := time.Second * 40 // assumes that multiple BLs are different shaman.
	sim.setCD(MagicIDBloodlust, dur)
	sim.cache.bloodlustCasts++ // TODO: will this break anything?
	sim.addAura(Aura{
		ID:      MagicIDBloodlust,
		Expires: sim.CurrentTime + dur,
		OnCast: func(sim *Simulation, c *Cast) {
			c.CastTime = (c.CastTime * 10) / 13 // 30% faster
		},
	})
}

func TryActivateRacial(sim *Simulation) {
	switch v := sim.Options.Buffs.Race; v {
	case RaceBonusOrc:
		if sim.isOnCD(MagicIDOrcBloodFury) {
			return
		}

		const spBonus = 143
		const dur = time.Second * 15
		const cd = time.Minute * 2

		sim.Stats[StatSpellDmg] += spBonus
		sim.setCD(MagicIDOrcBloodFury, cd)
		sim.addAura(AuraStatRemoval(sim.CurrentTime, dur, spBonus, StatSpellDmg, MagicIDOrcBloodFury))

	case RaceBonusTroll10, RaceBonusTroll30:
		if sim.isOnCD(MagicIDTrollBerserking) {
			return
		}

		hasteBonus := time.Duration(11) // 10% haste
		if v == RaceBonusTroll30 {
			hasteBonus = time.Duration(13) // 30% haste
		}
		const dur = time.Second * 10
		const cd = time.Minute * 3

		sim.setCD(MagicIDTrollBerserking, cd)
		sim.addAura(Aura{
			ID:      MagicIDTrollBerserking,
			Expires: sim.CurrentTime + dur,
			OnCast: func(sim *Simulation, c *Cast) {
				c.CastTime = (c.CastTime * 10) / hasteBonus
			},
		})
	}
}

func ActivateSkycall(sim *Simulation) Aura {
	const hasteBonus = 101
	const dur = time.Second * 10
	return Aura{
		ID:      MagicIDSkycall,
		Expires: neverExpires,
		OnCastComplete: func(sim *Simulation, c *Cast) {
			if c.Spell.ID == MagicIDLB12 && sim.Rando.Float64() < 0.15 {
				sim.Stats[StatHaste] += hasteBonus
				sim.addAura(Aura{
					ID:      MagicIDEnergized,
					Expires: sim.CurrentTime + dur,
					OnExpire: func(sim *Simulation, c *Cast) {
						sim.Stats[StatHaste] -= hasteBonus
					},
				})
			}
		},
	}
}

func ActivateNAC(sim *Simulation) Aura {
	const sp = 250
	sim.Stats[StatSpellDmg] += sp
	return Aura{
		ID:      MagicIDNAC,
		Expires: sim.CurrentTime + time.Second*20,
		OnCast: func(sim *Simulation, c *Cast) {
			c.ManaCost *= 1.2
		},
		OnExpire: func(sim *Simulation, c *Cast) {
			sim.Stats[StatSpellDmg] -= sp
		},
	}
}

func ActivateCSD(sim *Simulation) Aura {
	return Aura{
		ID:      MagicIDChaoticSkyfire,
		Expires: neverExpires,
		OnCastComplete: func(sim *Simulation, c *Cast) {
			c.CritBonus *= 1.03
		},
	}
}

func ActivateIED(sim *Simulation) Aura {
	icd := NewICD()
	const dur = time.Second * 15
	return Aura{
		ID:      MagicIDInsightfulEarthstorm,
		Expires: neverExpires,
		OnCastComplete: func(sim *Simulation, c *Cast) {
			if !icd.isOnCD(sim) && sim.Rando.Float64() < 0.04 {
				icd = InternalCD(sim.CurrentTime + dur)
				if sim.Debug != nil {
					sim.Debug(" *Insightful Earthstorm Mana Restore - 300\n")
				}
				sim.CurrentMana += 300
			}
		},
	}
}

func ActivateMSD(sim *Simulation) Aura {
	const hasteBonus = 320.0
	const dur = time.Second * 4
	const icdDur = time.Second * 35
	icd := NewICD()
	return Aura{
		ID:      MagicIDMysticSkyfire,
		Expires: neverExpires,
		OnCastComplete: func(sim *Simulation, c *Cast) {
			if !icd.isOnCD(sim) && sim.Rando.Float64() < 0.15 {
				icd = InternalCD(sim.CurrentTime + icdDur)
				sim.Stats[StatHaste] += hasteBonus
				sim.addAura(AuraStatRemoval(sim.CurrentTime, dur, hasteBonus, StatHaste, MagicIDMysticFocus))
			}
		},
	}
}

func ActivateESD(sim *Simulation) Aura {
	sim.Stats[StatInt] += sim.Stats[StatInt] * 0.02
	return Aura{
		ID:      MagicIDEmberSkyfire,
		Expires: neverExpires,
	}
}

func ActivateSpellstrike(sim *Simulation) Aura {
	const spellBonus = 92.0
	const duration = time.Second * 10
	return Aura{
		ID:      MagicIDSpellstrike,
		Expires: neverExpires,
		OnCastComplete: func(sim *Simulation, c *Cast) {
			if sim.Rando.Float64() < 0.05 {
				sim.Stats[StatSpellDmg] += spellBonus
				sim.addAura(Aura{
					ID:      MagicIDSpellstrikeInfusion,
					Expires: sim.CurrentTime + duration,
					OnExpire: func(sim *Simulation, c *Cast) {
						sim.Stats[StatSpellDmg] -= spellBonus
					},
				})
			}
		},
	}
}

func ActivateManaEtched(sim *Simulation) Aura {
	const spellBonus = 110.0
	const duration = time.Second * 15
	return Aura{
		ID:      MagicIDManaEtched,
		Expires: neverExpires,
		OnCastComplete: func(sim *Simulation, c *Cast) {
			if sim.Rando.Float64() < 0.02 {
				sim.Stats[StatSpellDmg] += spellBonus
				sim.addAura(Aura{
					ID:      MagicIDManaEtchedInsight,
					Expires: sim.CurrentTime + duration,
					OnExpire: func(sim *Simulation, c *Cast) {
						sim.Stats[StatSpellDmg] -= spellBonus
					},
				})
			}
		},
	}
}

func ActivateTLC(sim *Simulation) Aura {
	const spellBonus = 110.0
	const duration = time.Second * 15
	const icdDur = time.Millisecond * 2500
	tlcspell := spellmap[MagicIDTLCLB]

	charges := 0
	icd := NewICD()

	// Mods for TLC spells that don't change within the sim.
	var hitMod = (-0.02 * float64(sim.Options.Talents.ElementalPrecision)) + (-0.01 * float64(sim.Options.Talents.NaturesGuidance))
	var critMod = (-0.01 * float64(sim.Options.Talents.TidalMastery)) + (-0.01 * float64(sim.Options.Talents.CallOfThunder))

	return Aura{
		ID:      MagicIDTLC,
		Expires: neverExpires,
		OnSpellHit: func(sim *Simulation, c *Cast) {
			if icd.isOnCD(sim) {
				return
			}
			if !c.DidCrit {
				return
			}
			charges++
			if sim.Debug != nil {
				sim.Debug(" Lightning Capacitor Charges: %d\n", charges)
			}
			if charges >= 3 {
				if sim.Debug != nil {
					sim.Debug(" Lightning Capacitor Triggered!\n")
				}
				icd = InternalCD(sim.CurrentTime + icdDur)

				clone := sim.cache.NewCast()
				// TLC does not get hit talents bonus, subtract them here. (since we dont conditionally apply them)
				clone.Spell = tlcspell
				clone.CritBonus = 1.5
				clone.Hit = hitMod
				clone.Crit = critMod
				sim.Cast(clone)
				charges = 0
			}
		},
	}
}

func ActivateChainTO(sim *Simulation) Aura {
	if sim.Options.Buffs.TwilightOwl {
		return Aura{ID: MagicIDChainTO}
	}
	const bonus = 2 * 22.08 // 2% crit
	sim.Stats[StatSpellCrit] += bonus
	return Aura{
		ID:      MagicIDChainTO,
		Expires: sim.CurrentTime + time.Minute*30,
		// Technically this would never expire in any real sim... should I just make it neverExpires?
		OnExpire: func(sim *Simulation, c *Cast) {
			sim.Stats[StatSpellCrit] -= bonus
		},
	}
}

func ActivateCycloneManaReduce(sim *Simulation) Aura {
	return Aura{
		ID:      MagicIDCyclone4pc,
		Expires: neverExpires,
		OnSpellHit: func(sim *Simulation, c *Cast) {
			if c.DidCrit && sim.Rando.Float64() < 0.11 {
				sim.addAura(Aura{
					ID: MagicIDCycloneMana,
					OnCast: func(sim *Simulation, c *Cast) {
						// TODO: how to make sure this goes in before clearcasting?
						c.ManaCost -= 270
					},
					OnCastComplete: func(sim *Simulation, c *Cast) {
						sim.removeAura(MagicIDCycloneMana)
					},
				})
			}
		},
	}
}

func ActivateCataclysmLBDiscount(sim *Simulation) Aura {
	return Aura{
		ID:      MagicIDCataclysm4pc,
		Expires: neverExpires,
		OnSpellHit: func(sim *Simulation, c *Cast) {
			if c.DidCrit && sim.Rando.Float64() < 0.25 {
				sim.CurrentMana += 120
			}
		},
	}
}

func ActivateSkyshatterImpLB(sim *Simulation) Aura {
	return Aura{
		ID:      MagicIDSkyshatter4pc,
		Expires: neverExpires,
		OnSpellHit: func(sim *Simulation, c *Cast) {
			if c.Spell.ID == MagicIDLB12 {
				c.DidDmg *= 1.05
			}
		},
	}
}

func TryActivateDestructionPotion(sim *Simulation) {
	if !sim.Options.Consumes.DestructionPotion || sim.isOnCD(MagicIDPotion) {
		return
	}

	// Only use dest potion if not using mana or if we haven't used it once.
	// If we are using mana, only use destruction potion on the pull.
	if sim.cache.destructionPotion && sim.Options.Consumes.SuperManaPotion {
		return
	}

	const spBonus = 120
	const critBonus = 44.16
	const dur = time.Second * 15

	sim.cache.destructionPotion = true
	sim.setCD(MagicIDPotion, time.Second*120)
	sim.Stats[StatSpellDmg] += spBonus
	sim.Stats[StatSpellCrit] += critBonus

	sim.addAura(Aura{
		ID:      MagicIDDestructionPotion,
		Expires: sim.CurrentTime + dur,
		OnExpire: func(sim *Simulation, c *Cast) {
			sim.Stats[StatSpellDmg] -= spBonus
			sim.Stats[StatSpellCrit] -= critBonus
		},
	})
}

// TODO: This function doesn't really belong in auras.go, find a better home for it.
func TryActivateDarkRune(sim *Simulation) {
	if !sim.Options.Consumes.DarkRune || sim.isOnCD(MagicIDRune) {
		return
	}

	// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
	totalRegen := sim.manaRegenPerSecond() * 5
	if sim.Stats[StatMana]-(sim.CurrentMana+totalRegen) < 1500 {
		return
	}

	// Restores 900 to 1500 mana. (2 Min Cooldown)
	sim.CurrentMana += 900 + (sim.Rando.Float64() * 600)
	sim.setCD(MagicIDRune, time.Second*120)
	if sim.Debug != nil {
		sim.Debug("Used Dark Rune\n")
	}
	return
}

// TODO: This function doesn't really belong in auras.go, find a better home for it.
func TryActivateSuperManaPotion(sim *Simulation) {
	if !sim.Options.Consumes.SuperManaPotion || sim.isOnCD(MagicIDPotion) {
		return
	}

	// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
	totalRegen := sim.manaRegenPerSecond() * 5
	if sim.Stats[StatMana]-(sim.CurrentMana+totalRegen) < 3000 {
		return
	}

	// Restores 1800 to 3000 mana. (2 Min Cooldown)
	manaGain := 1800 + (sim.Rando.Float64() * 1200)

	if sim.hasAura(MagicIDAlchStone) {
		manaGain *= 1.4
	}

	sim.CurrentMana += manaGain
	sim.setCD(MagicIDPotion, time.Second*120)
	if sim.Debug != nil {
		sim.Debug("Used Mana Potion\n")
	}
	return
}

func ActivateSextant(sim *Simulation) Aura {
	icd := NewICD()
	const spellBonus = 190.0
	const dur = time.Second * 15
	const icdDur = time.Second * 45
	return Aura{
		ID:      MagicIDSextant,
		Expires: neverExpires,
		OnSpellHit: func(sim *Simulation, c *Cast) {
			if c.DidCrit && !icd.isOnCD(sim) && sim.Rando.Float64() < 0.2 {
				icd = InternalCD(sim.CurrentTime + icdDur)
				sim.Stats[StatSpellDmg] += spellBonus
				sim.addAura(AuraStatRemoval(sim.CurrentTime, dur, spellBonus, StatSpellDmg, MagicIDUnstableCurrents))
			}
		},
	}
}

func ActivateEyeOfMag(sim *Simulation) Aura {
	const spellBonus = 170.0
	const dur = time.Second * 10
	return Aura{
		ID:      MagicIDEyeOfMag,
		Expires: neverExpires,
		OnSpellMiss: func(sim *Simulation, c *Cast) {
			sim.Stats[StatSpellDmg] += spellBonus
			sim.addAura(Aura{
				ID:      MagicIDRecurringPower,
				Expires: sim.CurrentTime + dur,
				OnExpire: func(sim *Simulation, c *Cast) {
					sim.Stats[StatSpellDmg] -= spellBonus
				},
			})
		},
	}
}

func ActivateElderScribes(sim *Simulation) Aura {
	// Gives a chance when your harmful spells land to increase the damage of your spells and effects by up to 130 for 10 sec. (Proc chance: 20%, 50s cooldown)
	icd := NewICD()
	const spellBonus = 130.0
	const dur = time.Second * 10
	const icdDur = time.Second * 50
	const proc = 0.2
	return Aura{
		ID:      MagicIDElderScribe,
		Expires: neverExpires,
		OnSpellHit: func(sim *Simulation, c *Cast) {
			// This code is starting to look a lot like other ICD buff items. Perhaps we could DRY this out.
			if !icd.isOnCD(sim) && sim.Rando.Float64() < proc {
				icd = InternalCD(sim.CurrentTime + icdDur)
				sim.Stats[StatSpellDmg] += spellBonus
				sim.addAura(AuraStatRemoval(sim.CurrentTime, dur, spellBonus, StatSpellDmg, MagicIDElderScribeProc))
			}
		},
	}
}

func ActivateTotemOfPulsingEarth(sim *Simulation) Aura {
	return Aura{
		ID:      MagicIDTotemOfPulsingEarth,
		Expires: neverExpires,
		OnCast: func(sim *Simulation, c *Cast) {
			if c.Spell.ID == MagicIDLB12 {
				// TODO: how to make sure this goes in before clearcasting?
				c.ManaCost = math.Max(c.ManaCost-27, 0)
			}
		},
	}
}

// ActivateFathomBrooch adds an aura that has a chance on cast of nature spell
//  to restore 335 mana. 40s ICD
func ActivateFathomBrooch(sim *Simulation) Aura {
	icd := NewICD()
	const icdDur = time.Second * 40
	return Aura{
		ID:      MagicIDRegainMana,
		Expires: neverExpires,
		OnCastComplete: func(sim *Simulation, c *Cast) {
			if icd.isOnCD(sim) {
				return
			}
			if c.Spell.DamageType != DamageTypeNature {
				return
			}
			if sim.Rando.Float64() < 0.15 {
				icd = InternalCD(sim.CurrentTime + icdDur)
				sim.CurrentMana += 335
			}
		},
	}
}

// ActivateAlchStone adds the alch stone aura that has no effect on casts.
//  The usage for this aura is in the potion usage function.
func ActivateAlchStone(sim *Simulation) Aura {
	return Aura{
		ID:      MagicIDAlchStone,
		Expires: math.MaxInt32,
	}
}
