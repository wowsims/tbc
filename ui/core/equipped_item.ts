import { Enchant } from './api/newapi';
import { Gem } from './api/newapi';
import { Item } from './api/newapi';
import { ItemSlot } from './api/newapi';
import { EnchantAppliesToItem } from './api/utils';
import { GemEligibleForSocket } from './api/utils';
import { GemMatchesSocket } from './api/utils';

/**
 * Represents an equipped item along with enchants/gems attached to it.
 *
 * This is an immutable type.
 */
export class EquippedItem {
  readonly item: Item;
  readonly enchant: Enchant | null;
  readonly gems: Array<Gem | null>;

  constructor(item: Item, enchant?: Enchant | null, gems?: Array<Gem | null>) {
    this.item = item;
    this.enchant = enchant || null;
    this.gems = gems || [];

    // Fill gems with null so we always have the same number of gems as gem slots.
    if (this.gems.length < item.gemSockets.length) {
      this.gems = this.gems.concat(new Array(item.gemSockets.length - this.gems.length).fill(null));
    }
  }

  equals(other: EquippedItem) {
    if (Item.equals(this.item, other.item))
      return false;

    if ((this.enchant == null) != (other.enchant == null))
      return false;

    if (this.enchant && other.enchant && !Enchant.equals(this.enchant, other.enchant))
      return false;

    if (this.gems.length != other.gems.length)
      return false;

    for (let i = 0; i < this.gems.length; i++) {
      if ((this.gems[i] == null) != (other.gems[i] == null))
        return false;

      if (this.gems[i] && other.gems[i] && !Gem.equals(this.gems[i]!, other.gems[i]!))
        return false;
    }

    return true;
  }

  /**
   * Replaces the item and tries to keep the existing enchants/gems if possible.
   */
  withItem(item: Item): EquippedItem {
    let newEnchant = null;
    if (this.enchant && EnchantAppliesToItem(this.enchant, item))
      newEnchant = this.enchant;
    
    // Reorganize gems to match as many colors in the new item as possible.
    const newGems = new Array(item.gemSockets.length).fill(null);
    this.gems.filter(gem => gem != null).forEach(gem => {
      const firstMatchingIndex = item.gemSockets.findIndex((socketColor, socketIdx) => !newGems[socketIdx] && GemMatchesSocket(gem!, socketColor));
      const firstEligibleIndex = item.gemSockets.findIndex((socketColor, socketIdx) => !newGems[socketIdx] && GemEligibleForSocket(gem!, socketColor));
      if (firstMatchingIndex != -1) {
        newGems[firstMatchingIndex] = gem;
      } else if (firstEligibleIndex != -1) {
        newGems[firstEligibleIndex] = gem;
      }
    });

    return new EquippedItem(item, newEnchant, newGems);
  }

  /**
   * Returns a new EquippedItem with the given enchant applied.
   */
  withEnchant(enchant: Enchant | null): EquippedItem {
    return new EquippedItem(this.item, enchant, this.gems);
  }

  /**
   * Returns a new EquippedItem with the given gem socketed.
   */
  withGem(gem: Gem | null, socketIdx: number): EquippedItem {
    if (this.gems.length <= socketIdx) {
      throw new Error('No gem socket with index ' + socketIdx);
    }

    const newGems = this.gems.slice();
    newGems[socketIdx] = gem;

    return new EquippedItem(this.item, this.enchant, newGems);
  }
};
