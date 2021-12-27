import { GemColor } from '/tbc/core/proto/common.js';
import { ItemSlot } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { equalsOrBothNull } from '/tbc/core/utils.js';
import { getEnumValues } from '/tbc/core/utils.js';

import { EquippedItem } from './equipped_item.js';
import { validWeaponCombo } from './utils.js';

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

	/**
	 * Returns a new Gear set with the item equipped.
	 *
	 * Checks for validity and removes/exchanges items/gems as needed.
	 */
  withEquippedItem(newSlot: ItemSlot, newItem: EquippedItem | null): Gear {
		// Create a new identical set of gear
    const newInternalGear: Partial<InternalGear> = {};
		getEnumValues(ItemSlot).map(slot => Number(slot) as ItemSlot).forEach(slot => {
			newInternalGear[slot] = this.getEquippedItem(slot);
    });

		if (newItem) {
			// If the new item has unique gems, remove matching.
			newItem.gems
			.filter(gem => gem?.unique)
			.forEach(gem => {
				getEnumValues(ItemSlot).map(slot => Number(slot) as ItemSlot).forEach(slot => {
					newInternalGear[slot] = newInternalGear[slot]?.removeGemsWithId(gem!.id) || null;
				});
			});

			// If the new item is unique, remove matching items.
			if (newItem.item.unique) {
				getEnumValues(ItemSlot).map(slot => Number(slot) as ItemSlot).forEach(slot => {
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
			} else {
				newInternalGear[ItemSlot.ItemSlotMainHand] = null;
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

	metaGemEquipped(): boolean {
		const headItem = this.gear[ItemSlot.ItemSlotHead];
		return headItem?.gems.some(gem => gem.color == GemColor.GemColorMeta) || false;
	}
}
