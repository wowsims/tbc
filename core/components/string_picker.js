import { TypedEvent } from '/tbc/core/typed_event.js';
import { Input } from './input.js';
// UI element for picking an arbitrary string field.
export class StringPicker extends Input {
    constructor(parent, modObject, config) {
        super(parent, 'string-picker-root', modObject, config);
        this.inputElem = document.createElement('span');
        this.inputElem.setAttribute('contenteditable', '');
        this.inputElem.classList.add('string-picker-input');
        this.rootElem.appendChild(this.inputElem);
        this.init();
        this.inputElem.addEventListener('input', event => {
            this.inputChanged(TypedEvent.nextEventID());
        });
    }
    getInputElem() {
        return this.inputElem;
    }
    getInputValue() {
        return this.inputElem.textContent || '';
    }
    setInputValue(newValue) {
        this.inputElem.textContent = newValue;
    }
}
