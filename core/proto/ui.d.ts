import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { RaidTarget } from "./common";
/**
 * A buff bot placed in a raid.
 *
 * @generated from protobuf message proto.BuffBot
 */
export interface BuffBot {
    /**
     * Uniquely identifies which buffbot this is.
     *
     * @generated from protobuf field: string id = 1;
     */
    id: string;
    /**
     * @generated from protobuf field: int32 raid_index = 2;
     */
    raidIndex: number;
    /**
     * The assigned player to innervate. Only used for druid buffbots.
     *
     * @generated from protobuf field: proto.RaidTarget innervate_assignment = 3;
     */
    innervateAssignment?: RaidTarget;
    /**
     * The assigned player to PI. Only used for disc priest buffbots.
     *
     * @generated from protobuf field: proto.RaidTarget power_infusion_assignment = 4;
     */
    powerInfusionAssignment?: RaidTarget;
}
/**
 * @generated from protobuf message proto.BlessingsAssignment
 */
export interface BlessingsAssignment {
    /**
     * Index corresponds to Spec that the blessing should be applied to.
     *
     * @generated from protobuf field: repeated proto.Blessings blessings = 1;
     */
    blessings: Blessings[];
}
/**
 * @generated from protobuf message proto.BlessingsAssignments
 */
export interface BlessingsAssignments {
    /**
     * Assignments for each paladin.
     *
     * @generated from protobuf field: repeated proto.BlessingsAssignment paladins = 1;
     */
    paladins: BlessingsAssignment[];
}
/**
 * @generated from protobuf enum proto.Blessings
 */
export declare enum Blessings {
    /**
     * @generated from protobuf enum value: BlessingUnknown = 0;
     */
    BlessingUnknown = 0,
    /**
     * @generated from protobuf enum value: BlessingOfKings = 1;
     */
    BlessingOfKings = 1,
    /**
     * @generated from protobuf enum value: BlessingOfMight = 2;
     */
    BlessingOfMight = 2,
    /**
     * @generated from protobuf enum value: BlessingOfSalvation = 3;
     */
    BlessingOfSalvation = 3,
    /**
     * @generated from protobuf enum value: BlessingOfWisdom = 4;
     */
    BlessingOfWisdom = 4
}
declare class BuffBot$Type extends MessageType<BuffBot> {
    constructor();
    create(value?: PartialMessage<BuffBot>): BuffBot;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: BuffBot): BuffBot;
    internalBinaryWrite(message: BuffBot, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.BuffBot
 */
export declare const BuffBot: BuffBot$Type;
declare class BlessingsAssignment$Type extends MessageType<BlessingsAssignment> {
    constructor();
    create(value?: PartialMessage<BlessingsAssignment>): BlessingsAssignment;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: BlessingsAssignment): BlessingsAssignment;
    internalBinaryWrite(message: BlessingsAssignment, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.BlessingsAssignment
 */
export declare const BlessingsAssignment: BlessingsAssignment$Type;
declare class BlessingsAssignments$Type extends MessageType<BlessingsAssignments> {
    constructor();
    create(value?: PartialMessage<BlessingsAssignments>): BlessingsAssignments;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: BlessingsAssignments): BlessingsAssignments;
    internalBinaryWrite(message: BlessingsAssignments, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.BlessingsAssignments
 */
export declare const BlessingsAssignments: BlessingsAssignments$Type;
export {};
