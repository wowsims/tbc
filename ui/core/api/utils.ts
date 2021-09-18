import { intersection } from '../utils';

import { Class } from './newapi';
import { Enchant } from './newapi';
import { Gem } from './newapi';
import { GemColor } from './newapi';
import { HandType } from './newapi';
import { ItemSlot } from './newapi';
import { ItemType } from './newapi';
import { Item } from './newapi';
import { Race } from './newapi';
import { RangedWeaponType } from './newapi';
import { Druid, Druid_DruidAgent as DruidAgent, Druid_DruidTalents as DruidTalents, Druid_DruidOptions as DruidOptions} from './newapi';
import { Shaman, Shaman_ShamanAgent as ShamanAgent, Shaman_ShamanTalents as ShamanTalents, Shaman_ShamanOptions as ShamanOptions } from './newapi';
import { Spec } from './newapi';
import { WeaponType } from './newapi';

export type AgentUnion = DruidAgent | ShamanAgent;
export type ClassAgent<T extends Class> = T extends Class.ClassDruid ? DruidAgent : ShamanAgent;

export type TalentsUnion = DruidTalents | ShamanTalents;
export type ClassTalents<T extends Class> = T extends Class.ClassDruid ? DruidTalents: ShamanTalents;

export type ClassOptionsUnion = DruidOptions | ShamanOptions;
export type ClassOptions<T extends Class> = T extends Class.ClassDruid ? DruidOptions : ShamanOptions;

export type ClassProtoUnion = Druid | Shaman;
export type ClassProto<T extends Class> = T extends Class.ClassDruid ? Druid : Shaman;

export type ClassTypeFunctions<ClassType extends Class> = {
  agentCreate: () => ClassAgent<ClassType>;
  agentEquals: (a: ClassAgent<ClassType>, b: ClassAgent<ClassType>) => boolean;
  agentCopy: (a: ClassAgent<ClassType>) => ClassAgent<ClassType>;
  agentToJson: (a: ClassAgent<ClassType>) => any;
  agentFromJson: (obj: any) => ClassAgent<ClassType>;

  talentsCreate: () => ClassTalents<ClassType>;
  talentsEquals: (a: ClassTalents<ClassType>, b: ClassTalents<ClassType>) => boolean;
  talentsCopy: (a: ClassTalents<ClassType>) => ClassTalents<ClassType>;
  talentsToJson: (a: ClassTalents<ClassType>) => any;
  talentsFromJson: (obj: any) => ClassTalents<ClassType>;

  optionsCreate: () => ClassOptions<ClassType>;
  optionsEquals: (a: ClassOptions<ClassType>, b: ClassOptions<ClassType>) => boolean;
  optionsCopy: (a: ClassOptions<ClassType>) => ClassOptions<ClassType>;
  optionsToJson: (a: ClassOptions<ClassType>) => any;
  optionsFromJson: (obj: any) => ClassOptions<ClassType>;
};

export const classTypeFunctions: Partial<Record<Class, ClassTypeFunctions<any>>> = {
  [Class.ClassDruid]: {
    agentCreate: () => DruidAgent.create(),
    agentEquals: (a, b) => DruidAgent.equals(a as DruidAgent, b as DruidAgent),
    agentCopy: (a) => DruidAgent.clone(a as DruidAgent),
    agentToJson: (a) => DruidAgent.toJson(a as DruidAgent),
    agentFromJson: (obj) => DruidAgent.fromJson(obj),

    talentsCreate: () => DruidTalents.create(),
    talentsEquals: (a, b) => DruidTalents.equals(a as DruidTalents, b as DruidTalents),
    talentsCopy: (a) => DruidTalents.clone(a as DruidTalents),
    talentsToJson: (a) => DruidTalents.toJson(a as DruidTalents),
    talentsFromJson: (obj) => DruidTalents.fromJson(obj),

    optionsCreate: () => DruidOptions.create(),
    optionsEquals: (a, b) => DruidOptions.equals(a as DruidOptions, b as DruidOptions),
    optionsCopy: (a) => DruidOptions.clone(a as DruidOptions),
    optionsToJson: (a) => DruidOptions.toJson(a as DruidOptions),
    optionsFromJson: (obj) => DruidOptions.fromJson(obj),
  },
  [Class.ClassShaman]: {
    agentCreate: () => ShamanAgent.create(),
    agentEquals: (a, b) => ShamanAgent.equals(a as ShamanAgent, b as ShamanAgent),
    agentCopy: (a) => ShamanAgent.clone(a as ShamanAgent),
    agentToJson: (a) => ShamanAgent.toJson(a as ShamanAgent),
    agentFromJson: (obj) => ShamanAgent.fromJson(obj),

    talentsCreate: () => ShamanTalents.create(),
    talentsEquals: (a, b) => ShamanTalents.equals(a as ShamanTalents, b as ShamanTalents),
    talentsCopy: (a) => ShamanTalents.clone(a as ShamanTalents),
    talentsToJson: (a) => ShamanTalents.toJson(a as ShamanTalents),
    talentsFromJson: (obj) => ShamanTalents.fromJson(obj),

    optionsCreate: () => ShamanOptions.create(),
    optionsEquals: (a, b) => ShamanOptions.equals(a as ShamanOptions, b as ShamanOptions),
    optionsCopy: (a) => ShamanOptions.clone(a as ShamanOptions),
    optionsToJson: (a) => ShamanOptions.toJson(a as ShamanOptions),
    optionsFromJson: (obj) => ShamanOptions.fromJson(obj),
  },
};

export const specToClass: Record<Spec, Class> = {
  [Spec.ElementalShaman]: Class.ClassShaman,
};

const shamanRaces = [
    Race.RaceDraenei,
    Race.RaceOrc,
    Race.RaceTauren,
    Race.RaceTroll10,
    Race.RaceTroll30,
];

export const specToEligibleRaces: Record<Spec, Array<Race>> = {
  [Spec.ElementalShaman]: shamanRaces,
};

const itemTypeToSlotsMap: Partial<Record<ItemType, Array<ItemSlot>>> = {
  [ItemType.ItemTypeUnknown]: [],
  [ItemType.ItemTypeHead]: [ItemSlot.ItemSlotHead],
  [ItemType.ItemTypeNeck]: [ItemSlot.ItemSlotNeck],
  [ItemType.ItemTypeShoulder]: [ItemSlot.ItemSlotShoulder],
  [ItemType.ItemTypeBack]: [ItemSlot.ItemSlotBack],
  [ItemType.ItemTypeChest]: [ItemSlot.ItemSlotChest],
  [ItemType.ItemTypeWrist]: [ItemSlot.ItemSlotWrist],
  [ItemType.ItemTypeHands]: [ItemSlot.ItemSlotHead],
  [ItemType.ItemTypeWaist]: [ItemSlot.ItemSlotHead],
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

/**
 * Returns all item slots to which the enchant might be applied.
 *
 * Note that this alone is not enough; some items have further restrictions,
 * e.g. some weapon enchants may only be applied to 2H weapons.
 */
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
