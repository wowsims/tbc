import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
// @generated message type with reflection information, may provide speed optimized methods
class DruidTalents$Type extends MessageType {
    constructor() {
        super("api.DruidTalents", [
            { no: 1, name: "starlight_wrath", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value) {
        const message = { starlightWrath: 0 };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 starlight_wrath */ 1:
                    message.starlightWrath = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* int32 starlight_wrath = 1; */
        if (message.starlightWrath !== 0)
            writer.tag(1, WireType.Varint).int32(message.starlightWrath);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message api.DruidTalents
 */
export const DruidTalents = new DruidTalents$Type();
// @generated message type with reflection information, may provide speed optimized methods
class BalanceDruid$Type extends MessageType {
    constructor() {
        super("api.BalanceDruid", [
            { no: 1, name: "agent", kind: "message", T: () => BalanceDruid_Agent },
            { no: 2, name: "talents", kind: "message", T: () => DruidTalents },
            { no: 3, name: "options", kind: "message", T: () => BalanceDruid_Options }
        ]);
    }
    create(value) {
        const message = {};
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* api.BalanceDruid.Agent agent */ 1:
                    message.agent = BalanceDruid_Agent.internalBinaryRead(reader, reader.uint32(), options, message.agent);
                    break;
                case /* api.DruidTalents talents */ 2:
                    message.talents = DruidTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* api.BalanceDruid.Options options */ 3:
                    message.options = BalanceDruid_Options.internalBinaryRead(reader, reader.uint32(), options, message.options);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* api.BalanceDruid.Agent agent = 1; */
        if (message.agent)
            BalanceDruid_Agent.internalBinaryWrite(message.agent, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* api.DruidTalents talents = 2; */
        if (message.talents)
            DruidTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* api.BalanceDruid.Options options = 3; */
        if (message.options)
            BalanceDruid_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message api.BalanceDruid
 */
export const BalanceDruid = new BalanceDruid$Type();
// @generated message type with reflection information, may provide speed optimized methods
class BalanceDruid_Agent$Type extends MessageType {
    constructor() {
        super("api.BalanceDruid.Agent", []);
    }
    create(value) {
        const message = {};
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        return target ?? this.create();
    }
    internalBinaryWrite(message, writer, options) {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message api.BalanceDruid.Agent
 */
export const BalanceDruid_Agent = new BalanceDruid_Agent$Type();
// @generated message type with reflection information, may provide speed optimized methods
class BalanceDruid_Options$Type extends MessageType {
    constructor() {
        super("api.BalanceDruid.Options", []);
    }
    create(value) {
        const message = {};
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        return target ?? this.create();
    }
    internalBinaryWrite(message, writer, options) {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message api.BalanceDruid.Options
 */
export const BalanceDruid_Options = new BalanceDruid_Options$Type();
