import { Enchant } from './api/common';
import { Gem } from './api/common';
import { GemColor } from './api/common';
import { Item } from './api/common';
import { ItemQuality } from './api/common';
import { ItemSlot } from './api/common';
import { ItemSpec } from './api/common';
import { ItemType } from './api/common';
import { Stat } from './api/common';

import { ComputeStatsRequest, ComputeStatsResult } from './api/api';
import { GearListRequest, GearListResult } from './api/api';
import { IndividualSimRequest, IndividualSimResult } from './api/api';
import { StatWeightsRequest, StatWeightsResult } from './api/api';

import { wait } from './utils';

const SIM_WORKER_URL = '/sim_worker.js';

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
    return Promise.resolve(ComputeStatsResult.create());
  }

  async statWeights(request: StatWeightsRequest): Promise<StatWeightsResult> {
    const epValues = [];
    epValues[Stat.StatSpellPower] = Math.random() * 2;
    epValues[Stat.StatIntellect] = Math.random() * 2;
    epValues[Stat.StatMP5] = Math.random() * 2;
    epValues[Stat.StatNatureSpellPower] = Math.random() * 2;
    epValues[Stat.StatSpellHit] = Math.random() * 2;
    epValues[Stat.StatSpellCrit] = Math.random() * 2;
    epValues[Stat.StatSpellHaste] = Math.random() * 2;

    const epStDevs = [];
    epStDevs[Stat.StatSpellPower] = Math.random() * 0.5;
    epStDevs[Stat.StatIntellect] = Math.random() * 0.5;
    epStDevs[Stat.StatMP5] = Math.random() * 0.5;
    epStDevs[Stat.StatNatureSpellPower] = Math.random() * 0.5;
    epStDevs[Stat.StatSpellHit] = Math.random() * 0.5;
    epStDevs[Stat.StatSpellCrit] = Math.random() * 0.5;
    epStDevs[Stat.StatSpellHaste] = Math.random() * 0.5;

    return Promise.resolve(StatWeightsResult.create({
      weights: epValues,
      weightsStdev: epStDevs,
      epValues: epValues,
      epValuesStdev: epStDevs,
    }));
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
