import { getIconUrl } from '/tbc/core/resources.js';
import { ItemOrSpellId } from '/tbc/core/resources.js';
import { setWowheadHref } from '/tbc/core/resources.js';
import { isRightClick } from '/tbc/core/utils.js';

import { Component } from './component.js';
import { Input, InputConfig } from './input.js';

// Data for creating an icon-based input component.
// 
// E.g. one of these for arcane brilliance, another for kings, etc.
// ModObject is the object being modified (Sim, Player, or Target).
// ValueType is either number or boolean.
export interface IconPickerConfig<ModObject, ValueType> extends InputConfig<ModObject, ValueType> {
  id: ItemOrSpellId;
  
  // The number of possible 'states' this icon can have. Most inputs will use 2
  // for a bi-state icon (on or off). 0 indicates an unlimited number of states.
  states: number;

  // Only used if states == 3.
  improvedId?: ItemOrSpellId;
};

// Icon-based UI for picking buffs / consumes / etc
// ModObject is the object being modified (Sim, Player, or Target).
export class IconPicker<ModObject, ValueType> extends Input<ModObject, ValueType> {
  private readonly config: IconPickerConfig<ModObject, ValueType>;

  private readonly rootAnchor: HTMLAnchorElement;
  private readonly improvedAnchor: HTMLAnchorElement;
  private readonly counterElem: HTMLElement;

	private currentValue: number;

  constructor(parent: HTMLElement, modObj: ModObject, config: IconPickerConfig<ModObject, ValueType>) {
		config.rootElem = document.createElement('a');
    super(parent, 'icon-input', modObj, config);
    this.config = config;
		this.currentValue = 0;

    this.rootAnchor = this.rootElem as HTMLAnchorElement;
    this.rootAnchor.target = '_blank';
    this.rootAnchor.dataset.states = String(this.config.states);

    this.rootAnchor.innerHTML = `
    <div class="icon-input-level-container">
      <a class="icon-input-improved"></a>
      <span class="icon-input-counter"></span>
    </div>
    `;

    this.improvedAnchor = this.rootAnchor.getElementsByClassName('icon-input-improved')[0] as HTMLAnchorElement;
    this.counterElem = this.rootAnchor.getElementsByClassName('icon-input-counter')[0] as HTMLElement;

    setWowheadHref(this.rootAnchor, this.config.id);
    getIconUrl(this.config.id).then(url => {
      this.rootAnchor.style.backgroundImage = `url('${url}')`;
    });

    if (this.config.states == 3) {
      if (this.config.improvedId) {
        setWowheadHref(this.improvedAnchor, this.config.improvedId);
        getIconUrl(this.config.improvedId).then(url => {
          this.improvedAnchor.style.backgroundImage = `url('${url}')`;
        });
      } else {
        throw new Error('IconInput missing improved icon id');
      }
    }

		this.init();

    this.rootAnchor.addEventListener('click', event => {
      event.preventDefault();
    });
    this.rootAnchor.addEventListener('contextmenu', event => {
      event.preventDefault();
    });
    this.rootAnchor.addEventListener('mousedown', event => {
      const rightClick = isRightClick(event);

      if (rightClick) {
        if (this.currentValue > 0) {
					this.currentValue--;
					this.inputChanged();
        }
      } else {
        if (this.config.states == 0 || (this.currentValue + 1) < this.config.states) {
					this.currentValue++;
					this.inputChanged();
        }
      }
    });
  }

	getInputElem(): HTMLElement {
		return this.rootAnchor;
	}

	getInputValue(): ValueType {
		if (this.config.states == 2) {
			return Boolean(this.currentValue) as unknown as ValueType;
		} else {
			return this.currentValue as unknown as ValueType;
		}
	}

  setInputValue(newValue: ValueType) {
    this.currentValue = Number(newValue);

    if (this.currentValue > 0) {
      this.rootAnchor.classList.add('active');
      this.counterElem.classList.add('active');
    } else {
      this.rootAnchor.classList.remove('active');
      this.counterElem.classList.remove('active');
    }

    if (this.config.states == 3) {
      if (this.currentValue > 1) {
        this.improvedAnchor.classList.add('active');
      } else {
        this.improvedAnchor.classList.remove('active');
      }
    }

    if (this.config.states > 3 || this.config.states == 0) {
      this.counterElem.textContent = String(this.currentValue);
    }
  }
}
