import { Input } from './input.js';
export class EnumPicker extends Input {
    constructor(parent, sim, config) {
        super(parent, 'enum-picker-root', sim, config);
        this.selectElem = document.createElement('select');
        this.selectElem.classList.add('enum-picker-selector');
        this.rootElem.appendChild(this.selectElem);
        config.values.forEach((value) => {
            const option = document.createElement('option');
            option.value = String(value.value);
            option.textContent = value.name;
            this.selectElem.appendChild(option);
            if (value.tooltip) {
                option.title = value.tooltip;
            }
        });
        this.init();
        this.selectElem.addEventListener('change', event => {
            this.inputChanged();
        });
    }
    getInputElem() {
        return this.selectElem;
    }
    getInputValue() {
        return parseInt(this.selectElem.value);
    }
    setInputValue(newValue) {
        this.selectElem.value = String(newValue);
    }
}
