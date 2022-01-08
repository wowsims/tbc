import { TypedEvent } from '/tbc/core/typed_event.js';
import { isRightClick } from '/tbc/core/utils.js';
import { Input } from './input.js';
;
// Icon-based UI for picking buffs / consumes / etc
// ModObject is the object being modified (Sim, Player, or Target).
export class IconPicker extends Input {
    constructor(parent, modObj, config) {
        config.rootElem = document.createElement('a');
        super(parent, 'icon-input', modObj, config);
        this.config = config;
        this.currentValue = 0;
        this.rootAnchor = this.rootElem;
        this.rootAnchor.target = '_blank';
        this.rootAnchor.dataset.states = String(this.config.states);
        this.rootAnchor.innerHTML = `
    <div class="icon-input-level-container">
      <a class="icon-input-improved"></a>
      <span class="icon-input-counter"></span>
    </div>
    `;
        this.improvedAnchor = this.rootAnchor.getElementsByClassName('icon-input-improved')[0];
        this.counterElem = this.rootAnchor.getElementsByClassName('icon-input-counter')[0];
        this.config.id.fillAndSet(this.rootAnchor, true, true);
        if (this.config.states == 3) {
            if (this.config.improvedId) {
                this.config.improvedId.fillAndSet(this.improvedAnchor, true, true);
            }
            else {
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
                    this.inputChanged(TypedEvent.nextEventID());
                }
            }
            else {
                if (this.config.states == 0 || (this.currentValue + 1) < this.config.states) {
                    this.currentValue++;
                    this.inputChanged(TypedEvent.nextEventID());
                }
            }
        });
    }
    getInputElem() {
        return this.rootAnchor;
    }
    getInputValue() {
        if (this.config.states == 2) {
            return Boolean(this.currentValue);
        }
        else {
            return this.currentValue;
        }
    }
    setInputValue(newValue) {
        this.currentValue = Number(newValue);
        if (this.currentValue > 0) {
            this.rootAnchor.classList.add('active');
            this.counterElem.classList.add('active');
        }
        else {
            this.rootAnchor.classList.remove('active');
            this.counterElem.classList.remove('active');
        }
        if (this.config.states == 3) {
            if (this.currentValue > 1) {
                this.improvedAnchor.classList.add('active');
            }
            else {
                this.improvedAnchor.classList.remove('active');
            }
        }
        if (this.config.states > 3 || this.config.states == 0) {
            this.counterElem.textContent = String(this.currentValue);
        }
    }
}
