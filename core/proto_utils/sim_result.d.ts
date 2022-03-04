import { ActionMetrics as ActionMetricsProto } from '/tbc/core/proto/api.js';
import { AuraMetrics as AuraMetricsProto } from '/tbc/core/proto/api.js';
import { DistributionMetrics as DistributionMetricsProto } from '/tbc/core/proto/api.js';
import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { EncounterMetrics as EncounterMetricsProto } from '/tbc/core/proto/api.js';
import { Party as PartyProto } from '/tbc/core/proto/api.js';
import { PartyMetrics as PartyMetricsProto } from '/tbc/core/proto/api.js';
import { Player as PlayerProto } from '/tbc/core/proto/api.js';
import { PlayerMetrics as PlayerMetricsProto } from '/tbc/core/proto/api.js';
import { Raid as RaidProto } from '/tbc/core/proto/api.js';
import { RaidMetrics as RaidMetricsProto } from '/tbc/core/proto/api.js';
import { ResourceMetrics as ResourceMetricsProto, ResourceType } from '/tbc/core/proto/api.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { TargetMetrics as TargetMetricsProto } from '/tbc/core/proto/api.js';
import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';
import { Spec } from '/tbc/core/proto/common.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { AuraUptimeLog, CastLog, DamageDealtLog, DpsLog, MajorCooldownUsedLog, ResourceChangedLogGroup, SimLog } from './logs_parser.js';
export interface SimResultFilter {
    player?: number | null;
    target?: number | null;
}
export declare class SimResult {
    readonly request: RaidSimRequest;
    readonly result: RaidSimResult;
    readonly raidMetrics: RaidMetrics;
    readonly encounterMetrics: EncounterMetrics;
    readonly logs: Array<SimLog>;
    private constructor();
    getPlayers(filter?: SimResultFilter): Array<PlayerMetrics>;
    getFirstPlayer(): PlayerMetrics | null;
    getPlayerWithRaidIndex(raidIndex: number): PlayerMetrics | null;
    getTargets(filter?: SimResultFilter): Array<TargetMetrics>;
    getTargetWithIndex(index: number): TargetMetrics | null;
    getDamageMetrics(filter: SimResultFilter): DistributionMetricsProto;
    getActionMetrics(filter: SimResultFilter): Array<ActionMetrics>;
    getSpellMetrics(filter: SimResultFilter): Array<ActionMetrics>;
    getMeleeMetrics(filter: SimResultFilter): Array<ActionMetrics>;
    getResourceMetrics(filter: SimResultFilter, resourceType: ResourceType): Array<ResourceMetrics>;
    getBuffMetrics(filter: SimResultFilter): Array<AuraMetrics>;
    getDebuffMetrics(filter: SimResultFilter): Array<AuraMetrics>;
    toJson(): any;
    static fromJson(obj: any): Promise<SimResult>;
    static makeNew(request: RaidSimRequest, result: RaidSimResult): Promise<SimResult>;
}
export declare class RaidMetrics {
    private readonly raid;
    private readonly metrics;
    readonly dps: DistributionMetricsProto;
    readonly parties: Array<PartyMetrics>;
    private constructor();
    static makeNew(iterations: number, duration: number, raid: RaidProto, metrics: RaidMetricsProto, logs: Array<SimLog>): Promise<RaidMetrics>;
}
export declare class PartyMetrics {
    private readonly party;
    private readonly metrics;
    readonly partyIndex: number;
    readonly dps: DistributionMetricsProto;
    readonly players: Array<PlayerMetrics>;
    private constructor();
    static makeNew(iterations: number, duration: number, party: PartyProto, metrics: PartyMetricsProto, partyIndex: number, logs: Array<SimLog>): Promise<PartyMetrics>;
}
export declare class PlayerMetrics {
    private readonly player;
    private readonly metrics;
    readonly raidIndex: number;
    readonly name: string;
    readonly spec: Spec;
    readonly isPet: boolean;
    readonly iconUrl: string;
    readonly classColor: string;
    readonly dps: DistributionMetricsProto;
    readonly actions: Array<ActionMetrics>;
    readonly auras: Array<AuraMetrics>;
    readonly resources: Array<ResourceMetrics>;
    readonly pets: Array<PlayerMetrics>;
    private readonly iterations;
    private readonly duration;
    readonly logs: Array<SimLog>;
    readonly damageDealtLogs: Array<DamageDealtLog>;
    readonly manaChangedLogs: Array<ResourceChangedLogGroup>;
    readonly dpsLogs: Array<DpsLog>;
    readonly auraUptimeLogs: Array<AuraUptimeLog>;
    readonly majorCooldownLogs: Array<MajorCooldownUsedLog>;
    readonly castLogs: Array<CastLog>;
    readonly majorCooldownAuraUptimeLogs: Array<AuraUptimeLog>;
    private constructor();
    get label(): string;
    get secondsOomAvg(): number;
    get totalDamage(): number;
    getPlayerAndPetActions(): Array<ActionMetrics>;
    getMeleeActions(): Array<ActionMetrics>;
    getSpellActions(): Array<ActionMetrics>;
    static makeNew(iterations: number, duration: number, player: PlayerProto, metrics: PlayerMetricsProto, raidIndex: number, isPet: boolean, logs: Array<SimLog>): Promise<PlayerMetrics>;
}
export declare class EncounterMetrics {
    private readonly encounter;
    private readonly metrics;
    readonly targets: Array<TargetMetrics>;
    private constructor();
    static makeNew(iterations: number, duration: number, encounter: EncounterProto, metrics: EncounterMetricsProto, logs: Array<SimLog>): Promise<EncounterMetrics>;
    get durationSeconds(): number;
}
export declare class TargetMetrics {
    private readonly target;
    private readonly metrics;
    readonly index: number;
    readonly auras: Array<AuraMetrics>;
    readonly logs: Array<SimLog>;
    private constructor();
    static makeNew(iterations: number, duration: number, target: TargetProto, metrics: TargetMetricsProto, index: number, logs: Array<SimLog>): Promise<TargetMetrics>;
}
export declare class AuraMetrics {
    readonly actionId: ActionId;
    readonly name: string;
    readonly iconUrl: string;
    private readonly iterations;
    private readonly duration;
    private readonly data;
    private constructor();
    get uptimePercent(): number;
    static makeNew(iterations: number, duration: number, auraMetrics: AuraMetricsProto, playerIndex?: number): Promise<AuraMetrics>;
    static merge(auras: Array<AuraMetrics>): AuraMetrics;
    static joinById(auras: Array<AuraMetrics>): Array<AuraMetrics>;
}
export declare class ResourceMetrics {
    readonly actionId: ActionId;
    readonly name: string;
    readonly iconUrl: string;
    readonly type: ResourceType;
    private readonly iterations;
    private readonly duration;
    private readonly data;
    private constructor();
    get events(): number;
    get gain(): number;
    get gainPerSecond(): number;
    get avgGain(): number;
    get avgActualGain(): number;
    static makeNew(iterations: number, duration: number, resourceMetrics: ResourceMetricsProto, playerIndex?: number): Promise<ResourceMetrics>;
    static merge(resources: Array<ResourceMetrics>): ResourceMetrics;
    static joinById(resources: Array<ResourceMetrics>): Array<ResourceMetrics>;
}
export declare class ActionMetrics {
    readonly actionId: ActionId;
    readonly name: string;
    readonly iconUrl: string;
    private readonly iterations;
    private readonly duration;
    private readonly data;
    private constructor();
    get isMeleeAction(): boolean;
    get damage(): number;
    get dps(): number;
    get casts(): number;
    get castsPerMinute(): number;
    get avgCast(): number;
    get hits(): number;
    private get landedHitsRaw();
    get landedHits(): number;
    get hitAttempts(): number;
    get avgHit(): number;
    get critPercent(): number;
    get misses(): number;
    get missPercent(): number;
    get dodges(): number;
    get dodgePercent(): number;
    get parries(): number;
    get parryPercent(): number;
    get blocks(): number;
    get blockPercent(): number;
    get glances(): number;
    get glancePercent(): number;
    static makeNew(iterations: number, duration: number, actionMetrics: ActionMetricsProto, playerIndex?: number): Promise<ActionMetrics>;
    static merge(actions: Array<ActionMetrics>, removeTag?: boolean, actionIdOverride?: ActionId): ActionMetrics;
    static joinById(actions: Array<ActionMetrics>): Array<ActionMetrics>;
    static groupById(actions: Array<ActionMetrics>): Array<Array<ActionMetrics>>;
}
