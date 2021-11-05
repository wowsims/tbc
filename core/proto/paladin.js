import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
// @generated message type with reflection information, may provide speed optimized methods
class PaladinTalents$Type extends MessageType {
    constructor() {
        super("proto.PaladinTalents", [
            { no: 1, name: "divine_strength", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "divine_intellect", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "improved_seal_of_righteousness", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 34, name: "illumination", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "improved_blessing_of_wisdom", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "divine_favor", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 6, name: "purifying_power", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 7, name: "holy_power", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 8, name: "holy_shock", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 9, name: "holy_guidance", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 10, name: "divine_illumination", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 11, name: "precision", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 12, name: "blessing_of_kings", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 13, name: "reckoning", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 14, name: "sacred_duty", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 15, name: "one_handed_weapon_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 16, name: "combat_expertise", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 17, name: "avengers_shield", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 18, name: "improved_blessing_of_might", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 19, name: "benediction", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 20, name: "improved_judgement", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 21, name: "improved_seal_of_the_crusader", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 22, name: "vindication", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 23, name: "conviction", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 24, name: "seal_of_command", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 25, name: "crusade", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 26, name: "two_handed_weapon_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 27, name: "sanctity_aura", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 28, name: "improved_sanctity_aura", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 29, name: "vengeance", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 30, name: "sanctified_judgement", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 31, name: "sanctified_seals", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 32, name: "fanaticism", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 33, name: "crusader_strike", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { divineStrength: 0, divineIntellect: 0, improvedSealOfRighteousness: 0, illumination: 0, improvedBlessingOfWisdom: 0, divineFavor: false, purifyingPower: 0, holyPower: 0, holyShock: false, holyGuidance: 0, divineIllumination: false, precision: 0, blessingOfKings: false, reckoning: 0, sacredDuty: 0, oneHandedWeaponSpecialization: 0, combatExpertise: 0, avengersShield: false, improvedBlessingOfMight: 0, benediction: 0, improvedJudgement: 0, improvedSealOfTheCrusader: 0, vindication: 0, conviction: 0, sealOfCommand: false, crusade: 0, twoHandedWeaponSpecialization: 0, sanctityAura: false, improvedSanctityAura: 0, vengeance: 0, sanctifiedJudgement: 0, sanctifiedSeals: 0, fanaticism: 0, crusaderStrike: false };
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
                case /* int32 divine_strength */ 1:
                    message.divineStrength = reader.int32();
                    break;
                case /* int32 divine_intellect */ 2:
                    message.divineIntellect = reader.int32();
                    break;
                case /* int32 improved_seal_of_righteousness */ 3:
                    message.improvedSealOfRighteousness = reader.int32();
                    break;
                case /* int32 illumination */ 34:
                    message.illumination = reader.int32();
                    break;
                case /* int32 improved_blessing_of_wisdom */ 4:
                    message.improvedBlessingOfWisdom = reader.int32();
                    break;
                case /* bool divine_favor */ 5:
                    message.divineFavor = reader.bool();
                    break;
                case /* int32 purifying_power */ 6:
                    message.purifyingPower = reader.int32();
                    break;
                case /* int32 holy_power */ 7:
                    message.holyPower = reader.int32();
                    break;
                case /* bool holy_shock */ 8:
                    message.holyShock = reader.bool();
                    break;
                case /* int32 holy_guidance */ 9:
                    message.holyGuidance = reader.int32();
                    break;
                case /* bool divine_illumination */ 10:
                    message.divineIllumination = reader.bool();
                    break;
                case /* int32 precision */ 11:
                    message.precision = reader.int32();
                    break;
                case /* bool blessing_of_kings */ 12:
                    message.blessingOfKings = reader.bool();
                    break;
                case /* int32 reckoning */ 13:
                    message.reckoning = reader.int32();
                    break;
                case /* int32 sacred_duty */ 14:
                    message.sacredDuty = reader.int32();
                    break;
                case /* int32 one_handed_weapon_specialization */ 15:
                    message.oneHandedWeaponSpecialization = reader.int32();
                    break;
                case /* int32 combat_expertise */ 16:
                    message.combatExpertise = reader.int32();
                    break;
                case /* bool avengers_shield */ 17:
                    message.avengersShield = reader.bool();
                    break;
                case /* int32 improved_blessing_of_might */ 18:
                    message.improvedBlessingOfMight = reader.int32();
                    break;
                case /* int32 benediction */ 19:
                    message.benediction = reader.int32();
                    break;
                case /* int32 improved_judgement */ 20:
                    message.improvedJudgement = reader.int32();
                    break;
                case /* int32 improved_seal_of_the_crusader */ 21:
                    message.improvedSealOfTheCrusader = reader.int32();
                    break;
                case /* int32 vindication */ 22:
                    message.vindication = reader.int32();
                    break;
                case /* int32 conviction */ 23:
                    message.conviction = reader.int32();
                    break;
                case /* bool seal_of_command */ 24:
                    message.sealOfCommand = reader.bool();
                    break;
                case /* int32 crusade */ 25:
                    message.crusade = reader.int32();
                    break;
                case /* int32 two_handed_weapon_specialization */ 26:
                    message.twoHandedWeaponSpecialization = reader.int32();
                    break;
                case /* bool sanctity_aura */ 27:
                    message.sanctityAura = reader.bool();
                    break;
                case /* int32 improved_sanctity_aura */ 28:
                    message.improvedSanctityAura = reader.int32();
                    break;
                case /* int32 vengeance */ 29:
                    message.vengeance = reader.int32();
                    break;
                case /* int32 sanctified_judgement */ 30:
                    message.sanctifiedJudgement = reader.int32();
                    break;
                case /* int32 sanctified_seals */ 31:
                    message.sanctifiedSeals = reader.int32();
                    break;
                case /* int32 fanaticism */ 32:
                    message.fanaticism = reader.int32();
                    break;
                case /* bool crusader_strike */ 33:
                    message.crusaderStrike = reader.bool();
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
        /* int32 divine_strength = 1; */
        if (message.divineStrength !== 0)
            writer.tag(1, WireType.Varint).int32(message.divineStrength);
        /* int32 divine_intellect = 2; */
        if (message.divineIntellect !== 0)
            writer.tag(2, WireType.Varint).int32(message.divineIntellect);
        /* int32 improved_seal_of_righteousness = 3; */
        if (message.improvedSealOfRighteousness !== 0)
            writer.tag(3, WireType.Varint).int32(message.improvedSealOfRighteousness);
        /* int32 illumination = 34; */
        if (message.illumination !== 0)
            writer.tag(34, WireType.Varint).int32(message.illumination);
        /* int32 improved_blessing_of_wisdom = 4; */
        if (message.improvedBlessingOfWisdom !== 0)
            writer.tag(4, WireType.Varint).int32(message.improvedBlessingOfWisdom);
        /* bool divine_favor = 5; */
        if (message.divineFavor !== false)
            writer.tag(5, WireType.Varint).bool(message.divineFavor);
        /* int32 purifying_power = 6; */
        if (message.purifyingPower !== 0)
            writer.tag(6, WireType.Varint).int32(message.purifyingPower);
        /* int32 holy_power = 7; */
        if (message.holyPower !== 0)
            writer.tag(7, WireType.Varint).int32(message.holyPower);
        /* bool holy_shock = 8; */
        if (message.holyShock !== false)
            writer.tag(8, WireType.Varint).bool(message.holyShock);
        /* int32 holy_guidance = 9; */
        if (message.holyGuidance !== 0)
            writer.tag(9, WireType.Varint).int32(message.holyGuidance);
        /* bool divine_illumination = 10; */
        if (message.divineIllumination !== false)
            writer.tag(10, WireType.Varint).bool(message.divineIllumination);
        /* int32 precision = 11; */
        if (message.precision !== 0)
            writer.tag(11, WireType.Varint).int32(message.precision);
        /* bool blessing_of_kings = 12; */
        if (message.blessingOfKings !== false)
            writer.tag(12, WireType.Varint).bool(message.blessingOfKings);
        /* int32 reckoning = 13; */
        if (message.reckoning !== 0)
            writer.tag(13, WireType.Varint).int32(message.reckoning);
        /* int32 sacred_duty = 14; */
        if (message.sacredDuty !== 0)
            writer.tag(14, WireType.Varint).int32(message.sacredDuty);
        /* int32 one_handed_weapon_specialization = 15; */
        if (message.oneHandedWeaponSpecialization !== 0)
            writer.tag(15, WireType.Varint).int32(message.oneHandedWeaponSpecialization);
        /* int32 combat_expertise = 16; */
        if (message.combatExpertise !== 0)
            writer.tag(16, WireType.Varint).int32(message.combatExpertise);
        /* bool avengers_shield = 17; */
        if (message.avengersShield !== false)
            writer.tag(17, WireType.Varint).bool(message.avengersShield);
        /* int32 improved_blessing_of_might = 18; */
        if (message.improvedBlessingOfMight !== 0)
            writer.tag(18, WireType.Varint).int32(message.improvedBlessingOfMight);
        /* int32 benediction = 19; */
        if (message.benediction !== 0)
            writer.tag(19, WireType.Varint).int32(message.benediction);
        /* int32 improved_judgement = 20; */
        if (message.improvedJudgement !== 0)
            writer.tag(20, WireType.Varint).int32(message.improvedJudgement);
        /* int32 improved_seal_of_the_crusader = 21; */
        if (message.improvedSealOfTheCrusader !== 0)
            writer.tag(21, WireType.Varint).int32(message.improvedSealOfTheCrusader);
        /* int32 vindication = 22; */
        if (message.vindication !== 0)
            writer.tag(22, WireType.Varint).int32(message.vindication);
        /* int32 conviction = 23; */
        if (message.conviction !== 0)
            writer.tag(23, WireType.Varint).int32(message.conviction);
        /* bool seal_of_command = 24; */
        if (message.sealOfCommand !== false)
            writer.tag(24, WireType.Varint).bool(message.sealOfCommand);
        /* int32 crusade = 25; */
        if (message.crusade !== 0)
            writer.tag(25, WireType.Varint).int32(message.crusade);
        /* int32 two_handed_weapon_specialization = 26; */
        if (message.twoHandedWeaponSpecialization !== 0)
            writer.tag(26, WireType.Varint).int32(message.twoHandedWeaponSpecialization);
        /* bool sanctity_aura = 27; */
        if (message.sanctityAura !== false)
            writer.tag(27, WireType.Varint).bool(message.sanctityAura);
        /* int32 improved_sanctity_aura = 28; */
        if (message.improvedSanctityAura !== 0)
            writer.tag(28, WireType.Varint).int32(message.improvedSanctityAura);
        /* int32 vengeance = 29; */
        if (message.vengeance !== 0)
            writer.tag(29, WireType.Varint).int32(message.vengeance);
        /* int32 sanctified_judgement = 30; */
        if (message.sanctifiedJudgement !== 0)
            writer.tag(30, WireType.Varint).int32(message.sanctifiedJudgement);
        /* int32 sanctified_seals = 31; */
        if (message.sanctifiedSeals !== 0)
            writer.tag(31, WireType.Varint).int32(message.sanctifiedSeals);
        /* int32 fanaticism = 32; */
        if (message.fanaticism !== 0)
            writer.tag(32, WireType.Varint).int32(message.fanaticism);
        /* bool crusader_strike = 33; */
        if (message.crusaderStrike !== false)
            writer.tag(33, WireType.Varint).bool(message.crusaderStrike);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.PaladinTalents
 */
export const PaladinTalents = new PaladinTalents$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RetributionPaladin$Type extends MessageType {
    constructor() {
        super("proto.RetributionPaladin", [
            { no: 1, name: "rotation", kind: "message", T: () => RetributionPaladin_Rotation },
            { no: 2, name: "talents", kind: "message", T: () => PaladinTalents },
            { no: 3, name: "options", kind: "message", T: () => RetributionPaladin_Options }
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
                case /* proto.RetributionPaladin.Rotation rotation */ 1:
                    message.rotation = RetributionPaladin_Rotation.internalBinaryRead(reader, reader.uint32(), options, message.rotation);
                    break;
                case /* proto.PaladinTalents talents */ 2:
                    message.talents = PaladinTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* proto.RetributionPaladin.Options options */ 3:
                    message.options = RetributionPaladin_Options.internalBinaryRead(reader, reader.uint32(), options, message.options);
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
        /* proto.RetributionPaladin.Rotation rotation = 1; */
        if (message.rotation)
            RetributionPaladin_Rotation.internalBinaryWrite(message.rotation, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.PaladinTalents talents = 2; */
        if (message.talents)
            PaladinTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.RetributionPaladin.Options options = 3; */
        if (message.options)
            RetributionPaladin_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.RetributionPaladin
 */
export const RetributionPaladin = new RetributionPaladin$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RetributionPaladin_Rotation$Type extends MessageType {
    constructor() {
        super("proto.RetributionPaladin.Rotation", []);
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
 * @generated MessageType for protobuf message proto.RetributionPaladin.Rotation
 */
export const RetributionPaladin_Rotation = new RetributionPaladin_Rotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RetributionPaladin_Options$Type extends MessageType {
    constructor() {
        super("proto.RetributionPaladin.Options", []);
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
 * @generated MessageType for protobuf message proto.RetributionPaladin.Options
 */
export const RetributionPaladin_Options = new RetributionPaladin_Options$Type();
