import { Component } from './component.js';
// UI element for picking an arbitrary number field.
export class NumberPicker extends Component {
    constructor(parent, sim, config) {
        super(parent, 'number-picker-root');
        if (config.label) {
            const label = document.createElement('span');
            label.classList.add('number-picker-label');
            label.textContent = config.label;
            this.rootElem.appendChild(label);
        }
        const input = document.createElement('input');
        input.type = "number";
        input.classList.add('number-picker-input');
        this.rootElem.appendChild(input);
        input.value = String(config.getValue(sim));
        config.changedEvent(sim).on(() => {
            input.value = String(config.getValue(sim));
        });
        if (config.defaultValue) {
            config.setValue(sim, config.defaultValue);
        }
        input.addEventListener('input', event => {
            config.setValue(sim, parseInt(input.value || '') || 0);
        });
    }
}
