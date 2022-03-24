import { Stat } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { Player } from '/tbc/core/player.js';
import { Component } from './component.js';
export declare type StatBreakdown = Array<{
    label: string;
    value: number;
}>;
export declare class CharacterStats extends Component {
    readonly stats: Array<Stat>;
    readonly valueElems: Array<HTMLTableCellElement>;
    readonly tooltipElems: Array<HTMLElement>;
    private readonly player;
    private readonly modifyDisplayStats?;
    private readonly statBreakdowns?;
    constructor(parent: HTMLElement, player: Player<any>, stats: Array<Stat>, modifyDisplayStats?: (player: Player<any>, stats: Stats) => Stats, statBreakdowns?: (player: Player<any>, stats: Stats) => Partial<Record<Stat, StatBreakdown>>);
    private updateStats;
    static statDisplayString(stat: Stat, rawValue: number): string;
}
