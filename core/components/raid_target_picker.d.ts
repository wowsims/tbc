import { Input, InputConfig } from '/tbc/core/components/input.js';
import { Player } from '/tbc/core/player.js';
import { Raid } from '/tbc/core/raid.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
export interface RaidTargetPickerConfig<ModObject> extends InputConfig<ModObject, RaidTarget> {
    noTargetLabel: string;
    compChangeEmitter: TypedEvent<void>;
}
export interface RaidTargetElemOption {
    iconUrl: string;
    text: string;
    color: string;
    isDropdown: boolean;
}
export interface RaidTargetOption extends RaidTargetElemOption {
    value: Player<any> | null;
}
export declare class RaidTargetPicker<ModObject> extends Input<ModObject, RaidTarget> {
    private readonly config;
    private readonly raid;
    private readonly noTargetOption;
    private curPlayer;
    private curRaidTarget;
    private currentOptions;
    private readonly buttonElem;
    private readonly dropdownElem;
    constructor(parent: HTMLElement, raid: Raid, modObj: ModObject, config: RaidTargetPickerConfig<ModObject>);
    private makeTargetOptions;
    private updateOptions;
    private makeOption;
    getInputElem(): HTMLElement;
    getInputValue(): RaidTarget;
    setInputValue(newValue: RaidTarget): void;
    static makeOptionElem(data: RaidTargetElemOption): HTMLElement;
}
