import { RaidSimResult } from '/tbc/core/proto/api.js';
import { ResourceType } from '/tbc/core/proto/api.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
export declare class Entity {
    readonly name: string;
    readonly ownerName: string;
    readonly index: number;
    readonly isTarget: boolean;
    readonly isPet: boolean;
    constructor(name: string, ownerName: string, index: number, isTarget: boolean, isPet: boolean);
    equals(other: Entity): boolean;
    toString(): string;
    static parseRegex: RegExp;
    static parseAll(str: string): Array<Entity>;
}
interface SimLogParams {
    raw: string;
    logIndex: number;
    timestamp: number;
    source: Entity | null;
    target: Entity | null;
    actionId: ActionId | null;
    threat: number;
}
export declare class SimLog {
    readonly raw: string;
    readonly logIndex: number;
    readonly timestamp: number;
    readonly source: Entity | null;
    readonly target: Entity | null;
    readonly actionId: ActionId | null;
    readonly threat: number;
    activeAuras: Array<AuraUptimeLog>;
    constructor(params: SimLogParams);
    toString(): string;
    toStringPrefix(): string;
    static parseAll(result: RaidSimResult): Promise<Array<SimLog>>;
    isDamageDealt(): this is DamageDealtLog;
    isResourceChanged(): this is ResourceChangedLog;
    isAuraEvent(): this is AuraEventLog;
    isAuraStacksChange(): this is AuraStacksChangeLog;
    isMajorCooldownUsed(): this is MajorCooldownUsedLog;
    isCastBegan(): this is CastBeganLog;
    isCastCompleted(): this is CastCompletedLog;
    isStatChange(): this is StatChangeLog;
    static groupDuplicateTimestamps<LogType extends SimLog>(logs: Array<LogType>): Array<Array<LogType>>;
}
export declare class DamageDealtLog extends SimLog {
    readonly amount: number;
    readonly miss: boolean;
    readonly hit: boolean;
    readonly crit: boolean;
    readonly crush: boolean;
    readonly glance: boolean;
    readonly dodge: boolean;
    readonly parry: boolean;
    readonly block: boolean;
    readonly tick: boolean;
    readonly partialResist1_4: boolean;
    readonly partialResist2_4: boolean;
    readonly partialResist3_4: boolean;
    constructor(params: SimLogParams, amount: number, miss: boolean, crit: boolean, crush: boolean, glance: boolean, dodge: boolean, parry: boolean, block: boolean, tick: boolean, partialResist1_4: boolean, partialResist2_4: boolean, partialResist3_4: boolean);
    resultString(): string;
    toString(): string;
    static parse(params: SimLogParams): Promise<DamageDealtLog> | null;
}
export declare class DpsLog extends SimLog {
    readonly dps: number;
    readonly damageLogs: Array<DamageDealtLog>;
    constructor(params: SimLogParams, dps: number, damageLogs: Array<DamageDealtLog>);
    static DPS_WINDOW: number;
    static fromLogs(damageDealtLogs: Array<DamageDealtLog>): Array<DpsLog>;
}
export declare class ThreatLogGroup extends SimLog {
    readonly threatBefore: number;
    readonly threatAfter: number;
    readonly logs: Array<SimLog>;
    constructor(params: SimLogParams, threatBefore: number, threatAfter: number, logs: Array<SimLog>);
    static fromLogs(logs: Array<SimLog>): Array<ThreatLogGroup>;
}
export declare class AuraEventLog extends SimLog {
    readonly isGained: boolean;
    readonly isFaded: boolean;
    readonly isRefreshed: boolean;
    constructor(params: SimLogParams, isGained: boolean, isFaded: boolean, isRefreshed: boolean);
    toString(): string;
    static parse(params: SimLogParams): Promise<AuraEventLog> | null;
}
export declare class AuraStacksChangeLog extends SimLog {
    readonly oldStacks: number;
    readonly newStacks: number;
    constructor(params: SimLogParams, oldStacks: number, newStacks: number);
    toString(): string;
    static parse(params: SimLogParams): Promise<AuraStacksChangeLog> | null;
}
export declare class AuraUptimeLog extends SimLog {
    readonly gainedAt: number;
    readonly fadedAt: number;
    constructor(params: SimLogParams, fadedAt: number);
    static fromLogs(logs: Array<SimLog>, entity: Entity, encounterDuration: number): Array<AuraUptimeLog>;
    static populateActiveAuras(logs: Array<SimLog>, auraLogs: Array<AuraUptimeLog>): void;
}
export declare class ResourceChangedLog extends SimLog {
    readonly resourceType: ResourceType;
    readonly valueBefore: number;
    readonly valueAfter: number;
    readonly isSpend: boolean;
    constructor(params: SimLogParams, resourceType: ResourceType, valueBefore: number, valueAfter: number, isSpend: boolean);
    toString(): string;
    resultString(): string;
    static parse(params: SimLogParams): Promise<ResourceChangedLog> | null;
}
export declare class ResourceChangedLogGroup extends SimLog {
    readonly resourceType: ResourceType;
    readonly valueBefore: number;
    readonly valueAfter: number;
    readonly logs: Array<ResourceChangedLog>;
    constructor(params: SimLogParams, resourceType: ResourceType, valueBefore: number, valueAfter: number, logs: Array<ResourceChangedLog>);
    toString(): string;
    static fromLogs(logs: Array<SimLog>): Record<ResourceType, Array<ResourceChangedLogGroup>>;
}
export declare class MajorCooldownUsedLog extends SimLog {
    constructor(params: SimLogParams);
    toString(): string;
    static parse(params: SimLogParams): Promise<MajorCooldownUsedLog> | null;
}
export declare class CastBeganLog extends SimLog {
    readonly manaCost: number;
    readonly castTime: number;
    constructor(params: SimLogParams, manaCost: number, castTime: number);
    toString(): string;
    static parse(params: SimLogParams): Promise<CastBeganLog> | null;
}
export declare class CastCompletedLog extends SimLog {
    constructor(params: SimLogParams);
    toString(): string;
    static parse(params: SimLogParams): Promise<CastCompletedLog> | null;
}
export declare class CastLog extends SimLog {
    readonly castTime: number;
    readonly castBeganLog: CastBeganLog;
    readonly castCompletedLog: CastCompletedLog | null;
    readonly damageDealtLogs: Array<DamageDealtLog>;
    constructor(castBeganLog: CastBeganLog, castCompletedLog: CastCompletedLog | null, damageDealtLogs: Array<DamageDealtLog>);
    toString(): string;
    static fromLogs(logs: Array<SimLog>): Array<CastLog>;
}
export declare class StatChangeLog extends SimLog {
    readonly isGain: boolean;
    readonly stats: string;
    constructor(params: SimLogParams, isGain: boolean, stats: string);
    toString(): string;
    static parse(params: SimLogParams): Promise<StatChangeLog> | null;
}
export {};
