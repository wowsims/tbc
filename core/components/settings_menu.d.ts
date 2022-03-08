import { Spec } from '/tbc/core/proto/common.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Popup } from './popup.js';
export declare class SettingsMenu<SpecType extends Spec> extends Popup {
    private readonly simUI;
    constructor(parent: HTMLElement, simUI: IndividualSimUI<SpecType>);
    private setupEpWeightsSettings;
}
