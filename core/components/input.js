import { TypedEvent } from '/tbc/core/typed_event.js';
import { Component } from './component.js';
// Shared logic for UI elements that are mapped to a value for some modifiable object.
export class Input extends Component {
    constructor(parent, cssClass, modObject, config) {
        super(parent, 'input-root', config.rootElem);
        this.changeEmitter = new TypedEvent();
        this.inputConfig = config;
        this.modObject = modObject;
        this.rootElem.classList.add(cssClass);
        if (config.extraCssClasses) {
            this.rootElem.classList.add(...config.extraCssClasses);
        }
        if (config.label) {
            const labelDiv = document.createElement('div');
            labelDiv.classList.add('input-label-div');
            this.rootElem.appendChild(labelDiv);
            const label = document.createElement('span');
            label.classList.add('input-label');
            label.textContent = config.label;
            labelDiv.appendChild(label);
            if (config.labelTooltip) {
                const tooltip = document.createElement('span');
                tooltip.classList.add('input-tooltip', 'fa', 'fa-info-circle');
                tippy(tooltip, {
                    'content': config.labelTooltip,
                    'allowHTML': true,
                });
                labelDiv.appendChild(tooltip);
            }
        }
        config.changedEvent(this.modObject).on(() => {
            this.setInputValue(config.getValue(this.modObject));
            this.update();
        });
    }
    update() {
        const enable = !this.inputConfig.enableWhen || this.inputConfig.enableWhen(this.modObject);
        if (enable) {
            this.rootElem.classList.remove('disabled');
            this.getInputElem().removeAttribute('disabled');
        }
        else {
            this.rootElem.classList.add('disabled');
            this.getInputElem().setAttribute('disabled', '');
        }
    }
    // Can't call abstract functions in constructor, so need an init() call.
    init() {
        if (this.inputConfig.defaultValue) {
            this.setInputValue(this.inputConfig.defaultValue);
        }
        else {
            this.setInputValue(this.inputConfig.getValue(this.modObject));
        }
        this.update();
    }
    // Child classes should call this method when the value in the input element changes.
    inputChanged() {
        this.inputConfig.setValue(this.modObject, this.getInputValue());
        this.changeEmitter.emit();
    }
}
