import { Component } from '/tbc/core/components/component.js';
import { Input, InputConfig } from '/tbc/core/components/input.js';
import { Player } from '/tbc/core/player.js';
import { Raid } from '/tbc/core/raid.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { newRaidTarget, emptyRaidTarget } from '/tbc/core/proto_utils/utils.js';

declare var tippy: any;

export interface RaidTargetPickerConfig<ModObject> extends InputConfig<ModObject, RaidTarget> {
	noTargetLabel: string,
	compChangeEmitter: TypedEvent<void>,
	getOptions: () => Array<RaidTargetOption>,
}

export interface RaidTargetElemOption {
	iconUrl: string,
	text: string,
	color: string,
	isDropdown: boolean,
};

export interface RaidTargetOption extends RaidTargetElemOption {
	value: RaidTarget,
};

// Dropdown menu for selecting a player.
export class RaidTargetPicker<ModObject> extends Input<ModObject, RaidTarget> {
	private readonly config: RaidTargetPickerConfig<ModObject>;
	private readonly noTargetOption: RaidTargetOption;

	private raidTarget: RaidTarget;

	private currentOptions: Array<RaidTargetOption>;

	private readonly buttonElem: HTMLElement;
	private readonly dropdownElem: HTMLElement;

	constructor(parent: HTMLElement, modObj: ModObject, config: RaidTargetPickerConfig<ModObject>) {
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

		this.buttonElem = this.rootElem.getElementsByClassName('raid-target-picker-button')[0] as HTMLElement;
		this.dropdownElem = this.rootElem.getElementsByClassName('raid-target-picker-dropdown')[0] as HTMLElement;

		this.buttonElem.addEventListener('click', event => {
			event.preventDefault();
		});

		this.setOptions(TypedEvent.nextEventID(), config.getOptions());
		config.compChangeEmitter.on(eventID => {
			this.setOptions(eventID, config.getOptions());
		});

		this.init();
	}

	private setOptions(eventID: EventID, options: Array<RaidTargetOption>) {
		this.currentOptions = [this.noTargetOption].concat(options);

		const hasSameOption = this.currentOptions.find(option => RaidTarget.equals(option.value, this.getInputValue())) != null;
		if (!hasSameOption) {
			this.raidTarget = this.noTargetOption.value;
			this.inputChanged(eventID);
		}

		this.dropdownElem.innerHTML = '';
		this.currentOptions.forEach(option => this.dropdownElem.appendChild(this.makeOption(option)));
	}

	private makeOption(data: RaidTargetOption): HTMLElement {
		const option = RaidTargetPicker.makeOptionElem(data);

		option.addEventListener('click', event => {
			event.preventDefault();
			this.raidTarget = data.value;
			this.inputChanged(TypedEvent.nextEventID());
		});

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

		const optionData = this.currentOptions.find(optionData => RaidTarget.equals(optionData.value, newValue));
		if (!optionData) {
			return;
		}

		this.buttonElem.innerHTML = '';
		this.buttonElem.appendChild(RaidTargetPicker.makeOptionElem(optionData));
	}

	static makeOptionElem(data: RaidTargetElemOption): HTMLElement {
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
