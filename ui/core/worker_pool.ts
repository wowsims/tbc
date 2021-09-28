import { Enchant } from './api/common.js';
import { Gem } from './api/common.js';
import { GemColor } from './api/common.js';
import { Item } from './api/common.js';
import { ItemQuality } from './api/common.js';
import { ItemSlot } from './api/common.js';
import { ItemSpec } from './api/common.js';
import { ItemType } from './api/common.js';
import { Stat } from './api/common.js';

import { ComputeStatsRequest, ComputeStatsResult } from './api/api.js';
import { GearListRequest, GearListResult } from './api/api.js';
import { IndividualSimRequest, IndividualSimResult } from './api/api.js';
import { StatWeightsRequest, StatWeightsResult } from './api/api.js';

import { urlPathPrefix } from './resources.js';
import { wait } from './utils.js';

const SIM_WORKER_URL = `${urlPathPrefix}/sim_worker.js`;

export class WorkerPool {
	private workers: Array<SimWorker>;

  constructor(numWorkers: number) {
    this.workers = [];
    for (let i = 0; i < numWorkers; i++) {
      this.workers.push(new SimWorker());
    } }

  private getLeastBusyWorker(): SimWorker {
    return this.workers.reduce(
        (curMinWorker, nextWorker) => curMinWorker.numTasksRunning < nextWorker.numTasksRunning ?
            curMinWorker : nextWorker);
  }

  async makeApiCall(requestName: string, request: Uint8Array): Promise<Uint8Array> {
    return await this.getLeastBusyWorker().doApiCall(requestName, request);
  }

  async getGearList(request: GearListRequest): Promise<GearListResult> {
		const result = await this.makeApiCall('gearList', GearListRequest.toBinary(request));
		return GearListResult.fromBinary(result);
  }

  async computeStats(request: ComputeStatsRequest): Promise<ComputeStatsResult> {
		const result = await this.makeApiCall('computeStats', ComputeStatsRequest.toBinary(request));
		return ComputeStatsResult.fromBinary(result);
  }

  async statWeights(request: StatWeightsRequest): Promise<StatWeightsResult> {
		const result = await this.makeApiCall('statWeights', StatWeightsRequest.toBinary(request));
		return StatWeightsResult.fromBinary(result);
  }

  async individualSim(request: IndividualSimRequest): Promise<IndividualSimResult> {
    console.log('Individual sim request: ' + IndividualSimRequest.toJsonString(request));
		const resultData = await this.makeApiCall('individualSim', IndividualSimRequest.toBinary(request));
		const result = IndividualSimResult.fromBinary(resultData);
    console.log('Individual sim result: ' + IndividualSimResult.toJsonString(result));
		return result;
  }
}

class SimWorker {
	numTasksRunning: number;
	private taskIdsToPromiseFuncs: Record<string, [(result: any) => void, (error: any) => void]>;
	private worker: Worker;
	private onReady: Promise<void>;

  constructor() {
    this.numTasksRunning = 0;
    this.taskIdsToPromiseFuncs = {};
    this.worker = new window.Worker(SIM_WORKER_URL);

    let resolveReady: (() => void) | null = null;
    this.onReady = new Promise((_resolve, _reject) => {
      resolveReady = _resolve;
    });

    this.worker.onmessage = event => {
      if (event.data.msg == 'ready') {
				this.worker.postMessage({ msg: 'setID', id: '1' });
				resolveReady!();
      } else if (event.data.msg == 'idconfirm') {
        // Do nothing
      } else {
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

  makeTaskId(): string {
		let id = '';
		const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
		for (let i = 0; i < 16; i++) {
			id += characters.charAt(Math.floor(Math.random() * characters.length));
		}
		return id;
  }

  async doApiCall(requestName: string, request: Uint8Array): Promise<Uint8Array> {
    this.numTasksRunning++;
    await this.onReady;

    const taskPromise = new Promise<Uint8Array>((resolve, reject) => {
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
