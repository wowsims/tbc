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
import { Encounter } from "./common";
import { ActionID } from "./common";
import { RaidBuffs } from "./common";
import { PartyBuffs } from "./common";
import { Cooldowns } from "./common";
import { Warrior } from "./warrior";
import { Warlock } from "./warlock";
import { EnhancementShaman } from "./shaman";
import { ElementalShaman } from "./shaman";
import { Rogue } from "./rogue";
import { ShadowPriest } from "./priest";
import { RetributionPaladin } from "./paladin";
import { Mage } from "./mage";
import { Hunter } from "./hunter";
import { BalanceDruid } from "./druid";
import { IndividualBuffs } from "./common";
import { Consumes } from "./common";
import { EquipmentSpec } from "./common";
import { Class } from "./common";
import { Race } from "./common";
/**
 * @generated from protobuf message proto.Player
 */
export interface Player {
    /**
     * Label used for logging.
     *
     * @generated from protobuf field: string name = 16;
     */
    name: string;
    /**
     * @generated from protobuf field: proto.Race race = 1;
     */
    race: Race;
    /**
     * @generated from protobuf field: proto.Class class = 2;
     */
    class: Class;
    /**
     * @generated from protobuf field: proto.EquipmentSpec equipment = 3;
     */
    equipment?: EquipmentSpec;
    /**
     * @generated from protobuf field: proto.Consumes consumes = 4;
     */
    consumes?: Consumes;
    /**
     * @generated from protobuf field: repeated double bonus_stats = 5;
     */
    bonusStats: number[];
    /**
     * @generated from protobuf field: proto.IndividualBuffs buffs = 15;
     */
    buffs?: IndividualBuffs;
    /**
     * @generated from protobuf oneof: spec
     */
    spec: {
        oneofKind: "balanceDruid";
        /**
         * @generated from protobuf field: proto.BalanceDruid balance_druid = 6;
         */
        balanceDruid: BalanceDruid;
    } | {
        oneofKind: "hunter";
        /**
         * @generated from protobuf field: proto.Hunter hunter = 7;
         */
        hunter: Hunter;
    } | {
        oneofKind: "mage";
        /**
         * @generated from protobuf field: proto.Mage mage = 8;
         */
        mage: Mage;
    } | {
        oneofKind: "retributionPaladin";
        /**
         * @generated from protobuf field: proto.RetributionPaladin retribution_paladin = 9;
         */
        retributionPaladin: RetributionPaladin;
    } | {
        oneofKind: "shadowPriest";
        /**
         * @generated from protobuf field: proto.ShadowPriest shadow_priest = 10;
         */
        shadowPriest: ShadowPriest;
    } | {
        oneofKind: "rogue";
        /**
         * @generated from protobuf field: proto.Rogue rogue = 11;
         */
        rogue: Rogue;
    } | {
        oneofKind: "elementalShaman";
        /**
         * @generated from protobuf field: proto.ElementalShaman elemental_shaman = 12;
         */
        elementalShaman: ElementalShaman;
    } | {
        oneofKind: "enhancementShaman";
        /**
         * @generated from protobuf field: proto.EnhancementShaman enhancement_shaman = 18;
         */
        enhancementShaman: EnhancementShaman;
    } | {
        oneofKind: "warlock";
        /**
         * @generated from protobuf field: proto.Warlock warlock = 13;
         */
        warlock: Warlock;
    } | {
        oneofKind: "warrior";
        /**
         * @generated from protobuf field: proto.Warrior warrior = 14;
         */
        warrior: Warrior;
    } | {
        oneofKind: undefined;
    };
    /**
     * Only used by the UI. Sim uses talents within the spec protos.
     *
     * @generated from protobuf field: string talentsString = 17;
     */
    talentsString: string;
    /**
     * @generated from protobuf field: proto.Cooldowns cooldowns = 19;
     */
    cooldowns?: Cooldowns;
}
/**
 * @generated from protobuf message proto.Party
 */
export interface Party {
    /**
     * @generated from protobuf field: repeated proto.Player players = 1;
     */
    players: Player[];
    /**
     * @generated from protobuf field: proto.PartyBuffs buffs = 2;
     */
    buffs?: PartyBuffs;
}
/**
 * @generated from protobuf message proto.Raid
 */
export interface Raid {
    /**
     * @generated from protobuf field: repeated proto.Party parties = 1;
     */
    parties: Party[];
    /**
     * @generated from protobuf field: proto.RaidBuffs buffs = 2;
     */
    buffs?: RaidBuffs;
}
/**
 * @generated from protobuf message proto.SimOptions
 */
export interface SimOptions {
    /**
     * @generated from protobuf field: int32 iterations = 1;
     */
    iterations: number;
    /**
     * @generated from protobuf field: int64 random_seed = 2;
     */
    randomSeed: bigint;
    /**
     * @generated from protobuf field: bool debug = 3;
     */
    debug: boolean;
    /**
     * @generated from protobuf field: bool debug_first_iteration = 6;
     */
    debugFirstIteration: boolean;
    /**
     * @generated from protobuf field: bool is_test = 5;
     */
    isTest: boolean;
}
/**
 * The aggregated results from all uses of a particular action.
 *
 * @generated from protobuf message proto.ActionMetrics
 */
export interface ActionMetrics {
    /**
     * @generated from protobuf field: proto.ActionID id = 1;
     */
    id?: ActionID;
    /**
     * # of times this action was used by the agent.
     *
     * @generated from protobuf field: int32 casts = 2;
     */
    casts: number;
    /**
     * # of times this action hit a target. For cleave spells this can be larger than casts.
     *
     * @generated from protobuf field: int32 hits = 3;
     */
    hits: number;
    /**
     * # of times this action was a critical strike.
     *
     * @generated from protobuf field: int32 crits = 4;
     */
    crits: number;
    /**
     * # of times this action was a Miss or Resist.
     *
     * @generated from protobuf field: int32 misses = 5;
     */
    misses: number;
    /**
     * Total damage done to all targets by this action.
     *
     * @generated from protobuf field: double damage = 6;
     */
    damage: number;
}
/**
 * @generated from protobuf message proto.AuraMetrics
 */
export interface AuraMetrics {
    /**
     * @generated from protobuf field: proto.ActionID id = 1;
     */
    id?: ActionID;
    /**
     * @generated from protobuf field: double uptime_seconds_avg = 2;
     */
    uptimeSecondsAvg: number;
    /**
     * @generated from protobuf field: double uptime_seconds_stdev = 3;
     */
    uptimeSecondsStdev: number;
}
/**
 * @generated from protobuf message proto.DpsMetrics
 */
export interface DpsMetrics {
    /**
     * @generated from protobuf field: double avg = 1;
     */
    avg: number;
    /**
     * @generated from protobuf field: double stdev = 2;
     */
    stdev: number;
    /**
     * @generated from protobuf field: double max = 3;
     */
    max: number;
    /**
     * @generated from protobuf field: map<int32, int32> hist = 4;
     */
    hist: {
        [key: number]: number;
    };
}
/**
 * All the results for a single Player.
 *
 * @generated from protobuf message proto.PlayerMetrics
 */
export interface PlayerMetrics {
    /**
     * @generated from protobuf field: proto.DpsMetrics dps = 1;
     */
    dps?: DpsMetrics;
    /**
     * average seconds spent oom per iteration
     *
     * @generated from protobuf field: double seconds_oom_avg = 3;
     */
    secondsOomAvg: number;
    /**
     * @generated from protobuf field: repeated proto.ActionMetrics actions = 5;
     */
    actions: ActionMetrics[];
    /**
     * @generated from protobuf field: repeated proto.AuraMetrics auras = 6;
     */
    auras: AuraMetrics[];
    /**
     * @generated from protobuf field: repeated proto.PlayerMetrics pets = 7;
     */
    pets: PlayerMetrics[];
}
/**
 * Results for a whole raid.
 *
 * @generated from protobuf message proto.PartyMetrics
 */
export interface PartyMetrics {
    /**
     * @generated from protobuf field: proto.DpsMetrics dps = 1;
     */
    dps?: DpsMetrics;
    /**
     * @generated from protobuf field: repeated proto.PlayerMetrics players = 2;
     */
    players: PlayerMetrics[];
}
/**
 * Results for a whole raid.
 *
 * @generated from protobuf message proto.RaidMetrics
 */
export interface RaidMetrics {
    /**
     * @generated from protobuf field: proto.DpsMetrics dps = 1;
     */
    dps?: DpsMetrics;
    /**
     * @generated from protobuf field: repeated proto.PartyMetrics parties = 2;
     */
    parties: PartyMetrics[];
}
/**
 * @generated from protobuf message proto.TargetMetrics
 */
export interface TargetMetrics {
    /**
     * @generated from protobuf field: repeated proto.AuraMetrics auras = 1;
     */
    auras: AuraMetrics[];
}
/**
 * @generated from protobuf message proto.EncounterMetrics
 */
export interface EncounterMetrics {
    /**
     * @generated from protobuf field: repeated proto.TargetMetrics targets = 1;
     */
    targets: TargetMetrics[];
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
     * @generated from protobuf field: proto.SimOptions sim_options = 3;
     */
    simOptions?: SimOptions;
}
/**
 * Result from running the raid sim.
 *
 * @generated from protobuf message proto.RaidSimResult
 */
export interface RaidSimResult {
    /**
     * @generated from protobuf field: proto.RaidMetrics raid_metrics = 1;
     */
    raidMetrics?: RaidMetrics;
    /**
     * @generated from protobuf field: proto.EncounterMetrics encounter_metrics = 2;
     */
    encounterMetrics?: EncounterMetrics;
    /**
     * @generated from protobuf field: string logs = 3;
     */
    logs: string;
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
     * @generated from protobuf field: proto.RaidBuffs raid_buffs = 2;
     */
    raidBuffs?: RaidBuffs;
    /**
     * @generated from protobuf field: proto.PartyBuffs party_buffs = 3;
     */
    partyBuffs?: PartyBuffs;
    /**
     * @generated from protobuf field: proto.Encounter encounter = 5;
     */
    encounter?: Encounter;
    /**
     * @generated from protobuf field: proto.SimOptions sim_options = 6;
     */
    simOptions?: SimOptions;
}
/**
 * @generated from protobuf message proto.IndividualSimResult
 */
export interface IndividualSimResult {
    /**
     * @generated from protobuf field: proto.PlayerMetrics player_metrics = 1;
     */
    playerMetrics?: PlayerMetrics;
    /**
     * @generated from protobuf field: proto.EncounterMetrics encounter_metrics = 2;
     */
    encounterMetrics?: EncounterMetrics;
    /**
     * @generated from protobuf field: string logs = 3;
     */
    logs: string;
}
/**
 * RPC GearList
 *
 * @generated from protobuf message proto.GearListRequest
 */
export interface GearListRequest {
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
     * @generated from protobuf field: proto.Raid raid = 1;
     */
    raid?: Raid;
}
/**
 * @generated from protobuf message proto.PlayerStats
 */
export interface PlayerStats {
    /**
     * @generated from protobuf field: repeated double gear_only = 1;
     */
    gearOnly: number[];
    /**
     * @generated from protobuf field: repeated double final_stats = 2;
     */
    finalStats: number[];
    /**
     * @generated from protobuf field: repeated string sets = 3;
     */
    sets: string[];
    /**
     * @generated from protobuf field: proto.IndividualBuffs buffs = 4;
     */
    buffs?: IndividualBuffs;
    /**
     * @generated from protobuf field: repeated proto.ActionID cooldowns = 5;
     */
    cooldowns: ActionID[];
}
/**
 * @generated from protobuf message proto.PartyStats
 */
export interface PartyStats {
    /**
     * @generated from protobuf field: repeated proto.PlayerStats players = 1;
     */
    players: PlayerStats[];
}
/**
 * @generated from protobuf message proto.RaidStats
 */
export interface RaidStats {
    /**
     * @generated from protobuf field: repeated proto.PartyStats parties = 1;
     */
    parties: PartyStats[];
}
/**
 * @generated from protobuf message proto.ComputeStatsResult
 */
export interface ComputeStatsResult {
    /**
     * @generated from protobuf field: proto.RaidStats raid_stats = 1;
     */
    raidStats?: RaidStats;
}
/**
 * RPC StatWeights
 *
 * @generated from protobuf message proto.StatWeightsRequest
 */
export interface StatWeightsRequest {
    /**
     * @generated from protobuf field: proto.Player player = 1;
     */
    player?: Player;
    /**
     * @generated from protobuf field: proto.RaidBuffs raid_buffs = 2;
     */
    raidBuffs?: RaidBuffs;
    /**
     * @generated from protobuf field: proto.PartyBuffs party_buffs = 3;
     */
    partyBuffs?: PartyBuffs;
    /**
     * @generated from protobuf field: proto.Encounter encounter = 4;
     */
    encounter?: Encounter;
    /**
     * @generated from protobuf field: proto.SimOptions sim_options = 5;
     */
    simOptions?: SimOptions;
    /**
     * @generated from protobuf field: repeated proto.Stat stats_to_weigh = 6;
     */
    statsToWeigh: Stat[];
    /**
     * @generated from protobuf field: proto.Stat ep_reference_stat = 7;
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
declare class SimOptions$Type extends MessageType<SimOptions> {
    constructor();
    create(value?: PartialMessage<SimOptions>): SimOptions;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SimOptions): SimOptions;
    internalBinaryWrite(message: SimOptions, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.SimOptions
 */
export declare const SimOptions: SimOptions$Type;
declare class ActionMetrics$Type extends MessageType<ActionMetrics> {
    constructor();
    create(value?: PartialMessage<ActionMetrics>): ActionMetrics;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ActionMetrics): ActionMetrics;
    internalBinaryWrite(message: ActionMetrics, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ActionMetrics
 */
export declare const ActionMetrics: ActionMetrics$Type;
declare class AuraMetrics$Type extends MessageType<AuraMetrics> {
    constructor();
    create(value?: PartialMessage<AuraMetrics>): AuraMetrics;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: AuraMetrics): AuraMetrics;
    internalBinaryWrite(message: AuraMetrics, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.AuraMetrics
 */
export declare const AuraMetrics: AuraMetrics$Type;
declare class DpsMetrics$Type extends MessageType<DpsMetrics> {
    constructor();
    create(value?: PartialMessage<DpsMetrics>): DpsMetrics;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DpsMetrics): DpsMetrics;
    private binaryReadMap4;
    internalBinaryWrite(message: DpsMetrics, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.DpsMetrics
 */
export declare const DpsMetrics: DpsMetrics$Type;
declare class PlayerMetrics$Type extends MessageType<PlayerMetrics> {
    constructor();
    create(value?: PartialMessage<PlayerMetrics>): PlayerMetrics;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PlayerMetrics): PlayerMetrics;
    internalBinaryWrite(message: PlayerMetrics, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.PlayerMetrics
 */
export declare const PlayerMetrics: PlayerMetrics$Type;
declare class PartyMetrics$Type extends MessageType<PartyMetrics> {
    constructor();
    create(value?: PartialMessage<PartyMetrics>): PartyMetrics;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PartyMetrics): PartyMetrics;
    internalBinaryWrite(message: PartyMetrics, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.PartyMetrics
 */
export declare const PartyMetrics: PartyMetrics$Type;
declare class RaidMetrics$Type extends MessageType<RaidMetrics> {
    constructor();
    create(value?: PartialMessage<RaidMetrics>): RaidMetrics;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: RaidMetrics): RaidMetrics;
    internalBinaryWrite(message: RaidMetrics, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.RaidMetrics
 */
export declare const RaidMetrics: RaidMetrics$Type;
declare class TargetMetrics$Type extends MessageType<TargetMetrics> {
    constructor();
    create(value?: PartialMessage<TargetMetrics>): TargetMetrics;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: TargetMetrics): TargetMetrics;
    internalBinaryWrite(message: TargetMetrics, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.TargetMetrics
 */
export declare const TargetMetrics: TargetMetrics$Type;
declare class EncounterMetrics$Type extends MessageType<EncounterMetrics> {
    constructor();
    create(value?: PartialMessage<EncounterMetrics>): EncounterMetrics;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: EncounterMetrics): EncounterMetrics;
    internalBinaryWrite(message: EncounterMetrics, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.EncounterMetrics
 */
export declare const EncounterMetrics: EncounterMetrics$Type;
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
declare class RaidSimResult$Type extends MessageType<RaidSimResult> {
    constructor();
    create(value?: PartialMessage<RaidSimResult>): RaidSimResult;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: RaidSimResult): RaidSimResult;
    internalBinaryWrite(message: RaidSimResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.RaidSimResult
 */
export declare const RaidSimResult: RaidSimResult$Type;
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
    internalBinaryWrite(message: IndividualSimResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.IndividualSimResult
 */
export declare const IndividualSimResult: IndividualSimResult$Type;
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
declare class PlayerStats$Type extends MessageType<PlayerStats> {
    constructor();
    create(value?: PartialMessage<PlayerStats>): PlayerStats;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PlayerStats): PlayerStats;
    internalBinaryWrite(message: PlayerStats, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.PlayerStats
 */
export declare const PlayerStats: PlayerStats$Type;
declare class PartyStats$Type extends MessageType<PartyStats> {
    constructor();
    create(value?: PartialMessage<PartyStats>): PartyStats;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PartyStats): PartyStats;
    internalBinaryWrite(message: PartyStats, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.PartyStats
 */
export declare const PartyStats: PartyStats$Type;
declare class RaidStats$Type extends MessageType<RaidStats> {
    constructor();
    create(value?: PartialMessage<RaidStats>): RaidStats;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: RaidStats): RaidStats;
    internalBinaryWrite(message: RaidStats, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.RaidStats
 */
export declare const RaidStats: RaidStats$Type;
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
