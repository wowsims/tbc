import { Stat } from '../api/common.js';
import { Sim } from '../sim.js';
import { Component } from './component.js';
import { NumberPicker } from './number_picker.js';
export declare class CustomStatsPicker extends Component {
    readonly stats: Array<Stat>;
    readonly statPickers: Array<NumberPicker>;
    constructor(parent: HTMLElement, sim: Sim<any>, stats: Array<Stat>);
}
