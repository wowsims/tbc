import { ActionMetrics as ActionMetricsProto } from '/tbc/core/proto/api.js';
import { AuraMetrics as AuraMetricsProto } from '/tbc/core/proto/api.js';
import { DpsMetrics as DpsMetricsProto } from '/tbc/core/proto/api.js';
import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { EncounterMetrics as EncounterMetricsProto } from '/tbc/core/proto/api.js';
import { Party as PartyProto } from '/tbc/core/proto/api.js';
import { PartyMetrics as PartyMetricsProto } from '/tbc/core/proto/api.js';
import { Player as PlayerProto } from '/tbc/core/proto/api.js';
import { PlayerMetrics as PlayerMetricsProto } from '/tbc/core/proto/api.js';
import { Raid as RaidProto } from '/tbc/core/proto/api.js';
import { RaidMetrics as RaidMetricsProto } from '/tbc/core/proto/api.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { TargetMetrics as TargetMetricsProto } from '/tbc/core/proto/api.js';
import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';
import { Spec } from '/tbc/core/proto/common.js';
import { ActionId } from '/tbc/core/resources.js';
export interface SimResultFilter {
    player?: number | null;
    target?: number | null;
}
export declare class SimResult {
    private readonly request;
    private readonly result;
    readonly raidMetrics: RaidMetrics;
    readonly encounterMetrics: EncounterMetrics;
    private constructor();
    getPlayers(filter?: SimResultFilter): Array<PlayerMetrics>;
    getFirstPlayer(): PlayerMetrics | null;
    getPlayerWithRaidIndex(raidIndex: number): PlayerMetrics | null;
    getTargets(filter?: SimResultFilter): Array<TargetMetrics>;
    getTargetWithIndex(index: number): TargetMetrics | null;
    getDamageMetrics(filter: SimResultFilter): DpsMetricsProto;
    getActionMetrics(filter: SimResultFilter): Array<ActionMetrics>;
    getSpellMetrics(filter: SimResultFilter): Array<ActionMetrics>;
    getBuffMetrics(filter: SimResultFilter): Array<AuraMetrics>;
    getDebuffMetrics(filter: SimResultFilter): Array<AuraMetrics>;
    getLogs(): Array<string>;
    toJson(): any;
    static fromJson(obj: any): Promise<SimResult>;
    static makeNew(request: RaidSimRequest, result: RaidSimResult): Promise<SimResult>;
}
export declare class RaidMetrics {
    private readonly raid;
    private readonly metrics;
    readonly dps: DpsMetricsProto;
    readonly parties: Array<PartyMetrics>;
    private constructor();
    static makeNew(iterations: number, duration: number, raid: RaidProto, metrics: RaidMetricsProto): Promise<RaidMetrics>;
}
export declare class PartyMetrics {
    private readonly party;
    private readonly metrics;
    readonly partyIndex: number;
    readonly dps: DpsMetricsProto;
    readonly players: Array<PlayerMetrics>;
    private constructor();
    static makeNew(iterations: number, duration: number, party: PartyProto, metrics: PartyMetricsProto, partyIndex: number): Promise<PartyMetrics>;
}
export declare class PlayerMetrics {
    private readonly player;
    private readonly metrics;
    readonly raidIndex: number;
    readonly name: string;
    readonly spec: Spec;
    readonly dps: DpsMetricsProto;
    readonly actions: Array<ActionMetrics>;
    readonly auras: Array<AuraMetrics>;
    private readonly iterations;
    private readonly duration;
    private constructor();
    get oomPercent(): number;
    static makeNew(iterations: number, duration: number, player: PlayerProto, metrics: PlayerMetricsProto, raidIndex: number): Promise<PlayerMetrics>;
}
export declare class EncounterMetrics {
    private readonly encounter;
    private readonly metrics;
    readonly targets: Array<TargetMetrics>;
    private constructor();
    static makeNew(iterations: number, duration: number, encounter: EncounterProto, metrics: EncounterMetricsProto): Promise<EncounterMetrics>;
}
export declare class TargetMetrics {
    private readonly target;
    private readonly metrics;
    readonly index: number;
    readonly auras: Array<AuraMetrics>;
    private constructor();
    static makeNew(iterations: number, duration: number, target: TargetProto, metrics: TargetMetricsProto, index: number): Promise<TargetMetrics>;
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
    static makeNew(iterations: number, duration: number, auraMetrics: AuraMetricsProto): Promise<AuraMetrics>;
    static join(auras: Array<AuraMetrics>): Array<AuraMetrics>;
}
export declare class ActionMetrics {
    readonly actionId: ActionId;
    readonly name: string;
    readonly iconUrl: string;
    private readonly iterations;
    private readonly duration;
    private readonly data;
    private constructor();
    get damage(): number;
    get dps(): number;
    get casts(): number;
    get castsPerMinute(): number;
    get avgCast(): number;
    get hits(): number;
    get avgHit(): number;
    get critPercent(): number;
    get misses(): number;
    get missPercent(): number;
    static makeNew(iterations: number, duration: number, actionMetrics: ActionMetricsProto): Promise<ActionMetrics>;
    static join(actions: Array<ActionMetrics>): Array<ActionMetrics>;
}
