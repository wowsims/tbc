import { intersection } from '/tbc/core/utils.js';

import { Player } from '/tbc/core/proto/api.js';
import { PlayerOptions } from '/tbc/core/proto/api.js';
import { Class } from '/tbc/core/proto/common.js';
import { Enchant } from '/tbc/core/proto/common.js';
import { Gem } from '/tbc/core/proto/common.js';
import { GemColor } from '/tbc/core/proto/common.js';
import { HandType } from '/tbc/core/proto/common.js';
import { ItemSlot } from '/tbc/core/proto/common.js';
import { ItemType } from '/tbc/core/proto/common.js';
import { Item } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { RangedWeaponType } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { WeaponType } from '/tbc/core/proto/common.js';

import { BalanceDruid, BalanceDruid_Rotation as BalanceDruidRotation, DruidTalents, BalanceDruid_Options as BalanceDruidOptions} from '/tbc/core/proto/druid.js';
import { ElementalShaman, ElementalShaman_Rotation as ElementalShamanRotation, ShamanTalents, ElementalShaman_Options as ElementalShamanOptions } from '/tbc/core/proto/shaman.js';
import { Mage, Mage_Rotation as MageRotation, MageTalents, Mage_Options as MageOptions } from '/tbc/core/proto/mage.js';
import { ShadowPriest, ShadowPriest_Rotation as ShadowPriestRotation, PriestTalents, ShadowPriest_Options as ShadowPriestOptions } from '/tbc/core/proto/priest.js';
import { Warlock, Warlock_Rotation as WarlockRotation, WarlockTalents, Warlock_Options as WarlockOptions } from '/tbc/core/proto/warlock.js';

export type DruidSpecs = Spec.SpecBalanceDruid;
export type MageSpecs = Spec.SpecMage;
export type PriestSpecs = Spec.SpecShadowPriest;
export type ShamanSpecs = Spec.SpecElementalShaman;
export type WarlockSpecs = Spec.SpecWarlock;

export type RotationUnion =
		BalanceDruidRotation |
		ElementalShamanRotation |
		MageRotation |
		ShadowPriestRotation |
		WarlockRotation;
export type SpecRotation<T extends Spec> =
		T extends Spec.SpecBalanceDruid ? BalanceDruidRotation :
		T extends Spec.SpecElementalShaman ? ElementalShamanRotation :
		T extends Spec.SpecMage ? MageRotation :
		T extends Spec.SpecShadowPriest ? ShadowPriestRotation :
		T extends Spec.SpecWarlock ? WarlockRotation :
		ElementalShamanRotation; // Should never reach this case

export type TalentsUnion =
		DruidTalents |
		MageTalents |
		PriestTalents |
		ShamanTalents |
		WarlockTalents;
export type SpecTalents<T extends Spec> =
		T extends Spec.SpecBalanceDruid ? DruidTalents :
		T extends Spec.SpecElementalShaman ? ShamanTalents :
		T extends Spec.SpecMage ? MageTalents :
		T extends Spec.SpecShadowPriest ? PriestTalents :
		T extends Spec.SpecWarlock ? WarlockTalents :
		ShamanTalents; // Should never reach this case

export type SpecOptionsUnion =
		BalanceDruidOptions |
		ElementalShamanOptions |
		MageOptions |
		ShadowPriestOptions |
		WarlockOptions;
export type SpecOptions<T extends Spec> =
		T extends Spec.SpecBalanceDruid ? BalanceDruidOptions :
		T extends Spec.SpecElementalShaman ? ElementalShamanOptions :
		T extends Spec.SpecMage ? MageOptions :
		T extends Spec.SpecShadowPriest ? ShadowPriestOptions :
		T extends Spec.SpecWarlock ? WarlockOptions :
		ElementalShamanOptions; // Should never reach this case

export type SpecProtoUnion =
		BalanceDruid |
		ElementalShaman |
		Mage |
		ShadowPriest |
		Warlock;
export type SpecProto<T extends Spec> =
		T extends Spec.SpecBalanceDruid ? BalanceDruid :
		T extends Spec.SpecElementalShaman ? ElementalShaman :
		T extends Spec.SpecMage ? Mage :
		T extends Spec.SpecShadowPriest ? ShadowPriest :
		T extends Spec.SpecWarlock ? Warlock :
		ElementalShaman; // Should never reach this case

export type SpecTypeFunctions<SpecType extends Spec> = {
  rotationCreate: () => SpecRotation<SpecType>;
  rotationEquals: (a: SpecRotation<SpecType>, b: SpecRotation<SpecType>) => boolean;
  rotationCopy: (a: SpecRotation<SpecType>) => SpecRotation<SpecType>;
  rotationToJson: (a: SpecRotation<SpecType>) => any;
  rotationFromJson: (obj: any) => SpecRotation<SpecType>;

  talentsCreate: () => SpecTalents<SpecType>;
  talentsEquals: (a: SpecTalents<SpecType>, b: SpecTalents<SpecType>) => boolean;
  talentsCopy: (a: SpecTalents<SpecType>) => SpecTalents<SpecType>;
  talentsToJson: (a: SpecTalents<SpecType>) => any;
  talentsFromJson: (obj: any) => SpecTalents<SpecType>;

  optionsCreate: () => SpecOptions<SpecType>;
  optionsEquals: (a: SpecOptions<SpecType>, b: SpecOptions<SpecType>) => boolean;
  optionsCopy: (a: SpecOptions<SpecType>) => SpecOptions<SpecType>;
  optionsToJson: (a: SpecOptions<SpecType>) => any;
  optionsFromJson: (obj: any) => SpecOptions<SpecType>;
};

export const specTypeFunctions: Partial<Record<Spec, SpecTypeFunctions<any>>> = {
  [Spec.SpecBalanceDruid]: {
    rotationCreate: () => BalanceDruidRotation.create(),
    rotationEquals: (a, b) => BalanceDruidRotation.equals(a as BalanceDruidRotation, b as BalanceDruidRotation),
    rotationCopy: (a) => BalanceDruidRotation.clone(a as BalanceDruidRotation),
    rotationToJson: (a) => BalanceDruidRotation.toJson(a as BalanceDruidRotation),
    rotationFromJson: (obj) => BalanceDruidRotation.fromJson(obj),

    talentsCreate: () => DruidTalents.create(),
    talentsEquals: (a, b) => DruidTalents.equals(a as DruidTalents, b as DruidTalents),
    talentsCopy: (a) => DruidTalents.clone(a as DruidTalents),
    talentsToJson: (a) => DruidTalents.toJson(a as DruidTalents),
    talentsFromJson: (obj) => DruidTalents.fromJson(obj),

    optionsCreate: () => BalanceDruidOptions.create(),
    optionsEquals: (a, b) => BalanceDruidOptions.equals(a as BalanceDruidOptions, b as BalanceDruidOptions),
    optionsCopy: (a) => BalanceDruidOptions.clone(a as BalanceDruidOptions),
    optionsToJson: (a) => BalanceDruidOptions.toJson(a as BalanceDruidOptions),
    optionsFromJson: (obj) => BalanceDruidOptions.fromJson(obj),
  },
  [Spec.SpecElementalShaman]: {
    rotationCreate: () => ElementalShamanRotation.create(),
    rotationEquals: (a, b) => ElementalShamanRotation.equals(a as ElementalShamanRotation, b as ElementalShamanRotation),
    rotationCopy: (a) => ElementalShamanRotation.clone(a as ElementalShamanRotation),
    rotationToJson: (a) => ElementalShamanRotation.toJson(a as ElementalShamanRotation),
    rotationFromJson: (obj) => ElementalShamanRotation.fromJson(obj),

    talentsCreate: () => ShamanTalents.create(),
    talentsEquals: (a, b) => ShamanTalents.equals(a as ShamanTalents, b as ShamanTalents),
    talentsCopy: (a) => ShamanTalents.clone(a as ShamanTalents),
    talentsToJson: (a) => ShamanTalents.toJson(a as ShamanTalents),
    talentsFromJson: (obj) => ShamanTalents.fromJson(obj),

    optionsCreate: () => ElementalShamanOptions.create(),
    optionsEquals: (a, b) => ElementalShamanOptions.equals(a as ElementalShamanOptions, b as ElementalShamanOptions),
    optionsCopy: (a) => ElementalShamanOptions.clone(a as ElementalShamanOptions),
    optionsToJson: (a) => ElementalShamanOptions.toJson(a as ElementalShamanOptions),
    optionsFromJson: (obj) => ElementalShamanOptions.fromJson(obj),
  },
  [Spec.SpecMage]: {
    rotationCreate: () => MageRotation.create(),
    rotationEquals: (a, b) => MageRotation.equals(a as MageRotation, b as MageRotation),
    rotationCopy: (a) => MageRotation.clone(a as MageRotation),
    rotationToJson: (a) => MageRotation.toJson(a as MageRotation),
    rotationFromJson: (obj) => MageRotation.fromJson(obj),

    talentsCreate: () => MageTalents.create(),
    talentsEquals: (a, b) => MageTalents.equals(a as MageTalents, b as MageTalents),
    talentsCopy: (a) => MageTalents.clone(a as MageTalents),
    talentsToJson: (a) => MageTalents.toJson(a as MageTalents),
    talentsFromJson: (obj) => MageTalents.fromJson(obj),

    optionsCreate: () => MageOptions.create(),
    optionsEquals: (a, b) => MageOptions.equals(a as MageOptions, b as MageOptions),
    optionsCopy: (a) => MageOptions.clone(a as MageOptions),
    optionsToJson: (a) => MageOptions.toJson(a as MageOptions),
    optionsFromJson: (obj) => MageOptions.fromJson(obj),
  },
  [Spec.SpecShadowPriest]: {
    rotationCreate: () => ShadowPriestRotation.create(),
    rotationEquals: (a, b) => ShadowPriestRotation.equals(a as ShadowPriestRotation, b as ShadowPriestRotation),
    rotationCopy: (a) => ShadowPriestRotation.clone(a as ShadowPriestRotation),
    rotationToJson: (a) => ShadowPriestRotation.toJson(a as ShadowPriestRotation),
    rotationFromJson: (obj) => ShadowPriestRotation.fromJson(obj),

    talentsCreate: () => PriestTalents.create(),
    talentsEquals: (a, b) => PriestTalents.equals(a as PriestTalents, b as PriestTalents),
    talentsCopy: (a) => PriestTalents.clone(a as PriestTalents),
    talentsToJson: (a) => PriestTalents.toJson(a as PriestTalents),
    talentsFromJson: (obj) => PriestTalents.fromJson(obj),

    optionsCreate: () => ShadowPriestOptions.create(),
    optionsEquals: (a, b) => ShadowPriestOptions.equals(a as ShadowPriestOptions, b as ShadowPriestOptions),
    optionsCopy: (a) => ShadowPriestOptions.clone(a as ShadowPriestOptions),
    optionsToJson: (a) => ShadowPriestOptions.toJson(a as ShadowPriestOptions),
    optionsFromJson: (obj) => ShadowPriestOptions.fromJson(obj),
  },
  [Spec.SpecWarlock]: {
    rotationCreate: () => WarlockRotation.create(),
    rotationEquals: (a, b) => WarlockRotation.equals(a as WarlockRotation, b as WarlockRotation),
    rotationCopy: (a) => WarlockRotation.clone(a as WarlockRotation),
    rotationToJson: (a) => WarlockRotation.toJson(a as WarlockRotation),
    rotationFromJson: (obj) => WarlockRotation.fromJson(obj),

    talentsCreate: () => WarlockTalents.create(),
    talentsEquals: (a, b) => WarlockTalents.equals(a as WarlockTalents, b as WarlockTalents),
    talentsCopy: (a) => WarlockTalents.clone(a as WarlockTalents),
    talentsToJson: (a) => WarlockTalents.toJson(a as WarlockTalents),
    talentsFromJson: (obj) => WarlockTalents.fromJson(obj),

    optionsCreate: () => WarlockOptions.create(),
    optionsEquals: (a, b) => WarlockOptions.equals(a as WarlockOptions, b as WarlockOptions),
    optionsCopy: (a) => WarlockOptions.clone(a as WarlockOptions),
    optionsToJson: (a) => WarlockOptions.toJson(a as WarlockOptions),
    optionsFromJson: (obj) => WarlockOptions.fromJson(obj),
  },
};

export const specToClass: Record<Spec, Class> = {
  [Spec.SpecBalanceDruid]: Class.ClassDruid,
  [Spec.SpecElementalShaman]: Class.ClassShaman,
  [Spec.SpecMage]: Class.ClassMage,
  [Spec.SpecShadowPriest]: Class.ClassPriest,
  [Spec.SpecWarlock]: Class.ClassWarlock,
};

const druidRaces = [
    Race.RaceNightElf,
    Race.RaceTauren,
];
const mageRaces = [
    Race.RaceBloodElf,
    Race.RaceDraenei,
    Race.RaceGnome,
    Race.RaceHuman,
    Race.RaceTroll10,
    Race.RaceTroll30,
    Race.RaceUndead,
];
const priestRaces = [
    Race.RaceBloodElf,
    Race.RaceDraenei,
    Race.RaceDwarf,
    Race.RaceHuman,
    Race.RaceNightElf,
    Race.RaceOrc,
    Race.RaceTroll10,
    Race.RaceTroll30,
    Race.RaceUndead,
];
const shamanRaces = [
    Race.RaceDraenei,
    Race.RaceOrc,
    Race.RaceTauren,
    Race.RaceTroll10,
    Race.RaceTroll30,
];
const warlockRaces = [
    Race.RaceBloodElf,
    Race.RaceGnome,
    Race.RaceHuman,
    Race.RaceOrc,
    Race.RaceUndead,
];

export const specToEligibleRaces: Record<Spec, Array<Race>> = {
  [Spec.SpecBalanceDruid]: druidRaces,
  [Spec.SpecElementalShaman]: shamanRaces,
  [Spec.SpecMage]: mageRaces,
  [Spec.SpecShadowPriest]: priestRaces,
  [Spec.SpecWarlock]: warlockRaces,
};

// Prefixes used for storing browser data for each site. Even if a Spec is
// renamed, DO NOT change these values or people will lose their saved data.
export const specToLocalStorageKey: Record<Spec, string> = {
  [Spec.SpecBalanceDruid]: '__balance_druid',
  [Spec.SpecElementalShaman]: '__elemental_shaman',
  [Spec.SpecMage]: '__mage',
  [Spec.SpecShadowPriest]: '__shadow_priest',
  [Spec.SpecWarlock]: '__warlock',
};

// Returns a copy of playerOptions, with the class field set.
export function withSpecProto<SpecType extends Spec>(
    playerOptions: PlayerOptions,
    rotation: SpecRotation<SpecType>,
    talents: SpecTalents<SpecType>,
    specOptions: SpecOptions<SpecType>): PlayerOptions {
  const copy = PlayerOptions.clone(playerOptions);
  if (BalanceDruidRotation.is(rotation)) {
		copy.class = Class.ClassDruid;
    copy.spec = {
      oneofKind: 'balanceDruid',
      balanceDruid: BalanceDruid.create({
        rotation: rotation,
        talents: talents as DruidTalents,
        options: specOptions as BalanceDruidOptions,
      }),
    };
  } else if (ElementalShamanRotation.is(rotation)) {
		copy.class = Class.ClassShaman;
    copy.spec = {
      oneofKind: 'elementalShaman',
      elementalShaman: ElementalShaman.create({
        rotation: rotation,
        talents: talents as ShamanTalents,
        options: specOptions as ElementalShamanOptions,
      }),
    };
  } else if (MageRotation.is(rotation)) {
		copy.class = Class.ClassMage;
    copy.spec = {
      oneofKind: 'mage',
      mage: Mage.create({
        rotation: rotation,
        talents: talents as MageTalents,
        options: specOptions as MageOptions,
      }),
    };
  } else if (ShadowPriestRotation.is(rotation)) {
		copy.class = Class.ClassPriest;
    copy.spec = {
      oneofKind: 'shadowPriest',
      shadowPriest: ShadowPriest.create({
        rotation: rotation,
        talents: talents as PriestTalents,
        options: specOptions as ShadowPriestOptions,
      }),
    };
  } else if (WarlockRotation.is(rotation)) {
		copy.class = Class.ClassWarlock;
    copy.spec = {
      oneofKind: 'warlock',
      warlock: Warlock.create({
        rotation: rotation,
        talents: talents as WarlockTalents,
        options: specOptions as WarlockOptions,
      }),
    };
  } else {
    throw new Error('Unrecognized talents with options: ' + PlayerOptions.toJsonString(playerOptions));
  }
  return copy;
}

const itemTypeToSlotsMap: Partial<Record<ItemType, Array<ItemSlot>>> = {
  [ItemType.ItemTypeUnknown]: [],
  [ItemType.ItemTypeHead]: [ItemSlot.ItemSlotHead],
  [ItemType.ItemTypeNeck]: [ItemSlot.ItemSlotNeck],
  [ItemType.ItemTypeShoulder]: [ItemSlot.ItemSlotShoulder],
  [ItemType.ItemTypeBack]: [ItemSlot.ItemSlotBack],
  [ItemType.ItemTypeChest]: [ItemSlot.ItemSlotChest],
  [ItemType.ItemTypeWrist]: [ItemSlot.ItemSlotWrist],
  [ItemType.ItemTypeHands]: [ItemSlot.ItemSlotHands],
  [ItemType.ItemTypeWaist]: [ItemSlot.ItemSlotWaist],
  [ItemType.ItemTypeLegs]: [ItemSlot.ItemSlotLegs],
  [ItemType.ItemTypeFeet]: [ItemSlot.ItemSlotFeet],
  [ItemType.ItemTypeFinger]: [ItemSlot.ItemSlotFinger1, ItemSlot.ItemSlotFinger2],
  [ItemType.ItemTypeTrinket]: [ItemSlot.ItemSlotTrinket1, ItemSlot.ItemSlotTrinket2],
  [ItemType.ItemTypeRanged]: [ItemSlot.ItemSlotRanged],
};

export function getEligibleItemSlots(item: Item): Array<ItemSlot> {
  if (itemTypeToSlotsMap[item.type]) {
    return itemTypeToSlotsMap[item.type]!;
  }

  if (item.type == ItemType.ItemTypeWeapon) {
    if ([HandType.HandTypeMainHand, HandType.HandTypeTwoHand].includes(item.handType)) {
      return [ItemSlot.ItemSlotMainHand];
    } else if (item.handType == HandType.HandTypeOffHand) {
      return [ItemSlot.ItemSlotOffHand];
    } else {
      return [ItemSlot.ItemSlotMainHand, ItemSlot.ItemSlotOffHand];
    }
  }

  // Should never reach here
  throw new Error('Could not find item slots for item: ' + Item.toJsonString(item));
};

// Returns whether the given main-hand and off-hand items can be worn at the
// same time.
export function validWeaponCombo(mainHand: Item | null | undefined, offHand: Item | null | undefined): boolean {
	if (mainHand == null || offHand == null) {
		return true;
	}

	if (mainHand.handType == HandType.HandTypeTwoHand) {
		return false;
	}

	return true;
}

// Returns all item slots to which the enchant might be applied.
// 
// Note that this alone is not enough; some items have further restrictions,
// e.g. some weapon enchants may only be applied to 2H weapons.
export function getEligibleEnchantSlots(enchant: Enchant): Array<ItemSlot> {
  if (itemTypeToSlotsMap[enchant.type]) {
    return itemTypeToSlotsMap[enchant.type]!;
  }

  if (enchant.type == ItemType.ItemTypeWeapon) {
    return [ItemSlot.ItemSlotMainHand, ItemSlot.ItemSlotOffHand];
  }

  // Should never reach here
  throw new Error('Could not find item slots for enchant: ' + Enchant.toJsonString(enchant));
};

export function enchantAppliesToItem(enchant: Enchant, item: Item): boolean {
  const sharedSlots = intersection(getEligibleEnchantSlots(enchant), getEligibleItemSlots(item));
  if (sharedSlots.length == 0)
    return false;

  if (sharedSlots.includes(ItemSlot.ItemSlotMainHand)) {
    if (enchant.twoHandedOnly && item.handType != HandType.HandTypeTwoHand)
      return false;
  }

  if (sharedSlots.includes(ItemSlot.ItemSlotOffHand)) {
    if (enchant.shieldOnly && item.weaponType != WeaponType.WeaponTypeShield)
      return false;
  }

  if (sharedSlots.includes(ItemSlot.ItemSlotRanged)) {
    if (![
          RangedWeaponType.RangedWeaponTypeBow,
          RangedWeaponType.RangedWeaponTypeCrossbow,
          RangedWeaponType.RangedWeaponTypeGun,
        ].includes(item.rangedWeaponType))
      return false;
  }

  return true;
};

const socketToMatchingColors = new Map<GemColor, Array<GemColor>>();
socketToMatchingColors.set(GemColor.GemColorMeta,   [GemColor.GemColorMeta]);
socketToMatchingColors.set(GemColor.GemColorBlue,   [GemColor.GemColorBlue, GemColor.GemColorPurple, GemColor.GemColorGreen]);
socketToMatchingColors.set(GemColor.GemColorRed,    [GemColor.GemColorRed, GemColor.GemColorPurple, GemColor.GemColorOrange]);
socketToMatchingColors.set(GemColor.GemColorYellow, [GemColor.GemColorYellow, GemColor.GemColorOrange, GemColor.GemColorGreen]);

// Whether the gem matches the given socket color, for the purposes of gaining the socket bonuses.
export function gemMatchesSocket(gem: Gem, socketColor: GemColor) {
  return socketToMatchingColors.has(socketColor) && socketToMatchingColors.get(socketColor)!.includes(gem.color);
}

// Whether the gem is capable of slotting into a socket of the given color.
export function gemEligibleForSocket(gem: Gem, socketColor: GemColor) {
  return (gem.color == GemColor.GemColorMeta) == (socketColor == GemColor.GemColorMeta);
}
