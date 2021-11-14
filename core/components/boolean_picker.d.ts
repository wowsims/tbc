import { Input, InputConfig } from './input.js';
/**
 * Data for creating a boolean picker (checkbox).
 */
export interface BooleanPickerConfig<ModObject> extends InputConfig<ModObject, boolean> {
}
export declare class BooleanPicker<ModObject> extends Input<ModObject, boolean> {
    private readonly inputElem;
    constructor(parent: HTMLElement, modObject: ModObject, config: BooleanPickerConfig<ModObject>);
    getInputElem(): HTMLElement;
    getInputValue(): boolean;
    setInputValue(newValue: boolean): void;
}
