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
	CDs           [MagicIDLen]time.Duration // Map of MagicID to sim duration at which CD is done.
	Auras         [MagicIDLen]Aura          // this is array instead of map to speed up browser perf.
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
	case MagicIDEmberSkyfire:
		return "Ember Skyfire"
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

// Add stats to an agent, and return an aura that removes them when expired.
// NOTE: In most cases, use AddAuraWithTemporaryStats() instead.
func AddTemporaryStats(sim *Simulation, agent Agent, auraID int32, stat stats.Stat, amount float64, duration time.Duration) Aura {
	if sim.Log != nil {
		sim.Log(" +%0.0f %s from %s\n", amount, stat.StatName(), AuraName(auraID))
	}
	agent.GetCharacter().Stats[stat] += amount

	return Aura{
		ID:      auraID,
		Expires: sim.CurrentTime + duration,
		OnExpire: func(sim *Simulation) {
			if sim.Log != nil {
				sim.Log(" -%0.0f %s from %s\n", amount, stat.StatName(), AuraName(auraID))
			}
			agent.GetCharacter().Stats[stat] -= amount
		},
	}
}

// Like AddTemporaryStats(), but also adds the aura to the agent.
func AddAuraWithTemporaryStats(sim *Simulation, agent Agent, auraID int32, stat stats.Stat, amount float64, duration time.Duration) {
	agent.GetCharacter().AddAura(sim, AddTemporaryStats(sim, agent, auraID, stat, amount, duration))
}

func createHasteActivate(id int32, haste float64, duration time.Duration) ItemActivation {
	// Implemented haste activate as a buff so that the creation of a new cast gets the correct cast time
	return func(sim *Simulation, agent Agent) Aura {
		return AddTemporaryStats(sim, agent, id, stats.SpellHaste, haste, duration)
	}
}

// createSpellDmgActivate creates a function for trinket activation that uses +spellpower
//  This is so we don't need a separate function for every spell power trinket.
func createSpellDmgActivate(id int32, sp float64, duration time.Duration) ItemActivation {
	return func(sim *Simulation, agent Agent) Aura {
		return AddTemporaryStats(sim, agent, id, stats.SpellPower, sp, duration)
	}
}

func ActivateQuagsEye(sim *Simulation, agent Agent) Aura {
	const hasteBonus = 320.0
	const dur = time.Second * 45
	icd := NewICD()
	return Aura{
		ID:      MagicIDQuagsEye,
		Expires: NeverExpires,
		OnCastComplete: func(sim *Simulation, cast DirectCastAction) {
			if !icd.isOnCD(sim) && sim.Rando.Float64("quags") < 0.1 {
				icd = InternalCD(sim.CurrentTime + dur)
				AddAuraWithTemporaryStats(sim, agent, MagicIDFungalFrenzy, stats.SpellHaste, hasteBonus, time.Second*6)
			}
		},
	}
}

func ActivateNexusHorn(sim *Simulation, agent Agent) Aura {
	icd := NewICD()
	const spellBonus = 225.0
	const dur = time.Second * 45
	return Aura{
		ID:      MagicIDNexusHorn,
		Expires: NeverExpires,
		OnSpellHit: func(sim *Simulation, cast DirectCastAction, result *DirectCastDamageResult) {
			if cast.GetActionID().ItemID == ItemIDTLC {
				return // TLC can't proc Sextant
			}
			if !icd.isOnCD(sim) && result.Crit && sim.Rando.Float64("unmarked") < 0.2 {
				icd = InternalCD(sim.CurrentTime + dur)
				AddAuraWithTemporaryStats(sim, agent, MagicIDCallOfTheNexus, stats.SpellPower, spellBonus, time.Second*10)
			}
		},
	}
}

func ActivateDCC(sim *Simulation, agent Agent) Aura {
	const spellBonus = 8.0
	stacks := 0
	return Aura{
		ID:      MagicIDDCC,
		Expires: NeverExpires,
		OnCastComplete: func(sim *Simulation, cast DirectCastAction) {
			if stacks < 10 {
				stacks++
				agent.GetCharacter().Stats[stats.SpellPower] += spellBonus
			}
			// Removal aura will refresh with new total spellpower based on stacks.
			//  This will remove the old stack removal buff.
			agent.GetCharacter().AddAura(sim, Aura{
				ID:      MagicIDDCCBonus,
				Expires: sim.CurrentTime + time.Second*10,
				OnExpire: func(sim *Simulation) {
					agent.GetCharacter().Stats[stats.SpellPower] -= spellBonus * float64(stacks)
					stacks = 0
				},
			})
		},
	}
}

func ActivateCSD(sim *Simulation, agent Agent) Aura {
	return Aura{
		ID:      MagicIDChaoticSkyfire,
		Expires: NeverExpires,
		OnCast: func(sim *Simulation, cast DirectCastAction, input *DirectCastInput) {
			// For a normal spell with crit multiplier of 1.5, this will be 1.
			// For a spell with a multiplier of 2 (i.e. 100% increased critical damage) this will be 2.
			improvedCritRatio := (input.CritMultiplier - 1) / 0.5

			input.CritMultiplier += 0.045 * improvedCritRatio
		},
	}
}

func ActivateIED(sim *Simulation, agent Agent) Aura {
	icd := NewICD()
	const dur = time.Second * 15
	return Aura{
		ID:      MagicIDInsightfulEarthstorm,
		Expires: NeverExpires,
		OnCastComplete: func(sim *Simulation, cast DirectCastAction) {
			if !icd.isOnCD(sim) && sim.Rando.Float64("unmarked") < 0.04 {
				icd = InternalCD(sim.CurrentTime + dur)
				if sim.Log != nil {
					sim.Log(" *Insightful Earthstorm Mana Restore - 300\n")
				}
				agent.GetCharacter().Stats[stats.Mana] += 300
			}
		},
	}
}

func ActivateMSD(sim *Simulation, agent Agent) Aura {
	const hasteBonus = 320.0
	const dur = time.Second * 4
	const icdDur = time.Second * 35
	icd := NewICD()
	return Aura{
		ID:      MagicIDMysticSkyfire,
		Expires: NeverExpires,
		OnCastComplete: func(sim *Simulation, cast DirectCastAction) {
			if !icd.isOnCD(sim) && sim.Rando.Float64("unmarked") < 0.15 {
				icd = InternalCD(sim.CurrentTime + icdDur)
				AddAuraWithTemporaryStats(sim, agent, MagicIDMysticFocus, stats.SpellHaste, hasteBonus, dur)
			}
		},
	}
}

func ActivateESD(sim *Simulation, agent Agent) Aura {
	// FUTURE: this technically should be modified by blessing of kings?
	agent.GetCharacter().Stats[stats.Intellect] += agent.GetCharacter().Stats[stats.Intellect] * 0.02
	return Aura{
		ID:      MagicIDEmberSkyfire,
		Expires: NeverExpires,
	}
}

func ActivateSpellstrike(sim *Simulation, agent Agent) Aura {
	const spellBonus = 92.0
	const duration = time.Second * 10
	return Aura{
		ID:      MagicIDSpellstrike,
		Expires: NeverExpires,
		OnCastComplete: func(sim *Simulation, cast DirectCastAction) {
			if sim.Rando.Float64("unmarked") < 0.05 {
				AddAuraWithTemporaryStats(sim, agent, MagicIDSpellstrikeInfusion, stats.SpellPower, spellBonus, duration)
			}
		},
	}
}

func ActivateManaEtched(sim *Simulation, agent Agent) Aura {
	const spellBonus = 110.0
	const duration = time.Second * 15
	return Aura{
		ID:      MagicIDManaEtched,
		Expires: NeverExpires,
		OnCastComplete: func(sim *Simulation, cast DirectCastAction) {
			if sim.Rando.Float64("unmarked") < 0.02 {
				AddAuraWithTemporaryStats(sim, agent, MagicIDManaEtchedInsight, stats.SpellPower, spellBonus, duration)
			}
		},
	}
}

func ActivateTLC(sim *Simulation, agent Agent) Aura {
	charges := 0

	const icdDur = time.Millisecond * 2500
	icd := NewICD()

	return Aura{
		ID:      MagicIDTLC,
		Expires: NeverExpires,
		OnSpellHit: func(sim *Simulation, cast DirectCastAction, result *DirectCastDamageResult) {
			if icd.isOnCD(sim) {
				return
			}

			if !result.Crit {
				return
			}

			charges++
			if charges >= 3 {
				icd = InternalCD(sim.CurrentTime + icdDur)
				charges = 0
				castAction := NewLightningCapacitorCast(sim, agent)
				castAction.Act(sim)
			}
		},
	}
}

func ActivateSextant(sim *Simulation, agent Agent) Aura {
	icd := NewICD()
	const spellBonus = 190.0
	const dur = time.Second * 15
	const icdDur = time.Second * 45
	return Aura{
		ID:      MagicIDSextant,
		Expires: NeverExpires,
		OnSpellHit: func(sim *Simulation, cast DirectCastAction, result *DirectCastDamageResult) {
			if cast.GetActionID().ItemID == ItemIDTLC {
				return // TLC can't proc Sextant
			}
			if result.Crit && !icd.isOnCD(sim) && sim.Rando.Float64("unmarked") < 0.2 {
				icd = InternalCD(sim.CurrentTime + icdDur)
				AddAuraWithTemporaryStats(sim, agent, MagicIDUnstableCurrents, stats.SpellPower, spellBonus, dur)
			}
		},
	}
}

func ActivateEyeOfMag(sim *Simulation, agent Agent) Aura {
	const spellBonus = 170.0
	const dur = time.Second * 10
	return Aura{
		ID:      MagicIDEyeOfMag,
		Expires: NeverExpires,
		OnSpellMiss: func(sim *Simulation, cast DirectCastAction) {
			AddAuraWithTemporaryStats(sim, agent, MagicIDRecurringPower, stats.SpellPower, spellBonus, dur)
		},
	}
}

func ActivateElderScribes(sim *Simulation, agent Agent) Aura {
	// Gives a chance when your harmful spells land to increase the damage of your spells and effects by up to 130 for 10 sec. (Proc chance: 20%, 50s cooldown)
	icd := NewICD()
	const spellBonus = 130.0
	const dur = time.Second * 10
	const icdDur = time.Second * 50
	const proc = 0.2
	return Aura{
		ID:      MagicIDElderScribe,
		Expires: NeverExpires,
		OnSpellHit: func(sim *Simulation, cast DirectCastAction, result *DirectCastDamageResult) {
			// This code is starting to look a lot like other ICD buff items. Perhaps we could DRY this out.
			if !icd.isOnCD(sim) && sim.Rando.Float64("unmarked") < proc {
				icd = InternalCD(sim.CurrentTime + icdDur)
				AddAuraWithTemporaryStats(sim, agent, MagicIDElderScribeProc, stats.SpellPower, spellBonus, dur)
			}
		},
	}
}

// ActivateFathomBrooch adds an aura that has a chance on cast of nature spell
//  to restore 335 mana. 40s ICD
func ActivateFathomBrooch(sim *Simulation, agent Agent) Aura {
	icd := NewICD()
	const icdDur = time.Second * 40
	return Aura{
		ID:      MagicIDRegainMana,
		Expires: NeverExpires,
		OnCastComplete: func(sim *Simulation, cast DirectCastAction) {
			if icd.isOnCD(sim) {
				return
			}
			if cast.GetSpellSchool() != stats.NatureSpellPower {
				return
			}
			if sim.Rando.Float64("unmarked") < 0.15 {
				icd = InternalCD(sim.CurrentTime + icdDur)
				agent.GetCharacter().Stats[stats.Mana] += 335
			}
		},
	}
}

// ActivateAlchStone adds the alch stone aura that has no effect on casts.
//  The usage for this aura is in the potion usage function.
func ActivateAlchStone(sim *Simulation, agent Agent) Aura {
	return Aura{
		ID:      MagicIDAlchStone,
		Expires: math.MaxInt32,
	}
}

func ActivateMarkOfTheChampion(sim *Simulation, agent Agent) Aura {
	agent.GetCharacter().Stats[stats.SpellPower] += 85
	return Aura{}
}
