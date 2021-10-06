import { ItemSlot } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { equalsOrBothNull } from '/tbc/core/utils.js';
import { getEnumValues } from '/tbc/core/utils.js';

import { EquippedItem } from './equipped_item.js';

type InternalGear = Record<ItemSlot, EquippedItem | null>;

/**
 * Represents a full gear set, including items/enchants/gems for every slot.
 *
 * This is an immutable type.
 */
export class Gear {
  private readonly gear: InternalGear;

  constructor(gear: Partial<InternalGear>) {
		getEnumValues(ItemSlot).forEach(slot => {
      if (!gear[slot as ItemSlot])
        gear[slot as ItemSlot] = null;
    });
    this.gear = gear as InternalGear;
  }

  equals(other: Gear): boolean {
    return this.asArray().every((thisItem, slot) => equalsOrBothNull(thisItem, other.getEquippedItem(slot), (a, b) => a.equals(b)));
  }

  withEquippedItem(newSlot: ItemSlot, newItem: EquippedItem | null): Gear {
    const newInternalGear: Partial<InternalGear> = {};
		getEnumValues(ItemSlot).forEach(slot => {
      if (Number(slot) == newSlot) {
        newInternalGear[Number(slot) as ItemSlot] = newItem;
      } else {
        newInternalGear[Number(slot) as ItemSlot] = this.getEquippedItem(Number(slot));
      }
    });
    return new Gear(newInternalGear);
  }

  getEquippedItem(slot: ItemSlot): EquippedItem | null {
    return this.gear[slot];
  }

  asArray(): Array<EquippedItem | null> {
    return Object.values(this.gear);
  }

  asSpec(): EquipmentSpec {
    return EquipmentSpec.create({
      items: this.asArray().map(ei => ei ? ei.asSpec() : ItemSpec.create()),
    });
  }
}
