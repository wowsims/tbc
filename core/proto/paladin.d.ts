import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message proto.PaladinTalents
 */
export interface PaladinTalents {
    /**
     * Holy
     *
     * @generated from protobuf field: int32 divine_strength = 1;
     */
    divineStrength: number;
    /**
     * @generated from protobuf field: int32 divine_intellect = 2;
     */
    divineIntellect: number;
    /**
     * @generated from protobuf field: int32 improved_seal_of_righteousness = 3;
     */
    improvedSealOfRighteousness: number;
    /**
     * @generated from protobuf field: int32 illumination = 34;
     */
    illumination: number;
    /**
     * @generated from protobuf field: int32 improved_blessing_of_wisdom = 4;
     */
    improvedBlessingOfWisdom: number;
    /**
     * @generated from protobuf field: bool divine_favor = 5;
     */
    divineFavor: boolean;
    /**
     * @generated from protobuf field: int32 purifying_power = 6;
     */
    purifyingPower: number;
    /**
     * @generated from protobuf field: int32 holy_power = 7;
     */
    holyPower: number;
    /**
     * @generated from protobuf field: bool holy_shock = 8;
     */
    holyShock: boolean;
    /**
     * @generated from protobuf field: int32 blessed_life = 51;
     */
    blessedLife: number;
    /**
     * @generated from protobuf field: int32 holy_guidance = 9;
     */
    holyGuidance: number;
    /**
     * @generated from protobuf field: bool divine_illumination = 10;
     */
    divineIllumination: boolean;
    /**
     * Protection
     *
     * @generated from protobuf field: int32 improved_devotion_aura = 35;
     */
    improvedDevotionAura: number;
    /**
     * @generated from protobuf field: int32 redoubt = 36;
     */
    redoubt: number;
    /**
     * @generated from protobuf field: int32 precision = 11;
     */
    precision: number;
    /**
     * @generated from protobuf field: int32 toughness = 37;
     */
    toughness: number;
    /**
     * @generated from protobuf field: bool blessing_of_kings = 12;
     */
    blessingOfKings: boolean;
    /**
     * @generated from protobuf field: int32 improved_righteous_fury = 38;
     */
    improvedRighteousFury: number;
    /**
     * @generated from protobuf field: int32 shield_specialization = 39;
     */
    shieldSpecialization: number;
    /**
     * @generated from protobuf field: int32 anticipation = 40;
     */
    anticipation: number;
    /**
     * @generated from protobuf field: int32 spell_warding = 41;
     */
    spellWarding: number;
    /**
     * @generated from protobuf field: bool blessing_of_sanctuary = 42;
     */
    blessingOfSanctuary: boolean;
    /**
     * @generated from protobuf field: int32 reckoning = 13;
     */
    reckoning: number;
    /**
     * @generated from protobuf field: int32 sacred_duty = 14;
     */
    sacredDuty: number;
    /**
     * @generated from protobuf field: int32 one_handed_weapon_specialization = 15;
     */
    oneHandedWeaponSpecialization: number;
    /**
     * @generated from protobuf field: int32 improved_holy_shield = 43;
     */
    improvedHolyShield: number;
    /**
     * @generated from protobuf field: bool holy_shield = 44;
     */
    holyShield: boolean;
    /**
     * @generated from protobuf field: int32 ardent_defender = 45;
     */
    ardentDefender: number;
    /**
     * @generated from protobuf field: int32 combat_expertise = 16;
     */
    combatExpertise: number;
    /**
     * @generated from protobuf field: bool avengers_shield = 17;
     */
    avengersShield: boolean;
    /**
     * Retribution
     *
     * @generated from protobuf field: int32 improved_blessing_of_might = 18;
     */
    improvedBlessingOfMight: number;
    /**
     * @generated from protobuf field: int32 benediction = 19;
     */
    benediction: number;
    /**
     * @generated from protobuf field: int32 improved_judgement = 20;
     */
    improvedJudgement: number;
    /**
     * @generated from protobuf field: int32 improved_seal_of_the_crusader = 21;
     */
    improvedSealOfTheCrusader: number;
    /**
     * @generated from protobuf field: int32 deflection = 46;
     */
    deflection: number;
    /**
     * @generated from protobuf field: int32 vindication = 22;
     */
    vindication: number;
    /**
     * @generated from protobuf field: int32 conviction = 23;
     */
    conviction: number;
    /**
     * @generated from protobuf field: bool seal_of_command = 24;
     */
    sealOfCommand: boolean;
    /**
     * @generated from protobuf field: int32 pursuit_of_justice = 47;
     */
    pursuitOfJustice: number;
    /**
     * @generated from protobuf field: int32 eye_for_an_eye = 48;
     */
    eyeForAnEye: number;
    /**
     * @generated from protobuf field: int32 improved_retribution_aura = 49;
     */
    improvedRetributionAura: number;
    /**
     * @generated from protobuf field: int32 crusade = 25;
     */
    crusade: number;
    /**
     * @generated from protobuf field: int32 two_handed_weapon_specialization = 26;
     */
    twoHandedWeaponSpecialization: number;
    /**
     * @generated from protobuf field: bool sanctity_aura = 27;
     */
    sanctityAura: boolean;
    /**
     * @generated from protobuf field: int32 improved_sanctity_aura = 28;
     */
    improvedSanctityAura: number;
    /**
     * @generated from protobuf field: int32 vengeance = 29;
     */
    vengeance: number;
    /**
     * @generated from protobuf field: int32 sanctified_judgement = 30;
     */
    sanctifiedJudgement: number;
    /**
     * @generated from protobuf field: int32 sanctified_seals = 31;
     */
    sanctifiedSeals: number;
    /**
     * @generated from protobuf field: int32 divine_purpose = 50;
     */
    divinePurpose: number;
    /**
     * @generated from protobuf field: int32 fanaticism = 32;
     */
    fanaticism: number;
    /**
     * @generated from protobuf field: bool crusader_strike = 33;
     */
    crusaderStrike: boolean;
}
/**
 * @generated from protobuf message proto.RetributionPaladin
 */
export interface RetributionPaladin {
    /**
     * @generated from protobuf field: proto.RetributionPaladin.Rotation rotation = 1;
     */
    rotation?: RetributionPaladin_Rotation;
    /**
     * @generated from protobuf field: proto.PaladinTalents talents = 2;
     */
    talents?: PaladinTalents;
    /**
     * @generated from protobuf field: proto.RetributionPaladin.Options options = 3;
     */
    options?: RetributionPaladin_Options;
}
/**
 * @generated from protobuf message proto.RetributionPaladin.Rotation
 */
export interface RetributionPaladin_Rotation {
    /**
     * @generated from protobuf field: proto.RetributionPaladin.Rotation.ConsecrationRank consecration_rank = 1;
     */
    consecrationRank: RetributionPaladin_Rotation_ConsecrationRank;
    /**
     * @generated from protobuf field: bool use_exorcism = 2;
     */
    useExorcism: boolean;
}
/**
 * @generated from protobuf enum proto.RetributionPaladin.Rotation.ConsecrationRank
 */
export declare enum RetributionPaladin_Rotation_ConsecrationRank {
    /**
     * @generated from protobuf enum value: None = 0;
     */
    None = 0,
    /**
     * @generated from protobuf enum value: Rank1 = 1;
     */
    Rank1 = 1,
    /**
     * @generated from protobuf enum value: Rank4 = 2;
     */
    Rank4 = 2,
    /**
     * @generated from protobuf enum value: Rank6 = 3;
     */
    Rank6 = 3
}
/**
 * @generated from protobuf message proto.RetributionPaladin.Options
 */
export interface RetributionPaladin_Options {
    /**
     * @generated from protobuf field: proto.RetributionPaladin.Options.Judgement judgement = 1;
     */
    judgement: RetributionPaladin_Options_Judgement;
    /**
     * @generated from protobuf field: proto.PaladinAura aura = 5;
     */
    aura: PaladinAura;
    /**
     * @generated from protobuf field: int32 crusader_strike_delay_ms = 2;
     */
    crusaderStrikeDelayMs: number;
    /**
     * @generated from protobuf field: int32 haste_leeway_ms = 3;
     */
    hasteLeewayMs: number;
    /**
     * @generated from protobuf field: double damage_taken_per_second = 4;
     */
    damageTakenPerSecond: number;
}
/**
 * @generated from protobuf enum proto.RetributionPaladin.Options.Judgement
 */
export declare enum RetributionPaladin_Options_Judgement {
    /**
     * @generated from protobuf enum value: None = 0;
     */
    None = 0,
    /**
     * @generated from protobuf enum value: Wisdom = 1;
     */
    Wisdom = 1,
    /**
     * @generated from protobuf enum value: Crusader = 2;
     */
    Crusader = 2
}
/**
 * @generated from protobuf enum proto.PaladinAura
 */
export declare enum PaladinAura {
    /**
     * @generated from protobuf enum value: NoPaladinAura = 0;
     */
    NoPaladinAura = 0,
    /**
     * @generated from protobuf enum value: SanctityAura = 1;
     */
    SanctityAura = 1,
    /**
     * @generated from protobuf enum value: DevotionAura = 2;
     */
    DevotionAura = 2,
    /**
     * @generated from protobuf enum value: RetributionAura = 3;
     */
    RetributionAura = 3
}
declare class PaladinTalents$Type extends MessageType<PaladinTalents> {
    constructor();
    create(value?: PartialMessage<PaladinTalents>): PaladinTalents;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PaladinTalents): PaladinTalents;
    internalBinaryWrite(message: PaladinTalents, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.PaladinTalents
 */
export declare const PaladinTalents: PaladinTalents$Type;
declare class RetributionPaladin$Type extends MessageType<RetributionPaladin> {
    constructor();
    create(value?: PartialMessage<RetributionPaladin>): RetributionPaladin;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: RetributionPaladin): RetributionPaladin;
    internalBinaryWrite(message: RetributionPaladin, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.RetributionPaladin
 */
export declare const RetributionPaladin: RetributionPaladin$Type;
declare class RetributionPaladin_Rotation$Type extends MessageType<RetributionPaladin_Rotation> {
    constructor();
    create(value?: PartialMessage<RetributionPaladin_Rotation>): RetributionPaladin_Rotation;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: RetributionPaladin_Rotation): RetributionPaladin_Rotation;
    internalBinaryWrite(message: RetributionPaladin_Rotation, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.RetributionPaladin.Rotation
 */
export declare const RetributionPaladin_Rotation: RetributionPaladin_Rotation$Type;
declare class RetributionPaladin_Options$Type extends MessageType<RetributionPaladin_Options> {
    constructor();
    create(value?: PartialMessage<RetributionPaladin_Options>): RetributionPaladin_Options;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: RetributionPaladin_Options): RetributionPaladin_Options;
    internalBinaryWrite(message: RetributionPaladin_Options, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.RetributionPaladin.Options
 */
export declare const RetributionPaladin_Options: RetributionPaladin_Options$Type;
export {};
