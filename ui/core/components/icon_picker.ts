import { GetItemIconUrl } from '../resources';
import { GetSpellIconUrl } from '../resources';
import { Sim } from '../sim';
import { TypedEvent } from '../typed_event';
import { isRightClick } from '../utils';

import { Component } from './component';

// Icon-based UI for picking buffs / consumes / etc
export class IconPicker extends Component {
  private readonly _inputs: Array<IconInputComponent>;

  constructor(parent: HTMLElement, rootClass: string, sim: Sim, inputs: Array<IconInput>) {
    super(parent, 'icon-picker-root');
    this.rootElem.classList.add(rootClass);

    this._inputs = inputs.map(input => new IconInputComponent(this.rootElem, sim, input));
  }
}

class IconInputComponent extends Component {
  private readonly _input: IconInput;
  private readonly _sim: Sim;

  private readonly _rootAnchor: HTMLAnchorElement;
  private readonly _improvedAnchor: HTMLAnchorElement;
  private readonly _counterElem: HTMLElement;

  constructor(parent: HTMLElement, sim: Sim, input: IconInput) {
    super(parent, 'icon-input', document.createElement('a'));
    this._input = input;
    this._sim = sim;

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

    if (this._input.itemId) {
      this._rootAnchor.href = 'https://tbc.wowhead.com/item=' + this._input.itemId;
      GetItemIconUrl(this._input.itemId).then(url => {
        this._rootAnchor.style.backgroundImage = `url('${url}')`;
      });
    } else if (this._input.spellId) {
      this._rootAnchor.href = 'https://tbc.wowhead.com/spell=' + this._input.spellId;
      GetSpellIconUrl(this._input.spellId).then(url => {
        this._rootAnchor.style.backgroundImage = `url('${url}')`;
      });
    } else {
      throw new Error('IconInput missing icon id');
    }

    if (this._input.states == 3) {
      if (this._input.improvedItemId) {
        this._improvedAnchor.href = 'https://tbc.wowhead.com/item=' + this._input.improvedItemId;
        GetItemIconUrl(this._input.improvedItemId).then(url => {
          this._improvedAnchor.style.backgroundImage = `url('${url}')`;
        });
      } else if (this._input.improvedSpellId) {
        this._improvedAnchor.href = 'https://tbc.wowhead.com/spell=' + this._input.improvedSpellId;
        GetSpellIconUrl(this._input.improvedSpellId).then(url => {
          this._improvedAnchor.style.backgroundImage = `url('${url}')`;
        });
      } else {
        throw new Error('IconInput missing improved icon id');
      }
    }

    this.updateIcon();
    this._input.changedEvent(sim).on(() => this.updateIcon());

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
    });
  }

  // Instead of dealing with bool | number, just convert everything to numbers
  private getValue(): number {
    return Number(this._input.getValue(this._sim));
  }

  private setValue(newValue: number) {
    if (this._input.setBooleanValue) {
      this._input.setBooleanValue(this._sim, newValue > 0);
    } else if (this._input.setNumberValue) {
      this._input.setNumberValue(this._sim, newValue);
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

/**
 * Data for creating an icon-based input component.
 *
 * E.g. one of these for arcane brilliance, another for kings, etc.
 */
export type IconInput = {
  // Exactly one of these should be set.
  itemId?: number;
  spellId?: number;

  /**
   * The number of possible 'states' this icon can have. Most inputs will use 2
   * for a bi-state icon (on or off). 0 indicates an unlimited number of states.
   */
  states: number;

  // At most one of these should be set. Only used if states == 3.
  improvedItemId?: number;
  improvedSpellId?: number;

  changedEvent: (sim: Sim) => TypedEvent<any>;
  getValue: (sim: Sim) => boolean | number;

  // Exactly one of these should be set.
  setBooleanValue?: (sim: Sim, newValue: boolean) => void;
  setNumberValue?: (sim: Sim, newValue: number) => void;
};
