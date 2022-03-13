import { SimUI } from '/tbc/core/sim_ui.js';
import { Component } from './component.js';
import { RaidSimResultsManager } from './raid_sim_action.js';
export declare class DetailedResults extends Component {
    private readonly simUI;
    private readonly iframeElem;
    private tabWindow;
    private latestRun;
    constructor(parent: HTMLElement, simUI: SimUI, simResultsManager: RaidSimResultsManager);
    private setSimRunData;
    private updateSettings;
    private postMessage;
}
