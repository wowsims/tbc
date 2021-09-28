import { Sim } from '../sim.js';
import { TypedEvent } from '../typed_event.js';
import { Component } from './component.js';
export interface EnumPickerConfig {
    label?: string;
    defaultValue?: number;
    names: Array<string>;
    values: Array<number>;
    changedEvent: (sim: Sim<any>) => TypedEvent<any>;
    getValue: (sim: Sim<any>) => number;
    setValue: (sim: Sim<any>, newValue: number) => void;
}
export declare class EnumPicker extends Component {
    constructor(parent: HTMLElement, sim: Sim<any>, config: EnumPickerConfig);
}
