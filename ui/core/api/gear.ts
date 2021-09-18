import { equalsOrBothNull } from '../utils';

import { EquippedItem } from './equipped_item';
import { ItemSlot } from './newapi';
import { ItemSpec } from './newapi';
import { EquipmentSpec } from './newapi';

type InternalGear = Record<ItemSlot, EquippedItem | null>;

/**
 * Represents a full gear set, including items/enchants/gems for every slot.
 *
 * This is an immutable type.
 */
export class Gear {
  private readonly gear: InternalGear;

  constructor(gear: Partial<InternalGear>) {
    for (let slot in ItemSlot) {
      if (!gear[Number(slot) as ItemSlot])
        gear[Number(slot) as ItemSlot] = null;
    }
    this.gear = gear as InternalGear;
  }

  equals(other: Gear): boolean {
    return this.asArray().every((thisItem, slot) => equalsOrBothNull(thisItem, other.getEquippedItem(slot), (a, b) => a.equals(b)));
  }

  withEquippedItem(newSlot: ItemSlot, newItem: EquippedItem | null): Gear {
    const newInternalGear: Partial<InternalGear> = {};
    for (let slot in ItemSlot) {
      if (Number(slot) == newSlot) {
        newInternalGear[Number(slot) as ItemSlot] = newItem;
      } else {
        newInternalGear[Number(slot) as ItemSlot] = this.getEquippedItem(Number(slot));
      }
    }
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
