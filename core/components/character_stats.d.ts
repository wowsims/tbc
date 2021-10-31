import { Stat } from '/tbc/core/proto/common.js';
import { Player } from '/tbc/core/player.js';
import { Component } from './component.js';
export declare class CharacterStats extends Component {
    readonly stats: Array<Stat>;
    readonly valueElems: Array<HTMLTableCellElement>;
    constructor(parent: HTMLElement, stats: Array<Stat>, player: Player<any>);
    private updateStats;
}
