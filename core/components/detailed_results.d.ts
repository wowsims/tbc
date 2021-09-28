import { IndividualSimResult } from '../api/api.js';
import { Component } from './component.js';
export declare class DetailedResults extends Component {
    private readonly iframeElem;
    private tabWindow;
    private latestResult;
    constructor(parent: HTMLElement);
    setPending(): void;
    setSimResult(result: IndividualSimResult): void;
}
