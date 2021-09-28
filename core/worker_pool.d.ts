import { ComputeStatsRequest, ComputeStatsResult } from './api/api.js';
import { GearListRequest, GearListResult } from './api/api.js';
import { IndividualSimRequest, IndividualSimResult } from './api/api.js';
import { StatWeightsRequest, StatWeightsResult } from './api/api.js';
export declare class WorkerPool {
    private workers;
    constructor(numWorkers: number);
    private getLeastBusyWorker;
    makeApiCall(requestName: string, request: Uint8Array): Promise<Uint8Array>;
    getGearList(request: GearListRequest): Promise<GearListResult>;
    computeStats(request: ComputeStatsRequest): Promise<ComputeStatsResult>;
    statWeights(request: StatWeightsRequest): Promise<StatWeightsResult>;
    individualSim(request: IndividualSimRequest): Promise<IndividualSimResult>;
}
