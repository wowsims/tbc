import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { WireType } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
// @generated message type with reflection information, may provide speed optimized methods
class CharacterStatsTestResult$Type extends MessageType {
    constructor() {
        super("proto.CharacterStatsTestResult", [
            { no: 1, name: "final_stats", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { finalStats: [] };
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
                case /* repeated double final_stats */ 1:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.finalStats.push(reader.double());
                    else
                        message.finalStats.push(reader.double());
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
        /* repeated double final_stats = 1; */
        if (message.finalStats.length) {
            writer.tag(1, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.finalStats.length; i++)
                writer.double(message.finalStats[i]);
            writer.join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.CharacterStatsTestResult
 */
export const CharacterStatsTestResult = new CharacterStatsTestResult$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StatWeightsTestResult$Type extends MessageType {
    constructor() {
        super("proto.StatWeightsTestResult", [
            { no: 1, name: "weights", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { weights: [] };
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
                case /* repeated double weights */ 1:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.weights.push(reader.double());
                    else
                        message.weights.push(reader.double());
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
        /* repeated double weights = 1; */
        if (message.weights.length) {
            writer.tag(1, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.weights.length; i++)
                writer.double(message.weights[i]);
            writer.join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.StatWeightsTestResult
 */
export const StatWeightsTestResult = new StatWeightsTestResult$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DpsTestResult$Type extends MessageType {
    constructor() {
        super("proto.DpsTestResult", [
            { no: 1, name: "dps", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 2, name: "tps", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 3, name: "dtps", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { dps: 0, tps: 0, dtps: 0 };
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
                case /* double dps */ 1:
                    message.dps = reader.double();
                    break;
                case /* double tps */ 2:
                    message.tps = reader.double();
                    break;
                case /* double dtps */ 3:
                    message.dtps = reader.double();
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
        /* double dps = 1; */
        if (message.dps !== 0)
            writer.tag(1, WireType.Bit64).double(message.dps);
        /* double tps = 2; */
        if (message.tps !== 0)
            writer.tag(2, WireType.Bit64).double(message.tps);
        /* double dtps = 3; */
        if (message.dtps !== 0)
            writer.tag(3, WireType.Bit64).double(message.dtps);
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
            { no: 2, name: "character_stats_results", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "message", T: () => CharacterStatsTestResult } },
            { no: 3, name: "stat_weights_results", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "message", T: () => StatWeightsTestResult } },
            { no: 1, name: "dps_results", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "message", T: () => DpsTestResult } }
        ]);
    }
    create(value) {
        const message = { characterStatsResults: {}, statWeightsResults: {}, dpsResults: {} };
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
                case /* map<string, proto.CharacterStatsTestResult> character_stats_results */ 2:
                    this.binaryReadMap2(message.characterStatsResults, reader, options);
                    break;
                case /* map<string, proto.StatWeightsTestResult> stat_weights_results */ 3:
                    this.binaryReadMap3(message.statWeightsResults, reader, options);
                    break;
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
    binaryReadMap2(map, reader, options) {
        let len = reader.uint32(), end = reader.pos + len, key, val;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case 1:
                    key = reader.string();
                    break;
                case 2:
                    val = CharacterStatsTestResult.internalBinaryRead(reader, reader.uint32(), options);
                    break;
                default: throw new globalThis.Error("unknown map entry field for field proto.TestSuiteResult.character_stats_results");
            }
        }
        map[key ?? ""] = val ?? CharacterStatsTestResult.create();
    }
    binaryReadMap3(map, reader, options) {
        let len = reader.uint32(), end = reader.pos + len, key, val;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case 1:
                    key = reader.string();
                    break;
                case 2:
                    val = StatWeightsTestResult.internalBinaryRead(reader, reader.uint32(), options);
                    break;
                default: throw new globalThis.Error("unknown map entry field for field proto.TestSuiteResult.stat_weights_results");
            }
        }
        map[key ?? ""] = val ?? StatWeightsTestResult.create();
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
        /* map<string, proto.CharacterStatsTestResult> character_stats_results = 2; */
        for (let k of Object.keys(message.characterStatsResults)) {
            writer.tag(2, WireType.LengthDelimited).fork().tag(1, WireType.LengthDelimited).string(k);
            writer.tag(2, WireType.LengthDelimited).fork();
            CharacterStatsTestResult.internalBinaryWrite(message.characterStatsResults[k], writer, options);
            writer.join().join();
        }
        /* map<string, proto.StatWeightsTestResult> stat_weights_results = 3; */
        for (let k of Object.keys(message.statWeightsResults)) {
            writer.tag(3, WireType.LengthDelimited).fork().tag(1, WireType.LengthDelimited).string(k);
            writer.tag(2, WireType.LengthDelimited).fork();
            StatWeightsTestResult.internalBinaryWrite(message.statWeightsResults[k], writer, options);
            writer.join().join();
        }
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
