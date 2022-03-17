import { Spec } from '/tbc/core/proto/common.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';
import { Popup } from './popup.js';
export declare function newIndividualExporters<SpecType extends Spec>(simUI: IndividualSimUI<SpecType>): HTMLElement;
export declare abstract class Exporter extends Popup {
    private readonly textElem;
    constructor(parent: HTMLElement, title: string, allowDownload: boolean);
    protected init(): void;
    abstract getData(): string;
}
