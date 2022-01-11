import { Item } from '/tbc/core/proto/common.js';
import { ComputeStatsRequest } from '/tbc/core/proto/api.js';
import { GearListRequest } from '/tbc/core/proto/api.js';
import { RaidSimRequest } from '/tbc/core/proto/api.js';
import { SimOptions } from '/tbc/core/proto/api.js';
import { StatWeightsRequest, StatWeightsResult } from '/tbc/core/proto/api.js';
import { EquippedItem } from '/tbc/core/proto_utils/equipped_item.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { SimResult } from '/tbc/core/proto_utils/sim_result.js';
import { gemEligibleForSocket } from '/tbc/core/proto_utils/gems.js';
import { gemMatchesSocket } from '/tbc/core/proto_utils/gems.js';
import { getEligibleItemSlots } from '/tbc/core/proto_utils/utils.js';
import { getEligibleEnchantSlots } from '/tbc/core/proto_utils/utils.js';
import { Encounter } from './encounter.js';
import { Raid } from './raid.js';
import { TypedEvent } from './typed_event.js';
import { WorkerPool } from './worker_pool.js';
import * as OtherConstants from '/tbc/core/constants/other.js';
// Core Sim module which deals only with api types, no UI-related stuff.
export class Sim {
    constructor() {
        this.iterations = 3000;
        this.phase = OtherConstants.CURRENT_PHASE;
        // Database
        this.items = {};
        this.enchants = {};
        this.gems = {};
        this.iterationsChangeEmitter = new TypedEvent();
        this.phaseChangeEmitter = new TypedEvent();
        // Fires when a raid sim API call completes.
        this.simResultEmitter = new TypedEvent();
        // These callbacks are needed so we can apply BuffBot modifications automatically before sending requests.
        this.modifyRaidProto = () => { };
        this.modifyEncounterProto = () => { };
        this.workerPool = new WorkerPool(3);
        this._initPromise = this.workerPool.getGearList(GearListRequest.create()).then(result => {
            result.items.forEach(item => this.items[item.id] = item);
            result.enchants.forEach(enchant => this.enchants[enchant.id] = enchant);
            result.gems.forEach(gem => this.gems[gem.id] = gem);
        });
        this.raid = new Raid(this);
        this.encounter = new Encounter(this);
        this.changeEmitter = TypedEvent.onAny([
            this.iterationsChangeEmitter,
            this.phaseChangeEmitter,
            this.raid.changeEmitter,
            this.encounter.changeEmitter,
        ]);
        this.raid.changeEmitter.on(eventID => this.updateCharacterStats(eventID));
    }
    waitForInit() {
        return this._initPromise;
    }
    setModifyRaidProto(newModFn) {
        this.modifyRaidProto = newModFn;
    }
    getModifiedRaidProto() {
        const raidProto = this.raid.toProto();
        this.modifyRaidProto(raidProto);
        // Remove any inactive meta gems, since the backend doesn't have its own validation.
        raidProto.parties.forEach(party => {
            party.players.forEach(player => {
                if (!player.equipment) {
                    return;
                }
                const gear = this.lookupEquipmentSpec(player.equipment);
                if (gear.hasInactiveMetaGem()) {
                    player.equipment = gear.withoutMetaGem().asSpec();
                }
            });
        });
        return raidProto;
    }
    setModifyEncounterProto(newModFn) {
        this.modifyEncounterProto = newModFn;
    }
    getModifiedEncounterProto() {
        const encounterProto = this.encounter.toProto();
        this.modifyEncounterProto(encounterProto);
        return encounterProto;
    }
    makeRaidSimRequest(debug) {
        return RaidSimRequest.create({
            raid: this.getModifiedRaidProto(),
            encounter: this.getModifiedEncounterProto(),
            simOptions: SimOptions.create({
                iterations: debug ? 1 : this.getIterations(),
                debugFirstIteration: true,
            }),
        });
    }
    async runRaidSim(eventID) {
        if (this.raid.isEmpty()) {
            throw new Error('Raid is empty! Try adding some players first.');
        }
        else if (this.encounter.getNumTargets() < 1) {
            throw new Error('Encounter has no targets! Try adding some targets first.');
        }
        await this.waitForInit();
        const request = this.makeRaidSimRequest(false);
        const result = await this.workerPool.raidSim(request);
        const simResult = await SimResult.makeNew(request, result);
        this.simResultEmitter.emit(eventID, simResult);
        return simResult;
    }
    async runRaidSimWithLogs(eventID) {
        if (this.raid.isEmpty()) {
            throw new Error('Raid is empty! Try adding some players first.');
        }
        else if (this.encounter.getNumTargets() < 1) {
            throw new Error('Encounter has no targets! Try adding some targets first.');
        }
        await this.waitForInit();
        const request = this.makeRaidSimRequest(true);
        const result = await this.workerPool.raidSim(request);
        const simResult = await SimResult.makeNew(request, result);
        this.simResultEmitter.emit(eventID, simResult);
        return simResult;
    }
    // This should be invoked internally whenever stats might have changed.
    async updateCharacterStats(eventID) {
        if (eventID == 0) {
            // Skip the first event ID because it interferes with the loaded stats.
            return;
        }
        await this.waitForInit();
        // Capture the current players so we avoid issues if something changes while
        // request is in-flight.
        const players = this.raid.getPlayers();
        const result = await this.workerPool.computeStats(ComputeStatsRequest.create({
            raid: this.getModifiedRaidProto(),
        }));
        TypedEvent.freezeAllAndDo(() => {
            result.raidStats.parties
                .forEach((partyStats, partyIndex) => partyStats.players.forEach((playerStats, playerIndex) => players[partyIndex * 5 + playerIndex]?.setCurrentStats(eventID, playerStats)));
        });
    }
    async statWeights(player, epStats, epReferenceStat) {
        if (this.raid.isEmpty()) {
            throw new Error('Raid is empty! Try adding some players first.');
        }
        else if (this.encounter.getNumTargets() < 1) {
            throw new Error('Encounter has no targets! Try adding some targets first.');
        }
        await this.waitForInit();
        if (player.getParty() == null) {
            console.warn('Trying to get stat weights without a party!');
            return StatWeightsResult.create();
        }
        else {
            const request = StatWeightsRequest.create({
                player: player.toProto(),
                raidBuffs: this.raid.getBuffs(),
                partyBuffs: player.getParty().getBuffs(),
                encounter: this.encounter.toProto(),
                simOptions: SimOptions.create({
                    iterations: this.getIterations(),
                    debug: false,
                }),
                statsToWeigh: epStats,
                epReferenceStat: epReferenceStat,
            });
            return await this.workerPool.statWeights(request);
        }
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
    setPhase(eventID, newPhase) {
        if (newPhase != this.phase) {
            this.phase = newPhase;
            this.phaseChangeEmitter.emit(eventID);
        }
    }
    getIterations() {
        return this.iterations;
    }
    setIterations(eventID, newIterations) {
        if (newIterations != this.iterations) {
            this.iterations = newIterations;
            this.iterationsChangeEmitter.emit(eventID);
        }
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
            'raid': this.raid.toJson(),
            'encounter': this.encounter.toJson(),
        };
    }
    // Set all the current values, assumes obj is the same type returned by toJson().
    fromJson(eventID, obj, spec) {
        TypedEvent.freezeAllAndDo(() => {
            // For legacy format. Do not remove this until 2022/01/05 (1 month).
            if (obj['sim']) {
                if (!obj['raid']) {
                    obj['raid'] = {
                        'parties': [
                            {
                                'players': [
                                    {
                                        'spec': spec,
                                        'player': obj['player'],
                                    },
                                ],
                                'buffs': obj['sim']['partyBuffs'],
                            },
                        ],
                        'buffs': obj['sim']['raidBuffs'],
                    };
                    obj['raid']['parties'][0]['players'][0]['player']['buffs'] = obj['sim']['individualBuffs'];
                }
            }
            if (obj['raid']) {
                this.raid.fromJson(eventID, obj['raid']);
            }
            if (obj['encounter']) {
                this.encounter.fromJson(eventID, obj['encounter']);
            }
        });
    }
}
