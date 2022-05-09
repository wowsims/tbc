import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
/**
 * @generated from protobuf enum proto.Warrior.Rotation.SunderArmor
 */
export var Warrior_Rotation_SunderArmor;
(function (Warrior_Rotation_SunderArmor) {
    /**
     * @generated from protobuf enum value: SunderArmorNone = 0;
     */
    Warrior_Rotation_SunderArmor[Warrior_Rotation_SunderArmor["SunderArmorNone"] = 0] = "SunderArmorNone";
    /**
     * @generated from protobuf enum value: SunderArmorHelpStack = 1;
     */
    Warrior_Rotation_SunderArmor[Warrior_Rotation_SunderArmor["SunderArmorHelpStack"] = 1] = "SunderArmorHelpStack";
    /**
     * @generated from protobuf enum value: SunderArmorMaintain = 2;
     */
    Warrior_Rotation_SunderArmor[Warrior_Rotation_SunderArmor["SunderArmorMaintain"] = 2] = "SunderArmorMaintain";
})(Warrior_Rotation_SunderArmor || (Warrior_Rotation_SunderArmor = {}));
/**
 * @generated from protobuf enum proto.ProtectionWarrior.Rotation.DemoShout
 */
export var ProtectionWarrior_Rotation_DemoShout;
(function (ProtectionWarrior_Rotation_DemoShout) {
    /**
     * @generated from protobuf enum value: DemoShoutNone = 0;
     */
    ProtectionWarrior_Rotation_DemoShout[ProtectionWarrior_Rotation_DemoShout["DemoShoutNone"] = 0] = "DemoShoutNone";
    /**
     * @generated from protobuf enum value: DemoShoutMaintain = 1;
     */
    ProtectionWarrior_Rotation_DemoShout[ProtectionWarrior_Rotation_DemoShout["DemoShoutMaintain"] = 1] = "DemoShoutMaintain";
    /**
     * @generated from protobuf enum value: DemoShoutFiller = 2;
     */
    ProtectionWarrior_Rotation_DemoShout[ProtectionWarrior_Rotation_DemoShout["DemoShoutFiller"] = 2] = "DemoShoutFiller";
})(ProtectionWarrior_Rotation_DemoShout || (ProtectionWarrior_Rotation_DemoShout = {}));
/**
 * @generated from protobuf enum proto.ProtectionWarrior.Rotation.ThunderClap
 */
export var ProtectionWarrior_Rotation_ThunderClap;
(function (ProtectionWarrior_Rotation_ThunderClap) {
    /**
     * @generated from protobuf enum value: ThunderClapNone = 0;
     */
    ProtectionWarrior_Rotation_ThunderClap[ProtectionWarrior_Rotation_ThunderClap["ThunderClapNone"] = 0] = "ThunderClapNone";
    /**
     * @generated from protobuf enum value: ThunderClapMaintain = 1;
     */
    ProtectionWarrior_Rotation_ThunderClap[ProtectionWarrior_Rotation_ThunderClap["ThunderClapMaintain"] = 1] = "ThunderClapMaintain";
    /**
     * @generated from protobuf enum value: ThunderClapOnCD = 2;
     */
    ProtectionWarrior_Rotation_ThunderClap[ProtectionWarrior_Rotation_ThunderClap["ThunderClapOnCD"] = 2] = "ThunderClapOnCD";
})(ProtectionWarrior_Rotation_ThunderClap || (ProtectionWarrior_Rotation_ThunderClap = {}));
/**
 * @generated from protobuf enum proto.ProtectionWarrior.Rotation.ShieldBlock
 */
export var ProtectionWarrior_Rotation_ShieldBlock;
(function (ProtectionWarrior_Rotation_ShieldBlock) {
    /**
     * @generated from protobuf enum value: ShieldBlockNone = 0;
     */
    ProtectionWarrior_Rotation_ShieldBlock[ProtectionWarrior_Rotation_ShieldBlock["ShieldBlockNone"] = 0] = "ShieldBlockNone";
    /**
     * @generated from protobuf enum value: ShieldBlockToProcRevenge = 1;
     */
    ProtectionWarrior_Rotation_ShieldBlock[ProtectionWarrior_Rotation_ShieldBlock["ShieldBlockToProcRevenge"] = 1] = "ShieldBlockToProcRevenge";
    /**
     * @generated from protobuf enum value: ShieldBlockOnCD = 2;
     */
    ProtectionWarrior_Rotation_ShieldBlock[ProtectionWarrior_Rotation_ShieldBlock["ShieldBlockOnCD"] = 2] = "ShieldBlockOnCD";
})(ProtectionWarrior_Rotation_ShieldBlock || (ProtectionWarrior_Rotation_ShieldBlock = {}));
/**
 * @generated from protobuf enum proto.WarriorShout
 */
export var WarriorShout;
(function (WarriorShout) {
    /**
     * @generated from protobuf enum value: WarriorShoutNone = 0;
     */
    WarriorShout[WarriorShout["WarriorShoutNone"] = 0] = "WarriorShoutNone";
    /**
     * @generated from protobuf enum value: WarriorShoutBattle = 1;
     */
    WarriorShout[WarriorShout["WarriorShoutBattle"] = 1] = "WarriorShoutBattle";
    /**
     * @generated from protobuf enum value: WarriorShoutCommanding = 2;
     */
    WarriorShout[WarriorShout["WarriorShoutCommanding"] = 2] = "WarriorShoutCommanding";
})(WarriorShout || (WarriorShout = {}));
// @generated message type with reflection information, may provide speed optimized methods
class WarriorTalents$Type extends MessageType {
    constructor() {
        super("proto.WarriorTalents", [
            { no: 1, name: "improved_heroic_strike", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 45, name: "deflection", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "improved_rend", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "improved_charge", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "improved_thunder_clap", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "improved_overpower", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "anger_management", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 7, name: "deep_wounds", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 8, name: "two_handed_weapon_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 9, name: "impale", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 10, name: "poleaxe_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 11, name: "death_wish", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 12, name: "mace_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 13, name: "sword_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 14, name: "improved_disciplines", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 15, name: "blood_frenzy", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 16, name: "mortal_strike", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 17, name: "improved_mortal_strike", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 18, name: "endless_rage", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 19, name: "booming_voice", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 20, name: "cruelty", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 46, name: "improved_demoralizing_shout", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 21, name: "unbridled_wrath", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 22, name: "improved_cleave", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 23, name: "commanding_presence", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 24, name: "dual_wield_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 25, name: "improved_execute", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 26, name: "improved_slam", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 27, name: "sweeping_strikes", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 28, name: "weapon_mastery", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 29, name: "improved_berserker_rage", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 30, name: "flurry", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 31, name: "precision", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 32, name: "bloodthirst", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 33, name: "improved_whirlwind", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 34, name: "improved_berserker_stance", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 35, name: "rampage", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 36, name: "improved_bloodrage", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 37, name: "tactical_mastery", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 47, name: "anticipation", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 48, name: "shield_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 49, name: "toughness", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 50, name: "improved_shield_block", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 38, name: "defiance", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 39, name: "improved_sunder_armor", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 51, name: "shield_mastery", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 40, name: "one_handed_weapon_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 52, name: "improved_defensive_stance", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 41, name: "shield_slam", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 42, name: "focused_rage", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 43, name: "vitality", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 44, name: "devastate", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { improvedHeroicStrike: 0, deflection: 0, improvedRend: 0, improvedCharge: 0, improvedThunderClap: 0, improvedOverpower: 0, angerManagement: false, deepWounds: 0, twoHandedWeaponSpecialization: 0, impale: 0, poleaxeSpecialization: 0, deathWish: false, maceSpecialization: 0, swordSpecialization: 0, improvedDisciplines: 0, bloodFrenzy: 0, mortalStrike: false, improvedMortalStrike: 0, endlessRage: false, boomingVoice: 0, cruelty: 0, improvedDemoralizingShout: 0, unbridledWrath: 0, improvedCleave: 0, commandingPresence: 0, dualWieldSpecialization: 0, improvedExecute: 0, improvedSlam: 0, sweepingStrikes: false, weaponMastery: 0, improvedBerserkerRage: 0, flurry: 0, precision: 0, bloodthirst: false, improvedWhirlwind: 0, improvedBerserkerStance: 0, rampage: false, improvedBloodrage: 0, tacticalMastery: 0, anticipation: 0, shieldSpecialization: 0, toughness: 0, improvedShieldBlock: false, defiance: 0, improvedSunderArmor: 0, shieldMastery: 0, oneHandedWeaponSpecialization: 0, improvedDefensiveStance: 0, shieldSlam: false, focusedRage: 0, vitality: 0, devastate: false };
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
                case /* int32 improved_heroic_strike */ 1:
                    message.improvedHeroicStrike = reader.int32();
                    break;
                case /* int32 deflection */ 45:
                    message.deflection = reader.int32();
                    break;
                case /* int32 improved_rend */ 2:
                    message.improvedRend = reader.int32();
                    break;
                case /* int32 improved_charge */ 3:
                    message.improvedCharge = reader.int32();
                    break;
                case /* int32 improved_thunder_clap */ 4:
                    message.improvedThunderClap = reader.int32();
                    break;
                case /* int32 improved_overpower */ 5:
                    message.improvedOverpower = reader.int32();
                    break;
                case /* bool anger_management */ 6:
                    message.angerManagement = reader.bool();
                    break;
                case /* int32 deep_wounds */ 7:
                    message.deepWounds = reader.int32();
                    break;
                case /* int32 two_handed_weapon_specialization */ 8:
                    message.twoHandedWeaponSpecialization = reader.int32();
                    break;
                case /* int32 impale */ 9:
                    message.impale = reader.int32();
                    break;
                case /* int32 poleaxe_specialization */ 10:
                    message.poleaxeSpecialization = reader.int32();
                    break;
                case /* bool death_wish */ 11:
                    message.deathWish = reader.bool();
                    break;
                case /* int32 mace_specialization */ 12:
                    message.maceSpecialization = reader.int32();
                    break;
                case /* int32 sword_specialization */ 13:
                    message.swordSpecialization = reader.int32();
                    break;
                case /* int32 improved_disciplines */ 14:
                    message.improvedDisciplines = reader.int32();
                    break;
                case /* int32 blood_frenzy */ 15:
                    message.bloodFrenzy = reader.int32();
                    break;
                case /* bool mortal_strike */ 16:
                    message.mortalStrike = reader.bool();
                    break;
                case /* int32 improved_mortal_strike */ 17:
                    message.improvedMortalStrike = reader.int32();
                    break;
                case /* bool endless_rage */ 18:
                    message.endlessRage = reader.bool();
                    break;
                case /* int32 booming_voice */ 19:
                    message.boomingVoice = reader.int32();
                    break;
                case /* int32 cruelty */ 20:
                    message.cruelty = reader.int32();
                    break;
                case /* int32 improved_demoralizing_shout */ 46:
                    message.improvedDemoralizingShout = reader.int32();
                    break;
                case /* int32 unbridled_wrath */ 21:
                    message.unbridledWrath = reader.int32();
                    break;
                case /* int32 improved_cleave */ 22:
                    message.improvedCleave = reader.int32();
                    break;
                case /* int32 commanding_presence */ 23:
                    message.commandingPresence = reader.int32();
                    break;
                case /* int32 dual_wield_specialization */ 24:
                    message.dualWieldSpecialization = reader.int32();
                    break;
                case /* int32 improved_execute */ 25:
                    message.improvedExecute = reader.int32();
                    break;
                case /* int32 improved_slam */ 26:
                    message.improvedSlam = reader.int32();
                    break;
                case /* bool sweeping_strikes */ 27:
                    message.sweepingStrikes = reader.bool();
                    break;
                case /* int32 weapon_mastery */ 28:
                    message.weaponMastery = reader.int32();
                    break;
                case /* int32 improved_berserker_rage */ 29:
                    message.improvedBerserkerRage = reader.int32();
                    break;
                case /* int32 flurry */ 30:
                    message.flurry = reader.int32();
                    break;
                case /* int32 precision */ 31:
                    message.precision = reader.int32();
                    break;
                case /* bool bloodthirst */ 32:
                    message.bloodthirst = reader.bool();
                    break;
                case /* int32 improved_whirlwind */ 33:
                    message.improvedWhirlwind = reader.int32();
                    break;
                case /* int32 improved_berserker_stance */ 34:
                    message.improvedBerserkerStance = reader.int32();
                    break;
                case /* bool rampage */ 35:
                    message.rampage = reader.bool();
                    break;
                case /* int32 improved_bloodrage */ 36:
                    message.improvedBloodrage = reader.int32();
                    break;
                case /* int32 tactical_mastery */ 37:
                    message.tacticalMastery = reader.int32();
                    break;
                case /* int32 anticipation */ 47:
                    message.anticipation = reader.int32();
                    break;
                case /* int32 shield_specialization */ 48:
                    message.shieldSpecialization = reader.int32();
                    break;
                case /* int32 toughness */ 49:
                    message.toughness = reader.int32();
                    break;
                case /* bool improved_shield_block */ 50:
                    message.improvedShieldBlock = reader.bool();
                    break;
                case /* int32 defiance */ 38:
                    message.defiance = reader.int32();
                    break;
                case /* int32 improved_sunder_armor */ 39:
                    message.improvedSunderArmor = reader.int32();
                    break;
                case /* int32 shield_mastery */ 51:
                    message.shieldMastery = reader.int32();
                    break;
                case /* int32 one_handed_weapon_specialization */ 40:
                    message.oneHandedWeaponSpecialization = reader.int32();
                    break;
                case /* int32 improved_defensive_stance */ 52:
                    message.improvedDefensiveStance = reader.int32();
                    break;
                case /* bool shield_slam */ 41:
                    message.shieldSlam = reader.bool();
                    break;
                case /* int32 focused_rage */ 42:
                    message.focusedRage = reader.int32();
                    break;
                case /* int32 vitality */ 43:
                    message.vitality = reader.int32();
                    break;
                case /* bool devastate */ 44:
                    message.devastate = reader.bool();
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
        /* int32 improved_heroic_strike = 1; */
        if (message.improvedHeroicStrike !== 0)
            writer.tag(1, WireType.Varint).int32(message.improvedHeroicStrike);
        /* int32 deflection = 45; */
        if (message.deflection !== 0)
            writer.tag(45, WireType.Varint).int32(message.deflection);
        /* int32 improved_rend = 2; */
        if (message.improvedRend !== 0)
            writer.tag(2, WireType.Varint).int32(message.improvedRend);
        /* int32 improved_charge = 3; */
        if (message.improvedCharge !== 0)
            writer.tag(3, WireType.Varint).int32(message.improvedCharge);
        /* int32 improved_thunder_clap = 4; */
        if (message.improvedThunderClap !== 0)
            writer.tag(4, WireType.Varint).int32(message.improvedThunderClap);
        /* int32 improved_overpower = 5; */
        if (message.improvedOverpower !== 0)
            writer.tag(5, WireType.Varint).int32(message.improvedOverpower);
        /* bool anger_management = 6; */
        if (message.angerManagement !== false)
            writer.tag(6, WireType.Varint).bool(message.angerManagement);
        /* int32 deep_wounds = 7; */
        if (message.deepWounds !== 0)
            writer.tag(7, WireType.Varint).int32(message.deepWounds);
        /* int32 two_handed_weapon_specialization = 8; */
        if (message.twoHandedWeaponSpecialization !== 0)
            writer.tag(8, WireType.Varint).int32(message.twoHandedWeaponSpecialization);
        /* int32 impale = 9; */
        if (message.impale !== 0)
            writer.tag(9, WireType.Varint).int32(message.impale);
        /* int32 poleaxe_specialization = 10; */
        if (message.poleaxeSpecialization !== 0)
            writer.tag(10, WireType.Varint).int32(message.poleaxeSpecialization);
        /* bool death_wish = 11; */
        if (message.deathWish !== false)
            writer.tag(11, WireType.Varint).bool(message.deathWish);
        /* int32 mace_specialization = 12; */
        if (message.maceSpecialization !== 0)
            writer.tag(12, WireType.Varint).int32(message.maceSpecialization);
        /* int32 sword_specialization = 13; */
        if (message.swordSpecialization !== 0)
            writer.tag(13, WireType.Varint).int32(message.swordSpecialization);
        /* int32 improved_disciplines = 14; */
        if (message.improvedDisciplines !== 0)
            writer.tag(14, WireType.Varint).int32(message.improvedDisciplines);
        /* int32 blood_frenzy = 15; */
        if (message.bloodFrenzy !== 0)
            writer.tag(15, WireType.Varint).int32(message.bloodFrenzy);
        /* bool mortal_strike = 16; */
        if (message.mortalStrike !== false)
            writer.tag(16, WireType.Varint).bool(message.mortalStrike);
        /* int32 improved_mortal_strike = 17; */
        if (message.improvedMortalStrike !== 0)
            writer.tag(17, WireType.Varint).int32(message.improvedMortalStrike);
        /* bool endless_rage = 18; */
        if (message.endlessRage !== false)
            writer.tag(18, WireType.Varint).bool(message.endlessRage);
        /* int32 booming_voice = 19; */
        if (message.boomingVoice !== 0)
            writer.tag(19, WireType.Varint).int32(message.boomingVoice);
        /* int32 cruelty = 20; */
        if (message.cruelty !== 0)
            writer.tag(20, WireType.Varint).int32(message.cruelty);
        /* int32 improved_demoralizing_shout = 46; */
        if (message.improvedDemoralizingShout !== 0)
            writer.tag(46, WireType.Varint).int32(message.improvedDemoralizingShout);
        /* int32 unbridled_wrath = 21; */
        if (message.unbridledWrath !== 0)
            writer.tag(21, WireType.Varint).int32(message.unbridledWrath);
        /* int32 improved_cleave = 22; */
        if (message.improvedCleave !== 0)
            writer.tag(22, WireType.Varint).int32(message.improvedCleave);
        /* int32 commanding_presence = 23; */
        if (message.commandingPresence !== 0)
            writer.tag(23, WireType.Varint).int32(message.commandingPresence);
        /* int32 dual_wield_specialization = 24; */
        if (message.dualWieldSpecialization !== 0)
            writer.tag(24, WireType.Varint).int32(message.dualWieldSpecialization);
        /* int32 improved_execute = 25; */
        if (message.improvedExecute !== 0)
            writer.tag(25, WireType.Varint).int32(message.improvedExecute);
        /* int32 improved_slam = 26; */
        if (message.improvedSlam !== 0)
            writer.tag(26, WireType.Varint).int32(message.improvedSlam);
        /* bool sweeping_strikes = 27; */
        if (message.sweepingStrikes !== false)
            writer.tag(27, WireType.Varint).bool(message.sweepingStrikes);
        /* int32 weapon_mastery = 28; */
        if (message.weaponMastery !== 0)
            writer.tag(28, WireType.Varint).int32(message.weaponMastery);
        /* int32 improved_berserker_rage = 29; */
        if (message.improvedBerserkerRage !== 0)
            writer.tag(29, WireType.Varint).int32(message.improvedBerserkerRage);
        /* int32 flurry = 30; */
        if (message.flurry !== 0)
            writer.tag(30, WireType.Varint).int32(message.flurry);
        /* int32 precision = 31; */
        if (message.precision !== 0)
            writer.tag(31, WireType.Varint).int32(message.precision);
        /* bool bloodthirst = 32; */
        if (message.bloodthirst !== false)
            writer.tag(32, WireType.Varint).bool(message.bloodthirst);
        /* int32 improved_whirlwind = 33; */
        if (message.improvedWhirlwind !== 0)
            writer.tag(33, WireType.Varint).int32(message.improvedWhirlwind);
        /* int32 improved_berserker_stance = 34; */
        if (message.improvedBerserkerStance !== 0)
            writer.tag(34, WireType.Varint).int32(message.improvedBerserkerStance);
        /* bool rampage = 35; */
        if (message.rampage !== false)
            writer.tag(35, WireType.Varint).bool(message.rampage);
        /* int32 improved_bloodrage = 36; */
        if (message.improvedBloodrage !== 0)
            writer.tag(36, WireType.Varint).int32(message.improvedBloodrage);
        /* int32 tactical_mastery = 37; */
        if (message.tacticalMastery !== 0)
            writer.tag(37, WireType.Varint).int32(message.tacticalMastery);
        /* int32 anticipation = 47; */
        if (message.anticipation !== 0)
            writer.tag(47, WireType.Varint).int32(message.anticipation);
        /* int32 shield_specialization = 48; */
        if (message.shieldSpecialization !== 0)
            writer.tag(48, WireType.Varint).int32(message.shieldSpecialization);
        /* int32 toughness = 49; */
        if (message.toughness !== 0)
            writer.tag(49, WireType.Varint).int32(message.toughness);
        /* bool improved_shield_block = 50; */
        if (message.improvedShieldBlock !== false)
            writer.tag(50, WireType.Varint).bool(message.improvedShieldBlock);
        /* int32 defiance = 38; */
        if (message.defiance !== 0)
            writer.tag(38, WireType.Varint).int32(message.defiance);
        /* int32 improved_sunder_armor = 39; */
        if (message.improvedSunderArmor !== 0)
            writer.tag(39, WireType.Varint).int32(message.improvedSunderArmor);
        /* int32 shield_mastery = 51; */
        if (message.shieldMastery !== 0)
            writer.tag(51, WireType.Varint).int32(message.shieldMastery);
        /* int32 one_handed_weapon_specialization = 40; */
        if (message.oneHandedWeaponSpecialization !== 0)
            writer.tag(40, WireType.Varint).int32(message.oneHandedWeaponSpecialization);
        /* int32 improved_defensive_stance = 52; */
        if (message.improvedDefensiveStance !== 0)
            writer.tag(52, WireType.Varint).int32(message.improvedDefensiveStance);
        /* bool shield_slam = 41; */
        if (message.shieldSlam !== false)
            writer.tag(41, WireType.Varint).bool(message.shieldSlam);
        /* int32 focused_rage = 42; */
        if (message.focusedRage !== 0)
            writer.tag(42, WireType.Varint).int32(message.focusedRage);
        /* int32 vitality = 43; */
        if (message.vitality !== 0)
            writer.tag(43, WireType.Varint).int32(message.vitality);
        /* bool devastate = 44; */
        if (message.devastate !== false)
            writer.tag(44, WireType.Varint).bool(message.devastate);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.WarriorTalents
 */
export const WarriorTalents = new WarriorTalents$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Warrior$Type extends MessageType {
    constructor() {
        super("proto.Warrior", [
            { no: 1, name: "rotation", kind: "message", T: () => Warrior_Rotation },
            { no: 2, name: "talents", kind: "message", T: () => WarriorTalents },
            { no: 3, name: "options", kind: "message", T: () => Warrior_Options }
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
                case /* proto.Warrior.Rotation rotation */ 1:
                    message.rotation = Warrior_Rotation.internalBinaryRead(reader, reader.uint32(), options, message.rotation);
                    break;
                case /* proto.WarriorTalents talents */ 2:
                    message.talents = WarriorTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* proto.Warrior.Options options */ 3:
                    message.options = Warrior_Options.internalBinaryRead(reader, reader.uint32(), options, message.options);
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
        /* proto.Warrior.Rotation rotation = 1; */
        if (message.rotation)
            Warrior_Rotation.internalBinaryWrite(message.rotation, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.WarriorTalents talents = 2; */
        if (message.talents)
            WarriorTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.Warrior.Options options = 3; */
        if (message.options)
            Warrior_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
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
class Warrior_Rotation$Type extends MessageType {
    constructor() {
        super("proto.Warrior.Rotation", [
            { no: 14, name: "use_cleave", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 1, name: "use_overpower", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "use_hamstring", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "use_slam", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "prioritize_ww", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 15, name: "sunderArmor", kind: "enum", T: () => ["proto.Warrior.Rotation.SunderArmor", Warrior_Rotation_SunderArmor] },
            { no: 16, name: "maintain_demo_shout", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 17, name: "maintain_thunder_clap", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "hs_rage_threshold", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 6, name: "overpower_rage_threshold", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 7, name: "hamstring_rage_threshold", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 8, name: "rampage_cd_threshold", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 9, name: "slam_latency", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 19, name: "slam_gcd_delay", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 20, name: "slam_ms_ww_delay", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 10, name: "use_hs_during_execute", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 13, name: "use_bt_during_execute", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 12, name: "use_ms_during_execute", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 11, name: "use_ww_during_execute", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 18, name: "use_slam_during_execute", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { useCleave: false, useOverpower: false, useHamstring: false, useSlam: false, prioritizeWw: false, sunderArmor: 0, maintainDemoShout: false, maintainThunderClap: false, hsRageThreshold: 0, overpowerRageThreshold: 0, hamstringRageThreshold: 0, rampageCdThreshold: 0, slamLatency: 0, slamGcdDelay: 0, slamMsWwDelay: 0, useHsDuringExecute: false, useBtDuringExecute: false, useMsDuringExecute: false, useWwDuringExecute: false, useSlamDuringExecute: false };
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
                case /* bool use_cleave */ 14:
                    message.useCleave = reader.bool();
                    break;
                case /* bool use_overpower */ 1:
                    message.useOverpower = reader.bool();
                    break;
                case /* bool use_hamstring */ 2:
                    message.useHamstring = reader.bool();
                    break;
                case /* bool use_slam */ 3:
                    message.useSlam = reader.bool();
                    break;
                case /* bool prioritize_ww */ 4:
                    message.prioritizeWw = reader.bool();
                    break;
                case /* proto.Warrior.Rotation.SunderArmor sunderArmor */ 15:
                    message.sunderArmor = reader.int32();
                    break;
                case /* bool maintain_demo_shout */ 16:
                    message.maintainDemoShout = reader.bool();
                    break;
                case /* bool maintain_thunder_clap */ 17:
                    message.maintainThunderClap = reader.bool();
                    break;
                case /* double hs_rage_threshold */ 5:
                    message.hsRageThreshold = reader.double();
                    break;
                case /* double overpower_rage_threshold */ 6:
                    message.overpowerRageThreshold = reader.double();
                    break;
                case /* double hamstring_rage_threshold */ 7:
                    message.hamstringRageThreshold = reader.double();
                    break;
                case /* double rampage_cd_threshold */ 8:
                    message.rampageCdThreshold = reader.double();
                    break;
                case /* double slam_latency */ 9:
                    message.slamLatency = reader.double();
                    break;
                case /* double slam_gcd_delay */ 19:
                    message.slamGcdDelay = reader.double();
                    break;
                case /* double slam_ms_ww_delay */ 20:
                    message.slamMsWwDelay = reader.double();
                    break;
                case /* bool use_hs_during_execute */ 10:
                    message.useHsDuringExecute = reader.bool();
                    break;
                case /* bool use_bt_during_execute */ 13:
                    message.useBtDuringExecute = reader.bool();
                    break;
                case /* bool use_ms_during_execute */ 12:
                    message.useMsDuringExecute = reader.bool();
                    break;
                case /* bool use_ww_during_execute */ 11:
                    message.useWwDuringExecute = reader.bool();
                    break;
                case /* bool use_slam_during_execute */ 18:
                    message.useSlamDuringExecute = reader.bool();
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
        /* bool use_cleave = 14; */
        if (message.useCleave !== false)
            writer.tag(14, WireType.Varint).bool(message.useCleave);
        /* bool use_overpower = 1; */
        if (message.useOverpower !== false)
            writer.tag(1, WireType.Varint).bool(message.useOverpower);
        /* bool use_hamstring = 2; */
        if (message.useHamstring !== false)
            writer.tag(2, WireType.Varint).bool(message.useHamstring);
        /* bool use_slam = 3; */
        if (message.useSlam !== false)
            writer.tag(3, WireType.Varint).bool(message.useSlam);
        /* bool prioritize_ww = 4; */
        if (message.prioritizeWw !== false)
            writer.tag(4, WireType.Varint).bool(message.prioritizeWw);
        /* proto.Warrior.Rotation.SunderArmor sunderArmor = 15; */
        if (message.sunderArmor !== 0)
            writer.tag(15, WireType.Varint).int32(message.sunderArmor);
        /* bool maintain_demo_shout = 16; */
        if (message.maintainDemoShout !== false)
            writer.tag(16, WireType.Varint).bool(message.maintainDemoShout);
        /* bool maintain_thunder_clap = 17; */
        if (message.maintainThunderClap !== false)
            writer.tag(17, WireType.Varint).bool(message.maintainThunderClap);
        /* double hs_rage_threshold = 5; */
        if (message.hsRageThreshold !== 0)
            writer.tag(5, WireType.Bit64).double(message.hsRageThreshold);
        /* double overpower_rage_threshold = 6; */
        if (message.overpowerRageThreshold !== 0)
            writer.tag(6, WireType.Bit64).double(message.overpowerRageThreshold);
        /* double hamstring_rage_threshold = 7; */
        if (message.hamstringRageThreshold !== 0)
            writer.tag(7, WireType.Bit64).double(message.hamstringRageThreshold);
        /* double rampage_cd_threshold = 8; */
        if (message.rampageCdThreshold !== 0)
            writer.tag(8, WireType.Bit64).double(message.rampageCdThreshold);
        /* double slam_latency = 9; */
        if (message.slamLatency !== 0)
            writer.tag(9, WireType.Bit64).double(message.slamLatency);
        /* double slam_gcd_delay = 19; */
        if (message.slamGcdDelay !== 0)
            writer.tag(19, WireType.Bit64).double(message.slamGcdDelay);
        /* double slam_ms_ww_delay = 20; */
        if (message.slamMsWwDelay !== 0)
            writer.tag(20, WireType.Bit64).double(message.slamMsWwDelay);
        /* bool use_hs_during_execute = 10; */
        if (message.useHsDuringExecute !== false)
            writer.tag(10, WireType.Varint).bool(message.useHsDuringExecute);
        /* bool use_bt_during_execute = 13; */
        if (message.useBtDuringExecute !== false)
            writer.tag(13, WireType.Varint).bool(message.useBtDuringExecute);
        /* bool use_ms_during_execute = 12; */
        if (message.useMsDuringExecute !== false)
            writer.tag(12, WireType.Varint).bool(message.useMsDuringExecute);
        /* bool use_ww_during_execute = 11; */
        if (message.useWwDuringExecute !== false)
            writer.tag(11, WireType.Varint).bool(message.useWwDuringExecute);
        /* bool use_slam_during_execute = 18; */
        if (message.useSlamDuringExecute !== false)
            writer.tag(18, WireType.Varint).bool(message.useSlamDuringExecute);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Warrior.Rotation
 */
export const Warrior_Rotation = new Warrior_Rotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Warrior_Options$Type extends MessageType {
    constructor() {
        super("proto.Warrior.Options", [
            { no: 1, name: "starting_rage", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 2, name: "use_recklessness", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "shout", kind: "enum", T: () => ["proto.WarriorShout", WarriorShout] },
            { no: 4, name: "precast_shout", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "precast_shout_t2", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 6, name: "precast_shout_sapphire", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { startingRage: 0, useRecklessness: false, shout: 0, precastShout: false, precastShoutT2: false, precastShoutSapphire: false };
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
                case /* double starting_rage */ 1:
                    message.startingRage = reader.double();
                    break;
                case /* bool use_recklessness */ 2:
                    message.useRecklessness = reader.bool();
                    break;
                case /* proto.WarriorShout shout */ 3:
                    message.shout = reader.int32();
                    break;
                case /* bool precast_shout */ 4:
                    message.precastShout = reader.bool();
                    break;
                case /* bool precast_shout_t2 */ 5:
                    message.precastShoutT2 = reader.bool();
                    break;
                case /* bool precast_shout_sapphire */ 6:
                    message.precastShoutSapphire = reader.bool();
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
        /* double starting_rage = 1; */
        if (message.startingRage !== 0)
            writer.tag(1, WireType.Bit64).double(message.startingRage);
        /* bool use_recklessness = 2; */
        if (message.useRecklessness !== false)
            writer.tag(2, WireType.Varint).bool(message.useRecklessness);
        /* proto.WarriorShout shout = 3; */
        if (message.shout !== 0)
            writer.tag(3, WireType.Varint).int32(message.shout);
        /* bool precast_shout = 4; */
        if (message.precastShout !== false)
            writer.tag(4, WireType.Varint).bool(message.precastShout);
        /* bool precast_shout_t2 = 5; */
        if (message.precastShoutT2 !== false)
            writer.tag(5, WireType.Varint).bool(message.precastShoutT2);
        /* bool precast_shout_sapphire = 6; */
        if (message.precastShoutSapphire !== false)
            writer.tag(6, WireType.Varint).bool(message.precastShoutSapphire);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Warrior.Options
 */
export const Warrior_Options = new Warrior_Options$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ProtectionWarrior$Type extends MessageType {
    constructor() {
        super("proto.ProtectionWarrior", [
            { no: 1, name: "rotation", kind: "message", T: () => ProtectionWarrior_Rotation },
            { no: 2, name: "talents", kind: "message", T: () => WarriorTalents },
            { no: 3, name: "options", kind: "message", T: () => ProtectionWarrior_Options }
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
                case /* proto.ProtectionWarrior.Rotation rotation */ 1:
                    message.rotation = ProtectionWarrior_Rotation.internalBinaryRead(reader, reader.uint32(), options, message.rotation);
                    break;
                case /* proto.WarriorTalents talents */ 2:
                    message.talents = WarriorTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* proto.ProtectionWarrior.Options options */ 3:
                    message.options = ProtectionWarrior_Options.internalBinaryRead(reader, reader.uint32(), options, message.options);
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
        /* proto.ProtectionWarrior.Rotation rotation = 1; */
        if (message.rotation)
            ProtectionWarrior_Rotation.internalBinaryWrite(message.rotation, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.WarriorTalents talents = 2; */
        if (message.talents)
            WarriorTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.ProtectionWarrior.Options options = 3; */
        if (message.options)
            ProtectionWarrior_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ProtectionWarrior
 */
export const ProtectionWarrior = new ProtectionWarrior$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ProtectionWarrior_Rotation$Type extends MessageType {
    constructor() {
        super("proto.ProtectionWarrior.Rotation", [
            { no: 1, name: "demo_shout", kind: "enum", T: () => ["proto.ProtectionWarrior.Rotation.DemoShout", ProtectionWarrior_Rotation_DemoShout] },
            { no: 2, name: "thunder_clap", kind: "enum", T: () => ["proto.ProtectionWarrior.Rotation.ThunderClap", ProtectionWarrior_Rotation_ThunderClap] },
            { no: 5, name: "shield_block", kind: "enum", T: () => ["proto.ProtectionWarrior.Rotation.ShieldBlock", ProtectionWarrior_Rotation_ShieldBlock] },
            { no: 4, name: "use_cleave", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "hs_rage_threshold", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value) {
        const message = { demoShout: 0, thunderClap: 0, shieldBlock: 0, useCleave: false, hsRageThreshold: 0 };
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
                case /* proto.ProtectionWarrior.Rotation.DemoShout demo_shout */ 1:
                    message.demoShout = reader.int32();
                    break;
                case /* proto.ProtectionWarrior.Rotation.ThunderClap thunder_clap */ 2:
                    message.thunderClap = reader.int32();
                    break;
                case /* proto.ProtectionWarrior.Rotation.ShieldBlock shield_block */ 5:
                    message.shieldBlock = reader.int32();
                    break;
                case /* bool use_cleave */ 4:
                    message.useCleave = reader.bool();
                    break;
                case /* int32 hs_rage_threshold */ 3:
                    message.hsRageThreshold = reader.int32();
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
        /* proto.ProtectionWarrior.Rotation.DemoShout demo_shout = 1; */
        if (message.demoShout !== 0)
            writer.tag(1, WireType.Varint).int32(message.demoShout);
        /* proto.ProtectionWarrior.Rotation.ThunderClap thunder_clap = 2; */
        if (message.thunderClap !== 0)
            writer.tag(2, WireType.Varint).int32(message.thunderClap);
        /* proto.ProtectionWarrior.Rotation.ShieldBlock shield_block = 5; */
        if (message.shieldBlock !== 0)
            writer.tag(5, WireType.Varint).int32(message.shieldBlock);
        /* bool use_cleave = 4; */
        if (message.useCleave !== false)
            writer.tag(4, WireType.Varint).bool(message.useCleave);
        /* int32 hs_rage_threshold = 3; */
        if (message.hsRageThreshold !== 0)
            writer.tag(3, WireType.Varint).int32(message.hsRageThreshold);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ProtectionWarrior.Rotation
 */
export const ProtectionWarrior_Rotation = new ProtectionWarrior_Rotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ProtectionWarrior_Options$Type extends MessageType {
    constructor() {
        super("proto.ProtectionWarrior.Options", [
            { no: 1, name: "starting_rage", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 4, name: "shout", kind: "enum", T: () => ["proto.WarriorShout", WarriorShout] },
            { no: 5, name: "precast_shout", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "precast_shout_t2", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "precast_shout_sapphire", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { startingRage: 0, shout: 0, precastShout: false, precastShoutT2: false, precastShoutSapphire: false };
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
                case /* double starting_rage */ 1:
                    message.startingRage = reader.double();
                    break;
                case /* proto.WarriorShout shout */ 4:
                    message.shout = reader.int32();
                    break;
                case /* bool precast_shout */ 5:
                    message.precastShout = reader.bool();
                    break;
                case /* bool precast_shout_t2 */ 2:
                    message.precastShoutT2 = reader.bool();
                    break;
                case /* bool precast_shout_sapphire */ 3:
                    message.precastShoutSapphire = reader.bool();
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
        /* double starting_rage = 1; */
        if (message.startingRage !== 0)
            writer.tag(1, WireType.Bit64).double(message.startingRage);
        /* proto.WarriorShout shout = 4; */
        if (message.shout !== 0)
            writer.tag(4, WireType.Varint).int32(message.shout);
        /* bool precast_shout = 5; */
        if (message.precastShout !== false)
            writer.tag(5, WireType.Varint).bool(message.precastShout);
        /* bool precast_shout_t2 = 2; */
        if (message.precastShoutT2 !== false)
            writer.tag(2, WireType.Varint).bool(message.precastShoutT2);
        /* bool precast_shout_sapphire = 3; */
        if (message.precastShoutSapphire !== false)
            writer.tag(3, WireType.Varint).bool(message.precastShoutSapphire);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ProtectionWarrior.Options
 */
export const ProtectionWarrior_Options = new ProtectionWarrior_Options$Type();
