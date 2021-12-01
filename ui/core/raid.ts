import { Spec } from '/tbc/core/proto/common.js';

import { Party } from './party.js';
import { Player } from './player.js';
import { TypedEvent } from './typed_event.js';
import { Sim } from './sim.js';
import { sum } from './utils.js';

export const MAX_NUM_PARTIES = 5;

// Manages all the settings for a single Raid.
export class Raid {
  // Emits when anything in the raid changes.
  readonly changeEmitter = new TypedEvent<void>();

	// Should always hold exactly MAX_NUM_PARTIES elements.
	private parties: Array<Party>;

	private readonly sim: Sim;

  constructor(sim: Sim) {
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

	size(): number {
		return sum(this.parties.map(party => party.size()));
	}

	empty(): boolean {
		return this.size() == 0;
	}

	getParties(): Array<Party> {
		// Make defensive copy.
		return this.parties.slice();
	}

	//getPlayers(): Array<Player<any>> {
	//	// TODO: Flatten
	//	return this.parties.map(party => party.getPlayers());
	//}

  // Returns JSON representing all the current values.
  toJson(): Object {
		return this.parties.map(party => party.toJson());
  }

  // Set all the current values, assumes obj is the same type returned by toJson().
  fromJson(obj: any) {
		this.parties = (obj as Array<any>).map(partyObj => {
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
