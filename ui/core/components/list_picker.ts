
import { Component } from './component.js';

export interface ListPickerConfig {
	simpleTargetStats?: Array<Stat>;
	showNumTargets: boolean;
	showExecuteProportion: boolean;
}

export class ListPicker extends Component {
	constructor(parent: HTMLElement, modEncounter: Encounter, config: ListPickerConfig) {
		super(parent, 'list-picker-root');
	}
}
