import { ItemSlot } from './api/newapi';
import { Item } from './api/newapi';
import { RaceBonusType } from './api/newapi';
import { Spec } from './api/newapi';
import { SpecToEligibleRaces } from './api/utils';

import { ComputeStatsRequest, ComputeStatsResult } from './api/newapi';
import { GearListRequest, GearListResult } from './api/newapi';
import { IndividualSimRequest, IndividualSimResult } from './api/newapi';
import { StatWeightsRequest, StatWeightsResult } from './api/newapi';

export class Sim {
  readonly spec: Spec;
  race: RaceBonusType;
  gear: Partial<Record<ItemSlot, Item>>;

  constructor(spec: Spec) {
    this.spec = spec;
    this.race = SpecToEligibleRaces[this.spec][0];
    this.gear = {};
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
