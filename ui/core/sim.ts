import { Buffs } from './api/newapi';
import { Consumes } from './api/newapi';
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
import { equalsOrBothNull } from './utils';
import { wait } from './utils';

export type Gear = Array<EquippedItem | null>;

export class Sim {
  readonly spec: Spec;
  readonly gearListEmitter = new TypedEvent<GearListResult>();
  readonly raceChangeEmitter = new TypedEvent<Race>();
  readonly gearChangeEmitter = new TypedEvent<Gear>();
  readonly buffsChangeEmitter = new TypedEvent<Buffs>();
  readonly consumesChangeEmitter = new TypedEvent<Consumes>();

  // Database
  private _items: Record<number, Item> = {};
  private _enchants: Record<number, Enchant> = {};
  private _gems: Record<number, Gem> = {};

  // Current values
  private _race: Race;
  private _gear: Gear;
  private _buffs: Buffs;
  private _consumes: Consumes;

  private _init = false;

  constructor(spec: Spec) {
    this.spec = spec;
    this._race = SpecToEligibleRaces[this.spec][0];
    this._gear = [];
    this._buffs = Buffs.create();
    this._consumes = Consumes.create();
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

  get gems(): Array<Gem> {
    return Object.values(this._gems);
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

  get buffs(): Buffs {
    // Make a defensive copy
    return Buffs.clone(this._buffs);
  }

  set buffs(newBuffs: Buffs) {
    if (Buffs.equals(this._buffs, newBuffs))
      return;

    // Make a defensive copy
    this._buffs = Buffs.clone(newBuffs);
    this.buffsChangeEmitter.emit(this._buffs);
  }

  get consumes(): Consumes {
    // Make a defensive copy
    return Consumes.clone(this._consumes);
  }

  set consumes(newConsumes: Consumes) {
    if (Consumes.equals(this._consumes, newConsumes))
      return;

    // Make a defensive copy
    this._consumes = Consumes.clone(newConsumes);
    this.consumesChangeEmitter.emit(this._consumes);
  }

  equipItem(slot: ItemSlot, newItem: EquippedItem | null) {
    if (equalsOrBothNull(this._gear[slot], newItem, (a, b) => a.equals(b)))
      return;

    this._gear[slot] = newItem;
    this.gearChangeEmitter.emit(this._gear);
  }

  getEquippedItem(slot: ItemSlot): EquippedItem | null {
    return this._gear[slot];
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
          id: 29035,
          type: ItemType.ItemTypeHead,
          name: 'Cyclone Facegaurd',
          quality: ItemQuality.ItemQualityEpic,
          gemSockets: [GemColor.GemColorMeta, GemColor.GemColorYellow],
        }),
        Item.create({
          id: 30171,
          type: ItemType.ItemTypeHead,
          name: 'Cataclysm Headpiece',
          quality: ItemQuality.ItemQualityEpic,
          gemSockets: [GemColor.GemColorMeta, GemColor.GemColorYellow],
        }),
        Item.create({
          id: 30169,
          type: ItemType.ItemTypeChest,
          name: 'Cataclysm Chestpiece',
          quality: ItemQuality.ItemQualityEpic,
          gemSockets: [GemColor.GemColorBlue, GemColor.GemColorYellow, GemColor.GemColorYellow],
        }),
        Item.create({
          id: 32235,
          type: ItemType.ItemTypeHead,
          name: 'Cursed Vision of Sargeras',
          quality: ItemQuality.ItemQualityEpic,
          gemSockets: [GemColor.GemColorMeta, GemColor.GemColorYellow],
        }),
      ],
      enchants: [
        Enchant.create({
          id: 29191,
          effectId: 3002,
          name: 'Glyph of Power',
          type: ItemType.ItemTypeHead,
          quality: ItemQuality.ItemQualityUncommon,
        }),
      ],
      gems: [
        Gem.create({
          id: 34220,
          name: 'Chaotic Skyfire Diamond',
          quality: ItemQuality.ItemQualityRare,
          color: GemColor.GemColorMeta,
        }),
        Gem.create({
          id: 23096,
          name: 'Runed Blood Garnet',
          quality: ItemQuality.ItemQualityUncommon,
          color: GemColor.GemColorRed,
        }),
        Gem.create({
          id: 24030,
          name: 'Runed Living Ruby',
          quality: ItemQuality.ItemQualityRare,
          color: GemColor.GemColorRed,
        }),
        Gem.create({
          id: 24059,
          name: 'Potent Noble Topaz',
          quality: ItemQuality.ItemQualityRare,
          color: GemColor.GemColorOrange
        }),
      ],
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

  setWowheadData(equippedItem: EquippedItem, elem: HTMLElement) {
    let parts = [];
    if (equippedItem.gems.length > 0) {
      parts.push('gems=' + equippedItem.gems.map(gem => gem ? gem.id : 0).join(':'));
    }
    if (equippedItem.enchant != null) {
      parts.push('ench=' + equippedItem.enchant.effectId);
    }
    parts.push('pcs=' + this._gear.filter(ei => ei != null).map(ei => ei!.item.id).join(':'));

    elem.setAttribute('data-wowhead', parts.join('&'));
  }
}
