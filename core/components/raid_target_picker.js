import { Input } from '/tbc/core/components/input.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
import { emptyRaidTarget } from '/tbc/core/proto_utils/utils.js';
;
;
// Dropdown menu for selecting a player.
export class RaidTargetPicker extends Input {
    constructor(parent, modObj, config) {
        super(parent, 'raid-target-picker-root', modObj, config);
        this.rootElem.classList.add('dropdown-root');
        this.config = config;
        this.raidTarget = emptyRaidTarget();
        this.noTargetOption = {
            iconUrl: '',
            text: config.noTargetLabel,
            color: 'black',
            value: emptyRaidTarget(),
            isDropdown: true,
        };
        this.currentOptions = [this.noTargetOption];
        this.rootElem.innerHTML = `
			<div class="dropdown-button raid-target-picker-button"></div>
			<div class="dropdown-panel raid-target-picker-dropdown"></div>
    `;
        this.buttonElem = this.rootElem.getElementsByClassName('raid-target-picker-button')[0];
        this.dropdownElem = this.rootElem.getElementsByClassName('raid-target-picker-dropdown')[0];
        this.buttonElem.addEventListener('click', event => {
            event.preventDefault();
        });
        this.setOptions(config.getOptions());
        config.compChangeEmitter.on(() => {
            this.setOptions(config.getOptions());
        });
        this.init();
    }
    setOptions(options) {
        this.currentOptions = [this.noTargetOption].concat(options);
        const hasSameOption = this.currentOptions.find(option => RaidTarget.equals(option.value, this.getInputValue())) != null;
        if (!hasSameOption) {
            this.raidTarget = this.noTargetOption.value;
            this.inputChanged();
        }
        this.dropdownElem.innerHTML = '';
        this.currentOptions.forEach(option => this.dropdownElem.appendChild(this.makeOption(option)));
    }
    makeOption(data) {
        const option = RaidTargetPicker.makeOptionElem(data);
        option.addEventListener('click', event => {
            event.preventDefault();
            this.raidTarget = data.value;
            this.inputChanged();
        });
        return option;
    }
    getInputElem() {
        return this.buttonElem;
    }
    getInputValue() {
        return RaidTarget.clone(this.raidTarget);
    }
    setInputValue(newValue) {
        this.raidTarget = RaidTarget.clone(newValue);
        const optionData = this.currentOptions.find(optionData => RaidTarget.equals(optionData.value, newValue));
        if (!optionData) {
            return;
        }
        this.buttonElem.innerHTML = '';
        this.buttonElem.appendChild(RaidTargetPicker.makeOptionElem(optionData));
    }
    static makeOptionElem(data) {
        const option = document.createElement('div');
        option.classList.add('raid-target-picker-option');
        if (data.isDropdown) {
            option.classList.add('dropdown-option');
        }
        if (data.color) {
            option.style.backgroundColor = data.color;
        }
        if (data.iconUrl) {
            const icon = document.createElement('img');
            icon.src = data.iconUrl;
            icon.classList.add('raid-target-picker-icon');
            option.appendChild(icon);
        }
        if (data.text) {
            const label = document.createElement('span');
            label.textContent = data.text;
            label.classList.add('raid-target-picker-label');
            option.appendChild(label);
        }
        return option;
    }
}
