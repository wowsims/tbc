import { Sim } from '/tbc/core/sim.js';
import { TypedEvent } from '/tbc/core/typed_event.js';

import { Input, InputConfig } from './input.js';

/**
 * Data for creating a number picker.
 */
export interface NumberPickerConfig extends InputConfig<number> {
}

// UI element for picking an arbitrary number field.
export class NumberPicker extends Input<number> {
	private readonly inputElem: HTMLInputElement;

  constructor(parent: HTMLElement, sim: Sim<any>, config: NumberPickerConfig) {
    super(parent, 'number-picker-root', sim, config);

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
