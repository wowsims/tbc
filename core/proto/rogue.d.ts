import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message proto.RogueTalents
 */
export interface RogueTalents {
    /**
     * Assassination
     *
     * @generated from protobuf field: int32 improved_eviscerate = 1;
     */
    improvedEviscerate: number;
    /**
     * @generated from protobuf field: int32 malice = 2;
     */
    malice: number;
    /**
     * @generated from protobuf field: int32 ruthlessness = 3;
     */
    ruthlessness: number;
    /**
     * @generated from protobuf field: int32 murder = 4;
     */
    murder: number;
    /**
     * @generated from protobuf field: int32 puncturing_wounds = 5;
     */
    puncturingWounds: number;
    /**
     * @generated from protobuf field: bool relentless_strikes = 6;
     */
    relentlessStrikes: boolean;
    /**
     * @generated from protobuf field: int32 improved_expose_armor = 7;
     */
    improvedExposeArmor: number;
    /**
     * @generated from protobuf field: int32 lethality = 8;
     */
    lethality: number;
    /**
     * @generated from protobuf field: int32 vile_poisons = 9;
     */
    vilePoisons: number;
    /**
     * @generated from protobuf field: int32 improved_poisons = 10;
     */
    improvedPoisons: number;
    /**
     * @generated from protobuf field: bool cold_blood = 11;
     */
    coldBlood: boolean;
    /**
     * @generated from protobuf field: int32 quick_recovery = 12;
     */
    quickRecovery: number;
    /**
     * @generated from protobuf field: int32 seal_fate = 13;
     */
    sealFate: number;
    /**
     * @generated from protobuf field: int32 master_poisoner = 14;
     */
    masterPoisoner: number;
    /**
     * @generated from protobuf field: bool vigor = 15;
     */
    vigor: boolean;
    /**
     * @generated from protobuf field: int32 find_weakness = 16;
     */
    findWeakness: number;
    /**
     * @generated from protobuf field: bool mutilate = 17;
     */
    mutilate: boolean;
    /**
     * Combat
     *
     * @generated from protobuf field: int32 improved_sinister_strike = 18;
     */
    improvedSinisterStrike: number;
    /**
     * @generated from protobuf field: int32 improved_slice_and_dice = 19;
     */
    improvedSliceAndDice: number;
    /**
     * @generated from protobuf field: int32 precision = 20;
     */
    precision: number;
    /**
     * @generated from protobuf field: int32 dagger_specialization = 21;
     */
    daggerSpecialization: number;
    /**
     * @generated from protobuf field: int32 dual_wield_specialization = 22;
     */
    dualWieldSpecialization: number;
    /**
     * @generated from protobuf field: int32 mace_specialization = 23;
     */
    maceSpecialization: number;
    /**
     * @generated from protobuf field: bool blade_flurry = 24;
     */
    bladeFlurry: boolean;
    /**
     * @generated from protobuf field: int32 sword_specialization = 25;
     */
    swordSpecialization: number;
    /**
     * @generated from protobuf field: int32 fist_weapon_specialization = 26;
     */
    fistWeaponSpecialization: number;
    /**
     * @generated from protobuf field: int32 weapon_expertise = 27;
     */
    weaponExpertise: number;
    /**
     * @generated from protobuf field: int32 aggression = 28;
     */
    aggression: number;
    /**
     * @generated from protobuf field: int32 vitality = 29;
     */
    vitality: number;
    /**
     * @generated from protobuf field: bool adrenaline_rush = 30;
     */
    adrenalineRush: boolean;
    /**
     * @generated from protobuf field: int32 combat_potency = 31;
     */
    combatPotency: number;
    /**
     * @generated from protobuf field: bool surprise_attacks = 32;
     */
    surpriseAttacks: boolean;
    /**
     * Subtlety
     *
     * @generated from protobuf field: int32 opportunity = 33;
     */
    opportunity: number;
    /**
     * @generated from protobuf field: int32 sleight_of_hand = 46;
     */
    sleightOfHand: number;
    /**
     * @generated from protobuf field: int32 initiative = 34;
     */
    initiative: number;
    /**
     * @generated from protobuf field: bool ghostly_strike = 35;
     */
    ghostlyStrike: boolean;
    /**
     * @generated from protobuf field: int32 improved_ambush = 36;
     */
    improvedAmbush: number;
    /**
     * @generated from protobuf field: int32 elusiveness = 47;
     */
    elusiveness: number;
    /**
     * @generated from protobuf field: int32 serrated_blades = 37;
     */
    serratedBlades: number;
    /**
     * @generated from protobuf field: bool preparation = 38;
     */
    preparation: boolean;
    /**
     * @generated from protobuf field: int32 dirty_deeds = 39;
     */
    dirtyDeeds: number;
    /**
     * @generated from protobuf field: bool hemorrhage = 40;
     */
    hemorrhage: boolean;
    /**
     * @generated from protobuf field: int32 master_of_subtlety = 41;
     */
    masterOfSubtlety: number;
    /**
     * @generated from protobuf field: int32 deadliness = 42;
     */
    deadliness: number;
    /**
     * @generated from protobuf field: bool premeditation = 43;
     */
    premeditation: boolean;
    /**
     * @generated from protobuf field: bool sinister_calling = 44;
     */
    sinisterCalling: boolean;
    /**
     * @generated from protobuf field: bool shadowstep = 45;
     */
    shadowstep: boolean;
}
/**
 * @generated from protobuf message proto.Rogue
 */
export interface Rogue {
    /**
     * @generated from protobuf field: proto.Rogue.Rotation rotation = 1;
     */
    rotation?: Rogue_Rotation;
    /**
     * @generated from protobuf field: proto.RogueTalents talents = 2;
     */
    talents?: RogueTalents;
    /**
     * @generated from protobuf field: proto.Rogue.Options options = 3;
     */
    options?: Rogue_Options;
}
/**
 * @generated from protobuf message proto.Rogue.Rotation
 */
export interface Rogue_Rotation {
    /**
     * @generated from protobuf field: bool maintain_expose_armor = 1;
     */
    maintainExposeArmor: boolean;
}
/**
 * @generated from protobuf message proto.Rogue.Options
 */
export interface Rogue_Options {
}
declare class RogueTalents$Type extends MessageType<RogueTalents> {
    constructor();
    create(value?: PartialMessage<RogueTalents>): RogueTalents;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: RogueTalents): RogueTalents;
    internalBinaryWrite(message: RogueTalents, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.RogueTalents
 */
export declare const RogueTalents: RogueTalents$Type;
declare class Rogue$Type extends MessageType<Rogue> {
    constructor();
    create(value?: PartialMessage<Rogue>): Rogue;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Rogue): Rogue;
    internalBinaryWrite(message: Rogue, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Rogue
 */
export declare const Rogue: Rogue$Type;
declare class Rogue_Rotation$Type extends MessageType<Rogue_Rotation> {
    constructor();
    create(value?: PartialMessage<Rogue_Rotation>): Rogue_Rotation;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Rogue_Rotation): Rogue_Rotation;
    internalBinaryWrite(message: Rogue_Rotation, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Rogue.Rotation
 */
export declare const Rogue_Rotation: Rogue_Rotation$Type;
declare class Rogue_Options$Type extends MessageType<Rogue_Options> {
    constructor();
    create(value?: PartialMessage<Rogue_Options>): Rogue_Options;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Rogue_Options): Rogue_Options;
    internalBinaryWrite(message: Rogue_Options, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Rogue.Options
 */
export declare const Rogue_Options: Rogue_Options$Type;
export {};
