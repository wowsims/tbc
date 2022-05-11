import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { arrayEquals } from '/tbc/core/utils.js';

import { Input, InputConfig } from './input.js';

export interface ListPickerConfig<ModObject, ItemType> extends InputConfig<ModObject, Array<ItemType>> {
	itemLabel: string,
	newItem: () => ItemType,
	newItemPicker: ItemType => HTMLElement,
}

interface ItemPickerPair<ItemType> {
	item: ItemType,
	picker: HTMLElement,
}

export class ListPicker<ModObject, ItemType> extends Input<ModObject, Array<ItemType>> {
	private readonly config: ListPickerConfig<ModObject, ItemType>;
	private readonly itemsDiv: HTMLElement;

	private itemPickerPairs: Array<ItemPickerPair<ItemType>>;

	constructor(parent: HTMLElement, modObject: ModObject, config: ListPickerConfig<ModObject, ItemType>) {
		super(parent, 'list-picker-root', modObject, config);
		this.config = config;
		this.itemPickerPairs = [];

		this.rootElem.innerHTML = `
			<div class="list-picker-items"></div>
			<button class="list-picker-new-button sim-button">NEW ${config.itemLabel.toUpperCase()}</button>
		`;
		
		this.itemsDiv = this.rootElem.getElementsByClassName('list-picker-items')[0] as HTMLElement;

		const newItemButton = this.rootElem.getElementsByClassName('list-picker-new-button')[0] as HTMLElement;
		newItemButton.addEventListener('click', event => {
			const newItem = this.config.newItem();
			const newItemPicker = this.config.newItemPicker(newItem);
			const newPair = { item: newItem, picker: newItemPicker };
			this.itemsDiv.addChild(newItemPicker);

			this.itemPickerPairs.push(newPair);
			this.inputChanged(TypedEvent.nextEventID());
		});

		this.init();
	}

	getInputElem(): HTMLElement {
		return this.rootElem;
	}

	getInputValue(): Array<ItemType> {
		return this.itemPickerPairs.map(pair => pair.item);
	}

	setInputValue(newValue: Array<ItemType>): void {
		// Remove items that are no longer in the list.

		// Add items that were missing.

		// Reorder to match the new list.
	}
}
