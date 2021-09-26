import { EquippedItem } from './api/equipped_item';
import { Gear } from './api/gear';
import { Buffs } from './api/common';
import { Class } from './api/common';
import { Consumes } from './api/common';
import { Enchant } from './api/common';
import { Encounter } from './api/common';
import { EquipmentSpec } from './api/common';
import { Gem } from './api/common';
import { GemColor } from './api/common';
import { ItemQuality } from './api/common';
import { ItemSlot } from './api/common';
import { ItemSpec } from './api/common';
import { ItemType } from './api/common';
import { Item } from './api/common';
import { Race } from './api/common';
import { Spec } from './api/common';
import { Stat } from './api/common';
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
import { getEligibleEnchantSlots } from './api/utils';
import { gemEligibleForSocket } from './api/utils';
import { gemMatchesSocket } from './api/utils';

import { Player } from './api/api';
import { PlayerOptions } from './api/api';
import { ComputeStatsRequest, ComputeStatsResult } from './api/api';
import { GearListRequest, GearListResult } from './api/api';
import { IndividualSimRequest, IndividualSimResult } from './api/api';
import { StatWeightsRequest, StatWeightsResult } from './api/api';

import { Listener } from './typed_event';
import { TypedEvent } from './typed_event';
import { wait } from './utils';
import { WorkerPool } from './worker_pool';

export interface SimConfig<SpecType extends Spec> {
  spec: Spec;
  epStats: Array<Stat>;
  epReferenceStat: Stat;
  defaults: {
		epWeights: Stats,
    encounter: Encounter,
    buffs: Buffs,
    consumes: Consumes,
    agent: SpecAgent<SpecType>,
    talents: string,
    specOptions: SpecOptions<SpecType>,
  },
}

// Core Sim module which deals only with api types, no UI-related stuff.
export class Sim<SpecType extends Spec> extends WorkerPool {
  readonly spec: Spec;

  readonly buffsChangeEmitter = new TypedEvent<void>();
  readonly consumesChangeEmitter = new TypedEvent<void>();
  readonly customStatsChangeEmitter = new TypedEvent<void>();
  readonly encounterChangeEmitter = new TypedEvent<void>();
  readonly gearChangeEmitter = new TypedEvent<void>();
  readonly raceChangeEmitter = new TypedEvent<void>();
  readonly agentChangeEmitter = new TypedEvent<void>();
  readonly talentsChangeEmitter = new TypedEvent<void>();
  // Talents dont have all fields so we need this
  readonly talentsStringChangeEmitter = new TypedEvent<void>();
  readonly specOptionsChangeEmitter = new TypedEvent<void>();

  // Emits when any of the above emitters emit.
  readonly changeEmitter = new TypedEvent<void>();

  readonly gearListEmitter = new TypedEvent<void>();

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
	private _epWeights: Stats;

  readonly specTypeFunctions: SpecTypeFunctions<SpecType>;

  private _init = false;

  constructor(config: SimConfig<SpecType>) {
		super(3);

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
		this._epWeights = config.defaults.epWeights;
    this._specOptions = config.defaults.specOptions;

    [
      this.buffsChangeEmitter,
      this.consumesChangeEmitter,
      this.customStatsChangeEmitter,
      this.encounterChangeEmitter,
      this.gearChangeEmitter,
      this.raceChangeEmitter,
      this.agentChangeEmitter,
      this.talentsChangeEmitter,
      this.talentsStringChangeEmitter,
      this.specOptionsChangeEmitter,
    ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));
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

    this.gearListEmitter.emit();
  }

  async statWeights(request: StatWeightsRequest): Promise<StatWeightsResult> {
		const result = await super.statWeights(request);
		this._epWeights = new Stats(result.epValues);
		console.log('updating weights');
		return result;
	}

	getItems(slot: ItemSlot | undefined): Array<Item> {
		let items = Object.values(this._items);
		if (slot != undefined) {
			items = items.filter(item => getEligibleItemSlots(item).includes(slot));
		}
		return items;
	}

	getEnchants(slot: ItemSlot | undefined): Array<Enchant> {
		let enchants = Object.values(this._enchants);
		if (slot != undefined) {
			enchants = enchants.filter(enchant => getEligibleEnchantSlots(enchant).includes(slot));
		}
		return enchants;
	}

  getGems(socketColor: GemColor | undefined): Array<Gem> {
    let gems = Object.values(this._gems);
		if (socketColor) {
			gems = gems.filter(gem => gemEligibleForSocket(gem, socketColor));
		}
		return gems;
  }

	getMatchingGems(socketColor: GemColor): Array<Gem> {
    return Object.values(this._gems).filter(gem => gemMatchesSocket(gem, socketColor));
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

	computeGemEP(gem: Gem): number {
		return new Stats(gem.stats).computeEP(this._epWeights);
	}

	computeEnchantEP(enchant: Enchant): number {
		return new Stats(enchant.stats).computeEP(this._epWeights);
	}

	computeItemEP(item: Item): number {
		if (item == null)
			return 0;

		let ep = new Stats(item.stats).computeEP(this._epWeights);

		const slot = getEligibleItemSlots(item)[0];
		const enchants = this.getEnchants(slot);
		if (enchants.length > 0) {
			ep += Math.max(...enchants.map(enchant => this.computeEnchantEP(enchant)));
		}

		item.gemSockets.forEach(socketColor => {
			const gems = this.getGems(socketColor);
			if (gems.length > 0) {
				ep += Math.max(...gems.map(gem => this.computeGemEP(gem)));
			}
		});

		console.log(item.name + ': ' + ep);
		return ep;
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

  // Returns JSON representing all the current values.
  toJson(): Object {
    return {
      'buffs': Buffs.toJson(this._buffs),
      'consumes': Consumes.toJson(this._consumes),
      'customStats': this._customStats.toJson(),
      'encounter': Encounter.toJson(this._encounter),
      'gear': EquipmentSpec.toJson(this._gear.asSpec()),
      'race': this._race,
      'agent': this.specTypeFunctions.agentToJson(this._agent),
      'talents': this._talentsString,
      'specOptions': this.specTypeFunctions.optionsToJson(this._specOptions),
    };
  }

  // Set all the current values, assumes obj is the same type returned by toJson().
  fromJson(obj: any) {
    this.setBuffs(Buffs.fromJson(obj['buffs']));
    this.setConsumes(Consumes.fromJson(obj['consumes']));
    this.setCustomStats(Stats.fromJson(obj['customStats']));
    this.setEncounter(Encounter.fromJson(obj['encounter']));
    this.setGear(this.lookupEquipmentSpec(EquipmentSpec.fromJson(obj['gear'])));
    this.setRace(obj['race']);
    this.setAgent(this.specTypeFunctions.agentFromJson(obj['agent']));
    this.setTalentsString(obj['talents']);
    this.setSpecOptions(this.specTypeFunctions.optionsFromJson(obj['specOptions']));
  }
}
