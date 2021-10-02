import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message api.DruidTalents
 */
export interface DruidTalents {
    /**
     * @generated from protobuf field: int32 starlight_wrath = 1;
     */
    starlightWrath: number;
}
/**
 * @generated from protobuf message api.BalanceDruid
 */
export interface BalanceDruid {
    /**
     * @generated from protobuf field: api.BalanceDruid.Agent agent = 1;
     */
    agent?: BalanceDruid_Agent;
    /**
     * @generated from protobuf field: api.DruidTalents talents = 2;
     */
    talents?: DruidTalents;
    /**
     * @generated from protobuf field: api.BalanceDruid.Options options = 3;
     */
    options?: BalanceDruid_Options;
}
/**
 * @generated from protobuf message api.BalanceDruid.Agent
 */
export interface BalanceDruid_Agent {
}
/**
 * @generated from protobuf message api.BalanceDruid.Options
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
 * @generated MessageType for protobuf message api.DruidTalents
 */
export declare const DruidTalents: DruidTalents$Type;
declare class BalanceDruid$Type extends MessageType<BalanceDruid> {
    constructor();
    create(value?: PartialMessage<BalanceDruid>): BalanceDruid;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: BalanceDruid): BalanceDruid;
    internalBinaryWrite(message: BalanceDruid, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.BalanceDruid
 */
export declare const BalanceDruid: BalanceDruid$Type;
declare class BalanceDruid_Agent$Type extends MessageType<BalanceDruid_Agent> {
    constructor();
    create(value?: PartialMessage<BalanceDruid_Agent>): BalanceDruid_Agent;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: BalanceDruid_Agent): BalanceDruid_Agent;
    internalBinaryWrite(message: BalanceDruid_Agent, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.BalanceDruid.Agent
 */
export declare const BalanceDruid_Agent: BalanceDruid_Agent$Type;
declare class BalanceDruid_Options$Type extends MessageType<BalanceDruid_Options> {
    constructor();
    create(value?: PartialMessage<BalanceDruid_Options>): BalanceDruid_Options;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: BalanceDruid_Options): BalanceDruid_Options;
    internalBinaryWrite(message: BalanceDruid_Options, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.BalanceDruid.Options
 */
export declare const BalanceDruid_Options: BalanceDruid_Options$Type;
export {};
