import { getIconUrl } from '/tbc/core/resources.js';
import { setWowheadHref } from '/tbc/core/resources.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { isRightClick } from '/tbc/core/utils.js';
import { Component } from './component.js';
// Icon-based UI for picking buffs / consumes / etc
export class IconPicker extends Component {
    constructor(parent, rootClass, sim, inputs, simUI) {
        super(parent, 'icon-picker-root');
        this.rootElem.classList.add(rootClass);
        this._inputs = inputs.map(input => new IconInputComponent(this.rootElem, sim, input, simUI));
    }
}
class IconInputComponent extends Component {
    constructor(parent, sim, input, simUI) {
        super(parent, 'icon-input', document.createElement('a'));
        this._clickedEmitter = new TypedEvent();
        this._input = input;
        this._sim = sim;
        this._rootAnchor = this.rootElem;
        this._rootAnchor.target = '_blank';
        this._rootAnchor.dataset.states = String(this._input.states);
        this._rootAnchor.innerHTML = `
    <div class="icon-input-level-container">
      <a class="icon-input-improved"></a>
      <span class="icon-input-counter"></span>
    </div>
    `;
        this._improvedAnchor = this._rootAnchor.getElementsByClassName('icon-input-improved')[0];
        this._counterElem = this._rootAnchor.getElementsByClassName('icon-input-counter')[0];
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
            }
            else {
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
            }
            else {
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
    getValue() {
        return Number(this._input.getValue(this._sim));
    }
    setValue(newValue) {
        if (this._input.setBooleanValue) {
            this._input.setBooleanValue(this._sim, newValue > 0);
        }
        else if (this._input.setNumberValue) {
            this._input.setNumberValue(this._sim, newValue);
        }
    }
    updateIcon() {
        const value = this.getValue();
        if (value > 0) {
            this._rootAnchor.classList.add('active');
            this._counterElem.classList.add('active');
        }
        else {
            this._rootAnchor.classList.remove('active');
            this._counterElem.classList.remove('active');
        }
        if (this._input.states == 3) {
            if (value > 1) {
                this._improvedAnchor.classList.add('active');
            }
            else {
                this._improvedAnchor.classList.remove('active');
            }
        }
        if (this._input.states > 3 || this._input.states == 0) {
            this._counterElem.textContent = String(value);
        }
    }
}
