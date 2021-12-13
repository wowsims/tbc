import { getIconUrl } from '/tbc/core/resources.js';
import { ItemOrSpellId } from '/tbc/core/resources.js';
import { setWowheadHref } from '/tbc/core/resources.js';

import { Component } from './component.js';
import { Input, InputConfig } from './input.js';

export interface IconEnumValueConfig {
	// One of these should be set. If id is set, shows the icon for that id. If
	// color is set, shows that color.
  id?: ItemOrSpellId,
	color?: string,

	value: number,
}

export interface IconEnumPickerConfig<ModObject> extends InputConfig<ModObject, number> {
	values: Array<IconEnumValueConfig>;
}

// Icon-based UI for picking enum values.
// ModObject is the object being modified (Sim, Player, or Target).
export class IconEnumPicker<ModObject> extends Input<ModObject, number> {
  private readonly config: IconEnumPickerConfig<ModObject>;

	private currentValue: number;

	private readonly buttonElem: HTMLAnchorElement;

  constructor(parent: HTMLElement, modObj: ModObject, config: IconEnumPickerConfig<ModObject>) {
    super(parent, 'icon-enum-picker-root', modObj, config);
    this.config = config;
		this.currentValue = 0;

    this.rootElem.innerHTML = `
			<a class="icon-enum-picker-button"></a>
			<div class="icon-enum-picker-dropdown"></div>
    `;

		this.buttonElem = this.rootElem.getElementsByClassName('icon-enum-picker-button')[0] as HTMLAnchorElement;
		const dropdownElem = this.rootElem.getElementsByClassName('icon-enum-picker-dropdown')[0] as HTMLElement;

		this.buttonElem.addEventListener('click', event => {
			event.preventDefault();
		});

		config.values.forEach(valueConfig => {
			const option = document.createElement('a');
			option.classList.add('icon-enum-picker-option');
			dropdownElem.appendChild(option);
			this.setImage(option, valueConfig);

			option.addEventListener('click', event => {
				event.preventDefault();
				this.currentValue = valueConfig.value;
				this.inputChanged();
			});
		});

		this.init();
  }

	private setImage(elem: HTMLAnchorElement, valueConfig: IconEnumValueConfig) {
		if (valueConfig.id) {
			setWowheadHref(elem, valueConfig.id);
			getIconUrl(valueConfig.id).then(url => {
				elem.style.backgroundImage = `url('${url}')`;
			});
		} else {
			elem.style.backgroundImage = '';
			elem.style.backgroundColor = valueConfig.color!;
		}
	}

	getInputElem(): HTMLElement {
		return this.buttonElem;
	}

	getInputValue(): number {
		return this.currentValue;
	}

  setInputValue(newValue: number) {
    this.currentValue = newValue;

		const valueConfig = this.config.values.find(valueConfig => valueConfig.value == this.currentValue)!;
		this.setImage(this.buttonElem, valueConfig);
  }
}
