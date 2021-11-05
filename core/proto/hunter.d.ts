import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message proto.HunterTalents
 */
export interface HunterTalents {
    /**
     * Beast Mastery
     *
     * @generated from protobuf field: int32 improved_aspect_of_the_hawk = 1;
     */
    improvedAspectOfTheHawk: number;
    /**
     * @generated from protobuf field: int32 endurance_training = 2;
     */
    enduranceTraining: number;
    /**
     * @generated from protobuf field: int32 focused_fire = 3;
     */
    focusedFire: number;
    /**
     * @generated from protobuf field: int32 unleashed_fury = 4;
     */
    unleashedFury: number;
    /**
     * @generated from protobuf field: int32 ferocity = 5;
     */
    ferocity: number;
    /**
     * @generated from protobuf field: int32 bestial_discipline = 6;
     */
    bestialDiscipline: number;
    /**
     * @generated from protobuf field: int32 frenzy = 7;
     */
    frenzy: number;
    /**
     * @generated from protobuf field: int32 ferocious_inspiration = 8;
     */
    ferociousInspiration: number;
    /**
     * @generated from protobuf field: bool bestial_wrath = 9;
     */
    bestialWrath: boolean;
    /**
     * @generated from protobuf field: int32 serpents_swiftness = 10;
     */
    serpentsSwiftness: number;
    /**
     * @generated from protobuf field: bool the_beast_within = 11;
     */
    theBeastWithin: boolean;
    /**
     * Marksmanship
     *
     * @generated from protobuf field: int32 lethal_shots = 12;
     */
    lethalShots: number;
    /**
     * @generated from protobuf field: int32 improved_hunters_mark = 13;
     */
    improvedHuntersMark: number;
    /**
     * @generated from protobuf field: int32 efficiency = 14;
     */
    efficiency: number;
    /**
     * @generated from protobuf field: int32 go_for_the_throat = 15;
     */
    goForTheThroat: number;
    /**
     * @generated from protobuf field: int32 improved_arcane_shot = 16;
     */
    improvedArcaneShot: number;
    /**
     * @generated from protobuf field: bool aimed_shot = 17;
     */
    aimedShot: boolean;
    /**
     * @generated from protobuf field: int32 rapid_killing = 18;
     */
    rapidKilling: number;
    /**
     * @generated from protobuf field: int32 improved_stings = 19;
     */
    improvedStings: number;
    /**
     * @generated from protobuf field: int32 mortal_shots = 20;
     */
    mortalShots: number;
    /**
     * @generated from protobuf field: bool scatter_shot = 21;
     */
    scatterShot: boolean;
    /**
     * @generated from protobuf field: int32 barrage = 22;
     */
    barrage: number;
    /**
     * @generated from protobuf field: int32 combat_experience = 23;
     */
    combatExperience: number;
    /**
     * @generated from protobuf field: int32 ranged_weapon_specialization = 24;
     */
    rangedWeaponSpecialization: number;
    /**
     * @generated from protobuf field: int32 careful_aim = 25;
     */
    carefulAim: number;
    /**
     * @generated from protobuf field: bool trueshot_aura = 26;
     */
    trueshotAura: boolean;
    /**
     * @generated from protobuf field: int32 improved_barrage = 27;
     */
    improvedBarrage: number;
    /**
     * @generated from protobuf field: int32 master_marksman = 28;
     */
    masterMarksman: number;
    /**
     * @generated from protobuf field: bool silencing_shot = 29;
     */
    silencingShot: boolean;
    /**
     * Survival
     *
     * @generated from protobuf field: int32 monster_slaying = 30;
     */
    monsterSlaying: number;
    /**
     * @generated from protobuf field: int32 humanoid_slaying = 31;
     */
    humanoidSlaying: number;
    /**
     * @generated from protobuf field: int32 savage_strikes = 32;
     */
    savageStrikes: number;
    /**
     * @generated from protobuf field: int32 clever_traps = 33;
     */
    cleverTraps: number;
    /**
     * @generated from protobuf field: int32 survivalist = 34;
     */
    survivalist: number;
    /**
     * @generated from protobuf field: int32 trap_mastery = 35;
     */
    trapMastery: number;
    /**
     * @generated from protobuf field: int32 surefooted = 36;
     */
    surefooted: number;
    /**
     * @generated from protobuf field: int32 survival_instincts = 37;
     */
    survivalInstincts: number;
    /**
     * @generated from protobuf field: int32 killer_instinct = 38;
     */
    killerInstinct: number;
    /**
     * @generated from protobuf field: int32 resourcefulness = 39;
     */
    resourcefulness: number;
    /**
     * @generated from protobuf field: int32 lightning_reflexes = 40;
     */
    lightningReflexes: number;
    /**
     * @generated from protobuf field: int32 thrill_of_the_hunt = 41;
     */
    thrillOfTheHunt: number;
    /**
     * @generated from protobuf field: int32 expose_weakness = 42;
     */
    exposeWeakness: number;
    /**
     * @generated from protobuf field: int32 master_tactician = 43;
     */
    masterTactician: number;
    /**
     * @generated from protobuf field: bool readiness = 44;
     */
    readiness: boolean;
}
/**
 * @generated from protobuf message proto.Hunter
 */
export interface Hunter {
    /**
     * @generated from protobuf field: proto.Hunter.Rotation rotation = 1;
     */
    rotation?: Hunter_Rotation;
    /**
     * @generated from protobuf field: proto.HunterTalents talents = 2;
     */
    talents?: HunterTalents;
    /**
     * @generated from protobuf field: proto.Hunter.Options options = 3;
     */
    options?: Hunter_Options;
}
/**
 * @generated from protobuf message proto.Hunter.Rotation
 */
export interface Hunter_Rotation {
}
/**
 * @generated from protobuf message proto.Hunter.Options
 */
export interface Hunter_Options {
}
declare class HunterTalents$Type extends MessageType<HunterTalents> {
    constructor();
    create(value?: PartialMessage<HunterTalents>): HunterTalents;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: HunterTalents): HunterTalents;
    internalBinaryWrite(message: HunterTalents, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.HunterTalents
 */
export declare const HunterTalents: HunterTalents$Type;
declare class Hunter$Type extends MessageType<Hunter> {
    constructor();
    create(value?: PartialMessage<Hunter>): Hunter;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Hunter): Hunter;
    internalBinaryWrite(message: Hunter, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Hunter
 */
export declare const Hunter: Hunter$Type;
declare class Hunter_Rotation$Type extends MessageType<Hunter_Rotation> {
    constructor();
    create(value?: PartialMessage<Hunter_Rotation>): Hunter_Rotation;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Hunter_Rotation): Hunter_Rotation;
    internalBinaryWrite(message: Hunter_Rotation, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Hunter.Rotation
 */
export declare const Hunter_Rotation: Hunter_Rotation$Type;
declare class Hunter_Options$Type extends MessageType<Hunter_Options> {
    constructor();
    create(value?: PartialMessage<Hunter_Options>): Hunter_Options;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Hunter_Options): Hunter_Options;
    internalBinaryWrite(message: Hunter_Options, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Hunter.Options
 */
export declare const Hunter_Options: Hunter_Options$Type;
export {};
