package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/api"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func NewElementalShaman(character *core.Character, options *api.PlayerOptions, buffs core.Buffs) *Shaman {
	eleShamOptions := options.GetElementalShaman()
	talents := convertShamTalents(eleShamOptions.Talents)

	// TODO: Probably should get this from shaman options rather than buffs.
	// However, other classes will need totem buffs so it has to be on buffs too.
	totems := Totems{
		TotemOfWrath: int(buffs.TotemOfWrath),
		WrathOfAir:   buffs.WrathOfAirTotem != api.TristateEffect_TristateEffectMissing,
		ManaStream:   buffs.ManaSpringTotem != api.TristateEffect_TristateEffectMissing,
	}

	var agent shamanAgent

	switch eleShamOptions.Agent.Type {
	case api.ElementalShaman_Agent_Adaptive:
		agent = NewAdaptiveAgent()
	case api.ElementalShaman_Agent_CLOnClearcast:
		agent = NewCLOnClearcastAgent()
	case api.ElementalShaman_Agent_FixedLBCL:
		agent = NewLBOnlyAgent()
		// TODO: Add option for this
		//numLB := agentOptions["numLBtoCL"]
		//if numLB == -1 {
		//	agent = NewLBOnlyAgent()
		//} else {
		//	agent = NewFixedRotationAgent(numLB)
		//}
	case api.ElementalShaman_Agent_CLOnCD:
		agent = NewCLOnCDAgent()
	}

	return newShaman(character, talents, totems, eleShamOptions.Options.WaterShield, agent)
}

func loDmgMod(sim *core.Simulation, agent core.Agent, c *core.Cast) {
	c.DidDmg /= 2
}

const (
	CastTagLightningOverload int32 = 1 // This could be value or bitflag if we ended up needing multiple flags at the same time.
)

func AuraLightningOverload(lvl int) core.Aura {
	chance := 0.04 * float64(lvl)
	return core.Aura{
		ID:      core.MagicIDLOTalent,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, agent core.Agent, c *core.Cast) {
			if c.Spell.ID != core.MagicIDLB12 && c.Spell.ID != core.MagicIDCL6 {
				return
			}
			if c.Tag == CastTagLightningOverload {
				return // can't proc LO on LO
			}
			actualChance := chance
			if c.Spell.ID == core.MagicIDCL6 {
				actualChance /= 3 // 33% chance of regular for CL LO
			}
			if sim.Rando.Float64("LO") < actualChance {
				if sim.Log != nil {
					sim.Log(" - Lightning Overload -\n")
				}
				clone := sim.NewCast()
				// Don't set IsClBounce even if this is a bounce, so that the clone does a normal CL and bounces
				clone.Tag = CastTagLightningOverload
				clone.Spell = c.Spell

				// Clone dmg/hit/crit chance?
				clone.BonusHit = c.BonusHit
				clone.BonusCrit = c.BonusCrit
				clone.BonusSpellPower = c.BonusSpellPower

				clone.CritDamageMultipier = c.CritDamageMultipier
				clone.Effect = loDmgMod

				// Use the cast function from the original cast.
				clone.DoItNow = c.DoItNow
				clone.DoItNow(sim, agent, clone)
				if sim.Log != nil {
					sim.Log(" - Lightning Overload Complete -\n")
				}
			}
		},
	}
}

func TryActivateEleMastery(sim *core.Simulation, agent core.Agent) {
	if agent.GetCharacter().IsOnCD(core.MagicIDEleMastery, sim.CurrentTime) {
		return
	}

	agent.GetCharacter().AddAura(sim, core.Aura{
		ID:      core.MagicIDEleMastery,
		Expires: core.NeverExpires,
		OnCast: func(sim *core.Simulation, agent core.Agent, c *core.Cast) {
			c.ManaCost = 0
			c.BonusCrit += 1.01
		},
		OnCastComplete: func(sim *core.Simulation, agent core.Agent, c *core.Cast) {
			// Remove the buff and put skill on CD
			agent.GetCharacter().SetCD(core.MagicIDEleMastery, time.Second*180+sim.CurrentTime)
			agent.GetCharacter().RemoveAura(sim, &agent, core.MagicIDEleMastery)
		},
	})
}

// ################################################################
//                              LB ONLY
// ################################################################
type LBOnlyAgent struct {
	lb *core.Spell
}

func (agent *LBOnlyAgent) ChooseAction(shaman* Shaman, sim *core.Simulation) core.AgentAction {
	return NewCastAction(shaman, sim, agent.lb)
}

func (agent *LBOnlyAgent) OnActionAccepted(shaman *Shaman, sim *core.Simulation, action core.AgentAction) {
}
func (agent *LBOnlyAgent) Reset(shaman *Shaman, sim *core.Simulation) {}

func NewLBOnlyAgent() *LBOnlyAgent {
	return &LBOnlyAgent{
		lb: core.Spells[core.MagicIDLB12],
	}
}

// ################################################################
//                             CL ON CD
// ################################################################
type CLOnCDAgent struct {
	lb *core.Spell
	cl *core.Spell
}

func (agent *CLOnCDAgent) ChooseAction(shaman *Shaman, sim *core.Simulation) core.AgentAction {
	if shaman.IsOnCD(core.MagicIDCL6, sim.CurrentTime) {
		// sim.Log("[CLonCD] LB\n")
		return NewCastAction(shaman, sim, agent.lb)
	} else {
		// sim.Log("[CLonCD] CL\n")
		return NewCastAction(shaman, sim, agent.cl)
	}
}

func (agent *CLOnCDAgent) OnActionAccepted(shaman *Shaman, sim *core.Simulation, action core.AgentAction) {
}
func (agent *CLOnCDAgent) Reset(shaman *Shaman, sim *core.Simulation) {}

func NewCLOnCDAgent() *CLOnCDAgent {
	return &CLOnCDAgent{
		lb: core.Spells[core.MagicIDLB12],
		cl: core.Spells[core.MagicIDCL6],
	}
}

// ################################################################
//                          FIXED ROTATION
// ################################################################
type FixedRotationAgent struct {
	numLBsPerCL       int
	numLBsSinceLastCL int
	lb                *core.Spell
	cl                *core.Spell
}

// Returns if any temporary haste buff is currently active.
// TODO: Figure out a way to make this automatic
func (agent *FixedRotationAgent) temporaryHasteActive(shaman *Shaman) bool {
	return shaman.HasAura(core.MagicIDBloodlust) ||
		shaman.HasAura(core.MagicIDDrums) ||
		shaman.HasAura(core.MagicIDTrollBerserking) ||
		shaman.HasAura(core.MagicIDSkullGuldan) ||
		shaman.HasAura(core.MagicIDFungalFrenzy)
}

func (agent *FixedRotationAgent) ChooseAction(shaman *Shaman, sim *core.Simulation) core.AgentAction {
	if agent.numLBsSinceLastCL < agent.numLBsPerCL {
		return NewCastAction(shaman, sim, agent.lb)
	}

	if !shaman.IsOnCD(core.MagicIDCL6, sim.CurrentTime) {
		return NewCastAction(shaman, sim, agent.cl)
	}

	// If we have a temporary haste effect (like bloodlust or quags eye) then
	// we should add LB casts instead of waiting
	if agent.temporaryHasteActive(shaman) {
		return NewCastAction(shaman, sim, agent.lb)
	}

	return core.AgentAction{Wait: shaman.GetRemainingCD(core.MagicIDCL6, sim.CurrentTime)}
}

func (agent *FixedRotationAgent) OnActionAccepted(shaman *Shaman, sim *core.Simulation, action core.AgentAction) {
	if action.Cast == nil {
		return
	}

	if action.Cast.Spell.ID == core.MagicIDLB12 {
		agent.numLBsSinceLastCL++
	} else if action.Cast.Spell.ID == core.MagicIDCL6 {
		agent.numLBsSinceLastCL = 0
	}
}

func (agent *FixedRotationAgent) Reset(shaman *Shaman, sim *core.Simulation) {
	agent.numLBsSinceLastCL = agent.numLBsPerCL
}

func NewFixedRotationAgent(numLBsPerCL int) *FixedRotationAgent {
	return &FixedRotationAgent{
		numLBsPerCL:       numLBsPerCL,
		numLBsSinceLastCL: numLBsPerCL, // This lets us cast CL first
		lb:                core.Spells[core.MagicIDLB12],
		cl:                core.Spells[core.MagicIDCL6],
	}
}

// ################################################################
//                          CL ON CLEARCAST
// ################################################################
type CLOnClearcastAgent struct {
	// Whether the second-to-last spell procced clearcasting
	prevPrevCastProccedCC bool

	lb *core.Spell
	cl *core.Spell
}

func (agent *CLOnClearcastAgent) ChooseAction(shaman *Shaman, sim *core.Simulation) core.AgentAction {
	if shaman.IsOnCD(core.MagicIDCL6, sim.CurrentTime) || !agent.prevPrevCastProccedCC {
		// sim.Log("[CLonCC] - LB")
		return NewCastAction(shaman, sim, agent.lb)
	}

	// sim.Log("[CLonCC] - CL")
	return NewCastAction(shaman, sim, agent.cl)
}

func (agent *CLOnClearcastAgent) OnActionAccepted(shaman *Shaman, sim *core.Simulation, action core.AgentAction) {
	agent.prevPrevCastProccedCC = shaman.Auras[core.MagicIDEleFocus].Stacks == 2
}

func (agent *CLOnClearcastAgent) Reset(shaman *Shaman, sim *core.Simulation) {
	agent.prevPrevCastProccedCC = true // Lets us cast CL first
}

func NewCLOnClearcastAgent() *CLOnClearcastAgent {
	return &CLOnClearcastAgent{
		lb: core.Spells[core.MagicIDLB12],
		cl: core.Spells[core.MagicIDCL6],
	}
}

// ################################################################
//                             ADAPTIVE
// ################################################################
type AdaptiveAgent struct {
	// Circular array buffer for recent mana snapshots, within a time window
	manaSnapshots      [manaSnapshotsBufferSize]ManaSnapshot
	numSnapshots       int32
	firstSnapshotIndex int32
	timesOOM           int  // count of times gone oom.
	wentOOM            bool // if agent went OOM this time.

	baseAgent    shamanAgent // The agent used most of the time
	surplusAgent shamanAgent // The agent used when we have extra mana
}

const manaSpendingWindowNumSeconds = 60
const manaSpendingWindow = time.Second * manaSpendingWindowNumSeconds

// 2 * (# of seconds) should be plenty of slots
const manaSnapshotsBufferSize = manaSpendingWindowNumSeconds * 2

type ManaSnapshot struct {
	time      time.Duration // time this snapshot was taken
	manaSpent float64       // total amount of mana spent up to this time
}

func (agent *AdaptiveAgent) getOldestSnapshot() ManaSnapshot {
	return agent.manaSnapshots[agent.firstSnapshotIndex]
}

func (agent *AdaptiveAgent) purgeExpiredSnapshots(sim *core.Simulation) {
	expirationCutoff := sim.CurrentTime - manaSpendingWindow

	curIndex := agent.firstSnapshotIndex
	for agent.numSnapshots > 0 && agent.manaSnapshots[curIndex].time < expirationCutoff {
		curIndex = (curIndex + 1) % manaSnapshotsBufferSize
		agent.numSnapshots--
	}
	agent.firstSnapshotIndex = curIndex
}

func (agent *AdaptiveAgent) takeSnapshot(sim *core.Simulation, shaman *Shaman) {
	if agent.numSnapshots >= manaSnapshotsBufferSize {
		panic("Agent snapshot buffer full")
	}

	snapshot := ManaSnapshot{
		time:      sim.CurrentTime,
		manaSpent: sim.Metrics.IndividualMetrics[shaman.ID].ManaSpent,
	}

	nextIndex := (agent.firstSnapshotIndex + agent.numSnapshots) % manaSnapshotsBufferSize
	agent.manaSnapshots[nextIndex] = snapshot
	agent.numSnapshots++
}

func (agent *AdaptiveAgent) ChooseAction(shaman *Shaman, sim *core.Simulation) core.AgentAction {
	agent.purgeExpiredSnapshots(sim)
	oldestSnapshot := agent.getOldestSnapshot()

	manaSpent := 0.0
	if len(sim.Metrics.IndividualMetrics) > shaman.ID {
		manaSpent = sim.Metrics.IndividualMetrics[shaman.ID].ManaSpent - oldestSnapshot.manaSpent
	}
	timeDelta := sim.CurrentTime - oldestSnapshot.time
	if timeDelta == 0 {
		timeDelta = 1
	}

	timeRemaining := sim.Duration - sim.CurrentTime
	projectedManaCost := manaSpent * (timeRemaining.Seconds() / timeDelta.Seconds())

	if sim.Log != nil {
		manaSpendingRate := manaSpent / timeDelta.Seconds()
		sim.Log("[AI] CL Ready: Mana/s: %0.1f, Est Mana Cost: %0.1f, CurrentMana: %0.1f\n", manaSpendingRate, projectedManaCost, shaman.Stats[stats.Mana])
	}

	// If we have enough mana to burn, use the surplus agent.
	if projectedManaCost < shaman.Stats[stats.Mana] {
		return agent.surplusAgent.ChooseAction(shaman, sim)
	} else {
		return agent.baseAgent.ChooseAction(shaman, sim)
	}
}
func (agent *AdaptiveAgent) OnActionAccepted(shaman *Shaman, sim *core.Simulation, action core.AgentAction) {
	if !agent.wentOOM && action.Cast != nil && action.Cast.ManaCost > shaman.Stats[stats.Mana] {
		agent.timesOOM++
		agent.wentOOM = true
	}
	agent.takeSnapshot(sim, shaman)
	agent.baseAgent.OnActionAccepted(shaman, sim, action)
	agent.surplusAgent.OnActionAccepted(shaman, sim, action)
}

func (agent *AdaptiveAgent) Reset(shaman *Shaman, sim *core.Simulation) {
	if agent.timesOOM == 5 {
		agent.baseAgent = NewLBOnlyAgent()
		agent.surplusAgent = NewCLOnClearcastAgent()
	}
	agent.wentOOM = false
	agent.manaSnapshots = [manaSnapshotsBufferSize]ManaSnapshot{}
	agent.firstSnapshotIndex = 0
	agent.numSnapshots = 0
	agent.baseAgent.Reset(shaman, sim)
	agent.surplusAgent.Reset(shaman, sim)
}

func NewAdaptiveAgent() *AdaptiveAgent {
	agent := &AdaptiveAgent{}

	// TODO: Can we just start with more aggressive agent and drop to less aggressive if we go OOM 5 times?
	//   not as deterministic... but probably averages out the same?
	// Otherwise we need to figure out how to do this after all other agents are setup (in the eventual 'raid' sim setup)

	agent.baseAgent = NewCLOnClearcastAgent()
	agent.surplusAgent = NewCLOnCDAgent()

	return agent
}

// ChainCast is how to cast chain lightning.
func ChainCast(sim *core.Simulation, agent core.Agent, cast *core.Cast) {
	shaman := agent.(*Shaman)
	core.DirectCast(sim, agent, cast) // Start with a normal direct cast to start.

	// Now chain
	dmgCoeff := 1.0
	if cast.Tag == CastTagLightningOverload {
		dmgCoeff = 0.5
	}
	for i := 1; i < sim.Options.Encounter.NumTargets; i++ {
		if shaman.HasAura(core.MagicIDTidefury) {
			dmgCoeff *= 0.83
		} else {
			dmgCoeff *= 0.7
		}
		clone := &core.Cast{
			Tag:                 cast.Tag, // pass along lightning overload
			Spell:               cast.Spell,
			BonusCrit:           cast.BonusCrit,
			CritDamageMultipier: cast.CritDamageMultipier,
			Effect:              func(sim *core.Simulation, agent core.Agent, c *core.Cast) { cast.DidDmg *= dmgCoeff },
			DoItNow:             core.DirectCast,
		}
		clone.DoItNow(sim, shaman, clone)
	}
}
