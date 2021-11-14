import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { RaidTarget } from "./common";
/**
 * @generated from protobuf message proto.DruidTalents
 */
export interface DruidTalents {
    /**
     * Balance
     *
     * @generated from protobuf field: int32 starlight_wrath = 1;
     */
    starlightWrath: number;
    /**
     * @generated from protobuf field: int32 focused_starlight = 2;
     */
    focusedStarlight: number;
    /**
     * @generated from protobuf field: int32 improved_moonfire = 3;
     */
    improvedMoonfire: number;
    /**
     * @generated from protobuf field: int32 brambles = 4;
     */
    brambles: number;
    /**
     * @generated from protobuf field: bool insect_swarm = 5;
     */
    insectSwarm: boolean;
    /**
     * @generated from protobuf field: int32 vengeance = 6;
     */
    vengeance: number;
    /**
     * @generated from protobuf field: int32 lunar_guidance = 7;
     */
    lunarGuidance: number;
    /**
     * @generated from protobuf field: bool natures_grace = 8;
     */
    naturesGrace: boolean;
    /**
     * @generated from protobuf field: int32 moonglow = 9;
     */
    moonglow: number;
    /**
     * @generated from protobuf field: int32 moonfury = 10;
     */
    moonfury: number;
    /**
     * @generated from protobuf field: int32 balance_of_power = 11;
     */
    balanceOfPower: number;
    /**
     * @generated from protobuf field: int32 dreamstate = 12;
     */
    dreamstate: number;
    /**
     * @generated from protobuf field: bool moonkin_form = 13;
     */
    moonkinForm: boolean;
    /**
     * @generated from protobuf field: int32 improved_faerie_fire = 14;
     */
    improvedFaerieFire: number;
    /**
     * @generated from protobuf field: int32 wrath_of_cenarius = 15;
     */
    wrathOfCenarius: number;
    /**
     * @generated from protobuf field: bool force_of_nature = 16;
     */
    forceOfNature: boolean;
    /**
     * Feral Combat
     *
     * @generated from protobuf field: int32 ferocity = 17;
     */
    ferocity: number;
    /**
     * @generated from protobuf field: int32 feral_aggresion = 18;
     */
    feralAggresion: number;
    /**
     * @generated from protobuf field: int32 sharpened_claws = 19;
     */
    sharpenedClaws: number;
    /**
     * @generated from protobuf field: int32 shredding_attacks = 20;
     */
    shreddingAttacks: number;
    /**
     * @generated from protobuf field: int32 predatory_strikes = 21;
     */
    predatoryStrikes: number;
    /**
     * @generated from protobuf field: int32 primal_fury = 22;
     */
    primalFury: number;
    /**
     * @generated from protobuf field: int32 savage_fury = 23;
     */
    savageFury: number;
    /**
     * @generated from protobuf field: bool faerie_fire = 24;
     */
    faerieFire: boolean;
    /**
     * @generated from protobuf field: int32 heart_of_the_wild = 25;
     */
    heartOfTheWild: number;
    /**
     * @generated from protobuf field: int32 survival_of_the_fittest = 26;
     */
    survivalOfTheFittest: number;
    /**
     * @generated from protobuf field: bool leader_of_the_pack = 27;
     */
    leaderOfThePack: boolean;
    /**
     * @generated from protobuf field: int32 improved_leader_of_the_pack = 28;
     */
    improvedLeaderOfThePack: number;
    /**
     * @generated from protobuf field: int32 predatory_instincts = 29;
     */
    predatoryInstincts: number;
    /**
     * @generated from protobuf field: bool mangle = 30;
     */
    mangle: boolean;
    /**
     * Restoration
     *
     * @generated from protobuf field: int32 improved_mark_of_the_wild = 31;
     */
    improvedMarkOfTheWild: number;
    /**
     * @generated from protobuf field: int32 furor = 32;
     */
    furor: number;
    /**
     * @generated from protobuf field: int32 naturalist = 33;
     */
    naturalist: number;
    /**
     * @generated from protobuf field: int32 natural_shapeshifter = 34;
     */
    naturalShapeshifter: number;
    /**
     * @generated from protobuf field: int32 intensity = 35;
     */
    intensity: number;
    /**
     * @generated from protobuf field: bool omen_of_clarity = 36;
     */
    omenOfClarity: boolean;
    /**
     * @generated from protobuf field: bool natures_swiftness = 37;
     */
    naturesSwiftness: boolean;
    /**
     * @generated from protobuf field: int32 living_spirit = 38;
     */
    livingSpirit: number;
    /**
     * @generated from protobuf field: int32 natural_perfection = 39;
     */
    naturalPerfection: number;
}
/**
 * @generated from protobuf message proto.BalanceDruid
 */
export interface BalanceDruid {
    /**
     * @generated from protobuf field: proto.BalanceDruid.Rotation rotation = 1;
     */
    rotation?: BalanceDruid_Rotation;
    /**
     * @generated from protobuf field: proto.DruidTalents talents = 2;
     */
    talents?: DruidTalents;
    /**
     * @generated from protobuf field: proto.BalanceDruid.Options options = 3;
     */
    options?: BalanceDruid_Options;
}
/**
 * @generated from protobuf message proto.BalanceDruid.Rotation
 */
export interface BalanceDruid_Rotation {
    /**
     * @generated from protobuf field: proto.BalanceDruid.Rotation.PrimarySpell primary_spell = 1;
     */
    primarySpell: BalanceDruid_Rotation_PrimarySpell;
    /**
     * @generated from protobuf field: bool faerie_fire = 2;
     */
    faerieFire: boolean;
    /**
     * @generated from protobuf field: bool insect_swarm = 3;
     */
    insectSwarm: boolean;
    /**
     * @generated from protobuf field: bool moonfire = 4;
     */
    moonfire: boolean;
    /**
     * @generated from protobuf field: bool melee_while_oom = 5;
     */
    meleeWhileOom: boolean;
}
/**
 * @generated from protobuf enum proto.BalanceDruid.Rotation.PrimarySpell
 */
export declare enum BalanceDruid_Rotation_PrimarySpell {
    /**
     * @generated from protobuf enum value: Unknown = 0;
     */
    Unknown = 0,
    /**
     * @generated from protobuf enum value: Starfire = 1;
     */
    Starfire = 1,
    /**
     * @generated from protobuf enum value: Starfire6 = 2;
     */
    Starfire6 = 2,
    /**
     * @generated from protobuf enum value: Wrath = 3;
     */
    Wrath = 3
}
/**
 * @generated from protobuf message proto.BalanceDruid.Options
 */
export interface BalanceDruid_Options {
    /**
     * @generated from protobuf field: bool omen_of_clarity = 1;
     */
    omenOfClarity: boolean;
    /**
     * @generated from protobuf field: proto.RaidTarget innervate_target = 2;
     */
    innervateTarget?: RaidTarget;
}
declare class DruidTalents$Type extends MessageType<DruidTalents> {
    constructor();
    create(value?: PartialMessage<DruidTalents>): DruidTalents;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DruidTalents): DruidTalents;
    internalBinaryWrite(message: DruidTalents, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.DruidTalents
 */
export declare const DruidTalents: DruidTalents$Type;
declare class BalanceDruid$Type extends MessageType<BalanceDruid> {
    constructor();
    create(value?: PartialMessage<BalanceDruid>): BalanceDruid;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: BalanceDruid): BalanceDruid;
    internalBinaryWrite(message: BalanceDruid, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.BalanceDruid
 */
export declare const BalanceDruid: BalanceDruid$Type;
declare class BalanceDruid_Rotation$Type extends MessageType<BalanceDruid_Rotation> {
    constructor();
    create(value?: PartialMessage<BalanceDruid_Rotation>): BalanceDruid_Rotation;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: BalanceDruid_Rotation): BalanceDruid_Rotation;
    internalBinaryWrite(message: BalanceDruid_Rotation, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.BalanceDruid.Rotation
 */
export declare const BalanceDruid_Rotation: BalanceDruid_Rotation$Type;
declare class BalanceDruid_Options$Type extends MessageType<BalanceDruid_Options> {
    constructor();
    create(value?: PartialMessage<BalanceDruid_Options>): BalanceDruid_Options;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: BalanceDruid_Options): BalanceDruid_Options;
    internalBinaryWrite(message: BalanceDruid_Options, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.BalanceDruid.Options
 */
export declare const BalanceDruid_Options: BalanceDruid_Options$Type;
export {};
