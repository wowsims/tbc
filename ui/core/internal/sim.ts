import { ComputeStatsRequest, ComputeStatsResult } from '../api/newapi';
import { GearListRequest, GearListResult } from '../api/newapi';
import { IndividualSimRequest, IndividualSimResult } from '../api/newapi';
import { StatWeightsRequest, StatWeightsResult } from '../api/newapi';

export class Sim {
  constructor(numWorkers: number) {
  }

  async getGearList(request: GearListRequest): Promise<GearListResult> {
    return Promise.resolve({
      items: [],
      enchants: [],
      gems: [],
    });
  }

  async computeStats(request: ComputeStatsRequest): Promise<ComputeStatsResult> {
    return Promise.resolve(ComputeStatsResult.create());
  }

  async statWeights(request: StatWeightsRequest): Promise<StatWeightsResult> {
    return Promise.resolve(StatWeightsResult.create());
  }

  async individualSim(request: IndividualSimRequest): Promise<IndividualSimResult> {
    return Promise.resolve(IndividualSimResult.create());
  }
}
