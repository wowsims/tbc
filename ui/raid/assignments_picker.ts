import { Component } from '/tbc/core/components/component.js';
import { Player } from '/tbc/core/player.js';
import { Raid } from '/tbc/core/raid.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { Class } from '/tbc/core/proto/common.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { getEnumValues } from '/tbc/core/utils.js';
import { newRaidTarget, emptyRaidTarget } from '/tbc/core/proto_utils/utils.js';

import { BuffBot } from './buff_bot.js';
import { RaidSimUI } from './raid_sim_ui.js';

declare var tippy: any;

export class AssignmentsPicker extends Component {
	readonly raidSimUI: RaidSimUI;
	readonly changeEmitter: TypedEvent<void> = new TypedEvent<void>();

	private readonly innervatesPicker: InnervatesPicker;

  constructor(parentElem: HTMLElement, raidSimUI: RaidSimUI) {
    super(parentElem, 'assignments-picker-root');
		this.raidSimUI = raidSimUI;
		this.innervatesPicker = new InnervatesPicker(this.rootElem, raidSimUI);
	}
}

export class InnervatesPicker extends Component {
	readonly raidSimUI: RaidSimUI;
	readonly changeEmitter: TypedEvent<void> = new TypedEvent<void>();

	private readonly playersContainer: HTMLElement;

  constructor(parentElem: HTMLElement, raidSimUI: RaidSimUI) {
    super(parentElem, 'innervates-picker-root');
		this.raidSimUI = raidSimUI;

		this.playersContainer = document.createElement('div');
		this.playersContainer.classList.add('innervate-players-container');
		this.rootElem.appendChild(this.playersContainer);
	}

	private update(playersAndBots: Array<Player<any> | BuffBot | null>) {
		this.playersContainer.innerHTML = '';

		const druids = playersAndBots.filter(playerOrBot => playerOrBot?.getClass() == Class.ClassDruid);
	}
}


export interface RaidTargetPickerConfig<ModObject> extends InputConfig<ModObject, RaidTarget> {
	raid: Raid,
}

interface RaidTargetOption {
	iconUrl: string,
	text: string,
	color: string,
	value: RaidTarget,
};

const unassignedOption: RaidTargetOption = {
	iconUrl: '',
	text: 'Unassigned',
	color: 'black',
	value: emptyRaidTarget(),
};

// Dropdown menu for selecting a player.
class RaidTargetPicker<ModObject> extends Input<ModObject, RaidTarget> {
	private readonly config: RaidTargetPickerConfig<ModObject>;
	private readonly raidTarget: RaidTarget;

	private currentOptions: Array<RaidTargetOption>;

	private readonly buttonElem: HTMLElement;
	private readonly dropdownElem: HTMLElement;

  constructor(parent: HTMLElement, modObj: ModObject, config: RaidTargetPickerConfig<ModObject>) {
    super(parent, 'raid-target-picker-root', modObj, config);
    this.config = config;
		this.raidTarget = emptyRaidTarget();

		this.currentOptions = [unassignedOption];
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

		config.raidSimUI.compChangeEmitter.on(() => {
			this.setOptions(config.raid.getPlayers().filter(p => p != null));
		});

		this.init();
  }

	private setOptions(players: Array<Player<any>>) {
		this.currentOptions = [unassignedOption].concat(players.map(player => {
			return {
				iconUrl: player.getTalentTreeIcon(),
				text: player.getLabel(),
				color: player.getClassColor(),
				value: newRaidTarget(player.getRaidIndex()),
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

	getInputValue(): RaidTarget {
		return RaidTarget.clone(this.raidTarget);
	}

  setInputValue(newValue: RaidTarget) {
		this.raidTarget = RaidTarget.clone(newValue);

		const optionData = this.currentOptions.find(optionData => optionData.value == newValue);
		if (!optionData) {
			return;
		}

		this.buttonElem.innerHTML = '';
		this.buttonElem.appendChild(this.makeOptionElem(optionData));
  }
}
