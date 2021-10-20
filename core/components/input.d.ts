import { Sim } from '/tbc/core/sim.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { Component } from './component.js';
/**
 * Data for creating a new input UI element.
 */
export interface InputConfig<T> {
    label?: string;
    labelTooltip?: string;
    defaultValue?: T;
    changedEvent: (sim: Sim<any>) => TypedEvent<any>;
    getValue: (sim: Sim<any>) => T;
    setValue: (sim: Sim<any>, newValue: T) => void;
    enableWhen?: (sim: Sim<any>) => boolean;
}
export declare abstract class Input<T> extends Component {
    private readonly inputConfig;
    readonly sim: Sim<any>;
    constructor(parent: HTMLElement, cssClass: string, sim: Sim<any>, config: InputConfig<T>);
    private update;
    init(): void;
    abstract getInputElem(): HTMLElement;
    abstract getInputValue(): T;
    abstract setInputValue(newValue: T): void;
    inputChanged(): void;
}
