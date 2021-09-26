// The interface to the sim. All interactions with the sim should go through this file.
package api

import "github.com/wowsims/tbc/api/genapi"

/**
 * Returns all items, enchants, and gems recognized by the sim.
 */
func GetGearList(request *genapi.GearListRequest) *genapi.GearListResult {
	return getGearListImpl(request)
}

/**
 * Returns character stats taking into account gear / buffs / consumes / etc
 */
func ComputeStats(request *genapi.ComputeStatsRequest) *genapi.ComputeStatsResult {
	return computeStatsImpl(request)
}

/**
 * Returns stat weights and EP values, with standard deviations, for all stats.
 */
func StatWeights(request *genapi.StatWeightsRequest) *genapi.StatWeightsResult {
	return statWeightsImpl(request)
}

/**
 * Runs multiple iterations of the sim with a single set of options / gear.
 */
func RunSimulation(request *genapi.IndividualSimRequest) *genapi.IndividualSimResult {
	return runSimulationImpl(request)
}
