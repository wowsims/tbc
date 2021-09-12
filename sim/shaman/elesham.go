package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var ElementalItems = []int32{
	27471, 24266, 28278, 31330, 28415, 28758, 28349, 29504, 31107, 28193, 28169, 27488, 30297, 27993, 30946, 28245, 28134, 29333, 31692, 28254, 27758, 31693, 27464, 31338, 27473, 32078, 27796, 30925, 31797, 27778, 27802, 27994, 25777, 28269, 29813, 27981, 32541, 24252, 31140, 28379, 27469, 31340, 29129, 28231, 29341, 28191, 31297, 28342, 28232, 28229, 27824, 28391, 28638, 27522, 24250, 27462, 29240, 27746, 29243, 29955, 27465, 27793, 31149, 27470, 31280, 30924, 29317, 27493, 27508, 24452, 27537, 29784, 27743, 29244, 31461, 29257, 29241, 27783, 27795, 31513, 24262, 30541, 29141, 29142, 30531, 30709, 27492, 29343, 27472, 30532, 28185, 27838, 27907, 27909, 28266, 27948, 29314, 28406, 28640, 27914, 28179, 29245, 27821, 27845, 29808, 29242, 29258, 29313, 30519, 28227, 29126, 31922, 28394, 29320, 27784, 30366, 29172, 29352, 28555, 31339, 31921, 28248, 23199, 27543, 27868, 27741, 27937, 28412, 28260, 31287, 28187, 29330, 27910, 30984, 27534, 29355, 29130, 28341, 31308, 28188, 30011, 27842, 28346, 24557, 29389, 28744, 28586, 29035, 30171, 29986, 32476, 31014, 32525, 28530, 29368, 30008, 28762, 30015, 32349, 28726, 30024, 29037, 30079, 32338, 30173, 32587, 31023, 30884, 28766, 28570, 29369, 29992, 28797, 30735, 32331, 29033, 29519, 30056, 32327, 30913, 30169, 30107, 32592, 31017, 28515, 32351, 32270, 29521, 32259, 29918, 30870, 29034, 28507, 30170, 29987, 30725, 28780, 31008, 28565, 28639, 28654, 30044, 29520, 28799, 30064, 24256, 30914, 32256, 30038, 30888, 32276, 29036, 28594, 29972, 30172, 32367, 31020, 30734, 30916, 28670, 28585, 28810, 30894, 30037, 28517, 30043, 30067, 32242, 32352, 32239, 28793, 28510, 29922, 29367, 29287, 29286, 29285, 28753, 29305, 30109, 30667, 32247, 32527, 30832, 23554, 28770, 30723, 34009, 32237, 28633, 29988, 32374, 28734, 28611, 34011, 28781, 29268, 28603, 32361, 29273, 30049, 30909, 30872, 28297, 27683, 29370, 19344, 19379, 23046, 23207, 29132, 24126, 29179, 28418, 31856, 28785, 28789, 30626, 34429, 32483, 33829, 29376, 38290, 30663, 35749, 24116, 24121, 31075, 32664, 29522, 29524, 29523, 27510, 22730, 23070, 21709, 23031, 23025, 23057, 21608, 23664, 23665, 23050, 30682, 30677, 30686, 28583, 32586, 23049, 25778, 28174, 31283, 30004, 31290, 34336, 34179, 34350, 34542, 34186, 34566, 34437, 34230, 34362, 34204, 34332, 34242, 34396, 34390, 33970, 33965, 33588, 33537, 33534, 34359, 32330, 33506, 32086, 28602, 32963, 32524, 33357, 33533, 33354, 33283, 33466, 33591, 32817, 32792, 32328, 33281, 33334, 34344,
}

var ElementalGems = []int32{
	34220, 25897, 32641, 35503, 28557, 25893, 25901, 23096, 24030, 32196, 28118, 33133, 23121, 24037, 32202, 23113, 24047, 32204, 23114, 24050, 32207, 30551, 23101, 24059, 32218, 35760, 30588, 28123, 31866, 31867, 32221, 30564, 30560, 24065, 35759, 24056, 30555, 32215, 31116, 30600, 30605,
}

var ElementalEnchants = []int32{
	29191, 28909, 28886, 24421, 20076, 23545, 27960, 27917, 22534, 33997, 28272, 24274, 24273, 27975, 22555, 35445, 27945,
}

func loDmgMod(sim *core.Simulation, p *core.Player, c *core.Cast) {
	c.DidDmg /= 2
}

func AuraLightningOverload(lvl int) core.Aura {
	chance := 0.04 * float64(lvl)
	return core.Aura{
		ID:      core.MagicIDLOTalent,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, p *core.Player, c *core.Cast) {
			if c.Spell.ID != core.MagicIDLB12 && c.Spell.ID != core.MagicIDCL6 {
				return
			}
			if c.IsLO {
				return // can't proc LO on LO
			}
			actualChance := chance
			if c.Spell.ID == core.MagicIDCL6 {
				actualChance /= 3 // 33% chance of regular for CL LO
			}
			if sim.Rando.Float64() < actualChance {
				if sim.Debug != nil {
					sim.Debug(" +Lightning Overload\n")
				}
				clone := sim.NewCast()
				// Don't set IsClBounce even if this is a bounce, so that the clone does a normal CL and bounces
				clone.IsLO = true
				clone.Spell = c.Spell

				// Clone dmg/hit/crit chance?
				clone.Hit = c.Hit
				clone.Crit = c.Crit
				clone.Dmg = c.Dmg

				clone.CritBonus = c.CritBonus
				clone.Effect = loDmgMod
				c.DoItNow(sim, p, nil, clone)
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
		OnCast: func(sim *core.Simulation, p *core.Player, c *core.Cast) {
			c.ManaCost = 0
			c.Crit += 1.01
		},
		OnCastComplete: func(sim *core.Simulation, p *core.Player, c *core.Cast) {
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

func NewLBOnlyAgent(sim *core.Simulation) *LBOnlyAgent {
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
		return NewCastAction(sim, s, agent.lb)
	} else {
		return NewCastAction(sim, s, agent.cl)
	}
}

func (agent *CLOnCDAgent) OnActionAccepted(p *Shaman, sim *core.Simulation, action core.AgentAction) {
}
func (agent *CLOnCDAgent) Reset(sim *core.Simulation) {}

func NewCLOnCDAgent(sim *core.Simulation) *CLOnCDAgent {
	return &CLOnCDAgent{
		lb: core.Spells[core.MagicIDLB12],
		cl: core.Spells[core.MagicIDCL6],
	}
}

// ################################################################
//                          FIXED ROTATION
// ################################################################
// type FixedRotationAgent struct {
// 	numLBsPerCL       int
// 	numLBsSinceLastCL int
// 	lb                *core.Spell
// 	cl                *core.Spell
// }

// // Returns if any temporary haste buff is currently active.
// // TODO: Figure out a way to make this automatic
// func (agent *FixedRotationAgent) temporaryHasteActive(sim *core.Simulation) bool {
// 	return sim.hasAura(MagicIDBloodlust) ||
// 		sim.hasAura(MagicIDDrums) ||
// 		sim.hasAura(MagicIDTrollBerserking) ||
// 		sim.hasAura(MagicIDSkullGuldan) ||
// 		sim.hasAura(MagicIDFungalFrenzy)
// }

// func (agent *FixedRotationAgent) ChooseAction(sim *core.Simulation) AgentAction {
// 	if agent.numLBsSinceLastCL < agent.numLBsPerCL {
// 		return NewCastAction(sim, agent.lb)
// 	}

// 	if !sim.isOnCD(MagicIDCL6) {
// 		return NewCastAction(sim, agent.cl)
// 	}

// 	// If we have a temporary haste effect (like bloodlust or quags eye) then
// 	// we should add LB casts instead of waiting
// 	if agent.temporaryHasteActive(sim) {
// 		return NewCastAction(sim, agent.lb)
// 	}

// 	return NewWaitAction(sim.getRemainingCD(MagicIDCL6))
// }

// func (agent *FixedRotationAgent) OnActionAccepted(sim *core.Simulation, action AgentAction) {
// 	if action.Cast == nil {
// 		return
// 	}

// 	if action.Cast.Spell.ID == MagicIDLB12 {
// 		agent.numLBsSinceLastCL++
// 	} else if action.Cast.Spell.ID == MagicIDCL6 {
// 		agent.numLBsSinceLastCL = 0
// 	}
// }

// func (agent *FixedRotationAgent) Reset(sim *core.Simulation) {
// 	agent.numLBsSinceLastCL = agent.numLBsPerCL
// }

// func NewFixedRotationAgent(sim *core.Simulation, numLBsPerCL int) *FixedRotationAgent {
// 	return &FixedRotationAgent{
// 		numLBsPerCL:       numLBsPerCL,
// 		numLBsSinceLastCL: numLBsPerCL, // This lets us cast CL first
// 		lb:                spellmap[MagicIDLB12],
// 		cl:                spellmap[MagicIDCL6],
// 	}
// }

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
		return NewCastAction(sim, s, agent.lb)
	}

	return NewCastAction(sim, s, agent.cl)
}

func (agent *CLOnClearcastAgent) OnActionAccepted(p *Shaman, sim *core.Simulation, action core.AgentAction) {
	agent.prevPrevCastProccedCC = p.Auras[core.MagicIDEleFocus].Stacks == 2
}

func (agent *CLOnClearcastAgent) Reset(sim *core.Simulation) {
	agent.prevPrevCastProccedCC = true // Lets us cast CL first
}

func NewCLOnClearcastAgent(sim *core.Simulation) *CLOnClearcastAgent {
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
		sim.Debug("[AI] CL Ready: Mana/s: %0.1f, Est Mana Cost: %0.1f, CurrentMana: %0.1f\n", manaSpendingRate, projectedManaCost, s.Stats[core.StatMana])
	}

	// Before casting, activate shaman powers!
	TryActivateBloodlust(sim, party, s.Player)
	if s.Talents.ElementalMastery {
		TryActivateEleMastery(sim, s.Player)
	}

	// If we have enough mana to burn, use the surplus agent.
	if projectedManaCost < s.Stats[core.StatMana] {
		return agent.surplusAgent.ChooseAction(s, party, sim)
	} else {
		return agent.baseAgent.ChooseAction(s, party, sim)
	}
}
func (agent *AdaptiveAgent) OnActionAccepted(p *Shaman, sim *core.Simulation, action core.AgentAction) {
	agent.takeSnapshot(sim, p)
	agent.baseAgent.OnActionAccepted(p, sim, action)
	agent.surplusAgent.OnActionAccepted(p, sim, action)
}

func (agent *AdaptiveAgent) Reset(sim *core.Simulation) {
	agent.manaSnapshots = [manaSnapshotsBufferSize]ManaSnapshot{}
	agent.firstSnapshotIndex = 0
	agent.numSnapshots = 0
	agent.baseAgent.Reset(sim)
	agent.surplusAgent.Reset(sim)
}

func NewAdaptiveAgent(sim *core.Simulation) *AdaptiveAgent {
	agent := &AdaptiveAgent{}

	// TODO: Can we just start with more aggressive agent and drop to less aggressive if we go OOM 5 times?
	//   not as deterministic... but probably averages out the same?
	// Otherwise we need to figure out how to do this after all other agents are setup (in the eventual 'raid' sim setup)

	// clearcastSimRequest := core.SimRequest{
	// 	Options:    sim.Options,
	// 	Gear:       sim.EquipSpec,
	// 	Iterations: 100,
	// }
	// clearcastSimRequest.Options.AgentType = AGENT_TYPE_CL_ON_CLEARCAST
	// clearcastResult := core.RunSimulation(clearcastSimRequest)

	// if clearcastResult.NumOom >= 5 {
	// agent.baseAgent = NewLBOnlyAgent(sim)           //NewAgent(sim, AGENT_TYPE_FIXED_LB_ONLY)
	// agent.surplusAgent = NewCLOnClearcastAgent(sim) //NewAgent(sim, AGENT_TYPE_CL_ON_CLEARCAST)
	// } else {
	agent.baseAgent = NewCLOnClearcastAgent(sim)
	agent.surplusAgent = NewCLOnCDAgent(sim)
	// }

	return agent
}

// ChainCast is how to cast chain lightning.
func ChainCast(sim *core.Simulation, p *core.Player, a core.Agent, cast *core.Cast) {
	core.DirectCast(sim, p, a, cast) // Start with a normal direct cast to start.

	// Now chain
	dmgCoeff := 1.0
	if cast.IsLO {
		dmgCoeff = 0.5
	}
	for i := 1; i < sim.Options.Encounter.NumTargets; i++ {
		if p.HasAura(core.MagicIDTidefury) {
			dmgCoeff *= 0.83
		} else {
			dmgCoeff *= 0.7
		}
		clone := &core.Cast{
			IsLO:       cast.IsLO,
			IsClBounce: true,
			Spell:      cast.Spell,
			Crit:       cast.Crit,
			CritBonus:  cast.CritBonus,
			Effect:     func(sim *core.Simulation, p *core.Player, c *core.Cast) { cast.DidDmg *= dmgCoeff },
			DoItNow:    core.DirectCast,
		}
		clone.DoItNow(sim, p, a, clone)
	}
}
