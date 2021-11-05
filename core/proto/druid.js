import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
// @generated message type with reflection information, may provide speed optimized methods
class DruidTalents$Type extends MessageType {
    constructor() {
        super("proto.DruidTalents", [
            { no: 1, name: "starlight_wrath", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "focused_starlight", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "improved_moonfire", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "brambles", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "insect_swarm", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 6, name: "vengeance", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 7, name: "lunar_guidance", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 8, name: "natures_grace", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 9, name: "moonglow", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 10, name: "moonfury", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 11, name: "balance_of_power", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 12, name: "dreamstate", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 13, name: "moonkin_form", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 14, name: "improved_faerie_fire", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 15, name: "wrath_of_cenarius", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 16, name: "force_of_nature", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 17, name: "ferocity", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 18, name: "feral_aggresion", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 19, name: "sharpened_claws", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 20, name: "shredding_attacks", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 21, name: "predatory_strikes", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 22, name: "primal_fury", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 23, name: "savage_fury", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 24, name: "faerie_fire", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 25, name: "heart_of_the_wild", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 26, name: "survival_of_the_fittest", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 27, name: "leader_of_the_pack", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 28, name: "improved_leader_of_the_pack", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 29, name: "predatory_instincts", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 30, name: "mangle", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 31, name: "improved_mark_of_the_wild", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 32, name: "furor", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 33, name: "naturalist", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 34, name: "natural_shapeshifter", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 35, name: "intensity", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 36, name: "omen_of_clarity", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 37, name: "natures_swiftness", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 38, name: "living_spirit", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 39, name: "natural_perfection", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value) {
        const message = { starlightWrath: 0, focusedStarlight: 0, improvedMoonfire: 0, brambles: 0, insectSwarm: false, vengeance: 0, lunarGuidance: 0, naturesGrace: false, moonglow: 0, moonfury: 0, balanceOfPower: 0, dreamstate: 0, moonkinForm: false, improvedFaerieFire: 0, wrathOfCenarius: 0, forceOfNature: false, ferocity: 0, feralAggresion: 0, sharpenedClaws: 0, shreddingAttacks: 0, predatoryStrikes: 0, primalFury: 0, savageFury: 0, faerieFire: false, heartOfTheWild: 0, survivalOfTheFittest: 0, leaderOfThePack: 0, improvedLeaderOfThePack: 0, predatoryInstincts: 0, mangle: false, improvedMarkOfTheWild: 0, furor: 0, naturalist: 0, naturalShapeshifter: 0, intensity: 0, omenOfClarity: false, naturesSwiftness: false, livingSpirit: 0, naturalPerfection: 0 };
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
                case /* int32 starlight_wrath */ 1:
                    message.starlightWrath = reader.int32();
                    break;
                case /* int32 focused_starlight */ 2:
                    message.focusedStarlight = reader.int32();
                    break;
                case /* int32 improved_moonfire */ 3:
                    message.improvedMoonfire = reader.int32();
                    break;
                case /* int32 brambles */ 4:
                    message.brambles = reader.int32();
                    break;
                case /* bool insect_swarm */ 5:
                    message.insectSwarm = reader.bool();
                    break;
                case /* int32 vengeance */ 6:
                    message.vengeance = reader.int32();
                    break;
                case /* int32 lunar_guidance */ 7:
                    message.lunarGuidance = reader.int32();
                    break;
                case /* bool natures_grace */ 8:
                    message.naturesGrace = reader.bool();
                    break;
                case /* int32 moonglow */ 9:
                    message.moonglow = reader.int32();
                    break;
                case /* int32 moonfury */ 10:
                    message.moonfury = reader.int32();
                    break;
                case /* int32 balance_of_power */ 11:
                    message.balanceOfPower = reader.int32();
                    break;
                case /* int32 dreamstate */ 12:
                    message.dreamstate = reader.int32();
                    break;
                case /* bool moonkin_form */ 13:
                    message.moonkinForm = reader.bool();
                    break;
                case /* int32 improved_faerie_fire */ 14:
                    message.improvedFaerieFire = reader.int32();
                    break;
                case /* int32 wrath_of_cenarius */ 15:
                    message.wrathOfCenarius = reader.int32();
                    break;
                case /* bool force_of_nature */ 16:
                    message.forceOfNature = reader.bool();
                    break;
                case /* int32 ferocity */ 17:
                    message.ferocity = reader.int32();
                    break;
                case /* int32 feral_aggresion */ 18:
                    message.feralAggresion = reader.int32();
                    break;
                case /* int32 sharpened_claws */ 19:
                    message.sharpenedClaws = reader.int32();
                    break;
                case /* int32 shredding_attacks */ 20:
                    message.shreddingAttacks = reader.int32();
                    break;
                case /* int32 predatory_strikes */ 21:
                    message.predatoryStrikes = reader.int32();
                    break;
                case /* int32 primal_fury */ 22:
                    message.primalFury = reader.int32();
                    break;
                case /* int32 savage_fury */ 23:
                    message.savageFury = reader.int32();
                    break;
                case /* bool faerie_fire */ 24:
                    message.faerieFire = reader.bool();
                    break;
                case /* int32 heart_of_the_wild */ 25:
                    message.heartOfTheWild = reader.int32();
                    break;
                case /* int32 survival_of_the_fittest */ 26:
                    message.survivalOfTheFittest = reader.int32();
                    break;
                case /* int32 leader_of_the_pack */ 27:
                    message.leaderOfThePack = reader.int32();
                    break;
                case /* int32 improved_leader_of_the_pack */ 28:
                    message.improvedLeaderOfThePack = reader.int32();
                    break;
                case /* int32 predatory_instincts */ 29:
                    message.predatoryInstincts = reader.int32();
                    break;
                case /* bool mangle */ 30:
                    message.mangle = reader.bool();
                    break;
                case /* int32 improved_mark_of_the_wild */ 31:
                    message.improvedMarkOfTheWild = reader.int32();
                    break;
                case /* int32 furor */ 32:
                    message.furor = reader.int32();
                    break;
                case /* int32 naturalist */ 33:
                    message.naturalist = reader.int32();
                    break;
                case /* int32 natural_shapeshifter */ 34:
                    message.naturalShapeshifter = reader.int32();
                    break;
                case /* int32 intensity */ 35:
                    message.intensity = reader.int32();
                    break;
                case /* bool omen_of_clarity */ 36:
                    message.omenOfClarity = reader.bool();
                    break;
                case /* bool natures_swiftness */ 37:
                    message.naturesSwiftness = reader.bool();
                    break;
                case /* int32 living_spirit */ 38:
                    message.livingSpirit = reader.int32();
                    break;
                case /* int32 natural_perfection */ 39:
                    message.naturalPerfection = reader.int32();
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
        /* int32 starlight_wrath = 1; */
        if (message.starlightWrath !== 0)
            writer.tag(1, WireType.Varint).int32(message.starlightWrath);
        /* int32 focused_starlight = 2; */
        if (message.focusedStarlight !== 0)
            writer.tag(2, WireType.Varint).int32(message.focusedStarlight);
        /* int32 improved_moonfire = 3; */
        if (message.improvedMoonfire !== 0)
            writer.tag(3, WireType.Varint).int32(message.improvedMoonfire);
        /* int32 brambles = 4; */
        if (message.brambles !== 0)
            writer.tag(4, WireType.Varint).int32(message.brambles);
        /* bool insect_swarm = 5; */
        if (message.insectSwarm !== false)
            writer.tag(5, WireType.Varint).bool(message.insectSwarm);
        /* int32 vengeance = 6; */
        if (message.vengeance !== 0)
            writer.tag(6, WireType.Varint).int32(message.vengeance);
        /* int32 lunar_guidance = 7; */
        if (message.lunarGuidance !== 0)
            writer.tag(7, WireType.Varint).int32(message.lunarGuidance);
        /* bool natures_grace = 8; */
        if (message.naturesGrace !== false)
            writer.tag(8, WireType.Varint).bool(message.naturesGrace);
        /* int32 moonglow = 9; */
        if (message.moonglow !== 0)
            writer.tag(9, WireType.Varint).int32(message.moonglow);
        /* int32 moonfury = 10; */
        if (message.moonfury !== 0)
            writer.tag(10, WireType.Varint).int32(message.moonfury);
        /* int32 balance_of_power = 11; */
        if (message.balanceOfPower !== 0)
            writer.tag(11, WireType.Varint).int32(message.balanceOfPower);
        /* int32 dreamstate = 12; */
        if (message.dreamstate !== 0)
            writer.tag(12, WireType.Varint).int32(message.dreamstate);
        /* bool moonkin_form = 13; */
        if (message.moonkinForm !== false)
            writer.tag(13, WireType.Varint).bool(message.moonkinForm);
        /* int32 improved_faerie_fire = 14; */
        if (message.improvedFaerieFire !== 0)
            writer.tag(14, WireType.Varint).int32(message.improvedFaerieFire);
        /* int32 wrath_of_cenarius = 15; */
        if (message.wrathOfCenarius !== 0)
            writer.tag(15, WireType.Varint).int32(message.wrathOfCenarius);
        /* bool force_of_nature = 16; */
        if (message.forceOfNature !== false)
            writer.tag(16, WireType.Varint).bool(message.forceOfNature);
        /* int32 ferocity = 17; */
        if (message.ferocity !== 0)
            writer.tag(17, WireType.Varint).int32(message.ferocity);
        /* int32 feral_aggresion = 18; */
        if (message.feralAggresion !== 0)
            writer.tag(18, WireType.Varint).int32(message.feralAggresion);
        /* int32 sharpened_claws = 19; */
        if (message.sharpenedClaws !== 0)
            writer.tag(19, WireType.Varint).int32(message.sharpenedClaws);
        /* int32 shredding_attacks = 20; */
        if (message.shreddingAttacks !== 0)
            writer.tag(20, WireType.Varint).int32(message.shreddingAttacks);
        /* int32 predatory_strikes = 21; */
        if (message.predatoryStrikes !== 0)
            writer.tag(21, WireType.Varint).int32(message.predatoryStrikes);
        /* int32 primal_fury = 22; */
        if (message.primalFury !== 0)
            writer.tag(22, WireType.Varint).int32(message.primalFury);
        /* int32 savage_fury = 23; */
        if (message.savageFury !== 0)
            writer.tag(23, WireType.Varint).int32(message.savageFury);
        /* bool faerie_fire = 24; */
        if (message.faerieFire !== false)
            writer.tag(24, WireType.Varint).bool(message.faerieFire);
        /* int32 heart_of_the_wild = 25; */
        if (message.heartOfTheWild !== 0)
            writer.tag(25, WireType.Varint).int32(message.heartOfTheWild);
        /* int32 survival_of_the_fittest = 26; */
        if (message.survivalOfTheFittest !== 0)
            writer.tag(26, WireType.Varint).int32(message.survivalOfTheFittest);
        /* int32 leader_of_the_pack = 27; */
        if (message.leaderOfThePack !== 0)
            writer.tag(27, WireType.Varint).int32(message.leaderOfThePack);
        /* int32 improved_leader_of_the_pack = 28; */
        if (message.improvedLeaderOfThePack !== 0)
            writer.tag(28, WireType.Varint).int32(message.improvedLeaderOfThePack);
        /* int32 predatory_instincts = 29; */
        if (message.predatoryInstincts !== 0)
            writer.tag(29, WireType.Varint).int32(message.predatoryInstincts);
        /* bool mangle = 30; */
        if (message.mangle !== false)
            writer.tag(30, WireType.Varint).bool(message.mangle);
        /* int32 improved_mark_of_the_wild = 31; */
        if (message.improvedMarkOfTheWild !== 0)
            writer.tag(31, WireType.Varint).int32(message.improvedMarkOfTheWild);
        /* int32 furor = 32; */
        if (message.furor !== 0)
            writer.tag(32, WireType.Varint).int32(message.furor);
        /* int32 naturalist = 33; */
        if (message.naturalist !== 0)
            writer.tag(33, WireType.Varint).int32(message.naturalist);
        /* int32 natural_shapeshifter = 34; */
        if (message.naturalShapeshifter !== 0)
            writer.tag(34, WireType.Varint).int32(message.naturalShapeshifter);
        /* int32 intensity = 35; */
        if (message.intensity !== 0)
            writer.tag(35, WireType.Varint).int32(message.intensity);
        /* bool omen_of_clarity = 36; */
        if (message.omenOfClarity !== false)
            writer.tag(36, WireType.Varint).bool(message.omenOfClarity);
        /* bool natures_swiftness = 37; */
        if (message.naturesSwiftness !== false)
            writer.tag(37, WireType.Varint).bool(message.naturesSwiftness);
        /* int32 living_spirit = 38; */
        if (message.livingSpirit !== 0)
            writer.tag(38, WireType.Varint).int32(message.livingSpirit);
        /* int32 natural_perfection = 39; */
        if (message.naturalPerfection !== 0)
            writer.tag(39, WireType.Varint).int32(message.naturalPerfection);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.DruidTalents
 */
export const DruidTalents = new DruidTalents$Type();
// @generated message type with reflection information, may provide speed optimized methods
class BalanceDruid$Type extends MessageType {
    constructor() {
        super("proto.BalanceDruid", [
            { no: 1, name: "rotation", kind: "message", T: () => BalanceDruid_Rotation },
            { no: 2, name: "talents", kind: "message", T: () => DruidTalents },
            { no: 3, name: "options", kind: "message", T: () => BalanceDruid_Options }
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
                case /* proto.BalanceDruid.Rotation rotation */ 1:
                    message.rotation = BalanceDruid_Rotation.internalBinaryRead(reader, reader.uint32(), options, message.rotation);
                    break;
                case /* proto.DruidTalents talents */ 2:
                    message.talents = DruidTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* proto.BalanceDruid.Options options */ 3:
                    message.options = BalanceDruid_Options.internalBinaryRead(reader, reader.uint32(), options, message.options);
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
        /* proto.BalanceDruid.Rotation rotation = 1; */
        if (message.rotation)
            BalanceDruid_Rotation.internalBinaryWrite(message.rotation, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.DruidTalents talents = 2; */
        if (message.talents)
            DruidTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.BalanceDruid.Options options = 3; */
        if (message.options)
            BalanceDruid_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.BalanceDruid
 */
export const BalanceDruid = new BalanceDruid$Type();
// @generated message type with reflection information, may provide speed optimized methods
class BalanceDruid_Rotation$Type extends MessageType {
    constructor() {
        super("proto.BalanceDruid.Rotation", []);
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
 * @generated MessageType for protobuf message proto.BalanceDruid.Rotation
 */
export const BalanceDruid_Rotation = new BalanceDruid_Rotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class BalanceDruid_Options$Type extends MessageType {
    constructor() {
        super("proto.BalanceDruid.Options", []);
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
 * @generated MessageType for protobuf message proto.BalanceDruid.Options
 */
export const BalanceDruid_Options = new BalanceDruid_Options$Type();
