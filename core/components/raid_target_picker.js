import { Input } from '/tbc/core/components/input.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
import { emptyRaidTarget } from '/tbc/core/proto_utils/utils.js';
;
;
// Dropdown menu for selecting a player.
export class RaidTargetPicker extends Input {
    constructor(parent, raid, modObj, config) {
        super(parent, 'raid-target-picker-root', modObj, config);
        this.rootElem.classList.add('dropdown-root');
        this.config = config;
        this.raid = raid;
        this.curPlayer = this.raid.getPlayerFromRaidTarget(config.getValue(modObj));
        this.curRaidTarget = this.getInputValue();
        this.noTargetOption = {
            iconUrl: '',
            text: config.noTargetLabel,
            color: 'black',
            value: null,
            isDropdown: true,
        };
        this.rootElem.innerHTML = `
			<div class="dropdown-button raid-target-picker-button"></div>
			<div class="dropdown-panel raid-target-picker-dropdown"></div>
    `;
        this.buttonElem = this.rootElem.getElementsByClassName('raid-target-picker-button')[0];
        this.dropdownElem = this.rootElem.getElementsByClassName('raid-target-picker-dropdown')[0];
        this.buttonElem.addEventListener('click', event => {
            event.preventDefault();
        });
        this.currentOptions = [];
        this.updateOptions(TypedEvent.nextEventID());
        config.compChangeEmitter.on(eventID => {
            this.updateOptions(eventID);
        });
        this.init();
    }
    makeTargetOptions() {
        const playerOptions = this.raid.getPlayers().filter(player => player != null).map(player => {
            return {
                iconUrl: player.getTalentTreeIcon(),
                text: player.getLabel(),
                color: player.getClassColor(),
                isDropdown: true,
                value: player,
            };
        });
        return [this.noTargetOption].concat(playerOptions);
    }
    updateOptions(eventID) {
        this.currentOptions = this.makeTargetOptions();
        const prevRaidTarget = this.curRaidTarget;
        this.curRaidTarget = this.getInputValue();
        if (!RaidTarget.equals(prevRaidTarget, this.curRaidTarget)) {
            this.inputChanged(eventID);
        }
        this.dropdownElem.innerHTML = '';
        this.currentOptions.forEach(option => this.dropdownElem.appendChild(this.makeOption(option)));
    }
    makeOption(data) {
        const option = RaidTargetPicker.makeOptionElem(data);
        option.addEventListener('click', event => {
            event.preventDefault();
            this.curPlayer = data.value;
            this.curRaidTarget = this.getInputValue();
            this.inputChanged(TypedEvent.nextEventID());
        });
        return option;
    }
    getInputElem() {
        return this.buttonElem;
    }
    getInputValue() {
        if (this.curPlayer) {
            return this.curPlayer.makeRaidTarget();
        }
        else {
            return emptyRaidTarget();
        }
    }
    setInputValue(newValue) {
        this.curRaidTarget = RaidTarget.clone(newValue);
        this.curPlayer = this.raid.getPlayerFromRaidTarget(this.curRaidTarget);
        const optionData = this.currentOptions.find(optionData => optionData.value == this.curPlayer);
        if (!optionData) {
            return;
        }
        this.buttonElem.innerHTML = '';
        this.buttonElem.appendChild(RaidTargetPicker.makeOptionElem(optionData));
    }
    static makeOptionElem(data) {
        const optionContainer = document.createElement('div');
        optionContainer.classList.add('dropdown-option-container');
        const option = document.createElement('div');
        option.classList.add('raid-target-picker-option');
        optionContainer.appendChild(option);
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
        return optionContainer;
    }
}
