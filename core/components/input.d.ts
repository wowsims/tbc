import { TypedEvent } from '/tbc/core/typed_event.js';
import { Component } from './component.js';
/**
 * Data for creating a new input UI element.
 */
export interface InputConfig<ModObject, T> {
    label?: string;
    labelTooltip?: string;
    defaultValue?: T;
    changedEvent: (obj: ModObject) => TypedEvent<any>;
    getValue: (obj: ModObject) => T;
    setValue: (obj: ModObject, newValue: T) => void;
    enableWhen?: (obj: ModObject) => boolean;
}
export declare abstract class Input<ModObject, T> extends Component {
    private readonly inputConfig;
    readonly modObject: ModObject;
    constructor(parent: HTMLElement, cssClass: string, modObject: ModObject, config: InputConfig<ModObject, T>);
    private update;
    init(): void;
    abstract getInputElem(): HTMLElement;
    abstract getInputValue(): T;
    abstract setInputValue(newValue: T): void;
    inputChanged(): void;
}
