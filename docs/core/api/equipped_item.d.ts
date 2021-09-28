import { Enchant } from './common.js';
import { Gem } from './common.js';
import { Item } from './common.js';
import { ItemSpec } from './common.js';
/**
 * Represents an equipped item along with enchants/gems attached to it.
 *
 * This is an immutable type.
 */
export declare class EquippedItem {
    readonly _item: Item;
    readonly _enchant: Enchant | null;
    readonly _gems: Array<Gem | null>;
    constructor(item: Item, enchant?: Enchant | null, gems?: Array<Gem | null>);
    get item(): Item;
    get enchant(): Enchant | null;
    get gems(): Array<Gem | null>;
    equals(other: EquippedItem): boolean;
    /**
     * Replaces the item and tries to keep the existing enchants/gems if possible.
     */
    withItem(item: Item): EquippedItem;
    /**
     * Returns a new EquippedItem with the given enchant applied.
     */
    withEnchant(enchant: Enchant | null): EquippedItem;
    /**
     * Returns a new EquippedItem with the given gem socketed.
     */
    withGem(gem: Gem | null, socketIdx: number): EquippedItem;
    asSpec(): ItemSpec;
}
