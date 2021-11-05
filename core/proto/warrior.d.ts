import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message proto.WarriorTalents
 */
export interface WarriorTalents {
    /**
     * Arms
     *
     * @generated from protobuf field: int32 improved_heroic_strike = 1;
     */
    improvedHeroicStrike: number;
    /**
     * @generated from protobuf field: int32 improved_rend = 2;
     */
    improvedRend: number;
    /**
     * @generated from protobuf field: int32 improved_charge = 3;
     */
    improvedCharge: number;
    /**
     * @generated from protobuf field: int32 improved_thunder_clap = 4;
     */
    improvedThunderClap: number;
    /**
     * @generated from protobuf field: int32 improved_overpower = 5;
     */
    improvedOverpower: number;
    /**
     * @generated from protobuf field: bool anger_management = 6;
     */
    angerManagement: boolean;
    /**
     * @generated from protobuf field: int32 deep_wounds = 7;
     */
    deepWounds: number;
    /**
     * @generated from protobuf field: int32 two_handed_weapon_specialization = 8;
     */
    twoHandedWeaponSpecialization: number;
    /**
     * @generated from protobuf field: int32 impale = 9;
     */
    impale: number;
    /**
     * @generated from protobuf field: int32 poleaxe_specialization = 10;
     */
    poleaxeSpecialization: number;
    /**
     * @generated from protobuf field: bool death_wish = 11;
     */
    deathWish: boolean;
    /**
     * @generated from protobuf field: int32 mace_specialization = 12;
     */
    maceSpecialization: number;
    /**
     * @generated from protobuf field: int32 sword_specialization = 13;
     */
    swordSpecialization: number;
    /**
     * @generated from protobuf field: int32 improved_disciplines = 14;
     */
    improvedDisciplines: number;
    /**
     * @generated from protobuf field: int32 blood_frenzy = 15;
     */
    bloodFrenzy: number;
    /**
     * @generated from protobuf field: bool mortal_strike = 16;
     */
    mortalStrike: boolean;
    /**
     * @generated from protobuf field: int32 improved_mortal_strike = 17;
     */
    improvedMortalStrike: number;
    /**
     * @generated from protobuf field: bool endless_rage = 18;
     */
    endlessRage: boolean;
    /**
     * Fury
     *
     * @generated from protobuf field: int32 booming_voice = 19;
     */
    boomingVoice: number;
    /**
     * @generated from protobuf field: int32 cruelty = 20;
     */
    cruelty: number;
    /**
     * @generated from protobuf field: int32 unbridled_wrath = 21;
     */
    unbridledWrath: number;
    /**
     * @generated from protobuf field: int32 improved_cleave = 22;
     */
    improvedCleave: number;
    /**
     * @generated from protobuf field: int32 commanding_presence = 23;
     */
    commandingPresence: number;
    /**
     * @generated from protobuf field: int32 dual_wield_specialization = 24;
     */
    dualWieldSpecialization: number;
    /**
     * @generated from protobuf field: int32 improved_execute = 25;
     */
    improvedExecute: number;
    /**
     * @generated from protobuf field: int32 improved_slam = 26;
     */
    improvedSlam: number;
    /**
     * @generated from protobuf field: bool sweeping_strikes = 27;
     */
    sweepingStrikes: boolean;
    /**
     * @generated from protobuf field: int32 weapon_mastery = 28;
     */
    weaponMastery: number;
    /**
     * @generated from protobuf field: int32 improved_berserker_rage = 29;
     */
    improvedBerserkerRage: number;
    /**
     * @generated from protobuf field: int32 flurry = 30;
     */
    flurry: number;
    /**
     * @generated from protobuf field: int32 precision = 31;
     */
    precision: number;
    /**
     * @generated from protobuf field: bool bloodthirst = 32;
     */
    bloodthirst: boolean;
    /**
     * @generated from protobuf field: int32 improved_whirlwind = 33;
     */
    improvedWhirlwind: number;
    /**
     * @generated from protobuf field: int32 improved_berserker_stance = 34;
     */
    improvedBerserkerStance: number;
    /**
     * @generated from protobuf field: bool rampage = 35;
     */
    rampage: boolean;
    /**
     * Protection
     *
     * @generated from protobuf field: int32 improved_bloodrage = 36;
     */
    improvedBloodrage: number;
    /**
     * @generated from protobuf field: int32 tactical_mastery = 37;
     */
    tacticalMastery: number;
    /**
     * @generated from protobuf field: int32 defiance = 38;
     */
    defiance: number;
    /**
     * @generated from protobuf field: int32 improved_sunder_armor = 39;
     */
    improvedSunderArmor: number;
    /**
     * @generated from protobuf field: int32 one_handed_weapon_specialization = 40;
     */
    oneHandedWeaponSpecialization: number;
    /**
     * @generated from protobuf field: bool shield_slam = 41;
     */
    shieldSlam: boolean;
    /**
     * @generated from protobuf field: int32 focused_rage = 42;
     */
    focusedRage: number;
    /**
     * @generated from protobuf field: int32 vitality = 43;
     */
    vitality: number;
    /**
     * @generated from protobuf field: bool devastate = 44;
     */
    devastate: boolean;
}
/**
 * @generated from protobuf message proto.Warrior
 */
export interface Warrior {
    /**
     * @generated from protobuf field: proto.Warrior.Rotation rotation = 1;
     */
    rotation?: Warrior_Rotation;
    /**
     * @generated from protobuf field: proto.WarriorTalents talents = 2;
     */
    talents?: WarriorTalents;
    /**
     * @generated from protobuf field: proto.Warrior.Options options = 3;
     */
    options?: Warrior_Options;
}
/**
 * @generated from protobuf message proto.Warrior.Rotation
 */
export interface Warrior_Rotation {
}
/**
 * @generated from protobuf message proto.Warrior.Options
 */
export interface Warrior_Options {
}
declare class WarriorTalents$Type extends MessageType<WarriorTalents> {
    constructor();
    create(value?: PartialMessage<WarriorTalents>): WarriorTalents;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: WarriorTalents): WarriorTalents;
    internalBinaryWrite(message: WarriorTalents, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.WarriorTalents
 */
export declare const WarriorTalents: WarriorTalents$Type;
declare class Warrior$Type extends MessageType<Warrior> {
    constructor();
    create(value?: PartialMessage<Warrior>): Warrior;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Warrior): Warrior;
    internalBinaryWrite(message: Warrior, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Warrior
 */
export declare const Warrior: Warrior$Type;
declare class Warrior_Rotation$Type extends MessageType<Warrior_Rotation> {
    constructor();
    create(value?: PartialMessage<Warrior_Rotation>): Warrior_Rotation;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Warrior_Rotation): Warrior_Rotation;
    internalBinaryWrite(message: Warrior_Rotation, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Warrior.Rotation
 */
export declare const Warrior_Rotation: Warrior_Rotation$Type;
declare class Warrior_Options$Type extends MessageType<Warrior_Options> {
    constructor();
    create(value?: PartialMessage<Warrior_Options>): Warrior_Options;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Warrior_Options): Warrior_Options;
    internalBinaryWrite(message: Warrior_Options, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Warrior.Options
 */
export declare const Warrior_Options: Warrior_Options$Type;
export {};
