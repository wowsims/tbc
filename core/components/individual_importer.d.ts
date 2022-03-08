import { Spec } from '/tbc/core/proto/common.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Popup } from './popup.js';
export declare class IndividualImporter<SpecType extends Spec> extends Popup {
    private readonly simUI;
    private readonly importButton;
    constructor(parent: HTMLElement, simUI: IndividualSimUI<SpecType>);
    private setup70UpgradesImport;
    private setupAddonImport;
    private finishImport;
}
