import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
/**
 * @generated from protobuf enum proto.Warlock.Rotation.PrimarySpell
 */
export var Warlock_Rotation_PrimarySpell;
(function (Warlock_Rotation_PrimarySpell) {
    /**
     * @generated from protobuf enum value: UnknownSpell = 0;
     */
    Warlock_Rotation_PrimarySpell[Warlock_Rotation_PrimarySpell["UnknownSpell"] = 0] = "UnknownSpell";
    /**
     * @generated from protobuf enum value: Shadowbolt = 1;
     */
    Warlock_Rotation_PrimarySpell[Warlock_Rotation_PrimarySpell["Shadowbolt"] = 1] = "Shadowbolt";
    /**
     * @generated from protobuf enum value: Incinerate = 2;
     */
    Warlock_Rotation_PrimarySpell[Warlock_Rotation_PrimarySpell["Incinerate"] = 2] = "Incinerate";
    /**
     * @generated from protobuf enum value: Seed = 3;
     */
    Warlock_Rotation_PrimarySpell[Warlock_Rotation_PrimarySpell["Seed"] = 3] = "Seed";
})(Warlock_Rotation_PrimarySpell || (Warlock_Rotation_PrimarySpell = {}));
/**
 * @generated from protobuf enum proto.Warlock.Rotation.Curse
 */
export var Warlock_Rotation_Curse;
(function (Warlock_Rotation_Curse) {
    /**
     * @generated from protobuf enum value: NoCurse = 0;
     */
    Warlock_Rotation_Curse[Warlock_Rotation_Curse["NoCurse"] = 0] = "NoCurse";
    /**
     * @generated from protobuf enum value: Elements = 1;
     */
    Warlock_Rotation_Curse[Warlock_Rotation_Curse["Elements"] = 1] = "Elements";
    /**
     * @generated from protobuf enum value: Recklessness = 2;
     */
    Warlock_Rotation_Curse[Warlock_Rotation_Curse["Recklessness"] = 2] = "Recklessness";
    /**
     * @generated from protobuf enum value: Doom = 3;
     */
    Warlock_Rotation_Curse[Warlock_Rotation_Curse["Doom"] = 3] = "Doom";
    /**
     * @generated from protobuf enum value: Agony = 4;
     */
    Warlock_Rotation_Curse[Warlock_Rotation_Curse["Agony"] = 4] = "Agony";
    /**
     * @generated from protobuf enum value: Tongues = 5;
     */
    Warlock_Rotation_Curse[Warlock_Rotation_Curse["Tongues"] = 5] = "Tongues";
})(Warlock_Rotation_Curse || (Warlock_Rotation_Curse = {}));
/**
 * @generated from protobuf enum proto.Warlock.Options.Summon
 */
export var Warlock_Options_Summon;
(function (Warlock_Options_Summon) {
    /**
     * @generated from protobuf enum value: NoSummon = 0;
     */
    Warlock_Options_Summon[Warlock_Options_Summon["NoSummon"] = 0] = "NoSummon";
    /**
     * @generated from protobuf enum value: Imp = 1;
     */
    Warlock_Options_Summon[Warlock_Options_Summon["Imp"] = 1] = "Imp";
    /**
     * @generated from protobuf enum value: Voidwalker = 2;
     */
    Warlock_Options_Summon[Warlock_Options_Summon["Voidwalker"] = 2] = "Voidwalker";
    /**
     * @generated from protobuf enum value: Succubus = 3;
     */
    Warlock_Options_Summon[Warlock_Options_Summon["Succubus"] = 3] = "Succubus";
    /**
     * @generated from protobuf enum value: Felhound = 4;
     */
    Warlock_Options_Summon[Warlock_Options_Summon["Felhound"] = 4] = "Felhound";
    /**
     * @generated from protobuf enum value: Felgaurd = 5;
     */
    Warlock_Options_Summon[Warlock_Options_Summon["Felgaurd"] = 5] = "Felgaurd";
})(Warlock_Options_Summon || (Warlock_Options_Summon = {}));
/**
 * @generated from protobuf enum proto.Warlock.Options.Armor
 */
export var Warlock_Options_Armor;
(function (Warlock_Options_Armor) {
    /**
     * @generated from protobuf enum value: NoArmor = 0;
     */
    Warlock_Options_Armor[Warlock_Options_Armor["NoArmor"] = 0] = "NoArmor";
    /**
     * @generated from protobuf enum value: FelArmor = 1;
     */
    Warlock_Options_Armor[Warlock_Options_Armor["FelArmor"] = 1] = "FelArmor";
    /**
     * @generated from protobuf enum value: DemonArmor = 2;
     */
    Warlock_Options_Armor[Warlock_Options_Armor["DemonArmor"] = 2] = "DemonArmor";
})(Warlock_Options_Armor || (Warlock_Options_Armor = {}));
// @generated message type with reflection information, may provide speed optimized methods
class WarlockTalents$Type extends MessageType {
    constructor() {
        super("proto.WarlockTalents", [
            { no: 1, name: "suppression", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "improved_corruption", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 49, name: "improved_drain_soul", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "improved_life_tap", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "soul_siphon", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "improved_curse_of_agony", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "amplify_curse", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 7, name: "nightfall", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 8, name: "empowered_corruption", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 9, name: "siphon_life", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 10, name: "shadow_mastery", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 11, name: "contagion", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 12, name: "dark_pact", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 13, name: "malediction", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 14, name: "unstable_affliction", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 15, name: "improved_imp", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 16, name: "demonic_embrace", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 17, name: "improved_voidwalker", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 18, name: "fel_intellect", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 19, name: "improved_sayaad", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 20, name: "fel_stamina", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 21, name: "demonic_aegis", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 22, name: "unholy_power", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 23, name: "improved_enslave_demon", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 24, name: "demonic_sacrifice", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 25, name: "master_conjuror", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 26, name: "mana_feed", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 27, name: "master_demonologist", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 28, name: "soul_link", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 29, name: "demonic_knowledge", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 30, name: "demonic_tactics", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 31, name: "summon_felguard", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 32, name: "improved_shadow_bolt", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 33, name: "cataclysm", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 34, name: "bane", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 35, name: "improved_firebolt", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 36, name: "improved_lash_of_pain", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 48, name: "destructive_reach", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 37, name: "devastation", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 38, name: "shadowburn", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 39, name: "improved_searing_pain", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 40, name: "improved_immolate", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 41, name: "ruin", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 42, name: "emberstorm", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 43, name: "backlash", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 44, name: "conflagrate", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 45, name: "soul_leech", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 46, name: "shadow_and_flame", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 47, name: "shadowfury", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { suppression: 0, improvedCorruption: 0, improvedDrainSoul: 0, improvedLifeTap: 0, soulSiphon: 0, improvedCurseOfAgony: 0, amplifyCurse: false, nightfall: 0, empoweredCorruption: 0, siphonLife: false, shadowMastery: 0, contagion: 0, darkPact: false, malediction: 0, unstableAffliction: false, improvedImp: 0, demonicEmbrace: 0, improvedVoidwalker: 0, felIntellect: 0, improvedSayaad: 0, felStamina: 0, demonicAegis: 0, unholyPower: 0, improvedEnslaveDemon: 0, demonicSacrifice: false, masterConjuror: 0, manaFeed: 0, masterDemonologist: 0, soulLink: false, demonicKnowledge: 0, demonicTactics: 0, summonFelguard: false, improvedShadowBolt: 0, cataclysm: 0, bane: 0, improvedFirebolt: 0, improvedLashOfPain: 0, destructiveReach: 0, devastation: 0, shadowburn: false, improvedSearingPain: 0, improvedImmolate: 0, ruin: false, emberstorm: 0, backlash: 0, conflagrate: false, soulLeech: 0, shadowAndFlame: 0, shadowfury: false };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 suppression */ 1:
                    message.suppression = reader.int32();
                    break;
                case /* int32 improved_corruption */ 2:
                    message.improvedCorruption = reader.int32();
                    break;
                case /* int32 improved_drain_soul */ 49:
                    message.improvedDrainSoul = reader.int32();
                    break;
                case /* int32 improved_life_tap */ 3:
                    message.improvedLifeTap = reader.int32();
                    break;
                case /* int32 soul_siphon */ 4:
                    message.soulSiphon = reader.int32();
                    break;
                case /* int32 improved_curse_of_agony */ 5:
                    message.improvedCurseOfAgony = reader.int32();
                    break;
                case /* bool amplify_curse */ 6:
                    message.amplifyCurse = reader.bool();
                    break;
                case /* int32 nightfall */ 7:
                    message.nightfall = reader.int32();
                    break;
                case /* int32 empowered_corruption */ 8:
                    message.empoweredCorruption = reader.int32();
                    break;
                case /* bool siphon_life */ 9:
                    message.siphonLife = reader.bool();
                    break;
                case /* int32 shadow_mastery */ 10:
                    message.shadowMastery = reader.int32();
                    break;
                case /* int32 contagion */ 11:
                    message.contagion = reader.int32();
                    break;
                case /* bool dark_pact */ 12:
                    message.darkPact = reader.bool();
                    break;
                case /* int32 malediction */ 13:
                    message.malediction = reader.int32();
                    break;
                case /* bool unstable_affliction */ 14:
                    message.unstableAffliction = reader.bool();
                    break;
                case /* int32 improved_imp */ 15:
                    message.improvedImp = reader.int32();
                    break;
                case /* int32 demonic_embrace */ 16:
                    message.demonicEmbrace = reader.int32();
                    break;
                case /* int32 improved_voidwalker */ 17:
                    message.improvedVoidwalker = reader.int32();
                    break;
                case /* int32 fel_intellect */ 18:
                    message.felIntellect = reader.int32();
                    break;
                case /* int32 improved_sayaad */ 19:
                    message.improvedSayaad = reader.int32();
                    break;
                case /* int32 fel_stamina */ 20:
                    message.felStamina = reader.int32();
                    break;
                case /* int32 demonic_aegis */ 21:
                    message.demonicAegis = reader.int32();
                    break;
                case /* int32 unholy_power */ 22:
                    message.unholyPower = reader.int32();
                    break;
                case /* int32 improved_enslave_demon */ 23:
                    message.improvedEnslaveDemon = reader.int32();
                    break;
                case /* bool demonic_sacrifice */ 24:
                    message.demonicSacrifice = reader.bool();
                    break;
                case /* int32 master_conjuror */ 25:
                    message.masterConjuror = reader.int32();
                    break;
                case /* int32 mana_feed */ 26:
                    message.manaFeed = reader.int32();
                    break;
                case /* int32 master_demonologist */ 27:
                    message.masterDemonologist = reader.int32();
                    break;
                case /* bool soul_link */ 28:
                    message.soulLink = reader.bool();
                    break;
                case /* int32 demonic_knowledge */ 29:
                    message.demonicKnowledge = reader.int32();
                    break;
                case /* int32 demonic_tactics */ 30:
                    message.demonicTactics = reader.int32();
                    break;
                case /* bool summon_felguard */ 31:
                    message.summonFelguard = reader.bool();
                    break;
                case /* int32 improved_shadow_bolt */ 32:
                    message.improvedShadowBolt = reader.int32();
                    break;
                case /* int32 cataclysm */ 33:
                    message.cataclysm = reader.int32();
                    break;
                case /* int32 bane */ 34:
                    message.bane = reader.int32();
                    break;
                case /* int32 improved_firebolt */ 35:
                    message.improvedFirebolt = reader.int32();
                    break;
                case /* int32 improved_lash_of_pain */ 36:
                    message.improvedLashOfPain = reader.int32();
                    break;
                case /* int32 destructive_reach */ 48:
                    message.destructiveReach = reader.int32();
                    break;
                case /* int32 devastation */ 37:
                    message.devastation = reader.int32();
                    break;
                case /* bool shadowburn */ 38:
                    message.shadowburn = reader.bool();
                    break;
                case /* int32 improved_searing_pain */ 39:
                    message.improvedSearingPain = reader.int32();
                    break;
                case /* int32 improved_immolate */ 40:
                    message.improvedImmolate = reader.int32();
                    break;
                case /* bool ruin */ 41:
                    message.ruin = reader.bool();
                    break;
                case /* int32 emberstorm */ 42:
                    message.emberstorm = reader.int32();
                    break;
                case /* int32 backlash */ 43:
                    message.backlash = reader.int32();
                    break;
                case /* bool conflagrate */ 44:
                    message.conflagrate = reader.bool();
                    break;
                case /* int32 soul_leech */ 45:
                    message.soulLeech = reader.int32();
                    break;
                case /* int32 shadow_and_flame */ 46:
                    message.shadowAndFlame = reader.int32();
                    break;
                case /* bool shadowfury */ 47:
                    message.shadowfury = reader.bool();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* int32 suppression = 1; */
        if (message.suppression !== 0)
            writer.tag(1, WireType.Varint).int32(message.suppression);
        /* int32 improved_corruption = 2; */
        if (message.improvedCorruption !== 0)
            writer.tag(2, WireType.Varint).int32(message.improvedCorruption);
        /* int32 improved_drain_soul = 49; */
        if (message.improvedDrainSoul !== 0)
            writer.tag(49, WireType.Varint).int32(message.improvedDrainSoul);
        /* int32 improved_life_tap = 3; */
        if (message.improvedLifeTap !== 0)
            writer.tag(3, WireType.Varint).int32(message.improvedLifeTap);
        /* int32 soul_siphon = 4; */
        if (message.soulSiphon !== 0)
            writer.tag(4, WireType.Varint).int32(message.soulSiphon);
        /* int32 improved_curse_of_agony = 5; */
        if (message.improvedCurseOfAgony !== 0)
            writer.tag(5, WireType.Varint).int32(message.improvedCurseOfAgony);
        /* bool amplify_curse = 6; */
        if (message.amplifyCurse !== false)
            writer.tag(6, WireType.Varint).bool(message.amplifyCurse);
        /* int32 nightfall = 7; */
        if (message.nightfall !== 0)
            writer.tag(7, WireType.Varint).int32(message.nightfall);
        /* int32 empowered_corruption = 8; */
        if (message.empoweredCorruption !== 0)
            writer.tag(8, WireType.Varint).int32(message.empoweredCorruption);
        /* bool siphon_life = 9; */
        if (message.siphonLife !== false)
            writer.tag(9, WireType.Varint).bool(message.siphonLife);
        /* int32 shadow_mastery = 10; */
        if (message.shadowMastery !== 0)
            writer.tag(10, WireType.Varint).int32(message.shadowMastery);
        /* int32 contagion = 11; */
        if (message.contagion !== 0)
            writer.tag(11, WireType.Varint).int32(message.contagion);
        /* bool dark_pact = 12; */
        if (message.darkPact !== false)
            writer.tag(12, WireType.Varint).bool(message.darkPact);
        /* int32 malediction = 13; */
        if (message.malediction !== 0)
            writer.tag(13, WireType.Varint).int32(message.malediction);
        /* bool unstable_affliction = 14; */
        if (message.unstableAffliction !== false)
            writer.tag(14, WireType.Varint).bool(message.unstableAffliction);
        /* int32 improved_imp = 15; */
        if (message.improvedImp !== 0)
            writer.tag(15, WireType.Varint).int32(message.improvedImp);
        /* int32 demonic_embrace = 16; */
        if (message.demonicEmbrace !== 0)
            writer.tag(16, WireType.Varint).int32(message.demonicEmbrace);
        /* int32 improved_voidwalker = 17; */
        if (message.improvedVoidwalker !== 0)
            writer.tag(17, WireType.Varint).int32(message.improvedVoidwalker);
        /* int32 fel_intellect = 18; */
        if (message.felIntellect !== 0)
            writer.tag(18, WireType.Varint).int32(message.felIntellect);
        /* int32 improved_sayaad = 19; */
        if (message.improvedSayaad !== 0)
            writer.tag(19, WireType.Varint).int32(message.improvedSayaad);
        /* int32 fel_stamina = 20; */
        if (message.felStamina !== 0)
            writer.tag(20, WireType.Varint).int32(message.felStamina);
        /* int32 demonic_aegis = 21; */
        if (message.demonicAegis !== 0)
            writer.tag(21, WireType.Varint).int32(message.demonicAegis);
        /* int32 unholy_power = 22; */
        if (message.unholyPower !== 0)
            writer.tag(22, WireType.Varint).int32(message.unholyPower);
        /* int32 improved_enslave_demon = 23; */
        if (message.improvedEnslaveDemon !== 0)
            writer.tag(23, WireType.Varint).int32(message.improvedEnslaveDemon);
        /* bool demonic_sacrifice = 24; */
        if (message.demonicSacrifice !== false)
            writer.tag(24, WireType.Varint).bool(message.demonicSacrifice);
        /* int32 master_conjuror = 25; */
        if (message.masterConjuror !== 0)
            writer.tag(25, WireType.Varint).int32(message.masterConjuror);
        /* int32 mana_feed = 26; */
        if (message.manaFeed !== 0)
            writer.tag(26, WireType.Varint).int32(message.manaFeed);
        /* int32 master_demonologist = 27; */
        if (message.masterDemonologist !== 0)
            writer.tag(27, WireType.Varint).int32(message.masterDemonologist);
        /* bool soul_link = 28; */
        if (message.soulLink !== false)
            writer.tag(28, WireType.Varint).bool(message.soulLink);
        /* int32 demonic_knowledge = 29; */
        if (message.demonicKnowledge !== 0)
            writer.tag(29, WireType.Varint).int32(message.demonicKnowledge);
        /* int32 demonic_tactics = 30; */
        if (message.demonicTactics !== 0)
            writer.tag(30, WireType.Varint).int32(message.demonicTactics);
        /* bool summon_felguard = 31; */
        if (message.summonFelguard !== false)
            writer.tag(31, WireType.Varint).bool(message.summonFelguard);
        /* int32 improved_shadow_bolt = 32; */
        if (message.improvedShadowBolt !== 0)
            writer.tag(32, WireType.Varint).int32(message.improvedShadowBolt);
        /* int32 cataclysm = 33; */
        if (message.cataclysm !== 0)
            writer.tag(33, WireType.Varint).int32(message.cataclysm);
        /* int32 bane = 34; */
        if (message.bane !== 0)
            writer.tag(34, WireType.Varint).int32(message.bane);
        /* int32 improved_firebolt = 35; */
        if (message.improvedFirebolt !== 0)
            writer.tag(35, WireType.Varint).int32(message.improvedFirebolt);
        /* int32 improved_lash_of_pain = 36; */
        if (message.improvedLashOfPain !== 0)
            writer.tag(36, WireType.Varint).int32(message.improvedLashOfPain);
        /* int32 destructive_reach = 48; */
        if (message.destructiveReach !== 0)
            writer.tag(48, WireType.Varint).int32(message.destructiveReach);
        /* int32 devastation = 37; */
        if (message.devastation !== 0)
            writer.tag(37, WireType.Varint).int32(message.devastation);
        /* bool shadowburn = 38; */
        if (message.shadowburn !== false)
            writer.tag(38, WireType.Varint).bool(message.shadowburn);
        /* int32 improved_searing_pain = 39; */
        if (message.improvedSearingPain !== 0)
            writer.tag(39, WireType.Varint).int32(message.improvedSearingPain);
        /* int32 improved_immolate = 40; */
        if (message.improvedImmolate !== 0)
            writer.tag(40, WireType.Varint).int32(message.improvedImmolate);
        /* bool ruin = 41; */
        if (message.ruin !== false)
            writer.tag(41, WireType.Varint).bool(message.ruin);
        /* int32 emberstorm = 42; */
        if (message.emberstorm !== 0)
            writer.tag(42, WireType.Varint).int32(message.emberstorm);
        /* int32 backlash = 43; */
        if (message.backlash !== 0)
            writer.tag(43, WireType.Varint).int32(message.backlash);
        /* bool conflagrate = 44; */
        if (message.conflagrate !== false)
            writer.tag(44, WireType.Varint).bool(message.conflagrate);
        /* int32 soul_leech = 45; */
        if (message.soulLeech !== 0)
            writer.tag(45, WireType.Varint).int32(message.soulLeech);
        /* int32 shadow_and_flame = 46; */
        if (message.shadowAndFlame !== 0)
            writer.tag(46, WireType.Varint).int32(message.shadowAndFlame);
        /* bool shadowfury = 47; */
        if (message.shadowfury !== false)
            writer.tag(47, WireType.Varint).bool(message.shadowfury);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.WarlockTalents
 */
export const WarlockTalents = new WarlockTalents$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Warlock$Type extends MessageType {
    constructor() {
        super("proto.Warlock", [
            { no: 1, name: "rotation", kind: "message", T: () => Warlock_Rotation },
            { no: 2, name: "talents", kind: "message", T: () => WarlockTalents },
            { no: 3, name: "options", kind: "message", T: () => Warlock_Options }
        ]);
    }
    create(value) {
        const message = {};
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* proto.Warlock.Rotation rotation */ 1:
                    message.rotation = Warlock_Rotation.internalBinaryRead(reader, reader.uint32(), options, message.rotation);
                    break;
                case /* proto.WarlockTalents talents */ 2:
                    message.talents = WarlockTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* proto.Warlock.Options options */ 3:
                    message.options = Warlock_Options.internalBinaryRead(reader, reader.uint32(), options, message.options);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* proto.Warlock.Rotation rotation = 1; */
        if (message.rotation)
            Warlock_Rotation.internalBinaryWrite(message.rotation, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.WarlockTalents talents = 2; */
        if (message.talents)
            WarlockTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.Warlock.Options options = 3; */
        if (message.options)
            Warlock_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Warlock
 */
export const Warlock = new Warlock$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Warlock_Rotation$Type extends MessageType {
    constructor() {
        super("proto.Warlock.Rotation", [
            { no: 1, name: "primary_spell", kind: "enum", T: () => ["proto.Warlock.Rotation.PrimarySpell", Warlock_Rotation_PrimarySpell] },
            { no: 2, name: "curse", kind: "enum", T: () => ["proto.Warlock.Rotation.Curse", Warlock_Rotation_Curse] },
            { no: 3, name: "immolate", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "corruption", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "detonate_seed", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { primarySpell: 0, curse: 0, immolate: false, corruption: false, detonateSeed: false };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* proto.Warlock.Rotation.PrimarySpell primary_spell */ 1:
                    message.primarySpell = reader.int32();
                    break;
                case /* proto.Warlock.Rotation.Curse curse */ 2:
                    message.curse = reader.int32();
                    break;
                case /* bool immolate */ 3:
                    message.immolate = reader.bool();
                    break;
                case /* bool corruption */ 4:
                    message.corruption = reader.bool();
                    break;
                case /* bool detonate_seed */ 5:
                    message.detonateSeed = reader.bool();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* proto.Warlock.Rotation.PrimarySpell primary_spell = 1; */
        if (message.primarySpell !== 0)
            writer.tag(1, WireType.Varint).int32(message.primarySpell);
        /* proto.Warlock.Rotation.Curse curse = 2; */
        if (message.curse !== 0)
            writer.tag(2, WireType.Varint).int32(message.curse);
        /* bool immolate = 3; */
        if (message.immolate !== false)
            writer.tag(3, WireType.Varint).bool(message.immolate);
        /* bool corruption = 4; */
        if (message.corruption !== false)
            writer.tag(4, WireType.Varint).bool(message.corruption);
        /* bool detonate_seed = 5; */
        if (message.detonateSeed !== false)
            writer.tag(5, WireType.Varint).bool(message.detonateSeed);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Warlock.Rotation
 */
export const Warlock_Rotation = new Warlock_Rotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Warlock_Options$Type extends MessageType {
    constructor() {
        super("proto.Warlock.Options", [
            { no: 1, name: "armor", kind: "enum", T: () => ["proto.Warlock.Options.Armor", Warlock_Options_Armor] },
            { no: 2, name: "summon", kind: "enum", T: () => ["proto.Warlock.Options.Summon", Warlock_Options_Summon] },
            { no: 3, name: "sacrifice_summon", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { armor: 0, summon: 0, sacrificeSummon: false };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* proto.Warlock.Options.Armor armor */ 1:
                    message.armor = reader.int32();
                    break;
                case /* proto.Warlock.Options.Summon summon */ 2:
                    message.summon = reader.int32();
                    break;
                case /* bool sacrifice_summon */ 3:
                    message.sacrificeSummon = reader.bool();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* proto.Warlock.Options.Armor armor = 1; */
        if (message.armor !== 0)
            writer.tag(1, WireType.Varint).int32(message.armor);
        /* proto.Warlock.Options.Summon summon = 2; */
        if (message.summon !== 0)
            writer.tag(2, WireType.Varint).int32(message.summon);
        /* bool sacrifice_summon = 3; */
        if (message.sacrificeSummon !== false)
            writer.tag(3, WireType.Varint).bool(message.sacrificeSummon);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Warlock.Options
 */
export const Warlock_Options = new Warlock_Options$Type();
