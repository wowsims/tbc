import { EquippedItem } from './api/equipped_item';
import { Gear } from './api/gear';
import { Buffs } from './api/newapi';
import { Class } from './api/newapi';
import { Consumes } from './api/newapi';
import { Enchant } from './api/newapi';
import { Encounter } from './api/newapi';
import { EquipmentSpec } from './api/newapi';
import { Gem } from './api/newapi';
import { GemColor } from './api/newapi';
import { ItemQuality } from './api/newapi';
import { ItemSlot } from './api/newapi';
import { ItemSpec } from './api/newapi';
import { ItemType } from './api/newapi';
import { Item } from './api/newapi';
import { Player } from './api/newapi';
import { PlayerOptions } from './api/newapi';
import { Race } from './api/newapi';
import { Spec } from './api/newapi';
import { Stat } from './api/newapi';
import { makeIndividualSimRequest } from './api/request_helpers';
import { Stats } from './api/stats';
import { SpecAgent } from './api/utils';
import { SpecTalents } from './api/utils';
import { SpecTypeFunctions } from './api/utils';
import { specTypeFunctions } from './api/utils';
import { SpecOptions } from './api/utils';
import { specToClass } from './api/utils';
import { specToEligibleRaces } from './api/utils';
import { getEligibleItemSlots } from './api/utils';

import { ComputeStatsRequest, ComputeStatsResult } from './api/newapi';
import { GearListRequest, GearListResult } from './api/newapi';
import { IndividualSimRequest, IndividualSimResult } from './api/newapi';
import { StatWeightsRequest, StatWeightsResult } from './api/newapi';

import { Listener } from './typed_event';
import { TypedEvent } from './typed_event';
import { wait } from './utils';

export interface SimConfig<SpecType extends Spec> {
  spec: Spec;
  epStats: Array<Stat>;
  epReferenceStat: Stat;
  defaults: {
    encounter: Encounter,
    buffs: Buffs,
    consumes: Consumes,
    agent: SpecAgent<SpecType>,
    talents: string,
    specOptions: SpecOptions<SpecType>,
  },
}

export class Sim<SpecType extends Spec> {
  readonly spec: Spec;

  readonly buffsChangeEmitter = new TypedEvent<void>();
  readonly consumesChangeEmitter = new TypedEvent<void>();
  readonly customStatsChangeEmitter = new TypedEvent<void>();
  readonly encounterChangeEmitter = new TypedEvent<void>();
  readonly gearChangeEmitter = new TypedEvent<void>();
  readonly gearListEmitter = new TypedEvent<GearListResult>();
  readonly raceChangeEmitter = new TypedEvent<void>();
  readonly agentChangeEmitter = new TypedEvent<void>();
  readonly talentsChangeEmitter = new TypedEvent<void>();
  // Talents dont have all fields so we need this
  readonly talentsStringChangeEmitter = new TypedEvent<void>();
  readonly specOptionsChangeEmitter = new TypedEvent<void>();

  // Database
  private _items: Record<number, Item> = {};
  private _enchants: Record<number, Enchant> = {};
  private _gems: Record<number, Gem> = {};

  // Current values
  private _buffs: Buffs;
  private _consumes: Consumes;
  private _customStats: Stats;
  private _gear: Gear;
  private _encounter: Encounter;
  private _race: Race;
  private _agent: SpecAgent<SpecType>;
  private _talents: SpecTalents<SpecType>;
  private _talentsString: string;
  private _specOptions: SpecOptions<SpecType>;

  readonly specTypeFunctions: SpecTypeFunctions<SpecType>;

  private _init = false;

  constructor(config: SimConfig<SpecType>) {
    this.spec = config.spec;
    this._race = specToEligibleRaces[this.spec][0];

    this.specTypeFunctions = specTypeFunctions[this.spec] as SpecTypeFunctions<SpecType>;

    this._gear = new Gear({});
    this._encounter = config.defaults.encounter;
    this._buffs = config.defaults.buffs;
    this._consumes = config.defaults.consumes;
    this._customStats = new Stats();
    this._agent = config.defaults.agent;
    this._talents = this.specTypeFunctions.talentsCreate();
    this._talentsString = config.defaults.talents;
    this._specOptions = config.defaults.specOptions;
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

  getGems(): Array<Gem> {
    return Object.values(this._gems);
  }
  
  getRace() {
    return this._race;
  }
  setRace(newRace: Race) {
    if (newRace != this._race) {
      this._race = newRace;
      this.raceChangeEmitter.emit();
    }
  }

  getBuffs(): Buffs {
    // Make a defensive copy
    return Buffs.clone(this._buffs);
  }

  setBuffs(newBuffs: Buffs) {
    if (Buffs.equals(this._buffs, newBuffs))
      return;

    // Make a defensive copy
    this._buffs = Buffs.clone(newBuffs);
    this.buffsChangeEmitter.emit();
  }

  getConsumes(): Consumes {
    // Make a defensive copy
    return Consumes.clone(this._consumes);
  }

  setConsumes(newConsumes: Consumes) {
    if (Consumes.equals(this._consumes, newConsumes))
      return;

    // Make a defensive copy
    this._consumes = Consumes.clone(newConsumes);
    this.consumesChangeEmitter.emit();
  }

  getEncounter(): Encounter {
    // Make a defensive copy
    return Encounter.clone(this._encounter);
  }

  setEncounter(newEncounter: Encounter) {
    if (Encounter.equals(this._encounter, newEncounter))
      return;

    // Make a defensive copy
    this._encounter = Encounter.clone(newEncounter);
    this.encounterChangeEmitter.emit();
  }

  equipItem(slot: ItemSlot, newItem: EquippedItem | null) {
    const newGear = this._gear.withEquippedItem(slot, newItem);
    if (newGear.equals(this._gear))
      return;

    this._gear = newGear;
    this.gearChangeEmitter.emit();
  }

  getEquippedItem(slot: ItemSlot): EquippedItem | null {
    return this._gear.getEquippedItem(slot);
  }

  getGear(): Gear {
    return this._gear;
  }

  setGear(newGear: Gear) {
    if (newGear.equals(this._gear))
      return;

    this._gear = newGear;
    this.gearChangeEmitter.emit();
  }

  getCustomStats(): Stats {
    return this._customStats;
  }

  setCustomStats(newCustomStats: Stats) {
    if (newCustomStats.equals(this._customStats))
      return;

    this._customStats = newCustomStats;
    this.customStatsChangeEmitter.emit();
  }

  getAgent(): SpecAgent<SpecType> {
    return this.specTypeFunctions.agentCopy(this._agent);
  }

  setAgent(newAgent: SpecAgent<SpecType>) {
    if (this.specTypeFunctions.agentEquals(newAgent, this._agent))
      return;

    this._agent = this.specTypeFunctions.agentCopy(newAgent);
    this.agentChangeEmitter.emit();
  }

  // Commented because this should NOT be used; all external uses should be able to use getTalentsString()
  //getTalents(): SpecTalents<SpecType> {
  //  return this.specTypeFunctions.talentsCopy(this._talents);
  //}

  setTalents(newTalents: SpecTalents<SpecType>) {
    if (this.specTypeFunctions.talentsEquals(newTalents, this._talents))
      return;

    this._talents = this.specTypeFunctions.talentsCopy(newTalents);
    this.talentsChangeEmitter.emit();
  }

  getTalentsString(): string {
    return this._talentsString;
  }

  setTalentsString(newTalentsString: string) {
    if (newTalentsString == this._talentsString)
      return;

    this._talentsString = newTalentsString;
    this.talentsStringChangeEmitter.emit();
  }

  getSpecOptions(): SpecOptions<SpecType> {
    return this.specTypeFunctions.optionsCopy(this._specOptions);
  }

  setSpecOptions(newSpecOptions: SpecOptions<SpecType>) {
    if (this.specTypeFunctions.optionsEquals(newSpecOptions, this._specOptions))
      return;

    this._specOptions = this.specTypeFunctions.optionsCopy(newSpecOptions);
    this.specOptionsChangeEmitter.emit();
  }

  lookupItemSpec(itemSpec: ItemSpec): EquippedItem | null {
    const item = this._items[itemSpec.id];
    if (!item)
      return null;

    const enchant = this._enchants[itemSpec.enchant] || null;
    const gems = itemSpec.gems.map(gemId => this._gems[gemId] || null);

    return new EquippedItem(item, enchant, gems);
  }

  lookupEquipmentSpec(equipSpec: EquipmentSpec): Gear {
    // EquipmentSpec is supposed to be indexed by slot, but here we assume
    // it isn't just in case.
    const gearMap: Partial<Record<ItemSlot, EquippedItem | null>> = {};

    equipSpec.items.forEach(itemSpec => {
      const item = this.lookupItemSpec(itemSpec);
      if (!item)
        return;

      const itemSlots = getEligibleItemSlots(item.item);

      const assignedSlot = itemSlots.find(slot => !gearMap[slot]);
      if (assignedSlot == null)
        throw new Error('No slots left to equip ' + Item.toJsonString(item.item));

      gearMap[assignedSlot] = item;
    });

    return new Gear(gearMap);
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

  makeCurrentIndividualSimRequest(iterations: number, debug: boolean): IndividualSimRequest {
    return makeIndividualSimRequest(
      this._buffs,
      this._consumes,
      this._customStats,
      this._encounter,
      this._gear,
      this._race,
      this._agent,
      this._talents,
      this._specOptions,
      iterations,
      debug);
  }

  setWowheadData(equippedItem: EquippedItem, elem: HTMLElement) {
    let parts = [];
    if (equippedItem.gems.length > 0) {
      parts.push('gems=' + equippedItem.gems.map(gem => gem ? gem.id : 0).join(':'));
    }
    if (equippedItem.enchant != null) {
      parts.push('ench=' + equippedItem.enchant.effectId);
    }
    parts.push('pcs=' + this._gear.asArray().filter(ei => ei != null).map(ei => ei!.item.id).join(':'));

    elem.setAttribute('data-wowhead', parts.join('&'));
  }
}
