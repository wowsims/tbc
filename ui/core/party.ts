import { Spec } from '/tbc/core/proto/common.js';

import { Player } from './player.js';
import { TypedEvent } from './typed_event.js';
import { Sim } from './sim.js';

export interface PartyConfig {
}

export const MAX_PARTY_SIZE = 5;

// Manages all the settings for a single Party.
export class Party {
  // Emits when anything in the party changes.
  readonly changeEmitter = new TypedEvent<void>();

	// Should never hold more than MAX_PARTY_SIZE elements.
	private players: Array<Player<any>>;

	private readonly sim: Sim;

  constructor(sim: Sim) {
		this.sim = sim;
		this.players = [];
  }

	size(): number {
		return this.players.length;
	}

	empty(): boolean {
		return this.size() == 0;
	}

	getPlayers(): Array<Player<any>> {
		// Make defensive copy.
		return this.players.slice();
	}

	addPlayer(player: Player<any>) {
		if (this.size() >= MAX_PARTY_SIZE) {
			throw new Error('Cannot add player to full party');
		}

		this.players.push(player);
		player.changeEmitter.on(() => this.changeEmitter.emit());
		this.changeEmitter.emit();
	}

	removePlayer(playerToRemove: Player<any>) {
		const removeIndex = this.players.findIndex(partyPlayer => partyPlayer == playerToRemove);
		if (removeIndex == -1) {
			return;
		}

		// TODO: Might need to remove the player changeEmitter callback here.

		this.players = this.players.splice(removeIndex, 1);
		this.changeEmitter.emit();
	}

  // Returns JSON representing all the current values.
  toJson(): Object {
		return this.players.map(player => {
			return {
				'spec': player.spec,
				'player': player.toJson(),
			};
		});
  }

  // Set all the current values, assumes obj is the same type returned by toJson().
  fromJson(obj: any) {
		this.players = [];
		this.changeEmitter.emit();

		(obj as Array<any>).forEach(playerObj => {
			const newPlayer = new Player(playerObj['spec'] as Spec, this.sim);
			newPlayer.fromJson(playerObj['player']);
			this.addPlayer(newPlayer);
		});
  }
}
