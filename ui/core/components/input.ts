import { Sim } from '/tbc/core/sim.js';
import { TypedEvent } from '/tbc/core/typed_event.js';

import { Component } from './component.js';

declare var tippy: any;

/**
 * Data for creating a new input UI element.
 */
export interface InputConfig<T> {
  label?: string,
	labelTooltip?: string,

  defaultValue?: T,

	// Returns the event indicating the mapped Sim value has changed.
  changedEvent: (sim: Sim<any>) => TypedEvent<any>,

	// Get and set the mapped Sim value.
  getValue: (sim: Sim<any>) => T,
  setValue: (sim: Sim<any>, newValue: T) => void,

	// If set, will automatically disable the input when this evaluates to false.
	enableWhen?: (sim: Sim<any>) => boolean,
}

// Shared logic for UI elements that are mapped to Sim values.
export abstract class Input<T> extends Component {
	private readonly inputConfig: InputConfig<T>;
	readonly sim: Sim<any>;

  constructor(parent: HTMLElement, cssClass: string, sim: Sim<any>, config: InputConfig<T>) {
    super(parent, 'input-root');
		this.inputConfig = config;
		this.sim = sim;
		this.rootElem.classList.add(cssClass);

    if (config.label) {
      const label = document.createElement('span');
      label.classList.add('input-label');
      label.textContent = config.label;
      this.rootElem.appendChild(label);

			if (config.labelTooltip) {
				tippy(label, {
					'content': config.labelTooltip,
					'allowHTML': true,
				});
			}
    }

    config.changedEvent(this.sim).on(() => {
			this.setInputValue(config.getValue(this.sim));
			this.update();
    });
	}

	private update() {
		const enable = !this.inputConfig.enableWhen || this.inputConfig.enableWhen(this.sim);
		if (enable) {
			this.rootElem.classList.remove('disabled');
			this.getInputElem().removeAttribute('disabled');
		} else {
			this.rootElem.classList.add('disabled');
			this.getInputElem().setAttribute('disabled', '');
		}
	}

	// Can't call abstract functions in constructor, so need an init() call.
	init() {
		if (this.inputConfig.defaultValue) {
			this.setInputValue(this.inputConfig.defaultValue);
		} else {
			this.setInputValue(this.inputConfig.getValue(this.sim));
		}
		this.update();
	}

	abstract getInputElem(): HTMLElement;

	abstract getInputValue(): T;

	abstract setInputValue(newValue: T): void;

	// Child classes should call this method when the value in the input element changes.
	inputChanged() {
		this.inputConfig.setValue(this.sim, this.getInputValue());
	}
}
