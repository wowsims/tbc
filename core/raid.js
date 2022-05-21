import { Debuffs } from '/tbc/core/proto/common.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
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
        this.debuffs = Debuffs.create();
        this.tanks = [];
        this.staggerStormstrikes = false;
        // Emits when a raid member is added/removed/moved.
        this.compChangeEmitter = new TypedEvent();
        this.buffsChangeEmitter = new TypedEvent();
        this.debuffsChangeEmitter = new TypedEvent();
        this.tanksChangeEmitter = new TypedEvent();
        this.staggerStormstrikesChangeEmitter = new TypedEvent();
        this.sim = sim;
        this.parties = [...Array(MAX_NUM_PARTIES).keys()].map(i => {
            const newParty = new Party(this, sim);
            newParty.compChangeEmitter.on(eventID => this.compChangeEmitter.emit(eventID));
            newParty.changeEmitter.on(eventID => this.changeEmitter.emit(eventID));
            return newParty;
        });
        this.changeEmitter = TypedEvent.onAny([
            this.compChangeEmitter,
            this.buffsChangeEmitter,
            this.debuffsChangeEmitter,
            this.tanksChangeEmitter,
        ], 'RaidChange');
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
    getDebuffs() {
        // Make a defensive copy
        return Debuffs.clone(this.debuffs);
    }
    setDebuffs(eventID, newDebuffs) {
        if (Debuffs.equals(this.debuffs, newDebuffs))
            return;
        // Make a defensive copy
        this.debuffs = Debuffs.clone(newDebuffs);
        this.debuffsChangeEmitter.emit(eventID);
    }
    getTanks() {
        // Make a defensive copy
        return this.tanks.map(tank => RaidTarget.clone(tank));
    }
    setTanks(eventID, newTanks) {
        if (this.tanks.length == newTanks.length && this.tanks.every((tank, i) => RaidTarget.equals(tank, newTanks[i])))
            return;
        // Make a defensive copy
        this.tanks = newTanks.map(tank => RaidTarget.clone(tank));
        this.tanksChangeEmitter.emit(eventID);
    }
    getStaggerStormstrikes() {
        return this.staggerStormstrikes;
    }
    setStaggerStormstrikes(eventID, newValue) {
        if (this.staggerStormstrikes == newValue)
            return;
        this.staggerStormstrikes = newValue;
        this.staggerStormstrikesChangeEmitter.emit(eventID);
    }
    toProto(forExport) {
        return RaidProto.create({
            parties: this.parties.map(party => party.toProto(forExport)),
            buffs: this.getBuffs(),
            debuffs: this.getDebuffs(),
            tanks: this.getTanks(),
            staggerStormstrikes: this.getStaggerStormstrikes(),
        });
    }
    fromProto(eventID, proto) {
        TypedEvent.freezeAllAndDo(() => {
            this.setBuffs(eventID, proto.buffs || RaidBuffs.create());
            this.setDebuffs(eventID, proto.debuffs || Debuffs.create());
            this.setStaggerStormstrikes(eventID, proto.staggerStormstrikes);
            this.setTanks(eventID, proto.tanks);
            for (let i = 0; i < MAX_NUM_PARTIES; i++) {
                if (proto.parties[i]) {
                    this.parties[i].fromProto(eventID, proto.parties[i]);
                }
                else {
                    this.parties[i].clear(eventID);
                }
            }
        });
    }
    clear(eventID) {
        TypedEvent.freezeAllAndDo(() => {
            for (let i = 0; i < MAX_NUM_PARTIES; i++) {
                this.parties[i].clear(eventID);
            }
        });
    }
}
