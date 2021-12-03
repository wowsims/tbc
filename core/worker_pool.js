import { ComputeStatsRequest, ComputeStatsResult } from './proto/api.js';
import { GearListRequest, GearListResult } from './proto/api.js';
import { IndividualSimRequest, IndividualSimResult } from './proto/api.js';
import { RaidSimRequest, RaidSimResult } from './proto/api.js';
import { StatWeightsRequest, StatWeightsResult } from './proto/api.js';
import { repoName } from './resources.js';
const SIM_WORKER_URL = `/${repoName}/sim_worker.js`;
export class WorkerPool {
    constructor(numWorkers) {
        this.workers = [];
        for (let i = 0; i < numWorkers; i++) {
            this.workers.push(new SimWorker());
        }
    }
    getLeastBusyWorker() {
        return this.workers.reduce((curMinWorker, nextWorker) => curMinWorker.numTasksRunning < nextWorker.numTasksRunning ?
            curMinWorker : nextWorker);
    }
    async makeApiCall(requestName, request) {
        return await this.getLeastBusyWorker().doApiCall(requestName, request);
    }
    async getGearList(request) {
        const result = await this.makeApiCall('gearList', GearListRequest.toBinary(request));
        return GearListResult.fromBinary(result);
    }
    async computeStats(request) {
        const result = await this.makeApiCall('computeStats', ComputeStatsRequest.toBinary(request));
        return ComputeStatsResult.fromBinary(result);
    }
    async statWeights(request) {
        const result = await this.makeApiCall('statWeights', StatWeightsRequest.toBinary(request));
        return StatWeightsResult.fromBinary(result);
    }
    async individualSim(request) {
        console.log('Individual sim request: ' + IndividualSimRequest.toJsonString(request));
        const resultData = await this.makeApiCall('individualSim', IndividualSimRequest.toBinary(request));
        const result = IndividualSimResult.fromBinary(resultData);
        console.log('Individual sim result: ' + IndividualSimResult.toJsonString(result));
        return result;
    }
    async raidSim(request) {
        console.log('Raid sim request: ' + RaidSimRequest.toJsonString(request));
        const resultData = await this.makeApiCall('raidSim', RaidSimRequest.toBinary(request));
        const result = RaidSimResult.fromBinary(resultData);
        console.log('Raid sim result: ' + RaidSimResult.toJsonString(result));
        return result;
    }
}
class SimWorker {
    constructor() {
        this.numTasksRunning = 0;
        this.taskIdsToPromiseFuncs = {};
        this.worker = new window.Worker(SIM_WORKER_URL);
        let resolveReady = null;
        this.onReady = new Promise((_resolve, _reject) => {
            resolveReady = _resolve;
        });
        this.worker.onmessage = event => {
            if (event.data.msg == 'ready') {
                this.worker.postMessage({ msg: 'setID', id: '1' });
                resolveReady();
            }
            else if (event.data.msg == 'idconfirm') {
                // Do nothing
            }
            else {
                const id = event.data.id;
                if (!this.taskIdsToPromiseFuncs[id]) {
                    console.warn('Unrecognized result id: ' + id);
                    return;
                }
                const promiseFuncs = this.taskIdsToPromiseFuncs[id];
                delete this.taskIdsToPromiseFuncs[id];
                this.numTasksRunning--;
                promiseFuncs[0](event.data.outputData);
            }
        };
    }
    makeTaskId() {
        let id = '';
        const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
        for (let i = 0; i < 16; i++) {
            id += characters.charAt(Math.floor(Math.random() * characters.length));
        }
        return id;
    }
    async doApiCall(requestName, request) {
        this.numTasksRunning++;
        await this.onReady;
        const taskPromise = new Promise((resolve, reject) => {
            const id = this.makeTaskId();
            this.taskIdsToPromiseFuncs[id] = [resolve, reject];
            this.worker.postMessage({
                msg: requestName,
                id: id,
                inputData: request,
            });
        });
        return await taskPromise;
    }
}
