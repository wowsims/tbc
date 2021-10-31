import { Input } from './input.js';
// UI element for picking an arbitrary number field.
export class NumberPicker extends Input {
    constructor(parent, modObject, config) {
        super(parent, 'number-picker-root', modObject, config);
        this.inputElem = document.createElement('input');
        this.inputElem.type = "number";
        this.inputElem.classList.add('number-picker-input');
        this.rootElem.appendChild(this.inputElem);
        this.init();
        this.inputElem.addEventListener('input', event => {
            this.inputChanged();
        });
    }
    getInputElem() {
        return this.inputElem;
    }
    getInputValue() {
        return parseInt(this.inputElem.value || '') || 0;
    }
    setInputValue(newValue) {
        this.inputElem.value = String(newValue);
    }
}
