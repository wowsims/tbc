import { TypedEvent } from '/tbc/core/typed_event.js';

import { Input, InputConfig } from './input.js';

/**
 * Data for creating a number picker.
 */
export interface NumberPickerConfig<ModObject> extends InputConfig<ModObject, number> {
}

// UI element for picking an arbitrary number field.
export class NumberPicker<ModObject> extends Input<ModObject, number> {
	private readonly inputElem: HTMLInputElement;

  constructor(parent: HTMLElement, modObject: ModObject, config: NumberPickerConfig<ModObject>) {
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

	getInputElem(): HTMLElement {
		return this.inputElem;
	}

	getInputValue(): number {
		return parseInt(this.inputElem.value || '') || 0;
	}

	setInputValue(newValue: number) {
		this.inputElem.value = String(newValue);
	}
}
