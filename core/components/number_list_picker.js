import { TypedEvent } from '/tbc/core/typed_event.js';
import { arrayEquals } from '/tbc/core/utils.js';
import { Input } from './input.js';
// UI element for picking an arbitrary number list field.
export class NumberListPicker extends Input {
    constructor(parent, modObject, config) {
        super(parent, 'number-list-picker-root', modObject, config);
        this.inputElem = document.createElement('input');
        this.inputElem.type = 'text';
        this.inputElem.placeholder = config.placeholder || '';
        this.inputElem.classList.add('number-list-picker-input');
        this.rootElem.appendChild(this.inputElem);
        this.init();
        this.inputElem.addEventListener('change', event => {
            this.inputChanged(TypedEvent.nextEventID());
        });
    }
    getInputElem() {
        return this.inputElem;
    }
    getInputValue() {
        const str = this.inputElem.value;
        if (!str) {
            return [];
        }
        return str.split(',').map(parseFloat).filter(val => !isNaN(val));
    }
    setInputValue(newValue) {
        if (arrayEquals(this.getInputValue(), newValue)) {
            return;
        }
        this.inputElem.value = newValue.map(v => String(v)).join(',');
    }
}
