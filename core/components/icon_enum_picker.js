import { getIconUrl } from '/tbc/core/resources.js';
import { setWowheadHref } from '/tbc/core/resources.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { Input } from './input.js';
// Icon-based UI for picking enum values.
// ModObject is the object being modified (Sim, Player, or Target).
export class IconEnumPicker extends Input {
    constructor(parent, modObj, config) {
        super(parent, 'icon-enum-picker-root', modObj, config);
        this.config = config;
        this.currentValue = 0;
        this.rootElem.classList.add('dropdown-root');
        this.rootElem.innerHTML = `
			<a class="dropdown-button icon-enum-picker-button"></a>
			<div class="dropdown-panel icon-enum-picker-dropdown"></div>
    `;
        this.buttonElem = this.rootElem.getElementsByClassName('icon-enum-picker-button')[0];
        const dropdownElem = this.rootElem.getElementsByClassName('icon-enum-picker-dropdown')[0];
        this.buttonElem.addEventListener('click', event => {
            event.preventDefault();
        });
        config.values.forEach(valueConfig => {
            const option = document.createElement('a');
            option.classList.add('dropdown-option', 'icon-enum-picker-option');
            dropdownElem.appendChild(option);
            this.setImage(option, valueConfig);
            option.addEventListener('click', event => {
                event.preventDefault();
                this.currentValue = valueConfig.value;
                this.inputChanged(TypedEvent.nextEventID());
            });
        });
        this.init();
    }
    setImage(elem, valueConfig) {
        if (valueConfig.id) {
            setWowheadHref(elem, valueConfig.id);
            getIconUrl(valueConfig.id).then(url => {
                elem.style.backgroundImage = `url('${url}')`;
            });
        }
        else {
            elem.style.backgroundImage = '';
            elem.style.backgroundColor = valueConfig.color;
        }
    }
    getInputElem() {
        return this.buttonElem;
    }
    getInputValue() {
        return this.currentValue;
    }
    setInputValue(newValue) {
        this.currentValue = newValue;
        if (this.currentValue > 0) {
            this.rootElem.classList.add('active');
        }
        else {
            this.rootElem.classList.remove('active');
        }
        const valueConfig = this.config.values.find(valueConfig => valueConfig.value == this.currentValue);
        this.setImage(this.buttonElem, valueConfig);
    }
}
