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
	var RestrainedEssenceOfSapphironCooldownID = core.NewCooldownID()
	core.AddItemEffect(23046, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		28779,
		"Restrained Essence of Sapphiron",
		stats.SpellPower,
		130,
		time.Second*20,
		core.MajorCooldown{
			CooldownID:       RestrainedEssenceOfSapphironCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var LivingRubySerpentCooldownID = core.NewCooldownID()
	core.AddItemEffect(24126, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		31040,
		"Figurine - Living Ruby Serpent",
		stats.SpellPower,
		150,
		time.Second*20,
		core.MajorCooldown{
			CooldownID:       LivingRubySerpentCooldownID,
			Cooldown:         time.Minute * 5,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var ScryersBloodgemCooldownID = core.NewCooldownID()
	core.AddItemEffect(29132, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		35337,
		"Scryer's Bloodgem",
		stats.SpellPower,
		150,
		time.Second*15,
		core.MajorCooldown{
			CooldownID:       ScryersBloodgemCooldownID,
			Cooldown:         time.Second * 90,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var XirisGiftCooldownID = core.NewCooldownID()
	core.AddItemEffect(29179, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		35337,
		"Xi'ri's Gift",
		stats.SpellPower,
		150,
		time.Second*15,
		core.MajorCooldown{
			CooldownID:       XirisGiftCooldownID,
			Cooldown:         time.Second * 90,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var IconOfTheSilverCrescentCooldownID = core.NewCooldownID()
	core.AddItemEffect(29370, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		35163,
		"Blessing of the Silver Crescent",
		stats.SpellPower,
		155,
		time.Second*20,
		core.MajorCooldown{
			CooldownID:       IconOfTheSilverCrescentCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var EssenceOfTheMartyrCooldownID = core.NewCooldownID()
	core.AddItemEffect(29376, core.MakeTemporaryStatsOnUseCDRegistration(
		core.DefensiveTrinketActiveAuraID,
		35165,
		"Essence of the Martyr",
		stats.SpellPower,
		99,
		time.Second*20,
		core.MajorCooldown{
			CooldownID:       EssenceOfTheMartyrCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.DefensiveTrinketSharedCooldownID,
		},
	))

	var SkullOfGuldanCooldownID = core.NewCooldownID()
	core.AddItemEffect(32483, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		40396,
		"The Skull of Gul'dan",
		stats.SpellHaste,
		175,
		time.Second*20,
		core.MajorCooldown{
			CooldownID:       SkullOfGuldanCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var HexShrunkenHeadCooldownID = core.NewCooldownID()
	core.AddItemEffect(33829, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		43712,
		"Hex Shrunken Head",
		stats.SpellPower,
		211,
		time.Second*20,
		core.MajorCooldown{
			CooldownID:       HexShrunkenHeadCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var ShiftingNaaruSliverCooldownID = core.NewCooldownID()
	core.AddItemEffect(34429, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		45042,
		"Shifting Naaru Sliver",
		stats.SpellPower,
		320,
		time.Second*15,
		core.MajorCooldown{
			CooldownID:       ShiftingNaaruSliverCooldownID,
			Cooldown:         time.Second * 90,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	var DarkIronSmokingPipeCooldownID = core.NewCooldownID()
	core.AddItemEffect(38290, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		51953,
		"Dark Iron Smoking Pipe",
		stats.SpellPower,
		155,
		time.Second*20,
		core.MajorCooldown{
			CooldownID:       DarkIronSmokingPipeCooldownID,
			Cooldown:         time.Minute * 2,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))
}

var MarkOfTheChampionCasterAuraID = core.NewAuraID()

func ApplyMarkOfTheChampionCaster(agent core.Agent) {
	agent.GetCharacter().AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID:   MarkOfTheChampionCasterAuraID,
			Name: "Mark of the Champion (Caster)",
			OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellEffect.Target.MobType == proto.MobType_MobTypeDemon || spellEffect.Target.MobType == proto.MobType_MobTypeUndead {
					spellEffect.BonusSpellPower += 85
				}
			},
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

		return core.Aura{
			ID:   QuagmirransEyeAuraID,
			Name: "Quagmirran's Eye",
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				if icd.IsOnCD(sim) || sim.RandomFloat("Quagmirran's Eye") > 0.1 {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + dur)
				character.AddAuraWithTemporaryStats(sim, FungalFrenzyAuraID, 33370, "Fungal Frenzy", stats.SpellHaste, hasteBonus, time.Second*6)
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

		return core.Aura{
			ID:   ShiffarsNexusHornAuraID,
			Name: "Shiffar's Nexus-Horn",
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellCast.ActionID.ItemID == core.ItemIDTheLightningCapacitor {
					return // TLC can't proc Sextant
				}
				if icd.IsOnCD(sim) || spellEffect.Crit && sim.RandomFloat("Shiffar's Nexus-Horn") > 0.2 {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + dur)
				character.AddAuraWithTemporaryStats(sim, CallOfTheNexusAuraID, 34321, "Call of the Nexus", stats.SpellPower, spellBonus, time.Second*10)
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

		return core.Aura{
			ID:   EyeOfMagtheridonAuraID,
			Name: "Eye of Magtheridon",
			OnSpellMiss: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				character.AddAuraWithTemporaryStats(sim, RecurringPowerAuraID, 34747, "Recurring Power", stats.SpellPower, spellBonus, dur)
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

		return core.Aura{
			ID:   SextantOfUnstableCurrentsAuraID,
			Name: "Sextant of Unstable Currents",
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellCast.ActionID.ItemID == core.ItemIDTheLightningCapacitor {
					return // TLC can't proc Sextant
				}
				if !spellEffect.Crit || icd.IsOnCD(sim) || sim.RandomFloat("Sextant of Unstable Currents") > 0.2 {
					return // if not crit, or on cd, or didn't proc, dont activate
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				character.AddAuraWithTemporaryStats(sim, UnstableCurrentsAuraID, 38348, "Unstable Currents", stats.SpellPower, spellBonus, dur)
			},
		}
	})
}

var DarkmoonCardCrusadeAuraID = core.NewAuraID()
var AuraOfTheCrusadeAuraID = core.NewAuraID()

func ApplyDarkmoonCardCrusade(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const spellBonus = 8.0
		stacks := 0

		return core.Aura{
			ID:   DarkmoonCardCrusadeAuraID,
			Name: "Darkmoon Card Crusade",
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				if stacks < 10 {
					if sim.Log != nil {
						character.Log(sim, "Darkmoon Card Crusade: %d stacks.", stacks)
					}
					stacks++
					character.AddStat(stats.SpellPower, spellBonus)
				}

				// Removal aura will refresh with new total spellpower based on stacks.
				//  This will remove the old stack removal buff.
				character.ReplaceAura(sim, core.Aura{
					ID:      AuraOfTheCrusadeAuraID,
					SpellID: 39441,
					Name:    "Aura of the Crusade",
					Expires: sim.CurrentTime + time.Second*10,
					OnExpire: func(sim *core.Simulation) {
						character.AddStat(stats.SpellPower, -spellBonus*float64(stacks))
						stacks = 0
					},
				})
			},
		}
	})
}
