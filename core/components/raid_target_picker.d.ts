import { Input, InputConfig } from '/tbc/core/components/input.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
export interface RaidTargetPickerConfig<ModObject> extends InputConfig<ModObject, RaidTarget> {
    noTargetLabel: string;
    compChangeEmitter: TypedEvent<void>;
    getOptions: () => Array<RaidTargetOption>;
}
export interface RaidTargetElemOption {
    iconUrl: string;
    text: string;
    color: string;
    isDropdown: boolean;
}
export interface RaidTargetOption extends RaidTargetElemOption {
    value: RaidTarget;
}
export declare class RaidTargetPicker<ModObject> extends Input<ModObject, RaidTarget> {
    private readonly config;
    private readonly noTargetOption;
    private raidTarget;
    private currentOptions;
    private readonly buttonElem;
    private readonly dropdownElem;
    constructor(parent: HTMLElement, modObj: ModObject, config: RaidTargetPickerConfig<ModObject>);
    private setOptions;
    private makeOption;
    getInputElem(): HTMLElement;
    getInputValue(): RaidTarget;
    setInputValue(newValue: RaidTarget): void;
    static makeOptionElem(data: RaidTargetElemOption): HTMLElement;
}
