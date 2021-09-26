// Proto based function interface for the simulator
package api

import (
	"time"

	"github.com/wowsims/tbc/api/genapi"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/runner"
	"github.com/wowsims/tbc/sim/shaman"
)

func coreGemColorToColor(s []core.GemColor) []genapi.GemColor {
	newColors := make([]genapi.GemColor, len(s))
	for k, v := range s {
		newColors[k] = genapi.GemColor(v)
	}
	return newColors
}

func getGearListImpl(request *genapi.GearListRequest) *genapi.GearListResult {
	result := &genapi.GearListResult{}

	if request.Spec == genapi.Spec_SpecElementalShaman {
		for _, v := range shaman.ElementalItems {
			item := core.ItemsByID[v]

			result.Items = append(result.Items,
				&genapi.Item{
					Id:               item.ID,
					Type:             genapi.ItemType(item.ItemType),
					ArmorType:        genapi.ArmorType(item.ArmorType),
					WeaponType:       genapi.WeaponType(item.WeaponType),
					HandType:         genapi.HandType(item.HandType),
					RangedWeaponType: genapi.RangedWeaponType(item.RangedWeaponType),
					Name:             item.Name,
					Stats:            item.Stats[:],
					Phase:            int32(item.Phase),
					Quality:          genapi.ItemQuality(item.Quality + 1), // Hack until we use generated items
					GemSockets:       coreGemColorToColor(item.GemSockets),
				},
			)
		}
		for _, v := range shaman.ElementalGems {
			gem := core.GemsByID[v]
			result.Gems = append(result.Gems, &genapi.Gem{
				Id:      gem.ID,
				Name:    gem.Name,
				Stats:   gem.Stats[:],
				Color:   genapi.GemColor(gem.Color),
				Phase:   int32(gem.Phase),
				Quality: genapi.ItemQuality(gem.Quality + 1), // Hack until we use generated items
			})
		}
		for _, v := range shaman.ElementalEnchants {
			enchant := core.EnchantsByID[v]
			result.Enchants = append(result.Enchants, &genapi.Enchant{
				Id:       enchant.ID,
				EffectId: enchant.EffectID,
				Name:     enchant.Name,
				Type:     genapi.ItemType(enchant.ItemType),
				Stats:    enchant.Bonus[:],
				Quality:  genapi.ItemQuality(4),
			})
		}
	}
	// Enchants: Enchants,

	return result
}

func computeStatsImpl(request *genapi.ComputeStatsRequest) *genapi.ComputeStatsResult {
	return statsFromIndSimRequest(&genapi.IndividualSimRequest{Player: request.Player, Buffs: request.Buffs})
}

func statsFromIndSimRequest(isr *genapi.IndividualSimRequest) *genapi.ComputeStatsResult {
	sim := createSim(isr)
	gearStats := sim.Raid.Parties[0].Players[0].Equip.Stats()
	return &genapi.ComputeStatsResult{
		GearOnly:   gearStats[:],
		FinalStats: sim.Raid.Parties[0].Players[0].Stats[:], // createSim includes a call to buff up all party members.
		Sets:       []string{},
	}
}

func statWeightsImpl(request *genapi.StatWeightsRequest) *genapi.StatWeightsResult {
	statsToWeight := make([]core.Stat, len(request.StatsToWeigh))
	for i, v := range request.StatsToWeigh {
		statsToWeight[i] = core.Stat(v)
	}
	result := runner.CalcStatWeight(convertSimParams(request.Options), statsToWeight, core.Stat(request.EpReferenceStat))
	return &genapi.StatWeightsResult{
		Weights:       result.Weights,
		WeightsStdev:  result.WeightsStdev,
		EpValues:      result.EpValues,
		EpValuesStdev: result.EpValuesStdev,
	}
}

func convertSimParams(request *genapi.IndividualSimRequest) runner.IndividualParams {
	options := core.Options{
		Iterations: int(request.Iterations),
		RSeed:      request.RandomSeed,
		ExitOnOOM:  request.ExitOnOom,
		GCDMin:     time.Duration(request.GcdMin),
		Debug:      request.Debug,
	}
	if request.Encounter != nil {
		options.Encounter = core.Encounter{
			Duration:   request.Encounter.Duration,
			NumTargets: int(request.Encounter.NumTargets),
			Armor:      request.Encounter.TargetArmor,
		}
	}

	params := runner.IndividualParams{
		Equip:    convertEquip(request.Player.Equipment),
		Race:     core.RaceBonusType(request.Player.Options.Race),
		Consumes: convertConsumes(request.Player.Options.Consumes),
		Buffs:    convertBuffs(request.Buffs),
		Options:  options,
	}

	switch v := request.Player.Options.Spec.(type) {
	case *genapi.PlayerOptions_ElementalShaman:
		talents := convertShamTalents(v.ElementalShaman.Talents)
		totems := convertTotems(request.Buffs)
		params.Spec = shaman.ElementalSpec{
			Talents: talents,
			Totems:  totems,
			AgentID: shaman.AgentType(v.ElementalShaman.Agent.Type),
		}

	default:
		panic("class not supported")
	}

	return params
}

func createSim(request *genapi.IndividualSimRequest) *core.Simulation {
	params := convertSimParams(request)
	sim := runner.SetupIndividualSim(params)

	return sim
}

func runSimulationImpl(request *genapi.IndividualSimRequest) *genapi.IndividualSimResult {
	sim := createSim(request)
	result := runner.RunIndividualSim(sim)

	isr := &genapi.IndividualSimResult{
		DpsAvg:              result.DpsAvg,
		DpsStdev:            result.DpsStDev,
		DpsHist:             result.DpsHist,
		Logs:                result.Logs,
		DpsMax:              result.DpsMax,
		ExecutionDurationMs: result.ExecutionDurationMs,
		NumOom:              int32(result.NumOom),
		OomAtAvg:            result.OomAtAvg,
		DpsAtOomAvg:         result.DpsAtOomAvg,
		// TODO: convert casts
	}
	return isr
}

func convertTotems(inBuff *genapi.Buffs) shaman.Totems {
	return shaman.Totems{
		TotemOfWrath: int(inBuff.TotemOfWrath),
		WrathOfAir:   inBuff.WrathOfAirTotem != 0,
		ManaStream:   inBuff.ManaSpringTotem != 0,
	}
}

func convertShamTalents(t *genapi.ShamanTalents) shaman.Talents {
	return shaman.Talents{
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

func convertConsumes(c *genapi.Consumes) core.Consumes {
	cconsume := core.Consumes{
		BrilliantWizardOil:       c.BrilliantWizardOil,
		MajorMageblood:           c.ElixirOfMajorMageblood,
		FlaskOfBlindingLight:     c.FlaskOfBlindingLight,
		FlaskOfMightyRestoration: c.FlaskOfMightyRestoration,
		BlackendBasilisk:         c.BlackenedBasilisk,
		DestructionPotion:        c.DestructionPotion,
		SuperManaPotion:          c.SuperManaPotion,
		DarkRune:                 c.DarkRune,
		DrumsOfBattle:            c.DrumsOfBattle,
	}

	return cconsume
}

func convertEquip(es *genapi.EquipmentSpec) core.EquipmentSpec {
	coreEquip := core.EquipmentSpec{}

	for i, item := range es.Items {
		spec := core.ItemSpec{
			ID: item.Id,
		}
		spec.Gems = item.Gems
		spec.Enchant = item.Enchant

		coreEquip[i] = spec
	}

	return coreEquip
}

func convertBuffs(inBuff *genapi.Buffs) core.Buffs {
	// TODO: support tri-state better
	return core.Buffs{
		ArcaneInt:                 inBuff.ArcaneBrilliance,
		GiftOfTheWild:             inBuff.GiftOfTheWild != genapi.TristateEffect_TristateEffectMissing,
		BlessingOfKings:           inBuff.BlessingOfKings,
		ImprovedBlessingOfWisdom:  inBuff.BlessingOfWisdom != genapi.TristateEffect_TristateEffectMissing,
		ImprovedDivineSpirit:      inBuff.DivineSpirit != genapi.TristateEffect_TristateEffectMissing,
		Moonkin:                   inBuff.MoonkinAura != genapi.TristateEffect_TristateEffectMissing,
		MoonkinRavenGoddess:       inBuff.MoonkinAura == genapi.TristateEffect_TristateEffectImproved,
		SpriestDPS:                uint16(inBuff.ShadowPriestDps),
		EyeOfNight:                inBuff.EyeOfTheNight,
		TwilightOwl:               inBuff.ChainOfTheTwilightOwl,
		JudgementOfWisdom:         inBuff.JudgementOfWisdom,
		ImprovedSealOfTheCrusader: inBuff.ImprovedSealOfTheCrusader,
		Misery:                    inBuff.Misery,
	}
}
