import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message proto.DruidTalents
 */
export interface DruidTalents {
    /**
     * @generated from protobuf field: int32 starlight_wrath = 1;
     */
    starlightWrath: number;
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
}
/**
 * @generated from protobuf message proto.BalanceDruid.Options
 */
export interface BalanceDruid_Options {
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
