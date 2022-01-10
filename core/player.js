import { Cooldowns } from '/tbc/core/proto/common.js';
import { Conjured } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { PlayerStats } from '/tbc/core/proto/api.js';
import { Player as PlayerProto } from '/tbc/core/proto/api.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { gemMatchesSocket, } from '/tbc/core/proto_utils/gems.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { canEquipItem, classColors, getEligibleItemSlots, getTalentTreeIcon, getMetaGemEffectEP, raceToFaction, specToClass, specToEligibleRaces, specTypeFunctions, withSpecProto, } from '/tbc/core/proto_utils/utils.js';
import { TypedEvent } from './typed_event.js';
import { MAX_PARTY_SIZE } from './party.js';
import { sum } from './utils.js';
// Manages all the gear / consumes / other settings for a single Player.
export class Player {
    constructor(spec, sim) {
        this.name = '';
        this.buffs = IndividualBuffs.create();
        this.consumes = Consumes.create();
        this.bonusStats = new Stats();
        this.gear = new Gear({});
        this.talentsString = '';
        this.cooldowns = Cooldowns.create();
        this.epWeights = new Stats();
        this.currentStats = PlayerStats.create();
        this.nameChangeEmitter = new TypedEvent('PlayerName');
        this.buffsChangeEmitter = new TypedEvent('PlayerBuffs');
        this.consumesChangeEmitter = new TypedEvent('PlayerConsumes');
        this.bonusStatsChangeEmitter = new TypedEvent('PlayerBonusStats');
        this.gearChangeEmitter = new TypedEvent('PlayerGear');
        this.raceChangeEmitter = new TypedEvent('PlayerRace');
        this.rotationChangeEmitter = new TypedEvent('PlayerRotation');
        this.talentsChangeEmitter = new TypedEvent('PlayerTalents');
        // Talents dont have all fields so we need this.
        this.talentsStringChangeEmitter = new TypedEvent('PlayerTalentsString');
        this.specOptionsChangeEmitter = new TypedEvent('PlayerSpecOptions');
        this.cooldownsChangeEmitter = new TypedEvent('PlayerCooldowns');
        this.currentStatsEmitter = new TypedEvent('PlayerCurrentStats');
        this.sim = sim;
        this.party = null;
        this.raid = null;
        this.spec = spec;
        this.race = specToEligibleRaces[this.spec][0];
        this.specTypeFunctions = specTypeFunctions[this.spec];
        this.rotation = this.specTypeFunctions.rotationCreate();
        this.talents = this.specTypeFunctions.talentsCreate();
        this.specOptions = this.specTypeFunctions.optionsCreate();
        this.changeEmitter = TypedEvent.onAny([
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
            this.cooldownsChangeEmitter,
        ], 'PlayerChange');
    }
    getSpecIcon() {
        return getTalentTreeIcon(this.spec, this.getTalentsString());
    }
    getClass() {
        return specToClass[this.spec];
    }
    getClassColor() {
        return classColors[this.getClass()];
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
    // This should only ever be called from party.
    setParty(newParty) {
        if (newParty == null) {
            this.party = null;
            this.raid = null;
        }
        else {
            this.party = newParty;
            this.raid = newParty.raid;
        }
    }
    getOtherPartyMembers() {
        if (this.party == null) {
            return [];
        }
        return this.party.getPlayers().filter(player => player != null && player != this);
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
    async computeStatWeights(epStats, epReferenceStat) {
        const result = await this.sim.statWeights(this, epStats, epReferenceStat);
        this.epWeights = new Stats(result.epValues);
        return result;
    }
    getCurrentStats() {
        return PlayerStats.clone(this.currentStats);
    }
    setCurrentStats(eventID, newStats) {
        this.currentStats = newStats;
        this.currentStatsEmitter.emit(eventID);
        //// Remove item cooldowns if there is no cooldown available for the item.
        //const availableCooldowns = this.currentStats.cooldowns;
        //const newCooldowns = this.getCooldowns();
        //newCooldowns.cooldowns = newCooldowns.cooldowns.filter(cd => {
        //	if (cd.id && 'itemId' in cd.id.rawId) {
        //		return availableCooldowns.find(acd => ActionIdProto.equals(acd, cd.id)) != null;
        //	} else {
        //		return true;
        //	}
        //});
        //// TODO: Reference the parent event ID
        //this.setCooldowns(TypedEvent.nextEventID(), newCooldowns);
    }
    getName() {
        return this.name;
    }
    setName(eventID, newName) {
        if (newName != this.name) {
            this.name = newName;
            this.nameChangeEmitter.emit(eventID);
        }
    }
    getLabel() {
        if (this.party) {
            return `${this.name} (#${this.getRaidIndex() + 1})`;
        }
        else {
            return this.name;
        }
    }
    getRace() {
        return this.race;
    }
    setRace(eventID, newRace) {
        if (newRace != this.race) {
            this.race = newRace;
            this.raceChangeEmitter.emit(eventID);
        }
    }
    getFaction() {
        return raceToFaction[this.getRace()];
    }
    getBuffs() {
        // Make a defensive copy
        return IndividualBuffs.clone(this.buffs);
    }
    setBuffs(eventID, newBuffs) {
        if (IndividualBuffs.equals(this.buffs, newBuffs))
            return;
        // Make a defensive copy
        this.buffs = IndividualBuffs.clone(newBuffs);
        this.buffsChangeEmitter.emit(eventID);
    }
    getConsumes() {
        // Make a defensive copy
        return Consumes.clone(this.consumes);
    }
    setConsumes(eventID, newConsumes) {
        if (Consumes.equals(this.consumes, newConsumes))
            return;
        // Make a defensive copy
        this.consumes = Consumes.clone(newConsumes);
        this.consumesChangeEmitter.emit(eventID);
    }
    getCooldowns() {
        // Make a defensive copy
        return Cooldowns.clone(this.cooldowns);
    }
    setCooldowns(eventID, newCooldowns) {
        if (Cooldowns.equals(this.cooldowns, newCooldowns))
            return;
        // Make a defensive copy
        this.cooldowns = Cooldowns.clone(newCooldowns);
        this.cooldownsChangeEmitter.emit(eventID);
    }
    equipItem(eventID, slot, newItem) {
        this.setGear(eventID, this.gear.withEquippedItem(slot, newItem));
    }
    getEquippedItem(slot) {
        return this.gear.getEquippedItem(slot);
    }
    getGear() {
        return this.gear;
    }
    setGear(eventID, newGear) {
        if (newGear.equals(this.gear))
            return;
        // Commented for now because the UI for this is weird.
        //// If trinkets have changed and there were cooldowns assigned for those trinkets,
        //// try to match them up and switch to the new trinkets.
        //const newCooldowns = this.getCooldowns();
        //const oldTrinketIds = this.gear.getTrinkets().map(trinket => trinket?.asActionIdProto() || ActionIdProto.create());
        //const newTrinketIds = newGear.getTrinkets().map(trinket => trinket?.asActionIdProto() || ActionIdProto.create());
        //for (let i = 0; i < 2; i++) {
        //	const oldTrinketId = oldTrinketIds[i];
        //	const newTrinketId = newTrinketIds[i];
        //	if (ActionIdProto.equals(oldTrinketId, ActionIdProto.create())) {
        //		continue;
        //	}
        //	if (ActionIdProto.equals(newTrinketId, ActionIdProto.create())) {
        //		continue;
        //	}
        //	if (ActionIdProto.equals(oldTrinketId, newTrinketId)) {
        //		continue;
        //	}
        //	newCooldowns.cooldowns.forEach(cd => {
        //		if (ActionIdProto.equals(cd.id, oldTrinketId)) {
        //			cd.id = newTrinketId;
        //		}
        //	});
        //}
        TypedEvent.freezeAllAndDo(() => {
            this.gear = newGear;
            this.gearChangeEmitter.emit(eventID);
            //this.setCooldowns(eventID, newCooldowns);
        });
    }
    getBonusStats() {
        return this.bonusStats;
    }
    setBonusStats(eventID, newBonusStats) {
        if (newBonusStats.equals(this.bonusStats))
            return;
        this.bonusStats = newBonusStats;
        this.bonusStatsChangeEmitter.emit(eventID);
    }
    getRotation() {
        return this.specTypeFunctions.rotationCopy(this.rotation);
    }
    setRotation(eventID, newRotation) {
        if (this.specTypeFunctions.rotationEquals(newRotation, this.rotation))
            return;
        this.rotation = this.specTypeFunctions.rotationCopy(newRotation);
        this.rotationChangeEmitter.emit(eventID);
    }
    getTalents() {
        return this.specTypeFunctions.talentsCopy(this.talents);
    }
    setTalents(eventID, newTalents) {
        if (this.specTypeFunctions.talentsEquals(newTalents, this.talents))
            return;
        this.talents = this.specTypeFunctions.talentsCopy(newTalents);
        this.talentsChangeEmitter.emit(eventID);
    }
    getTalentsString() {
        return this.talentsString;
    }
    setTalentsString(eventID, newTalentsString) {
        if (newTalentsString == this.talentsString)
            return;
        this.talentsString = newTalentsString;
        this.talentsStringChangeEmitter.emit(eventID);
    }
    getTalentTreeIcon() {
        return getTalentTreeIcon(this.spec, this.getTalentsString());
    }
    getSpecOptions() {
        return this.specTypeFunctions.optionsCopy(this.specOptions);
    }
    setSpecOptions(eventID, newSpecOptions) {
        if (this.specTypeFunctions.optionsEquals(newSpecOptions, this.specOptions))
            return;
        this.specOptions = this.specTypeFunctions.optionsCopy(newSpecOptions);
        this.specOptionsChangeEmitter.emit(eventID);
    }
    computeStatsEP(stats) {
        if (stats == undefined) {
            return 0;
        }
        return new Stats(stats).computeEP(this.epWeights);
    }
    computeGemEP(gem) {
        const epFromStats = new Stats(gem.stats).computeEP(this.epWeights);
        const epFromEffect = getMetaGemEffectEP(this.spec, gem, new Stats(this.currentStats.finalStats));
        let bonusEP = 0;
        // unique items are slightly worse than non-unique because you can have only one.
        if (gem.unique) {
            bonusEP -= 0.01;
        }
        return epFromStats + epFromEffect + bonusEP;
    }
    computeEnchantEP(enchant) {
        return new Stats(enchant.stats).computeEP(this.epWeights);
    }
    computeItemEP(item) {
        if (item == null)
            return 0;
        let ep = new Stats(item.stats).computeEP(this.epWeights);
        // unique items are slightly worse than non-unique because you can have only one.
        if (item.unique) {
            ep -= 0.01;
        }
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
        return withSpecProto(this.spec, PlayerProto.create({
            name: this.getName(),
            race: this.getRace(),
            class: this.getClass(),
            equipment: this.getGear().asSpec(),
            consumes: this.getConsumes(),
            bonusStats: this.getBonusStats().asArray(),
            buffs: this.getBuffs(),
            cooldowns: this.getCooldowns(),
            talentsString: this.getTalentsString(),
        }), this.getRotation(), this.getTalents(), this.getSpecOptions());
    }
    fromProto(eventID, proto) {
        TypedEvent.freezeAllAndDo(() => {
            // TODO: Remove this on 1/31/2022 (1 month).
            if (proto.consumes && proto.consumes.darkRune) {
                proto.consumes.defaultConjured = Conjured.ConjuredDarkRune;
            }
            this.setName(eventID, proto.name);
            this.setRace(eventID, proto.race);
            this.setGear(eventID, proto.equipment ? this.sim.lookupEquipmentSpec(proto.equipment) : new Gear({}));
            this.setConsumes(eventID, proto.consumes || Consumes.create());
            this.setBonusStats(eventID, new Stats(proto.bonusStats));
            this.setBuffs(eventID, proto.buffs || IndividualBuffs.create());
            this.setCooldowns(eventID, proto.cooldowns || Cooldowns.create());
            this.setTalentsString(eventID, proto.talentsString);
            this.setRotation(eventID, this.specTypeFunctions.rotationFromPlayer(proto));
            this.setTalents(eventID, this.specTypeFunctions.talentsFromPlayer(proto));
            this.setSpecOptions(eventID, this.specTypeFunctions.optionsFromPlayer(proto));
        });
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
    fromJson(eventID, obj) {
        TypedEvent.freezeAllAndDo(() => {
            try {
                if (obj['name']) {
                    this.setName(eventID, obj['name']);
                }
            }
            catch (e) {
                console.warn('Failed to parse name: ' + e);
            }
            try {
                this.setBuffs(eventID, IndividualBuffs.fromJson(obj['buffs']));
            }
            catch (e) {
                console.warn('Failed to parse player buffs: ' + e);
            }
            try {
                const consumes = Consumes.fromJson(obj['consumes']);
                if (consumes.darkRune) {
                    consumes.defaultConjured = Conjured.ConjuredDarkRune;
                }
                this.setConsumes(eventID, consumes);
            }
            catch (e) {
                console.warn('Failed to parse consumes: ' + e);
            }
            // For legacy format. Do not remove this until 2022/01/02 (1 month).
            if (obj['customStats']) {
                obj['bonusStats'] = obj['customStats'];
            }
            try {
                this.setBonusStats(eventID, Stats.fromJson(obj['bonusStats']));
            }
            catch (e) {
                console.warn('Failed to parse bonus stats: ' + e);
            }
            try {
                this.setGear(eventID, this.sim.lookupEquipmentSpec(EquipmentSpec.fromJson(obj['gear'])));
            }
            catch (e) {
                console.warn('Failed to parse gear: ' + e);
            }
            try {
                this.setRace(eventID, obj['race']);
            }
            catch (e) {
                console.warn('Failed to parse race: ' + e);
            }
            try {
                this.setRotation(eventID, this.specTypeFunctions.rotationFromJson(obj['rotation']));
            }
            catch (e) {
                console.warn('Failed to parse rotation: ' + e);
            }
            try {
                this.setTalentsString(eventID, obj['talents']);
            }
            catch (e) {
                console.warn('Failed to parse talents: ' + e);
            }
            try {
                this.setSpecOptions(eventID, this.specTypeFunctions.optionsFromJson(obj['specOptions']));
            }
            catch (e) {
                console.warn('Failed to parse spec options: ' + e);
            }
        });
    }
    clone(eventID) {
        const newPlayer = new Player(this.spec, this.sim);
        newPlayer.fromProto(eventID, this.toProto());
        return newPlayer;
    }
}
