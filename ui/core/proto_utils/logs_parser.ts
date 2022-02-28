import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { stringComparator, sum } from '/tbc/core/utils.js';

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

	equals(other: Entity) {
		return this.isTarget == other.isTarget && this.isPet == other.isPet && this.index == other.index && this.name == other.name;
	}

	toString(): string {
		if (this.isTarget) {
			return 'Target ' + (this.index + 1);
		} else if (this.isPet) {
			return `${this.ownerName} (#${this.index + 1}) - ${this.name}`;
		} else {
			return `${this.name} (#${this.index + 1})`;
		}
	}

	// Parses one or more Entities from a string.
	// Each entity label should be one of:
	//   'Target 1' if a target,
	//   'PlayerName (#1)' if a player, or
	//   'PlayerName (#1) - PetName' if a pet.
	static parseRegex = /\[(Target (\d+))|(([a-zA-Z0-9]+) \(#(\d+)\) - ([a-zA-Z0-9\s]+))|(([a-zA-Z0-9]+) \(#(\d+)\))\]/g;
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

	// Logs for auras that were active at this timestamp.
	// This is only filled if populateActiveAuras() is called.
	activeAuras: Array<AuraUptimeLog>;

	constructor(params: SimLogParams) {
		this.raw = params.raw;
		this.timestamp = params.timestamp;
		this.source = params.source;
		this.target = params.target;
		this.activeAuras = [];
	}

	toString(): string {
		return this.raw;
	}

	toStringPrefix(): string {
		const timestampStr = `[${this.timestamp.toFixed(2)}]`;
		if (this.source) {
			return `${timestampStr} [${this.source}]`;
		} else {
			return timestampStr;
		}
	}

	static async parseAll(result: RaidSimResult): Promise<Array<SimLog>> {
		const lines = result.logs.split('\n');

		return Promise.all(lines.map(line => {
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
					|| ResourceChangedLog.parse(params)
					|| AuraGainedLog.parse(params)
					|| AuraFadedLog.parse(params)
					|| MajorCooldownUsedLog.parse(params)
					|| CastBeganLog.parse(params)
					|| StatChangeLog.parse(params)
					|| Promise.resolve(new SimLog(params));
		}));
	}

	isDamageDealt(): this is DamageDealtLog {
		return this instanceof DamageDealtLog;
	}

	isResourceChanged(): this is ResourceChangedLog {
		return this instanceof ResourceChangedLog;
	}

	isAuraGained(): this is AuraGainedLog {
		return this instanceof AuraGainedLog;
	}

	isAuraFaded(): this is AuraFadedLog {
		return this instanceof AuraFadedLog;
	}

	isMajorCooldownUsed(): this is MajorCooldownUsedLog {
		return this instanceof MajorCooldownUsedLog;
	}

	isCastBegan(): this is CastBeganLog {
		return this instanceof CastBeganLog;
	}

	isStatChange(): this is StatChangeLog {
		return this instanceof StatChangeLog;
	}

	// Group events that happen at the same time.
	static groupDuplicateTimestamps<LogType extends SimLog>(logs: Array<LogType>): Array<Array<LogType>> {
		const grouped: Array<Array<LogType>> = [];
		let curGroup: Array<LogType> = [];

		logs.forEach(log => {
			if (curGroup.length == 0 || log.timestamp == curGroup[0].timestamp) {
				curGroup.push(log);
			} else {
				grouped.push(curGroup);
				curGroup = [log];
			}
		});
		if (curGroup.length > 0) {
			grouped.push(curGroup);
		}

		return grouped;
	}
}

export class DamageDealtLog extends SimLog {
	readonly amount: number;
	readonly miss: boolean;
	readonly hit: boolean;
	readonly crit: boolean;
	readonly glance: boolean;
	readonly dodge: boolean;
	readonly parry: boolean;
	readonly block: boolean;
	readonly tick: boolean;
	readonly partialResist1_4: boolean;
	readonly partialResist2_4: boolean;
	readonly partialResist3_4: boolean;
	readonly cause: ActionId;

	constructor(params: SimLogParams, amount: number, miss: boolean, crit: boolean, glance: boolean, dodge: boolean, parry: boolean, block: boolean, tick: boolean, partialResist1_4: boolean, partialResist2_4: boolean, partialResist3_4: boolean, cause: ActionId) {
		super(params);
		this.amount = amount;
		this.miss = miss;
		this.glance = glance;
		this.dodge = dodge;
		this.parry = parry;
		this.block = block;
		this.hit = !miss && !crit;
		this.crit = crit;
		this.tick = tick;
		this.partialResist1_4 = partialResist1_4;
		this.partialResist2_4 = partialResist2_4;
		this.partialResist3_4 = partialResist3_4;
		this.cause = cause;
	}

	resultString(): string {
		let result = this.miss ? 'Miss'
				: this.dodge ? 'Dodge'
				: this.parry ? 'Parry'
				: this.glance ? 'Glance'
				: this.crit ? 'Crit'
				: this.block ? 'Block'
				: this.tick ? 'Tick'
				: 'Hit';
		if (!this.miss && !this.dodge && !this.parry) {
			result += ` for ${this.amount.toFixed(2)}`;
			if (this.partialResist1_4) {
				result += ' (25% Resist)';
			} else if (this.partialResist2_4) {
				result += ' (50% Resist)';
			} else if (this.partialResist3_4) {
				result += ' (75% Resist)';
			}
			result += '.'
		}
		return result;
	}

	toString(): string {
		return `${this.toStringPrefix()} ${this.cause.name} ${this.resultString()}`;
	}

	static parse(params: SimLogParams): Promise<DamageDealtLog> | null {
		const match = params.raw.match(/] (.*?) (tick )?((Miss)|(Hit)|(Crit)|(Glance)|(Dodge)|(Parry)|(Block))( \((\d+)% Resist\))?( for (\d+\.\d+) damage)?/);
		if (match) {
			return ActionId.fromLogString(match[1]).fill(params.source?.index).then(cause => {
				let amount = 0;
				if (match[14]) {
					amount = parseFloat(match[14]);
				}

				return new DamageDealtLog(
						params,
						amount,
						match[3] == 'Miss',
						match[3] == 'Crit',
						match[3] == 'Glance',
						match[3] == 'Dodge',
						match[3] == 'Parry',
						match[3] == 'Block',
						Boolean(match[2]) && match[2].includes('tick'),
						match[12] == '25',
						match[12] == '50',
						match[12] == '75',
						cause);
			});
		} else {
			return null;
		}
	}
}

export class DpsLog extends SimLog {
	readonly dps: number;

	// Damage events that occurred at the same time as this log.
	readonly damageLogs: Array<DamageDealtLog>;

	constructor(params: SimLogParams, dps: number, damageLogs: Array<DamageDealtLog>) {
		super(params);
		this.dps = dps;
		this.damageLogs = damageLogs;
	}

	static DPS_WINDOW = 15; // Window over which to calculate DPS.
	static fromLogs(damageDealtLogs: Array<DamageDealtLog>): Array<DpsLog> {
		const groupedDamageLogs = SimLog.groupDuplicateTimestamps(damageDealtLogs);

		let curDamageLogs: Array<DamageDealtLog> = [];
		let curDamageTotal = 0;

		return groupedDamageLogs.map(ddLogGroup => {
			ddLogGroup.forEach(ddLog => {
				curDamageLogs.push(ddLog);
				curDamageTotal += ddLog.amount;
			});

			const newStartIdx = curDamageLogs.findIndex(curLog => {
				const inWindow = curLog.timestamp > ddLogGroup[0].timestamp - DpsLog.DPS_WINDOW;
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
			if (isNaN(dps)) {
				console.warn('NaN dps!');
			}

			return new DpsLog({
				raw: '',
				timestamp: ddLogGroup[0].timestamp,
				source: ddLogGroup[0].source,
				target: null,
			}, dps, ddLogGroup);
		});
	}
}

export class AuraGainedLog extends SimLog {
	readonly aura: ActionId;

	constructor(params: SimLogParams, aura: ActionId) {
		super(params);
		this.aura = aura;
	}

	toString(): string {
		return `${this.toStringPrefix()} Aura gained: ${this.aura.name}.`;
	}

	static parse(params: SimLogParams): Promise<AuraGainedLog> | null {
		const match = params.raw.match(/Aura gained: (.*)/);
		if (match && match[1]) {
			return ActionId.fromLogString(match[1]).fill(params.source?.index).then(aura => new AuraGainedLog(params, aura));
		} else {
			return null;
		}
	}
}

export class AuraFadedLog extends SimLog {
	readonly aura: ActionId;

	constructor(params: SimLogParams, aura: ActionId) {
		super(params);
		this.aura = aura;
	}

	toString(): string {
		return `${this.toStringPrefix()} Aura faded: ${this.aura.name}.`;
	}

	static parse(params: SimLogParams): Promise<AuraFadedLog> | null {
		const match = params.raw.match(/Aura faded: (.*)/);
		if (match && match[1]) {
			return ActionId.fromLogString(match[1]).fill(params.source?.index).then(aura => new AuraFadedLog(params, aura));
		} else {
			return null;
		}
	}
}

export class AuraUptimeLog extends SimLog {
	readonly gainedAt: number;
	readonly fadedAt: number;
	readonly aura: ActionId;

	constructor(params: SimLogParams, fadedAt: number, aura: ActionId) {
		super(params);
		this.gainedAt = params.timestamp;
		this.fadedAt = fadedAt;
		this.aura = aura;
	}

	static fromLogs(logs: Array<SimLog>, entity: Entity, encounterDuration: number): Array<AuraUptimeLog> {
		let unmatchedGainedLogs: Array<AuraGainedLog> = [];
		const uptimeLogs: Array<AuraUptimeLog> = [];

		logs.forEach(log => {
			if (!log.source || !log.source.equals(entity)) {
				return;
			}
			if (log.isAuraGained()) {
				unmatchedGainedLogs.push(log);
				return;
			}
			if (!log.isAuraFaded()) {
				return;
			}

			const matchingGainedIdx = unmatchedGainedLogs.findIndex(gainedLog => gainedLog.aura.equals(log.aura));
			if (matchingGainedIdx == -1) {
				console.warn('Unmatched aura faded log: ' + log.aura.name);
				return;
			}
			const gainedLog = unmatchedGainedLogs.splice(matchingGainedIdx, 1)[0];

			uptimeLogs.push(new AuraUptimeLog({
				raw: log.raw,
				timestamp: gainedLog.timestamp,
				source: log.source,
				target: log.target,
			}, log.timestamp, gainedLog.aura));
		});

		// Auras active at the end won't have a faded log, so need to add them separately.
		unmatchedGainedLogs.forEach(gainedLog => {
			uptimeLogs.push(new AuraUptimeLog({
				raw: gainedLog.raw,
				timestamp: gainedLog.timestamp,
				source: gainedLog.source,
				target: gainedLog.target,
			}, encounterDuration, gainedLog.aura));
		});

		uptimeLogs.sort((a, b) => a.gainedAt - b.gainedAt);
		return uptimeLogs;
	}

	// Populates the activeAuras field for all logs using the provided auras.
	static populateActiveAuras(logs: Array<SimLog>, auraLogs: Array<AuraUptimeLog>) {
		let curAuras: Array<AuraUptimeLog> = [];
		let auraLogsIndex = 0;

		logs.forEach(log => {
			while (auraLogsIndex < auraLogs.length && auraLogs[auraLogsIndex].gainedAt <= log.timestamp) {
				curAuras.push(auraLogs[auraLogsIndex]);
				auraLogsIndex++;
			}
			curAuras = curAuras.filter(curAura => curAura.fadedAt > log.timestamp);

			const activeAuras = curAuras.slice();
			activeAuras.sort((a, b) => stringComparator(a.aura.name, b.aura.name));
			log.activeAuras = activeAuras;
		});
	}
}

export type Resource = 'mana' | 'energy' | 'focus' | 'rage';

export class ResourceChangedLog extends SimLog {
	readonly resource: Resource;
	readonly valueBefore: number;
	readonly valueAfter: number;
	readonly isSpend: boolean;
	readonly cause: ActionId;

	constructor(params: SimLogParams, resource: Resource, valueBefore: number, valueAfter: number, isSpend: boolean, cause: ActionId) {
		super(params);
		this.resource = resource;
		this.valueBefore = valueBefore;
		this.valueAfter = valueAfter;
		this.isSpend = isSpend;
		this.cause = cause;
	}

	toString(): string {
		const signedDiff = (this.valueAfter - this.valueBefore) * (this.isSpend ? -1 : 1);
		return `${this.toStringPrefix()} ${this.isSpend ? 'Spent' : 'Gained'} ${signedDiff.toFixed(1)} ${this.resource} from ${this.cause.name}. (${this.valueBefore.toFixed(1)} --> ${this.valueAfter.toFixed(1)})`;
	}

	resultString(): string {
		const delta = this.valueAfter - this.valueBefore;
		if (delta < 0) {
			return delta.toFixed(1);
		} else {
			return '+' + delta.toFixed(1);
		}
	}

	static parse(params: SimLogParams): Promise<ResourceChangedLog> | null {
		const match = params.raw.match(/((Gained)|(Spent)) \d+\.?\d* ((mana)|(energy)|(focus)|(rage)) from (.*) \((\d+\.?\d*) --> (\d+\.?\d*)\)/);
		if (match) {
			return ActionId.fromLogString(match[9]).fill(params.source?.index).then(cause => {
				return new ResourceChangedLog(params, match[4] as Resource, parseFloat(match[10]), parseFloat(match[11]), match[1] == 'Spent', cause);
			});
		} else {
			return null;
		}
	}
}

export class ResourceChangedLogGroup extends SimLog {
	readonly resource: Resource;
	readonly valueBefore: number;
	readonly valueAfter: number;
	readonly logs: Array<ResourceChangedLog>;

	constructor(params: SimLogParams, resource: Resource, valueBefore: number, valueAfter: number, logs: Array<ResourceChangedLog>) {
		super(params);
		this.resource = resource;
		this.valueBefore = valueBefore;
		this.valueAfter = valueAfter;
		this.logs = logs;
	}

	toString(): string {
		const capitalizedResource = this.resource.charAt(0).toUpperCase() + this.resource.slice(1);
		return `${this.toStringPrefix()} ${capitalizedResource}: ${this.valueBefore.toFixed(1)} --> ${this.valueAfter.toFixed(1)}`;
	}

	static fromLogs(logs: Array<SimLog>, resource: Resource): Array<ResourceChangedLogGroup> {
		const resourceChangedLogs = logs
				.filter((log): log is ResourceChangedLog => log.isResourceChanged())
				.filter(log => log.resource == resource);

		const groupedLogs = SimLog.groupDuplicateTimestamps(resourceChangedLogs);
		return groupedLogs.map(logGroup => new ResourceChangedLogGroup(
				{
					raw: '',
					timestamp: logGroup[0].timestamp,
					source: logGroup[0].source,
					target: logGroup[0].target,
				},
				resource,
				logGroup[0].valueBefore,
				logGroup[logGroup.length - 1].valueAfter,
				logGroup));
	}
}

export class MajorCooldownUsedLog extends SimLog {
	readonly cooldownId: ActionId;

	constructor(params: SimLogParams, cooldownId: ActionId) {
		super(params);
		this.cooldownId = cooldownId;
	}

	toString(): string {
		return `${this.toStringPrefix()} Major cooldown used: ${this.cooldownId.name}.`;
	}

	static parse(params: SimLogParams): Promise<MajorCooldownUsedLog> | null {
		const match = params.raw.match(/Major cooldown used: (.*)/);
		if (match) {
			return ActionId.fromLogString(match[1]).fill(params.source?.index).then(cooldownId => new MajorCooldownUsedLog(params, cooldownId));
		} else {
			return null;
		}
	}
}

export class CastBeganLog extends SimLog {
	readonly castId: ActionId;
	readonly manaCost: number;
	readonly castTime: number;

	constructor(params: SimLogParams, castId: ActionId, manaCost: number, castTime: number) {
		super(params);
		this.castId = castId;
		this.manaCost = manaCost;
		this.castTime = castTime;
	}

	toString(): string {
		return `${this.toStringPrefix()} Casting ${this.castId.name} (Cast time = ${this.castTime.toFixed(2)}s, Cost = ${this.manaCost.toFixed(1)}).`;
	}

	static parse(params: SimLogParams): Promise<CastBeganLog> | null {
		const match = params.raw.match(/Casting (.*) \(Cost = (\d+\.?\d*), Cast Time = (\d+\.?\d*)(m?s)\)/);
		if (match) {
			let castTime = parseFloat(match[3]);
			if (match[4] == 'ms') {
				castTime /= 1000;
			}
			return ActionId.fromLogString(match[1]).fill(params.source?.index).then(castId => new CastBeganLog(params, castId, parseFloat(match[2]), castTime));
		} else {
			return null;
		}
	}
}

export class StatChangeLog extends SimLog {
	readonly effectId: ActionId;
	readonly amount: number;
	readonly stat: string;

	constructor(params: SimLogParams, effectId: ActionId, amount: number, stat: string) {
		super(params);
		this.effectId = effectId;
		this.amount = amount;
		this.stat = stat;
	}

	toString(): string {
		if (this.amount > 0) {
			return `${this.toStringPrefix()} Gained ${this.amount.toFixed(0)} ${this.stat} from ${this.effectId.name}.`;
		} else {
			return `${this.toStringPrefix()} Lost ${(-this.amount).toFixed(0)} ${this.stat} from fading ${this.effectId.name}.`;
		}
	}

	static parse(params: SimLogParams): Promise<StatChangeLog> | null {
		const match = params.raw.match(/((Gained)|(Lost)) (\d+\.?\d*) (.*) from (fading )?(.*)/);
		if (match) {
			return ActionId.fromLogString(match[7]).fill(params.source?.index).then(effectId => {
				const sign = match[1] == 'Lost' ? -1 : 1;
				return new StatChangeLog(params, effectId, parseFloat(match[4]) * sign, match[5]);
			});
		} else {
			return null;
		}
	}
}
