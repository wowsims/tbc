import { Stat } from '/tbc/core/proto/common.js';
import { Sim } from '/tbc/core/sim.js';
import { Component } from './component.js';
export declare class CharacterStats extends Component {
    readonly stats: Array<Stat>;
    readonly valueElems: Array<HTMLTableCellElement>;
    constructor(parent: HTMLElement, stats: Array<Stat>, sim: Sim<any>);
    private updateStats;
}
