import { ResourceType } from '/tbc/core/proto/api.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { resourceNames, stringToResourceType } from '/tbc/core/proto_utils/names.js';
import { bucket, getEnumValues, stringComparator } from '/tbc/core/utils.js';
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
                || AuraRefreshedLog.parse(params)
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
    isAuraGained() {
        return this instanceof AuraGainedLog;
    }
    isAuraFaded() {
        return this instanceof AuraFadedLog;
    }
    isAuraRefreshed() {
        return this instanceof AuraRefreshedLog;
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
    constructor(params, amount, miss, crit, glance, dodge, parry, block, tick, partialResist1_4, partialResist2_4, partialResist3_4, cause) {
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
    resultString() {
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
        return `${this.toStringPrefix()} ${this.cause.name} ${this.resultString()}`;
    }
    static parse(params) {
        const match = params.raw.match(/] (.*?) (tick )?((Miss)|(Hit)|(Crit)|(Glance)|(Dodge)|(Parry)|(Block))( \((\d+)% Resist\))?( for (\d+\.\d+) damage)?/);
        if (match) {
            return ActionId.fromLogString(match[1]).fill(params.source?.index).then(cause => {
                let amount = 0;
                if (match[14]) {
                    amount = parseFloat(match[14]);
                }
                return new DamageDealtLog(params, amount, match[3] == 'Miss', match[3] == 'Crit', match[3] == 'Glance', match[3] == 'Dodge', match[3] == 'Parry', match[3] == 'Block', Boolean(match[2]) && match[2].includes('tick'), match[12] == '25', match[12] == '50', match[12] == '75', cause);
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
            }, dps, ddLogGroup);
        });
    }
}
DpsLog.DPS_WINDOW = 15; // Window over which to calculate DPS.
export class AuraGainedLog extends SimLog {
    constructor(params, aura) {
        super(params);
        this.aura = aura;
    }
    toString() {
        return `${this.toStringPrefix()} Aura gained: ${this.aura.name}.`;
    }
    static parse(params) {
        const match = params.raw.match(/Aura gained: (.*)/);
        if (match && match[1]) {
            return ActionId.fromLogString(match[1]).fill(params.source?.index).then(aura => new AuraGainedLog(params, aura));
        }
        else {
            return null;
        }
    }
}
export class AuraFadedLog extends SimLog {
    constructor(params, aura) {
        super(params);
        this.aura = aura;
    }
    toString() {
        return `${this.toStringPrefix()} Aura faded: ${this.aura.name}.`;
    }
    static parse(params) {
        const match = params.raw.match(/Aura faded: (.*)/);
        if (match && match[1]) {
            return ActionId.fromLogString(match[1]).fill(params.source?.index).then(aura => new AuraFadedLog(params, aura));
        }
        else {
            return null;
        }
    }
}
export class AuraRefreshedLog extends SimLog {
    constructor(params, aura) {
        super(params);
        this.aura = aura;
    }
    toString() {
        return `${this.toStringPrefix()} Aura refreshed: ${this.aura.name}.`;
    }
    static parse(params) {
        const match = params.raw.match(/Aura refreshed: (.*)/);
        if (match && match[1]) {
            return ActionId.fromLogString(match[1]).fill(params.source?.index).then(aura => new AuraRefreshedLog(params, aura));
        }
        else {
            return null;
        }
    }
}
export class AuraUptimeLog extends SimLog {
    constructor(params, fadedAt, aura) {
        super(params);
        this.gainedAt = params.timestamp;
        this.fadedAt = fadedAt;
        this.aura = aura;
    }
    static fromLogs(logs, entity, encounterDuration) {
        let unmatchedGainedLogs = [];
        const uptimeLogs = [];
        logs.forEach(log => {
            if (!log.source || !log.source.equals(entity)) {
                return;
            }
            if (log.isAuraGained()) {
                unmatchedGainedLogs.push(log);
                return;
            }
            if (!log.isAuraFaded() && !log.isAuraRefreshed()) {
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
                logIndex: gainedLog.logIndex,
                timestamp: gainedLog.timestamp,
                source: log.source,
                target: log.target,
            }, log.timestamp, gainedLog.aura));
            if (log.isAuraRefreshed()) {
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
            }, encounterDuration, gainedLog.aura));
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
            activeAuras.sort((a, b) => stringComparator(a.aura.name, b.aura.name));
            log.activeAuras = activeAuras;
        });
    }
}
export class ResourceChangedLog extends SimLog {
    constructor(params, resourceType, valueBefore, valueAfter, isSpend, cause) {
        super(params);
        this.resourceType = resourceType;
        this.valueBefore = valueBefore;
        this.valueAfter = valueAfter;
        this.isSpend = isSpend;
        this.cause = cause;
    }
    toString() {
        const signedDiff = (this.valueAfter - this.valueBefore) * (this.isSpend ? -1 : 1);
        return `${this.toStringPrefix()} ${this.isSpend ? 'Spent' : 'Gained'} ${signedDiff.toFixed(1)} ${resourceNames[this.resourceType]} from ${this.cause.name}. (${this.valueBefore.toFixed(1)} --> ${this.valueAfter.toFixed(1)})`;
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
                return new ResourceChangedLog(params, resourceType, parseFloat(match[11]), parseFloat(match[12]), match[1] == 'Spent', cause);
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
            }, resourceType, logGroup[0].valueBefore, logGroup[logGroup.length - 1].valueAfter, logGroup));
        });
        return results;
    }
}
export class MajorCooldownUsedLog extends SimLog {
    constructor(params, cooldownId) {
        super(params);
        this.cooldownId = cooldownId;
    }
    toString() {
        return `${this.toStringPrefix()} Major cooldown used: ${this.cooldownId.name}.`;
    }
    static parse(params) {
        const match = params.raw.match(/Major cooldown used: (.*)/);
        if (match) {
            return ActionId.fromLogString(match[1]).fill(params.source?.index).then(cooldownId => new MajorCooldownUsedLog(params, cooldownId));
        }
        else {
            return null;
        }
    }
}
export class CastBeganLog extends SimLog {
    constructor(params, castId, manaCost, castTime) {
        super(params);
        this.castId = castId;
        this.manaCost = manaCost;
        this.castTime = castTime;
    }
    toString() {
        return `${this.toStringPrefix()} Casting ${this.castId.name} (Cast time = ${this.castTime.toFixed(2)}s, Cost = ${this.manaCost.toFixed(1)}).`;
    }
    static parse(params) {
        const match = params.raw.match(/Casting (.*) \(Cost = (\d+\.?\d*), Cast Time = (\d+\.?\d*)(m?s)\)/);
        if (match) {
            let castTime = parseFloat(match[3]);
            if (match[4] == 'ms') {
                castTime /= 1000;
            }
            return ActionId.fromLogString(match[1]).fill(params.source?.index).then(castId => new CastBeganLog(params, castId, parseFloat(match[2]), castTime));
        }
        else {
            return null;
        }
    }
}
export class CastCompletedLog extends SimLog {
    constructor(params, castId) {
        super(params);
        this.castId = castId;
    }
    toString() {
        return `${this.toStringPrefix()} Completed cast ${this.castId.name}.`;
    }
    static parse(params) {
        const match = params.raw.match(/Completed cast (.*)/);
        if (match) {
            return ActionId.fromLogString(match[1]).fill(params.source?.index).then(castId => new CastCompletedLog(params, castId));
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
        });
        this.castId = castBeganLog.castId;
        this.castTime = castBeganLog.castTime;
        this.castBeganLog = castBeganLog;
        this.castCompletedLog = castCompletedLog;
        this.damageDealtLogs = damageDealtLogs;
    }
    toString() {
        return `${this.toStringPrefix()} Casting ${this.castId.name} (Cast time = ${this.castTime.toFixed(2)}s).`;
    }
    static fromLogs(logs) {
        const castBeganLogs = logs.filter((log) => log.isCastBegan());
        const castCompletedLogs = logs.filter((log) => log.isCastCompleted());
        const damageDealtLogs = logs.filter((log) => log.isDamageDealt());
        const castBeganLogsByAbility = Object.values(bucket(castBeganLogs, log => log.castId.toString()));
        const damageDealtLogsByAbility = Object.values(bucket(damageDealtLogs, log => log.cause.toString()));
        const castLogs = [];
        castBeganLogsByAbility.forEach(abilityCastsBegan => {
            const actionId = abilityCastsBegan[0].castId;
            const abilityDamageDealt = damageDealtLogsByAbility.find(ddl => ddl[0].cause.equals(actionId));
            const getCastCompleted = (cbIndex) => {
                if (cbIndex >= abilityCastsBegan.length) {
                    return null;
                }
                const nextBeganIndex = abilityCastsBegan[cbIndex + 1]?.logIndex || null;
                return castCompletedLogs.find(ccl => ccl.logIndex > abilityCastsBegan[cbIndex].logIndex
                    && (nextBeganIndex == null || ccl.logIndex < nextBeganIndex)
                    && ccl.castId.equals(actionId)) || null;
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
    constructor(params, effectId, amount, stat) {
        super(params);
        this.effectId = effectId;
        this.amount = amount;
        this.stat = stat;
    }
    toString() {
        if (this.amount > 0) {
            return `${this.toStringPrefix()} Gained ${this.amount.toFixed(0)} ${this.stat} from ${this.effectId.name}.`;
        }
        else {
            return `${this.toStringPrefix()} Lost ${(-this.amount).toFixed(0)} ${this.stat} from fading ${this.effectId.name}.`;
        }
    }
    static parse(params) {
        const match = params.raw.match(/((Gained)|(Lost)) (\d+\.?\d*) (.*) from (fading )?(.*)/);
        if (match) {
            return ActionId.fromLogString(match[7]).fill(params.source?.index).then(effectId => {
                const sign = match[1] == 'Lost' ? -1 : 1;
                return new StatChangeLog(params, effectId, parseFloat(match[4]) * sign, match[5]);
            });
        }
        else {
            return null;
        }
    }
}
