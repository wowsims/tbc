import { Class } from '/tbc/core/proto/common.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
import { Raid as RaidProto } from '/tbc/core/proto/api.js';
import { RaidStats as RaidStatsProto } from '/tbc/core/proto/api.js';
import { RaidBuffs } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { NO_TARGET } from '/tbc/core/proto_utils/utils.js';

import { Party, MAX_PARTY_SIZE } from './party.js';
import { Player } from './player.js';
import { EventID, TypedEvent } from './typed_event.js';
import { Sim } from './sim.js';
import { sum } from './utils.js';

export const MAX_NUM_PARTIES = 5;

// Manages all the settings for a single Raid.
export class Raid {
	private buffs: RaidBuffs = RaidBuffs.create();
	private staggerStormstrikes: boolean = false;

	// Emits when a raid member is added/removed/moved.
	readonly compChangeEmitter = new TypedEvent<void>();

	readonly buffsChangeEmitter = new TypedEvent<void>();
	readonly staggerStormstrikesChangeEmitter = new TypedEvent<void>();

	// Emits when anything in the raid changes.
	readonly changeEmitter: TypedEvent<void>;

	// Should always hold exactly MAX_NUM_PARTIES elements.
	private parties: Array<Party>;

	readonly sim: Sim;

	constructor(sim: Sim) {
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
		], 'RaidChange');
	}

	size(): number {
		return sum(this.parties.map(party => party.size()));
	}

	isEmpty(): boolean {
		return this.size() == 0;
	}

	getParties(): Array<Party> {
		// Make defensive copy.
		return this.parties.slice();
	}

	getParty(index: number): Party {
		return this.parties[index];
	}

	getPlayers(): Array<Player<any> | null> {
		return this.parties.map(party => party.getPlayers()).flat();
	}

	getPlayer(index: number): Player<any> | null {
		const party = this.parties[Math.floor(index / MAX_PARTY_SIZE)];
		return party.getPlayer(index % MAX_PARTY_SIZE);
	}

	getPlayerFromRaidTarget(raidTarget: RaidTarget): Player<any> | null {
		if (raidTarget.targetIndex == NO_TARGET) {
			return null;
		} else {
			return this.getPlayer(raidTarget.targetIndex);
		}
	}

	setPlayer(eventID: EventID, index: number, newPlayer: Player<any> | null) {
		const party = this.parties[Math.floor(index / MAX_PARTY_SIZE)];
		party.setPlayer(eventID, index % MAX_PARTY_SIZE, newPlayer);
	}

	getClassCount(playerClass: Class) {
		return this.getPlayers().filter(player => player != null && player.getClass() == playerClass).length;
	}

	getBuffs(): RaidBuffs {
		// Make a defensive copy
		return RaidBuffs.clone(this.buffs);
	}

	setBuffs(eventID: EventID, newBuffs: RaidBuffs) {
		if (RaidBuffs.equals(this.buffs, newBuffs))
			return;

		// Make a defensive copy
		this.buffs = RaidBuffs.clone(newBuffs);
		this.buffsChangeEmitter.emit(eventID);
	}

	getStaggerStormstrikes(): boolean {
		return this.staggerStormstrikes;
	}

	setStaggerStormstrikes(eventID: EventID, newValue: boolean) {
		if (this.staggerStormstrikes == newValue)
			return;

		this.staggerStormstrikes = newValue;
		this.staggerStormstrikesChangeEmitter.emit(eventID);
	}

	toProto(forExport?: boolean): RaidProto {
		return RaidProto.create({
			parties: this.parties.map(party => party.toProto(forExport)),
			buffs: this.buffs,
			staggerStormstrikes: this.staggerStormstrikes,
		});
	}

	fromProto(eventID: EventID, proto: RaidProto) {
		TypedEvent.freezeAllAndDo(() => {
			this.setBuffs(eventID, proto.buffs || RaidBuffs.create());
			this.setStaggerStormstrikes(eventID, proto.staggerStormstrikes);

			for (let i = 0; i < MAX_NUM_PARTIES; i++) {
				if (proto.parties[i]) {
					this.parties[i].fromProto(eventID, proto.parties[i]);
				} else {
					this.parties[i].clear(eventID);
				}
			}
		});
	}
}
