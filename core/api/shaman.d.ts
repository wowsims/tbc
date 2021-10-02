import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message api.ShamanTalents
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
     * Enhancement
     *
     * @generated from protobuf field: int32 ancestral_knowledge = 15;
     */
    ancestralKnowledge: number;
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
     * @generated from protobuf field: int32 flurry = 19;
     */
    flurry: number;
    /**
     * @generated from protobuf field: int32 improved_weapon_totems = 20;
     */
    improvedWeaponTotems: number;
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
 * @generated from protobuf message api.ElementalShaman
 */
export interface ElementalShaman {
    /**
     * @generated from protobuf field: api.ElementalShaman.Agent agent = 1;
     */
    agent?: ElementalShaman_Agent;
    /**
     * @generated from protobuf field: api.ShamanTalents talents = 2;
     */
    talents?: ShamanTalents;
    /**
     * @generated from protobuf field: api.ElementalShaman.Options options = 3;
     */
    options?: ElementalShaman_Options;
}
/**
 * @generated from protobuf message api.ElementalShaman.Agent
 */
export interface ElementalShaman_Agent {
    /**
     * @generated from protobuf field: api.ElementalShaman.Agent.AgentType type = 1;
     */
    type: ElementalShaman_Agent_AgentType;
}
/**
 * @generated from protobuf enum api.ElementalShaman.Agent.AgentType
 */
export declare enum ElementalShaman_Agent_AgentType {
    /**
     * @generated from protobuf enum value: Unknown = 0;
     */
    Unknown = 0,
    /**
     * @generated from protobuf enum value: FixedLBCL = 1;
     */
    FixedLBCL = 1,
    /**
     * @generated from protobuf enum value: CLOnClearcast = 2;
     */
    CLOnClearcast = 2,
    /**
     * @generated from protobuf enum value: Adaptive = 3;
     */
    Adaptive = 3,
    /**
     * @generated from protobuf enum value: CLOnCD = 4;
     */
    CLOnCD = 4
}
/**
 * @generated from protobuf message api.ElementalShaman.Options
 */
export interface ElementalShaman_Options {
    /**
     * @generated from protobuf field: bool water_shield = 1;
     */
    waterShield: boolean;
}
declare class ShamanTalents$Type extends MessageType<ShamanTalents> {
    constructor();
    create(value?: PartialMessage<ShamanTalents>): ShamanTalents;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ShamanTalents): ShamanTalents;
    internalBinaryWrite(message: ShamanTalents, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.ShamanTalents
 */
export declare const ShamanTalents: ShamanTalents$Type;
declare class ElementalShaman$Type extends MessageType<ElementalShaman> {
    constructor();
    create(value?: PartialMessage<ElementalShaman>): ElementalShaman;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ElementalShaman): ElementalShaman;
    internalBinaryWrite(message: ElementalShaman, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.ElementalShaman
 */
export declare const ElementalShaman: ElementalShaman$Type;
declare class ElementalShaman_Agent$Type extends MessageType<ElementalShaman_Agent> {
    constructor();
    create(value?: PartialMessage<ElementalShaman_Agent>): ElementalShaman_Agent;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ElementalShaman_Agent): ElementalShaman_Agent;
    internalBinaryWrite(message: ElementalShaman_Agent, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.ElementalShaman.Agent
 */
export declare const ElementalShaman_Agent: ElementalShaman_Agent$Type;
declare class ElementalShaman_Options$Type extends MessageType<ElementalShaman_Options> {
    constructor();
    create(value?: PartialMessage<ElementalShaman_Options>): ElementalShaman_Options;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ElementalShaman_Options): ElementalShaman_Options;
    internalBinaryWrite(message: ElementalShaman_Options, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message api.ElementalShaman.Options
 */
export declare const ElementalShaman_Options: ElementalShaman_Options$Type;
export {};
