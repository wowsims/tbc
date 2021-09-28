import { EquippedItem } from './equipped_item.js';
import { ItemSlot } from './common.js';
import { EquipmentSpec } from './common.js';
declare type InternalGear = Record<ItemSlot, EquippedItem | null>;
/**
 * Represents a full gear set, including items/enchants/gems for every slot.
 *
 * This is an immutable type.
 */
export declare class Gear {
    private readonly gear;
    constructor(gear: Partial<InternalGear>);
    equals(other: Gear): boolean;
    withEquippedItem(newSlot: ItemSlot, newItem: EquippedItem | null): Gear;
    getEquippedItem(slot: ItemSlot): EquippedItem | null;
    asArray(): Array<EquippedItem | null>;
    asSpec(): EquipmentSpec;
}
export {};
