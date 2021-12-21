import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { Component } from '/tbc/core/components/component.js';
export declare type SavedDataManagerConfig<ModObject, T> = {
    label: string;
    storageKey: string;
    changeEmitters: Array<TypedEvent<any>>;
    equals: (a: T, b: T) => boolean;
    getData: (modObject: ModObject) => T;
    setData: (eventID: EventID, modObject: ModObject, data: T) => void;
    toJson: (a: T) => any;
    fromJson: (obj: any) => T;
};
export declare type SavedDataConfig<ModObject, T> = {
    name: string;
    data: T;
    tooltip?: string;
    isPreset?: boolean;
    enableWhen?: (obj: ModObject) => boolean;
};
export declare class SavedDataManager<ModObject, T> extends Component {
    private readonly modObject;
    private readonly config;
    private readonly userData;
    private readonly presets;
    private readonly savedDataDiv;
    private readonly saveInput;
    private frozen;
    constructor(parent: HTMLElement, modObject: ModObject, config: SavedDataManagerConfig<ModObject, T>);
    addSavedData(config: SavedDataConfig<ModObject, T>): void;
    private makeSavedData;
    private saveUserData;
    loadUserData(): void;
    freeze(): void;
}
