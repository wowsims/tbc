import { Buffs } from '/tbc/core/proto/common.js';
import { Class } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Enchant } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Gem } from '/tbc/core/proto/common.js';
import { GemColor } from '/tbc/core/proto/common.js';
import { ItemQuality } from '/tbc/core/proto/common.js';
import { ItemSlot } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { ItemType } from '/tbc/core/proto/common.js';
import { Item } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/proto/api.js';
import { PlayerOptions } from '/tbc/core/proto/api.js';
import { ComputeStatsRequest, ComputeStatsResult } from '/tbc/core/proto/api.js';
import { GearListRequest, GearListResult } from '/tbc/core/proto/api.js';
import { IndividualSimRequest, IndividualSimResult } from '/tbc/core/proto/api.js';
import { StatWeightsRequest, StatWeightsResult } from '/tbc/core/proto/api.js';

import { EquippedItem } from '/tbc/core/proto_utils/equipped_item.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { makeComputeStatsRequest } from '/tbc/core/proto_utils/request_helpers.js';
import { makeIndividualSimRequest } from '/tbc/core/proto_utils/request_helpers.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { SpecRotation } from '/tbc/core/proto_utils/utils.js';
import { SpecTalents } from '/tbc/core/proto_utils/utils.js';
import { SpecTypeFunctions } from '/tbc/core/proto_utils/utils.js';
import { specTypeFunctions } from '/tbc/core/proto_utils/utils.js';
import { SpecOptions } from '/tbc/core/proto_utils/utils.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';
import { specToEligibleRaces } from '/tbc/core/proto_utils/utils.js';
import { getEligibleItemSlots } from '/tbc/core/proto_utils/utils.js';
import { getEligibleEnchantSlots } from '/tbc/core/proto_utils/utils.js';
import { gemEligibleForSocket } from '/tbc/core/proto_utils/utils.js';
import { gemMatchesSocket } from '/tbc/core/proto_utils/utils.js';

import { Listener } from './typed_event.js';
import { TypedEvent } from './typed_event.js';
import { sum } from './utils.js';
import { wait } from './utils.js';
import { WorkerPool } from './worker_pool.js';

export interface SimConfig {
  defaults: {
		phase: number,
    encounter: Encounter,
    buffs: Buffs,
  },
}

// Core Sim module which deals only with api types, no UI-related stuff.
export class Sim extends WorkerPool {
  readonly phaseChangeEmitter = new TypedEvent<void>();
  readonly buffsChangeEmitter = new TypedEvent<void>();
  readonly encounterChangeEmitter = new TypedEvent<void>();
  readonly numTargetsChangeEmitter = new TypedEvent<void>();

  // Emits when any of the above emitters emit.
  readonly changeEmitter = new TypedEvent<void>();

  // Database
  private _items: Record<number, Item> = {};
  private _enchants: Record<number, Enchant> = {};
  private _gems: Record<number, Gem> = {};

  readonly gearListEmitter = new TypedEvent<void>();

  // Current values
  private _phase: number;
  private _buffs: Buffs;
  private _encounter: Encounter;
  private _numTargets: number;

  private _init = false;

  constructor(config: SimConfig) {
		super(3);

		this._phase = config.defaults.phase;
    this._buffs = config.defaults.buffs;
    this._encounter = config.defaults.encounter;
    this._numTargets = 1;

    [
      this.buffsChangeEmitter,
      this.encounterChangeEmitter,
      this.numTargetsChangeEmitter,
      this.phaseChangeEmitter,
    ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));
  }

  async init(spec: Spec): Promise<void> {
    if (this._init)
      return;
    this._init = true;

    const result = await this.getGearList(GearListRequest.create({
      spec: spec,
    }));

    result.items.forEach(item => this._items[item.id] = item);
    result.enchants.forEach(enchant => this._enchants[enchant.id] = enchant);
    result.gems.forEach(gem => this._gems[gem.id] = gem);

    this.gearListEmitter.emit();
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
  
  getPhase(): number {
    return this._phase;
  }
  setPhase(newPhase: number) {
    if (newPhase != this._phase) {
      this._phase = newPhase;
      this.phaseChangeEmitter.emit();
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
  
  getNumTargets(): number {
    return this._numTargets;
  }
  setNumTargets(newNumTargets: number) {
    if (newNumTargets == this._numTargets)
			return;

		this._numTargets = newNumTargets;
		this.numTargetsChangeEmitter.emit();
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

  // Returns JSON representing all the current values.
  toJson(): Object {
    return {
      'buffs': Buffs.toJson(this._buffs),
      'encounter': Encounter.toJson(this._encounter),
      'numTargets': this._numTargets,
    };
  }

  // Set all the current values, assumes obj is the same type returned by toJson().
  fromJson(obj: any) {
		try {
			this.setBuffs(Buffs.fromJson(obj['buffs']));
		} catch (e) {
			console.warn('Failed to parse buffs: ' + e);
		}

		try {
			this.setEncounter(Encounter.fromJson(obj['encounter']));
		} catch (e) {
			console.warn('Failed to parse encounter: ' + e);
		}

		const parsedNumTargets = parseInt(obj['numTargets']);
		if (!isNaN(parsedNumTargets) && parsedNumTargets != 0) {
			this._numTargets = parsedNumTargets;
		}
  }
}
