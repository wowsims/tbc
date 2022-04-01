package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	// Proc effects. Keep these in order by item ID.
	core.AddItemEffect(23207, ApplyMarkOfTheChampionCaster)
	core.AddItemEffect(27683, ApplyQuagmirransEye)
	core.AddItemEffect(28418, ApplyShiffarsNexusHorn)
	core.AddItemEffect(28789, ApplyEyeOfMagtheridon)
	core.AddItemEffect(30626, ApplySextantOfUnstableCurrents)
	core.AddItemEffect(31856, ApplyDarkmoonCardCrusade)

	// Activatable effects. Keep these in order by item ID.
	var RestrainedEssenceOfSapphironAuraID = core.NewAuraID()
	var RestrainedEssenceOfSapphironCooldownID = core.NewCooldownID()
	core.AddItemEffect(23046, core.MakeTemporaryStatsOnUseCDRegistration(
		RestrainedEssenceOfSapphironAuraID,
		stats.Stats{stats.SpellPower: 130},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 23046},
			CooldownID:       RestrainedEssenceOfSapphironCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var LivingRubySerpentAuraID = core.NewAuraID()
	var LivingRubySerpentCooldownID = core.NewCooldownID()
	core.AddItemEffect(24126, core.MakeTemporaryStatsOnUseCDRegistration(
		LivingRubySerpentAuraID,
		stats.Stats{stats.SpellPower: 150},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 24126},
			CooldownID:       LivingRubySerpentCooldownID,
			Cooldown:         time.Minute * 5,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var ScryersBloodgemAuraID = core.NewAuraID()
	var ScryersBloodgemCooldownID = core.NewCooldownID()
	core.AddItemEffect(29132, core.MakeTemporaryStatsOnUseCDRegistration(
		ScryersBloodgemAuraID,
		stats.Stats{stats.SpellPower: 150},
		time.Second*15,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 29132},
			CooldownID:       ScryersBloodgemCooldownID,
			Cooldown:         time.Second * 90,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var XirisGiftAuraID = core.NewAuraID()
	var XirisGiftCooldownID = core.NewCooldownID()
	core.AddItemEffect(29179, core.MakeTemporaryStatsOnUseCDRegistration(
		XirisGiftAuraID,
		stats.Stats{stats.SpellPower: 150},
		time.Second*15,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 29179},
			CooldownID:       XirisGiftCooldownID,
			Cooldown:         time.Second * 90,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var IconOfTheSilverCrescentAuraID = core.NewAuraID()
	var IconOfTheSilverCrescentCooldownID = core.NewCooldownID()
	core.AddItemEffect(29370, core.MakeTemporaryStatsOnUseCDRegistration(
		IconOfTheSilverCrescentAuraID,
		stats.Stats{stats.SpellPower: 155},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 29370},
			CooldownID:       IconOfTheSilverCrescentCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var EssenceOfTheMartyrAuraID = core.NewAuraID()
	var EssenceOfTheMartyrCooldownID = core.NewCooldownID()
	core.AddItemEffect(29376, core.MakeTemporaryStatsOnUseCDRegistration(
		EssenceOfTheMartyrAuraID,
		stats.Stats{stats.SpellPower: 99},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 29376},
			CooldownID:       EssenceOfTheMartyrCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.DefensiveTrinketSharedCooldownID,
		},
	))

	var SkullOfGuldanAuraID = core.NewAuraID()
	var SkullOfGuldanCooldownID = core.NewCooldownID()
	core.AddItemEffect(32483, core.MakeTemporaryStatsOnUseCDRegistration(
		SkullOfGuldanAuraID,
		stats.Stats{stats.SpellHaste: 175},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 32483},
			CooldownID:       SkullOfGuldanCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var HexShrunkenHeadAuraID = core.NewAuraID()
	var HexShrunkenHeadCooldownID = core.NewCooldownID()
	core.AddItemEffect(33829, core.MakeTemporaryStatsOnUseCDRegistration(
		HexShrunkenHeadAuraID,
		stats.Stats{stats.SpellPower: 211},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 33829},
			CooldownID:       HexShrunkenHeadCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var ShiftingNaaruSliverAuraID = core.NewAuraID()
	var ShiftingNaaruSliverCooldownID = core.NewCooldownID()
	core.AddItemEffect(34429, core.MakeTemporaryStatsOnUseCDRegistration(
		ShiftingNaaruSliverAuraID,
		stats.Stats{stats.SpellPower: 320},
		time.Second*15,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 34429},
			CooldownID:       ShiftingNaaruSliverCooldownID,
			Cooldown:         time.Second * 90,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var DarkIronSmokingPipeAuraID = core.NewAuraID()
	var DarkIronSmokingPipeCooldownID = core.NewCooldownID()
	core.AddItemEffect(38290, core.MakeTemporaryStatsOnUseCDRegistration(
		DarkIronSmokingPipeAuraID,
		stats.Stats{stats.SpellPower: 155},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 38290},
			CooldownID:       DarkIronSmokingPipeCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	// Even though these item effects are handled elsewhere, add them so they are
	// detected for automatic testing.
	core.AddItemEffect(core.AlchStoneItemID, func(core.Agent) {})
}

var MarkOfTheChampionCasterAuraID = core.NewAuraID()

func ApplyMarkOfTheChampionCaster(agent core.Agent) {
	character := agent.GetCharacter()
	character.RegisterResetEffect(func(sim *core.Simulation) {
		if sim.GetPrimaryTarget().MobType == proto.MobType_MobTypeDemon || sim.GetPrimaryTarget().MobType == proto.MobType_MobTypeUndead {
			character.PseudoStats.MobTypeSpellPower += 85
		}
	})
}

var QuagmirransEyeAuraID = core.NewAuraID()
var FungalFrenzyAuraID = core.NewAuraID()

func ApplyQuagmirransEye(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 320.0
		const dur = time.Second * 45
		icd := core.NewICD()

		applyStatAura := character.NewTemporaryStatsAuraApplier(FungalFrenzyAuraID, core.ActionID{ItemID: 27683}, stats.Stats{stats.SpellHaste: hasteBonus}, time.Second*6)
		return core.Aura{
			ID: QuagmirransEyeAuraID,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				if icd.IsOnCD(sim) || sim.RandomFloat("Quagmirran's Eye") > 0.1 {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + dur)
				applyStatAura(sim)
			},
		}
	})
}

var ShiffarsNexusHornAuraID = core.NewAuraID()
var CallOfTheNexusAuraID = core.NewAuraID()

func ApplyShiffarsNexusHorn(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		icd := core.NewICD()
		const spellBonus = 225.0
		const dur = time.Second * 45
		applyStatAura := character.NewTemporaryStatsAuraApplier(CallOfTheNexusAuraID, core.ActionID{ItemID: 28418}, stats.Stats{stats.SpellPower: spellBonus}, time.Second*10)

		return core.Aura{
			ID: ShiffarsNexusHornAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}
				if icd.IsOnCD(sim) || !spellEffect.Outcome.Matches(core.OutcomeCrit) || spellEffect.IsPhantom {
					return
				}
				if sim.RandomFloat("Shiffar's Nexus-Horn") > 0.2 {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + dur)
				applyStatAura(sim)
			},
		}
	})
}

var EyeOfMagtheridonAuraID = core.NewAuraID()
var RecurringPowerAuraID = core.NewAuraID()

func ApplyEyeOfMagtheridon(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const spellBonus = 170.0
		const dur = time.Second * 10
		applyStatAura := character.NewTemporaryStatsAuraApplier(RecurringPowerAuraID, core.ActionID{ItemID: 28789}, stats.Stats{stats.SpellPower: spellBonus}, dur)

		return core.Aura{
			ID: EyeOfMagtheridonAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}
				if !spellEffect.Outcome.Matches(core.OutcomeMiss) {
					return
				}
				applyStatAura(sim)
			},
		}
	})
}

var SextantOfUnstableCurrentsAuraID = core.NewAuraID()
var UnstableCurrentsAuraID = core.NewAuraID()

func ApplySextantOfUnstableCurrents(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		icd := core.NewICD()
		const spellBonus = 190.0
		const dur = time.Second * 15
		const icdDur = time.Second * 45
		applyStatAura := character.NewTemporaryStatsAuraApplier(UnstableCurrentsAuraID, core.ActionID{ItemID: 30626}, stats.Stats{stats.SpellPower: spellBonus}, dur)

		return core.Aura{
			ID: SextantOfUnstableCurrentsAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}
				if !spellEffect.Outcome.Matches(core.OutcomeCrit) || icd.IsOnCD(sim) || spellEffect.IsPhantom {
					return
				}
				if sim.RandomFloat("Sextant of Unstable Currents") > 0.2 {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				applyStatAura(sim)
			},
		}
	})
}

var DarkmoonCardCrusadeAuraID = core.NewAuraID()
var AuraOfTheCrusadeMeleeAuraID = core.NewAuraID()
var AuraOfTheCrusadeSpellAuraID = core.NewAuraID()

func ApplyDarkmoonCardCrusade(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const meleeBonus = 6.0
		const spellBonus = 8.0
		meleeStacks := 0
		spellStacks := 0

		apAura := core.Aura{
			ID:       AuraOfTheCrusadeMeleeAuraID,
			ActionID: core.ActionID{ItemID: 31856, Tag: 1},
			Duration: time.Second * 10,
			OnExpire: func(sim *core.Simulation) {
				character.AddStat(stats.AttackPower, -meleeBonus*float64(meleeStacks))
				character.AddStat(stats.RangedAttackPower, -meleeBonus*float64(meleeStacks))
				meleeStacks = 0
			},
		}

		spAura := core.Aura{
			ID:       AuraOfTheCrusadeSpellAuraID,
			ActionID: core.ActionID{ItemID: 31856, Tag: 2},
			Duration: time.Second * 10,
			OnExpire: func(sim *core.Simulation) {
				character.AddStat(stats.SpellPower, -spellBonus*float64(spellStacks))
				spellStacks = 0
			},
		}

		return core.Aura{
			ID: DarkmoonCardCrusadeAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					if spellEffect.IsPhantom {
						return
					}

					if meleeStacks < 20 {
						meleeStacks++
						character.AddStat(stats.AttackPower, meleeBonus)
						character.AddStat(stats.RangedAttackPower, meleeBonus)
					}

					// Removal aura will refresh with new total spellpower based on stacks.
					//  This will remove the old stack removal buff.
					character.ReplaceAura(sim, apAura)
				} else {
					if !spellEffect.Landed() {
						return
					}
					if spellStacks < 10 {
						spellStacks++
						character.AddStat(stats.SpellPower, spellBonus)
					}

					// Removal aura will refresh with new total spellpower based on stacks.
					//  This will remove the old stack removal buff.
					character.ReplaceAura(sim, spAura)
				}
			},
		}
	})
}
