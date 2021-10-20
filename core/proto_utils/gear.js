import { ItemSlot } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { equalsOrBothNull } from '/tbc/core/utils.js';
import { getEnumValues } from '/tbc/core/utils.js';
import { validWeaponCombo } from './utils.js';
/**
 * Represents a full gear set, including items/enchants/gems for every slot.
 *
 * This is an immutable type.
 */
export class Gear {
    constructor(gear) {
        getEnumValues(ItemSlot).forEach(slot => {
            if (!gear[slot])
                gear[slot] = null;
        });
        this.gear = gear;
    }
    equals(other) {
        return this.asArray().every((thisItem, slot) => equalsOrBothNull(thisItem, other.getEquippedItem(slot), (a, b) => a.equals(b)));
    }
    /**
     * Returns a new Gear set with the item equipped.
     *
     * Checks for validity and removes/exchanges items/gems as needed.
     */
    withEquippedItem(newSlot, newItem) {
        // Create a new identical set of gear
        const newInternalGear = {};
        getEnumValues(ItemSlot).map(slot => Number(slot)).forEach(slot => {
            newInternalGear[slot] = this.getEquippedItem(slot);
        });
        if (newItem) {
            // If the new item has unique gems, remove matching.
            newItem.gems
                .filter(gem => gem?.unique)
                .forEach(gem => {
                getEnumValues(ItemSlot).map(slot => Number(slot)).forEach(slot => {
                    newInternalGear[slot] = newInternalGear[slot]?.removeGemsWithId(gem.id) || null;
                });
            });
            // If the new item is unique, remove matching items.
            if (newItem.item.unique) {
                getEnumValues(ItemSlot).map(slot => Number(slot)).forEach(slot => {
                    if (newInternalGear[slot]?.item.id == newItem.item.id) {
                        newInternalGear[slot] = null;
                    }
                });
            }
        }
        // Actually assign the new item.
        newInternalGear[newSlot] = newItem;
        // Check for valid weapon combos.
        if (!validWeaponCombo(newInternalGear[ItemSlot.ItemSlotMainHand]?.item, newInternalGear[ItemSlot.ItemSlotOffHand]?.item)) {
            if (newSlot == ItemSlot.ItemSlotMainHand) {
                newInternalGear[ItemSlot.ItemSlotOffHand] = null;
            }
            else {
                newInternalGear[ItemSlot.ItemSlotMainHand] = null;
            }
        }
        return new Gear(newInternalGear);
    }
    getEquippedItem(slot) {
        return this.gear[slot];
    }
    asArray() {
        return Object.values(this.gear);
    }
    asSpec() {
        return EquipmentSpec.create({
            items: this.asArray().map(ei => ei ? ei.asSpec() : ItemSpec.create()),
        });
    }
}
