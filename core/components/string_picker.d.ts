import { Input, InputConfig } from './input.js';
/**
 * Data for creating a string picker.
 */
export interface StringPickerConfig<ModObject> extends InputConfig<ModObject, string> {
}
export declare class StringPicker<ModObject> extends Input<ModObject, string> {
    private readonly inputElem;
    constructor(parent: HTMLElement, modObject: ModObject, config: StringPickerConfig<ModObject>);
    getInputElem(): HTMLElement;
    getInputValue(): string;
    setInputValue(newValue: string): void;
}
