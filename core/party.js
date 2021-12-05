import { Player } from './player.js';
import { TypedEvent } from './typed_event.js';
export const MAX_PARTY_SIZE = 5;
// Manages all the settings for a single Party.
export class Party {
    constructor(sim) {
        // Emits when a party member is added/removed/moved.
        this.compChangeEmitter = new TypedEvent();
        // Emits when anything in the party changes.
        this.changeEmitter = new TypedEvent();
        this.sim = sim;
        this.players = [...Array(MAX_PARTY_SIZE).keys()].map(i => null);
        this.playerChangeListener = () => this.changeEmitter.emit();
        this.compChangeEmitter.on(() => this.changeEmitter.emit());
    }
    size() {
        return this.players.filter(player => player != null).length;
    }
    empty() {
        return this.size() == 0;
    }
    getPlayers() {
        // Make defensive copy.
        return this.players.slice();
    }
    getPlayer(playerIndex) {
        return this.players[playerIndex];
    }
    setPlayer(playerIndex, newPlayer) {
        if (playerIndex < 0 || playerIndex >= MAX_PARTY_SIZE) {
            throw new Error('Invalid player index: ' + playerIndex);
        }
        if (newPlayer == this.players[playerIndex]) {
            return;
        }
        if (this.players[playerIndex] != null) {
            this.players[playerIndex].changeEmitter.off(this.playerChangeListener);
        }
        if (newPlayer != null) {
            newPlayer.changeEmitter.on(this.playerChangeListener);
        }
        this.players[playerIndex] = newPlayer;
        this.compChangeEmitter.emit();
    }
    // Returns JSON representing all the current values.
    toJson() {
        return this.players.map(player => {
            if (player == null) {
                return null;
            }
            else {
                return {
                    'spec': player.spec,
                    'player': player.toJson(),
                };
            }
        });
    }
    // Set all the current values, assumes obj is the same type returned by toJson().
    fromJson(obj) {
        this.players = [];
        this.changeEmitter.emit();
        obj.forEach((playerObj, i) => {
            if (playerObj == null) {
                this.setPlayer(i, null);
            }
            else {
                const newPlayer = new Player(playerObj['spec'], this.sim);
                newPlayer.fromJson(playerObj['player']);
                this.setPlayer(i, newPlayer);
            }
        });
    }
}
