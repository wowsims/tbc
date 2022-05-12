import { TypedEvent } from '/tbc/core/typed_event.js';
import { Input } from './input.js';
export class ListPicker extends Input {
    constructor(parent, modObject, config) {
        super(parent, 'list-picker-root', modObject, config);
        this.config = config;
        this.itemPickerPairs = [];
        this.rootElem.innerHTML = `
			<div class="list-picker-items"></div>
			<button class="list-picker-new-button sim-button">NEW ${config.itemLabel.toUpperCase()}</button>
		`;
        this.itemsDiv = this.rootElem.getElementsByClassName('list-picker-items')[0];
        const newItemButton = this.rootElem.getElementsByClassName('list-picker-new-button')[0];
        newItemButton.addEventListener('click', event => {
            const newItem = this.config.newItem();
            this.addNewPicker(newItem);
            this.inputChanged(TypedEvent.nextEventID());
        });
        this.init();
    }
    getInputElem() {
        return this.rootElem;
    }
    getInputValue() {
        return this.itemPickerPairs.map(pair => pair.item);
    }
    setInputValue(newValue) {
        // Remove items that are no longer in the list.
        const removePairs = this.itemPickerPairs.filter(ipp => !newValue.includes(ipp.item));
        removePairs.forEach(ipp => ipp.picker.remove());
        this.itemPickerPairs = this.itemPickerPairs.filter(ipp => !removePairs.includes(ipp));
        // Add items that were missing.
        const curItems = this.getInputValue();
        newValue
            .filter(newItem => !curItems.includes(newItem))
            .forEach(newItem => this.addNewPicker(newItem));
        // Reorder to match the new list.
        this.itemPickerPairs = newValue.map(item => this.itemPickerPairs.find(ipp => ipp.item == item));
        // Reorder item picker elements in the DOM if necessary.
        const curPickers = Array.from(this.itemsDiv.children);
        if (!curPickers.every((picker, i) => picker == this.itemPickerPairs[i].picker)) {
            this.itemPickerPairs.forEach(ipp => ipp.picker.remove());
            this.itemPickerPairs.forEach(ipp => this.itemsDiv.appendChild(ipp.picker));
        }
    }
    addNewPicker(item) {
        const itemContainer = document.createElement('div');
        itemContainer.classList.add('list-picker-item-container');
        itemContainer.innerHTML = `
			<div class="list-picker-item-header">
				<span class="list-picker-item-delete fa fa-times"></span>
			</div>
			<div class="list-picker-item">
			</div>
		`;
        const deleteButton = itemContainer.getElementsByClassName('list-picker-item-delete')[0];
        deleteButton.addEventListener('click', event => {
            const index = this.itemPickerPairs.findIndex(ipp => ipp.item == item);
            if (index == -1) {
                return;
            }
            this.itemPickerPairs[index].picker.remove();
            this.itemPickerPairs.splice(index, 1);
            this.inputChanged(TypedEvent.nextEventID());
        });
        const itemElem = itemContainer.getElementsByClassName('list-picker-item')[0];
        const itemPicker = this.config.newItemPicker(itemElem, item);
        this.itemsDiv.appendChild(itemContainer);
        this.itemPickerPairs.push({ item: item, picker: itemContainer });
    }
}
