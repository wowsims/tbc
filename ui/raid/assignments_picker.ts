import { Component } from '/tbc/core/components/component.js';
import { Player } from '/tbc/core/player.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { Class } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { BuffBot } from '/tbc/core/proto/ui.js';
import { getEnumValues } from '/tbc/core/utils.js';

import { RaidSimUI } from './raid_sim_ui.js';

declare var tippy: any;

export class AssignmentsPicker extends Component {
	readonly raidSimUI: RaidSimUI;
	readonly changeEmitter: TypedEvent<void> = new TypedEvent<void>();

  constructor(parentElem: HTMLElement, raidSimUI: RaidSimUI) {
    super(parentElem, 'assignments-picker-root');
		this.raidSimUI = raidSimUI;
	}
}

// Dropdown menu for selecting a player.
class PlayerPicker extends Input<FilterData, number> {
	private readonly filterData: FilterData;
	readonly changeEmitter: TypedEvent<void>;

	private currentOptions: Array<PlayerFilterOption>;

	private readonly buttonElem: HTMLElement;
	private readonly dropdownElem: HTMLElement;

  constructor(parent: HTMLElement, filterData: FilterData) {
		const changeEmitter = new TypedEvent<void>();
    super(parent, 'player-filter-root', filterData, {
			extraCssClasses: [
				'dropdown-root',
			],
			changedEvent: (filterData: FilterData) => changeEmitter,
			getValue: (filterData: FilterData) => filterData.player,
			setValue: (filterData: FilterData, newValue: number) => filterData.player = newValue,
		});
		this.filterData = filterData;
		this.currentOptions = [allPlayersOption];
		this.changeEmitter = changeEmitter;

    this.rootElem.innerHTML = `
			<div class="dropdown-button player-filter-button"></div>
			<div class="dropdown-panel player-filter-dropdown"></div>
    `;

		this.buttonElem = this.rootElem.getElementsByClassName('player-filter-button')[0] as HTMLElement;
		this.dropdownElem = this.rootElem.getElementsByClassName('player-filter-dropdown')[0] as HTMLElement;

		this.buttonElem.addEventListener('click', event => {
			event.preventDefault();
		});

		this.init();
  }

	setOptions(simResult: SimResult) {
		this.currentOptions = [allPlayersOption].concat(simResult.getPlayers().map(player => {
			return {
				iconUrl: player.iconUrl,
				text: player.label,
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

	private makeOption(data: PlayerFilterOption): HTMLElement {
		const option = this.makeOptionElem(data);

		option.addEventListener('click', event => {
			event.preventDefault();
			this.filterData.player = data.value;
			this.changeEmitter.emit();
		});

		return option;
	}

	private makeOptionElem(data: PlayerFilterOption): HTMLElement {
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

	getInputElem(): HTMLElement {
		return this.buttonElem;
	}

	getInputValue(): number {
		return this.filterData.player;
	}

  setInputValue(newValue: number) {
    this.filterData.player = newValue;

		const optionData = this.currentOptions.find(optionData => optionData.value == newValue);
		if (!optionData) {
			return;
		}

		this.buttonElem.innerHTML = '';
		this.buttonElem.appendChild(this.makeOptionElem(optionData));
  }
}
