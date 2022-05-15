import { Input, InputConfig } from './input.js';
export interface ListPickerConfig<ModObject, ItemType> extends InputConfig<ModObject, Array<ItemType>> {
    itemLabel: string;
    newItem: () => ItemType;
    copyItem: (oldItem: ItemType) => ItemType;
    newItemPicker: (parent: HTMLElement, item: ItemType) => void;
}
export declare class ListPicker<ModObject, ItemType> extends Input<ModObject, Array<ItemType>> {
    private readonly config;
    private readonly itemsDiv;
    private itemPickerPairs;
    constructor(parent: HTMLElement, modObject: ModObject, config: ListPickerConfig<ModObject, ItemType>);
    getInputElem(): HTMLElement;
    getInputValue(): Array<ItemType>;
    setInputValue(newValue: Array<ItemType>): void;
    private addNewPicker;
}
