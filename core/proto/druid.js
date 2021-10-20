import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
// @generated message type with reflection information, may provide speed optimized methods
class DruidTalents$Type extends MessageType {
    constructor() {
        super("proto.DruidTalents", [
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
 * @generated MessageType for protobuf message proto.DruidTalents
 */
export const DruidTalents = new DruidTalents$Type();
// @generated message type with reflection information, may provide speed optimized methods
class BalanceDruid$Type extends MessageType {
    constructor() {
        super("proto.BalanceDruid", [
            { no: 1, name: "rotation", kind: "message", T: () => BalanceDruid_Rotation },
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
                case /* proto.BalanceDruid.Rotation rotation */ 1:
                    message.rotation = BalanceDruid_Rotation.internalBinaryRead(reader, reader.uint32(), options, message.rotation);
                    break;
                case /* proto.DruidTalents talents */ 2:
                    message.talents = DruidTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* proto.BalanceDruid.Options options */ 3:
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
        /* proto.BalanceDruid.Rotation rotation = 1; */
        if (message.rotation)
            BalanceDruid_Rotation.internalBinaryWrite(message.rotation, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.DruidTalents talents = 2; */
        if (message.talents)
            DruidTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.BalanceDruid.Options options = 3; */
        if (message.options)
            BalanceDruid_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.BalanceDruid
 */
export const BalanceDruid = new BalanceDruid$Type();
// @generated message type with reflection information, may provide speed optimized methods
class BalanceDruid_Rotation$Type extends MessageType {
    constructor() {
        super("proto.BalanceDruid.Rotation", []);
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
 * @generated MessageType for protobuf message proto.BalanceDruid.Rotation
 */
export const BalanceDruid_Rotation = new BalanceDruid_Rotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class BalanceDruid_Options$Type extends MessageType {
    constructor() {
        super("proto.BalanceDruid.Options", []);
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
 * @generated MessageType for protobuf message proto.BalanceDruid.Options
 */
export const BalanceDruid_Options = new BalanceDruid_Options$Type();
