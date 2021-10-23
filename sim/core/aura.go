package core

import (
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

	// Invoked when a spell is fully resisted.
	OnSpellMiss    func(sim *Simulation, cast DirectCastAction)

	// Invoked when a spell hits, after results are calculated. Results can be modified by changing
	// properties of result.
	OnSpellHit     func(sim *Simulation, cast DirectCastAction, result *DirectCastDamageResult)

	// Invoked when this Aura expires.
	OnExpire       func(sim *Simulation)
}

func NewAuraTracker() *AuraTracker {
	return &AuraTracker{
		ActiveAuraIDs: make([]int32, 0, 5),
	}
}

// AuraTracker is a centralized implementation of CD and Aura tracking.
//  This is currently used by Player and Raid (for global debuffs)
type AuraTracker struct {
	PID           int                       // used to track which agent ID auras are coming on/off (mosly for debugging)
	CDs           [MagicIDLen]time.Duration // Maps MagicIDs to sim duration at which CD is done. Using array for perf.
	Auras         [MagicIDLen]Aura          // Maps MagicIDs to aura for that ID. Using array for perf.
	ActiveAuraIDs []int32                   // IDs of auras that are active, in no particular order
}

func (at *AuraTracker) ResetAuras() {
	at.Auras = [MagicIDLen]Aura{}
	at.CDs = [MagicIDLen]time.Duration{}
	at.ActiveAuraIDs = at.ActiveAuraIDs[:0]
}

func (at *AuraTracker) Advance(sim *Simulation, newTime time.Duration) {
	// Go in reverse order so we can safely delete while looping
	for i := len(at.ActiveAuraIDs) - 1; i >= 0; i-- {
		id := at.ActiveAuraIDs[i]
		if at.Auras[id].Expires != 0 && at.Auras[id].Expires <= newTime {
			at.RemoveAura(sim, id)
		}
	}
}

// addAura will add a new aura to the simulation. If there is a matching aura ID
// it will be replaced with the newer aura.
// Auras with duration of 0 will be logged as activating but never added to simulation auras.
func (at *AuraTracker) AddAura(sim *Simulation, newAura Aura) {
	if newAura.Expires <= sim.CurrentTime {
		return // no need to waste time adding aura that doesn't last.
	}

	if at.HasAura(newAura.ID) {
		at.RemoveAura(sim, newAura.ID)
	}

	at.Auras[newAura.ID] = newAura
	at.Auras[newAura.ID].activeIndex = int32(len(at.ActiveAuraIDs))
	at.ActiveAuraIDs = append(at.ActiveAuraIDs, newAura.ID)

	if sim.Log != nil {
		sim.Log("(%d) +%s\n", at.PID, AuraName(newAura.ID))
	}
}

// Remove an aura by its ID
func (at *AuraTracker) RemoveAura(sim *Simulation, id int32) {
	if at.Auras[id].OnExpire != nil {
		at.Auras[id].OnExpire(sim)
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

// Add stats to a character, and return an aura that removes them when expired.
// NOTE: In most cases, use AddAuraWithTemporaryStats() instead.
func AddTemporaryStats(sim *Simulation, character *Character, auraID int32, stat stats.Stat, amount float64, duration time.Duration) Aura {
	if sim.Log != nil {
		sim.Log(" +%0.0f %s from %s\n", amount, stat.StatName(), AuraName(auraID))
	}
	character.AddStat(stat, amount)

	return Aura{
		ID:      auraID,
		Expires: sim.CurrentTime + duration,
		OnExpire: func(sim *Simulation) {
			if sim.Log != nil {
				sim.Log(" -%0.0f %s from %s\n", amount, stat.StatName(), AuraName(auraID))
			}
			character.AddStat(stat, -amount)
		},
	}
}

// Like AddTemporaryStats(), but also adds the aura to the character.
func AddAuraWithTemporaryStats(sim *Simulation, character *Character, auraID int32, stat stats.Stat, amount float64, duration time.Duration) {
	character.AddAura(sim, AddTemporaryStats(sim, character, auraID, stat, amount, duration))
}
