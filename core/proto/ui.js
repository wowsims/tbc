import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
import { RaidTarget } from './common.js';
/**
 * @generated from protobuf enum proto.Blessings
 */
export var Blessings;
(function (Blessings) {
    /**
     * @generated from protobuf enum value: BlessingUnknown = 0;
     */
    Blessings[Blessings["BlessingUnknown"] = 0] = "BlessingUnknown";
    /**
     * @generated from protobuf enum value: BlessingOfKings = 1;
     */
    Blessings[Blessings["BlessingOfKings"] = 1] = "BlessingOfKings";
    /**
     * @generated from protobuf enum value: BlessingOfMight = 2;
     */
    Blessings[Blessings["BlessingOfMight"] = 2] = "BlessingOfMight";
    /**
     * @generated from protobuf enum value: BlessingOfSalvation = 3;
     */
    Blessings[Blessings["BlessingOfSalvation"] = 3] = "BlessingOfSalvation";
    /**
     * @generated from protobuf enum value: BlessingOfWisdom = 4;
     */
    Blessings[Blessings["BlessingOfWisdom"] = 4] = "BlessingOfWisdom";
})(Blessings || (Blessings = {}));
// @generated message type with reflection information, may provide speed optimized methods
class BuffBot$Type extends MessageType {
    constructor() {
        super("proto.BuffBot", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "raid_index", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "innervate_assignment", kind: "message", T: () => RaidTarget },
            { no: 4, name: "power_infusion_assignment", kind: "message", T: () => RaidTarget }
        ]);
    }
    create(value) {
        const message = { id: "", raidIndex: 0 };
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
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                case /* int32 raid_index */ 2:
                    message.raidIndex = reader.int32();
                    break;
                case /* proto.RaidTarget innervate_assignment */ 3:
                    message.innervateAssignment = RaidTarget.internalBinaryRead(reader, reader.uint32(), options, message.innervateAssignment);
                    break;
                case /* proto.RaidTarget power_infusion_assignment */ 4:
                    message.powerInfusionAssignment = RaidTarget.internalBinaryRead(reader, reader.uint32(), options, message.powerInfusionAssignment);
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
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        /* int32 raid_index = 2; */
        if (message.raidIndex !== 0)
            writer.tag(2, WireType.Varint).int32(message.raidIndex);
        /* proto.RaidTarget innervate_assignment = 3; */
        if (message.innervateAssignment)
            RaidTarget.internalBinaryWrite(message.innervateAssignment, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* proto.RaidTarget power_infusion_assignment = 4; */
        if (message.powerInfusionAssignment)
            RaidTarget.internalBinaryWrite(message.powerInfusionAssignment, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.BuffBot
 */
export const BuffBot = new BuffBot$Type();
// @generated message type with reflection information, may provide speed optimized methods
class BlessingsAssignment$Type extends MessageType {
    constructor() {
        super("proto.BlessingsAssignment", [
            { no: 1, name: "blessings", kind: "enum", repeat: 1 /*RepeatType.PACKED*/, T: () => ["proto.Blessings", Blessings] }
        ]);
    }
    create(value) {
        const message = { blessings: [] };
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
                case /* repeated proto.Blessings blessings */ 1:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.blessings.push(reader.int32());
                    else
                        message.blessings.push(reader.int32());
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
        /* repeated proto.Blessings blessings = 1; */
        if (message.blessings.length) {
            writer.tag(1, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.blessings.length; i++)
                writer.int32(message.blessings[i]);
            writer.join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.BlessingsAssignment
 */
export const BlessingsAssignment = new BlessingsAssignment$Type();
// @generated message type with reflection information, may provide speed optimized methods
class BlessingsAssignments$Type extends MessageType {
    constructor() {
        super("proto.BlessingsAssignments", [
            { no: 1, name: "paladins", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => BlessingsAssignment }
        ]);
    }
    create(value) {
        const message = { paladins: [] };
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
                case /* repeated proto.BlessingsAssignment paladins */ 1:
                    message.paladins.push(BlessingsAssignment.internalBinaryRead(reader, reader.uint32(), options));
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
        /* repeated proto.BlessingsAssignment paladins = 1; */
        for (let i = 0; i < message.paladins.length; i++)
            BlessingsAssignment.internalBinaryWrite(message.paladins[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.BlessingsAssignments
 */
export const BlessingsAssignments = new BlessingsAssignments$Type();
