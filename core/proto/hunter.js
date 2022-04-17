import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
/**
 * @generated from protobuf enum proto.Hunter.Rotation.StingType
 */
export var Hunter_Rotation_StingType;
(function (Hunter_Rotation_StingType) {
    /**
     * @generated from protobuf enum value: NoSting = 0;
     */
    Hunter_Rotation_StingType[Hunter_Rotation_StingType["NoSting"] = 0] = "NoSting";
    /**
     * @generated from protobuf enum value: ScorpidSting = 1;
     */
    Hunter_Rotation_StingType[Hunter_Rotation_StingType["ScorpidSting"] = 1] = "ScorpidSting";
    /**
     * @generated from protobuf enum value: SerpentSting = 2;
     */
    Hunter_Rotation_StingType[Hunter_Rotation_StingType["SerpentSting"] = 2] = "SerpentSting";
})(Hunter_Rotation_StingType || (Hunter_Rotation_StingType = {}));
/**
 * @generated from protobuf enum proto.Hunter.Rotation.WeaveType
 */
export var Hunter_Rotation_WeaveType;
(function (Hunter_Rotation_WeaveType) {
    /**
     * @generated from protobuf enum value: WeaveNone = 0;
     */
    Hunter_Rotation_WeaveType[Hunter_Rotation_WeaveType["WeaveNone"] = 0] = "WeaveNone";
    /**
     * @generated from protobuf enum value: WeaveAutosOnly = 1;
     */
    Hunter_Rotation_WeaveType[Hunter_Rotation_WeaveType["WeaveAutosOnly"] = 1] = "WeaveAutosOnly";
    /**
     * @generated from protobuf enum value: WeaveRaptorOnly = 2;
     */
    Hunter_Rotation_WeaveType[Hunter_Rotation_WeaveType["WeaveRaptorOnly"] = 2] = "WeaveRaptorOnly";
    /**
     * @generated from protobuf enum value: WeaveFull = 3;
     */
    Hunter_Rotation_WeaveType[Hunter_Rotation_WeaveType["WeaveFull"] = 3] = "WeaveFull";
})(Hunter_Rotation_WeaveType || (Hunter_Rotation_WeaveType = {}));
/**
 * @generated from protobuf enum proto.Hunter.Options.QuiverBonus
 */
export var Hunter_Options_QuiverBonus;
(function (Hunter_Options_QuiverBonus) {
    /**
     * @generated from protobuf enum value: QuiverNone = 0;
     */
    Hunter_Options_QuiverBonus[Hunter_Options_QuiverBonus["QuiverNone"] = 0] = "QuiverNone";
    /**
     * @generated from protobuf enum value: Speed10 = 1;
     */
    Hunter_Options_QuiverBonus[Hunter_Options_QuiverBonus["Speed10"] = 1] = "Speed10";
    /**
     * @generated from protobuf enum value: Speed11 = 2;
     */
    Hunter_Options_QuiverBonus[Hunter_Options_QuiverBonus["Speed11"] = 2] = "Speed11";
    /**
     * @generated from protobuf enum value: Speed12 = 3;
     */
    Hunter_Options_QuiverBonus[Hunter_Options_QuiverBonus["Speed12"] = 3] = "Speed12";
    /**
     * @generated from protobuf enum value: Speed13 = 4;
     */
    Hunter_Options_QuiverBonus[Hunter_Options_QuiverBonus["Speed13"] = 4] = "Speed13";
    /**
     * @generated from protobuf enum value: Speed14 = 5;
     */
    Hunter_Options_QuiverBonus[Hunter_Options_QuiverBonus["Speed14"] = 5] = "Speed14";
    /**
     * @generated from protobuf enum value: Speed15 = 6;
     */
    Hunter_Options_QuiverBonus[Hunter_Options_QuiverBonus["Speed15"] = 6] = "Speed15";
})(Hunter_Options_QuiverBonus || (Hunter_Options_QuiverBonus = {}));
/**
 * @generated from protobuf enum proto.Hunter.Options.Ammo
 */
export var Hunter_Options_Ammo;
(function (Hunter_Options_Ammo) {
    /**
     * @generated from protobuf enum value: AmmoNone = 0;
     */
    Hunter_Options_Ammo[Hunter_Options_Ammo["AmmoNone"] = 0] = "AmmoNone";
    /**
     * @generated from protobuf enum value: TimelessArrow = 1;
     */
    Hunter_Options_Ammo[Hunter_Options_Ammo["TimelessArrow"] = 1] = "TimelessArrow";
    /**
     * @generated from protobuf enum value: MysteriousArrow = 2;
     */
    Hunter_Options_Ammo[Hunter_Options_Ammo["MysteriousArrow"] = 2] = "MysteriousArrow";
    /**
     * @generated from protobuf enum value: AdamantiteStinger = 3;
     */
    Hunter_Options_Ammo[Hunter_Options_Ammo["AdamantiteStinger"] = 3] = "AdamantiteStinger";
    /**
     * @generated from protobuf enum value: WardensArrow = 4;
     */
    Hunter_Options_Ammo[Hunter_Options_Ammo["WardensArrow"] = 4] = "WardensArrow";
    /**
     * @generated from protobuf enum value: HalaaniRazorshaft = 5;
     */
    Hunter_Options_Ammo[Hunter_Options_Ammo["HalaaniRazorshaft"] = 5] = "HalaaniRazorshaft";
    /**
     * @generated from protobuf enum value: BlackflightArrow = 6;
     */
    Hunter_Options_Ammo[Hunter_Options_Ammo["BlackflightArrow"] = 6] = "BlackflightArrow";
})(Hunter_Options_Ammo || (Hunter_Options_Ammo = {}));
/**
 * @generated from protobuf enum proto.Hunter.Options.PetType
 */
export var Hunter_Options_PetType;
(function (Hunter_Options_PetType) {
    /**
     * @generated from protobuf enum value: PetNone = 0;
     */
    Hunter_Options_PetType[Hunter_Options_PetType["PetNone"] = 0] = "PetNone";
    /**
     * @generated from protobuf enum value: Bat = 5;
     */
    Hunter_Options_PetType[Hunter_Options_PetType["Bat"] = 5] = "Bat";
    /**
     * @generated from protobuf enum value: Bear = 7;
     */
    Hunter_Options_PetType[Hunter_Options_PetType["Bear"] = 7] = "Bear";
    /**
     * @generated from protobuf enum value: Cat = 2;
     */
    Hunter_Options_PetType[Hunter_Options_PetType["Cat"] = 2] = "Cat";
    /**
     * @generated from protobuf enum value: Crab = 8;
     */
    Hunter_Options_PetType[Hunter_Options_PetType["Crab"] = 8] = "Crab";
    /**
     * @generated from protobuf enum value: Owl = 4;
     */
    Hunter_Options_PetType[Hunter_Options_PetType["Owl"] = 4] = "Owl";
    /**
     * @generated from protobuf enum value: Raptor = 3;
     */
    Hunter_Options_PetType[Hunter_Options_PetType["Raptor"] = 3] = "Raptor";
    /**
     * @generated from protobuf enum value: Ravager = 1;
     */
    Hunter_Options_PetType[Hunter_Options_PetType["Ravager"] = 1] = "Ravager";
    /**
     * @generated from protobuf enum value: WindSerpent = 6;
     */
    Hunter_Options_PetType[Hunter_Options_PetType["WindSerpent"] = 6] = "WindSerpent";
})(Hunter_Options_PetType || (Hunter_Options_PetType = {}));
// @generated message type with reflection information, may provide speed optimized methods
class HunterTalents$Type extends MessageType {
    constructor() {
        super("proto.HunterTalents", [
            { no: 1, name: "improved_aspect_of_the_hawk", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "endurance_training", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "focused_fire", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "unleashed_fury", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "ferocity", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "bestial_discipline", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 45, name: "animal_handler", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 7, name: "frenzy", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 8, name: "ferocious_inspiration", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 9, name: "bestial_wrath", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 10, name: "serpents_swiftness", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 11, name: "the_beast_within", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 12, name: "lethal_shots", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 13, name: "improved_hunters_mark", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 14, name: "efficiency", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 15, name: "go_for_the_throat", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 16, name: "improved_arcane_shot", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 17, name: "aimed_shot", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 18, name: "rapid_killing", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 19, name: "improved_stings", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 20, name: "mortal_shots", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 21, name: "scatter_shot", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 22, name: "barrage", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 23, name: "combat_experience", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 24, name: "ranged_weapon_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 25, name: "careful_aim", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 26, name: "trueshot_aura", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 27, name: "improved_barrage", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 28, name: "master_marksman", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 29, name: "silencing_shot", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 30, name: "monster_slaying", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 31, name: "humanoid_slaying", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 32, name: "savage_strikes", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 46, name: "deflection", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 33, name: "clever_traps", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 34, name: "survivalist", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 35, name: "trap_mastery", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 36, name: "surefooted", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 37, name: "survival_instincts", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 38, name: "killer_instinct", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 39, name: "resourcefulness", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 40, name: "lightning_reflexes", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 41, name: "thrill_of_the_hunt", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 42, name: "expose_weakness", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 43, name: "master_tactician", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 44, name: "readiness", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { improvedAspectOfTheHawk: 0, enduranceTraining: 0, focusedFire: 0, unleashedFury: 0, ferocity: 0, bestialDiscipline: 0, animalHandler: 0, frenzy: 0, ferociousInspiration: 0, bestialWrath: false, serpentsSwiftness: 0, theBeastWithin: false, lethalShots: 0, improvedHuntersMark: 0, efficiency: 0, goForTheThroat: 0, improvedArcaneShot: 0, aimedShot: false, rapidKilling: 0, improvedStings: 0, mortalShots: 0, scatterShot: false, barrage: 0, combatExperience: 0, rangedWeaponSpecialization: 0, carefulAim: 0, trueshotAura: false, improvedBarrage: 0, masterMarksman: 0, silencingShot: false, monsterSlaying: 0, humanoidSlaying: 0, savageStrikes: 0, deflection: 0, cleverTraps: 0, survivalist: 0, trapMastery: 0, surefooted: 0, survivalInstincts: 0, killerInstinct: 0, resourcefulness: 0, lightningReflexes: 0, thrillOfTheHunt: 0, exposeWeakness: 0, masterTactician: 0, readiness: false };
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
                case /* int32 improved_aspect_of_the_hawk */ 1:
                    message.improvedAspectOfTheHawk = reader.int32();
                    break;
                case /* int32 endurance_training */ 2:
                    message.enduranceTraining = reader.int32();
                    break;
                case /* int32 focused_fire */ 3:
                    message.focusedFire = reader.int32();
                    break;
                case /* int32 unleashed_fury */ 4:
                    message.unleashedFury = reader.int32();
                    break;
                case /* int32 ferocity */ 5:
                    message.ferocity = reader.int32();
                    break;
                case /* int32 bestial_discipline */ 6:
                    message.bestialDiscipline = reader.int32();
                    break;
                case /* int32 animal_handler */ 45:
                    message.animalHandler = reader.int32();
                    break;
                case /* int32 frenzy */ 7:
                    message.frenzy = reader.int32();
                    break;
                case /* int32 ferocious_inspiration */ 8:
                    message.ferociousInspiration = reader.int32();
                    break;
                case /* bool bestial_wrath */ 9:
                    message.bestialWrath = reader.bool();
                    break;
                case /* int32 serpents_swiftness */ 10:
                    message.serpentsSwiftness = reader.int32();
                    break;
                case /* bool the_beast_within */ 11:
                    message.theBeastWithin = reader.bool();
                    break;
                case /* int32 lethal_shots */ 12:
                    message.lethalShots = reader.int32();
                    break;
                case /* int32 improved_hunters_mark */ 13:
                    message.improvedHuntersMark = reader.int32();
                    break;
                case /* int32 efficiency */ 14:
                    message.efficiency = reader.int32();
                    break;
                case /* int32 go_for_the_throat */ 15:
                    message.goForTheThroat = reader.int32();
                    break;
                case /* int32 improved_arcane_shot */ 16:
                    message.improvedArcaneShot = reader.int32();
                    break;
                case /* bool aimed_shot */ 17:
                    message.aimedShot = reader.bool();
                    break;
                case /* int32 rapid_killing */ 18:
                    message.rapidKilling = reader.int32();
                    break;
                case /* int32 improved_stings */ 19:
                    message.improvedStings = reader.int32();
                    break;
                case /* int32 mortal_shots */ 20:
                    message.mortalShots = reader.int32();
                    break;
                case /* bool scatter_shot */ 21:
                    message.scatterShot = reader.bool();
                    break;
                case /* int32 barrage */ 22:
                    message.barrage = reader.int32();
                    break;
                case /* int32 combat_experience */ 23:
                    message.combatExperience = reader.int32();
                    break;
                case /* int32 ranged_weapon_specialization */ 24:
                    message.rangedWeaponSpecialization = reader.int32();
                    break;
                case /* int32 careful_aim */ 25:
                    message.carefulAim = reader.int32();
                    break;
                case /* bool trueshot_aura */ 26:
                    message.trueshotAura = reader.bool();
                    break;
                case /* int32 improved_barrage */ 27:
                    message.improvedBarrage = reader.int32();
                    break;
                case /* int32 master_marksman */ 28:
                    message.masterMarksman = reader.int32();
                    break;
                case /* bool silencing_shot */ 29:
                    message.silencingShot = reader.bool();
                    break;
                case /* int32 monster_slaying */ 30:
                    message.monsterSlaying = reader.int32();
                    break;
                case /* int32 humanoid_slaying */ 31:
                    message.humanoidSlaying = reader.int32();
                    break;
                case /* int32 savage_strikes */ 32:
                    message.savageStrikes = reader.int32();
                    break;
                case /* int32 deflection */ 46:
                    message.deflection = reader.int32();
                    break;
                case /* int32 clever_traps */ 33:
                    message.cleverTraps = reader.int32();
                    break;
                case /* int32 survivalist */ 34:
                    message.survivalist = reader.int32();
                    break;
                case /* int32 trap_mastery */ 35:
                    message.trapMastery = reader.int32();
                    break;
                case /* int32 surefooted */ 36:
                    message.surefooted = reader.int32();
                    break;
                case /* int32 survival_instincts */ 37:
                    message.survivalInstincts = reader.int32();
                    break;
                case /* int32 killer_instinct */ 38:
                    message.killerInstinct = reader.int32();
                    break;
                case /* int32 resourcefulness */ 39:
                    message.resourcefulness = reader.int32();
                    break;
                case /* int32 lightning_reflexes */ 40:
                    message.lightningReflexes = reader.int32();
                    break;
                case /* int32 thrill_of_the_hunt */ 41:
                    message.thrillOfTheHunt = reader.int32();
                    break;
                case /* int32 expose_weakness */ 42:
                    message.exposeWeakness = reader.int32();
                    break;
                case /* int32 master_tactician */ 43:
                    message.masterTactician = reader.int32();
                    break;
                case /* bool readiness */ 44:
                    message.readiness = reader.bool();
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
        /* int32 improved_aspect_of_the_hawk = 1; */
        if (message.improvedAspectOfTheHawk !== 0)
            writer.tag(1, WireType.Varint).int32(message.improvedAspectOfTheHawk);
        /* int32 endurance_training = 2; */
        if (message.enduranceTraining !== 0)
            writer.tag(2, WireType.Varint).int32(message.enduranceTraining);
        /* int32 focused_fire = 3; */
        if (message.focusedFire !== 0)
            writer.tag(3, WireType.Varint).int32(message.focusedFire);
        /* int32 unleashed_fury = 4; */
        if (message.unleashedFury !== 0)
            writer.tag(4, WireType.Varint).int32(message.unleashedFury);
        /* int32 ferocity = 5; */
        if (message.ferocity !== 0)
            writer.tag(5, WireType.Varint).int32(message.ferocity);
        /* int32 bestial_discipline = 6; */
        if (message.bestialDiscipline !== 0)
            writer.tag(6, WireType.Varint).int32(message.bestialDiscipline);
        /* int32 animal_handler = 45; */
        if (message.animalHandler !== 0)
            writer.tag(45, WireType.Varint).int32(message.animalHandler);
        /* int32 frenzy = 7; */
        if (message.frenzy !== 0)
            writer.tag(7, WireType.Varint).int32(message.frenzy);
        /* int32 ferocious_inspiration = 8; */
        if (message.ferociousInspiration !== 0)
            writer.tag(8, WireType.Varint).int32(message.ferociousInspiration);
        /* bool bestial_wrath = 9; */
        if (message.bestialWrath !== false)
            writer.tag(9, WireType.Varint).bool(message.bestialWrath);
        /* int32 serpents_swiftness = 10; */
        if (message.serpentsSwiftness !== 0)
            writer.tag(10, WireType.Varint).int32(message.serpentsSwiftness);
        /* bool the_beast_within = 11; */
        if (message.theBeastWithin !== false)
            writer.tag(11, WireType.Varint).bool(message.theBeastWithin);
        /* int32 lethal_shots = 12; */
        if (message.lethalShots !== 0)
            writer.tag(12, WireType.Varint).int32(message.lethalShots);
        /* int32 improved_hunters_mark = 13; */
        if (message.improvedHuntersMark !== 0)
            writer.tag(13, WireType.Varint).int32(message.improvedHuntersMark);
        /* int32 efficiency = 14; */
        if (message.efficiency !== 0)
            writer.tag(14, WireType.Varint).int32(message.efficiency);
        /* int32 go_for_the_throat = 15; */
        if (message.goForTheThroat !== 0)
            writer.tag(15, WireType.Varint).int32(message.goForTheThroat);
        /* int32 improved_arcane_shot = 16; */
        if (message.improvedArcaneShot !== 0)
            writer.tag(16, WireType.Varint).int32(message.improvedArcaneShot);
        /* bool aimed_shot = 17; */
        if (message.aimedShot !== false)
            writer.tag(17, WireType.Varint).bool(message.aimedShot);
        /* int32 rapid_killing = 18; */
        if (message.rapidKilling !== 0)
            writer.tag(18, WireType.Varint).int32(message.rapidKilling);
        /* int32 improved_stings = 19; */
        if (message.improvedStings !== 0)
            writer.tag(19, WireType.Varint).int32(message.improvedStings);
        /* int32 mortal_shots = 20; */
        if (message.mortalShots !== 0)
            writer.tag(20, WireType.Varint).int32(message.mortalShots);
        /* bool scatter_shot = 21; */
        if (message.scatterShot !== false)
            writer.tag(21, WireType.Varint).bool(message.scatterShot);
        /* int32 barrage = 22; */
        if (message.barrage !== 0)
            writer.tag(22, WireType.Varint).int32(message.barrage);
        /* int32 combat_experience = 23; */
        if (message.combatExperience !== 0)
            writer.tag(23, WireType.Varint).int32(message.combatExperience);
        /* int32 ranged_weapon_specialization = 24; */
        if (message.rangedWeaponSpecialization !== 0)
            writer.tag(24, WireType.Varint).int32(message.rangedWeaponSpecialization);
        /* int32 careful_aim = 25; */
        if (message.carefulAim !== 0)
            writer.tag(25, WireType.Varint).int32(message.carefulAim);
        /* bool trueshot_aura = 26; */
        if (message.trueshotAura !== false)
            writer.tag(26, WireType.Varint).bool(message.trueshotAura);
        /* int32 improved_barrage = 27; */
        if (message.improvedBarrage !== 0)
            writer.tag(27, WireType.Varint).int32(message.improvedBarrage);
        /* int32 master_marksman = 28; */
        if (message.masterMarksman !== 0)
            writer.tag(28, WireType.Varint).int32(message.masterMarksman);
        /* bool silencing_shot = 29; */
        if (message.silencingShot !== false)
            writer.tag(29, WireType.Varint).bool(message.silencingShot);
        /* int32 monster_slaying = 30; */
        if (message.monsterSlaying !== 0)
            writer.tag(30, WireType.Varint).int32(message.monsterSlaying);
        /* int32 humanoid_slaying = 31; */
        if (message.humanoidSlaying !== 0)
            writer.tag(31, WireType.Varint).int32(message.humanoidSlaying);
        /* int32 savage_strikes = 32; */
        if (message.savageStrikes !== 0)
            writer.tag(32, WireType.Varint).int32(message.savageStrikes);
        /* int32 deflection = 46; */
        if (message.deflection !== 0)
            writer.tag(46, WireType.Varint).int32(message.deflection);
        /* int32 clever_traps = 33; */
        if (message.cleverTraps !== 0)
            writer.tag(33, WireType.Varint).int32(message.cleverTraps);
        /* int32 survivalist = 34; */
        if (message.survivalist !== 0)
            writer.tag(34, WireType.Varint).int32(message.survivalist);
        /* int32 trap_mastery = 35; */
        if (message.trapMastery !== 0)
            writer.tag(35, WireType.Varint).int32(message.trapMastery);
        /* int32 surefooted = 36; */
        if (message.surefooted !== 0)
            writer.tag(36, WireType.Varint).int32(message.surefooted);
        /* int32 survival_instincts = 37; */
        if (message.survivalInstincts !== 0)
            writer.tag(37, WireType.Varint).int32(message.survivalInstincts);
        /* int32 killer_instinct = 38; */
        if (message.killerInstinct !== 0)
            writer.tag(38, WireType.Varint).int32(message.killerInstinct);
        /* int32 resourcefulness = 39; */
        if (message.resourcefulness !== 0)
            writer.tag(39, WireType.Varint).int32(message.resourcefulness);
        /* int32 lightning_reflexes = 40; */
        if (message.lightningReflexes !== 0)
            writer.tag(40, WireType.Varint).int32(message.lightningReflexes);
        /* int32 thrill_of_the_hunt = 41; */
        if (message.thrillOfTheHunt !== 0)
            writer.tag(41, WireType.Varint).int32(message.thrillOfTheHunt);
        /* int32 expose_weakness = 42; */
        if (message.exposeWeakness !== 0)
            writer.tag(42, WireType.Varint).int32(message.exposeWeakness);
        /* int32 master_tactician = 43; */
        if (message.masterTactician !== 0)
            writer.tag(43, WireType.Varint).int32(message.masterTactician);
        /* bool readiness = 44; */
        if (message.readiness !== false)
            writer.tag(44, WireType.Varint).bool(message.readiness);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.HunterTalents
 */
export const HunterTalents = new HunterTalents$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Hunter$Type extends MessageType {
    constructor() {
        super("proto.Hunter", [
            { no: 1, name: "rotation", kind: "message", T: () => Hunter_Rotation },
            { no: 2, name: "talents", kind: "message", T: () => HunterTalents },
            { no: 3, name: "options", kind: "message", T: () => Hunter_Options }
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
                case /* proto.Hunter.Rotation rotation */ 1:
                    message.rotation = Hunter_Rotation.internalBinaryRead(reader, reader.uint32(), options, message.rotation);
                    break;
                case /* proto.HunterTalents talents */ 2:
                    message.talents = HunterTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* proto.Hunter.Options options */ 3:
                    message.options = Hunter_Options.internalBinaryRead(reader, reader.uint32(), options, message.options);
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
        /* proto.Hunter.Rotation rotation = 1; */
        if (message.rotation)
            Hunter_Rotation.internalBinaryWrite(message.rotation, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.HunterTalents talents = 2; */
        if (message.talents)
            HunterTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.Hunter.Options options = 3; */
        if (message.options)
            Hunter_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
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
class Hunter_Rotation$Type extends MessageType {
    constructor() {
        super("proto.Hunter.Rotation", [
            { no: 1, name: "use_multi_shot", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "use_arcane_shot", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "precast_aimed_shot", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 12, name: "lazy_rotation", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "sting", kind: "enum", T: () => ["proto.Hunter.Rotation.StingType", Hunter_Rotation_StingType] },
            { no: 6, name: "viper_start_mana_percent", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 7, name: "viper_stop_mana_percent", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 11, name: "weave", kind: "enum", T: () => ["proto.Hunter.Rotation.WeaveType", Hunter_Rotation_WeaveType] },
            { no: 9, name: "time_to_weave_ms", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 10, name: "percent_weaved", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { useMultiShot: false, useArcaneShot: false, precastAimedShot: false, lazyRotation: false, sting: 0, viperStartManaPercent: 0, viperStopManaPercent: 0, weave: 0, timeToWeaveMs: 0, percentWeaved: 0 };
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
                case /* bool use_multi_shot */ 1:
                    message.useMultiShot = reader.bool();
                    break;
                case /* bool use_arcane_shot */ 2:
                    message.useArcaneShot = reader.bool();
                    break;
                case /* bool precast_aimed_shot */ 3:
                    message.precastAimedShot = reader.bool();
                    break;
                case /* bool lazy_rotation */ 12:
                    message.lazyRotation = reader.bool();
                    break;
                case /* proto.Hunter.Rotation.StingType sting */ 5:
                    message.sting = reader.int32();
                    break;
                case /* double viper_start_mana_percent */ 6:
                    message.viperStartManaPercent = reader.double();
                    break;
                case /* double viper_stop_mana_percent */ 7:
                    message.viperStopManaPercent = reader.double();
                    break;
                case /* proto.Hunter.Rotation.WeaveType weave */ 11:
                    message.weave = reader.int32();
                    break;
                case /* int32 time_to_weave_ms */ 9:
                    message.timeToWeaveMs = reader.int32();
                    break;
                case /* double percent_weaved */ 10:
                    message.percentWeaved = reader.double();
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
        /* bool use_multi_shot = 1; */
        if (message.useMultiShot !== false)
            writer.tag(1, WireType.Varint).bool(message.useMultiShot);
        /* bool use_arcane_shot = 2; */
        if (message.useArcaneShot !== false)
            writer.tag(2, WireType.Varint).bool(message.useArcaneShot);
        /* bool precast_aimed_shot = 3; */
        if (message.precastAimedShot !== false)
            writer.tag(3, WireType.Varint).bool(message.precastAimedShot);
        /* bool lazy_rotation = 12; */
        if (message.lazyRotation !== false)
            writer.tag(12, WireType.Varint).bool(message.lazyRotation);
        /* proto.Hunter.Rotation.StingType sting = 5; */
        if (message.sting !== 0)
            writer.tag(5, WireType.Varint).int32(message.sting);
        /* double viper_start_mana_percent = 6; */
        if (message.viperStartManaPercent !== 0)
            writer.tag(6, WireType.Bit64).double(message.viperStartManaPercent);
        /* double viper_stop_mana_percent = 7; */
        if (message.viperStopManaPercent !== 0)
            writer.tag(7, WireType.Bit64).double(message.viperStopManaPercent);
        /* proto.Hunter.Rotation.WeaveType weave = 11; */
        if (message.weave !== 0)
            writer.tag(11, WireType.Varint).int32(message.weave);
        /* int32 time_to_weave_ms = 9; */
        if (message.timeToWeaveMs !== 0)
            writer.tag(9, WireType.Varint).int32(message.timeToWeaveMs);
        /* double percent_weaved = 10; */
        if (message.percentWeaved !== 0)
            writer.tag(10, WireType.Bit64).double(message.percentWeaved);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Hunter.Rotation
 */
export const Hunter_Rotation = new Hunter_Rotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Hunter_Options$Type extends MessageType {
    constructor() {
        super("proto.Hunter.Options", [
            { no: 1, name: "quiver_bonus", kind: "enum", T: () => ["proto.Hunter.Options.QuiverBonus", Hunter_Options_QuiverBonus] },
            { no: 2, name: "ammo", kind: "enum", T: () => ["proto.Hunter.Options.Ammo", Hunter_Options_Ammo] },
            { no: 3, name: "pet_type", kind: "enum", T: () => ["proto.Hunter.Options.PetType", Hunter_Options_PetType] },
            { no: 4, name: "pet_uptime", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 6, name: "pet_single_ability", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "latency_ms", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 7, name: "remove_randomness", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { quiverBonus: 0, ammo: 0, petType: 0, petUptime: 0, petSingleAbility: false, latencyMs: 0, removeRandomness: false };
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
                case /* proto.Hunter.Options.QuiverBonus quiver_bonus */ 1:
                    message.quiverBonus = reader.int32();
                    break;
                case /* proto.Hunter.Options.Ammo ammo */ 2:
                    message.ammo = reader.int32();
                    break;
                case /* proto.Hunter.Options.PetType pet_type */ 3:
                    message.petType = reader.int32();
                    break;
                case /* double pet_uptime */ 4:
                    message.petUptime = reader.double();
                    break;
                case /* bool pet_single_ability */ 6:
                    message.petSingleAbility = reader.bool();
                    break;
                case /* int32 latency_ms */ 5:
                    message.latencyMs = reader.int32();
                    break;
                case /* bool remove_randomness */ 7:
                    message.removeRandomness = reader.bool();
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
        /* proto.Hunter.Options.QuiverBonus quiver_bonus = 1; */
        if (message.quiverBonus !== 0)
            writer.tag(1, WireType.Varint).int32(message.quiverBonus);
        /* proto.Hunter.Options.Ammo ammo = 2; */
        if (message.ammo !== 0)
            writer.tag(2, WireType.Varint).int32(message.ammo);
        /* proto.Hunter.Options.PetType pet_type = 3; */
        if (message.petType !== 0)
            writer.tag(3, WireType.Varint).int32(message.petType);
        /* double pet_uptime = 4; */
        if (message.petUptime !== 0)
            writer.tag(4, WireType.Bit64).double(message.petUptime);
        /* bool pet_single_ability = 6; */
        if (message.petSingleAbility !== false)
            writer.tag(6, WireType.Varint).bool(message.petSingleAbility);
        /* int32 latency_ms = 5; */
        if (message.latencyMs !== 0)
            writer.tag(5, WireType.Varint).int32(message.latencyMs);
        /* bool remove_randomness = 7; */
        if (message.removeRandomness !== false)
            writer.tag(7, WireType.Varint).bool(message.removeRandomness);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Hunter.Options
 */
export const Hunter_Options = new Hunter_Options$Type();
