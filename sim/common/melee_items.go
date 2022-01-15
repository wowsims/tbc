package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	// Proc effects. Keep these in order by item ID.
	core.AddItemEffect(23541, ApplyKhoriumChampion)
	core.AddItemEffect(27901, ApplyBlackoutTruncheon)
	core.AddItemEffect(28429, ApplyLionheartChampion)
	core.AddItemEffect(28430, ApplyLionheartExecutioner)
	core.AddItemEffect(28437, ApplyDrakefistHammer)
	core.AddItemEffect(28438, ApplyDragonmaw)
	core.AddItemEffect(28439, ApplyDragonstrike)
	core.AddItemEffect(28774, ApplyGlaiveOfThePit)
	core.AddItemEffect(29348, ApplyTheBladefist)
	core.AddItemEffect(29996, ApplyRodOfTheSunKing)
	core.AddItemEffect(30090, ApplyWorldBreaker)
	core.AddItemEffect(30316, ApplyDevastation)
	core.AddItemEffect(31318, ApplySingingCrystalAxe)
	core.AddItemEffect(31331, ApplyTheNightBlade)
	// decapitator
	// despair
}

var KhoriumChampionAuraID = core.NewAuraID()
var KhoriumChampionProcAuraID = core.NewAuraID()

func ApplyKhoriumChampion(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const strBonus = 120.0
		const dur = time.Second * 30
		const procChance = 0.5 * 3.3 / 60.0

		return core.Aura{
			ID: KhoriumChampionAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					return
				}
				if sim.RandomFloat("KhoriumChampion") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, KhoriumChampionProcAuraID, core.ActionID{ItemID: 23541}, stats.Strength, strBonus, dur)
			},
		}
	})
}

var BlackoutTruncheonAuraID = core.NewAuraID()
var BlackoutTruncheonProcAuraID = core.NewAuraID()

func ApplyBlackoutTruncheon(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 132.0
		const dur = time.Second * 10
		const procChance = 1.5 * 0.8 / 60.0

		return core.Aura{
			ID: BlackoutTruncheonAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					return
				}
				if sim.RandomFloat("BlackoutTruncheon") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, BlackoutTruncheonProcAuraID, core.ActionID{ItemID: 27901}, stats.MeleeHaste, hasteBonus, dur)
			},
		}
	})
}

var LionheartChampionAuraID = core.NewAuraID()
var LionheartChampionProcAuraID = core.NewAuraID()

func ApplyLionheartChampion(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const strBonus = 100.0
		const dur = time.Second * 10
		const procChance = 3.6 / 60.0

		return core.Aura{
			ID: LionheartChampionAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					return
				}
				if sim.RandomFloat("LionheartChampion") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, LionheartChampionProcAuraID, core.ActionID{ItemID: 28429}, stats.Strength, strBonus, dur)
			},
		}
	})
}

var LionheartExecutionerAuraID = core.NewAuraID()
var LionheartExecutionerProcAuraID = core.NewAuraID()

func ApplyLionheartExecutioner(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const strBonus = 100.0
		const dur = time.Second * 10
		const procChance = 3.6 / 60.0

		return core.Aura{
			ID: LionheartExecutionerAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					return
				}
				if sim.RandomFloat("LionheartExecutioner") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, LionheartExecutionerProcAuraID, core.ActionID{ItemID: 28430}, stats.Strength, strBonus, dur)
			},
		}
	})
}

var DrakefistHammerAuraID = core.NewAuraID()
var DrakefistHammerProcAuraID = core.NewAuraID()

func ApplyDrakefistHammer(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 212.0
		const dur = time.Second * 10
		const procChance = 2.7 / 60.0

		return core.Aura{
			ID: DrakefistHammerAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					return
				}
				if sim.RandomFloat("DrakefistHammer") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, DrakefistHammerProcAuraID, core.ActionID{ItemID: 28437}, stats.MeleeHaste, hasteBonus, dur)
			},
		}
	})
}

var DragonmawAuraID = core.NewAuraID()
var DragonmawProcAuraID = core.NewAuraID()

func ApplyDragonmaw(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 212.0
		const dur = time.Second * 10
		const procChance = 2.7 / 60.0

		return core.Aura{
			ID: DragonmawAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					return
				}
				if sim.RandomFloat("Dragonmaw") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, DragonmawProcAuraID, core.ActionID{ItemID: 28438}, stats.MeleeHaste, hasteBonus, dur)
			},
		}
	})
}

var DragonstrikeAuraID = core.NewAuraID()
var DragonstrikeProcAuraID = core.NewAuraID()

func ApplyDragonstrike(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 212.0
		const dur = time.Second * 10
		const procChance = 2.7 / 60.0

		return core.Aura{
			ID: DragonstrikeAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					return
				}
				if sim.RandomFloat("Dragonstrike") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, DragonstrikeProcAuraID, core.ActionID{ItemID: 28439}, stats.MeleeHaste, hasteBonus, dur)
			},
		}
	})
}

var GlaiveOfThePitAuraID = core.NewAuraID()
var GlaiveOfThePitProcAuraID = core.NewAuraID()

func ApplyGlaiveOfThePit(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 212.0
		const procChance = 3.7 / 60.0

		spellObj := core.SimpleSpell{}
		castTemplate := core.NewSimpleSpellTemplate(core.SimpleSpell{
			SpellCast: core.SpellCast{
				Cast: core.Cast{
					ActionID:        core.ActionID{SpellID: 34696},
					Character:       character,
					IgnoreCooldowns: true,
					IgnoreManaCost:  true,
					IsPhantom:       true,
					SpellSchool:     stats.ShadowSpellPower,
					CritMultiplier:  1.5,
				},
			},
			Effect: core.SpellHitEffect{
				SpellEffect: core.SpellEffect{
					DamageMultiplier:       1,
					StaticDamageMultiplier: 1,
				},
				DirectInput: core.DirectDamageInput{
					MinBaseDamage: 285,
					MaxBaseDamage: 315,
				},
			},
		})

		return core.Aura{
			ID: GlaiveOfThePitAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					return
				}
				if sim.RandomFloat("GlaiveOfThePit") > procChance {
					return
				}

				castAction := &spellObj
				castTemplate.Apply(castAction)
				castAction.Effect.Target = hitEffect.Target
				castAction.Init(sim)
				castAction.Cast(sim)
			},
		}
	})
}

var TheBladefistAuraID = core.NewAuraID()
var TheBladefistProcAuraID = core.NewAuraID()

func ApplyTheBladefist(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 180.0
		const dur = time.Second * 10
		const procChance = 2.7 / 60.0

		return core.Aura{
			ID: TheBladefistAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					return
				}
				if sim.RandomFloat("The Bladefist") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, TheBladefistProcAuraID, core.ActionID{ItemID: 29348}, stats.MeleeHaste, hasteBonus, dur)
			},
		}
	})
}

var RodOfTheSunKingAuraID = core.NewAuraID()

func ApplyRodOfTheSunKing(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const procChance = 2.7 / 60.0

		return core.Aura{
			ID: RodOfTheSunKingAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					return
				}
				if sim.RandomFloat("Rod of the Sun King") > procChance {
					return
				}

				// TODO: Add 5 rage or 10 energy.
			},
		}
	})
}

var WorldBreakerAuraID = core.NewAuraID()
var WorldBreakerProcAuraID = core.NewAuraID()

func ApplyWorldBreaker(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const critBonus = 900.0
		const dur = time.Second * 4
		const procChance = 3.7 / 60.0

		return core.Aura{
			ID: WorldBreakerAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					if character.HasAura(WorldBreakerProcAuraID) {
						character.RemoveAura(sim, WorldBreakerProcAuraID)
					}
					return
				}
				if sim.RandomFloat("World Breaker") > procChance {
					if character.HasAura(WorldBreakerProcAuraID) {
						character.RemoveAura(sim, WorldBreakerProcAuraID)
					}
					return
				}

				character.AddAuraWithTemporaryStats(sim, WorldBreakerProcAuraID, core.ActionID{ItemID: 30090}, stats.MeleeCrit, critBonus, dur)
			},
		}
	})
}

var DevastationAuraID = core.NewAuraID()
var DevastationProcAuraID = core.NewAuraID()

func ApplyDevastation(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const bonus = 1.2
		const inverseBonus = 1 / 1.2
		const dur = time.Second * 30
		const procChance = 0.5

		return core.Aura{
			ID: DevastationAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					return
				}
				if sim.RandomFloat("Devastation") > procChance {
					return
				}

				character.MultiplyMeleeSpeed(sim, bonus)
				character.AddAura(sim, core.Aura{
					ID:       DevastationProcAuraID,
					ActionID: core.ActionID{ItemID: 30316},
					Expires:  sim.CurrentTime + dur,
					OnExpire: func(sim *core.Simulation) {
						character.MultiplyMeleeSpeed(sim, inverseBonus)
					},
				})
			},
		}
	})
}

var SingingCrystalAxeAuraID = core.NewAuraID()
var SingingCrystalAxeProcAuraID = core.NewAuraID()

func ApplySingingCrystalAxe(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 400.0
		const dur = time.Second * 10
		const procChance = 3.5 / 60.0

		return core.Aura{
			ID: SingingCrystalAxeAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					return
				}
				if sim.RandomFloat("SingingCrystalAxe") > procChance {
					return
				}

				character.AddAuraWithTemporaryStats(sim, SingingCrystalAxeProcAuraID, core.ActionID{ItemID: 31318}, stats.MeleeHaste, hasteBonus, dur)
			},
		}
	})
}

var TheNightBladeAuraID = core.NewAuraID()
var TheNightBladeProcAuraID = core.NewAuraID()

func ApplyTheNightBlade(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const arPenBonus = 435.0
		const dur = time.Second * 10
		const procChance = 2 * 1.8 / 60.0

		return core.Aura{
			ID: TheNightBladeAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() || !hitEffect.IsWeaponHit() {
					return
				}
				if sim.RandomFloat("The Night Blade") > procChance {
					return
				}

				stacks := character.NumStacks(TheNightBladeProcAuraID) + 1
				newBonus := arPenBonus * float64(stacks)
				character.AddAura(sim, core.Aura{
					ID:       TheNightBladeProcAuraID,
					ActionID: core.ActionID{ItemID: 31331},
					Expires:  sim.CurrentTime + dur,
					Stacks:   stacks,
					OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
						hitEffect.BonusArmorPenetration += newBonus
					},
				})
			},
		}
	})
}
