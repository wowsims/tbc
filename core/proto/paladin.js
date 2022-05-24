import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
/**
 * @generated from protobuf enum proto.RetributionPaladin.Rotation.ConsecrationRank
 */
export var RetributionPaladin_Rotation_ConsecrationRank;
(function (RetributionPaladin_Rotation_ConsecrationRank) {
    /**
     * @generated from protobuf enum value: None = 0;
     */
    RetributionPaladin_Rotation_ConsecrationRank[RetributionPaladin_Rotation_ConsecrationRank["None"] = 0] = "None";
    /**
     * @generated from protobuf enum value: Rank1 = 1;
     */
    RetributionPaladin_Rotation_ConsecrationRank[RetributionPaladin_Rotation_ConsecrationRank["Rank1"] = 1] = "Rank1";
    /**
     * @generated from protobuf enum value: Rank4 = 2;
     */
    RetributionPaladin_Rotation_ConsecrationRank[RetributionPaladin_Rotation_ConsecrationRank["Rank4"] = 2] = "Rank4";
    /**
     * @generated from protobuf enum value: Rank6 = 3;
     */
    RetributionPaladin_Rotation_ConsecrationRank[RetributionPaladin_Rotation_ConsecrationRank["Rank6"] = 3] = "Rank6";
})(RetributionPaladin_Rotation_ConsecrationRank || (RetributionPaladin_Rotation_ConsecrationRank = {}));
/**
 * @generated from protobuf enum proto.RetributionPaladin.Options.Judgement
 */
export var RetributionPaladin_Options_Judgement;
(function (RetributionPaladin_Options_Judgement) {
    /**
     * @generated from protobuf enum value: None = 0;
     */
    RetributionPaladin_Options_Judgement[RetributionPaladin_Options_Judgement["None"] = 0] = "None";
    /**
     * @generated from protobuf enum value: Wisdom = 1;
     */
    RetributionPaladin_Options_Judgement[RetributionPaladin_Options_Judgement["Wisdom"] = 1] = "Wisdom";
    /**
     * @generated from protobuf enum value: Crusader = 2;
     */
    RetributionPaladin_Options_Judgement[RetributionPaladin_Options_Judgement["Crusader"] = 2] = "Crusader";
})(RetributionPaladin_Options_Judgement || (RetributionPaladin_Options_Judgement = {}));
/**
 * @generated from protobuf enum proto.Blessings
 */
export var Blessings;
(function (Blessings) {
    /**
     * @generated from protobuf enum value: BlessingUnknown = 0;
     */
    Blessings[Blessings["BlessingUnknown"] = 0] = "BlessingUnknown";
    /**
     * @generated from protobuf enum value: BlessingOfKings = 1;
     */
    Blessings[Blessings["BlessingOfKings"] = 1] = "BlessingOfKings";
    /**
     * @generated from protobuf enum value: BlessingOfMight = 2;
     */
    Blessings[Blessings["BlessingOfMight"] = 2] = "BlessingOfMight";
    /**
     * @generated from protobuf enum value: BlessingOfSalvation = 3;
     */
    Blessings[Blessings["BlessingOfSalvation"] = 3] = "BlessingOfSalvation";
    /**
     * @generated from protobuf enum value: BlessingOfWisdom = 4;
     */
    Blessings[Blessings["BlessingOfWisdom"] = 4] = "BlessingOfWisdom";
    /**
     * @generated from protobuf enum value: BlessingOfSanctuary = 5;
     */
    Blessings[Blessings["BlessingOfSanctuary"] = 5] = "BlessingOfSanctuary";
    /**
     * @generated from protobuf enum value: BlessingOfLight = 6;
     */
    Blessings[Blessings["BlessingOfLight"] = 6] = "BlessingOfLight";
})(Blessings || (Blessings = {}));
/**
 * @generated from protobuf enum proto.PaladinAura
 */
export var PaladinAura;
(function (PaladinAura) {
    /**
     * @generated from protobuf enum value: NoPaladinAura = 0;
     */
    PaladinAura[PaladinAura["NoPaladinAura"] = 0] = "NoPaladinAura";
    /**
     * @generated from protobuf enum value: SanctityAura = 1;
     */
    PaladinAura[PaladinAura["SanctityAura"] = 1] = "SanctityAura";
    /**
     * @generated from protobuf enum value: DevotionAura = 2;
     */
    PaladinAura[PaladinAura["DevotionAura"] = 2] = "DevotionAura";
    /**
     * @generated from protobuf enum value: RetributionAura = 3;
     */
    PaladinAura[PaladinAura["RetributionAura"] = 3] = "RetributionAura";
})(PaladinAura || (PaladinAura = {}));
/**
 * @generated from protobuf enum proto.PaladinJudgement
 */
export var PaladinJudgement;
(function (PaladinJudgement) {
    /**
     * @generated from protobuf enum value: NoPaladinJudgement = 0;
     */
    PaladinJudgement[PaladinJudgement["NoPaladinJudgement"] = 0] = "NoPaladinJudgement";
    /**
     * @generated from protobuf enum value: JudgementOfWisdom = 1;
     */
    PaladinJudgement[PaladinJudgement["JudgementOfWisdom"] = 1] = "JudgementOfWisdom";
    /**
     * @generated from protobuf enum value: JudgementOfLight = 2;
     */
    PaladinJudgement[PaladinJudgement["JudgementOfLight"] = 2] = "JudgementOfLight";
    /**
     * @generated from protobuf enum value: JudgementOfCrusader = 3;
     */
    PaladinJudgement[PaladinJudgement["JudgementOfCrusader"] = 3] = "JudgementOfCrusader";
    /**
     * @generated from protobuf enum value: JudgementOfVengeance = 4;
     */
    PaladinJudgement[PaladinJudgement["JudgementOfVengeance"] = 4] = "JudgementOfVengeance";
    /**
     * @generated from protobuf enum value: JudgementOfRighteousness = 5;
     */
    PaladinJudgement[PaladinJudgement["JudgementOfRighteousness"] = 5] = "JudgementOfRighteousness";
})(PaladinJudgement || (PaladinJudgement = {}));
// @generated message type with reflection information, may provide speed optimized methods
class PaladinTalents$Type extends MessageType {
    constructor() {
        super("proto.PaladinTalents", [
            { no: 1, name: "divine_strength", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "divine_intellect", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "improved_seal_of_righteousness", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 34, name: "illumination", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "improved_blessing_of_wisdom", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "divine_favor", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 6, name: "purifying_power", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 7, name: "holy_power", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 8, name: "holy_shock", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 51, name: "blessed_life", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 9, name: "holy_guidance", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 10, name: "divine_illumination", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 35, name: "improved_devotion_aura", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 36, name: "redoubt", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 11, name: "precision", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 37, name: "toughness", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 12, name: "blessing_of_kings", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 38, name: "improved_righteous_fury", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 39, name: "shield_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 40, name: "anticipation", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 41, name: "spell_warding", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 42, name: "blessing_of_sanctuary", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 13, name: "reckoning", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 14, name: "sacred_duty", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 15, name: "one_handed_weapon_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 43, name: "improved_holy_shield", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 44, name: "holy_shield", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 45, name: "ardent_defender", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 16, name: "combat_expertise", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 17, name: "avengers_shield", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 18, name: "improved_blessing_of_might", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 19, name: "benediction", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 20, name: "improved_judgement", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 21, name: "improved_seal_of_the_crusader", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 46, name: "deflection", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 22, name: "vindication", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 23, name: "conviction", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 24, name: "seal_of_command", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 47, name: "pursuit_of_justice", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 48, name: "eye_for_an_eye", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 49, name: "improved_retribution_aura", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 25, name: "crusade", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 26, name: "two_handed_weapon_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 27, name: "sanctity_aura", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 28, name: "improved_sanctity_aura", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 29, name: "vengeance", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 30, name: "sanctified_judgement", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 31, name: "sanctified_seals", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 50, name: "divine_purpose", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 32, name: "fanaticism", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 33, name: "crusader_strike", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { divineStrength: 0, divineIntellect: 0, improvedSealOfRighteousness: 0, illumination: 0, improvedBlessingOfWisdom: 0, divineFavor: false, purifyingPower: 0, holyPower: 0, holyShock: false, blessedLife: 0, holyGuidance: 0, divineIllumination: false, improvedDevotionAura: 0, redoubt: 0, precision: 0, toughness: 0, blessingOfKings: false, improvedRighteousFury: 0, shieldSpecialization: 0, anticipation: 0, spellWarding: 0, blessingOfSanctuary: false, reckoning: 0, sacredDuty: 0, oneHandedWeaponSpecialization: 0, improvedHolyShield: 0, holyShield: false, ardentDefender: 0, combatExpertise: 0, avengersShield: false, improvedBlessingOfMight: 0, benediction: 0, improvedJudgement: 0, improvedSealOfTheCrusader: 0, deflection: 0, vindication: 0, conviction: 0, sealOfCommand: false, pursuitOfJustice: 0, eyeForAnEye: 0, improvedRetributionAura: 0, crusade: 0, twoHandedWeaponSpecialization: 0, sanctityAura: false, improvedSanctityAura: 0, vengeance: 0, sanctifiedJudgement: 0, sanctifiedSeals: 0, divinePurpose: 0, fanaticism: 0, crusaderStrike: false };
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
                case /* int32 divine_strength */ 1:
                    message.divineStrength = reader.int32();
                    break;
                case /* int32 divine_intellect */ 2:
                    message.divineIntellect = reader.int32();
                    break;
                case /* int32 improved_seal_of_righteousness */ 3:
                    message.improvedSealOfRighteousness = reader.int32();
                    break;
                case /* int32 illumination */ 34:
                    message.illumination = reader.int32();
                    break;
                case /* int32 improved_blessing_of_wisdom */ 4:
                    message.improvedBlessingOfWisdom = reader.int32();
                    break;
                case /* bool divine_favor */ 5:
                    message.divineFavor = reader.bool();
                    break;
                case /* int32 purifying_power */ 6:
                    message.purifyingPower = reader.int32();
                    break;
                case /* int32 holy_power */ 7:
                    message.holyPower = reader.int32();
                    break;
                case /* bool holy_shock */ 8:
                    message.holyShock = reader.bool();
                    break;
                case /* int32 blessed_life */ 51:
                    message.blessedLife = reader.int32();
                    break;
                case /* int32 holy_guidance */ 9:
                    message.holyGuidance = reader.int32();
                    break;
                case /* bool divine_illumination */ 10:
                    message.divineIllumination = reader.bool();
                    break;
                case /* int32 improved_devotion_aura */ 35:
                    message.improvedDevotionAura = reader.int32();
                    break;
                case /* int32 redoubt */ 36:
                    message.redoubt = reader.int32();
                    break;
                case /* int32 precision */ 11:
                    message.precision = reader.int32();
                    break;
                case /* int32 toughness */ 37:
                    message.toughness = reader.int32();
                    break;
                case /* bool blessing_of_kings */ 12:
                    message.blessingOfKings = reader.bool();
                    break;
                case /* int32 improved_righteous_fury */ 38:
                    message.improvedRighteousFury = reader.int32();
                    break;
                case /* int32 shield_specialization */ 39:
                    message.shieldSpecialization = reader.int32();
                    break;
                case /* int32 anticipation */ 40:
                    message.anticipation = reader.int32();
                    break;
                case /* int32 spell_warding */ 41:
                    message.spellWarding = reader.int32();
                    break;
                case /* bool blessing_of_sanctuary */ 42:
                    message.blessingOfSanctuary = reader.bool();
                    break;
                case /* int32 reckoning */ 13:
                    message.reckoning = reader.int32();
                    break;
                case /* int32 sacred_duty */ 14:
                    message.sacredDuty = reader.int32();
                    break;
                case /* int32 one_handed_weapon_specialization */ 15:
                    message.oneHandedWeaponSpecialization = reader.int32();
                    break;
                case /* int32 improved_holy_shield */ 43:
                    message.improvedHolyShield = reader.int32();
                    break;
                case /* bool holy_shield */ 44:
                    message.holyShield = reader.bool();
                    break;
                case /* int32 ardent_defender */ 45:
                    message.ardentDefender = reader.int32();
                    break;
                case /* int32 combat_expertise */ 16:
                    message.combatExpertise = reader.int32();
                    break;
                case /* bool avengers_shield */ 17:
                    message.avengersShield = reader.bool();
                    break;
                case /* int32 improved_blessing_of_might */ 18:
                    message.improvedBlessingOfMight = reader.int32();
                    break;
                case /* int32 benediction */ 19:
                    message.benediction = reader.int32();
                    break;
                case /* int32 improved_judgement */ 20:
                    message.improvedJudgement = reader.int32();
                    break;
                case /* int32 improved_seal_of_the_crusader */ 21:
                    message.improvedSealOfTheCrusader = reader.int32();
                    break;
                case /* int32 deflection */ 46:
                    message.deflection = reader.int32();
                    break;
                case /* int32 vindication */ 22:
                    message.vindication = reader.int32();
                    break;
                case /* int32 conviction */ 23:
                    message.conviction = reader.int32();
                    break;
                case /* bool seal_of_command */ 24:
                    message.sealOfCommand = reader.bool();
                    break;
                case /* int32 pursuit_of_justice */ 47:
                    message.pursuitOfJustice = reader.int32();
                    break;
                case /* int32 eye_for_an_eye */ 48:
                    message.eyeForAnEye = reader.int32();
                    break;
                case /* int32 improved_retribution_aura */ 49:
                    message.improvedRetributionAura = reader.int32();
                    break;
                case /* int32 crusade */ 25:
                    message.crusade = reader.int32();
                    break;
                case /* int32 two_handed_weapon_specialization */ 26:
                    message.twoHandedWeaponSpecialization = reader.int32();
                    break;
                case /* bool sanctity_aura */ 27:
                    message.sanctityAura = reader.bool();
                    break;
                case /* int32 improved_sanctity_aura */ 28:
                    message.improvedSanctityAura = reader.int32();
                    break;
                case /* int32 vengeance */ 29:
                    message.vengeance = reader.int32();
                    break;
                case /* int32 sanctified_judgement */ 30:
                    message.sanctifiedJudgement = reader.int32();
                    break;
                case /* int32 sanctified_seals */ 31:
                    message.sanctifiedSeals = reader.int32();
                    break;
                case /* int32 divine_purpose */ 50:
                    message.divinePurpose = reader.int32();
                    break;
                case /* int32 fanaticism */ 32:
                    message.fanaticism = reader.int32();
                    break;
                case /* bool crusader_strike */ 33:
                    message.crusaderStrike = reader.bool();
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
        /* int32 divine_strength = 1; */
        if (message.divineStrength !== 0)
            writer.tag(1, WireType.Varint).int32(message.divineStrength);
        /* int32 divine_intellect = 2; */
        if (message.divineIntellect !== 0)
            writer.tag(2, WireType.Varint).int32(message.divineIntellect);
        /* int32 improved_seal_of_righteousness = 3; */
        if (message.improvedSealOfRighteousness !== 0)
            writer.tag(3, WireType.Varint).int32(message.improvedSealOfRighteousness);
        /* int32 illumination = 34; */
        if (message.illumination !== 0)
            writer.tag(34, WireType.Varint).int32(message.illumination);
        /* int32 improved_blessing_of_wisdom = 4; */
        if (message.improvedBlessingOfWisdom !== 0)
            writer.tag(4, WireType.Varint).int32(message.improvedBlessingOfWisdom);
        /* bool divine_favor = 5; */
        if (message.divineFavor !== false)
            writer.tag(5, WireType.Varint).bool(message.divineFavor);
        /* int32 purifying_power = 6; */
        if (message.purifyingPower !== 0)
            writer.tag(6, WireType.Varint).int32(message.purifyingPower);
        /* int32 holy_power = 7; */
        if (message.holyPower !== 0)
            writer.tag(7, WireType.Varint).int32(message.holyPower);
        /* bool holy_shock = 8; */
        if (message.holyShock !== false)
            writer.tag(8, WireType.Varint).bool(message.holyShock);
        /* int32 blessed_life = 51; */
        if (message.blessedLife !== 0)
            writer.tag(51, WireType.Varint).int32(message.blessedLife);
        /* int32 holy_guidance = 9; */
        if (message.holyGuidance !== 0)
            writer.tag(9, WireType.Varint).int32(message.holyGuidance);
        /* bool divine_illumination = 10; */
        if (message.divineIllumination !== false)
            writer.tag(10, WireType.Varint).bool(message.divineIllumination);
        /* int32 improved_devotion_aura = 35; */
        if (message.improvedDevotionAura !== 0)
            writer.tag(35, WireType.Varint).int32(message.improvedDevotionAura);
        /* int32 redoubt = 36; */
        if (message.redoubt !== 0)
            writer.tag(36, WireType.Varint).int32(message.redoubt);
        /* int32 precision = 11; */
        if (message.precision !== 0)
            writer.tag(11, WireType.Varint).int32(message.precision);
        /* int32 toughness = 37; */
        if (message.toughness !== 0)
            writer.tag(37, WireType.Varint).int32(message.toughness);
        /* bool blessing_of_kings = 12; */
        if (message.blessingOfKings !== false)
            writer.tag(12, WireType.Varint).bool(message.blessingOfKings);
        /* int32 improved_righteous_fury = 38; */
        if (message.improvedRighteousFury !== 0)
            writer.tag(38, WireType.Varint).int32(message.improvedRighteousFury);
        /* int32 shield_specialization = 39; */
        if (message.shieldSpecialization !== 0)
            writer.tag(39, WireType.Varint).int32(message.shieldSpecialization);
        /* int32 anticipation = 40; */
        if (message.anticipation !== 0)
            writer.tag(40, WireType.Varint).int32(message.anticipation);
        /* int32 spell_warding = 41; */
        if (message.spellWarding !== 0)
            writer.tag(41, WireType.Varint).int32(message.spellWarding);
        /* bool blessing_of_sanctuary = 42; */
        if (message.blessingOfSanctuary !== false)
            writer.tag(42, WireType.Varint).bool(message.blessingOfSanctuary);
        /* int32 reckoning = 13; */
        if (message.reckoning !== 0)
            writer.tag(13, WireType.Varint).int32(message.reckoning);
        /* int32 sacred_duty = 14; */
        if (message.sacredDuty !== 0)
            writer.tag(14, WireType.Varint).int32(message.sacredDuty);
        /* int32 one_handed_weapon_specialization = 15; */
        if (message.oneHandedWeaponSpecialization !== 0)
            writer.tag(15, WireType.Varint).int32(message.oneHandedWeaponSpecialization);
        /* int32 improved_holy_shield = 43; */
        if (message.improvedHolyShield !== 0)
            writer.tag(43, WireType.Varint).int32(message.improvedHolyShield);
        /* bool holy_shield = 44; */
        if (message.holyShield !== false)
            writer.tag(44, WireType.Varint).bool(message.holyShield);
        /* int32 ardent_defender = 45; */
        if (message.ardentDefender !== 0)
            writer.tag(45, WireType.Varint).int32(message.ardentDefender);
        /* int32 combat_expertise = 16; */
        if (message.combatExpertise !== 0)
            writer.tag(16, WireType.Varint).int32(message.combatExpertise);
        /* bool avengers_shield = 17; */
        if (message.avengersShield !== false)
            writer.tag(17, WireType.Varint).bool(message.avengersShield);
        /* int32 improved_blessing_of_might = 18; */
        if (message.improvedBlessingOfMight !== 0)
            writer.tag(18, WireType.Varint).int32(message.improvedBlessingOfMight);
        /* int32 benediction = 19; */
        if (message.benediction !== 0)
            writer.tag(19, WireType.Varint).int32(message.benediction);
        /* int32 improved_judgement = 20; */
        if (message.improvedJudgement !== 0)
            writer.tag(20, WireType.Varint).int32(message.improvedJudgement);
        /* int32 improved_seal_of_the_crusader = 21; */
        if (message.improvedSealOfTheCrusader !== 0)
            writer.tag(21, WireType.Varint).int32(message.improvedSealOfTheCrusader);
        /* int32 deflection = 46; */
        if (message.deflection !== 0)
            writer.tag(46, WireType.Varint).int32(message.deflection);
        /* int32 vindication = 22; */
        if (message.vindication !== 0)
            writer.tag(22, WireType.Varint).int32(message.vindication);
        /* int32 conviction = 23; */
        if (message.conviction !== 0)
            writer.tag(23, WireType.Varint).int32(message.conviction);
        /* bool seal_of_command = 24; */
        if (message.sealOfCommand !== false)
            writer.tag(24, WireType.Varint).bool(message.sealOfCommand);
        /* int32 pursuit_of_justice = 47; */
        if (message.pursuitOfJustice !== 0)
            writer.tag(47, WireType.Varint).int32(message.pursuitOfJustice);
        /* int32 eye_for_an_eye = 48; */
        if (message.eyeForAnEye !== 0)
            writer.tag(48, WireType.Varint).int32(message.eyeForAnEye);
        /* int32 improved_retribution_aura = 49; */
        if (message.improvedRetributionAura !== 0)
            writer.tag(49, WireType.Varint).int32(message.improvedRetributionAura);
        /* int32 crusade = 25; */
        if (message.crusade !== 0)
            writer.tag(25, WireType.Varint).int32(message.crusade);
        /* int32 two_handed_weapon_specialization = 26; */
        if (message.twoHandedWeaponSpecialization !== 0)
            writer.tag(26, WireType.Varint).int32(message.twoHandedWeaponSpecialization);
        /* bool sanctity_aura = 27; */
        if (message.sanctityAura !== false)
            writer.tag(27, WireType.Varint).bool(message.sanctityAura);
        /* int32 improved_sanctity_aura = 28; */
        if (message.improvedSanctityAura !== 0)
            writer.tag(28, WireType.Varint).int32(message.improvedSanctityAura);
        /* int32 vengeance = 29; */
        if (message.vengeance !== 0)
            writer.tag(29, WireType.Varint).int32(message.vengeance);
        /* int32 sanctified_judgement = 30; */
        if (message.sanctifiedJudgement !== 0)
            writer.tag(30, WireType.Varint).int32(message.sanctifiedJudgement);
        /* int32 sanctified_seals = 31; */
        if (message.sanctifiedSeals !== 0)
            writer.tag(31, WireType.Varint).int32(message.sanctifiedSeals);
        /* int32 divine_purpose = 50; */
        if (message.divinePurpose !== 0)
            writer.tag(50, WireType.Varint).int32(message.divinePurpose);
        /* int32 fanaticism = 32; */
        if (message.fanaticism !== 0)
            writer.tag(32, WireType.Varint).int32(message.fanaticism);
        /* bool crusader_strike = 33; */
        if (message.crusaderStrike !== false)
            writer.tag(33, WireType.Varint).bool(message.crusaderStrike);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.PaladinTalents
 */
export const PaladinTalents = new PaladinTalents$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RetributionPaladin$Type extends MessageType {
    constructor() {
        super("proto.RetributionPaladin", [
            { no: 1, name: "rotation", kind: "message", T: () => RetributionPaladin_Rotation },
            { no: 2, name: "talents", kind: "message", T: () => PaladinTalents },
            { no: 3, name: "options", kind: "message", T: () => RetributionPaladin_Options }
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
                case /* proto.RetributionPaladin.Rotation rotation */ 1:
                    message.rotation = RetributionPaladin_Rotation.internalBinaryRead(reader, reader.uint32(), options, message.rotation);
                    break;
                case /* proto.PaladinTalents talents */ 2:
                    message.talents = PaladinTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* proto.RetributionPaladin.Options options */ 3:
                    message.options = RetributionPaladin_Options.internalBinaryRead(reader, reader.uint32(), options, message.options);
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
        /* proto.RetributionPaladin.Rotation rotation = 1; */
        if (message.rotation)
            RetributionPaladin_Rotation.internalBinaryWrite(message.rotation, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.PaladinTalents talents = 2; */
        if (message.talents)
            PaladinTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.RetributionPaladin.Options options = 3; */
        if (message.options)
            RetributionPaladin_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.RetributionPaladin
 */
export const RetributionPaladin = new RetributionPaladin$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RetributionPaladin_Rotation$Type extends MessageType {
    constructor() {
        super("proto.RetributionPaladin.Rotation", [
            { no: 1, name: "consecration_rank", kind: "enum", T: () => ["proto.RetributionPaladin.Rotation.ConsecrationRank", RetributionPaladin_Rotation_ConsecrationRank] },
            { no: 2, name: "use_exorcism", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { consecrationRank: 0, useExorcism: false };
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
                case /* proto.RetributionPaladin.Rotation.ConsecrationRank consecration_rank */ 1:
                    message.consecrationRank = reader.int32();
                    break;
                case /* bool use_exorcism */ 2:
                    message.useExorcism = reader.bool();
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
        /* proto.RetributionPaladin.Rotation.ConsecrationRank consecration_rank = 1; */
        if (message.consecrationRank !== 0)
            writer.tag(1, WireType.Varint).int32(message.consecrationRank);
        /* bool use_exorcism = 2; */
        if (message.useExorcism !== false)
            writer.tag(2, WireType.Varint).bool(message.useExorcism);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.RetributionPaladin.Rotation
 */
export const RetributionPaladin_Rotation = new RetributionPaladin_Rotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RetributionPaladin_Options$Type extends MessageType {
    constructor() {
        super("proto.RetributionPaladin.Options", [
            { no: 1, name: "judgement", kind: "enum", T: () => ["proto.RetributionPaladin.Options.Judgement", RetributionPaladin_Options_Judgement] },
            { no: 5, name: "aura", kind: "enum", T: () => ["proto.PaladinAura", PaladinAura] },
            { no: 2, name: "crusader_strike_delay_ms", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "haste_leeway_ms", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "damage_taken_per_second", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { judgement: 0, aura: 0, crusaderStrikeDelayMs: 0, hasteLeewayMs: 0, damageTakenPerSecond: 0 };
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
                case /* proto.RetributionPaladin.Options.Judgement judgement */ 1:
                    message.judgement = reader.int32();
                    break;
                case /* proto.PaladinAura aura */ 5:
                    message.aura = reader.int32();
                    break;
                case /* int32 crusader_strike_delay_ms */ 2:
                    message.crusaderStrikeDelayMs = reader.int32();
                    break;
                case /* int32 haste_leeway_ms */ 3:
                    message.hasteLeewayMs = reader.int32();
                    break;
                case /* double damage_taken_per_second */ 4:
                    message.damageTakenPerSecond = reader.double();
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
        /* proto.RetributionPaladin.Options.Judgement judgement = 1; */
        if (message.judgement !== 0)
            writer.tag(1, WireType.Varint).int32(message.judgement);
        /* proto.PaladinAura aura = 5; */
        if (message.aura !== 0)
            writer.tag(5, WireType.Varint).int32(message.aura);
        /* int32 crusader_strike_delay_ms = 2; */
        if (message.crusaderStrikeDelayMs !== 0)
            writer.tag(2, WireType.Varint).int32(message.crusaderStrikeDelayMs);
        /* int32 haste_leeway_ms = 3; */
        if (message.hasteLeewayMs !== 0)
            writer.tag(3, WireType.Varint).int32(message.hasteLeewayMs);
        /* double damage_taken_per_second = 4; */
        if (message.damageTakenPerSecond !== 0)
            writer.tag(4, WireType.Bit64).double(message.damageTakenPerSecond);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.RetributionPaladin.Options
 */
export const RetributionPaladin_Options = new RetributionPaladin_Options$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ProtectionPaladin$Type extends MessageType {
    constructor() {
        super("proto.ProtectionPaladin", [
            { no: 1, name: "rotation", kind: "message", T: () => ProtectionPaladin_Rotation },
            { no: 2, name: "talents", kind: "message", T: () => PaladinTalents },
            { no: 3, name: "options", kind: "message", T: () => ProtectionPaladin_Options }
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
                case /* proto.ProtectionPaladin.Rotation rotation */ 1:
                    message.rotation = ProtectionPaladin_Rotation.internalBinaryRead(reader, reader.uint32(), options, message.rotation);
                    break;
                case /* proto.PaladinTalents talents */ 2:
                    message.talents = PaladinTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* proto.ProtectionPaladin.Options options */ 3:
                    message.options = ProtectionPaladin_Options.internalBinaryRead(reader, reader.uint32(), options, message.options);
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
        /* proto.ProtectionPaladin.Rotation rotation = 1; */
        if (message.rotation)
            ProtectionPaladin_Rotation.internalBinaryWrite(message.rotation, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.PaladinTalents talents = 2; */
        if (message.talents)
            PaladinTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.ProtectionPaladin.Options options = 3; */
        if (message.options)
            ProtectionPaladin_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ProtectionPaladin
 */
export const ProtectionPaladin = new ProtectionPaladin$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ProtectionPaladin_Rotation$Type extends MessageType {
    constructor() {
        super("proto.ProtectionPaladin.Rotation", [
            { no: 1, name: "prioritize_holy_shield", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "consecration_rank", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "use_exorcism", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "maintain_judgement", kind: "enum", T: () => ["proto.PaladinJudgement", PaladinJudgement] }
        ]);
    }
    create(value) {
        const message = { prioritizeHolyShield: false, consecrationRank: 0, useExorcism: false, maintainJudgement: 0 };
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
                case /* bool prioritize_holy_shield */ 1:
                    message.prioritizeHolyShield = reader.bool();
                    break;
                case /* int32 consecration_rank */ 2:
                    message.consecrationRank = reader.int32();
                    break;
                case /* bool use_exorcism */ 3:
                    message.useExorcism = reader.bool();
                    break;
                case /* proto.PaladinJudgement maintain_judgement */ 4:
                    message.maintainJudgement = reader.int32();
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
        /* bool prioritize_holy_shield = 1; */
        if (message.prioritizeHolyShield !== false)
            writer.tag(1, WireType.Varint).bool(message.prioritizeHolyShield);
        /* int32 consecration_rank = 2; */
        if (message.consecrationRank !== 0)
            writer.tag(2, WireType.Varint).int32(message.consecrationRank);
        /* bool use_exorcism = 3; */
        if (message.useExorcism !== false)
            writer.tag(3, WireType.Varint).bool(message.useExorcism);
        /* proto.PaladinJudgement maintain_judgement = 4; */
        if (message.maintainJudgement !== 0)
            writer.tag(4, WireType.Varint).int32(message.maintainJudgement);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ProtectionPaladin.Rotation
 */
export const ProtectionPaladin_Rotation = new ProtectionPaladin_Rotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ProtectionPaladin_Options$Type extends MessageType {
    constructor() {
        super("proto.ProtectionPaladin.Options", [
            { no: 1, name: "aura", kind: "enum", T: () => ["proto.PaladinAura", PaladinAura] },
            { no: 2, name: "use_avenging_wrath", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { aura: 0, useAvengingWrath: false };
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
                case /* proto.PaladinAura aura */ 1:
                    message.aura = reader.int32();
                    break;
                case /* bool use_avenging_wrath */ 2:
                    message.useAvengingWrath = reader.bool();
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
        /* proto.PaladinAura aura = 1; */
        if (message.aura !== 0)
            writer.tag(1, WireType.Varint).int32(message.aura);
        /* bool use_avenging_wrath = 2; */
        if (message.useAvengingWrath !== false)
            writer.tag(2, WireType.Varint).bool(message.useAvengingWrath);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ProtectionPaladin.Options
 */
export const ProtectionPaladin_Options = new ProtectionPaladin_Options$Type();
