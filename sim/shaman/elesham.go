package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

type ElementalSpec struct {
	Talents      Talents
	Totems       Totems
	AgentID      AgentType
	AgentOptions map[string]int
}

func (es ElementalSpec) CreateAgent(player *core.Player, party *core.Party) core.Agent {
	return NewShaman(player, party, es.Talents, es.Totems, es.AgentID, es.AgentOptions)
}

func loDmgMod(sim *core.Simulation, p core.PlayerAgent, c *core.Cast) {
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
		OnSpellHit: func(sim *core.Simulation, p core.PlayerAgent, c *core.Cast) {
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
				if sim.Debug != nil {
					sim.Debug(" - Lightning Overload -\n")
				}
				clone := sim.NewCast()
				// Don't set IsClBounce even if this is a bounce, so that the clone does a normal CL and bounces
				clone.Tag = CastTagLightningOverload
				clone.Spell = c.Spell

				// Clone dmg/hit/crit chance?
				clone.Hit = c.Hit
				clone.Crit = c.Crit
				clone.Dmg = c.Dmg

				clone.CritBonus = c.CritBonus
				clone.Effect = loDmgMod

				// Use the cast function from the original cast.
				clone.DoItNow = c.DoItNow
				clone.DoItNow(sim, p, clone)
				if sim.Debug != nil {
					sim.Debug(" - Lightning Overload Complete -\n")
				}
			}
		},
	}
}

func TryActivateEleMastery(sim *core.Simulation, player *core.Player) {
	if player.IsOnCD(core.MagicIDEleMastery, sim.CurrentTime) {
		return
	}

	player.AddAura(sim, core.Aura{
		ID:      core.MagicIDEleMastery,
		Expires: core.NeverExpires,
		OnCast: func(sim *core.Simulation, p core.PlayerAgent, c *core.Cast) {
			c.ManaCost = 0
			c.Crit += 1.01
		},
		OnCastComplete: func(sim *core.Simulation, p core.PlayerAgent, c *core.Cast) {
			// Remove the buff and put skill on CD
			p.SetCD(core.MagicIDEleMastery, time.Second*180+sim.CurrentTime)
			p.RemoveAura(sim, p, core.MagicIDEleMastery)
		},
	})
}

// ################################################################
//                              LB ONLY
// ################################################################
type LBOnlyAgent struct {
	lb *core.Spell
}

func (agent *LBOnlyAgent) ChooseAction(s *Shaman, party *core.Party, sim *core.Simulation) core.AgentAction {
	return NewCastAction(sim, s, agent.lb)
}

func (agent *LBOnlyAgent) OnActionAccepted(p *Shaman, sim *core.Simulation, action core.AgentAction) {
}
func (agent *LBOnlyAgent) Reset(sim *core.Simulation) {}

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

func (agent *CLOnCDAgent) ChooseAction(s *Shaman, party *core.Party, sim *core.Simulation) core.AgentAction {
	if s.IsOnCD(core.MagicIDCL6, sim.CurrentTime) {
		// sim.Debug("[CLonCD] LB\n")
		return NewCastAction(sim, s, agent.lb)
	} else {
		// sim.Debug("[CLonCD] CL\n")
		return NewCastAction(sim, s, agent.cl)
	}
}

func (agent *CLOnCDAgent) OnActionAccepted(p *Shaman, sim *core.Simulation, action core.AgentAction) {
}
func (agent *CLOnCDAgent) Reset(sim *core.Simulation) {}

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
func (agent *FixedRotationAgent) temporaryHasteActive(s *Shaman) bool {
	return s.HasAura(core.MagicIDBloodlust) ||
		s.HasAura(core.MagicIDDrums) ||
		s.HasAura(core.MagicIDTrollBerserking) ||
		s.HasAura(core.MagicIDSkullGuldan) ||
		s.HasAura(core.MagicIDFungalFrenzy)
}

func (agent *FixedRotationAgent) ChooseAction(s *Shaman, party *core.Party, sim *core.Simulation) core.AgentAction {
	if agent.numLBsSinceLastCL < agent.numLBsPerCL {
		return NewCastAction(sim, s, agent.lb)
	}

	if !s.IsOnCD(core.MagicIDCL6, sim.CurrentTime) {
		return NewCastAction(sim, s, agent.cl)
	}

	// If we have a temporary haste effect (like bloodlust or quags eye) then
	// we should add LB casts instead of waiting
	if agent.temporaryHasteActive(s) {
		return NewCastAction(sim, s, agent.lb)
	}

	return core.AgentAction{Wait: s.GetRemainingCD(core.MagicIDCL6, sim.CurrentTime)}
}

func (agent *FixedRotationAgent) OnActionAccepted(s *Shaman, sim *core.Simulation, action core.AgentAction) {
	if action.Cast == nil {
		return
	}

	if action.Cast.Spell.ID == core.MagicIDLB12 {
		agent.numLBsSinceLastCL++
	} else if action.Cast.Spell.ID == core.MagicIDCL6 {
		agent.numLBsSinceLastCL = 0
	}
}

func (agent *FixedRotationAgent) Reset(sim *core.Simulation) {
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

func (agent *CLOnClearcastAgent) ChooseAction(s *Shaman, party *core.Party, sim *core.Simulation) core.AgentAction {
	if s.IsOnCD(core.MagicIDCL6, sim.CurrentTime) || !agent.prevPrevCastProccedCC {
		// sim.Debug("[CLonCC] - LB")
		return NewCastAction(sim, s, agent.lb)
	}

	// sim.Debug("[CLonCC] - CL")
	return NewCastAction(sim, s, agent.cl)
}

func (agent *CLOnClearcastAgent) OnActionAccepted(p *Shaman, sim *core.Simulation, action core.AgentAction) {
	agent.prevPrevCastProccedCC = p.Auras[core.MagicIDEleFocus].Stacks == 2
}

func (agent *CLOnClearcastAgent) Reset(sim *core.Simulation) {
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

func (agent *AdaptiveAgent) takeSnapshot(sim *core.Simulation, s *Shaman) {
	if agent.numSnapshots >= manaSnapshotsBufferSize {
		panic("Agent snapshot buffer full")
	}

	snapshot := ManaSnapshot{
		time:      sim.CurrentTime,
		manaSpent: sim.Metrics.IndividualMetrics[s.ID].ManaSpent,
	}

	nextIndex := (agent.firstSnapshotIndex + agent.numSnapshots) % manaSnapshotsBufferSize
	agent.manaSnapshots[nextIndex] = snapshot
	agent.numSnapshots++
}

func (agent *AdaptiveAgent) ChooseAction(s *Shaman, party *core.Party, sim *core.Simulation) core.AgentAction {
	agent.purgeExpiredSnapshots(sim)
	oldestSnapshot := agent.getOldestSnapshot()

	manaSpent := 0.0
	if len(sim.Metrics.IndividualMetrics) > s.ID {
		manaSpent = sim.Metrics.IndividualMetrics[s.ID].ManaSpent - oldestSnapshot.manaSpent
	}
	timeDelta := sim.CurrentTime - oldestSnapshot.time
	if timeDelta == 0 {
		timeDelta = 1
	}

	timeRemaining := sim.Duration - sim.CurrentTime
	projectedManaCost := manaSpent * (timeRemaining.Seconds() / timeDelta.Seconds())

	if sim.Debug != nil {
		manaSpendingRate := manaSpent / timeDelta.Seconds()
		sim.Debug("[AI] CL Ready: Mana/s: %0.1f, Est Mana Cost: %0.1f, CurrentMana: %0.1f\n", manaSpendingRate, projectedManaCost, s.Stats[stats.Mana])
	}

	// If we have enough mana to burn, use the surplus agent.
	if projectedManaCost < s.Stats[stats.Mana] {
		return agent.surplusAgent.ChooseAction(s, party, sim)
	} else {
		return agent.baseAgent.ChooseAction(s, party, sim)
	}
}
func (agent *AdaptiveAgent) OnActionAccepted(s *Shaman, sim *core.Simulation, action core.AgentAction) {
	if !agent.wentOOM && action.Cast != nil && action.Cast.ManaCost > s.Stats[stats.Mana] {
		agent.timesOOM++
		agent.wentOOM = true
	}
	agent.takeSnapshot(sim, s)
	agent.baseAgent.OnActionAccepted(s, sim, action)
	agent.surplusAgent.OnActionAccepted(s, sim, action)
}

func (agent *AdaptiveAgent) Reset(sim *core.Simulation) {
	if agent.timesOOM == 5 {
		agent.baseAgent = NewLBOnlyAgent()
		agent.surplusAgent = NewCLOnClearcastAgent()
	}
	agent.wentOOM = false
	agent.manaSnapshots = [manaSnapshotsBufferSize]ManaSnapshot{}
	agent.firstSnapshotIndex = 0
	agent.numSnapshots = 0
	agent.baseAgent.Reset(sim)
	agent.surplusAgent.Reset(sim)
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
func ChainCast(sim *core.Simulation, p core.PlayerAgent, cast *core.Cast) {
	core.DirectCast(sim, p, cast) // Start with a normal direct cast to start.

	// Now chain
	dmgCoeff := 1.0
	if cast.Tag == CastTagLightningOverload {
		dmgCoeff = 0.5
	}
	for i := 1; i < sim.Options.Encounter.NumTargets; i++ {
		if p.HasAura(core.MagicIDTidefury) {
			dmgCoeff *= 0.83
		} else {
			dmgCoeff *= 0.7
		}
		clone := &core.Cast{
			Tag:       cast.Tag, // pass along lightning overload
			Spell:     cast.Spell,
			Crit:      cast.Crit,
			CritBonus: cast.CritBonus,
			Effect:    func(sim *core.Simulation, p core.PlayerAgent, c *core.Cast) { cast.DidDmg *= dmgCoeff },
			DoItNow:   core.DirectCast,
		}
		clone.DoItNow(sim, p, clone)
	}
}
