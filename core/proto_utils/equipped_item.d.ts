import { Enchant } from '/tbc/core/proto/common.js';
import { Gem } from '/tbc/core/proto/common.js';
import { Item } from '/tbc/core/proto/common.js';
import { ItemSpec } from '/tbc/core/proto/common.js';
import { ActionId } from './action_id.js';
import { Stats } from './stats.js';
export declare function getWowheadItemId(item: Item): number;
export declare function getWeaponDPS(item: Item): number;
export declare function computeItemEP(item: Item, epWeights: Stats): number;
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
    private withGemHelper;
    /**
     * Returns a new EquippedItem with the given gem socketed.
       *
       * Also ensures validity of the item on its own. Currently this just means enforcing unique gems.
     */
    withGem(gem: Gem | null, socketIdx: number): EquippedItem;
    removeGemsWithId(gemId: number): EquippedItem;
    asActionId(): ActionId;
    asSpec(): ItemSpec;
}
