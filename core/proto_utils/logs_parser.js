import { ResourceType } from '/tbc/core/proto/api.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { resourceNames, stringToResourceType } from '/tbc/core/proto_utils/names.js';
import { bucket, getEnumValues, stringComparator, sum } from '/tbc/core/utils.js';
export class Entity {
    constructor(name, ownerName, index, isTarget, isPet) {
        this.name = name;
        this.ownerName = ownerName;
        this.index = index;
        this.isTarget = isTarget;
        this.isPet = isPet;
    }
    equals(other) {
        return this.isTarget == other.isTarget && this.isPet == other.isPet && this.index == other.index && this.name == other.name;
    }
    toString() {
        if (this.isTarget) {
            return 'Target ' + (this.index + 1);
        }
        else if (this.isPet) {
            return `${this.ownerName} (#${this.index + 1}) - ${this.name}`;
        }
        else {
            return `${this.name} (#${this.index + 1})`;
        }
    }
    static parseAll(str) {
        return Array.from(str.matchAll(Entity.parseRegex)).map(match => {
            if (match[1]) {
                return new Entity(match[1], '', parseInt(match[2]) - 1, true, false);
            }
            else if (match[3]) {
                return new Entity(match[6], match[4], parseInt(match[5]) - 1, false, true);
            }
            else if (match[7]) {
                return new Entity(match[8], '', parseInt(match[9]) - 1, false, false);
            }
            else {
                throw new Error('Invalid Entity match');
            }
        });
    }
}
// Parses one or more Entities from a string.
// Each entity label should be one of:
//   'Target 1' if a target,
//   'PlayerName (#1)' if a player, or
//   'PlayerName (#1) - PetName' if a pet.
Entity.parseRegex = /\[(Target (\d+))|(([a-zA-Z0-9]+) \(#(\d+)\) - ([a-zA-Z0-9\s]+))|(([a-zA-Z0-9]+) \(#(\d+)\))\]/g;
export class SimLog {
    constructor(params) {
        this.raw = params.raw;
        this.logIndex = params.logIndex;
        this.timestamp = params.timestamp;
        this.source = params.source;
        this.target = params.target;
        this.actionId = params.actionId;
        this.threat = params.threat;
        this.activeAuras = [];
    }
    toString() {
        return this.raw;
    }
    toStringPrefix() {
        const timestampStr = `[${this.timestamp.toFixed(2)}]`;
        if (this.source) {
            return `${timestampStr} [${this.source}]`;
        }
        else {
            return timestampStr;
        }
    }
    static async parseAll(result) {
        const lines = result.logs.split('\n');
        return Promise.all(lines.map((line, lineIndex) => {
            const params = {
                raw: line,
                logIndex: lineIndex,
                timestamp: 0,
                source: null,
                target: null,
                actionId: null,
                threat: 0,
            };
            const threatMatch = line.match(/ \(Threat: (-?[0-9]+\.[0-9]+)\)/);
            if (threatMatch) {
                params.threat = parseFloat(threatMatch[1]);
                line = line.substring(0, threatMatch.index);
            }
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
                || AuraEventLog.parse(params)
                || AuraStacksChangeLog.parse(params)
                || MajorCooldownUsedLog.parse(params)
                || CastBeganLog.parse(params)
                || CastCompletedLog.parse(params)
                || StatChangeLog.parse(params)
                || Promise.resolve(new SimLog(params));
        }));
    }
    isDamageDealt() {
        return this instanceof DamageDealtLog;
    }
    isResourceChanged() {
        return this instanceof ResourceChangedLog;
    }
    isAuraEvent() {
        return this instanceof AuraEventLog;
    }
    isAuraStacksChange() {
        return this instanceof AuraStacksChangeLog;
    }
    isMajorCooldownUsed() {
        return this instanceof MajorCooldownUsedLog;
    }
    isCastBegan() {
        return this instanceof CastBeganLog;
    }
    isCastCompleted() {
        return this instanceof CastCompletedLog;
    }
    isStatChange() {
        return this instanceof StatChangeLog;
    }
    // Group events that happen at the same time.
    static groupDuplicateTimestamps(logs) {
        const grouped = [];
        let curGroup = [];
        logs.forEach(log => {
            if (curGroup.length == 0 || log.timestamp == curGroup[0].timestamp) {
                curGroup.push(log);
            }
            else {
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
    constructor(params, amount, miss, crit, crush, glance, dodge, parry, block, tick, partialResist1_4, partialResist2_4, partialResist3_4) {
        super(params);
        this.amount = amount;
        this.miss = miss;
        this.glance = glance;
        this.dodge = dodge;
        this.parry = parry;
        this.block = block;
        this.hit = !miss && !crit;
        this.crit = crit;
        this.crush = crush;
        this.tick = tick;
        this.partialResist1_4 = partialResist1_4;
        this.partialResist2_4 = partialResist2_4;
        this.partialResist3_4 = partialResist3_4;
    }
    resultString() {
        let result = this.miss ? 'Miss'
            : this.dodge ? 'Dodge'
                : this.parry ? 'Parry'
                    : this.glance ? 'Glance'
                        : this.block ? (this.crit ? 'Critical Block' : 'Block')
                            : this.crit ? 'Crit'
                                : this.crush ? 'Crush'
                                    : this.tick ? 'Tick'
                                        : 'Hit';
        if (!this.miss && !this.dodge && !this.parry) {
            result += ` for ${this.amount.toFixed(2)}`;
            if (this.partialResist1_4) {
                result += ' (25% Resist)';
            }
            else if (this.partialResist2_4) {
                result += ' (50% Resist)';
            }
            else if (this.partialResist3_4) {
                result += ' (75% Resist)';
            }
            result += '.';
        }
        return result;
    }
    toString() {
        return `${this.toStringPrefix()} ${this.actionId.name} ${this.resultString()} (${this.threat.toFixed(2)} Threat)`;
    }
    static parse(params) {
        const match = params.raw.match(/] (.*?) (tick )?((Miss)|(Hit)|(Crit)|(Crush)|(Glance)|(Dodge)|(Parry)|(Block)|(CriticalBlock))( \((\d+)% Resist\))?( for (\d+\.\d+) damage)?/);
        if (match) {
            return ActionId.fromLogString(match[1]).fill(params.source?.index).then(cause => {
                params.actionId = cause;
                let amount = 0;
                if (match[16]) {
                    amount = parseFloat(match[16]);
                }
                return new DamageDealtLog(params, amount, match[3] == 'Miss', match[3] == 'Crit' || match[3] == 'CriticalBlock', match[3] == 'Crush', match[3] == 'Glance', match[3] == 'Dodge', match[3] == 'Parry', match[3] == 'Block' || match[3] == 'CriticalBlock', Boolean(match[2]) && match[2].includes('tick'), match[14] == '25', match[14] == '50', match[14] == '75');
            });
        }
        else {
            return null;
        }
    }
}
export class DpsLog extends SimLog {
    constructor(params, dps, damageLogs) {
        super(params);
        this.dps = dps;
        this.damageLogs = damageLogs;
    }
    static fromLogs(damageDealtLogs) {
        const groupedDamageLogs = SimLog.groupDuplicateTimestamps(damageDealtLogs);
        let curDamageLogs = [];
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
            }
            else {
                curDamageLogs = curDamageLogs.slice(newStartIdx);
            }
            const dps = curDamageTotal / DpsLog.DPS_WINDOW;
            if (isNaN(dps)) {
                console.warn('NaN dps!');
            }
            return new DpsLog({
                raw: '',
                logIndex: ddLogGroup[0].logIndex,
                timestamp: ddLogGroup[0].timestamp,
                source: ddLogGroup[0].source,
                target: null,
                actionId: null,
                threat: 0,
            }, dps, ddLogGroup);
        });
    }
}
DpsLog.DPS_WINDOW = 15; // Window over which to calculate DPS.
export class ThreatLogGroup extends SimLog {
    constructor(params, threatBefore, threatAfter, logs) {
        super(params);
        this.threatBefore = threatBefore;
        this.threatAfter = threatAfter;
        this.logs = logs;
    }
    static fromLogs(logs) {
        const groupedLogs = SimLog.groupDuplicateTimestamps(logs.filter(log => log.threat != 0));
        let curThreat = 0;
        return groupedLogs.map(logGroup => {
            const newThreat = sum(logGroup.map(log => log.threat));
            const threatLog = new ThreatLogGroup({
                raw: '',
                logIndex: logGroup[0].logIndex,
                timestamp: logGroup[0].timestamp,
                source: logGroup[0].source,
                target: logGroup[0].target,
                actionId: null,
                threat: newThreat,
            }, curThreat, curThreat + newThreat, logGroup);
            curThreat += newThreat;
            return threatLog;
        });
    }
}
export class AuraEventLog extends SimLog {
    constructor(params, isGained, isFaded, isRefreshed) {
        super(params);
        this.isGained = isGained;
        this.isFaded = isFaded;
        this.isRefreshed = isRefreshed;
    }
    toString() {
        return `${this.toStringPrefix()} Aura ${this.isGained ? 'gained' : this.isFaded ? 'faded' : 'refreshed'}: ${this.actionId.name}.`;
    }
    static parse(params) {
        const match = params.raw.match(/Aura ((gained)|(faded)|(refreshed)): (.*)/);
        if (match && match[5]) {
            return ActionId.fromLogString(match[5]).fill(params.source?.index).then(aura => {
                params.actionId = aura;
                const event = match[1];
                return new AuraEventLog(params, event == 'gained', event == 'faded', event == 'refreshed');
            });
        }
        else {
            return null;
        }
    }
}
export class AuraStacksChangeLog extends SimLog {
    constructor(params, oldStacks, newStacks) {
        super(params);
        this.oldStacks = oldStacks;
        this.newStacks = newStacks;
    }
    toString() {
        return `${this.toStringPrefix()} ${this.actionId.name} stacks: ${this.oldStacks} --> ${this.newStacks}.`;
    }
    static parse(params) {
        const match = params.raw.match(/(.*) stacks: ([0-9]+) --> ([0-9]+)/);
        if (match && match[1]) {
            return ActionId.fromLogString(match[1]).fill(params.source?.index).then(aura => {
                params.actionId = aura;
                return new AuraStacksChangeLog(params, parseInt(match[2]), parseInt(match[3]));
            });
        }
        else {
            return null;
        }
    }
}
export class AuraUptimeLog extends SimLog {
    constructor(params, fadedAt) {
        super(params);
        this.gainedAt = params.timestamp;
        this.fadedAt = fadedAt;
    }
    static fromLogs(logs, entity, encounterDuration) {
        let unmatchedGainedLogs = [];
        const uptimeLogs = [];
        logs.forEach((log) => {
            if (!log.source || !log.source.equals(entity) || !log.isAuraEvent()) {
                return;
            }
            if (log.isGained) {
                unmatchedGainedLogs.push(log);
                return;
            }
            const matchingGainedIdx = unmatchedGainedLogs.findIndex(gainedLog => gainedLog.actionId.equals(log.actionId));
            if (matchingGainedIdx == -1) {
                console.warn('Unmatched aura faded log: ' + log.actionId.name);
                return;
            }
            const gainedLog = unmatchedGainedLogs.splice(matchingGainedIdx, 1)[0];
            uptimeLogs.push(new AuraUptimeLog({
                raw: log.raw,
                logIndex: gainedLog.logIndex,
                timestamp: gainedLog.timestamp,
                source: log.source,
                target: log.target,
                actionId: gainedLog.actionId,
                threat: gainedLog.threat,
            }, log.timestamp));
            if (log.isRefreshed) {
                unmatchedGainedLogs.push(log);
            }
        });
        // Auras active at the end won't have a faded log, so need to add them separately.
        unmatchedGainedLogs.forEach(gainedLog => {
            uptimeLogs.push(new AuraUptimeLog({
                raw: gainedLog.raw,
                logIndex: gainedLog.logIndex,
                timestamp: gainedLog.timestamp,
                source: gainedLog.source,
                target: gainedLog.target,
                actionId: gainedLog.actionId,
                threat: gainedLog.threat,
            }, encounterDuration));
        });
        uptimeLogs.sort((a, b) => a.gainedAt - b.gainedAt);
        return uptimeLogs;
    }
    // Populates the activeAuras field for all logs using the provided auras.
    static populateActiveAuras(logs, auraLogs) {
        let curAuras = [];
        let auraLogsIndex = 0;
        logs.forEach(log => {
            while (auraLogsIndex < auraLogs.length && auraLogs[auraLogsIndex].gainedAt <= log.timestamp) {
                curAuras.push(auraLogs[auraLogsIndex]);
                auraLogsIndex++;
            }
            curAuras = curAuras.filter(curAura => curAura.fadedAt > log.timestamp);
            const activeAuras = curAuras.slice();
            activeAuras.sort((a, b) => stringComparator(a.actionId.name, b.actionId.name));
            log.activeAuras = activeAuras;
        });
    }
}
export class ResourceChangedLog extends SimLog {
    constructor(params, resourceType, valueBefore, valueAfter, isSpend) {
        super(params);
        this.resourceType = resourceType;
        this.valueBefore = valueBefore;
        this.valueAfter = valueAfter;
        this.isSpend = isSpend;
    }
    toString() {
        const signedDiff = (this.valueAfter - this.valueBefore) * (this.isSpend ? -1 : 1);
        return `${this.toStringPrefix()} ${this.isSpend ? 'Spent' : 'Gained'} ${signedDiff.toFixed(1)} ${resourceNames[this.resourceType]} from ${this.actionId.name}. (${this.valueBefore.toFixed(1)} --> ${this.valueAfter.toFixed(1)})`;
    }
    resultString() {
        const delta = this.valueAfter - this.valueBefore;
        if (delta < 0) {
            return delta.toFixed(1);
        }
        else {
            return '+' + delta.toFixed(1);
        }
    }
    static parse(params) {
        const match = params.raw.match(/((Gained)|(Spent)) \d+\.?\d* ((mana)|(energy)|(focus)|(rage)|(combo points)) from (.*) \((\d+\.?\d*) --> (\d+\.?\d*)\)/);
        if (match) {
            const resourceType = stringToResourceType(match[4]);
            return ActionId.fromLogString(match[10]).fill(params.source?.index).then(cause => {
                params.actionId = cause;
                return new ResourceChangedLog(params, resourceType, parseFloat(match[11]), parseFloat(match[12]), match[1] == 'Spent');
            });
        }
        else {
            return null;
        }
    }
}
export class ResourceChangedLogGroup extends SimLog {
    constructor(params, resourceType, valueBefore, valueAfter, logs) {
        super(params);
        this.resourceType = resourceType;
        this.valueBefore = valueBefore;
        this.valueAfter = valueAfter;
        this.logs = logs;
    }
    toString() {
        return `${this.toStringPrefix()} ${resourceNames[this.resourceType]}: ${this.valueBefore.toFixed(1)} --> ${this.valueAfter.toFixed(1)}`;
    }
    static fromLogs(logs) {
        const allResourceChangedLogs = logs.filter((log) => log.isResourceChanged());
        const results = {};
        const resourceTypes = getEnumValues(ResourceType).filter(val => val != ResourceType.ResourceTypeNone);
        resourceTypes.forEach(resourceType => {
            const resourceChangedLogs = allResourceChangedLogs.filter(log => log.resourceType == resourceType);
            const groupedLogs = SimLog.groupDuplicateTimestamps(resourceChangedLogs);
            results[resourceType] = groupedLogs.map(logGroup => new ResourceChangedLogGroup({
                raw: '',
                logIndex: logGroup[0].logIndex,
                timestamp: logGroup[0].timestamp,
                source: logGroup[0].source,
                target: logGroup[0].target,
                actionId: null,
                threat: 0,
            }, resourceType, logGroup[0].valueBefore, logGroup[logGroup.length - 1].valueAfter, logGroup));
        });
        return results;
    }
}
export class MajorCooldownUsedLog extends SimLog {
    constructor(params) {
        super(params);
    }
    toString() {
        return `${this.toStringPrefix()} Major cooldown used: ${this.actionId.name}.`;
    }
    static parse(params) {
        const match = params.raw.match(/Major cooldown used: (.*)/);
        if (match) {
            return ActionId.fromLogString(match[1]).fill(params.source?.index).then(cooldownId => {
                params.actionId = cooldownId;
                return new MajorCooldownUsedLog(params);
            });
        }
        else {
            return null;
        }
    }
}
export class CastBeganLog extends SimLog {
    constructor(params, manaCost, castTime) {
        super(params);
        this.manaCost = manaCost;
        this.castTime = castTime;
    }
    toString() {
        return `${this.toStringPrefix()} Casting ${this.actionId.name} (Cast time = ${this.castTime.toFixed(2)}s, Cost = ${this.manaCost.toFixed(1)}).`;
    }
    static parse(params) {
        const match = params.raw.match(/Casting (.*) \(Cost = (\d+\.?\d*), Cast Time = (\d+\.?\d*)(m?s)\)/);
        if (match) {
            let castTime = parseFloat(match[3]);
            if (match[4] == 'ms') {
                castTime /= 1000;
            }
            return ActionId.fromLogString(match[1]).fill(params.source?.index).then(castId => {
                params.actionId = castId;
                return new CastBeganLog(params, parseFloat(match[2]), castTime);
            });
        }
        else {
            return null;
        }
    }
}
export class CastCompletedLog extends SimLog {
    constructor(params) {
        super(params);
    }
    toString() {
        return `${this.toStringPrefix()} Completed cast ${this.actionId.name}.`;
    }
    static parse(params) {
        const match = params.raw.match(/Completed cast (.*)/);
        if (match) {
            return ActionId.fromLogString(match[1]).fill(params.source?.index).then(castId => {
                params.actionId = castId;
                return new CastCompletedLog(params);
            });
        }
        else {
            return null;
        }
    }
}
export class CastLog extends SimLog {
    constructor(castBeganLog, castCompletedLog, damageDealtLogs) {
        super({
            raw: castBeganLog.raw,
            logIndex: castBeganLog.logIndex,
            timestamp: castBeganLog.timestamp,
            source: castBeganLog.source,
            target: castBeganLog.target,
            actionId: castCompletedLog?.actionId || castBeganLog.actionId,
            threat: castCompletedLog?.threat || castBeganLog.threat,
        });
        this.castTime = castBeganLog.castTime;
        this.castBeganLog = castBeganLog;
        this.castCompletedLog = castCompletedLog;
        this.damageDealtLogs = damageDealtLogs;
    }
    toString() {
        return `${this.toStringPrefix()} Casting ${this.actionId.name} (Cast time = ${this.castTime.toFixed(2)}s).`;
    }
    static fromLogs(logs) {
        const castBeganLogs = logs.filter((log) => log.isCastBegan());
        const castCompletedLogs = logs.filter((log) => log.isCastCompleted());
        const damageDealtLogs = logs.filter((log) => log.isDamageDealt());
        const toBucketKey = (actionId) => {
            if (actionId.spellId == 30451) {
                // Arcane Blast is unique because it can finish its cast as a different
                // spell than it started (if stacks drop).
                return actionId.toStringIgnoringTag();
            }
            else {
                return actionId.toString();
            }
        };
        const castBeganLogsByAbility = bucket(castBeganLogs, log => toBucketKey(log.actionId));
        const castCompletedLogsByAbility = bucket(castCompletedLogs, log => toBucketKey(log.actionId));
        const damageDealtLogsByAbility = bucket(damageDealtLogs, log => toBucketKey(log.actionId));
        const castLogs = [];
        Object.keys(castBeganLogsByAbility).forEach(bucketKey => {
            const abilityCastsBegan = castBeganLogsByAbility[bucketKey];
            const abilityCastsCompleted = castCompletedLogsByAbility[bucketKey];
            const abilityDamageDealt = damageDealtLogsByAbility[bucketKey];
            const actionId = abilityCastsBegan[0].actionId;
            const getCastCompleted = (cbIndex) => {
                if (!abilityCastsCompleted) {
                    return null;
                }
                if (cbIndex >= abilityCastsBegan.length) {
                    return null;
                }
                const nextBeganIndex = abilityCastsBegan[cbIndex + 1]?.logIndex || null;
                return abilityCastsCompleted.find(ccl => ccl.logIndex > abilityCastsBegan[cbIndex].logIndex
                    && (nextBeganIndex == null || ccl.logIndex < nextBeganIndex)
                    && toBucketKey(ccl.actionId) == toBucketKey(actionId)) || null;
            };
            if (!abilityDamageDealt) {
                abilityCastsBegan.forEach((castBegan, i) => castLogs.push(new CastLog(castBegan, getCastCompleted(i), [])));
                return;
            }
            let curCCL = null;
            let nextCCL = getCastCompleted(0);
            let curDamageDealtLogs = [];
            let curCbIndex = -1;
            abilityDamageDealt.forEach(ddLog => {
                if (curCbIndex >= abilityCastsBegan.length) {
                    return;
                }
                if (nextCCL == null || ddLog.logIndex < nextCCL.logIndex) {
                    curDamageDealtLogs.push(ddLog);
                }
                else {
                    if (curCbIndex != -1) {
                        castLogs.push(new CastLog(abilityCastsBegan[curCbIndex], curCCL, curDamageDealtLogs));
                    }
                    curDamageDealtLogs = [ddLog];
                    curCbIndex++;
                    if (curCbIndex >= abilityCastsBegan.length) {
                        return;
                    }
                    curCCL = nextCCL;
                    nextCCL = getCastCompleted(curCbIndex + 1);
                }
            });
            castLogs.push(new CastLog(abilityCastsBegan[curCbIndex], curCCL, curDamageDealtLogs));
        });
        castLogs.sort((a, b) => a.timestamp - b.timestamp);
        return castLogs;
    }
}
export class StatChangeLog extends SimLog {
    constructor(params, isGain, stats) {
        super(params);
        this.isGain = isGain;
        this.stats = stats;
    }
    toString() {
        if (this.isGain) {
            return `${this.toStringPrefix()} Gained ${this.stats} from ${this.actionId.name}.`;
        }
        else {
            return `${this.toStringPrefix()} Lost ${this.stats} from fading ${this.actionId.name}.`;
        }
    }
    static parse(params) {
        const match = params.raw.match(/((Gained)|(Lost)) ({.*}) from (fading )?(.*)/);
        if (match) {
            return ActionId.fromLogString(match[6]).fill(params.source?.index).then(effectId => {
                params.actionId = effectId;
                const sign = match[1] == 'Lost' ? -1 : 1;
                return new StatChangeLog(params, sign == 1, match[4]);
            });
        }
        else {
            return null;
        }
    }
}
