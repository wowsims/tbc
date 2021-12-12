import { Component } from '/tbc/core/components/component.js';
import { Sim } from './sim.js';
import { TypedEvent } from './typed_event.js';
export interface SimUIConfig {
    title: string;
    knownIssues?: Array<string>;
}
export declare abstract class SimUI extends Component {
    readonly sim: Sim;
    readonly changeEmitter: TypedEvent<void>;
    readonly resultsPendingElem: HTMLElement;
    readonly resultsContentElem: HTMLElement;
    constructor(parentElem: HTMLElement, sim: Sim, config: SimUIConfig);
    addAction(name: string, cssClass: string, actFn: () => void): void;
    addTab(title: string, cssClass: string, innerHTML: string): void;
    setResultsPending(): void;
    setResultsContent(innerHTML: string): void;
    abstract getStorageKey(postfix: string): string;
    getSettingsStorageKey(): string;
}
