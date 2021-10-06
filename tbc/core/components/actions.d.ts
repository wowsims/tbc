import { Stat } from '/tbc/core/proto/common.js';
import { Sim } from '/tbc/core/sim.js';
import { Component } from './component.js';
import { DetailedResults } from './detailed_results.js';
import { Results } from './results.js';
export declare class Actions extends Component {
    constructor(parent: HTMLElement, sim: Sim<any>, epStats: Array<Stat>, epReferenceStat: Stat, results: Results, detailedResults: DetailedResults);
}
