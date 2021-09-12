import { Buffs } from './api/newapi';
import { Enchant } from './api/newapi';
import { Encounter } from './api/newapi';
import { Gem } from './api/newapi';
import { GemColor } from './api/newapi';
import { ItemQuality } from './api/newapi';
import { ItemSlot } from './api/newapi';
import { ItemType } from './api/newapi';
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

import { EquippedItem } from './equipped_item';
import { Listener } from './typed_event';
import { TypedEvent } from './typed_event';
import { wait } from './utils';

export type Gear = Array<EquippedItem | null>;

export class Sim {
  readonly spec: Spec;
  readonly gearListEmitter = new TypedEvent<GearListResult>();
  readonly raceChangeEmitter = new TypedEvent<Race>();
  readonly gearChangeEmitter = new TypedEvent<Gear>();

  private _items: Record<number, Item> = {};
  private _enchants: Record<number, Enchant> = {};
  private _gems: Record<number, Gem> = {};

  private _race: Race;
  private _gear: Gear;
  private _init = false;

  constructor(spec: Spec) {
    this.spec = spec;
    this._race = SpecToEligibleRaces[this.spec][0];
    this._gear = [];
  }

  async init(): Promise<void> {
    if (this._init)
      return;
    this._init = true;

    const result = await this.getGearList(GearListRequest.create({
      spec: this.spec,
    }));

    result.items.forEach(item => this._items[item.id] = item);
    result.enchants.forEach(enchant => this._enchants[enchant.id] = enchant);
    result.gems.forEach(gem => this._gems[gem.id] = gem);

    this.gearListEmitter.emit(result);
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

  equipItem(slot: ItemSlot, item: EquippedItem) {
    this._gear[slot] = item;
    this.gearChangeEmitter.emit(this._gear);
  }

  unequipItem(slot: ItemSlot) {
    this._gear[slot] = null;
    this.gearChangeEmitter.emit(this._gear);
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

  private async getGearList(request: GearListRequest): Promise<GearListResult> {
    return Promise.resolve({
      items: [
        Item.create({
          id: 32235,
          type: ItemType.ItemTypeHead,
          name: 'Cursed Vision of Sargeras',
          quality: ItemQuality.ItemQualityEpic,
          gemSockets: [GemColor.GemColorMeta, GemColor.GemColorYellow],
        }),
      ],
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
