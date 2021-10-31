import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
import { Stat } from './common.js';
import { Gem } from './common.js';
import { Enchant } from './common.js';
import { Item } from './common.js';
import { Spec } from './common.js';
import { Encounter } from './common.js';
import { Buffs } from './common.js';
import { EquipmentSpec } from './common.js';
import { Consumes } from './common.js';
import { ElementalShaman } from './shaman.js';
import { BalanceDruid } from './druid.js';
import { Class } from './common.js';
import { Race } from './common.js';
/**
 * ID for actions that aren't spells or items.
 *
 * @generated from protobuf enum proto.OtherAction
 */
export var OtherAction;
(function (OtherAction) {
    /**
     * @generated from protobuf enum value: OtherActionNone = 0;
     */
    OtherAction[OtherAction["OtherActionNone"] = 0] = "OtherActionNone";
    /**
     * @generated from protobuf enum value: OtherActionWait = 1;
     */
    OtherAction[OtherAction["OtherActionWait"] = 1] = "OtherActionWait";
})(OtherAction || (OtherAction = {}));
// @generated message type with reflection information, may provide speed optimized methods
class PlayerOptions$Type extends MessageType {
    constructor() {
        super("proto.PlayerOptions", [
            { no: 1, name: "race", kind: "enum", T: () => ["proto.Race", Race] },
            { no: 2, name: "class", kind: "enum", T: () => ["proto.Class", Class] },
            { no: 3, name: "balance_druid", kind: "message", oneof: "spec", T: () => BalanceDruid },
            { no: 4, name: "hunter", kind: "message", oneof: "spec", T: () => Hunter },
            { no: 5, name: "mage", kind: "message", oneof: "spec", T: () => Mage },
            { no: 6, name: "paladin", kind: "message", oneof: "spec", T: () => Paladin },
            { no: 7, name: "priest", kind: "message", oneof: "spec", T: () => Priest },
            { no: 8, name: "rogue", kind: "message", oneof: "spec", T: () => Rogue },
            { no: 9, name: "elemental_shaman", kind: "message", oneof: "spec", T: () => ElementalShaman },
            { no: 10, name: "warlock", kind: "message", oneof: "spec", T: () => Warlock },
            { no: 11, name: "warrior", kind: "message", oneof: "spec", T: () => Warrior },
            { no: 12, name: "consumes", kind: "message", T: () => Consumes }
        ]);
    }
    create(value) {
        const message = { race: 0, class: 0, spec: { oneofKind: undefined } };
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
                case /* proto.Race race */ 1:
                    message.race = reader.int32();
                    break;
                case /* proto.Class class */ 2:
                    message.class = reader.int32();
                    break;
                case /* proto.BalanceDruid balance_druid */ 3:
                    message.spec = {
                        oneofKind: "balanceDruid",
                        balanceDruid: BalanceDruid.internalBinaryRead(reader, reader.uint32(), options, message.spec.balanceDruid)
                    };
                    break;
                case /* proto.Hunter hunter */ 4:
                    message.spec = {
                        oneofKind: "hunter",
                        hunter: Hunter.internalBinaryRead(reader, reader.uint32(), options, message.spec.hunter)
                    };
                    break;
                case /* proto.Mage mage */ 5:
                    message.spec = {
                        oneofKind: "mage",
                        mage: Mage.internalBinaryRead(reader, reader.uint32(), options, message.spec.mage)
                    };
                    break;
                case /* proto.Paladin paladin */ 6:
                    message.spec = {
                        oneofKind: "paladin",
                        paladin: Paladin.internalBinaryRead(reader, reader.uint32(), options, message.spec.paladin)
                    };
                    break;
                case /* proto.Priest priest */ 7:
                    message.spec = {
                        oneofKind: "priest",
                        priest: Priest.internalBinaryRead(reader, reader.uint32(), options, message.spec.priest)
                    };
                    break;
                case /* proto.Rogue rogue */ 8:
                    message.spec = {
                        oneofKind: "rogue",
                        rogue: Rogue.internalBinaryRead(reader, reader.uint32(), options, message.spec.rogue)
                    };
                    break;
                case /* proto.ElementalShaman elemental_shaman */ 9:
                    message.spec = {
                        oneofKind: "elementalShaman",
                        elementalShaman: ElementalShaman.internalBinaryRead(reader, reader.uint32(), options, message.spec.elementalShaman)
                    };
                    break;
                case /* proto.Warlock warlock */ 10:
                    message.spec = {
                        oneofKind: "warlock",
                        warlock: Warlock.internalBinaryRead(reader, reader.uint32(), options, message.spec.warlock)
                    };
                    break;
                case /* proto.Warrior warrior */ 11:
                    message.spec = {
                        oneofKind: "warrior",
                        warrior: Warrior.internalBinaryRead(reader, reader.uint32(), options, message.spec.warrior)
                    };
                    break;
                case /* proto.Consumes consumes */ 12:
                    message.consumes = Consumes.internalBinaryRead(reader, reader.uint32(), options, message.consumes);
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
        /* proto.Race race = 1; */
        if (message.race !== 0)
            writer.tag(1, WireType.Varint).int32(message.race);
        /* proto.Class class = 2; */
        if (message.class !== 0)
            writer.tag(2, WireType.Varint).int32(message.class);
        /* proto.BalanceDruid balance_druid = 3; */
        if (message.spec.oneofKind === "balanceDruid")
            BalanceDruid.internalBinaryWrite(message.spec.balanceDruid, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* proto.Hunter hunter = 4; */
        if (message.spec.oneofKind === "hunter")
            Hunter.internalBinaryWrite(message.spec.hunter, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* proto.Mage mage = 5; */
        if (message.spec.oneofKind === "mage")
            Mage.internalBinaryWrite(message.spec.mage, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* proto.Paladin paladin = 6; */
        if (message.spec.oneofKind === "paladin")
            Paladin.internalBinaryWrite(message.spec.paladin, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        /* proto.Priest priest = 7; */
        if (message.spec.oneofKind === "priest")
            Priest.internalBinaryWrite(message.spec.priest, writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        /* proto.Rogue rogue = 8; */
        if (message.spec.oneofKind === "rogue")
            Rogue.internalBinaryWrite(message.spec.rogue, writer.tag(8, WireType.LengthDelimited).fork(), options).join();
        /* proto.ElementalShaman elemental_shaman = 9; */
        if (message.spec.oneofKind === "elementalShaman")
            ElementalShaman.internalBinaryWrite(message.spec.elementalShaman, writer.tag(9, WireType.LengthDelimited).fork(), options).join();
        /* proto.Warlock warlock = 10; */
        if (message.spec.oneofKind === "warlock")
            Warlock.internalBinaryWrite(message.spec.warlock, writer.tag(10, WireType.LengthDelimited).fork(), options).join();
        /* proto.Warrior warrior = 11; */
        if (message.spec.oneofKind === "warrior")
            Warrior.internalBinaryWrite(message.spec.warrior, writer.tag(11, WireType.LengthDelimited).fork(), options).join();
        /* proto.Consumes consumes = 12; */
        if (message.consumes)
            Consumes.internalBinaryWrite(message.consumes, writer.tag(12, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.PlayerOptions
 */
export const PlayerOptions = new PlayerOptions$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Hunter$Type extends MessageType {
    constructor() {
        super("proto.Hunter", []);
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
 * @generated MessageType for protobuf message proto.Hunter
 */
export const Hunter = new Hunter$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Mage$Type extends MessageType {
    constructor() {
        super("proto.Mage", []);
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
 * @generated MessageType for protobuf message proto.Mage
 */
export const Mage = new Mage$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Paladin$Type extends MessageType {
    constructor() {
        super("proto.Paladin", []);
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
 * @generated MessageType for protobuf message proto.Paladin
 */
export const Paladin = new Paladin$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Priest$Type extends MessageType {
    constructor() {
        super("proto.Priest", []);
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
 * @generated MessageType for protobuf message proto.Priest
 */
export const Priest = new Priest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Rogue$Type extends MessageType {
    constructor() {
        super("proto.Rogue", []);
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
 * @generated MessageType for protobuf message proto.Rogue
 */
export const Rogue = new Rogue$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Warlock$Type extends MessageType {
    constructor() {
        super("proto.Warlock", []);
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
 * @generated MessageType for protobuf message proto.Warlock
 */
export const Warlock = new Warlock$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Warrior$Type extends MessageType {
    constructor() {
        super("proto.Warrior", []);
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
 * @generated MessageType for protobuf message proto.Warrior
 */
export const Warrior = new Warrior$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Player$Type extends MessageType {
    constructor() {
        super("proto.Player", [
            { no: 1, name: "options", kind: "message", T: () => PlayerOptions },
            { no: 2, name: "equipment", kind: "message", T: () => EquipmentSpec },
            { no: 3, name: "custom_stats", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { customStats: [] };
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
                case /* proto.PlayerOptions options */ 1:
                    message.options = PlayerOptions.internalBinaryRead(reader, reader.uint32(), options, message.options);
                    break;
                case /* proto.EquipmentSpec equipment */ 2:
                    message.equipment = EquipmentSpec.internalBinaryRead(reader, reader.uint32(), options, message.equipment);
                    break;
                case /* repeated double custom_stats */ 3:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.customStats.push(reader.double());
                    else
                        message.customStats.push(reader.double());
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
        /* proto.PlayerOptions options = 1; */
        if (message.options)
            PlayerOptions.internalBinaryWrite(message.options, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.EquipmentSpec equipment = 2; */
        if (message.equipment)
            EquipmentSpec.internalBinaryWrite(message.equipment, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* repeated double custom_stats = 3; */
        if (message.customStats.length) {
            writer.tag(3, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.customStats.length; i++)
                writer.double(message.customStats[i]);
            writer.join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Player
 */
export const Player = new Player$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Party$Type extends MessageType {
    constructor() {
        super("proto.Party", [
            { no: 1, name: "players", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Player }
        ]);
    }
    create(value) {
        const message = { players: [] };
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
                case /* repeated proto.Player players */ 1:
                    message.players.push(Player.internalBinaryRead(reader, reader.uint32(), options));
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
        /* repeated proto.Player players = 1; */
        for (let i = 0; i < message.players.length; i++)
            Player.internalBinaryWrite(message.players[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Party
 */
export const Party = new Party$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Raid$Type extends MessageType {
    constructor() {
        super("proto.Raid", [
            { no: 1, name: "parties", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Party }
        ]);
    }
    create(value) {
        const message = { parties: [] };
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
                case /* repeated proto.Party parties */ 1:
                    message.parties.push(Party.internalBinaryRead(reader, reader.uint32(), options));
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
        /* repeated proto.Party parties = 1; */
        for (let i = 0; i < message.parties.length; i++)
            Party.internalBinaryWrite(message.parties[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Raid
 */
export const Raid = new Raid$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SimOptions$Type extends MessageType {
    constructor() {
        super("proto.SimOptions", [
            { no: 1, name: "iterations", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "random_seed", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 3, name: "debug", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "exit_on_oom", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { iterations: 0, randomSeed: 0n, debug: false, exitOnOom: false };
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
                case /* int32 iterations */ 1:
                    message.iterations = reader.int32();
                    break;
                case /* int64 random_seed */ 2:
                    message.randomSeed = reader.int64().toBigInt();
                    break;
                case /* bool debug */ 3:
                    message.debug = reader.bool();
                    break;
                case /* bool exit_on_oom */ 4:
                    message.exitOnOom = reader.bool();
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
        /* int32 iterations = 1; */
        if (message.iterations !== 0)
            writer.tag(1, WireType.Varint).int32(message.iterations);
        /* int64 random_seed = 2; */
        if (message.randomSeed !== 0n)
            writer.tag(2, WireType.Varint).int64(message.randomSeed);
        /* bool debug = 3; */
        if (message.debug !== false)
            writer.tag(3, WireType.Varint).bool(message.debug);
        /* bool exit_on_oom = 4; */
        if (message.exitOnOom !== false)
            writer.tag(4, WireType.Varint).bool(message.exitOnOom);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.SimOptions
 */
export const SimOptions = new SimOptions$Type();
// @generated message type with reflection information, may provide speed optimized methods
class IndividualSimRequest$Type extends MessageType {
    constructor() {
        super("proto.IndividualSimRequest", [
            { no: 1, name: "player", kind: "message", T: () => Player },
            { no: 2, name: "buffs", kind: "message", T: () => Buffs },
            { no: 3, name: "encounter", kind: "message", T: () => Encounter },
            { no: 4, name: "sim_options", kind: "message", T: () => SimOptions }
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
                case /* proto.Player player */ 1:
                    message.player = Player.internalBinaryRead(reader, reader.uint32(), options, message.player);
                    break;
                case /* proto.Buffs buffs */ 2:
                    message.buffs = Buffs.internalBinaryRead(reader, reader.uint32(), options, message.buffs);
                    break;
                case /* proto.Encounter encounter */ 3:
                    message.encounter = Encounter.internalBinaryRead(reader, reader.uint32(), options, message.encounter);
                    break;
                case /* proto.SimOptions sim_options */ 4:
                    message.simOptions = SimOptions.internalBinaryRead(reader, reader.uint32(), options, message.simOptions);
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
        /* proto.Player player = 1; */
        if (message.player)
            Player.internalBinaryWrite(message.player, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.Buffs buffs = 2; */
        if (message.buffs)
            Buffs.internalBinaryWrite(message.buffs, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.Encounter encounter = 3; */
        if (message.encounter)
            Encounter.internalBinaryWrite(message.encounter, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* proto.SimOptions sim_options = 4; */
        if (message.simOptions)
            SimOptions.internalBinaryWrite(message.simOptions, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.IndividualSimRequest
 */
export const IndividualSimRequest = new IndividualSimRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class IndividualSimResult$Type extends MessageType {
    constructor() {
        super("proto.IndividualSimResult", [
            { no: 1, name: "execution_duration_ms", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "logs", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "dps_avg", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 4, name: "dps_stdev", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 5, name: "dps_max", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 6, name: "dps_hist", kind: "map", K: 5 /*ScalarType.INT32*/, V: { kind: "scalar", T: 5 /*ScalarType.INT32*/ } },
            { no: 7, name: "num_oom", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 8, name: "oom_at_avg", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 9, name: "dps_at_oom_avg", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 10, name: "action_metrics", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => ActionMetric },
            { no: 11, name: "error", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value) {
        const message = { executionDurationMs: 0n, logs: "", dpsAvg: 0, dpsStdev: 0, dpsMax: 0, dpsHist: {}, numOom: 0, oomAtAvg: 0, dpsAtOomAvg: 0, actionMetrics: [], error: "" };
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
                case /* int64 execution_duration_ms */ 1:
                    message.executionDurationMs = reader.int64().toBigInt();
                    break;
                case /* string logs */ 2:
                    message.logs = reader.string();
                    break;
                case /* double dps_avg */ 3:
                    message.dpsAvg = reader.double();
                    break;
                case /* double dps_stdev */ 4:
                    message.dpsStdev = reader.double();
                    break;
                case /* double dps_max */ 5:
                    message.dpsMax = reader.double();
                    break;
                case /* map<int32, int32> dps_hist */ 6:
                    this.binaryReadMap6(message.dpsHist, reader, options);
                    break;
                case /* int32 num_oom */ 7:
                    message.numOom = reader.int32();
                    break;
                case /* double oom_at_avg */ 8:
                    message.oomAtAvg = reader.double();
                    break;
                case /* double dps_at_oom_avg */ 9:
                    message.dpsAtOomAvg = reader.double();
                    break;
                case /* repeated proto.ActionMetric action_metrics */ 10:
                    message.actionMetrics.push(ActionMetric.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* string error */ 11:
                    message.error = reader.string();
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
    binaryReadMap6(map, reader, options) {
        let len = reader.uint32(), end = reader.pos + len, key, val;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case 1:
                    key = reader.int32();
                    break;
                case 2:
                    val = reader.int32();
                    break;
                default: throw new globalThis.Error("unknown map entry field for field proto.IndividualSimResult.dps_hist");
            }
        }
        map[key ?? 0] = val ?? 0;
    }
    internalBinaryWrite(message, writer, options) {
        /* int64 execution_duration_ms = 1; */
        if (message.executionDurationMs !== 0n)
            writer.tag(1, WireType.Varint).int64(message.executionDurationMs);
        /* string logs = 2; */
        if (message.logs !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.logs);
        /* double dps_avg = 3; */
        if (message.dpsAvg !== 0)
            writer.tag(3, WireType.Bit64).double(message.dpsAvg);
        /* double dps_stdev = 4; */
        if (message.dpsStdev !== 0)
            writer.tag(4, WireType.Bit64).double(message.dpsStdev);
        /* double dps_max = 5; */
        if (message.dpsMax !== 0)
            writer.tag(5, WireType.Bit64).double(message.dpsMax);
        /* map<int32, int32> dps_hist = 6; */
        for (let k of Object.keys(message.dpsHist))
            writer.tag(6, WireType.LengthDelimited).fork().tag(1, WireType.Varint).int32(parseInt(k)).tag(2, WireType.Varint).int32(message.dpsHist[k]).join();
        /* int32 num_oom = 7; */
        if (message.numOom !== 0)
            writer.tag(7, WireType.Varint).int32(message.numOom);
        /* double oom_at_avg = 8; */
        if (message.oomAtAvg !== 0)
            writer.tag(8, WireType.Bit64).double(message.oomAtAvg);
        /* double dps_at_oom_avg = 9; */
        if (message.dpsAtOomAvg !== 0)
            writer.tag(9, WireType.Bit64).double(message.dpsAtOomAvg);
        /* repeated proto.ActionMetric action_metrics = 10; */
        for (let i = 0; i < message.actionMetrics.length; i++)
            ActionMetric.internalBinaryWrite(message.actionMetrics[i], writer.tag(10, WireType.LengthDelimited).fork(), options).join();
        /* string error = 11; */
        if (message.error !== "")
            writer.tag(11, WireType.LengthDelimited).string(message.error);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.IndividualSimResult
 */
export const IndividualSimResult = new IndividualSimResult$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ActionMetric$Type extends MessageType {
    constructor() {
        super("proto.ActionMetric", [
            { no: 1, name: "spell_id", kind: "scalar", oneof: "actionId", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "item_id", kind: "scalar", oneof: "actionId", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "other_id", kind: "enum", oneof: "actionId", T: () => ["proto.OtherAction", OtherAction] },
            { no: 4, name: "tag", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "casts", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "hits", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 7, name: "crits", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 8, name: "misses", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 9, name: "damage", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { actionId: { oneofKind: undefined }, tag: 0, casts: 0, hits: 0, crits: 0, misses: 0, damage: 0 };
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
                case /* int32 spell_id */ 1:
                    message.actionId = {
                        oneofKind: "spellId",
                        spellId: reader.int32()
                    };
                    break;
                case /* int32 item_id */ 2:
                    message.actionId = {
                        oneofKind: "itemId",
                        itemId: reader.int32()
                    };
                    break;
                case /* proto.OtherAction other_id */ 3:
                    message.actionId = {
                        oneofKind: "otherId",
                        otherId: reader.int32()
                    };
                    break;
                case /* int32 tag */ 4:
                    message.tag = reader.int32();
                    break;
                case /* int32 casts */ 5:
                    message.casts = reader.int32();
                    break;
                case /* int32 hits */ 6:
                    message.hits = reader.int32();
                    break;
                case /* int32 crits */ 7:
                    message.crits = reader.int32();
                    break;
                case /* int32 misses */ 8:
                    message.misses = reader.int32();
                    break;
                case /* double damage */ 9:
                    message.damage = reader.double();
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
        /* int32 spell_id = 1; */
        if (message.actionId.oneofKind === "spellId")
            writer.tag(1, WireType.Varint).int32(message.actionId.spellId);
        /* int32 item_id = 2; */
        if (message.actionId.oneofKind === "itemId")
            writer.tag(2, WireType.Varint).int32(message.actionId.itemId);
        /* proto.OtherAction other_id = 3; */
        if (message.actionId.oneofKind === "otherId")
            writer.tag(3, WireType.Varint).int32(message.actionId.otherId);
        /* int32 tag = 4; */
        if (message.tag !== 0)
            writer.tag(4, WireType.Varint).int32(message.tag);
        /* int32 casts = 5; */
        if (message.casts !== 0)
            writer.tag(5, WireType.Varint).int32(message.casts);
        /* int32 hits = 6; */
        if (message.hits !== 0)
            writer.tag(6, WireType.Varint).int32(message.hits);
        /* int32 crits = 7; */
        if (message.crits !== 0)
            writer.tag(7, WireType.Varint).int32(message.crits);
        /* int32 misses = 8; */
        if (message.misses !== 0)
            writer.tag(8, WireType.Varint).int32(message.misses);
        /* double damage = 9; */
        if (message.damage !== 0)
            writer.tag(9, WireType.Bit64).double(message.damage);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ActionMetric
 */
export const ActionMetric = new ActionMetric$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RaidSimRequest$Type extends MessageType {
    constructor() {
        super("proto.RaidSimRequest", [
            { no: 1, name: "raid", kind: "message", T: () => Raid },
            { no: 2, name: "encounter", kind: "message", T: () => Encounter },
            { no: 3, name: "random_seed", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 4, name: "gcd_min", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 5, name: "debug", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { randomSeed: 0n, gcdMin: 0, debug: false };
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
                case /* proto.Raid raid */ 1:
                    message.raid = Raid.internalBinaryRead(reader, reader.uint32(), options, message.raid);
                    break;
                case /* proto.Encounter encounter */ 2:
                    message.encounter = Encounter.internalBinaryRead(reader, reader.uint32(), options, message.encounter);
                    break;
                case /* int64 random_seed */ 3:
                    message.randomSeed = reader.int64().toBigInt();
                    break;
                case /* double gcd_min */ 4:
                    message.gcdMin = reader.double();
                    break;
                case /* bool debug */ 5:
                    message.debug = reader.bool();
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
        /* proto.Raid raid = 1; */
        if (message.raid)
            Raid.internalBinaryWrite(message.raid, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.Encounter encounter = 2; */
        if (message.encounter)
            Encounter.internalBinaryWrite(message.encounter, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* int64 random_seed = 3; */
        if (message.randomSeed !== 0n)
            writer.tag(3, WireType.Varint).int64(message.randomSeed);
        /* double gcd_min = 4; */
        if (message.gcdMin !== 0)
            writer.tag(4, WireType.Bit64).double(message.gcdMin);
        /* bool debug = 5; */
        if (message.debug !== false)
            writer.tag(5, WireType.Varint).bool(message.debug);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.RaidSimRequest
 */
export const RaidSimRequest = new RaidSimRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GearListRequest$Type extends MessageType {
    constructor() {
        super("proto.GearListRequest", [
            { no: 1, name: "spec", kind: "enum", T: () => ["proto.Spec", Spec] }
        ]);
    }
    create(value) {
        const message = { spec: 0 };
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
                case /* proto.Spec spec */ 1:
                    message.spec = reader.int32();
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
        /* proto.Spec spec = 1; */
        if (message.spec !== 0)
            writer.tag(1, WireType.Varint).int32(message.spec);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.GearListRequest
 */
export const GearListRequest = new GearListRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GearListResult$Type extends MessageType {
    constructor() {
        super("proto.GearListResult", [
            { no: 1, name: "items", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Item },
            { no: 2, name: "enchants", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Enchant },
            { no: 3, name: "gems", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Gem }
        ]);
    }
    create(value) {
        const message = { items: [], enchants: [], gems: [] };
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
                case /* repeated proto.Item items */ 1:
                    message.items.push(Item.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* repeated proto.Enchant enchants */ 2:
                    message.enchants.push(Enchant.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* repeated proto.Gem gems */ 3:
                    message.gems.push(Gem.internalBinaryRead(reader, reader.uint32(), options));
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
        /* repeated proto.Item items = 1; */
        for (let i = 0; i < message.items.length; i++)
            Item.internalBinaryWrite(message.items[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated proto.Enchant enchants = 2; */
        for (let i = 0; i < message.enchants.length; i++)
            Enchant.internalBinaryWrite(message.enchants[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* repeated proto.Gem gems = 3; */
        for (let i = 0; i < message.gems.length; i++)
            Gem.internalBinaryWrite(message.gems[i], writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.GearListResult
 */
export const GearListResult = new GearListResult$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ComputeStatsRequest$Type extends MessageType {
    constructor() {
        super("proto.ComputeStatsRequest", [
            { no: 1, name: "player", kind: "message", T: () => Player },
            { no: 2, name: "buffs", kind: "message", T: () => Buffs }
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
                case /* proto.Player player */ 1:
                    message.player = Player.internalBinaryRead(reader, reader.uint32(), options, message.player);
                    break;
                case /* proto.Buffs buffs */ 2:
                    message.buffs = Buffs.internalBinaryRead(reader, reader.uint32(), options, message.buffs);
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
        /* proto.Player player = 1; */
        if (message.player)
            Player.internalBinaryWrite(message.player, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.Buffs buffs = 2; */
        if (message.buffs)
            Buffs.internalBinaryWrite(message.buffs, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ComputeStatsRequest
 */
export const ComputeStatsRequest = new ComputeStatsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ComputeStatsResult$Type extends MessageType {
    constructor() {
        super("proto.ComputeStatsResult", [
            { no: 1, name: "gear_only", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 2, name: "finalStats", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 3, name: "sets", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value) {
        const message = { gearOnly: [], finalStats: [], sets: [] };
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
                case /* repeated double gear_only */ 1:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.gearOnly.push(reader.double());
                    else
                        message.gearOnly.push(reader.double());
                    break;
                case /* repeated double finalStats */ 2:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.finalStats.push(reader.double());
                    else
                        message.finalStats.push(reader.double());
                    break;
                case /* repeated string sets */ 3:
                    message.sets.push(reader.string());
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
        /* repeated double gear_only = 1; */
        if (message.gearOnly.length) {
            writer.tag(1, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.gearOnly.length; i++)
                writer.double(message.gearOnly[i]);
            writer.join();
        }
        /* repeated double finalStats = 2; */
        if (message.finalStats.length) {
            writer.tag(2, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.finalStats.length; i++)
                writer.double(message.finalStats[i]);
            writer.join();
        }
        /* repeated string sets = 3; */
        for (let i = 0; i < message.sets.length; i++)
            writer.tag(3, WireType.LengthDelimited).string(message.sets[i]);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ComputeStatsResult
 */
export const ComputeStatsResult = new ComputeStatsResult$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StatWeightsRequest$Type extends MessageType {
    constructor() {
        super("proto.StatWeightsRequest", [
            { no: 1, name: "options", kind: "message", T: () => IndividualSimRequest },
            { no: 2, name: "stats_to_weigh", kind: "enum", repeat: 1 /*RepeatType.PACKED*/, T: () => ["proto.Stat", Stat] },
            { no: 3, name: "ep_reference_stat", kind: "enum", T: () => ["proto.Stat", Stat] }
        ]);
    }
    create(value) {
        const message = { statsToWeigh: [], epReferenceStat: 0 };
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
                case /* proto.IndividualSimRequest options */ 1:
                    message.options = IndividualSimRequest.internalBinaryRead(reader, reader.uint32(), options, message.options);
                    break;
                case /* repeated proto.Stat stats_to_weigh */ 2:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.statsToWeigh.push(reader.int32());
                    else
                        message.statsToWeigh.push(reader.int32());
                    break;
                case /* proto.Stat ep_reference_stat */ 3:
                    message.epReferenceStat = reader.int32();
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
        /* proto.IndividualSimRequest options = 1; */
        if (message.options)
            IndividualSimRequest.internalBinaryWrite(message.options, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated proto.Stat stats_to_weigh = 2; */
        if (message.statsToWeigh.length) {
            writer.tag(2, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.statsToWeigh.length; i++)
                writer.int32(message.statsToWeigh[i]);
            writer.join();
        }
        /* proto.Stat ep_reference_stat = 3; */
        if (message.epReferenceStat !== 0)
            writer.tag(3, WireType.Varint).int32(message.epReferenceStat);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.StatWeightsRequest
 */
export const StatWeightsRequest = new StatWeightsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StatWeightsResult$Type extends MessageType {
    constructor() {
        super("proto.StatWeightsResult", [
            { no: 1, name: "weights", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 2, name: "weights_stdev", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 3, name: "ep_values", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 4, name: "ep_values_stdev", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { weights: [], weightsStdev: [], epValues: [], epValuesStdev: [] };
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
                case /* repeated double weights_stdev */ 2:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.weightsStdev.push(reader.double());
                    else
                        message.weightsStdev.push(reader.double());
                    break;
                case /* repeated double ep_values */ 3:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.epValues.push(reader.double());
                    else
                        message.epValues.push(reader.double());
                    break;
                case /* repeated double ep_values_stdev */ 4:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.epValuesStdev.push(reader.double());
                    else
                        message.epValuesStdev.push(reader.double());
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
        /* repeated double weights_stdev = 2; */
        if (message.weightsStdev.length) {
            writer.tag(2, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.weightsStdev.length; i++)
                writer.double(message.weightsStdev[i]);
            writer.join();
        }
        /* repeated double ep_values = 3; */
        if (message.epValues.length) {
            writer.tag(3, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.epValues.length; i++)
                writer.double(message.epValues[i]);
            writer.join();
        }
        /* repeated double ep_values_stdev = 4; */
        if (message.epValuesStdev.length) {
            writer.tag(4, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.epValuesStdev.length; i++)
                writer.double(message.epValuesStdev[i]);
            writer.join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.StatWeightsResult
 */
export const StatWeightsResult = new StatWeightsResult$Type();
