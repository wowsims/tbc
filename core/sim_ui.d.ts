import { Component } from '/tbc/core/components/component.js';
import { ResultsViewer } from '/tbc/core/components/results_viewer.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Sim } from './sim.js';
import { EventID, TypedEvent } from './typed_event.js';
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
    readonly resultsViewer: ResultsViewer;
    private warnings;
    private warningsTippy;
    constructor(parentElem: HTMLElement, sim: Sim, config: SimUIConfig);
    addAction(name: string, cssClass: string, actFn: () => void): void;
    addTab(title: string, cssClass: string, innerHTML: string): void;
    addToolbarItem(elem: HTMLElement): void;
    private updateWarnings;
    addWarning(warning: SimWarning): void;
    abstract getStorageKey(postfix: string): string;
    getSettingsStorageKey(): string;
    getSavedEncounterStorageKey(): string;
    isIndividualSim(): boolean;
    runSim(onProgress: Function): Promise<void>;
    runSimOnce(): Promise<void>;
    abstract applyDefaults(eventID: EventID): void;
}
