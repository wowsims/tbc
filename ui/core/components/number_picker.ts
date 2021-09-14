import { Sim } from '../sim.js';
import { TypedEvent } from '../typed_event.js';

import { Component } from './component.js';

// UI element for picking an arbitrary number field.
export class NumberPicker extends Component {
  constructor(parent: HTMLElement, sim: Sim, config: NumberPickerConfig) {
    super(parent, 'number-picker-root');

    const input = document.createElement('input');
    input.type = "number";
    input.classList.add('number-picker-input');
    this.rootElem.appendChild(input);

    input.value = String(config.getValue(sim));
    config.changedEvent(sim).on(() => {
      input.value = String(config.getValue(sim));
    });

    input.addEventListener('input', event => {
      config.setValue(sim, parseInt(input.value || '') || 0);
    });
  }
}

/**
 * Data for creating a number picker.
 */
export type NumberPickerConfig = {
  changedEvent: (sim: Sim) => TypedEvent<any>;
  getValue: (sim: Sim) => number;
  setValue: (sim: Sim, newValue: number) => void;
};
