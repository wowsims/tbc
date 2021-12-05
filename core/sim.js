import { RaidBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Item } from '/tbc/core/proto/common.js';
import { GearListRequest } from '/tbc/core/proto/api.js';
import { EquippedItem } from '/tbc/core/proto_utils/equipped_item.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { getEligibleItemSlots } from '/tbc/core/proto_utils/utils.js';
import { getEligibleEnchantSlots } from '/tbc/core/proto_utils/utils.js';
import { gemEligibleForSocket } from '/tbc/core/proto_utils/utils.js';
import { gemMatchesSocket } from '/tbc/core/proto_utils/utils.js';
import { TypedEvent } from './typed_event.js';
import { WorkerPool } from './worker_pool.js';
import * as OtherConstants from '/tbc/core/constants/other.js';
// Core Sim module which deals only with api types, no UI-related stuff.
export class Sim extends WorkerPool {
    constructor() {
        super(3);
        this.phase = OtherConstants.CURRENT_PHASE;
        this.raidBuffs = RaidBuffs.create();
        this.partyBuffs = PartyBuffs.create();
        this.individualBuffs = IndividualBuffs.create();
        // Database
        this.items = {};
        this.enchants = {};
        this.gems = {};
        this.phaseChangeEmitter = new TypedEvent();
        this.raidBuffsChangeEmitter = new TypedEvent();
        this.partyBuffsChangeEmitter = new TypedEvent();
        this.individualBuffsChangeEmitter = new TypedEvent();
        // Emits when any of the above emitters emit.
        this.changeEmitter = new TypedEvent();
        // Fires when the gear list is finished loading.
        this.gearListEmitter = new TypedEvent();
        this._init = false;
        [
            this.raidBuffsChangeEmitter,
            this.partyBuffsChangeEmitter,
            this.individualBuffsChangeEmitter,
            this.phaseChangeEmitter,
        ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));
    }
    async init() {
        if (this._init)
            return;
        this._init = true;
        const result = await this.getGearList(GearListRequest.create());
        result.items.forEach(item => this.items[item.id] = item);
        result.enchants.forEach(enchant => this.enchants[enchant.id] = enchant);
        result.gems.forEach(gem => this.gems[gem.id] = gem);
        this.gearListEmitter.emit();
    }
    getItems(slot) {
        let items = Object.values(this.items);
        if (slot != undefined) {
            items = items.filter(item => getEligibleItemSlots(item).includes(slot));
        }
        return items;
    }
    getEnchants(slot) {
        let enchants = Object.values(this.enchants);
        if (slot != undefined) {
            enchants = enchants.filter(enchant => getEligibleEnchantSlots(enchant).includes(slot));
        }
        return enchants;
    }
    getGems(socketColor) {
        let gems = Object.values(this.gems);
        if (socketColor) {
            gems = gems.filter(gem => gemEligibleForSocket(gem, socketColor));
        }
        return gems;
    }
    getMatchingGems(socketColor) {
        return Object.values(this.gems).filter(gem => gemMatchesSocket(gem, socketColor));
    }
    getPhase() {
        return this.phase;
    }
    setPhase(newPhase) {
        if (newPhase != this.phase) {
            this.phase = newPhase;
            this.phaseChangeEmitter.emit();
        }
    }
    getRaidBuffs() {
        // Make a defensive copy
        return RaidBuffs.clone(this.raidBuffs);
    }
    setRaidBuffs(newRaidBuffs) {
        if (RaidBuffs.equals(this.raidBuffs, newRaidBuffs))
            return;
        // Make a defensive copy
        this.raidBuffs = RaidBuffs.clone(newRaidBuffs);
        this.raidBuffsChangeEmitter.emit();
    }
    getPartyBuffs() {
        // Make a defensive copy
        return PartyBuffs.clone(this.partyBuffs);
    }
    setPartyBuffs(newPartyBuffs) {
        if (PartyBuffs.equals(this.partyBuffs, newPartyBuffs))
            return;
        // Make a defensive copy
        this.partyBuffs = PartyBuffs.clone(newPartyBuffs);
        this.partyBuffsChangeEmitter.emit();
    }
    getIndividualBuffs() {
        // Make a defensive copy
        return IndividualBuffs.clone(this.individualBuffs);
    }
    setIndividualBuffs(newIndividualBuffs) {
        if (IndividualBuffs.equals(this.individualBuffs, newIndividualBuffs))
            return;
        // Make a defensive copy
        this.individualBuffs = IndividualBuffs.clone(newIndividualBuffs);
        this.individualBuffsChangeEmitter.emit();
    }
    lookupItemSpec(itemSpec) {
        const item = this.items[itemSpec.id];
        if (!item)
            return null;
        const enchant = this.enchants[itemSpec.enchant] || null;
        const gems = itemSpec.gems.map(gemId => this.gems[gemId] || null);
        return new EquippedItem(item, enchant, gems);
    }
    lookupEquipmentSpec(equipSpec) {
        // EquipmentSpec is supposed to be indexed by slot, but here we assume
        // it isn't just in case.
        const gearMap = {};
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
    toJson() {
        return {
            'raidBuffs': RaidBuffs.toJson(this.raidBuffs),
            'partyBuffs': PartyBuffs.toJson(this.partyBuffs),
            'individualBuffs': IndividualBuffs.toJson(this.individualBuffs),
        };
    }
    // Set all the current values, assumes obj is the same type returned by toJson().
    fromJson(obj) {
        try {
            this.setRaidBuffs(RaidBuffs.fromJson(obj['raidBuffs']));
        }
        catch (e) {
            console.warn('Failed to parse raid buffs: ' + e);
        }
        try {
            this.setPartyBuffs(PartyBuffs.fromJson(obj['partyBuffs']));
        }
        catch (e) {
            console.warn('Failed to parse party buffs: ' + e);
        }
        try {
            this.setIndividualBuffs(IndividualBuffs.fromJson(obj['individualBuffs']));
        }
        catch (e) {
            console.warn('Failed to parse individual buffs: ' + e);
        }
    }
}
