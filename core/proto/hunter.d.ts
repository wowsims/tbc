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
     * @generated from protobuf field: int32 animal_handler = 45;
     */
    animalHandler: number;
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
     * @generated from protobuf field: int32 deflection = 46;
     */
    deflection: number;
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
    /**
     * @generated from protobuf field: bool use_multi_shot = 1;
     */
    useMultiShot: boolean;
    /**
     * @generated from protobuf field: bool use_arcane_shot = 2;
     */
    useArcaneShot: boolean;
    /**
     * @generated from protobuf field: bool precast_aimed_shot = 3;
     */
    precastAimedShot: boolean;
    /**
     * @generated from protobuf field: bool lazy_rotation = 12;
     */
    lazyRotation: boolean;
    /**
     * @generated from protobuf field: proto.Hunter.Rotation.StingType sting = 5;
     */
    sting: Hunter_Rotation_StingType;
    /**
     * Switch to Aspect of the Viper when mana goes below this percent.
     *
     * @generated from protobuf field: double viper_start_mana_percent = 6;
     */
    viperStartManaPercent: number;
    /**
     * Switch back to Aspect of the Hawk when mana goes above this percent.
     *
     * @generated from protobuf field: double viper_stop_mana_percent = 7;
     */
    viperStopManaPercent: number;
    /**
     * @generated from protobuf field: proto.Hunter.Rotation.WeaveType weave = 11;
     */
    weave: Hunter_Rotation_WeaveType;
    /**
     * Time in milliseconds from when the hunter starts moving to when they reengage their ranged autos.
     *
     * @generated from protobuf field: int32 time_to_weave_ms = 9;
     */
    timeToWeaveMs: number;
    /**
     * Percent of fight (0-1) to use melee weaving.
     *
     * @generated from protobuf field: double percent_weaved = 10;
     */
    percentWeaved: number;
}
/**
 * @generated from protobuf enum proto.Hunter.Rotation.StingType
 */
export declare enum Hunter_Rotation_StingType {
    /**
     * @generated from protobuf enum value: NoSting = 0;
     */
    NoSting = 0,
    /**
     * @generated from protobuf enum value: ScorpidSting = 1;
     */
    ScorpidSting = 1,
    /**
     * @generated from protobuf enum value: SerpentSting = 2;
     */
    SerpentSting = 2
}
/**
 * @generated from protobuf enum proto.Hunter.Rotation.WeaveType
 */
export declare enum Hunter_Rotation_WeaveType {
    /**
     * @generated from protobuf enum value: WeaveNone = 0;
     */
    WeaveNone = 0,
    /**
     * @generated from protobuf enum value: WeaveAutosOnly = 1;
     */
    WeaveAutosOnly = 1,
    /**
     * @generated from protobuf enum value: WeaveRaptorOnly = 2;
     */
    WeaveRaptorOnly = 2,
    /**
     * @generated from protobuf enum value: WeaveFull = 3;
     */
    WeaveFull = 3
}
/**
 * @generated from protobuf message proto.Hunter.Options
 */
export interface Hunter_Options {
    /**
     * @generated from protobuf field: proto.Hunter.Options.QuiverBonus quiver_bonus = 1;
     */
    quiverBonus: Hunter_Options_QuiverBonus;
    /**
     * @generated from protobuf field: proto.Hunter.Options.Ammo ammo = 2;
     */
    ammo: Hunter_Options_Ammo;
    /**
     * @generated from protobuf field: proto.Hunter.Options.PetType pet_type = 3;
     */
    petType: Hunter_Options_PetType;
    /**
     * @generated from protobuf field: double pet_uptime = 4;
     */
    petUptime: number;
    /**
     * @generated from protobuf field: bool pet_single_ability = 6;
     */
    petSingleAbility: boolean;
    /**
     * @generated from protobuf field: int32 latency_ms = 5;
     */
    latencyMs: number;
    /**
     * For internal use only.
     * Used for hunter presims to avoid artifacts caused by randomness when
     * calculating average ability damage from a low sample size.
     *
     * @generated from protobuf field: bool remove_randomness = 7;
     */
    removeRandomness: boolean;
}
/**
 * @generated from protobuf enum proto.Hunter.Options.QuiverBonus
 */
export declare enum Hunter_Options_QuiverBonus {
    /**
     * @generated from protobuf enum value: QuiverNone = 0;
     */
    QuiverNone = 0,
    /**
     * @generated from protobuf enum value: Speed10 = 1;
     */
    Speed10 = 1,
    /**
     * @generated from protobuf enum value: Speed11 = 2;
     */
    Speed11 = 2,
    /**
     * @generated from protobuf enum value: Speed12 = 3;
     */
    Speed12 = 3,
    /**
     * @generated from protobuf enum value: Speed13 = 4;
     */
    Speed13 = 4,
    /**
     * @generated from protobuf enum value: Speed14 = 5;
     */
    Speed14 = 5,
    /**
     * @generated from protobuf enum value: Speed15 = 6;
     */
    Speed15 = 6
}
/**
 * @generated from protobuf enum proto.Hunter.Options.Ammo
 */
export declare enum Hunter_Options_Ammo {
    /**
     * @generated from protobuf enum value: AmmoNone = 0;
     */
    AmmoNone = 0,
    /**
     * @generated from protobuf enum value: TimelessArrow = 1;
     */
    TimelessArrow = 1,
    /**
     * @generated from protobuf enum value: MysteriousArrow = 2;
     */
    MysteriousArrow = 2,
    /**
     * @generated from protobuf enum value: AdamantiteStinger = 3;
     */
    AdamantiteStinger = 3,
    /**
     * @generated from protobuf enum value: WardensArrow = 4;
     */
    WardensArrow = 4,
    /**
     * @generated from protobuf enum value: HalaaniRazorshaft = 5;
     */
    HalaaniRazorshaft = 5,
    /**
     * @generated from protobuf enum value: BlackflightArrow = 6;
     */
    BlackflightArrow = 6
}
/**
 * @generated from protobuf enum proto.Hunter.Options.PetType
 */
export declare enum Hunter_Options_PetType {
    /**
     * @generated from protobuf enum value: PetNone = 0;
     */
    PetNone = 0,
    /**
     * @generated from protobuf enum value: Bat = 5;
     */
    Bat = 5,
    /**
     * @generated from protobuf enum value: Bear = 7;
     */
    Bear = 7,
    /**
     * @generated from protobuf enum value: Cat = 2;
     */
    Cat = 2,
    /**
     * @generated from protobuf enum value: Crab = 8;
     */
    Crab = 8,
    /**
     * @generated from protobuf enum value: Owl = 4;
     */
    Owl = 4,
    /**
     * @generated from protobuf enum value: Raptor = 3;
     */
    Raptor = 3,
    /**
     * @generated from protobuf enum value: Ravager = 1;
     */
    Ravager = 1,
    /**
     * @generated from protobuf enum value: WindSerpent = 6;
     */
    WindSerpent = 6
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
