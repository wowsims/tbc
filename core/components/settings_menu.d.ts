import { Spec } from '/tbc/core/proto/common.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Component } from './component.js';
export declare class SettingsMenu<SpecType extends Spec> extends Component {
    private readonly simUI;
    constructor(parent: HTMLElement, simUI: IndividualSimUI<SpecType>);
    private setupEpWeightsSettings;
}
