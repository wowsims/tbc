import { ItemOrSpellId } from '/tbc/core/resources.js';
import { Input, InputConfig } from './input.js';
export interface IconEnumValueConfig {
    id?: ItemOrSpellId;
    color?: string;
    value: number;
}
export interface IconEnumPickerConfig<ModObject> extends InputConfig<ModObject, number> {
    values: Array<IconEnumValueConfig>;
}
export declare class IconEnumPicker<ModObject> extends Input<ModObject, number> {
    private readonly config;
    private currentValue;
    private readonly buttonElem;
    constructor(parent: HTMLElement, modObj: ModObject, config: IconEnumPickerConfig<ModObject>);
    private setImage;
    getInputElem(): HTMLElement;
    getInputValue(): number;
    setInputValue(newValue: number): void;
}
