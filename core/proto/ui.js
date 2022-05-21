import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
import { RaidSimResult } from './api.js';
import { RaidSimRequest } from './api.js';
import { Raid } from './api.js';
import { Blessings } from './paladin.js';
import { Cooldowns } from './common.js';
import { Race } from './common.js';
import { Consumes } from './common.js';
import { IndividualBuffs } from './common.js';
import { EquipmentSpec } from './common.js';
import { Encounter } from './common.js';
import { Player } from './api.js';
import { PartyBuffs } from './common.js';
import { RaidTarget } from './common.js';
import { Debuffs } from './common.js';
import { RaidBuffs } from './common.js';
// @generated message type with reflection information, may provide speed optimized methods
class SimSettings$Type extends MessageType {
    constructor() {
        super("proto.SimSettings", [
            { no: 1, name: "iterations", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "phase", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "fixed_rng_seed", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 4, name: "show_threat_metrics", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "show_experimental", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { iterations: 0, phase: 0, fixedRngSeed: 0n, showThreatMetrics: false, showExperimental: false };
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
                case /* int32 phase */ 2:
                    message.phase = reader.int32();
                    break;
                case /* int64 fixed_rng_seed */ 3:
                    message.fixedRngSeed = reader.int64().toBigInt();
                    break;
                case /* bool show_threat_metrics */ 4:
                    message.showThreatMetrics = reader.bool();
                    break;
                case /* bool show_experimental */ 5:
                    message.showExperimental = reader.bool();
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
        /* int32 phase = 2; */
        if (message.phase !== 0)
            writer.tag(2, WireType.Varint).int32(message.phase);
        /* int64 fixed_rng_seed = 3; */
        if (message.fixedRngSeed !== 0n)
            writer.tag(3, WireType.Varint).int64(message.fixedRngSeed);
        /* bool show_threat_metrics = 4; */
        if (message.showThreatMetrics !== false)
            writer.tag(4, WireType.Varint).bool(message.showThreatMetrics);
        /* bool show_experimental = 5; */
        if (message.showExperimental !== false)
            writer.tag(5, WireType.Varint).bool(message.showExperimental);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.SimSettings
 */
export const SimSettings = new SimSettings$Type();
// @generated message type with reflection information, may provide speed optimized methods
class IndividualSimSettings$Type extends MessageType {
    constructor() {
        super("proto.IndividualSimSettings", [
            { no: 5, name: "settings", kind: "message", T: () => SimSettings },
            { no: 1, name: "raid_buffs", kind: "message", T: () => RaidBuffs },
            { no: 8, name: "debuffs", kind: "message", T: () => Debuffs },
            { no: 7, name: "tanks", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => RaidTarget },
            { no: 2, name: "party_buffs", kind: "message", T: () => PartyBuffs },
            { no: 3, name: "player", kind: "message", T: () => Player },
            { no: 4, name: "encounter", kind: "message", T: () => Encounter },
            { no: 6, name: "ep_weights", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { tanks: [], epWeights: [] };
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
                case /* proto.SimSettings settings */ 5:
                    message.settings = SimSettings.internalBinaryRead(reader, reader.uint32(), options, message.settings);
                    break;
                case /* proto.RaidBuffs raid_buffs */ 1:
                    message.raidBuffs = RaidBuffs.internalBinaryRead(reader, reader.uint32(), options, message.raidBuffs);
                    break;
                case /* proto.Debuffs debuffs */ 8:
                    message.debuffs = Debuffs.internalBinaryRead(reader, reader.uint32(), options, message.debuffs);
                    break;
                case /* repeated proto.RaidTarget tanks */ 7:
                    message.tanks.push(RaidTarget.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* proto.PartyBuffs party_buffs */ 2:
                    message.partyBuffs = PartyBuffs.internalBinaryRead(reader, reader.uint32(), options, message.partyBuffs);
                    break;
                case /* proto.Player player */ 3:
                    message.player = Player.internalBinaryRead(reader, reader.uint32(), options, message.player);
                    break;
                case /* proto.Encounter encounter */ 4:
                    message.encounter = Encounter.internalBinaryRead(reader, reader.uint32(), options, message.encounter);
                    break;
                case /* repeated double ep_weights */ 6:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.epWeights.push(reader.double());
                    else
                        message.epWeights.push(reader.double());
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
        /* proto.SimSettings settings = 5; */
        if (message.settings)
            SimSettings.internalBinaryWrite(message.settings, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* proto.RaidBuffs raid_buffs = 1; */
        if (message.raidBuffs)
            RaidBuffs.internalBinaryWrite(message.raidBuffs, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.Debuffs debuffs = 8; */
        if (message.debuffs)
            Debuffs.internalBinaryWrite(message.debuffs, writer.tag(8, WireType.LengthDelimited).fork(), options).join();
        /* repeated proto.RaidTarget tanks = 7; */
        for (let i = 0; i < message.tanks.length; i++)
            RaidTarget.internalBinaryWrite(message.tanks[i], writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        /* proto.PartyBuffs party_buffs = 2; */
        if (message.partyBuffs)
            PartyBuffs.internalBinaryWrite(message.partyBuffs, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.Player player = 3; */
        if (message.player)
            Player.internalBinaryWrite(message.player, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* proto.Encounter encounter = 4; */
        if (message.encounter)
            Encounter.internalBinaryWrite(message.encounter, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* repeated double ep_weights = 6; */
        if (message.epWeights.length) {
            writer.tag(6, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.epWeights.length; i++)
                writer.double(message.epWeights[i]);
            writer.join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.IndividualSimSettings
 */
export const IndividualSimSettings = new IndividualSimSettings$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SavedGearSet$Type extends MessageType {
    constructor() {
        super("proto.SavedGearSet", [
            { no: 1, name: "gear", kind: "message", T: () => EquipmentSpec },
            { no: 2, name: "bonus_stats", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { bonusStats: [] };
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
                case /* proto.EquipmentSpec gear */ 1:
                    message.gear = EquipmentSpec.internalBinaryRead(reader, reader.uint32(), options, message.gear);
                    break;
                case /* repeated double bonus_stats */ 2:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.bonusStats.push(reader.double());
                    else
                        message.bonusStats.push(reader.double());
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
        /* proto.EquipmentSpec gear = 1; */
        if (message.gear)
            EquipmentSpec.internalBinaryWrite(message.gear, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated double bonus_stats = 2; */
        if (message.bonusStats.length) {
            writer.tag(2, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.bonusStats.length; i++)
                writer.double(message.bonusStats[i]);
            writer.join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.SavedGearSet
 */
export const SavedGearSet = new SavedGearSet$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SavedSettings$Type extends MessageType {
    constructor() {
        super("proto.SavedSettings", [
            { no: 1, name: "raid_buffs", kind: "message", T: () => RaidBuffs },
            { no: 2, name: "party_buffs", kind: "message", T: () => PartyBuffs },
            { no: 7, name: "debuffs", kind: "message", T: () => Debuffs },
            { no: 3, name: "player_buffs", kind: "message", T: () => IndividualBuffs },
            { no: 4, name: "consumes", kind: "message", T: () => Consumes },
            { no: 5, name: "race", kind: "enum", T: () => ["proto.Race", Race] },
            { no: 6, name: "cooldowns", kind: "message", T: () => Cooldowns }
        ]);
    }
    create(value) {
        const message = { race: 0 };
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
                case /* proto.RaidBuffs raid_buffs */ 1:
                    message.raidBuffs = RaidBuffs.internalBinaryRead(reader, reader.uint32(), options, message.raidBuffs);
                    break;
                case /* proto.PartyBuffs party_buffs */ 2:
                    message.partyBuffs = PartyBuffs.internalBinaryRead(reader, reader.uint32(), options, message.partyBuffs);
                    break;
                case /* proto.Debuffs debuffs */ 7:
                    message.debuffs = Debuffs.internalBinaryRead(reader, reader.uint32(), options, message.debuffs);
                    break;
                case /* proto.IndividualBuffs player_buffs */ 3:
                    message.playerBuffs = IndividualBuffs.internalBinaryRead(reader, reader.uint32(), options, message.playerBuffs);
                    break;
                case /* proto.Consumes consumes */ 4:
                    message.consumes = Consumes.internalBinaryRead(reader, reader.uint32(), options, message.consumes);
                    break;
                case /* proto.Race race */ 5:
                    message.race = reader.int32();
                    break;
                case /* proto.Cooldowns cooldowns */ 6:
                    message.cooldowns = Cooldowns.internalBinaryRead(reader, reader.uint32(), options, message.cooldowns);
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
        /* proto.RaidBuffs raid_buffs = 1; */
        if (message.raidBuffs)
            RaidBuffs.internalBinaryWrite(message.raidBuffs, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.PartyBuffs party_buffs = 2; */
        if (message.partyBuffs)
            PartyBuffs.internalBinaryWrite(message.partyBuffs, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.Debuffs debuffs = 7; */
        if (message.debuffs)
            Debuffs.internalBinaryWrite(message.debuffs, writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        /* proto.IndividualBuffs player_buffs = 3; */
        if (message.playerBuffs)
            IndividualBuffs.internalBinaryWrite(message.playerBuffs, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* proto.Consumes consumes = 4; */
        if (message.consumes)
            Consumes.internalBinaryWrite(message.consumes, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* proto.Race race = 5; */
        if (message.race !== 0)
            writer.tag(5, WireType.Varint).int32(message.race);
        /* proto.Cooldowns cooldowns = 6; */
        if (message.cooldowns)
            Cooldowns.internalBinaryWrite(message.cooldowns, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.SavedSettings
 */
export const SavedSettings = new SavedSettings$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SavedTalents$Type extends MessageType {
    constructor() {
        super("proto.SavedTalents", [
            { no: 1, name: "talents_string", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value) {
        const message = { talentsString: "" };
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
                case /* string talents_string */ 1:
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
        /* string talents_string = 1; */
        if (message.talentsString !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.talentsString);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.SavedTalents
 */
export const SavedTalents = new SavedTalents$Type();
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
// @generated message type with reflection information, may provide speed optimized methods
class SavedEncounter$Type extends MessageType {
    constructor() {
        super("proto.SavedEncounter", [
            { no: 1, name: "encounter", kind: "message", T: () => Encounter }
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
                case /* proto.Encounter encounter */ 1:
                    message.encounter = Encounter.internalBinaryRead(reader, reader.uint32(), options, message.encounter);
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
        /* proto.Encounter encounter = 1; */
        if (message.encounter)
            Encounter.internalBinaryWrite(message.encounter, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.SavedEncounter
 */
export const SavedEncounter = new SavedEncounter$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SavedRaid$Type extends MessageType {
    constructor() {
        super("proto.SavedRaid", [
            { no: 1, name: "raid", kind: "message", T: () => Raid },
            { no: 2, name: "buff_bots", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => BuffBot },
            { no: 3, name: "blessings", kind: "message", T: () => BlessingsAssignments }
        ]);
    }
    create(value) {
        const message = { buffBots: [] };
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
                case /* repeated proto.BuffBot buff_bots */ 2:
                    message.buffBots.push(BuffBot.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* proto.BlessingsAssignments blessings */ 3:
                    message.blessings = BlessingsAssignments.internalBinaryRead(reader, reader.uint32(), options, message.blessings);
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
        /* repeated proto.BuffBot buff_bots = 2; */
        for (let i = 0; i < message.buffBots.length; i++)
            BuffBot.internalBinaryWrite(message.buffBots[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.BlessingsAssignments blessings = 3; */
        if (message.blessings)
            BlessingsAssignments.internalBinaryWrite(message.blessings, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.SavedRaid
 */
export const SavedRaid = new SavedRaid$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RaidSimSettings$Type extends MessageType {
    constructor() {
        super("proto.RaidSimSettings", [
            { no: 5, name: "settings", kind: "message", T: () => SimSettings },
            { no: 1, name: "raid", kind: "message", T: () => Raid },
            { no: 2, name: "buff_bots", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => BuffBot },
            { no: 3, name: "blessings", kind: "message", T: () => BlessingsAssignments },
            { no: 4, name: "encounter", kind: "message", T: () => Encounter }
        ]);
    }
    create(value) {
        const message = { buffBots: [] };
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
                case /* proto.SimSettings settings */ 5:
                    message.settings = SimSettings.internalBinaryRead(reader, reader.uint32(), options, message.settings);
                    break;
                case /* proto.Raid raid */ 1:
                    message.raid = Raid.internalBinaryRead(reader, reader.uint32(), options, message.raid);
                    break;
                case /* repeated proto.BuffBot buff_bots */ 2:
                    message.buffBots.push(BuffBot.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* proto.BlessingsAssignments blessings */ 3:
                    message.blessings = BlessingsAssignments.internalBinaryRead(reader, reader.uint32(), options, message.blessings);
                    break;
                case /* proto.Encounter encounter */ 4:
                    message.encounter = Encounter.internalBinaryRead(reader, reader.uint32(), options, message.encounter);
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
        /* proto.SimSettings settings = 5; */
        if (message.settings)
            SimSettings.internalBinaryWrite(message.settings, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* proto.Raid raid = 1; */
        if (message.raid)
            Raid.internalBinaryWrite(message.raid, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated proto.BuffBot buff_bots = 2; */
        for (let i = 0; i < message.buffBots.length; i++)
            BuffBot.internalBinaryWrite(message.buffBots[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.BlessingsAssignments blessings = 3; */
        if (message.blessings)
            BlessingsAssignments.internalBinaryWrite(message.blessings, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* proto.Encounter encounter = 4; */
        if (message.encounter)
            Encounter.internalBinaryWrite(message.encounter, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.RaidSimSettings
 */
export const RaidSimSettings = new RaidSimSettings$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SimRun$Type extends MessageType {
    constructor() {
        super("proto.SimRun", [
            { no: 1, name: "request", kind: "message", T: () => RaidSimRequest },
            { no: 2, name: "result", kind: "message", T: () => RaidSimResult }
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
                case /* proto.RaidSimRequest request */ 1:
                    message.request = RaidSimRequest.internalBinaryRead(reader, reader.uint32(), options, message.request);
                    break;
                case /* proto.RaidSimResult result */ 2:
                    message.result = RaidSimResult.internalBinaryRead(reader, reader.uint32(), options, message.result);
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
        /* proto.RaidSimRequest request = 1; */
        if (message.request)
            RaidSimRequest.internalBinaryWrite(message.request, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.RaidSimResult result = 2; */
        if (message.result)
            RaidSimResult.internalBinaryWrite(message.result, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.SimRun
 */
export const SimRun = new SimRun$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SimRunData$Type extends MessageType {
    constructor() {
        super("proto.SimRunData", [
            { no: 1, name: "run", kind: "message", T: () => SimRun },
            { no: 2, name: "reference_run", kind: "message", T: () => SimRun }
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
                case /* proto.SimRun run */ 1:
                    message.run = SimRun.internalBinaryRead(reader, reader.uint32(), options, message.run);
                    break;
                case /* proto.SimRun reference_run */ 2:
                    message.referenceRun = SimRun.internalBinaryRead(reader, reader.uint32(), options, message.referenceRun);
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
        /* proto.SimRun run = 1; */
        if (message.run)
            SimRun.internalBinaryWrite(message.run, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.SimRun reference_run = 2; */
        if (message.referenceRun)
            SimRun.internalBinaryWrite(message.referenceRun, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.SimRunData
 */
export const SimRunData = new SimRunData$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DetailedResultsUpdate$Type extends MessageType {
    constructor() {
        super("proto.DetailedResultsUpdate", [
            { no: 1, name: "run_data", kind: "message", oneof: "data", T: () => SimRunData },
            { no: 2, name: "settings", kind: "message", oneof: "data", T: () => SimSettings }
        ]);
    }
    create(value) {
        const message = { data: { oneofKind: undefined } };
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
                case /* proto.SimRunData run_data */ 1:
                    message.data = {
                        oneofKind: "runData",
                        runData: SimRunData.internalBinaryRead(reader, reader.uint32(), options, message.data.runData)
                    };
                    break;
                case /* proto.SimSettings settings */ 2:
                    message.data = {
                        oneofKind: "settings",
                        settings: SimSettings.internalBinaryRead(reader, reader.uint32(), options, message.data.settings)
                    };
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
        /* proto.SimRunData run_data = 1; */
        if (message.data.oneofKind === "runData")
            SimRunData.internalBinaryWrite(message.data.runData, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.SimSettings settings = 2; */
        if (message.data.oneofKind === "settings")
            SimSettings.internalBinaryWrite(message.data.settings, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.DetailedResultsUpdate
 */
export const DetailedResultsUpdate = new DetailedResultsUpdate$Type();
