import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Stat } from "./common";
import { Gem } from "./common";
import { Enchant } from "./common";
import { Item } from "./common";
import { Spec } from "./common";
import { Encounter } from "./common";
import { Buffs } from "./common";
import { EquipmentSpec } from "./common";
import { Consumes } from "./common";
import { ElementalShaman } from "./shaman";
import { BalanceDruid } from "./druid";
import { Race } from "./common";
/**
 * @generated from protobuf message api.PlayerOptions
 */
export interface PlayerOptions {
    /**
     * @generated from protobuf field: api.Race race = 1;
     */
    race: Race;
    /**
     * @generated from protobuf oneof: spec
     */
    spec: {
        oneofKind: "balanceDruid";
        /**
         * @generated from protobuf field: api.BalanceDruid balance_druid = 2;
         */
        balanceDruid: BalanceDruid;
    } | {
        oneofKind: "hunter";
        /**
         * @generated from protobuf field: api.Hunter hunter = 3;
         */
        hunter: Hunter;
    } | {
        oneofKind: "mage";
        /**
         * @generated from protobuf field: api.Mage mage = 4;
         */
        mage: Mage;
    } | {
        oneofKind: "paladin";
        /**
         * @generated from protobuf field: api.Paladin paladin = 5;
         */
        paladin: Paladin;
    } | {
        oneofKind: "priest";
        /**
         * @generated from protobuf field: api.Priest priest = 6;
         */
        priest: Priest;
    } | {
        oneofKind: "rogue";
        /**
         * @generated from protobuf field: api.Rogue rogue = 7;
         */
        rogue: Rogue;
    } | {
        oneofKind: "elementalShaman";
        /**
         * @generated from protobuf field: api.ElementalShaman elemental_shaman = 8;
         */
        elementalShaman: ElementalShaman;
    } | {
        oneofKind: "warlock";
        /**
         * @generated from protobuf field: api.Warlock warlock = 9;
         */
        warlock: Warlock;
    } | {
        oneofKind: "warrior";
        /**
         * @generated from protobuf field: api.Warrior warrior = 10;
         */
        warrior: Warrior;
    } | {
        oneofKind: undefined;
    };
    /**
     * @generated from protobuf field: api.Consumes consumes = 11;
     */
    consumes?: Consumes;
}
/**
 * @generated from protobuf message api.Hunter
 */
export interface Hunter {
}
/**
 * @generated from protobuf message api.Mage
 */
export interface Mage {
}
/**
 * @generated from protobuf message api.Paladin
 */
export interface Paladin {
}
/**
 * @generated from protobuf message api.Priest
 */
export interface Priest {
}
/**
 * @generated from protobuf message api.Rogue
 */
export interface Rogue {
}
/**
 * @generated from protobuf message api.Warlock
 */
export interface Warlock {
}
/**
 * @generated from protobuf message api.Warrior
 */
export interface Warrior {
}
/**
 * @generated from protobuf message api.Player
 */
export interface Player {
    /**
     * @generated from protobuf field: api.PlayerOptions options = 1;
     */
    options?: PlayerOptions;
    /**
     * @generated from protobuf field: api.EquipmentSpec equipment = 2;
     */
    equipment?: EquipmentSpec;
    /**
     * @generated from protobuf field: repeated double custom_stats = 3;
     */
    customStats: number[];
}
/**
 * @generated from protobuf message api.Party
 */
export interface Party {
    /**
     * @generated from protobuf field: repeated api.Player players = 1;
     */
    players: Player[];
}
/**
 * @generated from protobuf message api.Raid
 */
export interface Raid {
    /**
     * @generated from protobuf field: repeated api.Party parties = 1;
     */
    parties: Party[];
}
/**
 * RPC IndividualSim
 *
 * @generated from protobuf message api.IndividualSimRequest
 */
export interface IndividualSimRequest {
    /**
     * @generated from protobuf field: api.Player player = 1;
     */
    player?: Player;
    /**
     * @generated from protobuf field: api.Buffs buffs = 2;
     */
    buffs?: Buffs;
    /**
     * @generated from protobuf field: api.Encounter encounter = 3;
     */
    encounter?: Encounter;
    /**
     * @generated from protobuf field: int32 iterations = 4;
     */
    iterations: number;
    /**
     * @generated from protobuf field: int64 random_seed = 5;
     */
    randomSeed: bigint;
    /**
     * @generated from protobuf field: double gcd_min = 6;
     */
    gcdMin: number;
    /**
     * @generated from protobuf field: bool debug = 7;
     */
    debug: boolean;
    /**
     * @generated from protobuf field: bool exit_on_oom = 8;
     */
    exitOnOom: boolean;
}
/**
 * @generated from protobuf message api.IndividualSimResult
 */
export interface IndividualSimResult {
    /**
     * @generated from protobuf field: int64 execution_duration_ms = 1;
     */
    executionDurationMs: bigint;
    /**
     * @generated from protobuf field: string logs = 2;
     */
    logs: string;
    /**
     * @generated from protobuf field: double dps_avg = 3;
     */
    dpsAvg: number;
    /**
     * @generated from protobuf field: double dps_stdev = 4;
     */
    dpsStdev: number;
    /**
     * @generated from protobuf field: double dps_max = 5;
     */
    dpsMax: number;
    /**
     * @generated from protobuf field: map<int32, int32> dps_hist = 6;
     */
    dpsHist: {
        [key: number]: number;
    };
    /**
     * @generated from protobuf field: int32 num_oom = 7;
     */
    numOom: number;
    /**
     * @generated from protobuf field: double oom_at_avg = 8;
     */
    oomAtAvg: number;
    /**
     * @generated from protobuf field: double dps_at_oom_avg = 9;
     */
    dpsAtOomAvg: number;
    /**
     * @generated from protobuf field: map<int32, api.CastMetric> casts = 10;
     */
    casts: {
        [key: number]: CastMetric;
    };
    /**
     * @generated from protobuf field: string error = 11;
     */
    error: string;
}
/**
 * CastMetric holds a collection of counts of casts and
 *
 *
 * @generated from protobuf message api.CastMetric
 */
export interface CastMetric {
    /**
     * @generated from protobuf field: repeated int32 counts = 1;
     */
    counts: number[];
    /**
     * @generated from protobuf field: repeated double dmgs = 2;
     */
    dmgs: number[];
    /**
     * @generated from protobuf field: repeated int32 tags = 3;
     */
    tags: number[];
}
/**
 * RPC RaidSim
 *
 * @generated from protobuf message api.RaidSimRequest
 */
export interface RaidSimRequest {
    /**
     * @generated from protobuf field: api.Raid raid = 1;
     */
    raid?: Raid;
    /**
     * @generated from protobuf field: api.Encounter encounter = 2;
     */
    encounter?: Encounter;
    /**
     * @generated from protobuf field: int64 random_seed = 3;
     */
    randomSeed: bigint;
    /**
     * @generated from protobuf field: double gcd_min = 4;
     */
    gcdMin: number;
    /**
     * @generated from protobuf field: bool debug = 5;
     */
    debug: boolean;
}
/**
 * RPC GearList
 *
 * @generated from protobuf message api.GearListRequest
 */
export interface GearListRequest {
    /**
     * @generated from protobuf field: api.Spec spec = 1;
     */
    spec: Spec;
}
/**
 * @generated from protobuf message api.GearListResult
 */
export interface GearListResult {
    /**
     * @generated from protobuf field: repeated api.Item items = 1;
     */
    items: Item[];
    /**
     * @generated from protobuf field: repeated api.Enchant enchants = 2;
     */
    enchants: Enchant[];
    /**
     * @generated from protobuf field: repeated api.Gem gems = 3;
     */
    gems: Gem[];
}
/**
 * RPC ComputeStats
 *
 * @generated from protobuf message api.ComputeStatsRequest
 */
export interface ComputeStatsRequest {
    /**
     * @generated from protobuf field: api.Player player = 1;
     */
    player?: Player;
    /**
     * @generated from protobuf field: api.Buffs buffs = 2;
     */
    buffs?: Buffs;
}
/**
 * @generated from protobuf message api.ComputeStatsResult
 */
export interface ComputeStatsResult {
    /**
     * @generated from protobuf field: repeated double gear_only = 1;
     */
    gearOnly: number[];
    /**
     * @generated from protobuf field: repeated double finalStats = 2;
     */
    finalStats: number[];
    /**
     * @generated from protobuf field: repeated string sets = 3;
     */
    sets: string[];
}
/**
 * RPC StatWeights
 *
 * @generated from protobuf message api.StatWeightsRequest
 */
export interface StatWeightsRequest {
    /**
     * @generated from protobuf field: api.IndividualSimRequest options = 1;
     */
    options?: IndividualSimRequest;
    /**
     * @generated from protobuf field: repeated api.Stat stats_to_weigh = 2;
     */
    statsToWeigh: Stat[];
    /**
     * @generated from protobuf field: api.Stat ep_reference_stat = 3;
     */
    epReferenceStat: Stat;
}
/**
 * @generated from protobuf message api.StatWeightsResult
 */
export interface StatWeightsResult {
    /**
     * @generated from protobuf field: repeated double weights = 1;
     */
    weights: number[];
    /**
     * @generated from protobuf field: repeated double weights_stdev = 2;
     */
    weightsStdev: number[];
    /**
     * @generated from protobuf field: repeated double ep_values = 3;
     */
    epValues: number[];
    /**
     * @generated from protobuf field: repeated double ep_values_stdev = 4;
     */
    epValuesStdev: number[];
}
declare class PlayerOptions$Type extends MessageType<PlayerOptions> {
    constructor();
    create(value?: PartialMessage<PlayerOptions>): PlayerOptions;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PlayerOptions): PlayerOptions;
    internalBinaryWrite(message: PlayerOptions, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.PlayerOptions
 */
export declare const PlayerOptions: PlayerOptions$Type;
declare class Hunter$Type extends MessageType<Hunter> {
    constructor();
    create(value?: PartialMessage<Hunter>): Hunter;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Hunter): Hunter;
    internalBinaryWrite(message: Hunter, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.Hunter
 */
export declare const Hunter: Hunter$Type;
declare class Mage$Type extends MessageType<Mage> {
    constructor();
    create(value?: PartialMessage<Mage>): Mage;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Mage): Mage;
    internalBinaryWrite(message: Mage, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.Mage
 */
export declare const Mage: Mage$Type;
declare class Paladin$Type extends MessageType<Paladin> {
    constructor();
    create(value?: PartialMessage<Paladin>): Paladin;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Paladin): Paladin;
    internalBinaryWrite(message: Paladin, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.Paladin
 */
export declare const Paladin: Paladin$Type;
declare class Priest$Type extends MessageType<Priest> {
    constructor();
    create(value?: PartialMessage<Priest>): Priest;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Priest): Priest;
    internalBinaryWrite(message: Priest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.Priest
 */
export declare const Priest: Priest$Type;
declare class Rogue$Type extends MessageType<Rogue> {
    constructor();
    create(value?: PartialMessage<Rogue>): Rogue;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Rogue): Rogue;
    internalBinaryWrite(message: Rogue, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.Rogue
 */
export declare const Rogue: Rogue$Type;
declare class Warlock$Type extends MessageType<Warlock> {
    constructor();
    create(value?: PartialMessage<Warlock>): Warlock;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Warlock): Warlock;
    internalBinaryWrite(message: Warlock, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.Warlock
 */
export declare const Warlock: Warlock$Type;
declare class Warrior$Type extends MessageType<Warrior> {
    constructor();
    create(value?: PartialMessage<Warrior>): Warrior;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Warrior): Warrior;
    internalBinaryWrite(message: Warrior, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.Warrior
 */
export declare const Warrior: Warrior$Type;
declare class Player$Type extends MessageType<Player> {
    constructor();
    create(value?: PartialMessage<Player>): Player;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Player): Player;
    internalBinaryWrite(message: Player, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.Player
 */
export declare const Player: Player$Type;
declare class Party$Type extends MessageType<Party> {
    constructor();
    create(value?: PartialMessage<Party>): Party;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Party): Party;
    internalBinaryWrite(message: Party, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.Party
 */
export declare const Party: Party$Type;
declare class Raid$Type extends MessageType<Raid> {
    constructor();
    create(value?: PartialMessage<Raid>): Raid;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Raid): Raid;
    internalBinaryWrite(message: Raid, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.Raid
 */
export declare const Raid: Raid$Type;
declare class IndividualSimRequest$Type extends MessageType<IndividualSimRequest> {
    constructor();
    create(value?: PartialMessage<IndividualSimRequest>): IndividualSimRequest;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: IndividualSimRequest): IndividualSimRequest;
    internalBinaryWrite(message: IndividualSimRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.IndividualSimRequest
 */
export declare const IndividualSimRequest: IndividualSimRequest$Type;
declare class IndividualSimResult$Type extends MessageType<IndividualSimResult> {
    constructor();
    create(value?: PartialMessage<IndividualSimResult>): IndividualSimResult;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: IndividualSimResult): IndividualSimResult;
    private binaryReadMap6;
    private binaryReadMap10;
    internalBinaryWrite(message: IndividualSimResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.IndividualSimResult
 */
export declare const IndividualSimResult: IndividualSimResult$Type;
declare class CastMetric$Type extends MessageType<CastMetric> {
    constructor();
    create(value?: PartialMessage<CastMetric>): CastMetric;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CastMetric): CastMetric;
    internalBinaryWrite(message: CastMetric, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.CastMetric
 */
export declare const CastMetric: CastMetric$Type;
declare class RaidSimRequest$Type extends MessageType<RaidSimRequest> {
    constructor();
    create(value?: PartialMessage<RaidSimRequest>): RaidSimRequest;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: RaidSimRequest): RaidSimRequest;
    internalBinaryWrite(message: RaidSimRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.RaidSimRequest
 */
export declare const RaidSimRequest: RaidSimRequest$Type;
declare class GearListRequest$Type extends MessageType<GearListRequest> {
    constructor();
    create(value?: PartialMessage<GearListRequest>): GearListRequest;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GearListRequest): GearListRequest;
    internalBinaryWrite(message: GearListRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.GearListRequest
 */
export declare const GearListRequest: GearListRequest$Type;
declare class GearListResult$Type extends MessageType<GearListResult> {
    constructor();
    create(value?: PartialMessage<GearListResult>): GearListResult;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GearListResult): GearListResult;
    internalBinaryWrite(message: GearListResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.GearListResult
 */
export declare const GearListResult: GearListResult$Type;
declare class ComputeStatsRequest$Type extends MessageType<ComputeStatsRequest> {
    constructor();
    create(value?: PartialMessage<ComputeStatsRequest>): ComputeStatsRequest;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ComputeStatsRequest): ComputeStatsRequest;
    internalBinaryWrite(message: ComputeStatsRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.ComputeStatsRequest
 */
export declare const ComputeStatsRequest: ComputeStatsRequest$Type;
declare class ComputeStatsResult$Type extends MessageType<ComputeStatsResult> {
    constructor();
    create(value?: PartialMessage<ComputeStatsResult>): ComputeStatsResult;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ComputeStatsResult): ComputeStatsResult;
    internalBinaryWrite(message: ComputeStatsResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.ComputeStatsResult
 */
export declare const ComputeStatsResult: ComputeStatsResult$Type;
declare class StatWeightsRequest$Type extends MessageType<StatWeightsRequest> {
    constructor();
    create(value?: PartialMessage<StatWeightsRequest>): StatWeightsRequest;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StatWeightsRequest): StatWeightsRequest;
    internalBinaryWrite(message: StatWeightsRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.StatWeightsRequest
 */
export declare const StatWeightsRequest: StatWeightsRequest$Type;
declare class StatWeightsResult$Type extends MessageType<StatWeightsResult> {
    constructor();
    create(value?: PartialMessage<StatWeightsResult>): StatWeightsResult;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StatWeightsResult): StatWeightsResult;
    internalBinaryWrite(message: StatWeightsResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.StatWeightsResult
 */
export declare const StatWeightsResult: StatWeightsResult$Type;
export {};
