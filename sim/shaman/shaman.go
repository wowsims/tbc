package shaman

import (
	"time"

	"github.com/wowsims/tbc/items"
	"github.com/wowsims/tbc/sim/api"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func newShaman(character *core.Character, talents Talents, totems Totems, waterShield bool, agent shamanAgent) *Shaman {
	character.InitialStats[stats.MP5] += 50

	character.InitialStats = character.InitialStats.Add(totems.Stats())

	shaman := &Shaman{
		Character: character,
		agent  :   agent,
		Talents:   talents,

		convectionBonus: 0.02 * float64(talents.Convection),
		concussionBonus: 1 + 0.01*float64(talents.Concussion),
	}
	character.Agent = shaman
	return shaman
}

type Totems struct {
	TotemOfWrath int
	WrathOfAir   bool
	ManaStream   bool
}

func (tt Totems) Stats() stats.Stats {
	totemStats := stats.Stats{}
	if tt.TotemOfWrath > 0 {
		totemStats[stats.SpellCrit] += 66.24 * float64(tt.TotemOfWrath)
		totemStats[stats.SpellHit] += 37.8 * float64(tt.TotemOfWrath)
	}
	if tt.WrathOfAir {
		totemStats[stats.SpellPower] += 101
	}
	if tt.ManaStream {
		totemStats[stats.MP5] += 50
	}
	return totemStats
}

// Agent is shaman specific agent for behavior.
type shamanAgent interface {
	// Returns the action this Agent would like to take next.
	ChooseAction(*Shaman, *core.Simulation) core.AgentAction

	// This will be invoked if the chosen action is actually executed, so the Agent can update its state.
	OnActionAccepted(*Shaman, *core.Simulation, core.AgentAction)

	// Returns this Agent to its initial state.
	Reset(*Shaman, *core.Simulation)
}

// Shaman represents a shaman character.
type Shaman struct {
	*core.Character

	agent shamanAgent

	Talents      Talents     // Shaman Talents

	// HACK HACK HACK
	// TODO: do we actually need a 'on start' method for agents?
	//   This particular use case could also be solved by the 'OnStatAdd' event...
	//    but are there other things we want to do once all buffs are applied right before starting?
	//   Unrelenting storm could also be calculated on the fly if we can allow agents to override the 'Advance' function.
	started bool

	// cache
	convectionBonus float64
	concussionBonus float64
}

func (shaman *Shaman) GetCharacter() *core.Character {
	return shaman.Character
}

// BuffUp lets you buff up all characters in sim (and yourself)
func (shaman *Shaman) BuffUp(sim *core.Simulation) {
	if shaman.Talents.LightningOverload > 0 {
		shaman.AddAura(sim, AuraLightningOverload(shaman.Talents.LightningOverload))
	}
}
func (shaman *Shaman) OnSpellHit(sim *core.Simulation, cast *core.Cast) {
	if cast.Spell.ID == core.MagicIDTLCLB { // TLC does not benefit from shaman talents
		return
	}
	cast.DidDmg *= shaman.concussionBonus // add concussion

	if cast.DidCrit && shaman.Talents.ElementalFocus {
		a := core.Aura{
			ID:             core.MagicIDEleFocus,
			Expires:        sim.CurrentTime + time.Second*15,
			Stacks:         2,
			OnCast:         elementalFocusOnCast,
			OnCastComplete: elementalFocusOnCastComplete,
		}
		shaman.AddAura(sim, a)
	}
}
func (shaman *Shaman) ChooseAction(sim *core.Simulation) core.AgentAction {
	if !shaman.started {
		shaman.started = true
		// we need to apply regen once all buffs are applied.
		shaman.Stats[stats.MP5] += shaman.Stats[stats.Intellect] * (0.02 * float64(shaman.Talents.UnrelentingStorm))
	}
	// Before casting, activate shaman powers!
	TryActivateBloodlust(sim, shaman)
	if shaman.Talents.ElementalMastery {
		TryActivateEleMastery(sim, shaman)
	}

	return shaman.agent.ChooseAction(shaman, sim)
}
func (shaman *Shaman) OnActionAccepted(sim *core.Simulation, action core.AgentAction) {
	shaman.agent.OnActionAccepted(shaman, sim, action)
}
func (shaman *Shaman) Reset(newsim *core.Simulation) {
	shaman.started = false
	shaman.agent.Reset(shaman, newsim)
}

func elementalFocusOnCast(sim *core.Simulation, agent core.Agent, cast *core.Cast) {
	cast.ManaCost *= .6 // reduced by 40%
}

func elementalFocusOnCastComplete(sim *core.Simulation, agent core.Agent, cast *core.Cast) {
	if cast.ManaCost <= 0 {
		return // Don't consume charges from free spells.
	}

	agent.GetCharacter().Auras[core.MagicIDEleFocus].Stacks--
	if agent.GetCharacter().Auras[core.MagicIDEleFocus].Stacks == 0 {
		agent.GetCharacter().RemoveAura(sim, &agent, core.MagicIDEleFocus)
	}
}

type Talents struct {
	ElementalFocus     bool
	LightningMastery   int
	LightningOverload  int
	ElementalPrecision int
	NaturesGuidance    int
	TidalMastery       int
	ElementalMastery   bool
	UnrelentingStorm   int
	CallOfThunder      int
	Convection         int
	Concussion         int
}

func convertShamTalents(t *api.ShamanTalents) Talents {
	return Talents{
		LightningOverload:  int(t.LightningOverload),
		ElementalPrecision: int(t.ElementalPrecision),
		NaturesGuidance:    int(t.NaturesGuidance),
		TidalMastery:       int(t.TidalMastery),
		ElementalMastery:   t.ElementalMastery,
		UnrelentingStorm:   int(t.UnrelentingStorm),
		CallOfThunder:      int(t.CallOfThunder),
		Convection:         int(t.Convection),
		Concussion:         int(t.Concussion),
		LightningMastery:   int(t.LightningMastery),
		ElementalFocus:     t.ElementalFocus,
	}
}

func TryActivateBloodlust(sim *core.Simulation, shaman *Shaman) {
	if shaman.IsOnCD(core.MagicIDBloodlust, sim.CurrentTime) {
		return
	}

	dur := time.Second * 40 // assumes that multiple BLs are different shaman.
	shaman.SetCD(core.MagicIDBloodlust, time.Minute*10+sim.CurrentTime)

	shaman.Party.AddAura(sim, core.Aura{
		ID:      core.MagicIDBloodlust,
		Expires: sim.CurrentTime + dur,
		OnCast: func(sim *core.Simulation, agent core.Agent, c *core.Cast) {
			c.CastTime = (c.CastTime * 10) / 13 // 30% faster
		},
	})
}

// FUTURE: We can cache like 75% of the calculation for a spell cast ahead of time.
//   First time we cast we should create and cache this cast object for better performance.
//   This would get rid of the individual cached floats on Shaman.

// func createBaseCast(character *Shaman, sim *core.Simulation, sp *core.Spell) *core.Cast {
// 	cast := core.NewCast(sim, sp)

// 	if character.Talents.ElementalPrecision > 0 {
// 		// FUTURE: This only impacts "frost fire and nature" spells.
// 		//  We know it doesnt impact TLC.
// 		//  Are there any other spells that a shaman can cast?
// 		cast.BonusHit += float64(character.Talents.ElementalPrecision) * 0.02
// 	}
// 	if character.Talents.NaturesGuidance > 0 {
// 		cast.BonusHit += float64(character.Talents.NaturesGuidance) * 0.01
// 	}
// 	if character.Talents.TidalMastery > 0 {
// 		cast.BonusCrit += float64(character.Talents.TidalMastery) * 0.01
// 	}

// 	// TODO: Should we change these to be full auras?
// 	//   Doesnt seem needed since they can only be used by shaman right here.
// 	if character.Equip[items.ItemSlotRanged].ID == 28248 {
// 		cast.BonusSpellPower += 55
// 	} else if character.Equip[items.ItemSlotRanged].ID == 23199 {
// 		cast.BonusSpellPower += 33
// 	} else if character.Equip[items.ItemSlotRanged].ID == 32330 {
// 		cast.BonusSpellPower += 85
// 	}
// 	if character.Talents.CallOfThunder > 0 { // only applies to CL and LB
// 		cast.BonusCrit += float64(character.Talents.CallOfThunder) * 0.01
// 	}
// 	if sim.Options.Encounter.NumTargets > 1 {
// 		cast.DoItNow = ChainCast
// 	}
// 	cast.ManaCost *= character.convectionBonus

// 	if character.Talents.LightningMastery > 0 {
// 		cast.CastTime -= time.Millisecond * 100 * time.Duration(character.Talents.LightningMastery)
// 	}

// 	return cast
// }

// Totem Item IDs
const (
	TotemOfTheVoid           = 28248
	TotemOfStorms            = 23199
	TotemOfAncestralGuidance = 32330
)

// NewCastAction is how a shaman creates a new spell
//  TODO: Decide if we need separate functions for elemental and enhancement?
func NewCastAction(shaman *Shaman, sim *core.Simulation, sp *core.Spell) core.AgentAction {
	cast := core.NewCast(sim, sp)

	itsElectric := sp.ID == core.MagicIDCL6 || sp.ID == core.MagicIDLB12

	if shaman.Talents.ElementalPrecision > 0 {
		// FUTURE: This only impacts "frost fire and nature" spells.
		//  We know it doesnt impact TLC.
		//  Are there any other spells that a shaman can cast?
		cast.BonusHit += float64(shaman.Talents.ElementalPrecision) * 0.02
	}
	if shaman.Talents.NaturesGuidance > 0 {
		cast.BonusHit += float64(shaman.Talents.NaturesGuidance) * 0.01
	}
	if shaman.Talents.TidalMastery > 0 {
		cast.BonusCrit += float64(shaman.Talents.TidalMastery) * 0.01
	}

	if itsElectric {
		// TODO: Should we change these to be full auras?
		//   Doesnt seem needed since they can only be used by shaman right here.
		if shaman.Equip[items.ItemSlotRanged].ID == TotemOfTheVoid {
			cast.BonusSpellPower += 55
		} else if shaman.Equip[items.ItemSlotRanged].ID == TotemOfStorms {
			cast.BonusSpellPower += 33
		} else if shaman.Equip[items.ItemSlotRanged].ID == TotemOfAncestralGuidance {
			cast.BonusSpellPower += 85
		}
		if shaman.Talents.CallOfThunder > 0 { // only applies to CL and LB
			cast.BonusCrit += float64(shaman.Talents.CallOfThunder) * 0.01
		}
		if sp.ID == core.MagicIDCL6 && sim.Options.Encounter.NumTargets > 1 {
			cast.DoItNow = ChainCast
		}
		if shaman.Talents.LightningMastery > 0 {
			cast.CastTime -= time.Millisecond * 100 * time.Duration(shaman.Talents.LightningMastery)
		}
	}
	cast.CastTime = time.Duration(float64(cast.CastTime) / shaman.HasteBonus())

	// Apply any on cast effects.
	for _, id := range shaman.ActiveAuraIDs {
		if shaman.Auras[id].OnCast != nil {
			shaman.Auras[id].OnCast(sim, shaman, cast)
		}
	}
	if itsElectric { // TODO: Add ElementalFury talent
		// This is written this way to deal with making CSD dmg increase correct.
		// The 'OnCast' auras include CSD
		cast.CritDamageMultipier *= 2 // This handles the 'Elemental Fury' talent which increases the crit bonus.
		cast.CritDamageMultipier -= 1 // reduce to multiplier instead of percent.

		// Convection applies against the base cost of the spell.
		cast.ManaCost -= sp.Mana * shaman.convectionBonus
	}

	return core.AgentAction{
		Wait: 0,
		Cast: cast,
	}
}
