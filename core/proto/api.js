import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { WireType } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
import { Stat } from './common.js';
import { Gem } from './common.js';
import { Enchant } from './common.js';
import { Item } from './common.js';
import { Encounter } from './common.js';
import { RaidBuffs } from './common.js';
import { PartyBuffs } from './common.js';
import { Warrior } from './warrior.js';
import { Warlock } from './warlock.js';
import { EnhancementShaman } from './shaman.js';
import { ElementalShaman } from './shaman.js';
import { Rogue } from './rogue.js';
import { ShadowPriest } from './priest.js';
import { RetributionPaladin } from './paladin.js';
import { Mage } from './mage.js';
import { Hunter } from './hunter.js';
import { BalanceDruid } from './druid.js';
import { IndividualBuffs } from './common.js';
import { Consumes } from './common.js';
import { EquipmentSpec } from './common.js';
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
class Player$Type extends MessageType {
    constructor() {
        super("proto.Player", [
            { no: 16, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 1, name: "race", kind: "enum", T: () => ["proto.Race", Race] },
            { no: 2, name: "class", kind: "enum", T: () => ["proto.Class", Class] },
            { no: 3, name: "equipment", kind: "message", T: () => EquipmentSpec },
            { no: 4, name: "consumes", kind: "message", T: () => Consumes },
            { no: 5, name: "bonus_stats", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 15, name: "buffs", kind: "message", T: () => IndividualBuffs },
            { no: 6, name: "balance_druid", kind: "message", oneof: "spec", T: () => BalanceDruid },
            { no: 7, name: "hunter", kind: "message", oneof: "spec", T: () => Hunter },
            { no: 8, name: "mage", kind: "message", oneof: "spec", T: () => Mage },
            { no: 9, name: "retribution_paladin", kind: "message", oneof: "spec", T: () => RetributionPaladin },
            { no: 10, name: "shadow_priest", kind: "message", oneof: "spec", T: () => ShadowPriest },
            { no: 11, name: "rogue", kind: "message", oneof: "spec", T: () => Rogue },
            { no: 12, name: "elemental_shaman", kind: "message", oneof: "spec", T: () => ElementalShaman },
            { no: 18, name: "enhancement_shaman", kind: "message", oneof: "spec", T: () => EnhancementShaman },
            { no: 13, name: "warlock", kind: "message", oneof: "spec", T: () => Warlock },
            { no: 14, name: "warrior", kind: "message", oneof: "spec", T: () => Warrior },
            { no: 17, name: "talentsString", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value) {
        const message = { name: "", race: 0, class: 0, bonusStats: [], spec: { oneofKind: undefined }, talentsString: "" };
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
                case /* string name */ 16:
                    message.name = reader.string();
                    break;
                case /* proto.Race race */ 1:
                    message.race = reader.int32();
                    break;
                case /* proto.Class class */ 2:
                    message.class = reader.int32();
                    break;
                case /* proto.EquipmentSpec equipment */ 3:
                    message.equipment = EquipmentSpec.internalBinaryRead(reader, reader.uint32(), options, message.equipment);
                    break;
                case /* proto.Consumes consumes */ 4:
                    message.consumes = Consumes.internalBinaryRead(reader, reader.uint32(), options, message.consumes);
                    break;
                case /* repeated double bonus_stats */ 5:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.bonusStats.push(reader.double());
                    else
                        message.bonusStats.push(reader.double());
                    break;
                case /* proto.IndividualBuffs buffs */ 15:
                    message.buffs = IndividualBuffs.internalBinaryRead(reader, reader.uint32(), options, message.buffs);
                    break;
                case /* proto.BalanceDruid balance_druid */ 6:
                    message.spec = {
                        oneofKind: "balanceDruid",
                        balanceDruid: BalanceDruid.internalBinaryRead(reader, reader.uint32(), options, message.spec.balanceDruid)
                    };
                    break;
                case /* proto.Hunter hunter */ 7:
                    message.spec = {
                        oneofKind: "hunter",
                        hunter: Hunter.internalBinaryRead(reader, reader.uint32(), options, message.spec.hunter)
                    };
                    break;
                case /* proto.Mage mage */ 8:
                    message.spec = {
                        oneofKind: "mage",
                        mage: Mage.internalBinaryRead(reader, reader.uint32(), options, message.spec.mage)
                    };
                    break;
                case /* proto.RetributionPaladin retribution_paladin */ 9:
                    message.spec = {
                        oneofKind: "retributionPaladin",
                        retributionPaladin: RetributionPaladin.internalBinaryRead(reader, reader.uint32(), options, message.spec.retributionPaladin)
                    };
                    break;
                case /* proto.ShadowPriest shadow_priest */ 10:
                    message.spec = {
                        oneofKind: "shadowPriest",
                        shadowPriest: ShadowPriest.internalBinaryRead(reader, reader.uint32(), options, message.spec.shadowPriest)
                    };
                    break;
                case /* proto.Rogue rogue */ 11:
                    message.spec = {
                        oneofKind: "rogue",
                        rogue: Rogue.internalBinaryRead(reader, reader.uint32(), options, message.spec.rogue)
                    };
                    break;
                case /* proto.ElementalShaman elemental_shaman */ 12:
                    message.spec = {
                        oneofKind: "elementalShaman",
                        elementalShaman: ElementalShaman.internalBinaryRead(reader, reader.uint32(), options, message.spec.elementalShaman)
                    };
                    break;
                case /* proto.EnhancementShaman enhancement_shaman */ 18:
                    message.spec = {
                        oneofKind: "enhancementShaman",
                        enhancementShaman: EnhancementShaman.internalBinaryRead(reader, reader.uint32(), options, message.spec.enhancementShaman)
                    };
                    break;
                case /* proto.Warlock warlock */ 13:
                    message.spec = {
                        oneofKind: "warlock",
                        warlock: Warlock.internalBinaryRead(reader, reader.uint32(), options, message.spec.warlock)
                    };
                    break;
                case /* proto.Warrior warrior */ 14:
                    message.spec = {
                        oneofKind: "warrior",
                        warrior: Warrior.internalBinaryRead(reader, reader.uint32(), options, message.spec.warrior)
                    };
                    break;
                case /* string talentsString */ 17:
                    message.talentsString = reader.string();
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
        /* string name = 16; */
        if (message.name !== "")
            writer.tag(16, WireType.LengthDelimited).string(message.name);
        /* proto.Race race = 1; */
        if (message.race !== 0)
            writer.tag(1, WireType.Varint).int32(message.race);
        /* proto.Class class = 2; */
        if (message.class !== 0)
            writer.tag(2, WireType.Varint).int32(message.class);
        /* proto.EquipmentSpec equipment = 3; */
        if (message.equipment)
            EquipmentSpec.internalBinaryWrite(message.equipment, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* proto.Consumes consumes = 4; */
        if (message.consumes)
            Consumes.internalBinaryWrite(message.consumes, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* repeated double bonus_stats = 5; */
        if (message.bonusStats.length) {
            writer.tag(5, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.bonusStats.length; i++)
                writer.double(message.bonusStats[i]);
            writer.join();
        }
        /* proto.IndividualBuffs buffs = 15; */
        if (message.buffs)
            IndividualBuffs.internalBinaryWrite(message.buffs, writer.tag(15, WireType.LengthDelimited).fork(), options).join();
        /* proto.BalanceDruid balance_druid = 6; */
        if (message.spec.oneofKind === "balanceDruid")
            BalanceDruid.internalBinaryWrite(message.spec.balanceDruid, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        /* proto.Hunter hunter = 7; */
        if (message.spec.oneofKind === "hunter")
            Hunter.internalBinaryWrite(message.spec.hunter, writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        /* proto.Mage mage = 8; */
        if (message.spec.oneofKind === "mage")
            Mage.internalBinaryWrite(message.spec.mage, writer.tag(8, WireType.LengthDelimited).fork(), options).join();
        /* proto.RetributionPaladin retribution_paladin = 9; */
        if (message.spec.oneofKind === "retributionPaladin")
            RetributionPaladin.internalBinaryWrite(message.spec.retributionPaladin, writer.tag(9, WireType.LengthDelimited).fork(), options).join();
        /* proto.ShadowPriest shadow_priest = 10; */
        if (message.spec.oneofKind === "shadowPriest")
            ShadowPriest.internalBinaryWrite(message.spec.shadowPriest, writer.tag(10, WireType.LengthDelimited).fork(), options).join();
        /* proto.Rogue rogue = 11; */
        if (message.spec.oneofKind === "rogue")
            Rogue.internalBinaryWrite(message.spec.rogue, writer.tag(11, WireType.LengthDelimited).fork(), options).join();
        /* proto.ElementalShaman elemental_shaman = 12; */
        if (message.spec.oneofKind === "elementalShaman")
            ElementalShaman.internalBinaryWrite(message.spec.elementalShaman, writer.tag(12, WireType.LengthDelimited).fork(), options).join();
        /* proto.EnhancementShaman enhancement_shaman = 18; */
        if (message.spec.oneofKind === "enhancementShaman")
            EnhancementShaman.internalBinaryWrite(message.spec.enhancementShaman, writer.tag(18, WireType.LengthDelimited).fork(), options).join();
        /* proto.Warlock warlock = 13; */
        if (message.spec.oneofKind === "warlock")
            Warlock.internalBinaryWrite(message.spec.warlock, writer.tag(13, WireType.LengthDelimited).fork(), options).join();
        /* proto.Warrior warrior = 14; */
        if (message.spec.oneofKind === "warrior")
            Warrior.internalBinaryWrite(message.spec.warrior, writer.tag(14, WireType.LengthDelimited).fork(), options).join();
        /* string talentsString = 17; */
        if (message.talentsString !== "")
            writer.tag(17, WireType.LengthDelimited).string(message.talentsString);
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
            { no: 1, name: "players", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Player },
            { no: 2, name: "buffs", kind: "message", T: () => PartyBuffs }
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
                case /* proto.PartyBuffs buffs */ 2:
                    message.buffs = PartyBuffs.internalBinaryRead(reader, reader.uint32(), options, message.buffs);
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
        /* proto.PartyBuffs buffs = 2; */
        if (message.buffs)
            PartyBuffs.internalBinaryWrite(message.buffs, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
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
            { no: 1, name: "parties", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Party },
            { no: 2, name: "buffs", kind: "message", T: () => RaidBuffs }
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
                case /* proto.RaidBuffs buffs */ 2:
                    message.buffs = RaidBuffs.internalBinaryRead(reader, reader.uint32(), options, message.buffs);
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
        /* proto.RaidBuffs buffs = 2; */
        if (message.buffs)
            RaidBuffs.internalBinaryWrite(message.buffs, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
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
            { no: 4, name: "exit_on_oom", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "is_test", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { iterations: 0, randomSeed: 0n, debug: false, exitOnOom: false, isTest: false };
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
                case /* bool is_test */ 5:
                    message.isTest = reader.bool();
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
        /* bool is_test = 5; */
        if (message.isTest !== false)
            writer.tag(5, WireType.Varint).bool(message.isTest);
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
class ActionID$Type extends MessageType {
    constructor() {
        super("proto.ActionID", [
            { no: 1, name: "spell_id", kind: "scalar", oneof: "rawId", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "item_id", kind: "scalar", oneof: "rawId", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "other_id", kind: "enum", oneof: "rawId", T: () => ["proto.OtherAction", OtherAction] },
            { no: 4, name: "tag", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value) {
        const message = { rawId: { oneofKind: undefined }, tag: 0 };
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
                    message.rawId = {
                        oneofKind: "spellId",
                        spellId: reader.int32()
                    };
                    break;
                case /* int32 item_id */ 2:
                    message.rawId = {
                        oneofKind: "itemId",
                        itemId: reader.int32()
                    };
                    break;
                case /* proto.OtherAction other_id */ 3:
                    message.rawId = {
                        oneofKind: "otherId",
                        otherId: reader.int32()
                    };
                    break;
                case /* int32 tag */ 4:
                    message.tag = reader.int32();
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
        if (message.rawId.oneofKind === "spellId")
            writer.tag(1, WireType.Varint).int32(message.rawId.spellId);
        /* int32 item_id = 2; */
        if (message.rawId.oneofKind === "itemId")
            writer.tag(2, WireType.Varint).int32(message.rawId.itemId);
        /* proto.OtherAction other_id = 3; */
        if (message.rawId.oneofKind === "otherId")
            writer.tag(3, WireType.Varint).int32(message.rawId.otherId);
        /* int32 tag = 4; */
        if (message.tag !== 0)
            writer.tag(4, WireType.Varint).int32(message.tag);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ActionID
 */
export const ActionID = new ActionID$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ActionMetrics$Type extends MessageType {
    constructor() {
        super("proto.ActionMetrics", [
            { no: 1, name: "id", kind: "message", T: () => ActionID },
            { no: 2, name: "casts", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "hits", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "crits", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "misses", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "damage", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { casts: 0, hits: 0, crits: 0, misses: 0, damage: 0 };
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
                case /* proto.ActionID id */ 1:
                    message.id = ActionID.internalBinaryRead(reader, reader.uint32(), options, message.id);
                    break;
                case /* int32 casts */ 2:
                    message.casts = reader.int32();
                    break;
                case /* int32 hits */ 3:
                    message.hits = reader.int32();
                    break;
                case /* int32 crits */ 4:
                    message.crits = reader.int32();
                    break;
                case /* int32 misses */ 5:
                    message.misses = reader.int32();
                    break;
                case /* double damage */ 6:
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
        /* proto.ActionID id = 1; */
        if (message.id)
            ActionID.internalBinaryWrite(message.id, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* int32 casts = 2; */
        if (message.casts !== 0)
            writer.tag(2, WireType.Varint).int32(message.casts);
        /* int32 hits = 3; */
        if (message.hits !== 0)
            writer.tag(3, WireType.Varint).int32(message.hits);
        /* int32 crits = 4; */
        if (message.crits !== 0)
            writer.tag(4, WireType.Varint).int32(message.crits);
        /* int32 misses = 5; */
        if (message.misses !== 0)
            writer.tag(5, WireType.Varint).int32(message.misses);
        /* double damage = 6; */
        if (message.damage !== 0)
            writer.tag(6, WireType.Bit64).double(message.damage);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ActionMetrics
 */
export const ActionMetrics = new ActionMetrics$Type();
// @generated message type with reflection information, may provide speed optimized methods
class AuraMetrics$Type extends MessageType {
    constructor() {
        super("proto.AuraMetrics", [
            { no: 1, name: "id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "uptime_seconds_avg", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 3, name: "uptime_seconds_stdev", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { id: 0, uptimeSecondsAvg: 0, uptimeSecondsStdev: 0 };
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
                case /* int32 id */ 1:
                    message.id = reader.int32();
                    break;
                case /* double uptime_seconds_avg */ 2:
                    message.uptimeSecondsAvg = reader.double();
                    break;
                case /* double uptime_seconds_stdev */ 3:
                    message.uptimeSecondsStdev = reader.double();
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
        /* int32 id = 1; */
        if (message.id !== 0)
            writer.tag(1, WireType.Varint).int32(message.id);
        /* double uptime_seconds_avg = 2; */
        if (message.uptimeSecondsAvg !== 0)
            writer.tag(2, WireType.Bit64).double(message.uptimeSecondsAvg);
        /* double uptime_seconds_stdev = 3; */
        if (message.uptimeSecondsStdev !== 0)
            writer.tag(3, WireType.Bit64).double(message.uptimeSecondsStdev);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.AuraMetrics
 */
export const AuraMetrics = new AuraMetrics$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DpsMetrics$Type extends MessageType {
    constructor() {
        super("proto.DpsMetrics", [
            { no: 1, name: "avg", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 2, name: "stdev", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 3, name: "max", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 4, name: "hist", kind: "map", K: 5 /*ScalarType.INT32*/, V: { kind: "scalar", T: 5 /*ScalarType.INT32*/ } }
        ]);
    }
    create(value) {
        const message = { avg: 0, stdev: 0, max: 0, hist: {} };
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
                case /* double avg */ 1:
                    message.avg = reader.double();
                    break;
                case /* double stdev */ 2:
                    message.stdev = reader.double();
                    break;
                case /* double max */ 3:
                    message.max = reader.double();
                    break;
                case /* map<int32, int32> hist */ 4:
                    this.binaryReadMap4(message.hist, reader, options);
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
    binaryReadMap4(map, reader, options) {
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
                default: throw new globalThis.Error("unknown map entry field for field proto.DpsMetrics.hist");
            }
        }
        map[key ?? 0] = val ?? 0;
    }
    internalBinaryWrite(message, writer, options) {
        /* double avg = 1; */
        if (message.avg !== 0)
            writer.tag(1, WireType.Bit64).double(message.avg);
        /* double stdev = 2; */
        if (message.stdev !== 0)
            writer.tag(2, WireType.Bit64).double(message.stdev);
        /* double max = 3; */
        if (message.max !== 0)
            writer.tag(3, WireType.Bit64).double(message.max);
        /* map<int32, int32> hist = 4; */
        for (let k of Object.keys(message.hist))
            writer.tag(4, WireType.LengthDelimited).fork().tag(1, WireType.Varint).int32(parseInt(k)).tag(2, WireType.Varint).int32(message.hist[k]).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.DpsMetrics
 */
export const DpsMetrics = new DpsMetrics$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PlayerMetrics$Type extends MessageType {
    constructor() {
        super("proto.PlayerMetrics", [
            { no: 1, name: "dps", kind: "message", T: () => DpsMetrics },
            { no: 2, name: "num_oom", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "oom_at_avg", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 4, name: "dps_at_oom_avg", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 5, name: "actions", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => ActionMetrics },
            { no: 6, name: "auras", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => AuraMetrics }
        ]);
    }
    create(value) {
        const message = { numOom: 0, oomAtAvg: 0, dpsAtOomAvg: 0, actions: [], auras: [] };
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
                case /* proto.DpsMetrics dps */ 1:
                    message.dps = DpsMetrics.internalBinaryRead(reader, reader.uint32(), options, message.dps);
                    break;
                case /* int32 num_oom */ 2:
                    message.numOom = reader.int32();
                    break;
                case /* double oom_at_avg */ 3:
                    message.oomAtAvg = reader.double();
                    break;
                case /* double dps_at_oom_avg */ 4:
                    message.dpsAtOomAvg = reader.double();
                    break;
                case /* repeated proto.ActionMetrics actions */ 5:
                    message.actions.push(ActionMetrics.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* repeated proto.AuraMetrics auras */ 6:
                    message.auras.push(AuraMetrics.internalBinaryRead(reader, reader.uint32(), options));
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
        /* proto.DpsMetrics dps = 1; */
        if (message.dps)
            DpsMetrics.internalBinaryWrite(message.dps, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* int32 num_oom = 2; */
        if (message.numOom !== 0)
            writer.tag(2, WireType.Varint).int32(message.numOom);
        /* double oom_at_avg = 3; */
        if (message.oomAtAvg !== 0)
            writer.tag(3, WireType.Bit64).double(message.oomAtAvg);
        /* double dps_at_oom_avg = 4; */
        if (message.dpsAtOomAvg !== 0)
            writer.tag(4, WireType.Bit64).double(message.dpsAtOomAvg);
        /* repeated proto.ActionMetrics actions = 5; */
        for (let i = 0; i < message.actions.length; i++)
            ActionMetrics.internalBinaryWrite(message.actions[i], writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* repeated proto.AuraMetrics auras = 6; */
        for (let i = 0; i < message.auras.length; i++)
            AuraMetrics.internalBinaryWrite(message.auras[i], writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.PlayerMetrics
 */
export const PlayerMetrics = new PlayerMetrics$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PartyMetrics$Type extends MessageType {
    constructor() {
        super("proto.PartyMetrics", [
            { no: 1, name: "dps", kind: "message", T: () => DpsMetrics },
            { no: 2, name: "players", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => PlayerMetrics }
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
                case /* proto.DpsMetrics dps */ 1:
                    message.dps = DpsMetrics.internalBinaryRead(reader, reader.uint32(), options, message.dps);
                    break;
                case /* repeated proto.PlayerMetrics players */ 2:
                    message.players.push(PlayerMetrics.internalBinaryRead(reader, reader.uint32(), options));
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
        /* proto.DpsMetrics dps = 1; */
        if (message.dps)
            DpsMetrics.internalBinaryWrite(message.dps, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated proto.PlayerMetrics players = 2; */
        for (let i = 0; i < message.players.length; i++)
            PlayerMetrics.internalBinaryWrite(message.players[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.PartyMetrics
 */
export const PartyMetrics = new PartyMetrics$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RaidMetrics$Type extends MessageType {
    constructor() {
        super("proto.RaidMetrics", [
            { no: 1, name: "dps", kind: "message", T: () => DpsMetrics },
            { no: 2, name: "parties", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => PartyMetrics }
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
                case /* proto.DpsMetrics dps */ 1:
                    message.dps = DpsMetrics.internalBinaryRead(reader, reader.uint32(), options, message.dps);
                    break;
                case /* repeated proto.PartyMetrics parties */ 2:
                    message.parties.push(PartyMetrics.internalBinaryRead(reader, reader.uint32(), options));
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
        /* proto.DpsMetrics dps = 1; */
        if (message.dps)
            DpsMetrics.internalBinaryWrite(message.dps, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated proto.PartyMetrics parties = 2; */
        for (let i = 0; i < message.parties.length; i++)
            PartyMetrics.internalBinaryWrite(message.parties[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.RaidMetrics
 */
export const RaidMetrics = new RaidMetrics$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TargetMetrics$Type extends MessageType {
    constructor() {
        super("proto.TargetMetrics", [
            { no: 1, name: "auras", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => AuraMetrics }
        ]);
    }
    create(value) {
        const message = { auras: [] };
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
                case /* repeated proto.AuraMetrics auras */ 1:
                    message.auras.push(AuraMetrics.internalBinaryRead(reader, reader.uint32(), options));
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
        /* repeated proto.AuraMetrics auras = 1; */
        for (let i = 0; i < message.auras.length; i++)
            AuraMetrics.internalBinaryWrite(message.auras[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.TargetMetrics
 */
export const TargetMetrics = new TargetMetrics$Type();
// @generated message type with reflection information, may provide speed optimized methods
class EncounterMetrics$Type extends MessageType {
    constructor() {
        super("proto.EncounterMetrics", [
            { no: 1, name: "targets", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => TargetMetrics }
        ]);
    }
    create(value) {
        const message = { targets: [] };
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
                case /* repeated proto.TargetMetrics targets */ 1:
                    message.targets.push(TargetMetrics.internalBinaryRead(reader, reader.uint32(), options));
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
        /* repeated proto.TargetMetrics targets = 1; */
        for (let i = 0; i < message.targets.length; i++)
            TargetMetrics.internalBinaryWrite(message.targets[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.EncounterMetrics
 */
export const EncounterMetrics = new EncounterMetrics$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RaidSimRequest$Type extends MessageType {
    constructor() {
        super("proto.RaidSimRequest", [
            { no: 1, name: "raid", kind: "message", T: () => Raid },
            { no: 2, name: "encounter", kind: "message", T: () => Encounter },
            { no: 3, name: "sim_options", kind: "message", T: () => SimOptions }
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
                case /* proto.Raid raid */ 1:
                    message.raid = Raid.internalBinaryRead(reader, reader.uint32(), options, message.raid);
                    break;
                case /* proto.Encounter encounter */ 2:
                    message.encounter = Encounter.internalBinaryRead(reader, reader.uint32(), options, message.encounter);
                    break;
                case /* proto.SimOptions sim_options */ 3:
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
        /* proto.Raid raid = 1; */
        if (message.raid)
            Raid.internalBinaryWrite(message.raid, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.Encounter encounter = 2; */
        if (message.encounter)
            Encounter.internalBinaryWrite(message.encounter, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.SimOptions sim_options = 3; */
        if (message.simOptions)
            SimOptions.internalBinaryWrite(message.simOptions, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
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
class RaidSimResult$Type extends MessageType {
    constructor() {
        super("proto.RaidSimResult", [
            { no: 1, name: "raid_metrics", kind: "message", T: () => RaidMetrics },
            { no: 2, name: "encounter_metrics", kind: "message", T: () => EncounterMetrics },
            { no: 3, name: "logs", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value) {
        const message = { logs: "" };
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
                case /* proto.RaidMetrics raid_metrics */ 1:
                    message.raidMetrics = RaidMetrics.internalBinaryRead(reader, reader.uint32(), options, message.raidMetrics);
                    break;
                case /* proto.EncounterMetrics encounter_metrics */ 2:
                    message.encounterMetrics = EncounterMetrics.internalBinaryRead(reader, reader.uint32(), options, message.encounterMetrics);
                    break;
                case /* string logs */ 3:
                    message.logs = reader.string();
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
        /* proto.RaidMetrics raid_metrics = 1; */
        if (message.raidMetrics)
            RaidMetrics.internalBinaryWrite(message.raidMetrics, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.EncounterMetrics encounter_metrics = 2; */
        if (message.encounterMetrics)
            EncounterMetrics.internalBinaryWrite(message.encounterMetrics, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* string logs = 3; */
        if (message.logs !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.logs);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.RaidSimResult
 */
export const RaidSimResult = new RaidSimResult$Type();
// @generated message type with reflection information, may provide speed optimized methods
class IndividualSimRequest$Type extends MessageType {
    constructor() {
        super("proto.IndividualSimRequest", [
            { no: 1, name: "player", kind: "message", T: () => Player },
            { no: 2, name: "raid_buffs", kind: "message", T: () => RaidBuffs },
            { no: 3, name: "party_buffs", kind: "message", T: () => PartyBuffs },
            { no: 5, name: "encounter", kind: "message", T: () => Encounter },
            { no: 6, name: "sim_options", kind: "message", T: () => SimOptions }
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
                case /* proto.RaidBuffs raid_buffs */ 2:
                    message.raidBuffs = RaidBuffs.internalBinaryRead(reader, reader.uint32(), options, message.raidBuffs);
                    break;
                case /* proto.PartyBuffs party_buffs */ 3:
                    message.partyBuffs = PartyBuffs.internalBinaryRead(reader, reader.uint32(), options, message.partyBuffs);
                    break;
                case /* proto.Encounter encounter */ 5:
                    message.encounter = Encounter.internalBinaryRead(reader, reader.uint32(), options, message.encounter);
                    break;
                case /* proto.SimOptions sim_options */ 6:
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
        /* proto.RaidBuffs raid_buffs = 2; */
        if (message.raidBuffs)
            RaidBuffs.internalBinaryWrite(message.raidBuffs, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.PartyBuffs party_buffs = 3; */
        if (message.partyBuffs)
            PartyBuffs.internalBinaryWrite(message.partyBuffs, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* proto.Encounter encounter = 5; */
        if (message.encounter)
            Encounter.internalBinaryWrite(message.encounter, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* proto.SimOptions sim_options = 6; */
        if (message.simOptions)
            SimOptions.internalBinaryWrite(message.simOptions, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
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
            { no: 1, name: "player_metrics", kind: "message", T: () => PlayerMetrics },
            { no: 2, name: "encounter_metrics", kind: "message", T: () => EncounterMetrics },
            { no: 3, name: "logs", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value) {
        const message = { logs: "" };
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
                case /* proto.PlayerMetrics player_metrics */ 1:
                    message.playerMetrics = PlayerMetrics.internalBinaryRead(reader, reader.uint32(), options, message.playerMetrics);
                    break;
                case /* proto.EncounterMetrics encounter_metrics */ 2:
                    message.encounterMetrics = EncounterMetrics.internalBinaryRead(reader, reader.uint32(), options, message.encounterMetrics);
                    break;
                case /* string logs */ 3:
                    message.logs = reader.string();
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
        /* proto.PlayerMetrics player_metrics = 1; */
        if (message.playerMetrics)
            PlayerMetrics.internalBinaryWrite(message.playerMetrics, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.EncounterMetrics encounter_metrics = 2; */
        if (message.encounterMetrics)
            EncounterMetrics.internalBinaryWrite(message.encounterMetrics, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* string logs = 3; */
        if (message.logs !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.logs);
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
class GearListRequest$Type extends MessageType {
    constructor() {
        super("proto.GearListRequest", []);
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
            { no: 1, name: "raid", kind: "message", T: () => Raid }
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
                case /* proto.Raid raid */ 1:
                    message.raid = Raid.internalBinaryRead(reader, reader.uint32(), options, message.raid);
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
class PlayerStats$Type extends MessageType {
    constructor() {
        super("proto.PlayerStats", [
            { no: 1, name: "gear_only", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 2, name: "final_stats", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 3, name: "sets", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "buffs", kind: "message", T: () => IndividualBuffs }
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
                case /* repeated double final_stats */ 2:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.finalStats.push(reader.double());
                    else
                        message.finalStats.push(reader.double());
                    break;
                case /* repeated string sets */ 3:
                    message.sets.push(reader.string());
                    break;
                case /* proto.IndividualBuffs buffs */ 4:
                    message.buffs = IndividualBuffs.internalBinaryRead(reader, reader.uint32(), options, message.buffs);
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
        /* repeated double final_stats = 2; */
        if (message.finalStats.length) {
            writer.tag(2, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.finalStats.length; i++)
                writer.double(message.finalStats[i]);
            writer.join();
        }
        /* repeated string sets = 3; */
        for (let i = 0; i < message.sets.length; i++)
            writer.tag(3, WireType.LengthDelimited).string(message.sets[i]);
        /* proto.IndividualBuffs buffs = 4; */
        if (message.buffs)
            IndividualBuffs.internalBinaryWrite(message.buffs, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.PlayerStats
 */
export const PlayerStats = new PlayerStats$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PartyStats$Type extends MessageType {
    constructor() {
        super("proto.PartyStats", [
            { no: 1, name: "players", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => PlayerStats }
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
                case /* repeated proto.PlayerStats players */ 1:
                    message.players.push(PlayerStats.internalBinaryRead(reader, reader.uint32(), options));
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
        /* repeated proto.PlayerStats players = 1; */
        for (let i = 0; i < message.players.length; i++)
            PlayerStats.internalBinaryWrite(message.players[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.PartyStats
 */
export const PartyStats = new PartyStats$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RaidStats$Type extends MessageType {
    constructor() {
        super("proto.RaidStats", [
            { no: 1, name: "parties", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => PartyStats }
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
                case /* repeated proto.PartyStats parties */ 1:
                    message.parties.push(PartyStats.internalBinaryRead(reader, reader.uint32(), options));
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
        /* repeated proto.PartyStats parties = 1; */
        for (let i = 0; i < message.parties.length; i++)
            PartyStats.internalBinaryWrite(message.parties[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.RaidStats
 */
export const RaidStats = new RaidStats$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ComputeStatsResult$Type extends MessageType {
    constructor() {
        super("proto.ComputeStatsResult", [
            { no: 1, name: "raid_stats", kind: "message", T: () => RaidStats }
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
                case /* proto.RaidStats raid_stats */ 1:
                    message.raidStats = RaidStats.internalBinaryRead(reader, reader.uint32(), options, message.raidStats);
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
        /* proto.RaidStats raid_stats = 1; */
        if (message.raidStats)
            RaidStats.internalBinaryWrite(message.raidStats, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
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
            { no: 1, name: "player", kind: "message", T: () => Player },
            { no: 2, name: "raid_buffs", kind: "message", T: () => RaidBuffs },
            { no: 3, name: "party_buffs", kind: "message", T: () => PartyBuffs },
            { no: 4, name: "encounter", kind: "message", T: () => Encounter },
            { no: 5, name: "sim_options", kind: "message", T: () => SimOptions },
            { no: 6, name: "stats_to_weigh", kind: "enum", repeat: 1 /*RepeatType.PACKED*/, T: () => ["proto.Stat", Stat] },
            { no: 7, name: "ep_reference_stat", kind: "enum", T: () => ["proto.Stat", Stat] }
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
                case /* proto.Player player */ 1:
                    message.player = Player.internalBinaryRead(reader, reader.uint32(), options, message.player);
                    break;
                case /* proto.RaidBuffs raid_buffs */ 2:
                    message.raidBuffs = RaidBuffs.internalBinaryRead(reader, reader.uint32(), options, message.raidBuffs);
                    break;
                case /* proto.PartyBuffs party_buffs */ 3:
                    message.partyBuffs = PartyBuffs.internalBinaryRead(reader, reader.uint32(), options, message.partyBuffs);
                    break;
                case /* proto.Encounter encounter */ 4:
                    message.encounter = Encounter.internalBinaryRead(reader, reader.uint32(), options, message.encounter);
                    break;
                case /* proto.SimOptions sim_options */ 5:
                    message.simOptions = SimOptions.internalBinaryRead(reader, reader.uint32(), options, message.simOptions);
                    break;
                case /* repeated proto.Stat stats_to_weigh */ 6:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.statsToWeigh.push(reader.int32());
                    else
                        message.statsToWeigh.push(reader.int32());
                    break;
                case /* proto.Stat ep_reference_stat */ 7:
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
        /* proto.Player player = 1; */
        if (message.player)
            Player.internalBinaryWrite(message.player, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.RaidBuffs raid_buffs = 2; */
        if (message.raidBuffs)
            RaidBuffs.internalBinaryWrite(message.raidBuffs, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.PartyBuffs party_buffs = 3; */
        if (message.partyBuffs)
            PartyBuffs.internalBinaryWrite(message.partyBuffs, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* proto.Encounter encounter = 4; */
        if (message.encounter)
            Encounter.internalBinaryWrite(message.encounter, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* proto.SimOptions sim_options = 5; */
        if (message.simOptions)
            SimOptions.internalBinaryWrite(message.simOptions, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* repeated proto.Stat stats_to_weigh = 6; */
        if (message.statsToWeigh.length) {
            writer.tag(6, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.statsToWeigh.length; i++)
                writer.int32(message.statsToWeigh[i]);
            writer.join();
        }
        /* proto.Stat ep_reference_stat = 7; */
        if (message.epReferenceStat !== 0)
            writer.tag(7, WireType.Varint).int32(message.epReferenceStat);
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
