import { Player } from './player.js';
import { TypedEvent } from './typed_event.js';
export const MAX_PARTY_SIZE = 5;
// Manages all the settings for a single Party.
export class Party {
    constructor(sim) {
        // Emits when anything in the party changes.
        this.changeEmitter = new TypedEvent();
        this.sim = sim;
        this.players = [];
    }
    size() {
        return this.players.length;
    }
    empty() {
        return this.size() == 0;
    }
    getPlayers() {
        // Make defensive copy.
        return this.players.slice();
    }
    addPlayer(player) {
        if (this.size() >= MAX_PARTY_SIZE) {
            throw new Error('Cannot add player to full party');
        }
        this.players.push(player);
        player.changeEmitter.on(() => this.changeEmitter.emit());
        this.changeEmitter.emit();
    }
    removePlayer(playerToRemove) {
        const removeIndex = this.players.findIndex(partyPlayer => partyPlayer == playerToRemove);
        if (removeIndex == -1) {
            return;
        }
        // TODO: Might need to remove the player changeEmitter callback here.
        this.players = this.players.splice(removeIndex, 1);
        this.changeEmitter.emit();
    }
    // Returns JSON representing all the current values.
    toJson() {
        return this.players.map(player => {
            return {
                'spec': player.spec,
                'player': player.toJson(),
            };
        });
    }
    // Set all the current values, assumes obj is the same type returned by toJson().
    fromJson(obj) {
        this.players = [];
        this.changeEmitter.emit();
        obj.forEach(playerObj => {
            const newPlayer = new Player(playerObj['spec'], this.sim);
            newPlayer.fromJson(playerObj['player']);
            this.addPlayer(newPlayer);
        });
    }
}
