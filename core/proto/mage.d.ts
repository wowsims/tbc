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
     * @generated from protobuf field: int32 magic_absorption = 48;
     */
    magicAbsorption: number;
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
     * @generated from protobuf field: int32 empowered_arcane_missiles = 11;
     */
    empoweredArcaneMissiles: number;
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
     * @generated from protobuf field: int32 burning_soul = 47;
     */
    burningSoul: number;
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
    /**
     * @generated from protobuf field: proto.Mage.Rotation.Type type = 1;
     */
    type: Mage_Rotation_Type;
    /**
     * @generated from protobuf field: proto.Mage.Rotation.ArcaneRotation arcane = 2;
     */
    arcane?: Mage_Rotation_ArcaneRotation;
    /**
     * @generated from protobuf field: proto.Mage.Rotation.FireRotation fire = 3;
     */
    fire?: Mage_Rotation_FireRotation;
    /**
     * @generated from protobuf field: proto.Mage.Rotation.FrostRotation frost = 4;
     */
    frost?: Mage_Rotation_FrostRotation;
    /**
     * @generated from protobuf field: proto.Mage.Rotation.AoeRotation aoe = 5;
     */
    aoe?: Mage_Rotation_AoeRotation;
    /**
     * @generated from protobuf field: bool multi_target_rotation = 6;
     */
    multiTargetRotation: boolean;
}
/**
 * @generated from protobuf message proto.Mage.Rotation.ArcaneRotation
 */
export interface Mage_Rotation_ArcaneRotation {
    /**
     * The spells to use to fill time while waiting for arcane blast stacks to drop.
     *
     * @generated from protobuf field: proto.Mage.Rotation.ArcaneRotation.Filler filler = 1;
     */
    filler: Mage_Rotation_ArcaneRotation_Filler;
    /**
     * Number of arcane blasts to cast before switching to filler.
     *
     * @generated from protobuf field: int32 arcane_blasts_between_fillers = 2;
     */
    arcaneBlastsBetweenFillers: number;
    /**
     * Percentage of mana (0-1) below which to switch to regen rotation.
     *
     * @generated from protobuf field: double start_regen_rotation_percent = 3;
     */
    startRegenRotationPercent: number;
    /**
     * Percentage of mana (0-1) above which to switch to regular rotation.
     *
     * @generated from protobuf field: double stop_regen_rotation_percent = 4;
     */
    stopRegenRotationPercent: number;
    /**
     * Prevents DPS cooldowns from being using during regen rotation.
     *
     * @generated from protobuf field: bool disable_dps_cooldowns_during_regen = 5;
     */
    disableDpsCooldownsDuringRegen: boolean;
}
/**
 * @generated from protobuf enum proto.Mage.Rotation.ArcaneRotation.Filler
 */
export declare enum Mage_Rotation_ArcaneRotation_Filler {
    /**
     * @generated from protobuf enum value: Frostbolt = 0;
     */
    Frostbolt = 0,
    /**
     * @generated from protobuf enum value: ArcaneMissiles = 1;
     */
    ArcaneMissiles = 1,
    /**
     * @generated from protobuf enum value: Scorch = 2;
     */
    Scorch = 2,
    /**
     * @generated from protobuf enum value: Fireball = 3;
     */
    Fireball = 3,
    /**
     * @generated from protobuf enum value: ArcaneMissilesFrostbolt = 4;
     */
    ArcaneMissilesFrostbolt = 4,
    /**
     * @generated from protobuf enum value: ArcaneMissilesScorch = 5;
     */
    ArcaneMissilesScorch = 5,
    /**
     * @generated from protobuf enum value: ScorchTwoFireball = 6;
     */
    ScorchTwoFireball = 6
}
/**
 * @generated from protobuf message proto.Mage.Rotation.FireRotation
 */
export interface Mage_Rotation_FireRotation {
    /**
     * @generated from protobuf field: proto.Mage.Rotation.FireRotation.PrimarySpell primary_spell = 1;
     */
    primarySpell: Mage_Rotation_FireRotation_PrimarySpell;
    /**
     * @generated from protobuf field: bool maintain_improved_scorch = 2;
     */
    maintainImprovedScorch: boolean;
    /**
     * @generated from protobuf field: bool weave_fire_blast = 3;
     */
    weaveFireBlast: boolean;
}
/**
 * @generated from protobuf enum proto.Mage.Rotation.FireRotation.PrimarySpell
 */
export declare enum Mage_Rotation_FireRotation_PrimarySpell {
    /**
     * @generated from protobuf enum value: Fireball = 0;
     */
    Fireball = 0,
    /**
     * @generated from protobuf enum value: Scorch = 1;
     */
    Scorch = 1
}
/**
 * @generated from protobuf message proto.Mage.Rotation.FrostRotation
 */
export interface Mage_Rotation_FrostRotation {
    /**
     * Chance for water elemental to disobey, doing nothing rather than cast.
     *
     * @generated from protobuf field: double water_elemental_disobey_chance = 3;
     */
    waterElementalDisobeyChance: number;
}
/**
 * @generated from protobuf message proto.Mage.Rotation.AoeRotation
 */
export interface Mage_Rotation_AoeRotation {
    /**
     * @generated from protobuf field: proto.Mage.Rotation.AoeRotation.Rotation rotation = 1;
     */
    rotation: Mage_Rotation_AoeRotation_Rotation;
}
/**
 * @generated from protobuf enum proto.Mage.Rotation.AoeRotation.Rotation
 */
export declare enum Mage_Rotation_AoeRotation_Rotation {
    /**
     * @generated from protobuf enum value: ArcaneExplosion = 0;
     */
    ArcaneExplosion = 0,
    /**
     * @generated from protobuf enum value: Flamestrike = 1;
     */
    Flamestrike = 1,
    /**
     * @generated from protobuf enum value: Blizzard = 2;
     */
    Blizzard = 2
}
/**
 * Just used for controlling which options are displayed in the UI. Is not
 * used by the sim.
 *
 * @generated from protobuf enum proto.Mage.Rotation.Type
 */
export declare enum Mage_Rotation_Type {
    /**
     * @generated from protobuf enum value: Arcane = 0;
     */
    Arcane = 0,
    /**
     * @generated from protobuf enum value: Fire = 1;
     */
    Fire = 1,
    /**
     * @generated from protobuf enum value: Frost = 2;
     */
    Frost = 2
}
/**
 * @generated from protobuf message proto.Mage.Options
 */
export interface Mage_Options {
    /**
     * @generated from protobuf field: proto.Mage.Options.ArmorType armor = 1;
     */
    armor: Mage_Options_ArmorType;
    /**
     * Number of Evocation ticks to use. If 0, use all of them.
     *
     * @generated from protobuf field: int32 evocation_ticks = 2;
     */
    evocationTicks: number;
}
/**
 * @generated from protobuf enum proto.Mage.Options.ArmorType
 */
export declare enum Mage_Options_ArmorType {
    /**
     * @generated from protobuf enum value: NoArmor = 0;
     */
    NoArmor = 0,
    /**
     * @generated from protobuf enum value: MageArmor = 1;
     */
    MageArmor = 1,
    /**
     * @generated from protobuf enum value: MoltenArmor = 2;
     */
    MoltenArmor = 2
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
declare class Mage_Rotation_ArcaneRotation$Type extends MessageType<Mage_Rotation_ArcaneRotation> {
    constructor();
    create(value?: PartialMessage<Mage_Rotation_ArcaneRotation>): Mage_Rotation_ArcaneRotation;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Mage_Rotation_ArcaneRotation): Mage_Rotation_ArcaneRotation;
    internalBinaryWrite(message: Mage_Rotation_ArcaneRotation, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Mage.Rotation.ArcaneRotation
 */
export declare const Mage_Rotation_ArcaneRotation: Mage_Rotation_ArcaneRotation$Type;
declare class Mage_Rotation_FireRotation$Type extends MessageType<Mage_Rotation_FireRotation> {
    constructor();
    create(value?: PartialMessage<Mage_Rotation_FireRotation>): Mage_Rotation_FireRotation;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Mage_Rotation_FireRotation): Mage_Rotation_FireRotation;
    internalBinaryWrite(message: Mage_Rotation_FireRotation, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Mage.Rotation.FireRotation
 */
export declare const Mage_Rotation_FireRotation: Mage_Rotation_FireRotation$Type;
declare class Mage_Rotation_FrostRotation$Type extends MessageType<Mage_Rotation_FrostRotation> {
    constructor();
    create(value?: PartialMessage<Mage_Rotation_FrostRotation>): Mage_Rotation_FrostRotation;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Mage_Rotation_FrostRotation): Mage_Rotation_FrostRotation;
    internalBinaryWrite(message: Mage_Rotation_FrostRotation, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Mage.Rotation.FrostRotation
 */
export declare const Mage_Rotation_FrostRotation: Mage_Rotation_FrostRotation$Type;
declare class Mage_Rotation_AoeRotation$Type extends MessageType<Mage_Rotation_AoeRotation> {
    constructor();
    create(value?: PartialMessage<Mage_Rotation_AoeRotation>): Mage_Rotation_AoeRotation;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Mage_Rotation_AoeRotation): Mage_Rotation_AoeRotation;
    internalBinaryWrite(message: Mage_Rotation_AoeRotation, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Mage.Rotation.AoeRotation
 */
export declare const Mage_Rotation_AoeRotation: Mage_Rotation_AoeRotation$Type;
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
