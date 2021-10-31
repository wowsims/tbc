import { Input, InputConfig } from './input.js';
export interface EnumValueConfig {
    name: string;
    value: number;
    tooltip?: string;
}
export interface EnumPickerConfig<ModObject> extends InputConfig<ModObject, number> {
    values: Array<EnumValueConfig>;
}
export declare class EnumPicker<ModObject> extends Input<ModObject, number> {
    private readonly selectElem;
    constructor(parent: HTMLElement, modObject: ModObject, config: EnumPickerConfig<ModObject>);
    getInputElem(): HTMLElement;
    getInputValue(): number;
    setInputValue(newValue: number): void;
}
