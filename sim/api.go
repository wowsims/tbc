// The interface to the sim. All interactions with the sim should go through this file.
package sim

import "github.com/wowsims/tbc/sim/api"

/**
 * Returns all items, enchants, and gems recognized by the sim.
 */
func GetGearList(request *api.GearListRequest) *api.GearListResult {
	return getGearListImpl(request)
}

/**
 * Returns character stats taking into account gear / buffs / consumes / etc
 */
func ComputeStats(request *api.ComputeStatsRequest) *api.ComputeStatsResult {
	return computeStatsImpl(request)
}

/**
 * Returns stat weights and EP values, with standard deviations, for all stats.
 */
func StatWeights(request *api.StatWeightsRequest) *api.StatWeightsResult {
	return statWeightsImpl(request)
}

/**
 * Runs multiple iterations of the sim with a single set of options / gear.
 */
func RunSimulation(request *api.IndividualSimRequest) *api.IndividualSimResult {
	return runSimulationImpl(request)
}
