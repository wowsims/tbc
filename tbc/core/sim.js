import { Buffs } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Item } from '/tbc/core/proto/common.js';
import { ComputeStatsResult } from '/tbc/core/proto/api.js';
import { GearListRequest } from '/tbc/core/proto/api.js';
import { EquippedItem } from '/tbc/core/proto_utils/equipped_item.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { makeComputeStatsRequest } from '/tbc/core/proto_utils/request_helpers.js';
import { makeIndividualSimRequest } from '/tbc/core/proto_utils/request_helpers.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { specTypeFunctions } from '/tbc/core/proto_utils/utils.js';
import { specToEligibleRaces } from '/tbc/core/proto_utils/utils.js';
import { getEligibleItemSlots } from '/tbc/core/proto_utils/utils.js';
import { getEligibleEnchantSlots } from '/tbc/core/proto_utils/utils.js';
import { gemEligibleForSocket } from '/tbc/core/proto_utils/utils.js';
import { gemMatchesSocket } from '/tbc/core/proto_utils/utils.js';
import { TypedEvent } from './typed_event.js';
import { sum } from './utils.js';
import { wait } from './utils.js';
import { WorkerPool } from './worker_pool.js';
// Core Sim module which deals only with api types, no UI-related stuff.
export class Sim extends WorkerPool {
    constructor(config) {
        super(3);
        this.phaseChangeEmitter = new TypedEvent();
        this.buffsChangeEmitter = new TypedEvent();
        this.consumesChangeEmitter = new TypedEvent();
        this.customStatsChangeEmitter = new TypedEvent();
        this.encounterChangeEmitter = new TypedEvent();
        this.gearChangeEmitter = new TypedEvent();
        this.raceChangeEmitter = new TypedEvent();
        this.agentChangeEmitter = new TypedEvent();
        this.talentsChangeEmitter = new TypedEvent();
        // Talents dont have all fields so we need this
        this.talentsStringChangeEmitter = new TypedEvent();
        this.specOptionsChangeEmitter = new TypedEvent();
        // Emits when any of the above emitters emit.
        this.changeEmitter = new TypedEvent();
        this.gearListEmitter = new TypedEvent();
        this.characterStatsEmitter = new TypedEvent();
        // Database
        this._items = {};
        this._enchants = {};
        this._gems = {};
        this._init = false;
        this.spec = config.spec;
        this._race = specToEligibleRaces[this.spec][0];
        this.specTypeFunctions = specTypeFunctions[this.spec];
        this._metaGemEffectEP = config.metaGemEffectEP || (() => 0);
        this._phase = config.defaults.phase;
        this._buffs = config.defaults.buffs;
        this._consumes = config.defaults.consumes;
        this._customStats = new Stats();
        this._encounter = config.defaults.encounter;
        this._gear = new Gear({});
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
        this._currentStats = ComputeStatsResult.create();
        this.changeEmitter.on(() => {
            this.updateCharacterStats();
        });
    }
    async init() {
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
    async statWeights(request) {
        const result = await super.statWeights(request);
        this._epWeights = new Stats(result.epValues);
        return result;
    }
    // This should be invoked internally whenever stats might have changed.
    async updateCharacterStats() {
        // Sometimes a ui change triggers other changes, so waiting a bit makes sure
        // we get all of them.
        await wait(10);
        const computeStatsResult = await this.computeStats(makeComputeStatsRequest(this._buffs, this._consumes, this._customStats, this._encounter, this._gear, this._race, this._agent, this._talents, this._specOptions));
        this._currentStats = computeStatsResult;
        this.characterStatsEmitter.emit();
    }
    getCurrentStats() {
        return ComputeStatsResult.clone(this._currentStats);
    }
    getItems(slot) {
        let items = Object.values(this._items);
        if (slot != undefined) {
            items = items.filter(item => getEligibleItemSlots(item).includes(slot));
        }
        return items;
    }
    getEnchants(slot) {
        let enchants = Object.values(this._enchants);
        if (slot != undefined) {
            enchants = enchants.filter(enchant => getEligibleEnchantSlots(enchant).includes(slot));
        }
        return enchants;
    }
    getGems(socketColor) {
        let gems = Object.values(this._gems);
        if (socketColor) {
            gems = gems.filter(gem => gemEligibleForSocket(gem, socketColor));
        }
        return gems;
    }
    getMatchingGems(socketColor) {
        return Object.values(this._gems).filter(gem => gemMatchesSocket(gem, socketColor));
    }
    getPhase() {
        return this._phase;
    }
    setPhase(newPhase) {
        if (newPhase != this._phase) {
            this._phase = newPhase;
            this.phaseChangeEmitter.emit();
        }
    }
    getRace() {
        return this._race;
    }
    setRace(newRace) {
        if (newRace != this._race) {
            this._race = newRace;
            this.raceChangeEmitter.emit();
        }
    }
    getBuffs() {
        // Make a defensive copy
        return Buffs.clone(this._buffs);
    }
    setBuffs(newBuffs) {
        if (Buffs.equals(this._buffs, newBuffs))
            return;
        // Make a defensive copy
        this._buffs = Buffs.clone(newBuffs);
        this.buffsChangeEmitter.emit();
    }
    getConsumes() {
        // Make a defensive copy
        return Consumes.clone(this._consumes);
    }
    setConsumes(newConsumes) {
        if (Consumes.equals(this._consumes, newConsumes))
            return;
        // Make a defensive copy
        this._consumes = Consumes.clone(newConsumes);
        this.consumesChangeEmitter.emit();
    }
    getEncounter() {
        // Make a defensive copy
        return Encounter.clone(this._encounter);
    }
    setEncounter(newEncounter) {
        if (Encounter.equals(this._encounter, newEncounter))
            return;
        // Make a defensive copy
        this._encounter = Encounter.clone(newEncounter);
        this.encounterChangeEmitter.emit();
    }
    equipItem(slot, newItem) {
        const newGear = this._gear.withEquippedItem(slot, newItem);
        if (newGear.equals(this._gear))
            return;
        this._gear = newGear;
        this.gearChangeEmitter.emit();
    }
    getEquippedItem(slot) {
        return this._gear.getEquippedItem(slot);
    }
    getGear() {
        return this._gear;
    }
    setGear(newGear) {
        if (newGear.equals(this._gear))
            return;
        this._gear = newGear;
        this.gearChangeEmitter.emit();
    }
    getCustomStats() {
        return this._customStats;
    }
    setCustomStats(newCustomStats) {
        if (newCustomStats.equals(this._customStats))
            return;
        this._customStats = newCustomStats;
        this.customStatsChangeEmitter.emit();
    }
    getAgent() {
        return this.specTypeFunctions.agentCopy(this._agent);
    }
    setAgent(newAgent) {
        if (this.specTypeFunctions.agentEquals(newAgent, this._agent))
            return;
        this._agent = this.specTypeFunctions.agentCopy(newAgent);
        this.agentChangeEmitter.emit();
    }
    // Commented because this should NOT be used; all external uses should be able to use getTalentsString()
    //getTalents(): SpecTalents<SpecType> {
    //  return this.specTypeFunctions.talentsCopy(this._talents);
    //}
    setTalents(newTalents) {
        if (this.specTypeFunctions.talentsEquals(newTalents, this._talents))
            return;
        this._talents = this.specTypeFunctions.talentsCopy(newTalents);
        this.talentsChangeEmitter.emit();
    }
    getTalentsString() {
        return this._talentsString;
    }
    setTalentsString(newTalentsString) {
        if (newTalentsString == this._talentsString)
            return;
        this._talentsString = newTalentsString;
        this.talentsStringChangeEmitter.emit();
    }
    getSpecOptions() {
        return this.specTypeFunctions.optionsCopy(this._specOptions);
    }
    setSpecOptions(newSpecOptions) {
        if (this.specTypeFunctions.optionsEquals(newSpecOptions, this._specOptions))
            return;
        this._specOptions = this.specTypeFunctions.optionsCopy(newSpecOptions);
        this.specOptionsChangeEmitter.emit();
    }
    lookupItemSpec(itemSpec) {
        const item = this._items[itemSpec.id];
        if (!item)
            return null;
        const enchant = this._enchants[itemSpec.enchant] || null;
        const gems = itemSpec.gems.map(gemId => this._gems[gemId] || null);
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
    computeGemEP(gem) {
        const epFromStats = new Stats(gem.stats).computeEP(this._epWeights);
        const epFromEffect = this._metaGemEffectEP(gem, this);
        return epFromStats + epFromEffect;
    }
    computeEnchantEP(enchant) {
        return new Stats(enchant.stats).computeEP(this._epWeights);
    }
    computeItemEP(item) {
        if (item == null)
            return 0;
        let ep = new Stats(item.stats).computeEP(this._epWeights);
        const slot = getEligibleItemSlots(item)[0];
        const enchants = this.getEnchants(slot);
        if (enchants.length > 0) {
            ep += Math.max(...enchants.map(enchant => this.computeEnchantEP(enchant)));
        }
        // Compare whether its better to match sockets + get socket bonus, or just use best gems.
        const bestGemEPNotMatchingSockets = sum(item.gemSockets.map(socketColor => {
            const gems = this.getGems(socketColor).filter(gem => !gem.unique && gem.phase <= this.getPhase());
            if (gems.length > 0) {
                return Math.max(...gems.map(gem => this.computeGemEP(gem)));
            }
            else {
                return 0;
            }
        }));
        const bestGemEPMatchingSockets = sum(item.gemSockets.map(socketColor => {
            const gems = this.getGems(socketColor).filter(gem => !gem.unique && gem.phase <= this.getPhase() && gemMatchesSocket(gem, socketColor));
            if (gems.length > 0) {
                return Math.max(...gems.map(gem => this.computeGemEP(gem)));
            }
            else {
                return 0;
            }
        })) + new Stats(item.socketBonus).computeEP(this._epWeights);
        ep += Math.max(bestGemEPMatchingSockets, bestGemEPNotMatchingSockets);
        return ep;
    }
    makeCurrentIndividualSimRequest(iterations, debug) {
        return makeIndividualSimRequest(this._buffs, this._consumes, this._customStats, this._encounter, this._gear, this._race, this._agent, this._talents, this._specOptions, iterations, debug);
    }
    setWowheadData(equippedItem, elem) {
        let parts = [];
        if (equippedItem.gems.length > 0) {
            parts.push('gems=' + equippedItem.gems.map(gem => gem ? gem.id : 0).join(':'));
        }
        if (equippedItem.enchant != null) {
            parts.push('ench=' + equippedItem.enchant.effectId);
        }
        parts.push('pcs=' + this._gear.asArray().filter(ei => ei != null).map(ei => ei.item.id).join(':'));
        elem.setAttribute('data-wowhead', parts.join('&'));
    }
    // Returns JSON representing all the current values.
    toJson() {
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
    fromJson(obj) {
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
