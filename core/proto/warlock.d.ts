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
     * @generated from protobuf field: int32 improved_succubus = 19;
     */
    improvedSuccubus: number;
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
}
/**
 * @generated from protobuf message proto.Warlock.Options
 */
export interface Warlock_Options {
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
