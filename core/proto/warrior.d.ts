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
     * @generated from protobuf field: int32 deflection = 45;
     */
    deflection: number;
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
     * @generated from protobuf field: int32 improved_demoralizing_shout = 46;
     */
    improvedDemoralizingShout: number;
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
     * @generated from protobuf field: int32 anticipation = 47;
     */
    anticipation: number;
    /**
     * @generated from protobuf field: int32 shield_specialization = 48;
     */
    shieldSpecialization: number;
    /**
     * @generated from protobuf field: int32 toughness = 49;
     */
    toughness: number;
    /**
     * @generated from protobuf field: bool improved_shield_block = 50;
     */
    improvedShieldBlock: boolean;
    /**
     * @generated from protobuf field: int32 defiance = 38;
     */
    defiance: number;
    /**
     * @generated from protobuf field: int32 improved_sunder_armor = 39;
     */
    improvedSunderArmor: number;
    /**
     * @generated from protobuf field: int32 shield_mastery = 51;
     */
    shieldMastery: number;
    /**
     * @generated from protobuf field: int32 one_handed_weapon_specialization = 40;
     */
    oneHandedWeaponSpecialization: number;
    /**
     * @generated from protobuf field: int32 improved_defensive_stance = 52;
     */
    improvedDefensiveStance: number;
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
    /**
     * @generated from protobuf field: bool use_cleave = 14;
     */
    useCleave: boolean;
    /**
     * @generated from protobuf field: bool use_overpower = 1;
     */
    useOverpower: boolean;
    /**
     * @generated from protobuf field: bool use_hamstring = 2;
     */
    useHamstring: boolean;
    /**
     * @generated from protobuf field: bool use_slam = 3;
     */
    useSlam: boolean;
    /**
     * @generated from protobuf field: bool prioritize_ww = 4;
     */
    prioritizeWw: boolean;
    /**
     * @generated from protobuf field: proto.Warrior.Rotation.SunderArmor sunderArmor = 15;
     */
    sunderArmor: Warrior_Rotation_SunderArmor;
    /**
     * @generated from protobuf field: bool maintain_demo_shout = 16;
     */
    maintainDemoShout: boolean;
    /**
     * @generated from protobuf field: bool maintain_thunder_clap = 17;
     */
    maintainThunderClap: boolean;
    /**
     * Queue HS or Cleave when over this threshold.
     *
     * @generated from protobuf field: double hs_rage_threshold = 5;
     */
    hsRageThreshold: number;
    /**
     * Swap for overpower after reducing rage below this threshold.
     *
     * @generated from protobuf field: double overpower_rage_threshold = 6;
     */
    overpowerRageThreshold: number;
    /**
     * Use Hamstring in unused GCDs when over this threshold.
     *
     * @generated from protobuf field: double hamstring_rage_threshold = 7;
     */
    hamstringRageThreshold: number;
    /**
     * Refresh Rampage when remaining duration is less than this threshold.
     *
     * @generated from protobuf field: double rampage_cd_threshold = 8;
     */
    rampageCdThreshold: number;
    /**
     * Time between MH swing and start of Slam cast.
     *
     * @generated from protobuf field: double slam_latency = 9;
     */
    slamLatency: number;
    /**
     * Amount of time Slam is allowed to  delay the GCD, and MS+WW, by.
     *
     * @generated from protobuf field: double slam_gcd_delay = 19;
     */
    slamGcdDelay: number;
    /**
     * @generated from protobuf field: double slam_ms_ww_delay = 20;
     */
    slamMsWwDelay: number;
    /**
     * @generated from protobuf field: bool use_hs_during_execute = 10;
     */
    useHsDuringExecute: boolean;
    /**
     * @generated from protobuf field: bool use_bt_during_execute = 13;
     */
    useBtDuringExecute: boolean;
    /**
     * @generated from protobuf field: bool use_ms_during_execute = 12;
     */
    useMsDuringExecute: boolean;
    /**
     * @generated from protobuf field: bool use_ww_during_execute = 11;
     */
    useWwDuringExecute: boolean;
    /**
     * @generated from protobuf field: bool use_slam_during_execute = 18;
     */
    useSlamDuringExecute: boolean;
}
/**
 * @generated from protobuf enum proto.Warrior.Rotation.SunderArmor
 */
export declare enum Warrior_Rotation_SunderArmor {
    /**
     * @generated from protobuf enum value: SunderArmorNone = 0;
     */
    SunderArmorNone = 0,
    /**
     * @generated from protobuf enum value: SunderArmorHelpStack = 1;
     */
    SunderArmorHelpStack = 1,
    /**
     * @generated from protobuf enum value: SunderArmorMaintain = 2;
     */
    SunderArmorMaintain = 2
}
/**
 * @generated from protobuf message proto.Warrior.Options
 */
export interface Warrior_Options {
    /**
     * @generated from protobuf field: double starting_rage = 1;
     */
    startingRage: number;
    /**
     * @generated from protobuf field: bool use_recklessness = 2;
     */
    useRecklessness: boolean;
    /**
     * @generated from protobuf field: proto.WarriorShout shout = 3;
     */
    shout: WarriorShout;
    /**
     * @generated from protobuf field: bool precast_shout = 4;
     */
    precastShout: boolean;
    /**
     * @generated from protobuf field: bool precast_shout_t2 = 5;
     */
    precastShoutT2: boolean;
    /**
     * @generated from protobuf field: bool precast_shout_sapphire = 6;
     */
    precastShoutSapphire: boolean;
}
/**
 * @generated from protobuf message proto.ProtectionWarrior
 */
export interface ProtectionWarrior {
    /**
     * @generated from protobuf field: proto.ProtectionWarrior.Rotation rotation = 1;
     */
    rotation?: ProtectionWarrior_Rotation;
    /**
     * @generated from protobuf field: proto.WarriorTalents talents = 2;
     */
    talents?: WarriorTalents;
    /**
     * @generated from protobuf field: proto.ProtectionWarrior.Options options = 3;
     */
    options?: ProtectionWarrior_Options;
}
/**
 * @generated from protobuf message proto.ProtectionWarrior.Rotation
 */
export interface ProtectionWarrior_Rotation {
    /**
     * @generated from protobuf field: proto.ProtectionWarrior.Rotation.DemoShout demo_shout = 1;
     */
    demoShout: ProtectionWarrior_Rotation_DemoShout;
    /**
     * @generated from protobuf field: proto.ProtectionWarrior.Rotation.ThunderClap thunder_clap = 2;
     */
    thunderClap: ProtectionWarrior_Rotation_ThunderClap;
    /**
     * @generated from protobuf field: proto.ProtectionWarrior.Rotation.ShieldBlock shield_block = 5;
     */
    shieldBlock: ProtectionWarrior_Rotation_ShieldBlock;
    /**
     * @generated from protobuf field: bool use_cleave = 4;
     */
    useCleave: boolean;
    /**
     * Minimum rage to queue HS or Cleave.
     *
     * @generated from protobuf field: int32 hs_rage_threshold = 3;
     */
    hsRageThreshold: number;
}
/**
 * @generated from protobuf enum proto.ProtectionWarrior.Rotation.DemoShout
 */
export declare enum ProtectionWarrior_Rotation_DemoShout {
    /**
     * @generated from protobuf enum value: DemoShoutNone = 0;
     */
    DemoShoutNone = 0,
    /**
     * @generated from protobuf enum value: DemoShoutMaintain = 1;
     */
    DemoShoutMaintain = 1,
    /**
     * @generated from protobuf enum value: DemoShoutFiller = 2;
     */
    DemoShoutFiller = 2
}
/**
 * @generated from protobuf enum proto.ProtectionWarrior.Rotation.ThunderClap
 */
export declare enum ProtectionWarrior_Rotation_ThunderClap {
    /**
     * @generated from protobuf enum value: ThunderClapNone = 0;
     */
    ThunderClapNone = 0,
    /**
     * @generated from protobuf enum value: ThunderClapMaintain = 1;
     */
    ThunderClapMaintain = 1,
    /**
     * @generated from protobuf enum value: ThunderClapOnCD = 2;
     */
    ThunderClapOnCD = 2
}
/**
 * @generated from protobuf enum proto.ProtectionWarrior.Rotation.ShieldBlock
 */
export declare enum ProtectionWarrior_Rotation_ShieldBlock {
    /**
     * @generated from protobuf enum value: ShieldBlockNone = 0;
     */
    ShieldBlockNone = 0,
    /**
     * @generated from protobuf enum value: ShieldBlockToProcRevenge = 1;
     */
    ShieldBlockToProcRevenge = 1,
    /**
     * @generated from protobuf enum value: ShieldBlockOnCD = 2;
     */
    ShieldBlockOnCD = 2
}
/**
 * @generated from protobuf message proto.ProtectionWarrior.Options
 */
export interface ProtectionWarrior_Options {
    /**
     * @generated from protobuf field: double starting_rage = 1;
     */
    startingRage: number;
    /**
     * @generated from protobuf field: proto.WarriorShout shout = 4;
     */
    shout: WarriorShout;
    /**
     * @generated from protobuf field: bool precast_shout = 5;
     */
    precastShout: boolean;
    /**
     * @generated from protobuf field: bool precast_shout_t2 = 2;
     */
    precastShoutT2: boolean;
    /**
     * @generated from protobuf field: bool precast_shout_sapphire = 3;
     */
    precastShoutSapphire: boolean;
}
/**
 * @generated from protobuf enum proto.WarriorShout
 */
export declare enum WarriorShout {
    /**
     * @generated from protobuf enum value: WarriorShoutNone = 0;
     */
    WarriorShoutNone = 0,
    /**
     * @generated from protobuf enum value: WarriorShoutBattle = 1;
     */
    WarriorShoutBattle = 1,
    /**
     * @generated from protobuf enum value: WarriorShoutCommanding = 2;
     */
    WarriorShoutCommanding = 2
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
declare class ProtectionWarrior$Type extends MessageType<ProtectionWarrior> {
    constructor();
    create(value?: PartialMessage<ProtectionWarrior>): ProtectionWarrior;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ProtectionWarrior): ProtectionWarrior;
    internalBinaryWrite(message: ProtectionWarrior, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ProtectionWarrior
 */
export declare const ProtectionWarrior: ProtectionWarrior$Type;
declare class ProtectionWarrior_Rotation$Type extends MessageType<ProtectionWarrior_Rotation> {
    constructor();
    create(value?: PartialMessage<ProtectionWarrior_Rotation>): ProtectionWarrior_Rotation;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ProtectionWarrior_Rotation): ProtectionWarrior_Rotation;
    internalBinaryWrite(message: ProtectionWarrior_Rotation, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ProtectionWarrior.Rotation
 */
export declare const ProtectionWarrior_Rotation: ProtectionWarrior_Rotation$Type;
declare class ProtectionWarrior_Options$Type extends MessageType<ProtectionWarrior_Options> {
    constructor();
    create(value?: PartialMessage<ProtectionWarrior_Options>): ProtectionWarrior_Options;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ProtectionWarrior_Options): ProtectionWarrior_Options;
    internalBinaryWrite(message: ProtectionWarrior_Options, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ProtectionWarrior.Options
 */
export declare const ProtectionWarrior_Options: ProtectionWarrior_Options$Type;
export {};
