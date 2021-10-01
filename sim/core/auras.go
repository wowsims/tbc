package core

import (
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

// AuraEffects will mutate a cast or simulation state.
type AuraEffect func(sim *Simulation, p PlayerAgent, c *Cast)

const NeverExpires = time.Duration(math.MaxInt64)

type Aura struct {
	ID          int32
	Expires     time.Duration // time at which aura will be removed
	activeIndex int32         // Position of this aura's index in the sim.activeAuraIDs array

	// The number of stacks, or charges, of this aura. If this aura doesn't care
	// about charges, is just 0.
	Stacks int32

	OnCast         AuraEffect
	OnCastComplete AuraEffect
	OnSpellHit     AuraEffect
	OnSpellMiss    AuraEffect
	OnExpire       AuraEffect
}

func NewAuraTracker() *AuraTracker {
	return &AuraTracker{
		ActiveAuraIDs: make([]int32, 0, 5),
	}
}

// AuraTracker is a centralized implementation of CD and Aura tracking.
//  This is currently used by Player and Raid (for global debuffs)
type AuraTracker struct {
	PID           int                       // used to track which player ID auras are coming on/off (mosly for debugging)
	CDs           [MagicIDLen]time.Duration // Map of MagicID to sim duration at which CD is done.
	Auras         [MagicIDLen]Aura          // this is array instead of map to speed up browser perf.
	ActiveAuraIDs []int32                   // IDs of auras that are active, in no particular order
}

func (at *AuraTracker) ResetAuras() {
	at.Auras = [MagicIDLen]Aura{}
	at.CDs = [MagicIDLen]time.Duration{}
	at.ActiveAuraIDs = at.ActiveAuraIDs[:0]
}

func (at *AuraTracker) Advance(sim *Simulation, player PlayerAgent, newTime time.Duration) {
	// Go in reverse order so we can safely delete while looping
	for i := len(at.ActiveAuraIDs) - 1; i >= 0; i-- {
		id := at.ActiveAuraIDs[i]
		if at.Auras[id].Expires != 0 && at.Auras[id].Expires <= newTime {
			at.RemoveAura(sim, player, id)
		}
	}
}

// addAura will add a new aura to the simulation. If there is a matching aura ID
// it will be replaced with the newer aura.
// Auras with duration of 0 will be logged as activating but never added to simulation auras.
func (at *AuraTracker) AddAura(sim *Simulation, player PlayerAgent, newAura Aura) {
	if newAura.Expires < sim.CurrentTime {
		return // no need to waste time adding aura that doesn't last.
	}

	if at.HasAura(newAura.ID) {
		at.RemoveAura(sim, player, newAura.ID)
	}

	at.Auras[newAura.ID] = newAura
	at.Auras[newAura.ID].activeIndex = int32(len(at.ActiveAuraIDs))
	at.ActiveAuraIDs = append(at.ActiveAuraIDs, newAura.ID)

	if sim.Log != nil {
		sim.Log("(%d) +%s\n", at.PID, AuraName(newAura.ID))
	}
}

// Remove an aura by its ID
func (at *AuraTracker) RemoveAura(sim *Simulation, player PlayerAgent, id int32) {
	if at.Auras[id].OnExpire != nil {
		at.Auras[id].OnExpire(sim, player, nil)
	}
	removeActiveIndex := at.Auras[id].activeIndex
	at.Auras[id] = Aura{}

	// Overwrite the element being removed with the last element
	otherAuraID := at.ActiveAuraIDs[len(at.ActiveAuraIDs)-1]
	if id != otherAuraID {
		at.ActiveAuraIDs[removeActiveIndex] = otherAuraID
		at.Auras[otherAuraID].activeIndex = removeActiveIndex
	}

	// Now we can remove the last element, in constant time
	at.ActiveAuraIDs = at.ActiveAuraIDs[:len(at.ActiveAuraIDs)-1]

	if sim.Log != nil {
		sim.Log("(%d) -%s\n", at.PID, AuraName(id))
	}
}

// Returns whether an aura with the given ID is currently active.
func (at *AuraTracker) HasAura(id int32) bool {
	return at.Auras[id].ID != 0
}

func (at *AuraTracker) IsOnCD(magicID int32, currentTime time.Duration) bool {
	return at.CDs[magicID] > currentTime
}

func (at *AuraTracker) GetRemainingCD(magicID int32, currentTime time.Duration) time.Duration {
	remainingCD := at.CDs[magicID] - currentTime
	if remainingCD > 0 {
		return remainingCD
	} else {
		return 0
	}
}

func (at *AuraTracker) SetCD(magicID int32, newCD time.Duration) {
	at.CDs[magicID] = newCD
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
	case MagicIDAtkTrinket:
		return "Shared Attack Trinkets CD"
	case MagicIDHealTrinket:
		return "Shared Heal Trinkets CD"
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

	return "<<Add Aura name to switch!!>>"
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
	MagicIDCurseOfElements

	// Items  (Usually individual trinket CDs)
	MagicIDISCTrink
	MagicIDNACTrink
	MagicIDPotion
	MagicIDRune
	MagicIDAtkTrinket
	MagicIDHealTrinket
	MagicIDScryerTrink
	MagicIDRubySerpentTrink
	MagicIDXiriTrink
	MagicIDHexTrink
	MagicIDShiftingNaaruTrink
	MagicIDSkullGuldanTrink
	MagicIDEssMartyrTrink
	MagicIDEssSappTrink
	MagicIDDITrink // Dark Iron pipe trinket CD

	// Always at end so we know how many magic IDs there are.
	MagicIDLen
)

func createHasteActivate(id int32, haste float64, duration time.Duration) ItemActivation {
	// Implemented haste activate as a buff so that the creation of a new cast gets the correct cast time
	return func(sim *Simulation, party *Party, player PlayerAgent) Aura {
		player.Stats[stats.SpellHaste] += haste
		return Aura{
			ID:      id,
			Expires: sim.CurrentTime + duration,
			OnExpire: func(sim *Simulation, player PlayerAgent, c *Cast) {
				player.Stats[stats.SpellHaste] -= haste
			},
		}
	}
}

// createSpellDmgActivate creates a function for trinket activation that uses +spellpower
//  This is so we don't need a separate function for every spell power trinket.
func createSpellDmgActivate(id int32, sp float64, duration time.Duration) ItemActivation {
	return func(sim *Simulation, party *Party, player PlayerAgent) Aura {
		player.Stats[stats.SpellPower] += sp
		return Aura{
			ID:      id,
			Expires: sim.CurrentTime + duration,
			OnExpire: func(sim *Simulation, player PlayerAgent, c *Cast) {
				player.Stats[stats.SpellPower] -= sp
			},
		}
	}
}

func ActivateQuagsEye(sim *Simulation, party *Party, player PlayerAgent) Aura {
	const hasteBonus = 320.0
	const dur = time.Second * 45
	icd := NewICD()
	return Aura{
		ID:      MagicIDQuagsEye,
		Expires: NeverExpires,
		OnCastComplete: func(sim *Simulation, player PlayerAgent, c *Cast) {
			if !icd.isOnCD(sim) && sim.Rando.Float64("quags") < 0.1 {
				icd = InternalCD(sim.CurrentTime + dur)
				player.Stats[stats.SpellHaste] += hasteBonus
				player.AddAura(sim, AuraStatRemoval(sim.CurrentTime, time.Second*6, hasteBonus, stats.SpellHaste, MagicIDFungalFrenzy))
			}
		},
	}
}

func ActivateNexusHorn(sim *Simulation, party *Party, player PlayerAgent) Aura {
	icd := NewICD()
	const spellBonus = 225.0
	const dur = time.Second * 45
	return Aura{
		ID:      MagicIDNexusHorn,
		Expires: NeverExpires,
		OnSpellHit: func(sim *Simulation, player PlayerAgent, c *Cast) {
			if c.Spell.ID == MagicIDTLCLB {
				return // TLC can't proc Sextant
			}
			if !icd.isOnCD(sim) && c.DidCrit && sim.Rando.Float64("unmarked") < 0.2 {
				icd = InternalCD(sim.CurrentTime + dur)
				player.Stats[stats.SpellPower] += spellBonus
				player.AddAura(sim, AuraStatRemoval(sim.CurrentTime, time.Second*10, spellBonus, stats.SpellPower, MagicIDCallOfTheNexus))
			}
		},
	}
}

func ActivateDCC(sim *Simulation, party *Party, player PlayerAgent) Aura {
	const spellBonus = 8.0
	stacks := 0
	return Aura{
		ID:      MagicIDDCC,
		Expires: NeverExpires,
		OnCastComplete: func(sim *Simulation, player PlayerAgent, c *Cast) {
			if stacks < 10 {
				stacks++
				player.Stats[stats.SpellPower] += spellBonus
			}
			// Removal aura will refresh with new total spellpower based on stacks.
			//  This will remove the old stack removal buff.
			player.AddAura(sim, Aura{
				ID:      MagicIDDCCBonus,
				Expires: sim.CurrentTime + time.Second*10,
				OnExpire: func(sim *Simulation, player PlayerAgent, c *Cast) {
					player.Stats[stats.SpellPower] -= spellBonus * float64(stacks)
					stacks = 0
				},
			})
		},
	}
}

// AuraStatRemoval creates a general aura for removing any buff stat on expiring.
// This is useful for activations / effects that give temp stats.
func AuraStatRemoval(currentTime time.Duration, duration time.Duration, amount float64, stat stats.Stat, id int32) Aura {
	return Aura{
		ID:      id,
		Expires: currentTime + duration,
		OnExpire: func(sim *Simulation, player PlayerAgent, c *Cast) {
			if sim.Log != nil {
				sim.Log(" -%0.0f %s from %s\n", amount, stat.StatName(), AuraName(id))
			}
			player.Stats[stat] -= amount
		},
	}
}

func ActivateCSD(sim *Simulation, party *Party, player PlayerAgent) Aura {
	return Aura{
		ID:      MagicIDChaoticSkyfire,
		Expires: NeverExpires,
		OnCast: func(sim *Simulation, p PlayerAgent, c *Cast) {
			// TODO: Figure out how to make this work properly/easily with crit bonus
			//  and classes with crit bonus modifiers
			c.CritBonus *= 1.03
		},
	}
}

func ActivateIED(sim *Simulation, party *Party, player PlayerAgent) Aura {
	icd := NewICD()
	const dur = time.Second * 15
	return Aura{
		ID:      MagicIDInsightfulEarthstorm,
		Expires: NeverExpires,
		OnCastComplete: func(sim *Simulation, player PlayerAgent, c *Cast) {
			if !icd.isOnCD(sim) && sim.Rando.Float64("unmarked") < 0.04 {
				icd = InternalCD(sim.CurrentTime + dur)
				if sim.Log != nil {
					sim.Log(" *Insightful Earthstorm Mana Restore - 300\n")
				}
				player.Stats[stats.Mana] += 300
			}
		},
	}
}

func ActivateMSD(sim *Simulation, party *Party, player PlayerAgent) Aura {
	const hasteBonus = 320.0
	const dur = time.Second * 4
	const icdDur = time.Second * 35
	icd := NewICD()
	return Aura{
		ID:      MagicIDMysticSkyfire,
		Expires: NeverExpires,
		OnCastComplete: func(sim *Simulation, player PlayerAgent, c *Cast) {
			if !icd.isOnCD(sim) && sim.Rando.Float64("unmarked") < 0.15 {
				icd = InternalCD(sim.CurrentTime + icdDur)
				player.Stats[stats.SpellHaste] += hasteBonus
				player.AddAura(sim, AuraStatRemoval(sim.CurrentTime, dur, hasteBonus, stats.SpellHaste, MagicIDMysticFocus))
			}
		},
	}
}

func ActivateESD(sim *Simulation, party *Party, player PlayerAgent) Aura {
	// FUTURE: this technically should be modified by blessing of kings?
	player.Stats[stats.Intellect] += player.Stats[stats.Intellect] * 0.02
	return Aura{
		ID:      MagicIDEmberSkyfire,
		Expires: NeverExpires,
	}
}

func ActivateSpellstrike(sim *Simulation, party *Party, player PlayerAgent) Aura {
	const spellBonus = 92.0
	const duration = time.Second * 10
	return Aura{
		ID:      MagicIDSpellstrike,
		Expires: NeverExpires,
		OnCastComplete: func(sim *Simulation, player PlayerAgent, c *Cast) {
			if sim.Rando.Float64("unmarked") < 0.05 {
				player.Stats[stats.SpellPower] += spellBonus
				player.AddAura(sim, Aura{
					ID:      MagicIDSpellstrikeInfusion,
					Expires: sim.CurrentTime + duration,
					OnExpire: func(sim *Simulation, player PlayerAgent, c *Cast) {
						player.Stats[stats.SpellPower] -= spellBonus
					},
				})
			}
		},
	}
}

func ActivateManaEtched(sim *Simulation, party *Party, player PlayerAgent) Aura {
	const spellBonus = 110.0
	const duration = time.Second * 15
	return Aura{
		ID:      MagicIDManaEtched,
		Expires: NeverExpires,
		OnCastComplete: func(sim *Simulation, player PlayerAgent, c *Cast) {
			if sim.Rando.Float64("unmarked") < 0.02 {
				player.Stats[stats.SpellPower] += spellBonus
				player.AddAura(sim, Aura{
					ID:      MagicIDManaEtchedInsight,
					Expires: sim.CurrentTime + duration,
					OnExpire: func(sim *Simulation, player PlayerAgent, c *Cast) {
						player.Stats[stats.SpellPower] -= spellBonus
					},
				})
			}
		},
	}
}

func ActivateTLC(sim *Simulation, party *Party, player PlayerAgent) Aura {
	const icdDur = time.Millisecond * 2500
	tlcspell := Spells[MagicIDTLCLB]

	charges := 0
	icd := NewICD()

	return Aura{
		ID:      MagicIDTLC,
		Expires: NeverExpires,
		OnSpellHit: func(sim *Simulation, player PlayerAgent, c *Cast) {
			if icd.isOnCD(sim) {
				return
			}
			if !c.DidCrit {
				return
			}
			charges++
			if sim.Log != nil {
				sim.Log(" Lightning Capacitor Charges: %d\n", charges)
			}
			if charges >= 3 {
				if sim.Log != nil {
					sim.Log(" Lightning Capacitor Triggered!\n")
				}
				icd = InternalCD(sim.CurrentTime + icdDur)
				clone := NewCast(sim, tlcspell)
				// TODO: handle CSD for TLC better...
				if player.HasAura(MagicIDChaoticSkyfire) {
					player.Auras[MagicIDChaoticSkyfire].OnCast(sim, player, clone)
				}
				DirectCast(sim, player, clone)
				charges = 0
			}
		},
	}
}

// Because this buff actually lasts 30min, lets just assume we never lose the buff.
func ActivateChainTO(sim *Simulation, party *Party, player PlayerAgent) Aura {
	const bonus = 2 * 22.08 // 2% crit
	party.AddStats(stats.Stats{stats.SpellCrit: bonus})
	return Aura{
		ID: MagicIDChainTO,
	}
}

// Because this buff actually lasts 30min, lets just assume we never lose the buff.
func ActivateEyeOfNight(sim *Simulation, party *Party, player PlayerAgent) Aura {
	const bonus = 34
	party.AddStats(stats.Stats{stats.SpellPower: bonus})
	return Aura{
		ID: MagicIDEyeOfTheNight,
	}
}

func ActivateCycloneManaReduce(sim *Simulation, party *Party, player PlayerAgent) Aura {
	return Aura{
		ID:      MagicIDCyclone4pc,
		Expires: NeverExpires,
		OnSpellHit: func(sim *Simulation, player PlayerAgent, c *Cast) {
			if c.DidCrit && sim.Rando.Float64("unmarked") < 0.11 {
				player.AddAura(sim, Aura{
					ID: MagicIDCycloneMana,
					OnCast: func(sim *Simulation, p PlayerAgent, c *Cast) {
						// TODO: how to make sure this goes in before clearcasting?
						c.ManaCost -= 270
					},
					OnCastComplete: func(sim *Simulation, player PlayerAgent, c *Cast) {
						player.RemoveAura(sim, player, MagicIDCycloneMana)
					},
				})
			}
		},
	}
}

func ActivateCataclysmLBDiscount(sim *Simulation, party *Party, player PlayerAgent) Aura {
	return Aura{
		ID:      MagicIDCataclysm4pc,
		Expires: NeverExpires,
		OnSpellHit: func(sim *Simulation, player PlayerAgent, c *Cast) {
			if c.DidCrit && sim.Rando.Float64("unmarked") < 0.25 {
				player.Stats[stats.Mana] += 120
			}
		},
	}
}

func ActivateSkyshatterImpLB(sim *Simulation, party *Party, player PlayerAgent) Aura {
	return Aura{
		ID:      MagicIDSkyshatter4pc,
		Expires: NeverExpires,
		OnSpellHit: func(sim *Simulation, player PlayerAgent, c *Cast) {
			if c.Spell.ID == MagicIDLB12 {
				c.DidDmg *= 1.05
			}
		},
	}
}

func ActivateSextant(sim *Simulation, party *Party, player PlayerAgent) Aura {
	icd := NewICD()
	const spellBonus = 190.0
	const dur = time.Second * 15
	const icdDur = time.Second * 45
	return Aura{
		ID:      MagicIDSextant,
		Expires: NeverExpires,
		OnSpellHit: func(sim *Simulation, player PlayerAgent, c *Cast) {
			if c.Spell.ID == MagicIDTLCLB {
				return // TLC can't proc Sextant
			}
			if c.DidCrit && !icd.isOnCD(sim) && sim.Rando.Float64("unmarked") < 0.2 {
				icd = InternalCD(sim.CurrentTime + icdDur)
				player.Stats[stats.SpellPower] += spellBonus
				player.AddAura(sim, AuraStatRemoval(sim.CurrentTime, dur, spellBonus, stats.SpellPower, MagicIDUnstableCurrents))
			}
		},
	}
}

func ActivateEyeOfMag(sim *Simulation, party *Party, player PlayerAgent) Aura {
	const spellBonus = 170.0
	const dur = time.Second * 10
	return Aura{
		ID:      MagicIDEyeOfMag,
		Expires: NeverExpires,
		OnSpellMiss: func(sim *Simulation, p PlayerAgent, c *Cast) {
			player.Stats[stats.SpellPower] += spellBonus
			player.AddAura(sim, Aura{
				ID:      MagicIDRecurringPower,
				Expires: sim.CurrentTime + dur,
				OnExpire: func(sim *Simulation, player PlayerAgent, c *Cast) {
					player.Stats[stats.SpellPower] -= spellBonus
				},
			})
		},
	}
}

func ActivateElderScribes(sim *Simulation, party *Party, player PlayerAgent) Aura {
	// Gives a chance when your harmful spells land to increase the damage of your spells and effects by up to 130 for 10 sec. (Proc chance: 20%, 50s cooldown)
	icd := NewICD()
	const spellBonus = 130.0
	const dur = time.Second * 10
	const icdDur = time.Second * 50
	const proc = 0.2
	return Aura{
		ID:      MagicIDElderScribe,
		Expires: NeverExpires,
		OnSpellHit: func(sim *Simulation, player PlayerAgent, c *Cast) {
			// This code is starting to look a lot like other ICD buff items. Perhaps we could DRY this out.
			if !icd.isOnCD(sim) && sim.Rando.Float64("unmarked") < proc {
				icd = InternalCD(sim.CurrentTime + icdDur)
				player.Stats[stats.SpellPower] += spellBonus
				player.AddAura(sim, AuraStatRemoval(sim.CurrentTime, dur, spellBonus, stats.SpellPower, MagicIDElderScribeProc))
			}
		},
	}
}

// ActivateFathomBrooch adds an aura that has a chance on cast of nature spell
//  to restore 335 mana. 40s ICD
func ActivateFathomBrooch(sim *Simulation, party *Party, player PlayerAgent) Aura {
	icd := NewICD()
	const icdDur = time.Second * 40
	return Aura{
		ID:      MagicIDRegainMana,
		Expires: NeverExpires,
		OnCastComplete: func(sim *Simulation, player PlayerAgent, c *Cast) {
			if icd.isOnCD(sim) {
				return
			}
			if c.Spell.DamageType != stats.NatureSpellPower {
				return
			}
			if sim.Rando.Float64("unmarked") < 0.15 {
				icd = InternalCD(sim.CurrentTime + icdDur)
				player.Stats[stats.Mana] += 335
			}
		},
	}
}

// ActivateAlchStone adds the alch stone aura that has no effect on casts.
//  The usage for this aura is in the potion usage function.
func ActivateAlchStone(sim *Simulation, party *Party, player PlayerAgent) Aura {
	return Aura{
		ID:      MagicIDAlchStone,
		Expires: math.MaxInt32,
	}
}
