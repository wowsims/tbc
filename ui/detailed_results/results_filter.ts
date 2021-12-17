import { SimResult, SimResultFilter } from '/tbc/core/proto_utils/sim_result.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { Input } from '/tbc/core/components/input.js';

import { ResultComponent, ResultComponentConfig, SimResultData } from './result_component.js';

const ALL_PLAYERS = -1;
const ALL_TARGETS = -1;

interface FilterData {
	player: number,
	target: number,
};

export class ResultsFilter extends ResultComponent {
	private readonly currentFilter: FilterData;

	readonly changeEmitter: TypedEvent<void>;

	private readonly playerFilter: PlayerFilter;

  constructor(config: ResultComponentConfig) {
		config.rootCssClass = 'results-filter-root';
    super(config);
		this.currentFilter = {
			player: ALL_PLAYERS,
			target: ALL_TARGETS,
		};
		this.changeEmitter = new TypedEvent<void>();

		this.playerFilter = new PlayerFilter(this.rootElem, this.currentFilter);
		this.playerFilter.changeEmitter.on(() => this.changeEmitter.emit());
  }

	getFilter(): SimResultFilter {
		return {
			player: this.currentFilter.player == ALL_PLAYERS ? null : this.currentFilter.player,
			target: this.currentFilter.target == ALL_TARGETS ? null : this.currentFilter.target,
		};
	}

	onSimResult(resultData: SimResultData) {
		this.playerFilter.setOptions(resultData.result);
	}

	setPlayer(newPlayer: number | null) {
		this.currentFilter.player = (newPlayer === null) ? ALL_PLAYERS : newPlayer;
		this.playerFilter.changeEmitter.emit();
	}
}

interface PlayerFilterOption {
	iconUrl: string,
	text: string,
	color: string,
	value: number,
};

const allPlayersOption: PlayerFilterOption = {
	iconUrl: '',
	text: 'All Players',
	color: 'black',
	value: ALL_PLAYERS,
};

// Dropdown menu for filtering by player.
class PlayerFilter extends Input<FilterData, number> {
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
