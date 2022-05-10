import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message proto.WarlockTalents
 */
export interface WarlockTalents {
    /**
     * Affliction
     *
     * @generated from protobuf field: int32 suppression = 1;
     */
    suppression: number;
    /**
     * @generated from protobuf field: int32 improved_corruption = 2;
     */
    improvedCorruption: number;
    /**
     * @generated from protobuf field: int32 improved_drain_soul = 49;
     */
    improvedDrainSoul: number;
    /**
     * @generated from protobuf field: int32 improved_life_tap = 3;
     */
    improvedLifeTap: number;
    /**
     * @generated from protobuf field: int32 soul_siphon = 4;
     */
    soulSiphon: number;
    /**
     * @generated from protobuf field: int32 improved_curse_of_agony = 5;
     */
    improvedCurseOfAgony: number;
    /**
     * @generated from protobuf field: bool amplify_curse = 6;
     */
    amplifyCurse: boolean;
    /**
     * @generated from protobuf field: int32 nightfall = 7;
     */
    nightfall: number;
    /**
     * @generated from protobuf field: int32 empowered_corruption = 8;
     */
    empoweredCorruption: number;
    /**
     * @generated from protobuf field: int32 shadow_embrace = 50;
     */
    shadowEmbrace: number;
    /**
     * @generated from protobuf field: bool siphon_life = 9;
     */
    siphonLife: boolean;
    /**
     * @generated from protobuf field: int32 shadow_mastery = 10;
     */
    shadowMastery: number;
    /**
     * @generated from protobuf field: int32 contagion = 11;
     */
    contagion: number;
    /**
     * @generated from protobuf field: bool dark_pact = 12;
     */
    darkPact: boolean;
    /**
     * @generated from protobuf field: int32 malediction = 13;
     */
    malediction: number;
    /**
     * @generated from protobuf field: bool unstable_affliction = 14;
     */
    unstableAffliction: boolean;
    /**
     * Demonology
     *
     * @generated from protobuf field: int32 improved_imp = 15;
     */
    improvedImp: number;
    /**
     * @generated from protobuf field: int32 demonic_embrace = 16;
     */
    demonicEmbrace: number;
    /**
     * @generated from protobuf field: int32 improved_voidwalker = 17;
     */
    improvedVoidwalker: number;
    /**
     * @generated from protobuf field: int32 fel_intellect = 18;
     */
    felIntellect: number;
    /**
     * @generated from protobuf field: int32 improved_sayaad = 19;
     */
    improvedSayaad: number;
    /**
     * @generated from protobuf field: int32 fel_stamina = 20;
     */
    felStamina: number;
    /**
     * @generated from protobuf field: int32 demonic_aegis = 21;
     */
    demonicAegis: number;
    /**
     * @generated from protobuf field: int32 unholy_power = 22;
     */
    unholyPower: number;
    /**
     * @generated from protobuf field: int32 improved_enslave_demon = 23;
     */
    improvedEnslaveDemon: number;
    /**
     * @generated from protobuf field: bool demonic_sacrifice = 24;
     */
    demonicSacrifice: boolean;
    /**
     * @generated from protobuf field: int32 master_conjuror = 25;
     */
    masterConjuror: number;
    /**
     * @generated from protobuf field: int32 mana_feed = 26;
     */
    manaFeed: number;
    /**
     * @generated from protobuf field: int32 master_demonologist = 27;
     */
    masterDemonologist: number;
    /**
     * @generated from protobuf field: bool soul_link = 28;
     */
    soulLink: boolean;
    /**
     * @generated from protobuf field: int32 demonic_knowledge = 29;
     */
    demonicKnowledge: number;
    /**
     * @generated from protobuf field: int32 demonic_tactics = 30;
     */
    demonicTactics: number;
    /**
     * @generated from protobuf field: bool summon_felguard = 31;
     */
    summonFelguard: boolean;
    /**
     * Destruction
     *
     * @generated from protobuf field: int32 improved_shadow_bolt = 32;
     */
    improvedShadowBolt: number;
    /**
     * @generated from protobuf field: int32 cataclysm = 33;
     */
    cataclysm: number;
    /**
     * @generated from protobuf field: int32 bane = 34;
     */
    bane: number;
    /**
     * @generated from protobuf field: int32 improved_firebolt = 35;
     */
    improvedFirebolt: number;
    /**
     * @generated from protobuf field: int32 improved_lash_of_pain = 36;
     */
    improvedLashOfPain: number;
    /**
     * @generated from protobuf field: int32 destructive_reach = 48;
     */
    destructiveReach: number;
    /**
     * @generated from protobuf field: int32 devastation = 37;
     */
    devastation: number;
    /**
     * @generated from protobuf field: bool shadowburn = 38;
     */
    shadowburn: boolean;
    /**
     * @generated from protobuf field: int32 improved_searing_pain = 39;
     */
    improvedSearingPain: number;
    /**
     * @generated from protobuf field: int32 improved_immolate = 40;
     */
    improvedImmolate: number;
    /**
     * @generated from protobuf field: bool ruin = 41;
     */
    ruin: boolean;
    /**
     * @generated from protobuf field: int32 emberstorm = 42;
     */
    emberstorm: number;
    /**
     * @generated from protobuf field: int32 backlash = 43;
     */
    backlash: number;
    /**
     * @generated from protobuf field: bool conflagrate = 44;
     */
    conflagrate: boolean;
    /**
     * @generated from protobuf field: int32 soul_leech = 45;
     */
    soulLeech: number;
    /**
     * @generated from protobuf field: int32 shadow_and_flame = 46;
     */
    shadowAndFlame: number;
    /**
     * @generated from protobuf field: bool shadowfury = 47;
     */
    shadowfury: boolean;
}
/**
 * @generated from protobuf message proto.Warlock
 */
export interface Warlock {
    /**
     * @generated from protobuf field: proto.Warlock.Rotation rotation = 1;
     */
    rotation?: Warlock_Rotation;
    /**
     * @generated from protobuf field: proto.WarlockTalents talents = 2;
     */
    talents?: WarlockTalents;
    /**
     * @generated from protobuf field: proto.Warlock.Options options = 3;
     */
    options?: Warlock_Options;
}
/**
 * @generated from protobuf message proto.Warlock.Rotation
 */
export interface Warlock_Rotation {
    /**
     * @generated from protobuf field: proto.Warlock.Rotation.PrimarySpell primary_spell = 1;
     */
    primarySpell: Warlock_Rotation_PrimarySpell;
    /**
     * @generated from protobuf field: proto.Warlock.Rotation.Curse curse = 2;
     */
    curse: Warlock_Rotation_Curse;
    /**
     * @generated from protobuf field: bool immolate = 3;
     */
    immolate: boolean;
    /**
     * @generated from protobuf field: bool corruption = 4;
     */
    corruption: boolean;
    /**
     * @generated from protobuf field: bool detonate_seed = 5;
     */
    detonateSeed: boolean;
}
/**
 * @generated from protobuf enum proto.Warlock.Rotation.PrimarySpell
 */
export declare enum Warlock_Rotation_PrimarySpell {
    /**
     * @generated from protobuf enum value: UnknownSpell = 0;
     */
    UnknownSpell = 0,
    /**
     * @generated from protobuf enum value: Shadowbolt = 1;
     */
    Shadowbolt = 1,
    /**
     * @generated from protobuf enum value: Incinerate = 2;
     */
    Incinerate = 2,
    /**
     * @generated from protobuf enum value: Seed = 3;
     */
    Seed = 3
}
/**
 * @generated from protobuf enum proto.Warlock.Rotation.Curse
 */
export declare enum Warlock_Rotation_Curse {
    /**
     * @generated from protobuf enum value: NoCurse = 0;
     */
    NoCurse = 0,
    /**
     * @generated from protobuf enum value: Elements = 1;
     */
    Elements = 1,
    /**
     * @generated from protobuf enum value: Recklessness = 2;
     */
    Recklessness = 2,
    /**
     * @generated from protobuf enum value: Doom = 3;
     */
    Doom = 3,
    /**
     * @generated from protobuf enum value: Agony = 4;
     */
    Agony = 4,
    /**
     * @generated from protobuf enum value: Tongues = 5;
     */
    Tongues = 5
}
/**
 * @generated from protobuf message proto.Warlock.Options
 */
export interface Warlock_Options {
    /**
     * @generated from protobuf field: proto.Warlock.Options.Armor armor = 1;
     */
    armor: Warlock_Options_Armor;
    /**
     * @generated from protobuf field: proto.Warlock.Options.Summon summon = 2;
     */
    summon: Warlock_Options_Summon;
    /**
     * @generated from protobuf field: bool sacrifice_summon = 3;
     */
    sacrificeSummon: boolean;
}
/**
 * @generated from protobuf enum proto.Warlock.Options.Summon
 */
export declare enum Warlock_Options_Summon {
    /**
     * @generated from protobuf enum value: NoSummon = 0;
     */
    NoSummon = 0,
    /**
     * @generated from protobuf enum value: Imp = 1;
     */
    Imp = 1,
    /**
     * @generated from protobuf enum value: Voidwalker = 2;
     */
    Voidwalker = 2,
    /**
     * @generated from protobuf enum value: Succubus = 3;
     */
    Succubus = 3,
    /**
     * @generated from protobuf enum value: Felhound = 4;
     */
    Felhound = 4,
    /**
     * @generated from protobuf enum value: Felgaurd = 5;
     */
    Felgaurd = 5
}
/**
 * @generated from protobuf enum proto.Warlock.Options.Armor
 */
export declare enum Warlock_Options_Armor {
    /**
     * @generated from protobuf enum value: NoArmor = 0;
     */
    NoArmor = 0,
    /**
     * @generated from protobuf enum value: FelArmor = 1;
     */
    FelArmor = 1,
    /**
     * @generated from protobuf enum value: DemonArmor = 2;
     */
    DemonArmor = 2
}
declare class WarlockTalents$Type extends MessageType<WarlockTalents> {
    constructor();
    create(value?: PartialMessage<WarlockTalents>): WarlockTalents;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: WarlockTalents): WarlockTalents;
    internalBinaryWrite(message: WarlockTalents, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.WarlockTalents
 */
export declare const WarlockTalents: WarlockTalents$Type;
declare class Warlock$Type extends MessageType<Warlock> {
    constructor();
    create(value?: PartialMessage<Warlock>): Warlock;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Warlock): Warlock;
    internalBinaryWrite(message: Warlock, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Warlock
 */
export declare const Warlock: Warlock$Type;
declare class Warlock_Rotation$Type extends MessageType<Warlock_Rotation> {
    constructor();
    create(value?: PartialMessage<Warlock_Rotation>): Warlock_Rotation;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Warlock_Rotation): Warlock_Rotation;
    internalBinaryWrite(message: Warlock_Rotation, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Warlock.Rotation
 */
export declare const Warlock_Rotation: Warlock_Rotation$Type;
declare class Warlock_Options$Type extends MessageType<Warlock_Options> {
    constructor();
    create(value?: PartialMessage<Warlock_Options>): Warlock_Options;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Warlock_Options): Warlock_Options;
    internalBinaryWrite(message: Warlock_Options, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Warlock.Options
 */
export declare const Warlock_Options: Warlock_Options$Type;
export {};
