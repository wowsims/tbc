// The interface to the sim. All interactions with the sim should go through this file.
package api

/**
 * Returns all items, enchants, and gems recognized by the sim.
 */
func GetGearList(request *GearListRequest) *GearListResult {
	return getGearListImpl(request)
}

/**
 * Returns character stats taking into account gear / buffs / consumes / etc
 */
func ComputeStats(request *ComputeStatsRequest) *ComputeStatsResult {
	return computeStatsImpl(request)
}

/**
 * Returns stat weights and EP values, with standard deviations, for all stats.
 */
func StatWeights(request *StatWeightsRequest) *StatWeightsResult {
	return statWeightsImpl(request)
}

/**
 * Runs multiple iterations of the sim with a single set of options / gear.
 */
func RunSimulation(request *IndividualSimRequest) *IndividualSimResult {
	return runSimulationImpl(request)
}
