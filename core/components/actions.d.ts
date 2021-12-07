import { Stat } from '/tbc/core/proto/common.js';
import { SimUI } from '/tbc/core/sim_ui.js';
import { Component } from './component.js';
import { DetailedResults } from './detailed_results.js';
import { Results } from './results.js';
export declare class Actions extends Component {
    constructor(parent: HTMLElement, simUI: SimUI<any>, epStats: Array<Stat>, epReferenceStat: Stat, results: Results, detailedResults: DetailedResults);
    private makeStatWeightsRequest;
}
