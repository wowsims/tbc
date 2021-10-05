import { Race } from '../proto/common.js';
import { specToEligibleRaces } from '../api/utils.js';
import { raceNames } from '../api/names.js';
import { Sim } from '../sim.js';
import { TypedEvent } from '../typed_event.js';

import { Component } from './component.js';

export interface EnumPickerConfig {
  label?: string,
  defaultValue?: number,

  // names and values are parallel arrays
  names: Array<string>;
  values: Array<number>;

  changedEvent: (sim: Sim<any>) => TypedEvent<any>;
  getValue: (sim: Sim<any>) => number;
  setValue: (sim: Sim<any>, newValue: number) => void;
}

export class EnumPicker extends Component {
  constructor(parent: HTMLElement, sim: Sim<any>, config: EnumPickerConfig) {
    super(parent, 'enum-picker-root');

    if (config.label) {
      const label = document.createElement('span');
      label.classList.add('enum-picker-label');
      label.textContent = config.label;
      this.rootElem.appendChild(label);
    }

    const selector = document.createElement('select');
    selector.classList.add('enum-picker-selector');
    this.rootElem.appendChild(selector);

    config.values.forEach((value, idx) => {
      const option = document.createElement('option');
      option.value = String(value);
      option.textContent = config.names[idx];
      selector.appendChild(option);
    });

    selector.value = String(config.getValue(sim));
    config.changedEvent(sim).on(() => {
      selector.value = String(config.getValue(sim));
    });

    if (config.defaultValue) {
      config.setValue(sim, config.defaultValue);
    }

    selector.addEventListener('change', event => {
      config.setValue(sim, parseInt(selector.value));
    });
  }
}
