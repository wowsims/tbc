import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';
import { Component } from './component.js';
export declare type RaidSimData = {
    request: RaidSimRequest;
    result: RaidSimResult;
};
export declare class DetailedResults extends Component {
    private readonly iframeElem;
    private tabWindow;
    private latestResult;
    constructor(parent: HTMLElement);
    setPending(): void;
    setSimResult(request: RaidSimRequest, result: RaidSimResult): void;
}
