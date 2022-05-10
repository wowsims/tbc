import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
import { RaidTarget } from './common.js';
/**
 * @generated from protobuf enum proto.ShadowPriest.Rotation.RotationType
 */
export var ShadowPriest_Rotation_RotationType;
(function (ShadowPriest_Rotation_RotationType) {
    /**
     * @generated from protobuf enum value: Unknown = 0;
     */
    ShadowPriest_Rotation_RotationType[ShadowPriest_Rotation_RotationType["Unknown"] = 0] = "Unknown";
    /**
     * @generated from protobuf enum value: Basic = 1;
     */
    ShadowPriest_Rotation_RotationType[ShadowPriest_Rotation_RotationType["Basic"] = 1] = "Basic";
    /**
     * @generated from protobuf enum value: Clipping = 2;
     */
    ShadowPriest_Rotation_RotationType[ShadowPriest_Rotation_RotationType["Clipping"] = 2] = "Clipping";
    /**
     * @generated from protobuf enum value: Ideal = 3;
     */
    ShadowPriest_Rotation_RotationType[ShadowPriest_Rotation_RotationType["Ideal"] = 3] = "Ideal";
})(ShadowPriest_Rotation_RotationType || (ShadowPriest_Rotation_RotationType = {}));
/**
 * @generated from protobuf enum proto.SmitePriest.Rotation.RotationType
 */
export var SmitePriest_Rotation_RotationType;
(function (SmitePriest_Rotation_RotationType) {
    /**
     * @generated from protobuf enum value: Unknown = 0;
     */
    SmitePriest_Rotation_RotationType[SmitePriest_Rotation_RotationType["Unknown"] = 0] = "Unknown";
    /**
     * @generated from protobuf enum value: Basic = 1;
     */
    SmitePriest_Rotation_RotationType[SmitePriest_Rotation_RotationType["Basic"] = 1] = "Basic";
    /**
     * @generated from protobuf enum value: HolyFireWeave = 2;
     */
    SmitePriest_Rotation_RotationType[SmitePriest_Rotation_RotationType["HolyFireWeave"] = 2] = "HolyFireWeave";
})(SmitePriest_Rotation_RotationType || (SmitePriest_Rotation_RotationType = {}));
// @generated message type with reflection information, may provide speed optimized methods
class PriestTalents$Type extends MessageType {
    constructor() {
        super("proto.PriestTalents", [
            { no: 1, name: "wand_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 31, name: "silent_resolve", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 34, name: "improved_power_word_fortitude", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "inner_focus", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "meditation", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "mental_agility", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "mental_strength", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "divine_spirit", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 7, name: "improved_divine_spirit", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 8, name: "focused_power", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 9, name: "force_of_will", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 10, name: "power_infusion", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 11, name: "enlightenment", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 12, name: "holy_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 13, name: "divine_fury", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 14, name: "holy_nova", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 15, name: "searing_light", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 16, name: "spiritual_guidance", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 17, name: "surge_of_light", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 33, name: "spirit_of_redemption", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 32, name: "shadow_affinity", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 18, name: "improved_shadow_word_pain", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 19, name: "shadow_focus", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 20, name: "improved_mind_blast", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 21, name: "mind_flay", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 22, name: "shadow_weaving", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 23, name: "vampiric_embrace", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 24, name: "improved_vampiric_embrace", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 25, name: "focused_mind", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 26, name: "darkness", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 27, name: "shadowform", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 28, name: "shadow_power", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 29, name: "misery", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 30, name: "vampiric_touch", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { wandSpecialization: 0, silentResolve: 0, improvedPowerWordFortitude: 0, innerFocus: false, meditation: 0, mentalAgility: 0, mentalStrength: 0, divineSpirit: false, improvedDivineSpirit: 0, focusedPower: 0, forceOfWill: 0, powerInfusion: false, enlightenment: 0, holySpecialization: 0, divineFury: 0, holyNova: false, searingLight: 0, spiritualGuidance: 0, surgeOfLight: 0, spiritOfRedemption: false, shadowAffinity: 0, improvedShadowWordPain: 0, shadowFocus: 0, improvedMindBlast: 0, mindFlay: false, shadowWeaving: 0, vampiricEmbrace: false, improvedVampiricEmbrace: 0, focusedMind: 0, darkness: 0, shadowform: false, shadowPower: 0, misery: 0, vampiricTouch: false };
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
                case /* int32 wand_specialization */ 1:
                    message.wandSpecialization = reader.int32();
                    break;
                case /* int32 silent_resolve */ 31:
                    message.silentResolve = reader.int32();
                    break;
                case /* int32 improved_power_word_fortitude */ 34:
                    message.improvedPowerWordFortitude = reader.int32();
                    break;
                case /* bool inner_focus */ 2:
                    message.innerFocus = reader.bool();
                    break;
                case /* int32 meditation */ 3:
                    message.meditation = reader.int32();
                    break;
                case /* int32 mental_agility */ 4:
                    message.mentalAgility = reader.int32();
                    break;
                case /* int32 mental_strength */ 5:
                    message.mentalStrength = reader.int32();
                    break;
                case /* bool divine_spirit */ 6:
                    message.divineSpirit = reader.bool();
                    break;
                case /* int32 improved_divine_spirit */ 7:
                    message.improvedDivineSpirit = reader.int32();
                    break;
                case /* int32 focused_power */ 8:
                    message.focusedPower = reader.int32();
                    break;
                case /* int32 force_of_will */ 9:
                    message.forceOfWill = reader.int32();
                    break;
                case /* bool power_infusion */ 10:
                    message.powerInfusion = reader.bool();
                    break;
                case /* int32 enlightenment */ 11:
                    message.enlightenment = reader.int32();
                    break;
                case /* int32 holy_specialization */ 12:
                    message.holySpecialization = reader.int32();
                    break;
                case /* int32 divine_fury */ 13:
                    message.divineFury = reader.int32();
                    break;
                case /* bool holy_nova */ 14:
                    message.holyNova = reader.bool();
                    break;
                case /* int32 searing_light */ 15:
                    message.searingLight = reader.int32();
                    break;
                case /* int32 spiritual_guidance */ 16:
                    message.spiritualGuidance = reader.int32();
                    break;
                case /* int32 surge_of_light */ 17:
                    message.surgeOfLight = reader.int32();
                    break;
                case /* bool spirit_of_redemption */ 33:
                    message.spiritOfRedemption = reader.bool();
                    break;
                case /* int32 shadow_affinity */ 32:
                    message.shadowAffinity = reader.int32();
                    break;
                case /* int32 improved_shadow_word_pain */ 18:
                    message.improvedShadowWordPain = reader.int32();
                    break;
                case /* int32 shadow_focus */ 19:
                    message.shadowFocus = reader.int32();
                    break;
                case /* int32 improved_mind_blast */ 20:
                    message.improvedMindBlast = reader.int32();
                    break;
                case /* bool mind_flay */ 21:
                    message.mindFlay = reader.bool();
                    break;
                case /* int32 shadow_weaving */ 22:
                    message.shadowWeaving = reader.int32();
                    break;
                case /* bool vampiric_embrace */ 23:
                    message.vampiricEmbrace = reader.bool();
                    break;
                case /* int32 improved_vampiric_embrace */ 24:
                    message.improvedVampiricEmbrace = reader.int32();
                    break;
                case /* int32 focused_mind */ 25:
                    message.focusedMind = reader.int32();
                    break;
                case /* int32 darkness */ 26:
                    message.darkness = reader.int32();
                    break;
                case /* bool shadowform */ 27:
                    message.shadowform = reader.bool();
                    break;
                case /* int32 shadow_power */ 28:
                    message.shadowPower = reader.int32();
                    break;
                case /* int32 misery */ 29:
                    message.misery = reader.int32();
                    break;
                case /* bool vampiric_touch */ 30:
                    message.vampiricTouch = reader.bool();
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
        /* int32 wand_specialization = 1; */
        if (message.wandSpecialization !== 0)
            writer.tag(1, WireType.Varint).int32(message.wandSpecialization);
        /* int32 silent_resolve = 31; */
        if (message.silentResolve !== 0)
            writer.tag(31, WireType.Varint).int32(message.silentResolve);
        /* int32 improved_power_word_fortitude = 34; */
        if (message.improvedPowerWordFortitude !== 0)
            writer.tag(34, WireType.Varint).int32(message.improvedPowerWordFortitude);
        /* bool inner_focus = 2; */
        if (message.innerFocus !== false)
            writer.tag(2, WireType.Varint).bool(message.innerFocus);
        /* int32 meditation = 3; */
        if (message.meditation !== 0)
            writer.tag(3, WireType.Varint).int32(message.meditation);
        /* int32 mental_agility = 4; */
        if (message.mentalAgility !== 0)
            writer.tag(4, WireType.Varint).int32(message.mentalAgility);
        /* int32 mental_strength = 5; */
        if (message.mentalStrength !== 0)
            writer.tag(5, WireType.Varint).int32(message.mentalStrength);
        /* bool divine_spirit = 6; */
        if (message.divineSpirit !== false)
            writer.tag(6, WireType.Varint).bool(message.divineSpirit);
        /* int32 improved_divine_spirit = 7; */
        if (message.improvedDivineSpirit !== 0)
            writer.tag(7, WireType.Varint).int32(message.improvedDivineSpirit);
        /* int32 focused_power = 8; */
        if (message.focusedPower !== 0)
            writer.tag(8, WireType.Varint).int32(message.focusedPower);
        /* int32 force_of_will = 9; */
        if (message.forceOfWill !== 0)
            writer.tag(9, WireType.Varint).int32(message.forceOfWill);
        /* bool power_infusion = 10; */
        if (message.powerInfusion !== false)
            writer.tag(10, WireType.Varint).bool(message.powerInfusion);
        /* int32 enlightenment = 11; */
        if (message.enlightenment !== 0)
            writer.tag(11, WireType.Varint).int32(message.enlightenment);
        /* int32 holy_specialization = 12; */
        if (message.holySpecialization !== 0)
            writer.tag(12, WireType.Varint).int32(message.holySpecialization);
        /* int32 divine_fury = 13; */
        if (message.divineFury !== 0)
            writer.tag(13, WireType.Varint).int32(message.divineFury);
        /* bool holy_nova = 14; */
        if (message.holyNova !== false)
            writer.tag(14, WireType.Varint).bool(message.holyNova);
        /* int32 searing_light = 15; */
        if (message.searingLight !== 0)
            writer.tag(15, WireType.Varint).int32(message.searingLight);
        /* int32 spiritual_guidance = 16; */
        if (message.spiritualGuidance !== 0)
            writer.tag(16, WireType.Varint).int32(message.spiritualGuidance);
        /* int32 surge_of_light = 17; */
        if (message.surgeOfLight !== 0)
            writer.tag(17, WireType.Varint).int32(message.surgeOfLight);
        /* bool spirit_of_redemption = 33; */
        if (message.spiritOfRedemption !== false)
            writer.tag(33, WireType.Varint).bool(message.spiritOfRedemption);
        /* int32 shadow_affinity = 32; */
        if (message.shadowAffinity !== 0)
            writer.tag(32, WireType.Varint).int32(message.shadowAffinity);
        /* int32 improved_shadow_word_pain = 18; */
        if (message.improvedShadowWordPain !== 0)
            writer.tag(18, WireType.Varint).int32(message.improvedShadowWordPain);
        /* int32 shadow_focus = 19; */
        if (message.shadowFocus !== 0)
            writer.tag(19, WireType.Varint).int32(message.shadowFocus);
        /* int32 improved_mind_blast = 20; */
        if (message.improvedMindBlast !== 0)
            writer.tag(20, WireType.Varint).int32(message.improvedMindBlast);
        /* bool mind_flay = 21; */
        if (message.mindFlay !== false)
            writer.tag(21, WireType.Varint).bool(message.mindFlay);
        /* int32 shadow_weaving = 22; */
        if (message.shadowWeaving !== 0)
            writer.tag(22, WireType.Varint).int32(message.shadowWeaving);
        /* bool vampiric_embrace = 23; */
        if (message.vampiricEmbrace !== false)
            writer.tag(23, WireType.Varint).bool(message.vampiricEmbrace);
        /* int32 improved_vampiric_embrace = 24; */
        if (message.improvedVampiricEmbrace !== 0)
            writer.tag(24, WireType.Varint).int32(message.improvedVampiricEmbrace);
        /* int32 focused_mind = 25; */
        if (message.focusedMind !== 0)
            writer.tag(25, WireType.Varint).int32(message.focusedMind);
        /* int32 darkness = 26; */
        if (message.darkness !== 0)
            writer.tag(26, WireType.Varint).int32(message.darkness);
        /* bool shadowform = 27; */
        if (message.shadowform !== false)
            writer.tag(27, WireType.Varint).bool(message.shadowform);
        /* int32 shadow_power = 28; */
        if (message.shadowPower !== 0)
            writer.tag(28, WireType.Varint).int32(message.shadowPower);
        /* int32 misery = 29; */
        if (message.misery !== 0)
            writer.tag(29, WireType.Varint).int32(message.misery);
        /* bool vampiric_touch = 30; */
        if (message.vampiricTouch !== false)
            writer.tag(30, WireType.Varint).bool(message.vampiricTouch);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.PriestTalents
 */
export const PriestTalents = new PriestTalents$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ShadowPriest$Type extends MessageType {
    constructor() {
        super("proto.ShadowPriest", [
            { no: 1, name: "rotation", kind: "message", T: () => ShadowPriest_Rotation },
            { no: 2, name: "talents", kind: "message", T: () => PriestTalents },
            { no: 3, name: "options", kind: "message", T: () => ShadowPriest_Options }
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
                case /* proto.ShadowPriest.Rotation rotation */ 1:
                    message.rotation = ShadowPriest_Rotation.internalBinaryRead(reader, reader.uint32(), options, message.rotation);
                    break;
                case /* proto.PriestTalents talents */ 2:
                    message.talents = PriestTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* proto.ShadowPriest.Options options */ 3:
                    message.options = ShadowPriest_Options.internalBinaryRead(reader, reader.uint32(), options, message.options);
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
        /* proto.ShadowPriest.Rotation rotation = 1; */
        if (message.rotation)
            ShadowPriest_Rotation.internalBinaryWrite(message.rotation, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.PriestTalents talents = 2; */
        if (message.talents)
            PriestTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.ShadowPriest.Options options = 3; */
        if (message.options)
            ShadowPriest_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ShadowPriest
 */
export const ShadowPriest = new ShadowPriest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ShadowPriest_Rotation$Type extends MessageType {
    constructor() {
        super("proto.ShadowPriest.Rotation", [
            { no: 1, name: "rotation_type", kind: "enum", T: () => ["proto.ShadowPriest.Rotation.RotationType", ShadowPriest_Rotation_RotationType] },
            { no: 3, name: "use_dev_plague", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "precast_vt", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "latency", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 6, name: "use_starshards", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { rotationType: 0, useDevPlague: false, precastVt: false, latency: 0, useStarshards: false };
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
                case /* proto.ShadowPriest.Rotation.RotationType rotation_type */ 1:
                    message.rotationType = reader.int32();
                    break;
                case /* bool use_dev_plague */ 3:
                    message.useDevPlague = reader.bool();
                    break;
                case /* bool precast_vt */ 4:
                    message.precastVt = reader.bool();
                    break;
                case /* double latency */ 5:
                    message.latency = reader.double();
                    break;
                case /* bool use_starshards */ 6:
                    message.useStarshards = reader.bool();
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
        /* proto.ShadowPriest.Rotation.RotationType rotation_type = 1; */
        if (message.rotationType !== 0)
            writer.tag(1, WireType.Varint).int32(message.rotationType);
        /* bool use_dev_plague = 3; */
        if (message.useDevPlague !== false)
            writer.tag(3, WireType.Varint).bool(message.useDevPlague);
        /* bool precast_vt = 4; */
        if (message.precastVt !== false)
            writer.tag(4, WireType.Varint).bool(message.precastVt);
        /* double latency = 5; */
        if (message.latency !== 0)
            writer.tag(5, WireType.Bit64).double(message.latency);
        /* bool use_starshards = 6; */
        if (message.useStarshards !== false)
            writer.tag(6, WireType.Varint).bool(message.useStarshards);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ShadowPriest.Rotation
 */
export const ShadowPriest_Rotation = new ShadowPriest_Rotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ShadowPriest_Options$Type extends MessageType {
    constructor() {
        super("proto.ShadowPriest.Options", [
            { no: 1, name: "use_shadowfiend", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { useShadowfiend: false };
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
                case /* bool use_shadowfiend */ 1:
                    message.useShadowfiend = reader.bool();
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
        /* bool use_shadowfiend = 1; */
        if (message.useShadowfiend !== false)
            writer.tag(1, WireType.Varint).bool(message.useShadowfiend);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ShadowPriest.Options
 */
export const ShadowPriest_Options = new ShadowPriest_Options$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SmitePriest$Type extends MessageType {
    constructor() {
        super("proto.SmitePriest", [
            { no: 1, name: "rotation", kind: "message", T: () => SmitePriest_Rotation },
            { no: 2, name: "talents", kind: "message", T: () => PriestTalents },
            { no: 3, name: "options", kind: "message", T: () => SmitePriest_Options }
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
                case /* proto.SmitePriest.Rotation rotation */ 1:
                    message.rotation = SmitePriest_Rotation.internalBinaryRead(reader, reader.uint32(), options, message.rotation);
                    break;
                case /* proto.PriestTalents talents */ 2:
                    message.talents = PriestTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* proto.SmitePriest.Options options */ 3:
                    message.options = SmitePriest_Options.internalBinaryRead(reader, reader.uint32(), options, message.options);
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
        /* proto.SmitePriest.Rotation rotation = 1; */
        if (message.rotation)
            SmitePriest_Rotation.internalBinaryWrite(message.rotation, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.PriestTalents talents = 2; */
        if (message.talents)
            PriestTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.SmitePriest.Options options = 3; */
        if (message.options)
            SmitePriest_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.SmitePriest
 */
export const SmitePriest = new SmitePriest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SmitePriest_Rotation$Type extends MessageType {
    constructor() {
        super("proto.SmitePriest.Rotation", [
            { no: 1, name: "rotation_type", kind: "enum", T: () => ["proto.SmitePriest.Rotation.RotationType", SmitePriest_Rotation_RotationType] },
            { no: 3, name: "use_dev_plague", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "use_starshards", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "use_mind_blast", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 6, name: "use_shadow_word_death", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { rotationType: 0, useDevPlague: false, useStarshards: false, useMindBlast: false, useShadowWordDeath: false };
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
                case /* proto.SmitePriest.Rotation.RotationType rotation_type */ 1:
                    message.rotationType = reader.int32();
                    break;
                case /* bool use_dev_plague */ 3:
                    message.useDevPlague = reader.bool();
                    break;
                case /* bool use_starshards */ 4:
                    message.useStarshards = reader.bool();
                    break;
                case /* bool use_mind_blast */ 5:
                    message.useMindBlast = reader.bool();
                    break;
                case /* bool use_shadow_word_death */ 6:
                    message.useShadowWordDeath = reader.bool();
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
        /* proto.SmitePriest.Rotation.RotationType rotation_type = 1; */
        if (message.rotationType !== 0)
            writer.tag(1, WireType.Varint).int32(message.rotationType);
        /* bool use_dev_plague = 3; */
        if (message.useDevPlague !== false)
            writer.tag(3, WireType.Varint).bool(message.useDevPlague);
        /* bool use_starshards = 4; */
        if (message.useStarshards !== false)
            writer.tag(4, WireType.Varint).bool(message.useStarshards);
        /* bool use_mind_blast = 5; */
        if (message.useMindBlast !== false)
            writer.tag(5, WireType.Varint).bool(message.useMindBlast);
        /* bool use_shadow_word_death = 6; */
        if (message.useShadowWordDeath !== false)
            writer.tag(6, WireType.Varint).bool(message.useShadowWordDeath);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.SmitePriest.Rotation
 */
export const SmitePriest_Rotation = new SmitePriest_Rotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SmitePriest_Options$Type extends MessageType {
    constructor() {
        super("proto.SmitePriest.Options", [
            { no: 1, name: "use_shadowfiend", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "power_infusion_target", kind: "message", T: () => RaidTarget }
        ]);
    }
    create(value) {
        const message = { useShadowfiend: false };
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
                case /* bool use_shadowfiend */ 1:
                    message.useShadowfiend = reader.bool();
                    break;
                case /* proto.RaidTarget power_infusion_target */ 2:
                    message.powerInfusionTarget = RaidTarget.internalBinaryRead(reader, reader.uint32(), options, message.powerInfusionTarget);
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
        /* bool use_shadowfiend = 1; */
        if (message.useShadowfiend !== false)
            writer.tag(1, WireType.Varint).bool(message.useShadowfiend);
        /* proto.RaidTarget power_infusion_target = 2; */
        if (message.powerInfusionTarget)
            RaidTarget.internalBinaryWrite(message.powerInfusionTarget, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.SmitePriest.Options
 */
export const SmitePriest_Options = new SmitePriest_Options$Type();
