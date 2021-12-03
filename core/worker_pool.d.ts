import { ComputeStatsRequest, ComputeStatsResult } from './proto/api.js';
import { GearListRequest, GearListResult } from './proto/api.js';
import { IndividualSimRequest, IndividualSimResult } from './proto/api.js';
import { RaidSimRequest, RaidSimResult } from './proto/api.js';
import { StatWeightsRequest, StatWeightsResult } from './proto/api.js';
export declare class WorkerPool {
    private workers;
    constructor(numWorkers: number);
    private getLeastBusyWorker;
    makeApiCall(requestName: string, request: Uint8Array): Promise<Uint8Array>;
    getGearList(request: GearListRequest): Promise<GearListResult>;
    computeStats(request: ComputeStatsRequest): Promise<ComputeStatsResult>;
    statWeights(request: StatWeightsRequest): Promise<StatWeightsResult>;
    individualSim(request: IndividualSimRequest): Promise<IndividualSimResult>;
    raidSim(request: RaidSimRequest): Promise<RaidSimResult>;
}
