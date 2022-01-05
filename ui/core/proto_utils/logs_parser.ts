

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

	// Parses an entity from its log label.
	// Label should be one of:
	//   'Target 1' if a target,
	//   'PlayerName (#1)' if a player, or
	//   'PlayerName (#1) - PetName' if a pet.
	static const parseRegex = /(Target (\d+))|([a-zA-Z0-9]+ \(#\))|()/;
	static parse(label: string): Entity {
		const match = 
		if (!label.contains('#')) {
			// Must be a target.
			const index = parseInt(label.split(' ')[1]) - 1;
			return new Entity(label, '', index, true, false);
		}

		const match = label.match(/(.*) \(#(\d+)\)( - (.*))?/);
		const playerName = match[1];
		const index = parseInt(match[2]);
		const petName = match[3] ? match[4] : '';
		const name = petName ? petName : playerName;
		
		return new Entity(name, petName ? playerName : '', index, false, Boolean(petName));
	}
}

export class SimLog {
	readonly raw: string;

	// Time in seconds from the encounter start.
	readonly timestamp: number;

	readonly source: Entity | null;
	readonly target: Entity | null;

	constructor(raw: string, timestamp: number, source: Entity | null, target: Entity | null) {
		this.raw = raw;
		this.timestamp = timestamp;
		this.source = source;
		this.target = target;
	}

	static parse(raw: string): SimLog {
		let match = raw.match(/\[([0-9]+\.[0-9]+)\]\w*(.*)/);
		if (!match[1]) {
			return new SimLog(raw);
		}

		const timestamp = parseFloat(match[1]);
		let remainder = match[2];
		const entityMatch = remainder.match(/\[([0-9]+\.[0-9]+)\]\w*(.*):\w*(.*)/);
	}
}

export class AuraGainedLog extends SimLog {
	readonly auraName: string,

	static parse(raw: string) {
		const match = remainder.match(/\[([0-9]+\.[0-9]+)\] (.*?): Aura gained: (.*)/);
	}
}
