// Top-level implementations for the go functions.
package api

import (
	"math"
	"math/rand"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/runner"
	"github.com/wowsims/tbc/sim/shaman"
)

func coreGemColorToColor(s []core.GemColor) []GemColor {
	newColors := make([]GemColor, len(s))
	for k, v := range s {
		newColors[k] = GemColor(v)
	}
	return newColors
}

func getGearListImpl(request *GearListRequest) *GearListResult {
	result := &GearListResult{}

	if request.Spec == Spec_SpecElementalShaman {
		for _, v := range shaman.ElementalItems {
			item := core.ItemsByID[v]
			result.Items = append(result.Items,
				&Item{
					Id:         item.ID,
					Type:       ItemType(item.ItemType),
					ArmorType:  ArmorType(item.ArmorType),
					WeaponType: WeaponType(item.WeaponType),
					HandType:   HandType(item.HandType),
					// RangedWeaponType:
					Name:       item.Name,
					Stats:      item.Stats[:],
					Phase:      int32(item.Phase),
					Quality:    ItemQuality(item.Quality),
					GemSockets: coreGemColorToColor(item.GemSockets),
				},
			)
		}
		for _, v := range shaman.ElementalGems {
			gem := core.GemsByID[v]
			result.Gems = append(result.Gems, &Gem{
				Id:      gem.ID,
				Name:    gem.Name,
				Stats:   gem.Stats[:],
				Color:   GemColor(gem.Color),
				Phase:   int32(gem.Phase),
				Quality: ItemQuality(gem.Quality),
			})
		}
	}
	// Enchants: Enchants,

	return result
}

func computeStatsImpl(request *ComputeStatsRequest) *ComputeStatsResult {
	panic("not implemented")
	// fakeSim := core.NewSim(request.Gear, request.Options)

	// sets := fakeSim.ActivateSets()
	// fakeSim.reset() // this will activate any perm-effect items as well

	// gearOnlyStats := fakeSim.Equip.Stats().CalculatedTotal()
	// finalStats := fakeSim.Stats

	// return &ComputeStatsResult{
	// 	GearOnly:   gearOnlyStats,
	// 	FinalStats: finalStats,
	// 	Sets:       sets,
	// }
}

func statWeightsImpl(request *StatWeightsRequest) *StatWeightsResult {
	panic("not implemented")

	// request.Options.AgentType = AGENT_TYPE_ADAPTIVE

	// baselineSimRequest := SimRequest{
	// 	Options:    request.Options,
	// 	Gear:       request.Gear,
	// 	Iterations: request.Iterations,
	// }
	// baselineResult := RunSimulation(baselineSimRequest)

	// var waitGroup sync.WaitGroup
	// result := StatWeightsResult{}
	// dpsHists := [StatLen]map[int]int{}

	// doStat := func(stat Stat, value float64) {
	// 	defer waitGroup.Done()

	// 	simRequest := baselineSimRequest
	// 	simRequest.Options.Buffs.Custom[stat] += value

	// 	simResult := RunSimulation(simRequest)
	// 	result.Weights[stat] = (simResult.DpsAvg - baselineResult.DpsAvg) / value
	// 	dpsHists[stat] = simResult.DpsHist
	// }

	// // Spell hit mod shouldn't go over hit cap.
	// computeStatsResult := ComputeStats(ComputeStatsRequest{
	// 	Options: request.Options,
	// 	Gear:    request.Gear,
	// })
	// spellHitMod := math.Max(0, math.Min(10, 202-computeStatsResult.FinalStats[StatSpellHit]))

	// statMods := Stats{
	// 	StatInt:       50,
	// 	StatSpellDmg:  50,
	// 	StatSpellCrit: 50,
	// 	StatSpellHit:  spellHitMod,
	// 	StatHaste:     50,
	// 	StatMP5:       50,
	// }

	// for stat, mod := range statMods {
	// 	if mod == 0 {
	// 		continue
	// 	}

	// 	waitGroup.Add(1)
	// 	go doStat(Stat(stat), mod)
	// }

	// waitGroup.Wait()

	// for stat, mod := range statMods {
	// 	if mod == 0 {
	// 		continue
	// 	}

	// 	result.EpValues[stat] = result.Weights[stat] / result.Weights[StatSpellPower]
	// 	result.WeightsStDev[stat] = computeStDevFromHists(request.Iterations, mod, dpsHists[stat], baselineResult.DpsHist, nil, statMods[StatSpellDmg])
	// 	result.EpValuesStDev[stat] = computeStDevFromHists(request.Iterations, mod, dpsHists[stat], baselineResult.DpsHist, dpsHists[StatSpellDmg], statMods[StatSpellDmg])
	// }
	// return result
}

func computeStDevFromHists(iters int, modValue float64, moddedStatDpsHist map[int]int, baselineDpsHist map[int]int, spellDmgDpsHist map[int]int, spellDmgModValue float64) float64 {
	sum := 0.0
	sumSquared := 0.0
	n := iters * 10
	for i := 0; i < n; {
		denominator := 1.0
		if spellDmgDpsHist != nil {
			denominator = float64(sampleFromDpsHist(spellDmgDpsHist, iters)-sampleFromDpsHist(baselineDpsHist, iters)) / spellDmgModValue
		}

		if denominator != 0 {
			ep := (float64(sampleFromDpsHist(moddedStatDpsHist, iters)-sampleFromDpsHist(baselineDpsHist, iters)) / modValue) / denominator
			sum += ep
			sumSquared += ep * ep
			i++
		}
	}
	epAvg := sum / float64(n)
	epStDev := math.Sqrt((sumSquared / float64(n)) - (epAvg * epAvg))
	return epStDev
}

func sampleFromDpsHist(hist map[int]int, histNumSamples int) int {
	r := rand.Float64()
	sampleIdx := int(math.Floor(float64(histNumSamples) * r))

	curSampleIdx := 0
	for roundedDps, count := range hist {
		curSampleIdx += count
		if curSampleIdx >= sampleIdx {
			return roundedDps
		}
	}

	panic("Invalid dps histogram")
}

func runSimulationImpl(request *IndividualSimRequest) *IndividualSimResult {

	player := core.NewPlayer(convertEquip(request.Player.Equipment), core.RaceBonusType(request.Player.Options.Race), convertConsumes(request.Player.Options.Consumes))

	// TODO: should this be moved into the player constructor?
	for k, v := range request.Player.CustomStats {
		player.Stats[k] += v
	}

	party := &core.Party{
		Players: []core.PlayerAgent{
			{Player: player},
		},
	}

	var agent core.Agent
	switch v := request.Player.Options.Spec.(type) {
	case *PlayerOptions_ElementalShaman:
		talents := convertShamTalents(v.ElementalShaman.Talents)
		totems := convertTotems(request.Buffs)
		agent = shaman.NewShaman(player, party, talents, totems, int(v.ElementalShaman.Agent.Type))
	default:
		panic("class not supported")
	}

	party.Players[0].Agent = agent
	raid := &core.Raid{Parties: []*core.Party{party}}

	options := core.Options{
		Encounter: core.Encounter{
			Duration:   request.Encounter.Duration,
			NumTargets: int(request.Encounter.NumTargets),
			Armor:      request.Encounter.TargetArmor,
		},
		Iterations: request.Iterations,
		RSeed:      request.RandomSeed,
		ExitOnOOM:  request.ExitOnOom,
		GCDMin:     time.Duration(request.GcdMin),
		Debug:      request.Debug,
	}

	buffs := convertBuffs(request.Buffs)
	sim := runner.SetupSim(raid, buffs, options)

	result := runner.RunIndividualSim(sim, options)

	isr := &IndividualSimResult{
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

func convertTotems(inBuff *Buffs) shaman.Totems {
	return shaman.Totems{
		TotemOfWrath: int(inBuff.TotemOfWrath),
		WrathOfAir:   inBuff.WrathOfAirTotem != 0,
		ManaStream:   inBuff.ManaSpringTotem != 0,
	}
}

func convertShamTalents(t *ShamanTalents) shaman.Talents {
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
	}
}

func convertConsumes(c *Consumes) core.Consumes {
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

func convertEquip(es *EquipmentSpec) core.EquipmentSpec {
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

func convertBuffs(inBuff *Buffs) core.Buffs {
	// TODO: support tri-state
	return core.Buffs{
		ArcaneInt:                 inBuff.ArcaneBrilliance,
		GiftOfTheWild:             inBuff.GiftOfTheWild != TristateEffect_TristateEffectMissing,
		BlessingOfKings:           inBuff.BlessingOfKings,
		ImprovedBlessingOfWisdom:  inBuff.BlessingOfWisdom != TristateEffect_TristateEffectMissing,
		ImprovedDivineSpirit:      inBuff.DivineSpirit != TristateEffect_TristateEffectMissing,
		Moonkin:                   inBuff.MoonkinAura != TristateEffect_TristateEffectMissing,
		MoonkinRavenGoddess:       inBuff.MoonkinAura == TristateEffect_TristateEffectImproved,
		SpriestDPS:                uint16(inBuff.ShadowPriestDps),
		EyeOfNight:                inBuff.EyeOfTheNight,
		TwilightOwl:               inBuff.ChainOfTheTwilightOwl,
		JudgementOfWisdom:         inBuff.JudgementOfWisdom,
		ImprovedSealOfTheCrusader: inBuff.ImprovedSealOfTheCrusader,
		Misery:                    inBuff.Misery,
	}
}
