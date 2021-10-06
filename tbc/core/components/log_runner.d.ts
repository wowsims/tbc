import { Sim } from '/tbc/core/sim.js';
import { Component } from './component.js';
import { DetailedResults } from './detailed_results.js';
import { Results } from './results.js';
export declare class LogRunner extends Component {
    constructor(parent: HTMLElement, sim: Sim<any>, results: Results, detailedResults: DetailedResults);
}
