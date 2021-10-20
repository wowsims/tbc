import { Sim } from '/tbc/core/sim.js';
import { Input, InputConfig } from './input.js';
/**
 * Data for creating a number picker.
 */
export interface NumberPickerConfig extends InputConfig<number> {
}
export declare class NumberPicker extends Input<number> {
    private readonly inputElem;
    constructor(parent: HTMLElement, sim: Sim<any>, config: NumberPickerConfig);
    getInputElem(): HTMLElement;
    getInputValue(): number;
    setInputValue(newValue: number): void;
}
