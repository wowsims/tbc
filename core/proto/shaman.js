import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
/**
 * @generated from protobuf enum proto.ElementalShaman.Rotation.RotationType
 */
export var ElementalShaman_Rotation_RotationType;
(function (ElementalShaman_Rotation_RotationType) {
    /**
     * @generated from protobuf enum value: Unknown = 0;
     */
    ElementalShaman_Rotation_RotationType[ElementalShaman_Rotation_RotationType["Unknown"] = 0] = "Unknown";
    /**
     * @generated from protobuf enum value: Adaptive = 1;
     */
    ElementalShaman_Rotation_RotationType[ElementalShaman_Rotation_RotationType["Adaptive"] = 1] = "Adaptive";
    /**
     * @generated from protobuf enum value: CLOnClearcast = 2;
     */
    ElementalShaman_Rotation_RotationType[ElementalShaman_Rotation_RotationType["CLOnClearcast"] = 2] = "CLOnClearcast";
    /**
     * @generated from protobuf enum value: CLOnCD = 3;
     */
    ElementalShaman_Rotation_RotationType[ElementalShaman_Rotation_RotationType["CLOnCD"] = 3] = "CLOnCD";
    /**
     * @generated from protobuf enum value: FixedLBCL = 4;
     */
    ElementalShaman_Rotation_RotationType[ElementalShaman_Rotation_RotationType["FixedLBCL"] = 4] = "FixedLBCL";
    /**
     * @generated from protobuf enum value: LBOnly = 5;
     */
    ElementalShaman_Rotation_RotationType[ElementalShaman_Rotation_RotationType["LBOnly"] = 5] = "LBOnly";
})(ElementalShaman_Rotation_RotationType || (ElementalShaman_Rotation_RotationType = {}));
/**
 * @generated from protobuf enum proto.EnhancementShaman.Rotation.PrimaryShock
 */
export var EnhancementShaman_Rotation_PrimaryShock;
(function (EnhancementShaman_Rotation_PrimaryShock) {
    /**
     * @generated from protobuf enum value: None = 0;
     */
    EnhancementShaman_Rotation_PrimaryShock[EnhancementShaman_Rotation_PrimaryShock["None"] = 0] = "None";
    /**
     * @generated from protobuf enum value: Earth = 1;
     */
    EnhancementShaman_Rotation_PrimaryShock[EnhancementShaman_Rotation_PrimaryShock["Earth"] = 1] = "Earth";
    /**
     * @generated from protobuf enum value: Frost = 2;
     */
    EnhancementShaman_Rotation_PrimaryShock[EnhancementShaman_Rotation_PrimaryShock["Frost"] = 2] = "Frost";
})(EnhancementShaman_Rotation_PrimaryShock || (EnhancementShaman_Rotation_PrimaryShock = {}));
/**
 * @generated from protobuf enum proto.EarthTotem
 */
export var EarthTotem;
(function (EarthTotem) {
    /**
     * @generated from protobuf enum value: NoEarthTotem = 0;
     */
    EarthTotem[EarthTotem["NoEarthTotem"] = 0] = "NoEarthTotem";
    /**
     * @generated from protobuf enum value: StrengthOfEarthTotem = 1;
     */
    EarthTotem[EarthTotem["StrengthOfEarthTotem"] = 1] = "StrengthOfEarthTotem";
    /**
     * @generated from protobuf enum value: TremorTotem = 2;
     */
    EarthTotem[EarthTotem["TremorTotem"] = 2] = "TremorTotem";
})(EarthTotem || (EarthTotem = {}));
/**
 * @generated from protobuf enum proto.AirTotem
 */
export var AirTotem;
(function (AirTotem) {
    /**
     * @generated from protobuf enum value: NoAirTotem = 0;
     */
    AirTotem[AirTotem["NoAirTotem"] = 0] = "NoAirTotem";
    /**
     * @generated from protobuf enum value: GraceOfAirTotem = 1;
     */
    AirTotem[AirTotem["GraceOfAirTotem"] = 1] = "GraceOfAirTotem";
    /**
     * @generated from protobuf enum value: TranquilAirTotem = 2;
     */
    AirTotem[AirTotem["TranquilAirTotem"] = 2] = "TranquilAirTotem";
    /**
     * @generated from protobuf enum value: WindfuryTotem = 3;
     */
    AirTotem[AirTotem["WindfuryTotem"] = 3] = "WindfuryTotem";
    /**
     * @generated from protobuf enum value: WrathOfAirTotem = 4;
     */
    AirTotem[AirTotem["WrathOfAirTotem"] = 4] = "WrathOfAirTotem";
})(AirTotem || (AirTotem = {}));
/**
 * @generated from protobuf enum proto.FireTotem
 */
export var FireTotem;
(function (FireTotem) {
    /**
     * @generated from protobuf enum value: NoFireTotem = 0;
     */
    FireTotem[FireTotem["NoFireTotem"] = 0] = "NoFireTotem";
    /**
     * @generated from protobuf enum value: MagmaTotem = 1;
     */
    FireTotem[FireTotem["MagmaTotem"] = 1] = "MagmaTotem";
    /**
     * @generated from protobuf enum value: SearingTotem = 2;
     */
    FireTotem[FireTotem["SearingTotem"] = 2] = "SearingTotem";
    /**
     * @generated from protobuf enum value: TotemOfWrath = 3;
     */
    FireTotem[FireTotem["TotemOfWrath"] = 3] = "TotemOfWrath";
    /**
     * @generated from protobuf enum value: FireNovaTotem = 4;
     */
    FireTotem[FireTotem["FireNovaTotem"] = 4] = "FireNovaTotem";
})(FireTotem || (FireTotem = {}));
/**
 * @generated from protobuf enum proto.WaterTotem
 */
export var WaterTotem;
(function (WaterTotem) {
    /**
     * @generated from protobuf enum value: NoWaterTotem = 0;
     */
    WaterTotem[WaterTotem["NoWaterTotem"] = 0] = "NoWaterTotem";
    /**
     * @generated from protobuf enum value: ManaSpringTotem = 1;
     */
    WaterTotem[WaterTotem["ManaSpringTotem"] = 1] = "ManaSpringTotem";
})(WaterTotem || (WaterTotem = {}));
// @generated message type with reflection information, may provide speed optimized methods
class ShamanTalents$Type extends MessageType {
    constructor() {
        super("proto.ShamanTalents", [
            { no: 1, name: "convection", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "concussion", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "call_of_flame", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "elemental_focus", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "reverberation", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "call_of_thunder", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 7, name: "improved_fire_totems", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 8, name: "elemental_devastation", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 9, name: "elemental_fury", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 10, name: "unrelenting_storm", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 11, name: "elemental_precision", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 12, name: "lightning_mastery", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 13, name: "elemental_mastery", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 14, name: "lightning_overload", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 33, name: "totemOfWrath", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 15, name: "ancestral_knowledge", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 37, name: "shield_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 16, name: "thundering_strikes", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 17, name: "enhancing_totems", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 18, name: "shamanistic_focus", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 38, name: "anticipation", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 19, name: "flurry", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 39, name: "toughness", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 20, name: "improved_weapon_totems", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 36, name: "spirit_weapons", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 21, name: "elemental_weapons", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 22, name: "mental_quickness", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 23, name: "weapon_mastery", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 24, name: "dual_wield_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 25, name: "unleashed_rage", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 34, name: "stormstrike", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 35, name: "shamanistic_rage", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 26, name: "totemic_focus", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 27, name: "natures_guidance", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 28, name: "restorative_totems", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 29, name: "tidal_mastery", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 30, name: "natures_swiftness", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 31, name: "mana_tide_totem", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 32, name: "natures_blessing", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value) {
        const message = { convection: 0, concussion: 0, callOfFlame: 0, elementalFocus: false, reverberation: 0, callOfThunder: 0, improvedFireTotems: 0, elementalDevastation: 0, elementalFury: false, unrelentingStorm: 0, elementalPrecision: 0, lightningMastery: 0, elementalMastery: false, lightningOverload: 0, totemOfWrath: false, ancestralKnowledge: 0, shieldSpecialization: 0, thunderingStrikes: 0, enhancingTotems: 0, shamanisticFocus: false, anticipation: 0, flurry: 0, toughness: 0, improvedWeaponTotems: 0, spiritWeapons: false, elementalWeapons: 0, mentalQuickness: 0, weaponMastery: 0, dualWieldSpecialization: 0, unleashedRage: 0, stormstrike: false, shamanisticRage: false, totemicFocus: 0, naturesGuidance: 0, restorativeTotems: 0, tidalMastery: 0, naturesSwiftness: false, manaTideTotem: false, naturesBlessing: 0 };
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
                case /* int32 convection */ 1:
                    message.convection = reader.int32();
                    break;
                case /* int32 concussion */ 2:
                    message.concussion = reader.int32();
                    break;
                case /* int32 call_of_flame */ 3:
                    message.callOfFlame = reader.int32();
                    break;
                case /* bool elemental_focus */ 4:
                    message.elementalFocus = reader.bool();
                    break;
                case /* int32 reverberation */ 5:
                    message.reverberation = reader.int32();
                    break;
                case /* int32 call_of_thunder */ 6:
                    message.callOfThunder = reader.int32();
                    break;
                case /* int32 improved_fire_totems */ 7:
                    message.improvedFireTotems = reader.int32();
                    break;
                case /* int32 elemental_devastation */ 8:
                    message.elementalDevastation = reader.int32();
                    break;
                case /* bool elemental_fury */ 9:
                    message.elementalFury = reader.bool();
                    break;
                case /* int32 unrelenting_storm */ 10:
                    message.unrelentingStorm = reader.int32();
                    break;
                case /* int32 elemental_precision */ 11:
                    message.elementalPrecision = reader.int32();
                    break;
                case /* int32 lightning_mastery */ 12:
                    message.lightningMastery = reader.int32();
                    break;
                case /* bool elemental_mastery */ 13:
                    message.elementalMastery = reader.bool();
                    break;
                case /* int32 lightning_overload */ 14:
                    message.lightningOverload = reader.int32();
                    break;
                case /* bool totemOfWrath */ 33:
                    message.totemOfWrath = reader.bool();
                    break;
                case /* int32 ancestral_knowledge */ 15:
                    message.ancestralKnowledge = reader.int32();
                    break;
                case /* int32 shield_specialization */ 37:
                    message.shieldSpecialization = reader.int32();
                    break;
                case /* int32 thundering_strikes */ 16:
                    message.thunderingStrikes = reader.int32();
                    break;
                case /* int32 enhancing_totems */ 17:
                    message.enhancingTotems = reader.int32();
                    break;
                case /* bool shamanistic_focus */ 18:
                    message.shamanisticFocus = reader.bool();
                    break;
                case /* int32 anticipation */ 38:
                    message.anticipation = reader.int32();
                    break;
                case /* int32 flurry */ 19:
                    message.flurry = reader.int32();
                    break;
                case /* int32 toughness */ 39:
                    message.toughness = reader.int32();
                    break;
                case /* int32 improved_weapon_totems */ 20:
                    message.improvedWeaponTotems = reader.int32();
                    break;
                case /* bool spirit_weapons */ 36:
                    message.spiritWeapons = reader.bool();
                    break;
                case /* int32 elemental_weapons */ 21:
                    message.elementalWeapons = reader.int32();
                    break;
                case /* int32 mental_quickness */ 22:
                    message.mentalQuickness = reader.int32();
                    break;
                case /* int32 weapon_mastery */ 23:
                    message.weaponMastery = reader.int32();
                    break;
                case /* int32 dual_wield_specialization */ 24:
                    message.dualWieldSpecialization = reader.int32();
                    break;
                case /* int32 unleashed_rage */ 25:
                    message.unleashedRage = reader.int32();
                    break;
                case /* bool stormstrike */ 34:
                    message.stormstrike = reader.bool();
                    break;
                case /* bool shamanistic_rage */ 35:
                    message.shamanisticRage = reader.bool();
                    break;
                case /* int32 totemic_focus */ 26:
                    message.totemicFocus = reader.int32();
                    break;
                case /* int32 natures_guidance */ 27:
                    message.naturesGuidance = reader.int32();
                    break;
                case /* int32 restorative_totems */ 28:
                    message.restorativeTotems = reader.int32();
                    break;
                case /* int32 tidal_mastery */ 29:
                    message.tidalMastery = reader.int32();
                    break;
                case /* bool natures_swiftness */ 30:
                    message.naturesSwiftness = reader.bool();
                    break;
                case /* bool mana_tide_totem */ 31:
                    message.manaTideTotem = reader.bool();
                    break;
                case /* int32 natures_blessing */ 32:
                    message.naturesBlessing = reader.int32();
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
        /* int32 convection = 1; */
        if (message.convection !== 0)
            writer.tag(1, WireType.Varint).int32(message.convection);
        /* int32 concussion = 2; */
        if (message.concussion !== 0)
            writer.tag(2, WireType.Varint).int32(message.concussion);
        /* int32 call_of_flame = 3; */
        if (message.callOfFlame !== 0)
            writer.tag(3, WireType.Varint).int32(message.callOfFlame);
        /* bool elemental_focus = 4; */
        if (message.elementalFocus !== false)
            writer.tag(4, WireType.Varint).bool(message.elementalFocus);
        /* int32 reverberation = 5; */
        if (message.reverberation !== 0)
            writer.tag(5, WireType.Varint).int32(message.reverberation);
        /* int32 call_of_thunder = 6; */
        if (message.callOfThunder !== 0)
            writer.tag(6, WireType.Varint).int32(message.callOfThunder);
        /* int32 improved_fire_totems = 7; */
        if (message.improvedFireTotems !== 0)
            writer.tag(7, WireType.Varint).int32(message.improvedFireTotems);
        /* int32 elemental_devastation = 8; */
        if (message.elementalDevastation !== 0)
            writer.tag(8, WireType.Varint).int32(message.elementalDevastation);
        /* bool elemental_fury = 9; */
        if (message.elementalFury !== false)
            writer.tag(9, WireType.Varint).bool(message.elementalFury);
        /* int32 unrelenting_storm = 10; */
        if (message.unrelentingStorm !== 0)
            writer.tag(10, WireType.Varint).int32(message.unrelentingStorm);
        /* int32 elemental_precision = 11; */
        if (message.elementalPrecision !== 0)
            writer.tag(11, WireType.Varint).int32(message.elementalPrecision);
        /* int32 lightning_mastery = 12; */
        if (message.lightningMastery !== 0)
            writer.tag(12, WireType.Varint).int32(message.lightningMastery);
        /* bool elemental_mastery = 13; */
        if (message.elementalMastery !== false)
            writer.tag(13, WireType.Varint).bool(message.elementalMastery);
        /* int32 lightning_overload = 14; */
        if (message.lightningOverload !== 0)
            writer.tag(14, WireType.Varint).int32(message.lightningOverload);
        /* bool totemOfWrath = 33; */
        if (message.totemOfWrath !== false)
            writer.tag(33, WireType.Varint).bool(message.totemOfWrath);
        /* int32 ancestral_knowledge = 15; */
        if (message.ancestralKnowledge !== 0)
            writer.tag(15, WireType.Varint).int32(message.ancestralKnowledge);
        /* int32 shield_specialization = 37; */
        if (message.shieldSpecialization !== 0)
            writer.tag(37, WireType.Varint).int32(message.shieldSpecialization);
        /* int32 thundering_strikes = 16; */
        if (message.thunderingStrikes !== 0)
            writer.tag(16, WireType.Varint).int32(message.thunderingStrikes);
        /* int32 enhancing_totems = 17; */
        if (message.enhancingTotems !== 0)
            writer.tag(17, WireType.Varint).int32(message.enhancingTotems);
        /* bool shamanistic_focus = 18; */
        if (message.shamanisticFocus !== false)
            writer.tag(18, WireType.Varint).bool(message.shamanisticFocus);
        /* int32 anticipation = 38; */
        if (message.anticipation !== 0)
            writer.tag(38, WireType.Varint).int32(message.anticipation);
        /* int32 flurry = 19; */
        if (message.flurry !== 0)
            writer.tag(19, WireType.Varint).int32(message.flurry);
        /* int32 toughness = 39; */
        if (message.toughness !== 0)
            writer.tag(39, WireType.Varint).int32(message.toughness);
        /* int32 improved_weapon_totems = 20; */
        if (message.improvedWeaponTotems !== 0)
            writer.tag(20, WireType.Varint).int32(message.improvedWeaponTotems);
        /* bool spirit_weapons = 36; */
        if (message.spiritWeapons !== false)
            writer.tag(36, WireType.Varint).bool(message.spiritWeapons);
        /* int32 elemental_weapons = 21; */
        if (message.elementalWeapons !== 0)
            writer.tag(21, WireType.Varint).int32(message.elementalWeapons);
        /* int32 mental_quickness = 22; */
        if (message.mentalQuickness !== 0)
            writer.tag(22, WireType.Varint).int32(message.mentalQuickness);
        /* int32 weapon_mastery = 23; */
        if (message.weaponMastery !== 0)
            writer.tag(23, WireType.Varint).int32(message.weaponMastery);
        /* int32 dual_wield_specialization = 24; */
        if (message.dualWieldSpecialization !== 0)
            writer.tag(24, WireType.Varint).int32(message.dualWieldSpecialization);
        /* int32 unleashed_rage = 25; */
        if (message.unleashedRage !== 0)
            writer.tag(25, WireType.Varint).int32(message.unleashedRage);
        /* bool stormstrike = 34; */
        if (message.stormstrike !== false)
            writer.tag(34, WireType.Varint).bool(message.stormstrike);
        /* bool shamanistic_rage = 35; */
        if (message.shamanisticRage !== false)
            writer.tag(35, WireType.Varint).bool(message.shamanisticRage);
        /* int32 totemic_focus = 26; */
        if (message.totemicFocus !== 0)
            writer.tag(26, WireType.Varint).int32(message.totemicFocus);
        /* int32 natures_guidance = 27; */
        if (message.naturesGuidance !== 0)
            writer.tag(27, WireType.Varint).int32(message.naturesGuidance);
        /* int32 restorative_totems = 28; */
        if (message.restorativeTotems !== 0)
            writer.tag(28, WireType.Varint).int32(message.restorativeTotems);
        /* int32 tidal_mastery = 29; */
        if (message.tidalMastery !== 0)
            writer.tag(29, WireType.Varint).int32(message.tidalMastery);
        /* bool natures_swiftness = 30; */
        if (message.naturesSwiftness !== false)
            writer.tag(30, WireType.Varint).bool(message.naturesSwiftness);
        /* bool mana_tide_totem = 31; */
        if (message.manaTideTotem !== false)
            writer.tag(31, WireType.Varint).bool(message.manaTideTotem);
        /* int32 natures_blessing = 32; */
        if (message.naturesBlessing !== 0)
            writer.tag(32, WireType.Varint).int32(message.naturesBlessing);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ShamanTalents
 */
export const ShamanTalents = new ShamanTalents$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ShamanTotems$Type extends MessageType {
    constructor() {
        super("proto.ShamanTotems", [
            { no: 1, name: "earth", kind: "enum", T: () => ["proto.EarthTotem", EarthTotem] },
            { no: 2, name: "air", kind: "enum", T: () => ["proto.AirTotem", AirTotem] },
            { no: 3, name: "fire", kind: "enum", T: () => ["proto.FireTotem", FireTotem] },
            { no: 4, name: "water", kind: "enum", T: () => ["proto.WaterTotem", WaterTotem] },
            { no: 5, name: "twist_windfury", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 11, name: "windfury_totem_rank", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "twist_fire_nova", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 7, name: "use_mana_tide", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 8, name: "use_fire_elemental", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 9, name: "recall_fire_elemental_on_oom", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 10, name: "recall_totems", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { earth: 0, air: 0, fire: 0, water: 0, twistWindfury: false, windfuryTotemRank: 0, twistFireNova: false, useManaTide: false, useFireElemental: false, recallFireElementalOnOom: false, recallTotems: false };
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
                case /* proto.EarthTotem earth */ 1:
                    message.earth = reader.int32();
                    break;
                case /* proto.AirTotem air */ 2:
                    message.air = reader.int32();
                    break;
                case /* proto.FireTotem fire */ 3:
                    message.fire = reader.int32();
                    break;
                case /* proto.WaterTotem water */ 4:
                    message.water = reader.int32();
                    break;
                case /* bool twist_windfury */ 5:
                    message.twistWindfury = reader.bool();
                    break;
                case /* int32 windfury_totem_rank */ 11:
                    message.windfuryTotemRank = reader.int32();
                    break;
                case /* bool twist_fire_nova */ 6:
                    message.twistFireNova = reader.bool();
                    break;
                case /* bool use_mana_tide */ 7:
                    message.useManaTide = reader.bool();
                    break;
                case /* bool use_fire_elemental */ 8:
                    message.useFireElemental = reader.bool();
                    break;
                case /* bool recall_fire_elemental_on_oom */ 9:
                    message.recallFireElementalOnOom = reader.bool();
                    break;
                case /* bool recall_totems */ 10:
                    message.recallTotems = reader.bool();
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
        /* proto.EarthTotem earth = 1; */
        if (message.earth !== 0)
            writer.tag(1, WireType.Varint).int32(message.earth);
        /* proto.AirTotem air = 2; */
        if (message.air !== 0)
            writer.tag(2, WireType.Varint).int32(message.air);
        /* proto.FireTotem fire = 3; */
        if (message.fire !== 0)
            writer.tag(3, WireType.Varint).int32(message.fire);
        /* proto.WaterTotem water = 4; */
        if (message.water !== 0)
            writer.tag(4, WireType.Varint).int32(message.water);
        /* bool twist_windfury = 5; */
        if (message.twistWindfury !== false)
            writer.tag(5, WireType.Varint).bool(message.twistWindfury);
        /* int32 windfury_totem_rank = 11; */
        if (message.windfuryTotemRank !== 0)
            writer.tag(11, WireType.Varint).int32(message.windfuryTotemRank);
        /* bool twist_fire_nova = 6; */
        if (message.twistFireNova !== false)
            writer.tag(6, WireType.Varint).bool(message.twistFireNova);
        /* bool use_mana_tide = 7; */
        if (message.useManaTide !== false)
            writer.tag(7, WireType.Varint).bool(message.useManaTide);
        /* bool use_fire_elemental = 8; */
        if (message.useFireElemental !== false)
            writer.tag(8, WireType.Varint).bool(message.useFireElemental);
        /* bool recall_fire_elemental_on_oom = 9; */
        if (message.recallFireElementalOnOom !== false)
            writer.tag(9, WireType.Varint).bool(message.recallFireElementalOnOom);
        /* bool recall_totems = 10; */
        if (message.recallTotems !== false)
            writer.tag(10, WireType.Varint).bool(message.recallTotems);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ShamanTotems
 */
export const ShamanTotems = new ShamanTotems$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ElementalShaman$Type extends MessageType {
    constructor() {
        super("proto.ElementalShaman", [
            { no: 1, name: "rotation", kind: "message", T: () => ElementalShaman_Rotation },
            { no: 2, name: "talents", kind: "message", T: () => ShamanTalents },
            { no: 3, name: "options", kind: "message", T: () => ElementalShaman_Options }
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
                case /* proto.ElementalShaman.Rotation rotation */ 1:
                    message.rotation = ElementalShaman_Rotation.internalBinaryRead(reader, reader.uint32(), options, message.rotation);
                    break;
                case /* proto.ShamanTalents talents */ 2:
                    message.talents = ShamanTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* proto.ElementalShaman.Options options */ 3:
                    message.options = ElementalShaman_Options.internalBinaryRead(reader, reader.uint32(), options, message.options);
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
        /* proto.ElementalShaman.Rotation rotation = 1; */
        if (message.rotation)
            ElementalShaman_Rotation.internalBinaryWrite(message.rotation, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.ShamanTalents talents = 2; */
        if (message.talents)
            ShamanTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.ElementalShaman.Options options = 3; */
        if (message.options)
            ElementalShaman_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ElementalShaman
 */
export const ElementalShaman = new ElementalShaman$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ElementalShaman_Rotation$Type extends MessageType {
    constructor() {
        super("proto.ElementalShaman.Rotation", [
            { no: 3, name: "totems", kind: "message", T: () => ShamanTotems },
            { no: 1, name: "type", kind: "enum", T: () => ["proto.ElementalShaman.Rotation.RotationType", ElementalShaman_Rotation_RotationType] },
            { no: 2, name: "lbs_per_cl", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value) {
        const message = { type: 0, lbsPerCl: 0 };
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
                case /* proto.ShamanTotems totems */ 3:
                    message.totems = ShamanTotems.internalBinaryRead(reader, reader.uint32(), options, message.totems);
                    break;
                case /* proto.ElementalShaman.Rotation.RotationType type */ 1:
                    message.type = reader.int32();
                    break;
                case /* int32 lbs_per_cl */ 2:
                    message.lbsPerCl = reader.int32();
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
        /* proto.ShamanTotems totems = 3; */
        if (message.totems)
            ShamanTotems.internalBinaryWrite(message.totems, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* proto.ElementalShaman.Rotation.RotationType type = 1; */
        if (message.type !== 0)
            writer.tag(1, WireType.Varint).int32(message.type);
        /* int32 lbs_per_cl = 2; */
        if (message.lbsPerCl !== 0)
            writer.tag(2, WireType.Varint).int32(message.lbsPerCl);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ElementalShaman.Rotation
 */
export const ElementalShaman_Rotation = new ElementalShaman_Rotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ElementalShaman_Options$Type extends MessageType {
    constructor() {
        super("proto.ElementalShaman.Options", [
            { no: 1, name: "water_shield", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "bloodlust", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 6, name: "snapshot_t4_2pc", kind: "scalar", jsonName: "snapshotT42pc", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { waterShield: false, bloodlust: false, snapshotT42Pc: false };
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
                case /* bool water_shield */ 1:
                    message.waterShield = reader.bool();
                    break;
                case /* bool bloodlust */ 2:
                    message.bloodlust = reader.bool();
                    break;
                case /* bool snapshot_t4_2pc = 6 [json_name = "snapshotT42pc"];*/ 6:
                    message.snapshotT42Pc = reader.bool();
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
        /* bool water_shield = 1; */
        if (message.waterShield !== false)
            writer.tag(1, WireType.Varint).bool(message.waterShield);
        /* bool bloodlust = 2; */
        if (message.bloodlust !== false)
            writer.tag(2, WireType.Varint).bool(message.bloodlust);
        /* bool snapshot_t4_2pc = 6 [json_name = "snapshotT42pc"]; */
        if (message.snapshotT42Pc !== false)
            writer.tag(6, WireType.Varint).bool(message.snapshotT42Pc);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ElementalShaman.Options
 */
export const ElementalShaman_Options = new ElementalShaman_Options$Type();
// @generated message type with reflection information, may provide speed optimized methods
class EnhancementShaman$Type extends MessageType {
    constructor() {
        super("proto.EnhancementShaman", [
            { no: 1, name: "rotation", kind: "message", T: () => EnhancementShaman_Rotation },
            { no: 2, name: "talents", kind: "message", T: () => ShamanTalents },
            { no: 3, name: "options", kind: "message", T: () => EnhancementShaman_Options }
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
                case /* proto.EnhancementShaman.Rotation rotation */ 1:
                    message.rotation = EnhancementShaman_Rotation.internalBinaryRead(reader, reader.uint32(), options, message.rotation);
                    break;
                case /* proto.ShamanTalents talents */ 2:
                    message.talents = ShamanTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* proto.EnhancementShaman.Options options */ 3:
                    message.options = EnhancementShaman_Options.internalBinaryRead(reader, reader.uint32(), options, message.options);
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
        /* proto.EnhancementShaman.Rotation rotation = 1; */
        if (message.rotation)
            EnhancementShaman_Rotation.internalBinaryWrite(message.rotation, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.ShamanTalents talents = 2; */
        if (message.talents)
            ShamanTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.EnhancementShaman.Options options = 3; */
        if (message.options)
            EnhancementShaman_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.EnhancementShaman
 */
export const EnhancementShaman = new EnhancementShaman$Type();
// @generated message type with reflection information, may provide speed optimized methods
class EnhancementShaman_Rotation$Type extends MessageType {
    constructor() {
        super("proto.EnhancementShaman.Rotation", [
            { no: 1, name: "totems", kind: "message", T: () => ShamanTotems },
            { no: 2, name: "primary_shock", kind: "enum", T: () => ["proto.EnhancementShaman.Rotation.PrimaryShock", EnhancementShaman_Rotation_PrimaryShock] },
            { no: 3, name: "weave_flame_shock", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "first_stormstrike_delay", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { primaryShock: 0, weaveFlameShock: false, firstStormstrikeDelay: 0 };
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
                case /* proto.ShamanTotems totems */ 1:
                    message.totems = ShamanTotems.internalBinaryRead(reader, reader.uint32(), options, message.totems);
                    break;
                case /* proto.EnhancementShaman.Rotation.PrimaryShock primary_shock */ 2:
                    message.primaryShock = reader.int32();
                    break;
                case /* bool weave_flame_shock */ 3:
                    message.weaveFlameShock = reader.bool();
                    break;
                case /* double first_stormstrike_delay */ 4:
                    message.firstStormstrikeDelay = reader.double();
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
        /* proto.ShamanTotems totems = 1; */
        if (message.totems)
            ShamanTotems.internalBinaryWrite(message.totems, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.EnhancementShaman.Rotation.PrimaryShock primary_shock = 2; */
        if (message.primaryShock !== 0)
            writer.tag(2, WireType.Varint).int32(message.primaryShock);
        /* bool weave_flame_shock = 3; */
        if (message.weaveFlameShock !== false)
            writer.tag(3, WireType.Varint).bool(message.weaveFlameShock);
        /* double first_stormstrike_delay = 4; */
        if (message.firstStormstrikeDelay !== 0)
            writer.tag(4, WireType.Bit64).double(message.firstStormstrikeDelay);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.EnhancementShaman.Rotation
 */
export const EnhancementShaman_Rotation = new EnhancementShaman_Rotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class EnhancementShaman_Options$Type extends MessageType {
    constructor() {
        super("proto.EnhancementShaman.Options", [
            { no: 1, name: "water_shield", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "bloodlust", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "delay_offhand_swings", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 6, name: "snapshot_t4_2pc", kind: "scalar", jsonName: "snapshotT42pc", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { waterShield: false, bloodlust: false, delayOffhandSwings: false, snapshotT42Pc: false };
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
                case /* bool water_shield */ 1:
                    message.waterShield = reader.bool();
                    break;
                case /* bool bloodlust */ 2:
                    message.bloodlust = reader.bool();
                    break;
                case /* bool delay_offhand_swings */ 5:
                    message.delayOffhandSwings = reader.bool();
                    break;
                case /* bool snapshot_t4_2pc = 6 [json_name = "snapshotT42pc"];*/ 6:
                    message.snapshotT42Pc = reader.bool();
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
        /* bool water_shield = 1; */
        if (message.waterShield !== false)
            writer.tag(1, WireType.Varint).bool(message.waterShield);
        /* bool bloodlust = 2; */
        if (message.bloodlust !== false)
            writer.tag(2, WireType.Varint).bool(message.bloodlust);
        /* bool delay_offhand_swings = 5; */
        if (message.delayOffhandSwings !== false)
            writer.tag(5, WireType.Varint).bool(message.delayOffhandSwings);
        /* bool snapshot_t4_2pc = 6 [json_name = "snapshotT42pc"]; */
        if (message.snapshotT42Pc !== false)
            writer.tag(6, WireType.Varint).bool(message.snapshotT42Pc);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.EnhancementShaman.Options
 */
export const EnhancementShaman_Options = new EnhancementShaman_Options$Type();
