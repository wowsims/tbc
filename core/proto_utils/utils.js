import { intersection } from '/tbc/core/utils.js';
import { Class } from '/tbc/core/proto/common.js';
import { Enchant } from '/tbc/core/proto/common.js';
import { GemColor } from '/tbc/core/proto/common.js';
import { HandType } from '/tbc/core/proto/common.js';
import { ItemSlot } from '/tbc/core/proto/common.js';
import { ItemType } from '/tbc/core/proto/common.js';
import { Item } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { RangedWeaponType } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { WeaponType } from '/tbc/core/proto/common.js';
import { BalanceDruid_Agent as BalanceDruidAgent, DruidTalents, BalanceDruid_Options as BalanceDruidOptions } from '/tbc/core/proto/druid.js';
import { ElementalShaman_Agent as ElementalShamanAgent, ShamanTalents, ElementalShaman_Options as ElementalShamanOptions } from '/tbc/core/proto/shaman.js';
export const specTypeFunctions = {
    [Spec.SpecBalanceDruid]: {
        agentCreate: () => BalanceDruidAgent.create(),
        agentEquals: (a, b) => BalanceDruidAgent.equals(a, b),
        agentCopy: (a) => BalanceDruidAgent.clone(a),
        agentToJson: (a) => BalanceDruidAgent.toJson(a),
        agentFromJson: (obj) => BalanceDruidAgent.fromJson(obj),
        talentsCreate: () => DruidTalents.create(),
        talentsEquals: (a, b) => DruidTalents.equals(a, b),
        talentsCopy: (a) => DruidTalents.clone(a),
        talentsToJson: (a) => DruidTalents.toJson(a),
        talentsFromJson: (obj) => DruidTalents.fromJson(obj),
        optionsCreate: () => BalanceDruidOptions.create(),
        optionsEquals: (a, b) => BalanceDruidOptions.equals(a, b),
        optionsCopy: (a) => BalanceDruidOptions.clone(a),
        optionsToJson: (a) => BalanceDruidOptions.toJson(a),
        optionsFromJson: (obj) => BalanceDruidOptions.fromJson(obj),
    },
    [Spec.SpecElementalShaman]: {
        agentCreate: () => ElementalShamanAgent.create(),
        agentEquals: (a, b) => ElementalShamanAgent.equals(a, b),
        agentCopy: (a) => ElementalShamanAgent.clone(a),
        agentToJson: (a) => ElementalShamanAgent.toJson(a),
        agentFromJson: (obj) => ElementalShamanAgent.fromJson(obj),
        talentsCreate: () => ShamanTalents.create(),
        talentsEquals: (a, b) => ShamanTalents.equals(a, b),
        talentsCopy: (a) => ShamanTalents.clone(a),
        talentsToJson: (a) => ShamanTalents.toJson(a),
        talentsFromJson: (obj) => ShamanTalents.fromJson(obj),
        optionsCreate: () => ElementalShamanOptions.create(),
        optionsEquals: (a, b) => ElementalShamanOptions.equals(a, b),
        optionsCopy: (a) => ElementalShamanOptions.clone(a),
        optionsToJson: (a) => ElementalShamanOptions.toJson(a),
        optionsFromJson: (obj) => ElementalShamanOptions.fromJson(obj),
    },
};
export const specToClass = {
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
export const specToEligibleRaces = {
    [Spec.SpecBalanceDruid]: druidRaces,
    [Spec.SpecElementalShaman]: shamanRaces,
};
const itemTypeToSlotsMap = {
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
export function getEligibleItemSlots(item) {
    if (itemTypeToSlotsMap[item.type]) {
        return itemTypeToSlotsMap[item.type];
    }
    if (item.type == ItemType.ItemTypeWeapon) {
        if ([HandType.HandTypeMainHand, HandType.HandTypeTwoHand].includes(item.handType)) {
            return [ItemSlot.ItemSlotMainHand];
        }
        else if (item.handType == HandType.HandTypeOffHand) {
            return [ItemSlot.ItemSlotOffHand];
        }
        else {
            return [ItemSlot.ItemSlotMainHand, ItemSlot.ItemSlotOffHand];
        }
    }
    // Should never reach here
    throw new Error('Could not find item slots for item: ' + Item.toJsonString(item));
}
;
// Returns whether the given main-hand and off-hand items can be worn at the
// same time.
export function validWeaponCombo(mainHand, offHand) {
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
export function getEligibleEnchantSlots(enchant) {
    if (itemTypeToSlotsMap[enchant.type]) {
        return itemTypeToSlotsMap[enchant.type];
    }
    if (enchant.type == ItemType.ItemTypeWeapon) {
        return [ItemSlot.ItemSlotMainHand, ItemSlot.ItemSlotOffHand];
    }
    // Should never reach here
    throw new Error('Could not find item slots for enchant: ' + Enchant.toJsonString(enchant));
}
;
export function enchantAppliesToItem(enchant, item) {
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
}
;
const socketToMatchingColors = new Map();
socketToMatchingColors.set(GemColor.GemColorMeta, [GemColor.GemColorMeta]);
socketToMatchingColors.set(GemColor.GemColorBlue, [GemColor.GemColorBlue, GemColor.GemColorPurple, GemColor.GemColorGreen]);
socketToMatchingColors.set(GemColor.GemColorRed, [GemColor.GemColorRed, GemColor.GemColorPurple, GemColor.GemColorOrange]);
socketToMatchingColors.set(GemColor.GemColorYellow, [GemColor.GemColorYellow, GemColor.GemColorOrange, GemColor.GemColorGreen]);
// Whether the gem matches the given socket color, for the purposes of gaining the socket bonuses.
export function gemMatchesSocket(gem, socketColor) {
    return socketToMatchingColors.has(socketColor) && socketToMatchingColors.get(socketColor).includes(gem.color);
}
// Whether the gem is capable of slotting into a socket of the given color.
export function gemEligibleForSocket(gem, socketColor) {
    return (gem.color == GemColor.GemColorMeta) == (socketColor == GemColor.GemColorMeta);
}
