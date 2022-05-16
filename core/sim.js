import { Item } from '/tbc/core/proto/common.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { ComputeStatsRequest } from '/tbc/core/proto/api.js';
import { GearListRequest } from '/tbc/core/proto/api.js';
import { RaidSimRequest } from '/tbc/core/proto/api.js';
import { SimOptions } from '/tbc/core/proto/api.js';
import { StatWeightsRequest, StatWeightsResult } from '/tbc/core/proto/api.js';
import { SimSettings as SimSettingsProto } from '/tbc/core/proto/ui.js';
import { EquippedItem } from '/tbc/core/proto_utils/equipped_item.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { SimResult } from '/tbc/core/proto_utils/sim_result.js';
import { gemEligibleForSocket } from '/tbc/core/proto_utils/gems.js';
import { gemMatchesSocket } from '/tbc/core/proto_utils/gems.js';
import { specTypeFunctions } from '/tbc/core/proto_utils/utils.js';
import { getEligibleItemSlots } from '/tbc/core/proto_utils/utils.js';
import { getEligibleEnchantSlots } from '/tbc/core/proto_utils/utils.js';
import { playerToSpec } from '/tbc/core/proto_utils/utils.js';
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
        this.fixedRngSeed = 0;
        this.show1hWeapons = true;
        this.show2hWeapons = true;
        this.showMatchingGems = true;
        this.showThreatMetrics = false;
        this.showExperimental = false;
        // Database
        this.items = {};
        this.enchants = {};
        this.gems = {};
        this.presetEncounters = {};
        this.presetTargets = {};
        this.iterationsChangeEmitter = new TypedEvent();
        this.phaseChangeEmitter = new TypedEvent();
        this.fixedRngSeedChangeEmitter = new TypedEvent();
        this.lastUsedRngSeedChangeEmitter = new TypedEvent();
        this.show1hWeaponsChangeEmitter = new TypedEvent();
        this.show2hWeaponsChangeEmitter = new TypedEvent();
        this.showMatchingGemsChangeEmitter = new TypedEvent();
        this.showThreatMetricsChangeEmitter = new TypedEvent();
        this.showExperimentalChangeEmitter = new TypedEvent();
        // Fires when a raid sim API call completes.
        this.simResultEmitter = new TypedEvent();
        this.lastUsedRngSeed = 0;
        // These callbacks are needed so we can apply BuffBot modifications automatically before sending requests.
        this.modifyRaidProto = () => { };
        this.workerPool = new WorkerPool(3);
        this._initPromise = this.workerPool.getGearList(GearListRequest.create()).then(result => {
            result.items.forEach(item => this.items[item.id] = item);
            result.enchants.forEach(enchant => this.enchants[enchant.id] = enchant);
            result.gems.forEach(gem => this.gems[gem.id] = gem);
            result.encounters.forEach(encounter => this.presetEncounters[encounter.path] = encounter);
            result.encounters.map(e => e.targets).flat().forEach(target => this.presetTargets[target.path] = target);
        });
        this.raid = new Raid(this);
        this.encounter = new Encounter(this);
        this.settingsChangeEmitter = TypedEvent.onAny([
            this.iterationsChangeEmitter,
            this.phaseChangeEmitter,
            this.fixedRngSeedChangeEmitter,
            this.show1hWeaponsChangeEmitter,
            this.show2hWeaponsChangeEmitter,
            this.showMatchingGemsChangeEmitter,
            this.showThreatMetricsChangeEmitter,
            this.showExperimentalChangeEmitter,
        ]);
        this.changeEmitter = TypedEvent.onAny([
            this.settingsChangeEmitter,
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
    makeRaidSimRequest(debug) {
        const raid = this.getModifiedRaidProto();
        const encounter = this.encounter.toProto();
        const hunters = raid.parties.map(party => party.players).flat().filter(player => player.name && playerToSpec(player) == Spec.SpecHunter);
        if (hunters.some(hunter => specTypeFunctions[Spec.SpecHunter].talentsFromPlayer(hunter).exposeWeakness > 0)) {
            if (raid.debuffs) {
                raid.debuffs.exposeWeaknessUptime = 0;
            }
        }
        return RaidSimRequest.create({
            raid: raid,
            encounter: encounter,
            simOptions: SimOptions.create({
                iterations: debug ? 1 : this.getIterations(),
                randomSeed: BigInt(this.nextRngSeed()),
                debugFirstIteration: true,
            }),
        });
    }
    async runRaidSim(eventID, onProgress) {
        if (this.raid.isEmpty()) {
            throw new Error('Raid is empty! Try adding some players first.');
        }
        else if (this.encounter.getNumTargets() < 1) {
            throw new Error('Encounter has no targets! Try adding some targets first.');
        }
        await this.waitForInit();
        const request = this.makeRaidSimRequest(false);
        var result = await this.workerPool.raidSimAsync(request, onProgress);
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
        const result = await this.workerPool.raidSimAsync(request, () => { });
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
        eventID = TypedEvent.nextEventID();
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
    async statWeights(player, epStats, epReferenceStat, onProgress) {
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
            const tanks = this.raid.getTanks().map(tank => tank.targetIndex).includes(player.getRaidIndex())
                ? [RaidTarget.create({ targetIndex: 0 })]
                : [];
            const request = StatWeightsRequest.create({
                player: player.toProto(),
                raidBuffs: this.raid.getBuffs(),
                partyBuffs: player.getParty().getBuffs(),
                debuffs: this.raid.getDebuffs(),
                encounter: this.encounter.toProto(),
                simOptions: SimOptions.create({
                    iterations: this.getIterations(),
                    randomSeed: BigInt(this.nextRngSeed()),
                    debug: false,
                }),
                tanks: tanks,
                statsToWeigh: epStats,
                epReferenceStat: epReferenceStat,
            });
            var result = await this.workerPool.statWeightsAsync(request, onProgress);
            return result;
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
    // ID can be the formula ID OR the effect ID.
    getEnchantFlexible(id) {
        return Object.values(this.enchants).find(enchant => enchant.id == id || enchant.effectId == id) || null;
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
    getPresetEncounter(path) {
        return this.presetEncounters[path] || null;
    }
    getPresetTarget(path) {
        return this.presetTargets[path] || null;
    }
    getAllPresetEncounters() {
        return Object.values(this.presetEncounters);
    }
    getAllPresetTargets() {
        return Object.values(this.presetTargets);
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
    getFixedRngSeed() {
        return this.fixedRngSeed;
    }
    setFixedRngSeed(eventID, newFixedRngSeed) {
        if (newFixedRngSeed != this.fixedRngSeed) {
            this.fixedRngSeed = newFixedRngSeed;
            this.fixedRngSeedChangeEmitter.emit(eventID);
        }
    }
    nextRngSeed() {
        let rngSeed = 0;
        if (this.fixedRngSeed) {
            rngSeed = this.fixedRngSeed;
        }
        else {
            rngSeed = Math.floor(Math.random() * Sim.MAX_RNG_SEED);
        }
        this.lastUsedRngSeed = rngSeed;
        this.lastUsedRngSeedChangeEmitter.emit(TypedEvent.nextEventID());
        return rngSeed;
    }
    getLastUsedRngSeed() {
        return this.lastUsedRngSeed;
    }
    getShow1hWeapons() {
        return this.show1hWeapons;
    }
    setShow1hWeapons(eventID, newShow1hWeapons) {
        if (newShow1hWeapons != this.show1hWeapons) {
            this.show1hWeapons = newShow1hWeapons;
            this.show1hWeaponsChangeEmitter.emit(eventID);
        }
    }
    getShow2hWeapons() {
        return this.show2hWeapons;
    }
    setShow2hWeapons(eventID, newShow2hWeapons) {
        if (newShow2hWeapons != this.show2hWeapons) {
            this.show2hWeapons = newShow2hWeapons;
            this.show2hWeaponsChangeEmitter.emit(eventID);
        }
    }
    getShowMatchingGems() {
        return this.showMatchingGems;
    }
    setShowMatchingGems(eventID, newShowMatchingGems) {
        if (newShowMatchingGems != this.showMatchingGems) {
            this.showMatchingGems = newShowMatchingGems;
            this.showMatchingGemsChangeEmitter.emit(eventID);
        }
    }
    getShowThreatMetrics() {
        return this.showThreatMetrics;
    }
    setShowThreatMetrics(eventID, newShowThreatMetrics) {
        if (newShowThreatMetrics != this.showThreatMetrics) {
            this.showThreatMetrics = newShowThreatMetrics;
            this.showThreatMetricsChangeEmitter.emit(eventID);
        }
    }
    getShowExperimental() {
        return this.showExperimental;
    }
    setShowExperimental(eventID, newShowExperimental) {
        if (newShowExperimental != this.showExperimental) {
            this.showExperimental = newShowExperimental;
            this.showExperimentalChangeEmitter.emit(eventID);
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
    toProto() {
        return SimSettingsProto.create({
            iterations: this.getIterations(),
            phase: this.getPhase(),
            fixedRngSeed: BigInt(this.getFixedRngSeed()),
            showThreatMetrics: this.getShowThreatMetrics(),
            showExperimental: this.getShowExperimental(),
        });
    }
    fromProto(eventID, proto) {
        TypedEvent.freezeAllAndDo(() => {
            this.setIterations(eventID, proto.iterations || 3000);
            this.setPhase(eventID, proto.phase || OtherConstants.CURRENT_PHASE);
            this.setFixedRngSeed(eventID, Number(proto.fixedRngSeed));
            this.setShowThreatMetrics(eventID, proto.showThreatMetrics);
            this.setShowExperimental(eventID, proto.showExperimental);
        });
    }
    applyDefaults(eventID, isTankSim) {
        this.fromProto(eventID, SimSettingsProto.create({
            iterations: 3000,
            phase: OtherConstants.CURRENT_PHASE,
            showThreatMetrics: isTankSim,
        }));
    }
}
Sim.MAX_RNG_SEED = Math.pow(2, 32) - 1;
