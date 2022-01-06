import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';
import { sum } from '/tbc/core/utils.js';

export class Entity {
	readonly name: string;
	readonly ownerName: string; // Blank if not a pet.

	// Either target index, player index, or owner index depending on what kind
	// of entity this is.
	readonly index: number;

	readonly isTarget: boolean;
	readonly isPet: boolean;

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
				return new Entity(match[1], '', parseInt(match[2]) - 1, true, false);
			} else if (match[3]) {
				return new Entity(match[6], match[4], parseInt(match[5]) - 1, false, true);
			} else if (match[7]) {
				return new Entity(match[8], '', parseInt(match[9]) - 1, false, false);
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

			// Order from most to least common to reduce number of checks.
			return DamageDealtLog.parse(params)
					|| ManaChangedLog.parse(params)
					|| AuraGainedLog.parse(params)
					|| AuraFadedLog.parse(params)
					|| new SimLog(params);
		});
	}

	// Remove events that happen at the same time.
	// Make sure we always keep the last log for each timestamp.
	static filterDuplicateTimestamps<LogType extends SimLog>(logs: Array<LogType>): Array<LogType> {
		const numLogs = logs.length;
		if (numLogs == 0) {
			return logs;
		}

		return logs.filter((log, i) => {
			return i == 0 || i == (numLogs - 1) || log.timestamp != logs[i + 1].timestamp;
		});
	}
}

export class DamageDealtLog extends SimLog {
	readonly amount: number;
	readonly miss: boolean;
	readonly hit: boolean;
	readonly crit: boolean;
	readonly partialResist1_4: boolean;
	readonly partialResist2_4: boolean;
	readonly partialResist3_4: boolean;
	readonly cause: string;

	constructor(params: SimLogParams, amount: number, miss: boolean, crit: boolean, partialResist1_4: boolean, partialResist2_4: boolean, partialResist3_4: boolean, cause: string) {
		super(params);
		this.amount = amount;
		this.miss = miss;
		this.hit = !miss && !crit;
		this.crit = crit;
		this.partialResist1_4 = partialResist1_4;
		this.partialResist2_4 = partialResist2_4;
		this.partialResist3_4 = partialResist3_4;
		this.cause = cause;
	}

	static parse(params: SimLogParams): DamageDealtLog | null {
		const match = params.raw.match(/] (.*?) ((Miss)|(Hit)|(Crit)|(ticked))( for (\d+\.\d+) damage.( \((\d+)% Resist\))?)?/);
		if (match) {
			const cause = match[1];
			if (match[2] == 'Miss') {
				return new DamageDealtLog(params, 0, true, false, false, false, false, cause);
			}

			const amount = parseFloat(match[8]);
			return new DamageDealtLog(params, amount, false, match[2] == 'Crit', match[10] == '25', match[10] == '50', match[10] == '75', cause);
		} else {
			return null;
		}
	}
}

export class DpsLog extends SimLog {
	readonly dps: number;

	constructor(params: SimLogParams, dps: number) {
		super(params);
		this.dps = dps;
	}

	static DPS_WINDOW = 15; // Window over which to calculate DPS.
	static fromDamageDealt(damageDealtLogs: Array<DamageDealtLog>): Array<DpsLog> {
		let curDamageLogs: Array<DamageDealtLog> = [];
		let curDamageTotal = 0;

		return damageDealtLogs.map(ddLog => {
			curDamageLogs.push(ddLog);
			curDamageTotal += ddLog.amount;

			const newStartIdx = curDamageLogs.findIndex(curLog => {
				const inWindow = curLog.timestamp > ddLog.timestamp - DpsLog.DPS_WINDOW;
				if (!inWindow) {
					curDamageTotal -= curLog.amount;
				}
				return inWindow;
			});
			if (newStartIdx == -1) {
				curDamageLogs = [];
			} else {
				curDamageLogs = curDamageLogs.slice(newStartIdx);
			}

			const dps = curDamageTotal / DpsLog.DPS_WINDOW;

			return new DpsLog({
				raw: ddLog.raw,
				timestamp: ddLog.timestamp,
				source: ddLog.source,
				target: ddLog.target,
			}, dps);
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

export class ManaChangedLog extends SimLog {
	readonly manaBefore: number;
	readonly manaAfter: number;
	readonly cause: string;

	constructor(params: SimLogParams, manaBefore: number, manaAfter: number, cause: string) {
		super(params);
		this.manaBefore = manaBefore;
		this.manaAfter = manaAfter;
		this.cause = cause;
	}

	static parse(params: SimLogParams): ManaChangedLog | null {
		const match = params.raw.match(/[Gained|Spent] \d+\.\d+ mana from (.*) \((\d+\.\d+) --> (\d+\.\d+)\)/);
		if (match && match[1]) {
			let cause = match[1];
			//if (cause.endsWith('s Regen')) {
			//	cause = 'Regen';
			//}
			return new ManaChangedLog(params, parseFloat(match[2]), parseFloat(match[3]), cause);
		} else {
			return null;
		}
	}
}
