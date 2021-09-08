package core

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func debugFunc(sim *Simulation) func(string, ...interface{}) {
	return func(s string, vals ...interface{}) {
		fmt.Printf("[%0.1f] "+s, append([]interface{}{sim.CurrentTime.Seconds()}, vals...)...)
	}
}

type PlayerAgent interface {
	BuffUp(*Simulation) // Any pre-start buffs to apply to the raid.

	// Returns the action this Agent would like to take next.
	ChooseAction(*Simulation) AgentAction

	// This will be invoked if the chosen action is actually executed, so the Agent can update its state.
	OnActionAccepted(*Simulation, AgentAction)

	// Returns this Agent to its initial state.
	Reset(*Simulation)
}

type Options struct {
	Encounter Encounter
	RSeed     int64
	ExitOnOOM bool
	GCDMin    time.Duration // sets the minimum GCD

	DPSReportTime time.Duration // how many seconds to calculate DPS for.
	Debug         bool          // enables debug printing.
}

type Encounter struct {
	Duration   float64
	NumTargets int
	Armor      int32
}

type Simulation struct {
	Raid     Raid
	Options  Options
	Duration time.Duration

	// Clears and regenerates on each Run call.
	Metrics SimMetrics

	Rando       *rand.Rand
	rseed       int64
	CurrentTime time.Duration // duration that has elapsed in the sim since starting

	Debug func(string, ...interface{})

	// caches to speed up perf and store temp state
	cache *cache
}

type SimMetrics struct {
	TotalDamage    float64
	ReportedDamage float64 // used when DPSReportTime is set
	DamageAtOOM    float64
	OOMAt          float64
	Casts          []*Cast
	ManaSpent      float64
	ManaAtEnd      int
}

// New sim contructs a simulator with the given equipment / options.
func NewSim(raid Raid, options Options) *Simulation {
	if options.GCDMin == 0 {
		options.GCDMin = durationFromSeconds(0.75) // default to 0.75s GCD
	}
	if options.RSeed == 0 {
		options.RSeed = time.Now().Unix()
	}
	// equip := NewEquipmentSet(equipSpec)
	// initialStats := CalculateTotalStats(options, equip)

	sim := &Simulation{
		Options:  options,
		Duration: durationFromSeconds(options.Encounter.Duration),
		rando:    rand.New(rand.NewSource(options.RSeed)),
		Debug:    nil,
		cache: &cache{
			castPool: make([]*Cast, 0, 1000),
		},
	}

	sim.cache.fillCasts()

	for i, eq := range equip {
		if eq.Activate != nil {
			sim.activeEquip = append(sim.activeEquip, &equip[i])
		}
		for _, g := range eq.Gems {
			if g.Activate != nil {
				sim.activeEquip = append(sim.activeEquip, &equip[i])
			}
		}
	}

	sim.Agent = NewAgent(sim, options.AgentType)
	if sim.Agent == nil {
		fmt.Printf("[ERROR] No rotation given to sim.\n")
		return nil
	}

	return sim
}

// reset will set sim back and erase all current state.
// This is automatically called before every 'Run'
//  This includes resetting and reactivating always on trinkets, auras, set bonuses, etc
func (sim *Simulation) reset() {
	// sim.rseed++
	// sim.rando.Seed(sim.rseed)

	sim.Stats = sim.InitialStats
	sim.cache.destructionPotion = false
	sim.cache.bloodlustCasts = 0
	sim.CurrentTime = 0.0
	sim.CurrentMana = sim.Stats[StatMana]
	sim.CDs = [MagicIDLen]time.Duration{}
	sim.auras = [MagicIDLen]Aura{}
	if len(sim.activeAuraIDs) > 0 {
		sim.activeAuraIDs = sim.activeAuraIDs[len(sim.activeAuraIDs)-1:] // chop off end of activeids slice, faster than making a new one
	}
	sim.metrics = SimMetrics{
		Casts: make([]*Cast, 0, 1000),
	}

	if sim.Debug != nil {
		sim.Debug("SIM RESET\n")
		sim.Debug("----------------------\n")
	}

	// Activate all talents
	if sim.Options.Talents.LightningOverload > 0 {
		sim.addAura(AuraLightningOverload(sim.Options.Talents.LightningOverload))
	}

	// Chain lightning bounces
	if sim.Options.Encounter.NumClTargets > 1 {
		sim.addAura(ActivateChainLightningBounce(sim))
	}

	// Judgement of Wisdom
	if sim.Options.Buffs.JudgementOfWisdom {
		sim.addAura(AuraJudgementOfWisdom())
	}

	// Activate all permanent item effects.
	for _, item := range sim.activeEquip {
		if item.Activate != nil && item.ActivateCD == neverExpires {
			sim.addAura(item.Activate(sim))
		}
		for _, g := range item.Gems {
			if g.Activate != nil {
				sim.addAura(g.Activate(sim))
			}
		}
	}

	sim.ActivateSets()
	sim.Agent.Reset(sim)

	// Currently no temp hit increase effects in the sim... so lets cache this!
	sim.cache.spellHit = 0.83 + sim.Stats[StatSpellHit]/1260.0 // 12.6 hit == 1% hit
}

// Activates set bonuses, returning the list of active bonuses.
func (sim *Simulation) ActivateSets() []string {
	active := []string{}
	// Activate Set Bonuses
	setItemCount := map[string]int{}
	for _, i := range sim.Equip {
		set := itemSetLookup[i.ID]
		if set != nil {
			setItemCount[set.Name]++
			if bonus, ok := set.Bonuses[setItemCount[set.Name]]; ok {
				active = append(active, set.Name+" ("+strconv.Itoa(setItemCount[set.Name])+"pc)")
				sim.addAura(bonus(sim))
			}
		}
	}

	return active
}

// Run will run the simulation for number of seconds.
// Returns metrics for what was cast and how much damage was done.
func (sim *Simulation) Run() SimMetrics {
	sim.reset()

	for sim.CurrentTime < sim.Duration {
		TryActivateDrums(sim)
		TryActivateBloodlust(sim)
		TryActivateEleMastery(sim)
		TryActivateRacial(sim)
		TryActivateDestructionPotion(sim)
		sim.TryActivateEquipment()

		TryActivateDarkRune(sim)
		TryActivateSuperManaPotion(sim)

		// Choose next action
		action := sim.Agent.ChooseAction(sim)
		if action.Wait != 0 {
			sim.Agent.OnActionAccepted(sim, action)
			sim.Advance(action.Wait)
			continue
		}

		castingSpell := action.Cast
		if castingSpell == nil {
			panic("Agent returned nil action")
		}
		if castingSpell.CastTime < sim.GCDMin {
			castingSpell.CastTime = sim.GCDMin
		}

		if sim.CurrentMana >= castingSpell.ManaCost {
			if sim.Debug != nil {
				sim.Debug("Start Casting %s Cast Time: %0.1fs\n", castingSpell.Spell.Name, castingSpell.CastTime.Seconds())
			}

			sim.Agent.OnActionAccepted(sim, action)
			sim.Advance(castingSpell.CastTime)
			sim.Cast(castingSpell)
		} else {
			// Not enough mana, wait until there is enough mana to cast the desired spell
			if sim.metrics.OOMAt == 0 {
				sim.metrics.OOMAt = sim.CurrentTime.Seconds()
				sim.metrics.DamageAtOOM = sim.metrics.TotalDamage
				if sim.Options.ExitOnOOM {
					return sim.metrics
				}
			}
			timeUntilRegen := durationFromSeconds((castingSpell.ManaCost-sim.CurrentMana)/sim.manaRegenPerSecond()) + 1
			sim.Advance(timeUntilRegen)
			// Don't actually cast; let the next iteration do the cast, so we recheck for pots/CDs/etc
		}
	}
	sim.metrics.ManaAtEnd = int(sim.CurrentMana)

	return sim.metrics
}

// Advance moves time forward counting down auras, CDs, mana regen, etc
func (sim *Simulation) Advance(elapsedTime time.Duration) {
	newTime := sim.CurrentTime + elapsedTime

	// MP5 regen
	sim.CurrentMana = math.Min(
		sim.Stats[StatMana],
		sim.CurrentMana+sim.manaRegenPerSecond()*elapsedTime.Seconds())

	// Go in reverse order so we can safely delete while looping
	for i := len(sim.activeAuraIDs) - 1; i >= 0; i-- {
		id := sim.activeAuraIDs[i]
		if sim.auras[id].Expires != 0 && sim.auras[id].Expires <= newTime {
			sim.removeAura(id)
		}
	}
	sim.CurrentTime = newTime
}

// Cast will actually cast and treat all casts as having no 'flight time'.
// This will activate any auras around casting, calculate hit/crit and add to sim metrics.
func (sim *Simulation) Cast(cast *Cast) {
	if sim.Debug != nil {
		sim.Debug("Current Mana %0.0f, Cast Cost: %0.0f\n", sim.CurrentMana, cast.ManaCost)
	}
	sim.CurrentMana -= cast.ManaCost
	sim.metrics.ManaSpent += cast.ManaCost

	for _, id := range sim.activeAuraIDs {
		if sim.auras[id].OnCastComplete != nil {
			sim.auras[id].OnCastComplete(sim, cast)
		}
	}
	hit := sim.cache.spellHit + cast.Hit // 12.6 hit == 1% hit
	hit = math.Min(hit, 0.99)            // can't get away from the 1% miss

	dbgCast := cast.Spell.Name
	if sim.Debug != nil {
		sim.Debug("Completed Cast (%s)\n", dbgCast)
	}
	if sim.rando.Float64() < hit {
		dmg := (sim.rando.Float64() * cast.Spell.DmgDiff) + cast.Spell.MinDmg + (sim.Stats[StatSpellDmg] * cast.Spell.Coeff)
		if cast.DidDmg != 0 { // use the pre-set dmg
			dmg = cast.DidDmg
		}
		cast.DidHit = true

		itsElectric := cast.Spell.ID == MagicIDCL6 || cast.Spell.ID == MagicIDLB12

		crit := (sim.Stats[StatSpellCrit] / 2208.0) + cast.Crit // 22.08 crit == 1% crit
		if sim.rando.Float64() < crit {
			cast.DidCrit = true
			critBonus := 1.5 // fall back crit damage
			if cast.CritBonus != 0 {
				critBonus = cast.CritBonus // This means we had pre-set the crit bonus when the spell was created. CSD will modify this.
			}
			if itsElectric {
				critBonus *= 2 // This handles the 'Elemental Fury' talent which increases the crit bonus.
				critBonus -= 1 // reduce to multiplier instead of percent.
			}
			dmg *= critBonus
			if cast.Spell.ID != MagicIDTLCLB {
				// TLC does not proc focus.
				sim.addAura(AuraElementalFocus(sim))
			}
			if sim.Debug != nil {
				dbgCast += " crit"
			}
		} else if sim.Debug != nil {
			dbgCast += " hit"
		}

		if itsElectric {
			dmg *= sim.cache.elcDmgBonus
		} else {
			dmg *= sim.cache.dmgBonus
		}

		// Average Resistance (AR) = (Target's Resistance / (Caster's Level * 5)) * 0.75
		// P(x) = 50% - 250%*|x - AR| <- where X is %resisted
		// Using these stats:
		//    13.6% chance of
		resVal := sim.rando.Float64()
		if resVal < 0.18 { // 13% chance for 25% resist, 4% for 50%, 1% for 75%
			if sim.Debug != nil {
				dbgCast += " (partial resist: "
			}
			if resVal < 0.01 {
				dmg *= .25
				if sim.Debug != nil {
					dbgCast += "75%)"
				}
			} else if resVal < 0.05 {
				dmg *= .5
				if sim.Debug != nil {
					dbgCast += "50%)"
				}
			} else {
				dmg *= .75
				if sim.Debug != nil {
					dbgCast += "25%)"
				}
			}
		}
		cast.DidDmg = dmg
		// Apply any effects specific to this cast.
		if cast.Effect != nil {
			cast.Effect(sim, cast)
		}
		// Apply any on spell hit effects.
		for _, id := range sim.activeAuraIDs {
			if sim.auras[id].OnSpellHit != nil {
				sim.auras[id].OnSpellHit(sim, cast)
			}
		}
	} else {
		if sim.Debug != nil {
			dbgCast += " miss"
		}
		cast.DidDmg = 0
		cast.DidCrit = false
		cast.DidHit = false
		for _, id := range sim.activeAuraIDs {
			if sim.auras[id].OnSpellMiss != nil {
				sim.auras[id].OnSpellMiss(sim, cast)
			}
		}
	}

	if cast.Spell.Cooldown > 0 {
		sim.setCD(cast.Spell.ID, cast.Spell.Cooldown)
	}

	if sim.Debug != nil {
		sim.Debug("%s: %0.0f\n", dbgCast, cast.DidDmg)
	}

	sim.metrics.Casts = append(sim.metrics.Casts, cast)

	sim.metrics.TotalDamage += cast.DidDmg
	if sim.Options.DPSReportTime > 0 && sim.CurrentTime <= sim.cache.dpsReportTime {
		sim.metrics.ReportedDamage += cast.DidDmg
	}
}

// addAura will add a new aura to the simulation. If there is a matching aura ID
// it will be replaced with the newer aura.
// Auras with duration of 0 will be logged as activating but never added to simulation auras.
func (sim *Simulation) addAura(newAura Aura) {
	if sim.Debug != nil {
		sim.Debug(" +%s\n", AuraName(newAura.ID))
	}
	if newAura.Expires < sim.CurrentTime {
		return // no need to waste time adding aura that doesn't last.
	}

	if sim.hasAura(newAura.ID) {
		sim.removeAura(newAura.ID)
	}

	sim.auras[newAura.ID] = newAura
	sim.auras[newAura.ID].activeIndex = int32(len(sim.activeAuraIDs))
	sim.activeAuraIDs = append(sim.activeAuraIDs, newAura.ID)
}

// Remove an aura by its ID
func (sim *Simulation) removeAura(id int32) {
	if sim.auras[id].OnExpire != nil {
		sim.auras[id].OnExpire(sim, nil)
	}
	removeActiveIndex := sim.auras[id].activeIndex
	sim.auras[id] = Aura{}

	// Overwrite the element being removed with the last element
	otherAuraID := sim.activeAuraIDs[len(sim.activeAuraIDs)-1]
	if id != otherAuraID {
		sim.activeAuraIDs[removeActiveIndex] = otherAuraID
		sim.auras[otherAuraID].activeIndex = removeActiveIndex
	}

	// Now we can remove the last element, in constant time
	sim.activeAuraIDs = sim.activeAuraIDs[:len(sim.activeAuraIDs)-1]

	if sim.Debug != nil {
		sim.Debug(" -%s\n", AuraName(id))
	}
}

// Returns whether an aura with the given ID is currently active.
func (sim *Simulation) hasAura(id int32) bool {
	return sim.auras[id].ID != 0
}

// Returns rate of mana regen, as mana / second
func (sim *Simulation) manaRegenPerSecond() float64 {
	return sim.Stats[StatMP5] / 5.0
}

func (sim *Simulation) isOnCD(magicID int32) bool {
	return sim.CDs[magicID] > sim.CurrentTime
}

func (sim *Simulation) getRemainingCD(magicID int32) time.Duration {
	remainingCD := sim.CDs[magicID] - sim.CurrentTime
	if remainingCD > 0 {
		return remainingCD
	} else {
		return 0
	}
}

func (sim *Simulation) setCD(magicID int32, newCD time.Duration) {
	sim.CDs[magicID] = sim.CurrentTime + newCD
}

// Pops any on-use trinkets / gear
func (sim *Simulation) TryActivateEquipment() {
	for _, item := range sim.activeEquip {
		if item.Activate == nil || item.ActivateCD == neverExpires { // ignore non-activatable, and always active items.
			continue
		}
		if sim.isOnCD(item.CoolID) {
			continue
		}
		if item.Slot == EquipTrinket && sim.isOnCD(MagicIDAllTrinket) {
			continue
		}
		sim.addAura(item.Activate(sim))
		sim.setCD(item.CoolID, item.ActivateCD)
		if item.Slot == EquipTrinket {
			sim.setCD(MagicIDAllTrinket, time.Second*30)
		}
	}
}

func durationFromSeconds(numSeconds float64) time.Duration {
	return time.Duration(float64(time.Second) * numSeconds)
}
