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
 * @generated from protobuf message proto.PlayerOptions
 */
export interface PlayerOptions {
    /**
     * @generated from protobuf field: proto.Race race = 1;
     */
    race: Race;
    /**
     * @generated from protobuf oneof: spec
     */
    spec: {
        oneofKind: "balanceDruid";
        /**
         * @generated from protobuf field: proto.BalanceDruid balance_druid = 2;
         */
        balanceDruid: BalanceDruid;
    } | {
        oneofKind: "hunter";
        /**
         * @generated from protobuf field: proto.Hunter hunter = 3;
         */
        hunter: Hunter;
    } | {
        oneofKind: "mage";
        /**
         * @generated from protobuf field: proto.Mage mage = 4;
         */
        mage: Mage;
    } | {
        oneofKind: "paladin";
        /**
         * @generated from protobuf field: proto.Paladin paladin = 5;
         */
        paladin: Paladin;
    } | {
        oneofKind: "priest";
        /**
         * @generated from protobuf field: proto.Priest priest = 6;
         */
        priest: Priest;
    } | {
        oneofKind: "rogue";
        /**
         * @generated from protobuf field: proto.Rogue rogue = 7;
         */
        rogue: Rogue;
    } | {
        oneofKind: "elementalShaman";
        /**
         * @generated from protobuf field: proto.ElementalShaman elemental_shaman = 8;
         */
        elementalShaman: ElementalShaman;
    } | {
        oneofKind: "warlock";
        /**
         * @generated from protobuf field: proto.Warlock warlock = 9;
         */
        warlock: Warlock;
    } | {
        oneofKind: "warrior";
        /**
         * @generated from protobuf field: proto.Warrior warrior = 10;
         */
        warrior: Warrior;
    } | {
        oneofKind: undefined;
    };
    /**
     * @generated from protobuf field: proto.Consumes consumes = 11;
     */
    consumes?: Consumes;
}
/**
 * @generated from protobuf message proto.Hunter
 */
export interface Hunter {
}
/**
 * @generated from protobuf message proto.Mage
 */
export interface Mage {
}
/**
 * @generated from protobuf message proto.Paladin
 */
export interface Paladin {
}
/**
 * @generated from protobuf message proto.Priest
 */
export interface Priest {
}
/**
 * @generated from protobuf message proto.Rogue
 */
export interface Rogue {
}
/**
 * @generated from protobuf message proto.Warlock
 */
export interface Warlock {
}
/**
 * @generated from protobuf message proto.Warrior
 */
export interface Warrior {
}
/**
 * @generated from protobuf message proto.Player
 */
export interface Player {
    /**
     * @generated from protobuf field: proto.PlayerOptions options = 1;
     */
    options?: PlayerOptions;
    /**
     * @generated from protobuf field: proto.EquipmentSpec equipment = 2;
     */
    equipment?: EquipmentSpec;
    /**
     * @generated from protobuf field: repeated double custom_stats = 3;
     */
    customStats: number[];
}
/**
 * @generated from protobuf message proto.Party
 */
export interface Party {
    /**
     * @generated from protobuf field: repeated proto.Player players = 1;
     */
    players: Player[];
}
/**
 * @generated from protobuf message proto.Raid
 */
export interface Raid {
    /**
     * @generated from protobuf field: repeated proto.Party parties = 1;
     */
    parties: Party[];
}
/**
 * RPC IndividualSim
 *
 * @generated from protobuf message proto.IndividualSimRequest
 */
export interface IndividualSimRequest {
    /**
     * @generated from protobuf field: proto.Player player = 1;
     */
    player?: Player;
    /**
     * @generated from protobuf field: proto.Buffs buffs = 2;
     */
    buffs?: Buffs;
    /**
     * @generated from protobuf field: proto.Encounter encounter = 3;
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
 * @generated from protobuf message proto.IndividualSimResult
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
     * @generated from protobuf field: repeated proto.ActionMetric action_metrics = 10;
     */
    actionMetrics: ActionMetric[];
    /**
     * @generated from protobuf field: string error = 11;
     */
    error: string;
}
/**
 * ActionMetric holds a collection of counts of casts and
 *
 *
 * @generated from protobuf message proto.ActionMetric
 */
export interface ActionMetric {
    /**
     * @generated from protobuf oneof: action_id
     */
    actionId: {
        oneofKind: "spellId";
        /**
         * @generated from protobuf field: int32 spell_id = 1;
         */
        spellId: number;
    } | {
        oneofKind: "itemId";
        /**
         * @generated from protobuf field: int32 item_id = 2;
         */
        itemId: number;
    } | {
        oneofKind: "otherId";
        /**
         * @generated from protobuf field: int32 other_id = 3;
         */
        otherId: number;
    } | {
        oneofKind: undefined;
    };
    /**
     * These arrays are indexed by tag (index 0 is untagged casts)
     *  Currently the only use-case for this is shaman Lightning Overload.
     *
     * @generated from protobuf field: repeated int32 casts = 4;
     */
    casts: number[];
    /**
     * @generated from protobuf field: repeated int32 crits = 5;
     */
    crits: number[];
    /**
     * @generated from protobuf field: repeated int32 misses = 6;
     */
    misses: number[];
    /**
     * @generated from protobuf field: repeated double dmgs = 7;
     */
    dmgs: number[];
}
/**
 * RPC RaidSim
 *
 * @generated from protobuf message proto.RaidSimRequest
 */
export interface RaidSimRequest {
    /**
     * @generated from protobuf field: proto.Raid raid = 1;
     */
    raid?: Raid;
    /**
     * @generated from protobuf field: proto.Encounter encounter = 2;
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
 * @generated from protobuf message proto.GearListRequest
 */
export interface GearListRequest {
    /**
     * @generated from protobuf field: proto.Spec spec = 1;
     */
    spec: Spec;
}
/**
 * @generated from protobuf message proto.GearListResult
 */
export interface GearListResult {
    /**
     * @generated from protobuf field: repeated proto.Item items = 1;
     */
    items: Item[];
    /**
     * @generated from protobuf field: repeated proto.Enchant enchants = 2;
     */
    enchants: Enchant[];
    /**
     * @generated from protobuf field: repeated proto.Gem gems = 3;
     */
    gems: Gem[];
}
/**
 * RPC ComputeStats
 *
 * @generated from protobuf message proto.ComputeStatsRequest
 */
export interface ComputeStatsRequest {
    /**
     * @generated from protobuf field: proto.Player player = 1;
     */
    player?: Player;
    /**
     * @generated from protobuf field: proto.Buffs buffs = 2;
     */
    buffs?: Buffs;
}
/**
 * @generated from protobuf message proto.ComputeStatsResult
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
 * @generated from protobuf message proto.StatWeightsRequest
 */
export interface StatWeightsRequest {
    /**
     * @generated from protobuf field: proto.IndividualSimRequest options = 1;
     */
    options?: IndividualSimRequest;
    /**
     * @generated from protobuf field: repeated proto.Stat stats_to_weigh = 2;
     */
    statsToWeigh: Stat[];
    /**
     * @generated from protobuf field: proto.Stat ep_reference_stat = 3;
     */
    epReferenceStat: Stat;
}
/**
 * @generated from protobuf message proto.StatWeightsResult
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
 * @generated MessageType for protobuf message proto.PlayerOptions
 */
export declare const PlayerOptions: PlayerOptions$Type;
declare class Hunter$Type extends MessageType<Hunter> {
    constructor();
    create(value?: PartialMessage<Hunter>): Hunter;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Hunter): Hunter;
    internalBinaryWrite(message: Hunter, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Hunter
 */
export declare const Hunter: Hunter$Type;
declare class Mage$Type extends MessageType<Mage> {
    constructor();
    create(value?: PartialMessage<Mage>): Mage;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Mage): Mage;
    internalBinaryWrite(message: Mage, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Mage
 */
export declare const Mage: Mage$Type;
declare class Paladin$Type extends MessageType<Paladin> {
    constructor();
    create(value?: PartialMessage<Paladin>): Paladin;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Paladin): Paladin;
    internalBinaryWrite(message: Paladin, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Paladin
 */
export declare const Paladin: Paladin$Type;
declare class Priest$Type extends MessageType<Priest> {
    constructor();
    create(value?: PartialMessage<Priest>): Priest;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Priest): Priest;
    internalBinaryWrite(message: Priest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Priest
 */
export declare const Priest: Priest$Type;
declare class Rogue$Type extends MessageType<Rogue> {
    constructor();
    create(value?: PartialMessage<Rogue>): Rogue;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Rogue): Rogue;
    internalBinaryWrite(message: Rogue, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Rogue
 */
export declare const Rogue: Rogue$Type;
declare class Warlock$Type extends MessageType<Warlock> {
    constructor();
    create(value?: PartialMessage<Warlock>): Warlock;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Warlock): Warlock;
    internalBinaryWrite(message: Warlock, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Warlock
 */
export declare const Warlock: Warlock$Type;
declare class Warrior$Type extends MessageType<Warrior> {
    constructor();
    create(value?: PartialMessage<Warrior>): Warrior;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Warrior): Warrior;
    internalBinaryWrite(message: Warrior, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Warrior
 */
export declare const Warrior: Warrior$Type;
declare class Player$Type extends MessageType<Player> {
    constructor();
    create(value?: PartialMessage<Player>): Player;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Player): Player;
    internalBinaryWrite(message: Player, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Player
 */
export declare const Player: Player$Type;
declare class Party$Type extends MessageType<Party> {
    constructor();
    create(value?: PartialMessage<Party>): Party;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Party): Party;
    internalBinaryWrite(message: Party, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Party
 */
export declare const Party: Party$Type;
declare class Raid$Type extends MessageType<Raid> {
    constructor();
    create(value?: PartialMessage<Raid>): Raid;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Raid): Raid;
    internalBinaryWrite(message: Raid, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Raid
 */
export declare const Raid: Raid$Type;
declare class IndividualSimRequest$Type extends MessageType<IndividualSimRequest> {
    constructor();
    create(value?: PartialMessage<IndividualSimRequest>): IndividualSimRequest;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: IndividualSimRequest): IndividualSimRequest;
    internalBinaryWrite(message: IndividualSimRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.IndividualSimRequest
 */
export declare const IndividualSimRequest: IndividualSimRequest$Type;
declare class IndividualSimResult$Type extends MessageType<IndividualSimResult> {
    constructor();
    create(value?: PartialMessage<IndividualSimResult>): IndividualSimResult;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: IndividualSimResult): IndividualSimResult;
    private binaryReadMap6;
    internalBinaryWrite(message: IndividualSimResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.IndividualSimResult
 */
export declare const IndividualSimResult: IndividualSimResult$Type;
declare class ActionMetric$Type extends MessageType<ActionMetric> {
    constructor();
    create(value?: PartialMessage<ActionMetric>): ActionMetric;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ActionMetric): ActionMetric;
    internalBinaryWrite(message: ActionMetric, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ActionMetric
 */
export declare const ActionMetric: ActionMetric$Type;
declare class RaidSimRequest$Type extends MessageType<RaidSimRequest> {
    constructor();
    create(value?: PartialMessage<RaidSimRequest>): RaidSimRequest;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: RaidSimRequest): RaidSimRequest;
    internalBinaryWrite(message: RaidSimRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.RaidSimRequest
 */
export declare const RaidSimRequest: RaidSimRequest$Type;
declare class GearListRequest$Type extends MessageType<GearListRequest> {
    constructor();
    create(value?: PartialMessage<GearListRequest>): GearListRequest;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GearListRequest): GearListRequest;
    internalBinaryWrite(message: GearListRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.GearListRequest
 */
export declare const GearListRequest: GearListRequest$Type;
declare class GearListResult$Type extends MessageType<GearListResult> {
    constructor();
    create(value?: PartialMessage<GearListResult>): GearListResult;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GearListResult): GearListResult;
    internalBinaryWrite(message: GearListResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.GearListResult
 */
export declare const GearListResult: GearListResult$Type;
declare class ComputeStatsRequest$Type extends MessageType<ComputeStatsRequest> {
    constructor();
    create(value?: PartialMessage<ComputeStatsRequest>): ComputeStatsRequest;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ComputeStatsRequest): ComputeStatsRequest;
    internalBinaryWrite(message: ComputeStatsRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ComputeStatsRequest
 */
export declare const ComputeStatsRequest: ComputeStatsRequest$Type;
declare class ComputeStatsResult$Type extends MessageType<ComputeStatsResult> {
    constructor();
    create(value?: PartialMessage<ComputeStatsResult>): ComputeStatsResult;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ComputeStatsResult): ComputeStatsResult;
    internalBinaryWrite(message: ComputeStatsResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ComputeStatsResult
 */
export declare const ComputeStatsResult: ComputeStatsResult$Type;
declare class StatWeightsRequest$Type extends MessageType<StatWeightsRequest> {
    constructor();
    create(value?: PartialMessage<StatWeightsRequest>): StatWeightsRequest;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StatWeightsRequest): StatWeightsRequest;
    internalBinaryWrite(message: StatWeightsRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.StatWeightsRequest
 */
export declare const StatWeightsRequest: StatWeightsRequest$Type;
declare class StatWeightsResult$Type extends MessageType<StatWeightsResult> {
    constructor();
    create(value?: PartialMessage<StatWeightsResult>): StatWeightsResult;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StatWeightsResult): StatWeightsResult;
    internalBinaryWrite(message: StatWeightsResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.StatWeightsResult
 */
export declare const StatWeightsResult: StatWeightsResult$Type;
export {};
