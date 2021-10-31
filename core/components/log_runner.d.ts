import { SimUI } from '/tbc/core/sim_ui.js';
import { Component } from './component.js';
import { DetailedResults } from './detailed_results.js';
import { Results } from './results.js';
export declare class LogRunner extends Component {
    constructor(parent: HTMLElement, simUI: SimUI<any>, results: Results, detailedResults: DetailedResults);
}
