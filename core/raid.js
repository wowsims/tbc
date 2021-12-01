import { Party } from './party.js';
import { TypedEvent } from './typed_event.js';
import { sum } from './utils.js';
export const MAX_NUM_PARTIES = 5;
// Manages all the settings for a single Raid.
export class Raid {
    constructor(sim) {
        // Emits when anything in the raid changes.
        this.changeEmitter = new TypedEvent();
        this.sim = sim;
        // TODO: Use MAX_NUM_PARTIES
        this.parties = [
            new Party(sim),
            new Party(sim),
            new Party(sim),
            new Party(sim),
            new Party(sim),
        ];
        this.parties.forEach(party => {
            party.changeEmitter.on(() => this.changeEmitter.emit());
        });
    }
    size() {
        return sum(this.parties.map(party => party.size()));
    }
    empty() {
        return this.size() == 0;
    }
    getParties() {
        // Make defensive copy.
        return this.parties.slice();
    }
    //getPlayers(): Array<Player<any>> {
    //	// TODO: Flatten
    //	return this.parties.map(party => party.getPlayers());
    //}
    // Returns JSON representing all the current values.
    toJson() {
        return this.parties.map(party => party.toJson());
    }
    // Set all the current values, assumes obj is the same type returned by toJson().
    fromJson(obj) {
        this.parties = obj.map(partyObj => {
            const newParty = new Party(this.sim);
            newParty.fromJson(partyObj);
            return newParty;
        });
        this.parties.forEach(party => {
            party.changeEmitter.on(() => this.changeEmitter.emit());
        });
        this.changeEmitter.emit();
    }
}
