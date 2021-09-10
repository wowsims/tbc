import { GearListRequest, GearListResult } from './api/newapi';

export class Sim {
  constructor(numWorkers: number) {
  }

  async getGearList(request: GearListRequest): Promise<GearListResult> {
    return Promise.resolve({
      'items': [],
      'enchants': [],
      'gems': [],
    });
  }

  async computeStats() {
  }

  async statWeights() {
  }

  async runSimulation() {
  }
}
