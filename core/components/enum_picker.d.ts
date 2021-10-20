import { Sim } from '/tbc/core/sim.js';
import { Input, InputConfig } from './input.js';
export interface EnumValueConfig {
    name: string;
    value: number;
    tooltip?: string;
}
export interface EnumPickerConfig extends InputConfig<number> {
    values: Array<EnumValueConfig>;
}
export declare class EnumPicker extends Input<number> {
    private readonly selectElem;
    constructor(parent: HTMLElement, sim: Sim<any>, config: EnumPickerConfig);
    getInputElem(): HTMLElement;
    getInputValue(): number;
    setInputValue(newValue: number): void;
}
