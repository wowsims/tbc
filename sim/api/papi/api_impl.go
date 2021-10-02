// Proto based function interface for the simulator
package papi

import (
	"time"

	"github.com/wowsims/tbc/items"
	"github.com/wowsims/tbc/sim/api"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
	"github.com/wowsims/tbc/sim/runner"
	"github.com/wowsims/tbc/sim/shaman"
)

func getGearListImpl(request *api.GearListRequest) *api.GearListResult {
	result := &api.GearListResult{}

	for i := range items.Items {
		item := items.Items[i]
		result.Items = append(result.Items,
			&api.Item{
				Id:               item.ID,
				Type:             api.ItemType(item.Type),
				ArmorType:        api.ArmorType(item.ArmorType),
				WeaponType:       api.WeaponType(item.WeaponType),
				HandType:         api.HandType(item.HandType),
				RangedWeaponType: api.RangedWeaponType(item.RangedWeaponType),
				Name:             item.Name,
				Stats:            item.Stats[:],
				Phase:            int32(item.Phase),
				Quality:          item.Quality, // Hack until we use generated items
				GemSockets:       item.GemSockets,
			},
		)
	}
	for i := range items.Gems {
		gem := items.Gems[i]
		result.Gems = append(result.Gems, &api.Gem{
			Id:      gem.ID,
			Name:    gem.Name,
			Stats:   gem.Stats[:],
			Color:   gem.Color,
			Phase:   int32(gem.Phase),
			Quality: gem.Quality, // Hack until we use generated items
		})
	}
	for i := range items.Enchants {
		enchant := items.Enchants[i]
		result.Enchants = append(result.Enchants, &api.Enchant{
			Id:       enchant.ID,
			EffectId: enchant.EffectID,
			Name:     enchant.Name,
			Type:     enchant.ItemType,
			Stats:    enchant.Bonus[:],
			Quality:  api.ItemQuality(4),
		})
	}

	return result
}

func computeStatsImpl(request *api.ComputeStatsRequest) *api.ComputeStatsResult {
	return statsFromIndSimRequest(&api.IndividualSimRequest{Player: request.Player, Buffs: request.Buffs})
}

func statsFromIndSimRequest(isr *api.IndividualSimRequest) *api.ComputeStatsResult {
	sim := createSim(isr)
	gearStats := sim.Raid.Parties[0].Players[0].Equip.Stats()
	return &api.ComputeStatsResult{
		GearOnly:   gearStats[:],
		FinalStats: sim.Raid.Parties[0].Players[0].Stats[:], // createSim includes a call to buff up all party members.
		Sets:       []string{},
	}
}

func statWeightsImpl(request *api.StatWeightsRequest) *api.StatWeightsResult {
	statsToWeight := make([]stats.Stat, len(request.StatsToWeigh))
	for i, v := range request.StatsToWeigh {
		statsToWeight[i] = stats.Stat(v)
	}
	result := runner.CalcStatWeight(convertSimParams(request.Options), statsToWeight, stats.Stat(request.EpReferenceStat))
	return &api.StatWeightsResult{
		Weights:       result.Weights,
		WeightsStdev:  result.WeightsStdev,
		EpValues:      result.EpValues,
		EpValuesStdev: result.EpValuesStdev,
	}
}

func convertSimParams(request *api.IndividualSimRequest) runner.IndividualParams {
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
		Equip:       convertEquip(request.Player.Equipment),
		Race:        core.RaceBonusType(request.Player.Options.Race),
		Consumes:    convertConsumes(request.Player.Options.Consumes),
		Buffs:       convertBuffs(request.Buffs),
		Options:     options,
		CustomStats: request.Player.CustomStats,
	}

	switch v := request.Player.Options.Spec.(type) {
	case *api.PlayerOptions_ElementalShaman:
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

func createSim(request *api.IndividualSimRequest) *core.Simulation {
	params := convertSimParams(request)
	sim := runner.SetupIndividualSim(params)

	return sim
}

func runSimulationImpl(request *api.IndividualSimRequest) *api.IndividualSimResult {
	sim := createSim(request)
	result := runner.RunIndividualSim(sim)

	castMetrics := map[int32]*api.CastMetric{}
	for k, v := range result.Casts {
		castMetrics[k] = &api.CastMetric{
			Casts:  v.Casts,
			Crits:  v.Crits,
			Misses: v.Misses,
			Dmgs:   v.Dmgs,
		}
	}
	isr := &api.IndividualSimResult{
		DpsAvg:              result.DpsAvg,
		DpsStdev:            result.DpsStDev,
		DpsHist:             result.DpsHist,
		Logs:                result.Logs,
		DpsMax:              result.DpsMax,
		ExecutionDurationMs: result.ExecutionDurationMs,
		NumOom:              int32(result.NumOom),
		OomAtAvg:            result.OomAtAvg,
		DpsAtOomAvg:         result.DpsAtOomAvg,
		Casts:               castMetrics,
	}
	return isr
}

func convertTotems(inBuff *api.Buffs) shaman.Totems {
	return shaman.Totems{
		TotemOfWrath: int(inBuff.TotemOfWrath),
		WrathOfAir:   inBuff.WrathOfAirTotem != api.TristateEffect_TristateEffectMissing,
		ManaStream:   inBuff.ManaSpringTotem != api.TristateEffect_TristateEffectMissing,
	}
}

func convertShamTalents(t *api.ShamanTalents) shaman.Talents {
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

func convertConsumes(c *api.Consumes) core.Consumes {
	cconsume := core.Consumes{
		FlaskOfBlindingLight:     c.FlaskOfBlindingLight,
		FlaskOfMightyRestoration: c.FlaskOfMightyRestoration,
		FlaskOfPureDeath:         c.FlaskOfPureDeath,
		FlaskOfSupremePower:      c.FlaskOfSupremePower,
		AdeptsElixir:             c.AdeptsElixir,
		ElixirOfMajorFirePower:   c.ElixirOfMajorFirePower,
		ElixirOfMajorFrostPower:  c.ElixirOfMajorFrostPower,
		ElixirOfMajorShadowPower: c.ElixirOfMajorShadowPower,
		ElixirOfDraenicWisdom:    c.ElixirOfDraenicWisdom,
		ElixirOfMajorMageblood:   c.ElixirOfMajorMageblood,
		BrilliantWizardOil:       c.BrilliantWizardOil,
		SuperiorWizardOil:        c.SuperiorWizardOil,
		BlackenedBasilisk:        c.BlackenedBasilisk,
		SkullfishSoup:            c.SkullfishSoup,
		DestructionPotion:        c.DestructionPotion,
		SuperManaPotion:          c.SuperManaPotion,
		DarkRune:                 c.DarkRune,
		DrumsOfBattle:            c.DrumsOfBattle,
		DrumsOfRestoration:       c.DrumsOfRestoration,
	}

	return cconsume
}

func convertEquip(es *api.EquipmentSpec) items.EquipmentSpec {
	coreEquip := items.EquipmentSpec{}

	for i, item := range es.Items {
		spec := items.ItemSpec{
			ID: item.Id,
		}
		spec.Gems = item.Gems
		spec.Enchant = item.Enchant
		coreEquip[i] = spec
	}

	return coreEquip
}

func convertBuffs(inBuff *api.Buffs) core.Buffs {
	// TODO: support tri-state better
	return core.Buffs{
		ArcaneInt:                 inBuff.ArcaneBrilliance,
		GiftOfTheWild:             inBuff.GiftOfTheWild != api.TristateEffect_TristateEffectMissing,
		BlessingOfKings:           inBuff.BlessingOfKings,
		ImprovedBlessingOfWisdom:  inBuff.BlessingOfWisdom != api.TristateEffect_TristateEffectMissing,
		ImprovedDivineSpirit:      inBuff.DivineSpirit != api.TristateEffect_TristateEffectMissing,
		Moonkin:                   inBuff.MoonkinAura != api.TristateEffect_TristateEffectMissing,
		MoonkinRavenGoddess:       inBuff.MoonkinAura == api.TristateEffect_TristateEffectImproved,
		SpriestDPS:                uint16(inBuff.ShadowPriestDps),
		EyeOfNight:                inBuff.EyeOfTheNight,
		TwilightOwl:               inBuff.ChainOfTheTwilightOwl,
		JudgementOfWisdom:         inBuff.JudgementOfWisdom,
		ImprovedSealOfTheCrusader: inBuff.ImprovedSealOfTheCrusader,
		Misery:                    inBuff.Misery,
	}
}
