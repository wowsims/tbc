import { intersection } from '../utils.js';

import { Class } from './common.js';
import { Enchant } from './common.js';
import { Gem } from './common.js';
import { GemColor } from './common.js';
import { HandType } from './common.js';
import { ItemSlot } from './common.js';
import { ItemType } from './common.js';
import { Item } from './common.js';
import { Race } from './common.js';
import { RangedWeaponType } from './common.js';
import { Spec } from './common.js';
import { WeaponType } from './common.js';

import { BalanceDruid, BalanceDruid_Agent as BalanceDruidAgent, DruidTalents, BalanceDruid_Options as BalanceDruidOptions} from './druid.js';
import { ElementalShaman, ElementalShaman_Agent as ElementalShamanAgent, ShamanTalents, ElementalShaman_Options as ElementalShamanOptions } from './shaman.js';

export type ShamanSpecs = Spec.SpecElementalShaman;

export type AgentUnion = BalanceDruidAgent | ElementalShamanAgent;
export type SpecAgent<T extends Spec> = T extends Spec.SpecBalanceDruid ? BalanceDruidAgent : ElementalShamanAgent;

export type TalentsUnion = DruidTalents | ShamanTalents;
export type SpecTalents<T extends Spec> = T extends Spec.SpecBalanceDruid ? DruidTalents: ShamanTalents;

export type SpecOptionsUnion = BalanceDruidOptions | ElementalShamanOptions;
export type SpecOptions<T extends Spec> = T extends Spec.SpecBalanceDruid ? BalanceDruidOptions : ElementalShamanOptions;

export type SpecProtoUnion = BalanceDruid | ElementalShaman;
export type SpecProto<T extends Spec> = T extends Spec.SpecBalanceDruid ? BalanceDruid : ElementalShaman;

export type SpecTypeFunctions<SpecType extends Spec> = {
  agentCreate: () => SpecAgent<SpecType>;
  agentEquals: (a: SpecAgent<SpecType>, b: SpecAgent<SpecType>) => boolean;
  agentCopy: (a: SpecAgent<SpecType>) => SpecAgent<SpecType>;
  agentToJson: (a: SpecAgent<SpecType>) => any;
  agentFromJson: (obj: any) => SpecAgent<SpecType>;

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
    agentCreate: () => BalanceDruidAgent.create(),
    agentEquals: (a, b) => BalanceDruidAgent.equals(a as BalanceDruidAgent, b as BalanceDruidAgent),
    agentCopy: (a) => BalanceDruidAgent.clone(a as BalanceDruidAgent),
    agentToJson: (a) => BalanceDruidAgent.toJson(a as BalanceDruidAgent),
    agentFromJson: (obj) => BalanceDruidAgent.fromJson(obj),

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
    agentCreate: () => ElementalShamanAgent.create(),
    agentEquals: (a, b) => ElementalShamanAgent.equals(a as ElementalShamanAgent, b as ElementalShamanAgent),
    agentCopy: (a) => ElementalShamanAgent.clone(a as ElementalShamanAgent),
    agentToJson: (a) => ElementalShamanAgent.toJson(a as ElementalShamanAgent),
    agentFromJson: (obj) => ElementalShamanAgent.fromJson(obj),

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
};

export const specToClass: Record<Spec, Class> = {
  [Spec.SpecBalanceDruid]: Class.ClassDruid,
  [Spec.SpecElementalShaman]: Class.ClassShaman,
};

const druidRaces = [
    Race.RaceNightElf,
    Race.RaceTauren,
];
const shamanRaces = [
    Race.RaceDraenei,
    Race.RaceOrc,
    Race.RaceTauren,
    Race.RaceTroll10,
    Race.RaceTroll30,
];

export const specToEligibleRaces: Record<Spec, Array<Race>> = {
  [Spec.SpecBalanceDruid]: druidRaces,
  [Spec.SpecElementalShaman]: shamanRaces,
};

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
