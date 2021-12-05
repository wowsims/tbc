import { RaidBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
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
import { GearListRequest, GearListResult } from '/tbc/core/proto/api.js';
import { IndividualSimRequest, IndividualSimResult } from '/tbc/core/proto/api.js';
import { StatWeightsRequest, StatWeightsResult } from '/tbc/core/proto/api.js';

import { EquippedItem } from '/tbc/core/proto_utils/equipped_item.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
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

import * as OtherConstants from '/tbc/core/constants/other.js';

// Core Sim module which deals only with api types, no UI-related stuff.
export class Sim extends WorkerPool {
  private phase: number = OtherConstants.CURRENT_PHASE;
  private raidBuffs: RaidBuffs = RaidBuffs.create();
  private partyBuffs: PartyBuffs = PartyBuffs.create();
  private individualBuffs: IndividualBuffs = IndividualBuffs.create();

  // Database
  private items: Record<number, Item> = {};
  private enchants: Record<number, Enchant> = {};
  private gems: Record<number, Gem> = {};

  readonly phaseChangeEmitter = new TypedEvent<void>();
  readonly raidBuffsChangeEmitter = new TypedEvent<void>();
  readonly partyBuffsChangeEmitter = new TypedEvent<void>();
  readonly individualBuffsChangeEmitter = new TypedEvent<void>();

  // Emits when any of the above emitters emit.
  readonly changeEmitter = new TypedEvent<void>();

	// Fires when the gear list is finished loading.
  readonly gearListEmitter = new TypedEvent<void>();

  private _init = false;

  constructor() {
		super(3);

    [
      this.raidBuffsChangeEmitter,
      this.partyBuffsChangeEmitter,
      this.individualBuffsChangeEmitter,
      this.phaseChangeEmitter,
    ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));
  }

  async init(): Promise<void> {
    if (this._init)
      return;
    this._init = true;

    const result = await this.getGearList(GearListRequest.create());

    result.items.forEach(item => this.items[item.id] = item);
    result.enchants.forEach(enchant => this.enchants[enchant.id] = enchant);
    result.gems.forEach(gem => this.gems[gem.id] = gem);

    this.gearListEmitter.emit();
  }

	getItems(slot: ItemSlot | undefined): Array<Item> {
		let items = Object.values(this.items);
		if (slot != undefined) {
			items = items.filter(item => getEligibleItemSlots(item).includes(slot));
		}
		return items;
	}

	getEnchants(slot: ItemSlot | undefined): Array<Enchant> {
		let enchants = Object.values(this.enchants);
		if (slot != undefined) {
			enchants = enchants.filter(enchant => getEligibleEnchantSlots(enchant).includes(slot));
		}
		return enchants;
	}

  getGems(socketColor: GemColor | undefined): Array<Gem> {
    let gems = Object.values(this.gems);
		if (socketColor) {
			gems = gems.filter(gem => gemEligibleForSocket(gem, socketColor));
		}
		return gems;
  }

	getMatchingGems(socketColor: GemColor): Array<Gem> {
    return Object.values(this.gems).filter(gem => gemMatchesSocket(gem, socketColor));
	}
  
  getPhase(): number {
    return this.phase;
  }
  setPhase(newPhase: number) {
    if (newPhase != this.phase) {
      this.phase = newPhase;
      this.phaseChangeEmitter.emit();
    }
  }

  getRaidBuffs(): RaidBuffs {
    // Make a defensive copy
    return RaidBuffs.clone(this.raidBuffs);
  }

  setRaidBuffs(newRaidBuffs: RaidBuffs) {
    if (RaidBuffs.equals(this.raidBuffs, newRaidBuffs))
      return;

    // Make a defensive copy
    this.raidBuffs = RaidBuffs.clone(newRaidBuffs);
    this.raidBuffsChangeEmitter.emit();
  }

  getPartyBuffs(): PartyBuffs {
    // Make a defensive copy
    return PartyBuffs.clone(this.partyBuffs);
  }

  setPartyBuffs(newPartyBuffs: PartyBuffs) {
    if (PartyBuffs.equals(this.partyBuffs, newPartyBuffs))
      return;

    // Make a defensive copy
    this.partyBuffs = PartyBuffs.clone(newPartyBuffs);
    this.partyBuffsChangeEmitter.emit();
  }

  getIndividualBuffs(): IndividualBuffs {
    // Make a defensive copy
    return IndividualBuffs.clone(this.individualBuffs);
  }

  setIndividualBuffs(newIndividualBuffs: IndividualBuffs) {
    if (IndividualBuffs.equals(this.individualBuffs, newIndividualBuffs))
      return;

    // Make a defensive copy
    this.individualBuffs = IndividualBuffs.clone(newIndividualBuffs);
    this.individualBuffsChangeEmitter.emit();
  }

  lookupItemSpec(itemSpec: ItemSpec): EquippedItem | null {
    const item = this.items[itemSpec.id];
    if (!item)
      return null;

    const enchant = this.enchants[itemSpec.enchant] || null;
    const gems = itemSpec.gems.map(gemId => this.gems[gemId] || null);

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
      'raidBuffs': RaidBuffs.toJson(this.raidBuffs),
      'partyBuffs': PartyBuffs.toJson(this.partyBuffs),
      'individualBuffs': IndividualBuffs.toJson(this.individualBuffs),
    };
  }

  // Set all the current values, assumes obj is the same type returned by toJson().
  fromJson(obj: any) {
		try {
			this.setRaidBuffs(RaidBuffs.fromJson(obj['raidBuffs']));
		} catch (e) {
			console.warn('Failed to parse raid buffs: ' + e);
		}

		try {
			this.setPartyBuffs(PartyBuffs.fromJson(obj['partyBuffs']));
		} catch (e) {
			console.warn('Failed to parse party buffs: ' + e);
		}

		try {
			this.setIndividualBuffs(IndividualBuffs.fromJson(obj['individualBuffs']));
		} catch (e) {
			console.warn('Failed to parse individual buffs: ' + e);
		}
  }
}
