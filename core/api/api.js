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
import { Race } from './common.js';
// @generated message type with reflection information, may provide speed optimized methods
class PlayerOptions$Type extends MessageType {
    constructor() {
        super("api.PlayerOptions", [
            { no: 1, name: "race", kind: "enum", T: () => ["api.Race", Race] },
            { no: 2, name: "balance_druid", kind: "message", oneof: "spec", T: () => BalanceDruid },
            { no: 3, name: "hunter", kind: "message", oneof: "spec", T: () => Hunter },
            { no: 4, name: "mage", kind: "message", oneof: "spec", T: () => Mage },
            { no: 5, name: "paladin", kind: "message", oneof: "spec", T: () => Paladin },
            { no: 6, name: "priest", kind: "message", oneof: "spec", T: () => Priest },
            { no: 7, name: "rogue", kind: "message", oneof: "spec", T: () => Rogue },
            { no: 8, name: "elemental_shaman", kind: "message", oneof: "spec", T: () => ElementalShaman },
            { no: 9, name: "warlock", kind: "message", oneof: "spec", T: () => Warlock },
            { no: 10, name: "warrior", kind: "message", oneof: "spec", T: () => Warrior },
            { no: 11, name: "consumes", kind: "message", T: () => Consumes }
        ]);
    }
    create(value) {
        const message = { race: 0, spec: { oneofKind: undefined } };
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
                case /* api.Race race */ 1:
                    message.race = reader.int32();
                    break;
                case /* api.BalanceDruid balance_druid */ 2:
                    message.spec = {
                        oneofKind: "balanceDruid",
                        balanceDruid: BalanceDruid.internalBinaryRead(reader, reader.uint32(), options, message.spec.balanceDruid)
                    };
                    break;
                case /* api.Hunter hunter */ 3:
                    message.spec = {
                        oneofKind: "hunter",
                        hunter: Hunter.internalBinaryRead(reader, reader.uint32(), options, message.spec.hunter)
                    };
                    break;
                case /* api.Mage mage */ 4:
                    message.spec = {
                        oneofKind: "mage",
                        mage: Mage.internalBinaryRead(reader, reader.uint32(), options, message.spec.mage)
                    };
                    break;
                case /* api.Paladin paladin */ 5:
                    message.spec = {
                        oneofKind: "paladin",
                        paladin: Paladin.internalBinaryRead(reader, reader.uint32(), options, message.spec.paladin)
                    };
                    break;
                case /* api.Priest priest */ 6:
                    message.spec = {
                        oneofKind: "priest",
                        priest: Priest.internalBinaryRead(reader, reader.uint32(), options, message.spec.priest)
                    };
                    break;
                case /* api.Rogue rogue */ 7:
                    message.spec = {
                        oneofKind: "rogue",
                        rogue: Rogue.internalBinaryRead(reader, reader.uint32(), options, message.spec.rogue)
                    };
                    break;
                case /* api.ElementalShaman elemental_shaman */ 8:
                    message.spec = {
                        oneofKind: "elementalShaman",
                        elementalShaman: ElementalShaman.internalBinaryRead(reader, reader.uint32(), options, message.spec.elementalShaman)
                    };
                    break;
                case /* api.Warlock warlock */ 9:
                    message.spec = {
                        oneofKind: "warlock",
                        warlock: Warlock.internalBinaryRead(reader, reader.uint32(), options, message.spec.warlock)
                    };
                    break;
                case /* api.Warrior warrior */ 10:
                    message.spec = {
                        oneofKind: "warrior",
                        warrior: Warrior.internalBinaryRead(reader, reader.uint32(), options, message.spec.warrior)
                    };
                    break;
                case /* api.Consumes consumes */ 11:
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
        /* api.Race race = 1; */
        if (message.race !== 0)
            writer.tag(1, WireType.Varint).int32(message.race);
        /* api.BalanceDruid balance_druid = 2; */
        if (message.spec.oneofKind === "balanceDruid")
            BalanceDruid.internalBinaryWrite(message.spec.balanceDruid, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* api.Hunter hunter = 3; */
        if (message.spec.oneofKind === "hunter")
            Hunter.internalBinaryWrite(message.spec.hunter, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* api.Mage mage = 4; */
        if (message.spec.oneofKind === "mage")
            Mage.internalBinaryWrite(message.spec.mage, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* api.Paladin paladin = 5; */
        if (message.spec.oneofKind === "paladin")
            Paladin.internalBinaryWrite(message.spec.paladin, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* api.Priest priest = 6; */
        if (message.spec.oneofKind === "priest")
            Priest.internalBinaryWrite(message.spec.priest, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        /* api.Rogue rogue = 7; */
        if (message.spec.oneofKind === "rogue")
            Rogue.internalBinaryWrite(message.spec.rogue, writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        /* api.ElementalShaman elemental_shaman = 8; */
        if (message.spec.oneofKind === "elementalShaman")
            ElementalShaman.internalBinaryWrite(message.spec.elementalShaman, writer.tag(8, WireType.LengthDelimited).fork(), options).join();
        /* api.Warlock warlock = 9; */
        if (message.spec.oneofKind === "warlock")
            Warlock.internalBinaryWrite(message.spec.warlock, writer.tag(9, WireType.LengthDelimited).fork(), options).join();
        /* api.Warrior warrior = 10; */
        if (message.spec.oneofKind === "warrior")
            Warrior.internalBinaryWrite(message.spec.warrior, writer.tag(10, WireType.LengthDelimited).fork(), options).join();
        /* api.Consumes consumes = 11; */
        if (message.consumes)
            Consumes.internalBinaryWrite(message.consumes, writer.tag(11, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message api.PlayerOptions
 */
export const PlayerOptions = new PlayerOptions$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Hunter$Type extends MessageType {
    constructor() {
        super("api.Hunter", []);
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
 * @generated MessageType for protobuf message api.Hunter
 */
export const Hunter = new Hunter$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Mage$Type extends MessageType {
    constructor() {
        super("api.Mage", []);
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
 * @generated MessageType for protobuf message api.Mage
 */
export const Mage = new Mage$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Paladin$Type extends MessageType {
    constructor() {
        super("api.Paladin", []);
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
 * @generated MessageType for protobuf message api.Paladin
 */
export const Paladin = new Paladin$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Priest$Type extends MessageType {
    constructor() {
        super("api.Priest", []);
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
 * @generated MessageType for protobuf message api.Priest
 */
export const Priest = new Priest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Rogue$Type extends MessageType {
    constructor() {
        super("api.Rogue", []);
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
 * @generated MessageType for protobuf message api.Rogue
 */
export const Rogue = new Rogue$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Warlock$Type extends MessageType {
    constructor() {
        super("api.Warlock", []);
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
 * @generated MessageType for protobuf message api.Warlock
 */
export const Warlock = new Warlock$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Warrior$Type extends MessageType {
    constructor() {
        super("api.Warrior", []);
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
 * @generated MessageType for protobuf message api.Warrior
 */
export const Warrior = new Warrior$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Player$Type extends MessageType {
    constructor() {
        super("api.Player", [
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
                case /* api.PlayerOptions options */ 1:
                    message.options = PlayerOptions.internalBinaryRead(reader, reader.uint32(), options, message.options);
                    break;
                case /* api.EquipmentSpec equipment */ 2:
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
        /* api.PlayerOptions options = 1; */
        if (message.options)
            PlayerOptions.internalBinaryWrite(message.options, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* api.EquipmentSpec equipment = 2; */
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
 * @generated MessageType for protobuf message api.Player
 */
export const Player = new Player$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Party$Type extends MessageType {
    constructor() {
        super("api.Party", [
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
                case /* repeated api.Player players */ 1:
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
        /* repeated api.Player players = 1; */
        for (let i = 0; i < message.players.length; i++)
            Player.internalBinaryWrite(message.players[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message api.Party
 */
export const Party = new Party$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Raid$Type extends MessageType {
    constructor() {
        super("api.Raid", [
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
                case /* repeated api.Party parties */ 1:
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
        /* repeated api.Party parties = 1; */
        for (let i = 0; i < message.parties.length; i++)
            Party.internalBinaryWrite(message.parties[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message api.Raid
 */
export const Raid = new Raid$Type();
// @generated message type with reflection information, may provide speed optimized methods
class IndividualSimRequest$Type extends MessageType {
    constructor() {
        super("api.IndividualSimRequest", [
            { no: 1, name: "player", kind: "message", T: () => Player },
            { no: 2, name: "buffs", kind: "message", T: () => Buffs },
            { no: 3, name: "encounter", kind: "message", T: () => Encounter },
            { no: 4, name: "iterations", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "random_seed", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 6, name: "gcd_min", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 7, name: "debug", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 8, name: "exit_on_oom", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { iterations: 0, randomSeed: 0n, gcdMin: 0, debug: false, exitOnOom: false };
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
                case /* api.Player player */ 1:
                    message.player = Player.internalBinaryRead(reader, reader.uint32(), options, message.player);
                    break;
                case /* api.Buffs buffs */ 2:
                    message.buffs = Buffs.internalBinaryRead(reader, reader.uint32(), options, message.buffs);
                    break;
                case /* api.Encounter encounter */ 3:
                    message.encounter = Encounter.internalBinaryRead(reader, reader.uint32(), options, message.encounter);
                    break;
                case /* int32 iterations */ 4:
                    message.iterations = reader.int32();
                    break;
                case /* int64 random_seed */ 5:
                    message.randomSeed = reader.int64().toBigInt();
                    break;
                case /* double gcd_min */ 6:
                    message.gcdMin = reader.double();
                    break;
                case /* bool debug */ 7:
                    message.debug = reader.bool();
                    break;
                case /* bool exit_on_oom */ 8:
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
        /* api.Player player = 1; */
        if (message.player)
            Player.internalBinaryWrite(message.player, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* api.Buffs buffs = 2; */
        if (message.buffs)
            Buffs.internalBinaryWrite(message.buffs, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* api.Encounter encounter = 3; */
        if (message.encounter)
            Encounter.internalBinaryWrite(message.encounter, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* int32 iterations = 4; */
        if (message.iterations !== 0)
            writer.tag(4, WireType.Varint).int32(message.iterations);
        /* int64 random_seed = 5; */
        if (message.randomSeed !== 0n)
            writer.tag(5, WireType.Varint).int64(message.randomSeed);
        /* double gcd_min = 6; */
        if (message.gcdMin !== 0)
            writer.tag(6, WireType.Bit64).double(message.gcdMin);
        /* bool debug = 7; */
        if (message.debug !== false)
            writer.tag(7, WireType.Varint).bool(message.debug);
        /* bool exit_on_oom = 8; */
        if (message.exitOnOom !== false)
            writer.tag(8, WireType.Varint).bool(message.exitOnOom);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message api.IndividualSimRequest
 */
export const IndividualSimRequest = new IndividualSimRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class IndividualSimResult$Type extends MessageType {
    constructor() {
        super("api.IndividualSimResult", [
            { no: 1, name: "execution_duration_ms", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "logs", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "dps_avg", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 4, name: "dps_stdev", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 5, name: "dps_max", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 6, name: "dps_hist", kind: "map", K: 5 /*ScalarType.INT32*/, V: { kind: "scalar", T: 5 /*ScalarType.INT32*/ } },
            { no: 7, name: "num_oom", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 8, name: "oom_at_avg", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 9, name: "dps_at_oom_avg", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 10, name: "casts", kind: "map", K: 5 /*ScalarType.INT32*/, V: { kind: "message", T: () => CastMetric } },
            { no: 11, name: "error", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value) {
        const message = { executionDurationMs: 0n, logs: "", dpsAvg: 0, dpsStdev: 0, dpsMax: 0, dpsHist: {}, numOom: 0, oomAtAvg: 0, dpsAtOomAvg: 0, casts: {}, error: "" };
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
                case /* map<int32, api.CastMetric> casts */ 10:
                    this.binaryReadMap10(message.casts, reader, options);
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
                default: throw new globalThis.Error("unknown map entry field for field api.IndividualSimResult.dps_hist");
            }
        }
        map[key ?? 0] = val ?? 0;
    }
    binaryReadMap10(map, reader, options) {
        let len = reader.uint32(), end = reader.pos + len, key, val;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case 1:
                    key = reader.int32();
                    break;
                case 2:
                    val = CastMetric.internalBinaryRead(reader, reader.uint32(), options);
                    break;
                default: throw new globalThis.Error("unknown map entry field for field api.IndividualSimResult.casts");
            }
        }
        map[key ?? 0] = val ?? CastMetric.create();
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
        /* map<int32, api.CastMetric> casts = 10; */
        for (let k of Object.keys(message.casts)) {
            writer.tag(10, WireType.LengthDelimited).fork().tag(1, WireType.Varint).int32(parseInt(k));
            writer.tag(2, WireType.LengthDelimited).fork();
            CastMetric.internalBinaryWrite(message.casts[k], writer, options);
            writer.join().join();
        }
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
 * @generated MessageType for protobuf message api.IndividualSimResult
 */
export const IndividualSimResult = new IndividualSimResult$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CastMetric$Type extends MessageType {
    constructor() {
        super("api.CastMetric", [
            { no: 1, name: "counts", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "dmgs", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 3, name: "tags", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value) {
        const message = { counts: [], dmgs: [], tags: [] };
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
                case /* repeated int32 counts */ 1:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.counts.push(reader.int32());
                    else
                        message.counts.push(reader.int32());
                    break;
                case /* repeated double dmgs */ 2:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.dmgs.push(reader.double());
                    else
                        message.dmgs.push(reader.double());
                    break;
                case /* repeated int32 tags */ 3:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.tags.push(reader.int32());
                    else
                        message.tags.push(reader.int32());
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
        /* repeated int32 counts = 1; */
        if (message.counts.length) {
            writer.tag(1, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.counts.length; i++)
                writer.int32(message.counts[i]);
            writer.join();
        }
        /* repeated double dmgs = 2; */
        if (message.dmgs.length) {
            writer.tag(2, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.dmgs.length; i++)
                writer.double(message.dmgs[i]);
            writer.join();
        }
        /* repeated int32 tags = 3; */
        if (message.tags.length) {
            writer.tag(3, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.tags.length; i++)
                writer.int32(message.tags[i]);
            writer.join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message api.CastMetric
 */
export const CastMetric = new CastMetric$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RaidSimRequest$Type extends MessageType {
    constructor() {
        super("api.RaidSimRequest", [
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
                case /* api.Raid raid */ 1:
                    message.raid = Raid.internalBinaryRead(reader, reader.uint32(), options, message.raid);
                    break;
                case /* api.Encounter encounter */ 2:
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
        /* api.Raid raid = 1; */
        if (message.raid)
            Raid.internalBinaryWrite(message.raid, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* api.Encounter encounter = 2; */
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
 * @generated MessageType for protobuf message api.RaidSimRequest
 */
export const RaidSimRequest = new RaidSimRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GearListRequest$Type extends MessageType {
    constructor() {
        super("api.GearListRequest", [
            { no: 1, name: "spec", kind: "enum", T: () => ["api.Spec", Spec] }
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
                case /* api.Spec spec */ 1:
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
        /* api.Spec spec = 1; */
        if (message.spec !== 0)
            writer.tag(1, WireType.Varint).int32(message.spec);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message api.GearListRequest
 */
export const GearListRequest = new GearListRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GearListResult$Type extends MessageType {
    constructor() {
        super("api.GearListResult", [
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
                case /* repeated api.Item items */ 1:
                    message.items.push(Item.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* repeated api.Enchant enchants */ 2:
                    message.enchants.push(Enchant.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* repeated api.Gem gems */ 3:
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
        /* repeated api.Item items = 1; */
        for (let i = 0; i < message.items.length; i++)
            Item.internalBinaryWrite(message.items[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated api.Enchant enchants = 2; */
        for (let i = 0; i < message.enchants.length; i++)
            Enchant.internalBinaryWrite(message.enchants[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* repeated api.Gem gems = 3; */
        for (let i = 0; i < message.gems.length; i++)
            Gem.internalBinaryWrite(message.gems[i], writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message api.GearListResult
 */
export const GearListResult = new GearListResult$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ComputeStatsRequest$Type extends MessageType {
    constructor() {
        super("api.ComputeStatsRequest", [
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
                case /* api.Player player */ 1:
                    message.player = Player.internalBinaryRead(reader, reader.uint32(), options, message.player);
                    break;
                case /* api.Buffs buffs */ 2:
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
        /* api.Player player = 1; */
        if (message.player)
            Player.internalBinaryWrite(message.player, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* api.Buffs buffs = 2; */
        if (message.buffs)
            Buffs.internalBinaryWrite(message.buffs, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message api.ComputeStatsRequest
 */
export const ComputeStatsRequest = new ComputeStatsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ComputeStatsResult$Type extends MessageType {
    constructor() {
        super("api.ComputeStatsResult", [
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
 * @generated MessageType for protobuf message api.ComputeStatsResult
 */
export const ComputeStatsResult = new ComputeStatsResult$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StatWeightsRequest$Type extends MessageType {
    constructor() {
        super("api.StatWeightsRequest", [
            { no: 1, name: "options", kind: "message", T: () => IndividualSimRequest },
            { no: 2, name: "stats_to_weigh", kind: "enum", repeat: 1 /*RepeatType.PACKED*/, T: () => ["api.Stat", Stat] },
            { no: 3, name: "ep_reference_stat", kind: "enum", T: () => ["api.Stat", Stat] }
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
                case /* api.IndividualSimRequest options */ 1:
                    message.options = IndividualSimRequest.internalBinaryRead(reader, reader.uint32(), options, message.options);
                    break;
                case /* repeated api.Stat stats_to_weigh */ 2:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.statsToWeigh.push(reader.int32());
                    else
                        message.statsToWeigh.push(reader.int32());
                    break;
                case /* api.Stat ep_reference_stat */ 3:
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
        /* api.IndividualSimRequest options = 1; */
        if (message.options)
            IndividualSimRequest.internalBinaryWrite(message.options, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated api.Stat stats_to_weigh = 2; */
        if (message.statsToWeigh.length) {
            writer.tag(2, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.statsToWeigh.length; i++)
                writer.int32(message.statsToWeigh[i]);
            writer.join();
        }
        /* api.Stat ep_reference_stat = 3; */
        if (message.epReferenceStat !== 0)
            writer.tag(3, WireType.Varint).int32(message.epReferenceStat);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message api.StatWeightsRequest
 */
export const StatWeightsRequest = new StatWeightsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StatWeightsResult$Type extends MessageType {
    constructor() {
        super("api.StatWeightsResult", [
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
 * @generated MessageType for protobuf message api.StatWeightsResult
 */
export const StatWeightsResult = new StatWeightsResult$Type();
