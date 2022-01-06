import { Component } from '/tbc/core/components/component.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Sim } from './sim.js';
import { TypedEvent } from './typed_event.js';
export interface SimWarning {
    updateOn: TypedEvent<any>;
    shouldDisplay: () => boolean;
    getContent: () => string;
}
export interface SimUIConfig {
    spec: Spec | null;
    knownIssues?: Array<string>;
}
export declare abstract class SimUI extends Component {
    readonly sim: Sim;
    readonly isWithinRaidSim: boolean;
    readonly changeEmitter: TypedEvent<void>;
    readonly resultsPendingElem: HTMLElement;
    readonly resultsContentElem: HTMLElement;
    private warnings;
    constructor(parentElem: HTMLElement, sim: Sim, config: SimUIConfig);
    addAction(name: string, cssClass: string, actFn: () => void): void;
    addTab(title: string, cssClass: string, innerHTML: string): void;
    addToolbarItem(elem: HTMLElement): void;
    hideAllResults(): void;
    setResultsPending(): void;
    setResultsContent(innerHTML: string): void;
    private updateWarnings;
    addWarning(warning: SimWarning): void;
    abstract getStorageKey(postfix: string): string;
    getSettingsStorageKey(): string;
    getSavedEncounterStorageKey(): string;
    isIndividualSim(): boolean;
    runSim(): Promise<void>;
    runSimOnce(): Promise<void>;
}
