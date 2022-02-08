import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * Buffs that affect the entire raid.
 *
 * @generated from protobuf message proto.RaidBuffs
 */
export interface RaidBuffs {
    /**
     * @generated from protobuf field: bool arcane_brilliance = 1;
     */
    arcaneBrilliance: boolean;
    /**
     * @generated from protobuf field: proto.TristateEffect divine_spirit = 4;
     */
    divineSpirit: TristateEffect;
    /**
     * @generated from protobuf field: proto.TristateEffect gift_of_the_wild = 5;
     */
    giftOfTheWild: TristateEffect;
}
/**
 * Buffs that affect a single party.
 *
 * @generated from protobuf message proto.PartyBuffs
 */
export interface PartyBuffs {
    /**
     * @generated from protobuf field: int32 bloodlust = 1;
     */
    bloodlust: number;
    /**
     * @generated from protobuf field: int32 ferocious_inspiration = 22;
     */
    ferociousInspiration: number;
    /**
     * @generated from protobuf field: int32 battle_chickens = 27;
     */
    battleChickens: number;
    /**
     * @generated from protobuf field: proto.TristateEffect moonkin_aura = 2;
     */
    moonkinAura: TristateEffect;
    /**
     * @generated from protobuf field: proto.TristateEffect leader_of_the_pack = 19;
     */
    leaderOfThePack: TristateEffect;
    /**
     * @generated from protobuf field: proto.TristateEffect sanctity_aura = 20;
     */
    sanctityAura: TristateEffect;
    /**
     * @generated from protobuf field: bool trueshot_aura = 21;
     */
    trueshotAura: boolean;
    /**
     * @generated from protobuf field: bool draenei_racial_melee = 3;
     */
    draeneiRacialMelee: boolean;
    /**
     * @generated from protobuf field: bool draenei_racial_caster = 4;
     */
    draeneiRacialCaster: boolean;
    /**
     * Drums
     *
     * @generated from protobuf field: proto.Drums drums = 5;
     */
    drums: Drums;
    /**
     * Item Buffs
     *
     * @generated from protobuf field: int32 atiesh_mage = 6;
     */
    atieshMage: number;
    /**
     * @generated from protobuf field: int32 atiesh_warlock = 7;
     */
    atieshWarlock: number;
    /**
     * @generated from protobuf field: bool braided_eternium_chain = 8;
     */
    braidedEterniumChain: boolean;
    /**
     * @generated from protobuf field: bool eye_of_the_night = 9;
     */
    eyeOfTheNight: boolean;
    /**
     * @generated from protobuf field: bool chain_of_the_twilight_owl = 10;
     */
    chainOfTheTwilightOwl: boolean;
    /**
     * @generated from protobuf field: bool jade_pendant_of_blasting = 11;
     */
    jadePendantOfBlasting: boolean;
    /**
     * Totems
     *
     * @generated from protobuf field: proto.TristateEffect mana_spring_totem = 12;
     */
    manaSpringTotem: TristateEffect;
    /**
     * @generated from protobuf field: int32 mana_tide_totems = 17;
     */
    manaTideTotems: number;
    /**
     * @generated from protobuf field: int32 totem_of_wrath = 13;
     */
    totemOfWrath: number;
    /**
     * @generated from protobuf field: proto.TristateEffect wrath_of_air_totem = 14;
     */
    wrathOfAirTotem: TristateEffect;
    /**
     * @generated from protobuf field: bool snapshot_improved_wrath_of_air_totem = 25;
     */
    snapshotImprovedWrathOfAirTotem: boolean;
    /**
     * @generated from protobuf field: proto.TristateEffect grace_of_air_totem = 15;
     */
    graceOfAirTotem: TristateEffect;
    /**
     * @generated from protobuf field: proto.StrengthOfEarthType strength_of_earth_totem = 16;
     */
    strengthOfEarthTotem: StrengthOfEarthType;
    /**
     * @generated from protobuf field: bool tranquil_air_totem = 26;
     */
    tranquilAirTotem: boolean;
    /**
     * @generated from protobuf field: int32 windfury_totem_rank = 23;
     */
    windfuryTotemRank: number;
    /**
     * @generated from protobuf field: int32 windfury_totem_iwt = 24;
     */
    windfuryTotemIwt: number;
    /**
     * @generated from protobuf field: proto.TristateEffect battle_shout = 18;
     */
    battleShout: TristateEffect;
    /**
     * @generated from protobuf field: bool bs_solarian_sapphire = 28;
     */
    bsSolarianSapphire: boolean;
    /**
     * @generated from protobuf field: bool snapshot_bs_solarian_sapphire = 29;
     */
    snapshotBsSolarianSapphire: boolean;
    /**
     * @generated from protobuf field: bool snapshot_bs_t2 = 30;
     */
    snapshotBsT2: boolean;
}
/**
 * Buffs are only used by individual sims, never the raid sim.
 * These are usually individuals of actions taken by other Characters.
 *
 * @generated from protobuf message proto.IndividualBuffs
 */
export interface IndividualBuffs {
    /**
     * @generated from protobuf field: bool blessing_of_kings = 1;
     */
    blessingOfKings: boolean;
    /**
     * @generated from protobuf field: bool blessing_of_salvation = 8;
     */
    blessingOfSalvation: boolean;
    /**
     * @generated from protobuf field: proto.TristateEffect blessing_of_wisdom = 2;
     */
    blessingOfWisdom: TristateEffect;
    /**
     * @generated from protobuf field: proto.TristateEffect blessing_of_might = 3;
     */
    blessingOfMight: TristateEffect;
    /**
     * @generated from protobuf field: int32 shadow_priest_dps = 4;
     */
    shadowPriestDps: number;
    /**
     * @generated from protobuf field: bool unleashed_rage = 7;
     */
    unleashedRage: boolean;
    /**
     * How many of each of these buffs the player will be receiving.
     *
     * @generated from protobuf field: int32 innervates = 5;
     */
    innervates: number;
    /**
     * @generated from protobuf field: int32 power_infusions = 6;
     */
    powerInfusions: number;
}
/**
 * @generated from protobuf message proto.Consumes
 */
export interface Consumes {
    /**
     * Deprecated in favor of flask.
     * Remove on 3/4/2022 (1 month).
     *
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
     * @generated from protobuf field: bool flask_of_relentless_assault = 21;
     */
    flaskOfRelentlessAssault: boolean;
    /**
     * @generated from protobuf field: proto.Flask flask = 38;
     */
    flask: Flask;
    /**
     * Deprecated in favor of battle_elixir.
     * Remove on 3/4/2022 (1 month).
     *
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
     * @generated from protobuf field: bool elixir_of_major_agility = 22;
     */
    elixirOfMajorAgility: boolean;
    /**
     * @generated from protobuf field: bool elixir_of_major_strength = 31;
     */
    elixirOfMajorStrength: boolean;
    /**
     * @generated from protobuf field: bool elixir_of_demonslaying = 23;
     */
    elixirOfDemonslaying: boolean;
    /**
     * @generated from protobuf field: bool elixir_of_the_mongoose = 30;
     */
    elixirOfTheMongoose: boolean;
    /**
     * @generated from protobuf field: proto.BattleElixir battle_elixir = 39;
     */
    battleElixir: BattleElixir;
    /**
     * Deprecated in favor of guardian_elixir.
     * Remove on 3/4/2022 (1 month).
     *
     * @generated from protobuf field: bool elixir_of_draenic_wisdom = 9;
     */
    elixirOfDraenicWisdom: boolean;
    /**
     * @generated from protobuf field: bool elixir_of_major_mageblood = 10;
     */
    elixirOfMajorMageblood: boolean;
    /**
     * @generated from protobuf field: proto.GuardianElixir guardian_elixir = 40;
     */
    guardianElixir: GuardianElixir;
    /**
     * Deprecated, use main_hand_imbue instead.
     * Remove on 2/18/2022 (1 month).
     *
     * @generated from protobuf field: bool brilliant_wizard_oil = 11;
     */
    brilliantWizardOil: boolean;
    /**
     * @generated from protobuf field: bool superior_wizard_oil = 12;
     */
    superiorWizardOil: boolean;
    /**
     * @generated from protobuf field: proto.WeaponImbue main_hand_imbue = 32;
     */
    mainHandImbue: WeaponImbue;
    /**
     * @generated from protobuf field: proto.WeaponImbue off_hand_imbue = 33;
     */
    offHandImbue: WeaponImbue;
    /**
     * Deprecated in favor of food.
     * Remove on 3/4/2022 (1 month).
     *
     * @generated from protobuf field: bool blackened_basilisk = 13;
     */
    blackenedBasilisk: boolean;
    /**
     * @generated from protobuf field: bool skullfish_soup = 14;
     */
    skullfishSoup: boolean;
    /**
     * @generated from protobuf field: bool roasted_clefthoof = 24;
     */
    roastedClefthoof: boolean;
    /**
     * @generated from protobuf field: bool spicy_hot_talbuk = 29;
     */
    spicyHotTalbuk: boolean;
    /**
     * @generated from protobuf field: bool grilled_mudfish = 35;
     */
    grilledMudfish: boolean;
    /**
     * @generated from protobuf field: bool ravager_dog = 36;
     */
    ravagerDog: boolean;
    /**
     * @generated from protobuf field: proto.Food food = 41;
     */
    food: Food;
    /**
     * @generated from protobuf field: proto.PetFood pet_food = 37;
     */
    petFood: PetFood;
    /**
     * Deprecated in favor of alchohol.
     * Remove on 3/4/2022 (1 month).
     *
     * @generated from protobuf field: bool kreegsStoutBeatdown = 20;
     */
    kreegsStoutBeatdown: boolean;
    /**
     * @generated from protobuf field: proto.Alchohol alchohol = 42;
     */
    alchohol: Alchohol;
    /**
     * Deprecated in favor of int32 versions below.
     * Remove on 3/4/2022 (1 month).
     *
     * @generated from protobuf field: bool scroll_of_strength_v = 25;
     */
    scrollOfStrengthV: boolean;
    /**
     * @generated from protobuf field: bool scroll_of_agility_v = 26;
     */
    scrollOfAgilityV: boolean;
    /**
     * @generated from protobuf field: bool scroll_of_spirit_v = 28;
     */
    scrollOfSpiritV: boolean;
    /**
     * 0 means no scroll, otherwise value is the scroll level.
     * E.g. 5 indicates Scroll of Agility V.
     *
     * @generated from protobuf field: int32 scroll_of_agility = 44;
     */
    scrollOfAgility: number;
    /**
     * @generated from protobuf field: int32 scroll_of_strength = 43;
     */
    scrollOfStrength: number;
    /**
     * @generated from protobuf field: int32 scroll_of_spirit = 45;
     */
    scrollOfSpirit: number;
    /**
     * @generated from protobuf field: int32 pet_scroll_of_agility = 46;
     */
    petScrollOfAgility: number;
    /**
     * @generated from protobuf field: int32 pet_scroll_of_strength = 47;
     */
    petScrollOfStrength: number;
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
     * @generated from protobuf field: proto.Conjured default_conjured = 27;
     */
    defaultConjured: Conjured;
    /**
     * @generated from protobuf field: proto.Drums drums = 19;
     */
    drums: Drums;
    /**
     * @generated from protobuf field: bool battle_chicken = 34;
     */
    battleChicken: boolean;
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
    /**
     * @generated from protobuf field: proto.TristateEffect curse_of_elements = 4;
     */
    curseOfElements: TristateEffect;
    /**
     * @generated from protobuf field: double isb_uptime = 5;
     */
    isbUptime: number;
    /**
     * @generated from protobuf field: bool improved_scorch = 6;
     */
    improvedScorch: boolean;
    /**
     * @generated from protobuf field: bool winters_chill = 7;
     */
    wintersChill: boolean;
    /**
     * @generated from protobuf field: bool blood_frenzy = 8;
     */
    bloodFrenzy: boolean;
    /**
     * @generated from protobuf field: proto.TristateEffect expose_armor = 9;
     */
    exposeArmor: TristateEffect;
    /**
     * @generated from protobuf field: proto.TristateEffect faerie_fire = 10;
     */
    faerieFire: TristateEffect;
    /**
     * @generated from protobuf field: bool sunder_armor = 11;
     */
    sunderArmor: boolean;
    /**
     * @generated from protobuf field: bool curse_of_recklessness = 12;
     */
    curseOfRecklessness: boolean;
    /**
     * @generated from protobuf field: proto.TristateEffect hunters_mark = 15;
     */
    huntersMark: TristateEffect;
    /**
     * @generated from protobuf field: double expose_weakness_uptime = 13;
     */
    exposeWeaknessUptime: number;
    /**
     * @generated from protobuf field: double expose_weakness_hunter_agility = 14;
     */
    exposeWeaknessHunterAgility: number;
}
/**
 * @generated from protobuf message proto.Target
 */
export interface Target {
    /**
     * @generated from protobuf field: double armor = 1;
     */
    armor: number;
    /**
     * @generated from protobuf field: int32 level = 4;
     */
    level: number;
    /**
     * @generated from protobuf field: proto.MobType mob_type = 3;
     */
    mobType: MobType;
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
     * Variation in the duration
     *
     * @generated from protobuf field: double duration_variation = 4;
     */
    durationVariation: number;
    /**
     * The ratio of the encounter duration, between 0 and 1, for which the targets
     * will be in execute range for the purposes of Warrior Execute, Mage Molten
     * Fury, etc.
     *
     * @generated from protobuf field: double execute_proportion = 3;
     */
    executeProportion: number;
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
     * This is unused by most items. For most items we set id to the
     * wowhead/in-game ID directly. For random enchant items though we need to
     * use unique hardcoded IDs so this field holds the wowhead ID instead.
     *
     * @generated from protobuf field: int32 wowhead_id = 16;
     */
    wowheadId: number;
    /**
     * @generated from protobuf field: string name = 2;
     */
    name: string;
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
     * Weapon stats, needed for computing proper EP for melee weapons
     *
     * @generated from protobuf field: double weapon_damage_min = 17;
     */
    weaponDamageMin: number;
    /**
     * @generated from protobuf field: double weapon_damage_max = 18;
     */
    weaponDamageMax: number;
    /**
     * @generated from protobuf field: double weapon_speed = 19;
     */
    weaponSpeed: number;
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
    /**
     * @generated from protobuf field: int32 ilvl = 20;
     */
    ilvl: number;
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
     * If true, then id is the ID of the enchant spell instead of the formula item.
     * This is used by enchants for which a formula doesn't exist (its taught by a trainer).
     *
     * @generated from protobuf field: bool is_spell_id = 10;
     */
    isSpellId: boolean;
    /**
     * @generated from protobuf field: proto.ItemType type = 4;
     */
    type: ItemType;
    /**
     * @generated from protobuf field: proto.EnchantType enchant_type = 9;
     */
    enchantType: EnchantType;
    /**
     * @generated from protobuf field: repeated double stats = 7;
     */
    stats: number[];
    /**
     * @generated from protobuf field: proto.ItemQuality quality = 8;
     */
    quality: ItemQuality;
    /**
     * @generated from protobuf field: int32 phase = 11;
     */
    phase: number;
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
 * @generated from protobuf message proto.RaidTarget
 */
export interface RaidTarget {
    /**
     * Raid index of the player to target. A value of -1 indicates no target.
     *
     * @generated from protobuf field: int32 target_index = 1;
     */
    targetIndex: number;
}
/**
 * @generated from protobuf message proto.ActionID
 */
export interface ActionID {
    /**
     * @generated from protobuf oneof: raw_id
     */
    rawId: {
        oneofKind: "spellId";
        /**
         * @generated from protobuf field: int32 spell_id = 1;
         */
        spellId: number;
    } | {
        oneofKind: "itemId";
        /**
         * @generated from protobuf field: int32 item_id = 2;
         */
        itemId: number;
    } | {
        oneofKind: "otherId";
        /**
         * @generated from protobuf field: proto.OtherAction other_id = 3;
         */
        otherId: OtherAction;
    } | {
        oneofKind: undefined;
    };
    /**
     * Distinguishes between different versions of the same action.
     * Currently the only use for this is Shaman Lightning Overload.
     *
     * @generated from protobuf field: int32 tag = 4;
     */
    tag: number;
}
/**
 * Custom options for a particular cooldown.
 *
 * @generated from protobuf message proto.Cooldown
 */
export interface Cooldown {
    /**
     * Identifies the cooldown to which these settings will apply.
     *
     * @generated from protobuf field: proto.ActionID id = 1;
     */
    id?: ActionID;
    /**
     * Fixed times at which to use this cooldown. Each value corresponds to a usage,
     * e.g. first value is the first usage, second value is the second usage.
     * Any usages after the specified timings will occur as soon as possible, subject
     * to the ShouldActivate() condition.
     *
     * @generated from protobuf field: repeated double timings = 2;
     */
    timings: number[];
}
/**
 * @generated from protobuf message proto.Cooldowns
 */
export interface Cooldowns {
    /**
     * @generated from protobuf field: repeated proto.Cooldown cooldowns = 1;
     */
    cooldowns: Cooldown[];
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
     * @generated from protobuf enum value: SpecEnhancementShaman = 9;
     */
    SpecEnhancementShaman = 9,
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
    StatArmor = 27,
    /**
     * @generated from protobuf enum value: StatRangedAttackPower = 28;
     */
    StatRangedAttackPower = 28
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
    SuperManaPotion = 2,
    /**
     * @generated from protobuf enum value: HastePotion = 3;
     */
    HastePotion = 3,
    /**
     * @generated from protobuf enum value: MightyRagePotion = 4;
     */
    MightyRagePotion = 4
}
/**
 * @generated from protobuf enum proto.Conjured
 */
export declare enum Conjured {
    /**
     * @generated from protobuf enum value: ConjuredUnknown = 0;
     */
    ConjuredUnknown = 0,
    /**
     * @generated from protobuf enum value: ConjuredDarkRune = 1;
     */
    ConjuredDarkRune = 1,
    /**
     * @generated from protobuf enum value: ConjuredFlameCap = 2;
     */
    ConjuredFlameCap = 2
}
/**
 * @generated from protobuf enum proto.WeaponImbue
 */
export declare enum WeaponImbue {
    /**
     * @generated from protobuf enum value: WeaponImbueUnknown = 0;
     */
    WeaponImbueUnknown = 0,
    /**
     * @generated from protobuf enum value: WeaponImbueAdamantiteSharpeningStone = 1;
     */
    WeaponImbueAdamantiteSharpeningStone = 1,
    /**
     * @generated from protobuf enum value: WeaponImbueAdamantiteWeightstone = 5;
     */
    WeaponImbueAdamantiteWeightstone = 5,
    /**
     * @generated from protobuf enum value: WeaponImbueElementalSharpeningStone = 2;
     */
    WeaponImbueElementalSharpeningStone = 2,
    /**
     * @generated from protobuf enum value: WeaponImbueBrilliantWizardOil = 3;
     */
    WeaponImbueBrilliantWizardOil = 3,
    /**
     * @generated from protobuf enum value: WeaponImbueSuperiorWizardOil = 4;
     */
    WeaponImbueSuperiorWizardOil = 4
}
/**
 * @generated from protobuf enum proto.Flask
 */
export declare enum Flask {
    /**
     * @generated from protobuf enum value: FlaskUnknown = 0;
     */
    FlaskUnknown = 0,
    /**
     * @generated from protobuf enum value: FlaskOfBlindingLight = 1;
     */
    FlaskOfBlindingLight = 1,
    /**
     * @generated from protobuf enum value: FlaskOfMightyRestoration = 2;
     */
    FlaskOfMightyRestoration = 2,
    /**
     * @generated from protobuf enum value: FlaskOfPureDeath = 3;
     */
    FlaskOfPureDeath = 3,
    /**
     * @generated from protobuf enum value: FlaskOfRelentlessAssault = 4;
     */
    FlaskOfRelentlessAssault = 4,
    /**
     * @generated from protobuf enum value: FlaskOfSupremePower = 5;
     */
    FlaskOfSupremePower = 5
}
/**
 * @generated from protobuf enum proto.BattleElixir
 */
export declare enum BattleElixir {
    /**
     * @generated from protobuf enum value: BattleElixirUnknown = 0;
     */
    BattleElixirUnknown = 0,
    /**
     * @generated from protobuf enum value: AdeptsElixir = 1;
     */
    AdeptsElixir = 1,
    /**
     * @generated from protobuf enum value: ElixirOfDemonslaying = 2;
     */
    ElixirOfDemonslaying = 2,
    /**
     * @generated from protobuf enum value: ElixirOfMajorAgility = 3;
     */
    ElixirOfMajorAgility = 3,
    /**
     * @generated from protobuf enum value: ElixirOfMajorFirePower = 4;
     */
    ElixirOfMajorFirePower = 4,
    /**
     * @generated from protobuf enum value: ElixirOfMajorFrostPower = 5;
     */
    ElixirOfMajorFrostPower = 5,
    /**
     * @generated from protobuf enum value: ElixirOfMajorShadowPower = 6;
     */
    ElixirOfMajorShadowPower = 6,
    /**
     * @generated from protobuf enum value: ElixirOfMajorStrength = 7;
     */
    ElixirOfMajorStrength = 7,
    /**
     * @generated from protobuf enum value: ElixirOfTheMongoose = 8;
     */
    ElixirOfTheMongoose = 8
}
/**
 * @generated from protobuf enum proto.GuardianElixir
 */
export declare enum GuardianElixir {
    /**
     * @generated from protobuf enum value: GuardianElixirUnknown = 0;
     */
    GuardianElixirUnknown = 0,
    /**
     * @generated from protobuf enum value: ElixirOfDraenicWisdom = 1;
     */
    ElixirOfDraenicWisdom = 1,
    /**
     * @generated from protobuf enum value: ElixirOfMajorMageblood = 2;
     */
    ElixirOfMajorMageblood = 2
}
/**
 * @generated from protobuf enum proto.Food
 */
export declare enum Food {
    /**
     * @generated from protobuf enum value: FoodUnknown = 0;
     */
    FoodUnknown = 0,
    /**
     * @generated from protobuf enum value: FoodBlackenedBasilisk = 1;
     */
    FoodBlackenedBasilisk = 1,
    /**
     * @generated from protobuf enum value: FoodGrilledMudfish = 2;
     */
    FoodGrilledMudfish = 2,
    /**
     * @generated from protobuf enum value: FoodRavagerDog = 3;
     */
    FoodRavagerDog = 3,
    /**
     * @generated from protobuf enum value: FoodRoastedClefthoof = 4;
     */
    FoodRoastedClefthoof = 4,
    /**
     * @generated from protobuf enum value: FoodSkullfishSoup = 5;
     */
    FoodSkullfishSoup = 5,
    /**
     * @generated from protobuf enum value: FoodSpicyHotTalbuk = 6;
     */
    FoodSpicyHotTalbuk = 6
}
/**
 * @generated from protobuf enum proto.PetFood
 */
export declare enum PetFood {
    /**
     * @generated from protobuf enum value: PetFoodUnknown = 0;
     */
    PetFoodUnknown = 0,
    /**
     * @generated from protobuf enum value: PetFoodKiblersBits = 1;
     */
    PetFoodKiblersBits = 1
}
/**
 * @generated from protobuf enum proto.Alchohol
 */
export declare enum Alchohol {
    /**
     * @generated from protobuf enum value: AlchoholUnknown = 0;
     */
    AlchoholUnknown = 0,
    /**
     * @generated from protobuf enum value: AlchoholKreegsStoutBeatdown = 1;
     */
    AlchoholKreegsStoutBeatdown = 1
}
/**
 * @generated from protobuf enum proto.StrengthOfEarthType
 */
export declare enum StrengthOfEarthType {
    /**
     * @generated from protobuf enum value: None = 0;
     */
    None = 0,
    /**
     * @generated from protobuf enum value: Basic = 1;
     */
    Basic = 1,
    /**
     * @generated from protobuf enum value: CycloneBonus = 2;
     */
    CycloneBonus = 2,
    /**
     * @generated from protobuf enum value: EnhancingTotems = 3;
     */
    EnhancingTotems = 3,
    /**
     * @generated from protobuf enum value: EnhancingAndCyclone = 4;
     */
    EnhancingAndCyclone = 4
}
/**
 * @generated from protobuf enum proto.MobType
 */
export declare enum MobType {
    /**
     * @generated from protobuf enum value: MobTypeUnknown = 0;
     */
    MobTypeUnknown = 0,
    /**
     * @generated from protobuf enum value: MobTypeBeast = 1;
     */
    MobTypeBeast = 1,
    /**
     * @generated from protobuf enum value: MobTypeDemon = 2;
     */
    MobTypeDemon = 2,
    /**
     * @generated from protobuf enum value: MobTypeDragonkin = 3;
     */
    MobTypeDragonkin = 3,
    /**
     * @generated from protobuf enum value: MobTypeElemental = 4;
     */
    MobTypeElemental = 4,
    /**
     * @generated from protobuf enum value: MobTypeGiant = 5;
     */
    MobTypeGiant = 5,
    /**
     * @generated from protobuf enum value: MobTypeHumanoid = 6;
     */
    MobTypeHumanoid = 6,
    /**
     * @generated from protobuf enum value: MobTypeMechanical = 7;
     */
    MobTypeMechanical = 7,
    /**
     * @generated from protobuf enum value: MobTypeUndead = 8;
     */
    MobTypeUndead = 8
}
/**
 * Extra enum for describing which items are eligible for an enchant, when
 * ItemType alone is not enough.
 *
 * @generated from protobuf enum proto.EnchantType
 */
export declare enum EnchantType {
    /**
     * @generated from protobuf enum value: EnchantTypeNormal = 0;
     */
    EnchantTypeNormal = 0,
    /**
     * @generated from protobuf enum value: EnchantTypeTwoHand = 1;
     */
    EnchantTypeTwoHand = 1,
    /**
     * @generated from protobuf enum value: EnchantTypeShield = 2;
     */
    EnchantTypeShield = 2
}
/**
 * ID for actions that aren't spells or items.
 *
 * @generated from protobuf enum proto.OtherAction
 */
export declare enum OtherAction {
    /**
     * @generated from protobuf enum value: OtherActionNone = 0;
     */
    OtherActionNone = 0,
    /**
     * @generated from protobuf enum value: OtherActionWait = 1;
     */
    OtherActionWait = 1,
    /**
     * @generated from protobuf enum value: OtherActionManaRegen = 2;
     */
    OtherActionManaRegen = 2,
    /**
     * @generated from protobuf enum value: OtherActionEnergyRegen = 5;
     */
    OtherActionEnergyRegen = 5,
    /**
     * @generated from protobuf enum value: OtherActionFocusRegen = 6;
     */
    OtherActionFocusRegen = 6,
    /**
     * A white hit, can be main hand or off hand.
     *
     * @generated from protobuf enum value: OtherActionAttack = 3;
     */
    OtherActionAttack = 3,
    /**
     * Default shoot action using a wand/bow/gun.
     *
     * @generated from protobuf enum value: OtherActionShoot = 4;
     */
    OtherActionShoot = 4,
    /**
     * Represents a grouping of all pet actions. Only used by the UI.
     *
     * @generated from protobuf enum value: OtherActionPet = 7;
     */
    OtherActionPet = 7
}
declare class RaidBuffs$Type extends MessageType<RaidBuffs> {
    constructor();
    create(value?: PartialMessage<RaidBuffs>): RaidBuffs;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: RaidBuffs): RaidBuffs;
    internalBinaryWrite(message: RaidBuffs, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.RaidBuffs
 */
export declare const RaidBuffs: RaidBuffs$Type;
declare class PartyBuffs$Type extends MessageType<PartyBuffs> {
    constructor();
    create(value?: PartialMessage<PartyBuffs>): PartyBuffs;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PartyBuffs): PartyBuffs;
    internalBinaryWrite(message: PartyBuffs, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.PartyBuffs
 */
export declare const PartyBuffs: PartyBuffs$Type;
declare class IndividualBuffs$Type extends MessageType<IndividualBuffs> {
    constructor();
    create(value?: PartialMessage<IndividualBuffs>): IndividualBuffs;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: IndividualBuffs): IndividualBuffs;
    internalBinaryWrite(message: IndividualBuffs, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.IndividualBuffs
 */
export declare const IndividualBuffs: IndividualBuffs$Type;
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
declare class RaidTarget$Type extends MessageType<RaidTarget> {
    constructor();
    create(value?: PartialMessage<RaidTarget>): RaidTarget;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: RaidTarget): RaidTarget;
    internalBinaryWrite(message: RaidTarget, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.RaidTarget
 */
export declare const RaidTarget: RaidTarget$Type;
declare class ActionID$Type extends MessageType<ActionID> {
    constructor();
    create(value?: PartialMessage<ActionID>): ActionID;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ActionID): ActionID;
    internalBinaryWrite(message: ActionID, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.ActionID
 */
export declare const ActionID: ActionID$Type;
declare class Cooldown$Type extends MessageType<Cooldown> {
    constructor();
    create(value?: PartialMessage<Cooldown>): Cooldown;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Cooldown): Cooldown;
    internalBinaryWrite(message: Cooldown, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Cooldown
 */
export declare const Cooldown: Cooldown$Type;
declare class Cooldowns$Type extends MessageType<Cooldowns> {
    constructor();
    create(value?: PartialMessage<Cooldowns>): Cooldowns;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Cooldowns): Cooldowns;
    internalBinaryWrite(message: Cooldowns, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}
/**
 * @generated MessageType for protobuf message proto.Cooldowns
 */
export declare const Cooldowns: Cooldowns$Type;
export {};