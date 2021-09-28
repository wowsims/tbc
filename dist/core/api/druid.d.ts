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
     * @generated from protobuf field: api.BalanceDruid.BalanceDruidAgent agent = 1;
     */
    agent?: BalanceDruid_BalanceDruidAgent;
    /**
     * @generated from protobuf field: api.DruidTalents talents = 2;
     */
    talents?: DruidTalents;
    /**
     * @generated from protobuf field: api.BalanceDruid.BalanceDruidOptions options = 3;
     */
    options?: BalanceDruid_BalanceDruidOptions;
}
/**
 * @generated from protobuf message api.BalanceDruid.BalanceDruidAgent
 */
export interface BalanceDruid_BalanceDruidAgent {
}
/**
 * @generated from protobuf message api.BalanceDruid.BalanceDruidOptions
 */
export interface BalanceDruid_BalanceDruidOptions {
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
declare class BalanceDruid_BalanceDruidAgent$Type extends MessageType<BalanceDruid_BalanceDruidAgent> {
    constructor();
    create(value?: PartialMessage<BalanceDruid_BalanceDruidAgent>): BalanceDruid_BalanceDruidAgent;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: BalanceDruid_BalanceDruidAgent): BalanceDruid_BalanceDruidAgent;
    internalBinaryWrite(message: BalanceDruid_BalanceDruidAgent, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.BalanceDruid.BalanceDruidAgent
 */
export declare const BalanceDruid_BalanceDruidAgent: BalanceDruid_BalanceDruidAgent$Type;
declare class BalanceDruid_BalanceDruidOptions$Type extends MessageType<BalanceDruid_BalanceDruidOptions> {
    constructor();
    create(value?: PartialMessage<BalanceDruid_BalanceDruidOptions>): BalanceDruid_BalanceDruidOptions;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: BalanceDruid_BalanceDruidOptions): BalanceDruid_BalanceDruidOptions;
    internalBinaryWrite(message: BalanceDruid_BalanceDruidOptions, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.BalanceDruid.BalanceDruidOptions
 */
export declare const BalanceDruid_BalanceDruidOptions: BalanceDruid_BalanceDruidOptions$Type;
export {};
