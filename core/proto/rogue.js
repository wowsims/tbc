import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
/**
 * @generated from protobuf enum proto.Rogue.Rotation.Builder
 */
export var Rogue_Rotation_Builder;
(function (Rogue_Rotation_Builder) {
    /**
     * @generated from protobuf enum value: Unknown = 0;
     */
    Rogue_Rotation_Builder[Rogue_Rotation_Builder["Unknown"] = 0] = "Unknown";
    /**
     * @generated from protobuf enum value: Auto = 1;
     */
    Rogue_Rotation_Builder[Rogue_Rotation_Builder["Auto"] = 1] = "Auto";
    /**
     * @generated from protobuf enum value: SinisterStrike = 2;
     */
    Rogue_Rotation_Builder[Rogue_Rotation_Builder["SinisterStrike"] = 2] = "SinisterStrike";
    /**
     * @generated from protobuf enum value: Backstab = 3;
     */
    Rogue_Rotation_Builder[Rogue_Rotation_Builder["Backstab"] = 3] = "Backstab";
    /**
     * @generated from protobuf enum value: Hemorrhage = 4;
     */
    Rogue_Rotation_Builder[Rogue_Rotation_Builder["Hemorrhage"] = 4] = "Hemorrhage";
    /**
     * @generated from protobuf enum value: Mutilate = 5;
     */
    Rogue_Rotation_Builder[Rogue_Rotation_Builder["Mutilate"] = 5] = "Mutilate";
})(Rogue_Rotation_Builder || (Rogue_Rotation_Builder = {}));
// @generated message type with reflection information, may provide speed optimized methods
class RogueTalents$Type extends MessageType {
    constructor() {
        super("proto.RogueTalents", [
            { no: 1, name: "improved_eviscerate", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "malice", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "ruthlessness", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "murder", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "puncturing_wounds", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "relentless_strikes", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 7, name: "improved_expose_armor", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 8, name: "lethality", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 9, name: "vile_poisons", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 10, name: "improved_poisons", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 11, name: "cold_blood", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 12, name: "quick_recovery", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 13, name: "seal_fate", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 14, name: "master_poisoner", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 15, name: "vigor", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 16, name: "find_weakness", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 17, name: "mutilate", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 18, name: "improved_sinister_strike", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 19, name: "improved_slice_and_dice", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 20, name: "precision", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 21, name: "dagger_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 22, name: "dual_wield_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 23, name: "mace_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 24, name: "blade_flurry", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 25, name: "sword_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 26, name: "fist_weapon_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 27, name: "weapon_expertise", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 28, name: "aggression", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 29, name: "vitality", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 30, name: "adrenaline_rush", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 31, name: "combat_potency", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 32, name: "surprise_attacks", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 33, name: "opportunity", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 46, name: "sleight_of_hand", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 34, name: "initiative", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 35, name: "ghostly_strike", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 36, name: "improved_ambush", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 47, name: "elusiveness", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 37, name: "serrated_blades", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 38, name: "preparation", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 39, name: "dirty_deeds", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 40, name: "hemorrhage", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 41, name: "master_of_subtlety", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 42, name: "deadliness", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 43, name: "premeditation", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 44, name: "sinister_calling", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 45, name: "shadowstep", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { improvedEviscerate: 0, malice: 0, ruthlessness: 0, murder: 0, puncturingWounds: 0, relentlessStrikes: false, improvedExposeArmor: 0, lethality: 0, vilePoisons: 0, improvedPoisons: 0, coldBlood: false, quickRecovery: 0, sealFate: 0, masterPoisoner: 0, vigor: false, findWeakness: 0, mutilate: false, improvedSinisterStrike: 0, improvedSliceAndDice: 0, precision: 0, daggerSpecialization: 0, dualWieldSpecialization: 0, maceSpecialization: 0, bladeFlurry: false, swordSpecialization: 0, fistWeaponSpecialization: 0, weaponExpertise: 0, aggression: 0, vitality: 0, adrenalineRush: false, combatPotency: 0, surpriseAttacks: false, opportunity: 0, sleightOfHand: 0, initiative: 0, ghostlyStrike: false, improvedAmbush: 0, elusiveness: 0, serratedBlades: 0, preparation: false, dirtyDeeds: 0, hemorrhage: false, masterOfSubtlety: 0, deadliness: 0, premeditation: false, sinisterCalling: 0, shadowstep: false };
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
                case /* int32 improved_eviscerate */ 1:
                    message.improvedEviscerate = reader.int32();
                    break;
                case /* int32 malice */ 2:
                    message.malice = reader.int32();
                    break;
                case /* int32 ruthlessness */ 3:
                    message.ruthlessness = reader.int32();
                    break;
                case /* int32 murder */ 4:
                    message.murder = reader.int32();
                    break;
                case /* int32 puncturing_wounds */ 5:
                    message.puncturingWounds = reader.int32();
                    break;
                case /* bool relentless_strikes */ 6:
                    message.relentlessStrikes = reader.bool();
                    break;
                case /* int32 improved_expose_armor */ 7:
                    message.improvedExposeArmor = reader.int32();
                    break;
                case /* int32 lethality */ 8:
                    message.lethality = reader.int32();
                    break;
                case /* int32 vile_poisons */ 9:
                    message.vilePoisons = reader.int32();
                    break;
                case /* int32 improved_poisons */ 10:
                    message.improvedPoisons = reader.int32();
                    break;
                case /* bool cold_blood */ 11:
                    message.coldBlood = reader.bool();
                    break;
                case /* int32 quick_recovery */ 12:
                    message.quickRecovery = reader.int32();
                    break;
                case /* int32 seal_fate */ 13:
                    message.sealFate = reader.int32();
                    break;
                case /* int32 master_poisoner */ 14:
                    message.masterPoisoner = reader.int32();
                    break;
                case /* bool vigor */ 15:
                    message.vigor = reader.bool();
                    break;
                case /* int32 find_weakness */ 16:
                    message.findWeakness = reader.int32();
                    break;
                case /* bool mutilate */ 17:
                    message.mutilate = reader.bool();
                    break;
                case /* int32 improved_sinister_strike */ 18:
                    message.improvedSinisterStrike = reader.int32();
                    break;
                case /* int32 improved_slice_and_dice */ 19:
                    message.improvedSliceAndDice = reader.int32();
                    break;
                case /* int32 precision */ 20:
                    message.precision = reader.int32();
                    break;
                case /* int32 dagger_specialization */ 21:
                    message.daggerSpecialization = reader.int32();
                    break;
                case /* int32 dual_wield_specialization */ 22:
                    message.dualWieldSpecialization = reader.int32();
                    break;
                case /* int32 mace_specialization */ 23:
                    message.maceSpecialization = reader.int32();
                    break;
                case /* bool blade_flurry */ 24:
                    message.bladeFlurry = reader.bool();
                    break;
                case /* int32 sword_specialization */ 25:
                    message.swordSpecialization = reader.int32();
                    break;
                case /* int32 fist_weapon_specialization */ 26:
                    message.fistWeaponSpecialization = reader.int32();
                    break;
                case /* int32 weapon_expertise */ 27:
                    message.weaponExpertise = reader.int32();
                    break;
                case /* int32 aggression */ 28:
                    message.aggression = reader.int32();
                    break;
                case /* int32 vitality */ 29:
                    message.vitality = reader.int32();
                    break;
                case /* bool adrenaline_rush */ 30:
                    message.adrenalineRush = reader.bool();
                    break;
                case /* int32 combat_potency */ 31:
                    message.combatPotency = reader.int32();
                    break;
                case /* bool surprise_attacks */ 32:
                    message.surpriseAttacks = reader.bool();
                    break;
                case /* int32 opportunity */ 33:
                    message.opportunity = reader.int32();
                    break;
                case /* int32 sleight_of_hand */ 46:
                    message.sleightOfHand = reader.int32();
                    break;
                case /* int32 initiative */ 34:
                    message.initiative = reader.int32();
                    break;
                case /* bool ghostly_strike */ 35:
                    message.ghostlyStrike = reader.bool();
                    break;
                case /* int32 improved_ambush */ 36:
                    message.improvedAmbush = reader.int32();
                    break;
                case /* int32 elusiveness */ 47:
                    message.elusiveness = reader.int32();
                    break;
                case /* int32 serrated_blades */ 37:
                    message.serratedBlades = reader.int32();
                    break;
                case /* bool preparation */ 38:
                    message.preparation = reader.bool();
                    break;
                case /* int32 dirty_deeds */ 39:
                    message.dirtyDeeds = reader.int32();
                    break;
                case /* bool hemorrhage */ 40:
                    message.hemorrhage = reader.bool();
                    break;
                case /* int32 master_of_subtlety */ 41:
                    message.masterOfSubtlety = reader.int32();
                    break;
                case /* int32 deadliness */ 42:
                    message.deadliness = reader.int32();
                    break;
                case /* bool premeditation */ 43:
                    message.premeditation = reader.bool();
                    break;
                case /* int32 sinister_calling */ 44:
                    message.sinisterCalling = reader.int32();
                    break;
                case /* bool shadowstep */ 45:
                    message.shadowstep = reader.bool();
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
        /* int32 improved_eviscerate = 1; */
        if (message.improvedEviscerate !== 0)
            writer.tag(1, WireType.Varint).int32(message.improvedEviscerate);
        /* int32 malice = 2; */
        if (message.malice !== 0)
            writer.tag(2, WireType.Varint).int32(message.malice);
        /* int32 ruthlessness = 3; */
        if (message.ruthlessness !== 0)
            writer.tag(3, WireType.Varint).int32(message.ruthlessness);
        /* int32 murder = 4; */
        if (message.murder !== 0)
            writer.tag(4, WireType.Varint).int32(message.murder);
        /* int32 puncturing_wounds = 5; */
        if (message.puncturingWounds !== 0)
            writer.tag(5, WireType.Varint).int32(message.puncturingWounds);
        /* bool relentless_strikes = 6; */
        if (message.relentlessStrikes !== false)
            writer.tag(6, WireType.Varint).bool(message.relentlessStrikes);
        /* int32 improved_expose_armor = 7; */
        if (message.improvedExposeArmor !== 0)
            writer.tag(7, WireType.Varint).int32(message.improvedExposeArmor);
        /* int32 lethality = 8; */
        if (message.lethality !== 0)
            writer.tag(8, WireType.Varint).int32(message.lethality);
        /* int32 vile_poisons = 9; */
        if (message.vilePoisons !== 0)
            writer.tag(9, WireType.Varint).int32(message.vilePoisons);
        /* int32 improved_poisons = 10; */
        if (message.improvedPoisons !== 0)
            writer.tag(10, WireType.Varint).int32(message.improvedPoisons);
        /* bool cold_blood = 11; */
        if (message.coldBlood !== false)
            writer.tag(11, WireType.Varint).bool(message.coldBlood);
        /* int32 quick_recovery = 12; */
        if (message.quickRecovery !== 0)
            writer.tag(12, WireType.Varint).int32(message.quickRecovery);
        /* int32 seal_fate = 13; */
        if (message.sealFate !== 0)
            writer.tag(13, WireType.Varint).int32(message.sealFate);
        /* int32 master_poisoner = 14; */
        if (message.masterPoisoner !== 0)
            writer.tag(14, WireType.Varint).int32(message.masterPoisoner);
        /* bool vigor = 15; */
        if (message.vigor !== false)
            writer.tag(15, WireType.Varint).bool(message.vigor);
        /* int32 find_weakness = 16; */
        if (message.findWeakness !== 0)
            writer.tag(16, WireType.Varint).int32(message.findWeakness);
        /* bool mutilate = 17; */
        if (message.mutilate !== false)
            writer.tag(17, WireType.Varint).bool(message.mutilate);
        /* int32 improved_sinister_strike = 18; */
        if (message.improvedSinisterStrike !== 0)
            writer.tag(18, WireType.Varint).int32(message.improvedSinisterStrike);
        /* int32 improved_slice_and_dice = 19; */
        if (message.improvedSliceAndDice !== 0)
            writer.tag(19, WireType.Varint).int32(message.improvedSliceAndDice);
        /* int32 precision = 20; */
        if (message.precision !== 0)
            writer.tag(20, WireType.Varint).int32(message.precision);
        /* int32 dagger_specialization = 21; */
        if (message.daggerSpecialization !== 0)
            writer.tag(21, WireType.Varint).int32(message.daggerSpecialization);
        /* int32 dual_wield_specialization = 22; */
        if (message.dualWieldSpecialization !== 0)
            writer.tag(22, WireType.Varint).int32(message.dualWieldSpecialization);
        /* int32 mace_specialization = 23; */
        if (message.maceSpecialization !== 0)
            writer.tag(23, WireType.Varint).int32(message.maceSpecialization);
        /* bool blade_flurry = 24; */
        if (message.bladeFlurry !== false)
            writer.tag(24, WireType.Varint).bool(message.bladeFlurry);
        /* int32 sword_specialization = 25; */
        if (message.swordSpecialization !== 0)
            writer.tag(25, WireType.Varint).int32(message.swordSpecialization);
        /* int32 fist_weapon_specialization = 26; */
        if (message.fistWeaponSpecialization !== 0)
            writer.tag(26, WireType.Varint).int32(message.fistWeaponSpecialization);
        /* int32 weapon_expertise = 27; */
        if (message.weaponExpertise !== 0)
            writer.tag(27, WireType.Varint).int32(message.weaponExpertise);
        /* int32 aggression = 28; */
        if (message.aggression !== 0)
            writer.tag(28, WireType.Varint).int32(message.aggression);
        /* int32 vitality = 29; */
        if (message.vitality !== 0)
            writer.tag(29, WireType.Varint).int32(message.vitality);
        /* bool adrenaline_rush = 30; */
        if (message.adrenalineRush !== false)
            writer.tag(30, WireType.Varint).bool(message.adrenalineRush);
        /* int32 combat_potency = 31; */
        if (message.combatPotency !== 0)
            writer.tag(31, WireType.Varint).int32(message.combatPotency);
        /* bool surprise_attacks = 32; */
        if (message.surpriseAttacks !== false)
            writer.tag(32, WireType.Varint).bool(message.surpriseAttacks);
        /* int32 opportunity = 33; */
        if (message.opportunity !== 0)
            writer.tag(33, WireType.Varint).int32(message.opportunity);
        /* int32 sleight_of_hand = 46; */
        if (message.sleightOfHand !== 0)
            writer.tag(46, WireType.Varint).int32(message.sleightOfHand);
        /* int32 initiative = 34; */
        if (message.initiative !== 0)
            writer.tag(34, WireType.Varint).int32(message.initiative);
        /* bool ghostly_strike = 35; */
        if (message.ghostlyStrike !== false)
            writer.tag(35, WireType.Varint).bool(message.ghostlyStrike);
        /* int32 improved_ambush = 36; */
        if (message.improvedAmbush !== 0)
            writer.tag(36, WireType.Varint).int32(message.improvedAmbush);
        /* int32 elusiveness = 47; */
        if (message.elusiveness !== 0)
            writer.tag(47, WireType.Varint).int32(message.elusiveness);
        /* int32 serrated_blades = 37; */
        if (message.serratedBlades !== 0)
            writer.tag(37, WireType.Varint).int32(message.serratedBlades);
        /* bool preparation = 38; */
        if (message.preparation !== false)
            writer.tag(38, WireType.Varint).bool(message.preparation);
        /* int32 dirty_deeds = 39; */
        if (message.dirtyDeeds !== 0)
            writer.tag(39, WireType.Varint).int32(message.dirtyDeeds);
        /* bool hemorrhage = 40; */
        if (message.hemorrhage !== false)
            writer.tag(40, WireType.Varint).bool(message.hemorrhage);
        /* int32 master_of_subtlety = 41; */
        if (message.masterOfSubtlety !== 0)
            writer.tag(41, WireType.Varint).int32(message.masterOfSubtlety);
        /* int32 deadliness = 42; */
        if (message.deadliness !== 0)
            writer.tag(42, WireType.Varint).int32(message.deadliness);
        /* bool premeditation = 43; */
        if (message.premeditation !== false)
            writer.tag(43, WireType.Varint).bool(message.premeditation);
        /* int32 sinister_calling = 44; */
        if (message.sinisterCalling !== 0)
            writer.tag(44, WireType.Varint).int32(message.sinisterCalling);
        /* bool shadowstep = 45; */
        if (message.shadowstep !== false)
            writer.tag(45, WireType.Varint).bool(message.shadowstep);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.RogueTalents
 */
export const RogueTalents = new RogueTalents$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Rogue$Type extends MessageType {
    constructor() {
        super("proto.Rogue", [
            { no: 1, name: "rotation", kind: "message", T: () => Rogue_Rotation },
            { no: 2, name: "talents", kind: "message", T: () => RogueTalents },
            { no: 3, name: "options", kind: "message", T: () => Rogue_Options }
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
                case /* proto.Rogue.Rotation rotation */ 1:
                    message.rotation = Rogue_Rotation.internalBinaryRead(reader, reader.uint32(), options, message.rotation);
                    break;
                case /* proto.RogueTalents talents */ 2:
                    message.talents = RogueTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* proto.Rogue.Options options */ 3:
                    message.options = Rogue_Options.internalBinaryRead(reader, reader.uint32(), options, message.options);
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
        /* proto.Rogue.Rotation rotation = 1; */
        if (message.rotation)
            Rogue_Rotation.internalBinaryWrite(message.rotation, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.RogueTalents talents = 2; */
        if (message.talents)
            RogueTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.Rogue.Options options = 3; */
        if (message.options)
            Rogue_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
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
class Rogue_Rotation$Type extends MessageType {
    constructor() {
        super("proto.Rogue.Rotation", [
            { no: 3, name: "builder", kind: "enum", T: () => ["proto.Rogue.Rotation.Builder", Rogue_Rotation_Builder] },
            { no: 1, name: "maintain_expose_armor", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "use_rupture", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "use_shiv", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "min_combo_points_for_damage_finisher", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value) {
        const message = { builder: 0, maintainExposeArmor: false, useRupture: false, useShiv: false, minComboPointsForDamageFinisher: 0 };
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
                case /* proto.Rogue.Rotation.Builder builder */ 3:
                    message.builder = reader.int32();
                    break;
                case /* bool maintain_expose_armor */ 1:
                    message.maintainExposeArmor = reader.bool();
                    break;
                case /* bool use_rupture */ 2:
                    message.useRupture = reader.bool();
                    break;
                case /* bool use_shiv */ 5:
                    message.useShiv = reader.bool();
                    break;
                case /* int32 min_combo_points_for_damage_finisher */ 4:
                    message.minComboPointsForDamageFinisher = reader.int32();
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
        /* proto.Rogue.Rotation.Builder builder = 3; */
        if (message.builder !== 0)
            writer.tag(3, WireType.Varint).int32(message.builder);
        /* bool maintain_expose_armor = 1; */
        if (message.maintainExposeArmor !== false)
            writer.tag(1, WireType.Varint).bool(message.maintainExposeArmor);
        /* bool use_rupture = 2; */
        if (message.useRupture !== false)
            writer.tag(2, WireType.Varint).bool(message.useRupture);
        /* bool use_shiv = 5; */
        if (message.useShiv !== false)
            writer.tag(5, WireType.Varint).bool(message.useShiv);
        /* int32 min_combo_points_for_damage_finisher = 4; */
        if (message.minComboPointsForDamageFinisher !== 0)
            writer.tag(4, WireType.Varint).int32(message.minComboPointsForDamageFinisher);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Rogue.Rotation
 */
export const Rogue_Rotation = new Rogue_Rotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Rogue_Options$Type extends MessageType {
    constructor() {
        super("proto.Rogue.Options", []);
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
 * @generated MessageType for protobuf message proto.Rogue.Options
 */
export const Rogue_Options = new Rogue_Options$Type();
