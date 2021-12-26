import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
// @generated message type with reflection information, may provide speed optimized methods
class DpsTestResult$Type extends MessageType {
    constructor() {
        super("proto.DpsTestResult", [
            { no: 2, name: "dps", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { dps: 0 };
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
                case /* double dps */ 2:
                    message.dps = reader.double();
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
        /* double dps = 2; */
        if (message.dps !== 0)
            writer.tag(2, WireType.Bit64).double(message.dps);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.DpsTestResult
 */
export const DpsTestResult = new DpsTestResult$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TestSuiteResult$Type extends MessageType {
    constructor() {
        super("proto.TestSuiteResult", [
            { no: 1, name: "dps_results", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "message", T: () => DpsTestResult } }
        ]);
    }
    create(value) {
        const message = { dpsResults: {} };
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
                case /* map<string, proto.DpsTestResult> dps_results */ 1:
                    this.binaryReadMap1(message.dpsResults, reader, options);
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
    binaryReadMap1(map, reader, options) {
        let len = reader.uint32(), end = reader.pos + len, key, val;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case 1:
                    key = reader.string();
                    break;
                case 2:
                    val = DpsTestResult.internalBinaryRead(reader, reader.uint32(), options);
                    break;
                default: throw new globalThis.Error("unknown map entry field for field proto.TestSuiteResult.dps_results");
            }
        }
        map[key ?? ""] = val ?? DpsTestResult.create();
    }
    internalBinaryWrite(message, writer, options) {
        /* map<string, proto.DpsTestResult> dps_results = 1; */
        for (let k of Object.keys(message.dpsResults)) {
            writer.tag(1, WireType.LengthDelimited).fork().tag(1, WireType.LengthDelimited).string(k);
            writer.tag(2, WireType.LengthDelimited).fork();
            DpsTestResult.internalBinaryWrite(message.dpsResults[k], writer, options);
            writer.join().join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.TestSuiteResult
 */
export const TestSuiteResult = new TestSuiteResult$Type();
