import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { ComputeStatsResult } from '/tbc/core/proto/api.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { makeComputeStatsRequest } from '/tbc/core/proto_utils/request_helpers.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { specTypeFunctions } from '/tbc/core/proto_utils/utils.js';
import { specToEligibleRaces } from '/tbc/core/proto_utils/utils.js';
import { getEligibleItemSlots } from '/tbc/core/proto_utils/utils.js';
import { gemMatchesSocket } from '/tbc/core/proto_utils/utils.js';
import { TypedEvent } from './typed_event.js';
import { sum } from './utils.js';
import { wait } from './utils.js';
// Manages all the gear / consumes / other settings for a single Player.
export class Player {
    constructor(config, sim) {
        this.phaseChangeEmitter = new TypedEvent();
        this.consumesChangeEmitter = new TypedEvent();
        this.customStatsChangeEmitter = new TypedEvent();
        this.gearChangeEmitter = new TypedEvent();
        this.raceChangeEmitter = new TypedEvent();
        this.rotationChangeEmitter = new TypedEvent();
        this.talentsChangeEmitter = new TypedEvent();
        // Talents dont have all fields so we need this
        this.talentsStringChangeEmitter = new TypedEvent();
        this.specOptionsChangeEmitter = new TypedEvent();
        // Emits when any of the above emitters emit.
        this.changeEmitter = new TypedEvent();
        this.gearListEmitter = new TypedEvent();
        this.characterStatsEmitter = new TypedEvent();
        this.sim = sim;
        this.spec = config.spec;
        this._race = specToEligibleRaces[this.spec][0];
        this.specTypeFunctions = specTypeFunctions[this.spec];
        this._metaGemEffectEP = config.metaGemEffectEP || (() => 0);
        this._consumes = config.defaults.consumes;
        this._customStats = new Stats();
        this._gear = new Gear({});
        this._rotation = config.defaults.rotation;
        this._talents = this.specTypeFunctions.talentsCreate();
        this._talentsString = config.defaults.talents;
        this._epWeights = config.defaults.epWeights;
        this._specOptions = config.defaults.specOptions;
        this.defaultGear = config.defaults.gear;
        [
            this.consumesChangeEmitter,
            this.customStatsChangeEmitter,
            this.gearChangeEmitter,
            this.raceChangeEmitter,
            this.rotationChangeEmitter,
            this.talentsChangeEmitter,
            this.talentsStringChangeEmitter,
            this.specOptionsChangeEmitter,
        ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));
        this._currentStats = ComputeStatsResult.create();
        this.sim.changeEmitter.on(() => {
            this.updateCharacterStats();
        });
        this.changeEmitter.on(() => {
            this.updateCharacterStats();
        });
    }
    async statWeights(request) {
        const result = await this.sim.statWeights(request);
        this._epWeights = new Stats(result.epValues);
        return result;
    }
    // This should be invoked internally whenever stats might have changed.
    async updateCharacterStats() {
        // Sometimes a ui change triggers other changes, so waiting a bit makes sure
        // we get all of them.
        await wait(10);
        const computeStatsResult = await this.sim.computeStats(makeComputeStatsRequest(this.sim.getBuffs(), this._consumes, this._customStats, this.sim.getEncounter(), this._gear, this._race, this._rotation, this._talents, this._specOptions));
        this._currentStats = computeStatsResult;
        this.characterStatsEmitter.emit();
    }
    getCurrentStats() {
        return ComputeStatsResult.clone(this._currentStats);
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
    getRotation() {
        return this.specTypeFunctions.rotationCopy(this._rotation);
    }
    setRotation(newRotation) {
        if (this.specTypeFunctions.rotationEquals(newRotation, this._rotation))
            return;
        this._rotation = this.specTypeFunctions.rotationCopy(newRotation);
        this.rotationChangeEmitter.emit();
    }
    getTalents() {
        return this.specTypeFunctions.talentsCopy(this._talents);
    }
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
        const enchants = this.sim.getEnchants(slot);
        if (enchants.length > 0) {
            ep += Math.max(...enchants.map(enchant => this.computeEnchantEP(enchant)));
        }
        // Compare whether its better to match sockets + get socket bonus, or just use best gems.
        const bestGemEPNotMatchingSockets = sum(item.gemSockets.map(socketColor => {
            const gems = this.sim.getGems(socketColor).filter(gem => !gem.unique && gem.phase <= this.sim.getPhase());
            if (gems.length > 0) {
                return Math.max(...gems.map(gem => this.computeGemEP(gem)));
            }
            else {
                return 0;
            }
        }));
        const bestGemEPMatchingSockets = sum(item.gemSockets.map(socketColor => {
            const gems = this.sim.getGems(socketColor).filter(gem => !gem.unique && gem.phase <= this.sim.getPhase() && gemMatchesSocket(gem, socketColor));
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
            'consumes': Consumes.toJson(this._consumes),
            'customStats': this._customStats.toJson(),
            'gear': EquipmentSpec.toJson(this._gear.asSpec()),
            'race': this._race,
            'rotation': this.specTypeFunctions.rotationToJson(this._rotation),
            'talents': this._talentsString,
            'specOptions': this.specTypeFunctions.optionsToJson(this._specOptions),
        };
    }
    // Set all the current values, assumes obj is the same type returned by toJson().
    fromJson(obj) {
        try {
            this.setConsumes(Consumes.fromJson(obj['consumes']));
        }
        catch (e) {
            console.warn('Failed to parse consumes: ' + e);
        }
        try {
            this.setCustomStats(Stats.fromJson(obj['customStats']));
        }
        catch (e) {
            console.warn('Failed to parse custom stats: ' + e);
        }
        try {
            this.setGear(this.sim.lookupEquipmentSpec(EquipmentSpec.fromJson(obj['gear'])));
        }
        catch (e) {
            console.warn('Failed to parse gear: ' + e);
        }
        try {
            this.setRace(obj['race']);
        }
        catch (e) {
            console.warn('Failed to parse race: ' + e);
        }
        try {
            this.setRotation(this.specTypeFunctions.rotationFromJson(obj['rotation']));
        }
        catch (e) {
            console.warn('Failed to parse rotation: ' + e);
        }
        try {
            this.setTalentsString(obj['talents']);
        }
        catch (e) {
            console.warn('Failed to parse talents: ' + e);
        }
        try {
            this.setSpecOptions(this.specTypeFunctions.optionsFromJson(obj['specOptions']));
        }
        catch (e) {
            console.warn('Failed to parse spec options: ' + e);
        }
    }
}
