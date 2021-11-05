import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message proto.MageTalents
 */
export interface MageTalents {
    /**
     * Arcane
     *
     * @generated from protobuf field: int32 arcane_subtlety = 1;
     */
    arcaneSubtlety: number;
    /**
     * @generated from protobuf field: int32 arcane_focus = 2;
     */
    arcaneFocus: number;
    /**
     * @generated from protobuf field: int32 wand_specialization = 3;
     */
    wandSpecialization: number;
    /**
     * @generated from protobuf field: int32 arcane_concentration = 4;
     */
    arcaneConcentration: number;
    /**
     * @generated from protobuf field: int32 arcane_impact = 5;
     */
    arcaneImpact: number;
    /**
     * @generated from protobuf field: int32 arcane_meditation = 6;
     */
    arcaneMeditation: number;
    /**
     * @generated from protobuf field: bool presence_of_mind = 7;
     */
    presenceOfMind: boolean;
    /**
     * @generated from protobuf field: int32 arcane_mind = 8;
     */
    arcaneMind: number;
    /**
     * @generated from protobuf field: int32 arcane_instability = 9;
     */
    arcaneInstability: number;
    /**
     * @generated from protobuf field: int32 arcane_potency = 10;
     */
    arcanePotency: number;
    /**
     * @generated from protobuf field: int32 empowered_arcane_missles = 11;
     */
    empoweredArcaneMissles: number;
    /**
     * @generated from protobuf field: bool arcane_power = 12;
     */
    arcanePower: boolean;
    /**
     * @generated from protobuf field: int32 spell_power = 13;
     */
    spellPower: number;
    /**
     * @generated from protobuf field: int32 mind_mastery = 14;
     */
    mindMastery: number;
    /**
     * Fire
     *
     * @generated from protobuf field: int32 improved_fireball = 15;
     */
    improvedFireball: number;
    /**
     * @generated from protobuf field: int32 ignite = 16;
     */
    ignite: number;
    /**
     * @generated from protobuf field: int32 improved_fire_blast = 17;
     */
    improvedFireBlast: number;
    /**
     * @generated from protobuf field: int32 incineration = 18;
     */
    incineration: number;
    /**
     * @generated from protobuf field: int32 improved_flamestrike = 19;
     */
    improvedFlamestrike: number;
    /**
     * @generated from protobuf field: bool pyroblast = 20;
     */
    pyroblast: boolean;
    /**
     * @generated from protobuf field: int32 improved_scorch = 21;
     */
    improvedScorch: number;
    /**
     * @generated from protobuf field: int32 master_of_elements = 22;
     */
    masterOfElements: number;
    /**
     * @generated from protobuf field: int32 playing_with_fire = 23;
     */
    playingWithFire: number;
    /**
     * @generated from protobuf field: int32 critical_mass = 24;
     */
    criticalMass: number;
    /**
     * @generated from protobuf field: bool blast_wave = 25;
     */
    blastWave: boolean;
    /**
     * @generated from protobuf field: int32 fire_power = 26;
     */
    firePower: number;
    /**
     * @generated from protobuf field: int32 pyromaniac = 27;
     */
    pyromaniac: number;
    /**
     * @generated from protobuf field: bool combustion = 28;
     */
    combustion: boolean;
    /**
     * @generated from protobuf field: int32 molten_fury = 29;
     */
    moltenFury: number;
    /**
     * @generated from protobuf field: int32 empowered_fireball = 30;
     */
    empoweredFireball: number;
    /**
     * @generated from protobuf field: bool dragons_breath = 31;
     */
    dragonsBreath: boolean;
    /**
     * Frost
     *
     * @generated from protobuf field: int32 improved_frostbolt = 32;
     */
    improvedFrostbolt: number;
    /**
     * @generated from protobuf field: int32 elemental_precision = 33;
     */
    elementalPrecision: number;
    /**
     * @generated from protobuf field: int32 ice_shards = 34;
     */
    iceShards: number;
    /**
     * @generated from protobuf field: int32 improved_frost_nova = 35;
     */
    improvedFrostNova: number;
    /**
     * @generated from protobuf field: int32 piercing_ice = 36;
     */
    piercingIce: number;
    /**
     * @generated from protobuf field: bool icy_veins = 37;
     */
    icyVeins: boolean;
    /**
     * @generated from protobuf field: int32 frost_channeling = 38;
     */
    frostChanneling: number;
    /**
     * @generated from protobuf field: int32 shatter = 39;
     */
    shatter: number;
    /**
     * @generated from protobuf field: bool cold_snap = 40;
     */
    coldSnap: boolean;
    /**
     * @generated from protobuf field: int32 improved_cone_of_cold = 41;
     */
    improvedConeOfCold: number;
    /**
     * @generated from protobuf field: int32 ice_floes = 42;
     */
    iceFloes: number;
    /**
     * @generated from protobuf field: int32 winters_chill = 43;
     */
    wintersChill: number;
    /**
     * @generated from protobuf field: int32 arctic_winds = 44;
     */
    arcticWinds: number;
    /**
     * @generated from protobuf field: int32 empowered_frostbolt = 45;
     */
    empoweredFrostbolt: number;
    /**
     * @generated from protobuf field: bool summon_water_elemental = 46;
     */
    summonWaterElemental: boolean;
}
/**
 * @generated from protobuf message proto.Mage
 */
export interface Mage {
    /**
     * @generated from protobuf field: proto.Mage.Rotation rotation = 1;
     */
    rotation?: Mage_Rotation;
    /**
     * @generated from protobuf field: proto.MageTalents talents = 2;
     */
    talents?: MageTalents;
    /**
     * @generated from protobuf field: proto.Mage.Options options = 3;
     */
    options?: Mage_Options;
}
/**
 * @generated from protobuf message proto.Mage.Rotation
 */
export interface Mage_Rotation {
}
/**
 * @generated from protobuf message proto.Mage.Options
 */
export interface Mage_Options {
}
declare class MageTalents$Type extends MessageType<MageTalents> {
    constructor();
    create(value?: PartialMessage<MageTalents>): MageTalents;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: MageTalents): MageTalents;
    internalBinaryWrite(message: MageTalents, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.MageTalents
 */
export declare const MageTalents: MageTalents$Type;
declare class Mage$Type extends MessageType<Mage> {
    constructor();
    create(value?: PartialMessage<Mage>): Mage;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Mage): Mage;
    internalBinaryWrite(message: Mage, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Mage
 */
export declare const Mage: Mage$Type;
declare class Mage_Rotation$Type extends MessageType<Mage_Rotation> {
    constructor();
    create(value?: PartialMessage<Mage_Rotation>): Mage_Rotation;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Mage_Rotation): Mage_Rotation;
    internalBinaryWrite(message: Mage_Rotation, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Mage.Rotation
 */
export declare const Mage_Rotation: Mage_Rotation$Type;
declare class Mage_Options$Type extends MessageType<Mage_Options> {
    constructor();
    create(value?: PartialMessage<Mage_Options>): Mage_Options;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Mage_Options): Mage_Options;
    internalBinaryWrite(message: Mage_Options, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Mage.Options
 */
export declare const Mage_Options: Mage_Options$Type;
export {};
