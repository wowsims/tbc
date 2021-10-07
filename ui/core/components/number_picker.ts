import { Sim } from '/tbc/core/sim.js';
import { TypedEvent } from '/tbc/core/typed_event.js';

import { Component } from './component.js';

/**
 * Data for creating a number picker.
 */
export type NumberPickerConfig = {
  label?: string,
  defaultValue?: number,

  changedEvent: (sim: Sim<any>) => TypedEvent<any>;
  getValue: (sim: Sim<any>) => number;
  setValue: (sim: Sim<any>, newValue: number) => void;
};

// UI element for picking an arbitrary number field.
export class NumberPicker extends Component {
  constructor(parent: HTMLElement, sim: Sim<any>, config: NumberPickerConfig) {
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
