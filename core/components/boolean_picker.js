import { TypedEvent } from '/tbc/core/typed_event.js';
import { Input } from './input.js';
// UI element for picking an arbitrary number field.
export class BooleanPicker extends Input {
    constructor(parent, modObject, config) {
        super(parent, 'boolean-picker-root', modObject, config);
        this.inputElem = document.createElement('input');
        this.inputElem.type = 'checkbox';
        this.inputElem.classList.add('boolean-picker-input');
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
        return Boolean(this.inputElem.checked);
    }
    setInputValue(newValue) {
        this.inputElem.checked = newValue;
    }
}
