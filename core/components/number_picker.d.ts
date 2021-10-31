import { Input, InputConfig } from './input.js';
/**
 * Data for creating a number picker.
 */
export interface NumberPickerConfig<ModObject> extends InputConfig<ModObject, number> {
}
export declare class NumberPicker<ModObject> extends Input<ModObject, number> {
    private readonly inputElem;
    constructor(parent: HTMLElement, modObject: ModObject, config: NumberPickerConfig<ModObject>);
    getInputElem(): HTMLElement;
    getInputValue(): number;
    setInputValue(newValue: number): void;
}
