import { Input, InputConfig } from './input.js';
/**
 * Data for creating a number picker.
 */
export interface NumberPickerConfig<ModObject> extends InputConfig<ModObject, number> {
    float?: boolean;
}
export declare class NumberPicker<ModObject> extends Input<ModObject, number> {
    private readonly inputElem;
    private float;
    constructor(parent: HTMLElement, modObject: ModObject, config: NumberPickerConfig<ModObject>);
    getInputElem(): HTMLElement;
    getInputValue(): number;
    setInputValue(newValue: number): void;
}
