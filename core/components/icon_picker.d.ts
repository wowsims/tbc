import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { Input, InputConfig } from './input.js';
export interface IconPickerConfig<ModObject, ValueType> extends InputConfig<ModObject, ValueType> {
    id: ActionId;
    states: number;
    improvedId?: ActionId;
    improvedId2?: ActionId;
}
export declare class IconPicker<ModObject, ValueType> extends Input<ModObject, ValueType> {
    private readonly config;
    private readonly rootAnchor;
    private readonly improvedAnchor;
    private readonly improvedAnchor2;
    private readonly counterElem;
    private currentValue;
    constructor(parent: HTMLElement, modObj: ModObject, config: IconPickerConfig<ModObject, ValueType>);
    getInputElem(): HTMLElement;
    getInputValue(): ValueType;
    setInputValue(newValue: ValueType): void;
}
