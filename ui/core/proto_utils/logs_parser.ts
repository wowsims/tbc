import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';

export class Entity {
	readonly name: string;
	readonly ownerName: string; // Blank if not a pet.

	// Either target index, player index, or owner index depending on what kind
	// of entity this is.
	private readonly index: number;

	private readonly isTarget: boolean;
	private readonly isPet: boolean;

	constructor(name: string, ownerName: string, index: number, isTarget: boolean, isPet: boolean) {
		this.name = name;
		this.ownerName = ownerName;
		this.index = index;
		this.isTarget = isTarget;
		this.isPet = isPet;
	}

	// Parses one or more Entities from a string.
	// Each entity label should be one of:
	//   'Target 1' if a target,
	//   'PlayerName (#1)' if a player, or
	//   'PlayerName (#1) - PetName' if a pet.
	static parseRegex = /(Target (\d+))|(([a-zA-Z0-9]+) \(#(\d+)\) - ([a-zA-Z0-9]+))|(([a-zA-Z0-9]+) \(#(\d+)\))/g;
	static parseAll(str: string): Array<Entity> {
		return Array.from(str.matchAll(Entity.parseRegex)).map(match => {
			if (match[1]) {
				return new Entity(match[1], '', parseInt(match[2]), true, false);
			} else if (match[3]) {
				return new Entity(match[6], match[4], parseInt(match[5]), false, true);
			} else if (match[7]) {
				return new Entity(match[8], '', parseInt(match[9]), false, false);
			} else {
				throw new Error('Invalid Entity match');
			}
		});
	}
}

interface SimLogParams {
	raw: string,
	timestamp: number,
	source: Entity | null,
	target: Entity | null,
}

export class SimLog {
	readonly raw: string;

	// Time in seconds from the encounter start.
	readonly timestamp: number;

	readonly source: Entity | null;
	readonly target: Entity | null;

	constructor(params: SimLogParams) {
		this.raw = params.raw;
		this.timestamp = params.timestamp;
		this.source = params.source;
		this.target = params.target;
	}

	static parseAll(result: RaidSimResult): Array<SimLog> {
		const lines = result.logs.split('\n');

		return lines.map(line => {
			const params: SimLogParams = {
				raw: line,
				timestamp: 0,
				source: null,
				target: null,
			};

			let match = line.match(/\[([0-9]+\.[0-9]+)\]\w*(.*)/);
			if (!match || !match[1]) {
				return new SimLog(params);
			}

			params.timestamp = parseFloat(match[1]);
			let remainder = match[2];

			const entities = Entity.parseAll(remainder);
			params.source = entities[0] || null;
			params.target = entities[1] || null;

			return AuraGainedLog.parse(params)
					|| AuraFadedLog.parse(params)
					|| new SimLog(params);
		});
	}
}

export class AuraGainedLog extends SimLog {
	readonly auraName: string;

	constructor(params: SimLogParams, auraName: string) {
		super(params);
		this.auraName = auraName;
	}

	static parse(params: SimLogParams): AuraGainedLog | null {
		const match = params.raw.match(/Aura gained: \[(.*)\]/);
		if (match && match[1]) {
			return new AuraGainedLog(params, match[1]);
		} else {
			return null;
		}
	}
}

export class AuraFadedLog extends SimLog {
	readonly auraName: string;

	constructor(params: SimLogParams, auraName: string) {
		super(params);
		this.auraName = auraName;
	}

	static parse(params: SimLogParams): AuraFadedLog | null {
		const match = params.raw.match(/Aura faded: \[(.*)\]/);
		if (match && match[1]) {
			return new AuraFadedLog(params, match[1]);
		} else {
			return null;
		}
	}
}
