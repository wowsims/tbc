package core

import (
	"fmt"
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

const NeverExpires = time.Duration(math.MaxInt64)

type Aura struct {
	ID          int32
	Expires     time.Duration // time at which aura will be removed
	activeIndex int32         // Position of this aura's index in the sim.activeAuraIDs array

	// The number of stacks, or charges, of this aura. If this aura doesn't care
	// about charges, is just 0.
	Stacks int32

	// Invoked at creation time for a spell cast.
	OnCast         func(sim *Simulation, cast DirectCastAction, castInput *DirectCastInput)

	// Invoked when a spell cast completes casting, before results are calculated.
	OnCastComplete func(sim *Simulation, cast DirectCastAction)

	// Invoked before a spell lands, but after the target is selected.
	OnBeforeSpellHit func(sim *Simulation, hitInput *DirectCastDamageInput)

	// Invoked when a spell is fully resisted.
	OnSpellMiss    func(sim *Simulation, cast DirectCastAction)

	// Invoked when a spell hits, after results are calculated. Results can be modified by changing
	// properties of result.
	OnSpellHit     func(sim *Simulation, cast DirectCastAction, result *DirectCastDamageResult)

	// Invoked when this Aura expires.
	OnExpire       func(sim *Simulation)
}

// This needs to be a function that returns an Aura rather than an Aura, so captured
// local variables can be reset on Sim reset.
type PermanentAura func(*Simulation) Aura

// auraTracker is a centralized implementation of CD and Aura tracking.
//  This is currently used by Player and Raid (for global debuffs)
type auraTracker struct {
	// Auras that never expire and should always be active.
	// These are automatically applied on each Sim reset.
	permanentAuras []PermanentAura

	// Used for logging.
	playerID int

	// Whether finalize() has been called for this object.
	finalized bool

  // Maps MagicIDs to sim duration at which CD is done. Using array for perf.
	cooldowns [MagicIDLen]time.Duration

	// Maps MagicIDs to aura for that ID. Using array for perf.
	auras [MagicIDLen]Aura

	// IDs of Auras that are active, in no particular order
	activeAuraIDs []int32
}

func newAuraTracker() auraTracker {
	return auraTracker{
		permanentAuras: []PermanentAura{},
		activeAuraIDs: make([]int32, 0, 5),
	}
}

// Registers a permanent aura to this Character which will be re-applied on
// every Sim reset.
func (at *auraTracker) AddPermanentAura(permAura PermanentAura) {
	if at.finalized {
		panic("Permanent auras may not be added once finalized!")
	}

	at.permanentAuras = append(at.permanentAuras, permAura)
}

func (at *auraTracker) finalize() {
	if at.finalized {
		return
	}
	at.finalized = true
}

func (at *auraTracker) reset(sim *Simulation) {
	//if at.playerID == -1 {
	//	fmt.Printf("Resetting\n")
	//}
	at.auras = [MagicIDLen]Aura{}
	at.cooldowns = [MagicIDLen]time.Duration{}
	at.activeAuraIDs = at.activeAuraIDs[:0]

	for _, permAura := range at.permanentAuras {
		aura := permAura(sim)
		aura.Expires = NeverExpires
		//if at.playerID == -1 {
		//	fmt.Printf("Re-adding perm aura %s\n", AuraName(aura.ID))
		//}
		at.AddAura(sim, aura)
	}
	//if at.playerID == -1 {
	//	fmt.Printf("Done Resetting\n")
	//}
}

func (at *auraTracker) advance(sim *Simulation, newTime time.Duration) {
	// Go in reverse order so we can safely delete while looping
	for i := len(at.activeAuraIDs) - 1; i >= 0; i-- {
		id := at.activeAuraIDs[i]
		if at.auras[id].Expires != 0 && at.auras[id].Expires <= newTime {
			at.RemoveAura(sim, id)
		}
	}
}

// addAura will add a new aura to the simulation. If there is a matching aura ID
// it will be replaced with the newer aura.
// Auras with duration of 0 will be logged as activating but never added to simulation Auras.
func (at *auraTracker) AddAura(sim *Simulation, newAura Aura) {
	if newAura.Expires <= sim.CurrentTime {
		return // no need to waste time adding aura that doesn't last.
	}

	if at.HasAura(newAura.ID) {
		at.RemoveAura(sim, newAura.ID)
	}

	at.auras[newAura.ID] = newAura
	at.auras[newAura.ID].activeIndex = int32(len(at.activeAuraIDs))
	at.activeAuraIDs = append(at.activeAuraIDs, newAura.ID)

	if sim.Log != nil {
		sim.Log("(%d) +%s\n", at.playerID, AuraName(newAura.ID))
	}
}

// Remove an aura by its ID
func (at *auraTracker) RemoveAura(sim *Simulation, id int32) {
	if at.auras[id].OnExpire != nil {
		at.auras[id].OnExpire(sim)
	}
	removeActiveIndex := at.auras[id].activeIndex
	at.auras[id] = Aura{}

	// Overwrite the element being removed with the last element
	otherAuraID := at.activeAuraIDs[len(at.activeAuraIDs)-1]
	if id != otherAuraID {
		at.activeAuraIDs[removeActiveIndex] = otherAuraID
		at.auras[otherAuraID].activeIndex = removeActiveIndex
	}

	// Now we can remove the last element, in constant time
	at.activeAuraIDs = at.activeAuraIDs[:len(at.activeAuraIDs)-1]

	if sim.Log != nil {
		sim.Log("(%d) -%s\n", at.playerID, AuraName(id))
	}
}

// Returns whether an aura with the given ID is currently active.
func (at *auraTracker) HasAura(id int32) bool {
	//if at.playerID == -1 && len(at.activeAuraIDs) == 0 {
	//	fmt.Printf("Checking for aura %s, cur id: %d\n", AuraName(id), at.auras[id].ID)
	//}
	return at.auras[id].ID != 0
}

func (at *auraTracker) IsOnCD(magicID int32, currentTime time.Duration) bool {
	return at.cooldowns[magicID] > currentTime
}

func (at *auraTracker) GetRemainingCD(magicID int32, currentTime time.Duration) time.Duration {
	remainingCD := at.cooldowns[magicID] - currentTime
	if remainingCD > 0 {
		return remainingCD
	} else {
		return 0
	}
}

func (at *auraTracker) SetCD(magicID int32, newCD time.Duration) {
	at.cooldowns[magicID] = newCD
}

// Invokes the OnCast event for all tracked Auras.
func (at *auraTracker) OnCast(sim *Simulation, cast DirectCastAction, castInput *DirectCastInput) {
	for _, id := range at.activeAuraIDs {
		if at.auras[id].OnCast != nil {
			at.auras[id].OnCast(sim, cast, castInput)
		}
	}
}

// Invokes the OnCastComplete event for all tracked Auras.
func (at *auraTracker) OnCastComplete(sim *Simulation, cast DirectCastAction) {
	for _, id := range at.activeAuraIDs {
		if at.auras[id].OnCastComplete != nil {
			at.auras[id].OnCastComplete(sim, cast)
		}
	}
}

// Invokes the OnBeforeSpellHit event for all tracked Auras.
func (at *auraTracker) OnBeforeSpellHit(sim *Simulation, hitInput *DirectCastDamageInput) {
	for _, id := range at.activeAuraIDs {
		if at.auras[id].OnBeforeSpellHit != nil {
			at.auras[id].OnBeforeSpellHit(sim, hitInput)
		}
	}
}

// Invokes the OnSpellMiss event for all tracked Auras.
func (at *auraTracker) OnSpellMiss(sim *Simulation, cast DirectCastAction) {
	for _, id := range at.activeAuraIDs {
		if at.auras[id].OnSpellMiss != nil {
			at.auras[id].OnSpellMiss(sim, cast)
		}
	}
}

// Invokes the OnSpellHit event for all tracked Auras.
func (at *auraTracker) OnSpellHit(sim *Simulation, cast DirectCastAction, result *DirectCastDamageResult) {
	for _, id := range at.activeAuraIDs {
		if at.auras[id].OnSpellHit != nil {
			at.auras[id].OnSpellHit(sim, cast, result)
		}
	}
}

func AuraName(a int32) string {
	switch a {
	case MagicIDUnknown:
		return "Unknown"
	case MagicIDLOTalent:
		return "Lightning Overload Talent"
	case MagicIDJoW:
		return "Judgement Of Wisdom Aura"
	case MagicIDEleMastery:
		return "Elemental Mastery"
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
		return "Drums"
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
	case MagicIDManaEtchedInsight:
		return "Mana-EtchedInsight"
	case MagicIDOrcBloodFury:
		return "Orc Blood Fury"
	case MagicIDTrollBerserking:
		return "Troll Berserking"
	case MagicIDEyeOfTheNight:
		return "EyeOfTheNight"
	case MagicIDChainTO:
		return "Chain of the Twilight Owl"
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
	case MagicIDSkyshatter4pc:
		return "Skyshatter 4pc Set Bonus"
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
	case MagicIDImprovedSealOfTheCrusader:
		return "Improved Seal of the Crusader"
	}

	return "<<Add Aura name to switch!!>>"
}

// Stored value is the time at which the ICD will be off CD
type InternalCD time.Duration

func (icd InternalCD) IsOnCD(sim *Simulation) bool {
	return time.Duration(icd) > sim.CurrentTime
}

func NewICD() InternalCD {
	return InternalCD(0)
}

// List of all magic effects and spells and items and stuff that can go on CD or have an aura.
const (
	MagicIDUnknown int32 = iota

	// Basic CDs
	MagicIDGCD
	MagicIDMainHandSwing
	MagicIDOffHandSwing
	MagicIDRangedSwing

	// Spells, used for tracking CDs
	MagicIDChainLightning6
	// MagicIDFlameShock

	// Auras
	MagicIDLOTalent
	MagicIDJoW
	MagicIDEleMastery
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
	MagicIDSpellPower
	MagicIDRubySerpent
	MagicIDCallOfTheNexus
	MagicIDDCC
	MagicIDDCCBonus
	MagicIDDrums // drums effect
	MagicIDTwinStars
	MagicIDTidefury
	MagicIDSpellstrike
	MagicIDSpellstrikeInfusion
	MagicIDManaEtched
	MagicIDManaEtchedInsight
	MagicIDMisery
	MagicIDEyeOfTheNight
	MagicIDChainTO
	MagicIDCyclone4pc
	MagicIDCycloneMana // proc from 4pc
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
	MagicIDSkyshatter4pc    // skyshatter 4pc aura
	MagicIDElderScribe      // elder scribe robe item aura
	MagicIDElderScribeProc  // elder scribe robe temp buff
	MagicIDRegainMana // effect from fathom brooch
	MagicIDCurseOfElements
	MagicIDImprovedSealOfTheCrusader

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

// Helper for the common case of adding an Aura that gives a temporary stat boost.
func (character *Character) AddAuraWithTemporaryStats(sim *Simulation, auraID int32, stat stats.Stat, amount float64, duration time.Duration) {
	if sim.Log != nil {
		sim.Log(" +%0.0f %s from %s\n", amount, stat.StatName(), AuraName(auraID))
	}
	character.AddStat(stat, amount)

	character.AddAura(sim, Aura{
		ID:      auraID,
		Expires: sim.CurrentTime + duration,
		OnExpire: func(sim *Simulation) {
			if sim.Log != nil {
				sim.Log(" -%0.0f %s from %s\n", amount, stat.StatName(), AuraName(auraID))
			}
			character.AddStat(stat, -amount)
		},
	})
}
