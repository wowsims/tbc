package retribution

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
	"github.com/wowsims/tbc/sim/paladin"
)

func RegisterRetributionPaladin() {
	core.RegisterAgentFactory(
		proto.Player_RetributionPaladin{},
		proto.Spec_SpecRetributionPaladin,
		func(character core.Character, options proto.Player) core.Agent {
			return NewRetributionPaladin(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_RetributionPaladin) // I don't really understand this line
			if !ok {
				panic("Invalid spec value for Retribution Paladin!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewRetributionPaladin(character core.Character, options proto.Player) *RetributionPaladin {
	retOptions := options.GetRetributionPaladin()

	ret := &RetributionPaladin{
		Paladin:     paladin.NewPaladin(character, *retOptions.Talents),
		Rotation:    *retOptions.Rotation,
		csDelay:     time.Duration(retOptions.Options.CrusaderStrikeDelayMs),
		hasteLeeway: time.Duration(retOptions.Options.HasteLeewayMs),
		judgement:   retOptions.Options.Judgement,
	}

	// Convert DTPS option to bonus MP5
	spAtt := retOptions.Options.DamageTakenPerSecond * 5.0 / 10.0
	ret.AddStat(stats.MP5, spAtt)

	ret.EnableAutoAttacks(ret, core.AutoAttackOptions{
		MainHand:       ret.WeaponFromMainHand(ret.DefaultMeleeCritMultiplier()),
		AutoSwingMelee: true,
	})

	ret.SetupSealOfCommand()

	return ret
}

type RetributionPaladin struct {
	*paladin.Paladin

	openerCompleted bool

	hasteLeeway time.Duration
	csDelay     time.Duration

	judgement proto.RetributionPaladin_Options_Judgement

	Rotation proto.RetributionPaladin_Rotation
}

func (ret *RetributionPaladin) GetPaladin() *paladin.Paladin {
	return ret.Paladin
}

func (ret *RetributionPaladin) Reset(sim *core.Simulation) {
	ret.Paladin.Reset(sim)

	switch ret.judgement {
	case proto.RetributionPaladin_Options_Wisdom:
		ret.UpdateSeal(sim, ret.SealOfWisdomAura)
	case proto.RetributionPaladin_Options_Crusader:
		ret.UpdateSeal(sim, ret.SealOfTheCrusaderAura)
	case proto.RetributionPaladin_Options_None:
		ret.UpdateSeal(sim, ret.SealOfCommandAura)
	}

	ret.AutoAttacks.CancelAutoSwing(sim)
	ret.openerCompleted = false
}

func (ret *RetributionPaladin) OnGCDReady(sim *core.Simulation) {
	ret.tryUseGCD(sim)
}

func (ret *RetributionPaladin) OnManaTick(sim *core.Simulation) {
	if ret.FinishedWaitingForManaAndGCDReady(sim) {
		ret.tryUseGCD(sim)
	}
}

func (ret *RetributionPaladin) tryUseGCD(sim *core.Simulation) {
	if !ret.openerCompleted {
		ret.openingRotation(sim)
		return
	}
	ret.ActRotation(sim)
}

func (ret *RetributionPaladin) openingRotation(sim *core.Simulation) {
	target := sim.GetPrimaryTarget()

	// Cast selected judgement to keep on the boss
	if !ret.IsOnCD(paladin.JudgementCD, sim.CurrentTime) &&
		ret.judgement != proto.RetributionPaladin_Options_None {
		var judge *core.SimpleSpell
		switch ret.judgement {
		case proto.RetributionPaladin_Options_Wisdom:
			judge = ret.NewJudgementOfWisdom(sim, target)
		case proto.RetributionPaladin_Options_Crusader:
			judge = ret.NewJudgementOfTheCrusader(sim, target)
		}
		if judge != nil {
			if success := judge.Cast(sim); !success {
				ret.WaitForMana(sim, judge.GetManaCost())
			}
		}
	}

	// Cast Seal of Command
	if !ret.HasAura(paladin.SealOfCommandAuraID) {
		soc := ret.NewSealOfCommand(sim)
		if success := soc.StartCast(sim); !success {
			ret.WaitForMana(sim, soc.GetManaCost())
		}
		return
	}

	// Cast Seal of Blood and enable attacks to twist
	if !ret.HasAura(paladin.SealOfBloodAuraID) {
		sob := ret.NewSealOfBlood(sim)
		if success := sob.StartCast(sim); !success {
			ret.WaitForMana(sim, sob.GetManaCost())
		}
		ret.AutoAttacks.EnableAutoSwing(sim)
		ret.openerCompleted = true
	}
}

// TO-DO maybe add low mana logic
func (ret *RetributionPaladin) ActRotation(sim *core.Simulation) {
	target := sim.GetPrimaryTarget()

	gcdCD := ret.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
	crusaderStrikeCD := ret.GetRemainingCD(paladin.CrusaderStrikeCD, sim.CurrentTime)
	judgmentCD := ret.GetRemainingCD(paladin.JudgementCD, sim.CurrentTime)

	sobActive := ret.RemainingAuraDuration(sim, paladin.SealOfBloodAuraID) > 0
	socActive := ret.RemainingAuraDuration(sim, paladin.SealOfCommandAuraID) > 0

	nextSwing := ret.AutoAttacks.NextAttackAt() - sim.CurrentTime - ret.hasteLeeway
	effWeaponSpeed := ret.AutoAttacks.MainhandSwingSpeed()
	spellGCD := ret.SpellGCD()

	twistWindow := ret.AutoAttacks.NextAttackAt() - 400*time.Millisecond
	canTwist := sim.CurrentTime >= twistWindow && sim.CurrentTime <= ret.AutoAttacks.NextAttackAt()

	// Check if we are on GCD
	if gcdCD > 0 {
		// Check if Judgement is off CD and Seal of Blood is active
		if judgmentCD == 0 && sobActive {
			// Check if Crusader Strike is off CD
			if crusaderStrikeCD == 0 {
				// Do nothing
				return
			} else {
				// Check if Crusader Strike will be ready this swing
				if crusaderStrikeCD < nextSwing {
					// Do nothing
					return
				} else {
					// Check if we can reseal after casting Crusader Strike
					if gcdCD < nextSwing {
						// Cast Judgement of Blood
						judge := ret.NewJudgementOfBlood(sim, target)
						if judge != nil {
							if success := judge.Cast(sim); !success {
								ret.WaitForMana(sim, judge.Cost.Value)
							}
						}
						return
					} else {
						// Do Nothing
						return
					}
				}
			}
		} else {
			// do nothing
			return
		}
	} else {
		// Check if Seal of Command is active
		if socActive {
			// Check if we can use fillers
			if effWeaponSpeed > spellGCD*2 &&
				sim.CurrentTime+spellGCD < nextSwing {
				ret.useFillers(sim)
			} else {
				// Check if we are in the twist window
				if canTwist {
					// Cast Seal of Blood
					sob := ret.NewSealOfBlood(sim)
					if success := sob.StartCast(sim); !success {
						ret.WaitForMana(sim, sob.GetManaCost())
					}
					return
				} else {
					// Wait until twist window
					ret.WaitUntil(sim, twistWindow)
				}
			}
		} else {
			// Check if Crusader Strike is ready
			if crusaderStrikeCD == 0 {
				// Check if Seal of Blood is active
				if sobActive {
					// Cast Crusader Strike
					cs := ret.NewCrusaderStrike(sim, target)
					if success := cs.Cast(sim); !success {
						ret.WaitForMana(sim, cs.Cost.Value)
					}
					return
				} else {
					// Check if we can Crusader Strike and reseal
					if sim.CurrentTime+crusaderStrikeCD > nextSwing {
						// Cast Crusader Strike
						cs := ret.NewCrusaderStrike(sim, target)
						if success := cs.Cast(sim); !success {
							ret.WaitForMana(sim, cs.Cost.Value)
						}
						return
					} else {
						// Cast Seal of Blood
						sob := ret.NewSealOfBlood(sim)
						if success := sob.StartCast(sim); !success {
							ret.WaitForMana(sim, sob.GetManaCost())
						}
					}
				}
			} else {

			}
		}
	}
}

func (ret *RetributionPaladin) useFillers(sim *core.Simulation) {
	return
}

func (ret *RetributionPaladin) _2007Rotation(sim *core.Simulation) {
	target := sim.GetPrimaryTarget()

	// judge blood whenever we can
	if !ret.IsOnCD(paladin.JudgementCD, sim.CurrentTime) {
		judge := ret.NewJudgementOfBlood(sim, target)
		if judge != nil {
			if success := judge.Cast(sim); !success {
				ret.WaitForMana(sim, judge.Cost.Value)
			}
		}
	}

	// roll seal of blood
	if !ret.HasAura(paladin.SealOfBloodAuraID) {
		sob := ret.NewSealOfBlood(sim)
		if success := sob.StartCast(sim); !success {
			ret.WaitForMana(sim, sob.GetManaCost())
		}
		return
	}

	// Crusader strike if we can
	if !ret.IsOnCD(paladin.CrusaderStrikeCD, sim.CurrentTime) {
		cs := ret.NewCrusaderStrike(sim, target)
		if success := cs.Cast(sim); !success {
			ret.WaitForMana(sim, cs.Cost.Value)
		}
		return
	}

	// Proceed until SoB expires, CrusaderStrike comes off GCD, or Judgement comes off GCD
	nextEventAt := ret.CDReadyAt(paladin.CrusaderStrikeCD)
	sobExpiration := sim.CurrentTime + ret.RemainingAuraDuration(sim, paladin.SealOfBloodAuraID)
	nextEventAt = core.MinDuration(nextEventAt, sobExpiration)
	// Waiting for judgement CD causes a bug that infinite loops for some reason
	// nextEventAt = core.MinDuration(nextEventAt, ret.CDReadyAt(paladin.JudgementCD))
	ret.WaitUntil(sim, nextEventAt)
}
