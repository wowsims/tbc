import { Sim } from '/tbc/core/sim.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { Component } from './component.js';
/**
 * Data for creating a number picker.
 */
export declare type NumberPickerConfig = {
    label?: string;
    defaultValue?: number;
    changedEvent: (sim: Sim<any>) => TypedEvent<any>;
    getValue: (sim: Sim<any>) => number;
    setValue: (sim: Sim<any>, newValue: number) => void;
};
export declare class NumberPicker extends Component {
    constructor(parent: HTMLElement, sim: Sim<any>, config: NumberPickerConfig);
}
