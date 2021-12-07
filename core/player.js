import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Player as PlayerProto } from '/tbc/core/proto/api.js';
import { ComputeStatsRequest, ComputeStatsResult } from '/tbc/core/proto/api.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { canEquipItem, getEligibleItemSlots, getMetaGemEffectEP, gemMatchesSocket, raceToFaction, specToClass, specToEligibleRaces, specTypeFunctions, withSpecProto, } from '/tbc/core/proto_utils/utils.js';
import { TypedEvent } from './typed_event.js';
import { MAX_PARTY_SIZE } from './party.js';
import { sum } from './utils.js';
import { wait } from './utils.js';
// Manages all the gear / consumes / other settings for a single Player.
export class Player {
    constructor(spec, sim) {
        this.name = '';
        this.buffs = IndividualBuffs.create();
        this.consumes = Consumes.create();
        this.bonusStats = new Stats();
        this.gear = new Gear({});
        this.talentsString = '';
        this.epWeights = new Stats();
        this.nameChangeEmitter = new TypedEvent();
        this.buffsChangeEmitter = new TypedEvent();
        this.consumesChangeEmitter = new TypedEvent();
        this.bonusStatsChangeEmitter = new TypedEvent();
        this.gearChangeEmitter = new TypedEvent();
        this.raceChangeEmitter = new TypedEvent();
        this.rotationChangeEmitter = new TypedEvent();
        this.talentsChangeEmitter = new TypedEvent();
        // Talents dont have all fields so we need this.
        this.talentsStringChangeEmitter = new TypedEvent();
        this.specOptionsChangeEmitter = new TypedEvent();
        this.currentStatsEmitter = new TypedEvent();
        // Emits when any of the above emitters emit.
        this.changeEmitter = new TypedEvent();
        this.sim = sim;
        this.party = null;
        this.raid = null;
        this.spec = spec;
        this.race = specToEligibleRaces[this.spec][0];
        this.specTypeFunctions = specTypeFunctions[this.spec];
        this.rotation = this.specTypeFunctions.rotationCreate();
        this.talents = this.specTypeFunctions.talentsCreate();
        this.specOptions = this.specTypeFunctions.optionsCreate();
        [
            this.nameChangeEmitter,
            this.buffsChangeEmitter,
            this.consumesChangeEmitter,
            this.bonusStatsChangeEmitter,
            this.gearChangeEmitter,
            this.raceChangeEmitter,
            this.rotationChangeEmitter,
            this.talentsChangeEmitter,
            this.talentsStringChangeEmitter,
            this.specOptionsChangeEmitter,
        ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));
        this.currentStats = ComputeStatsResult.create();
        this.sim.changeEmitter.on(() => {
            this.updateCharacterStats();
        });
        this.changeEmitter.on(() => {
            this.updateCharacterStats();
        });
    }
    getParty() {
        return this.party;
    }
    getRaid() {
        return this.raid;
    }
    // Returns this player's index within its party [0-4].
    getPartyIndex() {
        if (this.party == null) {
            throw new Error('Can\'t get party index for player without a party!');
        }
        return this.party.getPlayers().indexOf(this);
    }
    // Returns this player's index within its raid [0-24].
    getRaidIndex() {
        if (this.party == null) {
            throw new Error('Can\'t get raid index for player without a party!');
        }
        return this.party.getIndex() * MAX_PARTY_SIZE + this.getPartyIndex();
    }
    setParty(newParty) {
        if (this.party == newParty) {
            return;
        }
        // Remove player from its old party if there is one.
        if (this.party != null) {
            this.party.setPlayer(this.getPartyIndex(), null);
        }
        if (newParty == null) {
            this.party = null;
            this.raid = null;
        }
        else {
            this.party = newParty;
            this.raid = newParty.raid;
        }
    }
    // Returns all items that this player can wear in the given slot.
    getItems(slot) {
        return this.sim.getItems(slot).filter(item => canEquipItem(item, this.spec));
    }
    // Returns all enchants that this player can wear in the given slot.
    getEnchants(slot) {
        return this.sim.getEnchants(slot);
    }
    // Returns all gems that this player can wear of the given color.
    getGems(socketColor) {
        return this.sim.getGems(socketColor);
    }
    getEpWeights() {
        return this.epWeights;
    }
    setEpWeights(newEpWeights) {
        this.epWeights = newEpWeights;
    }
    async statWeights(request) {
        const result = await this.sim.statWeights(request);
        this.epWeights = new Stats(result.epValues);
        return result;
    }
    // This should be invoked internally whenever stats might have changed.
    async updateCharacterStats() {
        // Sometimes a ui change triggers other changes, so waiting a bit makes sure
        // we get all of them.
        await wait(10);
        const computeStatsResult = await this.sim.computeStats(ComputeStatsRequest.create({
            player: this.toProto(),
            raidBuffs: this.raid.getBuffs(),
            partyBuffs: this.party.getBuffs(),
        }));
        this.currentStats = computeStatsResult;
        this.currentStatsEmitter.emit();
    }
    getCurrentStats() {
        return ComputeStatsResult.clone(this.currentStats);
    }
    getName() {
        return this.name;
    }
    setName(newName) {
        if (newName != this.name) {
            this.name = newName;
            this.nameChangeEmitter.emit();
        }
    }
    getRace() {
        return this.race;
    }
    setRace(newRace) {
        if (newRace != this.race) {
            this.race = newRace;
            this.raceChangeEmitter.emit();
        }
    }
    getFaction() {
        return raceToFaction[this.getRace()];
    }
    getBuffs() {
        // Make a defensive copy
        return IndividualBuffs.clone(this.buffs);
    }
    setBuffs(newBuffs) {
        if (IndividualBuffs.equals(this.buffs, newBuffs))
            return;
        // Make a defensive copy
        this.buffs = IndividualBuffs.clone(newBuffs);
        this.buffsChangeEmitter.emit();
    }
    getConsumes() {
        // Make a defensive copy
        return Consumes.clone(this.consumes);
    }
    setConsumes(newConsumes) {
        if (Consumes.equals(this.consumes, newConsumes))
            return;
        // Make a defensive copy
        this.consumes = Consumes.clone(newConsumes);
        this.consumesChangeEmitter.emit();
    }
    equipItem(slot, newItem) {
        const newGear = this.gear.withEquippedItem(slot, newItem);
        if (newGear.equals(this.gear))
            return;
        this.gear = newGear;
        this.gearChangeEmitter.emit();
    }
    getEquippedItem(slot) {
        return this.gear.getEquippedItem(slot);
    }
    getGear() {
        return this.gear;
    }
    setGear(newGear) {
        if (newGear.equals(this.gear))
            return;
        this.gear = newGear;
        this.gearChangeEmitter.emit();
    }
    getBonusStats() {
        return this.bonusStats;
    }
    setBonusStats(newBonusStats) {
        if (newBonusStats.equals(this.bonusStats))
            return;
        this.bonusStats = newBonusStats;
        this.bonusStatsChangeEmitter.emit();
    }
    getRotation() {
        return this.specTypeFunctions.rotationCopy(this.rotation);
    }
    setRotation(newRotation) {
        if (this.specTypeFunctions.rotationEquals(newRotation, this.rotation))
            return;
        this.rotation = this.specTypeFunctions.rotationCopy(newRotation);
        this.rotationChangeEmitter.emit();
    }
    getTalents() {
        return this.specTypeFunctions.talentsCopy(this.talents);
    }
    setTalents(newTalents) {
        if (this.specTypeFunctions.talentsEquals(newTalents, this.talents))
            return;
        this.talents = this.specTypeFunctions.talentsCopy(newTalents);
        this.talentsChangeEmitter.emit();
    }
    getTalentsString() {
        return this.talentsString;
    }
    setTalentsString(newTalentsString) {
        if (newTalentsString == this.talentsString)
            return;
        this.talentsString = newTalentsString;
        this.talentsStringChangeEmitter.emit();
    }
    getSpecOptions() {
        return this.specTypeFunctions.optionsCopy(this.specOptions);
    }
    setSpecOptions(newSpecOptions) {
        if (this.specTypeFunctions.optionsEquals(newSpecOptions, this.specOptions))
            return;
        this.specOptions = this.specTypeFunctions.optionsCopy(newSpecOptions);
        this.specOptionsChangeEmitter.emit();
    }
    computeGemEP(gem) {
        const epFromStats = new Stats(gem.stats).computeEP(this.epWeights);
        const epFromEffect = getMetaGemEffectEP(this.spec, gem, new Stats(this.currentStats.finalStats));
        return epFromStats + epFromEffect;
    }
    computeEnchantEP(enchant) {
        return new Stats(enchant.stats).computeEP(this.epWeights);
    }
    computeItemEP(item) {
        if (item == null)
            return 0;
        let ep = new Stats(item.stats).computeEP(this.epWeights);
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
        })) + new Stats(item.socketBonus).computeEP(this.epWeights);
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
        parts.push('pcs=' + this.gear.asArray().filter(ei => ei != null).map(ei => ei.item.id).join(':'));
        elem.setAttribute('data-wowhead', parts.join('&'));
    }
    toProto() {
        return withSpecProto(PlayerProto.create({
            race: this.getRace(),
            class: specToClass[this.spec],
            equipment: this.getGear().asSpec(),
            consumes: this.getConsumes(),
            bonusStats: this.getBonusStats().asArray(),
            buffs: this.getBuffs(),
        }), this.getRotation(), this.getTalents(), this.getSpecOptions());
    }
    // TODO: Remove to/from json functions and use proto versions instead. This will require
    // some way to store all talents in the proto.
    // Returns JSON representing all the current values.
    toJson() {
        return {
            'name': this.name,
            'buffs': IndividualBuffs.toJson(this.buffs),
            'consumes': Consumes.toJson(this.consumes),
            'bonusStats': this.bonusStats.toJson(),
            'gear': EquipmentSpec.toJson(this.gear.asSpec()),
            'race': this.race,
            'rotation': this.specTypeFunctions.rotationToJson(this.rotation),
            'talents': this.talentsString,
            'specOptions': this.specTypeFunctions.optionsToJson(this.specOptions),
        };
    }
    // Set all the current values, assumes obj is the same type returned by toJson().
    fromJson(obj) {
        try {
            if (obj['name']) {
                this.setName(obj['name']);
            }
        }
        catch (e) {
            console.warn('Failed to parse name: ' + e);
        }
        try {
            this.setBuffs(IndividualBuffs.fromJson(obj['buffs']));
        }
        catch (e) {
            console.warn('Failed to parse player buffs: ' + e);
        }
        try {
            this.setConsumes(Consumes.fromJson(obj['consumes']));
        }
        catch (e) {
            console.warn('Failed to parse consumes: ' + e);
        }
        // For legacy format. Do not remove this until 2022/01/02 (1 month).
        if (obj['customStats']) {
            obj['bonusStats'] = obj['customStats'];
        }
        try {
            this.setBonusStats(Stats.fromJson(obj['bonusStats']));
        }
        catch (e) {
            console.warn('Failed to parse bonus stats: ' + e);
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
    clone() {
        const newPlayer = new Player(this.spec, this.sim);
        newPlayer.fromJson(this.toJson());
        return newPlayer;
    }
}
