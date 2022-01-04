import { Input, InputConfig } from './input.js';
/**
 * Data for creating a number list picker.
 */
export interface NumberListPickerConfig<ModObject> extends InputConfig<ModObject, Array<number>> {
    placeholder?: string;
}
export declare class NumberListPicker<ModObject> extends Input<ModObject, Array<number>> {
    private readonly inputElem;
    constructor(parent: HTMLElement, modObject: ModObject, config: NumberListPickerConfig<ModObject>);
    getInputElem(): HTMLElement;
    getInputValue(): Array<number>;
    setInputValue(newValue: Array<number>): void;
}
