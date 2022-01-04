import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { Input, InputConfig } from './input.js';
export interface IconEnumValueConfig<T> {
    actionId?: ActionId;
    color?: string;
    value: T;
}
export interface IconEnumPickerConfig<ModObject, T> extends InputConfig<ModObject, T> {
    numColumns: number;
    values: Array<IconEnumValueConfig<T>>;
    equals: (a: T, b: T) => boolean;
    zeroValue: T;
    backupIconUrl?: (value: T) => ActionId;
}
export declare class IconEnumPicker<ModObject, T> extends Input<ModObject, T> {
    private readonly config;
    private currentValue;
    private readonly buttonElem;
    constructor(parent: HTMLElement, modObj: ModObject, config: IconEnumPickerConfig<ModObject, T>);
    private setActionImage;
    private setImage;
    getInputElem(): HTMLElement;
    getInputValue(): T;
    setInputValue(newValue: T): void;
}
