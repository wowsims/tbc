import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
/**
 * @generated from protobuf enum proto.ElementalShaman.Agent.AgentType
 */
export var ElementalShaman_Agent_AgentType;
(function (ElementalShaman_Agent_AgentType) {
    /**
     * @generated from protobuf enum value: Unknown = 0;
     */
    ElementalShaman_Agent_AgentType[ElementalShaman_Agent_AgentType["Unknown"] = 0] = "Unknown";
    /**
     * @generated from protobuf enum value: FixedLBCL = 1;
     */
    ElementalShaman_Agent_AgentType[ElementalShaman_Agent_AgentType["FixedLBCL"] = 1] = "FixedLBCL";
    /**
     * @generated from protobuf enum value: CLOnClearcast = 2;
     */
    ElementalShaman_Agent_AgentType[ElementalShaman_Agent_AgentType["CLOnClearcast"] = 2] = "CLOnClearcast";
    /**
     * @generated from protobuf enum value: Adaptive = 3;
     */
    ElementalShaman_Agent_AgentType[ElementalShaman_Agent_AgentType["Adaptive"] = 3] = "Adaptive";
    /**
     * @generated from protobuf enum value: CLOnCD = 4;
     */
    ElementalShaman_Agent_AgentType[ElementalShaman_Agent_AgentType["CLOnCD"] = 4] = "CLOnCD";
})(ElementalShaman_Agent_AgentType || (ElementalShaman_Agent_AgentType = {}));
// @generated message type with reflection information, may provide speed optimized methods
class ShamanTalents$Type extends MessageType {
    constructor() {
        super("proto.ShamanTalents", [
            { no: 1, name: "convection", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "concussion", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "call_of_flame", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "elemental_focus", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "reverberation", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "call_of_thunder", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 7, name: "improved_fire_totems", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 8, name: "elemental_devastation", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 9, name: "elemental_fury", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 10, name: "unrelenting_storm", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 11, name: "elemental_precision", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 12, name: "lightning_mastery", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 13, name: "elemental_mastery", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 14, name: "lightning_overload", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 15, name: "ancestral_knowledge", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 16, name: "thundering_strikes", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 17, name: "enhancing_totems", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 18, name: "shamanistic_focus", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 19, name: "flurry", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 20, name: "improved_weapon_totems", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 21, name: "elemental_weapons", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 22, name: "mental_quickness", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 23, name: "weapon_mastery", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 24, name: "dual_wield_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 25, name: "unleashed_rage", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 26, name: "totemic_focus", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 27, name: "natures_guidance", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 28, name: "restorative_totems", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 29, name: "tidal_mastery", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 30, name: "natures_swiftness", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 31, name: "mana_tide_totem", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 32, name: "natures_blessing", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value) {
        const message = { convection: 0, concussion: 0, callOfFlame: 0, elementalFocus: false, reverberation: 0, callOfThunder: 0, improvedFireTotems: 0, elementalDevastation: 0, elementalFury: false, unrelentingStorm: 0, elementalPrecision: 0, lightningMastery: 0, elementalMastery: false, lightningOverload: 0, ancestralKnowledge: 0, thunderingStrikes: 0, enhancingTotems: 0, shamanisticFocus: false, flurry: 0, improvedWeaponTotems: 0, elementalWeapons: 0, mentalQuickness: 0, weaponMastery: 0, dualWieldSpecialization: 0, unleashedRage: 0, totemicFocus: 0, naturesGuidance: 0, restorativeTotems: 0, tidalMastery: 0, naturesSwiftness: false, manaTideTotem: false, naturesBlessing: 0 };
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
                case /* int32 convection */ 1:
                    message.convection = reader.int32();
                    break;
                case /* int32 concussion */ 2:
                    message.concussion = reader.int32();
                    break;
                case /* int32 call_of_flame */ 3:
                    message.callOfFlame = reader.int32();
                    break;
                case /* bool elemental_focus */ 4:
                    message.elementalFocus = reader.bool();
                    break;
                case /* int32 reverberation */ 5:
                    message.reverberation = reader.int32();
                    break;
                case /* int32 call_of_thunder */ 6:
                    message.callOfThunder = reader.int32();
                    break;
                case /* int32 improved_fire_totems */ 7:
                    message.improvedFireTotems = reader.int32();
                    break;
                case /* int32 elemental_devastation */ 8:
                    message.elementalDevastation = reader.int32();
                    break;
                case /* bool elemental_fury */ 9:
                    message.elementalFury = reader.bool();
                    break;
                case /* int32 unrelenting_storm */ 10:
                    message.unrelentingStorm = reader.int32();
                    break;
                case /* int32 elemental_precision */ 11:
                    message.elementalPrecision = reader.int32();
                    break;
                case /* int32 lightning_mastery */ 12:
                    message.lightningMastery = reader.int32();
                    break;
                case /* bool elemental_mastery */ 13:
                    message.elementalMastery = reader.bool();
                    break;
                case /* int32 lightning_overload */ 14:
                    message.lightningOverload = reader.int32();
                    break;
                case /* int32 ancestral_knowledge */ 15:
                    message.ancestralKnowledge = reader.int32();
                    break;
                case /* int32 thundering_strikes */ 16:
                    message.thunderingStrikes = reader.int32();
                    break;
                case /* int32 enhancing_totems */ 17:
                    message.enhancingTotems = reader.int32();
                    break;
                case /* bool shamanistic_focus */ 18:
                    message.shamanisticFocus = reader.bool();
                    break;
                case /* int32 flurry */ 19:
                    message.flurry = reader.int32();
                    break;
                case /* int32 improved_weapon_totems */ 20:
                    message.improvedWeaponTotems = reader.int32();
                    break;
                case /* int32 elemental_weapons */ 21:
                    message.elementalWeapons = reader.int32();
                    break;
                case /* int32 mental_quickness */ 22:
                    message.mentalQuickness = reader.int32();
                    break;
                case /* int32 weapon_mastery */ 23:
                    message.weaponMastery = reader.int32();
                    break;
                case /* int32 dual_wield_specialization */ 24:
                    message.dualWieldSpecialization = reader.int32();
                    break;
                case /* int32 unleashed_rage */ 25:
                    message.unleashedRage = reader.int32();
                    break;
                case /* int32 totemic_focus */ 26:
                    message.totemicFocus = reader.int32();
                    break;
                case /* int32 natures_guidance */ 27:
                    message.naturesGuidance = reader.int32();
                    break;
                case /* int32 restorative_totems */ 28:
                    message.restorativeTotems = reader.int32();
                    break;
                case /* int32 tidal_mastery */ 29:
                    message.tidalMastery = reader.int32();
                    break;
                case /* bool natures_swiftness */ 30:
                    message.naturesSwiftness = reader.bool();
                    break;
                case /* bool mana_tide_totem */ 31:
                    message.manaTideTotem = reader.bool();
                    break;
                case /* int32 natures_blessing */ 32:
                    message.naturesBlessing = reader.int32();
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
        /* int32 convection = 1; */
        if (message.convection !== 0)
            writer.tag(1, WireType.Varint).int32(message.convection);
        /* int32 concussion = 2; */
        if (message.concussion !== 0)
            writer.tag(2, WireType.Varint).int32(message.concussion);
        /* int32 call_of_flame = 3; */
        if (message.callOfFlame !== 0)
            writer.tag(3, WireType.Varint).int32(message.callOfFlame);
        /* bool elemental_focus = 4; */
        if (message.elementalFocus !== false)
            writer.tag(4, WireType.Varint).bool(message.elementalFocus);
        /* int32 reverberation = 5; */
        if (message.reverberation !== 0)
            writer.tag(5, WireType.Varint).int32(message.reverberation);
        /* int32 call_of_thunder = 6; */
        if (message.callOfThunder !== 0)
            writer.tag(6, WireType.Varint).int32(message.callOfThunder);
        /* int32 improved_fire_totems = 7; */
        if (message.improvedFireTotems !== 0)
            writer.tag(7, WireType.Varint).int32(message.improvedFireTotems);
        /* int32 elemental_devastation = 8; */
        if (message.elementalDevastation !== 0)
            writer.tag(8, WireType.Varint).int32(message.elementalDevastation);
        /* bool elemental_fury = 9; */
        if (message.elementalFury !== false)
            writer.tag(9, WireType.Varint).bool(message.elementalFury);
        /* int32 unrelenting_storm = 10; */
        if (message.unrelentingStorm !== 0)
            writer.tag(10, WireType.Varint).int32(message.unrelentingStorm);
        /* int32 elemental_precision = 11; */
        if (message.elementalPrecision !== 0)
            writer.tag(11, WireType.Varint).int32(message.elementalPrecision);
        /* int32 lightning_mastery = 12; */
        if (message.lightningMastery !== 0)
            writer.tag(12, WireType.Varint).int32(message.lightningMastery);
        /* bool elemental_mastery = 13; */
        if (message.elementalMastery !== false)
            writer.tag(13, WireType.Varint).bool(message.elementalMastery);
        /* int32 lightning_overload = 14; */
        if (message.lightningOverload !== 0)
            writer.tag(14, WireType.Varint).int32(message.lightningOverload);
        /* int32 ancestral_knowledge = 15; */
        if (message.ancestralKnowledge !== 0)
            writer.tag(15, WireType.Varint).int32(message.ancestralKnowledge);
        /* int32 thundering_strikes = 16; */
        if (message.thunderingStrikes !== 0)
            writer.tag(16, WireType.Varint).int32(message.thunderingStrikes);
        /* int32 enhancing_totems = 17; */
        if (message.enhancingTotems !== 0)
            writer.tag(17, WireType.Varint).int32(message.enhancingTotems);
        /* bool shamanistic_focus = 18; */
        if (message.shamanisticFocus !== false)
            writer.tag(18, WireType.Varint).bool(message.shamanisticFocus);
        /* int32 flurry = 19; */
        if (message.flurry !== 0)
            writer.tag(19, WireType.Varint).int32(message.flurry);
        /* int32 improved_weapon_totems = 20; */
        if (message.improvedWeaponTotems !== 0)
            writer.tag(20, WireType.Varint).int32(message.improvedWeaponTotems);
        /* int32 elemental_weapons = 21; */
        if (message.elementalWeapons !== 0)
            writer.tag(21, WireType.Varint).int32(message.elementalWeapons);
        /* int32 mental_quickness = 22; */
        if (message.mentalQuickness !== 0)
            writer.tag(22, WireType.Varint).int32(message.mentalQuickness);
        /* int32 weapon_mastery = 23; */
        if (message.weaponMastery !== 0)
            writer.tag(23, WireType.Varint).int32(message.weaponMastery);
        /* int32 dual_wield_specialization = 24; */
        if (message.dualWieldSpecialization !== 0)
            writer.tag(24, WireType.Varint).int32(message.dualWieldSpecialization);
        /* int32 unleashed_rage = 25; */
        if (message.unleashedRage !== 0)
            writer.tag(25, WireType.Varint).int32(message.unleashedRage);
        /* int32 totemic_focus = 26; */
        if (message.totemicFocus !== 0)
            writer.tag(26, WireType.Varint).int32(message.totemicFocus);
        /* int32 natures_guidance = 27; */
        if (message.naturesGuidance !== 0)
            writer.tag(27, WireType.Varint).int32(message.naturesGuidance);
        /* int32 restorative_totems = 28; */
        if (message.restorativeTotems !== 0)
            writer.tag(28, WireType.Varint).int32(message.restorativeTotems);
        /* int32 tidal_mastery = 29; */
        if (message.tidalMastery !== 0)
            writer.tag(29, WireType.Varint).int32(message.tidalMastery);
        /* bool natures_swiftness = 30; */
        if (message.naturesSwiftness !== false)
            writer.tag(30, WireType.Varint).bool(message.naturesSwiftness);
        /* bool mana_tide_totem = 31; */
        if (message.manaTideTotem !== false)
            writer.tag(31, WireType.Varint).bool(message.manaTideTotem);
        /* int32 natures_blessing = 32; */
        if (message.naturesBlessing !== 0)
            writer.tag(32, WireType.Varint).int32(message.naturesBlessing);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ShamanTalents
 */
export const ShamanTalents = new ShamanTalents$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ElementalShaman$Type extends MessageType {
    constructor() {
        super("proto.ElementalShaman", [
            { no: 1, name: "agent", kind: "message", T: () => ElementalShaman_Agent },
            { no: 2, name: "talents", kind: "message", T: () => ShamanTalents },
            { no: 3, name: "options", kind: "message", T: () => ElementalShaman_Options }
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
                case /* proto.ElementalShaman.Agent agent */ 1:
                    message.agent = ElementalShaman_Agent.internalBinaryRead(reader, reader.uint32(), options, message.agent);
                    break;
                case /* proto.ShamanTalents talents */ 2:
                    message.talents = ShamanTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* proto.ElementalShaman.Options options */ 3:
                    message.options = ElementalShaman_Options.internalBinaryRead(reader, reader.uint32(), options, message.options);
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
        /* proto.ElementalShaman.Agent agent = 1; */
        if (message.agent)
            ElementalShaman_Agent.internalBinaryWrite(message.agent, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.ShamanTalents talents = 2; */
        if (message.talents)
            ShamanTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.ElementalShaman.Options options = 3; */
        if (message.options)
            ElementalShaman_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ElementalShaman
 */
export const ElementalShaman = new ElementalShaman$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ElementalShaman_Agent$Type extends MessageType {
    constructor() {
        super("proto.ElementalShaman.Agent", [
            { no: 1, name: "type", kind: "enum", T: () => ["proto.ElementalShaman.Agent.AgentType", ElementalShaman_Agent_AgentType] }
        ]);
    }
    create(value) {
        const message = { type: 0 };
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
                case /* proto.ElementalShaman.Agent.AgentType type */ 1:
                    message.type = reader.int32();
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
        /* proto.ElementalShaman.Agent.AgentType type = 1; */
        if (message.type !== 0)
            writer.tag(1, WireType.Varint).int32(message.type);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ElementalShaman.Agent
 */
export const ElementalShaman_Agent = new ElementalShaman_Agent$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ElementalShaman_Options$Type extends MessageType {
    constructor() {
        super("proto.ElementalShaman.Options", [
            { no: 1, name: "water_shield", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "bloodlust", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "mana_spring_totem", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "totem_of_wrath", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "wrath_of_air_totem", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { waterShield: false, bloodlust: false, manaSpringTotem: false, totemOfWrath: false, wrathOfAirTotem: false };
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
                case /* bool water_shield */ 1:
                    message.waterShield = reader.bool();
                    break;
                case /* bool bloodlust */ 2:
                    message.bloodlust = reader.bool();
                    break;
                case /* bool mana_spring_totem */ 3:
                    message.manaSpringTotem = reader.bool();
                    break;
                case /* bool totem_of_wrath */ 4:
                    message.totemOfWrath = reader.bool();
                    break;
                case /* bool wrath_of_air_totem */ 5:
                    message.wrathOfAirTotem = reader.bool();
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
        /* bool water_shield = 1; */
        if (message.waterShield !== false)
            writer.tag(1, WireType.Varint).bool(message.waterShield);
        /* bool bloodlust = 2; */
        if (message.bloodlust !== false)
            writer.tag(2, WireType.Varint).bool(message.bloodlust);
        /* bool mana_spring_totem = 3; */
        if (message.manaSpringTotem !== false)
            writer.tag(3, WireType.Varint).bool(message.manaSpringTotem);
        /* bool totem_of_wrath = 4; */
        if (message.totemOfWrath !== false)
            writer.tag(4, WireType.Varint).bool(message.totemOfWrath);
        /* bool wrath_of_air_totem = 5; */
        if (message.wrathOfAirTotem !== false)
            writer.tag(5, WireType.Varint).bool(message.wrathOfAirTotem);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ElementalShaman.Options
 */
export const ElementalShaman_Options = new ElementalShaman_Options$Type();
