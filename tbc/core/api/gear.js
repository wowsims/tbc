import { ItemSlot } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { equalsOrBothNull } from '/tbc/core/utils.js';
import { getEnumValues } from '/tbc/core/utils.js';
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
    withEquippedItem(newSlot, newItem) {
        const newInternalGear = {};
        getEnumValues(ItemSlot).forEach(slot => {
            if (Number(slot) == newSlot) {
                newInternalGear[Number(slot)] = newItem;
            }
            else {
                newInternalGear[Number(slot)] = this.getEquippedItem(Number(slot));
            }
        });
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
