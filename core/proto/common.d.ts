import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message proto.Buffs
 */
export interface Buffs {
    /**
     * Raid buffs
     *
     * @generated from protobuf field: bool arcane_brilliance = 1;
     */
    arcaneBrilliance: boolean;
    /**
     * @generated from protobuf field: bool blessing_of_kings = 2;
     */
    blessingOfKings: boolean;
    /**
     * @generated from protobuf field: proto.TristateEffect blessing_of_wisdom = 3;
     */
    blessingOfWisdom: TristateEffect;
    /**
     * @generated from protobuf field: proto.TristateEffect divine_spirit = 4;
     */
    divineSpirit: TristateEffect;
    /**
     * @generated from protobuf field: proto.TristateEffect gift_of_the_wild = 5;
     */
    giftOfTheWild: TristateEffect;
    /**
     * Party Buffs
     *
     * @generated from protobuf field: int32 bloodlust = 6;
     */
    bloodlust: number;
    /**
     * @generated from protobuf field: proto.TristateEffect moonkin_aura = 7;
     */
    moonkinAura: TristateEffect;
    /**
     * @generated from protobuf field: bool draenei_racial_melee = 8;
     */
    draeneiRacialMelee: boolean;
    /**
     * @generated from protobuf field: bool draenei_racial_caster = 9;
     */
    draeneiRacialCaster: boolean;
    /**
     * @generated from protobuf field: int32 shadow_priest_dps = 10;
     */
    shadowPriestDps: number;
    /**
     * Drums
     *
     * @generated from protobuf field: proto.Drums drums = 11;
     */
    drums: Drums;
    /**
     * Item Buffs
     *
     * @generated from protobuf field: int32 atiesh_mage = 12;
     */
    atieshMage: number;
    /**
     * @generated from protobuf field: int32 atiesh_warlock = 13;
     */
    atieshWarlock: number;
    /**
     * @generated from protobuf field: bool braided_eternium_chain = 14;
     */
    braidedEterniumChain: boolean;
    /**
     * @generated from protobuf field: bool eye_of_the_night = 15;
     */
    eyeOfTheNight: boolean;
    /**
     * @generated from protobuf field: bool chain_of_the_twilight_owl = 16;
     */
    chainOfTheTwilightOwl: boolean;
    /**
     * @generated from protobuf field: bool jade_pendant_of_blasting = 17;
     */
    jadePendantOfBlasting: boolean;
    /**
     * Totems
     *
     * @generated from protobuf field: proto.TristateEffect mana_spring_totem = 18;
     */
    manaSpringTotem: TristateEffect;
    /**
     * @generated from protobuf field: bool mana_tide_totem = 19;
     */
    manaTideTotem: boolean;
    /**
     * @generated from protobuf field: int32 totem_of_wrath = 20;
     */
    totemOfWrath: number;
    /**
     * @generated from protobuf field: proto.TristateEffect wrath_of_air_totem = 21;
     */
    wrathOfAirTotem: TristateEffect;
}
/**
 * @generated from protobuf message proto.Consumes
 */
export interface Consumes {
    /**
     * @generated from protobuf field: bool flask_of_blinding_light = 1;
     */
    flaskOfBlindingLight: boolean;
    /**
     * @generated from protobuf field: bool flask_of_mighty_restoration = 2;
     */
    flaskOfMightyRestoration: boolean;
    /**
     * @generated from protobuf field: bool flask_of_pure_death = 3;
     */
    flaskOfPureDeath: boolean;
    /**
     * @generated from protobuf field: bool flask_of_supreme_power = 4;
     */
    flaskOfSupremePower: boolean;
    /**
     * @generated from protobuf field: bool adepts_elixir = 5;
     */
    adeptsElixir: boolean;
    /**
     * @generated from protobuf field: bool elixir_of_major_fire_power = 6;
     */
    elixirOfMajorFirePower: boolean;
    /**
     * @generated from protobuf field: bool elixir_of_major_frost_power = 7;
     */
    elixirOfMajorFrostPower: boolean;
    /**
     * @generated from protobuf field: bool elixir_of_major_shadow_power = 8;
     */
    elixirOfMajorShadowPower: boolean;
    /**
     * @generated from protobuf field: bool elixir_of_draenic_wisdom = 9;
     */
    elixirOfDraenicWisdom: boolean;
    /**
     * @generated from protobuf field: bool elixir_of_major_mageblood = 10;
     */
    elixirOfMajorMageblood: boolean;
    /**
     * @generated from protobuf field: bool brilliant_wizard_oil = 11;
     */
    brilliantWizardOil: boolean;
    /**
     * @generated from protobuf field: bool superior_wizard_oil = 12;
     */
    superiorWizardOil: boolean;
    /**
     * @generated from protobuf field: bool blackened_basilisk = 13;
     */
    blackenedBasilisk: boolean;
    /**
     * @generated from protobuf field: bool skullfish_soup = 14;
     */
    skullfishSoup: boolean;
    /**
     * @generated from protobuf field: proto.Potions default_potion = 15;
     */
    defaultPotion: Potions;
    /**
     * @generated from protobuf field: proto.Potions starting_potion = 16;
     */
    startingPotion: Potions;
    /**
     * @generated from protobuf field: int32 num_starting_potions = 17;
     */
    numStartingPotions: number;
    /**
     * @generated from protobuf field: bool dark_rune = 18;
     */
    darkRune: boolean;
    /**
     * @generated from protobuf field: proto.Drums drums = 19;
     */
    drums: Drums;
}
/**
 * @generated from protobuf message proto.Debuffs
 */
export interface Debuffs {
    /**
     * @generated from protobuf field: bool judgement_of_wisdom = 1;
     */
    judgementOfWisdom: boolean;
    /**
     * @generated from protobuf field: bool improved_seal_of_the_crusader = 2;
     */
    improvedSealOfTheCrusader: boolean;
    /**
     * @generated from protobuf field: bool misery = 3;
     */
    misery: boolean;
}
/**
 * @generated from protobuf message proto.Target
 */
export interface Target {
    /**
     * @generated from protobuf field: int32 armor = 1;
     */
    armor: number;
    /**
     * @generated from protobuf field: proto.Debuffs debuffs = 2;
     */
    debuffs?: Debuffs;
}
/**
 * @generated from protobuf message proto.Encounter
 */
export interface Encounter {
    /**
     * @generated from protobuf field: double duration = 1;
     */
    duration: number;
    /**
     * @generated from protobuf field: repeated proto.Target targets = 2;
     */
    targets: Target[];
}
/**
 * @generated from protobuf message proto.ItemSpec
 */
export interface ItemSpec {
    /**
     * @generated from protobuf field: int32 id = 2;
     */
    id: number;
    /**
     * @generated from protobuf field: int32 enchant = 3;
     */
    enchant: number;
    /**
     * @generated from protobuf field: repeated int32 gems = 4;
     */
    gems: number[];
}
/**
 * @generated from protobuf message proto.EquipmentSpec
 */
export interface EquipmentSpec {
    /**
     * @generated from protobuf field: repeated proto.ItemSpec items = 1;
     */
    items: ItemSpec[];
}
/**
 * @generated from protobuf message proto.Item
 */
export interface Item {
    /**
     * @generated from protobuf field: int32 id = 1;
     */
    id: number;
    /**
     * @generated from protobuf field: string name = 2;
     */
    name: string;
    /**
     * @generated from protobuf field: repeated proto.ItemCategory categories = 14;
     */
    categories: ItemCategory[];
    /**
     * Classes that are allowed to use the item. Empty indicates no special class restrictions.
     *
     * @generated from protobuf field: repeated proto.Class class_allowlist = 15;
     */
    classAllowlist: Class[];
    /**
     * @generated from protobuf field: proto.ItemType type = 3;
     */
    type: ItemType;
    /**
     * @generated from protobuf field: proto.ArmorType armor_type = 4;
     */
    armorType: ArmorType;
    /**
     * @generated from protobuf field: proto.WeaponType weapon_type = 5;
     */
    weaponType: WeaponType;
    /**
     * @generated from protobuf field: proto.HandType hand_type = 6;
     */
    handType: HandType;
    /**
     * @generated from protobuf field: proto.RangedWeaponType ranged_weapon_type = 7;
     */
    rangedWeaponType: RangedWeaponType;
    /**
     * @generated from protobuf field: repeated double stats = 8;
     */
    stats: number[];
    /**
     * @generated from protobuf field: repeated proto.GemColor gem_sockets = 9;
     */
    gemSockets: GemColor[];
    /**
     * @generated from protobuf field: repeated double socketBonus = 10;
     */
    socketBonus: number[];
    /**
     * @generated from protobuf field: int32 phase = 11;
     */
    phase: number;
    /**
     * @generated from protobuf field: proto.ItemQuality quality = 12;
     */
    quality: ItemQuality;
    /**
     * @generated from protobuf field: bool unique = 13;
     */
    unique: boolean;
}
/**
 * @generated from protobuf message proto.Enchant
 */
export interface Enchant {
    /**
     * @generated from protobuf field: int32 id = 1;
     */
    id: number;
    /**
     * @generated from protobuf field: int32 effect_id = 2;
     */
    effectId: number;
    /**
     * @generated from protobuf field: string name = 3;
     */
    name: string;
    /**
     * @generated from protobuf field: proto.ItemType type = 4;
     */
    type: ItemType;
    /**
     * @generated from protobuf field: bool two_handed_only = 5;
     */
    twoHandedOnly: boolean;
    /**
     * @generated from protobuf field: bool shield_only = 6;
     */
    shieldOnly: boolean;
    /**
     * @generated from protobuf field: repeated double stats = 7;
     */
    stats: number[];
    /**
     * @generated from protobuf field: proto.ItemQuality quality = 8;
     */
    quality: ItemQuality;
}
/**
 * @generated from protobuf message proto.Gem
 */
export interface Gem {
    /**
     * @generated from protobuf field: int32 id = 1;
     */
    id: number;
    /**
     * @generated from protobuf field: string name = 2;
     */
    name: string;
    /**
     * @generated from protobuf field: repeated double stats = 3;
     */
    stats: number[];
    /**
     * @generated from protobuf field: proto.GemColor color = 4;
     */
    color: GemColor;
    /**
     * @generated from protobuf field: int32 phase = 5;
     */
    phase: number;
    /**
     * @generated from protobuf field: proto.ItemQuality quality = 6;
     */
    quality: ItemQuality;
    /**
     * @generated from protobuf field: bool unique = 7;
     */
    unique: boolean;
}
/**
 * @generated from protobuf enum proto.Spec
 */
export declare enum Spec {
    /**
     * @generated from protobuf enum value: SpecBalanceDruid = 0;
     */
    SpecBalanceDruid = 0,
    /**
     * @generated from protobuf enum value: SpecElementalShaman = 1;
     */
    SpecElementalShaman = 1,
    /**
     * @generated from protobuf enum value: SpecHunter = 8;
     */
    SpecHunter = 8,
    /**
     * @generated from protobuf enum value: SpecMage = 2;
     */
    SpecMage = 2,
    /**
     * @generated from protobuf enum value: SpecRetributionPaladin = 3;
     */
    SpecRetributionPaladin = 3,
    /**
     * @generated from protobuf enum value: SpecRogue = 7;
     */
    SpecRogue = 7,
    /**
     * @generated from protobuf enum value: SpecShadowPriest = 4;
     */
    SpecShadowPriest = 4,
    /**
     * @generated from protobuf enum value: SpecWarlock = 5;
     */
    SpecWarlock = 5,
    /**
     * @generated from protobuf enum value: SpecWarrior = 6;
     */
    SpecWarrior = 6
}
/**
 * @generated from protobuf enum proto.Race
 */
export declare enum Race {
    /**
     * @generated from protobuf enum value: RaceUnknown = 0;
     */
    RaceUnknown = 0,
    /**
     * @generated from protobuf enum value: RaceBloodElf = 1;
     */
    RaceBloodElf = 1,
    /**
     * @generated from protobuf enum value: RaceDraenei = 2;
     */
    RaceDraenei = 2,
    /**
     * @generated from protobuf enum value: RaceDwarf = 3;
     */
    RaceDwarf = 3,
    /**
     * @generated from protobuf enum value: RaceGnome = 4;
     */
    RaceGnome = 4,
    /**
     * @generated from protobuf enum value: RaceHuman = 5;
     */
    RaceHuman = 5,
    /**
     * @generated from protobuf enum value: RaceNightElf = 6;
     */
    RaceNightElf = 6,
    /**
     * @generated from protobuf enum value: RaceOrc = 7;
     */
    RaceOrc = 7,
    /**
     * @generated from protobuf enum value: RaceTauren = 8;
     */
    RaceTauren = 8,
    /**
     * @generated from protobuf enum value: RaceTroll10 = 9;
     */
    RaceTroll10 = 9,
    /**
     * @generated from protobuf enum value: RaceTroll30 = 10;
     */
    RaceTroll30 = 10,
    /**
     * @generated from protobuf enum value: RaceUndead = 11;
     */
    RaceUndead = 11
}
/**
 * @generated from protobuf enum proto.Class
 */
export declare enum Class {
    /**
     * @generated from protobuf enum value: ClassUnknown = 0;
     */
    ClassUnknown = 0,
    /**
     * @generated from protobuf enum value: ClassDruid = 1;
     */
    ClassDruid = 1,
    /**
     * @generated from protobuf enum value: ClassHunter = 2;
     */
    ClassHunter = 2,
    /**
     * @generated from protobuf enum value: ClassMage = 3;
     */
    ClassMage = 3,
    /**
     * @generated from protobuf enum value: ClassPaladin = 4;
     */
    ClassPaladin = 4,
    /**
     * @generated from protobuf enum value: ClassPriest = 5;
     */
    ClassPriest = 5,
    /**
     * @generated from protobuf enum value: ClassRogue = 6;
     */
    ClassRogue = 6,
    /**
     * @generated from protobuf enum value: ClassShaman = 7;
     */
    ClassShaman = 7,
    /**
     * @generated from protobuf enum value: ClassWarlock = 8;
     */
    ClassWarlock = 8,
    /**
     * @generated from protobuf enum value: ClassWarrior = 9;
     */
    ClassWarrior = 9
}
/**
 * @generated from protobuf enum proto.Stat
 */
export declare enum Stat {
    /**
     * @generated from protobuf enum value: StatStrength = 0;
     */
    StatStrength = 0,
    /**
     * @generated from protobuf enum value: StatAgility = 1;
     */
    StatAgility = 1,
    /**
     * @generated from protobuf enum value: StatStamina = 2;
     */
    StatStamina = 2,
    /**
     * @generated from protobuf enum value: StatIntellect = 3;
     */
    StatIntellect = 3,
    /**
     * @generated from protobuf enum value: StatSpirit = 4;
     */
    StatSpirit = 4,
    /**
     * @generated from protobuf enum value: StatSpellPower = 5;
     */
    StatSpellPower = 5,
    /**
     * @generated from protobuf enum value: StatHealingPower = 6;
     */
    StatHealingPower = 6,
    /**
     * @generated from protobuf enum value: StatArcaneSpellPower = 7;
     */
    StatArcaneSpellPower = 7,
    /**
     * @generated from protobuf enum value: StatFireSpellPower = 8;
     */
    StatFireSpellPower = 8,
    /**
     * @generated from protobuf enum value: StatFrostSpellPower = 9;
     */
    StatFrostSpellPower = 9,
    /**
     * @generated from protobuf enum value: StatHolySpellPower = 10;
     */
    StatHolySpellPower = 10,
    /**
     * @generated from protobuf enum value: StatNatureSpellPower = 11;
     */
    StatNatureSpellPower = 11,
    /**
     * @generated from protobuf enum value: StatShadowSpellPower = 12;
     */
    StatShadowSpellPower = 12,
    /**
     * @generated from protobuf enum value: StatMP5 = 13;
     */
    StatMP5 = 13,
    /**
     * @generated from protobuf enum value: StatSpellHit = 14;
     */
    StatSpellHit = 14,
    /**
     * @generated from protobuf enum value: StatSpellCrit = 15;
     */
    StatSpellCrit = 15,
    /**
     * @generated from protobuf enum value: StatSpellHaste = 16;
     */
    StatSpellHaste = 16,
    /**
     * @generated from protobuf enum value: StatSpellPenetration = 17;
     */
    StatSpellPenetration = 17,
    /**
     * @generated from protobuf enum value: StatAttackPower = 18;
     */
    StatAttackPower = 18,
    /**
     * @generated from protobuf enum value: StatMeleeHit = 19;
     */
    StatMeleeHit = 19,
    /**
     * @generated from protobuf enum value: StatMeleeCrit = 20;
     */
    StatMeleeCrit = 20,
    /**
     * @generated from protobuf enum value: StatMeleeHaste = 21;
     */
    StatMeleeHaste = 21,
    /**
     * @generated from protobuf enum value: StatArmorPenetration = 22;
     */
    StatArmorPenetration = 22,
    /**
     * @generated from protobuf enum value: StatExpertise = 23;
     */
    StatExpertise = 23,
    /**
     * @generated from protobuf enum value: StatMana = 24;
     */
    StatMana = 24,
    /**
     * @generated from protobuf enum value: StatEnergy = 25;
     */
    StatEnergy = 25,
    /**
     * @generated from protobuf enum value: StatRage = 26;
     */
    StatRage = 26,
    /**
     * @generated from protobuf enum value: StatArmor = 27;
     */
    StatArmor = 27
}
/**
 * Does not correspond to anything in-game; just our own label to help filter
 * items in the UI.
 *
 * @generated from protobuf enum proto.ItemCategory
 */
export declare enum ItemCategory {
    /**
     * @generated from protobuf enum value: ItemCategoryUnknown = 0;
     */
    ItemCategoryUnknown = 0,
    /**
     * @generated from protobuf enum value: ItemCategoryCaster = 1;
     */
    ItemCategoryCaster = 1,
    /**
     * @generated from protobuf enum value: ItemCategoryMelee = 2;
     */
    ItemCategoryMelee = 2,
    /**
     * @generated from protobuf enum value: ItemCategoryHybrid = 3;
     */
    ItemCategoryHybrid = 3,
    /**
     * @generated from protobuf enum value: ItemCategoryHealer = 4;
     */
    ItemCategoryHealer = 4
}
/**
 * @generated from protobuf enum proto.ItemType
 */
export declare enum ItemType {
    /**
     * @generated from protobuf enum value: ItemTypeUnknown = 0;
     */
    ItemTypeUnknown = 0,
    /**
     * @generated from protobuf enum value: ItemTypeHead = 1;
     */
    ItemTypeHead = 1,
    /**
     * @generated from protobuf enum value: ItemTypeNeck = 2;
     */
    ItemTypeNeck = 2,
    /**
     * @generated from protobuf enum value: ItemTypeShoulder = 3;
     */
    ItemTypeShoulder = 3,
    /**
     * @generated from protobuf enum value: ItemTypeBack = 4;
     */
    ItemTypeBack = 4,
    /**
     * @generated from protobuf enum value: ItemTypeChest = 5;
     */
    ItemTypeChest = 5,
    /**
     * @generated from protobuf enum value: ItemTypeWrist = 6;
     */
    ItemTypeWrist = 6,
    /**
     * @generated from protobuf enum value: ItemTypeHands = 7;
     */
    ItemTypeHands = 7,
    /**
     * @generated from protobuf enum value: ItemTypeWaist = 8;
     */
    ItemTypeWaist = 8,
    /**
     * @generated from protobuf enum value: ItemTypeLegs = 9;
     */
    ItemTypeLegs = 9,
    /**
     * @generated from protobuf enum value: ItemTypeFeet = 10;
     */
    ItemTypeFeet = 10,
    /**
     * @generated from protobuf enum value: ItemTypeFinger = 11;
     */
    ItemTypeFinger = 11,
    /**
     * @generated from protobuf enum value: ItemTypeTrinket = 12;
     */
    ItemTypeTrinket = 12,
    /**
     * @generated from protobuf enum value: ItemTypeWeapon = 13;
     */
    ItemTypeWeapon = 13,
    /**
     * @generated from protobuf enum value: ItemTypeRanged = 14;
     */
    ItemTypeRanged = 14
}
/**
 * @generated from protobuf enum proto.ArmorType
 */
export declare enum ArmorType {
    /**
     * @generated from protobuf enum value: ArmorTypeUnknown = 0;
     */
    ArmorTypeUnknown = 0,
    /**
     * @generated from protobuf enum value: ArmorTypeCloth = 1;
     */
    ArmorTypeCloth = 1,
    /**
     * @generated from protobuf enum value: ArmorTypeLeather = 2;
     */
    ArmorTypeLeather = 2,
    /**
     * @generated from protobuf enum value: ArmorTypeMail = 3;
     */
    ArmorTypeMail = 3,
    /**
     * @generated from protobuf enum value: ArmorTypePlate = 4;
     */
    ArmorTypePlate = 4
}
/**
 * @generated from protobuf enum proto.WeaponType
 */
export declare enum WeaponType {
    /**
     * @generated from protobuf enum value: WeaponTypeUnknown = 0;
     */
    WeaponTypeUnknown = 0,
    /**
     * @generated from protobuf enum value: WeaponTypeAxe = 1;
     */
    WeaponTypeAxe = 1,
    /**
     * @generated from protobuf enum value: WeaponTypeDagger = 2;
     */
    WeaponTypeDagger = 2,
    /**
     * @generated from protobuf enum value: WeaponTypeFist = 3;
     */
    WeaponTypeFist = 3,
    /**
     * @generated from protobuf enum value: WeaponTypeMace = 4;
     */
    WeaponTypeMace = 4,
    /**
     * @generated from protobuf enum value: WeaponTypeOffHand = 5;
     */
    WeaponTypeOffHand = 5,
    /**
     * @generated from protobuf enum value: WeaponTypePolearm = 6;
     */
    WeaponTypePolearm = 6,
    /**
     * @generated from protobuf enum value: WeaponTypeShield = 7;
     */
    WeaponTypeShield = 7,
    /**
     * @generated from protobuf enum value: WeaponTypeStaff = 8;
     */
    WeaponTypeStaff = 8,
    /**
     * @generated from protobuf enum value: WeaponTypeSword = 9;
     */
    WeaponTypeSword = 9
}
/**
 * @generated from protobuf enum proto.HandType
 */
export declare enum HandType {
    /**
     * @generated from protobuf enum value: HandTypeUnknown = 0;
     */
    HandTypeUnknown = 0,
    /**
     * @generated from protobuf enum value: HandTypeMainHand = 1;
     */
    HandTypeMainHand = 1,
    /**
     * @generated from protobuf enum value: HandTypeOneHand = 2;
     */
    HandTypeOneHand = 2,
    /**
     * @generated from protobuf enum value: HandTypeOffHand = 3;
     */
    HandTypeOffHand = 3,
    /**
     * @generated from protobuf enum value: HandTypeTwoHand = 4;
     */
    HandTypeTwoHand = 4
}
/**
 * @generated from protobuf enum proto.RangedWeaponType
 */
export declare enum RangedWeaponType {
    /**
     * @generated from protobuf enum value: RangedWeaponTypeUnknown = 0;
     */
    RangedWeaponTypeUnknown = 0,
    /**
     * @generated from protobuf enum value: RangedWeaponTypeBow = 1;
     */
    RangedWeaponTypeBow = 1,
    /**
     * @generated from protobuf enum value: RangedWeaponTypeCrossbow = 2;
     */
    RangedWeaponTypeCrossbow = 2,
    /**
     * @generated from protobuf enum value: RangedWeaponTypeGun = 3;
     */
    RangedWeaponTypeGun = 3,
    /**
     * @generated from protobuf enum value: RangedWeaponTypeIdol = 4;
     */
    RangedWeaponTypeIdol = 4,
    /**
     * @generated from protobuf enum value: RangedWeaponTypeLibram = 5;
     */
    RangedWeaponTypeLibram = 5,
    /**
     * @generated from protobuf enum value: RangedWeaponTypeThrown = 6;
     */
    RangedWeaponTypeThrown = 6,
    /**
     * @generated from protobuf enum value: RangedWeaponTypeTotem = 7;
     */
    RangedWeaponTypeTotem = 7,
    /**
     * @generated from protobuf enum value: RangedWeaponTypeWand = 8;
     */
    RangedWeaponTypeWand = 8
}
/**
 * All slots on the gear menu where a single item can be worn.
 *
 * @generated from protobuf enum proto.ItemSlot
 */
export declare enum ItemSlot {
    /**
     * @generated from protobuf enum value: ItemSlotHead = 0;
     */
    ItemSlotHead = 0,
    /**
     * @generated from protobuf enum value: ItemSlotNeck = 1;
     */
    ItemSlotNeck = 1,
    /**
     * @generated from protobuf enum value: ItemSlotShoulder = 2;
     */
    ItemSlotShoulder = 2,
    /**
     * @generated from protobuf enum value: ItemSlotBack = 3;
     */
    ItemSlotBack = 3,
    /**
     * @generated from protobuf enum value: ItemSlotChest = 4;
     */
    ItemSlotChest = 4,
    /**
     * @generated from protobuf enum value: ItemSlotWrist = 5;
     */
    ItemSlotWrist = 5,
    /**
     * @generated from protobuf enum value: ItemSlotHands = 6;
     */
    ItemSlotHands = 6,
    /**
     * @generated from protobuf enum value: ItemSlotWaist = 7;
     */
    ItemSlotWaist = 7,
    /**
     * @generated from protobuf enum value: ItemSlotLegs = 8;
     */
    ItemSlotLegs = 8,
    /**
     * @generated from protobuf enum value: ItemSlotFeet = 9;
     */
    ItemSlotFeet = 9,
    /**
     * @generated from protobuf enum value: ItemSlotFinger1 = 10;
     */
    ItemSlotFinger1 = 10,
    /**
     * @generated from protobuf enum value: ItemSlotFinger2 = 11;
     */
    ItemSlotFinger2 = 11,
    /**
     * @generated from protobuf enum value: ItemSlotTrinket1 = 12;
     */
    ItemSlotTrinket1 = 12,
    /**
     * @generated from protobuf enum value: ItemSlotTrinket2 = 13;
     */
    ItemSlotTrinket2 = 13,
    /**
     * can be 1h or 2h
     *
     * @generated from protobuf enum value: ItemSlotMainHand = 14;
     */
    ItemSlotMainHand = 14,
    /**
     * @generated from protobuf enum value: ItemSlotOffHand = 15;
     */
    ItemSlotOffHand = 15,
    /**
     * @generated from protobuf enum value: ItemSlotRanged = 16;
     */
    ItemSlotRanged = 16
}
/**
 * @generated from protobuf enum proto.ItemQuality
 */
export declare enum ItemQuality {
    /**
     * @generated from protobuf enum value: ItemQualityJunk = 0;
     */
    ItemQualityJunk = 0,
    /**
     * @generated from protobuf enum value: ItemQualityCommon = 1;
     */
    ItemQualityCommon = 1,
    /**
     * @generated from protobuf enum value: ItemQualityUncommon = 2;
     */
    ItemQualityUncommon = 2,
    /**
     * @generated from protobuf enum value: ItemQualityRare = 3;
     */
    ItemQualityRare = 3,
    /**
     * @generated from protobuf enum value: ItemQualityEpic = 4;
     */
    ItemQualityEpic = 4,
    /**
     * @generated from protobuf enum value: ItemQualityLegendary = 5;
     */
    ItemQualityLegendary = 5
}
/**
 * @generated from protobuf enum proto.GemColor
 */
export declare enum GemColor {
    /**
     * @generated from protobuf enum value: GemColorUnknown = 0;
     */
    GemColorUnknown = 0,
    /**
     * @generated from protobuf enum value: GemColorMeta = 1;
     */
    GemColorMeta = 1,
    /**
     * @generated from protobuf enum value: GemColorRed = 2;
     */
    GemColorRed = 2,
    /**
     * @generated from protobuf enum value: GemColorBlue = 3;
     */
    GemColorBlue = 3,
    /**
     * @generated from protobuf enum value: GemColorYellow = 4;
     */
    GemColorYellow = 4,
    /**
     * @generated from protobuf enum value: GemColorGreen = 5;
     */
    GemColorGreen = 5,
    /**
     * @generated from protobuf enum value: GemColorOrange = 6;
     */
    GemColorOrange = 6,
    /**
     * @generated from protobuf enum value: GemColorPurple = 7;
     */
    GemColorPurple = 7,
    /**
     * @generated from protobuf enum value: GemColorPrismatic = 8;
     */
    GemColorPrismatic = 8
}
/**
 * @generated from protobuf enum proto.TristateEffect
 */
export declare enum TristateEffect {
    /**
     * @generated from protobuf enum value: TristateEffectMissing = 0;
     */
    TristateEffectMissing = 0,
    /**
     * @generated from protobuf enum value: TristateEffectRegular = 1;
     */
    TristateEffectRegular = 1,
    /**
     * @generated from protobuf enum value: TristateEffectImproved = 2;
     */
    TristateEffectImproved = 2
}
/**
 * @generated from protobuf enum proto.Drums
 */
export declare enum Drums {
    /**
     * @generated from protobuf enum value: DrumsUnknown = 0;
     */
    DrumsUnknown = 0,
    /**
     * @generated from protobuf enum value: DrumsOfBattle = 1;
     */
    DrumsOfBattle = 1,
    /**
     * @generated from protobuf enum value: DrumsOfRestoration = 2;
     */
    DrumsOfRestoration = 2
}
/**
 * @generated from protobuf enum proto.Potions
 */
export declare enum Potions {
    /**
     * @generated from protobuf enum value: UnknownPotion = 0;
     */
    UnknownPotion = 0,
    /**
     * @generated from protobuf enum value: DestructionPotion = 1;
     */
    DestructionPotion = 1,
    /**
     * @generated from protobuf enum value: SuperManaPotion = 2;
     */
    SuperManaPotion = 2
}
declare class Buffs$Type extends MessageType<Buffs> {
    constructor();
    create(value?: PartialMessage<Buffs>): Buffs;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Buffs): Buffs;
    internalBinaryWrite(message: Buffs, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Buffs
 */
export declare const Buffs: Buffs$Type;
declare class Consumes$Type extends MessageType<Consumes> {
    constructor();
    create(value?: PartialMessage<Consumes>): Consumes;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Consumes): Consumes;
    internalBinaryWrite(message: Consumes, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Consumes
 */
export declare const Consumes: Consumes$Type;
declare class Debuffs$Type extends MessageType<Debuffs> {
    constructor();
    create(value?: PartialMessage<Debuffs>): Debuffs;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Debuffs): Debuffs;
    internalBinaryWrite(message: Debuffs, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Debuffs
 */
export declare const Debuffs: Debuffs$Type;
declare class Target$Type extends MessageType<Target> {
    constructor();
    create(value?: PartialMessage<Target>): Target;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Target): Target;
    internalBinaryWrite(message: Target, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Target
 */
export declare const Target: Target$Type;
declare class Encounter$Type extends MessageType<Encounter> {
    constructor();
    create(value?: PartialMessage<Encounter>): Encounter;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Encounter): Encounter;
    internalBinaryWrite(message: Encounter, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Encounter
 */
export declare const Encounter: Encounter$Type;
declare class ItemSpec$Type extends MessageType<ItemSpec> {
    constructor();
    create(value?: PartialMessage<ItemSpec>): ItemSpec;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ItemSpec): ItemSpec;
    internalBinaryWrite(message: ItemSpec, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ItemSpec
 */
export declare const ItemSpec: ItemSpec$Type;
declare class EquipmentSpec$Type extends MessageType<EquipmentSpec> {
    constructor();
    create(value?: PartialMessage<EquipmentSpec>): EquipmentSpec;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: EquipmentSpec): EquipmentSpec;
    internalBinaryWrite(message: EquipmentSpec, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.EquipmentSpec
 */
export declare const EquipmentSpec: EquipmentSpec$Type;
declare class Item$Type extends MessageType<Item> {
    constructor();
    create(value?: PartialMessage<Item>): Item;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Item): Item;
    internalBinaryWrite(message: Item, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Item
 */
export declare const Item: Item$Type;
declare class Enchant$Type extends MessageType<Enchant> {
    constructor();
    create(value?: PartialMessage<Enchant>): Enchant;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Enchant): Enchant;
    internalBinaryWrite(message: Enchant, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Enchant
 */
export declare const Enchant: Enchant$Type;
declare class Gem$Type extends MessageType<Gem> {
    constructor();
    create(value?: PartialMessage<Gem>): Gem;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Gem): Gem;
    internalBinaryWrite(message: Gem, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Gem
 */
export declare const Gem: Gem$Type;
export {};
