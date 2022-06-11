import { TypedEvent } from '/tbc/core/typed_event.js';
import { Input } from '/tbc/core/components/input.js';
import { ResultComponent } from './result_component.js';
const ALL_UNITS = -1;
;
export class ResultsFilter extends ResultComponent {
    constructor(config) {
        config.rootCssClass = 'results-filter-root';
        super(config);
        this.currentFilter = {
            player: ALL_UNITS,
            target: ALL_UNITS,
        };
        this.changeEmitter = new TypedEvent();
        this.playerFilter = new PlayerFilter(this.rootElem, this.currentFilter);
        this.playerFilter.changeEmitter.on(eventID => this.changeEmitter.emit(eventID));
        this.targetFilter = new TargetFilter(this.rootElem, this.currentFilter);
        this.targetFilter.changeEmitter.on(eventID => this.changeEmitter.emit(eventID));
    }
    getFilter() {
        return {
            player: this.currentFilter.player == ALL_UNITS ? null : this.currentFilter.player,
            target: this.currentFilter.target == ALL_UNITS ? null : this.currentFilter.target,
        };
    }
    onSimResult(resultData) {
        this.playerFilter.setOptions(resultData.eventID, resultData.result);
        this.targetFilter.setOptions(resultData.eventID, resultData.result);
    }
    setPlayer(eventID, newPlayer) {
        this.currentFilter.player = (newPlayer === null) ? ALL_UNITS : newPlayer;
        this.playerFilter.changeEmitter.emit(eventID);
    }
    setTarget(eventID, newTarget) {
        this.currentFilter.target = (newTarget === null) ? ALL_UNITS : newTarget;
        this.targetFilter.changeEmitter.emit(eventID);
    }
}
;
// Dropdown menu for filtering by player.
class UnitGroupFilter extends Input {
    constructor(parent, filterData, allUnitsLabel) {
        const changeEmitter = new TypedEvent();
        super(parent, 'unit-filter-root', filterData, {
            extraCssClasses: [
                'dropdown-root',
            ],
            changedEvent: (filterData) => changeEmitter,
            getValue: (filterData) => this.getFilterDataValue(filterData),
            setValue: (eventID, filterData, newValue) => this.setFilterDataValue(filterData, newValue),
        });
        this.filterData = filterData;
        this.changeEmitter = changeEmitter;
        this.allUnitsOption = {
            iconUrl: '',
            text: allUnitsLabel,
            color: 'black',
            value: ALL_UNITS,
        };
        this.currentOptions = [this.allUnitsOption];
        this.rootElem.innerHTML = `
			<div class="dropdown-button unit-filter-button"></div>
			<div class="dropdown-panel unit-filter-dropdown"></div>
    `;
        this.buttonElem = this.rootElem.getElementsByClassName('unit-filter-button')[0];
        this.dropdownElem = this.rootElem.getElementsByClassName('unit-filter-dropdown')[0];
        this.buttonElem.addEventListener('click', event => {
            event.preventDefault();
        });
        this.init();
    }
    setOptions(eventID, simResult) {
        this.currentOptions = [this.allUnitsOption].concat(this.getAllUnits(simResult).map(unit => {
            return {
                iconUrl: unit.iconUrl || '',
                text: unit.label,
                color: unit.classColor || 'black',
                value: unit.index,
            };
        }));
        const hasSameOption = this.currentOptions.find(option => option.value == this.getInputValue()) != null;
        if (!hasSameOption) {
            this.setFilterDataValue(this.filterData, this.allUnitsOption.value);
            this.changeEmitter.emit(eventID);
        }
        this.dropdownElem.innerHTML = '';
        this.currentOptions.forEach(option => this.dropdownElem.appendChild(this.makeOption(option)));
    }
    makeOption(data) {
        const option = this.makeOptionElem(data);
        option.addEventListener('click', event => {
            event.preventDefault();
            this.setFilterDataValue(this.filterData, data.value);
            this.changeEmitter.emit(TypedEvent.nextEventID());
        });
        return option;
    }
    makeOptionElem(data) {
        const optionContainer = document.createElement('div');
        optionContainer.classList.add('dropdown-option-container');
        const option = document.createElement('div');
        option.classList.add('dropdown-option', 'unit-filter-option');
        optionContainer.appendChild(option);
        if (data.color) {
            option.style.backgroundColor = data.color;
        }
        if (data.iconUrl) {
            const icon = document.createElement('img');
            icon.src = data.iconUrl;
            icon.classList.add('unit-filter-icon');
            option.appendChild(icon);
        }
        if (data.text) {
            const label = document.createElement('span');
            label.textContent = data.text;
            label.classList.add('unit-filter-label');
            option.appendChild(label);
        }
        return optionContainer;
    }
    getInputElem() {
        return this.buttonElem;
    }
    getInputValue() {
        return this.getFilterDataValue(this.filterData);
    }
    setInputValue(newValue) {
        this.setFilterDataValue(this.filterData, newValue);
        const optionData = this.currentOptions.find(optionData => optionData.value == newValue);
        if (!optionData) {
            return;
        }
        this.buttonElem.innerHTML = '';
        this.buttonElem.appendChild(this.makeOptionElem(optionData));
    }
}
class PlayerFilter extends UnitGroupFilter {
    constructor(parent, filterData) {
        super(parent, filterData, 'All Players');
        this.rootElem.classList.add('player-filter-root');
    }
    getFilterDataValue(filterData) {
        return filterData.player;
    }
    setFilterDataValue(filterData, newValue) {
        filterData.player = newValue;
    }
    getAllUnits(simResult) {
        return simResult.getPlayers();
    }
}
class TargetFilter extends UnitGroupFilter {
    constructor(parent, filterData) {
        super(parent, filterData, 'All Targets');
        this.rootElem.classList.add('target-filter-root');
    }
    getFilterDataValue(filterData) {
        return filterData.target;
    }
    setFilterDataValue(filterData, newValue) {
        filterData.target = newValue;
    }
    getAllUnits(simResult) {
        return simResult.getTargets();
    }
}
