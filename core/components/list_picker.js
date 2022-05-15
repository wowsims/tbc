import { TypedEvent } from '/tbc/core/typed_event.js';
import { swap } from '/tbc/core/utils.js';
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
            const newList = this.config.getValue(this.modObject).concat([newItem]);
            this.config.setValue(TypedEvent.nextEventID(), this.modObject, newList);
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
				<span class="list-picker-item-up fa fa-angle-up"></span>
				<span class="list-picker-item-down fa fa-angle-down"></span>
				<span class="list-picker-item-copy fa fa-copy"></span>
				<span class="list-picker-item-delete fa fa-times"></span>
			</div>
			<div class="list-picker-item">
			</div>
		`;
        const upButton = itemContainer.getElementsByClassName('list-picker-item-up')[0];
        upButton.addEventListener('click', event => {
            const index = this.itemPickerPairs.findIndex(ipp => ipp.item == item);
            if (index == -1) {
                console.error('Could not find list picker item!');
                return;
            }
            if (index == 0) {
                return;
            }
            const newList = this.config.getValue(this.modObject);
            swap(newList, index, index - 1);
            this.config.setValue(TypedEvent.nextEventID(), this.modObject, newList);
        });
        tippy(upButton, {
            'content': `Move Up`,
            'allowHTML': true,
        });
        const downButton = itemContainer.getElementsByClassName('list-picker-item-down')[0];
        downButton.addEventListener('click', event => {
            const index = this.itemPickerPairs.findIndex(ipp => ipp.item == item);
            if (index == -1) {
                console.error('Could not find list picker item!');
                return;
            }
            if (index == this.itemPickerPairs.length - 1) {
                return;
            }
            const newList = this.config.getValue(this.modObject);
            swap(newList, index, index + 1);
            this.config.setValue(TypedEvent.nextEventID(), this.modObject, newList);
        });
        tippy(downButton, {
            'content': `Move Down`,
            'allowHTML': true,
        });
        const copyButton = itemContainer.getElementsByClassName('list-picker-item-copy')[0];
        copyButton.addEventListener('click', event => {
            const index = this.itemPickerPairs.findIndex(ipp => ipp.item == item);
            if (index == -1) {
                console.error('Could not find list picker item!');
                return;
            }
            const copiedItem = this.config.copyItem(item);
            const newList = this.config.getValue(this.modObject).concat([copiedItem]);
            this.config.setValue(TypedEvent.nextEventID(), this.modObject, newList);
        });
        tippy(copyButton, {
            'content': `Copy to New ${this.config.itemLabel}`,
            'allowHTML': true,
        });
        const deleteButton = itemContainer.getElementsByClassName('list-picker-item-delete')[0];
        deleteButton.addEventListener('click', event => {
            const index = this.itemPickerPairs.findIndex(ipp => ipp.item == item);
            if (index == -1) {
                console.error('Could not find list picker item!');
                return;
            }
            const newList = this.config.getValue(this.modObject);
            newList.splice(index, 1);
            this.config.setValue(TypedEvent.nextEventID(), this.modObject, newList);
        });
        tippy(deleteButton, {
            'content': `Delete`,
            'allowHTML': true,
        });
        const itemElem = itemContainer.getElementsByClassName('list-picker-item')[0];
        const itemPicker = this.config.newItemPicker(itemElem, item);
        this.itemsDiv.appendChild(itemContainer);
        this.itemPickerPairs.push({ item: item, picker: itemContainer });
    }
}
