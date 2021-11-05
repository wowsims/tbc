import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message proto.PriestTalents
 */
export interface PriestTalents {
    /**
     * Discipline
     *
     * @generated from protobuf field: int32 wand_specialization = 1;
     */
    wandSpecialization: number;
    /**
     * @generated from protobuf field: bool inner_focus = 2;
     */
    innerFocus: boolean;
    /**
     * @generated from protobuf field: int32 meditation = 3;
     */
    meditation: number;
    /**
     * @generated from protobuf field: int32 mental_agility = 4;
     */
    mentalAgility: number;
    /**
     * @generated from protobuf field: int32 mental_strength = 5;
     */
    mentalStrength: number;
    /**
     * @generated from protobuf field: bool divine_spirit = 6;
     */
    divineSpirit: boolean;
    /**
     * @generated from protobuf field: int32 improved_divine_spirit = 7;
     */
    improvedDivineSpirit: number;
    /**
     * @generated from protobuf field: int32 focused_power = 8;
     */
    focusedPower: number;
    /**
     * @generated from protobuf field: int32 force_of_will = 9;
     */
    forceOfWill: number;
    /**
     * @generated from protobuf field: bool power_infusion = 10;
     */
    powerInfusion: boolean;
    /**
     * @generated from protobuf field: int32 enlightenment = 11;
     */
    enlightenment: number;
    /**
     * Holy
     *
     * @generated from protobuf field: int32 holy_specialization = 12;
     */
    holySpecialization: number;
    /**
     * @generated from protobuf field: int32 divine_fury = 13;
     */
    divineFury: number;
    /**
     * @generated from protobuf field: bool holy_nova = 14;
     */
    holyNova: boolean;
    /**
     * @generated from protobuf field: int32 searing_light = 15;
     */
    searingLight: number;
    /**
     * @generated from protobuf field: int32 spiritual_guidance = 16;
     */
    spiritualGuidance: number;
    /**
     * @generated from protobuf field: int32 surge_of_light = 17;
     */
    surgeOfLight: number;
    /**
     * Shadow
     *
     * @generated from protobuf field: int32 improved_shadow_word_pain = 18;
     */
    improvedShadowWordPain: number;
    /**
     * @generated from protobuf field: int32 shadow_focus = 19;
     */
    shadowFocus: number;
    /**
     * @generated from protobuf field: int32 improved_mind_blast = 20;
     */
    improvedMindBlast: number;
    /**
     * @generated from protobuf field: bool mind_flay = 21;
     */
    mindFlay: boolean;
    /**
     * @generated from protobuf field: int32 shadow_weaving = 22;
     */
    shadowWeaving: number;
    /**
     * @generated from protobuf field: bool vampiric_embrace = 23;
     */
    vampiricEmbrace: boolean;
    /**
     * @generated from protobuf field: int32 improved_vampiric_embrace = 24;
     */
    improvedVampiricEmbrace: number;
    /**
     * @generated from protobuf field: int32 focused_mind = 25;
     */
    focusedMind: number;
    /**
     * @generated from protobuf field: int32 darkness = 26;
     */
    darkness: number;
    /**
     * @generated from protobuf field: bool shadowform = 27;
     */
    shadowform: boolean;
    /**
     * @generated from protobuf field: int32 shadow_power = 28;
     */
    shadowPower: number;
    /**
     * @generated from protobuf field: int32 misery = 29;
     */
    misery: number;
    /**
     * @generated from protobuf field: bool vampiric_touch = 30;
     */
    vampiricTouch: boolean;
}
/**
 * @generated from protobuf message proto.ShadowPriest
 */
export interface ShadowPriest {
    /**
     * @generated from protobuf field: proto.ShadowPriest.Rotation rotation = 1;
     */
    rotation?: ShadowPriest_Rotation;
    /**
     * @generated from protobuf field: proto.PriestTalents talents = 2;
     */
    talents?: PriestTalents;
    /**
     * @generated from protobuf field: proto.ShadowPriest.Options options = 3;
     */
    options?: ShadowPriest_Options;
}
/**
 * @generated from protobuf message proto.ShadowPriest.Rotation
 */
export interface ShadowPriest_Rotation {
    /**
     * @generated from protobuf field: proto.ShadowPriest.Rotation.RotationType type = 1;
     */
    type: ShadowPriest_Rotation_RotationType;
}
/**
 * @generated from protobuf enum proto.ShadowPriest.Rotation.RotationType
 */
export declare enum ShadowPriest_Rotation_RotationType {
    /**
     * @generated from protobuf enum value: Unknown = 0;
     */
    Unknown = 0,
    /**
     * @generated from protobuf enum value: Standard = 1;
     */
    Standard = 1
}
/**
 * @generated from protobuf message proto.ShadowPriest.Options
 */
export interface ShadowPriest_Options {
}
declare class PriestTalents$Type extends MessageType<PriestTalents> {
    constructor();
    create(value?: PartialMessage<PriestTalents>): PriestTalents;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PriestTalents): PriestTalents;
    internalBinaryWrite(message: PriestTalents, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.PriestTalents
 */
export declare const PriestTalents: PriestTalents$Type;
declare class ShadowPriest$Type extends MessageType<ShadowPriest> {
    constructor();
    create(value?: PartialMessage<ShadowPriest>): ShadowPriest;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ShadowPriest): ShadowPriest;
    internalBinaryWrite(message: ShadowPriest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ShadowPriest
 */
export declare const ShadowPriest: ShadowPriest$Type;
declare class ShadowPriest_Rotation$Type extends MessageType<ShadowPriest_Rotation> {
    constructor();
    create(value?: PartialMessage<ShadowPriest_Rotation>): ShadowPriest_Rotation;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ShadowPriest_Rotation): ShadowPriest_Rotation;
    internalBinaryWrite(message: ShadowPriest_Rotation, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ShadowPriest.Rotation
 */
export declare const ShadowPriest_Rotation: ShadowPriest_Rotation$Type;
declare class ShadowPriest_Options$Type extends MessageType<ShadowPriest_Options> {
    constructor();
    create(value?: PartialMessage<ShadowPriest_Options>): ShadowPriest_Options;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ShadowPriest_Options): ShadowPriest_Options;
    internalBinaryWrite(message: ShadowPriest_Options, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ShadowPriest.Options
 */
export declare const ShadowPriest_Options: ShadowPriest_Options$Type;
export {};
