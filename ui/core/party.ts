import { Party as PartyProto } from '/tbc/core/proto/api.js';
import { PartyStats as PartyStatsProto } from '/tbc/core/proto/api.js';
import { Player as PlayerProto } from '/tbc/core/proto/api.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { Class } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { playerToSpec } from '/tbc/core/proto_utils/utils.js';

import { Raid } from './raid.js';
import { Player } from './player.js';
import { EventID, TypedEvent } from './typed_event.js';
import { Sim } from './sim.js';

export const MAX_PARTY_SIZE = 5;

// Manages all the settings for a single Party.
export class Party {
	readonly sim: Sim;
	readonly raid: Raid;

  private buffs: PartyBuffs = PartyBuffs.create();

	// Emits when a party member is added/removed/moved.
  readonly compChangeEmitter = new TypedEvent<void>();

  readonly buffsChangeEmitter = new TypedEvent<void>();

  // Emits when anything in the party changes.
  readonly changeEmitter = new TypedEvent<void>();

	// Should always hold exactly MAX_PARTY_SIZE elements.
	private players: Array<Player<any> | null>;

	private readonly playerChangeListener: (eventID: EventID) => void;

  constructor(raid: Raid, sim: Sim) {
		this.sim = sim;
		this.raid = raid;
		this.players = [...Array(MAX_PARTY_SIZE).keys()].map(i => null);
		this.playerChangeListener = eventID => this.changeEmitter.emit(eventID);

		[
			this.compChangeEmitter,
			this.buffsChangeEmitter,
		].forEach(emitter => emitter.on(eventID => this.changeEmitter.emit(eventID)));
  }

	size(): number {
		return this.players.filter(player => player != null).length;
	}

	isEmpty(): boolean {
		return this.size() == 0;
	}

	clear(eventID: EventID) {
		this.setBuffs(eventID, PartyBuffs.create());
		for (let i = 0; i < MAX_PARTY_SIZE; i++) {
			this.setPlayer(eventID, i, null);
		}
	}

	// Returns this party's index within the raid [0-4].
	getIndex(): number {
		return this.raid.getParties().indexOf(this);
	}

	getPlayers(): Array<Player<any> | null> {
		// Make defensive copy.
		return this.players.slice();
	}

	getPlayer(playerIndex: number): Player<any> | null {
		return this.players[playerIndex];
	}

	setPlayer(eventID: EventID, playerIndex: number, newPlayer: Player<any> | null) {
		if (playerIndex < 0 || playerIndex >= MAX_PARTY_SIZE) {
			throw new Error('Invalid player index: ' + playerIndex);
		}

		if (newPlayer == this.players[playerIndex]) {
			return;
		}

		TypedEvent.freezeAll();
		const oldPlayer = this.players[playerIndex];
		this.players[playerIndex] = newPlayer;

		if (oldPlayer != null) {
			oldPlayer.changeEmitter.off(this.playerChangeListener);
			oldPlayer.setParty(null);
		}
		if (newPlayer != null) {
			const newPlayerOldParty = newPlayer.getParty();
			if (newPlayerOldParty) {
				newPlayerOldParty.setPlayer(eventID, newPlayer.getPartyIndex(), null);
			}
			newPlayer.changeEmitter.on(this.playerChangeListener);
			newPlayer.setParty(this);
		}

		this.compChangeEmitter.emit(eventID);
		TypedEvent.unfreezeAll();
	}

  getBuffs(): PartyBuffs {
    // Make a defensive copy
    return PartyBuffs.clone(this.buffs);
  }

  setBuffs(eventID: EventID, newBuffs: PartyBuffs) {
    if (PartyBuffs.equals(this.buffs, newBuffs))
      return;

    // Make a defensive copy
    this.buffs = PartyBuffs.clone(newBuffs);
    this.buffsChangeEmitter.emit(eventID);
  }

	toProto(): PartyProto {
		return PartyProto.create({
			players: this.players.map(player => player == null ? PlayerProto.create() : player.toProto()),
			buffs: this.buffs,
		});
	}

	fromProto(eventID: EventID, proto: PartyProto) {
		TypedEvent.freezeAll();
		this.setBuffs(eventID, proto.buffs || PartyBuffs.create());

		for (let i = 0; i < MAX_PARTY_SIZE; i++) {
			if (!proto.players[i] || proto.players[i].class == Class.ClassUnknown) {
				this.setPlayer(eventID, i, null);
				continue;
			}

			const playerProto = proto.players[i];
			const spec = playerToSpec(playerProto);
			const newPlayer = new Player(spec, this.sim);
			newPlayer.fromProto(eventID, playerProto);
			this.setPlayer(eventID, i, newPlayer);
		}
		TypedEvent.unfreezeAll();
	}

  // Returns JSON representing all the current values.
  toJson(): Object {
		return {
			'players': this.players.map(player => {
				if (player == null) {
					return null;
				} else {
					return {
						'spec': player.spec,
						'player': player.toJson(),
					};
				}
			}),
			'buffs': PartyBuffs.toJson(this.getBuffs()),
		};
  }

  // Set all the current values, assumes obj is the same type returned by toJson().
  fromJson(eventID: EventID, obj: any) {
		TypedEvent.unfreezeAll();
		try {
			this.setBuffs(eventID, PartyBuffs.fromJson(obj['buffs']));
		} catch (e) {
			console.warn('Failed to parse party buffs: ' + e);
		}

		if (obj['players']) {
			for (let i = 0; i < MAX_PARTY_SIZE; i++) {
				const playerObj = obj['players'][i];
				if (!playerObj) {
					this.setPlayer(eventID, i, null);
					continue;
				}

				const newSpec = playerObj['spec'] as Spec;
				if (this.players[i] != null && this.players[i]!.spec == newSpec) {
					this.players[i]!.fromJson(eventID, playerObj['player']);
				} else {
					const newPlayer = new Player(playerObj['spec'] as Spec, this.sim);
					newPlayer.fromJson(eventID, playerObj['player']);
					this.setPlayer(eventID, i, newPlayer);
				}
			}
		}
		TypedEvent.unfreezeAll();
  }
}
