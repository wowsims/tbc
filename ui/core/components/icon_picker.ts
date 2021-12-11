import { getIconUrl } from '/tbc/core/resources.js';
import { ItemOrSpellId } from '/tbc/core/resources.js';
import { setWowheadHref } from '/tbc/core/resources.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { isRightClick } from '/tbc/core/utils.js';
import { ExclusivityTag } from '/tbc/core/individual_sim_ui.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';

import { Component } from './component.js';

// Icon-based UI for picking buffs / consumes / etc
// ModObject is the object being modified (Sim, Player, or Target).
export class IconPicker<ModObject> extends Component {
  private readonly _input: IconInput<ModObject>;
  private readonly _modObject: ModObject;

  private readonly _rootAnchor: HTMLAnchorElement;
  private readonly _improvedAnchor: HTMLAnchorElement;
  private readonly _counterElem: HTMLElement;
  private readonly _clickedEmitter = new TypedEvent<void>();

  constructor(parent: HTMLElement, modObj: ModObject, input: IconInput<ModObject>, simUI: IndividualSimUI<any>) {
    super(parent, 'icon-input', document.createElement('a'));
    this._input = input;
    this._modObject = modObj;

    this._rootAnchor = this.rootElem as HTMLAnchorElement;
    this._rootAnchor.target = '_blank';
    this._rootAnchor.dataset.states = String(this._input.states);

    this._rootAnchor.innerHTML = `
    <div class="icon-input-level-container">
      <a class="icon-input-improved"></a>
      <span class="icon-input-counter"></span>
    </div>
    `;

    this._improvedAnchor = this._rootAnchor.getElementsByClassName('icon-input-improved')[0] as HTMLAnchorElement;
    this._counterElem = this._rootAnchor.getElementsByClassName('icon-input-counter')[0] as HTMLAnchorElement;

    setWowheadHref(this._rootAnchor, this._input.id);
    getIconUrl(this._input.id).then(url => {
      this._rootAnchor.style.backgroundImage = `url('${url}')`;
    });

    if (this._input.states == 3) {
      if (this._input.improvedId) {
        setWowheadHref(this._improvedAnchor, this._input.improvedId);
        getIconUrl(this._input.improvedId).then(url => {
          this._improvedAnchor.style.backgroundImage = `url('${url}')`;
        });
      } else {
        throw new Error('IconInput missing improved icon id');
      }
    }

    this.updateIcon();
    this._input.changedEvent(this._modObject).on(() => this.updateIcon());

    this._rootAnchor.addEventListener('click', event => {
      event.preventDefault();
    });
    this._rootAnchor.addEventListener('contextmenu', event => {
      event.preventDefault();
    });
    this._rootAnchor.addEventListener('mousedown', event => {
      const rightClick = isRightClick(event);
      const value = this.getValue();

      if (rightClick) {
        if (value > 0) {
          this.setValue(value - 1);
        }
      } else {
        if (this._input.states == 0 || (value + 1) < this._input.states) {
          this.setValue(value + 1);
        }
      }
      this._clickedEmitter.emit();
    });

    if (this._input.exclusivityTags) {
      simUI.registerExclusiveEffect({
        tags: this._input.exclusivityTags,
        changedEvent: this._clickedEmitter,
        isActive: () => Boolean(this.getValue()),
        deactivate: () => this.setValue(0),
      });
    }
  }

  // Instead of dealing with bool | number, just convert everything to numbers
  private getValue(): number {
    return Number(this._input.getValue(this._modObject));
  }

  private setValue(newValue: number) {
    if (this._input.setBooleanValue) {
      this._input.setBooleanValue(this._modObject, newValue > 0);
    } else if (this._input.setNumberValue) {
      this._input.setNumberValue(this._modObject, newValue);
    }
  }

  private updateIcon() {
    const value = this.getValue();
    if (value > 0) {
      this._rootAnchor.classList.add('active');
      this._counterElem.classList.add('active');
    } else {
      this._rootAnchor.classList.remove('active');
      this._counterElem.classList.remove('active');
    }

    if (this._input.states == 3) {
      if (value > 1) {
        this._improvedAnchor.classList.add('active');
      } else {
        this._improvedAnchor.classList.remove('active');
      }
    }

    if (this._input.states > 3 || this._input.states == 0) {
      this._counterElem.textContent = String(value);
    }
  }
}

// Data for creating an icon-based input component.
// 
// E.g. one of these for arcane brilliance, another for kings, etc.
// ModObject is the object being modified (Sim, Player, or Target).
export type IconInput<ModObject> = {
  id: ItemOrSpellId;
  
  // The number of possible 'states' this icon can have. Most inputs will use 2
  // for a bi-state icon (on or off). 0 indicates an unlimited number of states.
  states: number;

  // Only used if states == 3.
  improvedId?: ItemOrSpellId;
  
  // If set, all effects with matching tags will be deactivated when this
  // effect is enabled.
  exclusivityTags?: Array<ExclusivityTag>;

  changedEvent: (obj: ModObject) => TypedEvent<any>;
  getValue: (obj: ModObject) => boolean | number;

  // Exactly one of these should be set.
  setBooleanValue?: (obj: ModObject, newValue: boolean) => void;
  setNumberValue?: (obj: ModObject, newValue: number) => void;
};
