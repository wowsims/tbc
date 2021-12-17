import { TypedEvent } from '/tbc/core/typed_event.js';
import { Input } from '/tbc/core/components/input.js';
import { ResultComponent } from './result_component.js';
const ALL_PLAYERS = -1;
const ALL_TARGETS = -1;
;
export class ResultsFilter extends ResultComponent {
    constructor(config) {
        config.rootCssClass = 'results-filter-root';
        super(config);
        this.currentFilter = {
            player: ALL_PLAYERS,
            target: ALL_TARGETS,
        };
        this.changeEmitter = new TypedEvent();
        this.playerFilter = new PlayerFilter(this.rootElem, this.currentFilter);
        this.playerFilter.changeEmitter.on(() => this.changeEmitter.emit());
    }
    getFilter() {
        return {
            player: this.currentFilter.player == ALL_PLAYERS ? null : this.currentFilter.player,
            target: this.currentFilter.target == ALL_TARGETS ? null : this.currentFilter.target,
        };
    }
    onSimResult(resultData) {
        this.playerFilter.setOptions(resultData.result);
    }
}
;
const allPlayersOption = {
    iconUrl: '',
    text: 'All Players',
    color: '',
    value: ALL_PLAYERS,
};
// Dropdown menu for filtering by player.
class PlayerFilter extends Input {
    constructor(parent, filterData) {
        const changeEmitter = new TypedEvent();
        super(parent, 'player-filter-root', filterData, {
            extraCssClasses: [
                'dropdown-root',
            ],
            changedEvent: (filterData) => changeEmitter,
            getValue: (filterData) => filterData.player,
            setValue: (filterData, newValue) => filterData.player = newValue,
        });
        this.filterData = filterData;
        this.currentOptions = [allPlayersOption];
        this.changeEmitter = changeEmitter;
        this.rootElem.innerHTML = `
			<div class="dropdown-button player-filter-button"></div>
			<div class="dropdown-panel player-filter-dropdown"></div>
    `;
        this.buttonElem = this.rootElem.getElementsByClassName('player-filter-button')[0];
        this.dropdownElem = this.rootElem.getElementsByClassName('player-filter-dropdown')[0];
        this.buttonElem.addEventListener('click', event => {
            event.preventDefault();
        });
        this.init();
    }
    setOptions(simResult) {
        this.currentOptions = [allPlayersOption].concat(simResult.getPlayers().map(player => {
            return {
                iconUrl: player.iconUrl,
                text: `${player.name} (#${player.raidIndex + 1})`,
                color: player.classColor,
                value: player.raidIndex,
            };
        }));
        const hasSameOption = this.currentOptions.find(option => option.value == this.getInputValue()) != null;
        if (!hasSameOption) {
            this.filterData.player = allPlayersOption.value;
            this.changeEmitter.emit();
        }
        this.dropdownElem.innerHTML = '';
        this.currentOptions.forEach(option => this.dropdownElem.appendChild(this.makeOption(option)));
    }
    makeOption(data) {
        const option = this.makeOptionElem(data);
        option.addEventListener('click', event => {
            event.preventDefault();
            this.filterData.player = data.value;
            this.changeEmitter.emit();
        });
        return option;
    }
    makeOptionElem(data) {
        const option = document.createElement('div');
        option.classList.add('dropdown-option', 'player-filter-option');
        if (data.color) {
            option.style.backgroundColor = data.color;
        }
        if (data.iconUrl) {
            const icon = document.createElement('img');
            icon.src = data.iconUrl;
            icon.classList.add('player-filter-icon');
            option.appendChild(icon);
        }
        if (data.text) {
            const label = document.createElement('span');
            label.textContent = data.text;
            label.classList.add('player-filter-label');
            option.appendChild(label);
        }
        return option;
    }
    getInputElem() {
        return this.buttonElem;
    }
    getInputValue() {
        return this.filterData.player;
    }
    setInputValue(newValue) {
        this.filterData.player = newValue;
        const optionData = this.currentOptions.find(optionData => optionData.value == newValue);
        if (!optionData) {
            return;
        }
        this.buttonElem.innerHTML = '';
        this.buttonElem.appendChild(this.makeOptionElem(optionData));
    }
}
