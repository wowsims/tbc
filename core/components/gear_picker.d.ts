import { EquippedItem } from '/tbc/core/proto_utils/equipped_item.js';
import { ItemSlot } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { Component } from './component.js';
export declare class GearPicker extends Component {
    readonly itemPickers: Array<ItemPicker>;
    constructor(parent: HTMLElement, player: Player<any>);
}
declare class ItemPicker extends Component {
    readonly slot: ItemSlot;
    private readonly player;
    private readonly iconElem;
    private readonly nameElem;
    private readonly enchantElem;
    private readonly socketsContainerElem;
    private _items;
    private _enchants;
    private _equippedItem;
    constructor(parent: HTMLElement, player: Player<any>, slot: ItemSlot);
    set item(newItem: EquippedItem | null);
}
export {};
