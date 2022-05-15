import { TypedEvent } from '/tbc/core/typed_event.js';
import { Input } from './input.js';
// UI element for picking an arbitrary number field.
export class NumberPicker extends Input {
    constructor(parent, modObject, config) {
        super(parent, 'number-picker-root', modObject, config);
        this.float = config.float || false;
        this.inputElem = document.createElement('input');
        if (this.float) {
            this.inputElem.type = 'text';
            this.inputElem.inputMode = 'numeric';
        }
        else {
            this.inputElem.type = 'number';
        }
        this.inputElem.classList.add('number-picker-input');
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
        if (this.float) {
            return parseFloat(this.inputElem.value || '') || 0;
        }
        else {
            return parseInt(this.inputElem.value || '') || 0;
        }
    }
    setInputValue(newValue) {
        this.inputElem.value = String(newValue);
    }
}
