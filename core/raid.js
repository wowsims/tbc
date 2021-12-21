import { Raid as RaidProto } from '/tbc/core/proto/api.js';
import { RaidBuffs } from '/tbc/core/proto/common.js';
import { NO_TARGET } from '/tbc/core/proto_utils/utils.js';
import { Party, MAX_PARTY_SIZE } from './party.js';
import { TypedEvent } from './typed_event.js';
import { sum } from './utils.js';
export const MAX_NUM_PARTIES = 5;
// Manages all the settings for a single Raid.
export class Raid {
    constructor(sim) {
        this.buffs = RaidBuffs.create();
        // Emits when a raid member is added/removed/moved.
        this.compChangeEmitter = new TypedEvent();
        this.buffsChangeEmitter = new TypedEvent();
        // Emits when anything in the raid changes.
        this.changeEmitter = new TypedEvent();
        this.modifyRaidProto = () => { };
        this.sim = sim;
        this.parties = [...Array(MAX_NUM_PARTIES).keys()].map(i => {
            const newParty = new Party(this, sim);
            newParty.compChangeEmitter.on(eventID => this.compChangeEmitter.emit(eventID));
            newParty.changeEmitter.on(eventID => this.changeEmitter.emit(eventID));
            return newParty;
        });
        [
            this.compChangeEmitter,
            this.buffsChangeEmitter,
        ].forEach(emitter => emitter.on(eventID => this.changeEmitter.emit(eventID)));
    }
    size() {
        return sum(this.parties.map(party => party.size()));
    }
    isEmpty() {
        return this.size() == 0;
    }
    getParties() {
        // Make defensive copy.
        return this.parties.slice();
    }
    getParty(index) {
        return this.parties[index];
    }
    getPlayers() {
        return this.parties.map(party => party.getPlayers()).flat();
    }
    getPlayer(index) {
        const party = this.parties[Math.floor(index / MAX_PARTY_SIZE)];
        return party.getPlayer(index % MAX_PARTY_SIZE);
    }
    getPlayerFromRaidTarget(raidTarget) {
        if (raidTarget.targetIndex == NO_TARGET) {
            return null;
        }
        else {
            return this.getPlayer(raidTarget.targetIndex);
        }
    }
    setPlayer(eventID, index, newPlayer) {
        const party = this.parties[Math.floor(index / MAX_PARTY_SIZE)];
        party.setPlayer(eventID, index % MAX_PARTY_SIZE, newPlayer);
    }
    getClassCount(playerClass) {
        return this.getPlayers().filter(player => player != null && player.getClass() == playerClass).length;
    }
    getBuffs() {
        // Make a defensive copy
        return RaidBuffs.clone(this.buffs);
    }
    setBuffs(eventID, newBuffs) {
        if (RaidBuffs.equals(this.buffs, newBuffs))
            return;
        // Make a defensive copy
        this.buffs = RaidBuffs.clone(newBuffs);
        this.buffsChangeEmitter.emit(eventID);
    }
    setModifyRaidProto(newModFn) {
        this.modifyRaidProto = newModFn;
    }
    toProto() {
        const raidProto = RaidProto.create({
            parties: this.parties.map(party => party.toProto()),
            buffs: this.buffs,
        });
        this.modifyRaidProto(raidProto);
        return raidProto;
    }
    fromProto(eventID, proto) {
        TypedEvent.freezeAll();
        this.setBuffs(eventID, proto.buffs || RaidBuffs.create());
        for (let i = 0; i < MAX_NUM_PARTIES; i++) {
            if (proto.parties[i]) {
                this.parties[i].fromProto(eventID, proto.parties[i]);
            }
            else {
                this.parties[i].clear(eventID);
            }
        }
        TypedEvent.unfreezeAll();
    }
    // Returns JSON representing all the current values.
    toJson() {
        return {
            'parties': this.parties.map(party => party.toJson()),
            'buffs': RaidBuffs.toJson(this.buffs),
        };
    }
    // Set all the current values, assumes obj is the same type returned by toJson().
    fromJson(eventID, obj) {
        TypedEvent.freezeAll();
        try {
            this.setBuffs(eventID, RaidBuffs.fromJson(obj['buffs']));
        }
        catch (e) {
            console.warn('Failed to parse raid buffs: ' + e);
        }
        if (obj['parties']) {
            for (let i = 0; i < MAX_NUM_PARTIES; i++) {
                const partyObj = obj['parties'][i];
                if (!partyObj) {
                    this.parties[i].clear(eventID);
                    continue;
                }
                this.parties[i].fromJson(eventID, partyObj);
            }
        }
        TypedEvent.unfreezeAll();
    }
}
