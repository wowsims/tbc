import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message proto.ShamanTalents
 */
export interface ShamanTalents {
    /**
     * Elemental
     *
     * @generated from protobuf field: int32 convection = 1;
     */
    convection: number;
    /**
     * @generated from protobuf field: int32 concussion = 2;
     */
    concussion: number;
    /**
     * @generated from protobuf field: int32 call_of_flame = 3;
     */
    callOfFlame: number;
    /**
     * @generated from protobuf field: bool elemental_focus = 4;
     */
    elementalFocus: boolean;
    /**
     * @generated from protobuf field: int32 reverberation = 5;
     */
    reverberation: number;
    /**
     * @generated from protobuf field: int32 call_of_thunder = 6;
     */
    callOfThunder: number;
    /**
     * @generated from protobuf field: int32 improved_fire_totems = 7;
     */
    improvedFireTotems: number;
    /**
     * @generated from protobuf field: int32 elemental_devastation = 8;
     */
    elementalDevastation: number;
    /**
     * @generated from protobuf field: bool elemental_fury = 9;
     */
    elementalFury: boolean;
    /**
     * @generated from protobuf field: int32 unrelenting_storm = 10;
     */
    unrelentingStorm: number;
    /**
     * @generated from protobuf field: int32 elemental_precision = 11;
     */
    elementalPrecision: number;
    /**
     * @generated from protobuf field: int32 lightning_mastery = 12;
     */
    lightningMastery: number;
    /**
     * @generated from protobuf field: bool elemental_mastery = 13;
     */
    elementalMastery: boolean;
    /**
     * @generated from protobuf field: int32 lightning_overload = 14;
     */
    lightningOverload: number;
    /**
     * @generated from protobuf field: bool totemOfWrath = 33;
     */
    totemOfWrath: boolean;
    /**
     * Enhancement
     *
     * @generated from protobuf field: int32 ancestral_knowledge = 15;
     */
    ancestralKnowledge: number;
    /**
     * @generated from protobuf field: int32 shield_specialization = 37;
     */
    shieldSpecialization: number;
    /**
     * @generated from protobuf field: int32 thundering_strikes = 16;
     */
    thunderingStrikes: number;
    /**
     * @generated from protobuf field: int32 enhancing_totems = 17;
     */
    enhancingTotems: number;
    /**
     * @generated from protobuf field: bool shamanistic_focus = 18;
     */
    shamanisticFocus: boolean;
    /**
     * @generated from protobuf field: int32 anticipation = 38;
     */
    anticipation: number;
    /**
     * @generated from protobuf field: int32 flurry = 19;
     */
    flurry: number;
    /**
     * @generated from protobuf field: int32 toughness = 39;
     */
    toughness: number;
    /**
     * @generated from protobuf field: int32 improved_weapon_totems = 20;
     */
    improvedWeaponTotems: number;
    /**
     * @generated from protobuf field: bool spirit_weapons = 36;
     */
    spiritWeapons: boolean;
    /**
     * @generated from protobuf field: int32 elemental_weapons = 21;
     */
    elementalWeapons: number;
    /**
     * @generated from protobuf field: int32 mental_quickness = 22;
     */
    mentalQuickness: number;
    /**
     * @generated from protobuf field: int32 weapon_mastery = 23;
     */
    weaponMastery: number;
    /**
     * @generated from protobuf field: int32 dual_wield_specialization = 24;
     */
    dualWieldSpecialization: number;
    /**
     * @generated from protobuf field: int32 unleashed_rage = 25;
     */
    unleashedRage: number;
    /**
     * @generated from protobuf field: bool stormstrike = 34;
     */
    stormstrike: boolean;
    /**
     * @generated from protobuf field: bool shamanistic_rage = 35;
     */
    shamanisticRage: boolean;
    /**
     * Restoration
     *
     * @generated from protobuf field: int32 totemic_focus = 26;
     */
    totemicFocus: number;
    /**
     * @generated from protobuf field: int32 natures_guidance = 27;
     */
    naturesGuidance: number;
    /**
     * @generated from protobuf field: int32 restorative_totems = 28;
     */
    restorativeTotems: number;
    /**
     * @generated from protobuf field: int32 tidal_mastery = 29;
     */
    tidalMastery: number;
    /**
     * @generated from protobuf field: bool natures_swiftness = 30;
     */
    naturesSwiftness: boolean;
    /**
     * @generated from protobuf field: bool mana_tide_totem = 31;
     */
    manaTideTotem: boolean;
    /**
     * @generated from protobuf field: int32 natures_blessing = 32;
     */
    naturesBlessing: number;
}
/**
 * @generated from protobuf message proto.ShamanTotems
 */
export interface ShamanTotems {
    /**
     * @generated from protobuf field: proto.EarthTotem earth = 1;
     */
    earth: EarthTotem;
    /**
     * @generated from protobuf field: proto.AirTotem air = 2;
     */
    air: AirTotem;
    /**
     * @generated from protobuf field: proto.FireTotem fire = 3;
     */
    fire: FireTotem;
    /**
     * @generated from protobuf field: proto.WaterTotem water = 4;
     */
    water: WaterTotem;
    /**
     * If set, will twist windfury with whichever air totem is selected.
     *
     * @generated from protobuf field: bool twist_windfury = 5;
     */
    twistWindfury: boolean;
    /**
     * Rank of Windfury Totem to cast, if using Windfury Totem.
     *
     * @generated from protobuf field: int32 windfury_totem_rank = 11;
     */
    windfuryTotemRank: number;
    /**
     * If set, will twist fire nova with whichever fire totem is selected.
     *
     * @generated from protobuf field: bool twist_fire_nova = 6;
     */
    twistFireNova: boolean;
    /**
     * If set, will use mana tide when appropriate.
     *
     * @generated from protobuf field: bool use_mana_tide = 7;
     */
    useManaTide: boolean;
    /**
     * If set, will use fire elemental totem at the start and revert to regular
     * fire totems when it expires. If fire nova is also selected, fire nova
     * will be used once before dropping fire ele.
     *
     * @generated from protobuf field: bool use_fire_elemental = 8;
     */
    useFireElemental: boolean;
    /**
     * If set, will revert to regular fire totems when fire elemental goes OOM,
     * instead of waiting the full 2 minutes.
     *
     * @generated from protobuf field: bool recall_fire_elemental_on_oom = 9;
     */
    recallFireElementalOnOom: boolean;
    /**
     * If set, any time a 2-minute totem is about to expire, will recall and
     * replace all totems.
     *
     * @generated from protobuf field: bool recall_totems = 10;
     */
    recallTotems: boolean;
}
/**
 * @generated from protobuf message proto.ElementalShaman
 */
export interface ElementalShaman {
    /**
     * @generated from protobuf field: proto.ElementalShaman.Rotation rotation = 1;
     */
    rotation?: ElementalShaman_Rotation;
    /**
     * @generated from protobuf field: proto.ShamanTalents talents = 2;
     */
    talents?: ShamanTalents;
    /**
     * @generated from protobuf field: proto.ElementalShaman.Options options = 3;
     */
    options?: ElementalShaman_Options;
}
/**
 * @generated from protobuf message proto.ElementalShaman.Rotation
 */
export interface ElementalShaman_Rotation {
    /**
     * @generated from protobuf field: proto.ShamanTotems totems = 3;
     */
    totems?: ShamanTotems;
    /**
     * @generated from protobuf field: proto.ElementalShaman.Rotation.RotationType type = 1;
     */
    type: ElementalShaman_Rotation_RotationType;
    /**
     * Only used if type == FixedLBCL
     *
     * @generated from protobuf field: int32 lbs_per_cl = 2;
     */
    lbsPerCl: number;
}
/**
 * @generated from protobuf enum proto.ElementalShaman.Rotation.RotationType
 */
export declare enum ElementalShaman_Rotation_RotationType {
    /**
     * @generated from protobuf enum value: Unknown = 0;
     */
    Unknown = 0,
    /**
     * @generated from protobuf enum value: Adaptive = 1;
     */
    Adaptive = 1,
    /**
     * @generated from protobuf enum value: CLOnClearcast = 2;
     */
    CLOnClearcast = 2,
    /**
     * @generated from protobuf enum value: CLOnCD = 3;
     */
    CLOnCD = 3,
    /**
     * @generated from protobuf enum value: FixedLBCL = 4;
     */
    FixedLBCL = 4,
    /**
     * @generated from protobuf enum value: LBOnly = 5;
     */
    LBOnly = 5
}
/**
 * @generated from protobuf message proto.ElementalShaman.Options
 */
export interface ElementalShaman_Options {
    /**
     * @generated from protobuf field: bool water_shield = 1;
     */
    waterShield: boolean;
    /**
     * @generated from protobuf field: bool bloodlust = 2;
     */
    bloodlust: boolean;
    /**
     * Indicates the shaman will be dropping an improved wrath of air totem before
     * the fight by snapshotting the T4 2pc bonus.
     *
     * @generated from protobuf field: bool snapshot_t4_2pc = 6 [json_name = "snapshotT42pc"];
     */
    snapshotT42Pc: boolean;
}
/**
 * @generated from protobuf message proto.EnhancementShaman
 */
export interface EnhancementShaman {
    /**
     * @generated from protobuf field: proto.EnhancementShaman.Rotation rotation = 1;
     */
    rotation?: EnhancementShaman_Rotation;
    /**
     * @generated from protobuf field: proto.ShamanTalents talents = 2;
     */
    talents?: ShamanTalents;
    /**
     * @generated from protobuf field: proto.EnhancementShaman.Options options = 3;
     */
    options?: EnhancementShaman_Options;
}
/**
 * @generated from protobuf message proto.EnhancementShaman.Rotation
 */
export interface EnhancementShaman_Rotation {
    /**
     * @generated from protobuf field: proto.ShamanTotems totems = 1;
     */
    totems?: ShamanTotems;
    /**
     * @generated from protobuf field: proto.EnhancementShaman.Rotation.PrimaryShock primary_shock = 2;
     */
    primaryShock: EnhancementShaman_Rotation_PrimaryShock;
    /**
     * @generated from protobuf field: bool weave_flame_shock = 3;
     */
    weaveFlameShock: boolean;
    /**
     * For internal use only. Use to stagger SS casts between multiple Enhance
     * Shaman to optimize SS charge usage.
     *
     * @generated from protobuf field: double first_stormstrike_delay = 4;
     */
    firstStormstrikeDelay: number;
}
/**
 * @generated from protobuf enum proto.EnhancementShaman.Rotation.PrimaryShock
 */
export declare enum EnhancementShaman_Rotation_PrimaryShock {
    /**
     * @generated from protobuf enum value: None = 0;
     */
    None = 0,
    /**
     * @generated from protobuf enum value: Earth = 1;
     */
    Earth = 1,
    /**
     * @generated from protobuf enum value: Frost = 2;
     */
    Frost = 2
}
/**
 * @generated from protobuf message proto.EnhancementShaman.Options
 */
export interface EnhancementShaman_Options {
    /**
     * @generated from protobuf field: bool water_shield = 1;
     */
    waterShield: boolean;
    /**
     * @generated from protobuf field: bool bloodlust = 2;
     */
    bloodlust: boolean;
    /**
     * @generated from protobuf field: bool delay_offhand_swings = 5;
     */
    delayOffhandSwings: boolean;
    /**
     * Indicates the shaman will be dropping an improved strength of earth totem before
     * the fight by snapshotting the T4 2pc bonus.
     *
     * @generated from protobuf field: bool snapshot_t4_2pc = 6 [json_name = "snapshotT42pc"];
     */
    snapshotT42Pc: boolean;
}
/**
 * @generated from protobuf enum proto.EarthTotem
 */
export declare enum EarthTotem {
    /**
     * @generated from protobuf enum value: NoEarthTotem = 0;
     */
    NoEarthTotem = 0,
    /**
     * @generated from protobuf enum value: StrengthOfEarthTotem = 1;
     */
    StrengthOfEarthTotem = 1,
    /**
     * @generated from protobuf enum value: TremorTotem = 2;
     */
    TremorTotem = 2
}
/**
 * @generated from protobuf enum proto.AirTotem
 */
export declare enum AirTotem {
    /**
     * @generated from protobuf enum value: NoAirTotem = 0;
     */
    NoAirTotem = 0,
    /**
     * @generated from protobuf enum value: GraceOfAirTotem = 1;
     */
    GraceOfAirTotem = 1,
    /**
     * @generated from protobuf enum value: TranquilAirTotem = 2;
     */
    TranquilAirTotem = 2,
    /**
     * @generated from protobuf enum value: WindfuryTotem = 3;
     */
    WindfuryTotem = 3,
    /**
     * @generated from protobuf enum value: WrathOfAirTotem = 4;
     */
    WrathOfAirTotem = 4
}
/**
 * @generated from protobuf enum proto.FireTotem
 */
export declare enum FireTotem {
    /**
     * @generated from protobuf enum value: NoFireTotem = 0;
     */
    NoFireTotem = 0,
    /**
     * @generated from protobuf enum value: MagmaTotem = 1;
     */
    MagmaTotem = 1,
    /**
     * @generated from protobuf enum value: SearingTotem = 2;
     */
    SearingTotem = 2,
    /**
     * @generated from protobuf enum value: TotemOfWrath = 3;
     */
    TotemOfWrath = 3,
    /**
     * @generated from protobuf enum value: FireNovaTotem = 4;
     */
    FireNovaTotem = 4
}
/**
 * @generated from protobuf enum proto.WaterTotem
 */
export declare enum WaterTotem {
    /**
     * @generated from protobuf enum value: NoWaterTotem = 0;
     */
    NoWaterTotem = 0,
    /**
     * @generated from protobuf enum value: ManaSpringTotem = 1;
     */
    ManaSpringTotem = 1
}
declare class ShamanTalents$Type extends MessageType<ShamanTalents> {
    constructor();
    create(value?: PartialMessage<ShamanTalents>): ShamanTalents;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ShamanTalents): ShamanTalents;
    internalBinaryWrite(message: ShamanTalents, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ShamanTalents
 */
export declare const ShamanTalents: ShamanTalents$Type;
declare class ShamanTotems$Type extends MessageType<ShamanTotems> {
    constructor();
    create(value?: PartialMessage<ShamanTotems>): ShamanTotems;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ShamanTotems): ShamanTotems;
    internalBinaryWrite(message: ShamanTotems, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ShamanTotems
 */
export declare const ShamanTotems: ShamanTotems$Type;
declare class ElementalShaman$Type extends MessageType<ElementalShaman> {
    constructor();
    create(value?: PartialMessage<ElementalShaman>): ElementalShaman;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ElementalShaman): ElementalShaman;
    internalBinaryWrite(message: ElementalShaman, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ElementalShaman
 */
export declare const ElementalShaman: ElementalShaman$Type;
declare class ElementalShaman_Rotation$Type extends MessageType<ElementalShaman_Rotation> {
    constructor();
    create(value?: PartialMessage<ElementalShaman_Rotation>): ElementalShaman_Rotation;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ElementalShaman_Rotation): ElementalShaman_Rotation;
    internalBinaryWrite(message: ElementalShaman_Rotation, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ElementalShaman.Rotation
 */
export declare const ElementalShaman_Rotation: ElementalShaman_Rotation$Type;
declare class ElementalShaman_Options$Type extends MessageType<ElementalShaman_Options> {
    constructor();
    create(value?: PartialMessage<ElementalShaman_Options>): ElementalShaman_Options;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ElementalShaman_Options): ElementalShaman_Options;
    internalBinaryWrite(message: ElementalShaman_Options, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ElementalShaman.Options
 */
export declare const ElementalShaman_Options: ElementalShaman_Options$Type;
declare class EnhancementShaman$Type extends MessageType<EnhancementShaman> {
    constructor();
    create(value?: PartialMessage<EnhancementShaman>): EnhancementShaman;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: EnhancementShaman): EnhancementShaman;
    internalBinaryWrite(message: EnhancementShaman, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.EnhancementShaman
 */
export declare const EnhancementShaman: EnhancementShaman$Type;
declare class EnhancementShaman_Rotation$Type extends MessageType<EnhancementShaman_Rotation> {
    constructor();
    create(value?: PartialMessage<EnhancementShaman_Rotation>): EnhancementShaman_Rotation;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: EnhancementShaman_Rotation): EnhancementShaman_Rotation;
    internalBinaryWrite(message: EnhancementShaman_Rotation, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.EnhancementShaman.Rotation
 */
export declare const EnhancementShaman_Rotation: EnhancementShaman_Rotation$Type;
declare class EnhancementShaman_Options$Type extends MessageType<EnhancementShaman_Options> {
    constructor();
    create(value?: PartialMessage<EnhancementShaman_Options>): EnhancementShaman_Options;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: EnhancementShaman_Options): EnhancementShaman_Options;
    internalBinaryWrite(message: EnhancementShaman_Options, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.EnhancementShaman.Options
 */
export declare const EnhancementShaman_Options: EnhancementShaman_Options$Type;
export {};
