import { Spec } from '/tbc/core/proto/common.js';

import { Player } from './player.js';
import { TypedEvent } from './typed_event.js';
import { Sim } from './sim.js';

export interface PartyConfig {
}

export const MAX_PARTY_SIZE = 5;

// Manages all the settings for a single Party.
export class Party {
	// Emits when a party member is added/removed/moved.
  readonly compChangeEmitter = new TypedEvent<void>();

  // Emits when anything in the party changes.
  readonly changeEmitter = new TypedEvent<void>();

	// Should always hold exactly MAX_PARTY_SIZE elements.
	private players: Array<Player<any> | null>;

	readonly sim: Sim;

	private readonly playerChangeListener: () => void;

  constructor(sim: Sim) {
		this.sim = sim;
		this.players = [...Array(MAX_PARTY_SIZE).keys()].map(i => null);
		this.playerChangeListener = () => this.changeEmitter.emit();

		this.compChangeEmitter.on(() => this.changeEmitter.emit());
  }

	size(): number {
		return this.players.filter(player => player != null).length;
	}

	empty(): boolean {
		return this.size() == 0;
	}

	getPlayers(): Array<Player<any> | null> {
		// Make defensive copy.
		return this.players.slice();
	}

	getPlayer(playerIndex: number): Player<any> | null {
		return this.players[playerIndex];
	}

	setPlayer(playerIndex: number, newPlayer: Player<any> | null) {
		if (playerIndex < 0 || playerIndex >= MAX_PARTY_SIZE) {
			throw new Error('Invalid player index: ' + playerIndex);
		}

		if (newPlayer == this.players[playerIndex]) {
			return;
		}

		if (this.players[playerIndex] != null) {
			this.players[playerIndex]!.changeEmitter.off(this.playerChangeListener);
		}
		if (newPlayer != null) {
			newPlayer.changeEmitter.on(this.playerChangeListener);
		}

		this.players[playerIndex] = newPlayer;
		this.compChangeEmitter.emit();
	}

  // Returns JSON representing all the current values.
  toJson(): Object {
		return this.players.map(player => {
			if (player == null) {
				return null;
			} else {
				return {
					'spec': player.spec,
					'player': player.toJson(),
				};
			}
		});
  }

  // Set all the current values, assumes obj is the same type returned by toJson().
  fromJson(obj: any) {
		this.players = [];
		this.changeEmitter.emit();

		(obj as Array<any>).forEach((playerObj, i) => {
			if (playerObj == null) {
				this.setPlayer(i, null);
			} else {
				const newPlayer = new Player(playerObj['spec'] as Spec, this.sim);
				newPlayer.fromJson(playerObj['player']);
				this.setPlayer(i, newPlayer);
			}
		});
  }
}
