import { Buffs } from './api/newapi';
import { Encounter } from './api/newapi';
import { ItemSlot } from './api/newapi';
import { Item } from './api/newapi';
import { Player } from './api/newapi';
import { Race } from './api/newapi';
import { Spec } from './api/newapi';
import { SpecToEligibleRaces } from './api/utils';
import { Stat } from './api/newapi';

import { ComputeStatsRequest, ComputeStatsResult } from './api/newapi';
import { GearListRequest, GearListResult } from './api/newapi';
import { IndividualSimRequest, IndividualSimResult } from './api/newapi';
import { StatWeightsRequest, StatWeightsResult } from './api/newapi';

import { wait } from './utils';
import { Listener } from './typed_event';
import { TypedEvent } from './typed_event';

export class Sim {
  readonly spec: Spec;
  readonly raceChangeEmitter = new TypedEvent<Race>();

  private _race: Race;
  private _gear: Partial<Record<ItemSlot, Item>>;

  constructor(spec: Spec) {
    this.spec = spec;
    this._race = SpecToEligibleRaces[this.spec][0];
    this._gear = {};
  }
  
  get race() {
    return this._race;
  }
  set race(newRace: Race) {
    if (newRace != this._race) {
      this._race = newRace;
      this.raceChangeEmitter.emit(newRace);
    }
  }

  currentPlayer(): Player {
    return Player.create();
  }

  currentBuffs(): Buffs {
    return Buffs.create();
  }

  currentEncounter(): Encounter {
    return Encounter.create();
  }

  createSimRequest(): IndividualSimRequest {
    return IndividualSimRequest.create({
      player: this.currentPlayer(),
      buffs: this.currentBuffs(),
      encounter: this.currentEncounter(),
    });
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
    await wait(3000);
    const result = await Promise.resolve(IndividualSimResult.create());
    result.dpsAvg = Math.random() * 2000;
    result.dpsStdev = Math.random() * 200;
    console.log('Individual sim result: ' + IndividualSimResult.toJsonString(result));
    return result;
  }
}
