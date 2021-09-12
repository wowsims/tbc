import { Enchant } from './api/newapi';
import { Gem } from './api/newapi';
import { Item } from './api/newapi';
import { ItemSlot } from './api/newapi';

/**
 * Represents an equipped item along with enchants/gems attached to it.
 *
 * This is an immutable type.
 */
export class EquippedItem {
  readonly item: Item;
  readonly enchant: Enchant | null;
  readonly gems: Array<Gem | null>;

  constructor(item: Item, enchant?: Enchant, gems?: Array<Gem>) {
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

  setWowheadData(elem: HTMLElement) {
    let data = '';
    if (this.gems.length > 0) {
      data += 'gems=' + this.gems.map(gem => gem ? gem.id : 0).join(':');
    }
    if (this.enchant != null) {
      if (data.length > 0)
        data += '&';
      data += 'ench=' + this.enchant.effectId;
    }

    elem.setAttribute('data-wowhead', data);
  }
};
