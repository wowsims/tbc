import { Component } from '/tbc/core/components/component.js';
export declare class ResultsViewer extends Component {
    readonly pendingElem: HTMLElement;
    readonly contentElem: HTMLElement;
    constructor(parentElem: HTMLElement);
    hideAll(): void;
    setPending(): void;
    setContent(innerHTML: string): void;
}
