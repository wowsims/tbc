import { IndividualSimRequest, IndividualSimResult } from '/tbc/core/proto/api.js';
import { Component } from './component.js';
export declare type IndividualSimData = {
    request: IndividualSimRequest;
    result: IndividualSimResult;
};
export declare class DetailedResults extends Component {
    private readonly iframeElem;
    private tabWindow;
    private latestResult;
    constructor(parent: HTMLElement);
    setPending(): void;
    setSimResult(request: IndividualSimRequest, result: IndividualSimResult): void;
}
