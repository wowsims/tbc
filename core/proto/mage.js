import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
/**
 * @generated from protobuf enum proto.Mage.Rotation.ArcaneRotation.Filler
 */
export var Mage_Rotation_ArcaneRotation_Filler;
(function (Mage_Rotation_ArcaneRotation_Filler) {
    /**
     * @generated from protobuf enum value: Frostbolt = 0;
     */
    Mage_Rotation_ArcaneRotation_Filler[Mage_Rotation_ArcaneRotation_Filler["Frostbolt"] = 0] = "Frostbolt";
    /**
     * @generated from protobuf enum value: ArcaneMissiles = 1;
     */
    Mage_Rotation_ArcaneRotation_Filler[Mage_Rotation_ArcaneRotation_Filler["ArcaneMissiles"] = 1] = "ArcaneMissiles";
    /**
     * @generated from protobuf enum value: Scorch = 2;
     */
    Mage_Rotation_ArcaneRotation_Filler[Mage_Rotation_ArcaneRotation_Filler["Scorch"] = 2] = "Scorch";
    /**
     * @generated from protobuf enum value: Fireball = 3;
     */
    Mage_Rotation_ArcaneRotation_Filler[Mage_Rotation_ArcaneRotation_Filler["Fireball"] = 3] = "Fireball";
    /**
     * @generated from protobuf enum value: ArcaneMissilesFrostbolt = 4;
     */
    Mage_Rotation_ArcaneRotation_Filler[Mage_Rotation_ArcaneRotation_Filler["ArcaneMissilesFrostbolt"] = 4] = "ArcaneMissilesFrostbolt";
    /**
     * @generated from protobuf enum value: ArcaneMissilesScorch = 5;
     */
    Mage_Rotation_ArcaneRotation_Filler[Mage_Rotation_ArcaneRotation_Filler["ArcaneMissilesScorch"] = 5] = "ArcaneMissilesScorch";
    /**
     * @generated from protobuf enum value: ScorchTwoFireball = 6;
     */
    Mage_Rotation_ArcaneRotation_Filler[Mage_Rotation_ArcaneRotation_Filler["ScorchTwoFireball"] = 6] = "ScorchTwoFireball";
})(Mage_Rotation_ArcaneRotation_Filler || (Mage_Rotation_ArcaneRotation_Filler = {}));
/**
 * @generated from protobuf enum proto.Mage.Rotation.FireRotation.PrimarySpell
 */
export var Mage_Rotation_FireRotation_PrimarySpell;
(function (Mage_Rotation_FireRotation_PrimarySpell) {
    /**
     * @generated from protobuf enum value: Fireball = 0;
     */
    Mage_Rotation_FireRotation_PrimarySpell[Mage_Rotation_FireRotation_PrimarySpell["Fireball"] = 0] = "Fireball";
    /**
     * @generated from protobuf enum value: Scorch = 1;
     */
    Mage_Rotation_FireRotation_PrimarySpell[Mage_Rotation_FireRotation_PrimarySpell["Scorch"] = 1] = "Scorch";
})(Mage_Rotation_FireRotation_PrimarySpell || (Mage_Rotation_FireRotation_PrimarySpell = {}));
/**
 * @generated from protobuf enum proto.Mage.Rotation.AoeRotation.Rotation
 */
export var Mage_Rotation_AoeRotation_Rotation;
(function (Mage_Rotation_AoeRotation_Rotation) {
    /**
     * @generated from protobuf enum value: ArcaneExplosion = 0;
     */
    Mage_Rotation_AoeRotation_Rotation[Mage_Rotation_AoeRotation_Rotation["ArcaneExplosion"] = 0] = "ArcaneExplosion";
    /**
     * @generated from protobuf enum value: Flamestrike = 1;
     */
    Mage_Rotation_AoeRotation_Rotation[Mage_Rotation_AoeRotation_Rotation["Flamestrike"] = 1] = "Flamestrike";
    /**
     * @generated from protobuf enum value: Blizzard = 2;
     */
    Mage_Rotation_AoeRotation_Rotation[Mage_Rotation_AoeRotation_Rotation["Blizzard"] = 2] = "Blizzard";
})(Mage_Rotation_AoeRotation_Rotation || (Mage_Rotation_AoeRotation_Rotation = {}));
/**
 * Just used for controlling which options are displayed in the UI. Is not
 * used by the sim.
 *
 * @generated from protobuf enum proto.Mage.Rotation.Type
 */
export var Mage_Rotation_Type;
(function (Mage_Rotation_Type) {
    /**
     * @generated from protobuf enum value: Arcane = 0;
     */
    Mage_Rotation_Type[Mage_Rotation_Type["Arcane"] = 0] = "Arcane";
    /**
     * @generated from protobuf enum value: Fire = 1;
     */
    Mage_Rotation_Type[Mage_Rotation_Type["Fire"] = 1] = "Fire";
    /**
     * @generated from protobuf enum value: Frost = 2;
     */
    Mage_Rotation_Type[Mage_Rotation_Type["Frost"] = 2] = "Frost";
})(Mage_Rotation_Type || (Mage_Rotation_Type = {}));
/**
 * @generated from protobuf enum proto.Mage.Options.ArmorType
 */
export var Mage_Options_ArmorType;
(function (Mage_Options_ArmorType) {
    /**
     * @generated from protobuf enum value: NoArmor = 0;
     */
    Mage_Options_ArmorType[Mage_Options_ArmorType["NoArmor"] = 0] = "NoArmor";
    /**
     * @generated from protobuf enum value: MageArmor = 1;
     */
    Mage_Options_ArmorType[Mage_Options_ArmorType["MageArmor"] = 1] = "MageArmor";
    /**
     * @generated from protobuf enum value: MoltenArmor = 2;
     */
    Mage_Options_ArmorType[Mage_Options_ArmorType["MoltenArmor"] = 2] = "MoltenArmor";
})(Mage_Options_ArmorType || (Mage_Options_ArmorType = {}));
// @generated message type with reflection information, may provide speed optimized methods
class MageTalents$Type extends MessageType {
    constructor() {
        super("proto.MageTalents", [
            { no: 1, name: "arcane_subtlety", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "arcane_focus", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "wand_specialization", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 48, name: "magic_absorption", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "arcane_concentration", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "arcane_impact", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "arcane_meditation", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 7, name: "presence_of_mind", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 8, name: "arcane_mind", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 9, name: "arcane_instability", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 10, name: "arcane_potency", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 11, name: "empowered_arcane_missiles", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 12, name: "arcane_power", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 13, name: "spell_power", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 14, name: "mind_mastery", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 15, name: "improved_fireball", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 16, name: "ignite", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 17, name: "improved_fire_blast", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 18, name: "incineration", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 19, name: "improved_flamestrike", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 20, name: "pyroblast", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 47, name: "burning_soul", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 21, name: "improved_scorch", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 22, name: "master_of_elements", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 23, name: "playing_with_fire", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 24, name: "critical_mass", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 25, name: "blast_wave", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 26, name: "fire_power", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 27, name: "pyromaniac", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 28, name: "combustion", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 29, name: "molten_fury", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 30, name: "empowered_fireball", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 31, name: "dragons_breath", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 32, name: "improved_frostbolt", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 33, name: "elemental_precision", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 34, name: "ice_shards", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 35, name: "improved_frost_nova", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 36, name: "piercing_ice", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 37, name: "icy_veins", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 38, name: "frost_channeling", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 39, name: "shatter", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 40, name: "cold_snap", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 41, name: "improved_cone_of_cold", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 42, name: "ice_floes", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 43, name: "winters_chill", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 44, name: "arctic_winds", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 45, name: "empowered_frostbolt", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 46, name: "summon_water_elemental", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { arcaneSubtlety: 0, arcaneFocus: 0, wandSpecialization: 0, magicAbsorption: 0, arcaneConcentration: 0, arcaneImpact: 0, arcaneMeditation: 0, presenceOfMind: false, arcaneMind: 0, arcaneInstability: 0, arcanePotency: 0, empoweredArcaneMissiles: 0, arcanePower: false, spellPower: 0, mindMastery: 0, improvedFireball: 0, ignite: 0, improvedFireBlast: 0, incineration: 0, improvedFlamestrike: 0, pyroblast: false, burningSoul: 0, improvedScorch: 0, masterOfElements: 0, playingWithFire: 0, criticalMass: 0, blastWave: false, firePower: 0, pyromaniac: 0, combustion: false, moltenFury: 0, empoweredFireball: 0, dragonsBreath: false, improvedFrostbolt: 0, elementalPrecision: 0, iceShards: 0, improvedFrostNova: 0, piercingIce: 0, icyVeins: false, frostChanneling: 0, shatter: 0, coldSnap: false, improvedConeOfCold: 0, iceFloes: 0, wintersChill: 0, arcticWinds: 0, empoweredFrostbolt: 0, summonWaterElemental: false };
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
                case /* int32 arcane_subtlety */ 1:
                    message.arcaneSubtlety = reader.int32();
                    break;
                case /* int32 arcane_focus */ 2:
                    message.arcaneFocus = reader.int32();
                    break;
                case /* int32 wand_specialization */ 3:
                    message.wandSpecialization = reader.int32();
                    break;
                case /* int32 magic_absorption */ 48:
                    message.magicAbsorption = reader.int32();
                    break;
                case /* int32 arcane_concentration */ 4:
                    message.arcaneConcentration = reader.int32();
                    break;
                case /* int32 arcane_impact */ 5:
                    message.arcaneImpact = reader.int32();
                    break;
                case /* int32 arcane_meditation */ 6:
                    message.arcaneMeditation = reader.int32();
                    break;
                case /* bool presence_of_mind */ 7:
                    message.presenceOfMind = reader.bool();
                    break;
                case /* int32 arcane_mind */ 8:
                    message.arcaneMind = reader.int32();
                    break;
                case /* int32 arcane_instability */ 9:
                    message.arcaneInstability = reader.int32();
                    break;
                case /* int32 arcane_potency */ 10:
                    message.arcanePotency = reader.int32();
                    break;
                case /* int32 empowered_arcane_missiles */ 11:
                    message.empoweredArcaneMissiles = reader.int32();
                    break;
                case /* bool arcane_power */ 12:
                    message.arcanePower = reader.bool();
                    break;
                case /* int32 spell_power */ 13:
                    message.spellPower = reader.int32();
                    break;
                case /* int32 mind_mastery */ 14:
                    message.mindMastery = reader.int32();
                    break;
                case /* int32 improved_fireball */ 15:
                    message.improvedFireball = reader.int32();
                    break;
                case /* int32 ignite */ 16:
                    message.ignite = reader.int32();
                    break;
                case /* int32 improved_fire_blast */ 17:
                    message.improvedFireBlast = reader.int32();
                    break;
                case /* int32 incineration */ 18:
                    message.incineration = reader.int32();
                    break;
                case /* int32 improved_flamestrike */ 19:
                    message.improvedFlamestrike = reader.int32();
                    break;
                case /* bool pyroblast */ 20:
                    message.pyroblast = reader.bool();
                    break;
                case /* int32 burning_soul */ 47:
                    message.burningSoul = reader.int32();
                    break;
                case /* int32 improved_scorch */ 21:
                    message.improvedScorch = reader.int32();
                    break;
                case /* int32 master_of_elements */ 22:
                    message.masterOfElements = reader.int32();
                    break;
                case /* int32 playing_with_fire */ 23:
                    message.playingWithFire = reader.int32();
                    break;
                case /* int32 critical_mass */ 24:
                    message.criticalMass = reader.int32();
                    break;
                case /* bool blast_wave */ 25:
                    message.blastWave = reader.bool();
                    break;
                case /* int32 fire_power */ 26:
                    message.firePower = reader.int32();
                    break;
                case /* int32 pyromaniac */ 27:
                    message.pyromaniac = reader.int32();
                    break;
                case /* bool combustion */ 28:
                    message.combustion = reader.bool();
                    break;
                case /* int32 molten_fury */ 29:
                    message.moltenFury = reader.int32();
                    break;
                case /* int32 empowered_fireball */ 30:
                    message.empoweredFireball = reader.int32();
                    break;
                case /* bool dragons_breath */ 31:
                    message.dragonsBreath = reader.bool();
                    break;
                case /* int32 improved_frostbolt */ 32:
                    message.improvedFrostbolt = reader.int32();
                    break;
                case /* int32 elemental_precision */ 33:
                    message.elementalPrecision = reader.int32();
                    break;
                case /* int32 ice_shards */ 34:
                    message.iceShards = reader.int32();
                    break;
                case /* int32 improved_frost_nova */ 35:
                    message.improvedFrostNova = reader.int32();
                    break;
                case /* int32 piercing_ice */ 36:
                    message.piercingIce = reader.int32();
                    break;
                case /* bool icy_veins */ 37:
                    message.icyVeins = reader.bool();
                    break;
                case /* int32 frost_channeling */ 38:
                    message.frostChanneling = reader.int32();
                    break;
                case /* int32 shatter */ 39:
                    message.shatter = reader.int32();
                    break;
                case /* bool cold_snap */ 40:
                    message.coldSnap = reader.bool();
                    break;
                case /* int32 improved_cone_of_cold */ 41:
                    message.improvedConeOfCold = reader.int32();
                    break;
                case /* int32 ice_floes */ 42:
                    message.iceFloes = reader.int32();
                    break;
                case /* int32 winters_chill */ 43:
                    message.wintersChill = reader.int32();
                    break;
                case /* int32 arctic_winds */ 44:
                    message.arcticWinds = reader.int32();
                    break;
                case /* int32 empowered_frostbolt */ 45:
                    message.empoweredFrostbolt = reader.int32();
                    break;
                case /* bool summon_water_elemental */ 46:
                    message.summonWaterElemental = reader.bool();
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
        /* int32 arcane_subtlety = 1; */
        if (message.arcaneSubtlety !== 0)
            writer.tag(1, WireType.Varint).int32(message.arcaneSubtlety);
        /* int32 arcane_focus = 2; */
        if (message.arcaneFocus !== 0)
            writer.tag(2, WireType.Varint).int32(message.arcaneFocus);
        /* int32 wand_specialization = 3; */
        if (message.wandSpecialization !== 0)
            writer.tag(3, WireType.Varint).int32(message.wandSpecialization);
        /* int32 magic_absorption = 48; */
        if (message.magicAbsorption !== 0)
            writer.tag(48, WireType.Varint).int32(message.magicAbsorption);
        /* int32 arcane_concentration = 4; */
        if (message.arcaneConcentration !== 0)
            writer.tag(4, WireType.Varint).int32(message.arcaneConcentration);
        /* int32 arcane_impact = 5; */
        if (message.arcaneImpact !== 0)
            writer.tag(5, WireType.Varint).int32(message.arcaneImpact);
        /* int32 arcane_meditation = 6; */
        if (message.arcaneMeditation !== 0)
            writer.tag(6, WireType.Varint).int32(message.arcaneMeditation);
        /* bool presence_of_mind = 7; */
        if (message.presenceOfMind !== false)
            writer.tag(7, WireType.Varint).bool(message.presenceOfMind);
        /* int32 arcane_mind = 8; */
        if (message.arcaneMind !== 0)
            writer.tag(8, WireType.Varint).int32(message.arcaneMind);
        /* int32 arcane_instability = 9; */
        if (message.arcaneInstability !== 0)
            writer.tag(9, WireType.Varint).int32(message.arcaneInstability);
        /* int32 arcane_potency = 10; */
        if (message.arcanePotency !== 0)
            writer.tag(10, WireType.Varint).int32(message.arcanePotency);
        /* int32 empowered_arcane_missiles = 11; */
        if (message.empoweredArcaneMissiles !== 0)
            writer.tag(11, WireType.Varint).int32(message.empoweredArcaneMissiles);
        /* bool arcane_power = 12; */
        if (message.arcanePower !== false)
            writer.tag(12, WireType.Varint).bool(message.arcanePower);
        /* int32 spell_power = 13; */
        if (message.spellPower !== 0)
            writer.tag(13, WireType.Varint).int32(message.spellPower);
        /* int32 mind_mastery = 14; */
        if (message.mindMastery !== 0)
            writer.tag(14, WireType.Varint).int32(message.mindMastery);
        /* int32 improved_fireball = 15; */
        if (message.improvedFireball !== 0)
            writer.tag(15, WireType.Varint).int32(message.improvedFireball);
        /* int32 ignite = 16; */
        if (message.ignite !== 0)
            writer.tag(16, WireType.Varint).int32(message.ignite);
        /* int32 improved_fire_blast = 17; */
        if (message.improvedFireBlast !== 0)
            writer.tag(17, WireType.Varint).int32(message.improvedFireBlast);
        /* int32 incineration = 18; */
        if (message.incineration !== 0)
            writer.tag(18, WireType.Varint).int32(message.incineration);
        /* int32 improved_flamestrike = 19; */
        if (message.improvedFlamestrike !== 0)
            writer.tag(19, WireType.Varint).int32(message.improvedFlamestrike);
        /* bool pyroblast = 20; */
        if (message.pyroblast !== false)
            writer.tag(20, WireType.Varint).bool(message.pyroblast);
        /* int32 burning_soul = 47; */
        if (message.burningSoul !== 0)
            writer.tag(47, WireType.Varint).int32(message.burningSoul);
        /* int32 improved_scorch = 21; */
        if (message.improvedScorch !== 0)
            writer.tag(21, WireType.Varint).int32(message.improvedScorch);
        /* int32 master_of_elements = 22; */
        if (message.masterOfElements !== 0)
            writer.tag(22, WireType.Varint).int32(message.masterOfElements);
        /* int32 playing_with_fire = 23; */
        if (message.playingWithFire !== 0)
            writer.tag(23, WireType.Varint).int32(message.playingWithFire);
        /* int32 critical_mass = 24; */
        if (message.criticalMass !== 0)
            writer.tag(24, WireType.Varint).int32(message.criticalMass);
        /* bool blast_wave = 25; */
        if (message.blastWave !== false)
            writer.tag(25, WireType.Varint).bool(message.blastWave);
        /* int32 fire_power = 26; */
        if (message.firePower !== 0)
            writer.tag(26, WireType.Varint).int32(message.firePower);
        /* int32 pyromaniac = 27; */
        if (message.pyromaniac !== 0)
            writer.tag(27, WireType.Varint).int32(message.pyromaniac);
        /* bool combustion = 28; */
        if (message.combustion !== false)
            writer.tag(28, WireType.Varint).bool(message.combustion);
        /* int32 molten_fury = 29; */
        if (message.moltenFury !== 0)
            writer.tag(29, WireType.Varint).int32(message.moltenFury);
        /* int32 empowered_fireball = 30; */
        if (message.empoweredFireball !== 0)
            writer.tag(30, WireType.Varint).int32(message.empoweredFireball);
        /* bool dragons_breath = 31; */
        if (message.dragonsBreath !== false)
            writer.tag(31, WireType.Varint).bool(message.dragonsBreath);
        /* int32 improved_frostbolt = 32; */
        if (message.improvedFrostbolt !== 0)
            writer.tag(32, WireType.Varint).int32(message.improvedFrostbolt);
        /* int32 elemental_precision = 33; */
        if (message.elementalPrecision !== 0)
            writer.tag(33, WireType.Varint).int32(message.elementalPrecision);
        /* int32 ice_shards = 34; */
        if (message.iceShards !== 0)
            writer.tag(34, WireType.Varint).int32(message.iceShards);
        /* int32 improved_frost_nova = 35; */
        if (message.improvedFrostNova !== 0)
            writer.tag(35, WireType.Varint).int32(message.improvedFrostNova);
        /* int32 piercing_ice = 36; */
        if (message.piercingIce !== 0)
            writer.tag(36, WireType.Varint).int32(message.piercingIce);
        /* bool icy_veins = 37; */
        if (message.icyVeins !== false)
            writer.tag(37, WireType.Varint).bool(message.icyVeins);
        /* int32 frost_channeling = 38; */
        if (message.frostChanneling !== 0)
            writer.tag(38, WireType.Varint).int32(message.frostChanneling);
        /* int32 shatter = 39; */
        if (message.shatter !== 0)
            writer.tag(39, WireType.Varint).int32(message.shatter);
        /* bool cold_snap = 40; */
        if (message.coldSnap !== false)
            writer.tag(40, WireType.Varint).bool(message.coldSnap);
        /* int32 improved_cone_of_cold = 41; */
        if (message.improvedConeOfCold !== 0)
            writer.tag(41, WireType.Varint).int32(message.improvedConeOfCold);
        /* int32 ice_floes = 42; */
        if (message.iceFloes !== 0)
            writer.tag(42, WireType.Varint).int32(message.iceFloes);
        /* int32 winters_chill = 43; */
        if (message.wintersChill !== 0)
            writer.tag(43, WireType.Varint).int32(message.wintersChill);
        /* int32 arctic_winds = 44; */
        if (message.arcticWinds !== 0)
            writer.tag(44, WireType.Varint).int32(message.arcticWinds);
        /* int32 empowered_frostbolt = 45; */
        if (message.empoweredFrostbolt !== 0)
            writer.tag(45, WireType.Varint).int32(message.empoweredFrostbolt);
        /* bool summon_water_elemental = 46; */
        if (message.summonWaterElemental !== false)
            writer.tag(46, WireType.Varint).bool(message.summonWaterElemental);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.MageTalents
 */
export const MageTalents = new MageTalents$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Mage$Type extends MessageType {
    constructor() {
        super("proto.Mage", [
            { no: 1, name: "rotation", kind: "message", T: () => Mage_Rotation },
            { no: 2, name: "talents", kind: "message", T: () => MageTalents },
            { no: 3, name: "options", kind: "message", T: () => Mage_Options }
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
                case /* proto.Mage.Rotation rotation */ 1:
                    message.rotation = Mage_Rotation.internalBinaryRead(reader, reader.uint32(), options, message.rotation);
                    break;
                case /* proto.MageTalents talents */ 2:
                    message.talents = MageTalents.internalBinaryRead(reader, reader.uint32(), options, message.talents);
                    break;
                case /* proto.Mage.Options options */ 3:
                    message.options = Mage_Options.internalBinaryRead(reader, reader.uint32(), options, message.options);
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
        /* proto.Mage.Rotation rotation = 1; */
        if (message.rotation)
            Mage_Rotation.internalBinaryWrite(message.rotation, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* proto.MageTalents talents = 2; */
        if (message.talents)
            MageTalents.internalBinaryWrite(message.talents, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.Mage.Options options = 3; */
        if (message.options)
            Mage_Options.internalBinaryWrite(message.options, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Mage
 */
export const Mage = new Mage$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Mage_Rotation$Type extends MessageType {
    constructor() {
        super("proto.Mage.Rotation", [
            { no: 1, name: "type", kind: "enum", T: () => ["proto.Mage.Rotation.Type", Mage_Rotation_Type] },
            { no: 2, name: "arcane", kind: "message", T: () => Mage_Rotation_ArcaneRotation },
            { no: 3, name: "fire", kind: "message", T: () => Mage_Rotation_FireRotation },
            { no: 4, name: "frost", kind: "message", T: () => Mage_Rotation_FrostRotation },
            { no: 5, name: "aoe", kind: "message", T: () => Mage_Rotation_AoeRotation },
            { no: 6, name: "multi_target_rotation", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { type: 0, multiTargetRotation: false };
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
                case /* proto.Mage.Rotation.Type type */ 1:
                    message.type = reader.int32();
                    break;
                case /* proto.Mage.Rotation.ArcaneRotation arcane */ 2:
                    message.arcane = Mage_Rotation_ArcaneRotation.internalBinaryRead(reader, reader.uint32(), options, message.arcane);
                    break;
                case /* proto.Mage.Rotation.FireRotation fire */ 3:
                    message.fire = Mage_Rotation_FireRotation.internalBinaryRead(reader, reader.uint32(), options, message.fire);
                    break;
                case /* proto.Mage.Rotation.FrostRotation frost */ 4:
                    message.frost = Mage_Rotation_FrostRotation.internalBinaryRead(reader, reader.uint32(), options, message.frost);
                    break;
                case /* proto.Mage.Rotation.AoeRotation aoe */ 5:
                    message.aoe = Mage_Rotation_AoeRotation.internalBinaryRead(reader, reader.uint32(), options, message.aoe);
                    break;
                case /* bool multi_target_rotation */ 6:
                    message.multiTargetRotation = reader.bool();
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
        /* proto.Mage.Rotation.Type type = 1; */
        if (message.type !== 0)
            writer.tag(1, WireType.Varint).int32(message.type);
        /* proto.Mage.Rotation.ArcaneRotation arcane = 2; */
        if (message.arcane)
            Mage_Rotation_ArcaneRotation.internalBinaryWrite(message.arcane, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* proto.Mage.Rotation.FireRotation fire = 3; */
        if (message.fire)
            Mage_Rotation_FireRotation.internalBinaryWrite(message.fire, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* proto.Mage.Rotation.FrostRotation frost = 4; */
        if (message.frost)
            Mage_Rotation_FrostRotation.internalBinaryWrite(message.frost, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* proto.Mage.Rotation.AoeRotation aoe = 5; */
        if (message.aoe)
            Mage_Rotation_AoeRotation.internalBinaryWrite(message.aoe, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* bool multi_target_rotation = 6; */
        if (message.multiTargetRotation !== false)
            writer.tag(6, WireType.Varint).bool(message.multiTargetRotation);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Mage.Rotation
 */
export const Mage_Rotation = new Mage_Rotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Mage_Rotation_ArcaneRotation$Type extends MessageType {
    constructor() {
        super("proto.Mage.Rotation.ArcaneRotation", [
            { no: 1, name: "filler", kind: "enum", T: () => ["proto.Mage.Rotation.ArcaneRotation.Filler", Mage_Rotation_ArcaneRotation_Filler] },
            { no: 2, name: "arcane_blasts_between_fillers", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "start_regen_rotation_percent", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 4, name: "stop_regen_rotation_percent", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 5, name: "disable_dps_cooldowns_during_regen", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { filler: 0, arcaneBlastsBetweenFillers: 0, startRegenRotationPercent: 0, stopRegenRotationPercent: 0, disableDpsCooldownsDuringRegen: false };
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
                case /* proto.Mage.Rotation.ArcaneRotation.Filler filler */ 1:
                    message.filler = reader.int32();
                    break;
                case /* int32 arcane_blasts_between_fillers */ 2:
                    message.arcaneBlastsBetweenFillers = reader.int32();
                    break;
                case /* double start_regen_rotation_percent */ 3:
                    message.startRegenRotationPercent = reader.double();
                    break;
                case /* double stop_regen_rotation_percent */ 4:
                    message.stopRegenRotationPercent = reader.double();
                    break;
                case /* bool disable_dps_cooldowns_during_regen */ 5:
                    message.disableDpsCooldownsDuringRegen = reader.bool();
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
        /* proto.Mage.Rotation.ArcaneRotation.Filler filler = 1; */
        if (message.filler !== 0)
            writer.tag(1, WireType.Varint).int32(message.filler);
        /* int32 arcane_blasts_between_fillers = 2; */
        if (message.arcaneBlastsBetweenFillers !== 0)
            writer.tag(2, WireType.Varint).int32(message.arcaneBlastsBetweenFillers);
        /* double start_regen_rotation_percent = 3; */
        if (message.startRegenRotationPercent !== 0)
            writer.tag(3, WireType.Bit64).double(message.startRegenRotationPercent);
        /* double stop_regen_rotation_percent = 4; */
        if (message.stopRegenRotationPercent !== 0)
            writer.tag(4, WireType.Bit64).double(message.stopRegenRotationPercent);
        /* bool disable_dps_cooldowns_during_regen = 5; */
        if (message.disableDpsCooldownsDuringRegen !== false)
            writer.tag(5, WireType.Varint).bool(message.disableDpsCooldownsDuringRegen);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Mage.Rotation.ArcaneRotation
 */
export const Mage_Rotation_ArcaneRotation = new Mage_Rotation_ArcaneRotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Mage_Rotation_FireRotation$Type extends MessageType {
    constructor() {
        super("proto.Mage.Rotation.FireRotation", [
            { no: 1, name: "primary_spell", kind: "enum", T: () => ["proto.Mage.Rotation.FireRotation.PrimarySpell", Mage_Rotation_FireRotation_PrimarySpell] },
            { no: 2, name: "maintain_improved_scorch", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "weave_fire_blast", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { primarySpell: 0, maintainImprovedScorch: false, weaveFireBlast: false };
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
                case /* proto.Mage.Rotation.FireRotation.PrimarySpell primary_spell */ 1:
                    message.primarySpell = reader.int32();
                    break;
                case /* bool maintain_improved_scorch */ 2:
                    message.maintainImprovedScorch = reader.bool();
                    break;
                case /* bool weave_fire_blast */ 3:
                    message.weaveFireBlast = reader.bool();
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
        /* proto.Mage.Rotation.FireRotation.PrimarySpell primary_spell = 1; */
        if (message.primarySpell !== 0)
            writer.tag(1, WireType.Varint).int32(message.primarySpell);
        /* bool maintain_improved_scorch = 2; */
        if (message.maintainImprovedScorch !== false)
            writer.tag(2, WireType.Varint).bool(message.maintainImprovedScorch);
        /* bool weave_fire_blast = 3; */
        if (message.weaveFireBlast !== false)
            writer.tag(3, WireType.Varint).bool(message.weaveFireBlast);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Mage.Rotation.FireRotation
 */
export const Mage_Rotation_FireRotation = new Mage_Rotation_FireRotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Mage_Rotation_FrostRotation$Type extends MessageType {
    constructor() {
        super("proto.Mage.Rotation.FrostRotation", [
            { no: 3, name: "water_elemental_disobey_chance", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { waterElementalDisobeyChance: 0 };
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
                case /* double water_elemental_disobey_chance */ 3:
                    message.waterElementalDisobeyChance = reader.double();
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
        /* double water_elemental_disobey_chance = 3; */
        if (message.waterElementalDisobeyChance !== 0)
            writer.tag(3, WireType.Bit64).double(message.waterElementalDisobeyChance);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Mage.Rotation.FrostRotation
 */
export const Mage_Rotation_FrostRotation = new Mage_Rotation_FrostRotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Mage_Rotation_AoeRotation$Type extends MessageType {
    constructor() {
        super("proto.Mage.Rotation.AoeRotation", [
            { no: 1, name: "rotation", kind: "enum", T: () => ["proto.Mage.Rotation.AoeRotation.Rotation", Mage_Rotation_AoeRotation_Rotation] }
        ]);
    }
    create(value) {
        const message = { rotation: 0 };
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
                case /* proto.Mage.Rotation.AoeRotation.Rotation rotation */ 1:
                    message.rotation = reader.int32();
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
        /* proto.Mage.Rotation.AoeRotation.Rotation rotation = 1; */
        if (message.rotation !== 0)
            writer.tag(1, WireType.Varint).int32(message.rotation);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Mage.Rotation.AoeRotation
 */
export const Mage_Rotation_AoeRotation = new Mage_Rotation_AoeRotation$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Mage_Options$Type extends MessageType {
    constructor() {
        super("proto.Mage.Options", [
            { no: 1, name: "armor", kind: "enum", T: () => ["proto.Mage.Options.ArmorType", Mage_Options_ArmorType] },
            { no: 2, name: "evocation_ticks", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value) {
        const message = { armor: 0, evocationTicks: 0 };
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
                case /* proto.Mage.Options.ArmorType armor */ 1:
                    message.armor = reader.int32();
                    break;
                case /* int32 evocation_ticks */ 2:
                    message.evocationTicks = reader.int32();
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
        /* proto.Mage.Options.ArmorType armor = 1; */
        if (message.armor !== 0)
            writer.tag(1, WireType.Varint).int32(message.armor);
        /* int32 evocation_ticks = 2; */
        if (message.evocationTicks !== 0)
            writer.tag(2, WireType.Varint).int32(message.evocationTicks);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Mage.Options
 */
export const Mage_Options = new Mage_Options$Type();
