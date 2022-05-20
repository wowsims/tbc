import { Party as PartyProto } from '/tbc/core/proto/api.js';
import { Player as PlayerProto } from '/tbc/core/proto/api.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { Class } from '/tbc/core/proto/common.js';
import { playerToSpec } from '/tbc/core/proto_utils/utils.js';
import { Player } from './player.js';
import { TypedEvent } from './typed_event.js';
export const MAX_PARTY_SIZE = 5;
// Manages all the settings for a single Party.
export class Party {
    constructor(raid, sim) {
        this.buffs = PartyBuffs.create();
        // Emits when a party member is added/removed/moved.
        this.compChangeEmitter = new TypedEvent();
        this.buffsChangeEmitter = new TypedEvent();
        this.sim = sim;
        this.raid = raid;
        this.players = [...Array(MAX_PARTY_SIZE).keys()].map(i => null);
        this.playerChangeListener = eventID => this.changeEmitter.emit(eventID);
        this.changeEmitter = TypedEvent.onAny([
            this.compChangeEmitter,
            this.buffsChangeEmitter,
        ], 'PartyChange');
    }
    size() {
        return this.players.filter(player => player != null).length;
    }
    isEmpty() {
        return this.size() == 0;
    }
    clear(eventID) {
        this.setBuffs(eventID, PartyBuffs.create());
        for (let i = 0; i < MAX_PARTY_SIZE; i++) {
            this.setPlayer(eventID, i, null);
        }
    }
    // Returns this party's index within the raid [0-4].
    getIndex() {
        return this.raid.getParties().indexOf(this);
    }
    getPlayers() {
        // Make defensive copy.
        return this.players.slice();
    }
    getPlayer(playerIndex) {
        return this.players[playerIndex];
    }
    setPlayer(eventID, playerIndex, newPlayer) {
        if (playerIndex < 0 || playerIndex >= MAX_PARTY_SIZE) {
            throw new Error('Invalid player index: ' + playerIndex);
        }
        if (newPlayer == this.players[playerIndex]) {
            return;
        }
        TypedEvent.freezeAllAndDo(() => {
            const oldPlayer = this.players[playerIndex];
            if (oldPlayer != null) {
                oldPlayer.changeEmitter.off(this.playerChangeListener);
                oldPlayer.setParty(null);
            }
            if (newPlayer != null) {
                const newPlayerOldParty = newPlayer.getParty();
                if (newPlayerOldParty) {
                    newPlayerOldParty.setPlayer(eventID, newPlayer.getPartyIndex(), null);
                }
                this.players[playerIndex] = newPlayer;
                newPlayer.changeEmitter.on(this.playerChangeListener);
                newPlayer.setParty(this);
            }
            else {
                this.players[playerIndex] = null;
            }
            this.compChangeEmitter.emit(eventID);
        });
    }
    getBuffs() {
        // Make a defensive copy
        return PartyBuffs.clone(this.buffs);
    }
    setBuffs(eventID, newBuffs) {
        if (PartyBuffs.equals(this.buffs, newBuffs))
            return;
        // Make a defensive copy
        this.buffs = PartyBuffs.clone(newBuffs);
        this.buffsChangeEmitter.emit(eventID);
    }
    toProto(forExport) {
        return PartyProto.create({
            players: this.players.map(player => player == null ? PlayerProto.create() : player.toProto(forExport)),
            buffs: this.buffs,
        });
    }
    fromProto(eventID, proto) {
        TypedEvent.freezeAllAndDo(() => {
            this.setBuffs(eventID, proto.buffs || PartyBuffs.create());
            for (let i = 0; i < MAX_PARTY_SIZE; i++) {
                if (!proto.players[i] || proto.players[i].class == Class.ClassUnknown) {
                    this.setPlayer(eventID, i, null);
                    continue;
                }
                const playerProto = proto.players[i];
                const spec = playerToSpec(playerProto);
                const currentPlayer = this.players[i];
                // Reuse the current player if possible, so that event handlers are preserved.
                if (currentPlayer && spec == currentPlayer.spec) {
                    currentPlayer.fromProto(eventID, playerProto);
                }
                else {
                    const newPlayer = new Player(spec, this.sim);
                    newPlayer.fromProto(eventID, playerProto);
                    this.setPlayer(eventID, i, newPlayer);
                }
            }
        });
    }
}
