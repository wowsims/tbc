import { EquippedItem } from '../api/equipped_item.js';
import { Enchant } from '../proto/common.js';
import { Item } from '../proto/common.js';
import { ItemSlot } from '../proto/common.js';
import { Sim } from '../sim.js';
import { Component } from './component.js';
export declare class GearPicker extends Component {
    readonly itemPickers: Array<ItemPicker>;
    constructor(parent: HTMLElement, sim: Sim<any>);
}
declare class ItemPicker extends Component {
    readonly slot: ItemSlot;
    private readonly sim;
    private readonly iconElem;
    private readonly nameElem;
    private readonly enchantElem;
    private readonly socketsContainerElem;
    private _items;
    private _enchants;
    private _equippedItem;
    constructor(parent: HTMLElement, sim: Sim<any>, slot: ItemSlot, selectorModal: SelectorModal);
    set item(newItem: EquippedItem | null);
}
declare class SelectorModal extends Component {
    private readonly sim;
    private readonly tabsElem;
    private readonly contentElem;
    private readonly closeButton;
    constructor(parent: HTMLElement, sim: Sim<any>);
    setData(slot: ItemSlot, equippedItem: EquippedItem | null, eligibleItems: Array<Item>, eligibleEnchants: Array<Enchant>): void;
    private addGemTabs;
    /**
     * Adds one of the tabs for the item selector menu.
     *
     * T is expected to be Item, Enchant, or Gem. Tab menus for all 3 looks extremely
     * similar so this function uses extra functions to do it generically.
     */
    private addTab;
    private removeTabs;
}
export {};
