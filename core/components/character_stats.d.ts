import { Stat } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { Player } from '/tbc/core/player.js';
import { Component } from './component.js';
export declare class CharacterStats extends Component {
    readonly stats: Array<Stat>;
    readonly valueElems: Array<HTMLTableCellElement>;
    private readonly player;
    private readonly modifyDisplayStats?;
    constructor(parent: HTMLElement, player: Player<any>, stats: Array<Stat>, modifyDisplayStats?: (player: Player<any>, stats: Stats) => Stats);
    private updateStats;
}
