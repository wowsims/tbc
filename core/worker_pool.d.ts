import { ComputeStatsRequest, ComputeStatsResult } from './proto/api.js';
import { GearListRequest, GearListResult } from './proto/api.js';
import { RaidSimRequest, RaidSimResult } from './proto/api.js';
import { StatWeightsRequest, StatWeightsResult } from './proto/api.js';
export declare class WorkerPool {
    private workers;
    constructor(numWorkers: number);
    private getLeastBusyWorker;
    makeApiCall(requestName: string, request: Uint8Array): Promise<Uint8Array>;
    getGearList(request: GearListRequest): Promise<GearListResult>;
    computeStats(request: ComputeStatsRequest): Promise<ComputeStatsResult>;
    statWeightsAsync(request: StatWeightsRequest, onProgress: Function): Promise<StatWeightsResult>;
    raidSimAsync(request: RaidSimRequest, onProgress: Function): Promise<RaidSimResult>;
    newProgressHandler(id: string, worker: SimWorker, onProgress: Function): (progressData: any) => void;
}
declare class SimWorker {
    numTasksRunning: number;
    private taskIdsToPromiseFuncs;
    private worker;
    private onReady;
    constructor();
    addPromiseFunc(id: string, callback: (result: any) => void, onError: (error: any) => void): void;
    makeTaskId(): string;
    doApiCall(requestName: string, request: Uint8Array, id: string): Promise<Uint8Array>;
}
export {};
